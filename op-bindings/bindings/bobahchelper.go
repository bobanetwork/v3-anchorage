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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"DBG\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"Success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"EK\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"URL\",\"type\":\"string\"}],\"name\":\"EndpointRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"random\",\"type\":\"uint256\"}],\"name\":\"OffchainRandom\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"EK\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"CallOffchain\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Endpoints\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"URL\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cacheKey\",\"type\":\"bytes32\"}],\"name\":\"ExpiredResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rType\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"GetBounce\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rType\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"GetResponse\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetSimpleRandom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NextSimpleRandom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"a\",\"type\":\"uint32\"}],\"name\":\"Ping\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cacheKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"PutResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"PutSimpleRandom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"credits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"auth\",\"type\":\"bytes32\"}],\"name\":\"RegisterEndpoint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"EK\",\"type\":\"bytes32\"}],\"name\":\"UnregisterEndpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040527342000000000000000000000000000000000000fd608052606460035534801561002d57600080fd5b5060805161141a6100576000396000818161035b015281816108b401526109e3015261141a6000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c8063839f845711610081578063997e31cc1161005b578063997e31cc146101c4578063a101dc21146101e5578063dfc98ae8146101ee57600080fd5b8063839f8457146101765780639750097114610189578063981fce181461019c57600080fd5b80635b0dfe6f116100b25780635b0dfe6f1461012a5780637d93616c1461014d5780637dbf7c101461016057600080fd5b80631f999764146100d957806322a522a4146100ee578063279ecb9e14610101575b600080fd5b6100ec6100e7366004610c9d565b610201565b005b6100ec6100fc366004610c9d565b6102da565b61011461010f366004610de2565b6102f4565b6040516101219190610ed2565b60405180910390f35b61013d610138366004610ee5565b610356565b6040519015158152602001610121565b61011461015b366004610de2565b6105da565b6101686107e8565b604051908152602001610121565b610114610184366004610f33565b61089c565b6100ec610197366004610c9d565b600055565b6101af6101aa366004610f7a565b6109df565b60405163ffffffff9091168152602001610121565b6101d76101d2366004610c9d565b610ad0565b604051610121929190610f97565b61016860005481565b6100ec6101fc366004610f33565b610b8c565b60008181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff163314610293576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f6e6c79206f776e6572206d617920756e72656769737465720000000000000060448201526064015b60405180910390fd5b600081815260026020526040812080547fffffffffffffffffffffffff0000000000000000000000000000000000000000168155906102d56001830182610c4f565b505050565b60008181526001602052604081206102f191610c4f565b50565b60607f36eaa99bbc568a7704a0ec0cc2428796c95d28d08af99f4a3f0dc54c679520cd826040516103259190610ed2565b60405180910390a16040805160006020820181905291016040516020818303038152906040529150505b9392505050565b6000807f0000000000000000000000000000000000000000000000000000000000000000905060008560405160200161038f9190610fce565b6040516020818303038152906040528051906020012090506000600354861015610415576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f496e73756666696369656e7420726567697374726174696f6e20637265646974604482015260640161028a565b6040805160048082526024820183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f61f334d30000000000000000000000000000000000000000000000000000000017905291517f7d93616c000000000000000000000000000000000000000000000000000000008152909160009173ffffffffffffffffffffffffffffffffffffffff871691637d93616c916104c5916001918e91889101610fea565b6000604051808303816000875af11580156104e4573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405261052a9190810190611025565b90508051602014801561054257508051602082012087145b1561059057600084815260026020526040902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000163317815560010161058a8a8261112e565b50600192505b7f946870c115e26664fed95ef0844e3f2a093de34963e3a3a573bb793fa889a9488385338c6040516105c59493929190611248565b60405180910390a15090979650505050505050565b60608363ffffffff1660020361065c576000825111610655576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f547572696e673a476574526573706f6e73653a6e6f207061796c6f6164000000604482015260640161028a565b508061034f565b600083338460405160200161067393929190611285565b60405160208183030381529060405280519060200120905060006001600083815260200190815260200160002080546106ab90611093565b905011610714576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f4d697373696e6720636163686520656e74727900000000000000000000000000604482015260640161028a565b6000818152600160205260408120805461072d90611093565b80601f016020809104026020016040519081016040528092919081815260200182805461075990611093565b80156107a65780601f1061077b576101008083540402835291602001916107a6565b820191906000526020600020905b81548152906001019060200180831161078957829003601f168201915b505083519394506107c792506107c29150839050611388611313565b610c26565b60008381526001602052604081206107de91610c4f565b5095945050505050565b60008054600003610855576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f52616e646f6d206e756d626572206e6f742067656e6572617465640000000000604482015260640161028a565b60008054818055604080519283526020830182905290917f450d62889c3a6e19c9586840ce9c21040b90d81950fe31f2ba982090adaf53e8910160405180910390a1919050565b600082815260026020526040812060010180546060927f00000000000000000000000000000000000000000000000000000000000000009290916108df90611093565b80601f016020809104026020016040519081016040528092919081815260200182805461090b90611093565b80156109585780601f1061092d57610100808354040283529160200191610958565b820191906000526020600020905b81548152906001019060200180831161093b57829003601f168201915b5050505050905060008151116109ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f456e64706f696e74206e6f742072656769737465726564000000000000000000604482015260640161028a565b6109d6600182866105da565b95945050505050565b60007f000000000000000000000000000000000000000000000000000000000000000063ffffffff83168203610a185750600192915050565b73ffffffffffffffffffffffffffffffffffffffff811663981fce18610a3f600186611350565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815263ffffffff9190911660048201526024016020604051808303816000875af1158015610a9b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610abf9190611375565b61034f906002611392565b50919050565b6002602052600090815260409020805460018201805473ffffffffffffffffffffffffffffffffffffffff9092169291610b0990611093565b80601f0160208091040260200160405190810160405280929190818152602001828054610b3590611093565b8015610b825780601f10610b5757610100808354040283529160200191610b82565b820191906000526020600020905b815481529060010190602001808311610b6557829003601f168201915b5050505050905082565b60008281526001602052604090208054610ba590611093565b159050610c0e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4445425547202d20416c72656164792065786973747300000000000000000000604482015260640161028a565b60008281526001602052604090206102d5828261112e565b6000805a90505b825a610c3990836113be565b10156102d557610c48826113d5565b9150610c2d565b508054610c5b90611093565b6000825580601f10610c6b575050565b601f0160209004906000526020600020908101906102f191905b80821115610c995760008155600101610c85565b5090565b600060208284031215610caf57600080fd5b5035919050565b63ffffffff811681146102f157600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610d3e57610d3e610cc8565b604052919050565b600067ffffffffffffffff821115610d6057610d60610cc8565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112610d9d57600080fd5b8135610db0610dab82610d46565b610cf7565b818152846020838601011115610dc557600080fd5b816020850160208301376000918101602001919091529392505050565b600080600060608486031215610df757600080fd5b8335610e0281610cb6565b9250602084013567ffffffffffffffff80821115610e1f57600080fd5b610e2b87838801610d8c565b93506040860135915080821115610e4157600080fd5b50610e4e86828701610d8c565b9150509250925092565b60005b83811015610e73578181015183820152602001610e5b565b83811115610e82576000848401525b50505050565b60008151808452610ea0816020860160208601610e58565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061034f6020830184610e88565b600080600060608486031215610efa57600080fd5b833567ffffffffffffffff811115610f1157600080fd5b610f1d86828701610d8c565b9660208601359650604090950135949350505050565b60008060408385031215610f4657600080fd5b82359150602083013567ffffffffffffffff811115610f6457600080fd5b610f7085828601610d8c565b9150509250929050565b600060208284031215610f8c57600080fd5b813561034f81610cb6565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201526000610fc66040830184610e88565b949350505050565b60008251610fe0818460208701610e58565b9190910192915050565b63ffffffff841681526060602082015260006110096060830185610e88565b828103604084015261101b8185610e88565b9695505050505050565b60006020828403121561103757600080fd5b815167ffffffffffffffff81111561104e57600080fd5b8201601f8101841361105f57600080fd5b805161106d610dab82610d46565b81815285602083850101111561108257600080fd5b6109d6826020830160208601610e58565b600181811c908216806110a757607f821691505b602082108103610aca577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b601f8211156102d557600081815260208120601f850160051c810160208610156111075750805b601f850160051c820191505b8181101561112657828155600101611113565b505050505050565b815167ffffffffffffffff81111561114857611148610cc8565b61115c816111568454611093565b846110e0565b602080601f8311600181146111af57600084156111795750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611126565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b828110156111fc578886015182559484019460019091019084016111dd565b508582101561123857878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b841515815283602082015273ffffffffffffffffffffffffffffffffffffffff8316604082015260806060820152600061101b6080830184610e88565b60008451611297818460208901610e58565b80830190507fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008560601b16815283516112d7816014840160208801610e58565b0160140195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561134b5761134b6112e4565b500290565b600063ffffffff8381169083168181101561136d5761136d6112e4565b039392505050565b60006020828403121561138757600080fd5b815161034f81610cb6565b600063ffffffff808316818516818304811182151516156113b5576113b56112e4565b02949350505050565b6000828210156113d0576113d06112e4565b500390565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611406576114066112e4565b506001019056fea164736f6c634300080f000a",
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

