// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gateway

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

// MainnetGatewayContractABI is the input ABI used to generate the binding from.
const MainnetGatewayContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"toggleToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x15c75f89\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sig\",\"type\":\"bytes\"},{\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"withdrawERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x2cd2e930\"},{\"constant\":true,\"inputs\":[],\"name\":\"loomAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x37179db8\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"depositERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x392d661c\"},{\"constant\":true,\"inputs\":[{\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"getERC20\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x4e0dc557\"},{\"constant\":true,\"inputs\":[],\"name\":\"numValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x5d593f8d\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x715018a6\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"checkValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x797327ae\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x7ecebe00\"},{\"constant\":true,\"inputs\":[],\"name\":\"allowAnyToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x8025d570\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x8da5cb5b\"},{\"constant\":false,\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"},{\"name\":\"_v\",\"type\":\"uint8[]\"},{\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"name\":\"_s\",\"type\":\"bytes32[]\"}],\"name\":\"addValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x90b616c8\"},{\"constant\":false,\"inputs\":[{\"name\":\"_loomAddress\",\"type\":\"address\"},{\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"setLoomAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x94005ac3\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xaffed0e0\"},{\"constant\":false,\"inputs\":[{\"name\":\"_validator\",\"type\":\"address\"},{\"name\":\"_v\",\"type\":\"uint8[]\"},{\"name\":\"_r\",\"type\":\"bytes32[]\"},{\"name\":\"_s\",\"type\":\"bytes32[]\"}],\"name\":\"removeValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xc7e7f6f6\"},{\"constant\":false,\"inputs\":[{\"name\":\"_allow\",\"type\":\"bool\"}],\"name\":\"toggleAllowAnyToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xe402fbc8\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xe744092e\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xf2fde38b\"},{\"inputs\":[{\"name\":\"loomToken\",\"type\":\"address\"},{\"name\":\"_validators\",\"type\":\"address[]\"},{\"name\":\"_threshold_num\",\"type\":\"uint8\"},{\"name\":\"_threshold_denom\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"signature\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ETHReceived\",\"type\":\"event\",\"signature\":\"0xbfe611b001dfcd411432f7bf0d79b82b4b2ee81511edac123a3403c357fb972a\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ERC721Received\",\"type\":\"event\",\"signature\":\"0x691f4eac2b8850491851c72f70a121d76b20836d776658438f5b13dd9f8dbc6e\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ERC721XReceived\",\"type\":\"event\",\"signature\":\"0xc341982fb8843f55f2f7aae4eb89231a4ef94a199f370debe7bc5c07c2de2bab\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenTypes\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ERC721XBatchReceived\",\"type\":\"event\",\"signature\":\"0x48d67933be7b1e6d77d914145d793b5c9ced38156f34ebab23216e085435ac55\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"ERC20Received\",\"type\":\"event\",\"signature\":\"0xa13cf347fb36122550e414f6fd1a0c2e490cff76331c4dcc20f760891ecca12a\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TokenWithdrawn\",\"type\":\"event\",\"signature\":\"0x591f2d33d85291e32c4067b5a497caf3ddb5b1830eba9909e66006ec3a0051b4\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"loomCoinAddress\",\"type\":\"address\"}],\"name\":\"LoomCoinReceived\",\"type\":\"event\",\"signature\":\"0x91557346f7592c9279b67cc52709a00442f0597658ec38a5fe84568c016331d7\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"AddedValidator\",\"type\":\"event\",\"signature\":\"0x8e15bf46bd11add443414ada75aa9592a4af68f3f2ec02ae3d49572f9843c2a8\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"RemovedValidator\",\"type\":\"event\",\"signature\":\"0xb625c55cf7e37b54fcd18bc4edafdf3f4f9acd59a5ec824c77c795dcb2d65070\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\",\"signature\":\"0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\",\"signature\":\"0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sig\",\"type\":\"bytes\"},{\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"withdrawERC721X\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xe246e933\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"sig\",\"type\":\"bytes\"},{\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"withdrawERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xc899a86b\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x3ef32986\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC721XReceived\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x93ba7daa\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_types\",\"type\":\"uint256[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC721XBatchReceived\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xb3b0f4c7\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_uid\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0x150b7a02\"},{\"constant\":true,\"inputs\":[],\"name\":\"getETH\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x14f6c3be\"},{\"constant\":true,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"getERC721\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0x4e56ef52\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\"},{\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"getERC721X\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"signature\":\"0xb4c60342\"}]"

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

// AllowAnyToken is a free data retrieval call binding the contract method 0x8025d570.
//
// Solidity: function allowAnyToken() constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractCaller) AllowAnyToken(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MainnetGatewayContract.contract.Call(opts, out, "allowAnyToken")
	return *ret0, err
}

