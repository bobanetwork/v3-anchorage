package ether

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum-optimism/optimism/boba-chain-ops/utils"
)

type EthCrawler struct {
	RpcClient          *rpc.Client
	EthClient          *ethclient.Client
	EndBlock           int64
	RpcTimeout         time.Duration
	RpcPollingInterval time.Duration
	Output             string
	ctx                context.Context
	stop               chan struct{}
}

type EthAddresses struct {
	BlockNumber int64             `json:"blockNumber"`
	Addresses   []*common.Address `json:"addresses"`
}

type TraceTransaction struct {
	From    common.Address      `json:"from"`
	Value   hexutil.Big         `json:"value"`
	Gas     hexutil.Uint64      `json:"gas"`
	GasUsed hexutil.Uint64      `json:"gasUsed"`
	Input   hexutil.Bytes       `json:"input"`
	Output  hexutil.Bytes       `json:"output"`
	To      common.Address      `json:"to,omitempty"`
	Calls   []*TraceTransaction `json:"calls,omitempty"`
}

func NewEthCrawler(rpcClient *rpc.Client, ethClient *ethclient.Client, endBlock int64, rpcTimeout, rpcPollingInterval time.Duration, output string) *EthCrawler {
	return &EthCrawler{
		RpcClient:          rpcClient,
		EthClient:          ethClient,
		EndBlock:           endBlock,
		RpcTimeout:         rpcTimeout,
		RpcPollingInterval: rpcPollingInterval,
		Output:             output,
		ctx:                context.Background(),
		stop:               make(chan struct{}),
	}
}

func (e *EthCrawler) Start() error {
	currentBlock, addresses, err := e.LoadAddresses()
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	// insert the constant addresses to the addresses list
	addresses = append(addresses, &utils.SEQUENCER_FEE_WALLET_ADDRESS)

	go e.Loop(currentBlock, addresses)
	return nil
}

func (e *EthCrawler) Stop() {
	close(e.stop)
}

func (e *EthCrawler) Wait() {
	<-e.stop
}

func (e *EthCrawler) Loop(currentBlock int64, addresses []*common.Address) error {
	var err error
	timer := time.NewTicker(e.RpcPollingInterval)
	defer timer.Stop()
	mapAddresses := MapAddresses(addresses)
	rpcMethods := &RPC{}
	for {
		select {
		case <-timer.C:
			currentBlock, mapAddresses, err = e.StartCrawler(rpcMethods, currentBlock, mapAddresses)
			if err != nil {
				log.Error("error in crawler", "error", err)
			}
		case <-e.ctx.Done():
			e.Stop()
		}
	}
}

func (e *EthCrawler) StartCrawler(rpc RPCMethods, currentBlock int64, mapAddresses map[common.Address]bool) (int64, map[common.Address]bool, error) {
	var (
		err      error
		endBlock = big.NewInt(e.EndBlock)
	)
	if endBlock.Int64() == currentBlock {
		e.ctx.Done()
	}
	if endBlock.Cmp(common.Big0) == 0 {
		endBlock, err = rpc.GetBlockNumber(e.EthClient, e.RpcTimeout)
		if err != nil {
			return currentBlock, mapAddresses, err
		}
	}

	if currentBlock <= endBlock.Int64() {
		traceTransaction, err := e.GetTraceTransaction(big.NewInt(currentBlock), rpc)
		if err != nil {
			return currentBlock, mapAddresses, err
		}
		addresses, err := GetAddressesFromTrace(traceTransaction, true)
		if err != nil {
			return currentBlock, mapAddresses, err
		}
		mintAddress, err := e.GetToFromEthMintLogs(big.NewInt(currentBlock), rpc)
		if err != nil {
			return currentBlock, mapAddresses, err
		}
		log.Info("Crawled block", "block", currentBlock, "addresses", len(addresses), "mintAddress", len(mintAddress))
		addresses = append(addresses, mintAddress...)
		AddAddressesToMap(addresses, mapAddresses)
		e.SaveAddresses(currentBlock, MapToAddresses(mapAddresses))
		log.Info("Wrote addresses to file", "block", currentBlock, "addresses", len(mapAddresses))
		currentBlock++
	}

	return currentBlock, mapAddresses, nil
}

