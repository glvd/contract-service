// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
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

// VideoInfoABI is the input ABI used to generate the binding from.
const VideoInfoABI = "[{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"findNo\",\"type\":\"string\"},{\"name\":\"bangumi\",\"type\":\"string\"},{\"name\":\"intro\",\"type\":\"string\"},{\"name\":\"Alias\",\"type\":\"string[]\"},{\"name\":\"role\",\"type\":\"string[]\"},{\"name\":\"thumbHash\",\"type\":\"string\"},{\"name\":\"posterHash\",\"type\":\"string\"},{\"name\":\"uncensored\",\"type\":\"bool\"}],\"name\":\"v\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"bangumi\",\"type\":\"string\"},{\"name\":\"visit\",\"type\":\"uint32\"},{\"name\":\"series\",\"type\":\"string\"},{\"name\":\"tags\",\"type\":\"string[]\"},{\"name\":\"length\",\"type\":\"string\"},{\"name\":\"sharpness\",\"type\":\"string\"},{\"name\":\"format\",\"type\":\"string\"}],\"name\":\"d\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"bangumi\",\"type\":\"string\"},{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"m3u8\",\"type\":\"string\"},{\"name\":\"sourceHash\",\"type\":\"string\"},{\"name\":\"m3u8Hash\",\"type\":\"string\"}],\"name\":\"s\",\"type\":\"tuple\"}],\"name\":\"add\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"findNo\",\"type\":\"string\"},{\"name\":\"bangumi\",\"type\":\"string\"},{\"name\":\"intro\",\"type\":\"string\"},{\"name\":\"Alias\",\"type\":\"string[]\"},{\"name\":\"role\",\"type\":\"string[]\"},{\"name\":\"thumbHash\",\"type\":\"string\"},{\"name\":\"posterHash\",\"type\":\"string\"},{\"name\":\"uncensored\",\"type\":\"bool\"}],\"name\":\"v\",\"type\":\"tuple\"}],\"name\":\"addHot\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"date\",\"type\":\"string\"},{\"components\":[{\"name\":\"findNo\",\"type\":\"string\"},{\"name\":\"bangumi\",\"type\":\"string\"},{\"name\":\"intro\",\"type\":\"string\"},{\"name\":\"Alias\",\"type\":\"string[]\"},{\"name\":\"role\",\"type\":\"string[]\"},{\"name\":\"thumbHash\",\"type\":\"string\"},{\"name\":\"posterHash\",\"type\":\"string\"},{\"name\":\"uncensored\",\"type\":\"bool\"}],\"name\":\"v\",\"type\":\"tuple\"}],\"name\":\"addNew\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"findNo\",\"type\":\"string\"},{\"name\":\"bangumi\",\"type\":\"string\"},{\"name\":\"intro\",\"type\":\"string\"},{\"name\":\"Alias\",\"type\":\"string[]\"},{\"name\":\"role\",\"type\":\"string[]\"},{\"name\":\"thumbHash\",\"type\":\"string\"},{\"name\":\"posterHash\",\"type\":\"string\"},{\"name\":\"uncensored\",\"type\":\"bool\"}],\"name\":\"v\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"bangumi\",\"type\":\"string\"},{\"name\":\"visit\",\"type\":\"uint32\"},{\"name\":\"series\",\"type\":\"string\"},{\"name\":\"tags\",\"type\":\"string[]\"},{\"name\":\"length\",\"type\":\"string\"},{\"name\":\"sharpness\",\"type\":\"string\"},{\"name\":\"format\",\"type\":\"string\"}],\"name\":\"d\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"bangumi\",\"type\":\"string\"},{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"m3u8\",\"type\":\"string\"},{\"name\":\"sourceHash\",\"type\":\"string\"},{\"name\":\"m3u8Hash\",\"type\":\"string\"}],\"name\":\"s\",\"type\":\"tuple\"}],\"name\":\"replace\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"idx\",\"type\":\"uint32\"},{\"components\":[{\"name\":\"findNo\",\"type\":\"string\"},{\"name\":\"bangumi\",\"type\":\"string\"},{\"name\":\"intro\",\"type\":\"string\"},{\"name\":\"Alias\",\"type\":\"string[]\"},{\"name\":\"role\",\"type\":\"string[]\"},{\"name\":\"thumbHash\",\"type\":\"string\"},{\"name\":\"posterHash\",\"type\":\"string\"},{\"name\":\"uncensored\",\"type\":\"bool\"}],\"name\":\"v\",\"type\":\"tuple\"}],\"name\":\"replaceHot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"date\",\"type\":\"string\"},{\"name\":\"idx\",\"type\":\"uint32\"},{\"components\":[{\"name\":\"findNo\",\"type\":\"string\"},{\"name\":\"bangumi\",\"type\":\"string\"},{\"name\":\"intro\",\"type\":\"string\"},{\"name\":\"Alias\",\"type\":\"string[]\"},{\"name\":\"role\",\"type\":\"string[]\"},{\"name\":\"thumbHash\",\"type\":\"string\"},{\"name\":\"posterHash\",\"type\":\"string\"},{\"name\":\"uncensored\",\"type\":\"bool\"}],\"name\":\"v\",\"type\":\"tuple\"}],\"name\":\"replaceNew\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getHotVideos\",\"outputs\":[{\"components\":[{\"name\":\"findNo\",\"type\":\"string\"},{\"name\":\"bangumi\",\"type\":\"string\"},{\"name\":\"intro\",\"type\":\"string\"},{\"name\":\"Alias\",\"type\":\"string[]\"},{\"name\":\"role\",\"type\":\"string[]\"},{\"name\":\"thumbHash\",\"type\":\"string\"},{\"name\":\"posterHash\",\"type\":\"string\"},{\"name\":\"uncensored\",\"type\":\"bool\"}],\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"date\",\"type\":\"string\"}],\"name\":\"getNewVideos\",\"outputs\":[{\"components\":[{\"name\":\"findNo\",\"type\":\"string\"},{\"name\":\"bangumi\",\"type\":\"string\"},{\"name\":\"intro\",\"type\":\"string\"},{\"name\":\"Alias\",\"type\":\"string[]\"},{\"name\":\"role\",\"type\":\"string[]\"},{\"name\":\"thumbHash\",\"type\":\"string\"},{\"name\":\"posterHash\",\"type\":\"string\"},{\"name\":\"uncensored\",\"type\":\"bool\"}],\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVideoCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVideos\",\"outputs\":[{\"name\":\"\",\"type\":\"string[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"hotVideos\",\"outputs\":[{\"name\":\"findNo\",\"type\":\"string\"},{\"name\":\"bangumi\",\"type\":\"string\"},{\"name\":\"intro\",\"type\":\"string\"},{\"name\":\"thumbHash\",\"type\":\"string\"},{\"name\":\"posterHash\",\"type\":\"string\"},{\"name\":\"uncensored\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"bangumi\",\"type\":\"string\"}],\"name\":\"retrieval\",\"outputs\":[{\"components\":[{\"name\":\"findNo\",\"type\":\"string\"},{\"name\":\"bangumi\",\"type\":\"string\"},{\"name\":\"intro\",\"type\":\"string\"},{\"name\":\"Alias\",\"type\":\"string[]\"},{\"name\":\"role\",\"type\":\"string[]\"},{\"name\":\"thumbHash\",\"type\":\"string\"},{\"name\":\"posterHash\",\"type\":\"string\"},{\"name\":\"uncensored\",\"type\":\"bool\"}],\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"bangumi\",\"type\":\"string\"},{\"name\":\"visit\",\"type\":\"uint32\"},{\"name\":\"series\",\"type\":\"string\"},{\"name\":\"tags\",\"type\":\"string[]\"},{\"name\":\"length\",\"type\":\"string\"},{\"name\":\"sharpness\",\"type\":\"string\"},{\"name\":\"format\",\"type\":\"string\"}],\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"name\":\"bangumi\",\"type\":\"string\"},{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"m3u8\",\"type\":\"string\"},{\"name\":\"sourceHash\",\"type\":\"string\"},{\"name\":\"m3u8Hash\",\"type\":\"string\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"videos\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"videosCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VideoInfo is an auto generated Go binding around an Ethereum contract.
type VideoInfo struct {
	VideoInfoCaller     // Read-only binding to the contract
	VideoInfoTransactor // Write-only binding to the contract
	VideoInfoFilterer   // Log filterer for contract events
}