// AllowAnyToken is a free data retrieval call binding the contract method 0x8025d570.
//
// Solidity: function allowAnyToken() constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractSession) AllowAnyToken() (bool, error) {
	return _MainnetGatewayContract.Contract.AllowAnyToken(&_MainnetGatewayContract.CallOpts)
}

// AllowAnyToken is a free data retrieval call binding the contract method 0x8025d570.
//
// Solidity: function allowAnyToken() constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractCallerSession) AllowAnyToken() (bool, error) {
	return _MainnetGatewayContract.Contract.AllowAnyToken(&_MainnetGatewayContract.CallOpts)
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

// CheckValidator is a free data retrieval call binding the contract method 0x797327ae.
//
// Solidity: function checkValidator(_address address) constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractCaller) CheckValidator(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MainnetGatewayContract.contract.Call(opts, out, "checkValidator", _address)
	return *ret0, err
}

// CheckValidator is a free data retrieval call binding the contract method 0x797327ae.
//
// Solidity: function checkValidator(_address address) constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractSession) CheckValidator(_address common.Address) (bool, error) {
	return _MainnetGatewayContract.Contract.CheckValidator(&_MainnetGatewayContract.CallOpts, _address)
}

// CheckValidator is a free data retrieval call binding the contract method 0x797327ae.
//
// Solidity: function checkValidator(_address address) constant returns(bool)
func (_MainnetGatewayContract *MainnetGatewayContractCallerSession) CheckValidator(_address common.Address) (bool, error) {
	return _MainnetGatewayContract.Contract.CheckValidator(&_MainnetGatewayContract.CallOpts, _address)
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

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_MainnetGatewayContract *MainnetGatewayContractCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetGatewayContract.contract.Call(opts, out, "nonce")
	return *ret0, err
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_MainnetGatewayContract *MainnetGatewayContractSession) Nonce() (*big.Int, error) {
	return _MainnetGatewayContract.Contract.Nonce(&_MainnetGatewayContract.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() constant returns(uint256)
func (_MainnetGatewayContract *MainnetGatewayContractCallerSession) Nonce() (*big.Int, error) {
	return _MainnetGatewayContract.Contract.Nonce(&_MainnetGatewayContract.CallOpts)
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

// NumValidators is a free data retrieval call binding the contract method 0x5d593f8d.
//
// Solidity: function numValidators() constant returns(uint256)
func (_MainnetGatewayContract *MainnetGatewayContractCaller) NumValidators(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MainnetGatewayContract.contract.Call(opts, out, "numValidators")
	return *ret0, err
}

// NumValidators is a free data retrieval call binding the contract method 0x5d593f8d.
//
// Solidity: function numValidators() constant returns(uint256)
func (_MainnetGatewayContract *MainnetGatewayContractSession) NumValidators() (*big.Int, error) {
	return _MainnetGatewayContract.Contract.NumValidators(&_MainnetGatewayContract.CallOpts)
}

// NumValidators is a free data retrieval call binding the contract method 0x5d593f8d.
//
// Solidity: function numValidators() constant returns(uint256)
func (_MainnetGatewayContract *MainnetGatewayContractCallerSession) NumValidators() (*big.Int, error) {
	return _MainnetGatewayContract.Contract.NumValidators(&_MainnetGatewayContract.CallOpts)
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

// AddValidator is a paid mutator transaction binding the contract method 0x90b616c8.
//
// Solidity: function addValidator(_validator address, _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) AddValidator(opts *bind.TransactOpts, _validator common.Address, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "addValidator", _validator, _v, _r, _s)
}

// AddValidator is a paid mutator transaction binding the contract method 0x90b616c8.
//
// Solidity: function addValidator(_validator address, _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) AddValidator(_validator common.Address, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.AddValidator(&_MainnetGatewayContract.TransactOpts, _validator, _v, _r, _s)
}

// AddValidator is a paid mutator transaction binding the contract method 0x90b616c8.
//
// Solidity: function addValidator(_validator address, _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) AddValidator(_validator common.Address, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.AddValidator(&_MainnetGatewayContract.TransactOpts, _validator, _v, _r, _s)
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

// RemoveValidator is a paid mutator transaction binding the contract method 0xc7e7f6f6.
//
// Solidity: function removeValidator(_validator address, _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) RemoveValidator(opts *bind.TransactOpts, _validator common.Address, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "removeValidator", _validator, _v, _r, _s)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0xc7e7f6f6.
//
// Solidity: function removeValidator(_validator address, _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) RemoveValidator(_validator common.Address, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.RemoveValidator(&_MainnetGatewayContract.TransactOpts, _validator, _v, _r, _s)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0xc7e7f6f6.
//
// Solidity: function removeValidator(_validator address, _v uint8[], _r bytes32[], _s bytes32[]) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) RemoveValidator(_validator common.Address, _v []uint8, _r [][32]byte, _s [][32]byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.RemoveValidator(&_MainnetGatewayContract.TransactOpts, _validator, _v, _r, _s)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.RenounceOwnership(&_MainnetGatewayContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.RenounceOwnership(&_MainnetGatewayContract.TransactOpts)
}

// SetLoomAddress is a paid mutator transaction binding the contract method 0x94005ac3.
//
// Solidity: function setLoomAddress(_loomAddress address, sig bytes) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) SetLoomAddress(opts *bind.TransactOpts, _loomAddress common.Address, sig []byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "setLoomAddress", _loomAddress, sig)
}

