// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storage

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

// StorageMetaData contains all meta data concerning the Storage contract.
var StorageMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_escrowDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_verification\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_simulator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_surchargeRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"ESCROW_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SIMULATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SURCHARGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERIFICATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accessData\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"bonded\",\"type\":\"uint112\",\"internalType\":\"uint112\"},{\"name\":\"lastAccessedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"auctionWins\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"auctionFails\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"totalGasUsed\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bondedTotalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claims\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposits\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingSurchargeRecipient\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"solver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"solverLockData\",\"inputs\":[],\"outputs\":[{\"name\":\"currentSolver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"calledBack\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"surcharge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"surchargeRecipient\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Bond\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DAppDisabled\",\"inputs\":[{\"name\":\"controller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"governance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DAppGovernanceChanged\",\"inputs\":[{\"name\":\"controller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutionEnvironmentCreated\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executionEnvironment\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovernanceTransferStarted\",\"inputs\":[{\"name\":\"previousGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GovernanceTransferred\",\"inputs\":[{\"name\":\"previousGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetacallResult\",\"inputs\":[{\"name\":\"bundler\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"winningSolver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"ethPaidToBundler\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"netGasSurcharge\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewDAppSignatory\",\"inputs\":[{\"name\":\"controller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"governance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"signatory\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Redeem\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemovedDAppSignatory\",\"inputs\":[{\"name\":\"controller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"governance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"signatory\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SolverTxResult\",\"inputs\":[{\"name\":\"solverTo\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"solverFrom\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"result\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SurchargeRecipientTransferStarted\",\"inputs\":[{\"name\":\"currentRecipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRecipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SurchargeRecipientTransferred\",\"inputs\":[{\"name\":\"newRecipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SurchargeWithdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unbond\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"earliestAvailable\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlteredControl\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AtlasLockActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BalanceNotReconciled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BidFindSuccessful\",\"inputs\":[{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"BothUserAndDAppNoncesCannotBeSequenced\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallbackNotCalled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DAppNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DoubleReconcile\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnvironmentMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EscrowLockActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientAtlETHBalance\",\"inputs\":[{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientAvailableBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientBalanceForDeduction\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientBondedBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientFunds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientLocalFunds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientSolverBalance\",\"inputs\":[{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"holds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientSurchargeBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientTotalBalance\",\"inputs\":[{\"name\":\"shortfall\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientUnbondedBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientWithdrawableBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"IntentUnfulfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAccess\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidControl\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDAppControl\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEntry\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEntryFunction\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEnvironment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExecutionEnvironment\",\"inputs\":[{\"name\":\"correctEnvironment\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidLockState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSolverFrom\",\"inputs\":[{\"name\":\"solverFrom\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"LedgerBalancing\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"LedgerFinalized\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"MissingFunds\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"MustBeDelegatecalled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoAuctionWinner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoDelegatecall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoUnfilledRequests\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoUnusedNonceInBitmap\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotImplemented\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAccount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAtlas\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyGovernance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PermitDeadlineExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PostOpsFail\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PostOpsSimFail\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PostSolverFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PreOpsFail\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PreOpsSimFail\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PreSolverFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatoryActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SimulationPassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SolverBidUnpaid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SolverMustReconcile\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SolverOperationReverted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SolverSimFail\",\"inputs\":[{\"name\":\"solverOutcomeResult\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnbalancedAccounting\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UncoveredResult\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnexpectedNonRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unreachable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserNotFulfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserOpFail\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserOpSimFail\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserSimulationFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserSimulationSucceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserUnexpectedSuccess\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidCalls\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumValidCallsResult\"}]},{\"type\":\"error\",\"name\":\"ValueTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VerificationSimFail\",\"inputs\":[{\"name\":\"validCallsResult\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"WrongDepth\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongPhase\",\"inputs\":[]}]",
}

// StorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageMetaData.ABI instead.
var StorageABI = StorageMetaData.ABI

// Storage is an auto generated Go binding around an Ethereum contract.
type Storage struct {
	StorageCaller     // Read-only binding to the contract
	StorageTransactor // Write-only binding to the contract
	StorageFilterer   // Log filterer for contract events
}

// StorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSession struct {
	Contract     *Storage          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageCallerSession struct {
	Contract *StorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTransactorSession struct {
	Contract     *StorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageRaw struct {
	Contract *Storage // Generic contract binding to access the raw methods on
}

// StorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageCallerRaw struct {
	Contract *StorageCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTransactorRaw struct {
	Contract *StorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorage creates a new instance of Storage, bound to a specific deployed contract.
func NewStorage(address common.Address, backend bind.ContractBackend) (*Storage, error) {
	contract, err := bindStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Storage{StorageCaller: StorageCaller{contract: contract}, StorageTransactor: StorageTransactor{contract: contract}, StorageFilterer: StorageFilterer{contract: contract}}, nil
}

// NewStorageCaller creates a new read-only instance of Storage, bound to a specific deployed contract.
func NewStorageCaller(address common.Address, caller bind.ContractCaller) (*StorageCaller, error) {
	contract, err := bindStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageCaller{contract: contract}, nil
}

// NewStorageTransactor creates a new write-only instance of Storage, bound to a specific deployed contract.
func NewStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTransactor, error) {
	contract, err := bindStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTransactor{contract: contract}, nil
}

// NewStorageFilterer creates a new log filterer instance of Storage, bound to a specific deployed contract.
func NewStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageFilterer, error) {
	contract, err := bindStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageFilterer{contract: contract}, nil
}

// bindStorage binds a generic wrapper to an already deployed contract.
func bindStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.StorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.StorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Storage *StorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Storage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Storage *StorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Storage *StorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Storage.Contract.contract.Transact(opts, method, params...)
}

// ESCROWDURATION is a free data retrieval call binding the contract method 0xa6efccf9.
//
// Solidity: function ESCROW_DURATION() view returns(uint256)
func (_Storage *StorageCaller) ESCROWDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "ESCROW_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ESCROWDURATION is a free data retrieval call binding the contract method 0xa6efccf9.
//
// Solidity: function ESCROW_DURATION() view returns(uint256)
func (_Storage *StorageSession) ESCROWDURATION() (*big.Int, error) {
	return _Storage.Contract.ESCROWDURATION(&_Storage.CallOpts)
}

// ESCROWDURATION is a free data retrieval call binding the contract method 0xa6efccf9.
//
// Solidity: function ESCROW_DURATION() view returns(uint256)
func (_Storage *StorageCallerSession) ESCROWDURATION() (*big.Int, error) {
	return _Storage.Contract.ESCROWDURATION(&_Storage.CallOpts)
}

// SIMULATOR is a free data retrieval call binding the contract method 0x79b79765.
//
// Solidity: function SIMULATOR() view returns(address)
func (_Storage *StorageCaller) SIMULATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "SIMULATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SIMULATOR is a free data retrieval call binding the contract method 0x79b79765.
//
// Solidity: function SIMULATOR() view returns(address)
func (_Storage *StorageSession) SIMULATOR() (common.Address, error) {
	return _Storage.Contract.SIMULATOR(&_Storage.CallOpts)
}

// SIMULATOR is a free data retrieval call binding the contract method 0x79b79765.
//
// Solidity: function SIMULATOR() view returns(address)
func (_Storage *StorageCallerSession) SIMULATOR() (common.Address, error) {
	return _Storage.Contract.SIMULATOR(&_Storage.CallOpts)
}

