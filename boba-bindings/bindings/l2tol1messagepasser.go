// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"fmt"
	"math/big"
	"reflect"
	"strings"

	ethereum "github.com/ledgerwatch/erigon"
	libcommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon/accounts/abi"
	"github.com/ledgerwatch/erigon/accounts/abi/bind"
	"github.com/ledgerwatch/erigon/core/types"
	"github.com/ledgerwatch/erigon/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = libcommon.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// L2ToL1MessagePasserABI is the input ABI used to generate the binding from.
const L2ToL1MessagePasserABI = "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"MESSAGE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateETHERC20Withdrawal\",\"inputs\":[{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initiateWithdrawal\",\"inputs\":[{\"name\":\"_target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"messageNonce\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"sentMessages\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"MessagePassed\",\"inputs\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"withdrawalHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawerBalanceBurnt\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false}]"

// L2ToL1MessagePasserBin is the compiled bytecode used for deploying new contracts.
var L2ToL1MessagePasserBin = "0x608060405234801561001057600080fd5b50610a19806100206000396000f3fe6080604052600436106100745760003560e01c806382e3702d1161004e57806382e3702d14610135578063c2b3e5ac14610175578063dfd44e4c14610188578063ecc70428146101a857600080fd5b80633f827a5a1461009d57806344df8e70146100ca57806354fd4d50146100df57600080fd5b366100985761009633620186a06040518060200160405280600081525061020d565b005b600080fd5b3480156100a957600080fd5b506100b2600181565b60405161ffff90911681526020015b60405180910390f35b3480156100d657600080fd5b506100966103d1565b3480156100eb57600080fd5b506101286040518060400160405280600581526020017f312e312e3000000000000000000000000000000000000000000000000000000081525081565b6040516100c1919061078f565b34801561014157600080fd5b506101656101503660046107a9565b60006020819052908152604090205460ff1681565b60405190151581526020016100c1565b6100966101833660046108c5565b61020d565b34801561019457600080fd5b506100966101a336600461091c565b610409565b3480156101b457600080fd5b506101ff6001547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b6040519081526020016100c1565b60006102a36040518060c001604052806102676001547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b815233602082015273ffffffffffffffffffffffffffffffffffffffff871660408201523460608201526080810186905260a00184905261069c565b600081815260208190526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055905073ffffffffffffffffffffffffffffffffffffffff84163361033e6001547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b7f02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a274495505434878787604051610373949392919061097d565b60405180910390a45050600180547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8082168301167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b476103db816106e9565b60405181907f7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f90600090a250565b33734200000000000000000000000000000000000010146104d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604960248201527f4c32546f4c314d6573736167655061737365723a206f6e6c7920746865204c3260448201527f207374616e64617264206272696467652063616e20696e69746961746520776960648201527f746864726177616c730000000000000000000000000000000000000000000000608482015260a40160405180910390fd5b600061056d6040518060c001604052806105306001547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b815233602082015273ffffffffffffffffffffffffffffffffffffffff88166040820152606081018790526080810186905260a00184905261069c565b600081815260208190526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055905073ffffffffffffffffffffffffffffffffffffffff8516336106086001547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b7f02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a27449550548787878760405161063d949392919061097d565b60405180910390a45050600180547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8082168301167fffff000000000000000000000000000000000000000000000000000000000000909116179055505050565b80516020808301516040808501516060860151608087015160a088015193516000976106cc9790969591016109ad565b604051602081830303815290604052805190602001209050919050565b806040516106f690610718565b6040518091039082f0905080158015610713573d6000803e3d6000fd5b505050565b600880610a0583390190565b6000815180845260005b8181101561074a5760208185018101518683018201520161072e565b8181111561075c576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006107a26020830184610724565b9392505050565b6000602082840312156107bb57600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff811681146107e657600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261082b57600080fd5b813567ffffffffffffffff80821115610846576108466107eb565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561088c5761088c6107eb565b816040528381528660208588010111156108a557600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156108da57600080fd5b6108e3846107c2565b925060208401359150604084013567ffffffffffffffff81111561090657600080fd5b6109128682870161081a565b9150509250925092565b6000806000806080858703121561093257600080fd5b61093b856107c2565b93506020850135925060408501359150606085013567ffffffffffffffff81111561096557600080fd5b6109718782880161081a565b91505092959194509250565b84815283602082015260806040820152600061099c6080830185610724565b905082606083015295945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a08301526109f860c0830184610724565b9897505050505050505056fe608060405230fffea164736f6c634300080f000a"

