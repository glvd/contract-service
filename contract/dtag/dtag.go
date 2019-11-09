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
const DtagABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

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
