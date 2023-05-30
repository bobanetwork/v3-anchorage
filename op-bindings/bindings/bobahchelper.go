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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"DBG\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"URL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"EndpointRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"URL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"EndpointUnregistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"AddCredit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_callerAddress\",\"type\":\"address\"}],\"name\":\"AddPermittedCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"CallOffchain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_callerAddress\",\"type\":\"address\"}],\"name\":\"CheckPermittedCaller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Endpoints\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rType\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_random\",\"type\":\"uint256\"}],\"name\":\"GetRandom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"rType\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"GetResponse\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_cacheKey\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_returndata\",\"type\":\"bytes\"}],\"name\":\"PutResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_auth\",\"type\":\"bytes32\"}],\"name\":\"RegisterEndpoint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_callerAddress\",\"type\":\"address\"}],\"name\":\"RemovePermittedCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"session\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"myNum\",\"type\":\"uint256\"}],\"name\":\"SequentialRandom\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SimpleRandom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"UnregisterEndpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingCharge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"prepaidCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040527342000000000000000000000000000000000000fd60805273deaddeaddeaddeaddeaddeaddeaddeaddead990160a0527342000000000000000000000000000000000000fe60c05234801561005857600080fd5b5060805160a05160c05161269c61008f60003960008181610d3d0152610e300152600061148701526000610da8015261269c6000f3fe608060405234801561001057600080fd5b50600436106100ff5760003560e01c80637e12141311610097578063c1fd7e4611610066578063c1fd7e4614610287578063c6e4b0d3146102a8578063e2f94792146102b0578063eb6598b5146102c357600080fd5b80637e121413146101d657806387fc1b25146101f6578063997e31cc14610209578063a71adb3e1461026457600080fd5b806332be428f116100d357806332be428f14610168578063493d57d61461019057806362fccf6e146101a35780637d93616c146101b657600080fd5b8062292526146101045780631730f39b14610120578063248eaf62146101355780633042902714610155575b600080fd5b61010d60005481565b6040519081526020015b60405180910390f35b61013361012e366004611cd7565b6102d6565b005b61010d610143366004611d2b565b60026020526000908152604090205481565b610133610163366004611cd7565b6104a1565b61017b610176366004611d46565b610662565b60408051928352602083019190915201610117565b61010d61019e366004611d86565b6107a1565b6101336101b1366004611db0565b610937565b6101c96101c4366004611ef4565b610b10565b6040516101179190612009565b61010d6101e4366004611d2b565b60016020526000908152604090205481565b61013361020436600461201c565b610cb9565b61023f610217366004612038565b60036020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610117565b610277610272366004612051565b610da3565b6040519015158152602001610117565b61029a61029536600461209d565b611037565b604051610117929190612109565b61010d611235565b6102776102be366004611cd7565b611358565b6101336102d1366004612132565b61146f565b600083836040516020016102eb929190612182565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600390935291205490915073ffffffffffffffffffffffffffffffffffffffff166103b2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f74207265676973746572656400000000000060448201526064015b60405180910390fd5b60008181526003602052604090205473ffffffffffffffffffffffffffffffffffffffff16331461043f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4e6f742074686520456e64706f696e74206f776e65720000000000000000000060448201526064016103a9565b600090815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff90941683526001938401909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690911790555050565b600083836040516020016104b6929190612182565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600390935291205490915073ffffffffffffffffffffffffffffffffffffffff16610578576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f74207265676973746572656400000000000060448201526064016103a9565b60008181526003602052604090205473ffffffffffffffffffffffffffffffffffffffff163314610605576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4e6f742074686520456e64706f696e74206f776e65720000000000000000000060448201526064016103a9565b600090815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff9094168352600190930190522080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555050565b6040517f32be428f00000000000000000000000000000000000000000000000000000000602082015233606090811b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602483015260388201859052605882018490526078820183905260009182916332be428f9183919082906098016040516020818303038152906040528051906020012090506107048133601e6115f0565b9093509150828015610717575081516040145b61077d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f52616e646f6d206e756d6265722067656e65726174696f6e206661696c65640060448201526064016103a9565b818060200190518101906107919190612192565b9550955050505050935093915050565b60008263ffffffff16600114610813576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f545552494e473a204765746820696e74657263657074206661696c757265000060448201526064016103a9565b6040517f493d57d600000000000000000000000000000000000000000000000000000000602082015233606090811b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602483015263493d57d69160009190829060380160405160208183030381529060405280519060200120905061089c8133601e6115f0565b90935091508280156108af575081516020145b610915576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f52616e646f6d206e756d6265722067656e65726174696f6e206661696c65640060448201526064016103a9565b60008280602001905181019061092b91906121b6565b98975050505050505050565b6000828260405160200161094c929190612182565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600390935291205490915073ffffffffffffffffffffffffffffffffffffffff16610a0e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f74207265676973746572656400000000000060448201526064016103a9565b60008181526003602052604090205473ffffffffffffffffffffffffffffffffffffffff163314610a9b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4e6f742074686520456e64706f696e74206f776e65720000000000000000000060448201526064016103a9565b6000818152600360205260409081902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055517fddcfd504e796f664dc8ed3395209f4798379b47d8ee42d4ee5142a39f6d6d21a90610b0390859085903390612218565b60405180910390a1505050565b60608363ffffffff1660011480610b3057508363ffffffff166302000001145b610b96576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f545552494e473a204765746820696e74657263657074206661696c757265000060448201526064016103a9565b6000637d93616c90506000813386604051602001610bb49190612252565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052610bf3939291889060200161226e565b60405160208183030381529060405280519060200120905060006060610c1b833360326115f0565b909250905081610cac576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f545552494e473a204c656761637920476574526573706f6e736528292066616960448201527f6c7572650000000000000000000000000000000000000000000000000000000060648201526084016103a9565b93505050505b9392505050565b60008111610d23576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f496e76616c696420616d6f756e7400000000000000000000000000000000000060448201526064016103a9565b610d6573ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084611863565b73ffffffffffffffffffffffffffffffffffffffff821660009081526001602052604081208054839290610d9a90849061232a565b90915550505050565b6000807f0000000000000000000000000000000000000000000000000000000000000000905060008585604051602001610dde929190612182565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209050610e5973ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633306064611863565b6064600080828254610e6b919061232a565b90915550506040805160048082526024820183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f61f334d30000000000000000000000000000000000000000000000000000000017905291517fc1fd7e46000000000000000000000000000000000000000000000000000000008152909160609160009173ffffffffffffffffffffffffffffffffffffffff87169163c1fd7e4691610f22918d918d91899101612342565b6000604051808303816000875af1158015610f41573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610f879190810190612368565b92509050808015610f99575081516020145b8015610faa57508151602083012087145b15611028576000848152600360205260409081902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000163390811790915590517fd48d01edeb78cd5b560a64aebf74c66df4140fc9b5f72bff4ec03f43e7e9f8829161101b918c918c91612218565b60405180910390a161092b565b50600098975050505050505050565b60006060600063c1fd7e46905060008787604051602001611059929190612182565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012090503330146111a95760008181526003602052604090205473ffffffffffffffffffffffffffffffffffffffff16611123576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f456e64706f696e74206e6f74207265676973746572656400000000000000000060448201526064016103a9565b600081815260036020908152604080832033845260010190915290205460ff166111a9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616c6c6572206973206e6f74207065726d697474656400000000000000000060448201526064016103a9565b600082338a8a6040516020016111c0929190612182565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152908290526112019392918b908b906020016123f5565b604051602081830303815290604052805190602001209050611225813360326115f0565b9450945050505094509492505050565b6040517fc6e4b0d300000000000000000000000000000000000000000000000000000000602082015233606090811b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602483015260009163c6e4b0d391839182906038016040516020818303038152906040528051906020012090506112bf8133601e6115f0565b90935091508280156112d2575081516020145b611338576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f52616e646f6d206e756d6265722067656e65726174696f6e206661696c65640060448201526064016103a9565b60008280602001905181019061134e91906121b6565b9695505050505050565b600080848460405160200161136e929190612182565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600390935291205490915073ffffffffffffffffffffffffffffffffffffffff16611430576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f74207265676973746572656400000000000060448201526064016103a9565b600090815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616845260010190915290205460ff1690509392505050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461150e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f496e76616c696420507574526573706f6e736528292063616c6c65720000000060448201526064016103a9565b600084815260046020526040902080546115279061247b565b159050611590576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4445425547202d20416c7265616479206578697374730000000000000000000060448201526064016103a9565b8282826040516020016115a5939291906124ce565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526000868152600460205220906115e99082612541565b5050505050565b6000606073ffffffffffffffffffffffffffffffffffffffff841630146117275773ffffffffffffffffffffffffffffffffffffffff8416600090815260026020526040812054611641908561232a565b73ffffffffffffffffffffffffffffffffffffffff86166000908152600160205260409020549091508111156116d3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f496e73756666696369656e74206372656469740000000000000000000000000060448201526064016103a9565b73ffffffffffffffffffffffffffffffffffffffff85166000908152600160205260408120805486929061170890849061265b565b9250508190555083600080828254611720919061232a565b9091555050505b600085815260046020526040812080546117409061247b565b80601f016020809104026020016040519081016040528092919081815260200182805461176c9061247b565b80156117b95780601f1061178e576101008083540402835291602001916117b9565b820191906000526020600020905b81548152906001019060200180831161179c57829003601f168201915b50505050509050600081511161182b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f476574526573706f6e73653a204d697373696e6720636163686520656e74727960448201526064016103a9565b600086815260046020526040812061184291611c0f565b808060200190518101906118569190612368565b9250925050935093915050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd000000000000000000000000000000000000000000000000000000001790526118f89085906118fe565b50505050565b6000611960826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16611a0f9092919063ffffffff16565b805190915015611a0a578080602001905181019061197e9190612672565b611a0a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016103a9565b505050565b6060611a1e8484600085611a26565b949350505050565b606082471015611ab8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016103a9565b73ffffffffffffffffffffffffffffffffffffffff85163b611b36576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016103a9565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051611b5f9190612252565b60006040518083038185875af1925050503d8060008114611b9c576040519150601f19603f3d011682016040523d82523d6000602084013e611ba1565b606091505b5091509150611bb1828286611bbc565b979650505050505050565b60608315611bcb575081610cb2565b825115611bdb5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103a99190612009565b508054611c1b9061247b565b6000825580601f10611c2b575050565b601f016020900490600052602060002090810190611c499190611c4c565b50565b5b80821115611c615760008155600101611c4d565b5090565b60008083601f840112611c7757600080fd5b50813567ffffffffffffffff811115611c8f57600080fd5b602083019150836020828501011115611ca757600080fd5b9250929050565b803573ffffffffffffffffffffffffffffffffffffffff81168114611cd257600080fd5b919050565b600080600060408486031215611cec57600080fd5b833567ffffffffffffffff811115611d0357600080fd5b611d0f86828701611c65565b9094509250611d22905060208501611cae565b90509250925092565b600060208284031215611d3d57600080fd5b610cb282611cae565b600080600060608486031215611d5b57600080fd5b505081359360208301359350604090920135919050565b803563ffffffff81168114611cd257600080fd5b60008060408385031215611d9957600080fd5b611da283611d72565b946020939093013593505050565b60008060208385031215611dc357600080fd5b823567ffffffffffffffff811115611dda57600080fd5b611de685828601611c65565b90969095509350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611e6857611e68611df2565b604052919050565b600067ffffffffffffffff821115611e8a57611e8a611df2565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6000611ec9611ec484611e70565b611e21565b9050828152838383011115611edd57600080fd5b828260208301376000602084830101529392505050565b600080600060608486031215611f0957600080fd5b611f1284611d72565b9250602084013567ffffffffffffffff80821115611f2f57600080fd5b818601915086601f830112611f4357600080fd5b611f5287833560208501611eb6565b93506040860135915080821115611f6857600080fd5b508401601f81018613611f7a57600080fd5b611f8986823560208401611eb6565b9150509250925092565b60005b83811015611fae578181015183820152602001611f96565b838111156118f85750506000910152565b60008151808452611fd7816020860160208601611f93565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610cb26020830184611fbf565b6000806040838503121561202f57600080fd5b611da283611cae565b60006020828403121561204a57600080fd5b5035919050565b60008060006040848603121561206657600080fd5b833567ffffffffffffffff81111561207d57600080fd5b61208986828701611c65565b909790965060209590950135949350505050565b600080600080604085870312156120b357600080fd5b843567ffffffffffffffff808211156120cb57600080fd5b6120d788838901611c65565b909650945060208701359150808211156120f057600080fd5b506120fd87828801611c65565b95989497509550505050565b8215158152604060208201526000611a1e6040830184611fbf565b8015158114611c4957600080fd5b6000806000806060858703121561214857600080fd5b84359350602085013561215a81612124565b9250604085013567ffffffffffffffff81111561217657600080fd5b6120fd87828801611c65565b8183823760009101908152919050565b600080604083850312156121a557600080fd5b505080516020909101519092909150565b6000602082840312156121c857600080fd5b5051919050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60408152600061222c6040830185876121cf565b905073ffffffffffffffffffffffffffffffffffffffff83166020830152949350505050565b60008251612264818460208701611f93565b9190910192915050565b7fffffffff000000000000000000000000000000000000000000000000000000008560e01b1681527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008460601b166004820152600083516122d6816018850160208801611f93565b8351908301906122ed816018840160208801611f93565b016018019695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561233d5761233d6122fb565b500190565b6040815260006123566040830185876121cf565b828103602084015261134e8185611fbf565b6000806040838503121561237b57600080fd5b825161238681612124565b602084015190925067ffffffffffffffff8111156123a357600080fd5b8301601f810185136123b457600080fd5b80516123c2611ec482611e70565b8181528660208385010111156123d757600080fd5b6123e8826020830160208601611f93565b8093505050509250929050565b7fffffffff000000000000000000000000000000000000000000000000000000008660e01b1681527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008560601b1660048201526000845161245d816018850160208901611f93565b82018385601883013760009301601801928352509095945050505050565b600181811c9082168061248f57607f821691505b6020821081036124c8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b83151581526040602082015260006124ea6040830184866121cf565b95945050505050565b601f821115611a0a57600081815260208120601f850160051c8101602086101561251a5750805b601f850160051c820191505b8181101561253957828155600101612526565b505050505050565b815167ffffffffffffffff81111561255b5761255b611df2565b61256f81612569845461247b565b846124f3565b602080601f8311600181146125c2576000841561258c5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555612539565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b8281101561260f578886015182559484019460019091019084016125f0565b508582101561264b57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b60008282101561266d5761266d6122fb565b500390565b60006020828403121561268457600080fd5b8151610cb28161212456fea164736f6c634300080f000a",
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

