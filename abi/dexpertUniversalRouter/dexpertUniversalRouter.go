// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dexpertUniversalRouter

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

// DexpertUniversalRouterMetaData contains all meta data concerning the DexpertUniversalRouter contract.
var DexpertUniversalRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"uniswapV2Router02\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeBaseBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"permit2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"weth9\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v2Factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"v3Factory\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"pairInitCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"poolInitCodeHash\",\"type\":\"bytes32\"}],\"internalType\":\"structRouterParameters\",\"name\":\"params\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BalanceTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BuyPunkFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ContractLocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ETHNotAccepted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commandIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"ExecutionFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientAddressCannotBeZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeRecipientAddressCannotBeZeroAddress1\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FromAddressIsNotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBips\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"commandType\",\"type\":\"uint256\"}],\"name\":\"InvalidCommandType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeBaseBps\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeBaseBps\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeBps\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReserves\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSpender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"InvalidValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SliceOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionDeadlinePassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsafeCast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2InvalidPath\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2TooLittleReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V2TooMuchRequested\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3InvalidAmountOut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3InvalidSwap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3TooLittleReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"V3TooMuchRequested\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBaseBps\",\"type\":\"uint256\"}],\"name\":\"FeeBaseBpsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"}],\"name\":\"FeeBpsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBaseBps\",\"type\":\"uint256\"}],\"name\":\"Payment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"}],\"name\":\"Permit2Error\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"commands\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"inputs\",\"type\":\"bytes[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeBaseBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapType\",\"type\":\"uint256\"}],\"name\":\"feeBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeBaseBps\",\"type\":\"uint256\"}],\"name\":\"setFeeBaseBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"}],\"name\":\"setFeeBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"amount0Delta\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"amount1Delta\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV3SwapCallback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// DexpertUniversalRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use DexpertUniversalRouterMetaData.ABI instead.
var DexpertUniversalRouterABI = DexpertUniversalRouterMetaData.ABI

// DexpertUniversalRouter is an auto generated Go binding around an Ethereum contract.
type DexpertUniversalRouter struct {
	DexpertUniversalRouterCaller     // Read-only binding to the contract
	DexpertUniversalRouterTransactor // Write-only binding to the contract
	DexpertUniversalRouterFilterer   // Log filterer for contract events
}

// DexpertUniversalRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type DexpertUniversalRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexpertUniversalRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DexpertUniversalRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexpertUniversalRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DexpertUniversalRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexpertUniversalRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DexpertUniversalRouterSession struct {
	Contract     *DexpertUniversalRouter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DexpertUniversalRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DexpertUniversalRouterCallerSession struct {
	Contract *DexpertUniversalRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// DexpertUniversalRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DexpertUniversalRouterTransactorSession struct {
	Contract     *DexpertUniversalRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// DexpertUniversalRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type DexpertUniversalRouterRaw struct {
	Contract *DexpertUniversalRouter // Generic contract binding to access the raw methods on
}

// DexpertUniversalRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DexpertUniversalRouterCallerRaw struct {
	Contract *DexpertUniversalRouterCaller // Generic read-only contract binding to access the raw methods on
}

// DexpertUniversalRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DexpertUniversalRouterTransactorRaw struct {
	Contract *DexpertUniversalRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDexpertUniversalRouter creates a new instance of DexpertUniversalRouter, bound to a specific deployed contract.
func NewDexpertUniversalRouter(address common.Address, backend bind.ContractBackend) (*DexpertUniversalRouter, error) {
	contract, err := bindDexpertUniversalRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DexpertUniversalRouter{DexpertUniversalRouterCaller: DexpertUniversalRouterCaller{contract: contract}, DexpertUniversalRouterTransactor: DexpertUniversalRouterTransactor{contract: contract}, DexpertUniversalRouterFilterer: DexpertUniversalRouterFilterer{contract: contract}}, nil
}

// NewDexpertUniversalRouterCaller creates a new read-only instance of DexpertUniversalRouter, bound to a specific deployed contract.
func NewDexpertUniversalRouterCaller(address common.Address, caller bind.ContractCaller) (*DexpertUniversalRouterCaller, error) {
	contract, err := bindDexpertUniversalRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DexpertUniversalRouterCaller{contract: contract}, nil
}

