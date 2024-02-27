// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package atlETH

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

// AtlETHMetaData contains all meta data concerning the AtlETH contract.
var AtlETHMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"balanceOfBonded\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"}]",
}

// AtlETHABI is the input ABI used to generate the binding from.
// Deprecated: Use AtlETHMetaData.ABI instead.
var AtlETHABI = AtlETHMetaData.ABI

// AtlETH is an auto generated Go binding around an Ethereum contract.
type AtlETH struct {
	AtlETHCaller     // Read-only binding to the contract
	AtlETHTransactor // Write-only binding to the contract
	AtlETHFilterer   // Log filterer for contract events
}

// AtlETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type AtlETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtlETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AtlETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtlETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AtlETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtlETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AtlETHSession struct {
	Contract     *AtlETH           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AtlETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AtlETHCallerSession struct {
	Contract *AtlETHCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AtlETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AtlETHTransactorSession struct {
	Contract     *AtlETHTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AtlETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type AtlETHRaw struct {
	Contract *AtlETH // Generic contract binding to access the raw methods on
}

// AtlETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AtlETHCallerRaw struct {
	Contract *AtlETHCaller // Generic read-only contract binding to access the raw methods on
}

// AtlETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AtlETHTransactorRaw struct {
	Contract *AtlETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAtlETH creates a new instance of AtlETH, bound to a specific deployed contract.
func NewAtlETH(address common.Address, backend bind.ContractBackend) (*AtlETH, error) {
	contract, err := bindAtlETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AtlETH{AtlETHCaller: AtlETHCaller{contract: contract}, AtlETHTransactor: AtlETHTransactor{contract: contract}, AtlETHFilterer: AtlETHFilterer{contract: contract}}, nil
}

// NewAtlETHCaller creates a new read-only instance of AtlETH, bound to a specific deployed contract.
func NewAtlETHCaller(address common.Address, caller bind.ContractCaller) (*AtlETHCaller, error) {
	contract, err := bindAtlETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AtlETHCaller{contract: contract}, nil
}

// NewAtlETHTransactor creates a new write-only instance of AtlETH, bound to a specific deployed contract.
func NewAtlETHTransactor(address common.Address, transactor bind.ContractTransactor) (*AtlETHTransactor, error) {
	contract, err := bindAtlETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AtlETHTransactor{contract: contract}, nil
}

// NewAtlETHFilterer creates a new log filterer instance of AtlETH, bound to a specific deployed contract.
func NewAtlETHFilterer(address common.Address, filterer bind.ContractFilterer) (*AtlETHFilterer, error) {
	contract, err := bindAtlETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AtlETHFilterer{contract: contract}, nil
}

// bindAtlETH binds a generic wrapper to an already deployed contract.
func bindAtlETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AtlETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AtlETH *AtlETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AtlETH.Contract.AtlETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AtlETH *AtlETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AtlETH.Contract.AtlETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AtlETH *AtlETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AtlETH.Contract.AtlETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AtlETH *AtlETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AtlETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AtlETH *AtlETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AtlETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AtlETH *AtlETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AtlETH.Contract.contract.Transact(opts, method, params...)
}

// BalanceOfBonded is a free data retrieval call binding the contract method 0x825ad607.
//
// Solidity: function balanceOfBonded(address account) view returns(uint256)
func (_AtlETH *AtlETHCaller) BalanceOfBonded(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AtlETH.contract.Call(opts, &out, "balanceOfBonded", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfBonded is a free data retrieval call binding the contract method 0x825ad607.
//
// Solidity: function balanceOfBonded(address account) view returns(uint256)
func (_AtlETH *AtlETHSession) BalanceOfBonded(account common.Address) (*big.Int, error) {
	return _AtlETH.Contract.BalanceOfBonded(&_AtlETH.CallOpts, account)
}

// BalanceOfBonded is a free data retrieval call binding the contract method 0x825ad607.
//
// Solidity: function balanceOfBonded(address account) view returns(uint256)
func (_AtlETH *AtlETHCallerSession) BalanceOfBonded(account common.Address) (*big.Int, error) {
	return _AtlETH.Contract.BalanceOfBonded(&_AtlETH.CallOpts, account)
}
