// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package standardTokenFactory01

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

// TokenInitializeParams is an auto generated low-level Go binding around an user-defined struct.
type TokenInitializeParams struct {
	Name         string
	Symbol       string
	Decimals     uint8
	TotalSupply  *big.Int
	Description  string
	LogoLink     string
	TwitterLink  string
	TelegramLink string
	DiscordLink  string
	WebsiteLink  string
}

// TokenMetaData is an auto generated low-level Go binding around an user-defined struct.
type TokenMetaData struct {
	Description  string
	LogoLink     string
	TwitterLink  string
	TelegramLink string
	DiscordLink  string
	WebsiteLink  string
}

// StandardTokenFactory01MetaData contains all meta data concerning the StandardTokenFactory01 contract.
var StandardTokenFactory01MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factoryManager_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeTo_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxFee_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"InsufficientFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"InvalidFactoryManager\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"InvalidFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"InvalidFeeReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factoryManager\",\"type\":\"address\"}],\"name\":\"InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"}],\"name\":\"InvalidLevel\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"InvalidMaxFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeTo\",\"type\":\"address\"}],\"name\":\"FeeToUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"FeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"newLevels\",\"type\":\"uint256[]\"}],\"name\":\"LevelsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"tokenVersion\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"}],\"name\":\"TokenCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logoLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"twitterLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"telegramLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"discordLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"websiteLink\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structTokenMetaData\",\"name\":\"tokenMetaData\",\"type\":\"tuple\"}],\"name\":\"TokenMetaDataUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FACTORY_MANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logoLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"twitterLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"telegramLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"discordLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"websiteLink\",\"type\":\"string\"}],\"internalType\":\"structTokenInitializeParams\",\"name\":\"tokenInitializeParams\",\"type\":\"tuple\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLevels\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementationVersion\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeTo_\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"}],\"name\":\"setImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_levels\",\"type\":\"uint256[]\"}],\"name\":\"setLevels\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logoLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"twitterLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"telegramLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"discordLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"websiteLink\",\"type\":\"string\"}],\"internalType\":\"structTokenMetaData\",\"name\":\"tokenMetaData_\",\"type\":\"tuple\"}],\"name\":\"updateTokenMetaData\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// StandardTokenFactory01ABI is the input ABI used to generate the binding from.
// Deprecated: Use StandardTokenFactory01MetaData.ABI instead.
var StandardTokenFactory01ABI = StandardTokenFactory01MetaData.ABI

// StandardTokenFactory01 is an auto generated Go binding around an Ethereum contract.
type StandardTokenFactory01 struct {
	StandardTokenFactory01Caller     // Read-only binding to the contract
	StandardTokenFactory01Transactor // Write-only binding to the contract
	StandardTokenFactory01Filterer   // Log filterer for contract events
}

// StandardTokenFactory01Caller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenFactory01Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenFactory01Transactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenFactory01Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenFactory01Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardTokenFactory01Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenFactory01Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenFactory01Session struct {
	Contract     *StandardTokenFactory01 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StandardTokenFactory01CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenFactory01CallerSession struct {
	Contract *StandardTokenFactory01Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// StandardTokenFactory01TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenFactory01TransactorSession struct {
	Contract     *StandardTokenFactory01Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// StandardTokenFactory01Raw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenFactory01Raw struct {
	Contract *StandardTokenFactory01 // Generic contract binding to access the raw methods on
}

