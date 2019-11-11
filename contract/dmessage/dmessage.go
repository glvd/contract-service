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

// DMessageABI is the input ABI used to generate the binding from.
const DMessageABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldWriter\",\"type\":\"address\"}],\"name\":\"WritershipDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newWriter\",\"type\":\"address\"}],\"name\":\"WritershipIncreased\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"addMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"ids\",\"type\":\"string[]\"}],\"name\":\"checkAllMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"checkMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"ids\",\"type\":\"string[]\"}],\"name\":\"checkMessages\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldWriter\",\"type\":\"address\"}],\"name\":\"decreaseWritership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"delMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structDMessage.Message\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"ids\",\"type\":\"string[]\"}],\"name\":\"getIdsMessages\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structDMessage.Message[]\",\"name\":\"_value\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"getMessage\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structDMessage.Message\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"getMessageValue\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getMessages\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"internalType\":\"structDMessage.Message[]\",\"name\":\"_value\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_size\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getMsgId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_value\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMsgIds\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMsgLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWriter\",\"type\":\"address\"}],\"name\":\"increasedWritership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isWriter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceWritership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"name\":\"updateMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"writer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"writers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_value\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DMessageFuncSigs maps the 4-byte function signature to its string representation.
var DMessageFuncSigs = map[string]string{
	"b98464b3": "addMessage(string,string,string)",
	"16313e92": "checkAllMessage(string[])",
	"1873ec68": "checkMessage(string)",
	"834daf0c": "checkMessages(string[])",
	"9f217bdf": "decreaseWritership(address)",
	"2309bc96": "delMessage(string)",
	"6c4ef994": "getIdsMessages(string[])",
	"0cc4e8d8": "getMessage(string)",
	"c57bc642": "getMessageValue(string)",
	"4be8362d": "getMessages(uint256,uint256)",
	"95cb2638": "getMsgId(uint256)",
	"68bfc43c": "getMsgIds()",
	"085c00e1": "getMsgLength()",
	"14b4d411": "increasedWritership(address)",
	"8f32d59b": "isOwner()",
	"4d6ee9fd": "isWriter()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"4075e83e": "renounceWritership()",
	"f2fde38b": "transferOwnership(address)",
	"6a44815e": "updateMessage(string,string,string)",
	"453a2abc": "writer()",
	"8b2f5369": "writers()",
}

