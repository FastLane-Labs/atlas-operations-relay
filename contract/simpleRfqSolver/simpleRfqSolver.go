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

// SwapIntent is an auto generated low-level Go binding around an user-defined struct.
type SwapIntent struct {
	TokenUserBuys       common.Address
	AmountUserBuys      *big.Int
	TokenUserSells      common.Address
	AmountUserSells     *big.Int
	AuctionBaseCurrency common.Address
	Conditions          []Condition
}

// SimpleRfqSolverMetaData contains all meta data concerning the SimpleRfqSolver contract.
var SimpleRfqSolverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"weth\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"atlas\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"WETH_ADDRESS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"atlasSolverCall\",\"inputs\":[{\"name\":\"solverOpFrom\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"executionEnvironment\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"solverOpData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"fulfillRFQ\",\"inputs\":[{\"name\":\"swapIntent\",\"type\":\"tuple\",\"internalType\":\"structSwapIntent\",\"components\":[{\"name\":\"tokenUserBuys\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountUserBuys\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenUserSells\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountUserSells\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"auctionBaseCurrency\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"conditions\",\"type\":\"tuple[]\",\"internalType\":\"structCondition[]\",\"components\":[{\"name\":\"antecedent\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"context\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]},{\"name\":\"executionEnvironment\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
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

// AtlasSolverCall is a paid mutator transaction binding the contract method 0x024181a6.
//
// Solidity: function atlasSolverCall(address solverOpFrom, address executionEnvironment, address bidToken, uint256 bidAmount, bytes solverOpData, bytes ) payable returns()
func (_SimpleRfqSolver *SimpleRfqSolverTransactor) AtlasSolverCall(opts *bind.TransactOpts, solverOpFrom common.Address, executionEnvironment common.Address, bidToken common.Address, bidAmount *big.Int, solverOpData []byte, arg5 []byte) (*types.Transaction, error) {
	return _SimpleRfqSolver.contract.Transact(opts, "atlasSolverCall", solverOpFrom, executionEnvironment, bidToken, bidAmount, solverOpData, arg5)
}

// AtlasSolverCall is a paid mutator transaction binding the contract method 0x024181a6.
//
// Solidity: function atlasSolverCall(address solverOpFrom, address executionEnvironment, address bidToken, uint256 bidAmount, bytes solverOpData, bytes ) payable returns()
func (_SimpleRfqSolver *SimpleRfqSolverSession) AtlasSolverCall(solverOpFrom common.Address, executionEnvironment common.Address, bidToken common.Address, bidAmount *big.Int, solverOpData []byte, arg5 []byte) (*types.Transaction, error) {
	return _SimpleRfqSolver.Contract.AtlasSolverCall(&_SimpleRfqSolver.TransactOpts, solverOpFrom, executionEnvironment, bidToken, bidAmount, solverOpData, arg5)
}

// AtlasSolverCall is a paid mutator transaction binding the contract method 0x024181a6.
//
// Solidity: function atlasSolverCall(address solverOpFrom, address executionEnvironment, address bidToken, uint256 bidAmount, bytes solverOpData, bytes ) payable returns()
func (_SimpleRfqSolver *SimpleRfqSolverTransactorSession) AtlasSolverCall(solverOpFrom common.Address, executionEnvironment common.Address, bidToken common.Address, bidAmount *big.Int, solverOpData []byte, arg5 []byte) (*types.Transaction, error) {
	return _SimpleRfqSolver.Contract.AtlasSolverCall(&_SimpleRfqSolver.TransactOpts, solverOpFrom, executionEnvironment, bidToken, bidAmount, solverOpData, arg5)
}

// FulfillRFQ is a paid mutator transaction binding the contract method 0x2ad99656.
//
// Solidity: function fulfillRFQ((address,uint256,address,uint256,address,(address,bytes)[]) swapIntent, address executionEnvironment) returns()
func (_SimpleRfqSolver *SimpleRfqSolverTransactor) FulfillRFQ(opts *bind.TransactOpts, swapIntent SwapIntent, executionEnvironment common.Address) (*types.Transaction, error) {
	return _SimpleRfqSolver.contract.Transact(opts, "fulfillRFQ", swapIntent, executionEnvironment)
}

// FulfillRFQ is a paid mutator transaction binding the contract method 0x2ad99656.
//
// Solidity: function fulfillRFQ((address,uint256,address,uint256,address,(address,bytes)[]) swapIntent, address executionEnvironment) returns()
func (_SimpleRfqSolver *SimpleRfqSolverSession) FulfillRFQ(swapIntent SwapIntent, executionEnvironment common.Address) (*types.Transaction, error) {
	return _SimpleRfqSolver.Contract.FulfillRFQ(&_SimpleRfqSolver.TransactOpts, swapIntent, executionEnvironment)
}

// FulfillRFQ is a paid mutator transaction binding the contract method 0x2ad99656.
//
// Solidity: function fulfillRFQ((address,uint256,address,uint256,address,(address,bytes)[]) swapIntent, address executionEnvironment) returns()
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
