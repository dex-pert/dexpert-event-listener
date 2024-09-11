// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package universalrouter

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

// RouterParameters is an auto generated low-level Go binding around an user-defined struct.
type RouterParameters struct {
	UniswapV2Router02 common.Address
	FeeRecipient      common.Address
	FeeBaseBps        *big.Int
	Permit2           common.Address
	Weth9             common.Address
	V2Factory         common.Address
	V3Factory         common.Address
	PairInitCodeHash  [32]byte
	PoolInitCodeHash  [32]byte
}

// UniversalrouterMetaData contains all meta data concerning the Universalrouter contract.
var UniversalrouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"uniswapV2Router02\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBaseBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"permit2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth9\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v2Factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v3Factory\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"pairInitCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"poolInitCodeHash\",\"type\":\"bytes32\"}],\"internalType\":\"structRouterParameters\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BalanceTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BuyPunkFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContractLocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHNotAccepted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commandIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"ExecutionFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientAddressCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientAddressCannotBeZeroAddress1\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FromAddressIsNotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBips\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commandType\",\"type\":\"uint256\"}],\"name\":\"InvalidCommandType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeBaseBps\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeBaseBps\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeBps\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReserves\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SliceOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionDeadlinePassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsafeCast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2InvalidPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2TooLittleReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2TooMuchRequested\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3InvalidAmountOut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3InvalidSwap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3TooLittleReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3TooMuchRequested\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBaseBps\",\"type\":\"uint256\"}],\"name\":\"FeeBaseBpsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"}],\"name\":\"FeeBpsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBaseBps\",\"type\":\"uint256\"}],\"name\":\"Payment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"}],\"name\":\"Permit2Error\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBaseBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapType\",\"type\":\"uint256\"}],\"name\":\"feeBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeBaseBps\",\"type\":\"uint256\"}],\"name\":\"setFeeBaseBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"}],\"name\":\"setFeeBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// UniversalrouterABI is the input ABI used to generate the binding from.
// Deprecated: Use UniversalrouterMetaData.ABI instead.
var UniversalrouterABI = UniversalrouterMetaData.ABI

// Universalrouter is an auto generated Go binding around an Ethereum contract.
type Universalrouter struct {
	UniversalrouterCaller     // Read-only binding to the contract
	UniversalrouterTransactor // Write-only binding to the contract
	UniversalrouterFilterer   // Log filterer for contract events
}

// UniversalrouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniversalrouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalrouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniversalrouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalrouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniversalrouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniversalrouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniversalrouterSession struct {
	Contract     *Universalrouter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniversalrouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniversalrouterCallerSession struct {
	Contract *UniversalrouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UniversalrouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniversalrouterTransactorSession struct {
	Contract     *UniversalrouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UniversalrouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniversalrouterRaw struct {
	Contract *Universalrouter // Generic contract binding to access the raw methods on
}

// UniversalrouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniversalrouterCallerRaw struct {
	Contract *UniversalrouterCaller // Generic read-only contract binding to access the raw methods on
}

// UniversalrouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniversalrouterTransactorRaw struct {
	Contract *UniversalrouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniversalrouter creates a new instance of Universalrouter, bound to a specific deployed contract.
func NewUniversalrouter(address common.Address, backend bind.ContractBackend) (*Universalrouter, error) {
	contract, err := bindUniversalrouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Universalrouter{UniversalrouterCaller: UniversalrouterCaller{contract: contract}, UniversalrouterTransactor: UniversalrouterTransactor{contract: contract}, UniversalrouterFilterer: UniversalrouterFilterer{contract: contract}}, nil
}

// NewUniversalrouterCaller creates a new read-only instance of Universalrouter, bound to a specific deployed contract.
func NewUniversalrouterCaller(address common.Address, caller bind.ContractCaller) (*UniversalrouterCaller, error) {
	contract, err := bindUniversalrouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalrouterCaller{contract: contract}, nil
}

// NewUniversalrouterTransactor creates a new write-only instance of Universalrouter, bound to a specific deployed contract.
func NewUniversalrouterTransactor(address common.Address, transactor bind.ContractTransactor) (*UniversalrouterTransactor, error) {
	contract, err := bindUniversalrouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniversalrouterTransactor{contract: contract}, nil
}

// NewUniversalrouterFilterer creates a new log filterer instance of Universalrouter, bound to a specific deployed contract.
func NewUniversalrouterFilterer(address common.Address, filterer bind.ContractFilterer) (*UniversalrouterFilterer, error) {
	contract, err := bindUniversalrouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniversalrouterFilterer{contract: contract}, nil
}

// bindUniversalrouter binds a generic wrapper to an already deployed contract.
func bindUniversalrouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniversalrouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Universalrouter *UniversalrouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Universalrouter.Contract.UniversalrouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Universalrouter *UniversalrouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universalrouter.Contract.UniversalrouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Universalrouter *UniversalrouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Universalrouter.Contract.UniversalrouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Universalrouter *UniversalrouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Universalrouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Universalrouter *UniversalrouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universalrouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Universalrouter *UniversalrouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Universalrouter.Contract.contract.Transact(opts, method, params...)
}

