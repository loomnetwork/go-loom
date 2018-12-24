// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package timelock

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TokenTimelockABI is the input ABI used to generate the binding from.
const TokenTimelockABI = "[{\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"beneficiary\",\"type\":\"address\"},{\"name\":\"releaseTime\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"releaseTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"release\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenTimelock is an auto generated Go binding around an Ethereum contract.
type TokenTimelock struct {
	TokenTimelockCaller     // Read-only binding to the contract
	TokenTimelockTransactor // Write-only binding to the contract
	TokenTimelockFilterer   // Log filterer for contract events
}

// TokenTimelockCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenTimelockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTimelockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTimelockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTimelockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenTimelockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTimelockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenTimelockSession struct {
	Contract     *TokenTimelock    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenTimelockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenTimelockCallerSession struct {
	Contract *TokenTimelockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokenTimelockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTimelockTransactorSession struct {
	Contract     *TokenTimelockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokenTimelockRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenTimelockRaw struct {
	Contract *TokenTimelock // Generic contract binding to access the raw methods on
}

// TokenTimelockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenTimelockCallerRaw struct {
	Contract *TokenTimelockCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTimelockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTimelockTransactorRaw struct {
	Contract *TokenTimelockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenTimelock creates a new instance of TokenTimelock, bound to a specific deployed contract.
func NewTokenTimelock(address common.Address, backend bind.ContractBackend) (*TokenTimelock, error) {
	contract, err := bindTokenTimelock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenTimelock{TokenTimelockCaller: TokenTimelockCaller{contract: contract}, TokenTimelockTransactor: TokenTimelockTransactor{contract: contract}, TokenTimelockFilterer: TokenTimelockFilterer{contract: contract}}, nil
}

// NewTokenTimelockCaller creates a new read-only instance of TokenTimelock, bound to a specific deployed contract.
func NewTokenTimelockCaller(address common.Address, caller bind.ContractCaller) (*TokenTimelockCaller, error) {
	contract, err := bindTokenTimelock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTimelockCaller{contract: contract}, nil
}

// NewTokenTimelockTransactor creates a new write-only instance of TokenTimelock, bound to a specific deployed contract.
func NewTokenTimelockTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTimelockTransactor, error) {
	contract, err := bindTokenTimelock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTimelockTransactor{contract: contract}, nil
}

// NewTokenTimelockFilterer creates a new log filterer instance of TokenTimelock, bound to a specific deployed contract.
func NewTokenTimelockFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenTimelockFilterer, error) {
	contract, err := bindTokenTimelock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenTimelockFilterer{contract: contract}, nil
}

// bindTokenTimelock binds a generic wrapper to an already deployed contract.
func bindTokenTimelock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenTimelockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenTimelock *TokenTimelockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenTimelock.Contract.TokenTimelockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenTimelock *TokenTimelockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenTimelock.Contract.TokenTimelockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenTimelock *TokenTimelockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenTimelock.Contract.TokenTimelockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenTimelock *TokenTimelockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenTimelock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenTimelock *TokenTimelockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenTimelock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenTimelock *TokenTimelockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenTimelock.Contract.contract.Transact(opts, method, params...)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_TokenTimelock *TokenTimelockCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenTimelock.contract.Call(opts, out, "beneficiary")
	return *ret0, err
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_TokenTimelock *TokenTimelockSession) Beneficiary() (common.Address, error) {
	return _TokenTimelock.Contract.Beneficiary(&_TokenTimelock.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_TokenTimelock *TokenTimelockCallerSession) Beneficiary() (common.Address, error) {
	return _TokenTimelock.Contract.Beneficiary(&_TokenTimelock.CallOpts)
}

// ReleaseTime is a free data retrieval call binding the contract method 0xb91d4001.
//
// Solidity: function releaseTime() constant returns(uint256)
func (_TokenTimelock *TokenTimelockCaller) ReleaseTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenTimelock.contract.Call(opts, out, "releaseTime")
	return *ret0, err
}

// ReleaseTime is a free data retrieval call binding the contract method 0xb91d4001.
//
// Solidity: function releaseTime() constant returns(uint256)
func (_TokenTimelock *TokenTimelockSession) ReleaseTime() (*big.Int, error) {
	return _TokenTimelock.Contract.ReleaseTime(&_TokenTimelock.CallOpts)
}

// ReleaseTime is a free data retrieval call binding the contract method 0xb91d4001.
//
// Solidity: function releaseTime() constant returns(uint256)
func (_TokenTimelock *TokenTimelockCallerSession) ReleaseTime() (*big.Int, error) {
	return _TokenTimelock.Contract.ReleaseTime(&_TokenTimelock.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TokenTimelock *TokenTimelockCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenTimelock.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TokenTimelock *TokenTimelockSession) Token() (common.Address, error) {
	return _TokenTimelock.Contract.Token(&_TokenTimelock.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TokenTimelock *TokenTimelockCallerSession) Token() (common.Address, error) {
	return _TokenTimelock.Contract.Token(&_TokenTimelock.CallOpts)
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_TokenTimelock *TokenTimelockTransactor) Release(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenTimelock.contract.Transact(opts, "release")
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_TokenTimelock *TokenTimelockSession) Release() (*types.Transaction, error) {
	return _TokenTimelock.Contract.Release(&_TokenTimelock.TransactOpts)
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_TokenTimelock *TokenTimelockTransactorSession) Release() (*types.Transaction, error) {
	return _TokenTimelock.Contract.Release(&_TokenTimelock.TransactOpts)
}