// VideoInfoCaller is an auto generated read-only Go binding around an Ethereum contract.
type VideoInfoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VideoInfoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VideoInfoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VideoInfoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VideoInfoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VideoInfoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VideoInfoSession struct {
	Contract     *VideoInfo        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VideoInfoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VideoInfoCallerSession struct {
	Contract *VideoInfoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// VideoInfoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VideoInfoTransactorSession struct {
	Contract     *VideoInfoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VideoInfoRaw is an auto generated low-level Go binding around an Ethereum contract.
type VideoInfoRaw struct {
	Contract *VideoInfo // Generic contract binding to access the raw methods on
}

// VideoInfoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VideoInfoCallerRaw struct {
	Contract *VideoInfoCaller // Generic read-only contract binding to access the raw methods on
}

// VideoInfoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VideoInfoTransactorRaw struct {
	Contract *VideoInfoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVideoInfo creates a new instance of VideoInfo, bound to a specific deployed contract.
func NewVideoInfo(address common.Address, backend bind.ContractBackend) (*VideoInfo, error) {
	contract, err := bindVideoInfo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VideoInfo{VideoInfoCaller: VideoInfoCaller{contract: contract}, VideoInfoTransactor: VideoInfoTransactor{contract: contract}, VideoInfoFilterer: VideoInfoFilterer{contract: contract}}, nil
}

// NewVideoInfoCaller creates a new read-only instance of VideoInfo, bound to a specific deployed contract.
func NewVideoInfoCaller(address common.Address, caller bind.ContractCaller) (*VideoInfoCaller, error) {
	contract, err := bindVideoInfo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VideoInfoCaller{contract: contract}, nil
}

// NewVideoInfoTransactor creates a new write-only instance of VideoInfo, bound to a specific deployed contract.
func NewVideoInfoTransactor(address common.Address, transactor bind.ContractTransactor) (*VideoInfoTransactor, error) {
	contract, err := bindVideoInfo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VideoInfoTransactor{contract: contract}, nil
}

// NewVideoInfoFilterer creates a new log filterer instance of VideoInfo, bound to a specific deployed contract.
func NewVideoInfoFilterer(address common.Address, filterer bind.ContractFilterer) (*VideoInfoFilterer, error) {
	contract, err := bindVideoInfo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VideoInfoFilterer{contract: contract}, nil
}

// bindVideoInfo binds a generic wrapper to an already deployed contract.
func bindVideoInfo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VideoInfoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VideoInfo *VideoInfoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VideoInfo.Contract.VideoInfoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VideoInfo *VideoInfoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VideoInfo.Contract.VideoInfoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VideoInfo *VideoInfoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VideoInfo.Contract.VideoInfoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VideoInfo *VideoInfoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VideoInfo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VideoInfo *VideoInfoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VideoInfo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VideoInfo *VideoInfoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VideoInfo.Contract.contract.Transact(opts, method, params...)
}

