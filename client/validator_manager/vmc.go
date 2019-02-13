// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethcontract

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

// ValidatorManagerContractABI is the input ABI used to generate the binding from.
const ValidatorManagerContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validators\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loomAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"powers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"threshold_denom\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"threshold_num\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalPower\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_validators\",\"type\":\"address[]\"},{\"name\":\"_powers\",\"type\":\"uint64[]\"},{\"name\":\"_threshold_num\",\"type\":\"uint8\"},{\"name\":\"_threshold_denom\",\"type\":\"uint8\"},{\"name\":\"_loomAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_validators\",\"type\":\"address[]\"},{\"indexed\":false,\"name\":\"_powers\",\"type\":\"uint64[]\"}],\"name\":\"ValidatorSetChanged\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPowers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newValidators\",\"type\":\"address[]\"},{\"name\":\"_newPowers\",\"type\":\"uint64[]\"},{\"name\":\"_signIndexes\",\"type\":\"uint256[]\"},{\"name\":\"_v\",\"type\":\"uint8[]\"},{\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"name\":\"_s\",\"type\":\"bytes32[]\"}],\"name\":\"rotateValidators\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_message\",\"type\":\"bytes32\"},{\"name\":\"signersIndex\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"signedByValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_message\",\"type\":\"bytes32\"},{\"name\":\"signersIndex\",\"type\":\"uint256[]\"},{\"name\":\"_v\",\"type\":\"uint8[]\"},{\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"name\":\"_s\",\"type\":\"bytes32[]\"}],\"name\":\"checkThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"isTokenAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"allow\",\"type\":\"bool\"},{\"name\":\"validatorIndex\",\"type\":\"uint256\"}],\"name\":\"toggleAllowAnyToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"allow\",\"type\":\"bool\"},{\"name\":\"validatorIndex\",\"type\":\"uint256\"}],\"name\":\"toggleAllowToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ValidatorManagerContract is an auto generated Go binding around an Ethereum contract.
type ValidatorManagerContract struct {
	ValidatorManagerContractCaller     // Read-only binding to the contract
	ValidatorManagerContractTransactor // Write-only binding to the contract
	ValidatorManagerContractFilterer   // Log filterer for contract events
}

// ValidatorManagerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorManagerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorManagerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorManagerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorManagerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorManagerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorManagerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorManagerContractSession struct {
	Contract     *ValidatorManagerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ValidatorManagerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorManagerContractCallerSession struct {
	Contract *ValidatorManagerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ValidatorManagerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorManagerContractTransactorSession struct {
	Contract     *ValidatorManagerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ValidatorManagerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorManagerContractRaw struct {
	Contract *ValidatorManagerContract // Generic contract binding to access the raw methods on
}

// ValidatorManagerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorManagerContractCallerRaw struct {
	Contract *ValidatorManagerContractCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorManagerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorManagerContractTransactorRaw struct {
	Contract *ValidatorManagerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorManagerContract creates a new instance of ValidatorManagerContract, bound to a specific deployed contract.
func NewValidatorManagerContract(address common.Address, backend bind.ContractBackend) (*ValidatorManagerContract, error) {
	contract, err := bindValidatorManagerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerContract{ValidatorManagerContractCaller: ValidatorManagerContractCaller{contract: contract}, ValidatorManagerContractTransactor: ValidatorManagerContractTransactor{contract: contract}, ValidatorManagerContractFilterer: ValidatorManagerContractFilterer{contract: contract}}, nil
}

// NewValidatorManagerContractCaller creates a new read-only instance of ValidatorManagerContract, bound to a specific deployed contract.
func NewValidatorManagerContractCaller(address common.Address, caller bind.ContractCaller) (*ValidatorManagerContractCaller, error) {
	contract, err := bindValidatorManagerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerContractCaller{contract: contract}, nil
}

// NewValidatorManagerContractTransactor creates a new write-only instance of ValidatorManagerContract, bound to a specific deployed contract.
func NewValidatorManagerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorManagerContractTransactor, error) {
	contract, err := bindValidatorManagerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerContractTransactor{contract: contract}, nil
}

// NewValidatorManagerContractFilterer creates a new log filterer instance of ValidatorManagerContract, bound to a specific deployed contract.
func NewValidatorManagerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorManagerContractFilterer, error) {
	contract, err := bindValidatorManagerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerContractFilterer{contract: contract}, nil
}

// bindValidatorManagerContract binds a generic wrapper to an already deployed contract.
func bindValidatorManagerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorManagerContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorManagerContract *ValidatorManagerContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidatorManagerContract.Contract.ValidatorManagerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorManagerContract *ValidatorManagerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManagerContract.Contract.ValidatorManagerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorManagerContract *ValidatorManagerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorManagerContract.Contract.ValidatorManagerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorManagerContract *ValidatorManagerContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ValidatorManagerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorManagerContract *ValidatorManagerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorManagerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorManagerContract *ValidatorManagerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorManagerContract.Contract.contract.Transact(opts, method, params...)
}

