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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"DBG\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"Success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"EK\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"URL\",\"type\":\"string\"}],\"name\":\"EndpointRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"random\",\"type\":\"uint256\"}],\"name\":\"OffchainRandom\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_EK\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_callerAddress\",\"type\":\"address\"}],\"name\":\"AddPermittedCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"EK\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"CallOffchain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_EK\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_callerAddress\",\"type\":\"address\"}],\"name\":\"CheckPermittedCaller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Endpoints\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"URL\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cacheKey\",\"type\":\"bytes32\"}],\"name\":\"ExpiredResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_rType\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"GetResponse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cacheKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_code\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"PutResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"credits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"auth\",\"type\":\"bytes32\"}],\"name\":\"RegisterEndpoint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_EK\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_callerAddress\",\"type\":\"address\"}],\"name\":\"RemovePermittedCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SimpleRandom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"EK\",\"type\":\"bytes32\"}],\"name\":\"UnregisterEndpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c06040527342000000000000000000000000000000000000fd60805273deaddeaddeaddeaddeaddeaddeaddeaddead990160a052606460035534801561004557600080fd5b5060805160a051611ca66100806000396000818161036501526105ae01526000818161068201528181610aa601526111ed0152611ca66000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c806379d9c09011610081578063997e31cc1161005b578063997e31cc1461018b578063c3d8cae6146101ac578063c6e4b0d3146101bf57600080fd5b806379d9c09014610144578063839f8457146101575780638e5dc7651461017857600080fd5b80631f999764116100b25780631f999764146100f657806322a522a4146101095780635b0dfe6f1461011c57600080fd5b8063040db828146100ce57806311edaae0146100e3575b600080fd5b6100e16100dc36600461146a565b6101d5565b005b6100e16100f13660046115c4565b61034d565b6100e161010436600461161b565b6104c2565b6100e161011736600461161b565b610596565b61012f61012a366004611634565b61067d565b60405190151581526020015b60405180910390f35b6100e161015236600461146a565b610918565b61016a610165366004611682565b610a8e565b60405161013b92919061173f565b61016a610186366004611762565b610cfa565b61019e61019936600461161b565b611014565b60405161013b9291906117e7565b61012f6101ba36600461146a565b6110d0565b6101c7611197565b60405190815260200161013b565b60008281526002602052604090205473ffffffffffffffffffffffffffffffffffffffff16610265576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f74207265676973746572656400000000000060448201526064015b60405180910390fd5b60008281526002602052604090205473ffffffffffffffffffffffffffffffffffffffff1633146102f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4e6f742074686520456e64706f696e74206f776e657200000000000000000000604482015260640161025c565b600091825260026020818152604080852073ffffffffffffffffffffffffffffffffffffffff9094168552929091019052902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146103ec576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f496e76616c696420507574526573706f6e736528292063616c6c657200000000604482015260640161025c565b6000838152600160205260409020805461040590611816565b15905061046e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4445425547202d20416c72656164792065786973747300000000000000000000604482015260640161025c565b60008381526020818152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff8716179055600190915290206104bc82826118b7565b50505050565b60008181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff16331461054f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f6e6c79206f776e6572206d617920756e726567697374657200000000000000604482015260640161025c565b600081815260026020526040812080547fffffffffffffffffffffffff00000000000000000000000000000000000000001681559061059160018301826113f3565b505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610635576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f496e76616c696420507574526573706f6e736528292063616c6c657200000000604482015260640161025c565b60008181526020818152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001690556001909152812061067a916113f3565b50565b6000807f000000000000000000000000000000000000000000000000000000000000000090506000856040516020016106b691906119d1565b604051602081830303815290604052805190602001209050600060035486101561073c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f496e73756666696369656e7420726567697374726174696f6e20637265646974604482015260640161025c565b6040805160048082526024820183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f61f334d30000000000000000000000000000000000000000000000000000000017905291517f8e5dc765000000000000000000000000000000000000000000000000000000008152909160609173ffffffffffffffffffffffffffffffffffffffff871691638e5dc765916107ee9160019130918f918991016119ed565b6000604051808303816000875af115801561080d573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526108539190810190611a45565b9093509050828015610866575080516020145b801561087757508051602082012087145b156108c957600084815260026020526040902080547fffffffffffffffffffffffff000000000000000000000000000000000000000016331781556001016108bf8a826118b7565b50600192506108ce565b600092505b7f946870c115e26664fed95ef0844e3f2a093de34963e3a3a573bb793fa889a9488385338c6040516109039493929190611ad7565b60405180910390a15090979650505050505050565b60008281526002602052604090205473ffffffffffffffffffffffffffffffffffffffff166109a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f742072656769737465726564000000000000604482015260640161025c565b60008281526002602052604090205473ffffffffffffffffffffffffffffffffffffffff163314610a30576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4e6f742074686520456e64706f696e74206f776e657200000000000000000000604482015260640161025c565b600091825260026020818152604080852073ffffffffffffffffffffffffffffffffffffffff9094168552929091019052902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b600082815260026020526040812060010180546060917f000000000000000000000000000000000000000000000000000000000000000091849190610ad290611816565b80601f0160208091040260200160405190810160405280929190818152602001828054610afe90611816565b8015610b4b5780601f10610b2057610100808354040283529160200191610b4b565b820191906000526020600020905b815481529060010190602001808311610b2e57829003601f168201915b505050505090506000815111610bbd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f456e64706f696e74206e6f742072656769737465726564000000000000000000604482015260640161025c565b6000868152600260208181526040808420338552909201905290205460ff16610c42576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616c6c6572206973206e6f74207065726d6974746564000000000000000000604482015260640161025c565b600060608373ffffffffffffffffffffffffffffffffffffffff16638e5dc765600133868b6040518563ffffffff1660e01b8152600401610c8694939291906119ed565b6000604051808303816000875af1158015610ca5573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610ceb9190810190611a45565b90999098509650505050505050565b600060608181333014610d69576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f496e76616c696420476574526573706f6e736528292063616c6c657200000000604482015260640161025c565b600088888888604051602001610d829493929190611b1e565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181528151602092830120600081815280845282812054600190945291909120805491935063ffffffff9092169190610de590611816565b80601f0160208091040260200160405190810160405280929190818152602001828054610e1190611816565b8015610e5e5780601f10610e3357610100808354040283529160200191610e5e565b820191906000526020600020905b815481529060010190602001808311610e4157829003601f168201915b505050505092508063ffffffff16600003610efe576000835111610ede576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f476574526573706f6e73653a204d697373696e6720636163686520656e747279604482015260640161025c565b825160019450610ef8610ef382611388611bda565b6113ca565b50610fbb565b604080518082018252601381527f487962726964436f6d70757465204572726f720000000000000000000000000060208201529051610f41908290602401611c17565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f08c379a0000000000000000000000000000000000000000000000000000000001790529350505b6000828152600160205260408120610fd2916113f3565b50600090815260208190526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001690559097909650945050505050565b6002602052600090815260409020805460018201805473ffffffffffffffffffffffffffffffffffffffff909216929161104d90611816565b80601f016020809104026020016040519081016040528092919081815260200182805461107990611816565b80156110c65780601f1061109b576101008083540402835291602001916110c6565b820191906000526020600020905b8154815290600101906020018083116110a957829003601f168201915b5050505050905082565b60008281526002602052604081205473ffffffffffffffffffffffffffffffffffffffff1661115b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f742072656769737465726564000000000000604482015260640161025c565b50600082815260026020818152604080842073ffffffffffffffffffffffffffffffffffffffff86168552909201905290205460ff1692915050565b6040517f8e5dc7650000000000000000000000000000000000000000000000000000000081526201000160048201523360248201526080604482015260006084820181905260a0606483015260a48201819052907f000000000000000000000000000000000000000000000000000000000000000090829060609073ffffffffffffffffffffffffffffffffffffffff841690638e5dc7659060c4016000604051808303816000875af1158015611252573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526112989190810190611a45565b909250905081611304576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f4f7065726174696f6e206661696c656400000000000000000000000000000000604482015260640161025c565b805160201461136f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f42616420726573706f6e7365206c656e67746800000000000000000000000000604482015260640161025c565b6000818060200190518101906113859190611c31565b6040805160008152602081018390529192507f450d62889c3a6e19c9586840ce9c21040b90d81950fe31f2ba982090adaf53e8910160405180910390a1949350505050565b6000805a90505b825a6113dd9083611c4a565b1015610591576113ec82611c61565b91506113d1565b5080546113ff90611816565b6000825580601f1061140f575050565b601f01602090049060005260206000209081019061067a91905b8082111561143d5760008155600101611429565b5090565b803573ffffffffffffffffffffffffffffffffffffffff8116811461146557600080fd5b919050565b6000806040838503121561147d57600080fd5b8235915061148d60208401611441565b90509250929050565b803563ffffffff8116811461146557600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611520576115206114aa565b604052919050565b600067ffffffffffffffff821115611542576115426114aa565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f83011261157f57600080fd5b813561159261158d82611528565b6114d9565b8181528460208386010111156115a757600080fd5b816020850160208301376000918101602001919091529392505050565b6000806000606084860312156115d957600080fd5b833592506115e960208501611496565b9150604084013567ffffffffffffffff81111561160557600080fd5b6116118682870161156e565b9150509250925092565b60006020828403121561162d57600080fd5b5035919050565b60008060006060848603121561164957600080fd5b833567ffffffffffffffff81111561166057600080fd5b61166c8682870161156e565b9660208601359650604090950135949350505050565b6000806040838503121561169557600080fd5b82359150602083013567ffffffffffffffff8111156116b357600080fd5b6116bf8582860161156e565b9150509250929050565b60005b838110156116e45781810151838201526020016116cc565b838111156104bc5750506000910152565b6000815180845261170d8160208601602086016116c9565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b821515815260406020820152600061175a60408301846116f5565b949350505050565b6000806000806080858703121561177857600080fd5b61178185611496565b935061178f60208601611441565b9250604085013567ffffffffffffffff808211156117ac57600080fd5b6117b88883890161156e565b935060608701359150808211156117ce57600080fd5b506117db8782880161156e565b91505092959194509250565b73ffffffffffffffffffffffffffffffffffffffff8316815260406020820152600061175a60408301846116f5565b600181811c9082168061182a57607f821691505b602082108103611863577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f82111561059157600081815260208120601f850160051c810160208610156118905750805b601f850160051c820191505b818110156118af5782815560010161189c565b505050505050565b815167ffffffffffffffff8111156118d1576118d16114aa565b6118e5816118df8454611816565b84611869565b602080601f83116001811461193857600084156119025750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556118af565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b8281101561198557888601518255948401946001909101908401611966565b50858210156119c157878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b600082516119e38184602087016116c9565b9190910192915050565b63ffffffff8516815273ffffffffffffffffffffffffffffffffffffffff84166020820152608060408201526000611a2860808301856116f5565b8281036060840152611a3a81856116f5565b979650505050505050565b60008060408385031215611a5857600080fd5b82518015158114611a6857600080fd5b602084015190925067ffffffffffffffff811115611a8557600080fd5b8301601f81018513611a9657600080fd5b8051611aa461158d82611528565b818152866020838501011115611ab957600080fd5b611aca8260208301602086016116c9565b8093505050509250929050565b841515815283602082015273ffffffffffffffffffffffffffffffffffffffff83166040820152608060608201526000611b1460808301846116f5565b9695505050505050565b7fffffffff000000000000000000000000000000000000000000000000000000008560e01b1681527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008460601b16600482015260008351611b868160188501602088016116c9565b835190830190611b9d8160188401602088016116c9565b016018019695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611c1257611c12611bab565b500290565b602081526000611c2a60208301846116f5565b9392505050565b600060208284031215611c4357600080fd5b5051919050565b600082821015611c5c57611c5c611bab565b500390565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611c9257611c92611bab565b506001019056fea164736f6c634300080f000a",
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

