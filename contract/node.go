// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// AccelerateNodeABI is the input ABI used to generate the binding from.
const AccelerateNodeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"idx\",\"type\":\"uint32\"},{\"name\":\"n\",\"type\":\"string\"}],\"name\":\"replace\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"n\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// AccelerateNodeBin is the compiled bytecode used for deploying new contracts.
const AccelerateNodeBin = `0x60806040819052600080546001600160a01b03191633178082556001600160a01b0316917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3610a16806100576000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638f32d59b1161005b5780638f32d59b146100dd578063b0c8f9dc146100f2578063d826f88f14610105578063f2fde38b1461010d57610088565b80636d4ce63c1461008d578063715018a6146100ab578063853c51cc146100b55780638da5cb5b146100c8575b600080fd5b610095610120565b6040516100a291906108aa565b60405180910390f35b6100b36101f9565b005b6100b36100c33660046106f8565b610270565b6100d06102ca565b6040516100a2919061089c565b6100e56102d9565b6040516100a291906108bb565b6100956101003660046106c3565b6102ea565b6100b361042d565b6100b361011b36600461069d565b61045f565b60606001805480602002602001604051908101604052809291908181526020016000905b828210156101ef5760008481526020908190208301805460408051601f60026000196101006001871615020190941693909304928301859004850281018501909152818152928301828280156101db5780601f106101b0576101008083540402835291602001916101db565b820191906000526020600020905b8154815290600101906020018083116101be57829003601f168201915b505050505081526020019060010190610144565b5050505090505b90565b6102016102d9565b6102265760405162461bcd60e51b815260040161021d906108d9565b60405180910390fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6102786102d9565b6102945760405162461bcd60e51b815260040161021d906108d9565b8060018363ffffffff16815481106102a857fe5b9060005260206000200190805190602001906102c5929190610510565b505050565b6000546001600160a01b031690565b6000546001600160a01b0316331490565b60606102f46102d9565b6103105760405162461bcd60e51b815260040161021d906108d9565b600180548082018083556000929092528351610353917fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601906020860190610510565b50506001805480602002602001604051908101604052809291908181526020016000905b828210156104225760008481526020908190208301805460408051601f600260001961010060018716150201909416939093049283018590048502810185019091528181529283018282801561040e5780601f106103e35761010080835404028352916020019161040e565b820191906000526020600020905b8154815290600101906020018083116103f157829003601f168201915b505050505081526020019060010190610377565b505050509050919050565b6104356102d9565b6104515760405162461bcd60e51b815260040161021d906108d9565b61045d6001600061058e565b565b6104676102d9565b6104835760405162461bcd60e51b815260040161021d906108d9565b61048c8161048f565b50565b6001600160a01b0381166104b55760405162461bcd60e51b815260040161021d906108c9565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061055157805160ff191683800117855561057e565b8280016001018555821561057e579182015b8281111561057e578251825591602001919060010190610563565b5061058a9291506105ac565b5090565b508054600082559060005260206000209081019061048c91906105c6565b6101f691905b8082111561058a57600081556001016105b2565b6101f691905b8082111561058a5760006105e082826105e9565b506001016105cc565b50805460018160011615610100020316600290046000825580601f1061060f575061048c565b601f01602090049060005260206000209081019061048c91906105ac565b8035610638816109b6565b92915050565b600082601f83011261064f57600080fd5b813561066261065d82610910565b6108e9565b9150808252602083016020830185838301111561067e57600080fd5b610689838284610970565b50505092915050565b8035610638816109ca565b6000602082840312156106af57600080fd5b60006106bb848461062d565b949350505050565b6000602082840312156106d557600080fd5b813567ffffffffffffffff8111156106ec57600080fd5b6106bb8482850161063e565b6000806040838503121561070b57600080fd5b60006107178585610692565b925050602083013567ffffffffffffffff81111561073457600080fd5b6107408582860161063e565b9150509250929050565b600061075683836107e3565b9392505050565b6107668161094b565b82525050565b60006107778261093e565b6107818185610942565b93508360208202850161079385610938565b8060005b858110156107cd57848403895281516107b0858261074a565b94506107bb83610938565b60209a909a0199925050600101610797565b5091979650505050505050565b61076681610956565b60006107ee8261093e565b6107f88185610942565b935061080881856020860161097c565b610811816109ac565b9093019392505050565b6000610828602683610942565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181526564647265737360d01b602082015260400192915050565b6000610870602083610942565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572815260200192915050565b60208101610638828461075d565b60208082528101610756818461076c565b6020810161063882846107da565b602080825281016106388161081b565b6020808252810161063881610863565b60405181810167ffffffffffffffff8111828210171561090857600080fd5b604052919050565b600067ffffffffffffffff82111561092757600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b60006106388261095b565b151590565b6001600160a01b031690565b63ffffffff1690565b82818337506000910152565b60005b8381101561099757818101518382015260200161097f565b838111156109a6576000848401525b50505050565b601f01601f191690565b6109bf8161094b565b811461048c57600080fd5b6109bf8161096756fea365627a7a72305820996a6ded35cb5238668f632cc614843adc71fd6e3ca1180119e361a7de7a0a6c6c6578706572696d656e74616cf564736f6c63430005090040`

