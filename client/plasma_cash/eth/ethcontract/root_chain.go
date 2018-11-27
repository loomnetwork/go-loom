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

// RootChainABI is the input ABI used to generate the binding from.
const RootChainABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"bonded\",\"type\":\"uint256\"},{\"name\":\"withdrawable\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"childBlockInterval\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numCoins\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"childChain\",\"outputs\":[{\"name\":\"root\",\"type\":\"bytes32\"},{\"name\":\"createdAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_vmc\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"slot\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"denomination\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SubmittedBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"slot\",\"type\":\"uint64\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"StartedExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"slot\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"txHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"challengingBlockNumber\",\"type\":\"uint256\"}],\"name\":\"ChallengedExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"slot\",\"type\":\"uint64\"}],\"name\":\"RespondedExitChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"slot\",\"type\":\"uint64\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"CoinReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"slot\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"FinalizedExit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FreedBond\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SlashedBond\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrewBonds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"slot\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"mode\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"uid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"denomination\",\"type\":\"uint256\"}],\"name\":\"Withdrew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"submitBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"slot\",\"type\":\"uint64\"},{\"name\":\"prevTxBytes\",\"type\":\"bytes\"},{\"name\":\"exitingTxBytes\",\"type\":\"bytes\"},{\"name\":\"prevTxInclusionProof\",\"type\":\"bytes\"},{\"name\":\"exitingTxInclusionProof\",\"type\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\"},{\"name\":\"blocks\",\"type\":\"uint256[2]\"}],\"name\":\"startExit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"slot\",\"type\":\"uint64\"}],\"name\":\"finalizeExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"slots\",\"type\":\"uint64[]\"}],\"name\":\"finalizeExits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"slot\",\"type\":\"uint64\"}],\"name\":\"cancelExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"slots\",\"type\":\"uint64[]\"}],\"name\":\"cancelExits\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"slot\",\"type\":\"uint64\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"slot\",\"type\":\"uint64\"},{\"name\":\"txBytes\",\"type\":\"bytes\"},{\"name\":\"txInclusionProof\",\"type\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"challengeBefore\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"slot\",\"type\":\"uint64\"},{\"name\":\"challengingTxHash\",\"type\":\"bytes32\"},{\"name\":\"respondingBlockNumber\",\"type\":\"uint256\"},{\"name\":\"respondingTransaction\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"respondChallengeBefore\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"slot\",\"type\":\"uint64\"},{\"name\":\"challengingBlockNumber\",\"type\":\"uint256\"},{\"name\":\"challengingTransaction\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"challengeBetween\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"slot\",\"type\":\"uint64\"},{\"name\":\"challengingBlockNumber\",\"type\":\"uint256\"},{\"name\":\"challengingTransaction\",\"type\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"challengeAfter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawBonds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC20Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_uid\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"depositERC20\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uid\",\"type\":\"uint256\"},{\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"depositERC721\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"txHash\",\"type\":\"bytes32\"},{\"name\":\"root\",\"type\":\"bytes32\"},{\"name\":\"slot\",\"type\":\"uint64\"},{\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"checkMembership\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"slot\",\"type\":\"uint64\"}],\"name\":\"getPlasmaCoin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"slot\",\"type\":\"uint64\"},{\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"getChallenge\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"slot\",\"type\":\"uint64\"}],\"name\":\"getExit\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint8\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getBlockRoot\",\"outputs\":[{\"name\":\"root\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RootChain is an auto generated Go binding around an Ethereum contract.
type RootChain struct {
	RootChainCaller     // Read-only binding to the contract
	RootChainTransactor // Write-only binding to the contract
	RootChainFilterer   // Log filterer for contract events
}

// RootChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RootChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootChainSession struct {
	Contract     *RootChain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootChainCallerSession struct {
	Contract *RootChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RootChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootChainTransactorSession struct {
	Contract     *RootChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RootChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootChainRaw struct {
	Contract *RootChain // Generic contract binding to access the raw methods on
}

// RootChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootChainCallerRaw struct {
	Contract *RootChainCaller // Generic read-only contract binding to access the raw methods on
}

// RootChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootChainTransactorRaw struct {
	Contract *RootChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRootChain creates a new instance of RootChain, bound to a specific deployed contract.
func NewRootChain(address common.Address, backend bind.ContractBackend) (*RootChain, error) {
	contract, err := bindRootChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RootChain{RootChainCaller: RootChainCaller{contract: contract}, RootChainTransactor: RootChainTransactor{contract: contract}, RootChainFilterer: RootChainFilterer{contract: contract}}, nil
}

// NewRootChainCaller creates a new read-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainCaller(address common.Address, caller bind.ContractCaller) (*RootChainCaller, error) {
	contract, err := bindRootChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainCaller{contract: contract}, nil
}

// NewRootChainTransactor creates a new write-only instance of RootChain, bound to a specific deployed contract.
func NewRootChainTransactor(address common.Address, transactor bind.ContractTransactor) (*RootChainTransactor, error) {
	contract, err := bindRootChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RootChainTransactor{contract: contract}, nil
}

// NewRootChainFilterer creates a new log filterer instance of RootChain, bound to a specific deployed contract.
func NewRootChainFilterer(address common.Address, filterer bind.ContractFilterer) (*RootChainFilterer, error) {
	contract, err := bindRootChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RootChainFilterer{contract: contract}, nil
}