// CheckPermittedCaller is a free data retrieval call binding the contract method 0xc3d8cae6.
//
// Solidity: function CheckPermittedCaller(bytes32 _EK, address _callerAddress) view returns(bool)
func (_BobaHCHelper *BobaHCHelperCaller) CheckPermittedCaller(opts *bind.CallOpts, _EK [32]byte, _callerAddress common.Address) (bool, error) {
	var out []interface{}
	err := _BobaHCHelper.contract.Call(opts, &out, "CheckPermittedCaller", _EK, _callerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPermittedCaller is a free data retrieval call binding the contract method 0xc3d8cae6.
//
// Solidity: function CheckPermittedCaller(bytes32 _EK, address _callerAddress) view returns(bool)
func (_BobaHCHelper *BobaHCHelperSession) CheckPermittedCaller(_EK [32]byte, _callerAddress common.Address) (bool, error) {
	return _BobaHCHelper.Contract.CheckPermittedCaller(&_BobaHCHelper.CallOpts, _EK, _callerAddress)
}

// CheckPermittedCaller is a free data retrieval call binding the contract method 0xc3d8cae6.
//
// Solidity: function CheckPermittedCaller(bytes32 _EK, address _callerAddress) view returns(bool)
func (_BobaHCHelper *BobaHCHelperCallerSession) CheckPermittedCaller(_EK [32]byte, _callerAddress common.Address) (bool, error) {
	return _BobaHCHelper.Contract.CheckPermittedCaller(&_BobaHCHelper.CallOpts, _EK, _callerAddress)
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

// AddPermittedCaller is a paid mutator transaction binding the contract method 0x79d9c090.
//
// Solidity: function AddPermittedCaller(bytes32 _EK, address _callerAddress) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) AddPermittedCaller(opts *bind.TransactOpts, _EK [32]byte, _callerAddress common.Address) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "AddPermittedCaller", _EK, _callerAddress)
}

// AddPermittedCaller is a paid mutator transaction binding the contract method 0x79d9c090.
//
// Solidity: function AddPermittedCaller(bytes32 _EK, address _callerAddress) returns()
func (_BobaHCHelper *BobaHCHelperSession) AddPermittedCaller(_EK [32]byte, _callerAddress common.Address) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.AddPermittedCaller(&_BobaHCHelper.TransactOpts, _EK, _callerAddress)
}

// AddPermittedCaller is a paid mutator transaction binding the contract method 0x79d9c090.
//
// Solidity: function AddPermittedCaller(bytes32 _EK, address _callerAddress) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) AddPermittedCaller(_EK [32]byte, _callerAddress common.Address) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.AddPermittedCaller(&_BobaHCHelper.TransactOpts, _EK, _callerAddress)
}