// SURCHARGE is a free data retrieval call binding the contract method 0x37642394.
//
// Solidity: function SURCHARGE() view returns(uint256)
func (_Storage *StorageCaller) SURCHARGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "SURCHARGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SURCHARGE is a free data retrieval call binding the contract method 0x37642394.
//
// Solidity: function SURCHARGE() view returns(uint256)
func (_Storage *StorageSession) SURCHARGE() (*big.Int, error) {
	return _Storage.Contract.SURCHARGE(&_Storage.CallOpts)
}

// SURCHARGE is a free data retrieval call binding the contract method 0x37642394.
//
// Solidity: function SURCHARGE() view returns(uint256)
func (_Storage *StorageCallerSession) SURCHARGE() (*big.Int, error) {
	return _Storage.Contract.SURCHARGE(&_Storage.CallOpts)
}

// VERIFICATION is a free data retrieval call binding the contract method 0x791ae748.
//
// Solidity: function VERIFICATION() view returns(address)
func (_Storage *StorageCaller) VERIFICATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "VERIFICATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VERIFICATION is a free data retrieval call binding the contract method 0x791ae748.
//
// Solidity: function VERIFICATION() view returns(address)
func (_Storage *StorageSession) VERIFICATION() (common.Address, error) {
	return _Storage.Contract.VERIFICATION(&_Storage.CallOpts)
}

// VERIFICATION is a free data retrieval call binding the contract method 0x791ae748.
//
// Solidity: function VERIFICATION() view returns(address)
func (_Storage *StorageCallerSession) VERIFICATION() (common.Address, error) {
	return _Storage.Contract.VERIFICATION(&_Storage.CallOpts)
}

