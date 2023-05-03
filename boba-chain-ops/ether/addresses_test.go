package ether

import (
	"math/big"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/boba-chain-ops/utils"

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

type MappingTest struct {
	name     string
	input    interface{}
	expected interface{}
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
		name:             "Test 1",
		traceTransaction: traceTransactionType1,
		expected: []*common.Address{
			&traceTransactionType1.From,
			&traceTransactionType1.From,
			&traceTransactionType1.To,
		},
	}
	tests[0] = case1

	// calls: [calls: []
	case2 := &TraceTransactionTest{
		name:             "Test 2",
		traceTransaction: traceTransactionType2,
		expected: []*common.Address{
			&traceTransactionType2.From,
			&traceTransactionType2.From,
			&traceTransactionType2.To,
		},
	}
	case2.traceTransaction.Calls = []*TraceTransaction{
		traceTransactionType1,
	}
	case2.expected = append([]*common.Address{
		&traceTransactionType1.From,
		&traceTransactionType1.To,
	}, case2.expected...)
	tests[1] = case2

	// calls: [calls: [calls: []]]
	case3 := &TraceTransactionTest{
		name:             "Test 3",
		traceTransaction: traceTransactionType3,
		expected: []*common.Address{
			&traceTransactionType3.From,
			&traceTransactionType3.From,
			&traceTransactionType3.To,
		},
	}
	case3.traceTransaction.Calls = []*TraceTransaction{
		traceTransactionType2,
	}
	case3.expected = append([]*common.Address{
		&traceTransactionType1.From,
		&traceTransactionType1.To,
		&traceTransactionType2.From,
		&traceTransactionType2.To,
	}, case3.expected...)
	tests[2] = case3

	// calls: [calls:[], calls:[]]
	case4 := &TraceTransactionTest{
		name:             "Test 4",
		traceTransaction: traceTransactionType4,
		expected: []*common.Address{
			&traceTransactionType4.From,
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
		&traceTransactionType4.From,
		&traceTransactionType4.To,
		&traceTransactionType5.From,
		&traceTransactionType5.To,
		&traceTransactionType6.From,
		&traceTransactionType6.To,
	}
	tests[3] = case4

	// invalid payload
	case5 := &TraceTransactionTest{
		name:             "Test 5",
		traceTransaction: traceTransactionType7,
		expected:         []*common.Address{},
	}
	case5.traceTransaction.Calls = []*TraceTransaction{
		traceTransactionType8,
	}
	case5.expected = []*common.Address{
		&traceTransactionType7.From,
	}
	tests[4] = case5

	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			addresses, err := GetAddressesFromTrace(test.traceTransaction, true)
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
			for _, address := range addresses {
				if !slices.Contains(test.expected, address) {
					t.Errorf("Test %d: Unexpected address %s", i, address.String())
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
		OutputPath:         "invalid.json",
	}
	_, _, err := ethCrawler.LoadAddresses()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	ethCrawler.OutputPath = "../testdata/ethaddresses.json"
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
	address1, address2 := common.HexToAddress("0x4200000000000000000000000000000000000000"), common.HexToAddress("0x4200000000000000000000000000000000000001")
	expectedAddresses := []*common.Address{&address1, &address2}
	if !reflect.DeepEqual(addresses, expectedAddresses) {
		t.Errorf("Expected addresses %v, got %v", expectedAddresses, addresses)
	}
}

func TestSaveAddresses(t *testing.T) {
	ethCrawler := &EthCrawler{
		RpcClient:          nil,
		EndBlock:           100,
		RpcTimeout:         10 * time.Second,
		RpcPollingInterval: 1 * time.Second,
		OutputPath:         "../testdata/saveaddresses.json",
	}
	address1, address2 := common.Address{1}, common.Address{2}
	addresses := []*common.Address{&address1, &address2}
	if err := ethCrawler.SaveAddresses(2, addresses); err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	defer os.Remove("../testdata/saveaddresses.json")
	blockNumber, inputAddresses, err := ethCrawler.LoadAddresses()
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if blockNumber != 2 {
		t.Errorf("Expected block number 2, got %d", blockNumber)
	}
	if !reflect.DeepEqual(addresses, inputAddresses) {
		t.Errorf("Expected addresses %v, got %v", addresses, inputAddresses)
	}
}

func TestGetTraceTransaction(t *testing.T) {
	ethCrawler := &EthCrawler{
		RpcClient:          nil,
		EndBlock:           100,
		RpcTimeout:         10 * time.Second,
		RpcPollingInterval: 1 * time.Second,
		OutputPath:         "invalid.json",
	}
	traceTransaction := generateTraceTranscation(
		common.HexToAddress("0x4200000000000000000000000000000000000001"),
		common.HexToAddress("0x4200000000000000000000000000000000000002"),
		big.NewInt(100),
	)
	header := &types.Header{
		Number: big.NewInt(1),
	}
	block := types.NewBlock(header, nil, nil, nil, trie.NewStackTrie(nil))
	mockRPC := &MockRPC{
		traceTransaction:          traceTransaction,
		isDebugTransactionError:   false,
		block:                     block,
		isGetTransactionHashError: false,
		logs:                      []types.Log{},
		isGetLogsError:            false,
	}
	traceResult, err := ethCrawler.GetTraceTransaction(common.Big1, mockRPC)
	if err != nil {
		t.Errorf("Unexpected error: %s", err.Error())
	}
	if traceResult != nil {
		t.Errorf("Expected trace transaction nil, got %v", traceResult)
	}
	transaction := types.NewTransaction(
		0,
		common.HexToAddress("0x4200000000000000000000000000000000000001"),
		common.Big1,
		100,
		common.Big1,
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
	block = types.NewBlock(header, []*types.Transaction{transaction, transaction}, nil, nil, trie.NewStackTrie(nil))
	mockRPC.block = block
	traceResult, err = ethCrawler.GetTraceTransaction(common.Big1, mockRPC)
	if err == nil {
		t.Errorf("Unexpected error: %s", err.Error())
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

func TestMapAddresses(t *testing.T) {
	tests := []MappingTest{
		{
			name: "Test 1",
			input: []*common.Address{
				{1},
				{2},
			},
			expected: map[common.Address]bool{
				{1}: true,
				{2}: true,
			},
		},
		{
			name: "Test 2",
			input: []*common.Address{
				{1},
				{1},
			},
			expected: map[common.Address]bool{
				{1}: true,
			},
		},
	}
	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input, ok := test.input.([]*common.Address)
			if !ok {
				t.Errorf("Test %d: Invalid input", i)
			}
			expected, ok := test.expected.(map[common.Address]bool)
			if !ok {
				t.Errorf("Test %d: Invalid expected", i)
			}
			result := MapAddresses(input)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Test %d: Expected %v, got %v", i, expected, result)
			}
		})
	}
}

func TestAddAddressesToMap(t *testing.T) {
	type inputStruct struct {
		addresses  []*common.Address
		addressMap map[common.Address]bool
	}
	tests := []MappingTest{
		{
			name: "Test 1",
			input: inputStruct{
				addresses: []*common.Address{
					{1},
					{2},
				},
				addressMap: map[common.Address]bool{
					{1}: true,
				},
			},
			expected: map[common.Address]bool{
				{1}: true,
				{2}: true,
			},
		},
		{
			name: "Test 2",
			input: inputStruct{
				addresses: []*common.Address{
					{1},
					{1},
				},
				addressMap: map[common.Address]bool{
					{1}: true,
				},
			},
			expected: map[common.Address]bool{
				{1}: true,
			},
		},
	}
	for i, test := range tests {
		input, ok := test.input.(inputStruct)
		if !ok {
			t.Errorf("Test %d: Invalid input", i)
		}
		expected, ok := test.expected.(map[common.Address]bool)
		if !ok {
			t.Errorf("Test %d: Invalid expected", i)
		}
		t.Run(test.name, func(t *testing.T) {
			AddAddressesToMap(input.addresses, input.addressMap)
			if !reflect.DeepEqual(input.addressMap, expected) {
				t.Errorf("Test %d: Expected %v, got %v", i, expected, input.addressMap)
			}
		})

	}
}

func TestMapToAddresses(t *testing.T) {
	tests := []MappingTest{
		{
			name: "Test 1",
			input: map[common.Address]bool{
				{1}: true,
				{2}: true,
			},
			expected: []common.Address{
				{1},
				{2},
			},
		},
	}
	for i, test := range tests {
		input, ok := test.input.(map[common.Address]bool)
		if !ok {
			t.Errorf("Test %d: Invalid input", i)
		}
		expected, ok := test.expected.([]common.Address)
		if !ok {
			t.Errorf("Test %d: Invalid expected", i)
		}
		t.Run(test.name, func(t *testing.T) {
			result := MapToAddresses(input)
			output := make([]common.Address, len(result))
			for i, address := range result {
				output[i] = *address
			}
			for i, address := range output {
				if !slices.Contains(expected, address) {
					t.Errorf("Test %d: Expected %v, got %v", i, expected, output)
				}
			}
			for i, address := range expected {
				if !slices.Contains(output, address) {
					t.Errorf("Test %d: Expected %v, got %v", i, expected, output)
				}
			}
		})
	}
}

func TestGetToFromEthMintLogs(t *testing.T) {
	ABI, err := utils.GetABI("OVM_ETH")
	if err != nil {
		t.Errorf("Error getting ABI: %v", err)
	}
	ethCrawler := &EthCrawler{
		RpcClient:          nil,
		EndBlock:           100,
		RpcTimeout:         10 * time.Second,
		RpcPollingInterval: 1 * time.Second,
		OutputPath:         "invalid.json",
	}
	mockRPC := &MockRPC{
		traceTransaction:          &TraceTransaction{},
		isDebugTransactionError:   false,
		block:                     &types.Block{},
		isGetTransactionHashError: false,
		logs:                      []types.Log{},
		isGetLogsError:            false,
	}
	tests := []MappingTest{
		{
			name: "Test 1",
			input: []types.Log{
				{
					Address: utils.OVM_ETH_ADDRESS,
					Topics: []common.Hash{
						ABI.Events["Transfer"].ID,
						common.BytesToHash(common.Address{1}.Bytes()),
						common.BytesToHash(common.Address{2}.Bytes()),
					},
				},
			},
			expected: []common.Address{
				{2},
				{1},
			},
		},
		{
			name: "Test 2",
			input: []types.Log{
				{
					Address: utils.OVM_ETH_ADDRESS,
					Topics: []common.Hash{
						ABI.Events["Transfer"].ID,
						{0},
						common.BytesToHash(common.Address{1}.Bytes()),
					},
				},
			},
			expected: []common.Address{
				{1},
			},
		},
		{
			name: "Test 3",
			input: []types.Log{
				{
					Address: utils.OVM_ETH_ADDRESS,
					Topics: []common.Hash{
						ABI.Events["Transfer"].ID,
						common.BytesToHash(common.Address{1}.Bytes()),
						common.BytesToHash(common.Address{2}.Bytes()),
					},
				},
				{
					Address: utils.OVM_ETH_ADDRESS,
					Topics: []common.Hash{
						ABI.Events["Transfer"].ID,
						common.BytesToHash(common.Address{2}.Bytes()),
						common.BytesToHash(common.Address{1}.Bytes()),
					},
				},
			},
			expected: []common.Address{
				{2},
				{1},
				{1},
				{2},
			},
		},
		{
			name: "Test 4",
			input: []types.Log{
				{
					Address: utils.OVM_ETH_ADDRESS,
					Topics: []common.Hash{
						ABI.Events["Transfer"].ID,
						{1},
					},
				},
			},
			expected: []common.Address{},
		},
		{
			name: "Test 5",
			input: []types.Log{
				{
					Address: utils.OVM_ETH_ADDRESS,
					Topics: []common.Hash{
						{1},
						{2},
						{3},
					},
				},
			},
			expected: []common.Address{},
		},
	}
	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input, ok := test.input.([]types.Log)
			if !ok {
				t.Errorf("Test %d: Invalid input", i)
			}
			expected, ok := test.expected.([]common.Address)
			if !ok {
				t.Errorf("Test %d: Invalid expected", i)
			}
			mockRPC.logs = input
			result, err := ethCrawler.GetToFromEthMintLogs(common.Big0, mockRPC)
			if err != nil {
				t.Errorf("Test %d: Error getting to/from addresses: %v", i, err)
			}

			output := make([]common.Address, len(result))
			for i, address := range result {
				output[i] = *address
			}
			if !reflect.DeepEqual(expected, output) {
				t.Errorf("Test %d: Expected %v, got %v", i, expected, output)
			}
		})
	}

}

func generateTraceTranscation(from common.Address, to common.Address, value *big.Int) *TraceTransaction {
	return &TraceTransaction{
		From:  from,
		To:    to,
		Value: hexutil.Big(*value),
	}
}
