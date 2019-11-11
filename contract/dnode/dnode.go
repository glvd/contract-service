// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dnode

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// DNodeABI is the input ABI used to generate the binding from.
const DNodeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastNode\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"setLast\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"ver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"setLastVersion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"setVersion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"swarm\",\"type\":\"string\"}],\"name\":\"store\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DNodeFuncSigs maps the 4-byte function signature to its string representation.
var DNodeFuncSigs = map[string]string{
	"4d622831": "getLast()",
	"d1d72e3e": "getLastNode()",
	"558614c7": "getLastVersion()",
	"4f0f4aa9": "getNode(uint256)",
	"b88da759": "getVersion(uint256)",
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"4cc82215": "remove(uint256)",
	"715018a6": "renounceOwnership()",
	"5c57bb89": "setLast(uint256)",
	"e80d8c38": "setLastVersion(string,string)",
	"69cf7c9a": "setVersion(uint256,string,string)",
	"131a0680": "store(string)",
	"f2fde38b": "transferOwnership(address)",
}

// DNodeBin is the compiled bytecode used for deploying new contracts.
var DNodeBin = "0x608060405260006100176001600160e01b0361006616565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006a565b3390565b610f9f806100796000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063715018a61161008c578063b88da75911610066578063b88da759146101c4578063d1d72e3e146101d7578063e80d8c38146101df578063f2fde38b146101f2576100ea565b8063715018a6146101925780638da5cb5b1461019a5780638f32d59b146101af576100ea565b80634f0f4aa9116100c85780634f0f4aa914610135578063558614c7146101565780635c57bb891461016c57806369cf7c9a1461017f576100ea565b8063131a0680146100ef5780634cc82215146101045780634d62283114610117575b600080fd5b6101026100fd366004610a67565b610205565b005b610102610112366004610b05565b610280565b61011f610349565b60405161012c9190610e6d565b60405180910390f35b610148610143366004610b05565b610350565b60405161012c929190610dca565b61015e61047a565b60405161012c929190610df8565b61010261017a366004610b05565b6104b5565b61010261018d366004610b23565b6104ff565b61010261059b565b6101a2610609565b60405161012c9190610dbc565b6101b7610618565b60405161012c9190610dea565b61015e6101d2366004610b05565b61063c565b61014861077a565b6101026101ed366004610a9c565b610789565b610102610200366004610a41565b61080f565b61020d610618565b6102325760405162461bcd60e51b815260040161022990610e3d565b60405180910390fd5b60025460009081526001602081815260408084209283015484529181529120825161025f928401906108c4565b50506002546000908152600160208190526040909120810180549091019055565b610288610618565b6102a45760405162461bcd60e51b815260040161022990610e3d565b6002546000908152600160208190526040909120015481106102d85760405162461bcd60e51b815260040161022990610e2d565b6002805460009081526001602081815260408084208084015460001980820187529190935281852087865291909420815492956103279591949293928316156101000290910190911604610942565b5050600254600090815260016020819052604090912001805460001901905550565b6002545b90565b6060600080600160008581526020019081526020016000206001015490506060816040519080825280602002602001820160405280156103a457816020015b606081526020019060019003908161038f5790505b50905060005b8281101561047057600086815260016020818152604080842085855282529283902080548451600294821615610100026000190190911693909304601f810183900483028401830190945283835291929083018282801561044c5780601f106104215761010080835404028352916020019161044c565b820191906000526020600020905b81548152906001019060200180831161042f57829003601f168201915b505050505082828151811061045d57fe5b60209081029190910101526001016103aa565b5092509050915091565b60608060006004541161049f5760405162461bcd60e51b815260040161022990610e4d565b6104ad60016004540361063c565b915091509091565b6104bd610618565b6104d95760405162461bcd60e51b815260040161022990610e3d565b60025481116104fa5760405162461bcd60e51b815260040161022990610e5d565b600255565b610507610618565b6105235760405162461bcd60e51b815260040161022990610e3d565b60045483106105445760405162461bcd60e51b815260040161022990610e4d565b604080518082018252828152602080820185905260008681526003825292909220815180519293919261057a92849201906108c4565b50602082810151805161059392600185019201906108c4565b505050505050565b6105a3610618565b6105bf5760405162461bcd60e51b815260040161022990610e3d565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6000546001600160a01b031690565b600080546001600160a01b031661062d61083f565b6001600160a01b031614905090565b600081815260036020908152604091829020600180820180548551600293821615610100026000190190911692909204601f81018590048502830185019095528482526060948594919392918491908301828280156106dc5780601f106106b1576101008083540402835291602001916106dc565b820191906000526020600020905b8154815290600101906020018083116106bf57829003601f168201915b5050845460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529597508694509250840190508282801561076a5780601f1061073f5761010080835404028352916020019161076a565b820191906000526020600020905b81548152906001019060200180831161074d57829003601f168201915b5050505050905091509150915091565b606060006104ad600254610350565b610791610618565b6107ad5760405162461bcd60e51b815260040161022990610e3d565b60408051808201825282815260208082018590526004546000908152600382529290922081518051929391926107e692849201906108c4565b5060208281015180516107ff92600185019201906108c4565b5050600480546001019055505050565b610817610618565b6108335760405162461bcd60e51b815260040161022990610e3d565b61083c81610843565b50565b3390565b6001600160a01b0381166108695760405162461bcd60e51b815260040161022990610e1d565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061090557805160ff1916838001178555610932565b82800160010185558215610932579182015b82811115610932578251825591602001919060010190610917565b5061093e9291506109b7565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061097b5780548555610932565b8280016001018555821561093257600052602060002091601f016020900482015b8281111561093257825482559160010191906001019061099c565b61034d91905b8082111561093e57600081556001016109bd565b80356109dc81610f3f565b92915050565b600082601f8301126109f357600080fd5b8135610a06610a0182610ea2565b610e7b565b91508082526020830160208301858383011115610a2257600080fd5b610a2d838284610ef9565b50505092915050565b80356109dc81610f53565b600060208284031215610a5357600080fd5b6000610a5f84846109d1565b949350505050565b600060208284031215610a7957600080fd5b813567ffffffffffffffff811115610a9057600080fd5b610a5f848285016109e2565b60008060408385031215610aaf57600080fd5b823567ffffffffffffffff811115610ac657600080fd5b610ad2858286016109e2565b925050602083013567ffffffffffffffff811115610aef57600080fd5b610afb858286016109e2565b9150509250929050565b600060208284031215610b1757600080fd5b6000610a5f8484610a36565b600080600060608486031215610b3857600080fd5b6000610b448686610a36565b935050602084013567ffffffffffffffff811115610b6157600080fd5b610b6d868287016109e2565b925050604084013567ffffffffffffffff811115610b8a57600080fd5b610b96868287016109e2565b9150509250925092565b6000610bac8383610c39565b9392505050565b610bbc81610edd565b82525050565b6000610bcd82610ed0565b610bd78185610ed4565b935083602082028501610be985610eca565b8060005b85811015610c235784840389528151610c068582610ba0565b9450610c1183610eca565b60209a909a0199925050600101610bed565b5091979650505050505050565b610bbc81610ee8565b6000610c4482610ed0565b610c4e8185610ed4565b9350610c5e818560208601610f05565b610c6781610f35565b9093019392505050565b6000610c7e602683610ed4565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181526564647265737360d01b602082015260400192915050565b6000610cc6601e83610ed4565b7f4e6f64653a2064656c657465206e6f6465206973206e6f742065786973740000815260200192915050565b6000610cff602083610ed4565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572815260200192915050565b6000610d38602383610ed4565b7f56657273696f6e3a20696e6465782076657273696f6e206973206e6f742065788152621a5cdd60ea1b602082015260400192915050565b6000610d7d602183610ed4565b7f4e6f64653a206c617374206e6f6465206973206e65776572207468616e206e6f8152607760f81b602082015260400192915050565b610bbc8161034d565b602081016109dc8284610bb3565b60408082528101610ddb8185610bc2565b9050610bac6020830184610db3565b602081016109dc8284610c30565b60408082528101610e098185610c39565b90508181036020830152610a5f8184610c39565b602080825281016109dc81610c71565b602080825281016109dc81610cb9565b602080825281016109dc81610cf2565b602080825281016109dc81610d2b565b602080825281016109dc81610d70565b602081016109dc8284610db3565b60405181810167ffffffffffffffff81118282101715610e9a57600080fd5b604052919050565b600067ffffffffffffffff821115610eb957600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b60006109dc82610eed565b151590565b6001600160a01b031690565b82818337506000910152565b60005b83811015610f20578181015183820152602001610f08565b83811115610f2f576000848401525b50505050565b601f01601f191690565b610f4881610edd565b811461083c57600080fd5b610f488161034d56fea365627a7a7231582015269eec4fc665c2e295bb595e59cebad0b34ab4cf7d7e42676527b31a57243c6c6578706572696d656e74616cf564736f6c634300050c0040"