// DeployL2ToL1MessagePasser deploys a new Ethereum contract, binding an instance of L2ToL1MessagePasser to it.
func DeployL2ToL1MessagePasser(auth *bind.TransactOpts, backend bind.ContractBackend) (libcommon.Address, types.Transaction, *L2ToL1MessagePasser, error) {
	parsed, err := abi.JSON(strings.NewReader(L2ToL1MessagePasserABI))
	if err != nil {
		return libcommon.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, libcommon.FromHex(L2ToL1MessagePasserBin), backend)
	if err != nil {
		return libcommon.Address{}, nil, nil, err
	}
	return address, tx, &L2ToL1MessagePasser{L2ToL1MessagePasserCaller: L2ToL1MessagePasserCaller{contract: contract}, L2ToL1MessagePasserTransactor: L2ToL1MessagePasserTransactor{contract: contract}, L2ToL1MessagePasserFilterer: L2ToL1MessagePasserFilterer{contract: contract}}, nil
}

// L2ToL1MessagePasser is an auto generated Go binding around an Ethereum contract.
type L2ToL1MessagePasser struct {
	L2ToL1MessagePasserCaller     // Read-only binding to the contract
	L2ToL1MessagePasserTransactor // Write-only binding to the contract
	L2ToL1MessagePasserFilterer   // Log filterer for contract events
}

// L2ToL1MessagePasserCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ToL1MessagePasserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ToL1MessagePasserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2ToL1MessagePasserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2ToL1MessagePasserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2ToL1MessagePasserSession struct {
	Contract     *L2ToL1MessagePasser // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// L2ToL1MessagePasserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2ToL1MessagePasserCallerSession struct {
	Contract *L2ToL1MessagePasserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// L2ToL1MessagePasserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2ToL1MessagePasserTransactorSession struct {
	Contract     *L2ToL1MessagePasserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// L2ToL1MessagePasserRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2ToL1MessagePasserRaw struct {
	Contract *L2ToL1MessagePasser // Generic contract binding to access the raw methods on
}

// L2ToL1MessagePasserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserCallerRaw struct {
	Contract *L2ToL1MessagePasserCaller // Generic read-only contract binding to access the raw methods on
}

// L2ToL1MessagePasserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2ToL1MessagePasserTransactorRaw struct {
	Contract *L2ToL1MessagePasserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2ToL1MessagePasser creates a new instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasser(address libcommon.Address, backend bind.ContractBackend) (*L2ToL1MessagePasser, error) {
	contract, err := bindL2ToL1MessagePasser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasser{L2ToL1MessagePasserCaller: L2ToL1MessagePasserCaller{contract: contract}, L2ToL1MessagePasserTransactor: L2ToL1MessagePasserTransactor{contract: contract}, L2ToL1MessagePasserFilterer: L2ToL1MessagePasserFilterer{contract: contract}}, nil
}

// NewL2ToL1MessagePasserCaller creates a new read-only instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasserCaller(address libcommon.Address, caller bind.ContractCaller) (*L2ToL1MessagePasserCaller, error) {
	contract, err := bindL2ToL1MessagePasser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserCaller{contract: contract}, nil
}

// NewL2ToL1MessagePasserTransactor creates a new write-only instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasserTransactor(address libcommon.Address, transactor bind.ContractTransactor) (*L2ToL1MessagePasserTransactor, error) {
	contract, err := bindL2ToL1MessagePasser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserTransactor{contract: contract}, nil
}

