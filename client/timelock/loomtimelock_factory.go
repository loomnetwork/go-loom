// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package timelock

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

// LoomTimelockFactoryABI is the input ABI used to generate the binding from.
const LoomTimelockFactoryABI = "[{\"inputs\":[{\"name\":\"_loom\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"signature\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"validatorEthAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"timelockContractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"validatorName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"validatorPublicKey\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_releaseTime\",\"type\":\"uint256\"}],\"name\":\"LoomTimeLockCreated\",\"type\":\"event\",\"signature\":\"0xc54292f6524435faba788a1c757e3ced79c3a3a6e1d2bee3b13ee8d12d686123\"},{\"constant\":false,\"inputs\":[{\"name\":\"validatorEthAddress\",\"type\":\"address\"},{\"name\":\"validatorName\",\"type\":\"string\"},{\"name\":\"validatorPublicKey\",\"type\":\"string\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"deployTimeLock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"signature\":\"0xee66ca2b\"}]"

// LoomTimelockFactory is an auto generated Go binding around an Ethereum contract.
type LoomTimelockFactory struct {
	LoomTimelockFactoryCaller     // Read-only binding to the contract
	LoomTimelockFactoryTransactor // Write-only binding to the contract
	LoomTimelockFactoryFilterer   // Log filterer for contract events
}

// LoomTimelockFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoomTimelockFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoomTimelockFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoomTimelockFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoomTimelockFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoomTimelockFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoomTimelockFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoomTimelockFactorySession struct {
	Contract     *LoomTimelockFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LoomTimelockFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoomTimelockFactoryCallerSession struct {
	Contract *LoomTimelockFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// LoomTimelockFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoomTimelockFactoryTransactorSession struct {
	Contract     *LoomTimelockFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// LoomTimelockFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoomTimelockFactoryRaw struct {
	Contract *LoomTimelockFactory // Generic contract binding to access the raw methods on
}

// LoomTimelockFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoomTimelockFactoryCallerRaw struct {
	Contract *LoomTimelockFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// LoomTimelockFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoomTimelockFactoryTransactorRaw struct {
	Contract *LoomTimelockFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLoomTimelockFactory creates a new instance of LoomTimelockFactory, bound to a specific deployed contract.
func NewLoomTimelockFactory(address common.Address, backend bind.ContractBackend) (*LoomTimelockFactory, error) {
	contract, err := bindLoomTimelockFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LoomTimelockFactory{LoomTimelockFactoryCaller: LoomTimelockFactoryCaller{contract: contract}, LoomTimelockFactoryTransactor: LoomTimelockFactoryTransactor{contract: contract}, LoomTimelockFactoryFilterer: LoomTimelockFactoryFilterer{contract: contract}}, nil
}

// NewLoomTimelockFactoryCaller creates a new read-only instance of LoomTimelockFactory, bound to a specific deployed contract.
func NewLoomTimelockFactoryCaller(address common.Address, caller bind.ContractCaller) (*LoomTimelockFactoryCaller, error) {
	contract, err := bindLoomTimelockFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoomTimelockFactoryCaller{contract: contract}, nil
}

// NewLoomTimelockFactoryTransactor creates a new write-only instance of LoomTimelockFactory, bound to a specific deployed contract.
func NewLoomTimelockFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*LoomTimelockFactoryTransactor, error) {
	contract, err := bindLoomTimelockFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoomTimelockFactoryTransactor{contract: contract}, nil
}

// NewLoomTimelockFactoryFilterer creates a new log filterer instance of LoomTimelockFactory, bound to a specific deployed contract.
func NewLoomTimelockFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*LoomTimelockFactoryFilterer, error) {
	contract, err := bindLoomTimelockFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoomTimelockFactoryFilterer{contract: contract}, nil
}

// bindLoomTimelockFactory binds a generic wrapper to an already deployed contract.
func bindLoomTimelockFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LoomTimelockFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoomTimelockFactory *LoomTimelockFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LoomTimelockFactory.Contract.LoomTimelockFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoomTimelockFactory *LoomTimelockFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoomTimelockFactory.Contract.LoomTimelockFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoomTimelockFactory *LoomTimelockFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoomTimelockFactory.Contract.LoomTimelockFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoomTimelockFactory *LoomTimelockFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LoomTimelockFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoomTimelockFactory *LoomTimelockFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoomTimelockFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoomTimelockFactory *LoomTimelockFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoomTimelockFactory.Contract.contract.Transact(opts, method, params...)
}