// NewDexpertUniversalRouterTransactor creates a new write-only instance of DexpertUniversalRouter, bound to a specific deployed contract.
func NewDexpertUniversalRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*DexpertUniversalRouterTransactor, error) {
	contract, err := bindDexpertUniversalRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DexpertUniversalRouterTransactor{contract: contract}, nil
}

// NewDexpertUniversalRouterFilterer creates a new log filterer instance of DexpertUniversalRouter, bound to a specific deployed contract.
func NewDexpertUniversalRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*DexpertUniversalRouterFilterer, error) {
	contract, err := bindDexpertUniversalRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DexpertUniversalRouterFilterer{contract: contract}, nil
}

// bindDexpertUniversalRouter binds a generic wrapper to an already deployed contract.
func bindDexpertUniversalRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DexpertUniversalRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DexpertUniversalRouter *DexpertUniversalRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DexpertUniversalRouter.Contract.DexpertUniversalRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DexpertUniversalRouter *DexpertUniversalRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.DexpertUniversalRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DexpertUniversalRouter *DexpertUniversalRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.DexpertUniversalRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DexpertUniversalRouter *DexpertUniversalRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DexpertUniversalRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.contract.Transact(opts, method, params...)
}

// FeeBaseBps is a free data retrieval call binding the contract method 0x8ffb274f.
//
// Solidity: function feeBaseBps() view returns(uint256)
func (_DexpertUniversalRouter *DexpertUniversalRouterCaller) FeeBaseBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DexpertUniversalRouter.contract.Call(opts, &out, "feeBaseBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeBaseBps is a free data retrieval call binding the contract method 0x8ffb274f.
//
// Solidity: function feeBaseBps() view returns(uint256)
func (_DexpertUniversalRouter *DexpertUniversalRouterSession) FeeBaseBps() (*big.Int, error) {
	return _DexpertUniversalRouter.Contract.FeeBaseBps(&_DexpertUniversalRouter.CallOpts)
}

// FeeBaseBps is a free data retrieval call binding the contract method 0x8ffb274f.
//
// Solidity: function feeBaseBps() view returns(uint256)
func (_DexpertUniversalRouter *DexpertUniversalRouterCallerSession) FeeBaseBps() (*big.Int, error) {
	return _DexpertUniversalRouter.Contract.FeeBaseBps(&_DexpertUniversalRouter.CallOpts)
}

// FeeBps is a free data retrieval call binding the contract method 0xd9d852af.
//
// Solidity: function feeBps(uint256 level, uint256 swapType) view returns(uint256)
func (_DexpertUniversalRouter *DexpertUniversalRouterCaller) FeeBps(opts *bind.CallOpts, level *big.Int, swapType *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DexpertUniversalRouter.contract.Call(opts, &out, "feeBps", level, swapType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeBps is a free data retrieval call binding the contract method 0xd9d852af.
//
// Solidity: function feeBps(uint256 level, uint256 swapType) view returns(uint256)
func (_DexpertUniversalRouter *DexpertUniversalRouterSession) FeeBps(level *big.Int, swapType *big.Int) (*big.Int, error) {
	return _DexpertUniversalRouter.Contract.FeeBps(&_DexpertUniversalRouter.CallOpts, level, swapType)
}

// FeeBps is a free data retrieval call binding the contract method 0xd9d852af.
//
// Solidity: function feeBps(uint256 level, uint256 swapType) view returns(uint256)
func (_DexpertUniversalRouter *DexpertUniversalRouterCallerSession) FeeBps(level *big.Int, swapType *big.Int) (*big.Int, error) {
	return _DexpertUniversalRouter.Contract.FeeBps(&_DexpertUniversalRouter.CallOpts, level, swapType)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_DexpertUniversalRouter *DexpertUniversalRouterCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DexpertUniversalRouter.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_DexpertUniversalRouter *DexpertUniversalRouterSession) FeeRecipient() (common.Address, error) {
	return _DexpertUniversalRouter.Contract.FeeRecipient(&_DexpertUniversalRouter.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_DexpertUniversalRouter *DexpertUniversalRouterCallerSession) FeeRecipient() (common.Address, error) {
	return _DexpertUniversalRouter.Contract.FeeRecipient(&_DexpertUniversalRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DexpertUniversalRouter *DexpertUniversalRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DexpertUniversalRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DexpertUniversalRouter *DexpertUniversalRouterSession) Owner() (common.Address, error) {
	return _DexpertUniversalRouter.Contract.Owner(&_DexpertUniversalRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DexpertUniversalRouter *DexpertUniversalRouterCallerSession) Owner() (common.Address, error) {
	return _DexpertUniversalRouter.Contract.Owner(&_DexpertUniversalRouter.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactor) Execute(opts *bind.TransactOpts, commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _DexpertUniversalRouter.contract.Transact(opts, "execute", commands, inputs)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterSession) Execute(commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.Execute(&_DexpertUniversalRouter.TransactOpts, commands, inputs)
}

// Execute is a paid mutator transaction binding the contract method 0x24856bc3.
//
// Solidity: function execute(bytes commands, bytes[] inputs) payable returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactorSession) Execute(commands []byte, inputs [][]byte) (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.Execute(&_DexpertUniversalRouter.TransactOpts, commands, inputs)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactor) Execute0(opts *bind.TransactOpts, commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _DexpertUniversalRouter.contract.Transact(opts, "execute0", commands, inputs, deadline)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterSession) Execute0(commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.Execute0(&_DexpertUniversalRouter.TransactOpts, commands, inputs, deadline)
}

// Execute0 is a paid mutator transaction binding the contract method 0x3593564c.
//
// Solidity: function execute(bytes commands, bytes[] inputs, uint256 deadline) payable returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactorSession) Execute0(commands []byte, inputs [][]byte, deadline *big.Int) (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.Execute0(&_DexpertUniversalRouter.TransactOpts, commands, inputs, deadline)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DexpertUniversalRouter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.RenounceOwnership(&_DexpertUniversalRouter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.RenounceOwnership(&_DexpertUniversalRouter.TransactOpts)
}

// SetFeeBaseBps is a paid mutator transaction binding the contract method 0x086408db.
//
// Solidity: function setFeeBaseBps(uint256 feeBaseBps) returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactor) SetFeeBaseBps(opts *bind.TransactOpts, feeBaseBps *big.Int) (*types.Transaction, error) {
	return _DexpertUniversalRouter.contract.Transact(opts, "setFeeBaseBps", feeBaseBps)
}

// SetFeeBaseBps is a paid mutator transaction binding the contract method 0x086408db.
//
// Solidity: function setFeeBaseBps(uint256 feeBaseBps) returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterSession) SetFeeBaseBps(feeBaseBps *big.Int) (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.SetFeeBaseBps(&_DexpertUniversalRouter.TransactOpts, feeBaseBps)
}

// SetFeeBaseBps is a paid mutator transaction binding the contract method 0x086408db.
//
// Solidity: function setFeeBaseBps(uint256 feeBaseBps) returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactorSession) SetFeeBaseBps(feeBaseBps *big.Int) (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.SetFeeBaseBps(&_DexpertUniversalRouter.TransactOpts, feeBaseBps)
}

// SetFeeBps is a paid mutator transaction binding the contract method 0x8aa06386.
//
// Solidity: function setFeeBps(uint256 level, uint256 swapType, uint256 feeBps) returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactor) SetFeeBps(opts *bind.TransactOpts, level *big.Int, swapType *big.Int, feeBps *big.Int) (*types.Transaction, error) {
	return _DexpertUniversalRouter.contract.Transact(opts, "setFeeBps", level, swapType, feeBps)
}

// SetFeeBps is a paid mutator transaction binding the contract method 0x8aa06386.
//
// Solidity: function setFeeBps(uint256 level, uint256 swapType, uint256 feeBps) returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterSession) SetFeeBps(level *big.Int, swapType *big.Int, feeBps *big.Int) (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.SetFeeBps(&_DexpertUniversalRouter.TransactOpts, level, swapType, feeBps)
}

