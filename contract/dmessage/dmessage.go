// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dmessage

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

// DmessageABI is the input ABI used to generate the binding from.
const DmessageABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getMsgLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"getMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structDMessage.Message\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWriter\",\"type\":\"address\"}],\"name\":\"increasedWritership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"ids\",\"type\":\"string[]\"}],\"name\":\"checkAllMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"checkMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"delMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structDMessage.Message\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceWritership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"writer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getMessages\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structDMessage.Message[]\",\"name\":\"_value\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isWriter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMsgIds\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"updateMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"ids\",\"type\":\"string[]\"}],\"name\":\"getIdsMessages\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structDMessage.Message[]\",\"name\":\"_value\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"ids\",\"type\":\"string[]\"}],\"name\":\"checkMessages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"writers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_value\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getMsgId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_value\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldWriter\",\"type\":\"address\"}],\"name\":\"decreaseWritership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"addMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"getMessageValue\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newWriter\",\"type\":\"address\"}],\"name\":\"WritershipIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldWriter\",\"type\":\"address\"}],\"name\":\"WritershipDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Dmessage is an auto generated Go binding around an Ethereum contract.
type Dmessage struct {
	DmessageCaller     // Read-only binding to the contract
	DmessageTransactor // Write-only binding to the contract
	DmessageFilterer   // Log filterer for contract events
}

// DmessageCaller is an auto generated read-only Go binding around an Ethereum contract.
type DmessageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DmessageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DmessageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DmessageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DmessageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DmessageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DmessageSession struct {
	Contract     *Dmessage         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DmessageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DmessageCallerSession struct {
	Contract *DmessageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DmessageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DmessageTransactorSession struct {
	Contract     *DmessageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DmessageRaw is an auto generated low-level Go binding around an Ethereum contract.
type DmessageRaw struct {
	Contract *Dmessage // Generic contract binding to access the raw methods on
}

// DmessageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DmessageCallerRaw struct {
	Contract *DmessageCaller // Generic read-only contract binding to access the raw methods on
}

// DmessageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DmessageTransactorRaw struct {
	Contract *DmessageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDmessage creates a new instance of Dmessage, bound to a specific deployed contract.
func NewDmessage(address common.Address, backend bind.ContractBackend) (*Dmessage, error) {
	contract, err := bindDmessage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dmessage{DmessageCaller: DmessageCaller{contract: contract}, DmessageTransactor: DmessageTransactor{contract: contract}, DmessageFilterer: DmessageFilterer{contract: contract}}, nil
}

// NewDmessageCaller creates a new read-only instance of Dmessage, bound to a specific deployed contract.
func NewDmessageCaller(address common.Address, caller bind.ContractCaller) (*DmessageCaller, error) {
	contract, err := bindDmessage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DmessageCaller{contract: contract}, nil
}

// NewDmessageTransactor creates a new write-only instance of Dmessage, bound to a specific deployed contract.
func NewDmessageTransactor(address common.Address, transactor bind.ContractTransactor) (*DmessageTransactor, error) {
	contract, err := bindDmessage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DmessageTransactor{contract: contract}, nil
}

// NewDmessageFilterer creates a new log filterer instance of Dmessage, bound to a specific deployed contract.
func NewDmessageFilterer(address common.Address, filterer bind.ContractFilterer) (*DmessageFilterer, error) {
	contract, err := bindDmessage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DmessageFilterer{contract: contract}, nil
}

// bindDmessage binds a generic wrapper to an already deployed contract.
func bindDmessage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DmessageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dmessage *DmessageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dmessage.Contract.DmessageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dmessage *DmessageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dmessage.Contract.DmessageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dmessage *DmessageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dmessage.Contract.DmessageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dmessage *DmessageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Dmessage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dmessage *DmessageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dmessage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dmessage *DmessageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dmessage.Contract.contract.Transact(opts, method, params...)
}

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Id      string
	Content string
	Version string
}