// AllowedTokens is a free data retrieval call binding the contract method 0xe744092e.
//
// Solidity: function allowedTokens( address) constant returns(bool)
func (_ValidatorManagerContract *ValidatorManagerContractCaller) AllowedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ValidatorManagerContract.contract.Call(opts, out, "allowedTokens", arg0)
	return *ret0, err
}

// AllowedTokens is a free data retrieval call binding the contract method 0xe744092e.
//
// Solidity: function allowedTokens( address) constant returns(bool)
func (_ValidatorManagerContract *ValidatorManagerContractSession) AllowedTokens(arg0 common.Address) (bool, error) {
	return _ValidatorManagerContract.Contract.AllowedTokens(&_ValidatorManagerContract.CallOpts, arg0)
}

// AllowedTokens is a free data retrieval call binding the contract method 0xe744092e.
//
// Solidity: function allowedTokens( address) constant returns(bool)
func (_ValidatorManagerContract *ValidatorManagerContractCallerSession) AllowedTokens(arg0 common.Address) (bool, error) {
	return _ValidatorManagerContract.Contract.AllowedTokens(&_ValidatorManagerContract.CallOpts, arg0)
}

// CheckThreshold is a free data retrieval call binding the contract method 0x0fba29c3.
//
// Solidity: function checkThreshold(_message bytes32, signersIndex uint256[], _v uint8[], _r bytes32[], _s bytes32[]) constant returns()
func (_ValidatorManagerContract *ValidatorManagerContractCaller) CheckThreshold(opts *bind.CallOpts, _message [32]byte, signersIndex []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) error {
	var ()
	out := &[]interface{}{}
	err := _ValidatorManagerContract.contract.Call(opts, out, "checkThreshold", _message, signersIndex, _v, _r, _s)
	return err
}

// CheckThreshold is a free data retrieval call binding the contract method 0x0fba29c3.
//
// Solidity: function checkThreshold(_message bytes32, signersIndex uint256[], _v uint8[], _r bytes32[], _s bytes32[]) constant returns()
func (_ValidatorManagerContract *ValidatorManagerContractSession) CheckThreshold(_message [32]byte, signersIndex []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) error {
	return _ValidatorManagerContract.Contract.CheckThreshold(&_ValidatorManagerContract.CallOpts, _message, signersIndex, _v, _r, _s)
}

// CheckThreshold is a free data retrieval call binding the contract method 0x0fba29c3.
//
// Solidity: function checkThreshold(_message bytes32, signersIndex uint256[], _v uint8[], _r bytes32[], _s bytes32[]) constant returns()
func (_ValidatorManagerContract *ValidatorManagerContractCallerSession) CheckThreshold(_message [32]byte, signersIndex []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) error {
	return _ValidatorManagerContract.Contract.CheckThreshold(&_ValidatorManagerContract.CallOpts, _message, signersIndex, _v, _r, _s)
}

