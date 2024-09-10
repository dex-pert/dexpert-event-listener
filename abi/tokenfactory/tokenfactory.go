// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenfactory

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

// TokenfactoryMetaData contains all meta data concerning the Tokenfactory contract.
var TokenfactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factoryManager_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feeTo_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxFee_\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"InsufficientFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"InvalidFactoryManager\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"InvalidFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"InvalidFeeReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"factoryManager\",\"type\":\"address\"}],\"name\":\"InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"}],\"name\":\"InvalidLevel\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"InvalidMaxFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newFeeTo\",\"type\":\"address\"}],\"name\":\"FeeToUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"FeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"newLevels\",\"type\":\"uint256[]\"}],\"name\":\"LevelsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"tokenVersion\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"}],\"name\":\"TokenCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logoLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"twitterLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"telegramLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"discordLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"websiteLink\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structTokenMetaData\",\"name\":\"tokenMetaData\",\"type\":\"tuple\"}],\"name\":\"TokenMetaDataUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FACTORY_MANAGER\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logoLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"twitterLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"telegramLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"discordLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"websiteLink\",\"type\":\"string\"}],\"internalType\":\"structTokenInitializeParams\",\"name\":\"tokenInitializeParams\",\"type\":\"tuple\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLevels\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementationVersion\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeTo_\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"}],\"name\":\"setImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_levels\",\"type\":\"uint256[]\"}],\"name\":\"setLevels\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"level\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logoLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"twitterLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"telegramLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"discordLink\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"websiteLink\",\"type\":\"string\"}],\"internalType\":\"structTokenMetaData\",\"name\":\"tokenMetaData_\",\"type\":\"tuple\"}],\"name\":\"updateTokenMetaData\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TokenfactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenfactoryMetaData.ABI instead.
var TokenfactoryABI = TokenfactoryMetaData.ABI

// Tokenfactory is an auto generated Go binding around an Ethereum contract.
type Tokenfactory struct {
	TokenfactoryCaller     // Read-only binding to the contract
	TokenfactoryTransactor // Write-only binding to the contract
	TokenfactoryFilterer   // Log filterer for contract events
}

// TokenfactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenfactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenfactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenfactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenfactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenfactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenfactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenfactorySession struct {
	Contract     *Tokenfactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenfactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenfactoryCallerSession struct {
	Contract *TokenfactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenfactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenfactoryTransactorSession struct {
	Contract     *TokenfactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenfactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenfactoryRaw struct {
	Contract *Tokenfactory // Generic contract binding to access the raw methods on
}

// TokenfactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenfactoryCallerRaw struct {
	Contract *TokenfactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TokenfactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenfactoryTransactorRaw struct {
	Contract *TokenfactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenfactory creates a new instance of Tokenfactory, bound to a specific deployed contract.
func NewTokenfactory(address common.Address, backend bind.ContractBackend) (*Tokenfactory, error) {
	contract, err := bindTokenfactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenfactory{TokenfactoryCaller: TokenfactoryCaller{contract: contract}, TokenfactoryTransactor: TokenfactoryTransactor{contract: contract}, TokenfactoryFilterer: TokenfactoryFilterer{contract: contract}}, nil
}

// NewTokenfactoryCaller creates a new read-only instance of Tokenfactory, bound to a specific deployed contract.
func NewTokenfactoryCaller(address common.Address, caller bind.ContractCaller) (*TokenfactoryCaller, error) {
	contract, err := bindTokenfactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenfactoryCaller{contract: contract}, nil
}

// NewTokenfactoryTransactor creates a new write-only instance of Tokenfactory, bound to a specific deployed contract.
func NewTokenfactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenfactoryTransactor, error) {
	contract, err := bindTokenfactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenfactoryTransactor{contract: contract}, nil
}

// NewTokenfactoryFilterer creates a new log filterer instance of Tokenfactory, bound to a specific deployed contract.
func NewTokenfactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenfactoryFilterer, error) {
	contract, err := bindTokenfactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenfactoryFilterer{contract: contract}, nil
}

// bindTokenfactory binds a generic wrapper to an already deployed contract.
func bindTokenfactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenfactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenfactory *TokenfactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenfactory.Contract.TokenfactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenfactory *TokenfactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenfactory.Contract.TokenfactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenfactory *TokenfactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenfactory.Contract.TokenfactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenfactory *TokenfactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenfactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenfactory *TokenfactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenfactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenfactory *TokenfactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenfactory.Contract.contract.Transact(opts, method, params...)
}

