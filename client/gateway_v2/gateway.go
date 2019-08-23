// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gateway_v2

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

// MainnetGatewayContractABI is the input ABI used to generate the binding from.
const MainnetGatewayContractABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"vmc\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllowAnyToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loomAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"depositERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"enableGateway\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"getERC20\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"contractAddress\",\"type\":\"address\"},{\"name\":\"_signersIndexes\",\"type\":\"uint256[]\"},{\"name\":\"_v\",\"type\":\"uint8[]\"},{\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"name\":\"_s\",\"type\":\"bytes32[]\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"},{\"name\":\"allow\",\"type\":\"bool\"}],\"name\":\"toggleAllowToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGatewayEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"allow\",\"type\":\"bool\"}],\"name\":\"toggleAllowAnyToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"isTokenAllowed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_vmc\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ETHReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ERC721Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ERC721XReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenTypes\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ERC721XBatchReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TokenWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"loomCoinAddress\",\"type\":\"address\"}],\"name\":\"LoomCoinReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ERC20Received\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"contractAddress\",\"type\":\"address\"},{\"name\":\"_signersIndexes\",\"type\":\"uint256[]\"},{\"name\":\"_v\",\"type\":\"uint8[]\"},{\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"name\":\"_s\",\"type\":\"bytes32[]\"}],\"name\":\"withdrawERC721X\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"contractAddress\",\"type\":\"address\"},{\"name\":\"_signersIndexes\",\"type\":\"uint256[]\"},{\"name\":\"_v\",\"type\":\"uint8[]\"},{\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"name\":\"_s\",\"type\":\"bytes32[]\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"_signersIndexes\",\"type\":\"uint256[]\"},{\"name\":\"_v\",\"type\":\"uint8[]\"},{\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"name\":\"_s\",\"type\":\"bytes32[]\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC721XReceived\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_types\",\"type\":\"uint256[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC721XBatchReceived\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_uid\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getETH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"getERC721\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"getERC721X\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// MainnetGatewayContract is an auto generated Go binding around an Ethereum contract.
type MainnetGatewayContract struct {
	MainnetGatewayContractCaller     // Read-only binding to the contract
	MainnetGatewayContractTransactor // Write-only binding to the contract
	MainnetGatewayContractFilterer   // Log filterer for contract events
}

// MainnetGatewayContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainnetGatewayContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainnetGatewayContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainnetGatewayContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainnetGatewayContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainnetGatewayContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainnetGatewayContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainnetGatewayContractSession struct {
	Contract     *MainnetGatewayContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MainnetGatewayContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainnetGatewayContractCallerSession struct {
	Contract *MainnetGatewayContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// MainnetGatewayContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainnetGatewayContractTransactorSession struct {
	Contract     *MainnetGatewayContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// MainnetGatewayContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainnetGatewayContractRaw struct {
	Contract *MainnetGatewayContract // Generic contract binding to access the raw methods on
}

// MainnetGatewayContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainnetGatewayContractCallerRaw struct {
	Contract *MainnetGatewayContractCaller // Generic read-only contract binding to access the raw methods on
}

// MainnetGatewayContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainnetGatewayContractTransactorRaw struct {
	Contract *MainnetGatewayContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMainnetGatewayContract creates a new instance of MainnetGatewayContract, bound to a specific deployed contract.
func NewMainnetGatewayContract(address common.Address, backend bind.ContractBackend) (*MainnetGatewayContract, error) {
	contract, err := bindMainnetGatewayContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MainnetGatewayContract{MainnetGatewayContractCaller: MainnetGatewayContractCaller{contract: contract}, MainnetGatewayContractTransactor: MainnetGatewayContractTransactor{contract: contract}, MainnetGatewayContractFilterer: MainnetGatewayContractFilterer{contract: contract}}, nil
}

// NewMainnetGatewayContractCaller creates a new read-only instance of MainnetGatewayContract, bound to a specific deployed contract.
func NewMainnetGatewayContractCaller(address common.Address, caller bind.ContractCaller) (*MainnetGatewayContractCaller, error) {
	contract, err := bindMainnetGatewayContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainnetGatewayContractCaller{contract: contract}, nil
}

// NewMainnetGatewayContractTransactor creates a new write-only instance of MainnetGatewayContract, bound to a specific deployed contract.
func NewMainnetGatewayContractTransactor(address common.Address, transactor bind.ContractTransactor) (*MainnetGatewayContractTransactor, error) {
	contract, err := bindMainnetGatewayContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainnetGatewayContractTransactor{contract: contract}, nil
}

// NewMainnetGatewayContractFilterer creates a new log filterer instance of MainnetGatewayContract, bound to a specific deployed contract.
func NewMainnetGatewayContractFilterer(address common.Address, filterer bind.ContractFilterer) (*MainnetGatewayContractFilterer, error) {
	contract, err := bindMainnetGatewayContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainnetGatewayContractFilterer{contract: contract}, nil
}

// bindMainnetGatewayContract binds a generic wrapper to an already deployed contract.
func bindMainnetGatewayContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainnetGatewayContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainnetGatewayContract *MainnetGatewayContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MainnetGatewayContract.Contract.MainnetGatewayContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainnetGatewayContract *MainnetGatewayContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.MainnetGatewayContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainnetGatewayContract *MainnetGatewayContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.MainnetGatewayContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainnetGatewayContract *MainnetGatewayContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MainnetGatewayContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainnetGatewayContract *MainnetGatewayContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainnetGatewayContract *MainnetGatewayContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.contract.Transact(opts, method, params...)
}

