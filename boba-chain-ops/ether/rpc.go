package ether

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/tracers"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type RPC struct{}

type RPCMethods interface {
	GetBlockNumber(client *ethclient.Client, rpcTimeout time.Duration) (*big.Int, error)
	GetTransactionHash(client *ethclient.Client, rpcTimeout time.Duration, blockNumber *big.Int) (*types.Block, error)
	DebugTransaction(rpcClient *rpc.Client, rpcTimeout time.Duration, txHash *common.Hash) (*TraceTransaction, error)
	GetLogs(client *ethclient.Client, rpcTimeout time.Duration, filter *ethereum.FilterQuery) ([]types.Log, error)
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

func (r *RPC) GetTransactionHash(ethClient *ethclient.Client, rpcTimeout time.Duration, blockNumber *big.Int) (*types.Block, error) {
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()
	block, err := ethClient.BlockByNumber(ctx, blockNumber)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (r *RPC) GetBlockNumber(ethClient *ethclient.Client, rpcTimeout time.Duration) (*big.Int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()
	height, err := ethClient.BlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	return big.NewInt(int64(height)), nil
}

func (r *RPC) GetLogs(ethClient *ethclient.Client, rpcTimeout time.Duration, filter *ethereum.FilterQuery) ([]types.Log, error) {
	ctx, cancel := context.WithTimeout(context.Background(), rpcTimeout)
	defer cancel()
	logs, err := ethClient.FilterLogs(ctx, *filter)
	if err != nil {
		return nil, err
	}
	return logs, nil
}
