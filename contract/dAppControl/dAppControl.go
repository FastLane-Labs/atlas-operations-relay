// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dAppControl

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

// DAppConfig is an auto generated low-level Go binding around an user-defined struct.
type DAppConfig struct {
	To         common.Address
	CallConfig uint32
	BidToken   common.Address
}

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	From         common.Address
	To           common.Address
	Value        *big.Int
	Gas          *big.Int
	MaxFeePerGas *big.Int
	Nonce        *big.Int
	Deadline     *big.Int
	Dapp         common.Address
	Control      common.Address
	SessionKey   common.Address
	Data         []byte
	Signature    []byte
}

// DAppControlMetaData contains all meta data concerning the DAppControl contract.
var DAppControlMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getDAppConfig\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structUserOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"dConfig\",\"type\":\"tuple\",\"internalType\":\"structDAppConfig\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"preOpsCall\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structUserOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"}]",
}

// DAppControlABI is the input ABI used to generate the binding from.
// Deprecated: Use DAppControlMetaData.ABI instead.
var DAppControlABI = DAppControlMetaData.ABI

// DAppControl is an auto generated Go binding around an Ethereum contract.
type DAppControl struct {
	DAppControlCaller     // Read-only binding to the contract
	DAppControlTransactor // Write-only binding to the contract
	DAppControlFilterer   // Log filterer for contract events
}

// DAppControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type DAppControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAppControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DAppControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAppControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DAppControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DAppControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DAppControlSession struct {
	Contract     *DAppControl      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DAppControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DAppControlCallerSession struct {
	Contract *DAppControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DAppControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DAppControlTransactorSession struct {
	Contract     *DAppControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DAppControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type DAppControlRaw struct {
	Contract *DAppControl // Generic contract binding to access the raw methods on
}

// DAppControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DAppControlCallerRaw struct {
	Contract *DAppControlCaller // Generic read-only contract binding to access the raw methods on
}

// DAppControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DAppControlTransactorRaw struct {
	Contract *DAppControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDAppControl creates a new instance of DAppControl, bound to a specific deployed contract.
func NewDAppControl(address common.Address, backend bind.ContractBackend) (*DAppControl, error) {
	contract, err := bindDAppControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DAppControl{DAppControlCaller: DAppControlCaller{contract: contract}, DAppControlTransactor: DAppControlTransactor{contract: contract}, DAppControlFilterer: DAppControlFilterer{contract: contract}}, nil
}

// NewDAppControlCaller creates a new read-only instance of DAppControl, bound to a specific deployed contract.
func NewDAppControlCaller(address common.Address, caller bind.ContractCaller) (*DAppControlCaller, error) {
	contract, err := bindDAppControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DAppControlCaller{contract: contract}, nil
}

// NewDAppControlTransactor creates a new write-only instance of DAppControl, bound to a specific deployed contract.
func NewDAppControlTransactor(address common.Address, transactor bind.ContractTransactor) (*DAppControlTransactor, error) {
	contract, err := bindDAppControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DAppControlTransactor{contract: contract}, nil
}

// NewDAppControlFilterer creates a new log filterer instance of DAppControl, bound to a specific deployed contract.
func NewDAppControlFilterer(address common.Address, filterer bind.ContractFilterer) (*DAppControlFilterer, error) {
	contract, err := bindDAppControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DAppControlFilterer{contract: contract}, nil
}

// bindDAppControl binds a generic wrapper to an already deployed contract.
func bindDAppControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DAppControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAppControl *DAppControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAppControl.Contract.DAppControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAppControl *DAppControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAppControl.Contract.DAppControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAppControl *DAppControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAppControl.Contract.DAppControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DAppControl *DAppControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DAppControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DAppControl *DAppControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DAppControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DAppControl *DAppControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DAppControl.Contract.contract.Transact(opts, method, params...)
}

// GetDAppConfig is a free data retrieval call binding the contract method 0x845ccda0.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp) view returns((address,uint32,address) dConfig)
func (_DAppControl *DAppControlCaller) GetDAppConfig(opts *bind.CallOpts, userOp UserOperation) (DAppConfig, error) {
	var out []interface{}
	err := _DAppControl.contract.Call(opts, &out, "getDAppConfig", userOp)

	if err != nil {
		return *new(DAppConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(DAppConfig)).(*DAppConfig)

	return out0, err

}

// GetDAppConfig is a free data retrieval call binding the contract method 0x845ccda0.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp) view returns((address,uint32,address) dConfig)
func (_DAppControl *DAppControlSession) GetDAppConfig(userOp UserOperation) (DAppConfig, error) {
	return _DAppControl.Contract.GetDAppConfig(&_DAppControl.CallOpts, userOp)
}

// GetDAppConfig is a free data retrieval call binding the contract method 0x845ccda0.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp) view returns((address,uint32,address) dConfig)
func (_DAppControl *DAppControlCallerSession) GetDAppConfig(userOp UserOperation) (DAppConfig, error) {
	return _DAppControl.Contract.GetDAppConfig(&_DAppControl.CallOpts, userOp)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0x0c046445.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp) payable returns(bytes)
func (_DAppControl *DAppControlTransactor) PreOpsCall(opts *bind.TransactOpts, userOp UserOperation) (*types.Transaction, error) {
	return _DAppControl.contract.Transact(opts, "preOpsCall", userOp)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0x0c046445.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp) payable returns(bytes)
func (_DAppControl *DAppControlSession) PreOpsCall(userOp UserOperation) (*types.Transaction, error) {
	return _DAppControl.Contract.PreOpsCall(&_DAppControl.TransactOpts, userOp)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0x0c046445.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp) payable returns(bytes)
func (_DAppControl *DAppControlTransactorSession) PreOpsCall(userOp UserOperation) (*types.Transaction, error) {
	return _DAppControl.Contract.PreOpsCall(&_DAppControl.TransactOpts, userOp)
}