// Struct2 is an auto generated low-level Go binding around an user-defined struct.
type Struct2 struct {
	Bangumi    string
	Key        string
	M3u8       string
	SourceHash string
	M3u8Hash   string
}

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	FindNo     string
	Bangumi    string
	Intro      string
	Alias      []string
	Role       []string
	ThumbHash  string
	PosterHash string
	Uncensored bool
}

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	Bangumi   string
	Visit     uint32
	Series    string
	Tags      []string
	Length    string
	Sharpness string
	Format    string
}

// GetHotVideos is a free data retrieval call binding the contract method 0xf07b6b79.
//
// Solidity: function getHotVideos() constant returns([]Struct0)
func (_VideoInfo *VideoInfoCaller) GetHotVideos(opts *bind.CallOpts) ([]Struct0, error) {
	var (
		ret0 = new([]Struct0)
	)
	out := ret0
	err := _VideoInfo.contract.Call(opts, out, "getHotVideos")
	return *ret0, err
}

// GetHotVideos is a free data retrieval call binding the contract method 0xf07b6b79.
//
// Solidity: function getHotVideos() constant returns([]Struct0)
func (_VideoInfo *VideoInfoSession) GetHotVideos() ([]Struct0, error) {
	return _VideoInfo.Contract.GetHotVideos(&_VideoInfo.CallOpts)
}