// AllowedTokens is a free data retrieval call binding the contract method 0xe744092e.
//
// Solidity: function allowedTokens( address) constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractCaller) AllowedTokens(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MainnetGatewayContract.contract.Call(opts, out, "allowedTokens", arg0)
	return *ret0, err
}

// AllowedTokens is a free data retrieval call binding the contract method 0xe744092e.
//
// Solidity: function allowedTokens( address) constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractSession) AllowedTokens(arg0 common.Address) (bool, error) {
	return _MainnetGatewayContract.Contract.AllowedTokens(&_MainnetGatewayContract.CallOpts, arg0)
}

// AllowedTokens is a free data retrieval call binding the contract method 0xe744092e.
//
// Solidity: function allowedTokens( address) constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractCallerSession) AllowedTokens(arg0 common.Address) (bool, error) {
	return _MainnetGatewayContract.Contract.AllowedTokens(&_MainnetGatewayContract.CallOpts, arg0)
}

// GetAllowAnyToken is a free data retrieval call binding the contract method 0x2fc85c52.
//
// Solidity: function getAllowAnyToken() constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractCaller) GetAllowAnyToken(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MainnetGatewayContract.contract.Call(opts, out, "getAllowAnyToken")
	return *ret0, err
}

// GetAllowAnyToken is a free data retrieval call binding the contract method 0x2fc85c52.
//
// Solidity: function getAllowAnyToken() constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractSession) GetAllowAnyToken() (bool, error) {
	return _MainnetGatewayContract.Contract.GetAllowAnyToken(&_MainnetGatewayContract.CallOpts)
}

// GetAllowAnyToken is a free data retrieval call binding the contract method 0x2fc85c52.
//
// Solidity: function getAllowAnyToken() constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractCallerSession) GetAllowAnyToken() (bool, error) {
	return _MainnetGatewayContract.Contract.GetAllowAnyToken(&_MainnetGatewayContract.CallOpts)
}

// GetERC20 is a free data retrieval call binding the contract method 0x4e0dc557.
//
// Solidity: function getERC20(contractAddress address) constant returns(uint256)
func (_MainnetGatewayContract *MainnetGatewayContractCaller) GetERC20(opts *bind.CallOpts, contractAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetGatewayContract.contract.Call(opts, out, "getERC20", contractAddress)
	return *ret0, err
}

// GetERC20 is a free data retrieval call binding the contract method 0x4e0dc557.
//
// Solidity: function getERC20(contractAddress address) constant returns(uint256)
func (_MainnetGatewayContract *MainnetGatewayContractSession) GetERC20(contractAddress common.Address) (*big.Int, error) {
	return _MainnetGatewayContract.Contract.GetERC20(&_MainnetGatewayContract.CallOpts, contractAddress)
}

// GetERC20 is a free data retrieval call binding the contract method 0x4e0dc557.
//
// Solidity: function getERC20(contractAddress address) constant returns(uint256)
func (_MainnetGatewayContract *MainnetGatewayContractCallerSession) GetERC20(contractAddress common.Address) (*big.Int, error) {
	return _MainnetGatewayContract.Contract.GetERC20(&_MainnetGatewayContract.CallOpts, contractAddress)
}

// GetERC721 is a free data retrieval call binding the contract method 0x4e56ef52.
//
// Solidity: function getERC721(uid uint256, contractAddress address) constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractCaller) GetERC721(opts *bind.CallOpts, uid *big.Int, contractAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MainnetGatewayContract.contract.Call(opts, out, "getERC721", uid, contractAddress)
	return *ret0, err
}

// GetERC721 is a free data retrieval call binding the contract method 0x4e56ef52.
//
// Solidity: function getERC721(uid uint256, contractAddress address) constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractSession) GetERC721(uid *big.Int, contractAddress common.Address) (bool, error) {
	return _MainnetGatewayContract.Contract.GetERC721(&_MainnetGatewayContract.CallOpts, uid, contractAddress)
}

// GetERC721 is a free data retrieval call binding the contract method 0x4e56ef52.
//
// Solidity: function getERC721(uid uint256, contractAddress address) constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractCallerSession) GetERC721(uid *big.Int, contractAddress common.Address) (bool, error) {
	return _MainnetGatewayContract.Contract.GetERC721(&_MainnetGatewayContract.CallOpts, uid, contractAddress)
}

// GetERC721X is a free data retrieval call binding the contract method 0xb4c60342.
//
// Solidity: function getERC721X(tokenId uint256, contractAddress address) constant returns(uint256)
func (_MainnetGatewayContract *MainnetGatewayContractCaller) GetERC721X(opts *bind.CallOpts, tokenId *big.Int, contractAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetGatewayContract.contract.Call(opts, out, "getERC721X", tokenId, contractAddress)
	return *ret0, err
}

// GetERC721X is a free data retrieval call binding the contract method 0xb4c60342.
//
// Solidity: function getERC721X(tokenId uint256, contractAddress address) constant returns(uint256)
func (_MainnetGatewayContract *MainnetGatewayContractSession) GetERC721X(tokenId *big.Int, contractAddress common.Address) (*big.Int, error) {
	return _MainnetGatewayContract.Contract.GetERC721X(&_MainnetGatewayContract.CallOpts, tokenId, contractAddress)
}

// GetERC721X is a free data retrieval call binding the contract method 0xb4c60342.
//
// Solidity: function getERC721X(tokenId uint256, contractAddress address) constant returns(uint256)
func (_MainnetGatewayContract *MainnetGatewayContractCallerSession) GetERC721X(tokenId *big.Int, contractAddress common.Address) (*big.Int, error) {
	return _MainnetGatewayContract.Contract.GetERC721X(&_MainnetGatewayContract.CallOpts, tokenId, contractAddress)
}