// NewL2ToL1MessagePasserFilterer creates a new log filterer instance of L2ToL1MessagePasser, bound to a specific deployed contract.
func NewL2ToL1MessagePasserFilterer(address libcommon.Address, filterer bind.ContractFilterer) (*L2ToL1MessagePasserFilterer, error) {
	contract, err := bindL2ToL1MessagePasser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserFilterer{contract: contract}, nil
}

// bindL2ToL1MessagePasser binds a generic wrapper to an already deployed contract.
func bindL2ToL1MessagePasser(address libcommon.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2ToL1MessagePasserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ToL1MessagePasser.Contract.L2ToL1MessagePasserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserRaw) Transfer(opts *bind.TransactOpts) (types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.L2ToL1MessagePasserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.L2ToL1MessagePasserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2ToL1MessagePasser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorRaw) Transfer(opts *bind.TransactOpts) (types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.contract.Transact(opts, method, params...)
}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) MESSAGEVERSION(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "MESSAGE_VERSION")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) MESSAGEVERSION() (uint16, error) {
	return _L2ToL1MessagePasser.Contract.MESSAGEVERSION(&_L2ToL1MessagePasser.CallOpts)
}

// MESSAGEVERSION is a free data retrieval call binding the contract method 0x3f827a5a.
//
// Solidity: function MESSAGE_VERSION() view returns(uint16)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) MESSAGEVERSION() (uint16, error) {
	return _L2ToL1MessagePasser.Contract.MESSAGEVERSION(&_L2ToL1MessagePasser.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) MessageNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "messageNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) MessageNonce() (*big.Int, error) {
	return _L2ToL1MessagePasser.Contract.MessageNonce(&_L2ToL1MessagePasser.CallOpts)
}

// MessageNonce is a free data retrieval call binding the contract method 0xecc70428.
//
// Solidity: function messageNonce() view returns(uint256)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) MessageNonce() (*big.Int, error) {
	return _L2ToL1MessagePasser.Contract.MessageNonce(&_L2ToL1MessagePasser.CallOpts)
}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) SentMessages(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "sentMessages", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) SentMessages(arg0 [32]byte) (bool, error) {
	return _L2ToL1MessagePasser.Contract.SentMessages(&_L2ToL1MessagePasser.CallOpts, arg0)
}

