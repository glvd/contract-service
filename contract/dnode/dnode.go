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

// DnodeABI is the input ABI used to generate the binding from.
const DnodeABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"swarm\",\"type\":\"string\"}],\"name\":\"store\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"setLast\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"ver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"setVersion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastNode\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"ver\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"setLastVersion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Dnode is an auto generated Go binding around an Ethereum contract.
type Dnode struct {
	DnodeCaller     // Read-only binding to the contract
	DnodeTransactor // Write-only binding to the contract
	DnodeFilterer   // Log filterer for contract events
}

// DnodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type DnodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DnodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DnodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DnodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DnodeSession struct {
	Contract     *Dnode            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DnodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DnodeCallerSession struct {
	Contract *DnodeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DnodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DnodeTransactorSession struct {
	Contract     *DnodeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DnodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type DnodeRaw struct {
	Contract *Dnode // Generic contract binding to access the raw methods on
}

// DnodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DnodeCallerRaw struct {
	Contract *DnodeCaller // Generic read-only contract binding to access the raw methods on
}

// DnodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DnodeTransactorRaw struct {
	Contract *DnodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDnode creates a new instance of Dnode, bound to a specific deployed contract.
func NewDnode(address common.Address, backend bind.ContractBackend) (*Dnode, error) {
	contract, err := bindDnode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dnode{DnodeCaller: DnodeCaller{contract: contract}, DnodeTransactor: DnodeTransactor{contract: contract}, DnodeFilterer: DnodeFilterer{contract: contract}}, nil
}

// NewDnodeCaller creates a new read-only instance of Dnode, bound to a specific deployed contract.
func NewDnodeCaller(address common.Address, caller bind.ContractCaller) (*DnodeCaller, error) {
	contract, err := bindDnode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DnodeCaller{contract: contract}, nil
}

// NewDnodeTransactor creates a new write-only instance of Dnode, bound to a specific deployed contract.
func NewDnodeTransactor(address common.Address, transactor bind.ContractTransactor) (*DnodeTransactor, error) {
	contract, err := bindDnode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DnodeTransactor{contract: contract}, nil
}

// NewDnodeFilterer creates a new log filterer instance of Dnode, bound to a specific deployed contract.
func NewDnodeFilterer(address common.Address, filterer bind.ContractFilterer) (*DnodeFilterer, error) {
	contract, err := bindDnode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DnodeFilterer{contract: contract}, nil
}

// bindDnode binds a generic wrapper to an already deployed contract.
func bindDnode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DnodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dnode *DnodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dnode.Contract.DnodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dnode *DnodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dnode.Contract.DnodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dnode *DnodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dnode.Contract.DnodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dnode *DnodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dnode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dnode *DnodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dnode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dnode *DnodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dnode.Contract.contract.Transact(opts, method, params...)
}

// GetLast is a free data retrieval call binding the contract method 0x4d622831.
//
// Solidity: function getLast() constant returns(uint256)
func (_Dnode *DnodeCaller) GetLast(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dnode.contract.Call(opts, out, "getLast")
	return *ret0, err
}

// GetLast is a free data retrieval call binding the contract method 0x4d622831.
//
// Solidity: function getLast() constant returns(uint256)
func (_Dnode *DnodeSession) GetLast() (*big.Int, error) {
	return _Dnode.Contract.GetLast(&_Dnode.CallOpts)
}

// GetLast is a free data retrieval call binding the contract method 0x4d622831.
//
// Solidity: function getLast() constant returns(uint256)
func (_Dnode *DnodeCallerSession) GetLast() (*big.Int, error) {
	return _Dnode.Contract.GetLast(&_Dnode.CallOpts)
}

// GetLastNode is a free data retrieval call binding the contract method 0xd1d72e3e.
//
// Solidity: function getLastNode() constant returns(string[], uint256)
func (_Dnode *DnodeCaller) GetLastNode(opts *bind.CallOpts) ([]string, *big.Int, error) {
	var (
		ret0 = new([]string)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Dnode.contract.Call(opts, out, "getLastNode")
	return *ret0, *ret1, err
}

// GetLastNode is a free data retrieval call binding the contract method 0xd1d72e3e.
//
// Solidity: function getLastNode() constant returns(string[], uint256)
func (_Dnode *DnodeSession) GetLastNode() ([]string, *big.Int, error) {
	return _Dnode.Contract.GetLastNode(&_Dnode.CallOpts)
}

// GetLastNode is a free data retrieval call binding the contract method 0xd1d72e3e.
//
// Solidity: function getLastNode() constant returns(string[], uint256)
func (_Dnode *DnodeCallerSession) GetLastNode() ([]string, *big.Int, error) {
	return _Dnode.Contract.GetLastNode(&_Dnode.CallOpts)
}

// GetLastVersion is a free data retrieval call binding the contract method 0x558614c7.
//
// Solidity: function getLastVersion() constant returns(string, string)
func (_Dnode *DnodeCaller) GetLastVersion(opts *bind.CallOpts) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Dnode.contract.Call(opts, out, "getLastVersion")
	return *ret0, *ret1, err
}