// SetLoomAddress is a paid mutator transaction binding the contract method 0x94005ac3.
//
// Solidity: function setLoomAddress(_loomAddress address, sig bytes) returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) SetLoomAddress(_loomAddress common.Address, sig []byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.SetLoomAddress(&_MainnetGatewayContract.TransactOpts, _loomAddress, sig)
}

// SetLoomAddress is a paid mutator transaction binding the contract method 0x94005ac3.
//
// Solidity: function setLoomAddress(_loomAddress address, sig bytes) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) SetLoomAddress(_loomAddress common.Address, sig []byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.SetLoomAddress(&_MainnetGatewayContract.TransactOpts, _loomAddress, sig)
}

// ToggleAllowAnyToken is a paid mutator transaction binding the contract method 0xe402fbc8.
//
// Solidity: function toggleAllowAnyToken(_allow bool) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) ToggleAllowAnyToken(opts *bind.TransactOpts, _allow bool) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "toggleAllowAnyToken", _allow)
}

// ToggleAllowAnyToken is a paid mutator transaction binding the contract method 0xe402fbc8.
//
// Solidity: function toggleAllowAnyToken(_allow bool) returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) ToggleAllowAnyToken(_allow bool) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.ToggleAllowAnyToken(&_MainnetGatewayContract.TransactOpts, _allow)
}

// ToggleAllowAnyToken is a paid mutator transaction binding the contract method 0xe402fbc8.
//
// Solidity: function toggleAllowAnyToken(_allow bool) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) ToggleAllowAnyToken(_allow bool) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.ToggleAllowAnyToken(&_MainnetGatewayContract.TransactOpts, _allow)
}

// ToggleToken is a paid mutator transaction binding the contract method 0x15c75f89.
//
// Solidity: function toggleToken(_token address) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) ToggleToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "toggleToken", _token)
}

// ToggleToken is a paid mutator transaction binding the contract method 0x15c75f89.
//
// Solidity: function toggleToken(_token address) returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) ToggleToken(_token common.Address) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.ToggleToken(&_MainnetGatewayContract.TransactOpts, _token)
}

// ToggleToken is a paid mutator transaction binding the contract method 0x15c75f89.
//
// Solidity: function toggleToken(_token address) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) ToggleToken(_token common.Address) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.ToggleToken(&_MainnetGatewayContract.TransactOpts, _token)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.TransferOwnership(&_MainnetGatewayContract.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.TransferOwnership(&_MainnetGatewayContract.TransactOpts, _newOwner)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x2cd2e930.
//
// Solidity: function withdrawERC20(amount uint256, sig bytes, contractAddress address) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) WithdrawERC20(opts *bind.TransactOpts, amount *big.Int, sig []byte, contractAddress common.Address) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "withdrawERC20", amount, sig, contractAddress)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x2cd2e930.
//
// Solidity: function withdrawERC20(amount uint256, sig bytes, contractAddress address) returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) WithdrawERC20(amount *big.Int, sig []byte, contractAddress common.Address) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.WithdrawERC20(&_MainnetGatewayContract.TransactOpts, amount, sig, contractAddress)
}