// DeployTimeLock is a paid mutator transaction binding the contract method 0xee66ca2b.
//
// Solidity: function deployTimeLock(validatorEthAddress address, validatorName string, validatorPublicKey string, amount uint256, duration uint256) returns()
func (_LoomTimelockFactory *LoomTimelockFactoryTransactor) DeployTimeLock(opts *bind.TransactOpts, validatorEthAddress common.Address, validatorName string, validatorPublicKey string, amount *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _LoomTimelockFactory.contract.Transact(opts, "deployTimeLock", validatorEthAddress, validatorName, validatorPublicKey, amount, duration)
}

// DeployTimeLock is a paid mutator transaction binding the contract method 0xee66ca2b.
//
// Solidity: function deployTimeLock(validatorEthAddress address, validatorName string, validatorPublicKey string, amount uint256, duration uint256) returns()
func (_LoomTimelockFactory *LoomTimelockFactorySession) DeployTimeLock(validatorEthAddress common.Address, validatorName string, validatorPublicKey string, amount *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _LoomTimelockFactory.Contract.DeployTimeLock(&_LoomTimelockFactory.TransactOpts, validatorEthAddress, validatorName, validatorPublicKey, amount, duration)
}

// DeployTimeLock is a paid mutator transaction binding the contract method 0xee66ca2b.
//
// Solidity: function deployTimeLock(validatorEthAddress address, validatorName string, validatorPublicKey string, amount uint256, duration uint256) returns()
func (_LoomTimelockFactory *LoomTimelockFactoryTransactorSession) DeployTimeLock(validatorEthAddress common.Address, validatorName string, validatorPublicKey string, amount *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _LoomTimelockFactory.Contract.DeployTimeLock(&_LoomTimelockFactory.TransactOpts, validatorEthAddress, validatorName, validatorPublicKey, amount, duration)
}

// LoomTimelockFactoryLoomTimeLockCreatedIterator is returned from FilterLoomTimeLockCreated and is used to iterate over the raw logs and unpacked data for LoomTimeLockCreated events raised by the LoomTimelockFactory contract.
type LoomTimelockFactoryLoomTimeLockCreatedIterator struct {
	Event *LoomTimelockFactoryLoomTimeLockCreated // Event containing the contract specifics and raw log

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
func (it *LoomTimelockFactoryLoomTimeLockCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoomTimelockFactoryLoomTimeLockCreated)
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
		it.Event = new(LoomTimelockFactoryLoomTimeLockCreated)
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
func (it *LoomTimelockFactoryLoomTimeLockCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoomTimelockFactoryLoomTimeLockCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoomTimelockFactoryLoomTimeLockCreated represents a LoomTimeLockCreated event raised by the LoomTimelockFactory contract.
type LoomTimelockFactoryLoomTimeLockCreated struct {
	ValidatorEthAddress     common.Address
	TimelockContractAddress common.Address
	ValidatorName           string
	ValidatorPublicKey      string
	Amount                  *big.Int
	ReleaseTime             *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterLoomTimeLockCreated is a free log retrieval operation binding the contract event 0xc54292f6524435faba788a1c757e3ced79c3a3a6e1d2bee3b13ee8d12d686123.
//
// Solidity: e LoomTimeLockCreated(validatorEthAddress address, timelockContractAddress address, validatorName string, validatorPublicKey string, _amount uint256, _releaseTime uint256)
func (_LoomTimelockFactory *LoomTimelockFactoryFilterer) FilterLoomTimeLockCreated(opts *bind.FilterOpts) (*LoomTimelockFactoryLoomTimeLockCreatedIterator, error) {

	logs, sub, err := _LoomTimelockFactory.contract.FilterLogs(opts, "LoomTimeLockCreated")
	if err != nil {
		return nil, err
	}
	return &LoomTimelockFactoryLoomTimeLockCreatedIterator{contract: _LoomTimelockFactory.contract, event: "LoomTimeLockCreated", logs: logs, sub: sub}, nil
}

// WatchLoomTimeLockCreated is a free log subscription operation binding the contract event 0xc54292f6524435faba788a1c757e3ced79c3a3a6e1d2bee3b13ee8d12d686123.
//
// Solidity: e LoomTimeLockCreated(validatorEthAddress address, timelockContractAddress address, validatorName string, validatorPublicKey string, _amount uint256, _releaseTime uint256)
func (_LoomTimelockFactory *LoomTimelockFactoryFilterer) WatchLoomTimeLockCreated(opts *bind.WatchOpts, sink chan<- *LoomTimelockFactoryLoomTimeLockCreated) (event.Subscription, error) {

	logs, sub, err := _LoomTimelockFactory.contract.WatchLogs(opts, "LoomTimeLockCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoomTimelockFactoryLoomTimeLockCreated)
				if err := _LoomTimelockFactory.contract.UnpackLog(event, "LoomTimeLockCreated", log); err != nil {
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
