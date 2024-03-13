// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simpleRfqSolver

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

// Condition is an auto generated low-level Go binding around an user-defined struct.
type Condition struct {
	Antecedent common.Address
	Context    []byte
}

// StdInvariantFuzzInterface is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzInterface struct {
	Addr      common.Address
	Artifacts []string
}

// StdInvariantFuzzSelector is an auto generated low-level Go binding around an user-defined struct.
type StdInvariantFuzzSelector struct {
	Addr      common.Address
	Selectors [][4]byte
}

// SwapIntent is an auto generated low-level Go binding around an user-defined struct.
type SwapIntent struct {
	TokenUserBuys          common.Address
	AmountUserBuys         *big.Int
	TokenUserSells         common.Address
	AmountUserSells        *big.Int
	AuctionBaseCurrency    common.Address
	SolverMustReimburseGas bool
	Conditions             []Condition
}

// SimpleRfqSolverMetaData contains all meta data concerning the SimpleRfqSolver contract.
var SimpleRfqSolverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"weth\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"atlas\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"IS_TEST\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WETH_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"atlasSolverCall\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"solverOpData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraReturnData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"excludeArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"excludeSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"excludedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"failed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fulfillRFQ\",\"inputs\":[{\"name\":\"swapIntent\",\"type\":\"tuple\",\"internalType\":\"structSwapIntent\",\"components\":[{\"name\":\"tokenUserBuys\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountUserBuys\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenUserSells\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountUserSells\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"auctionBaseCurrency\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solverMustReimburseGas\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"conditions\",\"type\":\"tuple[]\",\"internalType\":\"structCondition[]\",\"components\":[{\"name\":\"antecedent\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"executionEnvironment\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"targetArtifactSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifactSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetArtifacts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedArtifacts_\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetContracts\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedContracts_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetInterfaces\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedInterfaces_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzInterface[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"artifacts\",\"type\":\"string[]\",\"internalType\":\"string[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSelectors\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSelectors_\",\"type\":\"tuple[]\",\"internalType\":\"structStdInvariant.FuzzSelector[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"targetSenders\",\"inputs\":[],\"outputs\":[{\"name\":\"targetedSenders_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"log\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_address\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_array\",\"inputs\":[{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_bytes32\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_int\",\"inputs\":[{\"name\":\"\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_address\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_array\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"address[]\",\"indexed\":false,\"internalType\":\"address[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_bytes32\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_decimal_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"decimals\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_int\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"int256\",\"indexed\":false,\"internalType\":\"int256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_string\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_named_uint\",\"inputs\":[{\"name\":\"key\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"val\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_string\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"log_uint\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"logs\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
}

// SimpleRfqSolverABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleRfqSolverMetaData.ABI instead.
var SimpleRfqSolverABI = SimpleRfqSolverMetaData.ABI

// SimpleRfqSolver is an auto generated Go binding around an Ethereum contract.
type SimpleRfqSolver struct {
	SimpleRfqSolverCaller     // Read-only binding to the contract
	SimpleRfqSolverTransactor // Write-only binding to the contract
	SimpleRfqSolverFilterer   // Log filterer for contract events
}

// SimpleRfqSolverCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleRfqSolverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleRfqSolverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleRfqSolverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleRfqSolverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleRfqSolverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleRfqSolverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleRfqSolverSession struct {
	Contract     *SimpleRfqSolver  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleRfqSolverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleRfqSolverCallerSession struct {
	Contract *SimpleRfqSolverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SimpleRfqSolverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleRfqSolverTransactorSession struct {
	Contract     *SimpleRfqSolverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SimpleRfqSolverRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleRfqSolverRaw struct {
	Contract *SimpleRfqSolver // Generic contract binding to access the raw methods on
}

// SimpleRfqSolverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleRfqSolverCallerRaw struct {
	Contract *SimpleRfqSolverCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleRfqSolverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleRfqSolverTransactorRaw struct {
	Contract *SimpleRfqSolverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleRfqSolver creates a new instance of SimpleRfqSolver, bound to a specific deployed contract.
func NewSimpleRfqSolver(address common.Address, backend bind.ContractBackend) (*SimpleRfqSolver, error) {
	contract, err := bindSimpleRfqSolver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolver{SimpleRfqSolverCaller: SimpleRfqSolverCaller{contract: contract}, SimpleRfqSolverTransactor: SimpleRfqSolverTransactor{contract: contract}, SimpleRfqSolverFilterer: SimpleRfqSolverFilterer{contract: contract}}, nil
}

// NewSimpleRfqSolverCaller creates a new read-only instance of SimpleRfqSolver, bound to a specific deployed contract.
func NewSimpleRfqSolverCaller(address common.Address, caller bind.ContractCaller) (*SimpleRfqSolverCaller, error) {
	contract, err := bindSimpleRfqSolver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverCaller{contract: contract}, nil
}

// NewSimpleRfqSolverTransactor creates a new write-only instance of SimpleRfqSolver, bound to a specific deployed contract.
func NewSimpleRfqSolverTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleRfqSolverTransactor, error) {
	contract, err := bindSimpleRfqSolver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverTransactor{contract: contract}, nil
}

// NewSimpleRfqSolverFilterer creates a new log filterer instance of SimpleRfqSolver, bound to a specific deployed contract.
func NewSimpleRfqSolverFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleRfqSolverFilterer, error) {
	contract, err := bindSimpleRfqSolver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverFilterer{contract: contract}, nil
}

// bindSimpleRfqSolver binds a generic wrapper to an already deployed contract.
func bindSimpleRfqSolver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleRfqSolverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleRfqSolver *SimpleRfqSolverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleRfqSolver.Contract.SimpleRfqSolverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleRfqSolver *SimpleRfqSolverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRfqSolver.Contract.SimpleRfqSolverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleRfqSolver *SimpleRfqSolverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleRfqSolver.Contract.SimpleRfqSolverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleRfqSolver *SimpleRfqSolverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleRfqSolver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleRfqSolver *SimpleRfqSolverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRfqSolver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleRfqSolver *SimpleRfqSolverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleRfqSolver.Contract.contract.Transact(opts, method, params...)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_SimpleRfqSolver *SimpleRfqSolverCaller) ISTEST(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SimpleRfqSolver.contract.Call(opts, &out, "IS_TEST")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_SimpleRfqSolver *SimpleRfqSolverSession) ISTEST() (bool, error) {
	return _SimpleRfqSolver.Contract.ISTEST(&_SimpleRfqSolver.CallOpts)
}