// GetLastVersion is a free data retrieval call binding the contract method 0x558614c7.
//
// Solidity: function getLastVersion() constant returns(string, string)
func (_Dnode *DnodeSession) GetLastVersion() (string, string, error) {
	return _Dnode.Contract.GetLastVersion(&_Dnode.CallOpts)
}

// GetLastVersion is a free data retrieval call binding the contract method 0x558614c7.
//
// Solidity: function getLastVersion() constant returns(string, string)
func (_Dnode *DnodeCallerSession) GetLastVersion() (string, string, error) {
	return _Dnode.Contract.GetLastVersion(&_Dnode.CallOpts)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 idx) constant returns(string[], uint256)
func (_Dnode *DnodeCaller) GetNode(opts *bind.CallOpts, idx *big.Int) ([]string, *big.Int, error) {
	var (
		ret0 = new([]string)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Dnode.contract.Call(opts, out, "getNode", idx)
	return *ret0, *ret1, err
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 idx) constant returns(string[], uint256)
func (_Dnode *DnodeSession) GetNode(idx *big.Int) ([]string, *big.Int, error) {
	return _Dnode.Contract.GetNode(&_Dnode.CallOpts, idx)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 idx) constant returns(string[], uint256)
func (_Dnode *DnodeCallerSession) GetNode(idx *big.Int) ([]string, *big.Int, error) {
	return _Dnode.Contract.GetNode(&_Dnode.CallOpts, idx)
}

// GetVersion is a free data retrieval call binding the contract method 0xb88da759.
//
// Solidity: function getVersion(uint256 idx) constant returns(string, string)
func (_Dnode *DnodeCaller) GetVersion(opts *bind.CallOpts, idx *big.Int) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Dnode.contract.Call(opts, out, "getVersion", idx)
	return *ret0, *ret1, err
}

// GetVersion is a free data retrieval call binding the contract method 0xb88da759.
//
// Solidity: function getVersion(uint256 idx) constant returns(string, string)
func (_Dnode *DnodeSession) GetVersion(idx *big.Int) (string, string, error) {
	return _Dnode.Contract.GetVersion(&_Dnode.CallOpts, idx)
}