// DMessageBin is the compiled bytecode used for deploying new contracts.
var DMessageBin = "0x608060405234801561001057600080fd5b5060006100246001600160e01b036100e216565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35060006100816001600160e01b036100e216565b6001600160a01b0381166000818152600160208190526040808320805460ff191690921790915560038290555192935090917f8f2adc2e2f2723d375f8bb1359cb1027f713945395ca36c42b83dd94c5bdc7f29190a25060006007556100e6565b3390565b61204080620000f66000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c80636a44815e116100c35780638f32d59b1161007c5780638f32d59b146102c057806395cb2638146102c85780639f217bdf146102e8578063b98464b3146102fb578063c57bc6421461030e578063f2fde38b146103305761014d565b80636a44815e146102545780636c4ef99414610267578063715018a61461027a578063834daf0c146102825780638b2f5369146102a35780638da5cb5b146102b85761014d565b80632309bc96116101155780632309bc96146101d85780634075e83e146101f9578063453a2abc146102015780634be8362d146102165780634d6ee9fd1461023757806368bfc43c1461023f5761014d565b8063085c00e1146101525780630cc4e8d81461017057806314b4d4111461019057806316313e92146101a55780631873ec68146101c5575b600080fd5b61015a610343565b6040516101679190611ed2565b60405180910390f35b61018361017e3660046117b4565b61034a565b6040516101679190611ea1565b6101a361019e366004611759565b610542565b005b6101b86101b336600461177f565b61057b565b6040516101679190611db9565b6101b86101d33660046117b4565b6105c3565b6101eb6101e63660046117b4565b6105ef565b604051610167929190611eb2565b6101a3610894565b61020961092d565b6040516101679190611d69565b61022961022436600461189b565b610956565b604051610167929190611d99565b6101b86109eb565b610247610a1b565b6040516101679190611d88565b6101b86102623660046117e9565b610b21565b61022961027536600461177f565b610c30565b6101a3610cc1565b61029561029036600461177f565b610d2f565b604051610167929190611ee0565b6102ab610d70565b6040516101679190611d77565b610209610dfb565b6101b8610e0a565b6102db6102d636600461187d565b610e2e565b6040516101679190611dc7565b6101a36102f6366004611759565b610ecf565b6101b86103093660046117e9565b610efc565b61032161031c3660046117b4565b611041565b60405161016793929190611dd8565b6101a361033e366004611759565b611257565b6007545b90565b610352611558565b6005826040516103629190611d5d565b9081526040805160209281900383018120805460026001821615610100026000190190911604601f810185900490940282016080908101909352606082018481529193909284929184918401828280156103fd5780601f106103d2576101008083540402835291602001916103fd565b820191906000526020600020905b8154815290600101906020018083116103e057829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561049f5780601f106104745761010080835404028352916020019161049f565b820191906000526020600020905b81548152906001019060200180831161048257829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f810183900483028501830190915280845293810193908301828280156105315780601f1061050657610100808354040283529160200191610531565b820191906000526020600020905b81548152906001019060200180831161051457829003601f168201915b50505050508152505090505b919050565b61054a610e0a565b61056f5760405162461bcd60e51b815260040161056690611e81565b60405180910390fd5b61057881611284565b50565b6000805b82518110156105ba576105a483828151811061059757fe5b60200260200101516105c3565b6105b257600091505061053d565b60010161057f565b50600192915050565b60006004826040516105d59190611d5d565b9081526040519081900360200190205460ff169050919050565b6105f7611558565b6000610601610e0a565b61061d5760405162461bcd60e51b815260040161056690611e81565b60048360405161062d9190611d5d565b9081526040519081900360200190205460ff1615156001146106615760405162461bcd60e51b815260040161056690611e71565b60006004846040516106739190611d5d565b908152604051908190036020019020805491151560ff1990921691909117905561069b611355565b6005836040516106ab9190611d5d565b908152604080516020928190038301812080546002600180831615610100026000190190921604601f810186900490950283016080908101909452606083018581529194909385928492909184919084018282801561074b5780601f106107205761010080835404028352916020019161074b565b820191906000526020600020905b81548152906001019060200180831161072e57829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107ed5780601f106107c2576101008083540402835291602001916107ed565b820191906000526020600020905b8154815290600101906020018083116107d057829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f8101839004830285018301909152808452938101939083018282801561087f5780601f106108545761010080835404028352916020019161087f565b820191906000526020600020905b81548152906001019060200180831161086257829003601f168201915b5050505050815250509150915091505b915091565b61089c610e0a565b6108b85760405162461bcd60e51b815260040161056690611e81565b6000600160006108c6611388565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790556108f6611388565b6001600160a01b03167fb24c36b70cc67df9ea22f053c15c4a41e7170d7be0637d2637a3445a266ac6af60405160405180910390a2565b60006109376109eb565b15156001141561095057610949611388565b9050610347565b50600090565b606060006007548385011061096d57836007540392505b826040519080825280602002602001820160405280156109a757816020015b610994611558565b81526020019060019003908161098c5790505b50915060005b838110156109e2576109c361017e828701610e2e565b8382815181106109cf57fe5b60209081029190910101526001016109ad565b50909391925050565b6000600160006109f9611388565b6001600160a01b0316815260208101919091526040016000205460ff16905090565b606080600754604051908082528060200260200182016040528015610a5457816020015b6060815260200190600190039081610a3f5790505b50905060005b600754811015610b1b5760008181526006602090815260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529091830182828015610af75780601f10610acc57610100808354040283529160200191610af7565b820191906000526020600020905b815481529060010190602001808311610ada57829003601f168201915b5050505050828281518110610b0857fe5b6020908102919091010152600101610a5a565b50905090565b6000610b2b6109eb565b610b475760405162461bcd60e51b815260040161056690611e61565b600484604051610b579190611d5d565b9081526040519081900360200190205460ff161515600114610b8b5760405162461bcd60e51b815260040161056690611e91565b610b93611558565b604051806060016040528086815260200185815260200184815250905080600586604051610bc19190611d5d565b90815260200160405180910390206000820151816000019080519060200190610beb929190611579565b506020828101518051610c049260018501920190611579565b5060408201518051610c20916002840191602090910190611579565b50600193505050505b9392505050565b606060008251905080604051908082528060200260200182016040528015610c7257816020015b610c5f611558565b815260200190600190039081610c575790505b50915060005b81811015610cbb57610c9c848281518110610c8f57fe5b602002602001015161034a565b838281518110610ca857fe5b6020908102919091010152600101610c78565b50915091565b610cc9610e0a565b610ce55760405162461bcd60e51b815260040161056690611e81565b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b600080805b8351811015610d6357610d4c84828151811061059757fe5b610d5b5791506000905061088f565b600101610d34565b5060009360019350915050565b6060600354604051908082528060200260200182016040528015610d9e578160200160208202803883390190505b50905060005b600354811015610df75760008181526002602052604090205482516001600160a01b0390911690839083908110610dd757fe5b6001600160a01b0390921660209283029190910190910152600101610da4565b5090565b6000546001600160a01b031690565b600080546001600160a01b0316610e1f611388565b6001600160a01b031614905090565b60008181526006602090815260409182902080548351601f6002600019610100600186161502019093169290920491820184900484028101840190945280845260609392830182828015610ec35780601f10610e9857610100808354040283529160200191610ec3565b820191906000526020600020905b815481529060010190602001808311610ea657829003601f168201915b50505050509050919050565b610ed7610e0a565b610ef35760405162461bcd60e51b815260040161056690611e81565b6105788161138c565b6000610f066109eb565b610f225760405162461bcd60e51b815260040161056690611e61565b600484604051610f329190611d5d565b9081526040519081900360200190205460ff1615610f625760405162461bcd60e51b815260040161056690611e21565b610f6a611558565b604051806060016040528086815260200185815260200184815250905080600586604051610f989190611d5d565b90815260200160405180910390206000820151816000019080519060200190610fc2929190611579565b506020828101518051610fdb9260018501920190611579565b5060408201518051610ff7916002840191602090910190611579565b50905050600160048660405161100d9190611d5d565b908152604051908190036020019020805491151560ff199092169190911790556110368561144f565b506001949350505050565b60608060606005846040516110569190611d5d565b90815260405190819003602001812090600590611074908790611d5d565b90815260200160405180910390206001016005866040516110959190611d5d565b90815260408051602092819003830181208554600260018216156101000260001901909116819004601f81018690048602840186019094528383520192909185918301828280156111275780601f106110fc57610100808354040283529160200191611127565b820191906000526020600020905b81548152906001019060200180831161110a57829003601f168201915b5050855460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959850879450925084019050828280156111b55780601f1061118a576101008083540402835291602001916111b5565b820191906000526020600020905b81548152906001019060200180831161119857829003601f168201915b5050845460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959750869450925084019050828280156112435780601f1061121857610100808354040283529160200191611243565b820191906000526020600020905b81548152906001019060200180831161122657829003601f168201915b505050505090509250925092509193909250565b61125f610e0a565b61127b5760405162461bcd60e51b815260040161056690611e81565b6105788161147e565b6001600160a01b0381166112aa5760405162461bcd60e51b815260040161056690611e51565b6001600160a01b03811660009081526001602052604090205460ff16156112e35760405162461bcd60e51b815260040161056690611e41565b6001600160a01b0381166000818152600160208181526040808420805460ff191684179055600380548552600290925280842080546001600160a01b0319168617905581549092019055517f8f2adc2e2f2723d375f8bb1359cb1027f713945395ca36c42b83dd94c5bdc7f29190a250565b60005b611360610343565b811015610578576113736101d382610e2e565b61138057611380816114ff565b600101611358565b3390565b6001600160a01b03811660009081526001602081905260409091205460ff161515146113ca5760405162461bcd60e51b815260040161056690611e11565b6001600160a01b0381166113f05760405162461bcd60e51b815260040161056690611e51565b6001600160a01b0381166000908152600160205260409020805460ff19169055611418611355565b6040516001600160a01b038216907fb24c36b70cc67df9ea22f053c15c4a41e7170d7be0637d2637a3445a266ac6af90600090a250565b6007546000908152600660209081526040909120825161147192840190611579565b5050600780546001019055565b6001600160a01b0381166114a45760405162461bcd60e51b815260040161056690611e31565b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b600754811061150d57610578565b6007546000199081016000908152600660205260408082208483529120815461154a939192916002610100600184161502909101909116046115f3565b506007805460001901905550565b60405180606001604052806060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106115ba57805160ff19168380011785556115e7565b828001600101855582156115e7579182015b828111156115e75782518255916020019190600101906115cc565b50610df7929150611668565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061162c57805485556115e7565b828001600101855582156115e757600052602060002091601f016020900482015b828111156115e757825482559160010191906001019061164d565b61034791905b80821115610df7576000815560010161166e565b803561168d81611fe0565b92915050565b600082601f8301126116a457600080fd5b81356116b76116b282611f22565b611efb565b81815260209384019390925082018360005b838110156116f557813586016116df88826116ff565b84525060209283019291909101906001016116c9565b5050505092915050565b600082601f83011261171057600080fd5b813561171e6116b282611f43565b9150808252602083016020830185838301111561173a57600080fd5b611745838284611f9a565b50505092915050565b803561168d81611ff4565b60006020828403121561176b57600080fd5b60006117778484611682565b949350505050565b60006020828403121561179157600080fd5b813567ffffffffffffffff8111156117a857600080fd5b61177784828501611693565b6000602082840312156117c657600080fd5b813567ffffffffffffffff8111156117dd57600080fd5b611777848285016116ff565b6000806000606084860312156117fe57600080fd5b833567ffffffffffffffff81111561181557600080fd5b611821868287016116ff565b935050602084013567ffffffffffffffff81111561183e57600080fd5b61184a868287016116ff565b925050604084013567ffffffffffffffff81111561186757600080fd5b611873868287016116ff565b9150509250925092565b60006020828403121561188f57600080fd5b6000611777848461174e565b600080604083850312156118ae57600080fd5b60006118ba858561174e565b92505060206118cb8582860161174e565b9150509250929050565b60006118e18383611901565b505060200190565b6000610c298383611a41565b6000610c298383611cff565b61190a81611f7e565b82525050565b600061191b82611f71565b6119258185611f75565b935061193083611f6b565b8060005b8381101561195e57815161194888826118d5565b975061195383611f6b565b925050600101611934565b509495945050505050565b600061197482611f71565b61197e8185611f75565b93508360208202850161199085611f6b565b8060005b858110156119ca57848403895281516119ad85826118e9565b94506119b883611f6b565b60209a909a0199925050600101611994565b5091979650505050505050565b60006119e282611f71565b6119ec8185611f75565b9350836020820285016119fe85611f6b565b8060005b858110156119ca5784840389528151611a1b85826118f5565b9450611a2683611f6b565b60209a909a0199925050600101611a02565b61190a81611f89565b6000611a4c82611f71565b611a568185611f75565b9350611a66818560208601611fa6565b611a6f81611fd6565b9093019392505050565b6000611a8482611f71565b611a8e818561053d565b9350611a9e818560208601611fa6565b9290920192915050565b6000611ab5602583611f75565b7f57726974657261626c653a206f6c642077726974657220776173206e6f7420778152643934ba32b960d91b602082015260400192915050565b6000611afc601d83611f75565b7f4d6573736167653a20616464206d657373616765206973206578697374000000815260200192915050565b6000611b35602683611f75565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181526564647265737360d01b602082015260400192915050565b6000611b7d601f83611f75565b7f57726974657261626c653a206e65772077726974657220697320657869737400815260200192915050565b6000611bb6602a83611f75565b7f57726974657261626c653a206e65772077726974657220697320746865207a65815269726f206164647265737360b01b602082015260400192915050565b6000611c02602383611f75565b7f577269746561626c653a2063616c6c6572206973206e6f7420746865207772698152623a32b960e91b602082015260400192915050565b6000611c47602483611f75565b7f4d6573736167653a2064656c657465206d657373616765206973206e6f7420658152631e1a5cdd60e21b602082015260400192915050565b6000611c8d602083611f75565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572815260200192915050565b6000611cc6602483611f75565b7f4d6573736167653a20757064617465206d657373616765206973206e6f7420658152631e1a5cdd60e21b602082015260400192915050565b8051606080845260009190840190611d178282611a41565b91505060208301518482036020860152611d318282611a41565b91505060408301518482036040860152611d4b8282611a41565b95945050505050565b61190a81610347565b6000610c298284611a79565b6020810161168d8284611901565b60208082528101610c298184611910565b60208082528101610c298184611969565b60408082528101611daa81856119d7565b9050610c296020830184611d54565b6020810161168d8284611a38565b60208082528101610c298184611a41565b60608082528101611de98186611a41565b90508181036020830152611dfd8185611a41565b90508181036040830152611d4b8184611a41565b6020808252810161168d81611aa8565b6020808252810161168d81611aef565b6020808252810161168d81611b28565b6020808252810161168d81611b70565b6020808252810161168d81611ba9565b6020808252810161168d81611bf5565b6020808252810161168d81611c3a565b6020808252810161168d81611c80565b6020808252810161168d81611cb9565b60208082528101610c298184611cff565b60408082528101611ec38185611cff565b9050610c296020830184611a38565b6020810161168d8284611d54565b60408101611eee8285611d54565b610c296020830184611a38565b60405181810167ffffffffffffffff81118282101715611f1a57600080fd5b604052919050565b600067ffffffffffffffff821115611f3957600080fd5b5060209081020190565b600067ffffffffffffffff821115611f5a57600080fd5b506020601f91909101601f19160190565b60200190565b5190565b90815260200190565b600061168d82611f8e565b151590565b6001600160a01b031690565b82818337506000910152565b60005b83811015611fc1578181015183820152602001611fa9565b83811115611fd0576000848401525b50505050565b601f01601f191690565b611fe981611f7e565b811461057857600080fd5b611fe98161034756fea365627a7a723158204b437980d744f1b001681758680543e40a15e459c475c8cfa28dff47edea7e696c6578706572696d656e74616cf564736f6c634300050c0040"

