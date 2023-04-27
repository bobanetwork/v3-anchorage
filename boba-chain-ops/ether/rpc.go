package ether

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/rpc"
)

type RPC struct{}

type RPCMethods interface {
	GetBlockNumber(rpcClient *rpc.Client, rpcTimeout time.Duration) (*hexutil.Big, error)
	GetTransactionHash(rpcClient *rpc.Client, rpcTimeout time.Duration, blockNumber *hexutil.Big) (*types.Block, error)
	DebugTransaction(rpcClient *rpc.Client, rpcTimeout time.Duration, txHash *common.Hash) (*TraceTransaction, error)
}

func (r *RPC) DebugTransaction(rpcClient *rpc.Client, rpcTimeout time.Duration, txHash *common.Hash) (*TraceTransaction, error) {
	tracerType := "callTracer"
	config := &tracers.TraceConfig{
		Tracer: &tracerType,
	}
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()
	var traceResult TraceTransaction
	if err := rpcClient.CallContext(ctx, &traceResult, "debug_traceTransaction", txHash, config); err != nil {
		return nil, err
	}
	return &traceResult, nil
}

func (r *RPC) GetTransactionHash(rpcClient *rpc.Client, rpcTimeout time.Duration, blockNumber *hexutil.Big) (*types.Block, error) {
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()
	var block *types.Block
	if err := rpcClient.CallContext(ctx, &block, "eth_getBlockByNumber", blockNumber, true); err != nil {
		return nil, err
	}
	return block, nil
}

func (r *RPC) GetBlockNumber(rpcClient *rpc.Client, rpcTimeout time.Duration) (*hexutil.Big, error) {
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()
	var height *hexutil.Big
	if err := rpcClient.CallContext(ctx, &height, "eth_blockNumber"); err != nil {
		return nil, err
	}
	return height, nil
}
