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
const DNodeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cid\",\"type\":\"uint256\"}],\"name\":\"getDeviceVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNodeLastData\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"removeNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"replaceVersion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_ver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"setDeviceVersion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"setNodeLast\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cid\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"setVersion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"swarm\",\"type\":\"string\"}],\"name\":\"storeNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"truncateNode\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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
	"69cf7c9a": "setVersion(uint256,string,string)",
	"0a73ba53": "storeNode(string)",
	"f2fde38b": "transferOwnership(address)",
	"5bf1f19e": "truncateNode(uint256)",
}

// DNodeBin is the compiled bytecode used for deploying new contracts.
var DNodeBin = "0x608060405260006100176001600160e01b0361006616565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35061006a565b3390565b6110a4806100796000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c80638da5cb5b11610097578063ed6c8d0b11610066578063ed6c8d0b1461020b578063f2fde38b1461021e578063f428efb014610231578063ffd740df1461023957610100565b80638da5cb5b146101ad5780638f32d59b146101c25780639c90028f146101d7578063b88da759146101ea57610100565b80635bf1f19e116100d35780635bf1f19e1461016a57806369cf7c9a1461017d578063715018a614610190578063812b5e121461019857610100565b8063011fc71a146101055780630a73ba531461011a578063342edfba1461012d5780634f0f4aa914610140575b600080fd5b610118610113366004610b97565b61024c565b005b610118610128366004610b44565b61029b565b61011861013b366004610c14565b61030e565b61015361014e366004610b79565b61039a565b604051610161929190610ecf565b60405180910390f35b610118610178366004610b79565b6104bb565b61011861018b366004610b97565b6104f3565b61011861058a565b6101a06105f8565b6040516101619190610f72565b6101b56105ff565b6040516101619190610ec1565b6101ca61060e565b6040516101619190610eef565b6101186101e5366004610b79565b610632565b6101fd6101f8366004610b79565b61067c565b604051610161929190610efd565b6101fd610219366004610b79565b6107be565b61011861022c366004610b1e565b61080f565b61015361083f565b610118610247366004610b79565b610856565b61025461060e565b6102795760405162461bcd60e51b815260040161027090610f42565b60405180910390fd5b6005546000848152600160205260409020556102968383836104f3565b505050565b6102a361060e565b6102bf5760405162461bcd60e51b815260040161027090610f42565b600354600090815260026020908152604080832060018101548452825290912082516102ed928401906109a1565b50506003546000908152600260205260409020600190810180549091019055565b61031661060e565b6103325760405162461bcd60e51b815260040161027090610f42565b60408051606081018252848152602080820184815282840186905260008881526004835293909320825181559251805192939261037592600185019201906109a1565b50604082015180516103919160028401916020909101906109a1565b50505050505050565b600081815260026020908152604080832060010154815181815281840281019093019091526060929183908280156103e657816020015b60608152602001906001900390816103d15790505b50905060005b828110156104b1576000868152600260208181526040808420858552825292839020805484516001821615610100026000190190911693909304601f810183900483028401830190945283835291929083018282801561048d5780601f106104625761010080835404028352916020019161048d565b820191906000526020600020905b81548152906001019060200180831161047057829003601f168201915b505050505082828151811061049e57fe5b60209081029190910101526001016103ec565b5092509050915091565b6104c361060e565b6104df5760405162461bcd60e51b815260040161027090610f42565b600090815260026020526040812060010155565b6104fb61060e565b6105175760405162461bcd60e51b815260040161027090610f42565b60408051606081018252848152602080820184815282840186905260055460009081526004835293909320825181559251805192939261055d92600185019201906109a1565b50604082015180516105799160028401916020909101906109a1565b505060058054600101905550505050565b61059261060e565b6105ae5760405162461bcd60e51b815260040161027090610f42565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6003545b90565b6000546001600160a01b031690565b600080546001600160a01b031661062361091c565b6001600160a01b031614905090565b61063a61060e565b6106565760405162461bcd60e51b815260040161027090610f42565b60035481116106775760405162461bcd60e51b815260040161027090610f62565b600355565b600081815260046020908152604091829020600280820180548551600180831615610100026000190190921693909304601f8101869004860284018601909652858352606095869592949190910192918491908301828280156107205780601f106106f557610100808354040283529160200191610720565b820191906000526020600020905b81548152906001019060200180831161070357829003601f168201915b5050845460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959750869450925084019050828280156107ae5780601f10610783576101008083540402835291602001916107ae565b820191906000526020600020905b81548152906001019060200180831161079157829003601f168201915b5050505050905091509150915091565b60008181526001602052604090205460609081906107ee5760405162461bcd60e51b815260040161027090610f52565b6000838152600160205260409020546108069061067c565b91509150915091565b61081761060e565b6108335760405162461bcd60e51b815260040161027090610f42565b61083c81610920565b50565b6060600061084e60035461039a565b915091509091565b61085e61060e565b61087a5760405162461bcd60e51b815260040161027090610f42565b60035460009081526002602052604090206001015481106108ad5760405162461bcd60e51b815260040161027090610f32565b600354600090815260026020818152604080842060018082015460001980820188529290945282862087875292909520825493956108fb959194908116156101000290920190911604610a1f565b50506003546000908152600260205260409020600101805460001901905550565b3390565b6001600160a01b0381166109465760405162461bcd60e51b815260040161027090610f22565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106109e257805160ff1916838001178555610a0f565b82800160010185558215610a0f579182015b82811115610a0f5782518255916020019190600101906109f4565b50610a1b929150610a94565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a585780548555610a0f565b82800160010185558215610a0f57600052602060002091601f016020900482015b82811115610a0f578254825591600101919060010190610a79565b6105fc91905b80821115610a1b5760008155600101610a9a565b8035610ab981611044565b92915050565b600082601f830112610ad057600080fd5b8135610ae3610ade82610fa7565b610f80565b91508082526020830160208301858383011115610aff57600080fd5b610b0a838284610ffe565b50505092915050565b8035610ab981611058565b600060208284031215610b3057600080fd5b6000610b3c8484610aae565b949350505050565b600060208284031215610b5657600080fd5b813567ffffffffffffffff811115610b6d57600080fd5b610b3c84828501610abf565b600060208284031215610b8b57600080fd5b6000610b3c8484610b13565b600080600060608486031215610bac57600080fd5b6000610bb88686610b13565b935050602084013567ffffffffffffffff811115610bd557600080fd5b610be186828701610abf565b925050604084013567ffffffffffffffff811115610bfe57600080fd5b610c0a86828701610abf565b9150509250925092565b60008060008060808587031215610c2a57600080fd5b6000610c368787610b13565b9450506020610c4787828801610b13565b935050604085013567ffffffffffffffff811115610c6457600080fd5b610c7087828801610abf565b925050606085013567ffffffffffffffff811115610c8d57600080fd5b610c9987828801610abf565b91505092959194509250565b6000610cb18383610d3e565b9392505050565b610cc181610fe2565b82525050565b6000610cd282610fd5565b610cdc8185610fd9565b935083602082028501610cee85610fcf565b8060005b85811015610d285784840389528151610d0b8582610ca5565b9450610d1683610fcf565b60209a909a0199925050600101610cf2565b5091979650505050505050565b610cc181610fed565b6000610d4982610fd5565b610d538185610fd9565b9350610d6381856020860161100a565b610d6c8161103a565b9093019392505050565b6000610d83602683610fd9565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181526564647265737360d01b602082015260400192915050565b6000610dcb601e83610fd9565b7f4e6f64653a2064656c657465206e6f6465206973206e6f742065786973740000815260200192915050565b6000610e04602083610fd9565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572815260200192915050565b6000610e3d602383610fd9565b7f56657273696f6e3a20696e6465782076657273696f6e206973206e6f742065788152621a5cdd60ea1b602082015260400192915050565b6000610e82602183610fd9565b7f4e6f64653a206c617374206e6f6465206973206e65776572207468616e206e6f8152607760f81b602082015260400192915050565b610cc1816105fc565b60208101610ab98284610cb8565b60408082528101610ee08185610cc7565b9050610cb16020830184610eb8565b60208101610ab98284610d35565b60408082528101610f0e8185610d3e565b90508181036020830152610b3c8184610d3e565b60208082528101610ab981610d76565b60208082528101610ab981610dbe565b60208082528101610ab981610df7565b60208082528101610ab981610e30565b60208082528101610ab981610e75565b60208101610ab98284610eb8565b60405181810167ffffffffffffffff81118282101715610f9f57600080fd5b604052919050565b600067ffffffffffffffff821115610fbe57600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b6000610ab982610ff2565b151590565b6001600160a01b031690565b82818337506000910152565b60005b8381101561102557818101518382015260200161100d565b83811115611034576000848401525b50505050565b601f01601f191690565b61104d81610fe2565b811461083c57600080fd5b61104d816105fc56fea365627a7a72315820a0193a6543d8a7d58ef7bef555559583036f248b079d9b00821d1863a1d5992d6c6578706572696d656e74616cf564736f6c634300050c0040"

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

