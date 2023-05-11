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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"DBG\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"AddCredit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_callerAddress\",\"type\":\"address\"}],\"name\":\"AddPermittedCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"CallOffchain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_callerAddress\",\"type\":\"address\"}],\"name\":\"CheckPermittedCaller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Endpoints\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_cacheKey\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_returndata\",\"type\":\"bytes\"}],\"name\":\"PutResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_auth\",\"type\":\"bytes32\"}],\"name\":\"RegisterEndpoint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_callerAddress\",\"type\":\"address\"}],\"name\":\"RemovePermittedCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"session\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"myNum\",\"type\":\"uint256\"}],\"name\":\"SequentialRandom\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SimpleRandom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerRevenue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingCharge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"prepaidCredit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040527342000000000000000000000000000000000000fd60805273deaddeaddeaddeaddeaddeaddeaddeaddead990160a0527342000000000000000000000000000000000000fe60c05234801561005857600080fd5b5060805160a05160c051611e8b61008f600039600081816107be01526108b101526000610ecd015260006108290152611e8b6000f3fe608060405234801561001057600080fd5b50600436106100de5760003560e01c806387fc1b251161008c578063c1fd7e4611610066578063c1fd7e4614610220578063c6e4b0d314610241578063e2f9479214610249578063eb6598b51461025c57600080fd5b806387fc1b251461018f578063997e31cc146101a2578063a71adb3e146101fd57600080fd5b806330429027116100bd578063304290271461013457806332be428f146101475780637e1214131461016f57600080fd5b8062292526146100e35780631730f39b146100ff578063248eaf6214610114575b600080fd5b6100ec60005481565b6040519081526020015b60405180910390f35b61011261010d36600461171d565b61026f565b005b6100ec610122366004611771565b60026020526000908152604090205481565b61011261014236600461171d565b61043a565b61015a61015536600461178c565b6105fb565b604080519283526020830191909152016100f6565b6100ec61017d366004611771565b60016020526000908152604090205481565b61011261019d3660046117b8565b61073a565b6101d86101b03660046117e2565b60036020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100f6565b61021061020b3660046117fb565b610824565b60405190151581526020016100f6565b61023361022e366004611847565b610a7d565b6040516100f6929190611929565b6100ec610c7b565b61021061025736600461171d565b610d9e565b61011261026a366004611952565b610eb5565b600083836040516020016102849291906119a2565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600390935291205490915073ffffffffffffffffffffffffffffffffffffffff1661034b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f74207265676973746572656400000000000060448201526064015b60405180910390fd5b60008181526003602052604090205473ffffffffffffffffffffffffffffffffffffffff1633146103d8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4e6f742074686520456e64706f696e74206f776e6572000000000000000000006044820152606401610342565b600090815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff90941683526001938401909152902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690911790555050565b6000838360405160200161044f9291906119a2565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600390935291205490915073ffffffffffffffffffffffffffffffffffffffff16610511576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f7420726567697374657265640000000000006044820152606401610342565b60008181526003602052604090205473ffffffffffffffffffffffffffffffffffffffff16331461059e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4e6f742074686520456e64706f696e74206f776e6572000000000000000000006044820152606401610342565b600090815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff9094168352600190930190522080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690555050565b6040517f32be428f00000000000000000000000000000000000000000000000000000000602082015233606090811b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602483015260388201859052605882018490526078820183905260009182916332be428f91839190829060980160405160208183030381529060405280519060200120905061069d8133601e611036565b90935091508280156106b0575081516040145b610716576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f52616e646f6d206e756d6265722067656e65726174696f6e206661696c6564006044820152606401610342565b8180602001905181019061072a91906119b2565b9550955050505050935093915050565b600081116107a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f496e76616c696420616d6f756e740000000000000000000000000000000000006044820152606401610342565b6107e673ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163330846112a9565b73ffffffffffffffffffffffffffffffffffffffff82166000908152600160205260408120805483929061081b908490611a05565b90915550505050565b6000807f000000000000000000000000000000000000000000000000000000000000000090506000858560405160200161085f9291906119a2565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052805160209091012090506108da73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016333060646112a9565b60646000808282546108ec9190611a05565b90915550506040805160048082526024820183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f61f334d30000000000000000000000000000000000000000000000000000000017905291517fc1fd7e46000000000000000000000000000000000000000000000000000000008152909160609160009173ffffffffffffffffffffffffffffffffffffffff87169163c1fd7e46916109a3918d918d91899101611a66565b6000604051808303816000875af11580156109c2573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610a089190810190611abb565b92509050808015610a1a575081516020145b8015610a2b57508151602083012087145b15610a6b57600084815260036020526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001633179055610a6f565b5060005b9450505050505b9392505050565b60006060600063c1fd7e46905060008787604051602001610a9f9291906119a2565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905280516020909101209050333014610bef5760008181526003602052604090205473ffffffffffffffffffffffffffffffffffffffff16610b69576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f456e64706f696e74206e6f7420726567697374657265640000000000000000006044820152606401610342565b600081815260036020908152604080832033845260010190915290205460ff16610bef576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616c6c6572206973206e6f74207065726d69747465640000000000000000006044820152606401610342565b600082338a8a604051602001610c069291906119a2565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052610c479392918b908b90602001611b9c565b604051602081830303815290604052805190602001209050610c6b81336032611036565b9450945050505094509492505050565b6040517fc6e4b0d300000000000000000000000000000000000000000000000000000000602082015233606090811b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016602483015260009163c6e4b0d39183918290603801604051602081830303815290604052805190602001209050610d058133601e611036565b9093509150828015610d18575081516020145b610d7e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f52616e646f6d206e756d6265722067656e65726174696f6e206661696c6564006044820152606401610342565b600082806020019051810190610d949190611c22565b9695505050505050565b6000808484604051602001610db49291906119a2565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815281516020928301206000818152600390935291205490915073ffffffffffffffffffffffffffffffffffffffff16610e76576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f456e64706f696e74206973206e6f7420726567697374657265640000000000006044820152606401610342565b600090815260036020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616845260010190915290205460ff1690509392505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610f54576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f496e76616c696420507574526573706f6e736528292063616c6c6572000000006044820152606401610342565b60008481526004602052604090208054610f6d90611c3b565b159050610fd6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4445425547202d20416c726561647920657869737473000000000000000000006044820152606401610342565b828282604051602001610feb93929190611c8e565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291815260008681526004602052209061102f9082611d01565b5050505050565b6000606073ffffffffffffffffffffffffffffffffffffffff8416301461116d5773ffffffffffffffffffffffffffffffffffffffff84166000908152600260205260408120546110879085611a05565b73ffffffffffffffffffffffffffffffffffffffff8616600090815260016020526040902054909150811115611119576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f496e73756666696369656e7420637265646974000000000000000000000000006044820152606401610342565b73ffffffffffffffffffffffffffffffffffffffff85166000908152600160205260408120805486929061114e908490611e1b565b92505081905550836000808282546111669190611a05565b9091555050505b6000858152600460205260408120805461118690611c3b565b80601f01602080910402602001604051908101604052809291908181526020018280546111b290611c3b565b80156111ff5780601f106111d4576101008083540402835291602001916111ff565b820191906000526020600020905b8154815290600101906020018083116111e257829003601f168201915b505050505090506000815111611271576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f476574526573706f6e73653a204d697373696e6720636163686520656e7472796044820152606401610342565b600086815260046020526040812061128891611655565b8080602001905181019061129c9190611abb565b9250925050935093915050565b6040805173ffffffffffffffffffffffffffffffffffffffff85811660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd0000000000000000000000000000000000000000000000000000000017905261133e908590611344565b50505050565b60006113a6826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166114559092919063ffffffff16565b80519091501561145057808060200190518101906113c49190611e32565b611450576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f742073756363656564000000000000000000000000000000000000000000006064820152608401610342565b505050565b6060611464848460008561146c565b949350505050565b6060824710156114fe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610342565b73ffffffffffffffffffffffffffffffffffffffff85163b61157c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610342565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516115a59190611e4f565b60006040518083038185875af1925050503d80600081146115e2576040519150601f19603f3d011682016040523d82523d6000602084013e6115e7565b606091505b50915091506115f7828286611602565b979650505050505050565b60608315611611575081610a76565b8251156116215782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103429190611e6b565b50805461166190611c3b565b6000825580601f10611671575050565b601f01602090049060005260206000209081019061168f9190611692565b50565b5b808211156116a75760008155600101611693565b5090565b60008083601f8401126116bd57600080fd5b50813567ffffffffffffffff8111156116d557600080fd5b6020830191508360208285010111156116ed57600080fd5b9250929050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461171857600080fd5b919050565b60008060006040848603121561173257600080fd5b833567ffffffffffffffff81111561174957600080fd5b611755868287016116ab565b90945092506117689050602085016116f4565b90509250925092565b60006020828403121561178357600080fd5b610a76826116f4565b6000806000606084860312156117a157600080fd5b505081359360208301359350604090920135919050565b600080604083850312156117cb57600080fd5b6117d4836116f4565b946020939093013593505050565b6000602082840312156117f457600080fd5b5035919050565b60008060006040848603121561181057600080fd5b833567ffffffffffffffff81111561182757600080fd5b611833868287016116ab565b909790965060209590950135949350505050565b6000806000806040858703121561185d57600080fd5b843567ffffffffffffffff8082111561187557600080fd5b611881888389016116ab565b9096509450602087013591508082111561189a57600080fd5b506118a7878288016116ab565b95989497509550505050565b60005b838110156118ce5781810151838201526020016118b6565b8381111561133e5750506000910152565b600081518084526118f78160208601602086016118b3565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b821515815260406020820152600061146460408301846118df565b801515811461168f57600080fd5b6000806000806060858703121561196857600080fd5b84359350602085013561197a81611944565b9250604085013567ffffffffffffffff81111561199657600080fd5b6118a7878288016116ab565b8183823760009101908152919050565b600080604083850312156119c557600080fd5b505080516020909101519092909150565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008219821115611a1857611a186119d6565b500190565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b604081526000611a7a604083018587611a1d565b8281036020840152610d9481856118df565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060408385031215611ace57600080fd5b8251611ad981611944565b602084015190925067ffffffffffffffff80821115611af757600080fd5b818501915085601f830112611b0b57600080fd5b815181811115611b1d57611b1d611a8c565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611b6357611b63611a8c565b81604052828152886020848701011115611b7c57600080fd5b611b8d8360208301602088016118b3565b80955050505050509250929050565b7fffffffff000000000000000000000000000000000000000000000000000000008660e01b1681527fffffffffffffffffffffffffffffffffffffffff0000000000000000000000008560601b16600482015260008451611c048160188501602089016118b3565b82018385601883013760009301601801928352509095945050505050565b600060208284031215611c3457600080fd5b5051919050565b600181811c90821680611c4f57607f821691505b602082108103611c88577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b8315158152604060208201526000611caa604083018486611a1d565b95945050505050565b601f82111561145057600081815260208120601f850160051c81016020861015611cda5750805b601f850160051c820191505b81811015611cf957828155600101611ce6565b505050505050565b815167ffffffffffffffff811115611d1b57611d1b611a8c565b611d2f81611d298454611c3b565b84611cb3565b602080601f831160018114611d825760008415611d4c5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611cf9565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015611dcf57888601518255948401946001909101908401611db0565b5085821015611e0b57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b600082821015611e2d57611e2d6119d6565b500390565b600060208284031215611e4457600080fd5b8151610a7681611944565b60008251611e618184602087016118b3565b9190910192915050565b602081526000610a7660208301846118df56fea164736f6c634300080f000a",
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