// bindRootChain binds a generic wrapper to an already deployed contract.
func bindRootChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChain *RootChainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChain.Contract.RootChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChain *RootChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.Contract.RootChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChain *RootChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChain.Contract.RootChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RootChain *RootChainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RootChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RootChain *RootChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RootChain *RootChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RootChain.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(bonded uint256, withdrawable uint256)
func (_RootChain *RootChainCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (struct {
	Bonded       *big.Int
	Withdrawable *big.Int
}, error) {
	ret := new(struct {
		Bonded       *big.Int
		Withdrawable *big.Int
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "balances", arg0)
	return *ret, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(bonded uint256, withdrawable uint256)
func (_RootChain *RootChainSession) Balances(arg0 common.Address) (struct {
	Bonded       *big.Int
	Withdrawable *big.Int
}, error) {
	return _RootChain.Contract.Balances(&_RootChain.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(bonded uint256, withdrawable uint256)
func (_RootChain *RootChainCallerSession) Balances(arg0 common.Address) (struct {
	Bonded       *big.Int
	Withdrawable *big.Int
}, error) {
	return _RootChain.Contract.Balances(&_RootChain.CallOpts, arg0)
}

// CheckMembership is a free data retrieval call binding the contract method 0xf586df65.
//
// Solidity: function checkMembership(txHash bytes32, root bytes32, slot uint64, proof bytes) constant returns(bool)
func (_RootChain *RootChainCaller) CheckMembership(opts *bind.CallOpts, txHash [32]byte, root [32]byte, slot uint64, proof []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "checkMembership", txHash, root, slot, proof)
	return *ret0, err
}

// CheckMembership is a free data retrieval call binding the contract method 0xf586df65.
//
// Solidity: function checkMembership(txHash bytes32, root bytes32, slot uint64, proof bytes) constant returns(bool)
func (_RootChain *RootChainSession) CheckMembership(txHash [32]byte, root [32]byte, slot uint64, proof []byte) (bool, error) {
	return _RootChain.Contract.CheckMembership(&_RootChain.CallOpts, txHash, root, slot, proof)
}

// CheckMembership is a free data retrieval call binding the contract method 0xf586df65.
//
// Solidity: function checkMembership(txHash bytes32, root bytes32, slot uint64, proof bytes) constant returns(bool)
func (_RootChain *RootChainCallerSession) CheckMembership(txHash [32]byte, root [32]byte, slot uint64, proof []byte) (bool, error) {
	return _RootChain.Contract.CheckMembership(&_RootChain.CallOpts, txHash, root, slot, proof)
}

// ChildBlockInterval is a free data retrieval call binding the contract method 0x38a9e0bc.
//
// Solidity: function childBlockInterval() constant returns(uint256)
func (_RootChain *RootChainCaller) ChildBlockInterval(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "childBlockInterval")
	return *ret0, err
}

// ChildBlockInterval is a free data retrieval call binding the contract method 0x38a9e0bc.
//
// Solidity: function childBlockInterval() constant returns(uint256)
func (_RootChain *RootChainSession) ChildBlockInterval() (*big.Int, error) {
	return _RootChain.Contract.ChildBlockInterval(&_RootChain.CallOpts)
}

// ChildBlockInterval is a free data retrieval call binding the contract method 0x38a9e0bc.
//
// Solidity: function childBlockInterval() constant returns(uint256)
func (_RootChain *RootChainCallerSession) ChildBlockInterval() (*big.Int, error) {
	return _RootChain.Contract.ChildBlockInterval(&_RootChain.CallOpts)
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(root bytes32, createdAt uint256)
func (_RootChain *RootChainCaller) ChildChain(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Root      [32]byte
	CreatedAt *big.Int
}, error) {
	ret := new(struct {
		Root      [32]byte
		CreatedAt *big.Int
	})
	out := ret
	err := _RootChain.contract.Call(opts, out, "childChain", arg0)
	return *ret, err
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(root bytes32, createdAt uint256)
func (_RootChain *RootChainSession) ChildChain(arg0 *big.Int) (struct {
	Root      [32]byte
	CreatedAt *big.Int
}, error) {
	return _RootChain.Contract.ChildChain(&_RootChain.CallOpts, arg0)
}

// ChildChain is a free data retrieval call binding the contract method 0xf95643b1.
//
// Solidity: function childChain( uint256) constant returns(root bytes32, createdAt uint256)
func (_RootChain *RootChainCallerSession) ChildChain(arg0 *big.Int) (struct {
	Root      [32]byte
	CreatedAt *big.Int
}, error) {
	return _RootChain.Contract.ChildChain(&_RootChain.CallOpts, arg0)
}

// CurrentBlock is a free data retrieval call binding the contract method 0xe12ed13c.
//
// Solidity: function currentBlock() constant returns(uint256)
func (_RootChain *RootChainCaller) CurrentBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "currentBlock")
	return *ret0, err
}

// CurrentBlock is a free data retrieval call binding the contract method 0xe12ed13c.
//
// Solidity: function currentBlock() constant returns(uint256)
func (_RootChain *RootChainSession) CurrentBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentBlock(&_RootChain.CallOpts)
}

// CurrentBlock is a free data retrieval call binding the contract method 0xe12ed13c.
//
// Solidity: function currentBlock() constant returns(uint256)
func (_RootChain *RootChainCallerSession) CurrentBlock() (*big.Int, error) {
	return _RootChain.Contract.CurrentBlock(&_RootChain.CallOpts)
}

// GetBlockRoot is a free data retrieval call binding the contract method 0xe41a5d17.
//
// Solidity: function getBlockRoot(blockNumber uint256) constant returns(root bytes32)
func (_RootChain *RootChainCaller) GetBlockRoot(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "getBlockRoot", blockNumber)
	return *ret0, err
}

// GetBlockRoot is a free data retrieval call binding the contract method 0xe41a5d17.
//
// Solidity: function getBlockRoot(blockNumber uint256) constant returns(root bytes32)
func (_RootChain *RootChainSession) GetBlockRoot(blockNumber *big.Int) ([32]byte, error) {
	return _RootChain.Contract.GetBlockRoot(&_RootChain.CallOpts, blockNumber)
}

// GetBlockRoot is a free data retrieval call binding the contract method 0xe41a5d17.
//
// Solidity: function getBlockRoot(blockNumber uint256) constant returns(root bytes32)
func (_RootChain *RootChainCallerSession) GetBlockRoot(blockNumber *big.Int) ([32]byte, error) {
	return _RootChain.Contract.GetBlockRoot(&_RootChain.CallOpts, blockNumber)
}