// GetDeviceVersion is a free data retrieval call binding the contract method 0xed6c8d0b.
//
// Solidity: function getDeviceVersion(uint256 cid) constant returns(string, string)
func (_DNode *DNodeCaller) GetDeviceVersion(opts *bind.CallOpts, cid *big.Int) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DNode.contract.Call(opts, out, "getDeviceVersion", cid)
	return *ret0, *ret1, err
}

// GetDeviceVersion is a free data retrieval call binding the contract method 0xed6c8d0b.
//
// Solidity: function getDeviceVersion(uint256 cid) constant returns(string, string)
func (_DNode *DNodeSession) GetDeviceVersion(cid *big.Int) (string, string, error) {
	return _DNode.Contract.GetDeviceVersion(&_DNode.CallOpts, cid)
}

// GetDeviceVersion is a free data retrieval call binding the contract method 0xed6c8d0b.
//
// Solidity: function getDeviceVersion(uint256 cid) constant returns(string, string)
func (_DNode *DNodeCallerSession) GetDeviceVersion(cid *big.Int) (string, string, error) {
	return _DNode.Contract.GetDeviceVersion(&_DNode.CallOpts, cid)
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

// RemoveNode is a paid mutator transaction binding the contract method 0xffd740df.
//
// Solidity: function removeNode(uint256 idx) returns()
func (_DNode *DNodeTransactor) RemoveNode(opts *bind.TransactOpts, idx *big.Int) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "removeNode", idx)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xffd740df.
//
// Solidity: function removeNode(uint256 idx) returns()
func (_DNode *DNodeSession) RemoveNode(idx *big.Int) (*types.Transaction, error) {
	return _DNode.Contract.RemoveNode(&_DNode.TransactOpts, idx)
}

