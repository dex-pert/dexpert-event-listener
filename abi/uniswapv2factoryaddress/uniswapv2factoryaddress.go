// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv2factoryaddress

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
	_ = abi.ConvertType
)

// Uniswapv2factoryaddressMetaData contains all meta data concerning the Uniswapv2factoryaddress contract.
var Uniswapv2factoryaddressMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToSetter\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// Uniswapv2factoryaddressABI is the input ABI used to generate the binding from.
// Deprecated: Use Uniswapv2factoryaddressMetaData.ABI instead.
var Uniswapv2factoryaddressABI = Uniswapv2factoryaddressMetaData.ABI

// Uniswapv2factoryaddress is an auto generated Go binding around an Ethereum contract.
type Uniswapv2factoryaddress struct {
	Uniswapv2factoryaddressCaller     // Read-only binding to the contract
	Uniswapv2factoryaddressTransactor // Write-only binding to the contract
	Uniswapv2factoryaddressFilterer   // Log filterer for contract events
}

// Uniswapv2factoryaddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type Uniswapv2factoryaddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2factoryaddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Uniswapv2factoryaddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2factoryaddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Uniswapv2factoryaddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2factoryaddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Uniswapv2factoryaddressSession struct {
	Contract     *Uniswapv2factoryaddress // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// Uniswapv2factoryaddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Uniswapv2factoryaddressCallerSession struct {
	Contract *Uniswapv2factoryaddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// Uniswapv2factoryaddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Uniswapv2factoryaddressTransactorSession struct {
	Contract     *Uniswapv2factoryaddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// Uniswapv2factoryaddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type Uniswapv2factoryaddressRaw struct {
	Contract *Uniswapv2factoryaddress // Generic contract binding to access the raw methods on
}

// Uniswapv2factoryaddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Uniswapv2factoryaddressCallerRaw struct {
	Contract *Uniswapv2factoryaddressCaller // Generic read-only contract binding to access the raw methods on
}

// Uniswapv2factoryaddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Uniswapv2factoryaddressTransactorRaw struct {
	Contract *Uniswapv2factoryaddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapv2factoryaddress creates a new instance of Uniswapv2factoryaddress, bound to a specific deployed contract.
func NewUniswapv2factoryaddress(address common.Address, backend bind.ContractBackend) (*Uniswapv2factoryaddress, error) {
	contract, err := bindUniswapv2factoryaddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2factoryaddress{Uniswapv2factoryaddressCaller: Uniswapv2factoryaddressCaller{contract: contract}, Uniswapv2factoryaddressTransactor: Uniswapv2factoryaddressTransactor{contract: contract}, Uniswapv2factoryaddressFilterer: Uniswapv2factoryaddressFilterer{contract: contract}}, nil
}

// NewUniswapv2factoryaddressCaller creates a new read-only instance of Uniswapv2factoryaddress, bound to a specific deployed contract.
func NewUniswapv2factoryaddressCaller(address common.Address, caller bind.ContractCaller) (*Uniswapv2factoryaddressCaller, error) {
	contract, err := bindUniswapv2factoryaddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2factoryaddressCaller{contract: contract}, nil
}

// NewUniswapv2factoryaddressTransactor creates a new write-only instance of Uniswapv2factoryaddress, bound to a specific deployed contract.
func NewUniswapv2factoryaddressTransactor(address common.Address, transactor bind.ContractTransactor) (*Uniswapv2factoryaddressTransactor, error) {
	contract, err := bindUniswapv2factoryaddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2factoryaddressTransactor{contract: contract}, nil
}

// NewUniswapv2factoryaddressFilterer creates a new log filterer instance of Uniswapv2factoryaddress, bound to a specific deployed contract.
func NewUniswapv2factoryaddressFilterer(address common.Address, filterer bind.ContractFilterer) (*Uniswapv2factoryaddressFilterer, error) {
	contract, err := bindUniswapv2factoryaddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2factoryaddressFilterer{contract: contract}, nil
}