// CheckPermittedCaller is a free data retrieval call binding the contract method 0xe2f94792.
//
// Solidity: function CheckPermittedCaller(string _url, address _callerAddress) view returns(bool)
func (_BobaHCHelper *BobaHCHelperCaller) CheckPermittedCaller(opts *bind.CallOpts, _url string, _callerAddress common.Address) (bool, error) {
	var out []interface{}
	err := _BobaHCHelper.contract.Call(opts, &out, "CheckPermittedCaller", _url, _callerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckPermittedCaller is a free data retrieval call binding the contract method 0xe2f94792.
//
// Solidity: function CheckPermittedCaller(string _url, address _callerAddress) view returns(bool)
func (_BobaHCHelper *BobaHCHelperSession) CheckPermittedCaller(_url string, _callerAddress common.Address) (bool, error) {
	return _BobaHCHelper.Contract.CheckPermittedCaller(&_BobaHCHelper.CallOpts, _url, _callerAddress)
}

// CheckPermittedCaller is a free data retrieval call binding the contract method 0xe2f94792.
//
// Solidity: function CheckPermittedCaller(string _url, address _callerAddress) view returns(bool)
func (_BobaHCHelper *BobaHCHelperCallerSession) CheckPermittedCaller(_url string, _callerAddress common.Address) (bool, error) {
	return _BobaHCHelper.Contract.CheckPermittedCaller(&_BobaHCHelper.CallOpts, _url, _callerAddress)
}

// Endpoints is a free data retrieval call binding the contract method 0x997e31cc.
//
// Solidity: function Endpoints(bytes32 ) view returns(address Owner)
func (_BobaHCHelper *BobaHCHelperCaller) Endpoints(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _BobaHCHelper.contract.Call(opts, &out, "Endpoints", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Endpoints is a free data retrieval call binding the contract method 0x997e31cc.
//
// Solidity: function Endpoints(bytes32 ) view returns(address Owner)
func (_BobaHCHelper *BobaHCHelperSession) Endpoints(arg0 [32]byte) (common.Address, error) {
	return _BobaHCHelper.Contract.Endpoints(&_BobaHCHelper.CallOpts, arg0)
}

// Endpoints is a free data retrieval call binding the contract method 0x997e31cc.
//
// Solidity: function Endpoints(bytes32 ) view returns(address Owner)
func (_BobaHCHelper *BobaHCHelperCallerSession) Endpoints(arg0 [32]byte) (common.Address, error) {
	return _BobaHCHelper.Contract.Endpoints(&_BobaHCHelper.CallOpts, arg0)
}

// OwnerRevenue is a free data retrieval call binding the contract method 0x00292526.
//
// Solidity: function ownerRevenue() view returns(uint256)
func (_BobaHCHelper *BobaHCHelperCaller) OwnerRevenue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BobaHCHelper.contract.Call(opts, &out, "ownerRevenue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerRevenue is a free data retrieval call binding the contract method 0x00292526.
//
// Solidity: function ownerRevenue() view returns(uint256)
func (_BobaHCHelper *BobaHCHelperSession) OwnerRevenue() (*big.Int, error) {
	return _BobaHCHelper.Contract.OwnerRevenue(&_BobaHCHelper.CallOpts)
}

// OwnerRevenue is a free data retrieval call binding the contract method 0x00292526.
//
// Solidity: function ownerRevenue() view returns(uint256)
func (_BobaHCHelper *BobaHCHelperCallerSession) OwnerRevenue() (*big.Int, error) {
	return _BobaHCHelper.Contract.OwnerRevenue(&_BobaHCHelper.CallOpts)
}

// PendingCharge is a free data retrieval call binding the contract method 0x248eaf62.
//
// Solidity: function pendingCharge(address ) view returns(uint256)
func (_BobaHCHelper *BobaHCHelperCaller) PendingCharge(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BobaHCHelper.contract.Call(opts, &out, "pendingCharge", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingCharge is a free data retrieval call binding the contract method 0x248eaf62.
//
// Solidity: function pendingCharge(address ) view returns(uint256)
func (_BobaHCHelper *BobaHCHelperSession) PendingCharge(arg0 common.Address) (*big.Int, error) {
	return _BobaHCHelper.Contract.PendingCharge(&_BobaHCHelper.CallOpts, arg0)
}

// PendingCharge is a free data retrieval call binding the contract method 0x248eaf62.
//
// Solidity: function pendingCharge(address ) view returns(uint256)
func (_BobaHCHelper *BobaHCHelperCallerSession) PendingCharge(arg0 common.Address) (*big.Int, error) {
	return _BobaHCHelper.Contract.PendingCharge(&_BobaHCHelper.CallOpts, arg0)
}

// PrepaidCredit is a free data retrieval call binding the contract method 0x7e121413.
//
// Solidity: function prepaidCredit(address ) view returns(uint256)
func (_BobaHCHelper *BobaHCHelperCaller) PrepaidCredit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BobaHCHelper.contract.Call(opts, &out, "prepaidCredit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrepaidCredit is a free data retrieval call binding the contract method 0x7e121413.
//
// Solidity: function prepaidCredit(address ) view returns(uint256)
func (_BobaHCHelper *BobaHCHelperSession) PrepaidCredit(arg0 common.Address) (*big.Int, error) {
	return _BobaHCHelper.Contract.PrepaidCredit(&_BobaHCHelper.CallOpts, arg0)
}

// PrepaidCredit is a free data retrieval call binding the contract method 0x7e121413.
//
// Solidity: function prepaidCredit(address ) view returns(uint256)
func (_BobaHCHelper *BobaHCHelperCallerSession) PrepaidCredit(arg0 common.Address) (*big.Int, error) {
	return _BobaHCHelper.Contract.PrepaidCredit(&_BobaHCHelper.CallOpts, arg0)
}

// AddCredit is a paid mutator transaction binding the contract method 0x87fc1b25.
//
// Solidity: function AddCredit(address _caller, uint256 _amount) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) AddCredit(opts *bind.TransactOpts, _caller common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "AddCredit", _caller, _amount)
}

// AddCredit is a paid mutator transaction binding the contract method 0x87fc1b25.
//
// Solidity: function AddCredit(address _caller, uint256 _amount) returns()
func (_BobaHCHelper *BobaHCHelperSession) AddCredit(_caller common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.AddCredit(&_BobaHCHelper.TransactOpts, _caller, _amount)
}

// AddCredit is a paid mutator transaction binding the contract method 0x87fc1b25.
//
// Solidity: function AddCredit(address _caller, uint256 _amount) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) AddCredit(_caller common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.AddCredit(&_BobaHCHelper.TransactOpts, _caller, _amount)
}

// AddPermittedCaller is a paid mutator transaction binding the contract method 0x1730f39b.
//
// Solidity: function AddPermittedCaller(string _url, address _callerAddress) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) AddPermittedCaller(opts *bind.TransactOpts, _url string, _callerAddress common.Address) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "AddPermittedCaller", _url, _callerAddress)
}

// AddPermittedCaller is a paid mutator transaction binding the contract method 0x1730f39b.
//
// Solidity: function AddPermittedCaller(string _url, address _callerAddress) returns()
func (_BobaHCHelper *BobaHCHelperSession) AddPermittedCaller(_url string, _callerAddress common.Address) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.AddPermittedCaller(&_BobaHCHelper.TransactOpts, _url, _callerAddress)
}

// AddPermittedCaller is a paid mutator transaction binding the contract method 0x1730f39b.
//
// Solidity: function AddPermittedCaller(string _url, address _callerAddress) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) AddPermittedCaller(_url string, _callerAddress common.Address) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.AddPermittedCaller(&_BobaHCHelper.TransactOpts, _url, _callerAddress)
}

// CallOffchain is a paid mutator transaction binding the contract method 0xc1fd7e46.
//
// Solidity: function CallOffchain(string _url, bytes _payload) returns(bool, bytes)
func (_BobaHCHelper *BobaHCHelperTransactor) CallOffchain(opts *bind.TransactOpts, _url string, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "CallOffchain", _url, _payload)
}

// CallOffchain is a paid mutator transaction binding the contract method 0xc1fd7e46.
//
// Solidity: function CallOffchain(string _url, bytes _payload) returns(bool, bytes)
func (_BobaHCHelper *BobaHCHelperSession) CallOffchain(_url string, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.CallOffchain(&_BobaHCHelper.TransactOpts, _url, _payload)
}

// CallOffchain is a paid mutator transaction binding the contract method 0xc1fd7e46.
//
// Solidity: function CallOffchain(string _url, bytes _payload) returns(bool, bytes)
func (_BobaHCHelper *BobaHCHelperTransactorSession) CallOffchain(_url string, _payload []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.CallOffchain(&_BobaHCHelper.TransactOpts, _url, _payload)
}

// GetRandom is a paid mutator transaction binding the contract method 0x493d57d6.
//
// Solidity: function GetRandom(uint32 rType, uint256 _random) returns(uint256)
func (_BobaHCHelper *BobaHCHelperTransactor) GetRandom(opts *bind.TransactOpts, rType uint32, _random *big.Int) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "GetRandom", rType, _random)
}