// GetETH is a free data retrieval call binding the contract method 0x14f6c3be.
//
// Solidity: function getETH() constant returns(uint256)
func (_MainnetGatewayContract *MainnetGatewayContractCaller) GetETH(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetGatewayContract.contract.Call(opts, out, "getETH")
	return *ret0, err
}

// GetETH is a free data retrieval call binding the contract method 0x14f6c3be.
//
// Solidity: function getETH() constant returns(uint256)
func (_MainnetGatewayContract *MainnetGatewayContractSession) GetETH() (*big.Int, error) {
	return _MainnetGatewayContract.Contract.GetETH(&_MainnetGatewayContract.CallOpts)
}

// GetETH is a free data retrieval call binding the contract method 0x14f6c3be.
//
// Solidity: function getETH() constant returns(uint256)
func (_MainnetGatewayContract *MainnetGatewayContractCallerSession) GetETH() (*big.Int, error) {
	return _MainnetGatewayContract.Contract.GetETH(&_MainnetGatewayContract.CallOpts)
}

// GetGatewayEnabled is a free data retrieval call binding the contract method 0xe32f3751.
//
// Solidity: function getGatewayEnabled() constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractCaller) GetGatewayEnabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MainnetGatewayContract.contract.Call(opts, out, "getGatewayEnabled")
	return *ret0, err
}

// GetGatewayEnabled is a free data retrieval call binding the contract method 0xe32f3751.
//
// Solidity: function getGatewayEnabled() constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractSession) GetGatewayEnabled() (bool, error) {
	return _MainnetGatewayContract.Contract.GetGatewayEnabled(&_MainnetGatewayContract.CallOpts)
}

// GetGatewayEnabled is a free data retrieval call binding the contract method 0xe32f3751.
//
// Solidity: function getGatewayEnabled() constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractCallerSession) GetGatewayEnabled() (bool, error) {
	return _MainnetGatewayContract.Contract.GetGatewayEnabled(&_MainnetGatewayContract.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_MainnetGatewayContract *MainnetGatewayContractCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MainnetGatewayContract.contract.Call(opts, out, "getOwner")
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_MainnetGatewayContract *MainnetGatewayContractSession) GetOwner() (common.Address, error) {
	return _MainnetGatewayContract.Contract.GetOwner(&_MainnetGatewayContract.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() constant returns(address)
func (_MainnetGatewayContract *MainnetGatewayContractCallerSession) GetOwner() (common.Address, error) {
	return _MainnetGatewayContract.Contract.GetOwner(&_MainnetGatewayContract.CallOpts)
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(tokenAddress address) constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractCaller) IsTokenAllowed(opts *bind.CallOpts, tokenAddress common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MainnetGatewayContract.contract.Call(opts, out, "isTokenAllowed", tokenAddress)
	return *ret0, err
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(tokenAddress address) constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractSession) IsTokenAllowed(tokenAddress common.Address) (bool, error) {
	return _MainnetGatewayContract.Contract.IsTokenAllowed(&_MainnetGatewayContract.CallOpts, tokenAddress)
}

// IsTokenAllowed is a free data retrieval call binding the contract method 0xf9eaee0d.
//
// Solidity: function isTokenAllowed(tokenAddress address) constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractCallerSession) IsTokenAllowed(tokenAddress common.Address) (bool, error) {
	return _MainnetGatewayContract.Contract.IsTokenAllowed(&_MainnetGatewayContract.CallOpts, tokenAddress)
}

// LoomAddress is a free data retrieval call binding the contract method 0x37179db8.
//
// Solidity: function loomAddress() constant returns(address)
func (_MainnetGatewayContract *MainnetGatewayContractCaller) LoomAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MainnetGatewayContract.contract.Call(opts, out, "loomAddress")
	return *ret0, err
}

// LoomAddress is a free data retrieval call binding the contract method 0x37179db8.
//
// Solidity: function loomAddress() constant returns(address)
func (_MainnetGatewayContract *MainnetGatewayContractSession) LoomAddress() (common.Address, error) {
	return _MainnetGatewayContract.Contract.LoomAddress(&_MainnetGatewayContract.CallOpts)
}

