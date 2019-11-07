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
const DtagABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getMsgLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"getMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structDMessage.Message\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWriter\",\"type\":\"address\"}],\"name\":\"increasedWritership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"ids\",\"type\":\"string[]\"}],\"name\":\"checkAllMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"checkMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"delMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structDMessage.Message\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceWritership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"writer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getMessages\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structDMessage.Message[]\",\"name\":\"_value\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isWriter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMsgIds\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"updateMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"ids\",\"type\":\"string[]\"}],\"name\":\"checkMessages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"writers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_value\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getMsgId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_value\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldWriter\",\"type\":\"address\"}],\"name\":\"decreaseWritership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"addMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newWriter\",\"type\":\"address\"}],\"name\":\"WritershipIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldWriter\",\"type\":\"address\"}],\"name\":\"WritershipDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Id      string
	Content string
	Version string
}

// CheckAllMessage is a free data retrieval call binding the contract method 0x16313e92.
//
// Solidity: function checkAllMessage(string[] ids) constant returns(bool)
func (_Dtag *DtagCaller) CheckAllMessage(opts *bind.CallOpts, ids []string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dtag.contract.Call(opts, out, "checkAllMessage", ids)
	return *ret0, err
}

// CheckAllMessage is a free data retrieval call binding the contract method 0x16313e92.
//
// Solidity: function checkAllMessage(string[] ids) constant returns(bool)
func (_Dtag *DtagSession) CheckAllMessage(ids []string) (bool, error) {
	return _Dtag.Contract.CheckAllMessage(&_Dtag.CallOpts, ids)
}

// CheckAllMessage is a free data retrieval call binding the contract method 0x16313e92.
//
// Solidity: function checkAllMessage(string[] ids) constant returns(bool)
func (_Dtag *DtagCallerSession) CheckAllMessage(ids []string) (bool, error) {
	return _Dtag.Contract.CheckAllMessage(&_Dtag.CallOpts, ids)
}

// CheckMessage is a free data retrieval call binding the contract method 0x1873ec68.
//
// Solidity: function checkMessage(string id) constant returns(bool)
func (_Dtag *DtagCaller) CheckMessage(opts *bind.CallOpts, id string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dtag.contract.Call(opts, out, "checkMessage", id)
	return *ret0, err
}

// CheckMessage is a free data retrieval call binding the contract method 0x1873ec68.
//
// Solidity: function checkMessage(string id) constant returns(bool)
func (_Dtag *DtagSession) CheckMessage(id string) (bool, error) {
	return _Dtag.Contract.CheckMessage(&_Dtag.CallOpts, id)
}

// CheckMessage is a free data retrieval call binding the contract method 0x1873ec68.
//
// Solidity: function checkMessage(string id) constant returns(bool)
func (_Dtag *DtagCallerSession) CheckMessage(id string) (bool, error) {
	return _Dtag.Contract.CheckMessage(&_Dtag.CallOpts, id)
}

// CheckMessages is a free data retrieval call binding the contract method 0x834daf0c.
//
// Solidity: function checkMessages(string[] ids) constant returns(uint256, bool)
func (_Dtag *DtagCaller) CheckMessages(opts *bind.CallOpts, ids []string) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Dtag.contract.Call(opts, out, "checkMessages", ids)
	return *ret0, *ret1, err
}

// CheckMessages is a free data retrieval call binding the contract method 0x834daf0c.
//
// Solidity: function checkMessages(string[] ids) constant returns(uint256, bool)
func (_Dtag *DtagSession) CheckMessages(ids []string) (*big.Int, bool, error) {
	return _Dtag.Contract.CheckMessages(&_Dtag.CallOpts, ids)
}

// CheckMessages is a free data retrieval call binding the contract method 0x834daf0c.
//
// Solidity: function checkMessages(string[] ids) constant returns(uint256, bool)
func (_Dtag *DtagCallerSession) CheckMessages(ids []string) (*big.Int, bool, error) {
	return _Dtag.Contract.CheckMessages(&_Dtag.CallOpts, ids)
}