// GetChallenge is a free data retrieval call binding the contract method 0x81686e6b.
//
// Solidity: function getChallenge(slot uint64, txHash bytes32) constant returns(address, address, bytes32, uint256)
func (_RootChain *RootChainCaller) GetChallenge(opts *bind.CallOpts, slot uint64, txHash [32]byte) (common.Address, common.Address, [32]byte, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
		ret2 = new([32]byte)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _RootChain.contract.Call(opts, out, "getChallenge", slot, txHash)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetChallenge is a free data retrieval call binding the contract method 0x81686e6b.
//
// Solidity: function getChallenge(slot uint64, txHash bytes32) constant returns(address, address, bytes32, uint256)
func (_RootChain *RootChainSession) GetChallenge(slot uint64, txHash [32]byte) (common.Address, common.Address, [32]byte, *big.Int, error) {
	return _RootChain.Contract.GetChallenge(&_RootChain.CallOpts, slot, txHash)
}

// GetChallenge is a free data retrieval call binding the contract method 0x81686e6b.
//
// Solidity: function getChallenge(slot uint64, txHash bytes32) constant returns(address, address, bytes32, uint256)
func (_RootChain *RootChainCallerSession) GetChallenge(slot uint64, txHash [32]byte) (common.Address, common.Address, [32]byte, *big.Int, error) {
	return _RootChain.Contract.GetChallenge(&_RootChain.CallOpts, slot, txHash)
}

// GetExit is a free data retrieval call binding the contract method 0xd157796e.
//
// Solidity: function getExit(slot uint64) constant returns(address, uint256, uint256, uint8, uint256)
func (_RootChain *RootChainCaller) GetExit(opts *bind.CallOpts, slot uint64) (common.Address, *big.Int, *big.Int, uint8, *big.Int, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(uint8)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _RootChain.contract.Call(opts, out, "getExit", slot)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetExit is a free data retrieval call binding the contract method 0xd157796e.
//
// Solidity: function getExit(slot uint64) constant returns(address, uint256, uint256, uint8, uint256)
func (_RootChain *RootChainSession) GetExit(slot uint64) (common.Address, *big.Int, *big.Int, uint8, *big.Int, error) {
	return _RootChain.Contract.GetExit(&_RootChain.CallOpts, slot)
}

// GetExit is a free data retrieval call binding the contract method 0xd157796e.
//
// Solidity: function getExit(slot uint64) constant returns(address, uint256, uint256, uint8, uint256)
func (_RootChain *RootChainCallerSession) GetExit(slot uint64) (common.Address, *big.Int, *big.Int, uint8, *big.Int, error) {
	return _RootChain.Contract.GetExit(&_RootChain.CallOpts, slot)
}

// GetPlasmaCoin is a free data retrieval call binding the contract method 0xf8353cf0.
//
// Solidity: function getPlasmaCoin(slot uint64) constant returns(uint256, uint256, uint256, address, uint8, uint8, address)
func (_RootChain *RootChainCaller) GetPlasmaCoin(opts *bind.CallOpts, slot uint64) (*big.Int, *big.Int, *big.Int, common.Address, uint8, uint8, common.Address, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(common.Address)
		ret4 = new(uint8)
		ret5 = new(uint8)
		ret6 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _RootChain.contract.Call(opts, out, "getPlasmaCoin", slot)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetPlasmaCoin is a free data retrieval call binding the contract method 0xf8353cf0.
//
// Solidity: function getPlasmaCoin(slot uint64) constant returns(uint256, uint256, uint256, address, uint8, uint8, address)
func (_RootChain *RootChainSession) GetPlasmaCoin(slot uint64) (*big.Int, *big.Int, *big.Int, common.Address, uint8, uint8, common.Address, error) {
	return _RootChain.Contract.GetPlasmaCoin(&_RootChain.CallOpts, slot)
}

// GetPlasmaCoin is a free data retrieval call binding the contract method 0xf8353cf0.
//
// Solidity: function getPlasmaCoin(slot uint64) constant returns(uint256, uint256, uint256, address, uint8, uint8, address)
func (_RootChain *RootChainCallerSession) GetPlasmaCoin(slot uint64) (*big.Int, *big.Int, *big.Int, common.Address, uint8, uint8, common.Address, error) {
	return _RootChain.Contract.GetPlasmaCoin(&_RootChain.CallOpts, slot)
}

// NumCoins is a free data retrieval call binding the contract method 0xa9737595.
//
// Solidity: function numCoins() constant returns(uint64)
func (_RootChain *RootChainCaller) NumCoins(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _RootChain.contract.Call(opts, out, "numCoins")
	return *ret0, err
}

// NumCoins is a free data retrieval call binding the contract method 0xa9737595.
//
// Solidity: function numCoins() constant returns(uint64)
func (_RootChain *RootChainSession) NumCoins() (uint64, error) {
	return _RootChain.Contract.NumCoins(&_RootChain.CallOpts)
}

// NumCoins is a free data retrieval call binding the contract method 0xa9737595.
//
// Solidity: function numCoins() constant returns(uint64)
func (_RootChain *RootChainCallerSession) NumCoins() (uint64, error) {
	return _RootChain.Contract.NumCoins(&_RootChain.CallOpts)
}

// CancelExit is a paid mutator transaction binding the contract method 0xcbd8da63.
//
// Solidity: function cancelExit(slot uint64) returns()
func (_RootChain *RootChainTransactor) CancelExit(opts *bind.TransactOpts, slot uint64) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "cancelExit", slot)
}

// CancelExit is a paid mutator transaction binding the contract method 0xcbd8da63.
//
// Solidity: function cancelExit(slot uint64) returns()
func (_RootChain *RootChainSession) CancelExit(slot uint64) (*types.Transaction, error) {
	return _RootChain.Contract.CancelExit(&_RootChain.TransactOpts, slot)
}

// CancelExit is a paid mutator transaction binding the contract method 0xcbd8da63.
//
// Solidity: function cancelExit(slot uint64) returns()
func (_RootChain *RootChainTransactorSession) CancelExit(slot uint64) (*types.Transaction, error) {
	return _RootChain.Contract.CancelExit(&_RootChain.TransactOpts, slot)
}

// CancelExits is a paid mutator transaction binding the contract method 0xdf3b8377.
//
// Solidity: function cancelExits(slots uint64[]) returns()
func (_RootChain *RootChainTransactor) CancelExits(opts *bind.TransactOpts, slots []uint64) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "cancelExits", slots)
}

// CancelExits is a paid mutator transaction binding the contract method 0xdf3b8377.
//
// Solidity: function cancelExits(slots uint64[]) returns()
func (_RootChain *RootChainSession) CancelExits(slots []uint64) (*types.Transaction, error) {
	return _RootChain.Contract.CancelExits(&_RootChain.TransactOpts, slots)
}

// CancelExits is a paid mutator transaction binding the contract method 0xdf3b8377.
//
// Solidity: function cancelExits(slots uint64[]) returns()
func (_RootChain *RootChainTransactorSession) CancelExits(slots []uint64) (*types.Transaction, error) {
	return _RootChain.Contract.CancelExits(&_RootChain.TransactOpts, slots)
}

// ChallengeAfter is a paid mutator transaction binding the contract method 0x058a6508.
//
// Solidity: function challengeAfter(slot uint64, challengingBlockNumber uint256, challengingTransaction bytes, proof bytes, signature bytes) returns()
func (_RootChain *RootChainTransactor) ChallengeAfter(opts *bind.TransactOpts, slot uint64, challengingBlockNumber *big.Int, challengingTransaction []byte, proof []byte, signature []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "challengeAfter", slot, challengingBlockNumber, challengingTransaction, proof, signature)
}

// ChallengeAfter is a paid mutator transaction binding the contract method 0x058a6508.
//
// Solidity: function challengeAfter(slot uint64, challengingBlockNumber uint256, challengingTransaction bytes, proof bytes, signature bytes) returns()
func (_RootChain *RootChainSession) ChallengeAfter(slot uint64, challengingBlockNumber *big.Int, challengingTransaction []byte, proof []byte, signature []byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeAfter(&_RootChain.TransactOpts, slot, challengingBlockNumber, challengingTransaction, proof, signature)
}

// ChallengeAfter is a paid mutator transaction binding the contract method 0x058a6508.
//
// Solidity: function challengeAfter(slot uint64, challengingBlockNumber uint256, challengingTransaction bytes, proof bytes, signature bytes) returns()
func (_RootChain *RootChainTransactorSession) ChallengeAfter(slot uint64, challengingBlockNumber *big.Int, challengingTransaction []byte, proof []byte, signature []byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeAfter(&_RootChain.TransactOpts, slot, challengingBlockNumber, challengingTransaction, proof, signature)
}

// ChallengeBefore is a paid mutator transaction binding the contract method 0x7eb7b0f6.
//
// Solidity: function challengeBefore(slot uint64, txBytes bytes, txInclusionProof bytes, signature bytes, blockNumber uint256) returns()
func (_RootChain *RootChainTransactor) ChallengeBefore(opts *bind.TransactOpts, slot uint64, txBytes []byte, txInclusionProof []byte, signature []byte, blockNumber *big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "challengeBefore", slot, txBytes, txInclusionProof, signature, blockNumber)
}

// ChallengeBefore is a paid mutator transaction binding the contract method 0x7eb7b0f6.
//
// Solidity: function challengeBefore(slot uint64, txBytes bytes, txInclusionProof bytes, signature bytes, blockNumber uint256) returns()
func (_RootChain *RootChainSession) ChallengeBefore(slot uint64, txBytes []byte, txInclusionProof []byte, signature []byte, blockNumber *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeBefore(&_RootChain.TransactOpts, slot, txBytes, txInclusionProof, signature, blockNumber)
}

// ChallengeBefore is a paid mutator transaction binding the contract method 0x7eb7b0f6.
//
// Solidity: function challengeBefore(slot uint64, txBytes bytes, txInclusionProof bytes, signature bytes, blockNumber uint256) returns()
func (_RootChain *RootChainTransactorSession) ChallengeBefore(slot uint64, txBytes []byte, txInclusionProof []byte, signature []byte, blockNumber *big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeBefore(&_RootChain.TransactOpts, slot, txBytes, txInclusionProof, signature, blockNumber)
}

// ChallengeBetween is a paid mutator transaction binding the contract method 0xf6d0ba1a.
//
// Solidity: function challengeBetween(slot uint64, challengingBlockNumber uint256, challengingTransaction bytes, proof bytes, signature bytes) returns()
func (_RootChain *RootChainTransactor) ChallengeBetween(opts *bind.TransactOpts, slot uint64, challengingBlockNumber *big.Int, challengingTransaction []byte, proof []byte, signature []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "challengeBetween", slot, challengingBlockNumber, challengingTransaction, proof, signature)
}