// FeeBaseBps is a free data retrieval call binding the contract method 0x8ffb274f.
//
// Solidity: function feeBaseBps() view returns(uint256)
func (_Universalrouter *UniversalrouterCaller) FeeBaseBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Universalrouter.contract.Call(opts, &out, "feeBaseBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeBaseBps is a free data retrieval call binding the contract method 0x8ffb274f.
//
// Solidity: function feeBaseBps() view returns(uint256)
func (_Universalrouter *UniversalrouterSession) FeeBaseBps() (*big.Int, error) {
	return _Universalrouter.Contract.FeeBaseBps(&_Universalrouter.CallOpts)
}

// FeeBaseBps is a free data retrieval call binding the contract method 0x8ffb274f.
//
// Solidity: function feeBaseBps() view returns(uint256)
func (_Universalrouter *UniversalrouterCallerSession) FeeBaseBps() (*big.Int, error) {
	return _Universalrouter.Contract.FeeBaseBps(&_Universalrouter.CallOpts)
}

// FeeBps is a free data retrieval call binding the contract method 0xd9d852af.
//
// Solidity: function feeBps(uint256 level, uint256 swapType) view returns(uint256)
func (_Universalrouter *UniversalrouterCaller) FeeBps(opts *bind.CallOpts, level *big.Int, swapType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Universalrouter.contract.Call(opts, &out, "feeBps", level, swapType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeBps is a free data retrieval call binding the contract method 0xd9d852af.
//
// Solidity: function feeBps(uint256 level, uint256 swapType) view returns(uint256)
func (_Universalrouter *UniversalrouterSession) FeeBps(level *big.Int, swapType *big.Int) (*big.Int, error) {
	return _Universalrouter.Contract.FeeBps(&_Universalrouter.CallOpts, level, swapType)
}

// FeeBps is a free data retrieval call binding the contract method 0xd9d852af.
//
// Solidity: function feeBps(uint256 level, uint256 swapType) view returns(uint256)
func (_Universalrouter *UniversalrouterCallerSession) FeeBps(level *big.Int, swapType *big.Int) (*big.Int, error) {
	return _Universalrouter.Contract.FeeBps(&_Universalrouter.CallOpts, level, swapType)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Universalrouter *UniversalrouterCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Universalrouter.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Universalrouter *UniversalrouterSession) FeeRecipient() (common.Address, error) {
	return _Universalrouter.Contract.FeeRecipient(&_Universalrouter.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Universalrouter *UniversalrouterCallerSession) FeeRecipient() (common.Address, error) {
	return _Universalrouter.Contract.FeeRecipient(&_Universalrouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Universalrouter *UniversalrouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Universalrouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Universalrouter *UniversalrouterSession) Owner() (common.Address, error) {
	return _Universalrouter.Contract.Owner(&_Universalrouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Universalrouter *UniversalrouterCallerSession) Owner() (common.Address, error) {
	return _Universalrouter.Contract.Owner(&_Universalrouter.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_Universalrouter *UniversalrouterTransactor) Execute(opts *bind.TransactOpts, commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _Universalrouter.contract.Transact(opts, "execute", commands, inputs)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_Universalrouter *UniversalrouterSession) Execute(commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _Universalrouter.Contract.Execute(&_Universalrouter.TransactOpts, commands, inputs)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_Universalrouter *UniversalrouterTransactorSession) Execute(commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _Universalrouter.Contract.Execute(&_Universalrouter.TransactOpts, commands, inputs)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_Universalrouter *UniversalrouterTransactor) Execute0(opts *bind.TransactOpts, commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _Universalrouter.contract.Transact(opts, "execute0", commands, inputs, deadline)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_Universalrouter *UniversalrouterSession) Execute0(commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _Universalrouter.Contract.Execute0(&_Universalrouter.TransactOpts, commands, inputs, deadline)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_Universalrouter *UniversalrouterTransactorSession) Execute0(commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _Universalrouter.Contract.Execute0(&_Universalrouter.TransactOpts, commands, inputs, deadline)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Universalrouter *UniversalrouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universalrouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Universalrouter *UniversalrouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _Universalrouter.Contract.RenounceOwnership(&_Universalrouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Universalrouter *UniversalrouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Universalrouter.Contract.RenounceOwnership(&_Universalrouter.TransactOpts)
}

// SetFeeBaseBps is a paid mutator transaction binding the contract method 0x086408db.
//
// Solidity: function setFeeBaseBps(uint256 feeBaseBps) returns()
func (_Universalrouter *UniversalrouterTransactor) SetFeeBaseBps(opts *bind.TransactOpts, feeBaseBps *big.Int) (*types.Transaction, error) {
	return _Universalrouter.contract.Transact(opts, "setFeeBaseBps", feeBaseBps)
}

// SetFeeBaseBps is a paid mutator transaction binding the contract method 0x086408db.
//
// Solidity: function setFeeBaseBps(uint256 feeBaseBps) returns()
func (_Universalrouter *UniversalrouterSession) SetFeeBaseBps(feeBaseBps *big.Int) (*types.Transaction, error) {
	return _Universalrouter.Contract.SetFeeBaseBps(&_Universalrouter.TransactOpts, feeBaseBps)
}

// SetFeeBaseBps is a paid mutator transaction binding the contract method 0x086408db.
//
// Solidity: function setFeeBaseBps(uint256 feeBaseBps) returns()
func (_Universalrouter *UniversalrouterTransactorSession) SetFeeBaseBps(feeBaseBps *big.Int) (*types.Transaction, error) {
	return _Universalrouter.Contract.SetFeeBaseBps(&_Universalrouter.TransactOpts, feeBaseBps)
}

// SetFeeBps is a paid mutator transaction binding the contract method 0x8aa06386.
//
// Solidity: function setFeeBps(uint256 level, uint256 swapType, uint256 feeBps) returns()
func (_Universalrouter *UniversalrouterTransactor) SetFeeBps(opts *bind.TransactOpts, level *big.Int, swapType *big.Int, feeBps *big.Int) (*types.Transaction, error) {
	return _Universalrouter.contract.Transact(opts, "setFeeBps", level, swapType, feeBps)
}

// SetFeeBps is a paid mutator transaction binding the contract method 0x8aa06386.
//
// Solidity: function setFeeBps(uint256 level, uint256 swapType, uint256 feeBps) returns()
func (_Universalrouter *UniversalrouterSession) SetFeeBps(level *big.Int, swapType *big.Int, feeBps *big.Int) (*types.Transaction, error) {
	return _Universalrouter.Contract.SetFeeBps(&_Universalrouter.TransactOpts, level, swapType, feeBps)
}

// SetFeeBps is a paid mutator transaction binding the contract method 0x8aa06386.
//
// Solidity: function setFeeBps(uint256 level, uint256 swapType, uint256 feeBps) returns()
func (_Universalrouter *UniversalrouterTransactorSession) SetFeeBps(level *big.Int, swapType *big.Int, feeBps *big.Int) (*types.Transaction, error) {
	return _Universalrouter.Contract.SetFeeBps(&_Universalrouter.TransactOpts, level, swapType, feeBps)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address feeRecipient) returns()
func (_Universalrouter *UniversalrouterTransactor) SetFeeRecipient(opts *bind.TransactOpts, feeRecipient common.Address) (*types.Transaction, error) {
	return _Universalrouter.contract.Transact(opts, "setFeeRecipient", feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address feeRecipient) returns()
func (_Universalrouter *UniversalrouterSession) SetFeeRecipient(feeRecipient common.Address) (*types.Transaction, error) {
	return _Universalrouter.Contract.SetFeeRecipient(&_Universalrouter.TransactOpts, feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address feeRecipient) returns()
func (_Universalrouter *UniversalrouterTransactorSession) SetFeeRecipient(feeRecipient common.Address) (*types.Transaction, error) {
	return _Universalrouter.Contract.SetFeeRecipient(&_Universalrouter.TransactOpts, feeRecipient)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Universalrouter *UniversalrouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Universalrouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Universalrouter *UniversalrouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Universalrouter.Contract.TransferOwnership(&_Universalrouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Universalrouter *UniversalrouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Universalrouter.Contract.TransferOwnership(&_Universalrouter.TransactOpts, newOwner)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_Universalrouter *UniversalrouterTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _Universalrouter.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_Universalrouter *UniversalrouterSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _Universalrouter.Contract.UniswapV3SwapCallback(&_Universalrouter.TransactOpts, amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_Universalrouter *UniversalrouterTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _Universalrouter.Contract.UniswapV3SwapCallback(&_Universalrouter.TransactOpts, amount0Delta, amount1Delta, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Universalrouter *UniversalrouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Universalrouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Universalrouter *UniversalrouterSession) Receive() (*types.Transaction, error) {
	return _Universalrouter.Contract.Receive(&_Universalrouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Universalrouter *UniversalrouterTransactorSession) Receive() (*types.Transaction, error) {
	return _Universalrouter.Contract.Receive(&_Universalrouter.TransactOpts)
}

// UniversalrouterFeeBaseBpsUpdatedIterator is returned from FilterFeeBaseBpsUpdated and is used to iterate over the raw logs and unpacked data for FeeBaseBpsUpdated events raised by the Universalrouter contract.
type UniversalrouterFeeBaseBpsUpdatedIterator struct {
	Event *UniversalrouterFeeBaseBpsUpdated // Event containing the contract specifics and raw log

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
func (it *UniversalrouterFeeBaseBpsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalrouterFeeBaseBpsUpdated)
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
		it.Event = new(UniversalrouterFeeBaseBpsUpdated)
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
func (it *UniversalrouterFeeBaseBpsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalrouterFeeBaseBpsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalrouterFeeBaseBpsUpdated represents a FeeBaseBpsUpdated event raised by the Universalrouter contract.
type UniversalrouterFeeBaseBpsUpdated struct {
	MsgSender  common.Address
	FeeBaseBps *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeBaseBpsUpdated is a free log retrieval operation binding the contract event 0xbed4dd7c9eb92506488e6892e1edc0cd8fe578bce7c3de4699223ef5ef34e297.
//
// Solidity: event FeeBaseBpsUpdated(address indexed msgSender, uint256 feeBaseBps)
func (_Universalrouter *UniversalrouterFilterer) FilterFeeBaseBpsUpdated(opts *bind.FilterOpts, msgSender []common.Address) (*UniversalrouterFeeBaseBpsUpdatedIterator, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}

	logs, sub, err := _Universalrouter.contract.FilterLogs(opts, "FeeBaseBpsUpdated", msgSenderRule)
	if err != nil {
		return nil, err
	}
	return &UniversalrouterFeeBaseBpsUpdatedIterator{contract: _Universalrouter.contract, event: "FeeBaseBpsUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeBaseBpsUpdated is a free log subscription operation binding the contract event 0xbed4dd7c9eb92506488e6892e1edc0cd8fe578bce7c3de4699223ef5ef34e297.
//
// Solidity: event FeeBaseBpsUpdated(address indexed msgSender, uint256 feeBaseBps)
func (_Universalrouter *UniversalrouterFilterer) WatchFeeBaseBpsUpdated(opts *bind.WatchOpts, sink chan<- *UniversalrouterFeeBaseBpsUpdated, msgSender []common.Address) (event.Subscription, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}

	logs, sub, err := _Universalrouter.contract.WatchLogs(opts, "FeeBaseBpsUpdated", msgSenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalrouterFeeBaseBpsUpdated)
				if err := _Universalrouter.contract.UnpackLog(event, "FeeBaseBpsUpdated", log); err != nil {
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

// ParseFeeBaseBpsUpdated is a log parse operation binding the contract event 0xbed4dd7c9eb92506488e6892e1edc0cd8fe578bce7c3de4699223ef5ef34e297.
//
// Solidity: event FeeBaseBpsUpdated(address indexed msgSender, uint256 feeBaseBps)
func (_Universalrouter *UniversalrouterFilterer) ParseFeeBaseBpsUpdated(log types.Log) (*UniversalrouterFeeBaseBpsUpdated, error) {
	event := new(UniversalrouterFeeBaseBpsUpdated)
	if err := _Universalrouter.contract.UnpackLog(event, "FeeBaseBpsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalrouterFeeBpsUpdatedIterator is returned from FilterFeeBpsUpdated and is used to iterate over the raw logs and unpacked data for FeeBpsUpdated events raised by the Universalrouter contract.
type UniversalrouterFeeBpsUpdatedIterator struct {
	Event *UniversalrouterFeeBpsUpdated // Event containing the contract specifics and raw log

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
func (it *UniversalrouterFeeBpsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalrouterFeeBpsUpdated)
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
		it.Event = new(UniversalrouterFeeBpsUpdated)
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
func (it *UniversalrouterFeeBpsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalrouterFeeBpsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalrouterFeeBpsUpdated represents a FeeBpsUpdated event raised by the Universalrouter contract.
type UniversalrouterFeeBpsUpdated struct {
	MsgSender common.Address
	Level     *big.Int
	SwapType  *big.Int
	FeeBps    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeeBpsUpdated is a free log retrieval operation binding the contract event 0xea2cde73613b525b182d8c53369636ecffc8b1b0747df51489d69fd290ccc55d.
//
// Solidity: event FeeBpsUpdated(address indexed msgSender, uint256 level, uint256 swapType, uint256 feeBps)
func (_Universalrouter *UniversalrouterFilterer) FilterFeeBpsUpdated(opts *bind.FilterOpts, msgSender []common.Address) (*UniversalrouterFeeBpsUpdatedIterator, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}

	logs, sub, err := _Universalrouter.contract.FilterLogs(opts, "FeeBpsUpdated", msgSenderRule)
	if err != nil {
		return nil, err
	}
	return &UniversalrouterFeeBpsUpdatedIterator{contract: _Universalrouter.contract, event: "FeeBpsUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeBpsUpdated is a free log subscription operation binding the contract event 0xea2cde73613b525b182d8c53369636ecffc8b1b0747df51489d69fd290ccc55d.
//
// Solidity: event FeeBpsUpdated(address indexed msgSender, uint256 level, uint256 swapType, uint256 feeBps)
func (_Universalrouter *UniversalrouterFilterer) WatchFeeBpsUpdated(opts *bind.WatchOpts, sink chan<- *UniversalrouterFeeBpsUpdated, msgSender []common.Address) (event.Subscription, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}

	logs, sub, err := _Universalrouter.contract.WatchLogs(opts, "FeeBpsUpdated", msgSenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalrouterFeeBpsUpdated)
				if err := _Universalrouter.contract.UnpackLog(event, "FeeBpsUpdated", log); err != nil {
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

// ParseFeeBpsUpdated is a log parse operation binding the contract event 0xea2cde73613b525b182d8c53369636ecffc8b1b0747df51489d69fd290ccc55d.
//
// Solidity: event FeeBpsUpdated(address indexed msgSender, uint256 level, uint256 swapType, uint256 feeBps)
func (_Universalrouter *UniversalrouterFilterer) ParseFeeBpsUpdated(log types.Log) (*UniversalrouterFeeBpsUpdated, error) {
	event := new(UniversalrouterFeeBpsUpdated)
	if err := _Universalrouter.contract.UnpackLog(event, "FeeBpsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalrouterFeeRecipientUpdatedIterator is returned from FilterFeeRecipientUpdated and is used to iterate over the raw logs and unpacked data for FeeRecipientUpdated events raised by the Universalrouter contract.
type UniversalrouterFeeRecipientUpdatedIterator struct {
	Event *UniversalrouterFeeRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *UniversalrouterFeeRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalrouterFeeRecipientUpdated)
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
		it.Event = new(UniversalrouterFeeRecipientUpdated)
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
func (it *UniversalrouterFeeRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalrouterFeeRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalrouterFeeRecipientUpdated represents a FeeRecipientUpdated event raised by the Universalrouter contract.
type UniversalrouterFeeRecipientUpdated struct {
	MsgSender    common.Address
	FeeRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientUpdated is a free log retrieval operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address indexed msgSender, address feeRecipient)
func (_Universalrouter *UniversalrouterFilterer) FilterFeeRecipientUpdated(opts *bind.FilterOpts, msgSender []common.Address) (*UniversalrouterFeeRecipientUpdatedIterator, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}

	logs, sub, err := _Universalrouter.contract.FilterLogs(opts, "FeeRecipientUpdated", msgSenderRule)
	if err != nil {
		return nil, err
	}
	return &UniversalrouterFeeRecipientUpdatedIterator{contract: _Universalrouter.contract, event: "FeeRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientUpdated is a free log subscription operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address indexed msgSender, address feeRecipient)
func (_Universalrouter *UniversalrouterFilterer) WatchFeeRecipientUpdated(opts *bind.WatchOpts, sink chan<- *UniversalrouterFeeRecipientUpdated, msgSender []common.Address) (event.Subscription, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}

	logs, sub, err := _Universalrouter.contract.WatchLogs(opts, "FeeRecipientUpdated", msgSenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalrouterFeeRecipientUpdated)
				if err := _Universalrouter.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
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

// ParseFeeRecipientUpdated is a log parse operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address indexed msgSender, address feeRecipient)
func (_Universalrouter *UniversalrouterFilterer) ParseFeeRecipientUpdated(log types.Log) (*UniversalrouterFeeRecipientUpdated, error) {
	event := new(UniversalrouterFeeRecipientUpdated)
	if err := _Universalrouter.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalrouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Universalrouter contract.
type UniversalrouterOwnershipTransferredIterator struct {
	Event *UniversalrouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UniversalrouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalrouterOwnershipTransferred)
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
		it.Event = new(UniversalrouterOwnershipTransferred)
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
func (it *UniversalrouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalrouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalrouterOwnershipTransferred represents a OwnershipTransferred event raised by the Universalrouter contract.
type UniversalrouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Universalrouter *UniversalrouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UniversalrouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Universalrouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UniversalrouterOwnershipTransferredIterator{contract: _Universalrouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Universalrouter *UniversalrouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UniversalrouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Universalrouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalrouterOwnershipTransferred)
				if err := _Universalrouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Universalrouter *UniversalrouterFilterer) ParseOwnershipTransferred(log types.Log) (*UniversalrouterOwnershipTransferred, error) {
	event := new(UniversalrouterOwnershipTransferred)
	if err := _Universalrouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalrouterPaymentIterator is returned from FilterPayment and is used to iterate over the raw logs and unpacked data for Payment events raised by the Universalrouter contract.
type UniversalrouterPaymentIterator struct {
	Event *UniversalrouterPayment // Event containing the contract specifics and raw log

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
func (it *UniversalrouterPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalrouterPayment)
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
		it.Event = new(UniversalrouterPayment)
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
func (it *UniversalrouterPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalrouterPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalrouterPayment represents a Payment event raised by the Universalrouter contract.
type UniversalrouterPayment struct {
	Payer      common.Address
	TokenIn    common.Address
	AmountIn   *big.Int
	FeeToken   common.Address
	FeeAmount  *big.Int
	Level      *big.Int
	SwapType   *big.Int
	FeeBps     *big.Int
	FeeBaseBps *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPayment is a free log retrieval operation binding the contract event 0xff4cac4d9f8ca0a2603dfc6a0dafa8be81c2c3cf70a7d25094ab220300315eda.
//
// Solidity: event Payment(address payer, address tokenIn, uint256 amountIn, address feeToken, uint256 feeAmount, uint256 level, uint256 swapType, uint256 feeBps, uint256 feeBaseBps)
func (_Universalrouter *UniversalrouterFilterer) FilterPayment(opts *bind.FilterOpts) (*UniversalrouterPaymentIterator, error) {

	logs, sub, err := _Universalrouter.contract.FilterLogs(opts, "Payment")
	if err != nil {
		return nil, err
	}
	return &UniversalrouterPaymentIterator{contract: _Universalrouter.contract, event: "Payment", logs: logs, sub: sub}, nil
}

// WatchPayment is a free log subscription operation binding the contract event 0xff4cac4d9f8ca0a2603dfc6a0dafa8be81c2c3cf70a7d25094ab220300315eda.
//
// Solidity: event Payment(address payer, address tokenIn, uint256 amountIn, address feeToken, uint256 feeAmount, uint256 level, uint256 swapType, uint256 feeBps, uint256 feeBaseBps)
func (_Universalrouter *UniversalrouterFilterer) WatchPayment(opts *bind.WatchOpts, sink chan<- *UniversalrouterPayment) (event.Subscription, error) {

	logs, sub, err := _Universalrouter.contract.WatchLogs(opts, "Payment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalrouterPayment)
				if err := _Universalrouter.contract.UnpackLog(event, "Payment", log); err != nil {
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

// ParsePayment is a log parse operation binding the contract event 0xff4cac4d9f8ca0a2603dfc6a0dafa8be81c2c3cf70a7d25094ab220300315eda.
//
// Solidity: event Payment(address payer, address tokenIn, uint256 amountIn, address feeToken, uint256 feeAmount, uint256 level, uint256 swapType, uint256 feeBps, uint256 feeBaseBps)
func (_Universalrouter *UniversalrouterFilterer) ParsePayment(log types.Log) (*UniversalrouterPayment, error) {
	event := new(UniversalrouterPayment)
	if err := _Universalrouter.contract.UnpackLog(event, "Payment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UniversalrouterPermit2ErrorIterator is returned from FilterPermit2Error and is used to iterate over the raw logs and unpacked data for Permit2Error events raised by the Universalrouter contract.
type UniversalrouterPermit2ErrorIterator struct {
	Event *UniversalrouterPermit2Error // Event containing the contract specifics and raw log

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
func (it *UniversalrouterPermit2ErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UniversalrouterPermit2Error)
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
		it.Event = new(UniversalrouterPermit2Error)
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
func (it *UniversalrouterPermit2ErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UniversalrouterPermit2ErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UniversalrouterPermit2Error represents a Permit2Error event raised by the Universalrouter contract.
type UniversalrouterPermit2Error struct {
	A   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPermit2Error is a free log retrieval operation binding the contract event 0x6e378e17e70ef12af2577d74c95bd54b67410761c7344725d6e61aed8a02baf0.
//
// Solidity: event Permit2Error(int256 a)
func (_Universalrouter *UniversalrouterFilterer) FilterPermit2Error(opts *bind.FilterOpts) (*UniversalrouterPermit2ErrorIterator, error) {

	logs, sub, err := _Universalrouter.contract.FilterLogs(opts, "Permit2Error")
	if err != nil {
		return nil, err
	}
	return &UniversalrouterPermit2ErrorIterator{contract: _Universalrouter.contract, event: "Permit2Error", logs: logs, sub: sub}, nil
}

// WatchPermit2Error is a free log subscription operation binding the contract event 0x6e378e17e70ef12af2577d74c95bd54b67410761c7344725d6e61aed8a02baf0.
//
// Solidity: event Permit2Error(int256 a)
func (_Universalrouter *UniversalrouterFilterer) WatchPermit2Error(opts *bind.WatchOpts, sink chan<- *UniversalrouterPermit2Error) (event.Subscription, error) {

	logs, sub, err := _Universalrouter.contract.WatchLogs(opts, "Permit2Error")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UniversalrouterPermit2Error)
				if err := _Universalrouter.contract.UnpackLog(event, "Permit2Error", log); err != nil {
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

// ParsePermit2Error is a log parse operation binding the contract event 0x6e378e17e70ef12af2577d74c95bd54b67410761c7344725d6e61aed8a02baf0.
//
// Solidity: event Permit2Error(int256 a)
func (_Universalrouter *UniversalrouterFilterer) ParsePermit2Error(log types.Log) (*UniversalrouterPermit2Error, error) {
	event := new(UniversalrouterPermit2Error)
	if err := _Universalrouter.contract.UnpackLog(event, "Permit2Error", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