// DeployDNode deploys a new Ethereum contract, binding an instance of DNode to it.
func DeployDNode(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DNode, error) {
	parsed, err := abi.JSON(strings.NewReader(DNodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DNodeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DNode{DNodeCaller: DNodeCaller{contract: contract}, DNodeTransactor: DNodeTransactor{contract: contract}, DNodeFilterer: DNodeFilterer{contract: contract}}, nil
}

// DNode is an auto generated Go binding around an Ethereum contract.
type DNode struct {
	DNodeCaller     // Read-only binding to the contract
	DNodeTransactor // Write-only binding to the contract
	DNodeFilterer   // Log filterer for contract events
}

// DNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type DNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DNodeSession struct {
	Contract     *DNode            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DNodeCallerSession struct {
	Contract *DNodeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DNodeTransactorSession struct {
	Contract     *DNodeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type DNodeRaw struct {
	Contract *DNode // Generic contract binding to access the raw methods on
}

// DNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DNodeCallerRaw struct {
	Contract *DNodeCaller // Generic read-only contract binding to access the raw methods on
}

// DNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DNodeTransactorRaw struct {
	Contract *DNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDNode creates a new instance of DNode, bound to a specific deployed contract.
func NewDNode(address common.Address, backend bind.ContractBackend) (*DNode, error) {
	contract, err := bindDNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DNode{DNodeCaller: DNodeCaller{contract: contract}, DNodeTransactor: DNodeTransactor{contract: contract}, DNodeFilterer: DNodeFilterer{contract: contract}}, nil
}

// NewDNodeCaller creates a new read-only instance of DNode, bound to a specific deployed contract.
func NewDNodeCaller(address common.Address, caller bind.ContractCaller) (*DNodeCaller, error) {
	contract, err := bindDNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DNodeCaller{contract: contract}, nil
}

// NewDNodeTransactor creates a new write-only instance of DNode, bound to a specific deployed contract.
func NewDNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*DNodeTransactor, error) {
	contract, err := bindDNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DNodeTransactor{contract: contract}, nil
}

// NewDNodeFilterer creates a new log filterer instance of DNode, bound to a specific deployed contract.
func NewDNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*DNodeFilterer, error) {
	contract, err := bindDNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DNodeFilterer{contract: contract}, nil
}

// bindDNode binds a generic wrapper to an already deployed contract.
func bindDNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DNode *DNodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DNode.Contract.DNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DNode *DNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DNode.Contract.DNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DNode *DNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DNode.Contract.DNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DNode *DNodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DNode *DNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DNode *DNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DNode.Contract.contract.Transact(opts, method, params...)
}

