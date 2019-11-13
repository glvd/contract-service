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
const DNodeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_device\",\"type\":\"uint256\"}],\"name\":\"getDeviceVersion\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"device\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structDNode.Version\",\"name\":\"_version\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeLastData\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"}],\"name\":\"getVersion\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"device\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structDNode.Version\",\"name\":\"_version\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"}],\"name\":\"removeNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_ver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"replaceVersion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_device\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_ver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"setDeviceVersion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_idx\",\"type\":\"uint256\"}],\"name\":\"setNodeLast\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_swarm\",\"type\":\"string\"}],\"name\":\"storeNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DNodeFuncSigs maps the 4-byte function signature to its string representation.
var DNodeFuncSigs = map[string]string{
	"ed6c8d0b": "getDeviceVersion(uint256)",
	"4f0f4aa9": "getNode(uint256)",
	"812b5e12": "getNodeLast()",
	"f428efb0": "getNodeLastData()",
	"b88da759": "getVersion(uint256)",
	"8f32d59b": "isOwner()",
	"8da5cb5b": "owner()",
	"ffd740df": "removeNode(uint256)",
	"715018a6": "renounceOwnership()",
	"342edfba": "replaceVersion(uint256,uint256,string,string)",
	"011fc71a": "setDeviceVersion(uint256,string,string)",
	"9c90028f": "setNodeLast(uint256)",
	"0a73ba53": "storeNode(string)",
	"f2fde38b": "transferOwnership(address)",
}