// Endpoints is a free data retrieval call binding the contract method 0x997e31cc.
//
// Solidity: function Endpoints(bytes32 ) view returns(address Owner, string URL)
func (_BobaHCHelper *BobaHCHelperCaller) Endpoints(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Owner common.Address
	URL   string
}, error) {
	var out []interface{}
	err := _BobaHCHelper.contract.Call(opts, &out, "Endpoints", arg0)

	outstruct := new(struct {
		Owner common.Address
		URL   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.URL = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Endpoints is a free data retrieval call binding the contract method 0x997e31cc.
//
// Solidity: function Endpoints(bytes32 ) view returns(address Owner, string URL)
func (_BobaHCHelper *BobaHCHelperSession) Endpoints(arg0 [32]byte) (struct {
	Owner common.Address
	URL   string
}, error) {
	return _BobaHCHelper.Contract.Endpoints(&_BobaHCHelper.CallOpts, arg0)
}

// Endpoints is a free data retrieval call binding the contract method 0x997e31cc.
//
// Solidity: function Endpoints(bytes32 ) view returns(address Owner, string URL)
func (_BobaHCHelper *BobaHCHelperCallerSession) Endpoints(arg0 [32]byte) (struct {
	Owner common.Address
	URL   string
}, error) {
	return _BobaHCHelper.Contract.Endpoints(&_BobaHCHelper.CallOpts, arg0)
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

// CallOffchain is a paid mutator transaction binding the contract method 0x839f8457.
//
// Solidity: function CallOffchain(bytes32 EK, bytes payload) returns(bytes)
func (_BobaHCHelper *BobaHCHelperTransactor) CallOffchain(opts *bind.TransactOpts, EK [32]byte, payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "CallOffchain", EK, payload)
}

// CallOffchain is a paid mutator transaction binding the contract method 0x839f8457.
//
// Solidity: function CallOffchain(bytes32 EK, bytes payload) returns(bytes)
func (_BobaHCHelper *BobaHCHelperSession) CallOffchain(EK [32]byte, payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.CallOffchain(&_BobaHCHelper.TransactOpts, EK, payload)
}

// CallOffchain is a paid mutator transaction binding the contract method 0x839f8457.
//
// Solidity: function CallOffchain(bytes32 EK, bytes payload) returns(bytes)
func (_BobaHCHelper *BobaHCHelperTransactorSession) CallOffchain(EK [32]byte, payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.CallOffchain(&_BobaHCHelper.TransactOpts, EK, payload)
}

// ExpiredResponse is a paid mutator transaction binding the contract method 0x22a522a4.
//
// Solidity: function ExpiredResponse(bytes32 cacheKey) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) ExpiredResponse(opts *bind.TransactOpts, cacheKey [32]byte) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "ExpiredResponse", cacheKey)
}

// ExpiredResponse is a paid mutator transaction binding the contract method 0x22a522a4.
//
// Solidity: function ExpiredResponse(bytes32 cacheKey) returns()
func (_BobaHCHelper *BobaHCHelperSession) ExpiredResponse(cacheKey [32]byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.ExpiredResponse(&_BobaHCHelper.TransactOpts, cacheKey)
}

// ExpiredResponse is a paid mutator transaction binding the contract method 0x22a522a4.
//
// Solidity: function ExpiredResponse(bytes32 cacheKey) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) ExpiredResponse(cacheKey [32]byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.ExpiredResponse(&_BobaHCHelper.TransactOpts, cacheKey)
}

// GetBounce is a paid mutator transaction binding the contract method 0x279ecb9e.
//
// Solidity: function GetBounce(uint32 rType, string _url, bytes _payload) returns(bytes)
func (_BobaHCHelper *BobaHCHelperTransactor) GetBounce(opts *bind.TransactOpts, rType uint32, _url string, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "GetBounce", rType, _url, _payload)
}

