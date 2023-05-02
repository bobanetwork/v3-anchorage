package genesis

import (
	"fmt"
	"testing"

	opether "github.com/ethereum-optimism/optimism/op-chain-ops/ether"
	"github.com/ethereum/go-ethereum/common"
)

type MappingTest struct {
	name     string
	input    interface{}
	expected error
}

func TestValidAddressInStorage(t *testing.T) {
	storage := make(map[common.Hash]string)
	storage[opether.CalcOVMETHStorageKey(common.Address{1})] = "foo"
	storage[opether.CalcOVMETHStorageKey(common.Address{2})] = "bar"
	storage[common.BytesToHash([]byte{2})] = "baz"

	tests := []MappingTest{
		{
			name: "test 1",
			input: []*common.Address{
				{1},
				{2},
			},
			expected: nil,
		},
		{
			name: "test 2",
			input: []*common.Address{
				{1},
			},
			expected: fmt.Errorf("Some addresses in eth addresses file are not valid, valid: 2, total: 3"),
		},
	}

	for _, test := range tests {
		input, ok := test.input.([]*common.Address)
		if !ok {
			t.Errorf("Test %s failed. Could not convert input to []*common.Address", test.name)
		}
		err := ValidAddressInStorage(storage, input)
		if err != nil && test.expected != nil {
			if err.Error() != test.expected.Error() {
				t.Errorf("Test %s failed. Expected %v, got %v", test.name, test.expected, err)
			}
		} else {
			if err != nil || test.expected != nil {
				t.Errorf("Test %s failed. Expected %v, got %v", test.name, test.expected, err)
			}
		}
	}
}