// GetLast is a free data retrieval call binding the contract method 0x4d622831.
//
// Solidity: function getLast() constant returns(uint256)
func (_DNode *DNodeCaller) GetLast(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DNode.contract.Call(opts, out, "getLast")
	return *ret0, err
}

// GetLast is a free data retrieval call binding the contract method 0x4d622831.
//
// Solidity: function getLast() constant returns(uint256)
func (_DNode *DNodeSession) GetLast() (*big.Int, error) {
	return _DNode.Contract.GetLast(&_DNode.CallOpts)
}

// GetLast is a free data retrieval call binding the contract method 0x4d622831.
//
// Solidity: function getLast() constant returns(uint256)
func (_DNode *DNodeCallerSession) GetLast() (*big.Int, error) {
	return _DNode.Contract.GetLast(&_DNode.CallOpts)
}

// GetLastNode is a free data retrieval call binding the contract method 0xd1d72e3e.
//
// Solidity: function getLastNode() constant returns(string[], uint256)
func (_DNode *DNodeCaller) GetLastNode(opts *bind.CallOpts) ([]string, *big.Int, error) {
	var (
		ret0 = new([]string)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DNode.contract.Call(opts, out, "getLastNode")
	return *ret0, *ret1, err
}

// GetLastNode is a free data retrieval call binding the contract method 0xd1d72e3e.
//
// Solidity: function getLastNode() constant returns(string[], uint256)
func (_DNode *DNodeSession) GetLastNode() ([]string, *big.Int, error) {
	return _DNode.Contract.GetLastNode(&_DNode.CallOpts)
}

// GetLastNode is a free data retrieval call binding the contract method 0xd1d72e3e.
//
// Solidity: function getLastNode() constant returns(string[], uint256)
func (_DNode *DNodeCallerSession) GetLastNode() ([]string, *big.Int, error) {
	return _DNode.Contract.GetLastNode(&_DNode.CallOpts)
}

// GetLastVersion is a free data retrieval call binding the contract method 0x558614c7.
//
// Solidity: function getLastVersion() constant returns(string, string)
func (_DNode *DNodeCaller) GetLastVersion(opts *bind.CallOpts) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DNode.contract.Call(opts, out, "getLastVersion")
	return *ret0, *ret1, err
}

