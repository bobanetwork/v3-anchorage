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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"AddBalanceTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"DBG\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"Success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"EK\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"URL\",\"type\":\"string\"}],\"name\":\"EndpointRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"random\",\"type\":\"uint256\"}],\"name\":\"OffchainRandom\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_EK\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_callerAddress\",\"type\":\"address\"}],\"name\":\"AddPermittedCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"EK\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"CallOffchain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_EK\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_callerAddress\",\"type\":\"address\"}],\"name\":\"CheckPermittedCaller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Endpoints\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"URL\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"CreditBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cacheKey\",\"type\":\"bytes32\"}],\"name\":\"ExpiredResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_rType\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"GetResponse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"session\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"myNum\",\"type\":\"uint256\"}],\"name\":\"NextRandom\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cacheKey\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_code\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"PutResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"credits\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"auth\",\"type\":\"bytes32\"}],\"name\":\"RegisterEndpoint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_EK\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_callerAddress\",\"type\":\"address\"}],\"name\":\"RemovePermittedCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SimpleRandom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"EK\",\"type\":\"bytes32\"}],\"name\":\"UnregisterEndpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_EK\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_addBalanceAmount\",\"type\":\"uint256\"}],\"name\":\"addBalanceTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040527342000000000000000000000000000000000000fd60805273deaddeaddeaddeaddeaddeaddeaddeaddead990160a0527342000000000000000000000000000000000000fe60c05234801561005857600080fd5b5060805160a05160c0516127e76100c0600039600081816106270152818161096e01528181610a2401528181610c87015281816116ed01526118000152600081816103d0015261066a01526000818161074001528181610ea801526116ae01526127e76000f3fe608060405234801561001057600080fd5b50600436106100e95760003560e01c806379d9c0901161008c578063997e31cc11610066578063997e31cc146101d5578063ade42320146101f7578063c3d8cae614610225578063c6e4b0d31461023857600080fd5b806379d9c0901461018e578063839f8457146101a15780638e5dc765146101c257600080fd5b80631f999764116100c85780631f9997641461013257806322a522a4146101455780635b0dfe6f146101585780636c6459c81461017b57600080fd5b8062292526146100ee578063040db8281461010a57806311edaae01461011f575b600080fd5b6100f760005481565b6040519081526020015b60405180910390f35b61011d610118366004611f2c565b610240565b005b61011d61012d366004612086565b6103b8565b61011d6101403660046120dd565b61052f565b61011d6101533660046120dd565b610652565b61016b6101663660046120f6565b61073b565b6040519015158152602001610101565b61011d610189366004612144565b610aeb565b61011d61019c366004611f2c565b610d1a565b6101b46101af366004612166565b610e90565b604051610101929190612223565b6101b46101d036600461223e565b6111d3565b6101e86101e33660046120dd565b6114ee565b604051610101939291906122c3565b610210610205366004612302565b600080935093915050565b60408051928352602083019190915201610101565b61016b610233366004611f2c565b6115b0565b6100f7611677565b60008281526003602052604090205473ffffffffffffffffffffffffffffffffffffffff166102d0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f74207265676973746572656400000000000060448201526064015b60405180910390fd5b60008281526003602052604090205473ffffffffffffffffffffffffffffffffffffffff16331461035d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4e6f742074686520456e64706f696e74206f776e65720000000000000000000060448201526064016102c7565b600091825260036020818152604080852073ffffffffffffffffffffffffffffffffffffffff9094168552929091019052902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610457576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f496e76616c696420507574526573706f6e736528292063616c6c65720000000060448201526064016102c7565b600083815260026020526040902080546104709061232e565b1590506104d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4445425547202d20416c7265616479206578697374730000000000000000000060448201526064016102c7565b600083815260016020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff87161790556002909152902061052982826123cf565b50505050565b60008181526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1633146105bc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4f6e6c79206f776e6572206d617920756e72656769737465720000000000000060448201526064016102c7565b6000818152600360205260408120600281015481547fffffffffffffffffffffffff0000000000000000000000000000000000000000168255916106036001830182611eb5565b5060006002919091015561064e73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163383611a49565b5050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016146106f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f496e76616c696420507574526573706f6e736528292063616c6c65720000000060448201526064016102c7565b600081815260016020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001690556002909152812061073891611eb5565b50565b6000807f0000000000000000000000000000000000000000000000000000000000000000905060008560405160200161077491906124e9565b604051602081830303815290604052805190602001209050600060648610156107f9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f496e73756666696369656e7420726567697374726174696f6e2063726564697460448201526064016102c7565b6040805160048082526024820183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f61f334d30000000000000000000000000000000000000000000000000000000017905291517f8e5dc765000000000000000000000000000000000000000000000000000000008152909160609173ffffffffffffffffffffffffffffffffffffffff871691638e5dc765916108ab9160019130918f91899101612505565b6000604051808303816000875af11580156108ca573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526109109190810190612562565b9093509050828015610923575080516020145b801561093457508051602082012087145b15610a9a576000848152600360205260409020805460029091015473ffffffffffffffffffffffffffffffffffffffff91821691610996907f00000000000000000000000000000000000000000000000000000000000000001633308d611b22565b6109a160648b61261b565b600087815260036020526040812060020180549091906109c2908490612632565b9250508190555060646000808282546109db9190612632565b909155505073ffffffffffffffffffffffffffffffffffffffff821615801590610a055750600081115b15610a4b57610a4b73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168383611a49565b600086815260036020526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001633178155600101610a8e8c826123cf565b50600194505050610a9f565b600092505b7f946870c115e26664fed95ef0844e3f2a093de34963e3a3a573bb793fa889a9488385338c604051610ad4949392919061264a565b60405180910390a1509093505050505b9392505050565b60008281526003602052604090205473ffffffffffffffffffffffffffffffffffffffff16610b76576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f74207265676973746572656400000000000060448201526064016102c7565b60008281526003602052604090205473ffffffffffffffffffffffffffffffffffffffff163314610c03576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4e6f742074686520456e64706f696e74206f776e65720000000000000000000060448201526064016102c7565b80600003610c6d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f496e76616c696420616d6f756e7400000000000000000000000000000000000060448201526064016102c7565b610caf73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333084611b22565b60008281526003602052604081206002018054839290610cd0908490612632565b909155505060408051338152602081018490529081018290527fcdff0c6cd6211a4d6233824ff3b2472af6ff5f3508a06df7404165f42870d31a9060600160405180910390a15050565b60008281526003602052604090205473ffffffffffffffffffffffffffffffffffffffff16610da5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f74207265676973746572656400000000000060448201526064016102c7565b60008281526003602052604090205473ffffffffffffffffffffffffffffffffffffffff163314610e32576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4e6f742074686520456e64706f696e74206f776e65720000000000000000000060448201526064016102c7565b600091825260036020818152604080852073ffffffffffffffffffffffffffffffffffffffff9094168552929091019052902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b600082815260036020526040812060010180546060917f000000000000000000000000000000000000000000000000000000000000000091849190610ed49061232e565b80601f0160208091040260200160405190810160405280929190818152602001828054610f009061232e565b8015610f4d5780601f10610f2257610100808354040283529160200191610f4d565b820191906000526020600020905b815481529060010190602001808311610f3057829003601f168201915b505050505090506000815111610fbf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f456e64706f696e74206e6f74207265676973746572656400000000000000000060448201526064016102c7565b6000868152600360208181526040808420338552909201905290205460ff16611044576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616c6c6572206973206e6f74207065726d697474656400000000000000000060448201526064016102c7565b600086815260036020526040902060020154603211156110c0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f496e73756666696369656e74206372656469742062616c616e6365000000000060448201526064016102c7565b60008681526003602052604081206002018054603292906110e290849061261b565b9250508190555060326000808282546110fb9190612632565b90915550506040517f8e5dc76500000000000000000000000000000000000000000000000000000000815260009060609073ffffffffffffffffffffffffffffffffffffffff851690638e5dc7659061115f90600190339088908d90600401612505565b6000604051808303816000875af115801561117e573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526111c49190810190612562565b90999098509650505050505050565b600060608181333014611242576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f496e76616c696420476574526573706f6e736528292063616c6c65720000000060448201526064016102c7565b60008888888860405160200161125b9493929190612691565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152815160209283012060008181526001845282812054600290945291909120805491935063ffffffff90921691906112bf9061232e565b80601f01602080910402602001604051908101604052809291908181526020018280546112eb9061232e565b80156113385780601f1061130d57610100808354040283529160200191611338565b820191906000526020600020905b81548152906001019060200180831161131b57829003601f168201915b505050505092508063ffffffff166000036113d85760008351116113b8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f476574526573706f6e73653a204d697373696e6720636163686520656e74727960448201526064016102c7565b8251600194506113d26113cd8261138861271e565b611b80565b50611495565b604080518082018252601381527f487962726964436f6d70757465204572726f72000000000000000000000000006020820152905161141b90829060240161275b565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f08c379a0000000000000000000000000000000000000000000000000000000001790529350505b60008281526002602052604081206114ac91611eb5565b50600090815260016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001690559097909650945050505050565b6003602052600090815260409020805460018201805473ffffffffffffffffffffffffffffffffffffffff90921692916115279061232e565b80601f01602080910402602001604051908101604052809291908181526020018280546115539061232e565b80156115a05780601f10611575576101008083540402835291602001916115a0565b820191906000526020600020905b81548152906001019060200180831161158357829003601f168201915b5050505050908060020154905083565b60008281526003602052604081205473ffffffffffffffffffffffffffffffffffffffff1661163b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f74207265676973746572656400000000000060448201526064016102c7565b50600082815260036020818152604080842073ffffffffffffffffffffffffffffffffffffffff86168552909201905290205460ff1692915050565b6040517fdd62ed3e0000000000000000000000000000000000000000000000000000000081523360048201523060248201526000907f0000000000000000000000000000000000000000000000000000000000000000908290606090601e9073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063dd62ed3e90604401602060405180830381865afa158015611734573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611758919061276e565b10156117e6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f496e73756666696369656e742063726564697420617574686f72697a6174696f60448201527f6e0000000000000000000000000000000000000000000000000000000000000060648201526084016102c7565b61182973ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163330601e611b22565b601e60008082825461183b9190612632565b90915550506040517f8e5dc7650000000000000000000000000000000000000000000000000000000081526201000160048201523360248201526080604482015260006084820181905260a0606483015260a482015273ffffffffffffffffffffffffffffffffffffffff841690638e5dc7659060c4016000604051808303816000875af11580156118d1573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526119179190810190612562565b909250905081611983576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f4f7065726174696f6e206661696c65640000000000000000000000000000000060448201526064016102c7565b80516020146119ee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f42616420726573706f6e7365206c656e6774680000000000000000000000000060448201526064016102c7565b600081806020019051810190611a04919061276e565b6040805160008152602081018390529192507f450d62889c3a6e19c9586840ce9c21040b90d81950fe31f2ba982090adaf53e8910160405180910390a1949350505050565b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052611b1d9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152611ba9565b505050565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526105299085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611a9b565b6000805a90505b825a611b93908361261b565b1015611b1d57611ba282612787565b9150611b87565b6000611c0b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16611cb59092919063ffffffff16565b805190915015611b1d5780806020019051810190611c2991906127bf565b611b1d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016102c7565b6060611cc48484600085611ccc565b949350505050565b606082471015611d5e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016102c7565b73ffffffffffffffffffffffffffffffffffffffff85163b611ddc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016102c7565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051611e0591906124e9565b60006040518083038185875af1925050503d8060008114611e42576040519150601f19603f3d011682016040523d82523d6000602084013e611e47565b606091505b5091509150611e57828286611e62565b979650505050505050565b60608315611e71575081610ae4565b825115611e815782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102c7919061275b565b508054611ec19061232e565b6000825580601f10611ed1575050565b601f01602090049060005260206000209081019061073891905b80821115611eff5760008155600101611eeb565b5090565b803573ffffffffffffffffffffffffffffffffffffffff81168114611f2757600080fd5b919050565b60008060408385031215611f3f57600080fd5b82359150611f4f60208401611f03565b90509250929050565b803563ffffffff81168114611f2757600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611fe257611fe2611f6c565b604052919050565b600067ffffffffffffffff82111561200457612004611f6c565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f83011261204157600080fd5b813561205461204f82611fea565b611f9b565b81815284602083860101111561206957600080fd5b816020850160208301376000918101602001919091529392505050565b60008060006060848603121561209b57600080fd5b833592506120ab60208501611f58565b9150604084013567ffffffffffffffff8111156120c757600080fd5b6120d386828701612030565b9150509250925092565b6000602082840312156120ef57600080fd5b5035919050565b60008060006060848603121561210b57600080fd5b833567ffffffffffffffff81111561212257600080fd5b61212e86828701612030565b9660208601359650604090950135949350505050565b6000806040838503121561215757600080fd5b50508035926020909101359150565b6000806040838503121561217957600080fd5b82359150602083013567ffffffffffffffff81111561219757600080fd5b6121a385828601612030565b9150509250929050565b60005b838110156121c85781810151838201526020016121b0565b838111156105295750506000910152565b600081518084526121f18160208601602086016121ad565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b8215158152604060208201526000611cc460408301846121d9565b6000806000806080858703121561225457600080fd5b61225d85611f58565b935061226b60208601611f03565b9250604085013567ffffffffffffffff8082111561228857600080fd5b61229488838901612030565b935060608701359150808211156122aa57600080fd5b506122b787828801612030565b91505092959194509250565b73ffffffffffffffffffffffffffffffffffffffff841681526060602082015260006122f260608301856121d9565b9050826040830152949350505050565b60008060006060848603121561231757600080fd5b505081359360208301359350604090920135919050565b600181811c9082168061234257607f821691505b60208210810361237b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b601f821115611b1d57600081815260208120601f850160051c810160208610156123a85750805b601f850160051c820191505b818110156123c7578281556001016123b4565b505050505050565b815167ffffffffffffffff8111156123e9576123e9611f6c565b6123fd816123f7845461232e565b84612381565b602080601f831160018114612450576000841561241a5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556123c7565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b8281101561249d5788860151825594840194600190910190840161247e565b50858210156124d957878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b600082516124fb8184602087016121ad565b9190910192915050565b63ffffffff8516815273ffffffffffffffffffffffffffffffffffffffff8416602082015260806040820152600061254060808301856121d9565b8281036060840152611e5781856121d9565b80518015158114611f2757600080fd5b6000806040838503121561257557600080fd5b61257e83612552565b9150602083015167ffffffffffffffff81111561259a57600080fd5b8301601f810185136125ab57600080fd5b80516125b961204f82611fea565b8181528660208385010111156125ce57600080fd5b6125df8260208301602086016121ad565b8093505050509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561262d5761262d6125ec565b500390565b60008219821115612645576126456125ec565b500190565b841515815283602082015273ffffffffffffffffffffffffffffffffffffffff8316604082015260806060820152600061268760808301846121d9565b9695505050505050565b7fffffffff000000000000000000000000000000000000000000000000000000008560e01b1681527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008460601b166004820152600083516126f98160188501602088016121ad565b8351908301906127108160188401602088016121ad565b016018019695505050505050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615612756576127566125ec565b500290565b602081526000610ae460208301846121d9565b60006020828403121561278057600080fd5b5051919050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036127b8576127b86125ec565b5060010190565b6000602082840312156127d157600080fd5b610ae48261255256fea164736f6c634300080f000a",
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
// Solidity: function Endpoints(bytes32 ) view returns(address Owner, string URL, uint256 CreditBalance)
func (_BobaHCHelper *BobaHCHelperCaller) Endpoints(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Owner         common.Address
	URL           string
	CreditBalance *big.Int
}, error) {
	var out []interface{}
	err := _BobaHCHelper.contract.Call(opts, &out, "Endpoints", arg0)

	outstruct := new(struct {
		Owner         common.Address
		URL           string
		CreditBalance *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.URL = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.CreditBalance = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Endpoints is a free data retrieval call binding the contract method 0x997e31cc.
//
// Solidity: function Endpoints(bytes32 ) view returns(address Owner, string URL, uint256 CreditBalance)
func (_BobaHCHelper *BobaHCHelperSession) Endpoints(arg0 [32]byte) (struct {
	Owner         common.Address
	URL           string
	CreditBalance *big.Int
}, error) {
	return _BobaHCHelper.Contract.Endpoints(&_BobaHCHelper.CallOpts, arg0)
}

// Endpoints is a free data retrieval call binding the contract method 0x997e31cc.
//
// Solidity: function Endpoints(bytes32 ) view returns(address Owner, string URL, uint256 CreditBalance)
func (_BobaHCHelper *BobaHCHelperCallerSession) Endpoints(arg0 [32]byte) (struct {
	Owner         common.Address
	URL           string
	CreditBalance *big.Int
}, error) {
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

// NextRandom is a paid mutator transaction binding the contract method 0xade42320.
//
// Solidity: function NextRandom(bytes32 session, bytes32 nextHash, uint256 myNum) returns(bytes32, uint256)
func (_BobaHCHelper *BobaHCHelperTransactor) NextRandom(opts *bind.TransactOpts, session [32]byte, nextHash [32]byte, myNum *big.Int) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "NextRandom", session, nextHash, myNum)
}

// NextRandom is a paid mutator transaction binding the contract method 0xade42320.
//
// Solidity: function NextRandom(bytes32 session, bytes32 nextHash, uint256 myNum) returns(bytes32, uint256)
func (_BobaHCHelper *BobaHCHelperSession) NextRandom(session [32]byte, nextHash [32]byte, myNum *big.Int) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.NextRandom(&_BobaHCHelper.TransactOpts, session, nextHash, myNum)
}

// NextRandom is a paid mutator transaction binding the contract method 0xade42320.
//
// Solidity: function NextRandom(bytes32 session, bytes32 nextHash, uint256 myNum) returns(bytes32, uint256)
func (_BobaHCHelper *BobaHCHelperTransactorSession) NextRandom(session [32]byte, nextHash [32]byte, myNum *big.Int) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.NextRandom(&_BobaHCHelper.TransactOpts, session, nextHash, myNum)
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

// AddBalanceTo is a paid mutator transaction binding the contract method 0x6c6459c8.
//
// Solidity: function addBalanceTo(bytes32 _EK, uint256 _addBalanceAmount) returns()
func (_BobaHCHelper *BobaHCHelperTransactor) AddBalanceTo(opts *bind.TransactOpts, _EK [32]byte, _addBalanceAmount *big.Int) (*types.Transaction, error) {
	return _BobaHCHelper.contract.Transact(opts, "addBalanceTo", _EK, _addBalanceAmount)
}

// AddBalanceTo is a paid mutator transaction binding the contract method 0x6c6459c8.
//
// Solidity: function addBalanceTo(bytes32 _EK, uint256 _addBalanceAmount) returns()
func (_BobaHCHelper *BobaHCHelperSession) AddBalanceTo(_EK [32]byte, _addBalanceAmount *big.Int) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.AddBalanceTo(&_BobaHCHelper.TransactOpts, _EK, _addBalanceAmount)
}