// GetPowers is a free data retrieval call binding the contract method 0xff13a1ac.
//
// Solidity: function getPowers() constant returns(uint64[])
func (_ValidatorManagerContract *ValidatorManagerContractCaller) GetPowers(opts *bind.CallOpts) ([]uint64, error) {
	var (
		ret0 = new([]uint64)
	)
	out := ret0
	err := _ValidatorManagerContract.contract.Call(opts, out, "getPowers")
	return *ret0, err
}

// GetPowers is a free data retrieval call binding the contract method 0xff13a1ac.
//
// Solidity: function getPowers() constant returns(uint64[])
func (_ValidatorManagerContract *ValidatorManagerContractSession) GetPowers() ([]uint64, error) {
	return _ValidatorManagerContract.Contract.GetPowers(&_ValidatorManagerContract.CallOpts)
}

// GetPowers is a free data retrieval call binding the contract method 0xff13a1ac.
//
// Solidity: function getPowers() constant returns(uint64[])
func (_ValidatorManagerContract *ValidatorManagerContractCallerSession) GetPowers() ([]uint64, error) {
	return _ValidatorManagerContract.Contract.GetPowers(&_ValidatorManagerContract.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() constant returns(address[])
func (_ValidatorManagerContract *ValidatorManagerContractCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _ValidatorManagerContract.contract.Call(opts, out, "getValidators")
	return *ret0, err
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() constant returns(address[])
func (_ValidatorManagerContract *ValidatorManagerContractSession) GetValidators() ([]common.Address, error) {
	return _ValidatorManagerContract.Contract.GetValidators(&_ValidatorManagerContract.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() constant returns(address[])
func (_ValidatorManagerContract *ValidatorManagerContractCallerSession) GetValidators() ([]common.Address, error) {
	return _ValidatorManagerContract.Contract.GetValidators(&_ValidatorManagerContract.CallOpts)
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(tokenAddress address) constant returns(bool)
func (_ValidatorManagerContract *ValidatorManagerContractCaller) IsTokenAllowed(opts *bind.CallOpts, tokenAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ValidatorManagerContract.contract.Call(opts, out, "isTokenAllowed", tokenAddress)
	return *ret0, err
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(tokenAddress address) constant returns(bool)
func (_ValidatorManagerContract *ValidatorManagerContractSession) IsTokenAllowed(tokenAddress common.Address) (bool, error) {
	return _ValidatorManagerContract.Contract.IsTokenAllowed(&_ValidatorManagerContract.CallOpts, tokenAddress)
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(tokenAddress address) constant returns(bool)
func (_ValidatorManagerContract *ValidatorManagerContractCallerSession) IsTokenAllowed(tokenAddress common.Address) (bool, error) {
	return _ValidatorManagerContract.Contract.IsTokenAllowed(&_ValidatorManagerContract.CallOpts, tokenAddress)
}

// LoomAddress is a free data retrieval call binding the contract method 0x37179db8.
//
// Solidity: function loomAddress() constant returns(address)
func (_ValidatorManagerContract *ValidatorManagerContractCaller) LoomAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ValidatorManagerContract.contract.Call(opts, out, "loomAddress")
	return *ret0, err
}

// LoomAddress is a free data retrieval call binding the contract method 0x37179db8.
//
// Solidity: function loomAddress() constant returns(address)
func (_ValidatorManagerContract *ValidatorManagerContractSession) LoomAddress() (common.Address, error) {
	return _ValidatorManagerContract.Contract.LoomAddress(&_ValidatorManagerContract.CallOpts)
}

// LoomAddress is a free data retrieval call binding the contract method 0x37179db8.
//
// Solidity: function loomAddress() constant returns(address)
func (_ValidatorManagerContract *ValidatorManagerContractCallerSession) LoomAddress() (common.Address, error) {
	return _ValidatorManagerContract.Contract.LoomAddress(&_ValidatorManagerContract.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_ValidatorManagerContract *ValidatorManagerContractCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorManagerContract.contract.Call(opts, out, "nonce")
	return *ret0, err
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_ValidatorManagerContract *ValidatorManagerContractSession) Nonce() (*big.Int, error) {
	return _ValidatorManagerContract.Contract.Nonce(&_ValidatorManagerContract.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_ValidatorManagerContract *ValidatorManagerContractCallerSession) Nonce() (*big.Int, error) {
	return _ValidatorManagerContract.Contract.Nonce(&_ValidatorManagerContract.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces( address) constant returns(uint256)
func (_ValidatorManagerContract *ValidatorManagerContractCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorManagerContract.contract.Call(opts, out, "nonces", arg0)
	return *ret0, err
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces( address) constant returns(uint256)
func (_ValidatorManagerContract *ValidatorManagerContractSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ValidatorManagerContract.Contract.Nonces(&_ValidatorManagerContract.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces( address) constant returns(uint256)
func (_ValidatorManagerContract *ValidatorManagerContractCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _ValidatorManagerContract.Contract.Nonces(&_ValidatorManagerContract.CallOpts, arg0)
}

// Powers is a free data retrieval call binding the contract method 0x40c1bfab.
//
// Solidity: function powers( uint256) constant returns(uint64)
func (_ValidatorManagerContract *ValidatorManagerContractCaller) Powers(opts *bind.CallOpts, arg0 *big.Int) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _ValidatorManagerContract.contract.Call(opts, out, "powers", arg0)
	return *ret0, err
}

// Powers is a free data retrieval call binding the contract method 0x40c1bfab.
//
// Solidity: function powers( uint256) constant returns(uint64)
func (_ValidatorManagerContract *ValidatorManagerContractSession) Powers(arg0 *big.Int) (uint64, error) {
	return _ValidatorManagerContract.Contract.Powers(&_ValidatorManagerContract.CallOpts, arg0)
}

// Powers is a free data retrieval call binding the contract method 0x40c1bfab.
//
// Solidity: function powers( uint256) constant returns(uint64)
func (_ValidatorManagerContract *ValidatorManagerContractCallerSession) Powers(arg0 *big.Int) (uint64, error) {
	return _ValidatorManagerContract.Contract.Powers(&_ValidatorManagerContract.CallOpts, arg0)
}

// SignedByValidator is a free data retrieval call binding the contract method 0xc47c479a.
//
// Solidity: function signedByValidator(_message bytes32, signersIndex uint256, _v uint8, _r bytes32, _s bytes32) constant returns()
func (_ValidatorManagerContract *ValidatorManagerContractCaller) SignedByValidator(opts *bind.CallOpts, _message [32]byte, signersIndex *big.Int, _v uint8, _r [32]byte, _s [32]byte) error {
	var ()
	out := &[]interface{}{}
	err := _ValidatorManagerContract.contract.Call(opts, out, "signedByValidator", _message, signersIndex, _v, _r, _s)
	return err
}

// SignedByValidator is a free data retrieval call binding the contract method 0xc47c479a.
//
// Solidity: function signedByValidator(_message bytes32, signersIndex uint256, _v uint8, _r bytes32, _s bytes32) constant returns()
func (_ValidatorManagerContract *ValidatorManagerContractSession) SignedByValidator(_message [32]byte, signersIndex *big.Int, _v uint8, _r [32]byte, _s [32]byte) error {
	return _ValidatorManagerContract.Contract.SignedByValidator(&_ValidatorManagerContract.CallOpts, _message, signersIndex, _v, _r, _s)
}

// SignedByValidator is a free data retrieval call binding the contract method 0xc47c479a.
//
// Solidity: function signedByValidator(_message bytes32, signersIndex uint256, _v uint8, _r bytes32, _s bytes32) constant returns()
func (_ValidatorManagerContract *ValidatorManagerContractCallerSession) SignedByValidator(_message [32]byte, signersIndex *big.Int, _v uint8, _r [32]byte, _s [32]byte) error {
	return _ValidatorManagerContract.Contract.SignedByValidator(&_ValidatorManagerContract.CallOpts, _message, signersIndex, _v, _r, _s)
}

// ThresholdDenom is a free data retrieval call binding the contract method 0x57d717d1.
//
// Solidity: function threshold_denom() constant returns(uint8)
func (_ValidatorManagerContract *ValidatorManagerContractCaller) ThresholdDenom(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ValidatorManagerContract.contract.Call(opts, out, "threshold_denom")
	return *ret0, err
}

// ThresholdDenom is a free data retrieval call binding the contract method 0x57d717d1.
//
// Solidity: function threshold_denom() constant returns(uint8)
func (_ValidatorManagerContract *ValidatorManagerContractSession) ThresholdDenom() (uint8, error) {
	return _ValidatorManagerContract.Contract.ThresholdDenom(&_ValidatorManagerContract.CallOpts)
}

// ThresholdDenom is a free data retrieval call binding the contract method 0x57d717d1.
//
// Solidity: function threshold_denom() constant returns(uint8)
func (_ValidatorManagerContract *ValidatorManagerContractCallerSession) ThresholdDenom() (uint8, error) {
	return _ValidatorManagerContract.Contract.ThresholdDenom(&_ValidatorManagerContract.CallOpts)
}

// ThresholdNum is a free data retrieval call binding the contract method 0xc57829d2.
//
// Solidity: function threshold_num() constant returns(uint8)
func (_ValidatorManagerContract *ValidatorManagerContractCaller) ThresholdNum(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _ValidatorManagerContract.contract.Call(opts, out, "threshold_num")
	return *ret0, err
}

// ThresholdNum is a free data retrieval call binding the contract method 0xc57829d2.
//
// Solidity: function threshold_num() constant returns(uint8)
func (_ValidatorManagerContract *ValidatorManagerContractSession) ThresholdNum() (uint8, error) {
	return _ValidatorManagerContract.Contract.ThresholdNum(&_ValidatorManagerContract.CallOpts)
}

// ThresholdNum is a free data retrieval call binding the contract method 0xc57829d2.
//
// Solidity: function threshold_num() constant returns(uint8)
func (_ValidatorManagerContract *ValidatorManagerContractCallerSession) ThresholdNum() (uint8, error) {
	return _ValidatorManagerContract.Contract.ThresholdNum(&_ValidatorManagerContract.CallOpts)
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() constant returns(uint256)
func (_ValidatorManagerContract *ValidatorManagerContractCaller) TotalPower(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ValidatorManagerContract.contract.Call(opts, out, "totalPower")
	return *ret0, err
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() constant returns(uint256)
func (_ValidatorManagerContract *ValidatorManagerContractSession) TotalPower() (*big.Int, error) {
	return _ValidatorManagerContract.Contract.TotalPower(&_ValidatorManagerContract.CallOpts)
}

// TotalPower is a free data retrieval call binding the contract method 0xdb3ad22c.
//
// Solidity: function totalPower() constant returns(uint256)
func (_ValidatorManagerContract *ValidatorManagerContractCallerSession) TotalPower() (*big.Int, error) {
	return _ValidatorManagerContract.Contract.TotalPower(&_ValidatorManagerContract.CallOpts)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators( uint256) constant returns(address)
func (_ValidatorManagerContract *ValidatorManagerContractCaller) Validators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ValidatorManagerContract.contract.Call(opts, out, "validators", arg0)
	return *ret0, err
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators( uint256) constant returns(address)
func (_ValidatorManagerContract *ValidatorManagerContractSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _ValidatorManagerContract.Contract.Validators(&_ValidatorManagerContract.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0x35aa2e44.
//
// Solidity: function validators( uint256) constant returns(address)
func (_ValidatorManagerContract *ValidatorManagerContractCallerSession) Validators(arg0 *big.Int) (common.Address, error) {
	return _ValidatorManagerContract.Contract.Validators(&_ValidatorManagerContract.CallOpts, arg0)
}

// RotateValidators is a paid mutator transaction binding the contract method 0xeb2eb0ef.
//
// Solidity: function rotateValidators(_newValidators address[], _newPowers uint64[], _signIndexes uint256[], _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_ValidatorManagerContract *ValidatorManagerContractTransactor) RotateValidators(opts *bind.TransactOpts, _newValidators []common.Address, _newPowers []uint64, _signIndexes []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _ValidatorManagerContract.contract.Transact(opts, "rotateValidators", _newValidators, _newPowers, _signIndexes, _v, _r, _s)
}

// RotateValidators is a paid mutator transaction binding the contract method 0xeb2eb0ef.
//
// Solidity: function rotateValidators(_newValidators address[], _newPowers uint64[], _signIndexes uint256[], _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_ValidatorManagerContract *ValidatorManagerContractSession) RotateValidators(_newValidators []common.Address, _newPowers []uint64, _signIndexes []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _ValidatorManagerContract.Contract.RotateValidators(&_ValidatorManagerContract.TransactOpts, _newValidators, _newPowers, _signIndexes, _v, _r, _s)
}

// RotateValidators is a paid mutator transaction binding the contract method 0xeb2eb0ef.
//
// Solidity: function rotateValidators(_newValidators address[], _newPowers uint64[], _signIndexes uint256[], _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_ValidatorManagerContract *ValidatorManagerContractTransactorSession) RotateValidators(_newValidators []common.Address, _newPowers []uint64, _signIndexes []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _ValidatorManagerContract.Contract.RotateValidators(&_ValidatorManagerContract.TransactOpts, _newValidators, _newPowers, _signIndexes, _v, _r, _s)
}

// ToggleAllowAnyToken is a paid mutator transaction binding the contract method 0x1a6be287.
//
// Solidity: function toggleAllowAnyToken(allow bool, validatorIndex uint256) returns()
func (_ValidatorManagerContract *ValidatorManagerContractTransactor) ToggleAllowAnyToken(opts *bind.TransactOpts, allow bool, validatorIndex *big.Int) (*types.Transaction, error) {
	return _ValidatorManagerContract.contract.Transact(opts, "toggleAllowAnyToken", allow, validatorIndex)
}

// ToggleAllowAnyToken is a paid mutator transaction binding the contract method 0x1a6be287.
//
// Solidity: function toggleAllowAnyToken(allow bool, validatorIndex uint256) returns()
func (_ValidatorManagerContract *ValidatorManagerContractSession) ToggleAllowAnyToken(allow bool, validatorIndex *big.Int) (*types.Transaction, error) {
	return _ValidatorManagerContract.Contract.ToggleAllowAnyToken(&_ValidatorManagerContract.TransactOpts, allow, validatorIndex)
}

// ToggleAllowAnyToken is a paid mutator transaction binding the contract method 0x1a6be287.
//
// Solidity: function toggleAllowAnyToken(allow bool, validatorIndex uint256) returns()
func (_ValidatorManagerContract *ValidatorManagerContractTransactorSession) ToggleAllowAnyToken(allow bool, validatorIndex *big.Int) (*types.Transaction, error) {
	return _ValidatorManagerContract.Contract.ToggleAllowAnyToken(&_ValidatorManagerContract.TransactOpts, allow, validatorIndex)
}

// ToggleAllowToken is a paid mutator transaction binding the contract method 0xe3ece440.
//
// Solidity: function toggleAllowToken(tokenAddress address, allow bool, validatorIndex uint256) returns()
func (_ValidatorManagerContract *ValidatorManagerContractTransactor) ToggleAllowToken(opts *bind.TransactOpts, tokenAddress common.Address, allow bool, validatorIndex *big.Int) (*types.Transaction, error) {
	return _ValidatorManagerContract.contract.Transact(opts, "toggleAllowToken", tokenAddress, allow, validatorIndex)
}

// ToggleAllowToken is a paid mutator transaction binding the contract method 0xe3ece440.
//
// Solidity: function toggleAllowToken(tokenAddress address, allow bool, validatorIndex uint256) returns()
func (_ValidatorManagerContract *ValidatorManagerContractSession) ToggleAllowToken(tokenAddress common.Address, allow bool, validatorIndex *big.Int) (*types.Transaction, error) {
	return _ValidatorManagerContract.Contract.ToggleAllowToken(&_ValidatorManagerContract.TransactOpts, tokenAddress, allow, validatorIndex)
}

// ToggleAllowToken is a paid mutator transaction binding the contract method 0xe3ece440.
//
// Solidity: function toggleAllowToken(tokenAddress address, allow bool, validatorIndex uint256) returns()
func (_ValidatorManagerContract *ValidatorManagerContractTransactorSession) ToggleAllowToken(tokenAddress common.Address, allow bool, validatorIndex *big.Int) (*types.Transaction, error) {
	return _ValidatorManagerContract.Contract.ToggleAllowToken(&_ValidatorManagerContract.TransactOpts, tokenAddress, allow, validatorIndex)
}

// ValidatorManagerContractValidatorSetChangedIterator is returned from FilterValidatorSetChanged and is used to iterate over the raw logs and unpacked data for ValidatorSetChanged events raised by the ValidatorManagerContract contract.
type ValidatorManagerContractValidatorSetChangedIterator struct {
	Event *ValidatorManagerContractValidatorSetChanged // Event containing the contract specifics and raw log

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
func (it *ValidatorManagerContractValidatorSetChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorManagerContractValidatorSetChanged)
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
		it.Event = new(ValidatorManagerContractValidatorSetChanged)
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
func (it *ValidatorManagerContractValidatorSetChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorManagerContractValidatorSetChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorManagerContractValidatorSetChanged represents a ValidatorSetChanged event raised by the ValidatorManagerContract contract.
type ValidatorManagerContractValidatorSetChanged struct {
	Validators []common.Address
	Powers     []uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetChanged is a free log retrieval operation binding the contract event 0x323c51e0ad42c317ff3b00c6ce354d799a4b5eaf3a25cf3169cf2efd339d4d54.
//
// Solidity: e ValidatorSetChanged(_validators address[], _powers uint64[])
func (_ValidatorManagerContract *ValidatorManagerContractFilterer) FilterValidatorSetChanged(opts *bind.FilterOpts) (*ValidatorManagerContractValidatorSetChangedIterator, error) {

	logs, sub, err := _ValidatorManagerContract.contract.FilterLogs(opts, "ValidatorSetChanged")
	if err != nil {
		return nil, err
	}
	return &ValidatorManagerContractValidatorSetChangedIterator{contract: _ValidatorManagerContract.contract, event: "ValidatorSetChanged", logs: logs, sub: sub}, nil
}

// WatchValidatorSetChanged is a free log subscription operation binding the contract event 0x323c51e0ad42c317ff3b00c6ce354d799a4b5eaf3a25cf3169cf2efd339d4d54.
//
// Solidity: e ValidatorSetChanged(_validators address[], _powers uint64[])
func (_ValidatorManagerContract *ValidatorManagerContractFilterer) WatchValidatorSetChanged(opts *bind.WatchOpts, sink chan<- *ValidatorManagerContractValidatorSetChanged) (event.Subscription, error) {

	logs, sub, err := _ValidatorManagerContract.contract.WatchLogs(opts, "ValidatorSetChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorManagerContractValidatorSetChanged)
				if err := _ValidatorManagerContract.contract.UnpackLog(event, "ValidatorSetChanged", log); err != nil {
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