// StandardTokenFactory01CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenFactory01CallerRaw struct {
	Contract *StandardTokenFactory01Caller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenFactory01TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenFactory01TransactorRaw struct {
	Contract *StandardTokenFactory01Transactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardTokenFactory01 creates a new instance of StandardTokenFactory01, bound to a specific deployed contract.
func NewStandardTokenFactory01(address common.Address, backend bind.ContractBackend) (*StandardTokenFactory01, error) {
	contract, err := bindStandardTokenFactory01(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFactory01{StandardTokenFactory01Caller: StandardTokenFactory01Caller{contract: contract}, StandardTokenFactory01Transactor: StandardTokenFactory01Transactor{contract: contract}, StandardTokenFactory01Filterer: StandardTokenFactory01Filterer{contract: contract}}, nil
}

// NewStandardTokenFactory01Caller creates a new read-only instance of StandardTokenFactory01, bound to a specific deployed contract.
func NewStandardTokenFactory01Caller(address common.Address, caller bind.ContractCaller) (*StandardTokenFactory01Caller, error) {
	contract, err := bindStandardTokenFactory01(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFactory01Caller{contract: contract}, nil
}

// NewStandardTokenFactory01Transactor creates a new write-only instance of StandardTokenFactory01, bound to a specific deployed contract.
func NewStandardTokenFactory01Transactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenFactory01Transactor, error) {
	contract, err := bindStandardTokenFactory01(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFactory01Transactor{contract: contract}, nil
}

// NewStandardTokenFactory01Filterer creates a new log filterer instance of StandardTokenFactory01, bound to a specific deployed contract.
func NewStandardTokenFactory01Filterer(address common.Address, filterer bind.ContractFilterer) (*StandardTokenFactory01Filterer, error) {
	contract, err := bindStandardTokenFactory01(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFactory01Filterer{contract: contract}, nil
}

// bindStandardTokenFactory01 binds a generic wrapper to an already deployed contract.
func bindStandardTokenFactory01(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StandardTokenFactory01MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardTokenFactory01 *StandardTokenFactory01Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StandardTokenFactory01.Contract.StandardTokenFactory01Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardTokenFactory01 *StandardTokenFactory01Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.StandardTokenFactory01Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardTokenFactory01 *StandardTokenFactory01Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.StandardTokenFactory01Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardTokenFactory01 *StandardTokenFactory01CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StandardTokenFactory01.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardTokenFactory01 *StandardTokenFactory01TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardTokenFactory01 *StandardTokenFactory01TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.contract.Transact(opts, method, params...)
}

// FACTORYMANAGER is a free data retrieval call binding the contract method 0x42246a57.
//
// Solidity: function FACTORY_MANAGER() view returns(address)
func (_StandardTokenFactory01 *StandardTokenFactory01Caller) FACTORYMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StandardTokenFactory01.contract.Call(opts, &out, "FACTORY_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FACTORYMANAGER is a free data retrieval call binding the contract method 0x42246a57.
//
// Solidity: function FACTORY_MANAGER() view returns(address)
func (_StandardTokenFactory01 *StandardTokenFactory01Session) FACTORYMANAGER() (common.Address, error) {
	return _StandardTokenFactory01.Contract.FACTORYMANAGER(&_StandardTokenFactory01.CallOpts)
}

// FACTORYMANAGER is a free data retrieval call binding the contract method 0x42246a57.
//
// Solidity: function FACTORY_MANAGER() view returns(address)
func (_StandardTokenFactory01 *StandardTokenFactory01CallerSession) FACTORYMANAGER() (common.Address, error) {
	return _StandardTokenFactory01.Contract.FACTORYMANAGER(&_StandardTokenFactory01.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_StandardTokenFactory01 *StandardTokenFactory01Caller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StandardTokenFactory01.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_StandardTokenFactory01 *StandardTokenFactory01Session) MAXFEE() (*big.Int, error) {
	return _StandardTokenFactory01.Contract.MAXFEE(&_StandardTokenFactory01.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_StandardTokenFactory01 *StandardTokenFactory01CallerSession) MAXFEE() (*big.Int, error) {
	return _StandardTokenFactory01.Contract.MAXFEE(&_StandardTokenFactory01.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_StandardTokenFactory01 *StandardTokenFactory01Caller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StandardTokenFactory01.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_StandardTokenFactory01 *StandardTokenFactory01Session) FeeTo() (common.Address, error) {
	return _StandardTokenFactory01.Contract.FeeTo(&_StandardTokenFactory01.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_StandardTokenFactory01 *StandardTokenFactory01CallerSession) FeeTo() (common.Address, error) {
	return _StandardTokenFactory01.Contract.FeeTo(&_StandardTokenFactory01.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x4acc79ed.
//
// Solidity: function fees(uint256 ) view returns(uint256)
func (_StandardTokenFactory01 *StandardTokenFactory01Caller) Fees(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StandardTokenFactory01.contract.Call(opts, &out, "fees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fees is a free data retrieval call binding the contract method 0x4acc79ed.
//
// Solidity: function fees(uint256 ) view returns(uint256)
func (_StandardTokenFactory01 *StandardTokenFactory01Session) Fees(arg0 *big.Int) (*big.Int, error) {
	return _StandardTokenFactory01.Contract.Fees(&_StandardTokenFactory01.CallOpts, arg0)
}

// Fees is a free data retrieval call binding the contract method 0x4acc79ed.
//
// Solidity: function fees(uint256 ) view returns(uint256)
func (_StandardTokenFactory01 *StandardTokenFactory01CallerSession) Fees(arg0 *big.Int) (*big.Int, error) {
	return _StandardTokenFactory01.Contract.Fees(&_StandardTokenFactory01.CallOpts, arg0)
}

// GetLevels is a free data retrieval call binding the contract method 0x0c394a60.
//
// Solidity: function getLevels() view returns(uint256[])
func (_StandardTokenFactory01 *StandardTokenFactory01Caller) GetLevels(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _StandardTokenFactory01.contract.Call(opts, &out, "getLevels")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetLevels is a free data retrieval call binding the contract method 0x0c394a60.
//
// Solidity: function getLevels() view returns(uint256[])
func (_StandardTokenFactory01 *StandardTokenFactory01Session) GetLevels() ([]*big.Int, error) {
	return _StandardTokenFactory01.Contract.GetLevels(&_StandardTokenFactory01.CallOpts)
}

// GetLevels is a free data retrieval call binding the contract method 0x0c394a60.
//
// Solidity: function getLevels() view returns(uint256[])
func (_StandardTokenFactory01 *StandardTokenFactory01CallerSession) GetLevels() ([]*big.Int, error) {
	return _StandardTokenFactory01.Contract.GetLevels(&_StandardTokenFactory01.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_StandardTokenFactory01 *StandardTokenFactory01Caller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StandardTokenFactory01.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_StandardTokenFactory01 *StandardTokenFactory01Session) Implementation() (common.Address, error) {
	return _StandardTokenFactory01.Contract.Implementation(&_StandardTokenFactory01.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_StandardTokenFactory01 *StandardTokenFactory01CallerSession) Implementation() (common.Address, error) {
	return _StandardTokenFactory01.Contract.Implementation(&_StandardTokenFactory01.CallOpts)
}

// ImplementationVersion is a free data retrieval call binding the contract method 0x06bfcec6.
//
// Solidity: function implementationVersion() view returns(uint96)
func (_StandardTokenFactory01 *StandardTokenFactory01Caller) ImplementationVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StandardTokenFactory01.contract.Call(opts, &out, "implementationVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ImplementationVersion is a free data retrieval call binding the contract method 0x06bfcec6.
//
// Solidity: function implementationVersion() view returns(uint96)
func (_StandardTokenFactory01 *StandardTokenFactory01Session) ImplementationVersion() (*big.Int, error) {
	return _StandardTokenFactory01.Contract.ImplementationVersion(&_StandardTokenFactory01.CallOpts)
}

// ImplementationVersion is a free data retrieval call binding the contract method 0x06bfcec6.
//
// Solidity: function implementationVersion() view returns(uint96)
func (_StandardTokenFactory01 *StandardTokenFactory01CallerSession) ImplementationVersion() (*big.Int, error) {
	return _StandardTokenFactory01.Contract.ImplementationVersion(&_StandardTokenFactory01.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StandardTokenFactory01 *StandardTokenFactory01Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StandardTokenFactory01.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StandardTokenFactory01 *StandardTokenFactory01Session) Owner() (common.Address, error) {
	return _StandardTokenFactory01.Contract.Owner(&_StandardTokenFactory01.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StandardTokenFactory01 *StandardTokenFactory01CallerSession) Owner() (common.Address, error) {
	return _StandardTokenFactory01.Contract.Owner(&_StandardTokenFactory01.CallOpts)
}

// Create is a paid mutator transaction binding the contract method 0xf9acb57f.
//
// Solidity: function create(uint256 level, (string,string,uint8,uint256,string,string,string,string,string,string) tokenInitializeParams) payable returns(address token)
func (_StandardTokenFactory01 *StandardTokenFactory01Transactor) Create(opts *bind.TransactOpts, level *big.Int, tokenInitializeParams TokenInitializeParams) (*types.Transaction, error) {
	return _StandardTokenFactory01.contract.Transact(opts, "create", level, tokenInitializeParams)
}

// Create is a paid mutator transaction binding the contract method 0xf9acb57f.
//
// Solidity: function create(uint256 level, (string,string,uint8,uint256,string,string,string,string,string,string) tokenInitializeParams) payable returns(address token)
func (_StandardTokenFactory01 *StandardTokenFactory01Session) Create(level *big.Int, tokenInitializeParams TokenInitializeParams) (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.Create(&_StandardTokenFactory01.TransactOpts, level, tokenInitializeParams)
}

// Create is a paid mutator transaction binding the contract method 0xf9acb57f.
//
// Solidity: function create(uint256 level, (string,string,uint8,uint256,string,string,string,string,string,string) tokenInitializeParams) payable returns(address token)
func (_StandardTokenFactory01 *StandardTokenFactory01TransactorSession) Create(level *big.Int, tokenInitializeParams TokenInitializeParams) (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.Create(&_StandardTokenFactory01.TransactOpts, level, tokenInitializeParams)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StandardTokenFactory01 *StandardTokenFactory01Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardTokenFactory01.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StandardTokenFactory01 *StandardTokenFactory01Session) RenounceOwnership() (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.RenounceOwnership(&_StandardTokenFactory01.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StandardTokenFactory01 *StandardTokenFactory01TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.RenounceOwnership(&_StandardTokenFactory01.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 level, uint256 fee) returns()
func (_StandardTokenFactory01 *StandardTokenFactory01Transactor) SetFee(opts *bind.TransactOpts, level *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _StandardTokenFactory01.contract.Transact(opts, "setFee", level, fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 level, uint256 fee) returns()
func (_StandardTokenFactory01 *StandardTokenFactory01Session) SetFee(level *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.SetFee(&_StandardTokenFactory01.TransactOpts, level, fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 level, uint256 fee) returns()
func (_StandardTokenFactory01 *StandardTokenFactory01TransactorSession) SetFee(level *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.SetFee(&_StandardTokenFactory01.TransactOpts, level, fee)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address feeTo_) returns()
func (_StandardTokenFactory01 *StandardTokenFactory01Transactor) SetFeeTo(opts *bind.TransactOpts, feeTo_ common.Address) (*types.Transaction, error) {
	return _StandardTokenFactory01.contract.Transact(opts, "setFeeTo", feeTo_)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address feeTo_) returns()
func (_StandardTokenFactory01 *StandardTokenFactory01Session) SetFeeTo(feeTo_ common.Address) (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.SetFeeTo(&_StandardTokenFactory01.TransactOpts, feeTo_)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address feeTo_) returns()
func (_StandardTokenFactory01 *StandardTokenFactory01TransactorSession) SetFeeTo(feeTo_ common.Address) (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.SetFeeTo(&_StandardTokenFactory01.TransactOpts, feeTo_)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xd784d426.
//
// Solidity: function setImplementation(address implementation_) returns()
func (_StandardTokenFactory01 *StandardTokenFactory01Transactor) SetImplementation(opts *bind.TransactOpts, implementation_ common.Address) (*types.Transaction, error) {
	return _StandardTokenFactory01.contract.Transact(opts, "setImplementation", implementation_)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xd784d426.
//
// Solidity: function setImplementation(address implementation_) returns()
func (_StandardTokenFactory01 *StandardTokenFactory01Session) SetImplementation(implementation_ common.Address) (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.SetImplementation(&_StandardTokenFactory01.TransactOpts, implementation_)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xd784d426.
//
// Solidity: function setImplementation(address implementation_) returns()
func (_StandardTokenFactory01 *StandardTokenFactory01TransactorSession) SetImplementation(implementation_ common.Address) (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.SetImplementation(&_StandardTokenFactory01.TransactOpts, implementation_)
}

// SetLevels is a paid mutator transaction binding the contract method 0x6cff29d7.
//
// Solidity: function setLevels(uint256[] _levels) returns()
func (_StandardTokenFactory01 *StandardTokenFactory01Transactor) SetLevels(opts *bind.TransactOpts, _levels []*big.Int) (*types.Transaction, error) {
	return _StandardTokenFactory01.contract.Transact(opts, "setLevels", _levels)
}

// SetLevels is a paid mutator transaction binding the contract method 0x6cff29d7.
//
// Solidity: function setLevels(uint256[] _levels) returns()
func (_StandardTokenFactory01 *StandardTokenFactory01Session) SetLevels(_levels []*big.Int) (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.SetLevels(&_StandardTokenFactory01.TransactOpts, _levels)
}

// SetLevels is a paid mutator transaction binding the contract method 0x6cff29d7.
//
// Solidity: function setLevels(uint256[] _levels) returns()
func (_StandardTokenFactory01 *StandardTokenFactory01TransactorSession) SetLevels(_levels []*big.Int) (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.SetLevels(&_StandardTokenFactory01.TransactOpts, _levels)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StandardTokenFactory01 *StandardTokenFactory01Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StandardTokenFactory01.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StandardTokenFactory01 *StandardTokenFactory01Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.TransferOwnership(&_StandardTokenFactory01.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StandardTokenFactory01 *StandardTokenFactory01TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.TransferOwnership(&_StandardTokenFactory01.TransactOpts, newOwner)
}

// UpdateTokenMetaData is a paid mutator transaction binding the contract method 0xe2a32e3e.
//
// Solidity: function updateTokenMetaData(uint256 level, address token, (string,string,string,string,string,string) tokenMetaData_) payable returns()
func (_StandardTokenFactory01 *StandardTokenFactory01Transactor) UpdateTokenMetaData(opts *bind.TransactOpts, level *big.Int, token common.Address, tokenMetaData_ TokenMetaData) (*types.Transaction, error) {
	return _StandardTokenFactory01.contract.Transact(opts, "updateTokenMetaData", level, token, tokenMetaData_)
}

// UpdateTokenMetaData is a paid mutator transaction binding the contract method 0xe2a32e3e.
//
// Solidity: function updateTokenMetaData(uint256 level, address token, (string,string,string,string,string,string) tokenMetaData_) payable returns()
func (_StandardTokenFactory01 *StandardTokenFactory01Session) UpdateTokenMetaData(level *big.Int, token common.Address, tokenMetaData_ TokenMetaData) (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.UpdateTokenMetaData(&_StandardTokenFactory01.TransactOpts, level, token, tokenMetaData_)
}

// UpdateTokenMetaData is a paid mutator transaction binding the contract method 0xe2a32e3e.
//
// Solidity: function updateTokenMetaData(uint256 level, address token, (string,string,string,string,string,string) tokenMetaData_) payable returns()
func (_StandardTokenFactory01 *StandardTokenFactory01TransactorSession) UpdateTokenMetaData(level *big.Int, token common.Address, tokenMetaData_ TokenMetaData) (*types.Transaction, error) {
	return _StandardTokenFactory01.Contract.UpdateTokenMetaData(&_StandardTokenFactory01.TransactOpts, level, token, tokenMetaData_)
}

// StandardTokenFactory01FeeToUpdatedIterator is returned from FilterFeeToUpdated and is used to iterate over the raw logs and unpacked data for FeeToUpdated events raised by the StandardTokenFactory01 contract.
type StandardTokenFactory01FeeToUpdatedIterator struct {
	Event *StandardTokenFactory01FeeToUpdated // Event containing the contract specifics and raw log

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
func (it *StandardTokenFactory01FeeToUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenFactory01FeeToUpdated)
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
		it.Event = new(StandardTokenFactory01FeeToUpdated)
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
func (it *StandardTokenFactory01FeeToUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenFactory01FeeToUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenFactory01FeeToUpdated represents a FeeToUpdated event raised by the StandardTokenFactory01 contract.
type StandardTokenFactory01FeeToUpdated struct {
	NewFeeTo common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeeToUpdated is a free log retrieval operation binding the contract event 0x9887df018dfa75103faf0f1388028e55a1124783c285d8d2cbcf62a067a6b844.
//
// Solidity: event FeeToUpdated(address newFeeTo)
func (_StandardTokenFactory01 *StandardTokenFactory01Filterer) FilterFeeToUpdated(opts *bind.FilterOpts) (*StandardTokenFactory01FeeToUpdatedIterator, error) {

	logs, sub, err := _StandardTokenFactory01.contract.FilterLogs(opts, "FeeToUpdated")
	if err != nil {
		return nil, err
	}
	return &StandardTokenFactory01FeeToUpdatedIterator{contract: _StandardTokenFactory01.contract, event: "FeeToUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeToUpdated is a free log subscription operation binding the contract event 0x9887df018dfa75103faf0f1388028e55a1124783c285d8d2cbcf62a067a6b844.
//
// Solidity: event FeeToUpdated(address newFeeTo)
func (_StandardTokenFactory01 *StandardTokenFactory01Filterer) WatchFeeToUpdated(opts *bind.WatchOpts, sink chan<- *StandardTokenFactory01FeeToUpdated) (event.Subscription, error) {

	logs, sub, err := _StandardTokenFactory01.contract.WatchLogs(opts, "FeeToUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenFactory01FeeToUpdated)
				if err := _StandardTokenFactory01.contract.UnpackLog(event, "FeeToUpdated", log); err != nil {
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

// ParseFeeToUpdated is a log parse operation binding the contract event 0x9887df018dfa75103faf0f1388028e55a1124783c285d8d2cbcf62a067a6b844.
//
// Solidity: event FeeToUpdated(address newFeeTo)
func (_StandardTokenFactory01 *StandardTokenFactory01Filterer) ParseFeeToUpdated(log types.Log) (*StandardTokenFactory01FeeToUpdated, error) {
	event := new(StandardTokenFactory01FeeToUpdated)
	if err := _StandardTokenFactory01.contract.UnpackLog(event, "FeeToUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardTokenFactory01FeeUpdatedIterator is returned from FilterFeeUpdated and is used to iterate over the raw logs and unpacked data for FeeUpdated events raised by the StandardTokenFactory01 contract.
type StandardTokenFactory01FeeUpdatedIterator struct {
	Event *StandardTokenFactory01FeeUpdated // Event containing the contract specifics and raw log

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
func (it *StandardTokenFactory01FeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenFactory01FeeUpdated)
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
		it.Event = new(StandardTokenFactory01FeeUpdated)
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
func (it *StandardTokenFactory01FeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenFactory01FeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenFactory01FeeUpdated represents a FeeUpdated event raised by the StandardTokenFactory01 contract.
type StandardTokenFactory01FeeUpdated struct {
	Level  *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdated is a free log retrieval operation binding the contract event 0x528d9479e9f9889a87a3c30c7f7ba537e5e59c4c85a37733b16e57c62df61302.
//
// Solidity: event FeeUpdated(uint256 level, uint256 newFee)
func (_StandardTokenFactory01 *StandardTokenFactory01Filterer) FilterFeeUpdated(opts *bind.FilterOpts) (*StandardTokenFactory01FeeUpdatedIterator, error) {

	logs, sub, err := _StandardTokenFactory01.contract.FilterLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return &StandardTokenFactory01FeeUpdatedIterator{contract: _StandardTokenFactory01.contract, event: "FeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeUpdated is a free log subscription operation binding the contract event 0x528d9479e9f9889a87a3c30c7f7ba537e5e59c4c85a37733b16e57c62df61302.
//
// Solidity: event FeeUpdated(uint256 level, uint256 newFee)
func (_StandardTokenFactory01 *StandardTokenFactory01Filterer) WatchFeeUpdated(opts *bind.WatchOpts, sink chan<- *StandardTokenFactory01FeeUpdated) (event.Subscription, error) {

	logs, sub, err := _StandardTokenFactory01.contract.WatchLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenFactory01FeeUpdated)
				if err := _StandardTokenFactory01.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
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

// ParseFeeUpdated is a log parse operation binding the contract event 0x528d9479e9f9889a87a3c30c7f7ba537e5e59c4c85a37733b16e57c62df61302.
//
// Solidity: event FeeUpdated(uint256 level, uint256 newFee)
func (_StandardTokenFactory01 *StandardTokenFactory01Filterer) ParseFeeUpdated(log types.Log) (*StandardTokenFactory01FeeUpdated, error) {
	event := new(StandardTokenFactory01FeeUpdated)
	if err := _StandardTokenFactory01.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardTokenFactory01LevelsUpdatedIterator is returned from FilterLevelsUpdated and is used to iterate over the raw logs and unpacked data for LevelsUpdated events raised by the StandardTokenFactory01 contract.
type StandardTokenFactory01LevelsUpdatedIterator struct {
	Event *StandardTokenFactory01LevelsUpdated // Event containing the contract specifics and raw log

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
func (it *StandardTokenFactory01LevelsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenFactory01LevelsUpdated)
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
		it.Event = new(StandardTokenFactory01LevelsUpdated)
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
func (it *StandardTokenFactory01LevelsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenFactory01LevelsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenFactory01LevelsUpdated represents a LevelsUpdated event raised by the StandardTokenFactory01 contract.
type StandardTokenFactory01LevelsUpdated struct {
	NewLevels []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLevelsUpdated is a free log retrieval operation binding the contract event 0xc83eb89b125b8e65762115d5be003eedcbcbc18d8ccb563044184e1ef45174e6.
//
// Solidity: event LevelsUpdated(uint256[] newLevels)
func (_StandardTokenFactory01 *StandardTokenFactory01Filterer) FilterLevelsUpdated(opts *bind.FilterOpts) (*StandardTokenFactory01LevelsUpdatedIterator, error) {

	logs, sub, err := _StandardTokenFactory01.contract.FilterLogs(opts, "LevelsUpdated")
	if err != nil {
		return nil, err
	}
	return &StandardTokenFactory01LevelsUpdatedIterator{contract: _StandardTokenFactory01.contract, event: "LevelsUpdated", logs: logs, sub: sub}, nil
}

// WatchLevelsUpdated is a free log subscription operation binding the contract event 0xc83eb89b125b8e65762115d5be003eedcbcbc18d8ccb563044184e1ef45174e6.
//
// Solidity: event LevelsUpdated(uint256[] newLevels)
func (_StandardTokenFactory01 *StandardTokenFactory01Filterer) WatchLevelsUpdated(opts *bind.WatchOpts, sink chan<- *StandardTokenFactory01LevelsUpdated) (event.Subscription, error) {

	logs, sub, err := _StandardTokenFactory01.contract.WatchLogs(opts, "LevelsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenFactory01LevelsUpdated)
				if err := _StandardTokenFactory01.contract.UnpackLog(event, "LevelsUpdated", log); err != nil {
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

// ParseLevelsUpdated is a log parse operation binding the contract event 0xc83eb89b125b8e65762115d5be003eedcbcbc18d8ccb563044184e1ef45174e6.
//
// Solidity: event LevelsUpdated(uint256[] newLevels)
func (_StandardTokenFactory01 *StandardTokenFactory01Filterer) ParseLevelsUpdated(log types.Log) (*StandardTokenFactory01LevelsUpdated, error) {
	event := new(StandardTokenFactory01LevelsUpdated)
	if err := _StandardTokenFactory01.contract.UnpackLog(event, "LevelsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardTokenFactory01OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StandardTokenFactory01 contract.
type StandardTokenFactory01OwnershipTransferredIterator struct {
	Event *StandardTokenFactory01OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StandardTokenFactory01OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenFactory01OwnershipTransferred)
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
		it.Event = new(StandardTokenFactory01OwnershipTransferred)
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
func (it *StandardTokenFactory01OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenFactory01OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenFactory01OwnershipTransferred represents a OwnershipTransferred event raised by the StandardTokenFactory01 contract.
type StandardTokenFactory01OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StandardTokenFactory01 *StandardTokenFactory01Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StandardTokenFactory01OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StandardTokenFactory01.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFactory01OwnershipTransferredIterator{contract: _StandardTokenFactory01.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StandardTokenFactory01 *StandardTokenFactory01Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StandardTokenFactory01OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StandardTokenFactory01.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenFactory01OwnershipTransferred)
				if err := _StandardTokenFactory01.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StandardTokenFactory01 *StandardTokenFactory01Filterer) ParseOwnershipTransferred(log types.Log) (*StandardTokenFactory01OwnershipTransferred, error) {
	event := new(StandardTokenFactory01OwnershipTransferred)
	if err := _StandardTokenFactory01.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardTokenFactory01TokenCreatedIterator is returned from FilterTokenCreated and is used to iterate over the raw logs and unpacked data for TokenCreated events raised by the StandardTokenFactory01 contract.
type StandardTokenFactory01TokenCreatedIterator struct {
	Event *StandardTokenFactory01TokenCreated // Event containing the contract specifics and raw log

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
func (it *StandardTokenFactory01TokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenFactory01TokenCreated)
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
		it.Event = new(StandardTokenFactory01TokenCreated)
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
func (it *StandardTokenFactory01TokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenFactory01TokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenFactory01TokenCreated represents a TokenCreated event raised by the StandardTokenFactory01 contract.
type StandardTokenFactory01TokenCreated struct {
	Owner        common.Address
	Token        common.Address
	TokenType    uint8
	TokenVersion *big.Int
	Level        *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenCreated is a free log retrieval operation binding the contract event 0xb18b475b9e42d7af048e225076f59d09cb6f4c4cfcdfe4a22739b0834ccff672.
//
// Solidity: event TokenCreated(address indexed owner, address indexed token, uint8 tokenType, uint96 tokenVersion, uint256 level)
func (_StandardTokenFactory01 *StandardTokenFactory01Filterer) FilterTokenCreated(opts *bind.FilterOpts, owner []common.Address, token []common.Address) (*StandardTokenFactory01TokenCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _StandardTokenFactory01.contract.FilterLogs(opts, "TokenCreated", ownerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFactory01TokenCreatedIterator{contract: _StandardTokenFactory01.contract, event: "TokenCreated", logs: logs, sub: sub}, nil
}

// WatchTokenCreated is a free log subscription operation binding the contract event 0xb18b475b9e42d7af048e225076f59d09cb6f4c4cfcdfe4a22739b0834ccff672.
//
// Solidity: event TokenCreated(address indexed owner, address indexed token, uint8 tokenType, uint96 tokenVersion, uint256 level)
func (_StandardTokenFactory01 *StandardTokenFactory01Filterer) WatchTokenCreated(opts *bind.WatchOpts, sink chan<- *StandardTokenFactory01TokenCreated, owner []common.Address, token []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _StandardTokenFactory01.contract.WatchLogs(opts, "TokenCreated", ownerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenFactory01TokenCreated)
				if err := _StandardTokenFactory01.contract.UnpackLog(event, "TokenCreated", log); err != nil {
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

// ParseTokenCreated is a log parse operation binding the contract event 0xb18b475b9e42d7af048e225076f59d09cb6f4c4cfcdfe4a22739b0834ccff672.
//
// Solidity: event TokenCreated(address indexed owner, address indexed token, uint8 tokenType, uint96 tokenVersion, uint256 level)
func (_StandardTokenFactory01 *StandardTokenFactory01Filterer) ParseTokenCreated(log types.Log) (*StandardTokenFactory01TokenCreated, error) {
	event := new(StandardTokenFactory01TokenCreated)
	if err := _StandardTokenFactory01.contract.UnpackLog(event, "TokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StandardTokenFactory01TokenMetaDataUpdatedIterator is returned from FilterTokenMetaDataUpdated and is used to iterate over the raw logs and unpacked data for TokenMetaDataUpdated events raised by the StandardTokenFactory01 contract.
type StandardTokenFactory01TokenMetaDataUpdatedIterator struct {
	Event *StandardTokenFactory01TokenMetaDataUpdated // Event containing the contract specifics and raw log

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
func (it *StandardTokenFactory01TokenMetaDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenFactory01TokenMetaDataUpdated)
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
		it.Event = new(StandardTokenFactory01TokenMetaDataUpdated)
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
func (it *StandardTokenFactory01TokenMetaDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenFactory01TokenMetaDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenFactory01TokenMetaDataUpdated represents a TokenMetaDataUpdated event raised by the StandardTokenFactory01 contract.
type StandardTokenFactory01TokenMetaDataUpdated struct {
	TokenMetaData TokenMetaData
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenMetaDataUpdated is a free log retrieval operation binding the contract event 0xd9fb697fa8ae059c60cd51834b683372711582a630b9e3beba4d0420e8be170e.
//
// Solidity: event TokenMetaDataUpdated((string,string,string,string,string,string) tokenMetaData)
func (_StandardTokenFactory01 *StandardTokenFactory01Filterer) FilterTokenMetaDataUpdated(opts *bind.FilterOpts) (*StandardTokenFactory01TokenMetaDataUpdatedIterator, error) {

	logs, sub, err := _StandardTokenFactory01.contract.FilterLogs(opts, "TokenMetaDataUpdated")
	if err != nil {
		return nil, err
	}
	return &StandardTokenFactory01TokenMetaDataUpdatedIterator{contract: _StandardTokenFactory01.contract, event: "TokenMetaDataUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenMetaDataUpdated is a free log subscription operation binding the contract event 0xd9fb697fa8ae059c60cd51834b683372711582a630b9e3beba4d0420e8be170e.
//
// Solidity: event TokenMetaDataUpdated((string,string,string,string,string,string) tokenMetaData)
func (_StandardTokenFactory01 *StandardTokenFactory01Filterer) WatchTokenMetaDataUpdated(opts *bind.WatchOpts, sink chan<- *StandardTokenFactory01TokenMetaDataUpdated) (event.Subscription, error) {

	logs, sub, err := _StandardTokenFactory01.contract.WatchLogs(opts, "TokenMetaDataUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenFactory01TokenMetaDataUpdated)
				if err := _StandardTokenFactory01.contract.UnpackLog(event, "TokenMetaDataUpdated", log); err != nil {
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

// ParseTokenMetaDataUpdated is a log parse operation binding the contract event 0xd9fb697fa8ae059c60cd51834b683372711582a630b9e3beba4d0420e8be170e.
//
// Solidity: event TokenMetaDataUpdated((string,string,string,string,string,string) tokenMetaData)
func (_StandardTokenFactory01 *StandardTokenFactory01Filterer) ParseTokenMetaDataUpdated(log types.Log) (*StandardTokenFactory01TokenMetaDataUpdated, error) {
	event := new(StandardTokenFactory01TokenMetaDataUpdated)
	if err := _StandardTokenFactory01.contract.UnpackLog(event, "TokenMetaDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