// GetMessage is a free data retrieval call binding the contract method 0x0cc4e8d8.
//
// Solidity: function getMessage(string id) constant returns(Struct0)
func (_Dtag *DtagCaller) GetMessage(opts *bind.CallOpts, id string) (Struct0, error) {
	var (
		ret0 = new(Struct0)
	)
	out := ret0
	err := _Dtag.contract.Call(opts, out, "getMessage", id)
	return *ret0, err
}

// GetMessage is a free data retrieval call binding the contract method 0x0cc4e8d8.
//
// Solidity: function getMessage(string id) constant returns(Struct0)
func (_Dtag *DtagSession) GetMessage(id string) (Struct0, error) {
	return _Dtag.Contract.GetMessage(&_Dtag.CallOpts, id)
}

// GetMessage is a free data retrieval call binding the contract method 0x0cc4e8d8.
//
// Solidity: function getMessage(string id) constant returns(Struct0)
func (_Dtag *DtagCallerSession) GetMessage(id string) (Struct0, error) {
	return _Dtag.Contract.GetMessage(&_Dtag.CallOpts, id)
}

// GetMessages is a free data retrieval call binding the contract method 0x4be8362d.
//
// Solidity: function getMessages(uint256 start, uint256 limit) constant returns([]Struct0 _value, uint256 _size)
func (_Dtag *DtagCaller) GetMessages(opts *bind.CallOpts, start *big.Int, limit *big.Int) (struct {
	Value []Struct0
	Size  *big.Int
}, error) {
	ret := new(struct {
		Value []Struct0
		Size  *big.Int
	})
	out := ret
	err := _Dtag.contract.Call(opts, out, "getMessages", start, limit)
	return *ret, err
}

// GetMessages is a free data retrieval call binding the contract method 0x4be8362d.
//
// Solidity: function getMessages(uint256 start, uint256 limit) constant returns([]Struct0 _value, uint256 _size)
func (_Dtag *DtagSession) GetMessages(start *big.Int, limit *big.Int) (struct {
	Value []Struct0
	Size  *big.Int
}, error) {
	return _Dtag.Contract.GetMessages(&_Dtag.CallOpts, start, limit)
}

// GetMessages is a free data retrieval call binding the contract method 0x4be8362d.
//
// Solidity: function getMessages(uint256 start, uint256 limit) constant returns([]Struct0 _value, uint256 _size)
func (_Dtag *DtagCallerSession) GetMessages(start *big.Int, limit *big.Int) (struct {
	Value []Struct0
	Size  *big.Int
}, error) {
	return _Dtag.Contract.GetMessages(&_Dtag.CallOpts, start, limit)
}

// GetMsgId is a free data retrieval call binding the contract method 0x95cb2638.
//
// Solidity: function getMsgId(uint256 _index) constant returns(string _value)
func (_Dtag *DtagCaller) GetMsgId(opts *bind.CallOpts, _index *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Dtag.contract.Call(opts, out, "getMsgId", _index)
	return *ret0, err
}

// GetMsgId is a free data retrieval call binding the contract method 0x95cb2638.
//
// Solidity: function getMsgId(uint256 _index) constant returns(string _value)
func (_Dtag *DtagSession) GetMsgId(_index *big.Int) (string, error) {
	return _Dtag.Contract.GetMsgId(&_Dtag.CallOpts, _index)
}

// GetMsgId is a free data retrieval call binding the contract method 0x95cb2638.
//
// Solidity: function getMsgId(uint256 _index) constant returns(string _value)
func (_Dtag *DtagCallerSession) GetMsgId(_index *big.Int) (string, error) {
	return _Dtag.Contract.GetMsgId(&_Dtag.CallOpts, _index)
}

// GetMsgIds is a free data retrieval call binding the contract method 0x68bfc43c.
//
// Solidity: function getMsgIds() constant returns(string[])
func (_Dtag *DtagCaller) GetMsgIds(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _Dtag.contract.Call(opts, out, "getMsgIds")
	return *ret0, err
}