// ChallengeBetween is a paid mutator transaction binding the contract method 0xf6d0ba1a.
//
// Solidity: function challengeBetween(slot uint64, challengingBlockNumber uint256, challengingTransaction bytes, proof bytes, signature bytes) returns()
func (_RootChain *RootChainSession) ChallengeBetween(slot uint64, challengingBlockNumber *big.Int, challengingTransaction []byte, proof []byte, signature []byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeBetween(&_RootChain.TransactOpts, slot, challengingBlockNumber, challengingTransaction, proof, signature)
}

// ChallengeBetween is a paid mutator transaction binding the contract method 0xf6d0ba1a.
//
// Solidity: function challengeBetween(slot uint64, challengingBlockNumber uint256, challengingTransaction bytes, proof bytes, signature bytes) returns()
func (_RootChain *RootChainTransactorSession) ChallengeBetween(slot uint64, challengingBlockNumber *big.Int, challengingTransaction []byte, proof []byte, signature []byte) (*types.Transaction, error) {
	return _RootChain.Contract.ChallengeBetween(&_RootChain.TransactOpts, slot, challengingBlockNumber, challengingTransaction, proof, signature)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x392d661c.
//
// Solidity: function depositERC20(amount uint256, contractAddress address) returns()
func (_RootChain *RootChainTransactor) DepositERC20(opts *bind.TransactOpts, amount *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "depositERC20", amount, contractAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x392d661c.
//
// Solidity: function depositERC20(amount uint256, contractAddress address) returns()
func (_RootChain *RootChainSession) DepositERC20(amount *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.DepositERC20(&_RootChain.TransactOpts, amount, contractAddress)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0x392d661c.
//
// Solidity: function depositERC20(amount uint256, contractAddress address) returns()
func (_RootChain *RootChainTransactorSession) DepositERC20(amount *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.DepositERC20(&_RootChain.TransactOpts, amount, contractAddress)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0xc8ac6542.
//
// Solidity: function depositERC721(uid uint256, contractAddress address) returns()
func (_RootChain *RootChainTransactor) DepositERC721(opts *bind.TransactOpts, uid *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "depositERC721", uid, contractAddress)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0xc8ac6542.
//
// Solidity: function depositERC721(uid uint256, contractAddress address) returns()
func (_RootChain *RootChainSession) DepositERC721(uid *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.DepositERC721(&_RootChain.TransactOpts, uid, contractAddress)
}

// DepositERC721 is a paid mutator transaction binding the contract method 0xc8ac6542.
//
// Solidity: function depositERC721(uid uint256, contractAddress address) returns()
func (_RootChain *RootChainTransactorSession) DepositERC721(uid *big.Int, contractAddress common.Address) (*types.Transaction, error) {
	return _RootChain.Contract.DepositERC721(&_RootChain.TransactOpts, uid, contractAddress)
}

// FinalizeExit is a paid mutator transaction binding the contract method 0x78417214.
//
// Solidity: function finalizeExit(slot uint64) returns()
func (_RootChain *RootChainTransactor) FinalizeExit(opts *bind.TransactOpts, slot uint64) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "finalizeExit", slot)
}

// FinalizeExit is a paid mutator transaction binding the contract method 0x78417214.
//
// Solidity: function finalizeExit(slot uint64) returns()
func (_RootChain *RootChainSession) FinalizeExit(slot uint64) (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeExit(&_RootChain.TransactOpts, slot)
}

// FinalizeExit is a paid mutator transaction binding the contract method 0x78417214.
//
// Solidity: function finalizeExit(slot uint64) returns()
func (_RootChain *RootChainTransactorSession) FinalizeExit(slot uint64) (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeExit(&_RootChain.TransactOpts, slot)
}

// FinalizeExits is a paid mutator transaction binding the contract method 0x74c0a56f.
//
// Solidity: function finalizeExits(slots uint64[]) returns()
func (_RootChain *RootChainTransactor) FinalizeExits(opts *bind.TransactOpts, slots []uint64) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "finalizeExits", slots)
}

// FinalizeExits is a paid mutator transaction binding the contract method 0x74c0a56f.
//
// Solidity: function finalizeExits(slots uint64[]) returns()
func (_RootChain *RootChainSession) FinalizeExits(slots []uint64) (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeExits(&_RootChain.TransactOpts, slots)
}

// FinalizeExits is a paid mutator transaction binding the contract method 0x74c0a56f.
//
// Solidity: function finalizeExits(slots uint64[]) returns()
func (_RootChain *RootChainTransactorSession) FinalizeExits(slots []uint64) (*types.Transaction, error) {
	return _RootChain.Contract.FinalizeExits(&_RootChain.TransactOpts, slots)
}

// OnERC20Received is a paid mutator transaction binding the contract method 0x65d83056.
//
// Solidity: function onERC20Received(_from address, _amount uint256,  bytes) returns(bytes4)
func (_RootChain *RootChainTransactor) OnERC20Received(opts *bind.TransactOpts, _from common.Address, _amount *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "onERC20Received", _from, _amount, arg2)
}

// OnERC20Received is a paid mutator transaction binding the contract method 0x65d83056.
//
// Solidity: function onERC20Received(_from address, _amount uint256,  bytes) returns(bytes4)
func (_RootChain *RootChainSession) OnERC20Received(_from common.Address, _amount *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _RootChain.Contract.OnERC20Received(&_RootChain.TransactOpts, _from, _amount, arg2)
}

