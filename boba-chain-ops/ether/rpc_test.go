package ether

import (
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type MockRPC struct {
	traceTransaction          *TraceTransaction
	isDebugTransactionError   bool
	block                     *types.Block
	isGetTransactionHashError bool
}

func (r *MockRPC) DebugTransaction(rpcClient *rpc.Client, rpcTimeout time.Duration, txHash *common.Hash) (*TraceTransaction, error) {
	if r.isDebugTransactionError {
		return nil, errors.New("Mock DebugTransaction error")
	}
	return r.traceTransaction, nil
}

func (r *MockRPC) GetTransactionHash(rpcClient *rpc.Client, rpcTimeout time.Duration, blockNumber *hexutil.Big) (*types.Block, error) {
	if r.isGetTransactionHashError {
		return nil, errors.New("Mock GetTransactionHash error")
	}
	return r.block, nil
}

func (r *MockRPC) GetBlockNumber(rpcClient *rpc.Client, rpcTimeout time.Duration) (*hexutil.Big, error) {
	return nil, nil
}
