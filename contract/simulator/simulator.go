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
	_ = abi.ConvertType
)

// DAppOperation is an auto generated low-level Go binding around an user-defined struct.
type DAppOperation struct {
	From          common.Address
	To            common.Address
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
	CallConfig   uint32
	SessionKey   common.Address
	Data         []byte
	Signature    []byte
}

// SimulatorMetaData contains all meta data concerning the Simulator contract.
var SimulatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AllocateValueDelegatecallFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllocateValueFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllocateValueSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlteredControl\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AtlasLockActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceNotReconciled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"BidFindSuccessful\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BidNotPaid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"indexInSolverOps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"BidTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BothPreOpsAndUserReturnDataCannotBeTracked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BothUserAndDAppNoncesCannotBeSequential\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallbackNotCalled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DAppNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DoubleReconcile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnvironmentMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EscrowLockActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionEnvironmentBalanceTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientAtlETHBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"InsufficientAvailableBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalanceForDeduction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"InsufficientBondedBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientEscrow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLocalFunds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"holds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientSolverBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"InsufficientSurchargeBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shortfall\",\"type\":\"uint256\"}],\"name\":\"InsufficientTotalBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"InsufficientUnbondedBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"InsufficientWithdrawableBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAccess\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCodeHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidControl\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDAppControl\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEntry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEntryFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEnvironment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEscrowDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"correctEnvironment\",\"type\":\"address\"}],\"name\":\"InvalidExecutionEnvironment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLockState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignatory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSolver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"solverFrom\",\"type\":\"address\"}],\"name\":\"InvalidSolverFrom\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvertBidValueCannotBeExPostBids\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"LedgerBalancing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"LedgerFinalized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"MissingFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeDelegatecalled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAuctionWinner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoDelegatecall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoUnfilledRequests\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoUnusedNonceInBitmap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnvironmentOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotImplemented\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAtlas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermitDeadlineExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PostOpsDelegatecallFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PostOpsDelegatecallReturnedFalse\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PostOpsFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PostOpsSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PostSolverFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreOpsDelegatecallFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreOpsFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreOpsSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreSolverFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatoryActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SimulationPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverMustReconcile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverOpReverted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"solverOutcomeResult\",\"type\":\"uint256\"}],\"name\":\"SolverSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnbalancedAccounting\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UncoveredResult\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedNonRevert\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unreachable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserNotFulfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserOpFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserOpSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserOpValueExceedsBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserSimulationFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserSimulationSucceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserUnexpectedSuccess\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserWrapperCallFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserWrapperDelegatecallFail\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumValidCallsResult\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"ValidCalls\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumValidCallsResult\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"VerificationSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongDepth\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongPhase\",\"type\":\"error\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"atlas\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation[]\",\"name\":\"solverOps\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bundler\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"callChainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structDAppOperation\",\"name\":\"dAppOp\",\"type\":\"tuple\"}],\"name\":\"metacallSimulation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_atlas\",\"type\":\"address\"}],\"name\":\"setAtlas\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bundler\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"callChainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structDAppOperation\",\"name\":\"dAppOp\",\"type\":\"tuple\"}],\"name\":\"simSolverCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"enumResult\",\"name\":\"simResult\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation[]\",\"name\":\"solverOps\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bundler\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"callChainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structDAppOperation\",\"name\":\"dAppOp\",\"type\":\"tuple\"}],\"name\":\"simSolverCalls\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"enumResult\",\"name\":\"simResult\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"simUserOperation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"enumResult\",\"name\":\"simResult\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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
	parsed, err := SimulatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// Atlas is a free data retrieval call binding the contract method 0x127f1b7d.