// GetRandom is a paid mutator transaction binding the contract method 0x493d57d6.
//
// Solidity: function GetRandom(uint32 rType, uint256 _random) returns(uint256)
func (_BobaHCHelper *BobaHCHelperSession) GetRandom(rType uint32, _random *big.Int) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.GetRandom(&_BobaHCHelper.TransactOpts, rType, _random)
}

// GetRandom is a paid mutator transaction binding the contract method 0x493d57d6.
//
// Solidity: function GetRandom(uint32 rType, uint256 _random) returns(uint256)
func (_BobaHCHelper *BobaHCHelperTransactorSession) GetRandom(rType uint32, _random *big.Int) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.GetRandom(&_BobaHCHelper.TransactOpts, rType, _random)
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

// PutResponse is a paid mutator transaction binding the contract method 0xeb6598b5.
//
// Solidity: function PutResponse(bytes32 _cacheKey, bool _success, bytes _returndata) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) PutResponse(opts *bind.TransactOpts, _cacheKey [32]byte, _success bool, _returndata []byte) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "PutResponse", _cacheKey, _success, _returndata)
}

// PutResponse is a paid mutator transaction binding the contract method 0xeb6598b5.
//
// Solidity: function PutResponse(bytes32 _cacheKey, bool _success, bytes _returndata) returns()
func (_BobaHCHelper *BobaHCHelperSession) PutResponse(_cacheKey [32]byte, _success bool, _returndata []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.PutResponse(&_BobaHCHelper.TransactOpts, _cacheKey, _success, _returndata)
}