// LoomAddress is a free data retrieval call binding the contract method 0x37179db8.
//
// Solidity: function loomAddress() constant returns(address)
func (_MainnetGatewayContract *MainnetGatewayContractCallerSession) LoomAddress() (common.Address, error) {
	return _MainnetGatewayContract.Contract.LoomAddress(&_MainnetGatewayContract.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces( address) constant returns(uint256)
func (_MainnetGatewayContract *MainnetGatewayContractCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetGatewayContract.contract.Call(opts, out, "nonces", arg0)
	return *ret0, err
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces( address) constant returns(uint256)
func (_MainnetGatewayContract *MainnetGatewayContractSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _MainnetGatewayContract.Contract.Nonces(&_MainnetGatewayContract.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces( address) constant returns(uint256)
func (_MainnetGatewayContract *MainnetGatewayContractCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _MainnetGatewayContract.Contract.Nonces(&_MainnetGatewayContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MainnetGatewayContract *MainnetGatewayContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MainnetGatewayContract.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MainnetGatewayContract *MainnetGatewayContractSession) Owner() (common.Address, error) {
	return _MainnetGatewayContract.Contract.Owner(&_MainnetGatewayContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MainnetGatewayContract *MainnetGatewayContractCallerSession) Owner() (common.Address, error) {
	return _MainnetGatewayContract.Contract.Owner(&_MainnetGatewayContract.CallOpts)
}

// Vmc is a free data retrieval call binding the contract method 0x20cc8e51.
//
// Solidity: function vmc() constant returns(address)
func (_MainnetGatewayContract *MainnetGatewayContractCaller) Vmc(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MainnetGatewayContract.contract.Call(opts, out, "vmc")
	return *ret0, err
}

// Vmc is a free data retrieval call binding the contract method 0x20cc8e51.
//
// Solidity: function vmc() constant returns(address)
func (_MainnetGatewayContract *MainnetGatewayContractSession) Vmc() (common.Address, error) {
	return _MainnetGatewayContract.Contract.Vmc(&_MainnetGatewayContract.CallOpts)
}

// Vmc is a free data retrieval call binding the contract method 0x20cc8e51.
//
// Solidity: function vmc() constant returns(address)
func (_MainnetGatewayContract *MainnetGatewayContractCallerSession) Vmc() (common.Address, error) {
	return _MainnetGatewayContract.Contract.Vmc(&_MainnetGatewayContract.CallOpts)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x392d661c.
//
// Solidity: function depositERC20(amount uint256, contractAddress address) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) DepositERC20(opts *bind.TransactOpts, amount *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "depositERC20", amount, contractAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x392d661c.
//
// Solidity: function depositERC20(amount uint256, contractAddress address) returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) DepositERC20(amount *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.DepositERC20(&_MainnetGatewayContract.TransactOpts, amount, contractAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x392d661c.
//
// Solidity: function depositERC20(amount uint256, contractAddress address) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) DepositERC20(amount *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.DepositERC20(&_MainnetGatewayContract.TransactOpts, amount, contractAddress)
}

// EnableGateway is a paid mutator transaction binding the contract method 0x41c25c4a.
//
// Solidity: function enableGateway(enable bool) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) EnableGateway(opts *bind.TransactOpts, enable bool) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "enableGateway", enable)
}

// EnableGateway is a paid mutator transaction binding the contract method 0x41c25c4a.
//
// Solidity: function enableGateway(enable bool) returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) EnableGateway(enable bool) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.EnableGateway(&_MainnetGatewayContract.TransactOpts, enable)
}

// EnableGateway is a paid mutator transaction binding the contract method 0x41c25c4a.
//
// Solidity: function enableGateway(enable bool) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) EnableGateway(enable bool) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.EnableGateway(&_MainnetGatewayContract.TransactOpts, enable)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(_operator address, _from address, _uid uint256, _data bytes) returns(bytes4)
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) OnERC721Received(opts *bind.TransactOpts, _operator common.Address, _from common.Address, _uid *big.Int, _data []byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "onERC721Received", _operator, _from, _uid, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(_operator address, _from address, _uid uint256, _data bytes) returns(bytes4)
func (_MainnetGatewayContract *MainnetGatewayContractSession) OnERC721Received(_operator common.Address, _from common.Address, _uid *big.Int, _data []byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.OnERC721Received(&_MainnetGatewayContract.TransactOpts, _operator, _from, _uid, _data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(_operator address, _from address, _uid uint256, _data bytes) returns(bytes4)
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) OnERC721Received(_operator common.Address, _from common.Address, _uid *big.Int, _data []byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.OnERC721Received(&_MainnetGatewayContract.TransactOpts, _operator, _from, _uid, _data)
}

// OnERC721XBatchReceived is a paid mutator transaction binding the contract method 0xb3b0f4c7.
//
// Solidity: function onERC721XBatchReceived(_operator address, _from address, _types uint256[], _amounts uint256[], _data bytes) returns(bytes4)
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) OnERC721XBatchReceived(opts *bind.TransactOpts, _operator common.Address, _from common.Address, _types []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "onERC721XBatchReceived", _operator, _from, _types, _amounts, _data)
}

// OnERC721XBatchReceived is a paid mutator transaction binding the contract method 0xb3b0f4c7.
//
// Solidity: function onERC721XBatchReceived(_operator address, _from address, _types uint256[], _amounts uint256[], _data bytes) returns(bytes4)
func (_MainnetGatewayContract *MainnetGatewayContractSession) OnERC721XBatchReceived(_operator common.Address, _from common.Address, _types []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.OnERC721XBatchReceived(&_MainnetGatewayContract.TransactOpts, _operator, _from, _types, _amounts, _data)
}

// OnERC721XBatchReceived is a paid mutator transaction binding the contract method 0xb3b0f4c7.
//
// Solidity: function onERC721XBatchReceived(_operator address, _from address, _types uint256[], _amounts uint256[], _data bytes) returns(bytes4)
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) OnERC721XBatchReceived(_operator common.Address, _from common.Address, _types []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.OnERC721XBatchReceived(&_MainnetGatewayContract.TransactOpts, _operator, _from, _types, _amounts, _data)
}

// OnERC721XReceived is a paid mutator transaction binding the contract method 0x93ba7daa.
//
// Solidity: function onERC721XReceived(_operator address, _from address, _tokenId uint256, _amount uint256, _data bytes) returns(bytes4)
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) OnERC721XReceived(opts *bind.TransactOpts, _operator common.Address, _from common.Address, _tokenId *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "onERC721XReceived", _operator, _from, _tokenId, _amount, _data)
}

// OnERC721XReceived is a paid mutator transaction binding the contract method 0x93ba7daa.
//
// Solidity: function onERC721XReceived(_operator address, _from address, _tokenId uint256, _amount uint256, _data bytes) returns(bytes4)
func (_MainnetGatewayContract *MainnetGatewayContractSession) OnERC721XReceived(_operator common.Address, _from common.Address, _tokenId *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.OnERC721XReceived(&_MainnetGatewayContract.TransactOpts, _operator, _from, _tokenId, _amount, _data)
}

// OnERC721XReceived is a paid mutator transaction binding the contract method 0x93ba7daa.
//
// Solidity: function onERC721XReceived(_operator address, _from address, _tokenId uint256, _amount uint256, _data bytes) returns(bytes4)
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) OnERC721XReceived(_operator common.Address, _from common.Address, _tokenId *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.OnERC721XReceived(&_MainnetGatewayContract.TransactOpts, _operator, _from, _tokenId, _amount, _data)
}