// DeployDMessage deploys a new Ethereum contract, binding an instance of DMessage to it.
func DeployDMessage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DMessage, error) {
	parsed, err := abi.JSON(strings.NewReader(DMessageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DMessageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DMessage{DMessageCaller: DMessageCaller{contract: contract}, DMessageTransactor: DMessageTransactor{contract: contract}, DMessageFilterer: DMessageFilterer{contract: contract}}, nil
}

// DMessage is an auto generated Go binding around an Ethereum contract.
type DMessage struct {
	DMessageCaller     // Read-only binding to the contract
	DMessageTransactor // Write-only binding to the contract
	DMessageFilterer   // Log filterer for contract events
}

// DMessageCaller is an auto generated read-only Go binding around an Ethereum contract.
type DMessageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DMessageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DMessageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DMessageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DMessageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DMessageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DMessageSession struct {
	Contract     *DMessage         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DMessageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DMessageCallerSession struct {
	Contract *DMessageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DMessageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DMessageTransactorSession struct {
	Contract     *DMessageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DMessageRaw is an auto generated low-level Go binding around an Ethereum contract.
type DMessageRaw struct {
	Contract *DMessage // Generic contract binding to access the raw methods on
}

// DMessageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DMessageCallerRaw struct {
	Contract *DMessageCaller // Generic read-only contract binding to access the raw methods on
}

// DMessageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DMessageTransactorRaw struct {
	Contract *DMessageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDMessage creates a new instance of DMessage, bound to a specific deployed contract.
func NewDMessage(address common.Address, backend bind.ContractBackend) (*DMessage, error) {
	contract, err := bindDMessage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DMessage{DMessageCaller: DMessageCaller{contract: contract}, DMessageTransactor: DMessageTransactor{contract: contract}, DMessageFilterer: DMessageFilterer{contract: contract}}, nil
}

// NewDMessageCaller creates a new read-only instance of DMessage, bound to a specific deployed contract.
func NewDMessageCaller(address common.Address, caller bind.ContractCaller) (*DMessageCaller, error) {
	contract, err := bindDMessage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DMessageCaller{contract: contract}, nil
}

// NewDMessageTransactor creates a new write-only instance of DMessage, bound to a specific deployed contract.
func NewDMessageTransactor(address common.Address, transactor bind.ContractTransactor) (*DMessageTransactor, error) {
	contract, err := bindDMessage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DMessageTransactor{contract: contract}, nil
}

// NewDMessageFilterer creates a new log filterer instance of DMessage, bound to a specific deployed contract.
func NewDMessageFilterer(address common.Address, filterer bind.ContractFilterer) (*DMessageFilterer, error) {
	contract, err := bindDMessage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DMessageFilterer{contract: contract}, nil
}

// bindDMessage binds a generic wrapper to an already deployed contract.
func bindDMessage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DMessageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DMessage *DMessageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DMessage.Contract.DMessageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DMessage *DMessageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DMessage.Contract.DMessageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DMessage *DMessageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DMessage.Contract.DMessageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DMessage *DMessageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DMessage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DMessage *DMessageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DMessage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DMessage *DMessageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DMessage.Contract.contract.Transact(opts, method, params...)
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
func (_DMessage *DMessageCaller) CheckAllMessage(opts *bind.CallOpts, ids []string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DMessage.contract.Call(opts, out, "checkAllMessage", ids)
	return *ret0, err
}

// CheckAllMessage is a free data retrieval call binding the contract method 0x16313e92.
//
// Solidity: function checkAllMessage(string[] ids) constant returns(bool)
func (_DMessage *DMessageSession) CheckAllMessage(ids []string) (bool, error) {
	return _DMessage.Contract.CheckAllMessage(&_DMessage.CallOpts, ids)
}

// CheckAllMessage is a free data retrieval call binding the contract method 0x16313e92.
//
// Solidity: function checkAllMessage(string[] ids) constant returns(bool)
func (_DMessage *DMessageCallerSession) CheckAllMessage(ids []string) (bool, error) {
	return _DMessage.Contract.CheckAllMessage(&_DMessage.CallOpts, ids)
}

// CheckMessage is a free data retrieval call binding the contract method 0x1873ec68.
//
// Solidity: function checkMessage(string id) constant returns(bool)
func (_DMessage *DMessageCaller) CheckMessage(opts *bind.CallOpts, id string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DMessage.contract.Call(opts, out, "checkMessage", id)
	return *ret0, err
}

// CheckMessage is a free data retrieval call binding the contract method 0x1873ec68.
//
// Solidity: function checkMessage(string id) constant returns(bool)
func (_DMessage *DMessageSession) CheckMessage(id string) (bool, error) {
	return _DMessage.Contract.CheckMessage(&_DMessage.CallOpts, id)
}

// CheckMessage is a free data retrieval call binding the contract method 0x1873ec68.
//
// Solidity: function checkMessage(string id) constant returns(bool)
func (_DMessage *DMessageCallerSession) CheckMessage(id string) (bool, error) {
	return _DMessage.Contract.CheckMessage(&_DMessage.CallOpts, id)
}

// CheckMessages is a free data retrieval call binding the contract method 0x834daf0c.
//
// Solidity: function checkMessages(string[] ids) constant returns(uint256, bool)
func (_DMessage *DMessageCaller) CheckMessages(opts *bind.CallOpts, ids []string) (*big.Int, bool, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(bool)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DMessage.contract.Call(opts, out, "checkMessages", ids)
	return *ret0, *ret1, err
}

// CheckMessages is a free data retrieval call binding the contract method 0x834daf0c.
//
// Solidity: function checkMessages(string[] ids) constant returns(uint256, bool)
func (_DMessage *DMessageSession) CheckMessages(ids []string) (*big.Int, bool, error) {
	return _DMessage.Contract.CheckMessages(&_DMessage.CallOpts, ids)
}

// CheckMessages is a free data retrieval call binding the contract method 0x834daf0c.
//
// Solidity: function checkMessages(string[] ids) constant returns(uint256, bool)
func (_DMessage *DMessageCallerSession) CheckMessages(ids []string) (*big.Int, bool, error) {
	return _DMessage.Contract.CheckMessages(&_DMessage.CallOpts, ids)
}

// GetIdsMessages is a free data retrieval call binding the contract method 0x6c4ef994.
//
// Solidity: function getIdsMessages(string[] ids) constant returns([]Struct0 _value, uint256 _size)
func (_DMessage *DMessageCaller) GetIdsMessages(opts *bind.CallOpts, ids []string) (struct {
	Value []Struct0
	Size  *big.Int
}, error) {
	ret := new(struct {
		Value []Struct0
		Size  *big.Int
	})
	out := ret
	err := _DMessage.contract.Call(opts, out, "getIdsMessages", ids)
	return *ret, err
}

// GetIdsMessages is a free data retrieval call binding the contract method 0x6c4ef994.
//
// Solidity: function getIdsMessages(string[] ids) constant returns([]Struct0 _value, uint256 _size)
func (_DMessage *DMessageSession) GetIdsMessages(ids []string) (struct {
	Value []Struct0
	Size  *big.Int
}, error) {
	return _DMessage.Contract.GetIdsMessages(&_DMessage.CallOpts, ids)
}

// GetIdsMessages is a free data retrieval call binding the contract method 0x6c4ef994.
//
// Solidity: function getIdsMessages(string[] ids) constant returns([]Struct0 _value, uint256 _size)
func (_DMessage *DMessageCallerSession) GetIdsMessages(ids []string) (struct {
	Value []Struct0
	Size  *big.Int
}, error) {
	return _DMessage.Contract.GetIdsMessages(&_DMessage.CallOpts, ids)
}

// GetMessage is a free data retrieval call binding the contract method 0x0cc4e8d8.
//
// Solidity: function getMessage(string id) constant returns(Struct0)
func (_DMessage *DMessageCaller) GetMessage(opts *bind.CallOpts, id string) (Struct0, error) {
	var (
		ret0 = new(Struct0)
	)
	out := ret0
	err := _DMessage.contract.Call(opts, out, "getMessage", id)
	return *ret0, err
}

// GetMessage is a free data retrieval call binding the contract method 0x0cc4e8d8.
//
// Solidity: function getMessage(string id) constant returns(Struct0)
func (_DMessage *DMessageSession) GetMessage(id string) (Struct0, error) {
	return _DMessage.Contract.GetMessage(&_DMessage.CallOpts, id)
}

// GetMessage is a free data retrieval call binding the contract method 0x0cc4e8d8.
//
// Solidity: function getMessage(string id) constant returns(Struct0)
func (_DMessage *DMessageCallerSession) GetMessage(id string) (Struct0, error) {
	return _DMessage.Contract.GetMessage(&_DMessage.CallOpts, id)
}

// GetMessageValue is a free data retrieval call binding the contract method 0xc57bc642.
//
// Solidity: function getMessageValue(string id) constant returns(string, string, string)
func (_DMessage *DMessageCaller) GetMessageValue(opts *bind.CallOpts, id string) (string, string, string, error) {
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
	err := _DMessage.contract.Call(opts, out, "getMessageValue", id)
	return *ret0, *ret1, *ret2, err
}

// GetMessageValue is a free data retrieval call binding the contract method 0xc57bc642.
//
// Solidity: function getMessageValue(string id) constant returns(string, string, string)
func (_DMessage *DMessageSession) GetMessageValue(id string) (string, string, string, error) {
	return _DMessage.Contract.GetMessageValue(&_DMessage.CallOpts, id)
}

// GetMessageValue is a free data retrieval call binding the contract method 0xc57bc642.
//
// Solidity: function getMessageValue(string id) constant returns(string, string, string)
func (_DMessage *DMessageCallerSession) GetMessageValue(id string) (string, string, string, error) {
	return _DMessage.Contract.GetMessageValue(&_DMessage.CallOpts, id)
}

// GetMessages is a free data retrieval call binding the contract method 0x4be8362d.
//
// Solidity: function getMessages(uint256 start, uint256 limit) constant returns([]Struct0 _value, uint256 _size)
func (_DMessage *DMessageCaller) GetMessages(opts *bind.CallOpts, start *big.Int, limit *big.Int) (struct {
	Value []Struct0
	Size  *big.Int
}, error) {
	ret := new(struct {
		Value []Struct0
		Size  *big.Int
	})
	out := ret
	err := _DMessage.contract.Call(opts, out, "getMessages", start, limit)
	return *ret, err
}

// GetMessages is a free data retrieval call binding the contract method 0x4be8362d.
//
// Solidity: function getMessages(uint256 start, uint256 limit) constant returns([]Struct0 _value, uint256 _size)
func (_DMessage *DMessageSession) GetMessages(start *big.Int, limit *big.Int) (struct {
	Value []Struct0
	Size  *big.Int
}, error) {
	return _DMessage.Contract.GetMessages(&_DMessage.CallOpts, start, limit)
}

// GetMessages is a free data retrieval call binding the contract method 0x4be8362d.
//
// Solidity: function getMessages(uint256 start, uint256 limit) constant returns([]Struct0 _value, uint256 _size)
func (_DMessage *DMessageCallerSession) GetMessages(start *big.Int, limit *big.Int) (struct {
	Value []Struct0
	Size  *big.Int
}, error) {
	return _DMessage.Contract.GetMessages(&_DMessage.CallOpts, start, limit)
}

// GetMsgId is a free data retrieval call binding the contract method 0x95cb2638.
//
// Solidity: function getMsgId(uint256 _index) constant returns(string _value)
func (_DMessage *DMessageCaller) GetMsgId(opts *bind.CallOpts, _index *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DMessage.contract.Call(opts, out, "getMsgId", _index)
	return *ret0, err
}

// GetMsgId is a free data retrieval call binding the contract method 0x95cb2638.
//
// Solidity: function getMsgId(uint256 _index) constant returns(string _value)
func (_DMessage *DMessageSession) GetMsgId(_index *big.Int) (string, error) {
	return _DMessage.Contract.GetMsgId(&_DMessage.CallOpts, _index)
}

// GetMsgId is a free data retrieval call binding the contract method 0x95cb2638.
//
// Solidity: function getMsgId(uint256 _index) constant returns(string _value)
func (_DMessage *DMessageCallerSession) GetMsgId(_index *big.Int) (string, error) {
	return _DMessage.Contract.GetMsgId(&_DMessage.CallOpts, _index)
}

// GetMsgIds is a free data retrieval call binding the contract method 0x68bfc43c.
//
// Solidity: function getMsgIds() constant returns(string[])
func (_DMessage *DMessageCaller) GetMsgIds(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _DMessage.contract.Call(opts, out, "getMsgIds")
	return *ret0, err
}

// GetMsgIds is a free data retrieval call binding the contract method 0x68bfc43c.
//
// Solidity: function getMsgIds() constant returns(string[])
func (_DMessage *DMessageSession) GetMsgIds() ([]string, error) {
	return _DMessage.Contract.GetMsgIds(&_DMessage.CallOpts)
}

// GetMsgIds is a free data retrieval call binding the contract method 0x68bfc43c.
//
// Solidity: function getMsgIds() constant returns(string[])
func (_DMessage *DMessageCallerSession) GetMsgIds() ([]string, error) {
	return _DMessage.Contract.GetMsgIds(&_DMessage.CallOpts)
}

// GetMsgLength is a free data retrieval call binding the contract method 0x085c00e1.
//
// Solidity: function getMsgLength() constant returns(uint256)
func (_DMessage *DMessageCaller) GetMsgLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DMessage.contract.Call(opts, out, "getMsgLength")
	return *ret0, err
}

// GetMsgLength is a free data retrieval call binding the contract method 0x085c00e1.
//
// Solidity: function getMsgLength() constant returns(uint256)
func (_DMessage *DMessageSession) GetMsgLength() (*big.Int, error) {
	return _DMessage.Contract.GetMsgLength(&_DMessage.CallOpts)
}

// GetMsgLength is a free data retrieval call binding the contract method 0x085c00e1.
//
// Solidity: function getMsgLength() constant returns(uint256)
func (_DMessage *DMessageCallerSession) GetMsgLength() (*big.Int, error) {
	return _DMessage.Contract.GetMsgLength(&_DMessage.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DMessage *DMessageCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DMessage.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DMessage *DMessageSession) IsOwner() (bool, error) {
	return _DMessage.Contract.IsOwner(&_DMessage.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_DMessage *DMessageCallerSession) IsOwner() (bool, error) {
	return _DMessage.Contract.IsOwner(&_DMessage.CallOpts)
}

// IsWriter is a free data retrieval call binding the contract method 0x4d6ee9fd.
//
// Solidity: function isWriter() constant returns(bool)
func (_DMessage *DMessageCaller) IsWriter(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DMessage.contract.Call(opts, out, "isWriter")
	return *ret0, err
}

// IsWriter is a free data retrieval call binding the contract method 0x4d6ee9fd.
//
// Solidity: function isWriter() constant returns(bool)
func (_DMessage *DMessageSession) IsWriter() (bool, error) {
	return _DMessage.Contract.IsWriter(&_DMessage.CallOpts)
}

// IsWriter is a free data retrieval call binding the contract method 0x4d6ee9fd.
//
// Solidity: function isWriter() constant returns(bool)
func (_DMessage *DMessageCallerSession) IsWriter() (bool, error) {
	return _DMessage.Contract.IsWriter(&_DMessage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DMessage *DMessageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DMessage.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DMessage *DMessageSession) Owner() (common.Address, error) {
	return _DMessage.Contract.Owner(&_DMessage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_DMessage *DMessageCallerSession) Owner() (common.Address, error) {
	return _DMessage.Contract.Owner(&_DMessage.CallOpts)
}

// Writer is a free data retrieval call binding the contract method 0x453a2abc.
//
// Solidity: function writer() constant returns(address)
func (_DMessage *DMessageCaller) Writer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DMessage.contract.Call(opts, out, "writer")
	return *ret0, err
}

// Writer is a free data retrieval call binding the contract method 0x453a2abc.
//
// Solidity: function writer() constant returns(address)
func (_DMessage *DMessageSession) Writer() (common.Address, error) {
	return _DMessage.Contract.Writer(&_DMessage.CallOpts)
}

// Writer is a free data retrieval call binding the contract method 0x453a2abc.
//
// Solidity: function writer() constant returns(address)
func (_DMessage *DMessageCallerSession) Writer() (common.Address, error) {
	return _DMessage.Contract.Writer(&_DMessage.CallOpts)
}

// Writers is a free data retrieval call binding the contract method 0x8b2f5369.
//
// Solidity: function writers() constant returns(address[] _value)
func (_DMessage *DMessageCaller) Writers(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _DMessage.contract.Call(opts, out, "writers")
	return *ret0, err
}

// Writers is a free data retrieval call binding the contract method 0x8b2f5369.
//
// Solidity: function writers() constant returns(address[] _value)
func (_DMessage *DMessageSession) Writers() ([]common.Address, error) {
	return _DMessage.Contract.Writers(&_DMessage.CallOpts)
}

// Writers is a free data retrieval call binding the contract method 0x8b2f5369.
//
// Solidity: function writers() constant returns(address[] _value)
func (_DMessage *DMessageCallerSession) Writers() ([]common.Address, error) {
	return _DMessage.Contract.Writers(&_DMessage.CallOpts)
}

// AddMessage is a paid mutator transaction binding the contract method 0xb98464b3.
//
// Solidity: function addMessage(string id, string content, string version) returns(bool)
func (_DMessage *DMessageTransactor) AddMessage(opts *bind.TransactOpts, id string, content string, version string) (*types.Transaction, error) {
	return _DMessage.contract.Transact(opts, "addMessage", id, content, version)
}

// AddMessage is a paid mutator transaction binding the contract method 0xb98464b3.
//
// Solidity: function addMessage(string id, string content, string version) returns(bool)
func (_DMessage *DMessageSession) AddMessage(id string, content string, version string) (*types.Transaction, error) {
	return _DMessage.Contract.AddMessage(&_DMessage.TransactOpts, id, content, version)
}

// AddMessage is a paid mutator transaction binding the contract method 0xb98464b3.
//
// Solidity: function addMessage(string id, string content, string version) returns(bool)
func (_DMessage *DMessageTransactorSession) AddMessage(id string, content string, version string) (*types.Transaction, error) {
	return _DMessage.Contract.AddMessage(&_DMessage.TransactOpts, id, content, version)
}

// DecreaseWritership is a paid mutator transaction binding the contract method 0x9f217bdf.
//
// Solidity: function decreaseWritership(address oldWriter) returns()
func (_DMessage *DMessageTransactor) DecreaseWritership(opts *bind.TransactOpts, oldWriter common.Address) (*types.Transaction, error) {
	return _DMessage.contract.Transact(opts, "decreaseWritership", oldWriter)
}

// DecreaseWritership is a paid mutator transaction binding the contract method 0x9f217bdf.
//
// Solidity: function decreaseWritership(address oldWriter) returns()
func (_DMessage *DMessageSession) DecreaseWritership(oldWriter common.Address) (*types.Transaction, error) {
	return _DMessage.Contract.DecreaseWritership(&_DMessage.TransactOpts, oldWriter)
}

// DecreaseWritership is a paid mutator transaction binding the contract method 0x9f217bdf.
//
// Solidity: function decreaseWritership(address oldWriter) returns()
func (_DMessage *DMessageTransactorSession) DecreaseWritership(oldWriter common.Address) (*types.Transaction, error) {
	return _DMessage.Contract.DecreaseWritership(&_DMessage.TransactOpts, oldWriter)
}

// DelMessage is a paid mutator transaction binding the contract method 0x2309bc96.
//
// Solidity: function delMessage(string id) returns(Struct0, bool)
func (_DMessage *DMessageTransactor) DelMessage(opts *bind.TransactOpts, id string) (*types.Transaction, error) {
	return _DMessage.contract.Transact(opts, "delMessage", id)
}

// DelMessage is a paid mutator transaction binding the contract method 0x2309bc96.
//
// Solidity: function delMessage(string id) returns(Struct0, bool)
func (_DMessage *DMessageSession) DelMessage(id string) (*types.Transaction, error) {
	return _DMessage.Contract.DelMessage(&_DMessage.TransactOpts, id)
}

// DelMessage is a paid mutator transaction binding the contract method 0x2309bc96.
//
// Solidity: function delMessage(string id) returns(Struct0, bool)
func (_DMessage *DMessageTransactorSession) DelMessage(id string) (*types.Transaction, error) {
	return _DMessage.Contract.DelMessage(&_DMessage.TransactOpts, id)
}

// IncreasedWritership is a paid mutator transaction binding the contract method 0x14b4d411.
//
// Solidity: function increasedWritership(address newWriter) returns()
func (_DMessage *DMessageTransactor) IncreasedWritership(opts *bind.TransactOpts, newWriter common.Address) (*types.Transaction, error) {
	return _DMessage.contract.Transact(opts, "increasedWritership", newWriter)
}

// IncreasedWritership is a paid mutator transaction binding the contract method 0x14b4d411.
//
// Solidity: function increasedWritership(address newWriter) returns()
func (_DMessage *DMessageSession) IncreasedWritership(newWriter common.Address) (*types.Transaction, error) {
	return _DMessage.Contract.IncreasedWritership(&_DMessage.TransactOpts, newWriter)
}

// IncreasedWritership is a paid mutator transaction binding the contract method 0x14b4d411.
//
// Solidity: function increasedWritership(address newWriter) returns()
func (_DMessage *DMessageTransactorSession) IncreasedWritership(newWriter common.Address) (*types.Transaction, error) {
	return _DMessage.Contract.IncreasedWritership(&_DMessage.TransactOpts, newWriter)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DMessage *DMessageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DMessage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DMessage *DMessageSession) RenounceOwnership() (*types.Transaction, error) {
	return _DMessage.Contract.RenounceOwnership(&_DMessage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DMessage *DMessageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DMessage.Contract.RenounceOwnership(&_DMessage.TransactOpts)
}

// RenounceWritership is a paid mutator transaction binding the contract method 0x4075e83e.
//
// Solidity: function renounceWritership() returns()
func (_DMessage *DMessageTransactor) RenounceWritership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DMessage.contract.Transact(opts, "renounceWritership")
}

// RenounceWritership is a paid mutator transaction binding the contract method 0x4075e83e.
//
// Solidity: function renounceWritership() returns()
func (_DMessage *DMessageSession) RenounceWritership() (*types.Transaction, error) {
	return _DMessage.Contract.RenounceWritership(&_DMessage.TransactOpts)
}

// RenounceWritership is a paid mutator transaction binding the contract method 0x4075e83e.
//
// Solidity: function renounceWritership() returns()
func (_DMessage *DMessageTransactorSession) RenounceWritership() (*types.Transaction, error) {
	return _DMessage.Contract.RenounceWritership(&_DMessage.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DMessage *DMessageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DMessage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DMessage *DMessageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DMessage.Contract.TransferOwnership(&_DMessage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DMessage *DMessageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DMessage.Contract.TransferOwnership(&_DMessage.TransactOpts, newOwner)
}

// UpdateMessage is a paid mutator transaction binding the contract method 0x6a44815e.
//
// Solidity: function updateMessage(string id, string content, string version) returns(bool)
func (_DMessage *DMessageTransactor) UpdateMessage(opts *bind.TransactOpts, id string, content string, version string) (*types.Transaction, error) {
	return _DMessage.contract.Transact(opts, "updateMessage", id, content, version)
}

// UpdateMessage is a paid mutator transaction binding the contract method 0x6a44815e.
//
// Solidity: function updateMessage(string id, string content, string version) returns(bool)
func (_DMessage *DMessageSession) UpdateMessage(id string, content string, version string) (*types.Transaction, error) {
	return _DMessage.Contract.UpdateMessage(&_DMessage.TransactOpts, id, content, version)
}

// UpdateMessage is a paid mutator transaction binding the contract method 0x6a44815e.
//
// Solidity: function updateMessage(string id, string content, string version) returns(bool)
func (_DMessage *DMessageTransactorSession) UpdateMessage(id string, content string, version string) (*types.Transaction, error) {
	return _DMessage.Contract.UpdateMessage(&_DMessage.TransactOpts, id, content, version)
}

// DMessageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DMessage contract.
type DMessageOwnershipTransferredIterator struct {
	Event *DMessageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DMessageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DMessageOwnershipTransferred)
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
		it.Event = new(DMessageOwnershipTransferred)
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
func (it *DMessageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DMessageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DMessageOwnershipTransferred represents a OwnershipTransferred event raised by the DMessage contract.
type DMessageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DMessage *DMessageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DMessageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DMessage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DMessageOwnershipTransferredIterator{contract: _DMessage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DMessage *DMessageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DMessageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DMessage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DMessageOwnershipTransferred)
				if err := _DMessage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DMessage *DMessageFilterer) ParseOwnershipTransferred(log types.Log) (*DMessageOwnershipTransferred, error) {
	event := new(DMessageOwnershipTransferred)
	if err := _DMessage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DMessageWritershipDecreasedIterator is returned from FilterWritershipDecreased and is used to iterate over the raw logs and unpacked data for WritershipDecreased events raised by the DMessage contract.
type DMessageWritershipDecreasedIterator struct {
	Event *DMessageWritershipDecreased // Event containing the contract specifics and raw log

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
func (it *DMessageWritershipDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DMessageWritershipDecreased)
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
		it.Event = new(DMessageWritershipDecreased)
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
func (it *DMessageWritershipDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DMessageWritershipDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DMessageWritershipDecreased represents a WritershipDecreased event raised by the DMessage contract.
type DMessageWritershipDecreased struct {
	OldWriter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWritershipDecreased is a free log retrieval operation binding the contract event 0xb24c36b70cc67df9ea22f053c15c4a41e7170d7be0637d2637a3445a266ac6af.
//
// Solidity: event WritershipDecreased(address indexed oldWriter)
func (_DMessage *DMessageFilterer) FilterWritershipDecreased(opts *bind.FilterOpts, oldWriter []common.Address) (*DMessageWritershipDecreasedIterator, error) {

	var oldWriterRule []interface{}
	for _, oldWriterItem := range oldWriter {
		oldWriterRule = append(oldWriterRule, oldWriterItem)
	}

	logs, sub, err := _DMessage.contract.FilterLogs(opts, "WritershipDecreased", oldWriterRule)
	if err != nil {
		return nil, err
	}
	return &DMessageWritershipDecreasedIterator{contract: _DMessage.contract, event: "WritershipDecreased", logs: logs, sub: sub}, nil
}

// WatchWritershipDecreased is a free log subscription operation binding the contract event 0xb24c36b70cc67df9ea22f053c15c4a41e7170d7be0637d2637a3445a266ac6af.
//
// Solidity: event WritershipDecreased(address indexed oldWriter)
func (_DMessage *DMessageFilterer) WatchWritershipDecreased(opts *bind.WatchOpts, sink chan<- *DMessageWritershipDecreased, oldWriter []common.Address) (event.Subscription, error) {

	var oldWriterRule []interface{}
	for _, oldWriterItem := range oldWriter {
		oldWriterRule = append(oldWriterRule, oldWriterItem)
	}

	logs, sub, err := _DMessage.contract.WatchLogs(opts, "WritershipDecreased", oldWriterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DMessageWritershipDecreased)
				if err := _DMessage.contract.UnpackLog(event, "WritershipDecreased", log); err != nil {
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
func (_DMessage *DMessageFilterer) ParseWritershipDecreased(log types.Log) (*DMessageWritershipDecreased, error) {
	event := new(DMessageWritershipDecreased)
	if err := _DMessage.contract.UnpackLog(event, "WritershipDecreased", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DMessageWritershipIncreasedIterator is returned from FilterWritershipIncreased and is used to iterate over the raw logs and unpacked data for WritershipIncreased events raised by the DMessage contract.
type DMessageWritershipIncreasedIterator struct {
	Event *DMessageWritershipIncreased // Event containing the contract specifics and raw log

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
func (it *DMessageWritershipIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DMessageWritershipIncreased)
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
		it.Event = new(DMessageWritershipIncreased)
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
func (it *DMessageWritershipIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DMessageWritershipIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DMessageWritershipIncreased represents a WritershipIncreased event raised by the DMessage contract.
type DMessageWritershipIncreased struct {
	NewWriter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWritershipIncreased is a free log retrieval operation binding the contract event 0x8f2adc2e2f2723d375f8bb1359cb1027f713945395ca36c42b83dd94c5bdc7f2.
//
// Solidity: event WritershipIncreased(address indexed newWriter)
func (_DMessage *DMessageFilterer) FilterWritershipIncreased(opts *bind.FilterOpts, newWriter []common.Address) (*DMessageWritershipIncreasedIterator, error) {

	var newWriterRule []interface{}
	for _, newWriterItem := range newWriter {
		newWriterRule = append(newWriterRule, newWriterItem)
	}

	logs, sub, err := _DMessage.contract.FilterLogs(opts, "WritershipIncreased", newWriterRule)
	if err != nil {
		return nil, err
	}
	return &DMessageWritershipIncreasedIterator{contract: _DMessage.contract, event: "WritershipIncreased", logs: logs, sub: sub}, nil
}

// WatchWritershipIncreased is a free log subscription operation binding the contract event 0x8f2adc2e2f2723d375f8bb1359cb1027f713945395ca36c42b83dd94c5bdc7f2.
//
// Solidity: event WritershipIncreased(address indexed newWriter)
func (_DMessage *DMessageFilterer) WatchWritershipIncreased(opts *bind.WatchOpts, sink chan<- *DMessageWritershipIncreased, newWriter []common.Address) (event.Subscription, error) {

	var newWriterRule []interface{}
	for _, newWriterItem := range newWriter {
		newWriterRule = append(newWriterRule, newWriterItem)
	}

	logs, sub, err := _DMessage.contract.WatchLogs(opts, "WritershipIncreased", newWriterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DMessageWritershipIncreased)
				if err := _DMessage.contract.UnpackLog(event, "WritershipIncreased", log); err != nil {
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
func (_DMessage *DMessageFilterer) ParseWritershipIncreased(log types.Log) (*DMessageWritershipIncreased, error) {
	event := new(DMessageWritershipIncreased)
	if err := _DMessage.contract.UnpackLog(event, "WritershipIncreased", log); err != nil {
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

// WriteableABI is the input ABI used to generate the binding from.
const WriteableABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldWriter\",\"type\":\"address\"}],\"name\":\"WritershipDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newWriter\",\"type\":\"address\"}],\"name\":\"WritershipIncreased\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldWriter\",\"type\":\"address\"}],\"name\":\"decreaseWritership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWriter\",\"type\":\"address\"}],\"name\":\"increasedWritership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isWriter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceWritership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"writer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"writers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_value\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// WriteableFuncSigs maps the 4-byte function signature to its string representation.
var WriteableFuncSigs = map[string]string{
	"9f217bdf": "decreaseWritership(address)",
	"14b4d411": "increasedWritership(address)",
	"8f32d59b": "isOwner()",
	"4d6ee9fd": "isWriter()",
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"4075e83e": "renounceWritership()",
	"f2fde38b": "transferOwnership(address)",
	"453a2abc": "writer()",
	"8b2f5369": "writers()",
}

// Writeable is an auto generated Go binding around an Ethereum contract.
type Writeable struct {
	WriteableCaller     // Read-only binding to the contract
	WriteableTransactor // Write-only binding to the contract
	WriteableFilterer   // Log filterer for contract events
}

// WriteableCaller is an auto generated read-only Go binding around an Ethereum contract.
type WriteableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WriteableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WriteableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WriteableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WriteableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WriteableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WriteableSession struct {
	Contract     *Writeable        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WriteableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WriteableCallerSession struct {
	Contract *WriteableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// WriteableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WriteableTransactorSession struct {
	Contract     *WriteableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// WriteableRaw is an auto generated low-level Go binding around an Ethereum contract.
type WriteableRaw struct {
	Contract *Writeable // Generic contract binding to access the raw methods on
}

// WriteableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WriteableCallerRaw struct {
	Contract *WriteableCaller // Generic read-only contract binding to access the raw methods on
}

// WriteableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WriteableTransactorRaw struct {
	Contract *WriteableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWriteable creates a new instance of Writeable, bound to a specific deployed contract.
func NewWriteable(address common.Address, backend bind.ContractBackend) (*Writeable, error) {
	contract, err := bindWriteable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Writeable{WriteableCaller: WriteableCaller{contract: contract}, WriteableTransactor: WriteableTransactor{contract: contract}, WriteableFilterer: WriteableFilterer{contract: contract}}, nil
}

// NewWriteableCaller creates a new read-only instance of Writeable, bound to a specific deployed contract.
func NewWriteableCaller(address common.Address, caller bind.ContractCaller) (*WriteableCaller, error) {
	contract, err := bindWriteable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WriteableCaller{contract: contract}, nil
}

// NewWriteableTransactor creates a new write-only instance of Writeable, bound to a specific deployed contract.
func NewWriteableTransactor(address common.Address, transactor bind.ContractTransactor) (*WriteableTransactor, error) {
	contract, err := bindWriteable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WriteableTransactor{contract: contract}, nil
}

// NewWriteableFilterer creates a new log filterer instance of Writeable, bound to a specific deployed contract.
func NewWriteableFilterer(address common.Address, filterer bind.ContractFilterer) (*WriteableFilterer, error) {
	contract, err := bindWriteable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WriteableFilterer{contract: contract}, nil
}

// bindWriteable binds a generic wrapper to an already deployed contract.
func bindWriteable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WriteableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Writeable *WriteableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Writeable.Contract.WriteableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Writeable *WriteableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Writeable.Contract.WriteableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Writeable *WriteableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Writeable.Contract.WriteableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Writeable *WriteableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Writeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Writeable *WriteableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Writeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Writeable *WriteableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Writeable.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Writeable *WriteableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Writeable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Writeable *WriteableSession) IsOwner() (bool, error) {
	return _Writeable.Contract.IsOwner(&_Writeable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Writeable *WriteableCallerSession) IsOwner() (bool, error) {
	return _Writeable.Contract.IsOwner(&_Writeable.CallOpts)
}

// IsWriter is a free data retrieval call binding the contract method 0x4d6ee9fd.
//
// Solidity: function isWriter() constant returns(bool)
func (_Writeable *WriteableCaller) IsWriter(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Writeable.contract.Call(opts, out, "isWriter")
	return *ret0, err
}

// IsWriter is a free data retrieval call binding the contract method 0x4d6ee9fd.
//
// Solidity: function isWriter() constant returns(bool)
func (_Writeable *WriteableSession) IsWriter() (bool, error) {
	return _Writeable.Contract.IsWriter(&_Writeable.CallOpts)
}

// IsWriter is a free data retrieval call binding the contract method 0x4d6ee9fd.
//
// Solidity: function isWriter() constant returns(bool)
func (_Writeable *WriteableCallerSession) IsWriter() (bool, error) {
	return _Writeable.Contract.IsWriter(&_Writeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Writeable *WriteableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Writeable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Writeable *WriteableSession) Owner() (common.Address, error) {
	return _Writeable.Contract.Owner(&_Writeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Writeable *WriteableCallerSession) Owner() (common.Address, error) {
	return _Writeable.Contract.Owner(&_Writeable.CallOpts)
}

// Writer is a free data retrieval call binding the contract method 0x453a2abc.
//
// Solidity: function writer() constant returns(address)
func (_Writeable *WriteableCaller) Writer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Writeable.contract.Call(opts, out, "writer")
	return *ret0, err
}

// Writer is a free data retrieval call binding the contract method 0x453a2abc.
//
// Solidity: function writer() constant returns(address)
func (_Writeable *WriteableSession) Writer() (common.Address, error) {
	return _Writeable.Contract.Writer(&_Writeable.CallOpts)
}

// Writer is a free data retrieval call binding the contract method 0x453a2abc.
//
// Solidity: function writer() constant returns(address)
func (_Writeable *WriteableCallerSession) Writer() (common.Address, error) {
	return _Writeable.Contract.Writer(&_Writeable.CallOpts)
}

// Writers is a free data retrieval call binding the contract method 0x8b2f5369.
//
// Solidity: function writers() constant returns(address[] _value)
func (_Writeable *WriteableCaller) Writers(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Writeable.contract.Call(opts, out, "writers")
	return *ret0, err
}

// Writers is a free data retrieval call binding the contract method 0x8b2f5369.
//
// Solidity: function writers() constant returns(address[] _value)
func (_Writeable *WriteableSession) Writers() ([]common.Address, error) {
	return _Writeable.Contract.Writers(&_Writeable.CallOpts)
}

// Writers is a free data retrieval call binding the contract method 0x8b2f5369.
//
// Solidity: function writers() constant returns(address[] _value)
func (_Writeable *WriteableCallerSession) Writers() ([]common.Address, error) {
	return _Writeable.Contract.Writers(&_Writeable.CallOpts)
}

// DecreaseWritership is a paid mutator transaction binding the contract method 0x9f217bdf.
//
// Solidity: function decreaseWritership(address oldWriter) returns()
func (_Writeable *WriteableTransactor) DecreaseWritership(opts *bind.TransactOpts, oldWriter common.Address) (*types.Transaction, error) {
	return _Writeable.contract.Transact(opts, "decreaseWritership", oldWriter)
}

// DecreaseWritership is a paid mutator transaction binding the contract method 0x9f217bdf.
//
// Solidity: function decreaseWritership(address oldWriter) returns()
func (_Writeable *WriteableSession) DecreaseWritership(oldWriter common.Address) (*types.Transaction, error) {
	return _Writeable.Contract.DecreaseWritership(&_Writeable.TransactOpts, oldWriter)
}

// DecreaseWritership is a paid mutator transaction binding the contract method 0x9f217bdf.
//
// Solidity: function decreaseWritership(address oldWriter) returns()
func (_Writeable *WriteableTransactorSession) DecreaseWritership(oldWriter common.Address) (*types.Transaction, error) {
	return _Writeable.Contract.DecreaseWritership(&_Writeable.TransactOpts, oldWriter)
}

// IncreasedWritership is a paid mutator transaction binding the contract method 0x14b4d411.
//
// Solidity: function increasedWritership(address newWriter) returns()
func (_Writeable *WriteableTransactor) IncreasedWritership(opts *bind.TransactOpts, newWriter common.Address) (*types.Transaction, error) {
	return _Writeable.contract.Transact(opts, "increasedWritership", newWriter)
}

// IncreasedWritership is a paid mutator transaction binding the contract method 0x14b4d411.
//
// Solidity: function increasedWritership(address newWriter) returns()
func (_Writeable *WriteableSession) IncreasedWritership(newWriter common.Address) (*types.Transaction, error) {
	return _Writeable.Contract.IncreasedWritership(&_Writeable.TransactOpts, newWriter)
}

// IncreasedWritership is a paid mutator transaction binding the contract method 0x14b4d411.
//
// Solidity: function increasedWritership(address newWriter) returns()
func (_Writeable *WriteableTransactorSession) IncreasedWritership(newWriter common.Address) (*types.Transaction, error) {
	return _Writeable.Contract.IncreasedWritership(&_Writeable.TransactOpts, newWriter)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Writeable *WriteableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Writeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Writeable *WriteableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Writeable.Contract.RenounceOwnership(&_Writeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Writeable *WriteableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Writeable.Contract.RenounceOwnership(&_Writeable.TransactOpts)
}

// RenounceWritership is a paid mutator transaction binding the contract method 0x4075e83e.
//
// Solidity: function renounceWritership() returns()
func (_Writeable *WriteableTransactor) RenounceWritership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Writeable.contract.Transact(opts, "renounceWritership")
}

// RenounceWritership is a paid mutator transaction binding the contract method 0x4075e83e.
//
// Solidity: function renounceWritership() returns()
func (_Writeable *WriteableSession) RenounceWritership() (*types.Transaction, error) {
	return _Writeable.Contract.RenounceWritership(&_Writeable.TransactOpts)
}

// RenounceWritership is a paid mutator transaction binding the contract method 0x4075e83e.
//
// Solidity: function renounceWritership() returns()
func (_Writeable *WriteableTransactorSession) RenounceWritership() (*types.Transaction, error) {
	return _Writeable.Contract.RenounceWritership(&_Writeable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Writeable *WriteableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Writeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Writeable *WriteableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Writeable.Contract.TransferOwnership(&_Writeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Writeable *WriteableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Writeable.Contract.TransferOwnership(&_Writeable.TransactOpts, newOwner)
}

// WriteableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Writeable contract.
type WriteableOwnershipTransferredIterator struct {
	Event *WriteableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WriteableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WriteableOwnershipTransferred)
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
		it.Event = new(WriteableOwnershipTransferred)
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
func (it *WriteableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WriteableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WriteableOwnershipTransferred represents a OwnershipTransferred event raised by the Writeable contract.
type WriteableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Writeable *WriteableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WriteableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Writeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WriteableOwnershipTransferredIterator{contract: _Writeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Writeable *WriteableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WriteableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Writeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WriteableOwnershipTransferred)
				if err := _Writeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Writeable *WriteableFilterer) ParseOwnershipTransferred(log types.Log) (*WriteableOwnershipTransferred, error) {
	event := new(WriteableOwnershipTransferred)
	if err := _Writeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WriteableWritershipDecreasedIterator is returned from FilterWritershipDecreased and is used to iterate over the raw logs and unpacked data for WritershipDecreased events raised by the Writeable contract.
type WriteableWritershipDecreasedIterator struct {
	Event *WriteableWritershipDecreased // Event containing the contract specifics and raw log

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
func (it *WriteableWritershipDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WriteableWritershipDecreased)
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
		it.Event = new(WriteableWritershipDecreased)
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
func (it *WriteableWritershipDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WriteableWritershipDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WriteableWritershipDecreased represents a WritershipDecreased event raised by the Writeable contract.
type WriteableWritershipDecreased struct {
	OldWriter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWritershipDecreased is a free log retrieval operation binding the contract event 0xb24c36b70cc67df9ea22f053c15c4a41e7170d7be0637d2637a3445a266ac6af.
//
// Solidity: event WritershipDecreased(address indexed oldWriter)
func (_Writeable *WriteableFilterer) FilterWritershipDecreased(opts *bind.FilterOpts, oldWriter []common.Address) (*WriteableWritershipDecreasedIterator, error) {

	var oldWriterRule []interface{}
	for _, oldWriterItem := range oldWriter {
		oldWriterRule = append(oldWriterRule, oldWriterItem)
	}

	logs, sub, err := _Writeable.contract.FilterLogs(opts, "WritershipDecreased", oldWriterRule)
	if err != nil {
		return nil, err
	}
	return &WriteableWritershipDecreasedIterator{contract: _Writeable.contract, event: "WritershipDecreased", logs: logs, sub: sub}, nil
}

// WatchWritershipDecreased is a free log subscription operation binding the contract event 0xb24c36b70cc67df9ea22f053c15c4a41e7170d7be0637d2637a3445a266ac6af.
//
// Solidity: event WritershipDecreased(address indexed oldWriter)
func (_Writeable *WriteableFilterer) WatchWritershipDecreased(opts *bind.WatchOpts, sink chan<- *WriteableWritershipDecreased, oldWriter []common.Address) (event.Subscription, error) {

	var oldWriterRule []interface{}
	for _, oldWriterItem := range oldWriter {
		oldWriterRule = append(oldWriterRule, oldWriterItem)
	}

	logs, sub, err := _Writeable.contract.WatchLogs(opts, "WritershipDecreased", oldWriterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WriteableWritershipDecreased)
				if err := _Writeable.contract.UnpackLog(event, "WritershipDecreased", log); err != nil {
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
func (_Writeable *WriteableFilterer) ParseWritershipDecreased(log types.Log) (*WriteableWritershipDecreased, error) {
	event := new(WriteableWritershipDecreased)
	if err := _Writeable.contract.UnpackLog(event, "WritershipDecreased", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WriteableWritershipIncreasedIterator is returned from FilterWritershipIncreased and is used to iterate over the raw logs and unpacked data for WritershipIncreased events raised by the Writeable contract.
type WriteableWritershipIncreasedIterator struct {
	Event *WriteableWritershipIncreased // Event containing the contract specifics and raw log

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
func (it *WriteableWritershipIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WriteableWritershipIncreased)
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
		it.Event = new(WriteableWritershipIncreased)
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
func (it *WriteableWritershipIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WriteableWritershipIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WriteableWritershipIncreased represents a WritershipIncreased event raised by the Writeable contract.
type WriteableWritershipIncreased struct {
	NewWriter common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWritershipIncreased is a free log retrieval operation binding the contract event 0x8f2adc2e2f2723d375f8bb1359cb1027f713945395ca36c42b83dd94c5bdc7f2.
//
// Solidity: event WritershipIncreased(address indexed newWriter)
func (_Writeable *WriteableFilterer) FilterWritershipIncreased(opts *bind.FilterOpts, newWriter []common.Address) (*WriteableWritershipIncreasedIterator, error) {

	var newWriterRule []interface{}
	for _, newWriterItem := range newWriter {
		newWriterRule = append(newWriterRule, newWriterItem)
	}

	logs, sub, err := _Writeable.contract.FilterLogs(opts, "WritershipIncreased", newWriterRule)
	if err != nil {
		return nil, err
	}
	return &WriteableWritershipIncreasedIterator{contract: _Writeable.contract, event: "WritershipIncreased", logs: logs, sub: sub}, nil
}

// WatchWritershipIncreased is a free log subscription operation binding the contract event 0x8f2adc2e2f2723d375f8bb1359cb1027f713945395ca36c42b83dd94c5bdc7f2.
//
// Solidity: event WritershipIncreased(address indexed newWriter)
func (_Writeable *WriteableFilterer) WatchWritershipIncreased(opts *bind.WatchOpts, sink chan<- *WriteableWritershipIncreased, newWriter []common.Address) (event.Subscription, error) {

	var newWriterRule []interface{}
	for _, newWriterItem := range newWriter {
		newWriterRule = append(newWriterRule, newWriterItem)
	}

	logs, sub, err := _Writeable.contract.WatchLogs(opts, "WritershipIncreased", newWriterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WriteableWritershipIncreased)
				if err := _Writeable.contract.UnpackLog(event, "WritershipIncreased", log); err != nil {
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
func (_Writeable *WriteableFilterer) ParseWritershipIncreased(log types.Log) (*WriteableWritershipIncreased, error) {
	event := new(WriteableWritershipIncreased)
	if err := _Writeable.contract.UnpackLog(event, "WritershipIncreased", log); err != nil {
		return nil, err
	}
	return event, nil
}