// PutResponse is a paid mutator transaction binding the contract method 0xeb6598b5.
//
// Solidity: function PutResponse(bytes32 _cacheKey, bool _success, bytes _returndata) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) PutResponse(_cacheKey [32]byte, _success bool, _returndata []byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.PutResponse(&_BobaHCHelper.TransactOpts, _cacheKey, _success, _returndata)
}

// RegisterEndpoint is a paid mutator transaction binding the contract method 0xa71adb3e.
//
// Solidity: function RegisterEndpoint(string _url, bytes32 _auth) returns(bool)
func (_BobaHCHelper *BobaHCHelperTransactor) RegisterEndpoint(opts *bind.TransactOpts, _url string, _auth [32]byte) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "RegisterEndpoint", _url, _auth)
}

// RegisterEndpoint is a paid mutator transaction binding the contract method 0xa71adb3e.
//
// Solidity: function RegisterEndpoint(string _url, bytes32 _auth) returns(bool)
func (_BobaHCHelper *BobaHCHelperSession) RegisterEndpoint(_url string, _auth [32]byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.RegisterEndpoint(&_BobaHCHelper.TransactOpts, _url, _auth)
}

// RegisterEndpoint is a paid mutator transaction binding the contract method 0xa71adb3e.
//
// Solidity: function RegisterEndpoint(string _url, bytes32 _auth) returns(bool)
func (_BobaHCHelper *BobaHCHelperTransactorSession) RegisterEndpoint(_url string, _auth [32]byte) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.RegisterEndpoint(&_BobaHCHelper.TransactOpts, _url, _auth)
}

