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
const DnodeABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

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