// DNodeBin is the compiled bytecode used for deploying new contracts.
var DNodeBin = "0x608060405260006100176001600160e01b0361006616565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006a565b3390565b6110ae806100796000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638f32d59b1161008c578063ed6c8d0b11610066578063ed6c8d0b146101ce578063f2fde38b146101e1578063f428efb0146101f4578063ffd740df146101fc576100ea565b80638f32d59b146101865780639c90028f1461019b578063b88da759146101ae576100ea565b80634f0f4aa9116100c85780634f0f4aa91461012a578063715018a614610154578063812b5e121461015c5780638da5cb5b14610171576100ea565b8063011fc71a146100ef5780630a73ba5314610104578063342edfba14610117575b600080fd5b6101026100fd366004610b66565b61020f565b005b610102610112366004610b13565b61025e565b610102610125366004610be3565b6102d1565b61013d610138366004610b48565b61035d565b60405161014b929190610eed565b60405180910390f35b61010261047e565b6101646104ec565b60405161014b9190610f7c565b6101796104f3565b60405161014b9190610edf565b61018e610502565b60405161014b9190610f0d565b6101026101a9366004610b48565b610526565b6101c16101bc366004610b48565b610570565b60405161014b9190610f6b565b6101c16101dc366004610b48565b6106db565b6101026101ef366004610aed565b61072c565b61013d61075c565b61010261020a366004610b48565b610773565b610217610502565b61023c5760405162461bcd60e51b815260040161023390610f3b565b60405180910390fd5b600554600084815260016020526040902055610259838383610839565b505050565b610266610502565b6102825760405162461bcd60e51b815260040161023390610f3b565b600354600090815260026020908152604080832060018101548452825290912082516102b092840190610955565b50506003546000908152600260205260409020600190810180549091019055565b6102d9610502565b6102f55760405162461bcd60e51b815260040161023390610f3b565b6040805160608101825284815260208082018481528284018690526000888152600483529390932082518155925180519293926103389260018501920190610955565b5060408201518051610354916002840191602090910190610955565b50505050505050565b600081815260026020908152604080832060010154815181815281840281019093019091526060929183908280156103a957816020015b60608152602001906001900390816103945790505b50905060005b82811015610474576000868152600260208181526040808420858552825292839020805484516001821615610100026000190190911693909304601f81018390048302840183019094528383529192908301828280156104505780601f1061042557610100808354040283529160200191610450565b820191906000526020600020905b81548152906001019060200180831161043357829003601f168201915b505050505082828151811061046157fe5b60209081029190910101526001016103af565b5092509050915091565b610486610502565b6104a25760405162461bcd60e51b815260040161023390610f3b565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6003545b90565b6000546001600160a01b031690565b600080546001600160a01b03166105176108d0565b6001600160a01b031614905090565b61052e610502565b61054a5760405162461bcd60e51b815260040161023390610f3b565b600354811161056b5760405162461bcd60e51b815260040161023390610f5b565b600355565b6105786109d3565b6004600083815260200190815260200160002060405180606001604052908160008201548152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106395780601f1061060e57610100808354040283529160200191610639565b820191906000526020600020905b81548152906001019060200180831161061c57829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f810183900483028501830190915280845293810193908301828280156106cb5780601f106106a0576101008083540402835291602001916106cb565b820191906000526020600020905b8154815290600101906020018083116106ae57829003601f168201915b5050505050815250509050919050565b6106e36109d3565b60008281526001602052604090205461070e5760405162461bcd60e51b815260040161023390610f4b565b60008281526001602052604090205461072690610570565b92915050565b610734610502565b6107505760405162461bcd60e51b815260040161023390610f3b565b610759816108d4565b50565b6060600061076b60035461035d565b915091509091565b61077b610502565b6107975760405162461bcd60e51b815260040161023390610f3b565b60035460009081526002602052604090206001015481106107ca5760405162461bcd60e51b815260040161023390610f2b565b600354600090815260026020818152604080842060018082015460001980820188529290945282862087875292909520825493956108189591949081161561010002909201909116046109f4565b50506003546000908152600260205260409020600101805460001901905550565b610841610502565b61085d5760405162461bcd60e51b815260040161023390610f3b565b6040805160608101825284815260208082018481528284018690526005546000908152600483529390932082518155925180519293926108a39260018501920190610955565b50604082015180516108bf916002840191602090910190610955565b505060058054600101905550505050565b3390565b6001600160a01b0381166108fa5760405162461bcd60e51b815260040161023390610f1b565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061099657805160ff19168380011785556109c3565b828001600101855582156109c3579182015b828111156109c35782518255916020019190600101906109a8565b506109cf929150610a69565b5090565b60405180606001604052806000815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a2d57805485556109c3565b828001600101855582156109c357600052602060002091601f016020900482015b828111156109c3578254825591600101919060010190610a4e565b6104f091905b808211156109cf5760008155600101610a6f565b80356107268161104e565b600082601f830112610a9f57600080fd5b8135610ab2610aad82610fb1565b610f8a565b91508082526020830160208301858383011115610ace57600080fd5b610ad9838284611008565b50505092915050565b803561072681611062565b600060208284031215610aff57600080fd5b6000610b0b8484610a83565b949350505050565b600060208284031215610b2557600080fd5b813567ffffffffffffffff811115610b3c57600080fd5b610b0b84828501610a8e565b600060208284031215610b5a57600080fd5b6000610b0b8484610ae2565b600080600060608486031215610b7b57600080fd5b6000610b878686610ae2565b935050602084013567ffffffffffffffff811115610ba457600080fd5b610bb086828701610a8e565b925050604084013567ffffffffffffffff811115610bcd57600080fd5b610bd986828701610a8e565b9150509250925092565b60008060008060808587031215610bf957600080fd5b6000610c058787610ae2565b9450506020610c1687828801610ae2565b935050604085013567ffffffffffffffff811115610c3357600080fd5b610c3f87828801610a8e565b925050606085013567ffffffffffffffff811115610c5c57600080fd5b610c6887828801610a8e565b91505092959194509250565b6000610c808383610d0d565b9392505050565b610c9081610fec565b82525050565b6000610ca182610fdf565b610cab8185610fe3565b935083602082028501610cbd85610fd9565b8060005b85811015610cf75784840389528151610cda8582610c74565b9450610ce583610fd9565b60209a909a0199925050600101610cc1565b5091979650505050505050565b610c9081610ff7565b6000610d1882610fdf565b610d228185610fe3565b9350610d32818560208601611014565b610d3b81611044565b9093019392505050565b6000610d52602683610fe3565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181526564647265737360d01b602082015260400192915050565b6000610d9a601e83610fe3565b7f4e6f64653a2064656c657465206e6f6465206973206e6f742065786973740000815260200192915050565b6000610dd3602083610fe3565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572815260200192915050565b6000610e0c602383610fe3565b7f56657273696f6e3a20696e6465782076657273696f6e206973206e6f742065788152621a5cdd60ea1b602082015260400192915050565b6000610e51602183610fe3565b7f4e6f64653a206c617374206e6f6465206973206e65776572207468616e206e6f8152607760f81b602082015260400192915050565b80516000906060840190610e9b8582610ed6565b5060208301518482036020860152610eb38282610d0d565b91505060408301518482036040860152610ecd8282610d0d565b95945050505050565b610c90816104f0565b602081016107268284610c87565b60408082528101610efe8185610c96565b9050610c806020830184610ed6565b602081016107268284610d04565b6020808252810161072681610d45565b6020808252810161072681610d8d565b6020808252810161072681610dc6565b6020808252810161072681610dff565b6020808252810161072681610e44565b60208082528101610c808184610e87565b602081016107268284610ed6565b60405181810167ffffffffffffffff81118282101715610fa957600080fd5b604052919050565b600067ffffffffffffffff821115610fc857600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b600061072682610ffc565b151590565b6001600160a01b031690565b82818337506000910152565b60005b8381101561102f578181015183820152602001611017565b8381111561103e576000848401525b50505050565b601f01601f191690565b61105781610fec565b811461075957600080fd5b611057816104f056fea365627a7a72315820bfad0b270c1713934ece7ea9de573b868e24f2f804556afb5ebbd9706a2b0c4d6c6578706572696d656e74616cf564736f6c634300050c0040"

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Device  *big.Int
	Hash    string
	Version string
}

