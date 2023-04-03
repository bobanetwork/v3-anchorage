// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BobaHCHelperMetaData contains all meta data concerning the BobaHCHelper contract.
var BobaHCHelperMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"random\",\"type\":\"uint256\"}],\"name\":\"OffchainRandom\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rType\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_method\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"GetResponse\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetSimpleRandom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NextSimpleRandom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cacheKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"PutResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"PutSimpleRandom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b031916301790556108ed806100326000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c8063a101dc2111610050578063a101dc211461009c578063d5054e9e146100a5578063dfc98ae8146100c557600080fd5b80637dbf7c101461006c5780639750097114610087575b600080fd5b6100746100d8565b6040519081526020015b60405180910390f35b61009a610095366004610483565b600155565b005b61007460015481565b6100b86100b3366004610576565b610195565b60405161007e919061064a565b61009a6100d336600461069b565b61038e565b600060015460000361014b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f52616e646f6d206e756d626572206e6f742067656e657261746564000000000060448201526064015b60405180910390fd5b60018054600091829055604080519283526020830182905290917f450d62889c3a6e19c9586840ce9c21040b90d81950fe31f2ba982090adaf53e8910160405180910390a1919050565b60608463ffffffff16600203610217576000825111610210576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f547572696e673a476574526573706f6e73653a6e6f207061796c6f61640000006044820152606401610142565b5080610386565b600084848460405160200161022e939291906106e2565b604051602081830303815290604052805190602001209050600060026000838152602001908152602001600020805461026690610725565b9050116102cf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f4d697373696e6720636163686520656e747279000000000000000000000000006044820152606401610142565b600081815260026020526040812080546102e890610725565b80601f016020809104026020016040519081016040528092919081815260200182805461031490610725565b80156103615780601f1061033657610100808354040283529160200191610361565b820191906000526020600020905b81548152906001019060200180831161034457829003601f168201915b5050506000858152600260205260408120939450610382939250905061042d565b9150505b949350505050565b600082815260026020526040902080546103a790610725565b159050610410576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4445425547202d20416c726561647920657869737473000000000000000000006044820152606401610142565b600082815260026020526040902061042882826107c6565b505050565b50805461043990610725565b6000825580601f10610449575050565b601f016020900490600052602060002090810190610467919061046a565b50565b5b8082111561047f576000815560010161046b565b5090565b60006020828403121561049557600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126104dc57600080fd5b813567ffffffffffffffff808211156104f7576104f761049c565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561053d5761053d61049c565b8160405283815286602085880101111561055657600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000806080858703121561058c57600080fd5b843563ffffffff811681146105a057600080fd5b9350602085013567ffffffffffffffff808211156105bd57600080fd5b6105c9888389016104cb565b945060408701359150808211156105df57600080fd5b6105eb888389016104cb565b9350606087013591508082111561060157600080fd5b5061060e878288016104cb565b91505092959194509250565b60005b8381101561063557818101518382015260200161061d565b83811115610644576000848401525b50505050565b602081526000825180602084015261066981604085016020870161061a565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b600080604083850312156106ae57600080fd5b82359150602083013567ffffffffffffffff8111156106cc57600080fd5b6106d8858286016104cb565b9150509250929050565b600084516106f481846020890161061a565b84519083019061070881836020890161061a565b845191019061071b81836020880161061a565b0195945050505050565b600181811c9082168061073957607f821691505b602082108103610772577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f82111561042857600081815260208120601f850160051c8101602086101561079f5750805b601f850160051c820191505b818110156107be578281556001016107ab565b505050505050565b815167ffffffffffffffff8111156107e0576107e061049c565b6107f4816107ee8454610725565b84610778565b602080601f83116001811461084757600084156108115750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556107be565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b8281101561089457888601518255948401946001909101908401610875565b50858210156108d057878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b0190555056fea164736f6c634300080f000a",
}

// BobaHCHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use BobaHCHelperMetaData.ABI instead.
var BobaHCHelperABI = BobaHCHelperMetaData.ABI

// BobaHCHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BobaHCHelperMetaData.Bin instead.
var BobaHCHelperBin = BobaHCHelperMetaData.Bin

// DeployBobaHCHelper deploys a new Ethereum contract, binding an instance of BobaHCHelper to it.
func DeployBobaHCHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BobaHCHelper, error) {
	parsed, err := BobaHCHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BobaHCHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BobaHCHelper{BobaHCHelperCaller: BobaHCHelperCaller{contract: contract}, BobaHCHelperTransactor: BobaHCHelperTransactor{contract: contract}, BobaHCHelperFilterer: BobaHCHelperFilterer{contract: contract}}, nil
}

// BobaHCHelper is an auto generated Go binding around an Ethereum contract.
type BobaHCHelper struct {
	BobaHCHelperCaller     // Read-only binding to the contract
	BobaHCHelperTransactor // Write-only binding to the contract
	BobaHCHelperFilterer   // Log filterer for contract events
}

// BobaHCHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type BobaHCHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BobaHCHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BobaHCHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BobaHCHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BobaHCHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BobaHCHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BobaHCHelperSession struct {
	Contract     *BobaHCHelper     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BobaHCHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BobaHCHelperCallerSession struct {
	Contract *BobaHCHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BobaHCHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BobaHCHelperTransactorSession struct {
	Contract     *BobaHCHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BobaHCHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type BobaHCHelperRaw struct {
	Contract *BobaHCHelper // Generic contract binding to access the raw methods on
}

// BobaHCHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BobaHCHelperCallerRaw struct {
	Contract *BobaHCHelperCaller // Generic read-only contract binding to access the raw methods on
}

// BobaHCHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BobaHCHelperTransactorRaw struct {
	Contract *BobaHCHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBobaHCHelper creates a new instance of BobaHCHelper, bound to a specific deployed contract.
func NewBobaHCHelper(address common.Address, backend bind.ContractBackend) (*BobaHCHelper, error) {
	contract, err := bindBobaHCHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BobaHCHelper{BobaHCHelperCaller: BobaHCHelperCaller{contract: contract}, BobaHCHelperTransactor: BobaHCHelperTransactor{contract: contract}, BobaHCHelperFilterer: BobaHCHelperFilterer{contract: contract}}, nil
}

// NewBobaHCHelperCaller creates a new read-only instance of BobaHCHelper, bound to a specific deployed contract.
func NewBobaHCHelperCaller(address common.Address, caller bind.ContractCaller) (*BobaHCHelperCaller, error) {
	contract, err := bindBobaHCHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BobaHCHelperCaller{contract: contract}, nil
}

// NewBobaHCHelperTransactor creates a new write-only instance of BobaHCHelper, bound to a specific deployed contract.
func NewBobaHCHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*BobaHCHelperTransactor, error) {
	contract, err := bindBobaHCHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BobaHCHelperTransactor{contract: contract}, nil
}

// NewBobaHCHelperFilterer creates a new log filterer instance of BobaHCHelper, bound to a specific deployed contract.
func NewBobaHCHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*BobaHCHelperFilterer, error) {
	contract, err := bindBobaHCHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BobaHCHelperFilterer{contract: contract}, nil
}

// bindBobaHCHelper binds a generic wrapper to an already deployed contract.
func bindBobaHCHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BobaHCHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BobaHCHelper *BobaHCHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BobaHCHelper.Contract.BobaHCHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BobaHCHelper *BobaHCHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.BobaHCHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BobaHCHelper *BobaHCHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.BobaHCHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BobaHCHelper *BobaHCHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BobaHCHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BobaHCHelper *BobaHCHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BobaHCHelper *BobaHCHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.contract.Transact(opts, method, params...)
}