// AccessData is a free data retrieval call binding the contract method 0x5e8edccc.
//
// Solidity: function accessData(address ) view returns(uint112 bonded, uint32 lastAccessedBlock, uint24 auctionWins, uint24 auctionFails, uint64 totalGasUsed)
func (_Storage *StorageCaller) AccessData(opts *bind.CallOpts, arg0 common.Address) (struct {
	Bonded            *big.Int
	LastAccessedBlock uint32
	AuctionWins       *big.Int
	AuctionFails      *big.Int
	TotalGasUsed      uint64
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "accessData", arg0)

	outstruct := new(struct {
		Bonded            *big.Int
		LastAccessedBlock uint32
		AuctionWins       *big.Int
		AuctionFails      *big.Int
		TotalGasUsed      uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bonded = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastAccessedBlock = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.AuctionWins = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.AuctionFails = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalGasUsed = *abi.ConvertType(out[4], new(uint64)).(*uint64)

	return *outstruct, err

}

// AccessData is a free data retrieval call binding the contract method 0x5e8edccc.
//
// Solidity: function accessData(address ) view returns(uint112 bonded, uint32 lastAccessedBlock, uint24 auctionWins, uint24 auctionFails, uint64 totalGasUsed)
func (_Storage *StorageSession) AccessData(arg0 common.Address) (struct {
	Bonded            *big.Int
	LastAccessedBlock uint32
	AuctionWins       *big.Int
	AuctionFails      *big.Int
	TotalGasUsed      uint64
}, error) {
	return _Storage.Contract.AccessData(&_Storage.CallOpts, arg0)
}

// AccessData is a free data retrieval call binding the contract method 0x5e8edccc.
//
// Solidity: function accessData(address ) view returns(uint112 bonded, uint32 lastAccessedBlock, uint24 auctionWins, uint24 auctionFails, uint64 totalGasUsed)
func (_Storage *StorageCallerSession) AccessData(arg0 common.Address) (struct {
	Bonded            *big.Int
	LastAccessedBlock uint32
	AuctionWins       *big.Int
	AuctionFails      *big.Int
	TotalGasUsed      uint64
}, error) {
	return _Storage.Contract.AccessData(&_Storage.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Storage *StorageCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Storage *StorageSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Storage.Contract.Allowance(&_Storage.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Storage *StorageCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Storage.Contract.Allowance(&_Storage.CallOpts, arg0, arg1)
}

// BondedTotalSupply is a free data retrieval call binding the contract method 0x890c2854.
//
// Solidity: function bondedTotalSupply() view returns(uint256)
func (_Storage *StorageCaller) BondedTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "bondedTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BondedTotalSupply is a free data retrieval call binding the contract method 0x890c2854.
//
// Solidity: function bondedTotalSupply() view returns(uint256)
func (_Storage *StorageSession) BondedTotalSupply() (*big.Int, error) {
	return _Storage.Contract.BondedTotalSupply(&_Storage.CallOpts)
}

// BondedTotalSupply is a free data retrieval call binding the contract method 0x890c2854.
//
// Solidity: function bondedTotalSupply() view returns(uint256)
func (_Storage *StorageCallerSession) BondedTotalSupply() (*big.Int, error) {
	return _Storage.Contract.BondedTotalSupply(&_Storage.CallOpts)
}

// Claims is a free data retrieval call binding the contract method 0xdcc59b6f.
//
// Solidity: function claims() view returns(uint256)
func (_Storage *StorageCaller) Claims(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "claims")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claims is a free data retrieval call binding the contract method 0xdcc59b6f.
//
// Solidity: function claims() view returns(uint256)
func (_Storage *StorageSession) Claims() (*big.Int, error) {
	return _Storage.Contract.Claims(&_Storage.CallOpts)
}

// Claims is a free data retrieval call binding the contract method 0xdcc59b6f.
//
// Solidity: function claims() view returns(uint256)
func (_Storage *StorageCallerSession) Claims() (*big.Int, error) {
	return _Storage.Contract.Claims(&_Storage.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Storage *StorageCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Storage *StorageSession) Decimals() (uint8, error) {
	return _Storage.Contract.Decimals(&_Storage.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Storage *StorageCallerSession) Decimals() (uint8, error) {
	return _Storage.Contract.Decimals(&_Storage.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0x323a5e0b.
//
// Solidity: function deposits() view returns(uint256)
func (_Storage *StorageCaller) Deposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "deposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0x323a5e0b.
//
// Solidity: function deposits() view returns(uint256)
func (_Storage *StorageSession) Deposits() (*big.Int, error) {
	return _Storage.Contract.Deposits(&_Storage.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0x323a5e0b.
//
// Solidity: function deposits() view returns(uint256)
func (_Storage *StorageCallerSession) Deposits() (*big.Int, error) {
	return _Storage.Contract.Deposits(&_Storage.CallOpts)
}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() view returns(address)
func (_Storage *StorageCaller) Lock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "lock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() view returns(address)
func (_Storage *StorageSession) Lock() (common.Address, error) {
	return _Storage.Contract.Lock(&_Storage.CallOpts)
}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() view returns(address)
func (_Storage *StorageCallerSession) Lock() (common.Address, error) {
	return _Storage.Contract.Lock(&_Storage.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Storage *StorageCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Storage *StorageSession) Name() (string, error) {
	return _Storage.Contract.Name(&_Storage.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Storage *StorageCallerSession) Name() (string, error) {
	return _Storage.Contract.Name(&_Storage.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Storage *StorageCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Storage *StorageSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.Nonces(&_Storage.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Storage *StorageCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Storage.Contract.Nonces(&_Storage.CallOpts, arg0)
}

// PendingSurchargeRecipient is a free data retrieval call binding the contract method 0x7c3c3160.
//
// Solidity: function pendingSurchargeRecipient() view returns(address)
func (_Storage *StorageCaller) PendingSurchargeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "pendingSurchargeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingSurchargeRecipient is a free data retrieval call binding the contract method 0x7c3c3160.
//
// Solidity: function pendingSurchargeRecipient() view returns(address)
func (_Storage *StorageSession) PendingSurchargeRecipient() (common.Address, error) {
	return _Storage.Contract.PendingSurchargeRecipient(&_Storage.CallOpts)
}

// PendingSurchargeRecipient is a free data retrieval call binding the contract method 0x7c3c3160.
//
// Solidity: function pendingSurchargeRecipient() view returns(address)
func (_Storage *StorageCallerSession) PendingSurchargeRecipient() (common.Address, error) {
	return _Storage.Contract.PendingSurchargeRecipient(&_Storage.CallOpts)
}

// Solver is a free data retrieval call binding the contract method 0x49a7a26d.
//
// Solidity: function solver() view returns(address)
func (_Storage *StorageCaller) Solver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "solver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Solver is a free data retrieval call binding the contract method 0x49a7a26d.
//
// Solidity: function solver() view returns(address)
func (_Storage *StorageSession) Solver() (common.Address, error) {
	return _Storage.Contract.Solver(&_Storage.CallOpts)
}

// Solver is a free data retrieval call binding the contract method 0x49a7a26d.
//
// Solidity: function solver() view returns(address)
func (_Storage *StorageCallerSession) Solver() (common.Address, error) {
	return _Storage.Contract.Solver(&_Storage.CallOpts)
}

// SolverLockData is a free data retrieval call binding the contract method 0xaa7d2dc8.
//
// Solidity: function solverLockData() view returns(address currentSolver, bool calledBack, bool fulfilled)
func (_Storage *StorageCaller) SolverLockData(opts *bind.CallOpts) (struct {
	CurrentSolver common.Address
	CalledBack    bool
	Fulfilled     bool
}, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "solverLockData")

	outstruct := new(struct {
		CurrentSolver common.Address
		CalledBack    bool
		Fulfilled     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentSolver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CalledBack = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Fulfilled = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// SolverLockData is a free data retrieval call binding the contract method 0xaa7d2dc8.
//
// Solidity: function solverLockData() view returns(address currentSolver, bool calledBack, bool fulfilled)
func (_Storage *StorageSession) SolverLockData() (struct {
	CurrentSolver common.Address
	CalledBack    bool
	Fulfilled     bool
}, error) {
	return _Storage.Contract.SolverLockData(&_Storage.CallOpts)
}

// SolverLockData is a free data retrieval call binding the contract method 0xaa7d2dc8.
//
// Solidity: function solverLockData() view returns(address currentSolver, bool calledBack, bool fulfilled)
func (_Storage *StorageCallerSession) SolverLockData() (struct {
	CurrentSolver common.Address
	CalledBack    bool
	Fulfilled     bool
}, error) {
	return _Storage.Contract.SolverLockData(&_Storage.CallOpts)
}

// Surcharge is a free data retrieval call binding the contract method 0x2ab26a56.
//
// Solidity: function surcharge() view returns(uint256)
func (_Storage *StorageCaller) Surcharge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "surcharge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Surcharge is a free data retrieval call binding the contract method 0x2ab26a56.
//
// Solidity: function surcharge() view returns(uint256)
func (_Storage *StorageSession) Surcharge() (*big.Int, error) {
	return _Storage.Contract.Surcharge(&_Storage.CallOpts)
}

// Surcharge is a free data retrieval call binding the contract method 0x2ab26a56.
//
// Solidity: function surcharge() view returns(uint256)
func (_Storage *StorageCallerSession) Surcharge() (*big.Int, error) {
	return _Storage.Contract.Surcharge(&_Storage.CallOpts)
}

// SurchargeRecipient is a free data retrieval call binding the contract method 0xfc61c541.
//
// Solidity: function surchargeRecipient() view returns(address)
func (_Storage *StorageCaller) SurchargeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "surchargeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SurchargeRecipient is a free data retrieval call binding the contract method 0xfc61c541.
//
// Solidity: function surchargeRecipient() view returns(address)
func (_Storage *StorageSession) SurchargeRecipient() (common.Address, error) {
	return _Storage.Contract.SurchargeRecipient(&_Storage.CallOpts)
}

// SurchargeRecipient is a free data retrieval call binding the contract method 0xfc61c541.
//
// Solidity: function surchargeRecipient() view returns(address)
func (_Storage *StorageCallerSession) SurchargeRecipient() (common.Address, error) {
	return _Storage.Contract.SurchargeRecipient(&_Storage.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Storage *StorageCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Storage *StorageSession) Symbol() (string, error) {
	return _Storage.Contract.Symbol(&_Storage.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Storage *StorageCallerSession) Symbol() (string, error) {
	return _Storage.Contract.Symbol(&_Storage.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Storage *StorageCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Storage *StorageSession) TotalSupply() (*big.Int, error) {
	return _Storage.Contract.TotalSupply(&_Storage.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Storage *StorageCallerSession) TotalSupply() (*big.Int, error) {
	return _Storage.Contract.TotalSupply(&_Storage.CallOpts)
}

// Withdrawals is a free data retrieval call binding the contract method 0xade0e93e.
//
// Solidity: function withdrawals() view returns(uint256)
func (_Storage *StorageCaller) Withdrawals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Storage.contract.Call(opts, &out, "withdrawals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Withdrawals is a free data retrieval call binding the contract method 0xade0e93e.
//
// Solidity: function withdrawals() view returns(uint256)
func (_Storage *StorageSession) Withdrawals() (*big.Int, error) {
	return _Storage.Contract.Withdrawals(&_Storage.CallOpts)
}

// Withdrawals is a free data retrieval call binding the contract method 0xade0e93e.
//
// Solidity: function withdrawals() view returns(uint256)
func (_Storage *StorageCallerSession) Withdrawals() (*big.Int, error) {
	return _Storage.Contract.Withdrawals(&_Storage.CallOpts)
}

// StorageApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Storage contract.
type StorageApprovalIterator struct {
	Event *StorageApproval // Event containing the contract specifics and raw log

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
func (it *StorageApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageApproval)
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
		it.Event = new(StorageApproval)
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
func (it *StorageApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageApproval represents a Approval event raised by the Storage contract.
type StorageApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Storage *StorageFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StorageApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StorageApprovalIterator{contract: _Storage.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Storage *StorageFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StorageApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageApproval)
				if err := _Storage.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Storage *StorageFilterer) ParseApproval(log types.Log) (*StorageApproval, error) {
	event := new(StorageApproval)
	if err := _Storage.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageBondIterator is returned from FilterBond and is used to iterate over the raw logs and unpacked data for Bond events raised by the Storage contract.
type StorageBondIterator struct {
	Event *StorageBond // Event containing the contract specifics and raw log

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
func (it *StorageBondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageBond)
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
		it.Event = new(StorageBond)
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
func (it *StorageBondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageBondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageBond represents a Bond event raised by the Storage contract.
type StorageBond struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBond is a free log retrieval operation binding the contract event 0x6b1d99469ed62a423d7e402bfa68a467261ca2229127c55045ee41e5d9a0f21d.
//
// Solidity: event Bond(address indexed owner, uint256 amount)
func (_Storage *StorageFilterer) FilterBond(opts *bind.FilterOpts, owner []common.Address) (*StorageBondIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Bond", ownerRule)
	if err != nil {
		return nil, err
	}
	return &StorageBondIterator{contract: _Storage.contract, event: "Bond", logs: logs, sub: sub}, nil
}

// WatchBond is a free log subscription operation binding the contract event 0x6b1d99469ed62a423d7e402bfa68a467261ca2229127c55045ee41e5d9a0f21d.
//
// Solidity: event Bond(address indexed owner, uint256 amount)
func (_Storage *StorageFilterer) WatchBond(opts *bind.WatchOpts, sink chan<- *StorageBond, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Bond", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageBond)
				if err := _Storage.contract.UnpackLog(event, "Bond", log); err != nil {
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

// ParseBond is a log parse operation binding the contract event 0x6b1d99469ed62a423d7e402bfa68a467261ca2229127c55045ee41e5d9a0f21d.
//
// Solidity: event Bond(address indexed owner, uint256 amount)
func (_Storage *StorageFilterer) ParseBond(log types.Log) (*StorageBond, error) {
	event := new(StorageBond)
	if err := _Storage.contract.UnpackLog(event, "Bond", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageDAppDisabledIterator is returned from FilterDAppDisabled and is used to iterate over the raw logs and unpacked data for DAppDisabled events raised by the Storage contract.
type StorageDAppDisabledIterator struct {
	Event *StorageDAppDisabled // Event containing the contract specifics and raw log

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
func (it *StorageDAppDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageDAppDisabled)
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
		it.Event = new(StorageDAppDisabled)
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
func (it *StorageDAppDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageDAppDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageDAppDisabled represents a DAppDisabled event raised by the Storage contract.
type StorageDAppDisabled struct {
	Controller common.Address
	Governance common.Address
	CallConfig uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDAppDisabled is a free log retrieval operation binding the contract event 0xfb402de0284ae6437f381fcd6b8da6639e155a56551b5a8b95d2ab6c4e007d0a.
//
// Solidity: event DAppDisabled(address indexed controller, address indexed governance, uint32 callConfig)
func (_Storage *StorageFilterer) FilterDAppDisabled(opts *bind.FilterOpts, controller []common.Address, governance []common.Address) (*StorageDAppDisabledIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}
	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "DAppDisabled", controllerRule, governanceRule)
	if err != nil {
		return nil, err
	}
	return &StorageDAppDisabledIterator{contract: _Storage.contract, event: "DAppDisabled", logs: logs, sub: sub}, nil
}

// WatchDAppDisabled is a free log subscription operation binding the contract event 0xfb402de0284ae6437f381fcd6b8da6639e155a56551b5a8b95d2ab6c4e007d0a.
//
// Solidity: event DAppDisabled(address indexed controller, address indexed governance, uint32 callConfig)
func (_Storage *StorageFilterer) WatchDAppDisabled(opts *bind.WatchOpts, sink chan<- *StorageDAppDisabled, controller []common.Address, governance []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}
	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "DAppDisabled", controllerRule, governanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageDAppDisabled)
				if err := _Storage.contract.UnpackLog(event, "DAppDisabled", log); err != nil {
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

// ParseDAppDisabled is a log parse operation binding the contract event 0xfb402de0284ae6437f381fcd6b8da6639e155a56551b5a8b95d2ab6c4e007d0a.
//
// Solidity: event DAppDisabled(address indexed controller, address indexed governance, uint32 callConfig)
func (_Storage *StorageFilterer) ParseDAppDisabled(log types.Log) (*StorageDAppDisabled, error) {
	event := new(StorageDAppDisabled)
	if err := _Storage.contract.UnpackLog(event, "DAppDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageDAppGovernanceChangedIterator is returned from FilterDAppGovernanceChanged and is used to iterate over the raw logs and unpacked data for DAppGovernanceChanged events raised by the Storage contract.
type StorageDAppGovernanceChangedIterator struct {
	Event *StorageDAppGovernanceChanged // Event containing the contract specifics and raw log

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
func (it *StorageDAppGovernanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageDAppGovernanceChanged)
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
		it.Event = new(StorageDAppGovernanceChanged)
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
func (it *StorageDAppGovernanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageDAppGovernanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageDAppGovernanceChanged represents a DAppGovernanceChanged event raised by the Storage contract.
type StorageDAppGovernanceChanged struct {
	Controller    common.Address
	OldGovernance common.Address
	NewGovernance common.Address
	CallConfig    uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDAppGovernanceChanged is a free log retrieval operation binding the contract event 0xcc02b55b78457369d35dc189a074d3fc4f96a1897b405c0f1c57789f782d8bd2.
//
// Solidity: event DAppGovernanceChanged(address indexed controller, address indexed oldGovernance, address indexed newGovernance, uint32 callConfig)
func (_Storage *StorageFilterer) FilterDAppGovernanceChanged(opts *bind.FilterOpts, controller []common.Address, oldGovernance []common.Address, newGovernance []common.Address) (*StorageDAppGovernanceChangedIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}
	var oldGovernanceRule []interface{}
	for _, oldGovernanceItem := range oldGovernance {
		oldGovernanceRule = append(oldGovernanceRule, oldGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "DAppGovernanceChanged", controllerRule, oldGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &StorageDAppGovernanceChangedIterator{contract: _Storage.contract, event: "DAppGovernanceChanged", logs: logs, sub: sub}, nil
}

// WatchDAppGovernanceChanged is a free log subscription operation binding the contract event 0xcc02b55b78457369d35dc189a074d3fc4f96a1897b405c0f1c57789f782d8bd2.
//
// Solidity: event DAppGovernanceChanged(address indexed controller, address indexed oldGovernance, address indexed newGovernance, uint32 callConfig)
func (_Storage *StorageFilterer) WatchDAppGovernanceChanged(opts *bind.WatchOpts, sink chan<- *StorageDAppGovernanceChanged, controller []common.Address, oldGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}
	var oldGovernanceRule []interface{}
	for _, oldGovernanceItem := range oldGovernance {
		oldGovernanceRule = append(oldGovernanceRule, oldGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "DAppGovernanceChanged", controllerRule, oldGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageDAppGovernanceChanged)
				if err := _Storage.contract.UnpackLog(event, "DAppGovernanceChanged", log); err != nil {
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

// ParseDAppGovernanceChanged is a log parse operation binding the contract event 0xcc02b55b78457369d35dc189a074d3fc4f96a1897b405c0f1c57789f782d8bd2.
//
// Solidity: event DAppGovernanceChanged(address indexed controller, address indexed oldGovernance, address indexed newGovernance, uint32 callConfig)
func (_Storage *StorageFilterer) ParseDAppGovernanceChanged(log types.Log) (*StorageDAppGovernanceChanged, error) {
	event := new(StorageDAppGovernanceChanged)
	if err := _Storage.contract.UnpackLog(event, "DAppGovernanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageExecutionEnvironmentCreatedIterator is returned from FilterExecutionEnvironmentCreated and is used to iterate over the raw logs and unpacked data for ExecutionEnvironmentCreated events raised by the Storage contract.
type StorageExecutionEnvironmentCreatedIterator struct {
	Event *StorageExecutionEnvironmentCreated // Event containing the contract specifics and raw log

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
func (it *StorageExecutionEnvironmentCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageExecutionEnvironmentCreated)
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
		it.Event = new(StorageExecutionEnvironmentCreated)
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
func (it *StorageExecutionEnvironmentCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageExecutionEnvironmentCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageExecutionEnvironmentCreated represents a ExecutionEnvironmentCreated event raised by the Storage contract.
type StorageExecutionEnvironmentCreated struct {
	User                 common.Address
	ExecutionEnvironment common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterExecutionEnvironmentCreated is a free log retrieval operation binding the contract event 0x6ed96358b086d2aca68c2e2e4dc23fb45421ac513a7fc3127e34569833b3c646.
//
// Solidity: event ExecutionEnvironmentCreated(address indexed user, address indexed executionEnvironment)
func (_Storage *StorageFilterer) FilterExecutionEnvironmentCreated(opts *bind.FilterOpts, user []common.Address, executionEnvironment []common.Address) (*StorageExecutionEnvironmentCreatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var executionEnvironmentRule []interface{}
	for _, executionEnvironmentItem := range executionEnvironment {
		executionEnvironmentRule = append(executionEnvironmentRule, executionEnvironmentItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "ExecutionEnvironmentCreated", userRule, executionEnvironmentRule)
	if err != nil {
		return nil, err
	}
	return &StorageExecutionEnvironmentCreatedIterator{contract: _Storage.contract, event: "ExecutionEnvironmentCreated", logs: logs, sub: sub}, nil
}

// WatchExecutionEnvironmentCreated is a free log subscription operation binding the contract event 0x6ed96358b086d2aca68c2e2e4dc23fb45421ac513a7fc3127e34569833b3c646.
//
// Solidity: event ExecutionEnvironmentCreated(address indexed user, address indexed executionEnvironment)
func (_Storage *StorageFilterer) WatchExecutionEnvironmentCreated(opts *bind.WatchOpts, sink chan<- *StorageExecutionEnvironmentCreated, user []common.Address, executionEnvironment []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var executionEnvironmentRule []interface{}
	for _, executionEnvironmentItem := range executionEnvironment {
		executionEnvironmentRule = append(executionEnvironmentRule, executionEnvironmentItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "ExecutionEnvironmentCreated", userRule, executionEnvironmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageExecutionEnvironmentCreated)
				if err := _Storage.contract.UnpackLog(event, "ExecutionEnvironmentCreated", log); err != nil {
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

// ParseExecutionEnvironmentCreated is a log parse operation binding the contract event 0x6ed96358b086d2aca68c2e2e4dc23fb45421ac513a7fc3127e34569833b3c646.
//
// Solidity: event ExecutionEnvironmentCreated(address indexed user, address indexed executionEnvironment)
func (_Storage *StorageFilterer) ParseExecutionEnvironmentCreated(log types.Log) (*StorageExecutionEnvironmentCreated, error) {
	event := new(StorageExecutionEnvironmentCreated)
	if err := _Storage.contract.UnpackLog(event, "ExecutionEnvironmentCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageGovernanceTransferStartedIterator is returned from FilterGovernanceTransferStarted and is used to iterate over the raw logs and unpacked data for GovernanceTransferStarted events raised by the Storage contract.
type StorageGovernanceTransferStartedIterator struct {
	Event *StorageGovernanceTransferStarted // Event containing the contract specifics and raw log

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
func (it *StorageGovernanceTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageGovernanceTransferStarted)
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
		it.Event = new(StorageGovernanceTransferStarted)
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
func (it *StorageGovernanceTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageGovernanceTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageGovernanceTransferStarted represents a GovernanceTransferStarted event raised by the Storage contract.
type StorageGovernanceTransferStarted struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferStarted is a free log retrieval operation binding the contract event 0x1c4bb4e3cab7b72da7eb9f0ae62554dda85dc7fb907c946ad2776095b95ac1ad.
//
// Solidity: event GovernanceTransferStarted(address indexed previousGovernance, address indexed newGovernance)
func (_Storage *StorageFilterer) FilterGovernanceTransferStarted(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*StorageGovernanceTransferStartedIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "GovernanceTransferStarted", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &StorageGovernanceTransferStartedIterator{contract: _Storage.contract, event: "GovernanceTransferStarted", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferStarted is a free log subscription operation binding the contract event 0x1c4bb4e3cab7b72da7eb9f0ae62554dda85dc7fb907c946ad2776095b95ac1ad.
//
// Solidity: event GovernanceTransferStarted(address indexed previousGovernance, address indexed newGovernance)
func (_Storage *StorageFilterer) WatchGovernanceTransferStarted(opts *bind.WatchOpts, sink chan<- *StorageGovernanceTransferStarted, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "GovernanceTransferStarted", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageGovernanceTransferStarted)
				if err := _Storage.contract.UnpackLog(event, "GovernanceTransferStarted", log); err != nil {
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

// ParseGovernanceTransferStarted is a log parse operation binding the contract event 0x1c4bb4e3cab7b72da7eb9f0ae62554dda85dc7fb907c946ad2776095b95ac1ad.
//
// Solidity: event GovernanceTransferStarted(address indexed previousGovernance, address indexed newGovernance)
func (_Storage *StorageFilterer) ParseGovernanceTransferStarted(log types.Log) (*StorageGovernanceTransferStarted, error) {
	event := new(StorageGovernanceTransferStarted)
	if err := _Storage.contract.UnpackLog(event, "GovernanceTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageGovernanceTransferredIterator is returned from FilterGovernanceTransferred and is used to iterate over the raw logs and unpacked data for GovernanceTransferred events raised by the Storage contract.
type StorageGovernanceTransferredIterator struct {
	Event *StorageGovernanceTransferred // Event containing the contract specifics and raw log

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
func (it *StorageGovernanceTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageGovernanceTransferred)
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
		it.Event = new(StorageGovernanceTransferred)
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
func (it *StorageGovernanceTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageGovernanceTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageGovernanceTransferred represents a GovernanceTransferred event raised by the Storage contract.
type StorageGovernanceTransferred struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferred is a free log retrieval operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_Storage *StorageFilterer) FilterGovernanceTransferred(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*StorageGovernanceTransferredIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &StorageGovernanceTransferredIterator{contract: _Storage.contract, event: "GovernanceTransferred", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferred is a free log subscription operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_Storage *StorageFilterer) WatchGovernanceTransferred(opts *bind.WatchOpts, sink chan<- *StorageGovernanceTransferred, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageGovernanceTransferred)
				if err := _Storage.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
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

// ParseGovernanceTransferred is a log parse operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_Storage *StorageFilterer) ParseGovernanceTransferred(log types.Log) (*StorageGovernanceTransferred, error) {
	event := new(StorageGovernanceTransferred)
	if err := _Storage.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageMetacallResultIterator is returned from FilterMetacallResult and is used to iterate over the raw logs and unpacked data for MetacallResult events raised by the Storage contract.
type StorageMetacallResultIterator struct {
	Event *StorageMetacallResult // Event containing the contract specifics and raw log

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
func (it *StorageMetacallResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageMetacallResult)
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
		it.Event = new(StorageMetacallResult)
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
func (it *StorageMetacallResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageMetacallResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageMetacallResult represents a MetacallResult event raised by the Storage contract.
type StorageMetacallResult struct {
	Bundler          common.Address
	User             common.Address
	WinningSolver    common.Address
	EthPaidToBundler *big.Int
	NetGasSurcharge  *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterMetacallResult is a free log retrieval operation binding the contract event 0x8795eacea74c6caf7c8d200679e03233b3cb1f7fa9cf78fa90ef33d80d941c6c.
//
// Solidity: event MetacallResult(address indexed bundler, address indexed user, address indexed winningSolver, uint256 ethPaidToBundler, uint256 netGasSurcharge)
func (_Storage *StorageFilterer) FilterMetacallResult(opts *bind.FilterOpts, bundler []common.Address, user []common.Address, winningSolver []common.Address) (*StorageMetacallResultIterator, error) {

	var bundlerRule []interface{}
	for _, bundlerItem := range bundler {
		bundlerRule = append(bundlerRule, bundlerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var winningSolverRule []interface{}
	for _, winningSolverItem := range winningSolver {
		winningSolverRule = append(winningSolverRule, winningSolverItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "MetacallResult", bundlerRule, userRule, winningSolverRule)
	if err != nil {
		return nil, err
	}
	return &StorageMetacallResultIterator{contract: _Storage.contract, event: "MetacallResult", logs: logs, sub: sub}, nil
}

// WatchMetacallResult is a free log subscription operation binding the contract event 0x8795eacea74c6caf7c8d200679e03233b3cb1f7fa9cf78fa90ef33d80d941c6c.
//
// Solidity: event MetacallResult(address indexed bundler, address indexed user, address indexed winningSolver, uint256 ethPaidToBundler, uint256 netGasSurcharge)
func (_Storage *StorageFilterer) WatchMetacallResult(opts *bind.WatchOpts, sink chan<- *StorageMetacallResult, bundler []common.Address, user []common.Address, winningSolver []common.Address) (event.Subscription, error) {

	var bundlerRule []interface{}
	for _, bundlerItem := range bundler {
		bundlerRule = append(bundlerRule, bundlerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var winningSolverRule []interface{}
	for _, winningSolverItem := range winningSolver {
		winningSolverRule = append(winningSolverRule, winningSolverItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "MetacallResult", bundlerRule, userRule, winningSolverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageMetacallResult)
				if err := _Storage.contract.UnpackLog(event, "MetacallResult", log); err != nil {
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

// ParseMetacallResult is a log parse operation binding the contract event 0x8795eacea74c6caf7c8d200679e03233b3cb1f7fa9cf78fa90ef33d80d941c6c.
//
// Solidity: event MetacallResult(address indexed bundler, address indexed user, address indexed winningSolver, uint256 ethPaidToBundler, uint256 netGasSurcharge)
func (_Storage *StorageFilterer) ParseMetacallResult(log types.Log) (*StorageMetacallResult, error) {
	event := new(StorageMetacallResult)
	if err := _Storage.contract.UnpackLog(event, "MetacallResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageNewDAppSignatoryIterator is returned from FilterNewDAppSignatory and is used to iterate over the raw logs and unpacked data for NewDAppSignatory events raised by the Storage contract.
type StorageNewDAppSignatoryIterator struct {
	Event *StorageNewDAppSignatory // Event containing the contract specifics and raw log

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
func (it *StorageNewDAppSignatoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageNewDAppSignatory)
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
		it.Event = new(StorageNewDAppSignatory)
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
func (it *StorageNewDAppSignatoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageNewDAppSignatoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageNewDAppSignatory represents a NewDAppSignatory event raised by the Storage contract.
type StorageNewDAppSignatory struct {
	Controller common.Address
	Governance common.Address
	Signatory  common.Address
	CallConfig uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewDAppSignatory is a free log retrieval operation binding the contract event 0xcacb2beeeba676b69bddabab0d5f66e2733cabc804f82afd92ea5beae71934b4.
//
// Solidity: event NewDAppSignatory(address indexed controller, address indexed governance, address indexed signatory, uint32 callConfig)
func (_Storage *StorageFilterer) FilterNewDAppSignatory(opts *bind.FilterOpts, controller []common.Address, governance []common.Address, signatory []common.Address) (*StorageNewDAppSignatoryIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}
	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}
	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "NewDAppSignatory", controllerRule, governanceRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return &StorageNewDAppSignatoryIterator{contract: _Storage.contract, event: "NewDAppSignatory", logs: logs, sub: sub}, nil
}

// WatchNewDAppSignatory is a free log subscription operation binding the contract event 0xcacb2beeeba676b69bddabab0d5f66e2733cabc804f82afd92ea5beae71934b4.
//
// Solidity: event NewDAppSignatory(address indexed controller, address indexed governance, address indexed signatory, uint32 callConfig)
func (_Storage *StorageFilterer) WatchNewDAppSignatory(opts *bind.WatchOpts, sink chan<- *StorageNewDAppSignatory, controller []common.Address, governance []common.Address, signatory []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}
	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}
	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "NewDAppSignatory", controllerRule, governanceRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageNewDAppSignatory)
				if err := _Storage.contract.UnpackLog(event, "NewDAppSignatory", log); err != nil {
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

// ParseNewDAppSignatory is a log parse operation binding the contract event 0xcacb2beeeba676b69bddabab0d5f66e2733cabc804f82afd92ea5beae71934b4.
//
// Solidity: event NewDAppSignatory(address indexed controller, address indexed governance, address indexed signatory, uint32 callConfig)
func (_Storage *StorageFilterer) ParseNewDAppSignatory(log types.Log) (*StorageNewDAppSignatory, error) {
	event := new(StorageNewDAppSignatory)
	if err := _Storage.contract.UnpackLog(event, "NewDAppSignatory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Storage contract.
type StorageRedeemIterator struct {
	Event *StorageRedeem // Event containing the contract specifics and raw log

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
func (it *StorageRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageRedeem)
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
		it.Event = new(StorageRedeem)
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
func (it *StorageRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageRedeem represents a Redeem event raised by the Storage contract.
type StorageRedeem struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0x222838db2794d11532d940e8dec38ae307ed0b63cd97c233322e221f998767a6.
//
// Solidity: event Redeem(address indexed owner, uint256 amount)
func (_Storage *StorageFilterer) FilterRedeem(opts *bind.FilterOpts, owner []common.Address) (*StorageRedeemIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Redeem", ownerRule)
	if err != nil {
		return nil, err
	}
	return &StorageRedeemIterator{contract: _Storage.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0x222838db2794d11532d940e8dec38ae307ed0b63cd97c233322e221f998767a6.
//
// Solidity: event Redeem(address indexed owner, uint256 amount)
func (_Storage *StorageFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *StorageRedeem, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Redeem", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageRedeem)
				if err := _Storage.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0x222838db2794d11532d940e8dec38ae307ed0b63cd97c233322e221f998767a6.
//
// Solidity: event Redeem(address indexed owner, uint256 amount)
func (_Storage *StorageFilterer) ParseRedeem(log types.Log) (*StorageRedeem, error) {
	event := new(StorageRedeem)
	if err := _Storage.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageRemovedDAppSignatoryIterator is returned from FilterRemovedDAppSignatory and is used to iterate over the raw logs and unpacked data for RemovedDAppSignatory events raised by the Storage contract.
type StorageRemovedDAppSignatoryIterator struct {
	Event *StorageRemovedDAppSignatory // Event containing the contract specifics and raw log

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
func (it *StorageRemovedDAppSignatoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageRemovedDAppSignatory)
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
		it.Event = new(StorageRemovedDAppSignatory)
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
func (it *StorageRemovedDAppSignatoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageRemovedDAppSignatoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageRemovedDAppSignatory represents a RemovedDAppSignatory event raised by the Storage contract.
type StorageRemovedDAppSignatory struct {
	Controller common.Address
	Governance common.Address
	Signatory  common.Address
	CallConfig uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemovedDAppSignatory is a free log retrieval operation binding the contract event 0x90a400fe7fe676a24908d5f730fe6d38072340bc8f74e18a543132d840bfe5ce.
//
// Solidity: event RemovedDAppSignatory(address indexed controller, address indexed governance, address indexed signatory, uint32 callConfig)
func (_Storage *StorageFilterer) FilterRemovedDAppSignatory(opts *bind.FilterOpts, controller []common.Address, governance []common.Address, signatory []common.Address) (*StorageRemovedDAppSignatoryIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}
	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}
	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "RemovedDAppSignatory", controllerRule, governanceRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return &StorageRemovedDAppSignatoryIterator{contract: _Storage.contract, event: "RemovedDAppSignatory", logs: logs, sub: sub}, nil
}

// WatchRemovedDAppSignatory is a free log subscription operation binding the contract event 0x90a400fe7fe676a24908d5f730fe6d38072340bc8f74e18a543132d840bfe5ce.
//
// Solidity: event RemovedDAppSignatory(address indexed controller, address indexed governance, address indexed signatory, uint32 callConfig)
func (_Storage *StorageFilterer) WatchRemovedDAppSignatory(opts *bind.WatchOpts, sink chan<- *StorageRemovedDAppSignatory, controller []common.Address, governance []common.Address, signatory []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}
	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}
	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "RemovedDAppSignatory", controllerRule, governanceRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageRemovedDAppSignatory)
				if err := _Storage.contract.UnpackLog(event, "RemovedDAppSignatory", log); err != nil {
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

// ParseRemovedDAppSignatory is a log parse operation binding the contract event 0x90a400fe7fe676a24908d5f730fe6d38072340bc8f74e18a543132d840bfe5ce.
//
// Solidity: event RemovedDAppSignatory(address indexed controller, address indexed governance, address indexed signatory, uint32 callConfig)
func (_Storage *StorageFilterer) ParseRemovedDAppSignatory(log types.Log) (*StorageRemovedDAppSignatory, error) {
	event := new(StorageRemovedDAppSignatory)
	if err := _Storage.contract.UnpackLog(event, "RemovedDAppSignatory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSolverTxResultIterator is returned from FilterSolverTxResult and is used to iterate over the raw logs and unpacked data for SolverTxResult events raised by the Storage contract.
type StorageSolverTxResultIterator struct {
	Event *StorageSolverTxResult // Event containing the contract specifics and raw log

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
func (it *StorageSolverTxResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageSolverTxResult)
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
		it.Event = new(StorageSolverTxResult)
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
func (it *StorageSolverTxResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageSolverTxResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageSolverTxResult represents a SolverTxResult event raised by the Storage contract.
type StorageSolverTxResult struct {
	SolverTo   common.Address
	SolverFrom common.Address
	Executed   bool
	Success    bool
	Result     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSolverTxResult is a free log retrieval operation binding the contract event 0xcb73ceef18491eb9eafd886c9b1e14b4cce79f4e50eb1168717912ff6174ef8a.
//
// Solidity: event SolverTxResult(address indexed solverTo, address indexed solverFrom, bool executed, bool success, uint256 result)
func (_Storage *StorageFilterer) FilterSolverTxResult(opts *bind.FilterOpts, solverTo []common.Address, solverFrom []common.Address) (*StorageSolverTxResultIterator, error) {

	var solverToRule []interface{}
	for _, solverToItem := range solverTo {
		solverToRule = append(solverToRule, solverToItem)
	}
	var solverFromRule []interface{}
	for _, solverFromItem := range solverFrom {
		solverFromRule = append(solverFromRule, solverFromItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "SolverTxResult", solverToRule, solverFromRule)
	if err != nil {
		return nil, err
	}
	return &StorageSolverTxResultIterator{contract: _Storage.contract, event: "SolverTxResult", logs: logs, sub: sub}, nil
}

// WatchSolverTxResult is a free log subscription operation binding the contract event 0xcb73ceef18491eb9eafd886c9b1e14b4cce79f4e50eb1168717912ff6174ef8a.
//
// Solidity: event SolverTxResult(address indexed solverTo, address indexed solverFrom, bool executed, bool success, uint256 result)
func (_Storage *StorageFilterer) WatchSolverTxResult(opts *bind.WatchOpts, sink chan<- *StorageSolverTxResult, solverTo []common.Address, solverFrom []common.Address) (event.Subscription, error) {

	var solverToRule []interface{}
	for _, solverToItem := range solverTo {
		solverToRule = append(solverToRule, solverToItem)
	}
	var solverFromRule []interface{}
	for _, solverFromItem := range solverFrom {
		solverFromRule = append(solverFromRule, solverFromItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "SolverTxResult", solverToRule, solverFromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageSolverTxResult)
				if err := _Storage.contract.UnpackLog(event, "SolverTxResult", log); err != nil {
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

// ParseSolverTxResult is a log parse operation binding the contract event 0xcb73ceef18491eb9eafd886c9b1e14b4cce79f4e50eb1168717912ff6174ef8a.
//
// Solidity: event SolverTxResult(address indexed solverTo, address indexed solverFrom, bool executed, bool success, uint256 result)
func (_Storage *StorageFilterer) ParseSolverTxResult(log types.Log) (*StorageSolverTxResult, error) {
	event := new(StorageSolverTxResult)
	if err := _Storage.contract.UnpackLog(event, "SolverTxResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSurchargeRecipientTransferStartedIterator is returned from FilterSurchargeRecipientTransferStarted and is used to iterate over the raw logs and unpacked data for SurchargeRecipientTransferStarted events raised by the Storage contract.
type StorageSurchargeRecipientTransferStartedIterator struct {
	Event *StorageSurchargeRecipientTransferStarted // Event containing the contract specifics and raw log

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
func (it *StorageSurchargeRecipientTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageSurchargeRecipientTransferStarted)
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
		it.Event = new(StorageSurchargeRecipientTransferStarted)
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
func (it *StorageSurchargeRecipientTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageSurchargeRecipientTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageSurchargeRecipientTransferStarted represents a SurchargeRecipientTransferStarted event raised by the Storage contract.
type StorageSurchargeRecipientTransferStarted struct {
	CurrentRecipient common.Address
	NewRecipient     common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSurchargeRecipientTransferStarted is a free log retrieval operation binding the contract event 0xfc722bcd56a71612cd14b287fbf50545429e6c9e8de86ea7c3febdecd34882fd.
//
// Solidity: event SurchargeRecipientTransferStarted(address currentRecipient, address newRecipient)
func (_Storage *StorageFilterer) FilterSurchargeRecipientTransferStarted(opts *bind.FilterOpts) (*StorageSurchargeRecipientTransferStartedIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "SurchargeRecipientTransferStarted")
	if err != nil {
		return nil, err
	}
	return &StorageSurchargeRecipientTransferStartedIterator{contract: _Storage.contract, event: "SurchargeRecipientTransferStarted", logs: logs, sub: sub}, nil
}

// WatchSurchargeRecipientTransferStarted is a free log subscription operation binding the contract event 0xfc722bcd56a71612cd14b287fbf50545429e6c9e8de86ea7c3febdecd34882fd.
//
// Solidity: event SurchargeRecipientTransferStarted(address currentRecipient, address newRecipient)
func (_Storage *StorageFilterer) WatchSurchargeRecipientTransferStarted(opts *bind.WatchOpts, sink chan<- *StorageSurchargeRecipientTransferStarted) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "SurchargeRecipientTransferStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageSurchargeRecipientTransferStarted)
				if err := _Storage.contract.UnpackLog(event, "SurchargeRecipientTransferStarted", log); err != nil {
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

// ParseSurchargeRecipientTransferStarted is a log parse operation binding the contract event 0xfc722bcd56a71612cd14b287fbf50545429e6c9e8de86ea7c3febdecd34882fd.
//
// Solidity: event SurchargeRecipientTransferStarted(address currentRecipient, address newRecipient)
func (_Storage *StorageFilterer) ParseSurchargeRecipientTransferStarted(log types.Log) (*StorageSurchargeRecipientTransferStarted, error) {
	event := new(StorageSurchargeRecipientTransferStarted)
	if err := _Storage.contract.UnpackLog(event, "SurchargeRecipientTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSurchargeRecipientTransferredIterator is returned from FilterSurchargeRecipientTransferred and is used to iterate over the raw logs and unpacked data for SurchargeRecipientTransferred events raised by the Storage contract.
type StorageSurchargeRecipientTransferredIterator struct {
	Event *StorageSurchargeRecipientTransferred // Event containing the contract specifics and raw log

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
func (it *StorageSurchargeRecipientTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageSurchargeRecipientTransferred)
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
		it.Event = new(StorageSurchargeRecipientTransferred)
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
func (it *StorageSurchargeRecipientTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageSurchargeRecipientTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageSurchargeRecipientTransferred represents a SurchargeRecipientTransferred event raised by the Storage contract.
type StorageSurchargeRecipientTransferred struct {
	NewRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSurchargeRecipientTransferred is a free log retrieval operation binding the contract event 0x53960c2e64e72b2c1326635f0c002d5cf63e7844d12ed107404693fedde43985.
//
// Solidity: event SurchargeRecipientTransferred(address newRecipient)
func (_Storage *StorageFilterer) FilterSurchargeRecipientTransferred(opts *bind.FilterOpts) (*StorageSurchargeRecipientTransferredIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "SurchargeRecipientTransferred")
	if err != nil {
		return nil, err
	}
	return &StorageSurchargeRecipientTransferredIterator{contract: _Storage.contract, event: "SurchargeRecipientTransferred", logs: logs, sub: sub}, nil
}

// WatchSurchargeRecipientTransferred is a free log subscription operation binding the contract event 0x53960c2e64e72b2c1326635f0c002d5cf63e7844d12ed107404693fedde43985.
//
// Solidity: event SurchargeRecipientTransferred(address newRecipient)
func (_Storage *StorageFilterer) WatchSurchargeRecipientTransferred(opts *bind.WatchOpts, sink chan<- *StorageSurchargeRecipientTransferred) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "SurchargeRecipientTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageSurchargeRecipientTransferred)
				if err := _Storage.contract.UnpackLog(event, "SurchargeRecipientTransferred", log); err != nil {
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

// ParseSurchargeRecipientTransferred is a log parse operation binding the contract event 0x53960c2e64e72b2c1326635f0c002d5cf63e7844d12ed107404693fedde43985.
//
// Solidity: event SurchargeRecipientTransferred(address newRecipient)
func (_Storage *StorageFilterer) ParseSurchargeRecipientTransferred(log types.Log) (*StorageSurchargeRecipientTransferred, error) {
	event := new(StorageSurchargeRecipientTransferred)
	if err := _Storage.contract.UnpackLog(event, "SurchargeRecipientTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSurchargeWithdrawnIterator is returned from FilterSurchargeWithdrawn and is used to iterate over the raw logs and unpacked data for SurchargeWithdrawn events raised by the Storage contract.
type StorageSurchargeWithdrawnIterator struct {
	Event *StorageSurchargeWithdrawn // Event containing the contract specifics and raw log

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
func (it *StorageSurchargeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageSurchargeWithdrawn)
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
		it.Event = new(StorageSurchargeWithdrawn)
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
func (it *StorageSurchargeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageSurchargeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageSurchargeWithdrawn represents a SurchargeWithdrawn event raised by the Storage contract.
type StorageSurchargeWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSurchargeWithdrawn is a free log retrieval operation binding the contract event 0x87fa2ce024d3bdae31517696c13905fc0882bc1b4f4508060eb29358056de22b.
//
// Solidity: event SurchargeWithdrawn(address to, uint256 amount)
func (_Storage *StorageFilterer) FilterSurchargeWithdrawn(opts *bind.FilterOpts) (*StorageSurchargeWithdrawnIterator, error) {

	logs, sub, err := _Storage.contract.FilterLogs(opts, "SurchargeWithdrawn")
	if err != nil {
		return nil, err
	}
	return &StorageSurchargeWithdrawnIterator{contract: _Storage.contract, event: "SurchargeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchSurchargeWithdrawn is a free log subscription operation binding the contract event 0x87fa2ce024d3bdae31517696c13905fc0882bc1b4f4508060eb29358056de22b.
//
// Solidity: event SurchargeWithdrawn(address to, uint256 amount)
func (_Storage *StorageFilterer) WatchSurchargeWithdrawn(opts *bind.WatchOpts, sink chan<- *StorageSurchargeWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Storage.contract.WatchLogs(opts, "SurchargeWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageSurchargeWithdrawn)
				if err := _Storage.contract.UnpackLog(event, "SurchargeWithdrawn", log); err != nil {
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

// ParseSurchargeWithdrawn is a log parse operation binding the contract event 0x87fa2ce024d3bdae31517696c13905fc0882bc1b4f4508060eb29358056de22b.
//
// Solidity: event SurchargeWithdrawn(address to, uint256 amount)
func (_Storage *StorageFilterer) ParseSurchargeWithdrawn(log types.Log) (*StorageSurchargeWithdrawn, error) {
	event := new(StorageSurchargeWithdrawn)
	if err := _Storage.contract.UnpackLog(event, "SurchargeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Storage contract.
type StorageTransferIterator struct {
	Event *StorageTransfer // Event containing the contract specifics and raw log

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
func (it *StorageTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageTransfer)
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
		it.Event = new(StorageTransfer)
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
func (it *StorageTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageTransfer represents a Transfer event raised by the Storage contract.
type StorageTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Storage *StorageFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StorageTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StorageTransferIterator{contract: _Storage.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Storage *StorageFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StorageTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageTransfer)
				if err := _Storage.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Storage *StorageFilterer) ParseTransfer(log types.Log) (*StorageTransfer, error) {
	event := new(StorageTransfer)
	if err := _Storage.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageUnbondIterator is returned from FilterUnbond and is used to iterate over the raw logs and unpacked data for Unbond events raised by the Storage contract.
type StorageUnbondIterator struct {
	Event *StorageUnbond // Event containing the contract specifics and raw log

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
func (it *StorageUnbondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageUnbond)
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
		it.Event = new(StorageUnbond)
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
func (it *StorageUnbondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageUnbondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageUnbond represents a Unbond event raised by the Storage contract.
type StorageUnbond struct {
	Owner             common.Address
	Amount            *big.Int
	EarliestAvailable *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUnbond is a free log retrieval operation binding the contract event 0x7659747cd8571f1071eea946fdc648adcf181cad916f32a05f82c3a525976048.
//
// Solidity: event Unbond(address indexed owner, uint256 amount, uint256 earliestAvailable)
func (_Storage *StorageFilterer) FilterUnbond(opts *bind.FilterOpts, owner []common.Address) (*StorageUnbondIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Storage.contract.FilterLogs(opts, "Unbond", ownerRule)
	if err != nil {
		return nil, err
	}
	return &StorageUnbondIterator{contract: _Storage.contract, event: "Unbond", logs: logs, sub: sub}, nil
}

// WatchUnbond is a free log subscription operation binding the contract event 0x7659747cd8571f1071eea946fdc648adcf181cad916f32a05f82c3a525976048.
//
// Solidity: event Unbond(address indexed owner, uint256 amount, uint256 earliestAvailable)
func (_Storage *StorageFilterer) WatchUnbond(opts *bind.WatchOpts, sink chan<- *StorageUnbond, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Storage.contract.WatchLogs(opts, "Unbond", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageUnbond)
				if err := _Storage.contract.UnpackLog(event, "Unbond", log); err != nil {
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

// ParseUnbond is a log parse operation binding the contract event 0x7659747cd8571f1071eea946fdc648adcf181cad916f32a05f82c3a525976048.
//
// Solidity: event Unbond(address indexed owner, uint256 amount, uint256 earliestAvailable)
func (_Storage *StorageFilterer) ParseUnbond(log types.Log) (*StorageUnbond, error) {
	event := new(StorageUnbond)
	if err := _Storage.contract.UnpackLog(event, "Unbond", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