// RemovePermittedCaller is a paid mutator transaction binding the contract method 0x30429027.
//
// Solidity: function RemovePermittedCaller(string _url, address _callerAddress) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) RemovePermittedCaller(opts *bind.TransactOpts, _url string, _callerAddress common.Address) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "RemovePermittedCaller", _url, _callerAddress)
}

// RemovePermittedCaller is a paid mutator transaction binding the contract method 0x30429027.
//
// Solidity: function RemovePermittedCaller(string _url, address _callerAddress) returns()
func (_BobaHCHelper *BobaHCHelperSession) RemovePermittedCaller(_url string, _callerAddress common.Address) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.RemovePermittedCaller(&_BobaHCHelper.TransactOpts, _url, _callerAddress)
}

// RemovePermittedCaller is a paid mutator transaction binding the contract method 0x30429027.
//
// Solidity: function RemovePermittedCaller(string _url, address _callerAddress) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) RemovePermittedCaller(_url string, _callerAddress common.Address) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.RemovePermittedCaller(&_BobaHCHelper.TransactOpts, _url, _callerAddress)
}

// SequentialRandom is a paid mutator transaction binding the contract method 0x32be428f.
//
// Solidity: function SequentialRandom(bytes32 session, bytes32 nextHash, uint256 myNum) returns(bytes32, uint256)
func (_BobaHCHelper *BobaHCHelperTransactor) SequentialRandom(opts *bind.TransactOpts, session [32]byte, nextHash [32]byte, myNum *big.Int) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "SequentialRandom", session, nextHash, myNum)
}