// NextSimpleRandom is a free data retrieval call binding the contract method 0xa101dc21.
//
// Solidity: function NextSimpleRandom() view returns(uint256)
func (_BobaHCHelper *BobaHCHelperCaller) NextSimpleRandom(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BobaHCHelper.contract.Call(opts, &out, "NextSimpleRandom")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextSimpleRandom is a free data retrieval call binding the contract method 0xa101dc21.
//
// Solidity: function NextSimpleRandom() view returns(uint256)
func (_BobaHCHelper *BobaHCHelperSession) NextSimpleRandom() (*big.Int, error) {
	return _BobaHCHelper.Contract.NextSimpleRandom(&_BobaHCHelper.CallOpts)
}

// NextSimpleRandom is a free data retrieval call binding the contract method 0xa101dc21.
//
// Solidity: function NextSimpleRandom() view returns(uint256)
func (_BobaHCHelper *BobaHCHelperCallerSession) NextSimpleRandom() (*big.Int, error) {
	return _BobaHCHelper.Contract.NextSimpleRandom(&_BobaHCHelper.CallOpts)
}

// GetResponse is a paid mutator transaction binding the contract method 0xd5054e9e.
//
// Solidity: function GetResponse(uint32 rType, string _url, string _method, bytes _payload) returns(bytes)
func (_BobaHCHelper *BobaHCHelperTransactor) GetResponse(opts *bind.TransactOpts, rType uint32, _url string, _method string, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "GetResponse", rType, _url, _method, _payload)
}

// GetResponse is a paid mutator transaction binding the contract method 0xd5054e9e.
//
// Solidity: function GetResponse(uint32 rType, string _url, string _method, bytes _payload) returns(bytes)
func (_BobaHCHelper *BobaHCHelperSession) GetResponse(rType uint32, _url string, _method string, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.GetResponse(&_BobaHCHelper.TransactOpts, rType, _url, _method, _payload)
}

// GetResponse is a paid mutator transaction binding the contract method 0xd5054e9e.
//
// Solidity: function GetResponse(uint32 rType, string _url, string _method, bytes _payload) returns(bytes)
func (_BobaHCHelper *BobaHCHelperTransactorSession) GetResponse(rType uint32, _url string, _method string, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.GetResponse(&_BobaHCHelper.TransactOpts, rType, _url, _method, _payload)
}

// GetSimpleRandom is a paid mutator transaction binding the contract method 0x7dbf7c10.
//
// Solidity: function GetSimpleRandom() returns(uint256)
func (_BobaHCHelper *BobaHCHelperTransactor) GetSimpleRandom(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "GetSimpleRandom")
}

// GetSimpleRandom is a paid mutator transaction binding the contract method 0x7dbf7c10.
//
// Solidity: function GetSimpleRandom() returns(uint256)
func (_BobaHCHelper *BobaHCHelperSession) GetSimpleRandom() (*types.Transaction, error) {
	return _BobaHCHelper.Contract.GetSimpleRandom(&_BobaHCHelper.TransactOpts)
}

// GetSimpleRandom is a paid mutator transaction binding the contract method 0x7dbf7c10.
//
// Solidity: function GetSimpleRandom() returns(uint256)
func (_BobaHCHelper *BobaHCHelperTransactorSession) GetSimpleRandom() (*types.Transaction, error) {
	return _BobaHCHelper.Contract.GetSimpleRandom(&_BobaHCHelper.TransactOpts)
}

// PutResponse is a paid mutator transaction binding the contract method 0xdfc98ae8.
//
// Solidity: function PutResponse(bytes32 cacheKey, bytes _payload) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) PutResponse(opts *bind.TransactOpts, cacheKey [32]byte, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "PutResponse", cacheKey, _payload)
}

// PutResponse is a paid mutator transaction binding the contract method 0xdfc98ae8.
//
// Solidity: function PutResponse(bytes32 cacheKey, bytes _payload) returns()
func (_BobaHCHelper *BobaHCHelperSession) PutResponse(cacheKey [32]byte, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.PutResponse(&_BobaHCHelper.TransactOpts, cacheKey, _payload)
}