// OnERC20Received is a paid mutator transaction binding the contract method 0x65d83056.
//
// Solidity: function onERC20Received(_from address, _amount uint256,  bytes) returns(bytes4)
func (_RootChain *RootChainTransactorSession) OnERC20Received(_from common.Address, _amount *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _RootChain.Contract.OnERC20Received(&_RootChain.TransactOpts, _from, _amount, arg2)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0xf0b9e5ba.
//
// Solidity: function onERC721Received(_from address, _uid uint256,  bytes) returns(bytes4)
func (_RootChain *RootChainTransactor) OnERC721Received(opts *bind.TransactOpts, _from common.Address, _uid *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "onERC721Received", _from, _uid, arg2)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0xf0b9e5ba.
//
// Solidity: function onERC721Received(_from address, _uid uint256,  bytes) returns(bytes4)
func (_RootChain *RootChainSession) OnERC721Received(_from common.Address, _uid *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _RootChain.Contract.OnERC721Received(&_RootChain.TransactOpts, _from, _uid, arg2)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0xf0b9e5ba.
//
// Solidity: function onERC721Received(_from address, _uid uint256,  bytes) returns(bytes4)
func (_RootChain *RootChainTransactorSession) OnERC721Received(_from common.Address, _uid *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _RootChain.Contract.OnERC721Received(&_RootChain.TransactOpts, _from, _uid, arg2)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RootChain *RootChainTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RootChain *RootChainSession) Pause() (*types.Transaction, error) {
	return _RootChain.Contract.Pause(&_RootChain.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RootChain *RootChainTransactorSession) Pause() (*types.Transaction, error) {
	return _RootChain.Contract.Pause(&_RootChain.TransactOpts)
}

// RespondChallengeBefore is a paid mutator transaction binding the contract method 0x4d69a8a1.
//
// Solidity: function respondChallengeBefore(slot uint64, challengingTxHash bytes32, respondingBlockNumber uint256, respondingTransaction bytes, proof bytes, signature bytes) returns()
func (_RootChain *RootChainTransactor) RespondChallengeBefore(opts *bind.TransactOpts, slot uint64, challengingTxHash [32]byte, respondingBlockNumber *big.Int, respondingTransaction []byte, proof []byte, signature []byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "respondChallengeBefore", slot, challengingTxHash, respondingBlockNumber, respondingTransaction, proof, signature)
}

// RespondChallengeBefore is a paid mutator transaction binding the contract method 0x4d69a8a1.
//
// Solidity: function respondChallengeBefore(slot uint64, challengingTxHash bytes32, respondingBlockNumber uint256, respondingTransaction bytes, proof bytes, signature bytes) returns()
func (_RootChain *RootChainSession) RespondChallengeBefore(slot uint64, challengingTxHash [32]byte, respondingBlockNumber *big.Int, respondingTransaction []byte, proof []byte, signature []byte) (*types.Transaction, error) {
	return _RootChain.Contract.RespondChallengeBefore(&_RootChain.TransactOpts, slot, challengingTxHash, respondingBlockNumber, respondingTransaction, proof, signature)
}

// RespondChallengeBefore is a paid mutator transaction binding the contract method 0x4d69a8a1.
//
// Solidity: function respondChallengeBefore(slot uint64, challengingTxHash bytes32, respondingBlockNumber uint256, respondingTransaction bytes, proof bytes, signature bytes) returns()
func (_RootChain *RootChainTransactorSession) RespondChallengeBefore(slot uint64, challengingTxHash [32]byte, respondingBlockNumber *big.Int, respondingTransaction []byte, proof []byte, signature []byte) (*types.Transaction, error) {
	return _RootChain.Contract.RespondChallengeBefore(&_RootChain.TransactOpts, slot, challengingTxHash, respondingBlockNumber, respondingTransaction, proof, signature)
}

// StartExit is a paid mutator transaction binding the contract method 0xe9a067c0.
//
// Solidity: function startExit(slot uint64, prevTxBytes bytes, exitingTxBytes bytes, prevTxInclusionProof bytes, exitingTxInclusionProof bytes, signature bytes, blocks uint256[2]) returns()
func (_RootChain *RootChainTransactor) StartExit(opts *bind.TransactOpts, slot uint64, prevTxBytes []byte, exitingTxBytes []byte, prevTxInclusionProof []byte, exitingTxInclusionProof []byte, signature []byte, blocks [2]*big.Int) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "startExit", slot, prevTxBytes, exitingTxBytes, prevTxInclusionProof, exitingTxInclusionProof, signature, blocks)
}

// StartExit is a paid mutator transaction binding the contract method 0xe9a067c0.
//
// Solidity: function startExit(slot uint64, prevTxBytes bytes, exitingTxBytes bytes, prevTxInclusionProof bytes, exitingTxInclusionProof bytes, signature bytes, blocks uint256[2]) returns()
func (_RootChain *RootChainSession) StartExit(slot uint64, prevTxBytes []byte, exitingTxBytes []byte, prevTxInclusionProof []byte, exitingTxInclusionProof []byte, signature []byte, blocks [2]*big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, slot, prevTxBytes, exitingTxBytes, prevTxInclusionProof, exitingTxInclusionProof, signature, blocks)
}