// CallOffchain is a paid mutator transaction binding the contract method 0x839f8457.
//
// Solidity: function CallOffchain(bytes32 EK, bytes payload) returns(bool, bytes)
func (_BobaHCHelper *BobaHCHelperTransactor) CallOffchain(opts *bind.TransactOpts, EK [32]byte, payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "CallOffchain", EK, payload)
}

// CallOffchain is a paid mutator transaction binding the contract method 0x839f8457.
//
// Solidity: function CallOffchain(bytes32 EK, bytes payload) returns(bool, bytes)
func (_BobaHCHelper *BobaHCHelperSession) CallOffchain(EK [32]byte, payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.CallOffchain(&_BobaHCHelper.TransactOpts, EK, payload)
}

// CallOffchain is a paid mutator transaction binding the contract method 0x839f8457.
//
// Solidity: function CallOffchain(bytes32 EK, bytes payload) returns(bool, bytes)
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

// GetResponse is a paid mutator transaction binding the contract method 0x8e5dc765.
//
// Solidity: function GetResponse(uint32 _rType, address _caller, string _url, bytes _payload) returns(bool, bytes)
func (_BobaHCHelper *BobaHCHelperTransactor) GetResponse(opts *bind.TransactOpts, _rType uint32, _caller common.Address, _url string, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "GetResponse", _rType, _caller, _url, _payload)
}

