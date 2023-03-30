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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"random\",\"type\":\"uint256\"}],\"name\":\"OffchainRandom\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rType\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"GetResponse\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetSimpleRandom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NextSimpleRandom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cacheKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"PutResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"PutSimpleRandom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b031916301790556108d5806100326000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c8063975009711161005057806397500971146100ab578063a101dc21146100c0578063dfc98ae8146100c957600080fd5b80637d93616c1461006c5780637dbf7c1014610095575b600080fd5b61007f61007a366004610567565b6100dc565b60405161008c919061062d565b60405180910390f35b61009d6102d7565b60405190815260200161008c565b6100be6100b936600461067e565b600155565b005b61009d60015481565b6100be6100d7366004610697565b61038f565b60608363ffffffff1660020361016357600082511161015c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f547572696e673a476574526573706f6e73653a6e6f207061796c6f616400000060448201526064015b60405180910390fd5b50806102d0565b600083836040516020016101789291906106de565b60405160208183030381529060405280519060200120905060006002600083815260200190815260200160002080546101b09061070d565b905011610219576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f4d697373696e6720636163686520656e747279000000000000000000000000006044820152606401610153565b600081815260026020526040812080546102329061070d565b80601f016020809104026020016040519081016040528092919081815260200182805461025e9061070d565b80156102ab5780601f10610280576101008083540402835291602001916102ab565b820191906000526020600020905b81548152906001019060200180831161028e57829003601f168201915b50505060008581526002602052604081209394506102cc939250905061042e565b9150505b9392505050565b6000600154600003610345576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f52616e646f6d206e756d626572206e6f742067656e65726174656400000000006044820152606401610153565b60018054600091829055604080519283526020830182905290917f450d62889c3a6e19c9586840ce9c21040b90d81950fe31f2ba982090adaf53e8910160405180910390a1919050565b600082815260026020526040902080546103a89061070d565b159050610411576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4445425547202d20416c726561647920657869737473000000000000000000006044820152606401610153565b600082815260026020526040902061042982826107ae565b505050565b50805461043a9061070d565b6000825580601f1061044a575050565b601f016020900490600052602060002090810190610468919061046b565b50565b5b80821115610480576000815560010161046c565b5090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600067ffffffffffffffff808411156104ce576104ce610484565b604051601f85017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561051457610514610484565b8160405280935085815286868601111561052d57600080fd5b858560208301376000602087830101525050509392505050565b600082601f83011261055857600080fd5b6102d0838335602085016104b3565b60008060006060848603121561057c57600080fd5b833563ffffffff8116811461059057600080fd5b9250602084013567ffffffffffffffff808211156105ad57600080fd5b818601915086601f8301126105c157600080fd5b6105d0878335602085016104b3565b935060408601359150808211156105e657600080fd5b506105f386828701610547565b9150509250925092565b60005b83811015610618578181015183820152602001610600565b83811115610627576000848401525b50505050565b602081526000825180602084015261064c8160408501602087016105fd565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b60006020828403121561069057600080fd5b5035919050565b600080604083850312156106aa57600080fd5b82359150602083013567ffffffffffffffff8111156106c857600080fd5b6106d485828601610547565b9150509250929050565b600083516106f08184602088016105fd565b8351908301906107048183602088016105fd565b01949350505050565b600181811c9082168061072157607f821691505b60208210810361075a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f82111561042957600081815260208120601f850160051c810160208610156107875750805b601f850160051c820191505b818110156107a657828155600101610793565b505050505050565b815167ffffffffffffffff8111156107c8576107c8610484565b6107dc816107d6845461070d565b84610760565b602080601f83116001811461082f57600084156107f95750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556107a6565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b8281101561087c5788860151825594840194600190910190840161085d565b50858210156108b857878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b0190555056fea164736f6c634300080f000a",
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

// GetResponse is a paid mutator transaction binding the contract method 0x7d93616c.
//
// Solidity: function GetResponse(uint32 rType, string _url, bytes _payload) returns(bytes)
func (_BobaHCHelper *BobaHCHelperTransactor) GetResponse(opts *bind.TransactOpts, rType uint32, _url string, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "GetResponse", rType, _url, _payload)
}

// GetResponse is a paid mutator transaction binding the contract method 0x7d93616c.
//
// Solidity: function GetResponse(uint32 rType, string _url, bytes _payload) returns(bytes)
func (_BobaHCHelper *BobaHCHelperSession) GetResponse(rType uint32, _url string, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.GetResponse(&_BobaHCHelper.TransactOpts, rType, _url, _payload)
}

// GetResponse is a paid mutator transaction binding the contract method 0x7d93616c.
//
// Solidity: function GetResponse(uint32 rType, string _url, bytes _payload) returns(bytes)
func (_BobaHCHelper *BobaHCHelperTransactorSession) GetResponse(rType uint32, _url string, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.GetResponse(&_BobaHCHelper.TransactOpts, rType, _url, _payload)
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
