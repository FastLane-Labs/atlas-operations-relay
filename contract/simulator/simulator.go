// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simulator

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

// DAppOperation is an auto generated low-level Go binding around an user-defined struct.
type DAppOperation struct {
	From          common.Address
	To            common.Address
	Value         *big.Int
	Gas           *big.Int
	MaxFeePerGas  *big.Int
	Nonce         *big.Int
	Deadline      *big.Int
	Control       common.Address
	Bundler       common.Address
	UserOpHash    [32]byte
	CallChainHash [32]byte
	Signature     []byte
}

// SolverOperation is an auto generated low-level Go binding around an user-defined struct.
type SolverOperation struct {
	From         common.Address
	To           common.Address
	Value        *big.Int
	Gas          *big.Int
	MaxFeePerGas *big.Int
	Deadline     *big.Int
	Solver       common.Address
	Control      common.Address
	UserOpHash   [32]byte
	BidToken     common.Address
	BidAmount    *big.Int
	Data         []byte
	Signature    []byte
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

// SimulatorMetaData contains all meta data concerning the Simulator contract.
var SimulatorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"simSolverCall\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structUserOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"solverOp\",\"type\":\"tuple\",\"internalType\":\"structSolverOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"dAppOp\",\"type\":\"tuple\",\"internalType\":\"structDAppOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bundler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callChainHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"simSolverCalls\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structUserOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"solverOps\",\"type\":\"tuple[]\",\"internalType\":\"structSolverOperation[]\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"dAppOp\",\"type\":\"tuple\",\"internalType\":\"structDAppOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bundler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callChainHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"simUserOperation\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structUserOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"success\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"}]",
}

// SimulatorABI is the input ABI used to generate the binding from.
// Deprecated: Use SimulatorMetaData.ABI instead.
var SimulatorABI = SimulatorMetaData.ABI

// Simulator is an auto generated Go binding around an Ethereum contract.
type Simulator struct {
	SimulatorCaller     // Read-only binding to the contract
	SimulatorTransactor // Write-only binding to the contract
	SimulatorFilterer   // Log filterer for contract events
}

// SimulatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimulatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimulatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimulatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimulatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimulatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimulatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimulatorSession struct {
	Contract     *Simulator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimulatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimulatorCallerSession struct {
	Contract *SimulatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SimulatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimulatorTransactorSession struct {
	Contract     *SimulatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SimulatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimulatorRaw struct {
	Contract *Simulator // Generic contract binding to access the raw methods on
}

// SimulatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimulatorCallerRaw struct {
	Contract *SimulatorCaller // Generic read-only contract binding to access the raw methods on
}

// SimulatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimulatorTransactorRaw struct {
	Contract *SimulatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimulator creates a new instance of Simulator, bound to a specific deployed contract.
func NewSimulator(address common.Address, backend bind.ContractBackend) (*Simulator, error) {
	contract, err := bindSimulator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simulator{SimulatorCaller: SimulatorCaller{contract: contract}, SimulatorTransactor: SimulatorTransactor{contract: contract}, SimulatorFilterer: SimulatorFilterer{contract: contract}}, nil
}

// NewSimulatorCaller creates a new read-only instance of Simulator, bound to a specific deployed contract.
func NewSimulatorCaller(address common.Address, caller bind.ContractCaller) (*SimulatorCaller, error) {
	contract, err := bindSimulator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimulatorCaller{contract: contract}, nil
}

// NewSimulatorTransactor creates a new write-only instance of Simulator, bound to a specific deployed contract.
func NewSimulatorTransactor(address common.Address, transactor bind.ContractTransactor) (*SimulatorTransactor, error) {
	contract, err := bindSimulator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimulatorTransactor{contract: contract}, nil
}

// NewSimulatorFilterer creates a new log filterer instance of Simulator, bound to a specific deployed contract.
func NewSimulatorFilterer(address common.Address, filterer bind.ContractFilterer) (*SimulatorFilterer, error) {
	contract, err := bindSimulator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimulatorFilterer{contract: contract}, nil
}

// bindSimulator binds a generic wrapper to an already deployed contract.
func bindSimulator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SimulatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simulator *SimulatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simulator.Contract.SimulatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simulator *SimulatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simulator.Contract.SimulatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simulator *SimulatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simulator.Contract.SimulatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simulator *SimulatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simulator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simulator *SimulatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simulator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simulator *SimulatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simulator.Contract.contract.Transact(opts, method, params...)
}

// SimSolverCall is a paid mutator transaction binding the contract method 0xf5f8bc61.
//
// Solidity: function simSolverCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool success)
func (_Simulator *SimulatorTransactor) SimSolverCall(opts *bind.TransactOpts, userOp UserOperation, solverOp SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.contract.Transact(opts, "simSolverCall", userOp, solverOp, dAppOp)
}

// SimSolverCall is a paid mutator transaction binding the contract method 0xf5f8bc61.
//
// Solidity: function simSolverCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool success)
func (_Simulator *SimulatorSession) SimSolverCall(userOp UserOperation, solverOp SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.Contract.SimSolverCall(&_Simulator.TransactOpts, userOp, solverOp, dAppOp)
}

// SimSolverCall is a paid mutator transaction binding the contract method 0xf5f8bc61.
//
// Solidity: function simSolverCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool success)
func (_Simulator *SimulatorTransactorSession) SimSolverCall(userOp UserOperation, solverOp SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.Contract.SimSolverCall(&_Simulator.TransactOpts, userOp, solverOp, dAppOp)
}

// SimSolverCalls is a paid mutator transaction binding the contract method 0x7c785b92.
//
// Solidity: function simSolverCalls((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool success)
func (_Simulator *SimulatorTransactor) SimSolverCalls(opts *bind.TransactOpts, userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.contract.Transact(opts, "simSolverCalls", userOp, solverOps, dAppOp)
}

// SimSolverCalls is a paid mutator transaction binding the contract method 0x7c785b92.
//
// Solidity: function simSolverCalls((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool success)
func (_Simulator *SimulatorSession) SimSolverCalls(userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.Contract.SimSolverCalls(&_Simulator.TransactOpts, userOp, solverOps, dAppOp)
}

// SimSolverCalls is a paid mutator transaction binding the contract method 0x7c785b92.
//
// Solidity: function simSolverCalls((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool success)
func (_Simulator *SimulatorTransactorSession) SimSolverCalls(userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.Contract.SimSolverCalls(&_Simulator.TransactOpts, userOp, solverOps, dAppOp)
}

// SimUserOperation is a paid mutator transaction binding the contract method 0xfb667144.
//
// Solidity: function simUserOperation((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp) payable returns(bool success)
func (_Simulator *SimulatorTransactor) SimUserOperation(opts *bind.TransactOpts, userOp UserOperation) (*types.Transaction, error) {
	return _Simulator.contract.Transact(opts, "simUserOperation", userOp)
}

// SimUserOperation is a paid mutator transaction binding the contract method 0xfb667144.
//
// Solidity: function simUserOperation((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp) payable returns(bool success)
func (_Simulator *SimulatorSession) SimUserOperation(userOp UserOperation) (*types.Transaction, error) {
	return _Simulator.Contract.SimUserOperation(&_Simulator.TransactOpts, userOp)
}

// SimUserOperation is a paid mutator transaction binding the contract method 0xfb667144.
//
// Solidity: function simUserOperation((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp) payable returns(bool success)
func (_Simulator *SimulatorTransactorSession) SimUserOperation(userOp UserOperation) (*types.Transaction, error) {
	return _Simulator.Contract.SimUserOperation(&_Simulator.TransactOpts, userOp)
}