// GetBounce is a paid mutator transaction binding the contract method 0x279ecb9e.
//
// Solidity: function GetBounce(uint32 rType, string _url, bytes _payload) returns(bytes)
func (_BobaHCHelper *BobaHCHelperSession) GetBounce(rType uint32, _url string, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.GetBounce(&_BobaHCHelper.TransactOpts, rType, _url, _payload)
}

// GetBounce is a paid mutator transaction binding the contract method 0x279ecb9e.
//
// Solidity: function GetBounce(uint32 rType, string _url, bytes _payload) returns(bytes)
func (_BobaHCHelper *BobaHCHelperTransactorSession) GetBounce(rType uint32, _url string, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.GetBounce(&_BobaHCHelper.TransactOpts, rType, _url, _payload)
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

// Ping is a paid mutator transaction binding the contract method 0x981fce18.
//
// Solidity: function Ping(uint32 a) returns(uint32)
func (_BobaHCHelper *BobaHCHelperTransactor) Ping(opts *bind.TransactOpts, a uint32) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "Ping", a)
}

// Ping is a paid mutator transaction binding the contract method 0x981fce18.
//
// Solidity: function Ping(uint32 a) returns(uint32)
func (_BobaHCHelper *BobaHCHelperSession) Ping(a uint32) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.Ping(&_BobaHCHelper.TransactOpts, a)
}