// GetResponse is a paid mutator transaction binding the contract method 0x8e5dc765.
//
// Solidity: function GetResponse(uint32 _rType, address _caller, string _url, bytes _payload) returns(bool, bytes)
func (_BobaHCHelper *BobaHCHelperSession) GetResponse(_rType uint32, _caller common.Address, _url string, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.GetResponse(&_BobaHCHelper.TransactOpts, _rType, _caller, _url, _payload)
}

// GetResponse is a paid mutator transaction binding the contract method 0x8e5dc765.
//
// Solidity: function GetResponse(uint32 _rType, address _caller, string _url, bytes _payload) returns(bool, bytes)
func (_BobaHCHelper *BobaHCHelperTransactorSession) GetResponse(_rType uint32, _caller common.Address, _url string, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.GetResponse(&_BobaHCHelper.TransactOpts, _rType, _caller, _url, _payload)
}

// PutResponse is a paid mutator transaction binding the contract method 0x11edaae0.
//
// Solidity: function PutResponse(bytes32 cacheKey, uint32 _code, bytes _payload) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) PutResponse(opts *bind.TransactOpts, cacheKey [32]byte, _code uint32, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "PutResponse", cacheKey, _code, _payload)
}

// PutResponse is a paid mutator transaction binding the contract method 0x11edaae0.
//
// Solidity: function PutResponse(bytes32 cacheKey, uint32 _code, bytes _payload) returns()
func (_BobaHCHelper *BobaHCHelperSession) PutResponse(cacheKey [32]byte, _code uint32, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.PutResponse(&_BobaHCHelper.TransactOpts, cacheKey, _code, _payload)
}