// GetHotVideos is a free data retrieval call binding the contract method 0xf07b6b79.
//
// Solidity: function getHotVideos() constant returns([]Struct0)
func (_VideoInfo *VideoInfoCallerSession) GetHotVideos() ([]Struct0, error) {
	return _VideoInfo.Contract.GetHotVideos(&_VideoInfo.CallOpts)
}

// GetNewVideos is a free data retrieval call binding the contract method 0xa20ddc12.
//
// Solidity: function getNewVideos(string date) constant returns([]Struct0)
func (_VideoInfo *VideoInfoCaller) GetNewVideos(opts *bind.CallOpts, date string) ([]Struct0, error) {
	var (
		ret0 = new([]Struct0)
	)
	out := ret0
	err := _VideoInfo.contract.Call(opts, out, "getNewVideos", date)
	return *ret0, err
}

// GetNewVideos is a free data retrieval call binding the contract method 0xa20ddc12.
//
// Solidity: function getNewVideos(string date) constant returns([]Struct0)
func (_VideoInfo *VideoInfoSession) GetNewVideos(date string) ([]Struct0, error) {
	return _VideoInfo.Contract.GetNewVideos(&_VideoInfo.CallOpts, date)
}

// GetNewVideos is a free data retrieval call binding the contract method 0xa20ddc12.
//
// Solidity: function getNewVideos(string date) constant returns([]Struct0)
func (_VideoInfo *VideoInfoCallerSession) GetNewVideos(date string) ([]Struct0, error) {
	return _VideoInfo.Contract.GetNewVideos(&_VideoInfo.CallOpts, date)
}

// GetVideoCount is a free data retrieval call binding the contract method 0xa286c0b9.
//
// Solidity: function getVideoCount() constant returns(uint32)
func (_VideoInfo *VideoInfoCaller) GetVideoCount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _VideoInfo.contract.Call(opts, out, "getVideoCount")
	return *ret0, err
}

// GetVideoCount is a free data retrieval call binding the contract method 0xa286c0b9.
//
// Solidity: function getVideoCount() constant returns(uint32)
func (_VideoInfo *VideoInfoSession) GetVideoCount() (uint32, error) {
	return _VideoInfo.Contract.GetVideoCount(&_VideoInfo.CallOpts)
}

// GetVideoCount is a free data retrieval call binding the contract method 0xa286c0b9.
//
// Solidity: function getVideoCount() constant returns(uint32)
func (_VideoInfo *VideoInfoCallerSession) GetVideoCount() (uint32, error) {
	return _VideoInfo.Contract.GetVideoCount(&_VideoInfo.CallOpts)
}

// GetVideos is a free data retrieval call binding the contract method 0x68f2d165.
//
// Solidity: function getVideos() constant returns(string[])
func (_VideoInfo *VideoInfoCaller) GetVideos(opts *bind.CallOpts) ([]string, error) {
	var (
		ret0 = new([]string)
	)
	out := ret0
	err := _VideoInfo.contract.Call(opts, out, "getVideos")
	return *ret0, err
}