// GetDeviceVersion is a free data retrieval call binding the contract method 0xed6c8d0b.
//
// Solidity: function getDeviceVersion(uint256 _device) constant returns(Struct0 _version)
func (_DNode *DNodeCaller) GetDeviceVersion(opts *bind.CallOpts, _device *big.Int) (Struct0, error) {
	var (
		ret0 = new(Struct0)
	)
	out := ret0
	err := _DNode.contract.Call(opts, out, "getDeviceVersion", _device)
	return *ret0, err
}

// GetDeviceVersion is a free data retrieval call binding the contract method 0xed6c8d0b.
//
// Solidity: function getDeviceVersion(uint256 _device) constant returns(Struct0 _version)
func (_DNode *DNodeSession) GetDeviceVersion(_device *big.Int) (Struct0, error) {
	return _DNode.Contract.GetDeviceVersion(&_DNode.CallOpts, _device)
}

// GetDeviceVersion is a free data retrieval call binding the contract method 0xed6c8d0b.
//
// Solidity: function getDeviceVersion(uint256 _device) constant returns(Struct0 _version)
func (_DNode *DNodeCallerSession) GetDeviceVersion(_device *big.Int) (Struct0, error) {
	return _DNode.Contract.GetDeviceVersion(&_DNode.CallOpts, _device)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 _idx) constant returns(string[], uint256)
