package ether

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/trie"
	"golang.org/x/exp/slices"
)

type TraceTransactionTest struct {
	name             string
	traceTransaction *TraceTransaction
	expected         []*common.Address
}

func TestGetAddressesFromTrace(t *testing.T) {
	traceTransactionType1 := generateTraceTranscation(
		common.HexToAddress("0x4200000000000000000000000000000000000001"),
		common.HexToAddress("0x4200000000000000000000000000000000000002"),
		big.NewInt(100),
	)
	traceTransactionType2 := generateTraceTranscation(
		common.HexToAddress("0x4200000000000000000000000000000000000003"),
		common.HexToAddress("0x4200000000000000000000000000000000000004"),
		big.NewInt(10),
	)
	traceTransactionType3 := generateTraceTranscation(
		common.HexToAddress("0x4200000000000000000000000000000000000005"),
		common.HexToAddress("0x4200000000000000000000000000000000000006"),
		big.NewInt(1),
	)
	traceTransactionType4 := generateTraceTranscation(
		common.HexToAddress("0x4200000000000000000000000000000000000007"),
		common.HexToAddress("0x4200000000000000000000000000000000000008"),
		big.NewInt(1),
	)
	traceTransactionType5 := generateTraceTranscation(
		common.HexToAddress("0x4200000000000000000000000000000000000009"),
		common.HexToAddress("0x4200000000000000000000000000000000000010"),
		big.NewInt(1),
	)
	traceTransactionType6 := generateTraceTranscation(
		common.HexToAddress("0x4200000000000000000000000000000000000011"),
		common.HexToAddress("0x4200000000000000000000000000000000000012"),
		big.NewInt(1),
	)
	traceTransactionType7 := generateTraceTranscation(
		common.HexToAddress("0x4200000000000000000000000000000000000011"),
		common.HexToAddress("0x4200000000000000000000000000000000000012"),
		big.NewInt(0),
	)
	traceTransactionType8 := generateTraceTranscation(
		common.HexToAddress("0x4200000000000000000000000000000000000011"),
		common.HexToAddress("0x4200000000000000000000000000000000000012"),
		big.NewInt(-1),
	)
	tests := make([]*TraceTransactionTest, 5)

	// calls: []
	case1 := &TraceTransactionTest{
		name:             "calls: []",
		traceTransaction: traceTransactionType1,
		expected: []*common.Address{
			&traceTransactionType1.From,
			&traceTransactionType1.To,
		},
	}
	tests[0] = case1

	// calls: [calls: []
	case2 := &TraceTransactionTest{
		name:             "calls: [calls: []",
		traceTransaction: traceTransactionType2,
		expected: []*common.Address{
			&traceTransactionType2.From,
			&traceTransactionType2.To,
		},
	}
	case2.traceTransaction.Calls = []*TraceTransaction{
		traceTransactionType1,
	}
	case2.expected = append(case1.expected, case2.expected...)
	tests[1] = case2

	// calls: [calls: [calls: []]]
	case3 := &TraceTransactionTest{
		name:             "calls: [calls: [calls: []]]",
		traceTransaction: traceTransactionType3,
		expected: []*common.Address{
			&traceTransactionType3.From,
			&traceTransactionType3.To,
		},
	}
	case3.traceTransaction.Calls = []*TraceTransaction{
		traceTransactionType2,
	}
	case3.expected = append(case2.expected, case3.expected...)
	tests[2] = case3

	// calls: [calls:[], calls:[]]
	case4 := &TraceTransactionTest{
		name:             "calls: [calls:[], calls:[]]",
		traceTransaction: traceTransactionType4,
		expected: []*common.Address{
			&traceTransactionType4.From,
			&traceTransactionType4.To,
		},
	}
	case4.traceTransaction.Calls = []*TraceTransaction{
		traceTransactionType5,
		traceTransactionType6,
	}
	case4.expected = []*common.Address{
		&traceTransactionType4.From,
		&traceTransactionType4.To,
		&traceTransactionType5.From,
		&traceTransactionType5.To,
		&traceTransactionType6.From,
		&traceTransactionType6.To,
	}
	tests[3] = case4

	// invalid test
	case5 := &TraceTransactionTest{
		name:             "Invalid test",
		traceTransaction: traceTransactionType7,
		expected:         []*common.Address{},
	}
	case5.traceTransaction.Calls = []*TraceTransaction{
		traceTransactionType8,
	}
	tests[4] = case5

	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			addresses, err := GetAddressesFromTrace(test.traceTransaction)
			if err != nil {
				t.Errorf("Test %d: Unexpected error: %s", i, err.Error())
			}
			if len(addresses) != len(test.expected) {
				t.Errorf("Test %d: Expected %d addresses, got %d", i, len(test.expected), len(addresses))
			}
			for _, address := range test.expected {
				if !slices.Contains(addresses, address) {
					t.Errorf("Test %d: Expected address %s not found", i, address.String())
				}
			}
		})
	}
}