// PutResponse is a paid mutator transaction binding the contract method 0xdfc98ae8.
//
// Solidity: function PutResponse(bytes32 cacheKey, bytes _payload) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) PutResponse(cacheKey [32]byte, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.PutResponse(&_BobaHCHelper.TransactOpts, cacheKey, _payload)
}

// PutSimpleRandom is a paid mutator transaction binding the contract method 0x97500971.
//
// Solidity: function PutSimpleRandom(uint256 val) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) PutSimpleRandom(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "PutSimpleRandom", val)
}

// PutSimpleRandom is a paid mutator transaction binding the contract method 0x97500971.
//
// Solidity: function PutSimpleRandom(uint256 val) returns()
func (_BobaHCHelper *BobaHCHelperSession) PutSimpleRandom(val *big.Int) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.PutSimpleRandom(&_BobaHCHelper.TransactOpts, val)
}

// PutSimpleRandom is a paid mutator transaction binding the contract method 0x97500971.
//
// Solidity: function PutSimpleRandom(uint256 val) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) PutSimpleRandom(val *big.Int) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.PutSimpleRandom(&_BobaHCHelper.TransactOpts, val)
}

// BobaHCHelperOffchainRandomIterator is returned from FilterOffchainRandom and is used to iterate over the raw logs and unpacked data for OffchainRandom events raised by the BobaHCHelper contract.
type BobaHCHelperOffchainRandomIterator struct {
	Event *BobaHCHelperOffchainRandom // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BobaHCHelperOffchainRandomIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BobaHCHelperOffchainRandom)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BobaHCHelperOffchainRandom)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BobaHCHelperOffchainRandomIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BobaHCHelperOffchainRandomIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BobaHCHelperOffchainRandom represents a OffchainRandom event raised by the BobaHCHelper contract.
type BobaHCHelperOffchainRandom struct {
	Version *big.Int
	Random  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOffchainRandom is a free log retrieval operation binding the contract event 0x450d62889c3a6e19c9586840ce9c21040b90d81950fe31f2ba982090adaf53e8.
//
// Solidity: event OffchainRandom(uint256 version, uint256 random)
func (_BobaHCHelper *BobaHCHelperFilterer) FilterOffchainRandom(opts *bind.FilterOpts) (*BobaHCHelperOffchainRandomIterator, error) {

	logs, sub, err := _BobaHCHelper.contract.FilterLogs(opts, "OffchainRandom")
	if err != nil {
		return nil, err
	}
	return &BobaHCHelperOffchainRandomIterator{contract: _BobaHCHelper.contract, event: "OffchainRandom", logs: logs, sub: sub}, nil
}

// WatchOffchainRandom is a free log subscription operation binding the contract event 0x450d62889c3a6e19c9586840ce9c21040b90d81950fe31f2ba982090adaf53e8.
//
// Solidity: event OffchainRandom(uint256 version, uint256 random)
func (_BobaHCHelper *BobaHCHelperFilterer) WatchOffchainRandom(opts *bind.WatchOpts, sink chan<- *BobaHCHelperOffchainRandom) (event.Subscription, error) {

	logs, sub, err := _BobaHCHelper.contract.WatchLogs(opts, "OffchainRandom")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BobaHCHelperOffchainRandom)
				if err := _BobaHCHelper.contract.UnpackLog(event, "OffchainRandom", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOffchainRandom is a log parse operation binding the contract event 0x450d62889c3a6e19c9586840ce9c21040b90d81950fe31f2ba982090adaf53e8.
//
// Solidity: event OffchainRandom(uint256 version, uint256 random)
func (_BobaHCHelper *BobaHCHelperFilterer) ParseOffchainRandom(log types.Log) (*BobaHCHelperOffchainRandom, error) {
	event := new(BobaHCHelperOffchainRandom)
	if err := _BobaHCHelper.contract.UnpackLog(event, "OffchainRandom", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