// WithdrawERC20 is a paid mutator transaction binding the contract method 0x2cd2e930.
//
// Solidity: function withdrawERC20(amount uint256, sig bytes, contractAddress address) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) WithdrawERC20(amount *big.Int, sig []byte, contractAddress common.Address) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.WithdrawERC20(&_MainnetGatewayContract.TransactOpts, amount, sig, contractAddress)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xc899a86b.
//
// Solidity: function withdrawERC721(uid uint256, sig bytes, contractAddress address) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) WithdrawERC721(opts *bind.TransactOpts, uid *big.Int, sig []byte, contractAddress common.Address) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "withdrawERC721", uid, sig, contractAddress)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xc899a86b.
//
// Solidity: function withdrawERC721(uid uint256, sig bytes, contractAddress address) returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) WithdrawERC721(uid *big.Int, sig []byte, contractAddress common.Address) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.WithdrawERC721(&_MainnetGatewayContract.TransactOpts, uid, sig, contractAddress)
}

// WithdrawERC721 is a paid mutator transaction binding the contract method 0xc899a86b.
//
// Solidity: function withdrawERC721(uid uint256, sig bytes, contractAddress address) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) WithdrawERC721(uid *big.Int, sig []byte, contractAddress common.Address) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.WithdrawERC721(&_MainnetGatewayContract.TransactOpts, uid, sig, contractAddress)
}

// WithdrawERC721X is a paid mutator transaction binding the contract method 0xe246e933.
//
// Solidity: function withdrawERC721X(tokenId uint256, amount uint256, sig bytes, contractAddress address) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) WithdrawERC721X(opts *bind.TransactOpts, tokenId *big.Int, amount *big.Int, sig []byte, contractAddress common.Address) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "withdrawERC721X", tokenId, amount, sig, contractAddress)
}

// WithdrawERC721X is a paid mutator transaction binding the contract method 0xe246e933.
//
// Solidity: function withdrawERC721X(tokenId uint256, amount uint256, sig bytes, contractAddress address) returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) WithdrawERC721X(tokenId *big.Int, amount *big.Int, sig []byte, contractAddress common.Address) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.WithdrawERC721X(&_MainnetGatewayContract.TransactOpts, tokenId, amount, sig, contractAddress)
}