func TestLoadAddresses(t *testing.T) {
	ethCrawler := &EthCrawler{
		RpcClient:          nil,
		EndBlock:           100,
		RpcTimeout:         10 * time.Second,
		RpcPollingInterval: 1 * time.Second,
		Output:             "invalid.json",
	}
	_, _, err := ethCrawler.LoadAddresses()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	ethCrawler.Output = "../testdata/ethaddresses.json"
	blockNumber, addresses, err := ethCrawler.LoadAddresses()
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if blockNumber != 2 {
		t.Errorf("Expected block number 2, got %d", blockNumber)
	}
	if len(addresses) != 2 {
		t.Errorf("Expected 2 addresses, got %d", len(addresses))
	}
	expectedAddresses := []common.Address{
		common.HexToAddress("0x4200000000000000000000000000000000000000"),
		common.HexToAddress("0x4200000000000000000000000000000000000001"),
	}
	output := make([]common.Address, len(addresses))
	for i, address := range addresses {
		output[i] = *address
	}
	for i, address := range expectedAddresses {
		if !slices.Contains(output, address) {
			t.Errorf("Test %d: Expected address %s not found", i, address.String())
		}
	}
}

func TestGetTraceTransaction(t *testing.T) {
	ethCrawler := &EthCrawler{
		RpcClient:          nil,
		EndBlock:           100,
		RpcTimeout:         10 * time.Second,
		RpcPollingInterval: 1 * time.Second,
		Output:             "invalid.json",
	}
	traceTransaction := generateTraceTranscation(
		common.HexToAddress("0x4200000000000000000000000000000000000001"),
		common.HexToAddress("0x4200000000000000000000000000000000000002"),
		big.NewInt(100),
	)
	header := &types.Header{
		Number: big.NewInt(1),
	}
	block := types.NewBlock(header, nil, nil, nil, nil)
	mockRPC := &MockRPC{
		traceTransaction:          traceTransaction,
		isDebugTransactionError:   false,
		block:                     block,
		isGetTransactionHashError: false,
	}
	traceResult, err := ethCrawler.GetTraceTransaction(common.Big1, mockRPC)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if traceResult != nil {
		t.Errorf("Expected trace transaction nil, got %v", traceResult)
	}
	transaction := types.NewTransaction(
		1,
		common.HexToAddress("0x4200000000000000000000000000000000000001"),
		big.NewInt(100),
		100,
		big.NewInt(100),
		[]byte{},
	)
	block = types.NewBlock(header, []*types.Transaction{transaction}, nil, nil, trie.NewStackTrie(nil))
	mockRPC.block = block
	traceResult, err = ethCrawler.GetTraceTransaction(common.Big1, mockRPC)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if traceResult == nil {
		t.Errorf("Expected trace transaction not nil, got nil")
	}
	if traceResult != traceTransaction {
		t.Errorf("Expected trace transaction %v, got %v", traceTransaction, traceResult)
	}
	mockRPC.isDebugTransactionError = true
	traceResult, err = ethCrawler.GetTraceTransaction(common.Big1, mockRPC)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if traceResult != nil {
		t.Errorf("Expected trace transaction nil, got %v", traceResult)
	}
	mockRPC.isDebugTransactionError = false
	mockRPC.isGetTransactionHashError = true
	traceResult, err = ethCrawler.GetTraceTransaction(common.Big1, mockRPC)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if traceResult != nil {
		t.Errorf("Expected trace transaction nil, got %v", traceResult)
	}
}

func generateTraceTranscation(from common.Address, to common.Address, value *big.Int) *TraceTransaction {
	return &TraceTransaction{
		From:  from,
		To:    to,
		Value: hexutil.Big(*value),
	}
}