// CheckAllMessage is a free data retrieval call binding the contract method 0x16313e92.
//
// Solidity: function checkAllMessage(string[] ids) constant returns(bool)
func (_Dmessage *DmessageCaller) CheckAllMessage(opts *bind.CallOpts, ids []string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dmessage.contract.Call(opts, out, "checkAllMessage", ids)
	return *ret0, err
}

// CheckAllMessage is a free data retrieval call binding the contract method 0x16313e92.
//
// Solidity: function checkAllMessage(string[] ids) constant returns(bool)
func (_Dmessage *DmessageSession) CheckAllMessage(ids []string) (bool, error) {
	return _Dmessage.Contract.CheckAllMessage(&_Dmessage.CallOpts, ids)
}

// CheckAllMessage is a free data retrieval call binding the contract method 0x16313e92.
//
// Solidity: function checkAllMessage(string[] ids) constant returns(bool)
func (_Dmessage *DmessageCallerSession) CheckAllMessage(ids []string) (bool, error) {
	return _Dmessage.Contract.CheckAllMessage(&_Dmessage.CallOpts, ids)
}

// CheckMessage is a free data retrieval call binding the contract method 0x1873ec68.
//
// Solidity: function checkMessage(string id) constant returns(bool)
func (_Dmessage *DmessageCaller) CheckMessage(opts *bind.CallOpts, id string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dmessage.contract.Call(opts, out, "checkMessage", id)
	return *ret0, err
}

// CheckMessage is a free data retrieval call binding the contract method 0x1873ec68.
//
// Solidity: function checkMessage(string id) constant returns(bool)
func (_Dmessage *DmessageSession) CheckMessage(id string) (bool, error) {
	return _Dmessage.Contract.CheckMessage(&_Dmessage.CallOpts, id)
}

// CheckMessage is a free data retrieval call binding the contract method 0x1873ec68.
//
// Solidity: function checkMessage(string id) constant returns(bool)
func (_Dmessage *DmessageCallerSession) CheckMessage(id string) (bool, error) {
	return _Dmessage.Contract.CheckMessage(&_Dmessage.CallOpts, id)
}

// CheckMessages is a free data retrieval call binding the contract method 0x834daf0c.
//
// Solidity: function checkMessages(string[] ids) constant returns(uint256, bool)
func (_Dmessage *DmessageCaller) CheckMessages(opts *bind.CallOpts, ids []string) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Dmessage.contract.Call(opts, out, "checkMessages", ids)
	return *ret0, *ret1, err
}

// CheckMessages is a free data retrieval call binding the contract method 0x834daf0c.
//
// Solidity: function checkMessages(string[] ids) constant returns(uint256, bool)
func (_Dmessage *DmessageSession) CheckMessages(ids []string) (*big.Int, bool, error) {
	return _Dmessage.Contract.CheckMessages(&_Dmessage.CallOpts, ids)
}

// CheckMessages is a free data retrieval call binding the contract method 0x834daf0c.
//
// Solidity: function checkMessages(string[] ids) constant returns(uint256, bool)
func (_Dmessage *DmessageCallerSession) CheckMessages(ids []string) (*big.Int, bool, error) {
	return _Dmessage.Contract.CheckMessages(&_Dmessage.CallOpts, ids)
}

// GetIdsMessages is a free data retrieval call binding the contract method 0x6c4ef994.
//
// Solidity: function getIdsMessages(string[] ids) constant returns([]Struct0 _value, uint256 _size)
func (_Dmessage *DmessageCaller) GetIdsMessages(opts *bind.CallOpts, ids []string) (struct {
	Value []Struct0
	Size  *big.Int
}, error) {
	ret := new(struct {
		Value []Struct0
		Size  *big.Int
	})
	out := ret
	err := _Dmessage.contract.Call(opts, out, "getIdsMessages", ids)
	return *ret, err
}

// GetIdsMessages is a free data retrieval call binding the contract method 0x6c4ef994.
//
// Solidity: function getIdsMessages(string[] ids) constant returns([]Struct0 _value, uint256 _size)
func (_Dmessage *DmessageSession) GetIdsMessages(ids []string) (struct {
	Value []Struct0
	Size  *big.Int
}, error) {
	return _Dmessage.Contract.GetIdsMessages(&_Dmessage.CallOpts, ids)
}