// WithdrawERC721X is a paid mutator transaction binding the contract method 0xe246e933.
//
// Solidity: function withdrawERC721X(tokenId uint256, amount uint256, sig bytes, contractAddress address) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) WithdrawERC721X(tokenId *big.Int, amount *big.Int, sig []byte, contractAddress common.Address) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.WithdrawERC721X(&_MainnetGatewayContract.TransactOpts, tokenId, amount, sig, contractAddress)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x3ef32986.
//
// Solidity: function withdrawETH(amount uint256, sig bytes) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactor) WithdrawETH(opts *bind.TransactOpts, amount *big.Int, sig []byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.contract.Transact(opts, "withdrawETH", amount, sig)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x3ef32986.
//
// Solidity: function withdrawETH(amount uint256, sig bytes) returns()
func (_MainnetGatewayContract *MainnetGatewayContractSession) WithdrawETH(amount *big.Int, sig []byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.WithdrawETH(&_MainnetGatewayContract.TransactOpts, amount, sig)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x3ef32986.
//
// Solidity: function withdrawETH(amount uint256, sig bytes) returns()
func (_MainnetGatewayContract *MainnetGatewayContractTransactorSession) WithdrawETH(amount *big.Int, sig []byte) (*types.Transaction, error) {
	return _MainnetGatewayContract.Contract.WithdrawETH(&_MainnetGatewayContract.TransactOpts, amount, sig)
}

// MainnetGatewayContractAddedValidatorIterator is returned from FilterAddedValidator and is used to iterate over the raw logs and unpacked data for AddedValidator events raised by the MainnetGatewayContract contract.
type MainnetGatewayContractAddedValidatorIterator struct {
	Event *MainnetGatewayContractAddedValidator // Event containing the contract specifics and raw log

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
func (it *MainnetGatewayContractAddedValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetGatewayContractAddedValidator)
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
		it.Event = new(MainnetGatewayContractAddedValidator)
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
func (it *MainnetGatewayContractAddedValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetGatewayContractAddedValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetGatewayContractAddedValidator represents a AddedValidator event raised by the MainnetGatewayContract contract.
type MainnetGatewayContractAddedValidator struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddedValidator is a free log retrieval operation binding the contract event 0x8e15bf46bd11add443414ada75aa9592a4af68f3f2ec02ae3d49572f9843c2a8.
//
// Solidity: e AddedValidator(validator address)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) FilterAddedValidator(opts *bind.FilterOpts) (*MainnetGatewayContractAddedValidatorIterator, error) {

	logs, sub, err := _MainnetGatewayContract.contract.FilterLogs(opts, "AddedValidator")
	if err != nil {
		return nil, err
	}
	return &MainnetGatewayContractAddedValidatorIterator{contract: _MainnetGatewayContract.contract, event: "AddedValidator", logs: logs, sub: sub}, nil
}

// WatchAddedValidator is a free log subscription operation binding the contract event 0x8e15bf46bd11add443414ada75aa9592a4af68f3f2ec02ae3d49572f9843c2a8.
//
// Solidity: e AddedValidator(validator address)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) WatchAddedValidator(opts *bind.WatchOpts, sink chan<- *MainnetGatewayContractAddedValidator) (event.Subscription, error) {

	logs, sub, err := _MainnetGatewayContract.contract.WatchLogs(opts, "AddedValidator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetGatewayContractAddedValidator)
				if err := _MainnetGatewayContract.contract.UnpackLog(event, "AddedValidator", log); err != nil {
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

// MainnetGatewayContractOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the MainnetGatewayContract contract.
type MainnetGatewayContractOwnershipRenouncedIterator struct {
	Event *MainnetGatewayContractOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *MainnetGatewayContractOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetGatewayContractOwnershipRenounced)
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
		it.Event = new(MainnetGatewayContractOwnershipRenounced)
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
func (it *MainnetGatewayContractOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetGatewayContractOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetGatewayContractOwnershipRenounced represents a OwnershipRenounced event raised by the MainnetGatewayContract contract.
type MainnetGatewayContractOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*MainnetGatewayContractOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _MainnetGatewayContract.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MainnetGatewayContractOwnershipRenouncedIterator{contract: _MainnetGatewayContract.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *MainnetGatewayContractOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _MainnetGatewayContract.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetGatewayContractOwnershipRenounced)
				if err := _MainnetGatewayContract.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// MainnetGatewayContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MainnetGatewayContract contract.
type MainnetGatewayContractOwnershipTransferredIterator struct {
	Event *MainnetGatewayContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MainnetGatewayContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetGatewayContractOwnershipTransferred)
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
		it.Event = new(MainnetGatewayContractOwnershipTransferred)
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
func (it *MainnetGatewayContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetGatewayContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetGatewayContractOwnershipTransferred represents a OwnershipTransferred event raised by the MainnetGatewayContract contract.
type MainnetGatewayContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MainnetGatewayContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MainnetGatewayContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MainnetGatewayContractOwnershipTransferredIterator{contract: _MainnetGatewayContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MainnetGatewayContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MainnetGatewayContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetGatewayContractOwnershipTransferred)
				if err := _MainnetGatewayContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// MainnetGatewayContractRemovedValidatorIterator is returned from FilterRemovedValidator and is used to iterate over the raw logs and unpacked data for RemovedValidator events raised by the MainnetGatewayContract contract.
type MainnetGatewayContractRemovedValidatorIterator struct {
	Event *MainnetGatewayContractRemovedValidator // Event containing the contract specifics and raw log

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
func (it *MainnetGatewayContractRemovedValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnetGatewayContractRemovedValidator)
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
		it.Event = new(MainnetGatewayContractRemovedValidator)
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
func (it *MainnetGatewayContractRemovedValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnetGatewayContractRemovedValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnetGatewayContractRemovedValidator represents a RemovedValidator event raised by the MainnetGatewayContract contract.
type MainnetGatewayContractRemovedValidator struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemovedValidator is a free log retrieval operation binding the contract event 0xb625c55cf7e37b54fcd18bc4edafdf3f4f9acd59a5ec824c77c795dcb2d65070.
//
// Solidity: e RemovedValidator(validator address)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) FilterRemovedValidator(opts *bind.FilterOpts) (*MainnetGatewayContractRemovedValidatorIterator, error) {

	logs, sub, err := _MainnetGatewayContract.contract.FilterLogs(opts, "RemovedValidator")
	if err != nil {
		return nil, err
	}
	return &MainnetGatewayContractRemovedValidatorIterator{contract: _MainnetGatewayContract.contract, event: "RemovedValidator", logs: logs, sub: sub}, nil
}

// WatchRemovedValidator is a free log subscription operation binding the contract event 0xb625c55cf7e37b54fcd18bc4edafdf3f4f9acd59a5ec824c77c795dcb2d65070.
//
// Solidity: e RemovedValidator(validator address)
func (_MainnetGatewayContract *MainnetGatewayContractFilterer) WatchRemovedValidator(opts *bind.WatchOpts, sink chan<- *MainnetGatewayContractRemovedValidator) (event.Subscription, error) {

	logs, sub, err := _MainnetGatewayContract.contract.WatchLogs(opts, "RemovedValidator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnetGatewayContractRemovedValidator)
				if err := _MainnetGatewayContract.contract.UnpackLog(event, "RemovedValidator", log); err != nil {
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