// GetLastVersion is a free data retrieval call binding the contract method 0x558614c7.
//
// Solidity: function getLastVersion() constant returns(string, string)
func (_DNode *DNodeSession) GetLastVersion() (string, string, error) {
	return _DNode.Contract.GetLastVersion(&_DNode.CallOpts)
}

// GetLastVersion is a free data retrieval call binding the contract method 0x558614c7.
//
// Solidity: function getLastVersion() constant returns(string, string)
func (_DNode *DNodeCallerSession) GetLastVersion() (string, string, error) {
	return _DNode.Contract.GetLastVersion(&_DNode.CallOpts)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 idx) constant returns(string[], uint256)
func (_DNode *DNodeCaller) GetNode(opts *bind.CallOpts, idx *big.Int) ([]string, *big.Int, error) {
	var (
		ret0 = new([]string)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DNode.contract.Call(opts, out, "getNode", idx)
	return *ret0, *ret1, err
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 idx) constant returns(string[], uint256)
func (_DNode *DNodeSession) GetNode(idx *big.Int) ([]string, *big.Int, error) {
	return _DNode.Contract.GetNode(&_DNode.CallOpts, idx)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 idx) constant returns(string[], uint256)
func (_DNode *DNodeCallerSession) GetNode(idx *big.Int) ([]string, *big.Int, error) {
	return _DNode.Contract.GetNode(&_DNode.CallOpts, idx)
}

// GetVersion is a free data retrieval call binding the contract method 0xb88da759.
//
// Solidity: function getVersion(uint256 idx) constant returns(string, string)
func (_DNode *DNodeCaller) GetVersion(opts *bind.CallOpts, idx *big.Int) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DNode.contract.Call(opts, out, "getVersion", idx)
	return *ret0, *ret1, err
}

// GetVersion is a free data retrieval call binding the contract method 0xb88da759.
//
// Solidity: function getVersion(uint256 idx) constant returns(string, string)
func (_DNode *DNodeSession) GetVersion(idx *big.Int) (string, string, error) {
	return _DNode.Contract.GetVersion(&_DNode.CallOpts, idx)
}

// GetVersion is a free data retrieval call binding the contract method 0xb88da759.
//
// Solidity: function getVersion(uint256 idx) constant returns(string, string)
func (_DNode *DNodeCallerSession) GetVersion(idx *big.Int) (string, string, error) {
	return _DNode.Contract.GetVersion(&_DNode.CallOpts, idx)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DNode *DNodeCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DNode.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DNode *DNodeSession) IsOwner() (bool, error) {
	return _DNode.Contract.IsOwner(&_DNode.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DNode *DNodeCallerSession) IsOwner() (bool, error) {
	return _DNode.Contract.IsOwner(&_DNode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DNode *DNodeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DNode.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DNode *DNodeSession) Owner() (common.Address, error) {
	return _DNode.Contract.Owner(&_DNode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DNode *DNodeCallerSession) Owner() (common.Address, error) {
	return _DNode.Contract.Owner(&_DNode.CallOpts)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 idx) returns()
func (_DNode *DNodeTransactor) Remove(opts *bind.TransactOpts, idx *big.Int) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "remove", idx)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 idx) returns()
func (_DNode *DNodeSession) Remove(idx *big.Int) (*types.Transaction, error) {
	return _DNode.Contract.Remove(&_DNode.TransactOpts, idx)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 idx) returns()
func (_DNode *DNodeTransactorSession) Remove(idx *big.Int) (*types.Transaction, error) {
	return _DNode.Contract.Remove(&_DNode.TransactOpts, idx)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DNode *DNodeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DNode *DNodeSession) RenounceOwnership() (*types.Transaction, error) {
	return _DNode.Contract.RenounceOwnership(&_DNode.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DNode *DNodeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DNode.Contract.RenounceOwnership(&_DNode.TransactOpts)
}

// SetLast is a paid mutator transaction binding the contract method 0x5c57bb89.
//
// Solidity: function setLast(uint256 idx) returns()
func (_DNode *DNodeTransactor) SetLast(opts *bind.TransactOpts, idx *big.Int) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "setLast", idx)
}

