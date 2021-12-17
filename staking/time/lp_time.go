// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lptime

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LptimeMetaData contains all meta data concerning the Lptime contract.
var LptimeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_lockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"UpdateRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"coin\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getUnsettlementReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getUserUnWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardDistribution\",\"type\":\"address\"}],\"name\":\"setRewardDistribution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"updateRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userAwardFlag\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"usersAward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"haveAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LptimeABI is the input ABI used to generate the binding from.
// Deprecated: Use LptimeMetaData.ABI instead.
var LptimeABI = LptimeMetaData.ABI

// Lptime is an auto generated Go binding around an Ethereum contract.
type Lptime struct {
	LptimeCaller     // Read-only binding to the contract
	LptimeTransactor // Write-only binding to the contract
	LptimeFilterer   // Log filterer for contract events
}

// LptimeCaller is an auto generated read-only Go binding around an Ethereum contract.
type LptimeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LptimeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LptimeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LptimeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LptimeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LptimeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LptimeSession struct {
	Contract     *Lptime           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LptimeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LptimeCallerSession struct {
	Contract *LptimeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LptimeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LptimeTransactorSession struct {
	Contract     *LptimeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LptimeRaw is an auto generated low-level Go binding around an Ethereum contract.
type LptimeRaw struct {
	Contract *Lptime // Generic contract binding to access the raw methods on
}

// LptimeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LptimeCallerRaw struct {
	Contract *LptimeCaller // Generic read-only contract binding to access the raw methods on
}

// LptimeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LptimeTransactorRaw struct {
	Contract *LptimeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLptime creates a new instance of Lptime, bound to a specific deployed contract.
func NewLptime(address common.Address, backend bind.ContractBackend) (*Lptime, error) {
	contract, err := bindLptime(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lptime{LptimeCaller: LptimeCaller{contract: contract}, LptimeTransactor: LptimeTransactor{contract: contract}, LptimeFilterer: LptimeFilterer{contract: contract}}, nil
}

// NewLptimeCaller creates a new read-only instance of Lptime, bound to a specific deployed contract.
func NewLptimeCaller(address common.Address, caller bind.ContractCaller) (*LptimeCaller, error) {
	contract, err := bindLptime(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LptimeCaller{contract: contract}, nil
}

// NewLptimeTransactor creates a new write-only instance of Lptime, bound to a specific deployed contract.
func NewLptimeTransactor(address common.Address, transactor bind.ContractTransactor) (*LptimeTransactor, error) {
	contract, err := bindLptime(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LptimeTransactor{contract: contract}, nil
}

// NewLptimeFilterer creates a new log filterer instance of Lptime, bound to a specific deployed contract.
func NewLptimeFilterer(address common.Address, filterer bind.ContractFilterer) (*LptimeFilterer, error) {
	contract, err := bindLptime(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LptimeFilterer{contract: contract}, nil
}

// bindLptime binds a generic wrapper to an already deployed contract.
func bindLptime(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LptimeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lptime *LptimeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lptime.Contract.LptimeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lptime *LptimeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lptime.Contract.LptimeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lptime *LptimeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lptime.Contract.LptimeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lptime *LptimeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lptime.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lptime *LptimeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lptime.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lptime *LptimeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lptime.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Lptime *LptimeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lptime.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Lptime *LptimeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Lptime.Contract.BalanceOf(&_Lptime.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Lptime *LptimeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Lptime.Contract.BalanceOf(&_Lptime.CallOpts, account)
}

// Coin is a free data retrieval call binding the contract method 0x11df9995.
//
// Solidity: function coin() view returns(address)
func (_Lptime *LptimeCaller) Coin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lptime.contract.Call(opts, &out, "coin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Coin is a free data retrieval call binding the contract method 0x11df9995.
//
// Solidity: function coin() view returns(address)
func (_Lptime *LptimeSession) Coin() (common.Address, error) {
	return _Lptime.Contract.Coin(&_Lptime.CallOpts)
}

// Coin is a free data retrieval call binding the contract method 0x11df9995.
//
// Solidity: function coin() view returns(address)
func (_Lptime *LptimeCallerSession) Coin() (common.Address, error) {
	return _Lptime.Contract.Coin(&_Lptime.CallOpts)
}

// GetUnsettlementReward is a free data retrieval call binding the contract method 0x8b5c432f.
//
// Solidity: function getUnsettlementReward(address account) view returns(uint256)
func (_Lptime *LptimeCaller) GetUnsettlementReward(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lptime.contract.Call(opts, &out, "getUnsettlementReward", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnsettlementReward is a free data retrieval call binding the contract method 0x8b5c432f.
//
// Solidity: function getUnsettlementReward(address account) view returns(uint256)
func (_Lptime *LptimeSession) GetUnsettlementReward(account common.Address) (*big.Int, error) {
	return _Lptime.Contract.GetUnsettlementReward(&_Lptime.CallOpts, account)
}

// GetUnsettlementReward is a free data retrieval call binding the contract method 0x8b5c432f.
//
// Solidity: function getUnsettlementReward(address account) view returns(uint256)
func (_Lptime *LptimeCallerSession) GetUnsettlementReward(account common.Address) (*big.Int, error) {
	return _Lptime.Contract.GetUnsettlementReward(&_Lptime.CallOpts, account)
}

// GetUserUnWithdraw is a free data retrieval call binding the contract method 0x2f7d1f62.
//
// Solidity: function getUserUnWithdraw(address account) view returns(uint256)
func (_Lptime *LptimeCaller) GetUserUnWithdraw(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lptime.contract.Call(opts, &out, "getUserUnWithdraw", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserUnWithdraw is a free data retrieval call binding the contract method 0x2f7d1f62.
//
// Solidity: function getUserUnWithdraw(address account) view returns(uint256)
func (_Lptime *LptimeSession) GetUserUnWithdraw(account common.Address) (*big.Int, error) {
	return _Lptime.Contract.GetUserUnWithdraw(&_Lptime.CallOpts, account)
}

// GetUserUnWithdraw is a free data retrieval call binding the contract method 0x2f7d1f62.
//
// Solidity: function getUserUnWithdraw(address account) view returns(uint256)
func (_Lptime *LptimeCallerSession) GetUserUnWithdraw(account common.Address) (*big.Int, error) {
	return _Lptime.Contract.GetUserUnWithdraw(&_Lptime.CallOpts, account)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Lptime *LptimeCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Lptime.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Lptime *LptimeSession) IsOwner() (bool, error) {
	return _Lptime.Contract.IsOwner(&_Lptime.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Lptime *LptimeCallerSession) IsOwner() (bool, error) {
	return _Lptime.Contract.IsOwner(&_Lptime.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_Lptime *LptimeCaller) LockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lptime.contract.Call(opts, &out, "lockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_Lptime *LptimeSession) LockTime() (*big.Int, error) {
	return _Lptime.Contract.LockTime(&_Lptime.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_Lptime *LptimeCallerSession) LockTime() (*big.Int, error) {
	return _Lptime.Contract.LockTime(&_Lptime.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lptime *LptimeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lptime.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lptime *LptimeSession) Owner() (common.Address, error) {
	return _Lptime.Contract.Owner(&_Lptime.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lptime *LptimeCallerSession) Owner() (common.Address, error) {
	return _Lptime.Contract.Owner(&_Lptime.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Lptime *LptimeCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lptime.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Lptime *LptimeSession) Rate() (*big.Int, error) {
	return _Lptime.Contract.Rate(&_Lptime.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Lptime *LptimeCallerSession) Rate() (*big.Int, error) {
	return _Lptime.Contract.Rate(&_Lptime.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Lptime *LptimeCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lptime.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Lptime *LptimeSession) StartTime() (*big.Int, error) {
	return _Lptime.Contract.StartTime(&_Lptime.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_Lptime *LptimeCallerSession) StartTime() (*big.Int, error) {
	return _Lptime.Contract.StartTime(&_Lptime.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Lptime *LptimeCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lptime.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Lptime *LptimeSession) Token() (common.Address, error) {
	return _Lptime.Contract.Token(&_Lptime.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Lptime *LptimeCallerSession) Token() (common.Address, error) {
	return _Lptime.Contract.Token(&_Lptime.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lptime *LptimeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lptime.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lptime *LptimeSession) TotalSupply() (*big.Int, error) {
	return _Lptime.Contract.TotalSupply(&_Lptime.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Lptime *LptimeCallerSession) TotalSupply() (*big.Int, error) {
	return _Lptime.Contract.TotalSupply(&_Lptime.CallOpts)
}

// UserAwardFlag is a free data retrieval call binding the contract method 0xa617461e.
//
// Solidity: function userAwardFlag(address ) view returns(bool)
func (_Lptime *LptimeCaller) UserAwardFlag(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Lptime.contract.Call(opts, &out, "userAwardFlag", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserAwardFlag is a free data retrieval call binding the contract method 0xa617461e.
//
// Solidity: function userAwardFlag(address ) view returns(bool)
func (_Lptime *LptimeSession) UserAwardFlag(arg0 common.Address) (bool, error) {
	return _Lptime.Contract.UserAwardFlag(&_Lptime.CallOpts, arg0)
}

// UserAwardFlag is a free data retrieval call binding the contract method 0xa617461e.
//
// Solidity: function userAwardFlag(address ) view returns(bool)
func (_Lptime *LptimeCallerSession) UserAwardFlag(arg0 common.Address) (bool, error) {
	return _Lptime.Contract.UserAwardFlag(&_Lptime.CallOpts, arg0)
}

// UsersAward is a free data retrieval call binding the contract method 0x401131b9.
//
// Solidity: function usersAward(address ) view returns(uint256 startTime, uint256 endTime, uint256 haveAmount)
func (_Lptime *LptimeCaller) UsersAward(opts *bind.CallOpts, arg0 common.Address) (struct {
	StartTime  *big.Int
	EndTime    *big.Int
	HaveAmount *big.Int
}, error) {
	var out []interface{}
	err := _Lptime.contract.Call(opts, &out, "usersAward", arg0)

	outstruct := new(struct {
		StartTime  *big.Int
		EndTime    *big.Int
		HaveAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartTime = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.HaveAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UsersAward is a free data retrieval call binding the contract method 0x401131b9.
//
// Solidity: function usersAward(address ) view returns(uint256 startTime, uint256 endTime, uint256 haveAmount)
func (_Lptime *LptimeSession) UsersAward(arg0 common.Address) (struct {
	StartTime  *big.Int
	EndTime    *big.Int
	HaveAmount *big.Int
}, error) {
	return _Lptime.Contract.UsersAward(&_Lptime.CallOpts, arg0)
}

// UsersAward is a free data retrieval call binding the contract method 0x401131b9.
//
// Solidity: function usersAward(address ) view returns(uint256 startTime, uint256 endTime, uint256 haveAmount)
func (_Lptime *LptimeCallerSession) UsersAward(arg0 common.Address) (struct {
	StartTime  *big.Int
	EndTime    *big.Int
	HaveAmount *big.Int
}, error) {
	return _Lptime.Contract.UsersAward(&_Lptime.CallOpts, arg0)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Lptime *LptimeTransactor) Exit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lptime.contract.Transact(opts, "exit")
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Lptime *LptimeSession) Exit() (*types.Transaction, error) {
	return _Lptime.Contract.Exit(&_Lptime.TransactOpts)
}

// Exit is a paid mutator transaction binding the contract method 0xe9fad8ee.
//
// Solidity: function exit() returns()
func (_Lptime *LptimeTransactorSession) Exit() (*types.Transaction, error) {
	return _Lptime.Contract.Exit(&_Lptime.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Lptime *LptimeTransactor) GetReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lptime.contract.Transact(opts, "getReward")
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Lptime *LptimeSession) GetReward() (*types.Transaction, error) {
	return _Lptime.Contract.GetReward(&_Lptime.TransactOpts)
}

// GetReward is a paid mutator transaction binding the contract method 0x3d18b912.
//
// Solidity: function getReward() returns()
func (_Lptime *LptimeTransactorSession) GetReward() (*types.Transaction, error) {
	return _Lptime.Contract.GetReward(&_Lptime.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lptime *LptimeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lptime.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lptime *LptimeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lptime.Contract.RenounceOwnership(&_Lptime.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Lptime *LptimeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Lptime.Contract.RenounceOwnership(&_Lptime.TransactOpts)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0x0d68b761.
//
// Solidity: function setRewardDistribution(address _rewardDistribution) returns()
func (_Lptime *LptimeTransactor) SetRewardDistribution(opts *bind.TransactOpts, _rewardDistribution common.Address) (*types.Transaction, error) {
	return _Lptime.contract.Transact(opts, "setRewardDistribution", _rewardDistribution)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0x0d68b761.
//
// Solidity: function setRewardDistribution(address _rewardDistribution) returns()
func (_Lptime *LptimeSession) SetRewardDistribution(_rewardDistribution common.Address) (*types.Transaction, error) {
	return _Lptime.Contract.SetRewardDistribution(&_Lptime.TransactOpts, _rewardDistribution)
}

// SetRewardDistribution is a paid mutator transaction binding the contract method 0x0d68b761.
//
// Solidity: function setRewardDistribution(address _rewardDistribution) returns()
func (_Lptime *LptimeTransactorSession) SetRewardDistribution(_rewardDistribution common.Address) (*types.Transaction, error) {
	return _Lptime.Contract.SetRewardDistribution(&_Lptime.TransactOpts, _rewardDistribution)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Lptime *LptimeTransactor) Stake(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Lptime.contract.Transact(opts, "stake", amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Lptime *LptimeSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Lptime.Contract.Stake(&_Lptime.TransactOpts, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 amount) returns()
func (_Lptime *LptimeTransactorSession) Stake(amount *big.Int) (*types.Transaction, error) {
	return _Lptime.Contract.Stake(&_Lptime.TransactOpts, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lptime *LptimeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lptime.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lptime *LptimeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lptime.Contract.TransferOwnership(&_Lptime.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Lptime *LptimeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Lptime.Contract.TransferOwnership(&_Lptime.TransactOpts, newOwner)
}

// UpdateRate is a paid mutator transaction binding the contract method 0x69ea1771.
//
// Solidity: function updateRate(uint256 _rate) returns()
func (_Lptime *LptimeTransactor) UpdateRate(opts *bind.TransactOpts, _rate *big.Int) (*types.Transaction, error) {
	return _Lptime.contract.Transact(opts, "updateRate", _rate)
}

// UpdateRate is a paid mutator transaction binding the contract method 0x69ea1771.
//
// Solidity: function updateRate(uint256 _rate) returns()
func (_Lptime *LptimeSession) UpdateRate(_rate *big.Int) (*types.Transaction, error) {
	return _Lptime.Contract.UpdateRate(&_Lptime.TransactOpts, _rate)
}

// UpdateRate is a paid mutator transaction binding the contract method 0x69ea1771.
//
// Solidity: function updateRate(uint256 _rate) returns()
func (_Lptime *LptimeTransactorSession) UpdateRate(_rate *big.Int) (*types.Transaction, error) {
	return _Lptime.Contract.UpdateRate(&_Lptime.TransactOpts, _rate)
}

// WithERC20 is a paid mutator transaction binding the contract method 0x53255858.
//
// Solidity: function withERC20(address tokenAddr, address recipient, uint256 amount) returns()
func (_Lptime *LptimeTransactor) WithERC20(opts *bind.TransactOpts, tokenAddr common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lptime.contract.Transact(opts, "withERC20", tokenAddr, recipient, amount)
}

// WithERC20 is a paid mutator transaction binding the contract method 0x53255858.
//
// Solidity: function withERC20(address tokenAddr, address recipient, uint256 amount) returns()
func (_Lptime *LptimeSession) WithERC20(tokenAddr common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lptime.Contract.WithERC20(&_Lptime.TransactOpts, tokenAddr, recipient, amount)
}

// WithERC20 is a paid mutator transaction binding the contract method 0x53255858.
//
// Solidity: function withERC20(address tokenAddr, address recipient, uint256 amount) returns()
func (_Lptime *LptimeTransactorSession) WithERC20(tokenAddr common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Lptime.Contract.WithERC20(&_Lptime.TransactOpts, tokenAddr, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Lptime *LptimeTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Lptime.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Lptime *LptimeSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Lptime.Contract.Withdraw(&_Lptime.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Lptime *LptimeTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Lptime.Contract.Withdraw(&_Lptime.TransactOpts, amount)
}

// LptimeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Lptime contract.
type LptimeOwnershipTransferredIterator struct {
	Event *LptimeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LptimeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LptimeOwnershipTransferred)
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
		it.Event = new(LptimeOwnershipTransferred)
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
func (it *LptimeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LptimeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LptimeOwnershipTransferred represents a OwnershipTransferred event raised by the Lptime contract.
type LptimeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lptime *LptimeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LptimeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lptime.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LptimeOwnershipTransferredIterator{contract: _Lptime.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Lptime *LptimeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LptimeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Lptime.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LptimeOwnershipTransferred)
				if err := _Lptime.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Lptime *LptimeFilterer) ParseOwnershipTransferred(log types.Log) (*LptimeOwnershipTransferred, error) {
	event := new(LptimeOwnershipTransferred)
	if err := _Lptime.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LptimeStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Lptime contract.
type LptimeStakedIterator struct {
	Event *LptimeStaked // Event containing the contract specifics and raw log

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
func (it *LptimeStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LptimeStaked)
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
		it.Event = new(LptimeStaked)
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
func (it *LptimeStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LptimeStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LptimeStaked represents a Staked event raised by the Lptime contract.
type LptimeStaked struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Lptime *LptimeFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*LptimeStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lptime.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &LptimeStakedIterator{contract: _Lptime.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Lptime *LptimeFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *LptimeStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lptime.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LptimeStaked)
				if err := _Lptime.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed user, uint256 amount)
func (_Lptime *LptimeFilterer) ParseStaked(log types.Log) (*LptimeStaked, error) {
	event := new(LptimeStaked)
	if err := _Lptime.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LptimeUpdateRateIterator is returned from FilterUpdateRate and is used to iterate over the raw logs and unpacked data for UpdateRate events raised by the Lptime contract.
type LptimeUpdateRateIterator struct {
	Event *LptimeUpdateRate // Event containing the contract specifics and raw log

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
func (it *LptimeUpdateRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LptimeUpdateRate)
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
		it.Event = new(LptimeUpdateRate)
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
func (it *LptimeUpdateRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LptimeUpdateRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LptimeUpdateRate represents a UpdateRate event raised by the Lptime contract.
type LptimeUpdateRate struct {
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUpdateRate is a free log retrieval operation binding the contract event 0xc04ab144c3f6fd32b71825eb92c20af3c8b2f2b6680dfaf3645b527f098c245b.
//
// Solidity: event UpdateRate(uint256 reward)
func (_Lptime *LptimeFilterer) FilterUpdateRate(opts *bind.FilterOpts) (*LptimeUpdateRateIterator, error) {

	logs, sub, err := _Lptime.contract.FilterLogs(opts, "UpdateRate")
	if err != nil {
		return nil, err
	}
	return &LptimeUpdateRateIterator{contract: _Lptime.contract, event: "UpdateRate", logs: logs, sub: sub}, nil
}

// WatchUpdateRate is a free log subscription operation binding the contract event 0xc04ab144c3f6fd32b71825eb92c20af3c8b2f2b6680dfaf3645b527f098c245b.
//
// Solidity: event UpdateRate(uint256 reward)
func (_Lptime *LptimeFilterer) WatchUpdateRate(opts *bind.WatchOpts, sink chan<- *LptimeUpdateRate) (event.Subscription, error) {

	logs, sub, err := _Lptime.contract.WatchLogs(opts, "UpdateRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LptimeUpdateRate)
				if err := _Lptime.contract.UnpackLog(event, "UpdateRate", log); err != nil {
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

// ParseUpdateRate is a log parse operation binding the contract event 0xc04ab144c3f6fd32b71825eb92c20af3c8b2f2b6680dfaf3645b527f098c245b.
//
// Solidity: event UpdateRate(uint256 reward)
func (_Lptime *LptimeFilterer) ParseUpdateRate(log types.Log) (*LptimeUpdateRate, error) {
	event := new(LptimeUpdateRate)
	if err := _Lptime.contract.UnpackLog(event, "UpdateRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LptimeWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Lptime contract.
type LptimeWithdrawnIterator struct {
	Event *LptimeWithdrawn // Event containing the contract specifics and raw log

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
func (it *LptimeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LptimeWithdrawn)
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
		it.Event = new(LptimeWithdrawn)
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
func (it *LptimeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LptimeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LptimeWithdrawn represents a Withdrawn event raised by the Lptime contract.
type LptimeWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Lptime *LptimeFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*LptimeWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lptime.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &LptimeWithdrawnIterator{contract: _Lptime.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Lptime *LptimeFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *LptimeWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Lptime.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LptimeWithdrawn)
				if err := _Lptime.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Lptime *LptimeFilterer) ParseWithdrawn(log types.Log) (*LptimeWithdrawn, error) {
	event := new(LptimeWithdrawn)
	if err := _Lptime.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
