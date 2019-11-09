// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dtag

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

// DtagABI is the input ABI used to generate the binding from.
const DtagABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sub\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"addTagId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sub\",\"type\":\"string\"}],\"name\":\"checkTag\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sub\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"addTagMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sub\",\"type\":\"string\"}],\"name\":\"fixTagIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sub\",\"type\":\"string\"}],\"name\":\"getTagIds\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sub\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"ids\",\"type\":\"string[]\"}],\"name\":\"setTagIds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sub\",\"type\":\"string\"}],\"name\":\"checkTagIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sub\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"ids\",\"type\":\"string[]\"}],\"name\":\"addTagIds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"tag\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sub\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"ids\",\"type\":\"string[]\"}],\"name\":\"replaceTagIds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"msgAddr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Dtag is an auto generated Go binding around an Ethereum contract.
type Dtag struct {
	DtagCaller     // Read-only binding to the contract
	DtagTransactor // Write-only binding to the contract
	DtagFilterer   // Log filterer for contract events
}

// DtagCaller is an auto generated read-only Go binding around an Ethereum contract.
type DtagCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtagTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DtagTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtagFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DtagFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DtagSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DtagSession struct {
	Contract     *Dtag             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DtagCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DtagCallerSession struct {
	Contract *DtagCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DtagTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DtagTransactorSession struct {
	Contract     *DtagTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DtagRaw is an auto generated low-level Go binding around an Ethereum contract.
type DtagRaw struct {
	Contract *Dtag // Generic contract binding to access the raw methods on
}

// DtagCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DtagCallerRaw struct {
	Contract *DtagCaller // Generic read-only contract binding to access the raw methods on
}

// DtagTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DtagTransactorRaw struct {
	Contract *DtagTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDtag creates a new instance of Dtag, bound to a specific deployed contract.
func NewDtag(address common.Address, backend bind.ContractBackend) (*Dtag, error) {
	contract, err := bindDtag(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dtag{DtagCaller: DtagCaller{contract: contract}, DtagTransactor: DtagTransactor{contract: contract}, DtagFilterer: DtagFilterer{contract: contract}}, nil
}

// NewDtagCaller creates a new read-only instance of Dtag, bound to a specific deployed contract.
func NewDtagCaller(address common.Address, caller bind.ContractCaller) (*DtagCaller, error) {
	contract, err := bindDtag(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DtagCaller{contract: contract}, nil
}

// NewDtagTransactor creates a new write-only instance of Dtag, bound to a specific deployed contract.
func NewDtagTransactor(address common.Address, transactor bind.ContractTransactor) (*DtagTransactor, error) {
	contract, err := bindDtag(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DtagTransactor{contract: contract}, nil
}

// NewDtagFilterer creates a new log filterer instance of Dtag, bound to a specific deployed contract.
func NewDtagFilterer(address common.Address, filterer bind.ContractFilterer) (*DtagFilterer, error) {
	contract, err := bindDtag(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DtagFilterer{contract: contract}, nil
}

// bindDtag binds a generic wrapper to an already deployed contract.
func bindDtag(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DtagABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dtag *DtagRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dtag.Contract.DtagCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dtag *DtagRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtag.Contract.DtagTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dtag *DtagRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dtag.Contract.DtagTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dtag *DtagCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dtag.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dtag *DtagTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtag.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dtag *DtagTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dtag.Contract.contract.Transact(opts, method, params...)
}

// CheckTag is a free data retrieval call binding the contract method 0x16896051.
//
// Solidity: function checkTag(string tag, string sub) constant returns(bool)
func (_Dtag *DtagCaller) CheckTag(opts *bind.CallOpts, tag string, sub string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dtag.contract.Call(opts, out, "checkTag", tag, sub)
	return *ret0, err
}

// CheckTag is a free data retrieval call binding the contract method 0x16896051.
//
// Solidity: function checkTag(string tag, string sub) constant returns(bool)
func (_Dtag *DtagSession) CheckTag(tag string, sub string) (bool, error) {
	return _Dtag.Contract.CheckTag(&_Dtag.CallOpts, tag, sub)
}

// CheckTag is a free data retrieval call binding the contract method 0x16896051.
//
// Solidity: function checkTag(string tag, string sub) constant returns(bool)
func (_Dtag *DtagCallerSession) CheckTag(tag string, sub string) (bool, error) {
	return _Dtag.Contract.CheckTag(&_Dtag.CallOpts, tag, sub)
}

// CheckTagIds is a free data retrieval call binding the contract method 0xaa79bd66.
//
// Solidity: function checkTagIds(string tag, string sub) constant returns(uint256, bool)
func (_Dtag *DtagCaller) CheckTagIds(opts *bind.CallOpts, tag string, sub string) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Dtag.contract.Call(opts, out, "checkTagIds", tag, sub)
	return *ret0, *ret1, err
}

// CheckTagIds is a free data retrieval call binding the contract method 0xaa79bd66.
//
// Solidity: function checkTagIds(string tag, string sub) constant returns(uint256, bool)
func (_Dtag *DtagSession) CheckTagIds(tag string, sub string) (*big.Int, bool, error) {
	return _Dtag.Contract.CheckTagIds(&_Dtag.CallOpts, tag, sub)
}

// CheckTagIds is a free data retrieval call binding the contract method 0xaa79bd66.
//
// Solidity: function checkTagIds(string tag, string sub) constant returns(uint256, bool)
func (_Dtag *DtagCallerSession) CheckTagIds(tag string, sub string) (*big.Int, bool, error) {
	return _Dtag.Contract.CheckTagIds(&_Dtag.CallOpts, tag, sub)
}

// GetTagIds is a free data retrieval call binding the contract method 0x8bb125e9.
//
// Solidity: function getTagIds(string tag, string sub) constant returns(string[])
func (_Dtag *DtagCaller) GetTagIds(opts *bind.CallOpts, tag string, sub string) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Dtag.contract.Call(opts, out, "getTagIds", tag, sub)
	return *ret0, err
}

// GetTagIds is a free data retrieval call binding the contract method 0x8bb125e9.
//
// Solidity: function getTagIds(string tag, string sub) constant returns(string[])
func (_Dtag *DtagSession) GetTagIds(tag string, sub string) ([]string, error) {
	return _Dtag.Contract.GetTagIds(&_Dtag.CallOpts, tag, sub)
}

// GetTagIds is a free data retrieval call binding the contract method 0x8bb125e9.
//
// Solidity: function getTagIds(string tag, string sub) constant returns(string[])
func (_Dtag *DtagCallerSession) GetTagIds(tag string, sub string) ([]string, error) {
	return _Dtag.Contract.GetTagIds(&_Dtag.CallOpts, tag, sub)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dtag *DtagCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dtag.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dtag *DtagSession) IsOwner() (bool, error) {
	return _Dtag.Contract.IsOwner(&_Dtag.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dtag *DtagCallerSession) IsOwner() (bool, error) {
	return _Dtag.Contract.IsOwner(&_Dtag.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dtag *DtagCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dtag.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dtag *DtagSession) Owner() (common.Address, error) {
	return _Dtag.Contract.Owner(&_Dtag.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dtag *DtagCallerSession) Owner() (common.Address, error) {
	return _Dtag.Contract.Owner(&_Dtag.CallOpts)
}

// AddTagId is a paid mutator transaction binding the contract method 0x16017c65.
//
// Solidity: function addTagId(string tag, string sub, string id) returns()
func (_Dtag *DtagTransactor) AddTagId(opts *bind.TransactOpts, tag string, sub string, id string) (*types.Transaction, error) {
	return _Dtag.contract.Transact(opts, "addTagId", tag, sub, id)
}

// AddTagId is a paid mutator transaction binding the contract method 0x16017c65.
//
// Solidity: function addTagId(string tag, string sub, string id) returns()
func (_Dtag *DtagSession) AddTagId(tag string, sub string, id string) (*types.Transaction, error) {
	return _Dtag.Contract.AddTagId(&_Dtag.TransactOpts, tag, sub, id)
}

// AddTagId is a paid mutator transaction binding the contract method 0x16017c65.
//
// Solidity: function addTagId(string tag, string sub, string id) returns()
func (_Dtag *DtagTransactorSession) AddTagId(tag string, sub string, id string) (*types.Transaction, error) {
	return _Dtag.Contract.AddTagId(&_Dtag.TransactOpts, tag, sub, id)
}

// AddTagIds is a paid mutator transaction binding the contract method 0xe314604e.
//
// Solidity: function addTagIds(string tag, string sub, string[] ids) returns()
func (_Dtag *DtagTransactor) AddTagIds(opts *bind.TransactOpts, tag string, sub string, ids []string) (*types.Transaction, error) {
	return _Dtag.contract.Transact(opts, "addTagIds", tag, sub, ids)
}

// AddTagIds is a paid mutator transaction binding the contract method 0xe314604e.
//
// Solidity: function addTagIds(string tag, string sub, string[] ids) returns()
func (_Dtag *DtagSession) AddTagIds(tag string, sub string, ids []string) (*types.Transaction, error) {
	return _Dtag.Contract.AddTagIds(&_Dtag.TransactOpts, tag, sub, ids)
}

// AddTagIds is a paid mutator transaction binding the contract method 0xe314604e.
//
// Solidity: function addTagIds(string tag, string sub, string[] ids) returns()
func (_Dtag *DtagTransactorSession) AddTagIds(tag string, sub string, ids []string) (*types.Transaction, error) {
	return _Dtag.Contract.AddTagIds(&_Dtag.TransactOpts, tag, sub, ids)
}

// AddTagMessage is a paid mutator transaction binding the contract method 0x2d2bb175.
//
// Solidity: function addTagMessage(string tag, string sub, string id, string content, string version) returns(bool)
func (_Dtag *DtagTransactor) AddTagMessage(opts *bind.TransactOpts, tag string, sub string, id string, content string, version string) (*types.Transaction, error) {
	return _Dtag.contract.Transact(opts, "addTagMessage", tag, sub, id, content, version)
}

// AddTagMessage is a paid mutator transaction binding the contract method 0x2d2bb175.
//
// Solidity: function addTagMessage(string tag, string sub, string id, string content, string version) returns(bool)
func (_Dtag *DtagSession) AddTagMessage(tag string, sub string, id string, content string, version string) (*types.Transaction, error) {
	return _Dtag.Contract.AddTagMessage(&_Dtag.TransactOpts, tag, sub, id, content, version)
}

// AddTagMessage is a paid mutator transaction binding the contract method 0x2d2bb175.
//
// Solidity: function addTagMessage(string tag, string sub, string id, string content, string version) returns(bool)
func (_Dtag *DtagTransactorSession) AddTagMessage(tag string, sub string, id string, content string, version string) (*types.Transaction, error) {
	return _Dtag.Contract.AddTagMessage(&_Dtag.TransactOpts, tag, sub, id, content, version)
}

// FixTagIds is a paid mutator transaction binding the contract method 0x44dfc0e1.
//
// Solidity: function fixTagIds(string tag, string sub) returns(uint256 _size)
func (_Dtag *DtagTransactor) FixTagIds(opts *bind.TransactOpts, tag string, sub string) (*types.Transaction, error) {
	return _Dtag.contract.Transact(opts, "fixTagIds", tag, sub)
}

// FixTagIds is a paid mutator transaction binding the contract method 0x44dfc0e1.
//
// Solidity: function fixTagIds(string tag, string sub) returns(uint256 _size)
func (_Dtag *DtagSession) FixTagIds(tag string, sub string) (*types.Transaction, error) {
	return _Dtag.Contract.FixTagIds(&_Dtag.TransactOpts, tag, sub)
}

// FixTagIds is a paid mutator transaction binding the contract method 0x44dfc0e1.
//
// Solidity: function fixTagIds(string tag, string sub) returns(uint256 _size)
func (_Dtag *DtagTransactorSession) FixTagIds(tag string, sub string) (*types.Transaction, error) {
	return _Dtag.Contract.FixTagIds(&_Dtag.TransactOpts, tag, sub)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtag *DtagTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtag.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtag *DtagSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dtag.Contract.RenounceOwnership(&_Dtag.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dtag *DtagTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dtag.Contract.RenounceOwnership(&_Dtag.TransactOpts)
}

// ReplaceTagIds is a paid mutator transaction binding the contract method 0xe748442e.
//
// Solidity: function replaceTagIds(string tag, string sub, string[] ids) returns()
func (_Dtag *DtagTransactor) ReplaceTagIds(opts *bind.TransactOpts, tag string, sub string, ids []string) (*types.Transaction, error) {
	return _Dtag.contract.Transact(opts, "replaceTagIds", tag, sub, ids)
}

// ReplaceTagIds is a paid mutator transaction binding the contract method 0xe748442e.
//
// Solidity: function replaceTagIds(string tag, string sub, string[] ids) returns()
func (_Dtag *DtagSession) ReplaceTagIds(tag string, sub string, ids []string) (*types.Transaction, error) {
	return _Dtag.Contract.ReplaceTagIds(&_Dtag.TransactOpts, tag, sub, ids)
}

// ReplaceTagIds is a paid mutator transaction binding the contract method 0xe748442e.
//
// Solidity: function replaceTagIds(string tag, string sub, string[] ids) returns()
func (_Dtag *DtagTransactorSession) ReplaceTagIds(tag string, sub string, ids []string) (*types.Transaction, error) {
	return _Dtag.Contract.ReplaceTagIds(&_Dtag.TransactOpts, tag, sub, ids)
}

// SetTagIds is a paid mutator transaction binding the contract method 0x8d540dfa.
//
// Solidity: function setTagIds(string tag, string sub, string[] ids) returns()
func (_Dtag *DtagTransactor) SetTagIds(opts *bind.TransactOpts, tag string, sub string, ids []string) (*types.Transaction, error) {
	return _Dtag.contract.Transact(opts, "setTagIds", tag, sub, ids)
}

// SetTagIds is a paid mutator transaction binding the contract method 0x8d540dfa.
//
// Solidity: function setTagIds(string tag, string sub, string[] ids) returns()
func (_Dtag *DtagSession) SetTagIds(tag string, sub string, ids []string) (*types.Transaction, error) {
	return _Dtag.Contract.SetTagIds(&_Dtag.TransactOpts, tag, sub, ids)
}

// SetTagIds is a paid mutator transaction binding the contract method 0x8d540dfa.
//
// Solidity: function setTagIds(string tag, string sub, string[] ids) returns()
func (_Dtag *DtagTransactorSession) SetTagIds(tag string, sub string, ids []string) (*types.Transaction, error) {
	return _Dtag.Contract.SetTagIds(&_Dtag.TransactOpts, tag, sub, ids)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtag *DtagTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dtag.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtag *DtagSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dtag.Contract.TransferOwnership(&_Dtag.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dtag *DtagTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dtag.Contract.TransferOwnership(&_Dtag.TransactOpts, newOwner)
}

// DtagOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dtag contract.
type DtagOwnershipTransferredIterator struct {
	Event *DtagOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DtagOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtagOwnershipTransferred)
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
		it.Event = new(DtagOwnershipTransferred)
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
func (it *DtagOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtagOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtagOwnershipTransferred represents a OwnershipTransferred event raised by the Dtag contract.
type DtagOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dtag *DtagFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DtagOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dtag.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DtagOwnershipTransferredIterator{contract: _Dtag.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dtag *DtagFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DtagOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dtag.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtagOwnershipTransferred)
				if err := _Dtag.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dtag *DtagFilterer) ParseOwnershipTransferred(log types.Log) (*DtagOwnershipTransferred, error) {
	event := new(DtagOwnershipTransferred)
	if err := _Dtag.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