// SentMessages is a free data retrieval call binding the contract method 0x82e3702d.
//
// Solidity: function sentMessages(bytes32 ) view returns(bool)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) SentMessages(arg0 [32]byte) (bool, error) {
	return _L2ToL1MessagePasser.Contract.SentMessages(&_L2ToL1MessagePasser.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2ToL1MessagePasser.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) Version() (string, error) {
	return _L2ToL1MessagePasser.Contract.Version(&_L2ToL1MessagePasser.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserCallerSession) Version() (string, error) {
	return _L2ToL1MessagePasser.Contract.Version(&_L2ToL1MessagePasser.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactor) Burn(opts *bind.TransactOpts) (types.Transaction, error) {
	return _L2ToL1MessagePasser.contract.Transact(opts, "burn")
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) Burn() (types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.Burn(&_L2ToL1MessagePasser.TransactOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x44df8e70.
//
// Solidity: function burn() returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession) Burn() (types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.Burn(&_L2ToL1MessagePasser.TransactOpts)
}

// InitiateETHERC20Withdrawal is a paid mutator transaction binding the contract method 0xdfd44e4c.
//
// Solidity: function initiateETHERC20Withdrawal(address _target, uint256 _value, uint256 _gasLimit, bytes _data) returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactor) InitiateETHERC20Withdrawal(opts *bind.TransactOpts, _target libcommon.Address, _value *big.Int, _gasLimit *big.Int, _data []byte) (types.Transaction, error) {
	return _L2ToL1MessagePasser.contract.Transact(opts, "initiateETHERC20Withdrawal", _target, _value, _gasLimit, _data)
}

// InitiateETHERC20Withdrawal is a paid mutator transaction binding the contract method 0xdfd44e4c.
//
// Solidity: function initiateETHERC20Withdrawal(address _target, uint256 _value, uint256 _gasLimit, bytes _data) returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) InitiateETHERC20Withdrawal(_target libcommon.Address, _value *big.Int, _gasLimit *big.Int, _data []byte) (types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.InitiateETHERC20Withdrawal(&_L2ToL1MessagePasser.TransactOpts, _target, _value, _gasLimit, _data)
}

// InitiateETHERC20Withdrawal is a paid mutator transaction binding the contract method 0xdfd44e4c.
//
// Solidity: function initiateETHERC20Withdrawal(address _target, uint256 _value, uint256 _gasLimit, bytes _data) returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession) InitiateETHERC20Withdrawal(_target libcommon.Address, _value *big.Int, _gasLimit *big.Int, _data []byte) (types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.InitiateETHERC20Withdrawal(&_L2ToL1MessagePasser.TransactOpts, _target, _value, _gasLimit, _data)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xc2b3e5ac.
//
// Solidity: function initiateWithdrawal(address _target, uint256 _gasLimit, bytes _data) payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactor) InitiateWithdrawal(opts *bind.TransactOpts, _target libcommon.Address, _gasLimit *big.Int, _data []byte) (types.Transaction, error) {
	return _L2ToL1MessagePasser.contract.Transact(opts, "initiateWithdrawal", _target, _gasLimit, _data)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xc2b3e5ac.
//
// Solidity: function initiateWithdrawal(address _target, uint256 _gasLimit, bytes _data) payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) InitiateWithdrawal(_target libcommon.Address, _gasLimit *big.Int, _data []byte) (types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.InitiateWithdrawal(&_L2ToL1MessagePasser.TransactOpts, _target, _gasLimit, _data)
}

// InitiateWithdrawal is a paid mutator transaction binding the contract method 0xc2b3e5ac.
//
// Solidity: function initiateWithdrawal(address _target, uint256 _gasLimit, bytes _data) payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession) InitiateWithdrawal(_target libcommon.Address, _gasLimit *big.Int, _data []byte) (types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.InitiateWithdrawal(&_L2ToL1MessagePasser.TransactOpts, _target, _gasLimit, _data)
}

// InitiateETHERC20WithdrawalParams is an auto generated read-only Go binding of transcaction calldata params
type InitiateETHERC20WithdrawalParams struct {
	Param__target   libcommon.Address
	Param__value    *big.Int
	Param__gasLimit *big.Int
	Param__data     []byte
}