// SetFeeBps is a paid mutator transaction binding the contract method 0x8aa06386.
//
// Solidity: function setFeeBps(uint256 level, uint256 swapType, uint256 feeBps) returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactorSession) SetFeeBps(level *big.Int, swapType *big.Int, feeBps *big.Int) (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.SetFeeBps(&_DexpertUniversalRouter.TransactOpts, level, swapType, feeBps)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address feeRecipient) returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactor) SetFeeRecipient(opts *bind.TransactOpts, feeRecipient common.Address) (*types.Transaction, error) {
	return _DexpertUniversalRouter.contract.Transact(opts, "setFeeRecipient", feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address feeRecipient) returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterSession) SetFeeRecipient(feeRecipient common.Address) (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.SetFeeRecipient(&_DexpertUniversalRouter.TransactOpts, feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address feeRecipient) returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactorSession) SetFeeRecipient(feeRecipient common.Address) (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.SetFeeRecipient(&_DexpertUniversalRouter.TransactOpts, feeRecipient)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DexpertUniversalRouter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.TransferOwnership(&_DexpertUniversalRouter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.TransferOwnership(&_DexpertUniversalRouter.TransactOpts, newOwner)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactor) UniswapV3SwapCallback(opts *bind.TransactOpts, amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _DexpertUniversalRouter.contract.Transact(opts, "uniswapV3SwapCallback", amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.UniswapV3SwapCallback(&_DexpertUniversalRouter.TransactOpts, amount0Delta, amount1Delta, data)
}

// UniswapV3SwapCallback is a paid mutator transaction binding the contract method 0xfa461e33.
//
// Solidity: function uniswapV3SwapCallback(int256 amount0Delta, int256 amount1Delta, bytes data) returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactorSession) UniswapV3SwapCallback(amount0Delta *big.Int, amount1Delta *big.Int, data []byte) (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.UniswapV3SwapCallback(&_DexpertUniversalRouter.TransactOpts, amount0Delta, amount1Delta, data)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DexpertUniversalRouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterSession) Receive() (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.Receive(&_DexpertUniversalRouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DexpertUniversalRouter *DexpertUniversalRouterTransactorSession) Receive() (*types.Transaction, error) {
	return _DexpertUniversalRouter.Contract.Receive(&_DexpertUniversalRouter.TransactOpts)
}

// DexpertUniversalRouterFeeBaseBpsUpdatedIterator is returned from FilterFeeBaseBpsUpdated and is used to iterate over the raw logs and unpacked data for FeeBaseBpsUpdated events raised by the DexpertUniversalRouter contract.
type DexpertUniversalRouterFeeBaseBpsUpdatedIterator struct {
	Event *DexpertUniversalRouterFeeBaseBpsUpdated // Event containing the contract specifics and raw log

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
func (it *DexpertUniversalRouterFeeBaseBpsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexpertUniversalRouterFeeBaseBpsUpdated)
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
		it.Event = new(DexpertUniversalRouterFeeBaseBpsUpdated)
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
func (it *DexpertUniversalRouterFeeBaseBpsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexpertUniversalRouterFeeBaseBpsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexpertUniversalRouterFeeBaseBpsUpdated represents a FeeBaseBpsUpdated event raised by the DexpertUniversalRouter contract.
type DexpertUniversalRouterFeeBaseBpsUpdated struct {
	MsgSender  common.Address
	FeeBaseBps *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeBaseBpsUpdated is a free log retrieval operation binding the contract event 0xbed4dd7c9eb92506488e6892e1edc0cd8fe578bce7c3de4699223ef5ef34e297.
//
// Solidity: event FeeBaseBpsUpdated(address indexed msgSender, uint256 feeBaseBps)
func (_DexpertUniversalRouter *DexpertUniversalRouterFilterer) FilterFeeBaseBpsUpdated(opts *bind.FilterOpts, msgSender []common.Address) (*DexpertUniversalRouterFeeBaseBpsUpdatedIterator, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}

	logs, sub, err := _DexpertUniversalRouter.contract.FilterLogs(opts, "FeeBaseBpsUpdated", msgSenderRule)
	if err != nil {
		return nil, err
	}
	return &DexpertUniversalRouterFeeBaseBpsUpdatedIterator{contract: _DexpertUniversalRouter.contract, event: "FeeBaseBpsUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeBaseBpsUpdated is a free log subscription operation binding the contract event 0xbed4dd7c9eb92506488e6892e1edc0cd8fe578bce7c3de4699223ef5ef34e297.
//
// Solidity: event FeeBaseBpsUpdated(address indexed msgSender, uint256 feeBaseBps)
func (_DexpertUniversalRouter *DexpertUniversalRouterFilterer) WatchFeeBaseBpsUpdated(opts *bind.WatchOpts, sink chan<- *DexpertUniversalRouterFeeBaseBpsUpdated, msgSender []common.Address) (event.Subscription, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}

	logs, sub, err := _DexpertUniversalRouter.contract.WatchLogs(opts, "FeeBaseBpsUpdated", msgSenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexpertUniversalRouterFeeBaseBpsUpdated)
				if err := _DexpertUniversalRouter.contract.UnpackLog(event, "FeeBaseBpsUpdated", log); err != nil {
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
func (_DexpertUniversalRouter *DexpertUniversalRouterFilterer) ParseFeeBaseBpsUpdated(log types.Log) (*DexpertUniversalRouterFeeBaseBpsUpdated, error) {
	event := new(DexpertUniversalRouterFeeBaseBpsUpdated)
	if err := _DexpertUniversalRouter.contract.UnpackLog(event, "FeeBaseBpsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexpertUniversalRouterFeeBpsUpdatedIterator is returned from FilterFeeBpsUpdated and is used to iterate over the raw logs and unpacked data for FeeBpsUpdated events raised by the DexpertUniversalRouter contract.
type DexpertUniversalRouterFeeBpsUpdatedIterator struct {
	Event *DexpertUniversalRouterFeeBpsUpdated // Event containing the contract specifics and raw log

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
func (it *DexpertUniversalRouterFeeBpsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexpertUniversalRouterFeeBpsUpdated)
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
		it.Event = new(DexpertUniversalRouterFeeBpsUpdated)
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
func (it *DexpertUniversalRouterFeeBpsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexpertUniversalRouterFeeBpsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexpertUniversalRouterFeeBpsUpdated represents a FeeBpsUpdated event raised by the DexpertUniversalRouter contract.
type DexpertUniversalRouterFeeBpsUpdated struct {
	MsgSender common.Address
	Level     *big.Int
	SwapType  *big.Int
	FeeBps    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeeBpsUpdated is a free log retrieval operation binding the contract event 0xea2cde73613b525b182d8c53369636ecffc8b1b0747df51489d69fd290ccc55d.
//
// Solidity: event FeeBpsUpdated(address indexed msgSender, uint256 level, uint256 swapType, uint256 feeBps)
func (_DexpertUniversalRouter *DexpertUniversalRouterFilterer) FilterFeeBpsUpdated(opts *bind.FilterOpts, msgSender []common.Address) (*DexpertUniversalRouterFeeBpsUpdatedIterator, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}

	logs, sub, err := _DexpertUniversalRouter.contract.FilterLogs(opts, "FeeBpsUpdated", msgSenderRule)
	if err != nil {
		return nil, err
	}
	return &DexpertUniversalRouterFeeBpsUpdatedIterator{contract: _DexpertUniversalRouter.contract, event: "FeeBpsUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeBpsUpdated is a free log subscription operation binding the contract event 0xea2cde73613b525b182d8c53369636ecffc8b1b0747df51489d69fd290ccc55d.
//
// Solidity: event FeeBpsUpdated(address indexed msgSender, uint256 level, uint256 swapType, uint256 feeBps)
func (_DexpertUniversalRouter *DexpertUniversalRouterFilterer) WatchFeeBpsUpdated(opts *bind.WatchOpts, sink chan<- *DexpertUniversalRouterFeeBpsUpdated, msgSender []common.Address) (event.Subscription, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}

	logs, sub, err := _DexpertUniversalRouter.contract.WatchLogs(opts, "FeeBpsUpdated", msgSenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexpertUniversalRouterFeeBpsUpdated)
				if err := _DexpertUniversalRouter.contract.UnpackLog(event, "FeeBpsUpdated", log); err != nil {
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
func (_DexpertUniversalRouter *DexpertUniversalRouterFilterer) ParseFeeBpsUpdated(log types.Log) (*DexpertUniversalRouterFeeBpsUpdated, error) {
	event := new(DexpertUniversalRouterFeeBpsUpdated)
	if err := _DexpertUniversalRouter.contract.UnpackLog(event, "FeeBpsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexpertUniversalRouterFeeRecipientUpdatedIterator is returned from FilterFeeRecipientUpdated and is used to iterate over the raw logs and unpacked data for FeeRecipientUpdated events raised by the DexpertUniversalRouter contract.
type DexpertUniversalRouterFeeRecipientUpdatedIterator struct {
	Event *DexpertUniversalRouterFeeRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *DexpertUniversalRouterFeeRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexpertUniversalRouterFeeRecipientUpdated)
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
		it.Event = new(DexpertUniversalRouterFeeRecipientUpdated)
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
func (it *DexpertUniversalRouterFeeRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexpertUniversalRouterFeeRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexpertUniversalRouterFeeRecipientUpdated represents a FeeRecipientUpdated event raised by the DexpertUniversalRouter contract.
type DexpertUniversalRouterFeeRecipientUpdated struct {
	MsgSender    common.Address
	FeeRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientUpdated is a free log retrieval operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address indexed msgSender, address feeRecipient)
func (_DexpertUniversalRouter *DexpertUniversalRouterFilterer) FilterFeeRecipientUpdated(opts *bind.FilterOpts, msgSender []common.Address) (*DexpertUniversalRouterFeeRecipientUpdatedIterator, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}

	logs, sub, err := _DexpertUniversalRouter.contract.FilterLogs(opts, "FeeRecipientUpdated", msgSenderRule)
	if err != nil {
		return nil, err
	}
	return &DexpertUniversalRouterFeeRecipientUpdatedIterator{contract: _DexpertUniversalRouter.contract, event: "FeeRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientUpdated is a free log subscription operation binding the contract event 0xaaebcf1bfa00580e41d966056b48521fa9f202645c86d4ddf28113e617c1b1d3.
//
// Solidity: event FeeRecipientUpdated(address indexed msgSender, address feeRecipient)
func (_DexpertUniversalRouter *DexpertUniversalRouterFilterer) WatchFeeRecipientUpdated(opts *bind.WatchOpts, sink chan<- *DexpertUniversalRouterFeeRecipientUpdated, msgSender []common.Address) (event.Subscription, error) {

	var msgSenderRule []interface{}
	for _, msgSenderItem := range msgSender {
		msgSenderRule = append(msgSenderRule, msgSenderItem)
	}

	logs, sub, err := _DexpertUniversalRouter.contract.WatchLogs(opts, "FeeRecipientUpdated", msgSenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexpertUniversalRouterFeeRecipientUpdated)
				if err := _DexpertUniversalRouter.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
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
func (_DexpertUniversalRouter *DexpertUniversalRouterFilterer) ParseFeeRecipientUpdated(log types.Log) (*DexpertUniversalRouterFeeRecipientUpdated, error) {
	event := new(DexpertUniversalRouterFeeRecipientUpdated)
	if err := _DexpertUniversalRouter.contract.UnpackLog(event, "FeeRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexpertUniversalRouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DexpertUniversalRouter contract.
type DexpertUniversalRouterOwnershipTransferredIterator struct {
	Event *DexpertUniversalRouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DexpertUniversalRouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexpertUniversalRouterOwnershipTransferred)
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
		it.Event = new(DexpertUniversalRouterOwnershipTransferred)
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
func (it *DexpertUniversalRouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexpertUniversalRouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexpertUniversalRouterOwnershipTransferred represents a OwnershipTransferred event raised by the DexpertUniversalRouter contract.
type DexpertUniversalRouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DexpertUniversalRouter *DexpertUniversalRouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DexpertUniversalRouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DexpertUniversalRouter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DexpertUniversalRouterOwnershipTransferredIterator{contract: _DexpertUniversalRouter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DexpertUniversalRouter *DexpertUniversalRouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DexpertUniversalRouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DexpertUniversalRouter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexpertUniversalRouterOwnershipTransferred)
				if err := _DexpertUniversalRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DexpertUniversalRouter *DexpertUniversalRouterFilterer) ParseOwnershipTransferred(log types.Log) (*DexpertUniversalRouterOwnershipTransferred, error) {
	event := new(DexpertUniversalRouterOwnershipTransferred)
	if err := _DexpertUniversalRouter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexpertUniversalRouterPaymentIterator is returned from FilterPayment and is used to iterate over the raw logs and unpacked data for Payment events raised by the DexpertUniversalRouter contract.
type DexpertUniversalRouterPaymentIterator struct {
	Event *DexpertUniversalRouterPayment // Event containing the contract specifics and raw log

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
func (it *DexpertUniversalRouterPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexpertUniversalRouterPayment)
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
		it.Event = new(DexpertUniversalRouterPayment)
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
func (it *DexpertUniversalRouterPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexpertUniversalRouterPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexpertUniversalRouterPayment represents a Payment event raised by the DexpertUniversalRouter contract.
type DexpertUniversalRouterPayment struct {
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
func (_DexpertUniversalRouter *DexpertUniversalRouterFilterer) FilterPayment(opts *bind.FilterOpts) (*DexpertUniversalRouterPaymentIterator, error) {

	logs, sub, err := _DexpertUniversalRouter.contract.FilterLogs(opts, "Payment")
	if err != nil {
		return nil, err
	}
	return &DexpertUniversalRouterPaymentIterator{contract: _DexpertUniversalRouter.contract, event: "Payment", logs: logs, sub: sub}, nil
}

// WatchPayment is a free log subscription operation binding the contract event 0xff4cac4d9f8ca0a2603dfc6a0dafa8be81c2c3cf70a7d25094ab220300315eda.
//
// Solidity: event Payment(address payer, address tokenIn, uint256 amountIn, address feeToken, uint256 feeAmount, uint256 level, uint256 swapType, uint256 feeBps, uint256 feeBaseBps)
func (_DexpertUniversalRouter *DexpertUniversalRouterFilterer) WatchPayment(opts *bind.WatchOpts, sink chan<- *DexpertUniversalRouterPayment) (event.Subscription, error) {

	logs, sub, err := _DexpertUniversalRouter.contract.WatchLogs(opts, "Payment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexpertUniversalRouterPayment)
				if err := _DexpertUniversalRouter.contract.UnpackLog(event, "Payment", log); err != nil {
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
func (_DexpertUniversalRouter *DexpertUniversalRouterFilterer) ParsePayment(log types.Log) (*DexpertUniversalRouterPayment, error) {
	event := new(DexpertUniversalRouterPayment)
	if err := _DexpertUniversalRouter.contract.UnpackLog(event, "Payment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DexpertUniversalRouterPermit2ErrorIterator is returned from FilterPermit2Error and is used to iterate over the raw logs and unpacked data for Permit2Error events raised by the DexpertUniversalRouter contract.
type DexpertUniversalRouterPermit2ErrorIterator struct {
	Event *DexpertUniversalRouterPermit2Error // Event containing the contract specifics and raw log

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
func (it *DexpertUniversalRouterPermit2ErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DexpertUniversalRouterPermit2Error)
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
		it.Event = new(DexpertUniversalRouterPermit2Error)
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
func (it *DexpertUniversalRouterPermit2ErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DexpertUniversalRouterPermit2ErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DexpertUniversalRouterPermit2Error represents a Permit2Error event raised by the DexpertUniversalRouter contract.
type DexpertUniversalRouterPermit2Error struct {
	A   *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPermit2Error is a free log retrieval operation binding the contract event 0x6e378e17e70ef12af2577d74c95bd54b67410761c7344725d6e61aed8a02baf0.
//
// Solidity: event Permit2Error(int256 a)
func (_DexpertUniversalRouter *DexpertUniversalRouterFilterer) FilterPermit2Error(opts *bind.FilterOpts) (*DexpertUniversalRouterPermit2ErrorIterator, error) {

	logs, sub, err := _DexpertUniversalRouter.contract.FilterLogs(opts, "Permit2Error")
	if err != nil {
		return nil, err
	}
	return &DexpertUniversalRouterPermit2ErrorIterator{contract: _DexpertUniversalRouter.contract, event: "Permit2Error", logs: logs, sub: sub}, nil
}

// WatchPermit2Error is a free log subscription operation binding the contract event 0x6e378e17e70ef12af2577d74c95bd54b67410761c7344725d6e61aed8a02baf0.
//
// Solidity: event Permit2Error(int256 a)
func (_DexpertUniversalRouter *DexpertUniversalRouterFilterer) WatchPermit2Error(opts *bind.WatchOpts, sink chan<- *DexpertUniversalRouterPermit2Error) (event.Subscription, error) {

	logs, sub, err := _DexpertUniversalRouter.contract.WatchLogs(opts, "Permit2Error")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DexpertUniversalRouterPermit2Error)
				if err := _DexpertUniversalRouter.contract.UnpackLog(event, "Permit2Error", log); err != nil {
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
func (_DexpertUniversalRouter *DexpertUniversalRouterFilterer) ParsePermit2Error(log types.Log) (*DexpertUniversalRouterPermit2Error, error) {
	event := new(DexpertUniversalRouterPermit2Error)
	if err := _DexpertUniversalRouter.contract.UnpackLog(event, "Permit2Error", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