// ToggleAllowAnyToken is a paid mutator transaction binding the contract method 0xe402fbc8.
//
// Solidity: function toggleAllowAnyToken(allow bool) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) ToggleAllowAnyToken(opts *bind.TransactOpts, allow bool) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "toggleAllowAnyToken", allow)
}

// ToggleAllowAnyToken is a paid mutator transaction binding the contract method 0xe402fbc8.
//
// Solidity: function toggleAllowAnyToken(allow bool) returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) ToggleAllowAnyToken(allow bool) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.ToggleAllowAnyToken(&_MainnetGatewayContract.TransactOpts, allow)
}

// ToggleAllowAnyToken is a paid mutator transaction binding the contract method 0xe402fbc8.
//
// Solidity: function toggleAllowAnyToken(allow bool) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) ToggleAllowAnyToken(allow bool) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.ToggleAllowAnyToken(&_MainnetGatewayContract.TransactOpts, allow)
}

// ToggleAllowToken is a paid mutator transaction binding the contract method 0xb82730ab.
//
// Solidity: function toggleAllowToken(tokenAddress address, allow bool) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) ToggleAllowToken(opts *bind.TransactOpts, tokenAddress common.Address, allow bool) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "toggleAllowToken", tokenAddress, allow)
}

// ToggleAllowToken is a paid mutator transaction binding the contract method 0xb82730ab.
//
// Solidity: function toggleAllowToken(tokenAddress address, allow bool) returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) ToggleAllowToken(tokenAddress common.Address, allow bool) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.ToggleAllowToken(&_MainnetGatewayContract.TransactOpts, tokenAddress, allow)
}

// ToggleAllowToken is a paid mutator transaction binding the contract method 0xb82730ab.
//
// Solidity: function toggleAllowToken(tokenAddress address, allow bool) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) ToggleAllowToken(tokenAddress common.Address, allow bool) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.ToggleAllowToken(&_MainnetGatewayContract.TransactOpts, tokenAddress, allow)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xb0116dc7.
//
// Solidity: function withdrawERC20(amount uint256, contractAddress address, _signersIndexes uint256[], _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) WithdrawERC20(opts *bind.TransactOpts, amount *big.Int, contractAddress common.Address, _signersIndexes []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "withdrawERC20", amount, contractAddress, _signersIndexes, _v, _r, _s)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xb0116dc7.
//
// Solidity: function withdrawERC20(amount uint256, contractAddress address, _signersIndexes uint256[], _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) WithdrawERC20(amount *big.Int, contractAddress common.Address, _signersIndexes []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.WithdrawERC20(&_MainnetGatewayContract.TransactOpts, amount, contractAddress, _signersIndexes, _v, _r, _s)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0xb0116dc7.
//
// Solidity: function withdrawERC20(amount uint256, contractAddress address, _signersIndexes uint256[], _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) WithdrawERC20(amount *big.Int, contractAddress common.Address, _signersIndexes []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.WithdrawERC20(&_MainnetGatewayContract.TransactOpts, amount, contractAddress, _signersIndexes, _v, _r, _s)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x03c0fe3a.
//
// Solidity: function withdrawERC721(uid uint256, contractAddress address, _signersIndexes uint256[], _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) WithdrawERC721(opts *bind.TransactOpts, uid *big.Int, contractAddress common.Address, _signersIndexes []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "withdrawERC721", uid, contractAddress, _signersIndexes, _v, _r, _s)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x03c0fe3a.
//
// Solidity: function withdrawERC721(uid uint256, contractAddress address, _signersIndexes uint256[], _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) WithdrawERC721(uid *big.Int, contractAddress common.Address, _signersIndexes []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.WithdrawERC721(&_MainnetGatewayContract.TransactOpts, uid, contractAddress, _signersIndexes, _v, _r, _s)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0x03c0fe3a.
//
// Solidity: function withdrawERC721(uid uint256, contractAddress address, _signersIndexes uint256[], _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) WithdrawERC721(uid *big.Int, contractAddress common.Address, _signersIndexes []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.WithdrawERC721(&_MainnetGatewayContract.TransactOpts, uid, contractAddress, _signersIndexes, _v, _r, _s)
}

// WithdrawERC721X is a paid mutator transaction binding the contract method 0xced7c92a.
//
// Solidity: function withdrawERC721X(tokenId uint256, amount uint256, contractAddress address, _signersIndexes uint256[], _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) WithdrawERC721X(opts *bind.TransactOpts, tokenId *big.Int, amount *big.Int, contractAddress common.Address, _signersIndexes []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "withdrawERC721X", tokenId, amount, contractAddress, _signersIndexes, _v, _r, _s)
}

// WithdrawERC721X is a paid mutator transaction binding the contract method 0xced7c92a.
//
// Solidity: function withdrawERC721X(tokenId uint256, amount uint256, contractAddress address, _signersIndexes uint256[], _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) WithdrawERC721X(tokenId *big.Int, amount *big.Int, contractAddress common.Address, _signersIndexes []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.WithdrawERC721X(&_MainnetGatewayContract.TransactOpts, tokenId, amount, contractAddress, _signersIndexes, _v, _r, _s)
}