// SequentialRandom is a paid mutator transaction binding the contract method 0x32be428f.
//
// Solidity: function SequentialRandom(bytes32 session, bytes32 nextHash, uint256 myNum) returns(bytes32, uint256)
func (_BobaHCHelper *BobaHCHelperSession) SequentialRandom(session [32]byte, nextHash [32]byte, myNum *big.Int) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.SequentialRandom(&_BobaHCHelper.TransactOpts, session, nextHash, myNum)
}

// SequentialRandom is a paid mutator transaction binding the contract method 0x32be428f.
//
// Solidity: function SequentialRandom(bytes32 session, bytes32 nextHash, uint256 myNum) returns(bytes32, uint256)
func (_BobaHCHelper *BobaHCHelperTransactorSession) SequentialRandom(session [32]byte, nextHash [32]byte, myNum *big.Int) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.SequentialRandom(&_BobaHCHelper.TransactOpts, session, nextHash, myNum)
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

// UnregisterEndpoint is a paid mutator transaction binding the contract method 0x62fccf6e.
//
// Solidity: function UnregisterEndpoint(string _url) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) UnregisterEndpoint(opts *bind.TransactOpts, _url string) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "UnregisterEndpoint", _url)
}

// UnregisterEndpoint is a paid mutator transaction binding the contract method 0x62fccf6e.
//
// Solidity: function UnregisterEndpoint(string _url) returns()
func (_BobaHCHelper *BobaHCHelperSession) UnregisterEndpoint(_url string) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.UnregisterEndpoint(&_BobaHCHelper.TransactOpts, _url)
}