// GetVersion is a free data retrieval call binding the contract method 0xb88da759.
//
// Solidity: function getVersion(uint256 idx) constant returns(string, string)
func (_Dnode *DnodeCallerSession) GetVersion(idx *big.Int) (string, string, error) {
	return _Dnode.Contract.GetVersion(&_Dnode.CallOpts, idx)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dnode *DnodeCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dnode.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dnode *DnodeSession) IsOwner() (bool, error) {
	return _Dnode.Contract.IsOwner(&_Dnode.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dnode *DnodeCallerSession) IsOwner() (bool, error) {
	return _Dnode.Contract.IsOwner(&_Dnode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dnode *DnodeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dnode.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dnode *DnodeSession) Owner() (common.Address, error) {
	return _Dnode.Contract.Owner(&_Dnode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dnode *DnodeCallerSession) Owner() (common.Address, error) {
	return _Dnode.Contract.Owner(&_Dnode.CallOpts)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 idx) returns()
func (_Dnode *DnodeTransactor) Remove(opts *bind.TransactOpts, idx *big.Int) (*types.Transaction, error) {
	return _Dnode.contract.Transact(opts, "remove", idx)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 idx) returns()
func (_Dnode *DnodeSession) Remove(idx *big.Int) (*types.Transaction, error) {
	return _Dnode.Contract.Remove(&_Dnode.TransactOpts, idx)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 idx) returns()
func (_Dnode *DnodeTransactorSession) Remove(idx *big.Int) (*types.Transaction, error) {
	return _Dnode.Contract.Remove(&_Dnode.TransactOpts, idx)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dnode *DnodeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dnode.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dnode *DnodeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dnode.Contract.RenounceOwnership(&_Dnode.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dnode *DnodeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dnode.Contract.RenounceOwnership(&_Dnode.TransactOpts)
}

// SetLast is a paid mutator transaction binding the contract method 0x5c57bb89.
//
// Solidity: function setLast(uint256 idx) returns()
func (_Dnode *DnodeTransactor) SetLast(opts *bind.TransactOpts, idx *big.Int) (*types.Transaction, error) {
	return _Dnode.contract.Transact(opts, "setLast", idx)
}

// SetLast is a paid mutator transaction binding the contract method 0x5c57bb89.
//
// Solidity: function setLast(uint256 idx) returns()
func (_Dnode *DnodeSession) SetLast(idx *big.Int) (*types.Transaction, error) {
	return _Dnode.Contract.SetLast(&_Dnode.TransactOpts, idx)
}

// SetLast is a paid mutator transaction binding the contract method 0x5c57bb89.
//
// Solidity: function setLast(uint256 idx) returns()
func (_Dnode *DnodeTransactorSession) SetLast(idx *big.Int) (*types.Transaction, error) {
	return _Dnode.Contract.SetLast(&_Dnode.TransactOpts, idx)
}

// SetLastVersion is a paid mutator transaction binding the contract method 0xe80d8c38.
//
// Solidity: function setLastVersion(string ver, string hash) returns()
func (_Dnode *DnodeTransactor) SetLastVersion(opts *bind.TransactOpts, ver string, hash string) (*types.Transaction, error) {
	return _Dnode.contract.Transact(opts, "setLastVersion", ver, hash)
}

// SetLastVersion is a paid mutator transaction binding the contract method 0xe80d8c38.
//
// Solidity: function setLastVersion(string ver, string hash) returns()
func (_Dnode *DnodeSession) SetLastVersion(ver string, hash string) (*types.Transaction, error) {
	return _Dnode.Contract.SetLastVersion(&_Dnode.TransactOpts, ver, hash)
}

// SetLastVersion is a paid mutator transaction binding the contract method 0xe80d8c38.
//
// Solidity: function setLastVersion(string ver, string hash) returns()
func (_Dnode *DnodeTransactorSession) SetLastVersion(ver string, hash string) (*types.Transaction, error) {
	return _Dnode.Contract.SetLastVersion(&_Dnode.TransactOpts, ver, hash)
}

// SetVersion is a paid mutator transaction binding the contract method 0x69cf7c9a.
//
// Solidity: function setVersion(uint256 idx, string ver, string hash) returns()
func (_Dnode *DnodeTransactor) SetVersion(opts *bind.TransactOpts, idx *big.Int, ver string, hash string) (*types.Transaction, error) {
	return _Dnode.contract.Transact(opts, "setVersion", idx, ver, hash)
}

// SetVersion is a paid mutator transaction binding the contract method 0x69cf7c9a.
//
// Solidity: function setVersion(uint256 idx, string ver, string hash) returns()
func (_Dnode *DnodeSession) SetVersion(idx *big.Int, ver string, hash string) (*types.Transaction, error) {
	return _Dnode.Contract.SetVersion(&_Dnode.TransactOpts, idx, ver, hash)
}

// SetVersion is a paid mutator transaction binding the contract method 0x69cf7c9a.
//
// Solidity: function setVersion(uint256 idx, string ver, string hash) returns()
func (_Dnode *DnodeTransactorSession) SetVersion(idx *big.Int, ver string, hash string) (*types.Transaction, error) {
	return _Dnode.Contract.SetVersion(&_Dnode.TransactOpts, idx, ver, hash)
}

// Store is a paid mutator transaction binding the contract method 0x131a0680.
//
// Solidity: function store(string swarm) returns()
func (_Dnode *DnodeTransactor) Store(opts *bind.TransactOpts, swarm string) (*types.Transaction, error) {
	return _Dnode.contract.Transact(opts, "store", swarm)
}

// Store is a paid mutator transaction binding the contract method 0x131a0680.
//
// Solidity: function store(string swarm) returns()
func (_Dnode *DnodeSession) Store(swarm string) (*types.Transaction, error) {
	return _Dnode.Contract.Store(&_Dnode.TransactOpts, swarm)
}

// Store is a paid mutator transaction binding the contract method 0x131a0680.
//
// Solidity: function store(string swarm) returns()
func (_Dnode *DnodeTransactorSession) Store(swarm string) (*types.Transaction, error) {
	return _Dnode.Contract.Store(&_Dnode.TransactOpts, swarm)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dnode *DnodeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dnode.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dnode *DnodeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dnode.Contract.TransferOwnership(&_Dnode.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dnode *DnodeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dnode.Contract.TransferOwnership(&_Dnode.TransactOpts, newOwner)
}

// DnodeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dnode contract.
type DnodeOwnershipTransferredIterator struct {
	Event *DnodeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DnodeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DnodeOwnershipTransferred)
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
		it.Event = new(DnodeOwnershipTransferred)
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
func (it *DnodeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DnodeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DnodeOwnershipTransferred represents a OwnershipTransferred event raised by the Dnode contract.
type DnodeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dnode *DnodeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DnodeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dnode.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DnodeOwnershipTransferredIterator{contract: _Dnode.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dnode *DnodeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DnodeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dnode.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DnodeOwnershipTransferred)
				if err := _Dnode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dnode *DnodeFilterer) ParseOwnershipTransferred(log types.Log) (*DnodeOwnershipTransferred, error) {
	event := new(DnodeOwnershipTransferred)
	if err := _Dnode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