// Ping is a paid mutator transaction binding the contract method 0x981fce18.
//
// Solidity: function Ping(uint32 a) returns(uint32)
func (_BobaHCHelper *BobaHCHelperTransactorSession) Ping(a uint32) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.Ping(&_BobaHCHelper.TransactOpts, a)
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

// RegisterEndpoint is a paid mutator transaction binding the contract method 0x5b0dfe6f.
//
// Solidity: function RegisterEndpoint(string url, uint256 credits, bytes32 auth) returns(bool)
func (_BobaHCHelper *BobaHCHelperTransactor) RegisterEndpoint(opts *bind.TransactOpts, url string, credits *big.Int, auth [32]byte) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "RegisterEndpoint", url, credits, auth)
}

// RegisterEndpoint is a paid mutator transaction binding the contract method 0x5b0dfe6f.
//
// Solidity: function RegisterEndpoint(string url, uint256 credits, bytes32 auth) returns(bool)
func (_BobaHCHelper *BobaHCHelperSession) RegisterEndpoint(url string, credits *big.Int, auth [32]byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.RegisterEndpoint(&_BobaHCHelper.TransactOpts, url, credits, auth)
}

// RegisterEndpoint is a paid mutator transaction binding the contract method 0x5b0dfe6f.
//
// Solidity: function RegisterEndpoint(string url, uint256 credits, bytes32 auth) returns(bool)
func (_BobaHCHelper *BobaHCHelperTransactorSession) RegisterEndpoint(url string, credits *big.Int, auth [32]byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.RegisterEndpoint(&_BobaHCHelper.TransactOpts, url, credits, auth)
}

// UnregisterEndpoint is a paid mutator transaction binding the contract method 0x1f999764.
//
// Solidity: function UnregisterEndpoint(bytes32 EK) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) UnregisterEndpoint(opts *bind.TransactOpts, EK [32]byte) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "UnregisterEndpoint", EK)
}

// UnregisterEndpoint is a paid mutator transaction binding the contract method 0x1f999764.
//
// Solidity: function UnregisterEndpoint(bytes32 EK) returns()
func (_BobaHCHelper *BobaHCHelperSession) UnregisterEndpoint(EK [32]byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.UnregisterEndpoint(&_BobaHCHelper.TransactOpts, EK)
}

// UnregisterEndpoint is a paid mutator transaction binding the contract method 0x1f999764.
//
// Solidity: function UnregisterEndpoint(bytes32 EK) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) UnregisterEndpoint(EK [32]byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.UnregisterEndpoint(&_BobaHCHelper.TransactOpts, EK)
}