// WithdrawERC721X is a paid mutator transaction binding the contract method 0xced7c92a.
//
// Solidity: function withdrawERC721X(tokenId uint256, amount uint256, contractAddress address, _signersIndexes uint256[], _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) WithdrawERC721X(tokenId *big.Int, amount *big.Int, contractAddress common.Address, _signersIndexes []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.WithdrawERC721X(&_MainnetGatewayContract.TransactOpts, tokenId, amount, contractAddress, _signersIndexes, _v, _r, _s)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xda293eb3.
//
// Solidity: function withdrawETH(amount uint256, _signersIndexes uint256[], _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) WithdrawETH(opts *bind.TransactOpts, amount *big.Int, _signersIndexes []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "withdrawETH", amount, _signersIndexes, _v, _r, _s)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xda293eb3.
//
// Solidity: function withdrawETH(amount uint256, _signersIndexes uint256[], _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) WithdrawETH(amount *big.Int, _signersIndexes []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.WithdrawETH(&_MainnetGatewayContract.TransactOpts, amount, _signersIndexes, _v, _r, _s)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0xda293eb3.
//
// Solidity: function withdrawETH(amount uint256, _signersIndexes uint256[], _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) WithdrawETH(amount *big.Int, _signersIndexes []*big.Int, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.WithdrawETH(&_MainnetGatewayContract.TransactOpts, amount, _signersIndexes, _v, _r, _s)
}

