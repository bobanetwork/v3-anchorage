package utils

import (
	"errors"
	"testing"
)

type Test struct {
	name  string
	input string
	err   error
}

func TestGetABI(t *testing.T) {
	tests := []Test{
		{
			name:  "OVM_ETH",
			input: "OVM_ETH",
			err:   nil,
		},
		{
			name:  "Invalid",
			input: "Invalid",
			err:   errors.New("abi not found"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := GetABI(test.input)
			if err != nil && test.err == nil {
				t.Errorf("Expected no error, got %v", err)
			}
			if err == nil && test.err != nil {
				t.Errorf("Expected error %v, got nil", test.err)
			}
			if err != nil && test.err != nil && err.Error() != test.err.Error() {
				t.Errorf("Expected error %v, got %v", test.err, err)
			}
		})
	}
}