// GetVideos is a free data retrieval call binding the contract method 0x68f2d165.
//
// Solidity: function getVideos() constant returns(string[])
func (_VideoInfo *VideoInfoSession) GetVideos() ([]string, error) {
	return _VideoInfo.Contract.GetVideos(&_VideoInfo.CallOpts)
}

// GetVideos is a free data retrieval call binding the contract method 0x68f2d165.
//
// Solidity: function getVideos() constant returns(string[])
func (_VideoInfo *VideoInfoCallerSession) GetVideos() ([]string, error) {
	return _VideoInfo.Contract.GetVideos(&_VideoInfo.CallOpts)
}

// HotVideos is a free data retrieval call binding the contract method 0xb332a7bc.
//
// Solidity: function hotVideos(uint256 ) constant returns(string findNo, string bangumi, string intro, string thumbHash, string posterHash, bool uncensored)
func (_VideoInfo *VideoInfoCaller) HotVideos(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FindNo     string
	Bangumi    string
	Intro      string
	ThumbHash  string
	PosterHash string
	Uncensored bool
}, error) {
	ret := new(struct {
		FindNo     string
		Bangumi    string
		Intro      string
		ThumbHash  string
		PosterHash string
		Uncensored bool
	})
	out := ret
	err := _VideoInfo.contract.Call(opts, out, "hotVideos", arg0)
	return *ret, err
}

// HotVideos is a free data retrieval call binding the contract method 0xb332a7bc.
//
// Solidity: function hotVideos(uint256 ) constant returns(string findNo, string bangumi, string intro, string thumbHash, string posterHash, bool uncensored)
func (_VideoInfo *VideoInfoSession) HotVideos(arg0 *big.Int) (struct {
	FindNo     string
	Bangumi    string
	Intro      string
	ThumbHash  string
	PosterHash string
	Uncensored bool
}, error) {
	return _VideoInfo.Contract.HotVideos(&_VideoInfo.CallOpts, arg0)
}

