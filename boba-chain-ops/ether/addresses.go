package ether

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

type EthCrawler struct {
	RpcClient          *rpc.Client
	EndBlock           int64
	RpcTimeout         time.Duration
	RpcPollingInterval time.Duration
	Output             string
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

func NewEthCrawler(rpcClient *rpc.Client, endBlock int64, rpcTimeout, rpcPollingInterval time.Duration, output string) *EthCrawler {
	return &EthCrawler{
		RpcClient:          rpcClient,
		EndBlock:           endBlock,
		RpcTimeout:         rpcTimeout,
		RpcPollingInterval: rpcPollingInterval,
		Output:             output,
	}
}

func (e *EthCrawler) Start() {
	go e.Loop()
}

func (e *EthCrawler) Loop() {
	timer := time.NewTicker(e.RpcPollingInterval)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:

		}
	}
}

func (e *EthCrawler) LoadAddresses() (int64, []*common.Address, error) {
	file, err := os.Open(e.Output)
	defer file.Close()
	if err != nil {
		return 0, nil, err
	}
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return 0, nil, err
	}
	var history EthAddresses
	if err := json.Unmarshal(byteValue, &history); err != nil {
		return 0, nil, err
	}
	return history.BlockNumber, history.Addresses, nil
}

func (e *EthCrawler) GetTraceTransaction(blockNumber *big.Int, rpc RPCMethods) (*TraceTransaction, error) {
	block, err := rpc.GetTransactionHash(e.RpcClient, e.RpcTimeout, (*hexutil.Big)(blockNumber))
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
	// This should not never happened on Boba legacy chain
	// more than one transaction in a block
	return nil, fmt.Errorf("block %d has more than one transaction", blockNumber)
}

func GetAddressesFromTrace(traceTransaction *TraceTransaction) ([]*common.Address, error) {
	var addresses []*common.Address
	calls := traceTransaction
	value := calls.Value.ToInt()
	if value.Cmp(common.Big0) == 1 {
		addresses = append(addresses, &calls.From)
		addresses = append(addresses, &calls.To)
	}
	if len(calls.Calls) == 0 {
		return addresses, nil
	}
	for _, call := range calls.Calls {
		innerAddress, err := GetAddressesFromTrace(call)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, innerAddress...)
	}
	return addresses, nil
}