// bindUniswapv2factoryaddress binds a generic wrapper to an already deployed contract.
func bindUniswapv2factoryaddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Uniswapv2factoryaddressMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv2factoryaddress.Contract.Uniswapv2factoryaddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv2factoryaddress.Contract.Uniswapv2factoryaddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv2factoryaddress.Contract.Uniswapv2factoryaddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv2factoryaddress.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv2factoryaddress.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv2factoryaddress.Contract.contract.Transact(opts, method, params...)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv2factoryaddress.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Uniswapv2factoryaddress.Contract.AllPairs(&_Uniswapv2factoryaddress.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Uniswapv2factoryaddress.Contract.AllPairs(&_Uniswapv2factoryaddress.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Uniswapv2factoryaddress.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressSession) AllPairsLength() (*big.Int, error) {
	return _Uniswapv2factoryaddress.Contract.AllPairsLength(&_Uniswapv2factoryaddress.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressCallerSession) AllPairsLength() (*big.Int, error) {
	return _Uniswapv2factoryaddress.Contract.AllPairsLength(&_Uniswapv2factoryaddress.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv2factoryaddress.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressSession) FeeTo() (common.Address, error) {
	return _Uniswapv2factoryaddress.Contract.FeeTo(&_Uniswapv2factoryaddress.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressCallerSession) FeeTo() (common.Address, error) {
	return _Uniswapv2factoryaddress.Contract.FeeTo(&_Uniswapv2factoryaddress.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv2factoryaddress.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressSession) FeeToSetter() (common.Address, error) {
	return _Uniswapv2factoryaddress.Contract.FeeToSetter(&_Uniswapv2factoryaddress.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressCallerSession) FeeToSetter() (common.Address, error) {
	return _Uniswapv2factoryaddress.Contract.FeeToSetter(&_Uniswapv2factoryaddress.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressCaller) GetPair(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv2factoryaddress.contract.Call(opts, &out, "getPair", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _Uniswapv2factoryaddress.Contract.GetPair(&_Uniswapv2factoryaddress.CallOpts, arg0, arg1)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address , address ) view returns(address)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressCallerSession) GetPair(arg0 common.Address, arg1 common.Address) (common.Address, error) {
	return _Uniswapv2factoryaddress.Contract.GetPair(&_Uniswapv2factoryaddress.CallOpts, arg0, arg1)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Uniswapv2factoryaddress.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Uniswapv2factoryaddress.Contract.CreatePair(&_Uniswapv2factoryaddress.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Uniswapv2factoryaddress.Contract.CreatePair(&_Uniswapv2factoryaddress.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressTransactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _Uniswapv2factoryaddress.contract.Transact(opts, "setFeeTo", _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _Uniswapv2factoryaddress.Contract.SetFeeTo(&_Uniswapv2factoryaddress.TransactOpts, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressTransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _Uniswapv2factoryaddress.Contract.SetFeeTo(&_Uniswapv2factoryaddress.TransactOpts, _feeTo)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressTransactor) SetFeeToSetter(opts *bind.TransactOpts, _feeToSetter common.Address) (*types.Transaction, error) {
	return _Uniswapv2factoryaddress.contract.Transact(opts, "setFeeToSetter", _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _Uniswapv2factoryaddress.Contract.SetFeeToSetter(&_Uniswapv2factoryaddress.TransactOpts, _feeToSetter)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address _feeToSetter) returns()
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressTransactorSession) SetFeeToSetter(_feeToSetter common.Address) (*types.Transaction, error) {
	return _Uniswapv2factoryaddress.Contract.SetFeeToSetter(&_Uniswapv2factoryaddress.TransactOpts, _feeToSetter)
}

// Uniswapv2factoryaddressPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the Uniswapv2factoryaddress contract.
type Uniswapv2factoryaddressPairCreatedIterator struct {
	Event *Uniswapv2factoryaddressPairCreated // Event containing the contract specifics and raw log

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
func (it *Uniswapv2factoryaddressPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Uniswapv2factoryaddressPairCreated)
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
		it.Event = new(Uniswapv2factoryaddressPairCreated)
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
func (it *Uniswapv2factoryaddressPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Uniswapv2factoryaddressPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Uniswapv2factoryaddressPairCreated represents a PairCreated event raised by the Uniswapv2factoryaddress contract.
type Uniswapv2factoryaddressPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*Uniswapv2factoryaddressPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Uniswapv2factoryaddress.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2factoryaddressPairCreatedIterator{contract: _Uniswapv2factoryaddress.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *Uniswapv2factoryaddressPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Uniswapv2factoryaddress.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Uniswapv2factoryaddressPairCreated)
				if err := _Uniswapv2factoryaddress.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Uniswapv2factoryaddress *Uniswapv2factoryaddressFilterer) ParsePairCreated(log types.Log) (*Uniswapv2factoryaddressPairCreated, error) {
	event := new(Uniswapv2factoryaddressPairCreated)
	if err := _Uniswapv2factoryaddress.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
