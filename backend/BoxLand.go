// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package boxland

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

// BoxlandMetaData contains all meta data concerning the Boxland contract.
var BoxlandMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"TokenTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MaxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Supply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// BoxlandABI is the input ABI used to generate the binding from.
// Deprecated: Use BoxlandMetaData.ABI instead.
var BoxlandABI = BoxlandMetaData.ABI

// Boxland is an auto generated Go binding around an Ethereum contract.
type Boxland struct {
	BoxlandCaller     // Read-only binding to the contract
	BoxlandTransactor // Write-only binding to the contract
	BoxlandFilterer   // Log filterer for contract events
}

// BoxlandCaller is an auto generated read-only Go binding around an Ethereum contract.
type BoxlandCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoxlandTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BoxlandTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoxlandFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BoxlandFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BoxlandSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BoxlandSession struct {
	Contract     *Boxland          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BoxlandCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BoxlandCallerSession struct {
	Contract *BoxlandCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BoxlandTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BoxlandTransactorSession struct {
	Contract     *BoxlandTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BoxlandRaw is an auto generated low-level Go binding around an Ethereum contract.
type BoxlandRaw struct {
	Contract *Boxland // Generic contract binding to access the raw methods on
}

// BoxlandCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BoxlandCallerRaw struct {
	Contract *BoxlandCaller // Generic read-only contract binding to access the raw methods on
}

// BoxlandTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BoxlandTransactorRaw struct {
	Contract *BoxlandTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBoxland creates a new instance of Boxland, bound to a specific deployed contract.
func NewBoxland(address common.Address, backend bind.ContractBackend) (*Boxland, error) {
	contract, err := bindBoxland(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Boxland{BoxlandCaller: BoxlandCaller{contract: contract}, BoxlandTransactor: BoxlandTransactor{contract: contract}, BoxlandFilterer: BoxlandFilterer{contract: contract}}, nil
}

// NewBoxlandCaller creates a new read-only instance of Boxland, bound to a specific deployed contract.
func NewBoxlandCaller(address common.Address, caller bind.ContractCaller) (*BoxlandCaller, error) {
	contract, err := bindBoxland(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BoxlandCaller{contract: contract}, nil
}

// NewBoxlandTransactor creates a new write-only instance of Boxland, bound to a specific deployed contract.
func NewBoxlandTransactor(address common.Address, transactor bind.ContractTransactor) (*BoxlandTransactor, error) {
	contract, err := bindBoxland(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BoxlandTransactor{contract: contract}, nil
}

// NewBoxlandFilterer creates a new log filterer instance of Boxland, bound to a specific deployed contract.
func NewBoxlandFilterer(address common.Address, filterer bind.ContractFilterer) (*BoxlandFilterer, error) {
	contract, err := bindBoxland(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BoxlandFilterer{contract: contract}, nil
}

// bindBoxland binds a generic wrapper to an already deployed contract.
func bindBoxland(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BoxlandABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Boxland *BoxlandRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Boxland.Contract.BoxlandCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Boxland *BoxlandRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Boxland.Contract.BoxlandTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Boxland *BoxlandRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Boxland.Contract.BoxlandTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Boxland *BoxlandCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Boxland.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Boxland *BoxlandTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Boxland.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Boxland *BoxlandTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Boxland.Contract.contract.Transact(opts, method, params...)
}

// MaxSupply is a free data retrieval call binding the contract method 0xb36c1284.
//
// Solidity: function MaxSupply() view returns(uint256)
func (_Boxland *BoxlandCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Boxland.contract.Call(opts, &out, "MaxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xb36c1284.
//
// Solidity: function MaxSupply() view returns(uint256)
func (_Boxland *BoxlandSession) MaxSupply() (*big.Int, error) {
	return _Boxland.Contract.MaxSupply(&_Boxland.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xb36c1284.
//
// Solidity: function MaxSupply() view returns(uint256)
func (_Boxland *BoxlandCallerSession) MaxSupply() (*big.Int, error) {
	return _Boxland.Contract.MaxSupply(&_Boxland.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0xee07bf22.
//
// Solidity: function Supply() view returns(uint256)
func (_Boxland *BoxlandCaller) Supply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Boxland.contract.Call(opts, &out, "Supply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Supply is a free data retrieval call binding the contract method 0xee07bf22.
//
// Solidity: function Supply() view returns(uint256)
func (_Boxland *BoxlandSession) Supply() (*big.Int, error) {
	return _Boxland.Contract.Supply(&_Boxland.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0xee07bf22.
//
// Solidity: function Supply() view returns(uint256)
func (_Boxland *BoxlandCallerSession) Supply() (*big.Int, error) {
	return _Boxland.Contract.Supply(&_Boxland.CallOpts)
}

// Verify is a free data retrieval call binding the contract method 0x21fb335c.
//
// Solidity: function verify(bytes32[] proof, bytes32 root, bytes32 leaf, uint256 index) pure returns(bool)
func (_Boxland *BoxlandCaller) Verify(opts *bind.CallOpts, proof [][32]byte, root [32]byte, leaf [32]byte, index *big.Int) (bool, error) {
	var out []interface{}
	err := _Boxland.contract.Call(opts, &out, "verify", proof, root, leaf, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x21fb335c.
//
// Solidity: function verify(bytes32[] proof, bytes32 root, bytes32 leaf, uint256 index) pure returns(bool)
func (_Boxland *BoxlandSession) Verify(proof [][32]byte, root [32]byte, leaf [32]byte, index *big.Int) (bool, error) {
	return _Boxland.Contract.Verify(&_Boxland.CallOpts, proof, root, leaf, index)
}

// Verify is a free data retrieval call binding the contract method 0x21fb335c.
//
// Solidity: function verify(bytes32[] proof, bytes32 root, bytes32 leaf, uint256 index) pure returns(bool)
func (_Boxland *BoxlandCallerSession) Verify(proof [][32]byte, root [32]byte, leaf [32]byte, index *big.Int) (bool, error) {
	return _Boxland.Contract.Verify(&_Boxland.CallOpts, proof, root, leaf, index)
}

// AddToken is a paid mutator transaction binding the contract method 0x179d375c.
//
// Solidity: function addToken(uint256 price) payable returns()
func (_Boxland *BoxlandTransactor) AddToken(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _Boxland.contract.Transact(opts, "addToken", price)
}

// AddToken is a paid mutator transaction binding the contract method 0x179d375c.
//
// Solidity: function addToken(uint256 price) payable returns()
func (_Boxland *BoxlandSession) AddToken(price *big.Int) (*types.Transaction, error) {
	return _Boxland.Contract.AddToken(&_Boxland.TransactOpts, price)
}

// AddToken is a paid mutator transaction binding the contract method 0x179d375c.
//
// Solidity: function addToken(uint256 price) payable returns()
func (_Boxland *BoxlandTransactorSession) AddToken(price *big.Int) (*types.Transaction, error) {
	return _Boxland.Contract.AddToken(&_Boxland.TransactOpts, price)
}

// BuyToken is a paid mutator transaction binding the contract method 0x057466ea.
//
// Solidity: function buyToken(uint256 id, uint256 price) payable returns()
func (_Boxland *BoxlandTransactor) BuyToken(opts *bind.TransactOpts, id *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Boxland.contract.Transact(opts, "buyToken", id, price)
}

// BuyToken is a paid mutator transaction binding the contract method 0x057466ea.
//
// Solidity: function buyToken(uint256 id, uint256 price) payable returns()
func (_Boxland *BoxlandSession) BuyToken(id *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Boxland.Contract.BuyToken(&_Boxland.TransactOpts, id, price)
}

// BuyToken is a paid mutator transaction binding the contract method 0x057466ea.
//
// Solidity: function buyToken(uint256 id, uint256 price) payable returns()
func (_Boxland *BoxlandTransactorSession) BuyToken(id *big.Int, price *big.Int) (*types.Transaction, error) {
	return _Boxland.Contract.BuyToken(&_Boxland.TransactOpts, id, price)
}

// Transfer is a paid mutator transaction binding the contract method 0x8a4068dd.
//
// Solidity: function transfer() payable returns()
func (_Boxland *BoxlandTransactor) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Boxland.contract.Transact(opts, "transfer")
}

// Transfer is a paid mutator transaction binding the contract method 0x8a4068dd.
//
// Solidity: function transfer() payable returns()
func (_Boxland *BoxlandSession) Transfer() (*types.Transaction, error) {
	return _Boxland.Contract.Transfer(&_Boxland.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x8a4068dd.
//
// Solidity: function transfer() payable returns()
func (_Boxland *BoxlandTransactorSession) Transfer() (*types.Transaction, error) {
	return _Boxland.Contract.Transfer(&_Boxland.TransactOpts)
}

// BoxlandTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the Boxland contract.
type BoxlandTokenAddedIterator struct {
	Event *BoxlandTokenAdded // Event containing the contract specifics and raw log

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
func (it *BoxlandTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoxlandTokenAdded)
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
		it.Event = new(BoxlandTokenAdded)
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
func (it *BoxlandTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoxlandTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoxlandTokenAdded represents a TokenAdded event raised by the Boxland contract.
type BoxlandTokenAdded struct {
	Id    *big.Int
	Owner common.Address
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x8eebdda427f4613147238037e21b19a9094003ef7d1ca1185ac0ebf8d5c18bd1.
//
// Solidity: event TokenAdded(uint256 id, address owner, uint256 price)
func (_Boxland *BoxlandFilterer) FilterTokenAdded(opts *bind.FilterOpts) (*BoxlandTokenAddedIterator, error) {

	logs, sub, err := _Boxland.contract.FilterLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return &BoxlandTokenAddedIterator{contract: _Boxland.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x8eebdda427f4613147238037e21b19a9094003ef7d1ca1185ac0ebf8d5c18bd1.
//
// Solidity: event TokenAdded(uint256 id, address owner, uint256 price)
func (_Boxland *BoxlandFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *BoxlandTokenAdded) (event.Subscription, error) {

	logs, sub, err := _Boxland.contract.WatchLogs(opts, "TokenAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoxlandTokenAdded)
				if err := _Boxland.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0x8eebdda427f4613147238037e21b19a9094003ef7d1ca1185ac0ebf8d5c18bd1.
//
// Solidity: event TokenAdded(uint256 id, address owner, uint256 price)
func (_Boxland *BoxlandFilterer) ParseTokenAdded(log types.Log) (*BoxlandTokenAdded, error) {
	event := new(BoxlandTokenAdded)
	if err := _Boxland.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BoxlandTokenTransferredIterator is returned from FilterTokenTransferred and is used to iterate over the raw logs and unpacked data for TokenTransferred events raised by the Boxland contract.
type BoxlandTokenTransferredIterator struct {
	Event *BoxlandTokenTransferred // Event containing the contract specifics and raw log

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
func (it *BoxlandTokenTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BoxlandTokenTransferred)
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
		it.Event = new(BoxlandTokenTransferred)
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
func (it *BoxlandTokenTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BoxlandTokenTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BoxlandTokenTransferred represents a TokenTransferred event raised by the Boxland contract.
type BoxlandTokenTransferred struct {
	Id       *big.Int
	NewOwner common.Address
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenTransferred is a free log retrieval operation binding the contract event 0x60826d98c92424c12bac5dc944e3c0d83cd739c00e5ed7dc5baf3031098b5a7f.
//
// Solidity: event TokenTransferred(uint256 id, address newOwner, uint256 price)
func (_Boxland *BoxlandFilterer) FilterTokenTransferred(opts *bind.FilterOpts) (*BoxlandTokenTransferredIterator, error) {

	logs, sub, err := _Boxland.contract.FilterLogs(opts, "TokenTransferred")
	if err != nil {
		return nil, err
	}
	return &BoxlandTokenTransferredIterator{contract: _Boxland.contract, event: "TokenTransferred", logs: logs, sub: sub}, nil
}

// WatchTokenTransferred is a free log subscription operation binding the contract event 0x60826d98c92424c12bac5dc944e3c0d83cd739c00e5ed7dc5baf3031098b5a7f.
//
// Solidity: event TokenTransferred(uint256 id, address newOwner, uint256 price)
func (_Boxland *BoxlandFilterer) WatchTokenTransferred(opts *bind.WatchOpts, sink chan<- *BoxlandTokenTransferred) (event.Subscription, error) {

	logs, sub, err := _Boxland.contract.WatchLogs(opts, "TokenTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BoxlandTokenTransferred)
				if err := _Boxland.contract.UnpackLog(event, "TokenTransferred", log); err != nil {
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

// ParseTokenTransferred is a log parse operation binding the contract event 0x60826d98c92424c12bac5dc944e3c0d83cd739c00e5ed7dc5baf3031098b5a7f.
//
// Solidity: event TokenTransferred(uint256 id, address newOwner, uint256 price)
func (_Boxland *BoxlandFilterer) ParseTokenTransferred(log types.Log) (*BoxlandTokenTransferred, error) {
	event := new(BoxlandTokenTransferred)
	if err := _Boxland.contract.UnpackLog(event, "TokenTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