// RemoveNode is a paid mutator transaction binding the contract method 0xffd740df.
//
// Solidity: function removeNode(uint256 idx) returns()
func (_DNode *DNodeTransactorSession) RemoveNode(idx *big.Int) (*types.Transaction, error) {
	return _DNode.Contract.RemoveNode(&_DNode.TransactOpts, idx)
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
// Solidity: function replaceVersion(uint256 idx, uint256 cid, string ver, string hash) returns()
func (_DNode *DNodeTransactor) ReplaceVersion(opts *bind.TransactOpts, idx *big.Int, cid *big.Int, ver string, hash string) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "replaceVersion", idx, cid, ver, hash)
}

// ReplaceVersion is a paid mutator transaction binding the contract method 0x342edfba.
//
// Solidity: function replaceVersion(uint256 idx, uint256 cid, string ver, string hash) returns()
func (_DNode *DNodeSession) ReplaceVersion(idx *big.Int, cid *big.Int, ver string, hash string) (*types.Transaction, error) {
	return _DNode.Contract.ReplaceVersion(&_DNode.TransactOpts, idx, cid, ver, hash)
}

// ReplaceVersion is a paid mutator transaction binding the contract method 0x342edfba.
//
// Solidity: function replaceVersion(uint256 idx, uint256 cid, string ver, string hash) returns()
func (_DNode *DNodeTransactorSession) ReplaceVersion(idx *big.Int, cid *big.Int, ver string, hash string) (*types.Transaction, error) {
	return _DNode.Contract.ReplaceVersion(&_DNode.TransactOpts, idx, cid, ver, hash)
}

// SetDeviceVersion is a paid mutator transaction binding the contract method 0x011fc71a.
//
// Solidity: function setDeviceVersion(uint256 _cid, string _ver, string _hash) returns()
func (_DNode *DNodeTransactor) SetDeviceVersion(opts *bind.TransactOpts, _cid *big.Int, _ver string, _hash string) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "setDeviceVersion", _cid, _ver, _hash)
}

// SetDeviceVersion is a paid mutator transaction binding the contract method 0x011fc71a.
//
// Solidity: function setDeviceVersion(uint256 _cid, string _ver, string _hash) returns()
func (_DNode *DNodeSession) SetDeviceVersion(_cid *big.Int, _ver string, _hash string) (*types.Transaction, error) {
	return _DNode.Contract.SetDeviceVersion(&_DNode.TransactOpts, _cid, _ver, _hash)
}

