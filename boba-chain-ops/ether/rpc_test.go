package ether

import (
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type MockRPC struct {
	traceTransaction          *TraceTransaction
	isDebugTransactionError   bool
	block                     *types.Block
	isGetTransactionHashError bool
	logs                      []types.Log
	isGetLogsError            bool
}

func (r *MockRPC) GetBlockNumber(ethClient *ethclient.Client, rpcTimeout time.Duration) (*big.Int, error) {
	return nil, nil
}

func (r *MockRPC) GetTransactionHash(ethClient *ethclient.Client, rpcTimeout time.Duration, blockNumber *big.Int) (*types.Block, error) {
	if r.isGetTransactionHashError {
		return nil, errors.New("Mock GetTransactionHash error")
	}
	return r.block, nil
}

func (r *MockRPC) DebugTransaction(rpcClient *rpc.Client, rpcTimeout time.Duration, txHash *common.Hash) (*TraceTransaction, error) {
	if r.isDebugTransactionError {
		return nil, errors.New("Mock DebugTransaction error")
	}
	return r.traceTransaction, nil
}

func (r *MockRPC) GetLogs(ethClient *ethclient.Client, rpcTimeout time.Duration, filter *ethereum.FilterQuery) ([]types.Log, error) {
	return r.logs, nil
}