func (_DNode *DNodeCaller) GetNode(opts *bind.CallOpts, _idx *big.Int) ([]string, *big.Int, error) {
	var (
		ret0 = new([]string)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DNode.contract.Call(opts, out, "getNode", _idx)
	return *ret0, *ret1, err
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 _idx) constant returns(string[], uint256)
func (_DNode *DNodeSession) GetNode(_idx *big.Int) ([]string, *big.Int, error) {
	return _DNode.Contract.GetNode(&_DNode.CallOpts, _idx)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 _idx) constant returns(string[], uint256)
func (_DNode *DNodeCallerSession) GetNode(_idx *big.Int) ([]string, *big.Int, error) {
	return _DNode.Contract.GetNode(&_DNode.CallOpts, _idx)
}

// GetNodeLast is a free data retrieval call binding the contract method 0x812b5e12.
//
// Solidity: function getNodeLast() constant returns(uint256)
func (_DNode *DNodeCaller) GetNodeLast(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DNode.contract.Call(opts, out, "getNodeLast")
	return *ret0, err
}

// GetNodeLast is a free data retrieval call binding the contract method 0x812b5e12.
//
// Solidity: function getNodeLast() constant returns(uint256)
func (_DNode *DNodeSession) GetNodeLast() (*big.Int, error) {
	return _DNode.Contract.GetNodeLast(&_DNode.CallOpts)
}

// GetNodeLast is a free data retrieval call binding the contract method 0x812b5e12.
//
// Solidity: function getNodeLast() constant returns(uint256)
func (_DNode *DNodeCallerSession) GetNodeLast() (*big.Int, error) {
	return _DNode.Contract.GetNodeLast(&_DNode.CallOpts)
}

// GetNodeLastData is a free data retrieval call binding the contract method 0xf428efb0.
//
// Solidity: function getNodeLastData() constant returns(string[], uint256)
func (_DNode *DNodeCaller) GetNodeLastData(opts *bind.CallOpts) ([]string, *big.Int, error) {
	var (
		ret0 = new([]string)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DNode.contract.Call(opts, out, "getNodeLastData")
	return *ret0, *ret1, err
}

// GetNodeLastData is a free data retrieval call binding the contract method 0xf428efb0.
//
// Solidity: function getNodeLastData() constant returns(string[], uint256)
func (_DNode *DNodeSession) GetNodeLastData() ([]string, *big.Int, error) {
	return _DNode.Contract.GetNodeLastData(&_DNode.CallOpts)
}

// GetNodeLastData is a free data retrieval call binding the contract method 0xf428efb0.
//
// Solidity: function getNodeLastData() constant returns(string[], uint256)
func (_DNode *DNodeCallerSession) GetNodeLastData() ([]string, *big.Int, error) {
	return _DNode.Contract.GetNodeLastData(&_DNode.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0xb88da759.
//
// Solidity: function getVersion(uint256 _idx) constant returns(Struct0 _version)
func (_DNode *DNodeCaller) GetVersion(opts *bind.CallOpts, _idx *big.Int) (Struct0, error) {
	var (
		ret0 = new(Struct0)
	)
	out := ret0
	err := _DNode.contract.Call(opts, out, "getVersion", _idx)
	return *ret0, err
}

// GetVersion is a free data retrieval call binding the contract method 0xb88da759.
//
// Solidity: function getVersion(uint256 _idx) constant returns(Struct0 _version)
func (_DNode *DNodeSession) GetVersion(_idx *big.Int) (Struct0, error) {
	return _DNode.Contract.GetVersion(&_DNode.CallOpts, _idx)
}

// GetVersion is a free data retrieval call binding the contract method 0xb88da759.
//
// Solidity: function getVersion(uint256 _idx) constant returns(Struct0 _version)
func (_DNode *DNodeCallerSession) GetVersion(_idx *big.Int) (Struct0, error) {
	return _DNode.Contract.GetVersion(&_DNode.CallOpts, _idx)
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

// RemoveNode is a paid mutator transaction binding the contract method 0xffd740df.
//
// Solidity: function removeNode(uint256 _idx) returns()
func (_DNode *DNodeTransactor) RemoveNode(opts *bind.TransactOpts, _idx *big.Int) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "removeNode", _idx)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xffd740df.
//
// Solidity: function removeNode(uint256 _idx) returns()
func (_DNode *DNodeSession) RemoveNode(_idx *big.Int) (*types.Transaction, error) {
	return _DNode.Contract.RemoveNode(&_DNode.TransactOpts, _idx)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xffd740df.
//
// Solidity: function removeNode(uint256 _idx) returns()
func (_DNode *DNodeTransactorSession) RemoveNode(_idx *big.Int) (*types.Transaction, error) {
	return _DNode.Contract.RemoveNode(&_DNode.TransactOpts, _idx)
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

// ReplaceVersion is a paid mutator transaction binding the contract method 0x342edfba.
//
// Solidity: function replaceVersion(uint256 _idx, uint256 _cid, string _ver, string _hash) returns()
func (_DNode *DNodeTransactor) ReplaceVersion(opts *bind.TransactOpts, _idx *big.Int, _cid *big.Int, _ver string, _hash string) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "replaceVersion", _idx, _cid, _ver, _hash)
}

// ReplaceVersion is a paid mutator transaction binding the contract method 0x342edfba.
//
// Solidity: function replaceVersion(uint256 _idx, uint256 _cid, string _ver, string _hash) returns()
func (_DNode *DNodeSession) ReplaceVersion(_idx *big.Int, _cid *big.Int, _ver string, _hash string) (*types.Transaction, error) {
	return _DNode.Contract.ReplaceVersion(&_DNode.TransactOpts, _idx, _cid, _ver, _hash)
}

// ReplaceVersion is a paid mutator transaction binding the contract method 0x342edfba.
//
// Solidity: function replaceVersion(uint256 _idx, uint256 _cid, string _ver, string _hash) returns()
func (_DNode *DNodeTransactorSession) ReplaceVersion(_idx *big.Int, _cid *big.Int, _ver string, _hash string) (*types.Transaction, error) {
	return _DNode.Contract.ReplaceVersion(&_DNode.TransactOpts, _idx, _cid, _ver, _hash)
}

// SetDeviceVersion is a paid mutator transaction binding the contract method 0x011fc71a.
//
// Solidity: function setDeviceVersion(uint256 _device, string _ver, string _hash) returns()
func (_DNode *DNodeTransactor) SetDeviceVersion(opts *bind.TransactOpts, _device *big.Int, _ver string, _hash string) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "setDeviceVersion", _device, _ver, _hash)
}

// SetDeviceVersion is a paid mutator transaction binding the contract method 0x011fc71a.
//
// Solidity: function setDeviceVersion(uint256 _device, string _ver, string _hash) returns()
func (_DNode *DNodeSession) SetDeviceVersion(_device *big.Int, _ver string, _hash string) (*types.Transaction, error) {
	return _DNode.Contract.SetDeviceVersion(&_DNode.TransactOpts, _device, _ver, _hash)
}

// SetDeviceVersion is a paid mutator transaction binding the contract method 0x011fc71a.
//
// Solidity: function setDeviceVersion(uint256 _device, string _ver, string _hash) returns()
func (_DNode *DNodeTransactorSession) SetDeviceVersion(_device *big.Int, _ver string, _hash string) (*types.Transaction, error) {
	return _DNode.Contract.SetDeviceVersion(&_DNode.TransactOpts, _device, _ver, _hash)
}

// SetNodeLast is a paid mutator transaction binding the contract method 0x9c90028f.
//
// Solidity: function setNodeLast(uint256 _idx) returns()
func (_DNode *DNodeTransactor) SetNodeLast(opts *bind.TransactOpts, _idx *big.Int) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "setNodeLast", _idx)
}