// GetMsgIds is a free data retrieval call binding the contract method 0x68bfc43c.
//
// Solidity: function getMsgIds() constant returns(string[])
func (_Dtag *DtagSession) GetMsgIds() ([]string, error) {
	return _Dtag.Contract.GetMsgIds(&_Dtag.CallOpts)
}

// GetMsgIds is a free data retrieval call binding the contract method 0x68bfc43c.
//
// Solidity: function getMsgIds() constant returns(string[])
func (_Dtag *DtagCallerSession) GetMsgIds() ([]string, error) {
	return _Dtag.Contract.GetMsgIds(&_Dtag.CallOpts)
}

// GetMsgLength is a free data retrieval call binding the contract method 0x085c00e1.
//
// Solidity: function getMsgLength() constant returns(uint256)
func (_Dtag *DtagCaller) GetMsgLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Dtag.contract.Call(opts, out, "getMsgLength")
	return *ret0, err
}

// GetMsgLength is a free data retrieval call binding the contract method 0x085c00e1.
//
// Solidity: function getMsgLength() constant returns(uint256)
func (_Dtag *DtagSession) GetMsgLength() (*big.Int, error) {
	return _Dtag.Contract.GetMsgLength(&_Dtag.CallOpts)
}

// GetMsgLength is a free data retrieval call binding the contract method 0x085c00e1.
//
// Solidity: function getMsgLength() constant returns(uint256)
func (_Dtag *DtagCallerSession) GetMsgLength() (*big.Int, error) {
	return _Dtag.Contract.GetMsgLength(&_Dtag.CallOpts)
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

// IsWriter is a free data retrieval call binding the contract method 0x4d6ee9fd.
//
// Solidity: function isWriter() constant returns(bool)
func (_Dtag *DtagCaller) IsWriter(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Dtag.contract.Call(opts, out, "isWriter")
	return *ret0, err
}

// IsWriter is a free data retrieval call binding the contract method 0x4d6ee9fd.
//
// Solidity: function isWriter() constant returns(bool)
func (_Dtag *DtagSession) IsWriter() (bool, error) {
	return _Dtag.Contract.IsWriter(&_Dtag.CallOpts)
}

// IsWriter is a free data retrieval call binding the contract method 0x4d6ee9fd.
//
// Solidity: function isWriter() constant returns(bool)
func (_Dtag *DtagCallerSession) IsWriter() (bool, error) {
	return _Dtag.Contract.IsWriter(&_Dtag.CallOpts)
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

// Writer is a free data retrieval call binding the contract method 0x453a2abc.
//
// Solidity: function writer() constant returns(address)
func (_Dtag *DtagCaller) Writer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Dtag.contract.Call(opts, out, "writer")
	return *ret0, err
}

// Writer is a free data retrieval call binding the contract method 0x453a2abc.
//
// Solidity: function writer() constant returns(address)
func (_Dtag *DtagSession) Writer() (common.Address, error) {
	return _Dtag.Contract.Writer(&_Dtag.CallOpts)
}

// Writer is a free data retrieval call binding the contract method 0x453a2abc.
//
// Solidity: function writer() constant returns(address)
func (_Dtag *DtagCallerSession) Writer() (common.Address, error) {
	return _Dtag.Contract.Writer(&_Dtag.CallOpts)
}

// Writers is a free data retrieval call binding the contract method 0x8b2f5369.
//
// Solidity: function writers() constant returns(address[] _value)
func (_Dtag *DtagCaller) Writers(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Dtag.contract.Call(opts, out, "writers")
	return *ret0, err
}

// Writers is a free data retrieval call binding the contract method 0x8b2f5369.
//
// Solidity: function writers() constant returns(address[] _value)
func (_Dtag *DtagSession) Writers() ([]common.Address, error) {
	return _Dtag.Contract.Writers(&_Dtag.CallOpts)
}

// Writers is a free data retrieval call binding the contract method 0x8b2f5369.
//
// Solidity: function writers() constant returns(address[] _value)
func (_Dtag *DtagCallerSession) Writers() ([]common.Address, error) {
	return _Dtag.Contract.Writers(&_Dtag.CallOpts)
}

// AddMessage is a paid mutator transaction binding the contract method 0xb98464b3.
//
// Solidity: function addMessage(string id, string content, string version) returns(bool)
func (_Dtag *DtagTransactor) AddMessage(opts *bind.TransactOpts, id string, content string, version string) (*types.Transaction, error) {
	return _Dtag.contract.Transact(opts, "addMessage", id, content, version)
}

// AddMessage is a paid mutator transaction binding the contract method 0xb98464b3.
//
// Solidity: function addMessage(string id, string content, string version) returns(bool)
func (_Dtag *DtagSession) AddMessage(id string, content string, version string) (*types.Transaction, error) {
	return _Dtag.Contract.AddMessage(&_Dtag.TransactOpts, id, content, version)
}

// AddMessage is a paid mutator transaction binding the contract method 0xb98464b3.
//
// Solidity: function addMessage(string id, string content, string version) returns(bool)
func (_Dtag *DtagTransactorSession) AddMessage(id string, content string, version string) (*types.Transaction, error) {
	return _Dtag.Contract.AddMessage(&_Dtag.TransactOpts, id, content, version)
}

// DecreaseWritership is a paid mutator transaction binding the contract method 0x9f217bdf.
//
// Solidity: function decreaseWritership(address oldWriter) returns()
func (_Dtag *DtagTransactor) DecreaseWritership(opts *bind.TransactOpts, oldWriter common.Address) (*types.Transaction, error) {
	return _Dtag.contract.Transact(opts, "decreaseWritership", oldWriter)
}

// DecreaseWritership is a paid mutator transaction binding the contract method 0x9f217bdf.
//
// Solidity: function decreaseWritership(address oldWriter) returns()
func (_Dtag *DtagSession) DecreaseWritership(oldWriter common.Address) (*types.Transaction, error) {
	return _Dtag.Contract.DecreaseWritership(&_Dtag.TransactOpts, oldWriter)
}

// DecreaseWritership is a paid mutator transaction binding the contract method 0x9f217bdf.
//
// Solidity: function decreaseWritership(address oldWriter) returns()
func (_Dtag *DtagTransactorSession) DecreaseWritership(oldWriter common.Address) (*types.Transaction, error) {
	return _Dtag.Contract.DecreaseWritership(&_Dtag.TransactOpts, oldWriter)
}

// DelMessage is a paid mutator transaction binding the contract method 0x2309bc96.
//
// Solidity: function delMessage(string id) returns(Struct0, bool)
func (_Dtag *DtagTransactor) DelMessage(opts *bind.TransactOpts, id string) (*types.Transaction, error) {
	return _Dtag.contract.Transact(opts, "delMessage", id)
}

// DelMessage is a paid mutator transaction binding the contract method 0x2309bc96.
//
// Solidity: function delMessage(string id) returns(Struct0, bool)
func (_Dtag *DtagSession) DelMessage(id string) (*types.Transaction, error) {
	return _Dtag.Contract.DelMessage(&_Dtag.TransactOpts, id)
}

// DelMessage is a paid mutator transaction binding the contract method 0x2309bc96.
//
// Solidity: function delMessage(string id) returns(Struct0, bool)
func (_Dtag *DtagTransactorSession) DelMessage(id string) (*types.Transaction, error) {
	return _Dtag.Contract.DelMessage(&_Dtag.TransactOpts, id)
}

// IncreasedWritership is a paid mutator transaction binding the contract method 0x14b4d411.
//
// Solidity: function increasedWritership(address newWriter) returns()
func (_Dtag *DtagTransactor) IncreasedWritership(opts *bind.TransactOpts, newWriter common.Address) (*types.Transaction, error) {
	return _Dtag.contract.Transact(opts, "increasedWritership", newWriter)
}

// IncreasedWritership is a paid mutator transaction binding the contract method 0x14b4d411.
//
// Solidity: function increasedWritership(address newWriter) returns()
func (_Dtag *DtagSession) IncreasedWritership(newWriter common.Address) (*types.Transaction, error) {
	return _Dtag.Contract.IncreasedWritership(&_Dtag.TransactOpts, newWriter)
}

// IncreasedWritership is a paid mutator transaction binding the contract method 0x14b4d411.
//
// Solidity: function increasedWritership(address newWriter) returns()
func (_Dtag *DtagTransactorSession) IncreasedWritership(newWriter common.Address) (*types.Transaction, error) {
	return _Dtag.Contract.IncreasedWritership(&_Dtag.TransactOpts, newWriter)
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

// RenounceWritership is a paid mutator transaction binding the contract method 0x4075e83e.
//
// Solidity: function renounceWritership() returns()
func (_Dtag *DtagTransactor) RenounceWritership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dtag.contract.Transact(opts, "renounceWritership")
}

// RenounceWritership is a paid mutator transaction binding the contract method 0x4075e83e.
//
// Solidity: function renounceWritership() returns()
func (_Dtag *DtagSession) RenounceWritership() (*types.Transaction, error) {
	return _Dtag.Contract.RenounceWritership(&_Dtag.TransactOpts)
}

// RenounceWritership is a paid mutator transaction binding the contract method 0x4075e83e.
//
// Solidity: function renounceWritership() returns()
func (_Dtag *DtagTransactorSession) RenounceWritership() (*types.Transaction, error) {
	return _Dtag.Contract.RenounceWritership(&_Dtag.TransactOpts)
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

// UpdateMessage is a paid mutator transaction binding the contract method 0x6a44815e.
//
// Solidity: function updateMessage(string id, string content, string version) returns(bool)
func (_Dtag *DtagTransactor) UpdateMessage(opts *bind.TransactOpts, id string, content string, version string) (*types.Transaction, error) {
	return _Dtag.contract.Transact(opts, "updateMessage", id, content, version)
}

// UpdateMessage is a paid mutator transaction binding the contract method 0x6a44815e.
//
// Solidity: function updateMessage(string id, string content, string version) returns(bool)
func (_Dtag *DtagSession) UpdateMessage(id string, content string, version string) (*types.Transaction, error) {
	return _Dtag.Contract.UpdateMessage(&_Dtag.TransactOpts, id, content, version)
}

// UpdateMessage is a paid mutator transaction binding the contract method 0x6a44815e.
//
// Solidity: function updateMessage(string id, string content, string version) returns(bool)
func (_Dtag *DtagTransactorSession) UpdateMessage(id string, content string, version string) (*types.Transaction, error) {
	return _Dtag.Contract.UpdateMessage(&_Dtag.TransactOpts, id, content, version)
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

// DtagWritershipDecreasedIterator is returned from FilterWritershipDecreased and is used to iterate over the raw logs and unpacked data for WritershipDecreased events raised by the Dtag contract.
type DtagWritershipDecreasedIterator struct {
	Event *DtagWritershipDecreased // Event containing the contract specifics and raw log

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
func (it *DtagWritershipDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtagWritershipDecreased)
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
		it.Event = new(DtagWritershipDecreased)
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
func (it *DtagWritershipDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtagWritershipDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtagWritershipDecreased represents a WritershipDecreased event raised by the Dtag contract.
type DtagWritershipDecreased struct {
	OldWriter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWritershipDecreased is a free log retrieval operation binding the contract event 0xb24c36b70cc67df9ea22f053c15c4a41e7170d7be0637d2637a3445a266ac6af.
//
// Solidity: event WritershipDecreased(address indexed oldWriter)
func (_Dtag *DtagFilterer) FilterWritershipDecreased(opts *bind.FilterOpts, oldWriter []common.Address) (*DtagWritershipDecreasedIterator, error) {

	var oldWriterRule []interface{}
	for _, oldWriterItem := range oldWriter {
		oldWriterRule = append(oldWriterRule, oldWriterItem)
	}

	logs, sub, err := _Dtag.contract.FilterLogs(opts, "WritershipDecreased", oldWriterRule)
	if err != nil {
		return nil, err
	}
	return &DtagWritershipDecreasedIterator{contract: _Dtag.contract, event: "WritershipDecreased", logs: logs, sub: sub}, nil
}

// WatchWritershipDecreased is a free log subscription operation binding the contract event 0xb24c36b70cc67df9ea22f053c15c4a41e7170d7be0637d2637a3445a266ac6af.
//
// Solidity: event WritershipDecreased(address indexed oldWriter)
func (_Dtag *DtagFilterer) WatchWritershipDecreased(opts *bind.WatchOpts, sink chan<- *DtagWritershipDecreased, oldWriter []common.Address) (event.Subscription, error) {

	var oldWriterRule []interface{}
	for _, oldWriterItem := range oldWriter {
		oldWriterRule = append(oldWriterRule, oldWriterItem)
	}

	logs, sub, err := _Dtag.contract.WatchLogs(opts, "WritershipDecreased", oldWriterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtagWritershipDecreased)
				if err := _Dtag.contract.UnpackLog(event, "WritershipDecreased", log); err != nil {
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
func (_Dtag *DtagFilterer) ParseWritershipDecreased(log types.Log) (*DtagWritershipDecreased, error) {
	event := new(DtagWritershipDecreased)
	if err := _Dtag.contract.UnpackLog(event, "WritershipDecreased", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DtagWritershipIncreasedIterator is returned from FilterWritershipIncreased and is used to iterate over the raw logs and unpacked data for WritershipIncreased events raised by the Dtag contract.
type DtagWritershipIncreasedIterator struct {
	Event *DtagWritershipIncreased // Event containing the contract specifics and raw log

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
func (it *DtagWritershipIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DtagWritershipIncreased)
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
		it.Event = new(DtagWritershipIncreased)
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
func (it *DtagWritershipIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DtagWritershipIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DtagWritershipIncreased represents a WritershipIncreased event raised by the Dtag contract.
type DtagWritershipIncreased struct {
	NewWriter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWritershipIncreased is a free log retrieval operation binding the contract event 0x8f2adc2e2f2723d375f8bb1359cb1027f713945395ca36c42b83dd94c5bdc7f2.
//
// Solidity: event WritershipIncreased(address indexed newWriter)
func (_Dtag *DtagFilterer) FilterWritershipIncreased(opts *bind.FilterOpts, newWriter []common.Address) (*DtagWritershipIncreasedIterator, error) {

	var newWriterRule []interface{}
	for _, newWriterItem := range newWriter {
		newWriterRule = append(newWriterRule, newWriterItem)
	}

	logs, sub, err := _Dtag.contract.FilterLogs(opts, "WritershipIncreased", newWriterRule)
	if err != nil {
		return nil, err
	}
	return &DtagWritershipIncreasedIterator{contract: _Dtag.contract, event: "WritershipIncreased", logs: logs, sub: sub}, nil
}

// WatchWritershipIncreased is a free log subscription operation binding the contract event 0x8f2adc2e2f2723d375f8bb1359cb1027f713945395ca36c42b83dd94c5bdc7f2.
//
// Solidity: event WritershipIncreased(address indexed newWriter)
func (_Dtag *DtagFilterer) WatchWritershipIncreased(opts *bind.WatchOpts, sink chan<- *DtagWritershipIncreased, newWriter []common.Address) (event.Subscription, error) {

	var newWriterRule []interface{}
	for _, newWriterItem := range newWriter {
		newWriterRule = append(newWriterRule, newWriterItem)
	}

	logs, sub, err := _Dtag.contract.WatchLogs(opts, "WritershipIncreased", newWriterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DtagWritershipIncreased)
				if err := _Dtag.contract.UnpackLog(event, "WritershipIncreased", log); err != nil {
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
func (_Dtag *DtagFilterer) ParseWritershipIncreased(log types.Log) (*DtagWritershipIncreased, error) {
	event := new(DtagWritershipIncreased)
	if err := _Dtag.contract.UnpackLog(event, "WritershipIncreased", log); err != nil {
		return nil, err
	}
	return event, nil
}
