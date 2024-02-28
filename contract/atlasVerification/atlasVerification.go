// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package atlasVerification

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

// AtlasVerificationMetaData contains all meta data concerning the AtlasVerification contract.
var AtlasVerificationMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"getDomainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"domainSeparator\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDAppSignatories\",\"inputs\":[{\"name\":\"dAppControl\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"}]",
}

// AtlasVerificationABI is the input ABI used to generate the binding from.
// Deprecated: Use AtlasVerificationMetaData.ABI instead.
var AtlasVerificationABI = AtlasVerificationMetaData.ABI

// AtlasVerification is an auto generated Go binding around an Ethereum contract.
type AtlasVerification struct {
	AtlasVerificationCaller     // Read-only binding to the contract
	AtlasVerificationTransactor // Write-only binding to the contract
	AtlasVerificationFilterer   // Log filterer for contract events
}

// AtlasVerificationCaller is an auto generated read-only Go binding around an Ethereum contract.
type AtlasVerificationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtlasVerificationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AtlasVerificationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtlasVerificationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AtlasVerificationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtlasVerificationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AtlasVerificationSession struct {
	Contract     *AtlasVerification // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AtlasVerificationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AtlasVerificationCallerSession struct {
	Contract *AtlasVerificationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// AtlasVerificationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AtlasVerificationTransactorSession struct {
	Contract     *AtlasVerificationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AtlasVerificationRaw is an auto generated low-level Go binding around an Ethereum contract.
type AtlasVerificationRaw struct {
	Contract *AtlasVerification // Generic contract binding to access the raw methods on
}

// AtlasVerificationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AtlasVerificationCallerRaw struct {
	Contract *AtlasVerificationCaller // Generic read-only contract binding to access the raw methods on
}

// AtlasVerificationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AtlasVerificationTransactorRaw struct {
	Contract *AtlasVerificationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAtlasVerification creates a new instance of AtlasVerification, bound to a specific deployed contract.
func NewAtlasVerification(address common.Address, backend bind.ContractBackend) (*AtlasVerification, error) {
	contract, err := bindAtlasVerification(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AtlasVerification{AtlasVerificationCaller: AtlasVerificationCaller{contract: contract}, AtlasVerificationTransactor: AtlasVerificationTransactor{contract: contract}, AtlasVerificationFilterer: AtlasVerificationFilterer{contract: contract}}, nil
}

// NewAtlasVerificationCaller creates a new read-only instance of AtlasVerification, bound to a specific deployed contract.
func NewAtlasVerificationCaller(address common.Address, caller bind.ContractCaller) (*AtlasVerificationCaller, error) {
	contract, err := bindAtlasVerification(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AtlasVerificationCaller{contract: contract}, nil
}

// NewAtlasVerificationTransactor creates a new write-only instance of AtlasVerification, bound to a specific deployed contract.
func NewAtlasVerificationTransactor(address common.Address, transactor bind.ContractTransactor) (*AtlasVerificationTransactor, error) {
	contract, err := bindAtlasVerification(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AtlasVerificationTransactor{contract: contract}, nil
}

// NewAtlasVerificationFilterer creates a new log filterer instance of AtlasVerification, bound to a specific deployed contract.
func NewAtlasVerificationFilterer(address common.Address, filterer bind.ContractFilterer) (*AtlasVerificationFilterer, error) {
	contract, err := bindAtlasVerification(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AtlasVerificationFilterer{contract: contract}, nil
}

// bindAtlasVerification binds a generic wrapper to an already deployed contract.
func bindAtlasVerification(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AtlasVerificationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AtlasVerification *AtlasVerificationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AtlasVerification.Contract.AtlasVerificationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AtlasVerification *AtlasVerificationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AtlasVerification.Contract.AtlasVerificationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AtlasVerification *AtlasVerificationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AtlasVerification.Contract.AtlasVerificationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AtlasVerification *AtlasVerificationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AtlasVerification.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AtlasVerification *AtlasVerificationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AtlasVerification.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AtlasVerification *AtlasVerificationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AtlasVerification.Contract.contract.Transact(opts, method, params...)
}

// GetDAppSignatories is a free data retrieval call binding the contract method 0x6c18e10b.
//
// Solidity: function getDAppSignatories(address dAppControl) view returns(address[])
func (_AtlasVerification *AtlasVerificationCaller) GetDAppSignatories(opts *bind.CallOpts, dAppControl common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getDAppSignatories", dAppControl)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDAppSignatories is a free data retrieval call binding the contract method 0x6c18e10b.
//
// Solidity: function getDAppSignatories(address dAppControl) view returns(address[])
func (_AtlasVerification *AtlasVerificationSession) GetDAppSignatories(dAppControl common.Address) ([]common.Address, error) {
	return _AtlasVerification.Contract.GetDAppSignatories(&_AtlasVerification.CallOpts, dAppControl)
}

// GetDAppSignatories is a free data retrieval call binding the contract method 0x6c18e10b.
//
// Solidity: function getDAppSignatories(address dAppControl) view returns(address[])
func (_AtlasVerification *AtlasVerificationCallerSession) GetDAppSignatories(dAppControl common.Address) ([]common.Address, error) {
	return _AtlasVerification.Contract.GetDAppSignatories(&_AtlasVerification.CallOpts, dAppControl)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32 domainSeparator)
func (_AtlasVerification *AtlasVerificationCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32 domainSeparator)
func (_AtlasVerification *AtlasVerificationSession) GetDomainSeparator() ([32]byte, error) {
	return _AtlasVerification.Contract.GetDomainSeparator(&_AtlasVerification.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32 domainSeparator)
func (_AtlasVerification *AtlasVerificationCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _AtlasVerification.Contract.GetDomainSeparator(&_AtlasVerification.CallOpts)
}