// GetIdsMessages is a free data retrieval call binding the contract method 0x6c4ef994.
//
// Solidity: function getIdsMessages(string[] ids) constant returns([]Struct0 _value, uint256 _size)
func (_Dmessage *DmessageCallerSession) GetIdsMessages(ids []string) (struct {
	Value []Struct0
	Size  *big.Int
}, error) {
	return _Dmessage.Contract.GetIdsMessages(&_Dmessage.CallOpts, ids)
}

// GetMessage is a free data retrieval call binding the contract method 0x0cc4e8d8.
//
// Solidity: function getMessage(string id) constant returns(Struct0)
func (_Dmessage *DmessageCaller) GetMessage(opts *bind.CallOpts, id string) (Struct0, error) {
	var (
		ret0 = new(Struct0)
	)
	out := ret0
	err := _Dmessage.contract.Call(opts, out, "getMessage", id)
	return *ret0, err
}

// GetMessage is a free data retrieval call binding the contract method 0x0cc4e8d8.
//
// Solidity: function getMessage(string id) constant returns(Struct0)
func (_Dmessage *DmessageSession) GetMessage(id string) (Struct0, error) {
	return _Dmessage.Contract.GetMessage(&_Dmessage.CallOpts, id)
}

// GetMessage is a free data retrieval call binding the contract method 0x0cc4e8d8.
//
// Solidity: function getMessage(string id) constant returns(Struct0)
func (_Dmessage *DmessageCallerSession) GetMessage(id string) (Struct0, error) {
	return _Dmessage.Contract.GetMessage(&_Dmessage.CallOpts, id)
}

// GetMessageValue is a free data retrieval call binding the contract method 0xc57bc642.
//
// Solidity: function getMessageValue(string id) constant returns(string, string, string)
func (_Dmessage *DmessageCaller) GetMessageValue(opts *bind.CallOpts, id string) (string, string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Dmessage.contract.Call(opts, out, "getMessageValue", id)
	return *ret0, *ret1, *ret2, err
}

// GetMessageValue is a free data retrieval call binding the contract method 0xc57bc642.
//
// Solidity: function getMessageValue(string id) constant returns(string, string, string)
func (_Dmessage *DmessageSession) GetMessageValue(id string) (string, string, string, error) {
	return _Dmessage.Contract.GetMessageValue(&_Dmessage.CallOpts, id)
}

// GetMessageValue is a free data retrieval call binding the contract method 0xc57bc642.
//
// Solidity: function getMessageValue(string id) constant returns(string, string, string)
func (_Dmessage *DmessageCallerSession) GetMessageValue(id string) (string, string, string, error) {
	return _Dmessage.Contract.GetMessageValue(&_Dmessage.CallOpts, id)
}

// GetMessages is a free data retrieval call binding the contract method 0x4be8362d.
//
// Solidity: function getMessages(uint256 start, uint256 limit) constant returns([]Struct0 _value, uint256 _size)
func (_Dmessage *DmessageCaller) GetMessages(opts *bind.CallOpts, start *big.Int, limit *big.Int) (struct {
	Value []Struct0
	Size  *big.Int
}, error) {
	ret := new(struct {
		Value []Struct0
		Size  *big.Int
	})
	out := ret
	err := _Dmessage.contract.Call(opts, out, "getMessages", start, limit)
	return *ret, err
}

// GetMessages is a free data retrieval call binding the contract method 0x4be8362d.
//
// Solidity: function getMessages(uint256 start, uint256 limit) constant returns([]Struct0 _value, uint256 _size)
func (_Dmessage *DmessageSession) GetMessages(start *big.Int, limit *big.Int) (struct {
	Value []Struct0
	Size  *big.Int
}, error) {
	return _Dmessage.Contract.GetMessages(&_Dmessage.CallOpts, start, limit)
}

// GetMessages is a free data retrieval call binding the contract method 0x4be8362d.
//
// Solidity: function getMessages(uint256 start, uint256 limit) constant returns([]Struct0 _value, uint256 _size)
func (_Dmessage *DmessageCallerSession) GetMessages(start *big.Int, limit *big.Int) (struct {
	Value []Struct0
	Size  *big.Int
}, error) {
	return _Dmessage.Contract.GetMessages(&_Dmessage.CallOpts, start, limit)
}