// Parse InitiateETHERC20Withdrawal method from calldata of a transaction
//
// Solidity: function initiateETHERC20Withdrawal(address _target, uint256 _value, uint256 _gasLimit, bytes _data) returns()
func ParseInitiateETHERC20Withdrawal(calldata []byte) (*InitiateETHERC20WithdrawalParams, error) {
	if len(calldata) <= 4 {
		return nil, fmt.Errorf("invalid calldata input")
	}

	_abi, err := abi.JSON(strings.NewReader(L2ToL1MessagePasserABI))
	if err != nil {
		return nil, fmt.Errorf("failed to get abi of registry metadata: %w", err)
	}

	out, err := _abi.Methods["initiateETHERC20Withdrawal"].Inputs.Unpack(calldata[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack initiateETHERC20Withdrawal params data: %w", err)
	}

	var paramsResult = new(InitiateETHERC20WithdrawalParams)
	value := reflect.ValueOf(paramsResult).Elem()

	if value.NumField() != len(out) {
		return nil, fmt.Errorf("failed to match calldata with param field number")
	}

	out0 := *abi.ConvertType(out[0], new(libcommon.Address)).(*libcommon.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return &InitiateETHERC20WithdrawalParams{
		Param__target: out0, Param__value: out1, Param__gasLimit: out2, Param__data: out3,
	}, nil
}

// InitiateWithdrawalParams is an auto generated read-only Go binding of transcaction calldata params
type InitiateWithdrawalParams struct {
	Param__target   libcommon.Address
	Param__gasLimit *big.Int
	Param__data     []byte
}

// Parse InitiateWithdrawal method from calldata of a transaction
//
// Solidity: function initiateWithdrawal(address _target, uint256 _gasLimit, bytes _data) payable returns()
func ParseInitiateWithdrawal(calldata []byte) (*InitiateWithdrawalParams, error) {
	if len(calldata) <= 4 {
		return nil, fmt.Errorf("invalid calldata input")
	}

	_abi, err := abi.JSON(strings.NewReader(L2ToL1MessagePasserABI))
	if err != nil {
		return nil, fmt.Errorf("failed to get abi of registry metadata: %w", err)
	}

	out, err := _abi.Methods["initiateWithdrawal"].Inputs.Unpack(calldata[4:])
	if err != nil {
		return nil, fmt.Errorf("failed to unpack initiateWithdrawal params data: %w", err)
	}

	var paramsResult = new(InitiateWithdrawalParams)
	value := reflect.ValueOf(paramsResult).Elem()

	if value.NumField() != len(out) {
		return nil, fmt.Errorf("failed to match calldata with param field number")
	}

	out0 := *abi.ConvertType(out[0], new(libcommon.Address)).(*libcommon.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return &InitiateWithdrawalParams{
		Param__target: out0, Param__gasLimit: out1, Param__data: out2,
	}, nil
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactor) Receive(opts *bind.TransactOpts) (types.Transaction, error) {
	return _L2ToL1MessagePasser.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserSession) Receive() (types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.Receive(&_L2ToL1MessagePasser.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2ToL1MessagePasser *L2ToL1MessagePasserTransactorSession) Receive() (types.Transaction, error) {
	return _L2ToL1MessagePasser.Contract.Receive(&_L2ToL1MessagePasser.TransactOpts)
}

// L2ToL1MessagePasserMessagePassedIterator is returned from FilterMessagePassed and is used to iterate over the raw logs and unpacked data for MessagePassed events raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserMessagePassedIterator struct {
	Event *L2ToL1MessagePasserMessagePassed // Event containing the contract specifics and raw log

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
func (it *L2ToL1MessagePasserMessagePassedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ToL1MessagePasserMessagePassed)
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
		it.Event = new(L2ToL1MessagePasserMessagePassed)
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
func (it *L2ToL1MessagePasserMessagePassedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ToL1MessagePasserMessagePassedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ToL1MessagePasserMessagePassed represents a MessagePassed event raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserMessagePassed struct {
	Nonce          *big.Int
	Sender         libcommon.Address
	Target         libcommon.Address
	Value          *big.Int
	GasLimit       *big.Int
	Data           []byte
	WithdrawalHash [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMessagePassed is a free log retrieval operation binding the contract event 0x02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a2744955054.
//
// Solidity: event MessagePassed(uint256 indexed nonce, address indexed sender, address indexed target, uint256 value, uint256 gasLimit, bytes data, bytes32 withdrawalHash)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) FilterMessagePassed(opts *bind.FilterOpts, nonce []*big.Int, sender []libcommon.Address, target []libcommon.Address) (*L2ToL1MessagePasserMessagePassedIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L2ToL1MessagePasser.contract.FilterLogs(opts, "MessagePassed", nonceRule, senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserMessagePassedIterator{contract: _L2ToL1MessagePasser.contract, event: "MessagePassed", logs: logs, sub: sub}, nil
}

// WatchMessagePassed is a free log subscription operation binding the contract event 0x02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a2744955054.
//
// Solidity: event MessagePassed(uint256 indexed nonce, address indexed sender, address indexed target, uint256 value, uint256 gasLimit, bytes data, bytes32 withdrawalHash)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) WatchMessagePassed(opts *bind.WatchOpts, sink chan<- *L2ToL1MessagePasserMessagePassed, nonce []*big.Int, sender []libcommon.Address, target []libcommon.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}

	logs, sub, err := _L2ToL1MessagePasser.contract.WatchLogs(opts, "MessagePassed", nonceRule, senderRule, targetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ToL1MessagePasserMessagePassed)
				if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "MessagePassed", log); err != nil {
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

// ParseMessagePassed is a log parse operation binding the contract event 0x02a52367d10742d8032712c1bb8e0144ff1ec5ffda1ed7d70bb05a2744955054.
//
// Solidity: event MessagePassed(uint256 indexed nonce, address indexed sender, address indexed target, uint256 value, uint256 gasLimit, bytes data, bytes32 withdrawalHash)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) ParseMessagePassed(log types.Log) (*L2ToL1MessagePasserMessagePassed, error) {
	event := new(L2ToL1MessagePasserMessagePassed)
	if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "MessagePassed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2ToL1MessagePasserWithdrawerBalanceBurntIterator is returned from FilterWithdrawerBalanceBurnt and is used to iterate over the raw logs and unpacked data for WithdrawerBalanceBurnt events raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserWithdrawerBalanceBurntIterator struct {
	Event *L2ToL1MessagePasserWithdrawerBalanceBurnt // Event containing the contract specifics and raw log

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
func (it *L2ToL1MessagePasserWithdrawerBalanceBurntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2ToL1MessagePasserWithdrawerBalanceBurnt)
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
		it.Event = new(L2ToL1MessagePasserWithdrawerBalanceBurnt)
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
func (it *L2ToL1MessagePasserWithdrawerBalanceBurntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ToL1MessagePasserWithdrawerBalanceBurntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2ToL1MessagePasserWithdrawerBalanceBurnt represents a WithdrawerBalanceBurnt event raised by the L2ToL1MessagePasser contract.
type L2ToL1MessagePasserWithdrawerBalanceBurnt struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawerBalanceBurnt is a free log retrieval operation binding the contract event 0x7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f.
//
// Solidity: event WithdrawerBalanceBurnt(uint256 indexed amount)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) FilterWithdrawerBalanceBurnt(opts *bind.FilterOpts, amount []*big.Int) (*L2ToL1MessagePasserWithdrawerBalanceBurntIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _L2ToL1MessagePasser.contract.FilterLogs(opts, "WithdrawerBalanceBurnt", amountRule)
	if err != nil {
		return nil, err
	}
	return &L2ToL1MessagePasserWithdrawerBalanceBurntIterator{contract: _L2ToL1MessagePasser.contract, event: "WithdrawerBalanceBurnt", logs: logs, sub: sub}, nil
}

// WatchWithdrawerBalanceBurnt is a free log subscription operation binding the contract event 0x7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f.
//
// Solidity: event WithdrawerBalanceBurnt(uint256 indexed amount)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) WatchWithdrawerBalanceBurnt(opts *bind.WatchOpts, sink chan<- *L2ToL1MessagePasserWithdrawerBalanceBurnt, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _L2ToL1MessagePasser.contract.WatchLogs(opts, "WithdrawerBalanceBurnt", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2ToL1MessagePasserWithdrawerBalanceBurnt)
				if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "WithdrawerBalanceBurnt", log); err != nil {
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

// ParseWithdrawerBalanceBurnt is a log parse operation binding the contract event 0x7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f.
//
// Solidity: event WithdrawerBalanceBurnt(uint256 indexed amount)
func (_L2ToL1MessagePasser *L2ToL1MessagePasserFilterer) ParseWithdrawerBalanceBurnt(log types.Log) (*L2ToL1MessagePasserWithdrawerBalanceBurnt, error) {
	event := new(L2ToL1MessagePasserWithdrawerBalanceBurnt)
	if err := _L2ToL1MessagePasser.contract.UnpackLog(event, "WithdrawerBalanceBurnt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