// MainnetGatewayContractERC20ReceivedIterator is returned from FilterERC20Received and is used to iterate over the raw logs and unpacked data for ERC20Received events raised by the MainnetGatewayContract contract.
type MainnetGatewayContractERC20ReceivedIterator struct {
	Event *MainnetGatewayContractERC20Received // Event containing the contract specifics and raw log

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
func (it *MainnetGatewayContractERC20ReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetGatewayContractERC20Received)
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
		it.Event = new(MainnetGatewayContractERC20Received)
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
func (it *MainnetGatewayContractERC20ReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetGatewayContractERC20ReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetGatewayContractERC20Received represents a ERC20Received event raised by the MainnetGatewayContract contract.
type MainnetGatewayContractERC20Received struct {
	From            common.Address
	Amount          *big.Int
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterERC20Received is a free log retrieval operation binding the contract event 0xa13cf347fb36122550e414f6fd1a0c2e490cff76331c4dcc20f760891ecca12a.
//
// Solidity: e ERC20Received(from address, amount uint256, contractAddress address)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) FilterERC20Received(opts *bind.FilterOpts) (*MainnetGatewayContractERC20ReceivedIterator, error) {

	logs, sub, err := _MainnetGatewayContract.contract.FilterLogs(opts, "ERC20Received")
	if err != nil {
		return nil, err
	}
	return &MainnetGatewayContractERC20ReceivedIterator{contract: _MainnetGatewayContract.contract, event: "ERC20Received", logs: logs, sub: sub}, nil
}

// WatchERC20Received is a free log subscription operation binding the contract event 0xa13cf347fb36122550e414f6fd1a0c2e490cff76331c4dcc20f760891ecca12a.
//
// Solidity: e ERC20Received(from address, amount uint256, contractAddress address)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) WatchERC20Received(opts *bind.WatchOpts, sink chan<- *MainnetGatewayContractERC20Received) (event.Subscription, error) {

	logs, sub, err := _MainnetGatewayContract.contract.WatchLogs(opts, "ERC20Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetGatewayContractERC20Received)
				if err := _MainnetGatewayContract.contract.UnpackLog(event, "ERC20Received", log); err != nil {
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

// MainnetGatewayContractERC721ReceivedIterator is returned from FilterERC721Received and is used to iterate over the raw logs and unpacked data for ERC721Received events raised by the MainnetGatewayContract contract.
type MainnetGatewayContractERC721ReceivedIterator struct {
	Event *MainnetGatewayContractERC721Received // Event containing the contract specifics and raw log

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
func (it *MainnetGatewayContractERC721ReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetGatewayContractERC721Received)
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
		it.Event = new(MainnetGatewayContractERC721Received)
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
func (it *MainnetGatewayContractERC721ReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetGatewayContractERC721ReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetGatewayContractERC721Received represents a ERC721Received event raised by the MainnetGatewayContract contract.
type MainnetGatewayContractERC721Received struct {
	Operator        common.Address
	From            common.Address
	TokenId         *big.Int
	ContractAddress common.Address
	Data            []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterERC721Received is a free log retrieval operation binding the contract event 0x691f4eac2b8850491851c72f70a121d76b20836d776658438f5b13dd9f8dbc6e.
//
// Solidity: e ERC721Received(operator address, from address, tokenId uint256, contractAddress address, data bytes)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) FilterERC721Received(opts *bind.FilterOpts) (*MainnetGatewayContractERC721ReceivedIterator, error) {

	logs, sub, err := _MainnetGatewayContract.contract.FilterLogs(opts, "ERC721Received")
	if err != nil {
		return nil, err
	}
	return &MainnetGatewayContractERC721ReceivedIterator{contract: _MainnetGatewayContract.contract, event: "ERC721Received", logs: logs, sub: sub}, nil
}

// WatchERC721Received is a free log subscription operation binding the contract event 0x691f4eac2b8850491851c72f70a121d76b20836d776658438f5b13dd9f8dbc6e.
//
// Solidity: e ERC721Received(operator address, from address, tokenId uint256, contractAddress address, data bytes)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) WatchERC721Received(opts *bind.WatchOpts, sink chan<- *MainnetGatewayContractERC721Received) (event.Subscription, error) {

	logs, sub, err := _MainnetGatewayContract.contract.WatchLogs(opts, "ERC721Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetGatewayContractERC721Received)
				if err := _MainnetGatewayContract.contract.UnpackLog(event, "ERC721Received", log); err != nil {
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

// MainnetGatewayContractERC721XBatchReceivedIterator is returned from FilterERC721XBatchReceived and is used to iterate over the raw logs and unpacked data for ERC721XBatchReceived events raised by the MainnetGatewayContract contract.
type MainnetGatewayContractERC721XBatchReceivedIterator struct {
	Event *MainnetGatewayContractERC721XBatchReceived // Event containing the contract specifics and raw log

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
func (it *MainnetGatewayContractERC721XBatchReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetGatewayContractERC721XBatchReceived)
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
		it.Event = new(MainnetGatewayContractERC721XBatchReceived)
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
func (it *MainnetGatewayContractERC721XBatchReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetGatewayContractERC721XBatchReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetGatewayContractERC721XBatchReceived represents a ERC721XBatchReceived event raised by the MainnetGatewayContract contract.
type MainnetGatewayContractERC721XBatchReceived struct {
	Operator        common.Address
	To              common.Address
	TokenTypes      []*big.Int
	Amounts         []*big.Int
	ContractAddress common.Address
	Data            []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterERC721XBatchReceived is a free log retrieval operation binding the contract event 0x48d67933be7b1e6d77d914145d793b5c9ced38156f34ebab23216e085435ac55.
//
// Solidity: e ERC721XBatchReceived(operator address, to address, tokenTypes uint256[], amounts uint256[], contractAddress address, data bytes)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) FilterERC721XBatchReceived(opts *bind.FilterOpts) (*MainnetGatewayContractERC721XBatchReceivedIterator, error) {

	logs, sub, err := _MainnetGatewayContract.contract.FilterLogs(opts, "ERC721XBatchReceived")
	if err != nil {
		return nil, err
	}
	return &MainnetGatewayContractERC721XBatchReceivedIterator{contract: _MainnetGatewayContract.contract, event: "ERC721XBatchReceived", logs: logs, sub: sub}, nil
}

// WatchERC721XBatchReceived is a free log subscription operation binding the contract event 0x48d67933be7b1e6d77d914145d793b5c9ced38156f34ebab23216e085435ac55.
//
// Solidity: e ERC721XBatchReceived(operator address, to address, tokenTypes uint256[], amounts uint256[], contractAddress address, data bytes)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) WatchERC721XBatchReceived(opts *bind.WatchOpts, sink chan<- *MainnetGatewayContractERC721XBatchReceived) (event.Subscription, error) {

	logs, sub, err := _MainnetGatewayContract.contract.WatchLogs(opts, "ERC721XBatchReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetGatewayContractERC721XBatchReceived)
				if err := _MainnetGatewayContract.contract.UnpackLog(event, "ERC721XBatchReceived", log); err != nil {
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

// MainnetGatewayContractERC721XReceivedIterator is returned from FilterERC721XReceived and is used to iterate over the raw logs and unpacked data for ERC721XReceived events raised by the MainnetGatewayContract contract.
type MainnetGatewayContractERC721XReceivedIterator struct {
	Event *MainnetGatewayContractERC721XReceived // Event containing the contract specifics and raw log

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
func (it *MainnetGatewayContractERC721XReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetGatewayContractERC721XReceived)
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
		it.Event = new(MainnetGatewayContractERC721XReceived)
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
func (it *MainnetGatewayContractERC721XReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetGatewayContractERC721XReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetGatewayContractERC721XReceived represents a ERC721XReceived event raised by the MainnetGatewayContract contract.
type MainnetGatewayContractERC721XReceived struct {
	Operator        common.Address
	From            common.Address
	TokenId         *big.Int
	Amount          *big.Int
	ContractAddress common.Address
	Data            []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterERC721XReceived is a free log retrieval operation binding the contract event 0xc341982fb8843f55f2f7aae4eb89231a4ef94a199f370debe7bc5c07c2de2bab.
//
// Solidity: e ERC721XReceived(operator address, from address, tokenId uint256, amount uint256, contractAddress address, data bytes)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) FilterERC721XReceived(opts *bind.FilterOpts) (*MainnetGatewayContractERC721XReceivedIterator, error) {

	logs, sub, err := _MainnetGatewayContract.contract.FilterLogs(opts, "ERC721XReceived")
	if err != nil {
		return nil, err
	}
	return &MainnetGatewayContractERC721XReceivedIterator{contract: _MainnetGatewayContract.contract, event: "ERC721XReceived", logs: logs, sub: sub}, nil
}

// WatchERC721XReceived is a free log subscription operation binding the contract event 0xc341982fb8843f55f2f7aae4eb89231a4ef94a199f370debe7bc5c07c2de2bab.
//
// Solidity: e ERC721XReceived(operator address, from address, tokenId uint256, amount uint256, contractAddress address, data bytes)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) WatchERC721XReceived(opts *bind.WatchOpts, sink chan<- *MainnetGatewayContractERC721XReceived) (event.Subscription, error) {

	logs, sub, err := _MainnetGatewayContract.contract.WatchLogs(opts, "ERC721XReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetGatewayContractERC721XReceived)
				if err := _MainnetGatewayContract.contract.UnpackLog(event, "ERC721XReceived", log); err != nil {
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

// MainnetGatewayContractETHReceivedIterator is returned from FilterETHReceived and is used to iterate over the raw logs and unpacked data for ETHReceived events raised by the MainnetGatewayContract contract.
type MainnetGatewayContractETHReceivedIterator struct {
	Event *MainnetGatewayContractETHReceived // Event containing the contract specifics and raw log

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
func (it *MainnetGatewayContractETHReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetGatewayContractETHReceived)
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
		it.Event = new(MainnetGatewayContractETHReceived)
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
func (it *MainnetGatewayContractETHReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetGatewayContractETHReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetGatewayContractETHReceived represents a ETHReceived event raised by the MainnetGatewayContract contract.
type MainnetGatewayContractETHReceived struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterETHReceived is a free log retrieval operation binding the contract event 0xbfe611b001dfcd411432f7bf0d79b82b4b2ee81511edac123a3403c357fb972a.
//
// Solidity: e ETHReceived(from address, amount uint256)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) FilterETHReceived(opts *bind.FilterOpts) (*MainnetGatewayContractETHReceivedIterator, error) {

	logs, sub, err := _MainnetGatewayContract.contract.FilterLogs(opts, "ETHReceived")
	if err != nil {
		return nil, err
	}
	return &MainnetGatewayContractETHReceivedIterator{contract: _MainnetGatewayContract.contract, event: "ETHReceived", logs: logs, sub: sub}, nil
}

// WatchETHReceived is a free log subscription operation binding the contract event 0xbfe611b001dfcd411432f7bf0d79b82b4b2ee81511edac123a3403c357fb972a.
//
// Solidity: e ETHReceived(from address, amount uint256)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) WatchETHReceived(opts *bind.WatchOpts, sink chan<- *MainnetGatewayContractETHReceived) (event.Subscription, error) {

	logs, sub, err := _MainnetGatewayContract.contract.WatchLogs(opts, "ETHReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetGatewayContractETHReceived)
				if err := _MainnetGatewayContract.contract.UnpackLog(event, "ETHReceived", log); err != nil {
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

// MainnetGatewayContractLoomCoinReceivedIterator is returned from FilterLoomCoinReceived and is used to iterate over the raw logs and unpacked data for LoomCoinReceived events raised by the MainnetGatewayContract contract.
type MainnetGatewayContractLoomCoinReceivedIterator struct {
	Event *MainnetGatewayContractLoomCoinReceived // Event containing the contract specifics and raw log

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
func (it *MainnetGatewayContractLoomCoinReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetGatewayContractLoomCoinReceived)
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
		it.Event = new(MainnetGatewayContractLoomCoinReceived)
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
func (it *MainnetGatewayContractLoomCoinReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetGatewayContractLoomCoinReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetGatewayContractLoomCoinReceived represents a LoomCoinReceived event raised by the MainnetGatewayContract contract.
type MainnetGatewayContractLoomCoinReceived struct {
	From            common.Address
	Amount          *big.Int
	LoomCoinAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLoomCoinReceived is a free log retrieval operation binding the contract event 0x91557346f7592c9279b67cc52709a00442f0597658ec38a5fe84568c016331d7.
//
// Solidity: e LoomCoinReceived(from indexed address, amount uint256, loomCoinAddress address)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) FilterLoomCoinReceived(opts *bind.FilterOpts, from []common.Address) (*MainnetGatewayContractLoomCoinReceivedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MainnetGatewayContract.contract.FilterLogs(opts, "LoomCoinReceived", fromRule)
	if err != nil {
		return nil, err
	}
	return &MainnetGatewayContractLoomCoinReceivedIterator{contract: _MainnetGatewayContract.contract, event: "LoomCoinReceived", logs: logs, sub: sub}, nil
}

// WatchLoomCoinReceived is a free log subscription operation binding the contract event 0x91557346f7592c9279b67cc52709a00442f0597658ec38a5fe84568c016331d7.
//
// Solidity: e LoomCoinReceived(from indexed address, amount uint256, loomCoinAddress address)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) WatchLoomCoinReceived(opts *bind.WatchOpts, sink chan<- *MainnetGatewayContractLoomCoinReceived, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _MainnetGatewayContract.contract.WatchLogs(opts, "LoomCoinReceived", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetGatewayContractLoomCoinReceived)
				if err := _MainnetGatewayContract.contract.UnpackLog(event, "LoomCoinReceived", log); err != nil {
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

// MainnetGatewayContractTokenWithdrawnIterator is returned from FilterTokenWithdrawn and is used to iterate over the raw logs and unpacked data for TokenWithdrawn events raised by the MainnetGatewayContract contract.
type MainnetGatewayContractTokenWithdrawnIterator struct {
	Event *MainnetGatewayContractTokenWithdrawn // Event containing the contract specifics and raw log

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
func (it *MainnetGatewayContractTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetGatewayContractTokenWithdrawn)
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
		it.Event = new(MainnetGatewayContractTokenWithdrawn)
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
func (it *MainnetGatewayContractTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetGatewayContractTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetGatewayContractTokenWithdrawn represents a TokenWithdrawn event raised by the MainnetGatewayContract contract.
type MainnetGatewayContractTokenWithdrawn struct {
	Owner           common.Address
	Kind            uint8
	ContractAddress common.Address
	Value           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawn is a free log retrieval operation binding the contract event 0x591f2d33d85291e32c4067b5a497caf3ddb5b1830eba9909e66006ec3a0051b4.
//
// Solidity: e TokenWithdrawn(owner indexed address, kind uint8, contractAddress address, value uint256)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) FilterTokenWithdrawn(opts *bind.FilterOpts, owner []common.Address) (*MainnetGatewayContractTokenWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MainnetGatewayContract.contract.FilterLogs(opts, "TokenWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MainnetGatewayContractTokenWithdrawnIterator{contract: _MainnetGatewayContract.contract, event: "TokenWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawn is a free log subscription operation binding the contract event 0x591f2d33d85291e32c4067b5a497caf3ddb5b1830eba9909e66006ec3a0051b4.
//
// Solidity: e TokenWithdrawn(owner indexed address, kind uint8, contractAddress address, value uint256)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) WatchTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *MainnetGatewayContractTokenWithdrawn, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MainnetGatewayContract.contract.WatchLogs(opts, "TokenWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetGatewayContractTokenWithdrawn)
				if err := _MainnetGatewayContract.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
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