// UnregisterEndpoint is a paid mutator transaction binding the contract method 0x62fccf6e.
//
// Solidity: function UnregisterEndpoint(string _url) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) UnregisterEndpoint(_url string) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.UnregisterEndpoint(&_BobaHCHelper.TransactOpts, _url)
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
	URL   string
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEndpointRegistered is a free log retrieval operation binding the contract event 0xd48d01edeb78cd5b560a64aebf74c66df4140fc9b5f72bff4ec03f43e7e9f882.
//
// Solidity: event EndpointRegistered(string URL, address owner)
func (_BobaHCHelper *BobaHCHelperFilterer) FilterEndpointRegistered(opts *bind.FilterOpts) (*BobaHCHelperEndpointRegisteredIterator, error) {

	logs, sub, err := _BobaHCHelper.contract.FilterLogs(opts, "EndpointRegistered")
	if err != nil {
		return nil, err
	}
	return &BobaHCHelperEndpointRegisteredIterator{contract: _BobaHCHelper.contract, event: "EndpointRegistered", logs: logs, sub: sub}, nil
}

// WatchEndpointRegistered is a free log subscription operation binding the contract event 0xd48d01edeb78cd5b560a64aebf74c66df4140fc9b5f72bff4ec03f43e7e9f882.
//
// Solidity: event EndpointRegistered(string URL, address owner)
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