// AddBalanceTo is a paid mutator transaction binding the contract method 0x6c6459c8.
//
// Solidity: function addBalanceTo(bytes32 _EK, uint256 _addBalanceAmount) returns()
func (_BobaHCHelper *BobaHCHelperTransactorSession) AddBalanceTo(_EK [32]byte, _addBalanceAmount *big.Int) (*types.Transaction, error) {
	return _BobaHCHelper.Contract.AddBalanceTo(&_BobaHCHelper.TransactOpts, _EK, _addBalanceAmount)
}

// BobaHCHelperAddBalanceToIterator is returned from FilterAddBalanceTo and is used to iterate over the raw logs and unpacked data for AddBalanceTo events raised by the BobaHCHelper contract.
type BobaHCHelperAddBalanceToIterator struct {
	Event *BobaHCHelperAddBalanceTo // Event containing the contract specifics and raw log

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
func (it *BobaHCHelperAddBalanceToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BobaHCHelperAddBalanceTo)
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
		it.Event = new(BobaHCHelperAddBalanceTo)
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
func (it *BobaHCHelperAddBalanceToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BobaHCHelperAddBalanceToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BobaHCHelperAddBalanceTo represents a AddBalanceTo event raised by the BobaHCHelper contract.
type BobaHCHelperAddBalanceTo struct {
	Arg0 common.Address
	Arg1 [32]byte
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddBalanceTo is a free log retrieval operation binding the contract event 0xcdff0c6cd6211a4d6233824ff3b2472af6ff5f3508a06df7404165f42870d31a.
//
// Solidity: event AddBalanceTo(address arg0, bytes32 arg1, uint256 arg2)
func (_BobaHCHelper *BobaHCHelperFilterer) FilterAddBalanceTo(opts *bind.FilterOpts) (*BobaHCHelperAddBalanceToIterator, error) {

	logs, sub, err := _BobaHCHelper.contract.FilterLogs(opts, "AddBalanceTo")
	if err != nil {
		return nil, err
	}
	return &BobaHCHelperAddBalanceToIterator{contract: _BobaHCHelper.contract, event: "AddBalanceTo", logs: logs, sub: sub}, nil
}

// WatchAddBalanceTo is a free log subscription operation binding the contract event 0xcdff0c6cd6211a4d6233824ff3b2472af6ff5f3508a06df7404165f42870d31a.
//
// Solidity: event AddBalanceTo(address arg0, bytes32 arg1, uint256 arg2)
func (_BobaHCHelper *BobaHCHelperFilterer) WatchAddBalanceTo(opts *bind.WatchOpts, sink chan<- *BobaHCHelperAddBalanceTo) (event.Subscription, error) {

	logs, sub, err := _BobaHCHelper.contract.WatchLogs(opts, "AddBalanceTo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BobaHCHelperAddBalanceTo)
				if err := _BobaHCHelper.contract.UnpackLog(event, "AddBalanceTo", log); err != nil {
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

// ParseAddBalanceTo is a log parse operation binding the contract event 0xcdff0c6cd6211a4d6233824ff3b2472af6ff5f3508a06df7404165f42870d31a.
//
// Solidity: event AddBalanceTo(address arg0, bytes32 arg1, uint256 arg2)
func (_BobaHCHelper *BobaHCHelperFilterer) ParseAddBalanceTo(log types.Log) (*BobaHCHelperAddBalanceTo, error) {
	event := new(BobaHCHelperAddBalanceTo)
	if err := _BobaHCHelper.contract.UnpackLog(event, "AddBalanceTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