// HotVideos is a free data retrieval call binding the contract method 0xb332a7bc.
//
// Solidity: function hotVideos(uint256 ) constant returns(string findNo, string bangumi, string intro, string thumbHash, string posterHash, bool uncensored)
func (_VideoInfo *VideoInfoCallerSession) HotVideos(arg0 *big.Int) (struct {
	FindNo     string
	Bangumi    string
	Intro      string
	ThumbHash  string
	PosterHash string
	Uncensored bool
}, error) {
	return _VideoInfo.Contract.HotVideos(&_VideoInfo.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_VideoInfo *VideoInfoCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VideoInfo.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_VideoInfo *VideoInfoSession) IsOwner() (bool, error) {
	return _VideoInfo.Contract.IsOwner(&_VideoInfo.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_VideoInfo *VideoInfoCallerSession) IsOwner() (bool, error) {
	return _VideoInfo.Contract.IsOwner(&_VideoInfo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VideoInfo *VideoInfoCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VideoInfo.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VideoInfo *VideoInfoSession) Owner() (common.Address, error) {
	return _VideoInfo.Contract.Owner(&_VideoInfo.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VideoInfo *VideoInfoCallerSession) Owner() (common.Address, error) {
	return _VideoInfo.Contract.Owner(&_VideoInfo.CallOpts)
}

// Retrieval is a free data retrieval call binding the contract method 0x4478bacf.
//
// Solidity: function retrieval(string bangumi) constant returns(Struct0, Struct1, Struct2)
func (_VideoInfo *VideoInfoCaller) Retrieval(opts *bind.CallOpts, bangumi string) (Struct0, Struct1, Struct2, error) {
	var (
		ret0 = new(Struct0)
		ret1 = new(Struct1)
		ret2 = new(Struct2)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _VideoInfo.contract.Call(opts, out, "retrieval", bangumi)
	return *ret0, *ret1, *ret2, err
}

// Retrieval is a free data retrieval call binding the contract method 0x4478bacf.
//
// Solidity: function retrieval(string bangumi) constant returns(Struct0, Struct1, Struct2)
func (_VideoInfo *VideoInfoSession) Retrieval(bangumi string) (Struct0, Struct1, Struct2, error) {
	return _VideoInfo.Contract.Retrieval(&_VideoInfo.CallOpts, bangumi)
}

// Retrieval is a free data retrieval call binding the contract method 0x4478bacf.
//
// Solidity: function retrieval(string bangumi) constant returns(Struct0, Struct1, Struct2)
func (_VideoInfo *VideoInfoCallerSession) Retrieval(bangumi string) (Struct0, Struct1, Struct2, error) {
	return _VideoInfo.Contract.Retrieval(&_VideoInfo.CallOpts, bangumi)
}

// Videos is a free data retrieval call binding the contract method 0xe6821bf5.
//
// Solidity: function videos(uint256 ) constant returns(string)
func (_VideoInfo *VideoInfoCaller) Videos(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _VideoInfo.contract.Call(opts, out, "videos", arg0)
	return *ret0, err
}

// Videos is a free data retrieval call binding the contract method 0xe6821bf5.
//
// Solidity: function videos(uint256 ) constant returns(string)
func (_VideoInfo *VideoInfoSession) Videos(arg0 *big.Int) (string, error) {
	return _VideoInfo.Contract.Videos(&_VideoInfo.CallOpts, arg0)
}

// Videos is a free data retrieval call binding the contract method 0xe6821bf5.
//
// Solidity: function videos(uint256 ) constant returns(string)
func (_VideoInfo *VideoInfoCallerSession) Videos(arg0 *big.Int) (string, error) {
	return _VideoInfo.Contract.Videos(&_VideoInfo.CallOpts, arg0)
}

// VideosCount is a free data retrieval call binding the contract method 0xef08523b.
//
// Solidity: function videosCount() constant returns(uint32)
func (_VideoInfo *VideoInfoCaller) VideosCount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _VideoInfo.contract.Call(opts, out, "videosCount")
	return *ret0, err
}

// VideosCount is a free data retrieval call binding the contract method 0xef08523b.
//
// Solidity: function videosCount() constant returns(uint32)
func (_VideoInfo *VideoInfoSession) VideosCount() (uint32, error) {
	return _VideoInfo.Contract.VideosCount(&_VideoInfo.CallOpts)
}

// VideosCount is a free data retrieval call binding the contract method 0xef08523b.
//
// Solidity: function videosCount() constant returns(uint32)
func (_VideoInfo *VideoInfoCallerSession) VideosCount() (uint32, error) {
	return _VideoInfo.Contract.VideosCount(&_VideoInfo.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0x5eec0ee4.
//
// Solidity: function add(Struct0 v, Struct1 d, Struct2 s) returns(bool)
func (_VideoInfo *VideoInfoTransactor) Add(opts *bind.TransactOpts, v Struct0, d Struct1, s Struct2) (*types.Transaction, error) {
	return _VideoInfo.contract.Transact(opts, "add", v, d, s)
}

// Add is a paid mutator transaction binding the contract method 0x5eec0ee4.
//
// Solidity: function add(Struct0 v, Struct1 d, Struct2 s) returns(bool)
func (_VideoInfo *VideoInfoSession) Add(v Struct0, d Struct1, s Struct2) (*types.Transaction, error) {
	return _VideoInfo.Contract.Add(&_VideoInfo.TransactOpts, v, d, s)
}

// Add is a paid mutator transaction binding the contract method 0x5eec0ee4.
//
// Solidity: function add(Struct0 v, Struct1 d, Struct2 s) returns(bool)
func (_VideoInfo *VideoInfoTransactorSession) Add(v Struct0, d Struct1, s Struct2) (*types.Transaction, error) {
	return _VideoInfo.Contract.Add(&_VideoInfo.TransactOpts, v, d, s)
}

// AddHot is a paid mutator transaction binding the contract method 0xab288f38.
//
// Solidity: function addHot(Struct0 v) returns(bool)
func (_VideoInfo *VideoInfoTransactor) AddHot(opts *bind.TransactOpts, v Struct0) (*types.Transaction, error) {
	return _VideoInfo.contract.Transact(opts, "addHot", v)
}

// AddHot is a paid mutator transaction binding the contract method 0xab288f38.
//
// Solidity: function addHot(Struct0 v) returns(bool)
func (_VideoInfo *VideoInfoSession) AddHot(v Struct0) (*types.Transaction, error) {
	return _VideoInfo.Contract.AddHot(&_VideoInfo.TransactOpts, v)
}

// AddHot is a paid mutator transaction binding the contract method 0xab288f38.
//
// Solidity: function addHot(Struct0 v) returns(bool)
func (_VideoInfo *VideoInfoTransactorSession) AddHot(v Struct0) (*types.Transaction, error) {
	return _VideoInfo.Contract.AddHot(&_VideoInfo.TransactOpts, v)
}

// AddNew is a paid mutator transaction binding the contract method 0x86f5798d.
//
// Solidity: function addNew(string date, Struct0 v) returns(bool)
func (_VideoInfo *VideoInfoTransactor) AddNew(opts *bind.TransactOpts, date string, v Struct0) (*types.Transaction, error) {
	return _VideoInfo.contract.Transact(opts, "addNew", date, v)
}

// AddNew is a paid mutator transaction binding the contract method 0x86f5798d.
//
// Solidity: function addNew(string date, Struct0 v) returns(bool)
func (_VideoInfo *VideoInfoSession) AddNew(date string, v Struct0) (*types.Transaction, error) {
	return _VideoInfo.Contract.AddNew(&_VideoInfo.TransactOpts, date, v)
}

// AddNew is a paid mutator transaction binding the contract method 0x86f5798d.
//
// Solidity: function addNew(string date, Struct0 v) returns(bool)
func (_VideoInfo *VideoInfoTransactorSession) AddNew(date string, v Struct0) (*types.Transaction, error) {
	return _VideoInfo.Contract.AddNew(&_VideoInfo.TransactOpts, date, v)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VideoInfo *VideoInfoTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VideoInfo.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VideoInfo *VideoInfoSession) RenounceOwnership() (*types.Transaction, error) {
	return _VideoInfo.Contract.RenounceOwnership(&_VideoInfo.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_VideoInfo *VideoInfoTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _VideoInfo.Contract.RenounceOwnership(&_VideoInfo.TransactOpts)
}

// Replace is a paid mutator transaction binding the contract method 0xf388a31e.
//
// Solidity: function replace(Struct0 v, Struct1 d, Struct2 s) returns(bool)
func (_VideoInfo *VideoInfoTransactor) Replace(opts *bind.TransactOpts, v Struct0, d Struct1, s Struct2) (*types.Transaction, error) {
	return _VideoInfo.contract.Transact(opts, "replace", v, d, s)
}

// Replace is a paid mutator transaction binding the contract method 0xf388a31e.
//
// Solidity: function replace(Struct0 v, Struct1 d, Struct2 s) returns(bool)
func (_VideoInfo *VideoInfoSession) Replace(v Struct0, d Struct1, s Struct2) (*types.Transaction, error) {
	return _VideoInfo.Contract.Replace(&_VideoInfo.TransactOpts, v, d, s)
}

// Replace is a paid mutator transaction binding the contract method 0xf388a31e.
//
// Solidity: function replace(Struct0 v, Struct1 d, Struct2 s) returns(bool)
func (_VideoInfo *VideoInfoTransactorSession) Replace(v Struct0, d Struct1, s Struct2) (*types.Transaction, error) {
	return _VideoInfo.Contract.Replace(&_VideoInfo.TransactOpts, v, d, s)
}

// ReplaceHot is a paid mutator transaction binding the contract method 0x4fbc40e3.
//
// Solidity: function replaceHot(uint32 idx, Struct0 v) returns()
func (_VideoInfo *VideoInfoTransactor) ReplaceHot(opts *bind.TransactOpts, idx uint32, v Struct0) (*types.Transaction, error) {
	return _VideoInfo.contract.Transact(opts, "replaceHot", idx, v)
}

// ReplaceHot is a paid mutator transaction binding the contract method 0x4fbc40e3.
//
// Solidity: function replaceHot(uint32 idx, Struct0 v) returns()
func (_VideoInfo *VideoInfoSession) ReplaceHot(idx uint32, v Struct0) (*types.Transaction, error) {
	return _VideoInfo.Contract.ReplaceHot(&_VideoInfo.TransactOpts, idx, v)
}

// ReplaceHot is a paid mutator transaction binding the contract method 0x4fbc40e3.
//
// Solidity: function replaceHot(uint32 idx, Struct0 v) returns()
func (_VideoInfo *VideoInfoTransactorSession) ReplaceHot(idx uint32, v Struct0) (*types.Transaction, error) {
	return _VideoInfo.Contract.ReplaceHot(&_VideoInfo.TransactOpts, idx, v)
}

// ReplaceNew is a paid mutator transaction binding the contract method 0xb0145537.
//
// Solidity: function replaceNew(string date, uint32 idx, Struct0 v) returns()
func (_VideoInfo *VideoInfoTransactor) ReplaceNew(opts *bind.TransactOpts, date string, idx uint32, v Struct0) (*types.Transaction, error) {
	return _VideoInfo.contract.Transact(opts, "replaceNew", date, idx, v)
}

// ReplaceNew is a paid mutator transaction binding the contract method 0xb0145537.
//
// Solidity: function replaceNew(string date, uint32 idx, Struct0 v) returns()
func (_VideoInfo *VideoInfoSession) ReplaceNew(date string, idx uint32, v Struct0) (*types.Transaction, error) {
	return _VideoInfo.Contract.ReplaceNew(&_VideoInfo.TransactOpts, date, idx, v)
}

// ReplaceNew is a paid mutator transaction binding the contract method 0xb0145537.
//
// Solidity: function replaceNew(string date, uint32 idx, Struct0 v) returns()
func (_VideoInfo *VideoInfoTransactorSession) ReplaceNew(date string, idx uint32, v Struct0) (*types.Transaction, error) {
	return _VideoInfo.Contract.ReplaceNew(&_VideoInfo.TransactOpts, date, idx, v)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VideoInfo *VideoInfoTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _VideoInfo.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VideoInfo *VideoInfoSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VideoInfo.Contract.TransferOwnership(&_VideoInfo.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_VideoInfo *VideoInfoTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _VideoInfo.Contract.TransferOwnership(&_VideoInfo.TransactOpts, newOwner)
}

// VideoInfoOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VideoInfo contract.
type VideoInfoOwnershipTransferredIterator struct {
	Event *VideoInfoOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VideoInfoOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VideoInfoOwnershipTransferred)
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
		it.Event = new(VideoInfoOwnershipTransferred)
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
func (it *VideoInfoOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VideoInfoOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VideoInfoOwnershipTransferred represents a OwnershipTransferred event raised by the VideoInfo contract.
type VideoInfoOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VideoInfo *VideoInfoFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VideoInfoOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VideoInfo.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VideoInfoOwnershipTransferredIterator{contract: _VideoInfo.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_VideoInfo *VideoInfoFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VideoInfoOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _VideoInfo.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VideoInfoOwnershipTransferred)
				if err := _VideoInfo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VideoInfo *VideoInfoFilterer) ParseOwnershipTransferred(log types.Log) (*VideoInfoOwnershipTransferred, error) {
	event := new(VideoInfoOwnershipTransferred)
	if err := _VideoInfo.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