// ParseEndpointRegistered is a log parse operation binding the contract event 0xd48d01edeb78cd5b560a64aebf74c66df4140fc9b5f72bff4ec03f43e7e9f882.
//
// Solidity: event EndpointRegistered(string URL, address owner)
func (_BobaHCHelper *BobaHCHelperFilterer) ParseEndpointRegistered(log types.Log) (*BobaHCHelperEndpointRegistered, error) {
	event := new(BobaHCHelperEndpointRegistered)
	if err := _BobaHCHelper.contract.UnpackLog(event, "EndpointRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BobaHCHelperEndpointUnregisteredIterator is returned from FilterEndpointUnregistered and is used to iterate over the raw logs and unpacked data for EndpointUnregistered events raised by the BobaHCHelper contract.
type BobaHCHelperEndpointUnregisteredIterator struct {
	Event *BobaHCHelperEndpointUnregistered // Event containing the contract specifics and raw log

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
func (it *BobaHCHelperEndpointUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BobaHCHelperEndpointUnregistered)
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
		it.Event = new(BobaHCHelperEndpointUnregistered)
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
func (it *BobaHCHelperEndpointUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BobaHCHelperEndpointUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BobaHCHelperEndpointUnregistered represents a EndpointUnregistered event raised by the BobaHCHelper contract.
type BobaHCHelperEndpointUnregistered struct {
	URL   string
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEndpointUnregistered is a free log retrieval operation binding the contract event 0xddcfd504e796f664dc8ed3395209f4798379b47d8ee42d4ee5142a39f6d6d21a.
//
// Solidity: event EndpointUnregistered(string URL, address owner)
func (_BobaHCHelper *BobaHCHelperFilterer) FilterEndpointUnregistered(opts *bind.FilterOpts) (*BobaHCHelperEndpointUnregisteredIterator, error) {

	logs, sub, err := _BobaHCHelper.contract.FilterLogs(opts, "EndpointUnregistered")
	if err != nil {
		return nil, err
	}
	return &BobaHCHelperEndpointUnregisteredIterator{contract: _BobaHCHelper.contract, event: "EndpointUnregistered", logs: logs, sub: sub}, nil
}

// WatchEndpointUnregistered is a free log subscription operation binding the contract event 0xddcfd504e796f664dc8ed3395209f4798379b47d8ee42d4ee5142a39f6d6d21a.
//
// Solidity: event EndpointUnregistered(string URL, address owner)
func (_BobaHCHelper *BobaHCHelperFilterer) WatchEndpointUnregistered(opts *bind.WatchOpts, sink chan<- *BobaHCHelperEndpointUnregistered) (event.Subscription, error) {

	logs, sub, err := _BobaHCHelper.contract.WatchLogs(opts, "EndpointUnregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BobaHCHelperEndpointUnregistered)
				if err := _BobaHCHelper.contract.UnpackLog(event, "EndpointUnregistered", log); err != nil {
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

// ParseEndpointUnregistered is a log parse operation binding the contract event 0xddcfd504e796f664dc8ed3395209f4798379b47d8ee42d4ee5142a39f6d6d21a.
//
// Solidity: event EndpointUnregistered(string URL, address owner)
func (_BobaHCHelper *BobaHCHelperFilterer) ParseEndpointUnregistered(log types.Log) (*BobaHCHelperEndpointUnregistered, error) {
	event := new(BobaHCHelperEndpointUnregistered)
	if err := _BobaHCHelper.contract.UnpackLog(event, "EndpointUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
