package utils

import (
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const EthInterface = `
[
	{"anonymous": false,"inputs": [{"indexed": true,"internalType": "address","name": "_account","type": "address"},{"indexed": false,"internalType": "uint256","name": "_amount","type": "uint256"}],"name": "Mint","type": "event"},
	{"anonymous": false,"inputs": [{"indexed": true,"internalType": "address","name": "from","type": "address"},{"indexed": true,"internalType": "address","name": "to","type": "address"},{"indexed": false,"internalType": "uint256","name": "value","type": "uint256"}],"name": "Transfer","type": "event"}
]`

func GetABI(name string) (*abi.ABI, error) {
	switch name {
	case "OVM_ETH":
		abi, err := abi.JSON(strings.NewReader(EthInterface))
		if err != nil {
			return nil, err
		}
		return &abi, nil
	}
	return nil, errors.New("abi not found")
}