func (e *EthCrawler) LoadAddresses() (int64, []*common.Address, error) {
	file, err := os.Open(e.Output)
	defer file.Close()
	if err != nil {
		return 1, nil, err
	}
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return 1, nil, err
	}
	var history EthAddresses
	if err := json.Unmarshal(byteValue, &history); err != nil {
		return 1, nil, err
	}
	return history.BlockNumber, history.Addresses, nil
}

func (e *EthCrawler) SaveAddresses(blockNumber int64, addresses []*common.Address) error {
	ethAddresses := EthAddresses{
		BlockNumber: blockNumber,
		Addresses:   addresses,
	}
	byteValue, err := json.Marshal(ethAddresses)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(e.Output, byteValue, 0644); err != nil {
		return err
	}
	return nil
}

func (e *EthCrawler) GetTraceTransaction(blockNumber *big.Int, rpc RPCMethods) (*TraceTransaction, error) {
	block, err := rpc.GetTransactionHash(e.EthClient, e.RpcTimeout, blockNumber)
	if err != nil {
		return nil, err
	}
	transactions := block.Transactions()
	if len(transactions) == 0 {
		return nil, nil
	}
	if len(transactions) == 1 {
		transactionHash := transactions[0].Hash()
		traceTransaction, err := rpc.DebugTransaction(e.RpcClient, e.RpcTimeout, &transactionHash)
		if err != nil {
			return nil, err
		}
		return traceTransaction, nil
	}
	// This should not never happen on Boba legacy chain
	// more than one transaction in a block
	return nil, fmt.Errorf("block %d has more than one transaction", blockNumber)
}

func (e *EthCrawler) GetToFromEthMintLogs(blockNumber *big.Int, rpc RPCMethods) ([]*common.Address, error) {
	ABI, err := utils.GetABI("OVM_ETH")
	if err != nil {
		return nil, err
	}
	filter := ethereum.FilterQuery{
		FromBlock: blockNumber,
		ToBlock:   blockNumber,
		Addresses: []common.Address{
			utils.OVM_ETH_ADDRESS,
		},
	}
	var addresses []*common.Address
	logs, err := rpc.GetLogs(e.EthClient, e.RpcTimeout, &filter)
	if err != nil {
		return nil, err
	}
	for _, log := range logs {
		if log.Topics[0] == ABI.Events["Transfer"].ID && len(log.Topics) == 3 {
			to := common.BytesToAddress(log.Topics[2].Bytes())
			addresses = append(addresses, &to)
			// This case is for BOBA V1 that ETH can be transferred via
			// calling the OVM_ETH contract directly
			if log.Topics[1] != (common.Hash{}) {
				from := common.BytesToAddress(log.Topics[1].Bytes())
				addresses = append(addresses, &from)
			}
		}
	}
	return addresses, nil
}

func GetAddressesFromTrace(traceTransaction *TraceTransaction, sender bool) ([]*common.Address, error) {
	if traceTransaction == nil {
		return nil, nil
	}
	var addresses []*common.Address
	calls := traceTransaction
	value := calls.Value.ToInt()
	if sender {
		addresses = append(addresses, &calls.From)
	}
	if value.Cmp(common.Big0) == 1 {
		addresses = append(addresses, &calls.From)
		addresses = append(addresses, &calls.To)
	}
	if len(calls.Calls) == 0 {
		return addresses, nil
	}
	for _, call := range calls.Calls {
		innerAddress, err := GetAddressesFromTrace(call, false)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, innerAddress...)
	}
	return addresses, nil
}

func MapAddresses(addresses []*common.Address) map[common.Address]bool {
	addressMap := make(map[common.Address]bool)
	for _, address := range addresses {
		addressMap[*address] = true
	}
	return addressMap
}

func AddAddressesToMap(addresses []*common.Address, addressMap map[common.Address]bool) {
	for _, address := range addresses {
		addressMap[*address] = true
	}
}

func MapToAddresses(addressMap map[common.Address]bool) []*common.Address {
	var addresses []*common.Address
	for address := range addressMap {
		addr := address
		addresses = append(addresses, &addr)
	}
	return addresses
}