// DeployAccelerateNode deploys a new Ethereum contract, binding an instance of AccelerateNode to it.
func DeployAccelerateNode(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccelerateNode, error) {
	parsed, err := abi.JSON(strings.NewReader(AccelerateNodeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccelerateNodeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccelerateNode{AccelerateNodeCaller: AccelerateNodeCaller{contract: contract}, AccelerateNodeTransactor: AccelerateNodeTransactor{contract: contract}, AccelerateNodeFilterer: AccelerateNodeFilterer{contract: contract}}, nil
}

// AccelerateNode is an auto generated Go binding around an Ethereum contract.
type AccelerateNode struct {
	AccelerateNodeCaller     // Read-only binding to the contract
	AccelerateNodeTransactor // Write-only binding to the contract
	AccelerateNodeFilterer   // Log filterer for contract events
}

// AccelerateNodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccelerateNodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccelerateNodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccelerateNodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccelerateNodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccelerateNodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccelerateNodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccelerateNodeSession struct {
	Contract     *AccelerateNode   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccelerateNodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccelerateNodeCallerSession struct {
	Contract *AccelerateNodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AccelerateNodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccelerateNodeTransactorSession struct {
	Contract     *AccelerateNodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AccelerateNodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccelerateNodeRaw struct {
	Contract *AccelerateNode // Generic contract binding to access the raw methods on
}

// AccelerateNodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccelerateNodeCallerRaw struct {
	Contract *AccelerateNodeCaller // Generic read-only contract binding to access the raw methods on
}

// AccelerateNodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccelerateNodeTransactorRaw struct {
	Contract *AccelerateNodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccelerateNode creates a new instance of AccelerateNode, bound to a specific deployed contract.
func NewAccelerateNode(address common.Address, backend bind.ContractBackend) (*AccelerateNode, error) {
	contract, err := bindAccelerateNode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccelerateNode{AccelerateNodeCaller: AccelerateNodeCaller{contract: contract}, AccelerateNodeTransactor: AccelerateNodeTransactor{contract: contract}, AccelerateNodeFilterer: AccelerateNodeFilterer{contract: contract}}, nil
}

// NewAccelerateNodeCaller creates a new read-only instance of AccelerateNode, bound to a specific deployed contract.
func NewAccelerateNodeCaller(address common.Address, caller bind.ContractCaller) (*AccelerateNodeCaller, error) {
	contract, err := bindAccelerateNode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccelerateNodeCaller{contract: contract}, nil
}

// NewAccelerateNodeTransactor creates a new write-only instance of AccelerateNode, bound to a specific deployed contract.
func NewAccelerateNodeTransactor(address common.Address, transactor bind.ContractTransactor) (*AccelerateNodeTransactor, error) {
	contract, err := bindAccelerateNode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccelerateNodeTransactor{contract: contract}, nil
}

// NewAccelerateNodeFilterer creates a new log filterer instance of AccelerateNode, bound to a specific deployed contract.
func NewAccelerateNodeFilterer(address common.Address, filterer bind.ContractFilterer) (*AccelerateNodeFilterer, error) {
	contract, err := bindAccelerateNode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccelerateNodeFilterer{contract: contract}, nil
}

// bindAccelerateNode binds a generic wrapper to an already deployed contract.
func bindAccelerateNode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccelerateNodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccelerateNode *AccelerateNodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccelerateNode.Contract.AccelerateNodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccelerateNode *AccelerateNodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccelerateNode.Contract.AccelerateNodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccelerateNode *AccelerateNodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccelerateNode.Contract.AccelerateNodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccelerateNode *AccelerateNodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccelerateNode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccelerateNode *AccelerateNodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccelerateNode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccelerateNode *AccelerateNodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccelerateNode.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(string[])
func (_AccelerateNode *AccelerateNodeCaller) Get(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _AccelerateNode.contract.Call(opts, out, "get")
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(string[])
func (_AccelerateNode *AccelerateNodeSession) Get() ([]string, error) {
	return _AccelerateNode.Contract.Get(&_AccelerateNode.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(string[])
func (_AccelerateNode *AccelerateNodeCallerSession) Get() ([]string, error) {
	return _AccelerateNode.Contract.Get(&_AccelerateNode.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_AccelerateNode *AccelerateNodeCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AccelerateNode.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_AccelerateNode *AccelerateNodeSession) IsOwner() (bool, error) {
	return _AccelerateNode.Contract.IsOwner(&_AccelerateNode.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_AccelerateNode *AccelerateNodeCallerSession) IsOwner() (bool, error) {
	return _AccelerateNode.Contract.IsOwner(&_AccelerateNode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AccelerateNode *AccelerateNodeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccelerateNode.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AccelerateNode *AccelerateNodeSession) Owner() (common.Address, error) {
	return _AccelerateNode.Contract.Owner(&_AccelerateNode.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AccelerateNode *AccelerateNodeCallerSession) Owner() (common.Address, error) {
	return _AccelerateNode.Contract.Owner(&_AccelerateNode.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string n) returns(string[])
func (_AccelerateNode *AccelerateNodeTransactor) Add(opts *bind.TransactOpts, n string) (*types.Transaction, error) {
	return _AccelerateNode.contract.Transact(opts, "add", n)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string n) returns(string[])
func (_AccelerateNode *AccelerateNodeSession) Add(n string) (*types.Transaction, error) {
	return _AccelerateNode.Contract.Add(&_AccelerateNode.TransactOpts, n)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string n) returns(string[])
func (_AccelerateNode *AccelerateNodeTransactorSession) Add(n string) (*types.Transaction, error) {
	return _AccelerateNode.Contract.Add(&_AccelerateNode.TransactOpts, n)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccelerateNode *AccelerateNodeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccelerateNode.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccelerateNode *AccelerateNodeSession) RenounceOwnership() (*types.Transaction, error) {
	return _AccelerateNode.Contract.RenounceOwnership(&_AccelerateNode.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AccelerateNode *AccelerateNodeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AccelerateNode.Contract.RenounceOwnership(&_AccelerateNode.TransactOpts)
}

// Replace is a paid mutator transaction binding the contract method 0x853c51cc.
//
// Solidity: function replace(uint32 idx, string n) returns()
func (_AccelerateNode *AccelerateNodeTransactor) Replace(opts *bind.TransactOpts, idx uint32, n string) (*types.Transaction, error) {
	return _AccelerateNode.contract.Transact(opts, "replace", idx, n)
}

// Replace is a paid mutator transaction binding the contract method 0x853c51cc.
//
// Solidity: function replace(uint32 idx, string n) returns()
func (_AccelerateNode *AccelerateNodeSession) Replace(idx uint32, n string) (*types.Transaction, error) {
	return _AccelerateNode.Contract.Replace(&_AccelerateNode.TransactOpts, idx, n)
}

// Replace is a paid mutator transaction binding the contract method 0x853c51cc.
//
// Solidity: function replace(uint32 idx, string n) returns()
func (_AccelerateNode *AccelerateNodeTransactorSession) Replace(idx uint32, n string) (*types.Transaction, error) {
	return _AccelerateNode.Contract.Replace(&_AccelerateNode.TransactOpts, idx, n)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_AccelerateNode *AccelerateNodeTransactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccelerateNode.contract.Transact(opts, "reset")
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_AccelerateNode *AccelerateNodeSession) Reset() (*types.Transaction, error) {
	return _AccelerateNode.Contract.Reset(&_AccelerateNode.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns()
func (_AccelerateNode *AccelerateNodeTransactorSession) Reset() (*types.Transaction, error) {
	return _AccelerateNode.Contract.Reset(&_AccelerateNode.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccelerateNode *AccelerateNodeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AccelerateNode.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccelerateNode *AccelerateNodeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccelerateNode.Contract.TransferOwnership(&_AccelerateNode.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AccelerateNode *AccelerateNodeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AccelerateNode.Contract.TransferOwnership(&_AccelerateNode.TransactOpts, newOwner)
}

// AccelerateNodeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AccelerateNode contract.
type AccelerateNodeOwnershipTransferredIterator struct {
	Event *AccelerateNodeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccelerateNodeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccelerateNodeOwnershipTransferred)
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
		it.Event = new(AccelerateNodeOwnershipTransferred)
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
func (it *AccelerateNodeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccelerateNodeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccelerateNodeOwnershipTransferred represents a OwnershipTransferred event raised by the AccelerateNode contract.
type AccelerateNodeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccelerateNode *AccelerateNodeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccelerateNodeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccelerateNode.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccelerateNodeOwnershipTransferredIterator{contract: _AccelerateNode.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AccelerateNode *AccelerateNodeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccelerateNodeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AccelerateNode.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccelerateNodeOwnershipTransferred)
				if err := _AccelerateNode.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