//
// Solidity: function atlas() view returns(address)
func (_Simulator *SimulatorCaller) Atlas(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Simulator.contract.Call(opts, &out, "atlas")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Atlas is a free data retrieval call binding the contract method 0x127f1b7d.
//
// Solidity: function atlas() view returns(address)
func (_Simulator *SimulatorSession) Atlas() (common.Address, error) {
	return _Simulator.Contract.Atlas(&_Simulator.CallOpts)
}

// Atlas is a free data retrieval call binding the contract method 0x127f1b7d.
//
// Solidity: function atlas() view returns(address)
func (_Simulator *SimulatorCallerSession) Atlas() (common.Address, error) {
	return _Simulator.Contract.Atlas(&_Simulator.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_Simulator *SimulatorCaller) Deployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Simulator.contract.Call(opts, &out, "deployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_Simulator *SimulatorSession) Deployer() (common.Address, error) {
	return _Simulator.Contract.Deployer(&_Simulator.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_Simulator *SimulatorCallerSession) Deployer() (common.Address, error) {
	return _Simulator.Contract.Deployer(&_Simulator.CallOpts)
}

// MetacallSimulation is a paid mutator transaction binding the contract method 0x776c302a.
//
// Solidity: function metacallSimulation((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns()
func (_Simulator *SimulatorTransactor) MetacallSimulation(opts *bind.TransactOpts, userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.contract.Transact(opts, "metacallSimulation", userOp, solverOps, dAppOp)
}

// MetacallSimulation is a paid mutator transaction binding the contract method 0x776c302a.
//
// Solidity: function metacallSimulation((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns()
func (_Simulator *SimulatorSession) MetacallSimulation(userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.Contract.MetacallSimulation(&_Simulator.TransactOpts, userOp, solverOps, dAppOp)
}

// MetacallSimulation is a paid mutator transaction binding the contract method 0x776c302a.
//
// Solidity: function metacallSimulation((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns()
func (_Simulator *SimulatorTransactorSession) MetacallSimulation(userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.Contract.MetacallSimulation(&_Simulator.TransactOpts, userOp, solverOps, dAppOp)
}

// SetAtlas is a paid mutator transaction binding the contract method 0x133bff36.
//
// Solidity: function setAtlas(address _atlas) returns()
func (_Simulator *SimulatorTransactor) SetAtlas(opts *bind.TransactOpts, _atlas common.Address) (*types.Transaction, error) {
	return _Simulator.contract.Transact(opts, "setAtlas", _atlas)
}

// SetAtlas is a paid mutator transaction binding the contract method 0x133bff36.
//
// Solidity: function setAtlas(address _atlas) returns()
func (_Simulator *SimulatorSession) SetAtlas(_atlas common.Address) (*types.Transaction, error) {
	return _Simulator.Contract.SetAtlas(&_Simulator.TransactOpts, _atlas)
}

// SetAtlas is a paid mutator transaction binding the contract method 0x133bff36.
//
// Solidity: function setAtlas(address _atlas) returns()
func (_Simulator *SimulatorTransactorSession) SetAtlas(_atlas common.Address) (*types.Transaction, error) {
	return _Simulator.Contract.SetAtlas(&_Simulator.TransactOpts, _atlas)
}

// SimSolverCall is a paid mutator transaction binding the contract method 0xc40ff638.
//
// Solidity: function simSolverCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, (address,address,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool success, uint8 simResult, uint256)
func (_Simulator *SimulatorTransactor) SimSolverCall(opts *bind.TransactOpts, userOp UserOperation, solverOp SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.contract.Transact(opts, "simSolverCall", userOp, solverOp, dAppOp)
}

// SimSolverCall is a paid mutator transaction binding the contract method 0xc40ff638.
//
// Solidity: function simSolverCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, (address,address,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool success, uint8 simResult, uint256)
func (_Simulator *SimulatorSession) SimSolverCall(userOp UserOperation, solverOp SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.Contract.SimSolverCall(&_Simulator.TransactOpts, userOp, solverOp, dAppOp)
}

// SimSolverCall is a paid mutator transaction binding the contract method 0xc40ff638.
//
// Solidity: function simSolverCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, (address,address,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool success, uint8 simResult, uint256)
func (_Simulator *SimulatorTransactorSession) SimSolverCall(userOp UserOperation, solverOp SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.Contract.SimSolverCall(&_Simulator.TransactOpts, userOp, solverOp, dAppOp)
}

// SimSolverCalls is a paid mutator transaction binding the contract method 0x041faf77.
//
// Solidity: function simSolverCalls((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool success, uint8 simResult, uint256)
func (_Simulator *SimulatorTransactor) SimSolverCalls(opts *bind.TransactOpts, userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.contract.Transact(opts, "simSolverCalls", userOp, solverOps, dAppOp)
}

// SimSolverCalls is a paid mutator transaction binding the contract method 0x041faf77.
//
// Solidity: function simSolverCalls((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool success, uint8 simResult, uint256)
func (_Simulator *SimulatorSession) SimSolverCalls(userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.Contract.SimSolverCalls(&_Simulator.TransactOpts, userOp, solverOps, dAppOp)
}

// SimSolverCalls is a paid mutator transaction binding the contract method 0x041faf77.
//
// Solidity: function simSolverCalls((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool success, uint8 simResult, uint256)
func (_Simulator *SimulatorTransactorSession) SimSolverCalls(userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Simulator.Contract.SimSolverCalls(&_Simulator.TransactOpts, userOp, solverOps, dAppOp)
}

// SimUserOperation is a paid mutator transaction binding the contract method 0x6731c9c5.
//
// Solidity: function simUserOperation((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) payable returns(bool success, uint8 simResult, uint256)
func (_Simulator *SimulatorTransactor) SimUserOperation(opts *bind.TransactOpts, userOp UserOperation) (*types.Transaction, error) {
	return _Simulator.contract.Transact(opts, "simUserOperation", userOp)
}

// SimUserOperation is a paid mutator transaction binding the contract method 0x6731c9c5.
//
// Solidity: function simUserOperation((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) payable returns(bool success, uint8 simResult, uint256)
func (_Simulator *SimulatorSession) SimUserOperation(userOp UserOperation) (*types.Transaction, error) {
	return _Simulator.Contract.SimUserOperation(&_Simulator.TransactOpts, userOp)
}

// SimUserOperation is a paid mutator transaction binding the contract method 0x6731c9c5.
//
// Solidity: function simUserOperation((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) payable returns(bool success, uint8 simResult, uint256)
func (_Simulator *SimulatorTransactorSession) SimUserOperation(userOp UserOperation) (*types.Transaction, error) {
	return _Simulator.Contract.SimUserOperation(&_Simulator.TransactOpts, userOp)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Simulator *SimulatorTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Simulator.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Simulator *SimulatorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Simulator.Contract.Fallback(&_Simulator.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Simulator *SimulatorTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Simulator.Contract.Fallback(&_Simulator.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Simulator *SimulatorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simulator.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Simulator *SimulatorSession) Receive() (*types.Transaction, error) {
	return _Simulator.Contract.Receive(&_Simulator.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Simulator *SimulatorTransactorSession) Receive() (*types.Transaction, error) {
	return _Simulator.Contract.Receive(&_Simulator.TransactOpts)
}