// PutResponse is a paid mutator transaction binding the contract method 0x11edaae0.
//
// Solidity: function PutResponse(bytes32 cacheKey, uint32 _code, bytes _payload) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) PutResponse(cacheKey [32]byte, _code uint32, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.PutResponse(&_BobaHCHelper.TransactOpts, cacheKey, _code, _payload)
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

// RemovePermittedCaller is a paid mutator transaction binding the contract method 0x040db828.
//
// Solidity: function RemovePermittedCaller(bytes32 _EK, address _callerAddress) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) RemovePermittedCaller(opts *bind.TransactOpts, _EK [32]byte, _callerAddress common.Address) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "RemovePermittedCaller", _EK, _callerAddress)
}

// RemovePermittedCaller is a paid mutator transaction binding the contract method 0x040db828.
//
// Solidity: function RemovePermittedCaller(bytes32 _EK, address _callerAddress) returns()
func (_BobaHCHelper *BobaHCHelperSession) RemovePermittedCaller(_EK [32]byte, _callerAddress common.Address) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.RemovePermittedCaller(&_BobaHCHelper.TransactOpts, _EK, _callerAddress)
}

// RemovePermittedCaller is a paid mutator transaction binding the contract method 0x040db828.
//
// Solidity: function RemovePermittedCaller(bytes32 _EK, address _callerAddress) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) RemovePermittedCaller(_EK [32]byte, _callerAddress common.Address) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.RemovePermittedCaller(&_BobaHCHelper.TransactOpts, _EK, _callerAddress)
}

// SimpleRandom is a paid mutator transaction binding the contract method 0xc6e4b0d3.
//
// Solidity: function SimpleRandom() returns(uint256)
func (_BobaHCHelper *BobaHCHelperTransactor) SimpleRandom(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "SimpleRandom")
}

// SimpleRandom is a paid mutator transaction binding the contract method 0xc6e4b0d3.
//
// Solidity: function SimpleRandom() returns(uint256)
func (_BobaHCHelper *BobaHCHelperSession) SimpleRandom() (*types.Transaction, error) {
	return _BobaHCHelper.Contract.SimpleRandom(&_BobaHCHelper.TransactOpts)
}

// SimpleRandom is a paid mutator transaction binding the contract method 0xc6e4b0d3.
//
// Solidity: function SimpleRandom() returns(uint256)
func (_BobaHCHelper *BobaHCHelperTransactorSession) SimpleRandom() (*types.Transaction, error) {
	return _BobaHCHelper.Contract.SimpleRandom(&_BobaHCHelper.TransactOpts)
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