// StartExit is a paid mutator transaction binding the contract method 0xe9a067c0.
//
// Solidity: function startExit(slot uint64, prevTxBytes bytes, exitingTxBytes bytes, prevTxInclusionProof bytes, exitingTxInclusionProof bytes, signature bytes, blocks uint256[2]) returns()
func (_RootChain *RootChainTransactorSession) StartExit(slot uint64, prevTxBytes []byte, exitingTxBytes []byte, prevTxInclusionProof []byte, exitingTxInclusionProof []byte, signature []byte, blocks [2]*big.Int) (*types.Transaction, error) {
	return _RootChain.Contract.StartExit(&_RootChain.TransactOpts, slot, prevTxBytes, exitingTxBytes, prevTxInclusionProof, exitingTxInclusionProof, signature, blocks)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0x149be19f.
//
// Solidity: function submitBlock(blockNumber uint256, root bytes32) returns()
func (_RootChain *RootChainTransactor) SubmitBlock(opts *bind.TransactOpts, blockNumber *big.Int, root [32]byte) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "submitBlock", blockNumber, root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0x149be19f.
//
// Solidity: function submitBlock(blockNumber uint256, root bytes32) returns()
func (_RootChain *RootChainSession) SubmitBlock(blockNumber *big.Int, root [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitBlock(&_RootChain.TransactOpts, blockNumber, root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0x149be19f.
//
// Solidity: function submitBlock(blockNumber uint256, root bytes32) returns()
func (_RootChain *RootChainTransactorSession) SubmitBlock(blockNumber *big.Int, root [32]byte) (*types.Transaction, error) {
	return _RootChain.Contract.SubmitBlock(&_RootChain.TransactOpts, blockNumber, root)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RootChain *RootChainTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RootChain *RootChainSession) Unpause() (*types.Transaction, error) {
	return _RootChain.Contract.Unpause(&_RootChain.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RootChain *RootChainTransactorSession) Unpause() (*types.Transaction, error) {
	return _RootChain.Contract.Unpause(&_RootChain.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x750f0acc.
//
// Solidity: function withdraw(slot uint64) returns()
func (_RootChain *RootChainTransactor) Withdraw(opts *bind.TransactOpts, slot uint64) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "withdraw", slot)
}

// Withdraw is a paid mutator transaction binding the contract method 0x750f0acc.
//
// Solidity: function withdraw(slot uint64) returns()
func (_RootChain *RootChainSession) Withdraw(slot uint64) (*types.Transaction, error) {
	return _RootChain.Contract.Withdraw(&_RootChain.TransactOpts, slot)
}

// Withdraw is a paid mutator transaction binding the contract method 0x750f0acc.
//
// Solidity: function withdraw(slot uint64) returns()
func (_RootChain *RootChainTransactorSession) Withdraw(slot uint64) (*types.Transaction, error) {
	return _RootChain.Contract.Withdraw(&_RootChain.TransactOpts, slot)
}

// WithdrawBonds is a paid mutator transaction binding the contract method 0x1cc6ffa0.
//
// Solidity: function withdrawBonds() returns()
func (_RootChain *RootChainTransactor) WithdrawBonds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RootChain.contract.Transact(opts, "withdrawBonds")
}

// WithdrawBonds is a paid mutator transaction binding the contract method 0x1cc6ffa0.
//
// Solidity: function withdrawBonds() returns()
func (_RootChain *RootChainSession) WithdrawBonds() (*types.Transaction, error) {
	return _RootChain.Contract.WithdrawBonds(&_RootChain.TransactOpts)
}

// WithdrawBonds is a paid mutator transaction binding the contract method 0x1cc6ffa0.
//
// Solidity: function withdrawBonds() returns()
func (_RootChain *RootChainTransactorSession) WithdrawBonds() (*types.Transaction, error) {
	return _RootChain.Contract.WithdrawBonds(&_RootChain.TransactOpts)
}

// RootChainChallengedExitIterator is returned from FilterChallengedExit and is used to iterate over the raw logs and unpacked data for ChallengedExit events raised by the RootChain contract.
type RootChainChallengedExitIterator struct {
	Event *RootChainChallengedExit // Event containing the contract specifics and raw log

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
func (it *RootChainChallengedExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainChallengedExit)
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
		it.Event = new(RootChainChallengedExit)
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
func (it *RootChainChallengedExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainChallengedExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainChallengedExit represents a ChallengedExit event raised by the RootChain contract.
type RootChainChallengedExit struct {
	Slot                   uint64
	TxHash                 [32]byte
	ChallengingBlockNumber *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterChallengedExit is a free log retrieval operation binding the contract event 0x057d34b2360e71f2764a7189966401bc058621905c0bef7123a6bdfe9a13284b.
//
// Solidity: e ChallengedExit(slot indexed uint64, txHash bytes32, challengingBlockNumber uint256)
func (_RootChain *RootChainFilterer) FilterChallengedExit(opts *bind.FilterOpts, slot []uint64) (*RootChainChallengedExitIterator, error) {

	var slotRule []interface{}
	for _, slotItem := range slot {
		slotRule = append(slotRule, slotItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "ChallengedExit", slotRule)
	if err != nil {
		return nil, err
	}
	return &RootChainChallengedExitIterator{contract: _RootChain.contract, event: "ChallengedExit", logs: logs, sub: sub}, nil
}

// WatchChallengedExit is a free log subscription operation binding the contract event 0x057d34b2360e71f2764a7189966401bc058621905c0bef7123a6bdfe9a13284b.
//
// Solidity: e ChallengedExit(slot indexed uint64, txHash bytes32, challengingBlockNumber uint256)
func (_RootChain *RootChainFilterer) WatchChallengedExit(opts *bind.WatchOpts, sink chan<- *RootChainChallengedExit, slot []uint64) (event.Subscription, error) {

	var slotRule []interface{}
	for _, slotItem := range slot {
		slotRule = append(slotRule, slotItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "ChallengedExit", slotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainChallengedExit)
				if err := _RootChain.contract.UnpackLog(event, "ChallengedExit", log); err != nil {
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

// RootChainCoinResetIterator is returned from FilterCoinReset and is used to iterate over the raw logs and unpacked data for CoinReset events raised by the RootChain contract.
type RootChainCoinResetIterator struct {
	Event *RootChainCoinReset // Event containing the contract specifics and raw log

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
func (it *RootChainCoinResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainCoinReset)
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
		it.Event = new(RootChainCoinReset)
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
func (it *RootChainCoinResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainCoinResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainCoinReset represents a CoinReset event raised by the RootChain contract.
type RootChainCoinReset struct {
	Slot  uint64
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCoinReset is a free log retrieval operation binding the contract event 0x35cc9c4d0d9a2ce748d131382582af17dbf1f0d72167ea7906ea3f18597b31a4.
//
// Solidity: e CoinReset(slot indexed uint64, owner indexed address)
func (_RootChain *RootChainFilterer) FilterCoinReset(opts *bind.FilterOpts, slot []uint64, owner []common.Address) (*RootChainCoinResetIterator, error) {

	var slotRule []interface{}
	for _, slotItem := range slot {
		slotRule = append(slotRule, slotItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "CoinReset", slotRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &RootChainCoinResetIterator{contract: _RootChain.contract, event: "CoinReset", logs: logs, sub: sub}, nil
}

// WatchCoinReset is a free log subscription operation binding the contract event 0x35cc9c4d0d9a2ce748d131382582af17dbf1f0d72167ea7906ea3f18597b31a4.
//
// Solidity: e CoinReset(slot indexed uint64, owner indexed address)
func (_RootChain *RootChainFilterer) WatchCoinReset(opts *bind.WatchOpts, sink chan<- *RootChainCoinReset, slot []uint64, owner []common.Address) (event.Subscription, error) {

	var slotRule []interface{}
	for _, slotItem := range slot {
		slotRule = append(slotRule, slotItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "CoinReset", slotRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainCoinReset)
				if err := _RootChain.contract.UnpackLog(event, "CoinReset", log); err != nil {
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

// RootChainDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the RootChain contract.
type RootChainDepositIterator struct {
	Event *RootChainDeposit // Event containing the contract specifics and raw log

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
func (it *RootChainDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainDeposit)
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
		it.Event = new(RootChainDeposit)
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
func (it *RootChainDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainDeposit represents a Deposit event raised by the RootChain contract.
type RootChainDeposit struct {
	Slot            uint64
	BlockNumber     *big.Int
	Denomination    *big.Int
	From            common.Address
	ContractAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xaa231413cf2b61ec7d73eeef7c722cd8ef6ed7a76e4d9e533867281e94f9a823.
//
// Solidity: e Deposit(slot indexed uint64, blockNumber uint256, denomination uint256, from indexed address, contractAddress indexed address)
func (_RootChain *RootChainFilterer) FilterDeposit(opts *bind.FilterOpts, slot []uint64, from []common.Address, contractAddress []common.Address) (*RootChainDepositIterator, error) {

	var slotRule []interface{}
	for _, slotItem := range slot {
		slotRule = append(slotRule, slotItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "Deposit", slotRule, fromRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return &RootChainDepositIterator{contract: _RootChain.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xaa231413cf2b61ec7d73eeef7c722cd8ef6ed7a76e4d9e533867281e94f9a823.
//
// Solidity: e Deposit(slot indexed uint64, blockNumber uint256, denomination uint256, from indexed address, contractAddress indexed address)
func (_RootChain *RootChainFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *RootChainDeposit, slot []uint64, from []common.Address, contractAddress []common.Address) (event.Subscription, error) {

	var slotRule []interface{}
	for _, slotItem := range slot {
		slotRule = append(slotRule, slotItem)
	}

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "Deposit", slotRule, fromRule, contractAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainDeposit)
				if err := _RootChain.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// RootChainFinalizedExitIterator is returned from FilterFinalizedExit and is used to iterate over the raw logs and unpacked data for FinalizedExit events raised by the RootChain contract.
type RootChainFinalizedExitIterator struct {
	Event *RootChainFinalizedExit // Event containing the contract specifics and raw log

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
func (it *RootChainFinalizedExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainFinalizedExit)
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
		it.Event = new(RootChainFinalizedExit)
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
func (it *RootChainFinalizedExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainFinalizedExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainFinalizedExit represents a FinalizedExit event raised by the RootChain contract.
type RootChainFinalizedExit struct {
	Slot  uint64
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFinalizedExit is a free log retrieval operation binding the contract event 0x432647a5fb9bdea356d78f8e3d83b6ddc2e78b4e4a702ac7eb968f7fe03d9dda.
//
// Solidity: e FinalizedExit(slot indexed uint64, owner address)
func (_RootChain *RootChainFilterer) FilterFinalizedExit(opts *bind.FilterOpts, slot []uint64) (*RootChainFinalizedExitIterator, error) {

	var slotRule []interface{}
	for _, slotItem := range slot {
		slotRule = append(slotRule, slotItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "FinalizedExit", slotRule)
	if err != nil {
		return nil, err
	}
	return &RootChainFinalizedExitIterator{contract: _RootChain.contract, event: "FinalizedExit", logs: logs, sub: sub}, nil
}

// WatchFinalizedExit is a free log subscription operation binding the contract event 0x432647a5fb9bdea356d78f8e3d83b6ddc2e78b4e4a702ac7eb968f7fe03d9dda.
//
// Solidity: e FinalizedExit(slot indexed uint64, owner address)
func (_RootChain *RootChainFilterer) WatchFinalizedExit(opts *bind.WatchOpts, sink chan<- *RootChainFinalizedExit, slot []uint64) (event.Subscription, error) {

	var slotRule []interface{}
	for _, slotItem := range slot {
		slotRule = append(slotRule, slotItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "FinalizedExit", slotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainFinalizedExit)
				if err := _RootChain.contract.UnpackLog(event, "FinalizedExit", log); err != nil {
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

// RootChainFreedBondIterator is returned from FilterFreedBond and is used to iterate over the raw logs and unpacked data for FreedBond events raised by the RootChain contract.
type RootChainFreedBondIterator struct {
	Event *RootChainFreedBond // Event containing the contract specifics and raw log

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
func (it *RootChainFreedBondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainFreedBond)
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
		it.Event = new(RootChainFreedBond)
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
func (it *RootChainFreedBondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainFreedBondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainFreedBond represents a FreedBond event raised by the RootChain contract.
type RootChainFreedBond struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFreedBond is a free log retrieval operation binding the contract event 0x5406cdbb33189b887da81a90c7ff307f3af1854b2504d0bde438b8bb5b7d8b52.
//
// Solidity: e FreedBond(from indexed address, amount uint256)
func (_RootChain *RootChainFilterer) FilterFreedBond(opts *bind.FilterOpts, from []common.Address) (*RootChainFreedBondIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "FreedBond", fromRule)
	if err != nil {
		return nil, err
	}
	return &RootChainFreedBondIterator{contract: _RootChain.contract, event: "FreedBond", logs: logs, sub: sub}, nil
}

// WatchFreedBond is a free log subscription operation binding the contract event 0x5406cdbb33189b887da81a90c7ff307f3af1854b2504d0bde438b8bb5b7d8b52.
//
// Solidity: e FreedBond(from indexed address, amount uint256)
func (_RootChain *RootChainFilterer) WatchFreedBond(opts *bind.WatchOpts, sink chan<- *RootChainFreedBond, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "FreedBond", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainFreedBond)
				if err := _RootChain.contract.UnpackLog(event, "FreedBond", log); err != nil {
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

// RootChainPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the RootChain contract.
type RootChainPausedIterator struct {
	Event *RootChainPaused // Event containing the contract specifics and raw log

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
func (it *RootChainPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainPaused)
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
		it.Event = new(RootChainPaused)
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
func (it *RootChainPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainPaused represents a Paused event raised by the RootChain contract.
type RootChainPaused struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2.
//
// Solidity: e Paused(status bool)
func (_RootChain *RootChainFilterer) FilterPaused(opts *bind.FilterOpts) (*RootChainPausedIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RootChainPausedIterator{contract: _RootChain.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x0e2fb031ee032dc02d8011dc50b816eb450cf856abd8261680dac74f72165bd2.
//
// Solidity: e Paused(status bool)
func (_RootChain *RootChainFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RootChainPaused) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainPaused)
				if err := _RootChain.contract.UnpackLog(event, "Paused", log); err != nil {
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

// RootChainRespondedExitChallengeIterator is returned from FilterRespondedExitChallenge and is used to iterate over the raw logs and unpacked data for RespondedExitChallenge events raised by the RootChain contract.
type RootChainRespondedExitChallengeIterator struct {
	Event *RootChainRespondedExitChallenge // Event containing the contract specifics and raw log

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
func (it *RootChainRespondedExitChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainRespondedExitChallenge)
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
		it.Event = new(RootChainRespondedExitChallenge)
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
func (it *RootChainRespondedExitChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainRespondedExitChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainRespondedExitChallenge represents a RespondedExitChallenge event raised by the RootChain contract.
type RootChainRespondedExitChallenge struct {
	Slot uint64
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRespondedExitChallenge is a free log retrieval operation binding the contract event 0x7deeb80e1644537d3ba4f917cfcbaba62b54354e9f2598d6885d52e359885e25.
//
// Solidity: e RespondedExitChallenge(slot indexed uint64)
func (_RootChain *RootChainFilterer) FilterRespondedExitChallenge(opts *bind.FilterOpts, slot []uint64) (*RootChainRespondedExitChallengeIterator, error) {

	var slotRule []interface{}
	for _, slotItem := range slot {
		slotRule = append(slotRule, slotItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "RespondedExitChallenge", slotRule)
	if err != nil {
		return nil, err
	}
	return &RootChainRespondedExitChallengeIterator{contract: _RootChain.contract, event: "RespondedExitChallenge", logs: logs, sub: sub}, nil
}

// WatchRespondedExitChallenge is a free log subscription operation binding the contract event 0x7deeb80e1644537d3ba4f917cfcbaba62b54354e9f2598d6885d52e359885e25.
//
// Solidity: e RespondedExitChallenge(slot indexed uint64)
func (_RootChain *RootChainFilterer) WatchRespondedExitChallenge(opts *bind.WatchOpts, sink chan<- *RootChainRespondedExitChallenge, slot []uint64) (event.Subscription, error) {

	var slotRule []interface{}
	for _, slotItem := range slot {
		slotRule = append(slotRule, slotItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "RespondedExitChallenge", slotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainRespondedExitChallenge)
				if err := _RootChain.contract.UnpackLog(event, "RespondedExitChallenge", log); err != nil {
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

// RootChainSlashedBondIterator is returned from FilterSlashedBond and is used to iterate over the raw logs and unpacked data for SlashedBond events raised by the RootChain contract.
type RootChainSlashedBondIterator struct {
	Event *RootChainSlashedBond // Event containing the contract specifics and raw log

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
func (it *RootChainSlashedBondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainSlashedBond)
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
		it.Event = new(RootChainSlashedBond)
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
func (it *RootChainSlashedBondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainSlashedBondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainSlashedBond represents a SlashedBond event raised by the RootChain contract.
type RootChainSlashedBond struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSlashedBond is a free log retrieval operation binding the contract event 0xbd6bf2f753dc1a7dee94a7abd7645adc6c228c2af17fd72d9104674cc5d21273.
//
// Solidity: e SlashedBond(from indexed address, to indexed address, amount uint256)
func (_RootChain *RootChainFilterer) FilterSlashedBond(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RootChainSlashedBondIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "SlashedBond", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RootChainSlashedBondIterator{contract: _RootChain.contract, event: "SlashedBond", logs: logs, sub: sub}, nil
}

// WatchSlashedBond is a free log subscription operation binding the contract event 0xbd6bf2f753dc1a7dee94a7abd7645adc6c228c2af17fd72d9104674cc5d21273.
//
// Solidity: e SlashedBond(from indexed address, to indexed address, amount uint256)
func (_RootChain *RootChainFilterer) WatchSlashedBond(opts *bind.WatchOpts, sink chan<- *RootChainSlashedBond, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "SlashedBond", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainSlashedBond)
				if err := _RootChain.contract.UnpackLog(event, "SlashedBond", log); err != nil {
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

// RootChainStartedExitIterator is returned from FilterStartedExit and is used to iterate over the raw logs and unpacked data for StartedExit events raised by the RootChain contract.
type RootChainStartedExitIterator struct {
	Event *RootChainStartedExit // Event containing the contract specifics and raw log

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
func (it *RootChainStartedExitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainStartedExit)
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
		it.Event = new(RootChainStartedExit)
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
func (it *RootChainStartedExitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainStartedExitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainStartedExit represents a StartedExit event raised by the RootChain contract.
type RootChainStartedExit struct {
	Slot  uint64
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStartedExit is a free log retrieval operation binding the contract event 0xbaf912dedc1b0ee4647f945b0432e694bce1aa2c4e21052d9776876415874956.
//
// Solidity: e StartedExit(slot indexed uint64, owner indexed address)
func (_RootChain *RootChainFilterer) FilterStartedExit(opts *bind.FilterOpts, slot []uint64, owner []common.Address) (*RootChainStartedExitIterator, error) {

	var slotRule []interface{}
	for _, slotItem := range slot {
		slotRule = append(slotRule, slotItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "StartedExit", slotRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &RootChainStartedExitIterator{contract: _RootChain.contract, event: "StartedExit", logs: logs, sub: sub}, nil
}

// WatchStartedExit is a free log subscription operation binding the contract event 0xbaf912dedc1b0ee4647f945b0432e694bce1aa2c4e21052d9776876415874956.
//
// Solidity: e StartedExit(slot indexed uint64, owner indexed address)
func (_RootChain *RootChainFilterer) WatchStartedExit(opts *bind.WatchOpts, sink chan<- *RootChainStartedExit, slot []uint64, owner []common.Address) (event.Subscription, error) {

	var slotRule []interface{}
	for _, slotItem := range slot {
		slotRule = append(slotRule, slotItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "StartedExit", slotRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainStartedExit)
				if err := _RootChain.contract.UnpackLog(event, "StartedExit", log); err != nil {
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

// RootChainSubmittedBlockIterator is returned from FilterSubmittedBlock and is used to iterate over the raw logs and unpacked data for SubmittedBlock events raised by the RootChain contract.
type RootChainSubmittedBlockIterator struct {
	Event *RootChainSubmittedBlock // Event containing the contract specifics and raw log

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
func (it *RootChainSubmittedBlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainSubmittedBlock)
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
		it.Event = new(RootChainSubmittedBlock)
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
func (it *RootChainSubmittedBlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainSubmittedBlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainSubmittedBlock represents a SubmittedBlock event raised by the RootChain contract.
type RootChainSubmittedBlock struct {
	BlockNumber *big.Int
	Root        [32]byte
	Timestamp   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSubmittedBlock is a free log retrieval operation binding the contract event 0x6d91cd6ccac8368394df514e6aee19a55264f5ab49a891af91ca86da27bedd4f.
//
// Solidity: e SubmittedBlock(blockNumber uint256, root bytes32, timestamp uint256)
func (_RootChain *RootChainFilterer) FilterSubmittedBlock(opts *bind.FilterOpts) (*RootChainSubmittedBlockIterator, error) {

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "SubmittedBlock")
	if err != nil {
		return nil, err
	}
	return &RootChainSubmittedBlockIterator{contract: _RootChain.contract, event: "SubmittedBlock", logs: logs, sub: sub}, nil
}

// WatchSubmittedBlock is a free log subscription operation binding the contract event 0x6d91cd6ccac8368394df514e6aee19a55264f5ab49a891af91ca86da27bedd4f.
//
// Solidity: e SubmittedBlock(blockNumber uint256, root bytes32, timestamp uint256)
func (_RootChain *RootChainFilterer) WatchSubmittedBlock(opts *bind.WatchOpts, sink chan<- *RootChainSubmittedBlock) (event.Subscription, error) {

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "SubmittedBlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainSubmittedBlock)
				if err := _RootChain.contract.UnpackLog(event, "SubmittedBlock", log); err != nil {
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

// RootChainWithdrewIterator is returned from FilterWithdrew and is used to iterate over the raw logs and unpacked data for Withdrew events raised by the RootChain contract.
type RootChainWithdrewIterator struct {
	Event *RootChainWithdrew // Event containing the contract specifics and raw log

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
func (it *RootChainWithdrewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainWithdrew)
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
		it.Event = new(RootChainWithdrew)
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
func (it *RootChainWithdrewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainWithdrewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainWithdrew represents a Withdrew event raised by the RootChain contract.
type RootChainWithdrew struct {
	Owner           common.Address
	Slot            uint64
	Mode            uint8
	ContractAddress common.Address
	Uid             *big.Int
	Denomination    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterWithdrew is a free log retrieval operation binding the contract event 0xc753e3ccb1413ef7152e8d33fd0abfb140e236ce17bf9ae40a2a6940bcb7ed7b.
//
// Solidity: e Withdrew(owner indexed address, slot indexed uint64, mode uint8, contractAddress address, uid uint256, denomination uint256)
func (_RootChain *RootChainFilterer) FilterWithdrew(opts *bind.FilterOpts, owner []common.Address, slot []uint64) (*RootChainWithdrewIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var slotRule []interface{}
	for _, slotItem := range slot {
		slotRule = append(slotRule, slotItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "Withdrew", ownerRule, slotRule)
	if err != nil {
		return nil, err
	}
	return &RootChainWithdrewIterator{contract: _RootChain.contract, event: "Withdrew", logs: logs, sub: sub}, nil
}

// WatchWithdrew is a free log subscription operation binding the contract event 0xc753e3ccb1413ef7152e8d33fd0abfb140e236ce17bf9ae40a2a6940bcb7ed7b.
//
// Solidity: e Withdrew(owner indexed address, slot indexed uint64, mode uint8, contractAddress address, uid uint256, denomination uint256)
func (_RootChain *RootChainFilterer) WatchWithdrew(opts *bind.WatchOpts, sink chan<- *RootChainWithdrew, owner []common.Address, slot []uint64) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var slotRule []interface{}
	for _, slotItem := range slot {
		slotRule = append(slotRule, slotItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "Withdrew", ownerRule, slotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainWithdrew)
				if err := _RootChain.contract.UnpackLog(event, "Withdrew", log); err != nil {
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

// RootChainWithdrewBondsIterator is returned from FilterWithdrewBonds and is used to iterate over the raw logs and unpacked data for WithdrewBonds events raised by the RootChain contract.
type RootChainWithdrewBondsIterator struct {
	Event *RootChainWithdrewBonds // Event containing the contract specifics and raw log

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
func (it *RootChainWithdrewBondsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RootChainWithdrewBonds)
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
		it.Event = new(RootChainWithdrewBonds)
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
func (it *RootChainWithdrewBondsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RootChainWithdrewBondsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RootChainWithdrewBonds represents a WithdrewBonds event raised by the RootChain contract.
type RootChainWithdrewBonds struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrewBonds is a free log retrieval operation binding the contract event 0xba49a755e2e4831fb6ef34b384178b03af45679532c9ac5a2199b27b4dddea8d.
//
// Solidity: e WithdrewBonds(from indexed address, amount uint256)
func (_RootChain *RootChainFilterer) FilterWithdrewBonds(opts *bind.FilterOpts, from []common.Address) (*RootChainWithdrewBondsIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RootChain.contract.FilterLogs(opts, "WithdrewBonds", fromRule)
	if err != nil {
		return nil, err
	}
	return &RootChainWithdrewBondsIterator{contract: _RootChain.contract, event: "WithdrewBonds", logs: logs, sub: sub}, nil
}

// WatchWithdrewBonds is a free log subscription operation binding the contract event 0xba49a755e2e4831fb6ef34b384178b03af45679532c9ac5a2199b27b4dddea8d.
//
// Solidity: e WithdrewBonds(from indexed address, amount uint256)
func (_RootChain *RootChainFilterer) WatchWithdrewBonds(opts *bind.WatchOpts, sink chan<- *RootChainWithdrewBonds, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _RootChain.contract.WatchLogs(opts, "WithdrewBonds", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RootChainWithdrewBonds)
				if err := _RootChain.contract.UnpackLog(event, "WithdrewBonds", log); err != nil {
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