// SetDeviceVersion is a paid mutator transaction binding the contract method 0x011fc71a.
//
// Solidity: function setDeviceVersion(uint256 _cid, string _ver, string _hash) returns()
func (_DNode *DNodeTransactorSession) SetDeviceVersion(_cid *big.Int, _ver string, _hash string) (*types.Transaction, error) {
	return _DNode.Contract.SetDeviceVersion(&_DNode.TransactOpts, _cid, _ver, _hash)
}

// SetNodeLast is a paid mutator transaction binding the contract method 0x9c90028f.
//
// Solidity: function setNodeLast(uint256 idx) returns()
func (_DNode *DNodeTransactor) SetNodeLast(opts *bind.TransactOpts, idx *big.Int) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "setNodeLast", idx)
}

// SetNodeLast is a paid mutator transaction binding the contract method 0x9c90028f.
//
// Solidity: function setNodeLast(uint256 idx) returns()
func (_DNode *DNodeSession) SetNodeLast(idx *big.Int) (*types.Transaction, error) {
	return _DNode.Contract.SetNodeLast(&_DNode.TransactOpts, idx)
}

// SetNodeLast is a paid mutator transaction binding the contract method 0x9c90028f.
//
// Solidity: function setNodeLast(uint256 idx) returns()
func (_DNode *DNodeTransactorSession) SetNodeLast(idx *big.Int) (*types.Transaction, error) {
	return _DNode.Contract.SetNodeLast(&_DNode.TransactOpts, idx)
}

// SetVersion is a paid mutator transaction binding the contract method 0x69cf7c9a.
//
// Solidity: function setVersion(uint256 cid, string ver, string hash) returns()
func (_DNode *DNodeTransactor) SetVersion(opts *bind.TransactOpts, cid *big.Int, ver string, hash string) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "setVersion", cid, ver, hash)
}

// SetVersion is a paid mutator transaction binding the contract method 0x69cf7c9a.
//
// Solidity: function setVersion(uint256 cid, string ver, string hash) returns()
func (_DNode *DNodeSession) SetVersion(cid *big.Int, ver string, hash string) (*types.Transaction, error) {
	return _DNode.Contract.SetVersion(&_DNode.TransactOpts, cid, ver, hash)
}

// SetVersion is a paid mutator transaction binding the contract method 0x69cf7c9a.
//
// Solidity: function setVersion(uint256 cid, string ver, string hash) returns()
func (_DNode *DNodeTransactorSession) SetVersion(cid *big.Int, ver string, hash string) (*types.Transaction, error) {
	return _DNode.Contract.SetVersion(&_DNode.TransactOpts, cid, ver, hash)
}

// StoreNode is a paid mutator transaction binding the contract method 0x0a73ba53.
//
// Solidity: function storeNode(string swarm) returns()
func (_DNode *DNodeTransactor) StoreNode(opts *bind.TransactOpts, swarm string) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "storeNode", swarm)
}

// StoreNode is a paid mutator transaction binding the contract method 0x0a73ba53.
//
// Solidity: function storeNode(string swarm) returns()
func (_DNode *DNodeSession) StoreNode(swarm string) (*types.Transaction, error) {
	return _DNode.Contract.StoreNode(&_DNode.TransactOpts, swarm)
}

// StoreNode is a paid mutator transaction binding the contract method 0x0a73ba53.
//
// Solidity: function storeNode(string swarm) returns()
func (_DNode *DNodeTransactorSession) StoreNode(swarm string) (*types.Transaction, error) {
	return _DNode.Contract.StoreNode(&_DNode.TransactOpts, swarm)
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

// TruncateNode is a paid mutator transaction binding the contract method 0x5bf1f19e.
//
// Solidity: function truncateNode(uint256 idx) returns()
func (_DNode *DNodeTransactor) TruncateNode(opts *bind.TransactOpts, idx *big.Int) (*types.Transaction, error) {
	return _DNode.contract.Transact(opts, "truncateNode", idx)
}

// TruncateNode is a paid mutator transaction binding the contract method 0x5bf1f19e.
//
// Solidity: function truncateNode(uint256 idx) returns()
func (_DNode *DNodeSession) TruncateNode(idx *big.Int) (*types.Transaction, error) {
	return _DNode.Contract.TruncateNode(&_DNode.TransactOpts, idx)
}

// TruncateNode is a paid mutator transaction binding the contract method 0x5bf1f19e.
//
// Solidity: function truncateNode(uint256 idx) returns()
func (_DNode *DNodeTransactorSession) TruncateNode(idx *big.Int) (*types.Transaction, error) {
	return _DNode.Contract.TruncateNode(&_DNode.TransactOpts, idx)
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