// FACTORYMANAGER is a free data retrieval call binding the contract method 0x42246a57.
//
// Solidity: function FACTORY_MANAGER() view returns(address)
func (_Tokenfactory *TokenfactoryCaller) FACTORYMANAGER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "FACTORY_MANAGER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FACTORYMANAGER is a free data retrieval call binding the contract method 0x42246a57.
//
// Solidity: function FACTORY_MANAGER() view returns(address)
func (_Tokenfactory *TokenfactorySession) FACTORYMANAGER() (common.Address, error) {
	return _Tokenfactory.Contract.FACTORYMANAGER(&_Tokenfactory.CallOpts)
}

// FACTORYMANAGER is a free data retrieval call binding the contract method 0x42246a57.
//
// Solidity: function FACTORY_MANAGER() view returns(address)
func (_Tokenfactory *TokenfactoryCallerSession) FACTORYMANAGER() (common.Address, error) {
	return _Tokenfactory.Contract.FACTORYMANAGER(&_Tokenfactory.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Tokenfactory *TokenfactoryCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Tokenfactory *TokenfactorySession) MAXFEE() (*big.Int, error) {
	return _Tokenfactory.Contract.MAXFEE(&_Tokenfactory.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Tokenfactory *TokenfactoryCallerSession) MAXFEE() (*big.Int, error) {
	return _Tokenfactory.Contract.MAXFEE(&_Tokenfactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Tokenfactory *TokenfactoryCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Tokenfactory *TokenfactorySession) FeeTo() (common.Address, error) {
	return _Tokenfactory.Contract.FeeTo(&_Tokenfactory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Tokenfactory *TokenfactoryCallerSession) FeeTo() (common.Address, error) {
	return _Tokenfactory.Contract.FeeTo(&_Tokenfactory.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x4acc79ed.
//
// Solidity: function fees(uint256 ) view returns(uint256)
func (_Tokenfactory *TokenfactoryCaller) Fees(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "fees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fees is a free data retrieval call binding the contract method 0x4acc79ed.
//
// Solidity: function fees(uint256 ) view returns(uint256)
func (_Tokenfactory *TokenfactorySession) Fees(arg0 *big.Int) (*big.Int, error) {
	return _Tokenfactory.Contract.Fees(&_Tokenfactory.CallOpts, arg0)
}

// Fees is a free data retrieval call binding the contract method 0x4acc79ed.
//
// Solidity: function fees(uint256 ) view returns(uint256)
func (_Tokenfactory *TokenfactoryCallerSession) Fees(arg0 *big.Int) (*big.Int, error) {
	return _Tokenfactory.Contract.Fees(&_Tokenfactory.CallOpts, arg0)
}

// GetLevels is a free data retrieval call binding the contract method 0x0c394a60.
//
// Solidity: function getLevels() view returns(uint256[])
func (_Tokenfactory *TokenfactoryCaller) GetLevels(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "getLevels")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetLevels is a free data retrieval call binding the contract method 0x0c394a60.
//
// Solidity: function getLevels() view returns(uint256[])
func (_Tokenfactory *TokenfactorySession) GetLevels() ([]*big.Int, error) {
	return _Tokenfactory.Contract.GetLevels(&_Tokenfactory.CallOpts)
}

// GetLevels is a free data retrieval call binding the contract method 0x0c394a60.
//
// Solidity: function getLevels() view returns(uint256[])
func (_Tokenfactory *TokenfactoryCallerSession) GetLevels() ([]*big.Int, error) {
	return _Tokenfactory.Contract.GetLevels(&_Tokenfactory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Tokenfactory *TokenfactoryCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Tokenfactory *TokenfactorySession) Implementation() (common.Address, error) {
	return _Tokenfactory.Contract.Implementation(&_Tokenfactory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Tokenfactory *TokenfactoryCallerSession) Implementation() (common.Address, error) {
	return _Tokenfactory.Contract.Implementation(&_Tokenfactory.CallOpts)
}

// ImplementationVersion is a free data retrieval call binding the contract method 0x06bfcec6.
//
// Solidity: function implementationVersion() view returns(uint96)
func (_Tokenfactory *TokenfactoryCaller) ImplementationVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "implementationVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ImplementationVersion is a free data retrieval call binding the contract method 0x06bfcec6.
//
// Solidity: function implementationVersion() view returns(uint96)
func (_Tokenfactory *TokenfactorySession) ImplementationVersion() (*big.Int, error) {
	return _Tokenfactory.Contract.ImplementationVersion(&_Tokenfactory.CallOpts)
}

// ImplementationVersion is a free data retrieval call binding the contract method 0x06bfcec6.
//
// Solidity: function implementationVersion() view returns(uint96)
func (_Tokenfactory *TokenfactoryCallerSession) ImplementationVersion() (*big.Int, error) {
	return _Tokenfactory.Contract.ImplementationVersion(&_Tokenfactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenfactory *TokenfactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenfactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenfactory *TokenfactorySession) Owner() (common.Address, error) {
	return _Tokenfactory.Contract.Owner(&_Tokenfactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tokenfactory *TokenfactoryCallerSession) Owner() (common.Address, error) {
	return _Tokenfactory.Contract.Owner(&_Tokenfactory.CallOpts)
}

// Create is a paid mutator transaction binding the contract method 0xf9acb57f.
//
// Solidity: function create(uint256 level, (string,string,uint8,uint256,string,string,string,string,string,string) tokenInitializeParams) payable returns(address token)
func (_Tokenfactory *TokenfactoryTransactor) Create(opts *bind.TransactOpts, level *big.Int, tokenInitializeParams TokenInitializeParams) (*types.Transaction, error) {
	return _Tokenfactory.contract.Transact(opts, "create", level, tokenInitializeParams)
}

// Create is a paid mutator transaction binding the contract method 0xf9acb57f.
//
// Solidity: function create(uint256 level, (string,string,uint8,uint256,string,string,string,string,string,string) tokenInitializeParams) payable returns(address token)
func (_Tokenfactory *TokenfactorySession) Create(level *big.Int, tokenInitializeParams TokenInitializeParams) (*types.Transaction, error) {
	return _Tokenfactory.Contract.Create(&_Tokenfactory.TransactOpts, level, tokenInitializeParams)
}

// Create is a paid mutator transaction binding the contract method 0xf9acb57f.
//
// Solidity: function create(uint256 level, (string,string,uint8,uint256,string,string,string,string,string,string) tokenInitializeParams) payable returns(address token)
func (_Tokenfactory *TokenfactoryTransactorSession) Create(level *big.Int, tokenInitializeParams TokenInitializeParams) (*types.Transaction, error) {
	return _Tokenfactory.Contract.Create(&_Tokenfactory.TransactOpts, level, tokenInitializeParams)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenfactory *TokenfactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenfactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenfactory *TokenfactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokenfactory.Contract.RenounceOwnership(&_Tokenfactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenfactory *TokenfactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokenfactory.Contract.RenounceOwnership(&_Tokenfactory.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 level, uint256 fee) returns()
func (_Tokenfactory *TokenfactoryTransactor) SetFee(opts *bind.TransactOpts, level *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Tokenfactory.contract.Transact(opts, "setFee", level, fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 level, uint256 fee) returns()
func (_Tokenfactory *TokenfactorySession) SetFee(level *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Tokenfactory.Contract.SetFee(&_Tokenfactory.TransactOpts, level, fee)
}

// SetFee is a paid mutator transaction binding the contract method 0x52f7c988.
//
// Solidity: function setFee(uint256 level, uint256 fee) returns()
func (_Tokenfactory *TokenfactoryTransactorSession) SetFee(level *big.Int, fee *big.Int) (*types.Transaction, error) {
	return _Tokenfactory.Contract.SetFee(&_Tokenfactory.TransactOpts, level, fee)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address feeTo_) returns()
func (_Tokenfactory *TokenfactoryTransactor) SetFeeTo(opts *bind.TransactOpts, feeTo_ common.Address) (*types.Transaction, error) {
	return _Tokenfactory.contract.Transact(opts, "setFeeTo", feeTo_)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address feeTo_) returns()
func (_Tokenfactory *TokenfactorySession) SetFeeTo(feeTo_ common.Address) (*types.Transaction, error) {
	return _Tokenfactory.Contract.SetFeeTo(&_Tokenfactory.TransactOpts, feeTo_)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address feeTo_) returns()
func (_Tokenfactory *TokenfactoryTransactorSession) SetFeeTo(feeTo_ common.Address) (*types.Transaction, error) {
	return _Tokenfactory.Contract.SetFeeTo(&_Tokenfactory.TransactOpts, feeTo_)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xd784d426.
//
// Solidity: function setImplementation(address implementation_) returns()
func (_Tokenfactory *TokenfactoryTransactor) SetImplementation(opts *bind.TransactOpts, implementation_ common.Address) (*types.Transaction, error) {
	return _Tokenfactory.contract.Transact(opts, "setImplementation", implementation_)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xd784d426.
//
// Solidity: function setImplementation(address implementation_) returns()
func (_Tokenfactory *TokenfactorySession) SetImplementation(implementation_ common.Address) (*types.Transaction, error) {
	return _Tokenfactory.Contract.SetImplementation(&_Tokenfactory.TransactOpts, implementation_)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xd784d426.
//
// Solidity: function setImplementation(address implementation_) returns()
func (_Tokenfactory *TokenfactoryTransactorSession) SetImplementation(implementation_ common.Address) (*types.Transaction, error) {
	return _Tokenfactory.Contract.SetImplementation(&_Tokenfactory.TransactOpts, implementation_)
}

// SetLevels is a paid mutator transaction binding the contract method 0x6cff29d7.
//
// Solidity: function setLevels(uint256[] _levels) returns()
func (_Tokenfactory *TokenfactoryTransactor) SetLevels(opts *bind.TransactOpts, _levels []*big.Int) (*types.Transaction, error) {
	return _Tokenfactory.contract.Transact(opts, "setLevels", _levels)
}

// SetLevels is a paid mutator transaction binding the contract method 0x6cff29d7.
//
// Solidity: function setLevels(uint256[] _levels) returns()
func (_Tokenfactory *TokenfactorySession) SetLevels(_levels []*big.Int) (*types.Transaction, error) {
	return _Tokenfactory.Contract.SetLevels(&_Tokenfactory.TransactOpts, _levels)
}

// SetLevels is a paid mutator transaction binding the contract method 0x6cff29d7.
//
// Solidity: function setLevels(uint256[] _levels) returns()
func (_Tokenfactory *TokenfactoryTransactorSession) SetLevels(_levels []*big.Int) (*types.Transaction, error) {
	return _Tokenfactory.Contract.SetLevels(&_Tokenfactory.TransactOpts, _levels)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenfactory *TokenfactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tokenfactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenfactory *TokenfactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokenfactory.Contract.TransferOwnership(&_Tokenfactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tokenfactory *TokenfactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokenfactory.Contract.TransferOwnership(&_Tokenfactory.TransactOpts, newOwner)
}

// UpdateTokenMetaData is a paid mutator transaction binding the contract method 0xe2a32e3e.
//
// Solidity: function updateTokenMetaData(uint256 level, address token, (string,string,string,string,string,string) tokenMetaData_) payable returns()
func (_Tokenfactory *TokenfactoryTransactor) UpdateTokenMetaData(opts *bind.TransactOpts, level *big.Int, token common.Address, tokenMetaData_ TokenMetaData) (*types.Transaction, error) {
	return _Tokenfactory.contract.Transact(opts, "updateTokenMetaData", level, token, tokenMetaData_)
}

// UpdateTokenMetaData is a paid mutator transaction binding the contract method 0xe2a32e3e.
//
// Solidity: function updateTokenMetaData(uint256 level, address token, (string,string,string,string,string,string) tokenMetaData_) payable returns()
func (_Tokenfactory *TokenfactorySession) UpdateTokenMetaData(level *big.Int, token common.Address, tokenMetaData_ TokenMetaData) (*types.Transaction, error) {
	return _Tokenfactory.Contract.UpdateTokenMetaData(&_Tokenfactory.TransactOpts, level, token, tokenMetaData_)
}

// UpdateTokenMetaData is a paid mutator transaction binding the contract method 0xe2a32e3e.
//
// Solidity: function updateTokenMetaData(uint256 level, address token, (string,string,string,string,string,string) tokenMetaData_) payable returns()
func (_Tokenfactory *TokenfactoryTransactorSession) UpdateTokenMetaData(level *big.Int, token common.Address, tokenMetaData_ TokenMetaData) (*types.Transaction, error) {
	return _Tokenfactory.Contract.UpdateTokenMetaData(&_Tokenfactory.TransactOpts, level, token, tokenMetaData_)
}

// TokenfactoryFeeToUpdatedIterator is returned from FilterFeeToUpdated and is used to iterate over the raw logs and unpacked data for FeeToUpdated events raised by the Tokenfactory contract.
type TokenfactoryFeeToUpdatedIterator struct {
	Event *TokenfactoryFeeToUpdated // Event containing the contract specifics and raw log

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
func (it *TokenfactoryFeeToUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenfactoryFeeToUpdated)
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
		it.Event = new(TokenfactoryFeeToUpdated)
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
func (it *TokenfactoryFeeToUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenfactoryFeeToUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenfactoryFeeToUpdated represents a FeeToUpdated event raised by the Tokenfactory contract.
type TokenfactoryFeeToUpdated struct {
	NewFeeTo common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFeeToUpdated is a free log retrieval operation binding the contract event 0x9887df018dfa75103faf0f1388028e55a1124783c285d8d2cbcf62a067a6b844.
//
// Solidity: event FeeToUpdated(address newFeeTo)
func (_Tokenfactory *TokenfactoryFilterer) FilterFeeToUpdated(opts *bind.FilterOpts) (*TokenfactoryFeeToUpdatedIterator, error) {

	logs, sub, err := _Tokenfactory.contract.FilterLogs(opts, "FeeToUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenfactoryFeeToUpdatedIterator{contract: _Tokenfactory.contract, event: "FeeToUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeToUpdated is a free log subscription operation binding the contract event 0x9887df018dfa75103faf0f1388028e55a1124783c285d8d2cbcf62a067a6b844.
//
// Solidity: event FeeToUpdated(address newFeeTo)
func (_Tokenfactory *TokenfactoryFilterer) WatchFeeToUpdated(opts *bind.WatchOpts, sink chan<- *TokenfactoryFeeToUpdated) (event.Subscription, error) {

	logs, sub, err := _Tokenfactory.contract.WatchLogs(opts, "FeeToUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenfactoryFeeToUpdated)
				if err := _Tokenfactory.contract.UnpackLog(event, "FeeToUpdated", log); err != nil {
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
func (_Tokenfactory *TokenfactoryFilterer) ParseFeeToUpdated(log types.Log) (*TokenfactoryFeeToUpdated, error) {
	event := new(TokenfactoryFeeToUpdated)
	if err := _Tokenfactory.contract.UnpackLog(event, "FeeToUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenfactoryFeeUpdatedIterator is returned from FilterFeeUpdated and is used to iterate over the raw logs and unpacked data for FeeUpdated events raised by the Tokenfactory contract.
type TokenfactoryFeeUpdatedIterator struct {
	Event *TokenfactoryFeeUpdated // Event containing the contract specifics and raw log

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
func (it *TokenfactoryFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenfactoryFeeUpdated)
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
		it.Event = new(TokenfactoryFeeUpdated)
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
func (it *TokenfactoryFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenfactoryFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenfactoryFeeUpdated represents a FeeUpdated event raised by the Tokenfactory contract.
type TokenfactoryFeeUpdated struct {
	Level  *big.Int
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeUpdated is a free log retrieval operation binding the contract event 0x528d9479e9f9889a87a3c30c7f7ba537e5e59c4c85a37733b16e57c62df61302.
//
// Solidity: event FeeUpdated(uint256 level, uint256 newFee)
func (_Tokenfactory *TokenfactoryFilterer) FilterFeeUpdated(opts *bind.FilterOpts) (*TokenfactoryFeeUpdatedIterator, error) {

	logs, sub, err := _Tokenfactory.contract.FilterLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenfactoryFeeUpdatedIterator{contract: _Tokenfactory.contract, event: "FeeUpdated", logs: logs, sub: sub}, nil
}

// WatchFeeUpdated is a free log subscription operation binding the contract event 0x528d9479e9f9889a87a3c30c7f7ba537e5e59c4c85a37733b16e57c62df61302.
//
// Solidity: event FeeUpdated(uint256 level, uint256 newFee)
func (_Tokenfactory *TokenfactoryFilterer) WatchFeeUpdated(opts *bind.WatchOpts, sink chan<- *TokenfactoryFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Tokenfactory.contract.WatchLogs(opts, "FeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenfactoryFeeUpdated)
				if err := _Tokenfactory.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
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
func (_Tokenfactory *TokenfactoryFilterer) ParseFeeUpdated(log types.Log) (*TokenfactoryFeeUpdated, error) {
	event := new(TokenfactoryFeeUpdated)
	if err := _Tokenfactory.contract.UnpackLog(event, "FeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenfactoryLevelsUpdatedIterator is returned from FilterLevelsUpdated and is used to iterate over the raw logs and unpacked data for LevelsUpdated events raised by the Tokenfactory contract.
type TokenfactoryLevelsUpdatedIterator struct {
	Event *TokenfactoryLevelsUpdated // Event containing the contract specifics and raw log

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
func (it *TokenfactoryLevelsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenfactoryLevelsUpdated)
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
		it.Event = new(TokenfactoryLevelsUpdated)
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
func (it *TokenfactoryLevelsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenfactoryLevelsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenfactoryLevelsUpdated represents a LevelsUpdated event raised by the Tokenfactory contract.
type TokenfactoryLevelsUpdated struct {
	NewLevels []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLevelsUpdated is a free log retrieval operation binding the contract event 0xc83eb89b125b8e65762115d5be003eedcbcbc18d8ccb563044184e1ef45174e6.
//
// Solidity: event LevelsUpdated(uint256[] newLevels)
func (_Tokenfactory *TokenfactoryFilterer) FilterLevelsUpdated(opts *bind.FilterOpts) (*TokenfactoryLevelsUpdatedIterator, error) {

	logs, sub, err := _Tokenfactory.contract.FilterLogs(opts, "LevelsUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenfactoryLevelsUpdatedIterator{contract: _Tokenfactory.contract, event: "LevelsUpdated", logs: logs, sub: sub}, nil
}

// WatchLevelsUpdated is a free log subscription operation binding the contract event 0xc83eb89b125b8e65762115d5be003eedcbcbc18d8ccb563044184e1ef45174e6.
//
// Solidity: event LevelsUpdated(uint256[] newLevels)
func (_Tokenfactory *TokenfactoryFilterer) WatchLevelsUpdated(opts *bind.WatchOpts, sink chan<- *TokenfactoryLevelsUpdated) (event.Subscription, error) {

	logs, sub, err := _Tokenfactory.contract.WatchLogs(opts, "LevelsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenfactoryLevelsUpdated)
				if err := _Tokenfactory.contract.UnpackLog(event, "LevelsUpdated", log); err != nil {
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
func (_Tokenfactory *TokenfactoryFilterer) ParseLevelsUpdated(log types.Log) (*TokenfactoryLevelsUpdated, error) {
	event := new(TokenfactoryLevelsUpdated)
	if err := _Tokenfactory.contract.UnpackLog(event, "LevelsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenfactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tokenfactory contract.
type TokenfactoryOwnershipTransferredIterator struct {
	Event *TokenfactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenfactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenfactoryOwnershipTransferred)
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
		it.Event = new(TokenfactoryOwnershipTransferred)
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
func (it *TokenfactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenfactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenfactoryOwnershipTransferred represents a OwnershipTransferred event raised by the Tokenfactory contract.
type TokenfactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokenfactory *TokenfactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenfactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokenfactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenfactoryOwnershipTransferredIterator{contract: _Tokenfactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tokenfactory *TokenfactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenfactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokenfactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenfactoryOwnershipTransferred)
				if err := _Tokenfactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Tokenfactory *TokenfactoryFilterer) ParseOwnershipTransferred(log types.Log) (*TokenfactoryOwnershipTransferred, error) {
	event := new(TokenfactoryOwnershipTransferred)
	if err := _Tokenfactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenfactoryTokenCreatedIterator is returned from FilterTokenCreated and is used to iterate over the raw logs and unpacked data for TokenCreated events raised by the Tokenfactory contract.
type TokenfactoryTokenCreatedIterator struct {
	Event *TokenfactoryTokenCreated // Event containing the contract specifics and raw log

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
func (it *TokenfactoryTokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenfactoryTokenCreated)
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
		it.Event = new(TokenfactoryTokenCreated)
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
func (it *TokenfactoryTokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenfactoryTokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenfactoryTokenCreated represents a TokenCreated event raised by the Tokenfactory contract.
type TokenfactoryTokenCreated struct {
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
func (_Tokenfactory *TokenfactoryFilterer) FilterTokenCreated(opts *bind.FilterOpts, owner []common.Address, token []common.Address) (*TokenfactoryTokenCreatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Tokenfactory.contract.FilterLogs(opts, "TokenCreated", ownerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenfactoryTokenCreatedIterator{contract: _Tokenfactory.contract, event: "TokenCreated", logs: logs, sub: sub}, nil
}

// WatchTokenCreated is a free log subscription operation binding the contract event 0xb18b475b9e42d7af048e225076f59d09cb6f4c4cfcdfe4a22739b0834ccff672.
//
// Solidity: event TokenCreated(address indexed owner, address indexed token, uint8 tokenType, uint96 tokenVersion, uint256 level)
func (_Tokenfactory *TokenfactoryFilterer) WatchTokenCreated(opts *bind.WatchOpts, sink chan<- *TokenfactoryTokenCreated, owner []common.Address, token []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Tokenfactory.contract.WatchLogs(opts, "TokenCreated", ownerRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenfactoryTokenCreated)
				if err := _Tokenfactory.contract.UnpackLog(event, "TokenCreated", log); err != nil {
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
func (_Tokenfactory *TokenfactoryFilterer) ParseTokenCreated(log types.Log) (*TokenfactoryTokenCreated, error) {
	event := new(TokenfactoryTokenCreated)
	if err := _Tokenfactory.contract.UnpackLog(event, "TokenCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenfactoryTokenMetaDataUpdatedIterator is returned from FilterTokenMetaDataUpdated and is used to iterate over the raw logs and unpacked data for TokenMetaDataUpdated events raised by the Tokenfactory contract.
type TokenfactoryTokenMetaDataUpdatedIterator struct {
	Event *TokenfactoryTokenMetaDataUpdated // Event containing the contract specifics and raw log

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
func (it *TokenfactoryTokenMetaDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenfactoryTokenMetaDataUpdated)
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
		it.Event = new(TokenfactoryTokenMetaDataUpdated)
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
func (it *TokenfactoryTokenMetaDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenfactoryTokenMetaDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenfactoryTokenMetaDataUpdated represents a TokenMetaDataUpdated event raised by the Tokenfactory contract.
type TokenfactoryTokenMetaDataUpdated struct {
	TokenMetaData TokenMetaData
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenMetaDataUpdated is a free log retrieval operation binding the contract event 0xd9fb697fa8ae059c60cd51834b683372711582a630b9e3beba4d0420e8be170e.
//
// Solidity: event TokenMetaDataUpdated((string,string,string,string,string,string) tokenMetaData)
func (_Tokenfactory *TokenfactoryFilterer) FilterTokenMetaDataUpdated(opts *bind.FilterOpts) (*TokenfactoryTokenMetaDataUpdatedIterator, error) {

	logs, sub, err := _Tokenfactory.contract.FilterLogs(opts, "TokenMetaDataUpdated")
	if err != nil {
		return nil, err
	}
	return &TokenfactoryTokenMetaDataUpdatedIterator{contract: _Tokenfactory.contract, event: "TokenMetaDataUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenMetaDataUpdated is a free log subscription operation binding the contract event 0xd9fb697fa8ae059c60cd51834b683372711582a630b9e3beba4d0420e8be170e.
//
// Solidity: event TokenMetaDataUpdated((string,string,string,string,string,string) tokenMetaData)
func (_Tokenfactory *TokenfactoryFilterer) WatchTokenMetaDataUpdated(opts *bind.WatchOpts, sink chan<- *TokenfactoryTokenMetaDataUpdated) (event.Subscription, error) {

	logs, sub, err := _Tokenfactory.contract.WatchLogs(opts, "TokenMetaDataUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenfactoryTokenMetaDataUpdated)
				if err := _Tokenfactory.contract.UnpackLog(event, "TokenMetaDataUpdated", log); err != nil {
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
func (_Tokenfactory *TokenfactoryFilterer) ParseTokenMetaDataUpdated(log types.Log) (*TokenfactoryTokenMetaDataUpdated, error) {
	event := new(TokenfactoryTokenMetaDataUpdated)
	if err := _Tokenfactory.contract.UnpackLog(event, "TokenMetaDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