// SetLast is a paid mutator transaction binding the contract method 0x5c57bb89.
//
// Solidity: function setLast(uint256 idx) returns()
func (_DNode *DNodeSession) SetLast(idx *big.Int) (*types.Transaction, error) {
	return _DNode.Contract.SetLast(&_DNode.TransactOpts, idx)
}

// SetLast is a paid mutator transaction binding the contract method 0x5c57bb89.
//
// Solidity: function setLast(uint256 idx) returns()
func (_DNode *DNodeTransactorSession) SetLast(idx *big.Int) (*types.Transaction, error) {
	return _DNode.Contract.SetLast(&_DNode.TransactOpts, idx)
}

// SetLastVersion is a paid mutator transaction binding the contract method 0xe80d8c38.
//
// Solidity: function setLastVersion(string ver, string hash) returns()
func (_DNode *DNodeTransactor) SetLastVersion(opts *bind.TransactOpts, ver string, hash string) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "setLastVersion", ver, hash)
}

// SetLastVersion is a paid mutator transaction binding the contract method 0xe80d8c38.
//
// Solidity: function setLastVersion(string ver, string hash) returns()
func (_DNode *DNodeSession) SetLastVersion(ver string, hash string) (*types.Transaction, error) {
	return _DNode.Contract.SetLastVersion(&_DNode.TransactOpts, ver, hash)
}

// SetLastVersion is a paid mutator transaction binding the contract method 0xe80d8c38.
//
// Solidity: function setLastVersion(string ver, string hash) returns()
func (_DNode *DNodeTransactorSession) SetLastVersion(ver string, hash string) (*types.Transaction, error) {
	return _DNode.Contract.SetLastVersion(&_DNode.TransactOpts, ver, hash)
}

// SetVersion is a paid mutator transaction binding the contract method 0x69cf7c9a.
//
// Solidity: function setVersion(uint256 idx, string ver, string hash) returns()
func (_DNode *DNodeTransactor) SetVersion(opts *bind.TransactOpts, idx *big.Int, ver string, hash string) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "setVersion", idx, ver, hash)
}

// SetVersion is a paid mutator transaction binding the contract method 0x69cf7c9a.
//
// Solidity: function setVersion(uint256 idx, string ver, string hash) returns()
func (_DNode *DNodeSession) SetVersion(idx *big.Int, ver string, hash string) (*types.Transaction, error) {
	return _DNode.Contract.SetVersion(&_DNode.TransactOpts, idx, ver, hash)
}

// SetVersion is a paid mutator transaction binding the contract method 0x69cf7c9a.
//
// Solidity: function setVersion(uint256 idx, string ver, string hash) returns()
func (_DNode *DNodeTransactorSession) SetVersion(idx *big.Int, ver string, hash string) (*types.Transaction, error) {
	return _DNode.Contract.SetVersion(&_DNode.TransactOpts, idx, ver, hash)
}

// Store is a paid mutator transaction binding the contract method 0x131a0680.
//
// Solidity: function store(string swarm) returns()
func (_DNode *DNodeTransactor) Store(opts *bind.TransactOpts, swarm string) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "store", swarm)
}

// Store is a paid mutator transaction binding the contract method 0x131a0680.
//
// Solidity: function store(string swarm) returns()
func (_DNode *DNodeSession) Store(swarm string) (*types.Transaction, error) {
	return _DNode.Contract.Store(&_DNode.TransactOpts, swarm)
}

// Store is a paid mutator transaction binding the contract method 0x131a0680.
//
// Solidity: function store(string swarm) returns()
func (_DNode *DNodeTransactorSession) Store(swarm string) (*types.Transaction, error) {
	return _DNode.Contract.Store(&_DNode.TransactOpts, swarm)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DNode *DNodeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DNode *DNodeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DNode.Contract.TransferOwnership(&_DNode.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DNode *DNodeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DNode.Contract.TransferOwnership(&_DNode.TransactOpts, newOwner)
}

// DNodeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DNode contract.
type DNodeOwnershipTransferredIterator struct {
	Event *DNodeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DNodeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DNodeOwnershipTransferred)
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
		it.Event = new(DNodeOwnershipTransferred)
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
func (it *DNodeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DNodeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DNodeOwnershipTransferred represents a OwnershipTransferred event raised by the DNode contract.
type DNodeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DNode *DNodeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DNodeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DNode.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DNodeOwnershipTransferredIterator{contract: _DNode.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DNode *DNodeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DNodeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DNode.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DNodeOwnershipTransferred)
				if err := _DNode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DNode *DNodeFilterer) ParseOwnershipTransferred(log types.Log) (*DNodeOwnershipTransferred, error) {
	event := new(DNodeOwnershipTransferred)
	if err := _DNode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