// GetMsgId is a free data retrieval call binding the contract method 0x95cb2638.
//
// Solidity: function getMsgId(uint256 _index) constant returns(string _value)
func (_Dmessage *DmessageCaller) GetMsgId(opts *bind.CallOpts, _index *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Dmessage.contract.Call(opts, out, "getMsgId", _index)
	return *ret0, err
}

// GetMsgId is a free data retrieval call binding the contract method 0x95cb2638.
//
// Solidity: function getMsgId(uint256 _index) constant returns(string _value)
func (_Dmessage *DmessageSession) GetMsgId(_index *big.Int) (string, error) {
	return _Dmessage.Contract.GetMsgId(&_Dmessage.CallOpts, _index)
}

// GetMsgId is a free data retrieval call binding the contract method 0x95cb2638.
//
// Solidity: function getMsgId(uint256 _index) constant returns(string _value)
func (_Dmessage *DmessageCallerSession) GetMsgId(_index *big.Int) (string, error) {
	return _Dmessage.Contract.GetMsgId(&_Dmessage.CallOpts, _index)
}

// GetMsgIds is a free data retrieval call binding the contract method 0x68bfc43c.
//
// Solidity: function getMsgIds() constant returns(string[])
func (_Dmessage *DmessageCaller) GetMsgIds(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Dmessage.contract.Call(opts, out, "getMsgIds")
	return *ret0, err
}

// GetMsgIds is a free data retrieval call binding the contract method 0x68bfc43c.
//
// Solidity: function getMsgIds() constant returns(string[])
func (_Dmessage *DmessageSession) GetMsgIds() ([]string, error) {
	return _Dmessage.Contract.GetMsgIds(&_Dmessage.CallOpts)
}

// GetMsgIds is a free data retrieval call binding the contract method 0x68bfc43c.
//
// Solidity: function getMsgIds() constant returns(string[])
func (_Dmessage *DmessageCallerSession) GetMsgIds() ([]string, error) {
	return _Dmessage.Contract.GetMsgIds(&_Dmessage.CallOpts)
}

// GetMsgLength is a free data retrieval call binding the contract method 0x085c00e1.
//
// Solidity: function getMsgLength() constant returns(uint256)
func (_Dmessage *DmessageCaller) GetMsgLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dmessage.contract.Call(opts, out, "getMsgLength")
	return *ret0, err
}

// GetMsgLength is a free data retrieval call binding the contract method 0x085c00e1.
//
// Solidity: function getMsgLength() constant returns(uint256)
func (_Dmessage *DmessageSession) GetMsgLength() (*big.Int, error) {
	return _Dmessage.Contract.GetMsgLength(&_Dmessage.CallOpts)
}