// BobaHCHelperDBGIterator is returned from FilterDBG and is used to iterate over the raw logs and unpacked data for DBG events raised by the BobaHCHelper contract.
type BobaHCHelperDBGIterator struct {
	Event *BobaHCHelperDBG // Event containing the contract specifics and raw log

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
func (it *BobaHCHelperDBGIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BobaHCHelperDBG)
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
		it.Event = new(BobaHCHelperDBG)
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
func (it *BobaHCHelperDBGIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BobaHCHelperDBGIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BobaHCHelperDBG represents a DBG event raised by the BobaHCHelper contract.
type BobaHCHelperDBG struct {
	B   []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDBG is a free log retrieval operation binding the contract event 0x36eaa99bbc568a7704a0ec0cc2428796c95d28d08af99f4a3f0dc54c679520cd.
//
// Solidity: event DBG(bytes b)
func (_BobaHCHelper *BobaHCHelperFilterer) FilterDBG(opts *bind.FilterOpts) (*BobaHCHelperDBGIterator, error) {

	logs, sub, err := _BobaHCHelper.contract.FilterLogs(opts, "DBG")
	if err != nil {
		return nil, err
	}
	return &BobaHCHelperDBGIterator{contract: _BobaHCHelper.contract, event: "DBG", logs: logs, sub: sub}, nil
}

// WatchDBG is a free log subscription operation binding the contract event 0x36eaa99bbc568a7704a0ec0cc2428796c95d28d08af99f4a3f0dc54c679520cd.
//
// Solidity: event DBG(bytes b)
func (_BobaHCHelper *BobaHCHelperFilterer) WatchDBG(opts *bind.WatchOpts, sink chan<- *BobaHCHelperDBG) (event.Subscription, error) {

	logs, sub, err := _BobaHCHelper.contract.WatchLogs(opts, "DBG")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BobaHCHelperDBG)
				if err := _BobaHCHelper.contract.UnpackLog(event, "DBG", log); err != nil {
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

// ParseDBG is a log parse operation binding the contract event 0x36eaa99bbc568a7704a0ec0cc2428796c95d28d08af99f4a3f0dc54c679520cd.
//
// Solidity: event DBG(bytes b)
func (_BobaHCHelper *BobaHCHelperFilterer) ParseDBG(log types.Log) (*BobaHCHelperDBG, error) {
	event := new(BobaHCHelperDBG)
	if err := _BobaHCHelper.contract.UnpackLog(event, "DBG", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BobaHCHelperEndpointRegisteredIterator is returned from FilterEndpointRegistered and is used to iterate over the raw logs and unpacked data for EndpointRegistered events raised by the BobaHCHelper contract.
type BobaHCHelperEndpointRegisteredIterator struct {
	Event *BobaHCHelperEndpointRegistered // Event containing the contract specifics and raw log

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
func (it *BobaHCHelperEndpointRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BobaHCHelperEndpointRegistered)
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
		it.Event = new(BobaHCHelperEndpointRegistered)
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
func (it *BobaHCHelperEndpointRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BobaHCHelperEndpointRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BobaHCHelperEndpointRegistered represents a EndpointRegistered event raised by the BobaHCHelper contract.
type BobaHCHelperEndpointRegistered struct {
	Success bool
	EK      [32]byte
	Owner   common.Address
	URL     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEndpointRegistered is a free log retrieval operation binding the contract event 0x946870c115e26664fed95ef0844e3f2a093de34963e3a3a573bb793fa889a948.
//
// Solidity: event EndpointRegistered(bool Success, bytes32 EK, address Owner, string URL)
func (_BobaHCHelper *BobaHCHelperFilterer) FilterEndpointRegistered(opts *bind.FilterOpts) (*BobaHCHelperEndpointRegisteredIterator, error) {

	logs, sub, err := _BobaHCHelper.contract.FilterLogs(opts, "EndpointRegistered")
	if err != nil {
		return nil, err
	}
	return &BobaHCHelperEndpointRegisteredIterator{contract: _BobaHCHelper.contract, event: "EndpointRegistered", logs: logs, sub: sub}, nil
}

// WatchEndpointRegistered is a free log subscription operation binding the contract event 0x946870c115e26664fed95ef0844e3f2a093de34963e3a3a573bb793fa889a948.
//
// Solidity: event EndpointRegistered(bool Success, bytes32 EK, address Owner, string URL)
func (_BobaHCHelper *BobaHCHelperFilterer) WatchEndpointRegistered(opts *bind.WatchOpts, sink chan<- *BobaHCHelperEndpointRegistered) (event.Subscription, error) {

	logs, sub, err := _BobaHCHelper.contract.WatchLogs(opts, "EndpointRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BobaHCHelperEndpointRegistered)
				if err := _BobaHCHelper.contract.UnpackLog(event, "EndpointRegistered", log); err != nil {
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

// ParseEndpointRegistered is a log parse operation binding the contract event 0x946870c115e26664fed95ef0844e3f2a093de34963e3a3a573bb793fa889a948.
//
// Solidity: event EndpointRegistered(bool Success, bytes32 EK, address Owner, string URL)
func (_BobaHCHelper *BobaHCHelperFilterer) ParseEndpointRegistered(log types.Log) (*BobaHCHelperEndpointRegistered, error) {
	event := new(BobaHCHelperEndpointRegistered)
	if err := _BobaHCHelper.contract.UnpackLog(event, "EndpointRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
