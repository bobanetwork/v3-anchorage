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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"DBG\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"URL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"EndpointRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"URL\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"EndpointUnregistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"AddCredit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_callerAddress\",\"type\":\"address\"}],\"name\":\"AddPermittedCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"CallOffchain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_callerAddress\",\"type\":\"address\"}],\"name\":\"CheckPermittedCaller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Endpoints\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_cacheKey\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_returndata\",\"type\":\"bytes\"}],\"name\":\"PutResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_auth\",\"type\":\"bytes32\"}],\"name\":\"RegisterEndpoint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_callerAddress\",\"type\":\"address\"}],\"name\":\"RemovePermittedCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"session\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"myNum\",\"type\":\"uint256\"}],\"name\":\"SequentialRandom\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SimpleRandom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"UnregisterEndpoint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingCharge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"prepaidCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040527342000000000000000000000000000000000000fd60805273deaddeaddeaddeaddeaddeaddeaddeaddead990160a0527342000000000000000000000000000000000000fe60c05234801561005857600080fd5b5060805160a05160c05161213c61008f600039600081816109b50152610aa80152600061110201526000610a20015261213c6000f3fe608060405234801561001057600080fd5b50600436106100e95760003560e01c806387fc1b251161008c578063c1fd7e4611610066578063c1fd7e461461023e578063c6e4b0d31461025f578063e2f9479214610267578063eb6598b51461027a57600080fd5b806387fc1b25146101ad578063997e31cc146101c0578063a71adb3e1461021b57600080fd5b806330429027116100c8578063304290271461013f57806332be428f1461015257806362fccf6e1461017a5780637e1214131461018d57600080fd5b8062292526146100ee5780631730f39b1461010a578063248eaf621461011f575b600080fd5b6100f760005481565b6040519081526020015b60405180910390f35b61011d610118366004611952565b61028d565b005b6100f761012d3660046119a6565b60026020526000908152604090205481565b61011d61014d366004611952565b610458565b6101656101603660046119c1565b610619565b60408051928352602083019190915201610101565b61011d6101883660046119ed565b610758565b6100f761019b3660046119a6565b60016020526000908152604090205481565b61011d6101bb366004611a2f565b610931565b6101f66101ce366004611a59565b60036020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610101565b61022e610229366004611a72565b610a1b565b6040519015158152602001610101565b61025161024c366004611abe565b610cb2565b604051610101929190611ba0565b6100f7610eb0565b61022e610275366004611952565b610fd3565b61011d610288366004611bc9565b6110ea565b600083836040516020016102a2929190611c19565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600390935291205490915073ffffffffffffffffffffffffffffffffffffffff16610369576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f74207265676973746572656400000000000060448201526064015b60405180910390fd5b60008181526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1633146103f6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4e6f742074686520456e64706f696e74206f776e6572000000000000000000006044820152606401610360565b600090815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff90941683526001938401909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690911790555050565b6000838360405160200161046d929190611c19565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600390935291205490915073ffffffffffffffffffffffffffffffffffffffff1661052f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f7420726567697374657265640000000000006044820152606401610360565b60008181526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1633146105bc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4e6f742074686520456e64706f696e74206f776e6572000000000000000000006044820152606401610360565b600090815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff9094168352600190930190522080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555050565b6040517f32be428f00000000000000000000000000000000000000000000000000000000602082015233606090811b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602483015260388201859052605882018490526078820183905260009182916332be428f9183919082906098016040516020818303038152906040528051906020012090506106bb8133601e61126b565b90935091508280156106ce575081516040145b610734576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f52616e646f6d206e756d6265722067656e65726174696f6e206661696c6564006044820152606401610360565b818060200190518101906107489190611c29565b9550955050505050935093915050565b6000828260405160200161076d929190611c19565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600390935291205490915073ffffffffffffffffffffffffffffffffffffffff1661082f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f7420726567697374657265640000000000006044820152606401610360565b60008181526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1633146108bc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4e6f742074686520456e64706f696e74206f776e6572000000000000000000006044820152606401610360565b6000818152600360205260409081902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055517fddcfd504e796f664dc8ed3395209f4798379b47d8ee42d4ee5142a39f6d6d21a9061092490859085903390611c96565b60405180910390a1505050565b6000811161099b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f496e76616c696420616d6f756e740000000000000000000000000000000000006044820152606401610360565b6109dd73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163330846114de565b73ffffffffffffffffffffffffffffffffffffffff821660009081526001602052604081208054839290610a12908490611cff565b90915550505050565b6000807f0000000000000000000000000000000000000000000000000000000000000000905060008585604051602001610a56929190611c19565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209050610ad173ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333060646114de565b6064600080828254610ae39190611cff565b90915550506040805160048082526024820183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f61f334d30000000000000000000000000000000000000000000000000000000017905291517fc1fd7e46000000000000000000000000000000000000000000000000000000008152909160609160009173ffffffffffffffffffffffffffffffffffffffff87169163c1fd7e4691610b9a918d918d91899101611d17565b6000604051808303816000875af1158015610bb9573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610bff9190810190611d6c565b92509050808015610c11575081516020145b8015610c2257508151602083012087145b15610ca0576000848152600360205260409081902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000163390811790915590517fd48d01edeb78cd5b560a64aebf74c66df4140fc9b5f72bff4ec03f43e7e9f88291610c93918c918c91611c96565b60405180910390a1610ca4565b5060005b9450505050505b9392505050565b60006060600063c1fd7e46905060008787604051602001610cd4929190611c19565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209050333014610e245760008181526003602052604090205473ffffffffffffffffffffffffffffffffffffffff16610d9e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f456e64706f696e74206e6f7420726567697374657265640000000000000000006044820152606401610360565b600081815260036020908152604080832033845260010190915290205460ff16610e24576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616c6c6572206973206e6f74207065726d69747465640000000000000000006044820152606401610360565b600082338a8a604051602001610e3b929190611c19565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052610e7c9392918b908b90602001611e4d565b604051602081830303815290604052805190602001209050610ea08133603261126b565b9450945050505094509492505050565b6040517fc6e4b0d300000000000000000000000000000000000000000000000000000000602082015233606090811b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602483015260009163c6e4b0d39183918290603801604051602081830303815290604052805190602001209050610f3a8133601e61126b565b9093509150828015610f4d575081516020145b610fb3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f52616e646f6d206e756d6265722067656e65726174696f6e206661696c6564006044820152606401610360565b600082806020019051810190610fc99190611ed3565b9695505050505050565b6000808484604051602001610fe9929190611c19565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600390935291205490915073ffffffffffffffffffffffffffffffffffffffff166110ab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f7420726567697374657265640000000000006044820152606401610360565b600090815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616845260010190915290205460ff1690509392505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614611189576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f496e76616c696420507574526573706f6e736528292063616c6c6572000000006044820152606401610360565b600084815260046020526040902080546111a290611eec565b15905061120b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4445425547202d20416c726561647920657869737473000000000000000000006044820152606401610360565b82828260405160200161122093929190611f3f565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526000868152600460205220906112649082611fb2565b5050505050565b6000606073ffffffffffffffffffffffffffffffffffffffff841630146113a25773ffffffffffffffffffffffffffffffffffffffff84166000908152600260205260408120546112bc9085611cff565b73ffffffffffffffffffffffffffffffffffffffff861660009081526001602052604090205490915081111561134e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f496e73756666696369656e7420637265646974000000000000000000000000006044820152606401610360565b73ffffffffffffffffffffffffffffffffffffffff8516600090815260016020526040812080548692906113839084906120cc565b925050819055508360008082825461139b9190611cff565b9091555050505b600085815260046020526040812080546113bb90611eec565b80601f01602080910402602001604051908101604052809291908181526020018280546113e790611eec565b80156114345780601f1061140957610100808354040283529160200191611434565b820191906000526020600020905b81548152906001019060200180831161141757829003601f168201915b5050505050905060008151116114a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f476574526573706f6e73653a204d697373696e6720636163686520656e7472796044820152606401610360565b60008681526004602052604081206114bd9161188a565b808060200190518101906114d19190611d6c565b9250925050935093915050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052611573908590611579565b50505050565b60006115db826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff1661168a9092919063ffffffff16565b80519091501561168557808060200190518101906115f991906120e3565b611685576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610360565b505050565b606061169984846000856116a1565b949350505050565b606082471015611733576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610360565b73ffffffffffffffffffffffffffffffffffffffff85163b6117b1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610360565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516117da9190612100565b60006040518083038185875af1925050503d8060008114611817576040519150601f19603f3d011682016040523d82523d6000602084013e61181c565b606091505b509150915061182c828286611837565b979650505050505050565b60608315611846575081610cab565b8251156118565782518084602001fd5b816040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610360919061211c565b50805461189690611eec565b6000825580601f106118a6575050565b601f0160209004906000526020600020908101906118c491906118c7565b50565b5b808211156118dc57600081556001016118c8565b5090565b60008083601f8401126118f257600080fd5b50813567ffffffffffffffff81111561190a57600080fd5b60208301915083602082850101111561192257600080fd5b9250929050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461194d57600080fd5b919050565b60008060006040848603121561196757600080fd5b833567ffffffffffffffff81111561197e57600080fd5b61198a868287016118e0565b909450925061199d905060208501611929565b90509250925092565b6000602082840312156119b857600080fd5b610cab82611929565b6000806000606084860312156119d657600080fd5b505081359360208301359350604090920135919050565b60008060208385031215611a0057600080fd5b823567ffffffffffffffff811115611a1757600080fd5b611a23858286016118e0565b90969095509350505050565b60008060408385031215611a4257600080fd5b611a4b83611929565b946020939093013593505050565b600060208284031215611a6b57600080fd5b5035919050565b600080600060408486031215611a8757600080fd5b833567ffffffffffffffff811115611a9e57600080fd5b611aaa868287016118e0565b909790965060209590950135949350505050565b60008060008060408587031215611ad457600080fd5b843567ffffffffffffffff80821115611aec57600080fd5b611af8888389016118e0565b90965094506020870135915080821115611b1157600080fd5b50611b1e878288016118e0565b95989497509550505050565b60005b83811015611b45578181015183820152602001611b2d565b838111156115735750506000910152565b60008151808452611b6e816020860160208601611b2a565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b82151581526040602082015260006116996040830184611b56565b80151581146118c457600080fd5b60008060008060608587031215611bdf57600080fd5b843593506020850135611bf181611bbb565b9250604085013567ffffffffffffffff811115611c0d57600080fd5b611b1e878288016118e0565b8183823760009101908152919050565b60008060408385031215611c3c57600080fd5b505080516020909101519092909150565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000611caa604083018587611c4d565b905073ffffffffffffffffffffffffffffffffffffffff83166020830152949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115611d1257611d12611cd0565b500190565b604081526000611d2b604083018587611c4d565b8281036020840152610fc98185611b56565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060408385031215611d7f57600080fd5b8251611d8a81611bbb565b602084015190925067ffffffffffffffff80821115611da857600080fd5b818501915085601f830112611dbc57600080fd5b815181811115611dce57611dce611d3d565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611e1457611e14611d3d565b81604052828152886020848701011115611e2d57600080fd5b611e3e836020830160208801611b2a565b80955050505050509250929050565b7fffffffff000000000000000000000000000000000000000000000000000000008660e01b1681527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008560601b16600482015260008451611eb5816018850160208901611b2a565b82018385601883013760009301601801928352509095945050505050565b600060208284031215611ee557600080fd5b5051919050565b600181811c90821680611f0057607f821691505b602082108103611f39577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8315158152604060208201526000611f5b604083018486611c4d565b95945050505050565b601f82111561168557600081815260208120601f850160051c81016020861015611f8b5750805b601f850160051c820191505b81811015611faa57828155600101611f97565b505050505050565b815167ffffffffffffffff811115611fcc57611fcc611d3d565b611fe081611fda8454611eec565b84611f64565b602080601f8311600181146120335760008415611ffd5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611faa565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b8281101561208057888601518255948401946001909101908401612061565b50858210156120bc57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b6000828210156120de576120de611cd0565b500390565b6000602082840312156120f557600080fd5b8151610cab81611bbb565b60008251612112818460208701611b2a565b9190910192915050565b602081526000610cab6020830184611b5656fea164736f6c634300080f000a",
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