// GetMsgLength is a free data retrieval call binding the contract method 0x085c00e1.
//
// Solidity: function getMsgLength() constant returns(uint256)
func (_Dmessage *DmessageCallerSession) GetMsgLength() (*big.Int, error) {
	return _Dmessage.Contract.GetMsgLength(&_Dmessage.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dmessage *DmessageCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dmessage.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dmessage *DmessageSession) IsOwner() (bool, error) {
	return _Dmessage.Contract.IsOwner(&_Dmessage.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Dmessage *DmessageCallerSession) IsOwner() (bool, error) {
	return _Dmessage.Contract.IsOwner(&_Dmessage.CallOpts)
}

// IsWriter is a free data retrieval call binding the contract method 0x4d6ee9fd.
//
// Solidity: function isWriter() constant returns(bool)
func (_Dmessage *DmessageCaller) IsWriter(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dmessage.contract.Call(opts, out, "isWriter")
	return *ret0, err
}

// IsWriter is a free data retrieval call binding the contract method 0x4d6ee9fd.
//
// Solidity: function isWriter() constant returns(bool)
func (_Dmessage *DmessageSession) IsWriter() (bool, error) {
	return _Dmessage.Contract.IsWriter(&_Dmessage.CallOpts)
}

// IsWriter is a free data retrieval call binding the contract method 0x4d6ee9fd.
//
// Solidity: function isWriter() constant returns(bool)
func (_Dmessage *DmessageCallerSession) IsWriter() (bool, error) {
	return _Dmessage.Contract.IsWriter(&_Dmessage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dmessage *DmessageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dmessage.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dmessage *DmessageSession) Owner() (common.Address, error) {
	return _Dmessage.Contract.Owner(&_Dmessage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Dmessage *DmessageCallerSession) Owner() (common.Address, error) {
	return _Dmessage.Contract.Owner(&_Dmessage.CallOpts)
}

// Writer is a free data retrieval call binding the contract method 0x453a2abc.
//
// Solidity: function writer() constant returns(address)
func (_Dmessage *DmessageCaller) Writer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dmessage.contract.Call(opts, out, "writer")
	return *ret0, err
}

// Writer is a free data retrieval call binding the contract method 0x453a2abc.
//
// Solidity: function writer() constant returns(address)
func (_Dmessage *DmessageSession) Writer() (common.Address, error) {
	return _Dmessage.Contract.Writer(&_Dmessage.CallOpts)
}

// Writer is a free data retrieval call binding the contract method 0x453a2abc.
//
// Solidity: function writer() constant returns(address)
func (_Dmessage *DmessageCallerSession) Writer() (common.Address, error) {
	return _Dmessage.Contract.Writer(&_Dmessage.CallOpts)
}

// Writers is a free data retrieval call binding the contract method 0x8b2f5369.
//
// Solidity: function writers() constant returns(address[] _value)
func (_Dmessage *DmessageCaller) Writers(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Dmessage.contract.Call(opts, out, "writers")
	return *ret0, err
}

// Writers is a free data retrieval call binding the contract method 0x8b2f5369.
//
// Solidity: function writers() constant returns(address[] _value)
func (_Dmessage *DmessageSession) Writers() ([]common.Address, error) {
	return _Dmessage.Contract.Writers(&_Dmessage.CallOpts)
}

// Writers is a free data retrieval call binding the contract method 0x8b2f5369.
//
// Solidity: function writers() constant returns(address[] _value)
func (_Dmessage *DmessageCallerSession) Writers() ([]common.Address, error) {
	return _Dmessage.Contract.Writers(&_Dmessage.CallOpts)
}

// AddMessage is a paid mutator transaction binding the contract method 0xb98464b3.
//
// Solidity: function addMessage(string id, string content, string version) returns(bool)
func (_Dmessage *DmessageTransactor) AddMessage(opts *bind.TransactOpts, id string, content string, version string) (*types.Transaction, error) {
	return _Dmessage.contract.Transact(opts, "addMessage", id, content, version)
}

// AddMessage is a paid mutator transaction binding the contract method 0xb98464b3.
//
// Solidity: function addMessage(string id, string content, string version) returns(bool)
func (_Dmessage *DmessageSession) AddMessage(id string, content string, version string) (*types.Transaction, error) {
	return _Dmessage.Contract.AddMessage(&_Dmessage.TransactOpts, id, content, version)
}

// AddMessage is a paid mutator transaction binding the contract method 0xb98464b3.
//
// Solidity: function addMessage(string id, string content, string version) returns(bool)
func (_Dmessage *DmessageTransactorSession) AddMessage(id string, content string, version string) (*types.Transaction, error) {
	return _Dmessage.Contract.AddMessage(&_Dmessage.TransactOpts, id, content, version)
}

// DecreaseWritership is a paid mutator transaction binding the contract method 0x9f217bdf.
//
// Solidity: function decreaseWritership(address oldWriter) returns()
func (_Dmessage *DmessageTransactor) DecreaseWritership(opts *bind.TransactOpts, oldWriter common.Address) (*types.Transaction, error) {
	return _Dmessage.contract.Transact(opts, "decreaseWritership", oldWriter)
}

// DecreaseWritership is a paid mutator transaction binding the contract method 0x9f217bdf.
//
// Solidity: function decreaseWritership(address oldWriter) returns()
func (_Dmessage *DmessageSession) DecreaseWritership(oldWriter common.Address) (*types.Transaction, error) {
	return _Dmessage.Contract.DecreaseWritership(&_Dmessage.TransactOpts, oldWriter)
}

// DecreaseWritership is a paid mutator transaction binding the contract method 0x9f217bdf.
//
// Solidity: function decreaseWritership(address oldWriter) returns()
func (_Dmessage *DmessageTransactorSession) DecreaseWritership(oldWriter common.Address) (*types.Transaction, error) {
	return _Dmessage.Contract.DecreaseWritership(&_Dmessage.TransactOpts, oldWriter)
}

// DelMessage is a paid mutator transaction binding the contract method 0x2309bc96.
//
// Solidity: function delMessage(string id) returns(Struct0, bool)
func (_Dmessage *DmessageTransactor) DelMessage(opts *bind.TransactOpts, id string) (*types.Transaction, error) {
	return _Dmessage.contract.Transact(opts, "delMessage", id)
}

// DelMessage is a paid mutator transaction binding the contract method 0x2309bc96.
//
// Solidity: function delMessage(string id) returns(Struct0, bool)
func (_Dmessage *DmessageSession) DelMessage(id string) (*types.Transaction, error) {
	return _Dmessage.Contract.DelMessage(&_Dmessage.TransactOpts, id)
}

// DelMessage is a paid mutator transaction binding the contract method 0x2309bc96.
//
// Solidity: function delMessage(string id) returns(Struct0, bool)
func (_Dmessage *DmessageTransactorSession) DelMessage(id string) (*types.Transaction, error) {
	return _Dmessage.Contract.DelMessage(&_Dmessage.TransactOpts, id)
}

// IncreasedWritership is a paid mutator transaction binding the contract method 0x14b4d411.
//
// Solidity: function increasedWritership(address newWriter) returns()
func (_Dmessage *DmessageTransactor) IncreasedWritership(opts *bind.TransactOpts, newWriter common.Address) (*types.Transaction, error) {
	return _Dmessage.contract.Transact(opts, "increasedWritership", newWriter)
}

// IncreasedWritership is a paid mutator transaction binding the contract method 0x14b4d411.
//
// Solidity: function increasedWritership(address newWriter) returns()
func (_Dmessage *DmessageSession) IncreasedWritership(newWriter common.Address) (*types.Transaction, error) {
	return _Dmessage.Contract.IncreasedWritership(&_Dmessage.TransactOpts, newWriter)
}

// IncreasedWritership is a paid mutator transaction binding the contract method 0x14b4d411.
//
// Solidity: function increasedWritership(address newWriter) returns()
func (_Dmessage *DmessageTransactorSession) IncreasedWritership(newWriter common.Address) (*types.Transaction, error) {
	return _Dmessage.Contract.IncreasedWritership(&_Dmessage.TransactOpts, newWriter)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dmessage *DmessageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dmessage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dmessage *DmessageSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dmessage.Contract.RenounceOwnership(&_Dmessage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dmessage *DmessageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dmessage.Contract.RenounceOwnership(&_Dmessage.TransactOpts)
}

// RenounceWritership is a paid mutator transaction binding the contract method 0x4075e83e.
//
// Solidity: function renounceWritership() returns()
func (_Dmessage *DmessageTransactor) RenounceWritership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dmessage.contract.Transact(opts, "renounceWritership")
}

// RenounceWritership is a paid mutator transaction binding the contract method 0x4075e83e.
//
// Solidity: function renounceWritership() returns()
func (_Dmessage *DmessageSession) RenounceWritership() (*types.Transaction, error) {
	return _Dmessage.Contract.RenounceWritership(&_Dmessage.TransactOpts)
}

// RenounceWritership is a paid mutator transaction binding the contract method 0x4075e83e.
//
// Solidity: function renounceWritership() returns()
func (_Dmessage *DmessageTransactorSession) RenounceWritership() (*types.Transaction, error) {
	return _Dmessage.Contract.RenounceWritership(&_Dmessage.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dmessage *DmessageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dmessage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dmessage *DmessageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dmessage.Contract.TransferOwnership(&_Dmessage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dmessage *DmessageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dmessage.Contract.TransferOwnership(&_Dmessage.TransactOpts, newOwner)
}

// UpdateMessage is a paid mutator transaction binding the contract method 0x6a44815e.
//
// Solidity: function updateMessage(string id, string content, string version) returns(bool)
func (_Dmessage *DmessageTransactor) UpdateMessage(opts *bind.TransactOpts, id string, content string, version string) (*types.Transaction, error) {
	return _Dmessage.contract.Transact(opts, "updateMessage", id, content, version)
}

// UpdateMessage is a paid mutator transaction binding the contract method 0x6a44815e.
//
// Solidity: function updateMessage(string id, string content, string version) returns(bool)
func (_Dmessage *DmessageSession) UpdateMessage(id string, content string, version string) (*types.Transaction, error) {
	return _Dmessage.Contract.UpdateMessage(&_Dmessage.TransactOpts, id, content, version)
}

// UpdateMessage is a paid mutator transaction binding the contract method 0x6a44815e.
//
// Solidity: function updateMessage(string id, string content, string version) returns(bool)
func (_Dmessage *DmessageTransactorSession) UpdateMessage(id string, content string, version string) (*types.Transaction, error) {
	return _Dmessage.Contract.UpdateMessage(&_Dmessage.TransactOpts, id, content, version)
}

// DmessageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dmessage contract.
type DmessageOwnershipTransferredIterator struct {
	Event *DmessageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DmessageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmessageOwnershipTransferred)
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
		it.Event = new(DmessageOwnershipTransferred)
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
func (it *DmessageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmessageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmessageOwnershipTransferred represents a OwnershipTransferred event raised by the Dmessage contract.
type DmessageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dmessage *DmessageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DmessageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dmessage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DmessageOwnershipTransferredIterator{contract: _Dmessage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dmessage *DmessageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DmessageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dmessage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmessageOwnershipTransferred)
				if err := _Dmessage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Dmessage *DmessageFilterer) ParseOwnershipTransferred(log types.Log) (*DmessageOwnershipTransferred, error) {
	event := new(DmessageOwnershipTransferred)
	if err := _Dmessage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DmessageWritershipDecreasedIterator is returned from FilterWritershipDecreased and is used to iterate over the raw logs and unpacked data for WritershipDecreased events raised by the Dmessage contract.
type DmessageWritershipDecreasedIterator struct {
	Event *DmessageWritershipDecreased // Event containing the contract specifics and raw log

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
func (it *DmessageWritershipDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmessageWritershipDecreased)
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
		it.Event = new(DmessageWritershipDecreased)
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
func (it *DmessageWritershipDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmessageWritershipDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmessageWritershipDecreased represents a WritershipDecreased event raised by the Dmessage contract.
type DmessageWritershipDecreased struct {
	OldWriter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWritershipDecreased is a free log retrieval operation binding the contract event 0xb24c36b70cc67df9ea22f053c15c4a41e7170d7be0637d2637a3445a266ac6af.
//
// Solidity: event WritershipDecreased(address indexed oldWriter)
func (_Dmessage *DmessageFilterer) FilterWritershipDecreased(opts *bind.FilterOpts, oldWriter []common.Address) (*DmessageWritershipDecreasedIterator, error) {

	var oldWriterRule []interface{}
	for _, oldWriterItem := range oldWriter {
		oldWriterRule = append(oldWriterRule, oldWriterItem)
	}

	logs, sub, err := _Dmessage.contract.FilterLogs(opts, "WritershipDecreased", oldWriterRule)
	if err != nil {
		return nil, err
	}
	return &DmessageWritershipDecreasedIterator{contract: _Dmessage.contract, event: "WritershipDecreased", logs: logs, sub: sub}, nil
}

// WatchWritershipDecreased is a free log subscription operation binding the contract event 0xb24c36b70cc67df9ea22f053c15c4a41e7170d7be0637d2637a3445a266ac6af.
//
// Solidity: event WritershipDecreased(address indexed oldWriter)
func (_Dmessage *DmessageFilterer) WatchWritershipDecreased(opts *bind.WatchOpts, sink chan<- *DmessageWritershipDecreased, oldWriter []common.Address) (event.Subscription, error) {

	var oldWriterRule []interface{}
	for _, oldWriterItem := range oldWriter {
		oldWriterRule = append(oldWriterRule, oldWriterItem)
	}

	logs, sub, err := _Dmessage.contract.WatchLogs(opts, "WritershipDecreased", oldWriterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmessageWritershipDecreased)
				if err := _Dmessage.contract.UnpackLog(event, "WritershipDecreased", log); err != nil {
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

// ParseWritershipDecreased is a log parse operation binding the contract event 0xb24c36b70cc67df9ea22f053c15c4a41e7170d7be0637d2637a3445a266ac6af.
//
// Solidity: event WritershipDecreased(address indexed oldWriter)
func (_Dmessage *DmessageFilterer) ParseWritershipDecreased(log types.Log) (*DmessageWritershipDecreased, error) {
	event := new(DmessageWritershipDecreased)
	if err := _Dmessage.contract.UnpackLog(event, "WritershipDecreased", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DmessageWritershipIncreasedIterator is returned from FilterWritershipIncreased and is used to iterate over the raw logs and unpacked data for WritershipIncreased events raised by the Dmessage contract.
type DmessageWritershipIncreasedIterator struct {
	Event *DmessageWritershipIncreased // Event containing the contract specifics and raw log

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
func (it *DmessageWritershipIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DmessageWritershipIncreased)
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
		it.Event = new(DmessageWritershipIncreased)
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
func (it *DmessageWritershipIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DmessageWritershipIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DmessageWritershipIncreased represents a WritershipIncreased event raised by the Dmessage contract.
type DmessageWritershipIncreased struct {
	NewWriter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWritershipIncreased is a free log retrieval operation binding the contract event 0x8f2adc2e2f2723d375f8bb1359cb1027f713945395ca36c42b83dd94c5bdc7f2.
//
// Solidity: event WritershipIncreased(address indexed newWriter)
func (_Dmessage *DmessageFilterer) FilterWritershipIncreased(opts *bind.FilterOpts, newWriter []common.Address) (*DmessageWritershipIncreasedIterator, error) {

	var newWriterRule []interface{}
	for _, newWriterItem := range newWriter {
		newWriterRule = append(newWriterRule, newWriterItem)
	}

	logs, sub, err := _Dmessage.contract.FilterLogs(opts, "WritershipIncreased", newWriterRule)
	if err != nil {
		return nil, err
	}
	return &DmessageWritershipIncreasedIterator{contract: _Dmessage.contract, event: "WritershipIncreased", logs: logs, sub: sub}, nil
}

// WatchWritershipIncreased is a free log subscription operation binding the contract event 0x8f2adc2e2f2723d375f8bb1359cb1027f713945395ca36c42b83dd94c5bdc7f2.
//
// Solidity: event WritershipIncreased(address indexed newWriter)
func (_Dmessage *DmessageFilterer) WatchWritershipIncreased(opts *bind.WatchOpts, sink chan<- *DmessageWritershipIncreased, newWriter []common.Address) (event.Subscription, error) {

	var newWriterRule []interface{}
	for _, newWriterItem := range newWriter {
		newWriterRule = append(newWriterRule, newWriterItem)
	}

	logs, sub, err := _Dmessage.contract.WatchLogs(opts, "WritershipIncreased", newWriterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DmessageWritershipIncreased)
				if err := _Dmessage.contract.UnpackLog(event, "WritershipIncreased", log); err != nil {
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

// ParseWritershipIncreased is a log parse operation binding the contract event 0x8f2adc2e2f2723d375f8bb1359cb1027f713945395ca36c42b83dd94c5bdc7f2.
//
// Solidity: event WritershipIncreased(address indexed newWriter)
func (_Dmessage *DmessageFilterer) ParseWritershipIncreased(log types.Log) (*DmessageWritershipIncreased, error) {
	event := new(DmessageWritershipIncreased)
	if err := _Dmessage.contract.UnpackLog(event, "WritershipIncreased", log); err != nil {
		return nil, err
	}
	return event, nil
}