// SetNodeLast is a paid mutator transaction binding the contract method 0x9c90028f.
//
// Solidity: function setNodeLast(uint256 _idx) returns()
func (_DNode *DNodeSession) SetNodeLast(_idx *big.Int) (*types.Transaction, error) {
	return _DNode.Contract.SetNodeLast(&_DNode.TransactOpts, _idx)
}

// SetNodeLast is a paid mutator transaction binding the contract method 0x9c90028f.
//
// Solidity: function setNodeLast(uint256 _idx) returns()
func (_DNode *DNodeTransactorSession) SetNodeLast(_idx *big.Int) (*types.Transaction, error) {
	return _DNode.Contract.SetNodeLast(&_DNode.TransactOpts, _idx)
}

// StoreNode is a paid mutator transaction binding the contract method 0x0a73ba53.
//
// Solidity: function storeNode(string _swarm) returns()
func (_DNode *DNodeTransactor) StoreNode(opts *bind.TransactOpts, _swarm string) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "storeNode", _swarm)
}

// StoreNode is a paid mutator transaction binding the contract method 0x0a73ba53.
//
// Solidity: function storeNode(string _swarm) returns()
func (_DNode *DNodeSession) StoreNode(_swarm string) (*types.Transaction, error) {
	return _DNode.Contract.StoreNode(&_DNode.TransactOpts, _swarm)
}

// StoreNode is a paid mutator transaction binding the contract method 0x0a73ba53.
//
// Solidity: function storeNode(string _swarm) returns()
func (_DNode *DNodeTransactorSession) StoreNode(_swarm string) (*types.Transaction, error) {
	return _DNode.Contract.StoreNode(&_DNode.TransactOpts, _swarm)
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
