// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc721x

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

// ERC721XABI is the input ABI used to generate the binding from.
const ERC721XABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementsERC721\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"InterfaceId_ERC165\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOwned\",\"outputs\":[{\"name\":\"indexes\",\"type\":\"uint256[]\"},{\"name\":\"balances\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_operator\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"tokenUri\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"isOperator\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenTypes\",\"type\":\"uint256[]\"},{\"indexed\":false,\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"BatchTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"quantity\",\"type\":\"uint256\"}],\"name\":\"TransferWithQuantity\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementsERC721X\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC721X is an auto generated Go binding around an Ethereum contract.
type ERC721X struct {
	ERC721XCaller     // Read-only binding to the contract
	ERC721XTransactor // Write-only binding to the contract
	ERC721XFilterer   // Log filterer for contract events
}

// ERC721XCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721XCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721XTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721XTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721XFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721XFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721XSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721XSession struct {
	Contract     *ERC721X          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721XCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721XCallerSession struct {
	Contract *ERC721XCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ERC721XTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721XTransactorSession struct {
	Contract     *ERC721XTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC721XRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721XRaw struct {
	Contract *ERC721X // Generic contract binding to access the raw methods on
}

// ERC721XCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721XCallerRaw struct {
	Contract *ERC721XCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721XTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721XTransactorRaw struct {
	Contract *ERC721XTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721X creates a new instance of ERC721X, bound to a specific deployed contract.
func NewERC721X(address common.Address, backend bind.ContractBackend) (*ERC721X, error) {
	contract, err := bindERC721X(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721X{ERC721XCaller: ERC721XCaller{contract: contract}, ERC721XTransactor: ERC721XTransactor{contract: contract}, ERC721XFilterer: ERC721XFilterer{contract: contract}}, nil
}

// NewERC721XCaller creates a new read-only instance of ERC721X, bound to a specific deployed contract.
func NewERC721XCaller(address common.Address, caller bind.ContractCaller) (*ERC721XCaller, error) {
	contract, err := bindERC721X(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721XCaller{contract: contract}, nil
}

// NewERC721XTransactor creates a new write-only instance of ERC721X, bound to a specific deployed contract.
func NewERC721XTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721XTransactor, error) {
	contract, err := bindERC721X(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721XTransactor{contract: contract}, nil
}

// NewERC721XFilterer creates a new log filterer instance of ERC721X, bound to a specific deployed contract.
func NewERC721XFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721XFilterer, error) {
	contract, err := bindERC721X(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721XFilterer{contract: contract}, nil
}

// bindERC721X binds a generic wrapper to an already deployed contract.
func bindERC721X(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721XABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721X *ERC721XRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721X.Contract.ERC721XCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721X *ERC721XRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721X.Contract.ERC721XTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721X *ERC721XRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721X.Contract.ERC721XTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721X *ERC721XCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC721X.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721X *ERC721XTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721X.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721X *ERC721XTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721X.Contract.contract.Transact(opts, method, params...)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_ERC721X *ERC721XCaller) InterfaceIdERC165(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _ERC721X.contract.Call(opts, out, "InterfaceId_ERC165")
	return *ret0, err
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_ERC721X *ERC721XSession) InterfaceIdERC165() ([4]byte, error) {
	return _ERC721X.Contract.InterfaceIdERC165(&_ERC721X.CallOpts)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_ERC721X *ERC721XCallerSession) InterfaceIdERC165() ([4]byte, error) {
	return _ERC721X.Contract.InterfaceIdERC165(&_ERC721X.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC721X *ERC721XCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721X.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC721X *ERC721XSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721X.Contract.BalanceOf(&_ERC721X.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ERC721X *ERC721XCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721X.Contract.BalanceOf(&_ERC721X.CallOpts, _owner)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_ERC721X *ERC721XCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721X.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_ERC721X *ERC721XSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721X.Contract.Exists(&_ERC721X.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_ERC721X *ERC721XCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _ERC721X.Contract.Exists(&_ERC721X.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_ERC721X *ERC721XCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721X.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_ERC721X *ERC721XSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721X.Contract.GetApproved(&_ERC721X.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_ERC721X *ERC721XCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _ERC721X.Contract.GetApproved(&_ERC721X.CallOpts, _tokenId)
}

// ImplementsERC721 is a free data retrieval call binding the contract method 0x1051db34.
//
// Solidity: function implementsERC721() constant returns(bool)
func (_ERC721X *ERC721XCaller) ImplementsERC721(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721X.contract.Call(opts, out, "implementsERC721")
	return *ret0, err
}

// ImplementsERC721 is a free data retrieval call binding the contract method 0x1051db34.
//
// Solidity: function implementsERC721() constant returns(bool)
func (_ERC721X *ERC721XSession) ImplementsERC721() (bool, error) {
	return _ERC721X.Contract.ImplementsERC721(&_ERC721X.CallOpts)
}

// ImplementsERC721 is a free data retrieval call binding the contract method 0x1051db34.
//
// Solidity: function implementsERC721() constant returns(bool)
func (_ERC721X *ERC721XCallerSession) ImplementsERC721() (bool, error) {
	return _ERC721X.Contract.ImplementsERC721(&_ERC721X.CallOpts)
}

// ImplementsERC721X is a free data retrieval call binding the contract method 0x7fb42a36.
//
// Solidity: function implementsERC721X() constant returns(bool)
func (_ERC721X *ERC721XCaller) ImplementsERC721X(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721X.contract.Call(opts, out, "implementsERC721X")
	return *ret0, err
}

// ImplementsERC721X is a free data retrieval call binding the contract method 0x7fb42a36.
//
// Solidity: function implementsERC721X() constant returns(bool)
func (_ERC721X *ERC721XSession) ImplementsERC721X() (bool, error) {
	return _ERC721X.Contract.ImplementsERC721X(&_ERC721X.CallOpts)
}

// ImplementsERC721X is a free data retrieval call binding the contract method 0x7fb42a36.
//
// Solidity: function implementsERC721X() constant returns(bool)
func (_ERC721X *ERC721XCallerSession) ImplementsERC721X() (bool, error) {
	return _ERC721X.Contract.ImplementsERC721X(&_ERC721X.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(isOperator bool)
func (_ERC721X *ERC721XCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721X.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(isOperator bool)
func (_ERC721X *ERC721XSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721X.Contract.IsApprovedForAll(&_ERC721X.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(isOperator bool)
func (_ERC721X *ERC721XCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _ERC721X.Contract.IsApprovedForAll(&_ERC721X.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721X *ERC721XCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721X.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721X *ERC721XSession) Name() (string, error) {
	return _ERC721X.Contract.Name(&_ERC721X.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ERC721X *ERC721XCallerSession) Name() (string, error) {
	return _ERC721X.Contract.Name(&_ERC721X.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC721X *ERC721XCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ERC721X.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC721X *ERC721XSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721X.Contract.OwnerOf(&_ERC721X.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_ERC721X *ERC721XCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721X.Contract.OwnerOf(&_ERC721X.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721X *ERC721XCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC721X.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721X *ERC721XSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721X.Contract.SupportsInterface(&_ERC721X.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_ERC721X *ERC721XCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _ERC721X.Contract.SupportsInterface(&_ERC721X.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721X *ERC721XCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721X.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721X *ERC721XSession) Symbol() (string, error) {
	return _ERC721X.Contract.Symbol(&_ERC721X.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_ERC721X *ERC721XCallerSession) Symbol() (string, error) {
	return _ERC721X.Contract.Symbol(&_ERC721X.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721X *ERC721XCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721X.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721X *ERC721XSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721X.Contract.TokenByIndex(&_ERC721X.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_ERC721X *ERC721XCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _ERC721X.Contract.TokenByIndex(&_ERC721X.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId uint256)
func (_ERC721X *ERC721XCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721X.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId uint256)
func (_ERC721X *ERC721XSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721X.Contract.TokenOfOwnerByIndex(&_ERC721X.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(_tokenId uint256)
func (_ERC721X *ERC721XCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _ERC721X.Contract.TokenOfOwnerByIndex(&_ERC721X.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(tokenUri string)
func (_ERC721X *ERC721XCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ERC721X.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(tokenUri string)
func (_ERC721X *ERC721XSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721X.Contract.TokenURI(&_ERC721X.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(tokenUri string)
func (_ERC721X *ERC721XCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _ERC721X.Contract.TokenURI(&_ERC721X.CallOpts, _tokenId)
}

// TokensOwned is a free data retrieval call binding the contract method 0x21cda790.
//
// Solidity: function tokensOwned(_owner address) constant returns(indexes uint256[], balances uint256[])
func (_ERC721X *ERC721XCaller) TokensOwned(opts *bind.CallOpts, _owner common.Address) (struct {
	Indexes  []*big.Int
	Balances []*big.Int
}, error) {
	ret := new(struct {
		Indexes  []*big.Int
		Balances []*big.Int
	})
	out := ret
	err := _ERC721X.contract.Call(opts, out, "tokensOwned", _owner)
	return *ret, err
}

// TokensOwned is a free data retrieval call binding the contract method 0x21cda790.
//
// Solidity: function tokensOwned(_owner address) constant returns(indexes uint256[], balances uint256[])
func (_ERC721X *ERC721XSession) TokensOwned(_owner common.Address) (struct {
	Indexes  []*big.Int
	Balances []*big.Int
}, error) {
	return _ERC721X.Contract.TokensOwned(&_ERC721X.CallOpts, _owner)
}

// TokensOwned is a free data retrieval call binding the contract method 0x21cda790.
//
// Solidity: function tokensOwned(_owner address) constant returns(indexes uint256[], balances uint256[])
func (_ERC721X *ERC721XCallerSession) TokensOwned(_owner common.Address) (struct {
	Indexes  []*big.Int
	Balances []*big.Int
}, error) {
	return _ERC721X.Contract.TokensOwned(&_ERC721X.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721X *ERC721XCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC721X.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721X *ERC721XSession) TotalSupply() (*big.Int, error) {
	return _ERC721X.Contract.TotalSupply(&_ERC721X.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC721X *ERC721XCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721X.Contract.TotalSupply(&_ERC721X.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721X *ERC721XTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721X.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721X *ERC721XSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721X.Contract.Approve(&_ERC721X.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_ERC721X *ERC721XTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721X.Contract.Approve(&_ERC721X.TransactOpts, _to, _tokenId)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0x17fad7fc.
//
// Solidity: function batchTransferFrom(_from address, _to address, _tokenIds uint256[], _amounts uint256[]) returns()
func (_ERC721X *ERC721XTransactor) BatchTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _ERC721X.contract.Transact(opts, "batchTransferFrom", _from, _to, _tokenIds, _amounts)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0x17fad7fc.
//
// Solidity: function batchTransferFrom(_from address, _to address, _tokenIds uint256[], _amounts uint256[]) returns()
func (_ERC721X *ERC721XSession) BatchTransferFrom(_from common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _ERC721X.Contract.BatchTransferFrom(&_ERC721X.TransactOpts, _from, _to, _tokenIds, _amounts)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0x17fad7fc.
//
// Solidity: function batchTransferFrom(_from address, _to address, _tokenIds uint256[], _amounts uint256[]) returns()
func (_ERC721X *ERC721XTransactorSession) BatchTransferFrom(_from common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int) (*types.Transaction, error) {
	return _ERC721X.Contract.BatchTransferFrom(&_ERC721X.TransactOpts, _from, _to, _tokenIds, _amounts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(_from address, _to address, _tokenIds uint256[], _amounts uint256[], _data bytes) returns()
func (_ERC721X *ERC721XTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721X.contract.Transact(opts, "safeBatchTransferFrom", _from, _to, _tokenIds, _amounts, _data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(_from address, _to address, _tokenIds uint256[], _amounts uint256[], _data bytes) returns()
func (_ERC721X *ERC721XSession) SafeBatchTransferFrom(_from common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721X.Contract.SafeBatchTransferFrom(&_ERC721X.TransactOpts, _from, _to, _tokenIds, _amounts, _data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(_from address, _to address, _tokenIds uint256[], _amounts uint256[], _data bytes) returns()
func (_ERC721X *ERC721XTransactorSession) SafeBatchTransferFrom(_from common.Address, _to common.Address, _tokenIds []*big.Int, _amounts []*big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721X.Contract.SafeBatchTransferFrom(&_ERC721X.TransactOpts, _from, _to, _tokenIds, _amounts, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _amount uint256, _data bytes) returns()
func (_ERC721X *ERC721XTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721X.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _amount, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _amount uint256, _data bytes) returns()
func (_ERC721X *ERC721XSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721X.Contract.SafeTransferFrom(&_ERC721X.TransactOpts, _from, _to, _tokenId, _amount, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _amount uint256, _data bytes) returns()
func (_ERC721X *ERC721XTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int, _data []byte) (*types.Transaction, error) {
	return _ERC721X.Contract.SafeTransferFrom(&_ERC721X.TransactOpts, _from, _to, _tokenId, _amount, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721X *ERC721XTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721X.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721X *ERC721XSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721X.Contract.SetApprovalForAll(&_ERC721X.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_operator address, _approved bool) returns()
func (_ERC721X *ERC721XTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _ERC721X.Contract.SetApprovalForAll(&_ERC721X.TransactOpts, _operator, _approved)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(_to address, _tokenId uint256, _amount uint256) returns()
func (_ERC721X *ERC721XTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ERC721X.contract.Transact(opts, "transfer", _to, _tokenId, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(_to address, _tokenId uint256, _amount uint256) returns()
func (_ERC721X *ERC721XSession) Transfer(_to common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ERC721X.Contract.Transfer(&_ERC721X.TransactOpts, _to, _tokenId, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(_to address, _tokenId uint256, _amount uint256) returns()
func (_ERC721X *ERC721XTransactorSession) Transfer(_to common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ERC721X.Contract.Transfer(&_ERC721X.TransactOpts, _to, _tokenId, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xfe99049a.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256, _amount uint256) returns()
func (_ERC721X *ERC721XTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ERC721X.contract.Transact(opts, "transferFrom", _from, _to, _tokenId, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xfe99049a.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256, _amount uint256) returns()
func (_ERC721X *ERC721XSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ERC721X.Contract.TransferFrom(&_ERC721X.TransactOpts, _from, _to, _tokenId, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xfe99049a.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256, _amount uint256) returns()
func (_ERC721X *ERC721XTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _ERC721X.Contract.TransferFrom(&_ERC721X.TransactOpts, _from, _to, _tokenId, _amount)
}

// ERC721XApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721X contract.
type ERC721XApprovalIterator struct {
	Event *ERC721XApproval // Event containing the contract specifics and raw log

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
func (it *ERC721XApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721XApproval)
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
		it.Event = new(ERC721XApproval)
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
func (it *ERC721XApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721XApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721XApproval represents a Approval event raised by the ERC721X contract.
type ERC721XApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_ERC721X *ERC721XFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*ERC721XApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721X.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721XApprovalIterator{contract: _ERC721X.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_ERC721X *ERC721XFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721XApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721X.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721XApproval)
				if err := _ERC721X.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC721XApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ERC721X contract.
type ERC721XApprovalForAllIterator struct {
	Event *ERC721XApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ERC721XApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721XApprovalForAll)
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
		it.Event = new(ERC721XApprovalForAll)
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
func (it *ERC721XApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721XApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721XApprovalForAll represents a ApprovalForAll event raised by the ERC721X contract.
type ERC721XApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721X *ERC721XFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ERC721XApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721X.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ERC721XApprovalForAllIterator{contract: _ERC721X.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_ERC721X *ERC721XFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ERC721XApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _ERC721X.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721XApprovalForAll)
				if err := _ERC721X.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ERC721XBatchTransferIterator is returned from FilterBatchTransfer and is used to iterate over the raw logs and unpacked data for BatchTransfer events raised by the ERC721X contract.
type ERC721XBatchTransferIterator struct {
	Event *ERC721XBatchTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721XBatchTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721XBatchTransfer)
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
		it.Event = new(ERC721XBatchTransfer)
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
func (it *ERC721XBatchTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721XBatchTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721XBatchTransfer represents a BatchTransfer event raised by the ERC721X contract.
type ERC721XBatchTransfer struct {
	From       common.Address
	To         common.Address
	TokenTypes []*big.Int
	Amounts    []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBatchTransfer is a free log retrieval operation binding the contract event 0xf59807b2c31ca3ba212e90599175c120c556422950bac5be656274483e8581df.
//
// Solidity: e BatchTransfer(from address, to address, tokenTypes uint256[], amounts uint256[])
func (_ERC721X *ERC721XFilterer) FilterBatchTransfer(opts *bind.FilterOpts) (*ERC721XBatchTransferIterator, error) {

	logs, sub, err := _ERC721X.contract.FilterLogs(opts, "BatchTransfer")
	if err != nil {
		return nil, err
	}
	return &ERC721XBatchTransferIterator{contract: _ERC721X.contract, event: "BatchTransfer", logs: logs, sub: sub}, nil
}

// WatchBatchTransfer is a free log subscription operation binding the contract event 0xf59807b2c31ca3ba212e90599175c120c556422950bac5be656274483e8581df.
//
// Solidity: e BatchTransfer(from address, to address, tokenTypes uint256[], amounts uint256[])
func (_ERC721X *ERC721XFilterer) WatchBatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721XBatchTransfer) (event.Subscription, error) {

	logs, sub, err := _ERC721X.contract.WatchLogs(opts, "BatchTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721XBatchTransfer)
				if err := _ERC721X.contract.UnpackLog(event, "BatchTransfer", log); err != nil {
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

// ERC721XTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721X contract.
type ERC721XTransferIterator struct {
	Event *ERC721XTransfer // Event containing the contract specifics and raw log

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
func (it *ERC721XTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721XTransfer)
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
		it.Event = new(ERC721XTransfer)
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
func (it *ERC721XTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721XTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721XTransfer represents a Transfer event raised by the ERC721X contract.
type ERC721XTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_ERC721X *ERC721XFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*ERC721XTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721X.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721XTransferIterator{contract: _ERC721X.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_ERC721X *ERC721XFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721XTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _ERC721X.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721XTransfer)
				if err := _ERC721X.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC721XTransferWithQuantityIterator is returned from FilterTransferWithQuantity and is used to iterate over the raw logs and unpacked data for TransferWithQuantity events raised by the ERC721X contract.
type ERC721XTransferWithQuantityIterator struct {
	Event *ERC721XTransferWithQuantity // Event containing the contract specifics and raw log

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
func (it *ERC721XTransferWithQuantityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721XTransferWithQuantity)
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
		it.Event = new(ERC721XTransferWithQuantity)
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
func (it *ERC721XTransferWithQuantityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721XTransferWithQuantityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721XTransferWithQuantity represents a TransferWithQuantity event raised by the ERC721X contract.
type ERC721XTransferWithQuantity struct {
	From     common.Address
	To       common.Address
	TokenId  *big.Int
	Quantity *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferWithQuantity is a free log retrieval operation binding the contract event 0x2114851a3e2a54429989f46c1ab0743e37ded205d9bbdfd85635aed5bd595a06.
//
// Solidity: e TransferWithQuantity(from indexed address, to indexed address, tokenId indexed uint256, quantity uint256)
func (_ERC721X *ERC721XFilterer) FilterTransferWithQuantity(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ERC721XTransferWithQuantityIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721X.contract.FilterLogs(opts, "TransferWithQuantity", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ERC721XTransferWithQuantityIterator{contract: _ERC721X.contract, event: "TransferWithQuantity", logs: logs, sub: sub}, nil
}

// WatchTransferWithQuantity is a free log subscription operation binding the contract event 0x2114851a3e2a54429989f46c1ab0743e37ded205d9bbdfd85635aed5bd595a06.
//
// Solidity: e TransferWithQuantity(from indexed address, to indexed address, tokenId indexed uint256, quantity uint256)
func (_ERC721X *ERC721XFilterer) WatchTransferWithQuantity(opts *bind.WatchOpts, sink chan<- *ERC721XTransferWithQuantity, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ERC721X.contract.WatchLogs(opts, "TransferWithQuantity", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721XTransferWithQuantity)
				if err := _ERC721X.contract.UnpackLog(event, "TransferWithQuantity", log); err != nil {
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