// ISTEST is a free data retrieval call binding the contract method 0xfa7626d4.
//
// Solidity: function IS_TEST() view returns(bool)
func (_SimpleRfqSolver *SimpleRfqSolverCallerSession) ISTEST() (bool, error) {
	return _SimpleRfqSolver.Contract.ISTEST(&_SimpleRfqSolver.CallOpts)
}

// WETHADDRESS is a free data retrieval call binding the contract method 0x040141e5.
//
// Solidity: function WETH_ADDRESS() view returns(address)
func (_SimpleRfqSolver *SimpleRfqSolverCaller) WETHADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleRfqSolver.contract.Call(opts, &out, "WETH_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETHADDRESS is a free data retrieval call binding the contract method 0x040141e5.
//
// Solidity: function WETH_ADDRESS() view returns(address)
func (_SimpleRfqSolver *SimpleRfqSolverSession) WETHADDRESS() (common.Address, error) {
	return _SimpleRfqSolver.Contract.WETHADDRESS(&_SimpleRfqSolver.CallOpts)
}

// WETHADDRESS is a free data retrieval call binding the contract method 0x040141e5.
//
// Solidity: function WETH_ADDRESS() view returns(address)
func (_SimpleRfqSolver *SimpleRfqSolverCallerSession) WETHADDRESS() (common.Address, error) {
	return _SimpleRfqSolver.Contract.WETHADDRESS(&_SimpleRfqSolver.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_SimpleRfqSolver *SimpleRfqSolverCaller) ExcludeArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _SimpleRfqSolver.contract.Call(opts, &out, "excludeArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_SimpleRfqSolver *SimpleRfqSolverSession) ExcludeArtifacts() ([]string, error) {
	return _SimpleRfqSolver.Contract.ExcludeArtifacts(&_SimpleRfqSolver.CallOpts)
}

// ExcludeArtifacts is a free data retrieval call binding the contract method 0xb5508aa9.
//
// Solidity: function excludeArtifacts() view returns(string[] excludedArtifacts_)
func (_SimpleRfqSolver *SimpleRfqSolverCallerSession) ExcludeArtifacts() ([]string, error) {
	return _SimpleRfqSolver.Contract.ExcludeArtifacts(&_SimpleRfqSolver.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_SimpleRfqSolver *SimpleRfqSolverCaller) ExcludeContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SimpleRfqSolver.contract.Call(opts, &out, "excludeContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_SimpleRfqSolver *SimpleRfqSolverSession) ExcludeContracts() ([]common.Address, error) {
	return _SimpleRfqSolver.Contract.ExcludeContracts(&_SimpleRfqSolver.CallOpts)
}

// ExcludeContracts is a free data retrieval call binding the contract method 0xe20c9f71.
//
// Solidity: function excludeContracts() view returns(address[] excludedContracts_)
func (_SimpleRfqSolver *SimpleRfqSolverCallerSession) ExcludeContracts() ([]common.Address, error) {
	return _SimpleRfqSolver.Contract.ExcludeContracts(&_SimpleRfqSolver.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_SimpleRfqSolver *SimpleRfqSolverCaller) ExcludeSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SimpleRfqSolver.contract.Call(opts, &out, "excludeSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_SimpleRfqSolver *SimpleRfqSolverSession) ExcludeSenders() ([]common.Address, error) {
	return _SimpleRfqSolver.Contract.ExcludeSenders(&_SimpleRfqSolver.CallOpts)
}

// ExcludeSenders is a free data retrieval call binding the contract method 0x1ed7831c.
//
// Solidity: function excludeSenders() view returns(address[] excludedSenders_)
func (_SimpleRfqSolver *SimpleRfqSolverCallerSession) ExcludeSenders() ([]common.Address, error) {
	return _SimpleRfqSolver.Contract.ExcludeSenders(&_SimpleRfqSolver.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((address,bytes4[])[] targetedArtifactSelectors_)
func (_SimpleRfqSolver *SimpleRfqSolverCaller) TargetArtifactSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _SimpleRfqSolver.contract.Call(opts, &out, "targetArtifactSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((address,bytes4[])[] targetedArtifactSelectors_)
func (_SimpleRfqSolver *SimpleRfqSolverSession) TargetArtifactSelectors() ([]StdInvariantFuzzSelector, error) {
	return _SimpleRfqSolver.Contract.TargetArtifactSelectors(&_SimpleRfqSolver.CallOpts)
}

// TargetArtifactSelectors is a free data retrieval call binding the contract method 0x66d9a9a0.
//
// Solidity: function targetArtifactSelectors() view returns((address,bytes4[])[] targetedArtifactSelectors_)
func (_SimpleRfqSolver *SimpleRfqSolverCallerSession) TargetArtifactSelectors() ([]StdInvariantFuzzSelector, error) {
	return _SimpleRfqSolver.Contract.TargetArtifactSelectors(&_SimpleRfqSolver.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_SimpleRfqSolver *SimpleRfqSolverCaller) TargetArtifacts(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _SimpleRfqSolver.contract.Call(opts, &out, "targetArtifacts")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_SimpleRfqSolver *SimpleRfqSolverSession) TargetArtifacts() ([]string, error) {
	return _SimpleRfqSolver.Contract.TargetArtifacts(&_SimpleRfqSolver.CallOpts)
}

// TargetArtifacts is a free data retrieval call binding the contract method 0x85226c81.
//
// Solidity: function targetArtifacts() view returns(string[] targetedArtifacts_)
func (_SimpleRfqSolver *SimpleRfqSolverCallerSession) TargetArtifacts() ([]string, error) {
	return _SimpleRfqSolver.Contract.TargetArtifacts(&_SimpleRfqSolver.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_SimpleRfqSolver *SimpleRfqSolverCaller) TargetContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SimpleRfqSolver.contract.Call(opts, &out, "targetContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_SimpleRfqSolver *SimpleRfqSolverSession) TargetContracts() ([]common.Address, error) {
	return _SimpleRfqSolver.Contract.TargetContracts(&_SimpleRfqSolver.CallOpts)
}

// TargetContracts is a free data retrieval call binding the contract method 0x3f7286f4.
//
// Solidity: function targetContracts() view returns(address[] targetedContracts_)
func (_SimpleRfqSolver *SimpleRfqSolverCallerSession) TargetContracts() ([]common.Address, error) {
	return _SimpleRfqSolver.Contract.TargetContracts(&_SimpleRfqSolver.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_SimpleRfqSolver *SimpleRfqSolverCaller) TargetInterfaces(opts *bind.CallOpts) ([]StdInvariantFuzzInterface, error) {
	var out []interface{}
	err := _SimpleRfqSolver.contract.Call(opts, &out, "targetInterfaces")

	if err != nil {
		return *new([]StdInvariantFuzzInterface), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzInterface)).(*[]StdInvariantFuzzInterface)

	return out0, err

}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_SimpleRfqSolver *SimpleRfqSolverSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _SimpleRfqSolver.Contract.TargetInterfaces(&_SimpleRfqSolver.CallOpts)
}

// TargetInterfaces is a free data retrieval call binding the contract method 0x2ade3880.
//
// Solidity: function targetInterfaces() view returns((address,string[])[] targetedInterfaces_)
func (_SimpleRfqSolver *SimpleRfqSolverCallerSession) TargetInterfaces() ([]StdInvariantFuzzInterface, error) {
	return _SimpleRfqSolver.Contract.TargetInterfaces(&_SimpleRfqSolver.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_SimpleRfqSolver *SimpleRfqSolverCaller) TargetSelectors(opts *bind.CallOpts) ([]StdInvariantFuzzSelector, error) {
	var out []interface{}
	err := _SimpleRfqSolver.contract.Call(opts, &out, "targetSelectors")

	if err != nil {
		return *new([]StdInvariantFuzzSelector), err
	}

	out0 := *abi.ConvertType(out[0], new([]StdInvariantFuzzSelector)).(*[]StdInvariantFuzzSelector)

	return out0, err

}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_SimpleRfqSolver *SimpleRfqSolverSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _SimpleRfqSolver.Contract.TargetSelectors(&_SimpleRfqSolver.CallOpts)
}

// TargetSelectors is a free data retrieval call binding the contract method 0x916a17c6.
//
// Solidity: function targetSelectors() view returns((address,bytes4[])[] targetedSelectors_)
func (_SimpleRfqSolver *SimpleRfqSolverCallerSession) TargetSelectors() ([]StdInvariantFuzzSelector, error) {
	return _SimpleRfqSolver.Contract.TargetSelectors(&_SimpleRfqSolver.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_SimpleRfqSolver *SimpleRfqSolverCaller) TargetSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SimpleRfqSolver.contract.Call(opts, &out, "targetSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_SimpleRfqSolver *SimpleRfqSolverSession) TargetSenders() ([]common.Address, error) {
	return _SimpleRfqSolver.Contract.TargetSenders(&_SimpleRfqSolver.CallOpts)
}

// TargetSenders is a free data retrieval call binding the contract method 0x3e5e3c23.
//
// Solidity: function targetSenders() view returns(address[] targetedSenders_)
func (_SimpleRfqSolver *SimpleRfqSolverCallerSession) TargetSenders() ([]common.Address, error) {
	return _SimpleRfqSolver.Contract.TargetSenders(&_SimpleRfqSolver.CallOpts)
}

// AtlasSolverCall is a paid mutator transaction binding the contract method 0x24634204.
//
// Solidity: function atlasSolverCall(address sender, address bidToken, uint256 bidAmount, bytes solverOpData, bytes extraReturnData) payable returns(bool success, bytes data)
func (_SimpleRfqSolver *SimpleRfqSolverTransactor) AtlasSolverCall(opts *bind.TransactOpts, sender common.Address, bidToken common.Address, bidAmount *big.Int, solverOpData []byte, extraReturnData []byte) (*types.Transaction, error) {
	return _SimpleRfqSolver.contract.Transact(opts, "atlasSolverCall", sender, bidToken, bidAmount, solverOpData, extraReturnData)
}

// AtlasSolverCall is a paid mutator transaction binding the contract method 0x24634204.
//
// Solidity: function atlasSolverCall(address sender, address bidToken, uint256 bidAmount, bytes solverOpData, bytes extraReturnData) payable returns(bool success, bytes data)
func (_SimpleRfqSolver *SimpleRfqSolverSession) AtlasSolverCall(sender common.Address, bidToken common.Address, bidAmount *big.Int, solverOpData []byte, extraReturnData []byte) (*types.Transaction, error) {
	return _SimpleRfqSolver.Contract.AtlasSolverCall(&_SimpleRfqSolver.TransactOpts, sender, bidToken, bidAmount, solverOpData, extraReturnData)
}

// AtlasSolverCall is a paid mutator transaction binding the contract method 0x24634204.
//
// Solidity: function atlasSolverCall(address sender, address bidToken, uint256 bidAmount, bytes solverOpData, bytes extraReturnData) payable returns(bool success, bytes data)
func (_SimpleRfqSolver *SimpleRfqSolverTransactorSession) AtlasSolverCall(sender common.Address, bidToken common.Address, bidAmount *big.Int, solverOpData []byte, extraReturnData []byte) (*types.Transaction, error) {
	return _SimpleRfqSolver.Contract.AtlasSolverCall(&_SimpleRfqSolver.TransactOpts, sender, bidToken, bidAmount, solverOpData, extraReturnData)
}

// Failed is a paid mutator transaction binding the contract method 0xba414fa6.
//
// Solidity: function failed() returns(bool)
func (_SimpleRfqSolver *SimpleRfqSolverTransactor) Failed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRfqSolver.contract.Transact(opts, "failed")
}

// Failed is a paid mutator transaction binding the contract method 0xba414fa6.
//
// Solidity: function failed() returns(bool)
func (_SimpleRfqSolver *SimpleRfqSolverSession) Failed() (*types.Transaction, error) {
	return _SimpleRfqSolver.Contract.Failed(&_SimpleRfqSolver.TransactOpts)
}

// Failed is a paid mutator transaction binding the contract method 0xba414fa6.
//
// Solidity: function failed() returns(bool)
func (_SimpleRfqSolver *SimpleRfqSolverTransactorSession) Failed() (*types.Transaction, error) {
	return _SimpleRfqSolver.Contract.Failed(&_SimpleRfqSolver.TransactOpts)
}

// FulfillRFQ is a paid mutator transaction binding the contract method 0x491274c5.
//
// Solidity: function fulfillRFQ((address,uint256,address,uint256,address,bool,(address,bytes)[]) swapIntent, address executionEnvironment) returns()
func (_SimpleRfqSolver *SimpleRfqSolverTransactor) FulfillRFQ(opts *bind.TransactOpts, swapIntent SwapIntent, executionEnvironment common.Address) (*types.Transaction, error) {
	return _SimpleRfqSolver.contract.Transact(opts, "fulfillRFQ", swapIntent, executionEnvironment)
}

// FulfillRFQ is a paid mutator transaction binding the contract method 0x491274c5.
//
// Solidity: function fulfillRFQ((address,uint256,address,uint256,address,bool,(address,bytes)[]) swapIntent, address executionEnvironment) returns()
func (_SimpleRfqSolver *SimpleRfqSolverSession) FulfillRFQ(swapIntent SwapIntent, executionEnvironment common.Address) (*types.Transaction, error) {
	return _SimpleRfqSolver.Contract.FulfillRFQ(&_SimpleRfqSolver.TransactOpts, swapIntent, executionEnvironment)
}

// FulfillRFQ is a paid mutator transaction binding the contract method 0x491274c5.
//
// Solidity: function fulfillRFQ((address,uint256,address,uint256,address,bool,(address,bytes)[]) swapIntent, address executionEnvironment) returns()
func (_SimpleRfqSolver *SimpleRfqSolverTransactorSession) FulfillRFQ(swapIntent SwapIntent, executionEnvironment common.Address) (*types.Transaction, error) {
	return _SimpleRfqSolver.Contract.FulfillRFQ(&_SimpleRfqSolver.TransactOpts, swapIntent, executionEnvironment)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SimpleRfqSolver *SimpleRfqSolverTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SimpleRfqSolver.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SimpleRfqSolver *SimpleRfqSolverSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SimpleRfqSolver.Contract.Fallback(&_SimpleRfqSolver.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SimpleRfqSolver *SimpleRfqSolverTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SimpleRfqSolver.Contract.Fallback(&_SimpleRfqSolver.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleRfqSolver *SimpleRfqSolverTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRfqSolver.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleRfqSolver *SimpleRfqSolverSession) Receive() (*types.Transaction, error) {
	return _SimpleRfqSolver.Contract.Receive(&_SimpleRfqSolver.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleRfqSolver *SimpleRfqSolverTransactorSession) Receive() (*types.Transaction, error) {
	return _SimpleRfqSolver.Contract.Receive(&_SimpleRfqSolver.TransactOpts)
}

// SimpleRfqSolverLogIterator is returned from FilterLog and is used to iterate over the raw logs and unpacked data for Log events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogIterator struct {
	Event *SimpleRfqSolverLog // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLog)
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
		it.Event = new(SimpleRfqSolverLog)
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
func (it *SimpleRfqSolverLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLog represents a Log event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLog is a free log retrieval operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLog(opts *bind.FilterOpts) (*SimpleRfqSolverLogIterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogIterator{contract: _SimpleRfqSolver.contract, event: "log", logs: logs, sub: sub}, nil
}

// WatchLog is a free log subscription operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLog(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLog) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLog)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log", log); err != nil {
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

// ParseLog is a log parse operation binding the contract event 0x41304facd9323d75b11bcdd609cb38effffdb05710f7caf0e9b16c6d9d709f50.
//
// Solidity: event log(string arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLog(log types.Log) (*SimpleRfqSolverLog, error) {
	event := new(SimpleRfqSolverLog)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogAddressIterator is returned from FilterLogAddress and is used to iterate over the raw logs and unpacked data for LogAddress events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogAddressIterator struct {
	Event *SimpleRfqSolverLogAddress // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogAddress)
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
		it.Event = new(SimpleRfqSolverLogAddress)
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
func (it *SimpleRfqSolverLogAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogAddress represents a LogAddress event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogAddress struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddress is a free log retrieval operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogAddress(opts *bind.FilterOpts) (*SimpleRfqSolverLogAddressIterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogAddressIterator{contract: _SimpleRfqSolver.contract, event: "log_address", logs: logs, sub: sub}, nil
}

// WatchLogAddress is a free log subscription operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogAddress(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogAddress) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogAddress)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_address", log); err != nil {
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

// ParseLogAddress is a log parse operation binding the contract event 0x7ae74c527414ae135fd97047b12921a5ec3911b804197855d67e25c7b75ee6f3.
//
// Solidity: event log_address(address arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogAddress(log types.Log) (*SimpleRfqSolverLogAddress, error) {
	event := new(SimpleRfqSolverLogAddress)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogArrayIterator is returned from FilterLogArray and is used to iterate over the raw logs and unpacked data for LogArray events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogArrayIterator struct {
	Event *SimpleRfqSolverLogArray // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogArray)
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
		it.Event = new(SimpleRfqSolverLogArray)
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
func (it *SimpleRfqSolverLogArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogArray represents a LogArray event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogArray struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray is a free log retrieval operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogArray(opts *bind.FilterOpts) (*SimpleRfqSolverLogArrayIterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogArrayIterator{contract: _SimpleRfqSolver.contract, event: "log_array", logs: logs, sub: sub}, nil
}

// WatchLogArray is a free log subscription operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogArray(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogArray) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogArray)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_array", log); err != nil {
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

// ParseLogArray is a log parse operation binding the contract event 0xfb102865d50addddf69da9b5aa1bced66c80cf869a5c8d0471a467e18ce9cab1.
//
// Solidity: event log_array(uint256[] val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogArray(log types.Log) (*SimpleRfqSolverLogArray, error) {
	event := new(SimpleRfqSolverLogArray)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogArray0Iterator is returned from FilterLogArray0 and is used to iterate over the raw logs and unpacked data for LogArray0 events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogArray0Iterator struct {
	Event *SimpleRfqSolverLogArray0 // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogArray0)
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
		it.Event = new(SimpleRfqSolverLogArray0)
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
func (it *SimpleRfqSolverLogArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogArray0 represents a LogArray0 event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogArray0 struct {
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray0 is a free log retrieval operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogArray0(opts *bind.FilterOpts) (*SimpleRfqSolverLogArray0Iterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogArray0Iterator{contract: _SimpleRfqSolver.contract, event: "log_array0", logs: logs, sub: sub}, nil
}

// WatchLogArray0 is a free log subscription operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogArray0(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogArray0) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogArray0)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_array0", log); err != nil {
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

// ParseLogArray0 is a log parse operation binding the contract event 0x890a82679b470f2bd82816ed9b161f97d8b967f37fa3647c21d5bf39749e2dd5.
//
// Solidity: event log_array(int256[] val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogArray0(log types.Log) (*SimpleRfqSolverLogArray0, error) {
	event := new(SimpleRfqSolverLogArray0)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogArray1Iterator is returned from FilterLogArray1 and is used to iterate over the raw logs and unpacked data for LogArray1 events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogArray1Iterator struct {
	Event *SimpleRfqSolverLogArray1 // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogArray1)
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
		it.Event = new(SimpleRfqSolverLogArray1)
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
func (it *SimpleRfqSolverLogArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogArray1 represents a LogArray1 event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogArray1 struct {
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogArray1 is a free log retrieval operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogArray1(opts *bind.FilterOpts) (*SimpleRfqSolverLogArray1Iterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogArray1Iterator{contract: _SimpleRfqSolver.contract, event: "log_array1", logs: logs, sub: sub}, nil
}

// WatchLogArray1 is a free log subscription operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogArray1(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogArray1) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogArray1)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_array1", log); err != nil {
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

// ParseLogArray1 is a log parse operation binding the contract event 0x40e1840f5769073d61bd01372d9b75baa9842d5629a0c99ff103be1178a8e9e2.
//
// Solidity: event log_array(address[] val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogArray1(log types.Log) (*SimpleRfqSolverLogArray1, error) {
	event := new(SimpleRfqSolverLogArray1)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogBytesIterator is returned from FilterLogBytes and is used to iterate over the raw logs and unpacked data for LogBytes events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogBytesIterator struct {
	Event *SimpleRfqSolverLogBytes // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogBytes)
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
		it.Event = new(SimpleRfqSolverLogBytes)
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
func (it *SimpleRfqSolverLogBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogBytes represents a LogBytes event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogBytes struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes is a free log retrieval operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogBytes(opts *bind.FilterOpts) (*SimpleRfqSolverLogBytesIterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogBytesIterator{contract: _SimpleRfqSolver.contract, event: "log_bytes", logs: logs, sub: sub}, nil
}

// WatchLogBytes is a free log subscription operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogBytes(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogBytes) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogBytes)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_bytes", log); err != nil {
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

// ParseLogBytes is a log parse operation binding the contract event 0x23b62ad0584d24a75f0bf3560391ef5659ec6db1269c56e11aa241d637f19b20.
//
// Solidity: event log_bytes(bytes arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogBytes(log types.Log) (*SimpleRfqSolverLogBytes, error) {
	event := new(SimpleRfqSolverLogBytes)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogBytes32Iterator is returned from FilterLogBytes32 and is used to iterate over the raw logs and unpacked data for LogBytes32 events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogBytes32Iterator struct {
	Event *SimpleRfqSolverLogBytes32 // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogBytes32)
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
		it.Event = new(SimpleRfqSolverLogBytes32)
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
func (it *SimpleRfqSolverLogBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogBytes32 represents a LogBytes32 event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogBytes32 struct {
	Arg0 [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogBytes32 is a free log retrieval operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogBytes32(opts *bind.FilterOpts) (*SimpleRfqSolverLogBytes32Iterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogBytes32Iterator{contract: _SimpleRfqSolver.contract, event: "log_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogBytes32 is a free log subscription operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogBytes32(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogBytes32) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogBytes32)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_bytes32", log); err != nil {
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

// ParseLogBytes32 is a log parse operation binding the contract event 0xe81699b85113eea1c73e10588b2b035e55893369632173afd43feb192fac64e3.
//
// Solidity: event log_bytes32(bytes32 arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogBytes32(log types.Log) (*SimpleRfqSolverLogBytes32, error) {
	event := new(SimpleRfqSolverLogBytes32)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogIntIterator is returned from FilterLogInt and is used to iterate over the raw logs and unpacked data for LogInt events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogIntIterator struct {
	Event *SimpleRfqSolverLogInt // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogInt)
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
		it.Event = new(SimpleRfqSolverLogInt)
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
func (it *SimpleRfqSolverLogIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogInt represents a LogInt event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogInt struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogInt is a free log retrieval operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogInt(opts *bind.FilterOpts) (*SimpleRfqSolverLogIntIterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogIntIterator{contract: _SimpleRfqSolver.contract, event: "log_int", logs: logs, sub: sub}, nil
}

// WatchLogInt is a free log subscription operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogInt(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogInt) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogInt)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_int", log); err != nil {
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

// ParseLogInt is a log parse operation binding the contract event 0x0eb5d52624c8d28ada9fc55a8c502ed5aa3fbe2fb6e91b71b5f376882b1d2fb8.
//
// Solidity: event log_int(int256 arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogInt(log types.Log) (*SimpleRfqSolverLogInt, error) {
	event := new(SimpleRfqSolverLogInt)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogNamedAddressIterator is returned from FilterLogNamedAddress and is used to iterate over the raw logs and unpacked data for LogNamedAddress events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedAddressIterator struct {
	Event *SimpleRfqSolverLogNamedAddress // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogNamedAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogNamedAddress)
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
		it.Event = new(SimpleRfqSolverLogNamedAddress)
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
func (it *SimpleRfqSolverLogNamedAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogNamedAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogNamedAddress represents a LogNamedAddress event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedAddress struct {
	Key string
	Val common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedAddress is a free log retrieval operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogNamedAddress(opts *bind.FilterOpts) (*SimpleRfqSolverLogNamedAddressIterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogNamedAddressIterator{contract: _SimpleRfqSolver.contract, event: "log_named_address", logs: logs, sub: sub}, nil
}

// WatchLogNamedAddress is a free log subscription operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogNamedAddress(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogNamedAddress) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_named_address")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogNamedAddress)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_address", log); err != nil {
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

// ParseLogNamedAddress is a log parse operation binding the contract event 0x9c4e8541ca8f0dc1c413f9108f66d82d3cecb1bddbce437a61caa3175c4cc96f.
//
// Solidity: event log_named_address(string key, address val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogNamedAddress(log types.Log) (*SimpleRfqSolverLogNamedAddress, error) {
	event := new(SimpleRfqSolverLogNamedAddress)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_address", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogNamedArrayIterator is returned from FilterLogNamedArray and is used to iterate over the raw logs and unpacked data for LogNamedArray events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedArrayIterator struct {
	Event *SimpleRfqSolverLogNamedArray // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogNamedArrayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogNamedArray)
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
		it.Event = new(SimpleRfqSolverLogNamedArray)
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
func (it *SimpleRfqSolverLogNamedArrayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogNamedArrayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogNamedArray represents a LogNamedArray event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedArray struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray is a free log retrieval operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogNamedArray(opts *bind.FilterOpts) (*SimpleRfqSolverLogNamedArrayIterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogNamedArrayIterator{contract: _SimpleRfqSolver.contract, event: "log_named_array", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray is a free log subscription operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogNamedArray(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogNamedArray) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_named_array")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogNamedArray)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_array", log); err != nil {
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

// ParseLogNamedArray is a log parse operation binding the contract event 0x00aaa39c9ffb5f567a4534380c737075702e1f7f14107fc95328e3b56c0325fb.
//
// Solidity: event log_named_array(string key, uint256[] val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogNamedArray(log types.Log) (*SimpleRfqSolverLogNamedArray, error) {
	event := new(SimpleRfqSolverLogNamedArray)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_array", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogNamedArray0Iterator is returned from FilterLogNamedArray0 and is used to iterate over the raw logs and unpacked data for LogNamedArray0 events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedArray0Iterator struct {
	Event *SimpleRfqSolverLogNamedArray0 // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogNamedArray0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogNamedArray0)
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
		it.Event = new(SimpleRfqSolverLogNamedArray0)
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
func (it *SimpleRfqSolverLogNamedArray0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogNamedArray0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogNamedArray0 represents a LogNamedArray0 event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedArray0 struct {
	Key string
	Val []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray0 is a free log retrieval operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogNamedArray0(opts *bind.FilterOpts) (*SimpleRfqSolverLogNamedArray0Iterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogNamedArray0Iterator{contract: _SimpleRfqSolver.contract, event: "log_named_array0", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray0 is a free log subscription operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogNamedArray0(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogNamedArray0) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_named_array0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogNamedArray0)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_array0", log); err != nil {
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

// ParseLogNamedArray0 is a log parse operation binding the contract event 0xa73eda09662f46dde729be4611385ff34fe6c44fbbc6f7e17b042b59a3445b57.
//
// Solidity: event log_named_array(string key, int256[] val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogNamedArray0(log types.Log) (*SimpleRfqSolverLogNamedArray0, error) {
	event := new(SimpleRfqSolverLogNamedArray0)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_array0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogNamedArray1Iterator is returned from FilterLogNamedArray1 and is used to iterate over the raw logs and unpacked data for LogNamedArray1 events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedArray1Iterator struct {
	Event *SimpleRfqSolverLogNamedArray1 // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogNamedArray1Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogNamedArray1)
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
		it.Event = new(SimpleRfqSolverLogNamedArray1)
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
func (it *SimpleRfqSolverLogNamedArray1Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogNamedArray1Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogNamedArray1 represents a LogNamedArray1 event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedArray1 struct {
	Key string
	Val []common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedArray1 is a free log retrieval operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogNamedArray1(opts *bind.FilterOpts) (*SimpleRfqSolverLogNamedArray1Iterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogNamedArray1Iterator{contract: _SimpleRfqSolver.contract, event: "log_named_array1", logs: logs, sub: sub}, nil
}

// WatchLogNamedArray1 is a free log subscription operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogNamedArray1(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogNamedArray1) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_named_array1")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogNamedArray1)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_array1", log); err != nil {
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

// ParseLogNamedArray1 is a log parse operation binding the contract event 0x3bcfb2ae2e8d132dd1fce7cf278a9a19756a9fceabe470df3bdabb4bc577d1bd.
//
// Solidity: event log_named_array(string key, address[] val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogNamedArray1(log types.Log) (*SimpleRfqSolverLogNamedArray1, error) {
	event := new(SimpleRfqSolverLogNamedArray1)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_array1", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogNamedBytesIterator is returned from FilterLogNamedBytes and is used to iterate over the raw logs and unpacked data for LogNamedBytes events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedBytesIterator struct {
	Event *SimpleRfqSolverLogNamedBytes // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogNamedBytesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogNamedBytes)
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
		it.Event = new(SimpleRfqSolverLogNamedBytes)
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
func (it *SimpleRfqSolverLogNamedBytesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogNamedBytesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogNamedBytes represents a LogNamedBytes event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedBytes struct {
	Key string
	Val []byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes is a free log retrieval operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogNamedBytes(opts *bind.FilterOpts) (*SimpleRfqSolverLogNamedBytesIterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogNamedBytesIterator{contract: _SimpleRfqSolver.contract, event: "log_named_bytes", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes is a free log subscription operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogNamedBytes(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogNamedBytes) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_named_bytes")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogNamedBytes)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
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

// ParseLogNamedBytes is a log parse operation binding the contract event 0xd26e16cad4548705e4c9e2d94f98ee91c289085ee425594fd5635fa2964ccf18.
//
// Solidity: event log_named_bytes(string key, bytes val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogNamedBytes(log types.Log) (*SimpleRfqSolverLogNamedBytes, error) {
	event := new(SimpleRfqSolverLogNamedBytes)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_bytes", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogNamedBytes32Iterator is returned from FilterLogNamedBytes32 and is used to iterate over the raw logs and unpacked data for LogNamedBytes32 events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedBytes32Iterator struct {
	Event *SimpleRfqSolverLogNamedBytes32 // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogNamedBytes32Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogNamedBytes32)
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
		it.Event = new(SimpleRfqSolverLogNamedBytes32)
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
func (it *SimpleRfqSolverLogNamedBytes32Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogNamedBytes32Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogNamedBytes32 represents a LogNamedBytes32 event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedBytes32 struct {
	Key string
	Val [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedBytes32 is a free log retrieval operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogNamedBytes32(opts *bind.FilterOpts) (*SimpleRfqSolverLogNamedBytes32Iterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogNamedBytes32Iterator{contract: _SimpleRfqSolver.contract, event: "log_named_bytes32", logs: logs, sub: sub}, nil
}

// WatchLogNamedBytes32 is a free log subscription operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogNamedBytes32(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogNamedBytes32) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_named_bytes32")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogNamedBytes32)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
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

// ParseLogNamedBytes32 is a log parse operation binding the contract event 0xafb795c9c61e4fe7468c386f925d7a5429ecad9c0495ddb8d38d690614d32f99.
//
// Solidity: event log_named_bytes32(string key, bytes32 val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogNamedBytes32(log types.Log) (*SimpleRfqSolverLogNamedBytes32, error) {
	event := new(SimpleRfqSolverLogNamedBytes32)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_bytes32", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogNamedDecimalIntIterator is returned from FilterLogNamedDecimalInt and is used to iterate over the raw logs and unpacked data for LogNamedDecimalInt events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedDecimalIntIterator struct {
	Event *SimpleRfqSolverLogNamedDecimalInt // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogNamedDecimalIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogNamedDecimalInt)
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
		it.Event = new(SimpleRfqSolverLogNamedDecimalInt)
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
func (it *SimpleRfqSolverLogNamedDecimalIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogNamedDecimalIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogNamedDecimalInt represents a LogNamedDecimalInt event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedDecimalInt struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalInt is a free log retrieval operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogNamedDecimalInt(opts *bind.FilterOpts) (*SimpleRfqSolverLogNamedDecimalIntIterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogNamedDecimalIntIterator{contract: _SimpleRfqSolver.contract, event: "log_named_decimal_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalInt is a free log subscription operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogNamedDecimalInt(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogNamedDecimalInt) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_named_decimal_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogNamedDecimalInt)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
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

// ParseLogNamedDecimalInt is a log parse operation binding the contract event 0x5da6ce9d51151ba10c09a559ef24d520b9dac5c5b8810ae8434e4d0d86411a95.
//
// Solidity: event log_named_decimal_int(string key, int256 val, uint256 decimals)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogNamedDecimalInt(log types.Log) (*SimpleRfqSolverLogNamedDecimalInt, error) {
	event := new(SimpleRfqSolverLogNamedDecimalInt)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_decimal_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogNamedDecimalUintIterator is returned from FilterLogNamedDecimalUint and is used to iterate over the raw logs and unpacked data for LogNamedDecimalUint events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedDecimalUintIterator struct {
	Event *SimpleRfqSolverLogNamedDecimalUint // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogNamedDecimalUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogNamedDecimalUint)
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
		it.Event = new(SimpleRfqSolverLogNamedDecimalUint)
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
func (it *SimpleRfqSolverLogNamedDecimalUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogNamedDecimalUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogNamedDecimalUint represents a LogNamedDecimalUint event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedDecimalUint struct {
	Key      string
	Val      *big.Int
	Decimals *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNamedDecimalUint is a free log retrieval operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogNamedDecimalUint(opts *bind.FilterOpts) (*SimpleRfqSolverLogNamedDecimalUintIterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogNamedDecimalUintIterator{contract: _SimpleRfqSolver.contract, event: "log_named_decimal_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedDecimalUint is a free log subscription operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogNamedDecimalUint(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogNamedDecimalUint) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_named_decimal_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogNamedDecimalUint)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
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

// ParseLogNamedDecimalUint is a log parse operation binding the contract event 0xeb8ba43ced7537421946bd43e828b8b2b8428927aa8f801c13d934bf11aca57b.
//
// Solidity: event log_named_decimal_uint(string key, uint256 val, uint256 decimals)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogNamedDecimalUint(log types.Log) (*SimpleRfqSolverLogNamedDecimalUint, error) {
	event := new(SimpleRfqSolverLogNamedDecimalUint)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_decimal_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogNamedIntIterator is returned from FilterLogNamedInt and is used to iterate over the raw logs and unpacked data for LogNamedInt events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedIntIterator struct {
	Event *SimpleRfqSolverLogNamedInt // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogNamedIntIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogNamedInt)
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
		it.Event = new(SimpleRfqSolverLogNamedInt)
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
func (it *SimpleRfqSolverLogNamedIntIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogNamedIntIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogNamedInt represents a LogNamedInt event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedInt struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedInt is a free log retrieval operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogNamedInt(opts *bind.FilterOpts) (*SimpleRfqSolverLogNamedIntIterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogNamedIntIterator{contract: _SimpleRfqSolver.contract, event: "log_named_int", logs: logs, sub: sub}, nil
}

// WatchLogNamedInt is a free log subscription operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogNamedInt(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogNamedInt) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_named_int")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogNamedInt)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_int", log); err != nil {
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

// ParseLogNamedInt is a log parse operation binding the contract event 0x2fe632779174374378442a8e978bccfbdcc1d6b2b0d81f7e8eb776ab2286f168.
//
// Solidity: event log_named_int(string key, int256 val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogNamedInt(log types.Log) (*SimpleRfqSolverLogNamedInt, error) {
	event := new(SimpleRfqSolverLogNamedInt)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_int", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogNamedStringIterator is returned from FilterLogNamedString and is used to iterate over the raw logs and unpacked data for LogNamedString events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedStringIterator struct {
	Event *SimpleRfqSolverLogNamedString // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogNamedStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogNamedString)
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
		it.Event = new(SimpleRfqSolverLogNamedString)
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
func (it *SimpleRfqSolverLogNamedStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogNamedStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogNamedString represents a LogNamedString event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedString struct {
	Key string
	Val string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedString is a free log retrieval operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogNamedString(opts *bind.FilterOpts) (*SimpleRfqSolverLogNamedStringIterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogNamedStringIterator{contract: _SimpleRfqSolver.contract, event: "log_named_string", logs: logs, sub: sub}, nil
}

// WatchLogNamedString is a free log subscription operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogNamedString(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogNamedString) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_named_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogNamedString)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_string", log); err != nil {
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

// ParseLogNamedString is a log parse operation binding the contract event 0x280f4446b28a1372417dda658d30b95b2992b12ac9c7f378535f29a97acf3583.
//
// Solidity: event log_named_string(string key, string val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogNamedString(log types.Log) (*SimpleRfqSolverLogNamedString, error) {
	event := new(SimpleRfqSolverLogNamedString)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogNamedUintIterator is returned from FilterLogNamedUint and is used to iterate over the raw logs and unpacked data for LogNamedUint events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedUintIterator struct {
	Event *SimpleRfqSolverLogNamedUint // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogNamedUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogNamedUint)
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
		it.Event = new(SimpleRfqSolverLogNamedUint)
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
func (it *SimpleRfqSolverLogNamedUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogNamedUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogNamedUint represents a LogNamedUint event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogNamedUint struct {
	Key string
	Val *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogNamedUint is a free log retrieval operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogNamedUint(opts *bind.FilterOpts) (*SimpleRfqSolverLogNamedUintIterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogNamedUintIterator{contract: _SimpleRfqSolver.contract, event: "log_named_uint", logs: logs, sub: sub}, nil
}

// WatchLogNamedUint is a free log subscription operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogNamedUint(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogNamedUint) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_named_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogNamedUint)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_uint", log); err != nil {
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

// ParseLogNamedUint is a log parse operation binding the contract event 0xb2de2fbe801a0df6c0cbddfd448ba3c41d48a040ca35c56c8196ef0fcae721a8.
//
// Solidity: event log_named_uint(string key, uint256 val)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogNamedUint(log types.Log) (*SimpleRfqSolverLogNamedUint, error) {
	event := new(SimpleRfqSolverLogNamedUint)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_named_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogStringIterator struct {
	Event *SimpleRfqSolverLogString // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogString)
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
		it.Event = new(SimpleRfqSolverLogString)
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
func (it *SimpleRfqSolverLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogString represents a LogString event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogString struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogString(opts *bind.FilterOpts) (*SimpleRfqSolverLogStringIterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogStringIterator{contract: _SimpleRfqSolver.contract, event: "log_string", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogString) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_string")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogString)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_string", log); err != nil {
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

// ParseLogString is a log parse operation binding the contract event 0x0b2e13ff20ac7b474198655583edf70dedd2c1dc980e329c4fbb2fc0748b796b.
//
// Solidity: event log_string(string arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogString(log types.Log) (*SimpleRfqSolverLogString, error) {
	event := new(SimpleRfqSolverLogString)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_string", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogUintIterator is returned from FilterLogUint and is used to iterate over the raw logs and unpacked data for LogUint events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogUintIterator struct {
	Event *SimpleRfqSolverLogUint // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogUintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogUint)
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
		it.Event = new(SimpleRfqSolverLogUint)
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
func (it *SimpleRfqSolverLogUintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogUintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogUint represents a LogUint event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogUint struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogUint is a free log retrieval operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogUint(opts *bind.FilterOpts) (*SimpleRfqSolverLogUintIterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogUintIterator{contract: _SimpleRfqSolver.contract, event: "log_uint", logs: logs, sub: sub}, nil
}

// WatchLogUint is a free log subscription operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogUint(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogUint) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "log_uint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogUint)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_uint", log); err != nil {
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

// ParseLogUint is a log parse operation binding the contract event 0x2cab9790510fd8bdfbd2115288db33fec66691d476efc5427cfd4c0969301755.
//
// Solidity: event log_uint(uint256 arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogUint(log types.Log) (*SimpleRfqSolverLogUint, error) {
	event := new(SimpleRfqSolverLogUint)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "log_uint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SimpleRfqSolverLogsIterator is returned from FilterLogs and is used to iterate over the raw logs and unpacked data for Logs events raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogsIterator struct {
	Event *SimpleRfqSolverLogs // Event containing the contract specifics and raw log

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
func (it *SimpleRfqSolverLogsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRfqSolverLogs)
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
		it.Event = new(SimpleRfqSolverLogs)
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
func (it *SimpleRfqSolverLogsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRfqSolverLogsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRfqSolverLogs represents a Logs event raised by the SimpleRfqSolver contract.
type SimpleRfqSolverLogs struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogs is a free log retrieval operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) FilterLogs(opts *bind.FilterOpts) (*SimpleRfqSolverLogsIterator, error) {

	logs, sub, err := _SimpleRfqSolver.contract.FilterLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return &SimpleRfqSolverLogsIterator{contract: _SimpleRfqSolver.contract, event: "logs", logs: logs, sub: sub}, nil
}

// WatchLogs is a free log subscription operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) WatchLogs(opts *bind.WatchOpts, sink chan<- *SimpleRfqSolverLogs) (event.Subscription, error) {

	logs, sub, err := _SimpleRfqSolver.contract.WatchLogs(opts, "logs")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRfqSolverLogs)
				if err := _SimpleRfqSolver.contract.UnpackLog(event, "logs", log); err != nil {
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

// ParseLogs is a log parse operation binding the contract event 0xe7950ede0394b9f2ce4a5a1bf5a7e1852411f7e6661b4308c913c4bfd11027e4.
//
// Solidity: event logs(bytes arg0)
func (_SimpleRfqSolver *SimpleRfqSolverFilterer) ParseLogs(log types.Log) (*SimpleRfqSolverLogs, error) {
	event := new(SimpleRfqSolverLogs)
	if err := _SimpleRfqSolver.contract.UnpackLog(event, "logs", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
