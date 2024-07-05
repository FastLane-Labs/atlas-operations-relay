// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package atlas

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

// Context is an auto generated low-level Go binding around an user-defined struct.
type Context struct {
	UserOpHash           [32]byte
	ExecutionEnvironment common.Address
	SolverOutcome        *big.Int
	SolverIndex          uint8
	SolverCount          uint8
	CallDepth            uint8
	Phase                uint8
	SolverSuccessful     bool
	PaymentsSuccessful   bool
	BidFind              bool
	IsSimulation         bool
	Bundler              common.Address
}

// DAppConfig is an auto generated low-level Go binding around an user-defined struct.
type DAppConfig struct {
	To             common.Address
	CallConfig     uint32
	BidToken       common.Address
	SolverGasLimit uint32
}

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

// SolverTracker is an auto generated low-level Go binding around an user-defined struct.
type SolverTracker struct {
	BidAmount       *big.Int
	Floor           *big.Int
	Ceiling         *big.Int
	EtherIsBidToken bool
	InvertsBidValue bool
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

// AtlasMetaData contains all meta data concerning the Atlas contract.
var AtlasMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"escrowDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verification\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"simulator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"initialSurchargeRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"executionTemplate\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AllocateValueDelegatecallFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllocateValueFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllocateValueSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlteredControl\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AtlasLockActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceNotReconciled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"BidFindSuccessful\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BidNotPaid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"indexInSolverOps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"}],\"name\":\"BidTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BothPreOpsAndUserReturnDataCannotBeTracked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BothUserAndDAppNoncesCannotBeSequential\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallbackNotCalled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DAppNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DoubleReconcile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnvironmentMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EscrowLockActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExecutionEnvironmentBalanceTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientAtlETHBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"InsufficientAvailableBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalanceForDeduction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"InsufficientBondedBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientEscrow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLocalFunds\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"holds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientSolverBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"InsufficientSurchargeBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shortfall\",\"type\":\"uint256\"}],\"name\":\"InsufficientTotalBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"InsufficientUnbondedBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requested\",\"type\":\"uint256\"}],\"name\":\"InsufficientWithdrawableBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAccess\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCodeHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidControl\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDAppControl\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEntry\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEntryFunction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEnvironment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEscrowDuration\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"correctEnvironment\",\"type\":\"address\"}],\"name\":\"InvalidExecutionEnvironment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLockState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignatory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSolver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"solverFrom\",\"type\":\"address\"}],\"name\":\"InvalidSolverFrom\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidUser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvertBidValueCannotBeExPostBids\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"LedgerBalancing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"LedgerFinalized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"name\":\"MissingFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeDelegatecalled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAuctionWinner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoDelegatecall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoUnfilledRequests\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoUnusedNonceInBitmap\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnvironmentOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotImplemented\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAtlas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PermitDeadlineExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PostOpsDelegatecallFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PostOpsDelegatecallReturnedFalse\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PostOpsFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PostOpsSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PostSolverFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreOpsDelegatecallFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreOpsFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreOpsSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PreSolverFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatoryActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SimulationPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverMustReconcile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SolverOpReverted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"solverOutcomeResult\",\"type\":\"uint256\"}],\"name\":\"SolverSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnbalancedAccounting\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UncoveredResult\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedNonRevert\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unreachable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserNotFulfilled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserOpFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserOpSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserOpValueExceedsBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserSimulationFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserSimulationSucceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserUnexpectedSuccess\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserWrapperCallFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserWrapperDelegatecallFail\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumValidCallsResult\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"ValidCalls\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumValidCallsResult\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"VerificationSimFail\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongDepth\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongPhase\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Bond\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"name\":\"DAppDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"name\":\"DAppGovernanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executionEnvironment\",\"type\":\"address\"}],\"name\":\"ExecutionEnvironmentCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bundler\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"solverSuccessful\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"disbursementSuccessful\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ethPaidToBundler\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"netGasSurcharge\",\"type\":\"uint256\"}],\"name\":\"MetacallResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"name\":\"NewDAppSignatory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"name\":\"RemovedDAppSignatory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"solverTo\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"solverFrom\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"name\":\"SolverTxResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currentRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"}],\"name\":\"SurchargeRecipientTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"}],\"name\":\"SurchargeRecipientTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SurchargeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"earliestAvailable\",\"type\":\"uint256\"}],\"name\":\"Unbond\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ATLAS_SURCHARGE_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BUNDLER_SURCHARGE_RATE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ESCROW_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXECUTION_ENV_TEMPLATE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FIXED_GAS_OFFSET\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIMULATOR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SURCHARGE_SCALE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"S_accessData\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"bonded\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"lastAccessedBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"auctionWins\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"auctionFails\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"totalGasUsed\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"S_bondedTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"S_totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFICATION\",\"outputs\":[{\"internalType\":\"contractAtlasVerification\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accessData\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"bonded\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"lastAccessedBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint24\",\"name\":\"auctionWins\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"auctionFails\",\"type\":\"uint24\"},{\"internalType\":\"uint64\",\"name\":\"totalGasUsed\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"accountLastActiveBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeEnvironment\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOfBonded\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOfUnbonding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"becomeSurchargeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bondedTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claims\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contribute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"}],\"name\":\"createExecutionEnvironment\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"executionEnvironment\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cumulativeSurcharge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToBond\",\"type\":\"uint256\"}],\"name\":\"depositAndBond\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"solverGasLimit\",\"type\":\"uint32\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation[]\",\"name\":\"solverOps\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"executionEnvironment\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bundler\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"executionEnvironment\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"solverOutcome\",\"type\":\"uint24\"},{\"internalType\":\"uint8\",\"name\":\"solverIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"solverCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"callDepth\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"solverSuccessful\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paymentsSuccessful\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"bidFind\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSimulation\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"bundler\",\"type\":\"address\"}],\"internalType\":\"structContext\",\"name\":\"ctx\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"}],\"name\":\"getExecutionEnvironment\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"executionEnvironment\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isUnlocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"activeEnvironment\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation[]\",\"name\":\"solverOps\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bundler\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"callChainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structDAppOperation\",\"name\":\"dAppOp\",\"type\":\"tuple\"}],\"name\":\"metacall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"auctionWon\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingSurchargeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"phase\",\"outputs\":[{\"internalType\":\"enumExecutionPhase\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxApprovedGasSpend\",\"type\":\"uint256\"}],\"name\":\"reconcile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"owed\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shortfall\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"solver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"executionEnvironment\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"solverOutcome\",\"type\":\"uint24\"},{\"internalType\":\"uint8\",\"name\":\"solverIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"solverCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"callDepth\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"phase\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"solverSuccessful\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paymentsSuccessful\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"bidFind\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isSimulation\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"bundler\",\"type\":\"address\"}],\"internalType\":\"structContext\",\"name\":\"ctx\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"solverCall\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"floor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ceiling\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"etherIsBidToken\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"invertsBidValue\",\"type\":\"bool\"}],\"internalType\":\"structSolverTracker\",\"name\":\"solverTracker\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"solverLock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"solverLockData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"currentSolver\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"calledBack\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"fulfilled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"opHash\",\"type\":\"bytes32\"}],\"name\":\"solverOpHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"surchargeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"}],\"name\":\"transferDAppERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"}],\"name\":\"transferSurchargeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"}],\"name\":\"transferUserERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unbond\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"unbondingCompleteBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawSurcharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"writeoffs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AtlasABI is the input ABI used to generate the binding from.
// Deprecated: Use AtlasMetaData.ABI instead.
var AtlasABI = AtlasMetaData.ABI

// Atlas is an auto generated Go binding around an Ethereum contract.
type Atlas struct {
	AtlasCaller     // Read-only binding to the contract
	AtlasTransactor // Write-only binding to the contract
	AtlasFilterer   // Log filterer for contract events
}

// AtlasCaller is an auto generated read-only Go binding around an Ethereum contract.
type AtlasCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtlasTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AtlasTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtlasFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AtlasFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AtlasSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AtlasSession struct {
	Contract     *Atlas            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AtlasCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AtlasCallerSession struct {
	Contract *AtlasCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AtlasTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AtlasTransactorSession struct {
	Contract     *AtlasTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AtlasRaw is an auto generated low-level Go binding around an Ethereum contract.
type AtlasRaw struct {
	Contract *Atlas // Generic contract binding to access the raw methods on
}

// AtlasCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AtlasCallerRaw struct {
	Contract *AtlasCaller // Generic read-only contract binding to access the raw methods on
}

// AtlasTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AtlasTransactorRaw struct {
	Contract *AtlasTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAtlas creates a new instance of Atlas, bound to a specific deployed contract.
func NewAtlas(address common.Address, backend bind.ContractBackend) (*Atlas, error) {
	contract, err := bindAtlas(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Atlas{AtlasCaller: AtlasCaller{contract: contract}, AtlasTransactor: AtlasTransactor{contract: contract}, AtlasFilterer: AtlasFilterer{contract: contract}}, nil
}

// NewAtlasCaller creates a new read-only instance of Atlas, bound to a specific deployed contract.
func NewAtlasCaller(address common.Address, caller bind.ContractCaller) (*AtlasCaller, error) {
	contract, err := bindAtlas(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AtlasCaller{contract: contract}, nil
}

// NewAtlasTransactor creates a new write-only instance of Atlas, bound to a specific deployed contract.
func NewAtlasTransactor(address common.Address, transactor bind.ContractTransactor) (*AtlasTransactor, error) {
	contract, err := bindAtlas(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AtlasTransactor{contract: contract}, nil
}

// NewAtlasFilterer creates a new log filterer instance of Atlas, bound to a specific deployed contract.
func NewAtlasFilterer(address common.Address, filterer bind.ContractFilterer) (*AtlasFilterer, error) {
	contract, err := bindAtlas(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AtlasFilterer{contract: contract}, nil
}

// bindAtlas binds a generic wrapper to an already deployed contract.
func bindAtlas(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AtlasMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Atlas *AtlasRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Atlas.Contract.AtlasCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Atlas *AtlasRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atlas.Contract.AtlasTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Atlas *AtlasRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Atlas.Contract.AtlasTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Atlas *AtlasCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Atlas.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Atlas *AtlasTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atlas.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Atlas *AtlasTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Atlas.Contract.contract.Transact(opts, method, params...)
}

// ATLASSURCHARGERATE is a free data retrieval call binding the contract method 0x3173d012.
//
// Solidity: function ATLAS_SURCHARGE_RATE() view returns(uint256)
func (_Atlas *AtlasCaller) ATLASSURCHARGERATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "ATLAS_SURCHARGE_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ATLASSURCHARGERATE is a free data retrieval call binding the contract method 0x3173d012.
//
// Solidity: function ATLAS_SURCHARGE_RATE() view returns(uint256)
func (_Atlas *AtlasSession) ATLASSURCHARGERATE() (*big.Int, error) {
	return _Atlas.Contract.ATLASSURCHARGERATE(&_Atlas.CallOpts)
}

// ATLASSURCHARGERATE is a free data retrieval call binding the contract method 0x3173d012.
//
// Solidity: function ATLAS_SURCHARGE_RATE() view returns(uint256)
func (_Atlas *AtlasCallerSession) ATLASSURCHARGERATE() (*big.Int, error) {
	return _Atlas.Contract.ATLASSURCHARGERATE(&_Atlas.CallOpts)
}

// BUNDLERSURCHARGERATE is a free data retrieval call binding the contract method 0x1df634ad.
//
// Solidity: function BUNDLER_SURCHARGE_RATE() view returns(uint256)
func (_Atlas *AtlasCaller) BUNDLERSURCHARGERATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "BUNDLER_SURCHARGE_RATE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BUNDLERSURCHARGERATE is a free data retrieval call binding the contract method 0x1df634ad.
//
// Solidity: function BUNDLER_SURCHARGE_RATE() view returns(uint256)
func (_Atlas *AtlasSession) BUNDLERSURCHARGERATE() (*big.Int, error) {
	return _Atlas.Contract.BUNDLERSURCHARGERATE(&_Atlas.CallOpts)
}

// BUNDLERSURCHARGERATE is a free data retrieval call binding the contract method 0x1df634ad.
//
// Solidity: function BUNDLER_SURCHARGE_RATE() view returns(uint256)
func (_Atlas *AtlasCallerSession) BUNDLERSURCHARGERATE() (*big.Int, error) {
	return _Atlas.Contract.BUNDLERSURCHARGERATE(&_Atlas.CallOpts)
}

// ESCROWDURATION is a free data retrieval call binding the contract method 0xa6efccf9.
//
// Solidity: function ESCROW_DURATION() view returns(uint256)
func (_Atlas *AtlasCaller) ESCROWDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "ESCROW_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ESCROWDURATION is a free data retrieval call binding the contract method 0xa6efccf9.
//
// Solidity: function ESCROW_DURATION() view returns(uint256)
func (_Atlas *AtlasSession) ESCROWDURATION() (*big.Int, error) {
	return _Atlas.Contract.ESCROWDURATION(&_Atlas.CallOpts)
}

// ESCROWDURATION is a free data retrieval call binding the contract method 0xa6efccf9.
//
// Solidity: function ESCROW_DURATION() view returns(uint256)
func (_Atlas *AtlasCallerSession) ESCROWDURATION() (*big.Int, error) {
	return _Atlas.Contract.ESCROWDURATION(&_Atlas.CallOpts)
}

// EXECUTIONENVTEMPLATE is a free data retrieval call binding the contract method 0xd57c4dbd.
//
// Solidity: function EXECUTION_ENV_TEMPLATE() view returns(address)
func (_Atlas *AtlasCaller) EXECUTIONENVTEMPLATE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "EXECUTION_ENV_TEMPLATE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EXECUTIONENVTEMPLATE is a free data retrieval call binding the contract method 0xd57c4dbd.
//
// Solidity: function EXECUTION_ENV_TEMPLATE() view returns(address)
func (_Atlas *AtlasSession) EXECUTIONENVTEMPLATE() (common.Address, error) {
	return _Atlas.Contract.EXECUTIONENVTEMPLATE(&_Atlas.CallOpts)
}

// EXECUTIONENVTEMPLATE is a free data retrieval call binding the contract method 0xd57c4dbd.
//
// Solidity: function EXECUTION_ENV_TEMPLATE() view returns(address)
func (_Atlas *AtlasCallerSession) EXECUTIONENVTEMPLATE() (common.Address, error) {
	return _Atlas.Contract.EXECUTIONENVTEMPLATE(&_Atlas.CallOpts)
}

// FIXEDGASOFFSET is a free data retrieval call binding the contract method 0x0019f274.
//
// Solidity: function FIXED_GAS_OFFSET() view returns(uint256)
func (_Atlas *AtlasCaller) FIXEDGASOFFSET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "FIXED_GAS_OFFSET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FIXEDGASOFFSET is a free data retrieval call binding the contract method 0x0019f274.
//
// Solidity: function FIXED_GAS_OFFSET() view returns(uint256)
func (_Atlas *AtlasSession) FIXEDGASOFFSET() (*big.Int, error) {
	return _Atlas.Contract.FIXEDGASOFFSET(&_Atlas.CallOpts)
}

// FIXEDGASOFFSET is a free data retrieval call binding the contract method 0x0019f274.
//
// Solidity: function FIXED_GAS_OFFSET() view returns(uint256)
func (_Atlas *AtlasCallerSession) FIXEDGASOFFSET() (*big.Int, error) {
	return _Atlas.Contract.FIXEDGASOFFSET(&_Atlas.CallOpts)
}

// SIMULATOR is a free data retrieval call binding the contract method 0x79b79765.
//
// Solidity: function SIMULATOR() view returns(address)
func (_Atlas *AtlasCaller) SIMULATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "SIMULATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SIMULATOR is a free data retrieval call binding the contract method 0x79b79765.
//
// Solidity: function SIMULATOR() view returns(address)
func (_Atlas *AtlasSession) SIMULATOR() (common.Address, error) {
	return _Atlas.Contract.SIMULATOR(&_Atlas.CallOpts)
}

// SIMULATOR is a free data retrieval call binding the contract method 0x79b79765.
//
// Solidity: function SIMULATOR() view returns(address)
func (_Atlas *AtlasCallerSession) SIMULATOR() (common.Address, error) {
	return _Atlas.Contract.SIMULATOR(&_Atlas.CallOpts)
}

// SURCHARGESCALE is a free data retrieval call binding the contract method 0xfa4b7e98.
//
// Solidity: function SURCHARGE_SCALE() view returns(uint256)
func (_Atlas *AtlasCaller) SURCHARGESCALE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "SURCHARGE_SCALE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SURCHARGESCALE is a free data retrieval call binding the contract method 0xfa4b7e98.
//
// Solidity: function SURCHARGE_SCALE() view returns(uint256)
func (_Atlas *AtlasSession) SURCHARGESCALE() (*big.Int, error) {
	return _Atlas.Contract.SURCHARGESCALE(&_Atlas.CallOpts)
}

// SURCHARGESCALE is a free data retrieval call binding the contract method 0xfa4b7e98.
//
// Solidity: function SURCHARGE_SCALE() view returns(uint256)
func (_Atlas *AtlasCallerSession) SURCHARGESCALE() (*big.Int, error) {
	return _Atlas.Contract.SURCHARGESCALE(&_Atlas.CallOpts)
}

// SAccessData is a free data retrieval call binding the contract method 0xc4730b1c.
//
// Solidity: function S_accessData(address ) view returns(uint112 bonded, uint32 lastAccessedBlock, uint24 auctionWins, uint24 auctionFails, uint64 totalGasUsed)
func (_Atlas *AtlasCaller) SAccessData(opts *bind.CallOpts, arg0 common.Address) (struct {
	Bonded            *big.Int
	LastAccessedBlock uint32
	AuctionWins       *big.Int
	AuctionFails      *big.Int
	TotalGasUsed      uint64
}, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "S_accessData", arg0)

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

// SAccessData is a free data retrieval call binding the contract method 0xc4730b1c.
//
// Solidity: function S_accessData(address ) view returns(uint112 bonded, uint32 lastAccessedBlock, uint24 auctionWins, uint24 auctionFails, uint64 totalGasUsed)
func (_Atlas *AtlasSession) SAccessData(arg0 common.Address) (struct {
	Bonded            *big.Int
	LastAccessedBlock uint32
	AuctionWins       *big.Int
	AuctionFails      *big.Int
	TotalGasUsed      uint64
}, error) {
	return _Atlas.Contract.SAccessData(&_Atlas.CallOpts, arg0)
}

// SAccessData is a free data retrieval call binding the contract method 0xc4730b1c.
//
// Solidity: function S_accessData(address ) view returns(uint112 bonded, uint32 lastAccessedBlock, uint24 auctionWins, uint24 auctionFails, uint64 totalGasUsed)
func (_Atlas *AtlasCallerSession) SAccessData(arg0 common.Address) (struct {
	Bonded            *big.Int
	LastAccessedBlock uint32
	AuctionWins       *big.Int
	AuctionFails      *big.Int
	TotalGasUsed      uint64
}, error) {
	return _Atlas.Contract.SAccessData(&_Atlas.CallOpts, arg0)
}

// SBondedTotalSupply is a free data retrieval call binding the contract method 0xaca06e89.
//
// Solidity: function S_bondedTotalSupply() view returns(uint256)
func (_Atlas *AtlasCaller) SBondedTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "S_bondedTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SBondedTotalSupply is a free data retrieval call binding the contract method 0xaca06e89.
//
// Solidity: function S_bondedTotalSupply() view returns(uint256)
func (_Atlas *AtlasSession) SBondedTotalSupply() (*big.Int, error) {
	return _Atlas.Contract.SBondedTotalSupply(&_Atlas.CallOpts)
}

// SBondedTotalSupply is a free data retrieval call binding the contract method 0xaca06e89.
//
// Solidity: function S_bondedTotalSupply() view returns(uint256)
func (_Atlas *AtlasCallerSession) SBondedTotalSupply() (*big.Int, error) {
	return _Atlas.Contract.SBondedTotalSupply(&_Atlas.CallOpts)
}

// STotalSupply is a free data retrieval call binding the contract method 0x31a52770.
//
// Solidity: function S_totalSupply() view returns(uint256)
func (_Atlas *AtlasCaller) STotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "S_totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STotalSupply is a free data retrieval call binding the contract method 0x31a52770.
//
// Solidity: function S_totalSupply() view returns(uint256)
func (_Atlas *AtlasSession) STotalSupply() (*big.Int, error) {
	return _Atlas.Contract.STotalSupply(&_Atlas.CallOpts)
}

// STotalSupply is a free data retrieval call binding the contract method 0x31a52770.
//
// Solidity: function S_totalSupply() view returns(uint256)
func (_Atlas *AtlasCallerSession) STotalSupply() (*big.Int, error) {
	return _Atlas.Contract.STotalSupply(&_Atlas.CallOpts)
}

// VERIFICATION is a free data retrieval call binding the contract method 0x791ae748.
//
// Solidity: function VERIFICATION() view returns(address)
func (_Atlas *AtlasCaller) VERIFICATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "VERIFICATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VERIFICATION is a free data retrieval call binding the contract method 0x791ae748.
//
// Solidity: function VERIFICATION() view returns(address)
func (_Atlas *AtlasSession) VERIFICATION() (common.Address, error) {
	return _Atlas.Contract.VERIFICATION(&_Atlas.CallOpts)
}

// VERIFICATION is a free data retrieval call binding the contract method 0x791ae748.
//
// Solidity: function VERIFICATION() view returns(address)
func (_Atlas *AtlasCallerSession) VERIFICATION() (common.Address, error) {
	return _Atlas.Contract.VERIFICATION(&_Atlas.CallOpts)
}

// AccessData is a free data retrieval call binding the contract method 0x5e8edccc.
//
// Solidity: function accessData(address account) view returns(uint112 bonded, uint32 lastAccessedBlock, uint24 auctionWins, uint24 auctionFails, uint64 totalGasUsed)
func (_Atlas *AtlasCaller) AccessData(opts *bind.CallOpts, account common.Address) (struct {
	Bonded            *big.Int
	LastAccessedBlock uint32
	AuctionWins       *big.Int
	AuctionFails      *big.Int
	TotalGasUsed      uint64
}, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "accessData", account)

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
// Solidity: function accessData(address account) view returns(uint112 bonded, uint32 lastAccessedBlock, uint24 auctionWins, uint24 auctionFails, uint64 totalGasUsed)
func (_Atlas *AtlasSession) AccessData(account common.Address) (struct {
	Bonded            *big.Int
	LastAccessedBlock uint32
	AuctionWins       *big.Int
	AuctionFails      *big.Int
	TotalGasUsed      uint64
}, error) {
	return _Atlas.Contract.AccessData(&_Atlas.CallOpts, account)
}

// AccessData is a free data retrieval call binding the contract method 0x5e8edccc.
//
// Solidity: function accessData(address account) view returns(uint112 bonded, uint32 lastAccessedBlock, uint24 auctionWins, uint24 auctionFails, uint64 totalGasUsed)
func (_Atlas *AtlasCallerSession) AccessData(account common.Address) (struct {
	Bonded            *big.Int
	LastAccessedBlock uint32
	AuctionWins       *big.Int
	AuctionFails      *big.Int
	TotalGasUsed      uint64
}, error) {
	return _Atlas.Contract.AccessData(&_Atlas.CallOpts, account)
}

// AccountLastActiveBlock is a free data retrieval call binding the contract method 0x7c20857a.
//
// Solidity: function accountLastActiveBlock(address account) view returns(uint256)
func (_Atlas *AtlasCaller) AccountLastActiveBlock(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "accountLastActiveBlock", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountLastActiveBlock is a free data retrieval call binding the contract method 0x7c20857a.
//
// Solidity: function accountLastActiveBlock(address account) view returns(uint256)
func (_Atlas *AtlasSession) AccountLastActiveBlock(account common.Address) (*big.Int, error) {
	return _Atlas.Contract.AccountLastActiveBlock(&_Atlas.CallOpts, account)
}

// AccountLastActiveBlock is a free data retrieval call binding the contract method 0x7c20857a.
//
// Solidity: function accountLastActiveBlock(address account) view returns(uint256)
func (_Atlas *AtlasCallerSession) AccountLastActiveBlock(account common.Address) (*big.Int, error) {
	return _Atlas.Contract.AccountLastActiveBlock(&_Atlas.CallOpts, account)
}

// ActiveEnvironment is a free data retrieval call binding the contract method 0x6ea43423.
//
// Solidity: function activeEnvironment() view returns(address)
func (_Atlas *AtlasCaller) ActiveEnvironment(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "activeEnvironment")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ActiveEnvironment is a free data retrieval call binding the contract method 0x6ea43423.
//
// Solidity: function activeEnvironment() view returns(address)
func (_Atlas *AtlasSession) ActiveEnvironment() (common.Address, error) {
	return _Atlas.Contract.ActiveEnvironment(&_Atlas.CallOpts)
}

// ActiveEnvironment is a free data retrieval call binding the contract method 0x6ea43423.
//
// Solidity: function activeEnvironment() view returns(address)
func (_Atlas *AtlasCallerSession) ActiveEnvironment() (common.Address, error) {
	return _Atlas.Contract.ActiveEnvironment(&_Atlas.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Atlas *AtlasCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Atlas *AtlasSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Atlas.Contract.BalanceOf(&_Atlas.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Atlas *AtlasCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Atlas.Contract.BalanceOf(&_Atlas.CallOpts, account)
}

// BalanceOfBonded is a free data retrieval call binding the contract method 0x825ad607.
//
// Solidity: function balanceOfBonded(address account) view returns(uint256)
func (_Atlas *AtlasCaller) BalanceOfBonded(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "balanceOfBonded", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfBonded is a free data retrieval call binding the contract method 0x825ad607.
//
// Solidity: function balanceOfBonded(address account) view returns(uint256)
func (_Atlas *AtlasSession) BalanceOfBonded(account common.Address) (*big.Int, error) {
	return _Atlas.Contract.BalanceOfBonded(&_Atlas.CallOpts, account)
}

// BalanceOfBonded is a free data retrieval call binding the contract method 0x825ad607.
//
// Solidity: function balanceOfBonded(address account) view returns(uint256)
func (_Atlas *AtlasCallerSession) BalanceOfBonded(account common.Address) (*big.Int, error) {
	return _Atlas.Contract.BalanceOfBonded(&_Atlas.CallOpts, account)
}

// BalanceOfUnbonding is a free data retrieval call binding the contract method 0xaebaa5f7.
//
// Solidity: function balanceOfUnbonding(address account) view returns(uint256)
func (_Atlas *AtlasCaller) BalanceOfUnbonding(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "balanceOfUnbonding", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfUnbonding is a free data retrieval call binding the contract method 0xaebaa5f7.
//
// Solidity: function balanceOfUnbonding(address account) view returns(uint256)
func (_Atlas *AtlasSession) BalanceOfUnbonding(account common.Address) (*big.Int, error) {
	return _Atlas.Contract.BalanceOfUnbonding(&_Atlas.CallOpts, account)
}

// BalanceOfUnbonding is a free data retrieval call binding the contract method 0xaebaa5f7.
//
// Solidity: function balanceOfUnbonding(address account) view returns(uint256)
func (_Atlas *AtlasCallerSession) BalanceOfUnbonding(account common.Address) (*big.Int, error) {
	return _Atlas.Contract.BalanceOfUnbonding(&_Atlas.CallOpts, account)
}

// BondedTotalSupply is a free data retrieval call binding the contract method 0x890c2854.
//
// Solidity: function bondedTotalSupply() view returns(uint256)
func (_Atlas *AtlasCaller) BondedTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "bondedTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BondedTotalSupply is a free data retrieval call binding the contract method 0x890c2854.
//
// Solidity: function bondedTotalSupply() view returns(uint256)
func (_Atlas *AtlasSession) BondedTotalSupply() (*big.Int, error) {
	return _Atlas.Contract.BondedTotalSupply(&_Atlas.CallOpts)
}

// BondedTotalSupply is a free data retrieval call binding the contract method 0x890c2854.
//
// Solidity: function bondedTotalSupply() view returns(uint256)
func (_Atlas *AtlasCallerSession) BondedTotalSupply() (*big.Int, error) {
	return _Atlas.Contract.BondedTotalSupply(&_Atlas.CallOpts)
}

// Claims is a free data retrieval call binding the contract method 0xdcc59b6f.
//
// Solidity: function claims() view returns(uint256)
func (_Atlas *AtlasCaller) Claims(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "claims")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claims is a free data retrieval call binding the contract method 0xdcc59b6f.
//
// Solidity: function claims() view returns(uint256)
func (_Atlas *AtlasSession) Claims() (*big.Int, error) {
	return _Atlas.Contract.Claims(&_Atlas.CallOpts)
}

// Claims is a free data retrieval call binding the contract method 0xdcc59b6f.
//
// Solidity: function claims() view returns(uint256)
func (_Atlas *AtlasCallerSession) Claims() (*big.Int, error) {
	return _Atlas.Contract.Claims(&_Atlas.CallOpts)
}

// CumulativeSurcharge is a free data retrieval call binding the contract method 0xc5471d9e.
//
// Solidity: function cumulativeSurcharge() view returns(uint256)
func (_Atlas *AtlasCaller) CumulativeSurcharge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "cumulativeSurcharge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CumulativeSurcharge is a free data retrieval call binding the contract method 0xc5471d9e.
//
// Solidity: function cumulativeSurcharge() view returns(uint256)
func (_Atlas *AtlasSession) CumulativeSurcharge() (*big.Int, error) {
	return _Atlas.Contract.CumulativeSurcharge(&_Atlas.CallOpts)
}

// CumulativeSurcharge is a free data retrieval call binding the contract method 0xc5471d9e.
//
// Solidity: function cumulativeSurcharge() view returns(uint256)
func (_Atlas *AtlasCallerSession) CumulativeSurcharge() (*big.Int, error) {
	return _Atlas.Contract.CumulativeSurcharge(&_Atlas.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Atlas *AtlasCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Atlas *AtlasSession) Decimals() (uint8, error) {
	return _Atlas.Contract.Decimals(&_Atlas.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Atlas *AtlasCallerSession) Decimals() (uint8, error) {
	return _Atlas.Contract.Decimals(&_Atlas.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0x323a5e0b.
//
// Solidity: function deposits() view returns(uint256)
func (_Atlas *AtlasCaller) Deposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "deposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposits is a free data retrieval call binding the contract method 0x323a5e0b.
//
// Solidity: function deposits() view returns(uint256)
func (_Atlas *AtlasSession) Deposits() (*big.Int, error) {
	return _Atlas.Contract.Deposits(&_Atlas.CallOpts)
}

// Deposits is a free data retrieval call binding the contract method 0x323a5e0b.
//
// Solidity: function deposits() view returns(uint256)
func (_Atlas *AtlasCallerSession) Deposits() (*big.Int, error) {
	return _Atlas.Contract.Deposits(&_Atlas.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint256)
func (_Atlas *AtlasCaller) Fees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "fees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint256)
func (_Atlas *AtlasSession) Fees() (*big.Int, error) {
	return _Atlas.Contract.Fees(&_Atlas.CallOpts)
}

// Fees is a free data retrieval call binding the contract method 0x9af1d35a.
//
// Solidity: function fees() view returns(uint256)
func (_Atlas *AtlasCallerSession) Fees() (*big.Int, error) {
	return _Atlas.Contract.Fees(&_Atlas.CallOpts)
}

// GetExecutionEnvironment is a free data retrieval call binding the contract method 0x45112906.
//
// Solidity: function getExecutionEnvironment(address user, address control) view returns(address executionEnvironment, uint32 callConfig, bool exists)
func (_Atlas *AtlasCaller) GetExecutionEnvironment(opts *bind.CallOpts, user common.Address, control common.Address) (struct {
	ExecutionEnvironment common.Address
	CallConfig           uint32
	Exists               bool
}, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "getExecutionEnvironment", user, control)

	outstruct := new(struct {
		ExecutionEnvironment common.Address
		CallConfig           uint32
		Exists               bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ExecutionEnvironment = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CallConfig = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Exists = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetExecutionEnvironment is a free data retrieval call binding the contract method 0x45112906.
//
// Solidity: function getExecutionEnvironment(address user, address control) view returns(address executionEnvironment, uint32 callConfig, bool exists)
func (_Atlas *AtlasSession) GetExecutionEnvironment(user common.Address, control common.Address) (struct {
	ExecutionEnvironment common.Address
	CallConfig           uint32
	Exists               bool
}, error) {
	return _Atlas.Contract.GetExecutionEnvironment(&_Atlas.CallOpts, user, control)
}

// GetExecutionEnvironment is a free data retrieval call binding the contract method 0x45112906.
//
// Solidity: function getExecutionEnvironment(address user, address control) view returns(address executionEnvironment, uint32 callConfig, bool exists)
func (_Atlas *AtlasCallerSession) GetExecutionEnvironment(user common.Address, control common.Address) (struct {
	ExecutionEnvironment common.Address
	CallConfig           uint32
	Exists               bool
}, error) {
	return _Atlas.Contract.GetExecutionEnvironment(&_Atlas.CallOpts, user, control)
}

// IsUnlocked is a free data retrieval call binding the contract method 0x8380edb7.
//
// Solidity: function isUnlocked() view returns(bool)
func (_Atlas *AtlasCaller) IsUnlocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "isUnlocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUnlocked is a free data retrieval call binding the contract method 0x8380edb7.
//
// Solidity: function isUnlocked() view returns(bool)
func (_Atlas *AtlasSession) IsUnlocked() (bool, error) {
	return _Atlas.Contract.IsUnlocked(&_Atlas.CallOpts)
}

// IsUnlocked is a free data retrieval call binding the contract method 0x8380edb7.
//
// Solidity: function isUnlocked() view returns(bool)
func (_Atlas *AtlasCallerSession) IsUnlocked() (bool, error) {
	return _Atlas.Contract.IsUnlocked(&_Atlas.CallOpts)
}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() view returns(address activeEnvironment, uint32 callConfig, uint8 phase)
func (_Atlas *AtlasCaller) Lock(opts *bind.CallOpts) (struct {
	ActiveEnvironment common.Address
	CallConfig        uint32
	Phase             uint8
}, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "lock")

	outstruct := new(struct {
		ActiveEnvironment common.Address
		CallConfig        uint32
		Phase             uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ActiveEnvironment = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CallConfig = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Phase = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() view returns(address activeEnvironment, uint32 callConfig, uint8 phase)
func (_Atlas *AtlasSession) Lock() (struct {
	ActiveEnvironment common.Address
	CallConfig        uint32
	Phase             uint8
}, error) {
	return _Atlas.Contract.Lock(&_Atlas.CallOpts)
}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() view returns(address activeEnvironment, uint32 callConfig, uint8 phase)
func (_Atlas *AtlasCallerSession) Lock() (struct {
	ActiveEnvironment common.Address
	CallConfig        uint32
	Phase             uint8
}, error) {
	return _Atlas.Contract.Lock(&_Atlas.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Atlas *AtlasCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Atlas *AtlasSession) Name() (string, error) {
	return _Atlas.Contract.Name(&_Atlas.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Atlas *AtlasCallerSession) Name() (string, error) {
	return _Atlas.Contract.Name(&_Atlas.CallOpts)
}

// PendingSurchargeRecipient is a free data retrieval call binding the contract method 0x7c3c3160.
//
// Solidity: function pendingSurchargeRecipient() view returns(address)
func (_Atlas *AtlasCaller) PendingSurchargeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "pendingSurchargeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingSurchargeRecipient is a free data retrieval call binding the contract method 0x7c3c3160.
//
// Solidity: function pendingSurchargeRecipient() view returns(address)
func (_Atlas *AtlasSession) PendingSurchargeRecipient() (common.Address, error) {
	return _Atlas.Contract.PendingSurchargeRecipient(&_Atlas.CallOpts)
}

// PendingSurchargeRecipient is a free data retrieval call binding the contract method 0x7c3c3160.
//
// Solidity: function pendingSurchargeRecipient() view returns(address)
func (_Atlas *AtlasCallerSession) PendingSurchargeRecipient() (common.Address, error) {
	return _Atlas.Contract.PendingSurchargeRecipient(&_Atlas.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() view returns(uint8)
func (_Atlas *AtlasCaller) Phase(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "phase")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() view returns(uint8)
func (_Atlas *AtlasSession) Phase() (uint8, error) {
	return _Atlas.Contract.Phase(&_Atlas.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() view returns(uint8)
func (_Atlas *AtlasCallerSession) Phase() (uint8, error) {
	return _Atlas.Contract.Phase(&_Atlas.CallOpts)
}

// Shortfall is a free data retrieval call binding the contract method 0x19b1faef.
//
// Solidity: function shortfall() view returns(uint256)
func (_Atlas *AtlasCaller) Shortfall(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "shortfall")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Shortfall is a free data retrieval call binding the contract method 0x19b1faef.
//
// Solidity: function shortfall() view returns(uint256)
func (_Atlas *AtlasSession) Shortfall() (*big.Int, error) {
	return _Atlas.Contract.Shortfall(&_Atlas.CallOpts)
}

// Shortfall is a free data retrieval call binding the contract method 0x19b1faef.
//
// Solidity: function shortfall() view returns(uint256)
func (_Atlas *AtlasCallerSession) Shortfall() (*big.Int, error) {
	return _Atlas.Contract.Shortfall(&_Atlas.CallOpts)
}

// Solver is a free data retrieval call binding the contract method 0x49a7a26d.
//
// Solidity: function solver() view returns(address)
func (_Atlas *AtlasCaller) Solver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "solver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Solver is a free data retrieval call binding the contract method 0x49a7a26d.
//
// Solidity: function solver() view returns(address)
func (_Atlas *AtlasSession) Solver() (common.Address, error) {
	return _Atlas.Contract.Solver(&_Atlas.CallOpts)
}

// Solver is a free data retrieval call binding the contract method 0x49a7a26d.
//
// Solidity: function solver() view returns(address)
func (_Atlas *AtlasCallerSession) Solver() (common.Address, error) {
	return _Atlas.Contract.Solver(&_Atlas.CallOpts)
}

// SolverLock is a free data retrieval call binding the contract method 0xa66520a3.
//
// Solidity: function solverLock() view returns(uint256)
func (_Atlas *AtlasCaller) SolverLock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "solverLock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SolverLock is a free data retrieval call binding the contract method 0xa66520a3.
//
// Solidity: function solverLock() view returns(uint256)
func (_Atlas *AtlasSession) SolverLock() (*big.Int, error) {
	return _Atlas.Contract.SolverLock(&_Atlas.CallOpts)
}

// SolverLock is a free data retrieval call binding the contract method 0xa66520a3.
//
// Solidity: function solverLock() view returns(uint256)
func (_Atlas *AtlasCallerSession) SolverLock() (*big.Int, error) {
	return _Atlas.Contract.SolverLock(&_Atlas.CallOpts)
}

// SolverLockData is a free data retrieval call binding the contract method 0xaa7d2dc8.
//
// Solidity: function solverLockData() view returns(address currentSolver, bool calledBack, bool fulfilled)
func (_Atlas *AtlasCaller) SolverLockData(opts *bind.CallOpts) (struct {
	CurrentSolver common.Address
	CalledBack    bool
	Fulfilled     bool
}, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "solverLockData")

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
func (_Atlas *AtlasSession) SolverLockData() (struct {
	CurrentSolver common.Address
	CalledBack    bool
	Fulfilled     bool
}, error) {
	return _Atlas.Contract.SolverLockData(&_Atlas.CallOpts)
}

// SolverLockData is a free data retrieval call binding the contract method 0xaa7d2dc8.
//
// Solidity: function solverLockData() view returns(address currentSolver, bool calledBack, bool fulfilled)
func (_Atlas *AtlasCallerSession) SolverLockData() (struct {
	CurrentSolver common.Address
	CalledBack    bool
	Fulfilled     bool
}, error) {
	return _Atlas.Contract.SolverLockData(&_Atlas.CallOpts)
}

// SolverOpHashes is a free data retrieval call binding the contract method 0x6ef5ac7a.
//
// Solidity: function solverOpHashes(bytes32 opHash) view returns(bool)
func (_Atlas *AtlasCaller) SolverOpHashes(opts *bind.CallOpts, opHash [32]byte) (bool, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "solverOpHashes", opHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SolverOpHashes is a free data retrieval call binding the contract method 0x6ef5ac7a.
//
// Solidity: function solverOpHashes(bytes32 opHash) view returns(bool)
func (_Atlas *AtlasSession) SolverOpHashes(opHash [32]byte) (bool, error) {
	return _Atlas.Contract.SolverOpHashes(&_Atlas.CallOpts, opHash)
}

// SolverOpHashes is a free data retrieval call binding the contract method 0x6ef5ac7a.
//
// Solidity: function solverOpHashes(bytes32 opHash) view returns(bool)
func (_Atlas *AtlasCallerSession) SolverOpHashes(opHash [32]byte) (bool, error) {
	return _Atlas.Contract.SolverOpHashes(&_Atlas.CallOpts, opHash)
}

// SurchargeRecipient is a free data retrieval call binding the contract method 0xfc61c541.
//
// Solidity: function surchargeRecipient() view returns(address)
func (_Atlas *AtlasCaller) SurchargeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "surchargeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SurchargeRecipient is a free data retrieval call binding the contract method 0xfc61c541.
//
// Solidity: function surchargeRecipient() view returns(address)
func (_Atlas *AtlasSession) SurchargeRecipient() (common.Address, error) {
	return _Atlas.Contract.SurchargeRecipient(&_Atlas.CallOpts)
}

// SurchargeRecipient is a free data retrieval call binding the contract method 0xfc61c541.
//
// Solidity: function surchargeRecipient() view returns(address)
func (_Atlas *AtlasCallerSession) SurchargeRecipient() (common.Address, error) {
	return _Atlas.Contract.SurchargeRecipient(&_Atlas.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Atlas *AtlasCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Atlas *AtlasSession) Symbol() (string, error) {
	return _Atlas.Contract.Symbol(&_Atlas.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Atlas *AtlasCallerSession) Symbol() (string, error) {
	return _Atlas.Contract.Symbol(&_Atlas.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Atlas *AtlasCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Atlas *AtlasSession) TotalSupply() (*big.Int, error) {
	return _Atlas.Contract.TotalSupply(&_Atlas.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Atlas *AtlasCallerSession) TotalSupply() (*big.Int, error) {
	return _Atlas.Contract.TotalSupply(&_Atlas.CallOpts)
}

// UnbondingCompleteBlock is a free data retrieval call binding the contract method 0x5270182c.
//
// Solidity: function unbondingCompleteBlock(address account) view returns(uint256)
func (_Atlas *AtlasCaller) UnbondingCompleteBlock(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "unbondingCompleteBlock", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnbondingCompleteBlock is a free data retrieval call binding the contract method 0x5270182c.
//
// Solidity: function unbondingCompleteBlock(address account) view returns(uint256)
func (_Atlas *AtlasSession) UnbondingCompleteBlock(account common.Address) (*big.Int, error) {
	return _Atlas.Contract.UnbondingCompleteBlock(&_Atlas.CallOpts, account)
}

// UnbondingCompleteBlock is a free data retrieval call binding the contract method 0x5270182c.
//
// Solidity: function unbondingCompleteBlock(address account) view returns(uint256)
func (_Atlas *AtlasCallerSession) UnbondingCompleteBlock(account common.Address) (*big.Int, error) {
	return _Atlas.Contract.UnbondingCompleteBlock(&_Atlas.CallOpts, account)
}

// Withdrawals is a free data retrieval call binding the contract method 0xade0e93e.
//
// Solidity: function withdrawals() view returns(uint256)
func (_Atlas *AtlasCaller) Withdrawals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "withdrawals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Withdrawals is a free data retrieval call binding the contract method 0xade0e93e.
//
// Solidity: function withdrawals() view returns(uint256)
func (_Atlas *AtlasSession) Withdrawals() (*big.Int, error) {
	return _Atlas.Contract.Withdrawals(&_Atlas.CallOpts)
}

// Withdrawals is a free data retrieval call binding the contract method 0xade0e93e.
//
// Solidity: function withdrawals() view returns(uint256)
func (_Atlas *AtlasCallerSession) Withdrawals() (*big.Int, error) {
	return _Atlas.Contract.Withdrawals(&_Atlas.CallOpts)
}

// Writeoffs is a free data retrieval call binding the contract method 0x78448bcc.
//
// Solidity: function writeoffs() view returns(uint256)
func (_Atlas *AtlasCaller) Writeoffs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "writeoffs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Writeoffs is a free data retrieval call binding the contract method 0x78448bcc.
//
// Solidity: function writeoffs() view returns(uint256)
func (_Atlas *AtlasSession) Writeoffs() (*big.Int, error) {
	return _Atlas.Contract.Writeoffs(&_Atlas.CallOpts)
}

// Writeoffs is a free data retrieval call binding the contract method 0x78448bcc.
//
// Solidity: function writeoffs() view returns(uint256)
func (_Atlas *AtlasCallerSession) Writeoffs() (*big.Int, error) {
	return _Atlas.Contract.Writeoffs(&_Atlas.CallOpts)
}

// BecomeSurchargeRecipient is a paid mutator transaction binding the contract method 0x8ebf091f.
//
// Solidity: function becomeSurchargeRecipient() returns()
func (_Atlas *AtlasTransactor) BecomeSurchargeRecipient(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "becomeSurchargeRecipient")
}

// BecomeSurchargeRecipient is a paid mutator transaction binding the contract method 0x8ebf091f.
//
// Solidity: function becomeSurchargeRecipient() returns()
func (_Atlas *AtlasSession) BecomeSurchargeRecipient() (*types.Transaction, error) {
	return _Atlas.Contract.BecomeSurchargeRecipient(&_Atlas.TransactOpts)
}

// BecomeSurchargeRecipient is a paid mutator transaction binding the contract method 0x8ebf091f.
//
// Solidity: function becomeSurchargeRecipient() returns()
func (_Atlas *AtlasTransactorSession) BecomeSurchargeRecipient() (*types.Transaction, error) {
	return _Atlas.Contract.BecomeSurchargeRecipient(&_Atlas.TransactOpts)
}

// Bond is a paid mutator transaction binding the contract method 0x9940686e.
//
// Solidity: function bond(uint256 amount) returns()
func (_Atlas *AtlasTransactor) Bond(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "bond", amount)
}

// Bond is a paid mutator transaction binding the contract method 0x9940686e.
//
// Solidity: function bond(uint256 amount) returns()
func (_Atlas *AtlasSession) Bond(amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Bond(&_Atlas.TransactOpts, amount)
}

// Bond is a paid mutator transaction binding the contract method 0x9940686e.
//
// Solidity: function bond(uint256 amount) returns()
func (_Atlas *AtlasTransactorSession) Bond(amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Bond(&_Atlas.TransactOpts, amount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 amount) payable returns()
func (_Atlas *AtlasTransactor) Borrow(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "borrow", amount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 amount) payable returns()
func (_Atlas *AtlasSession) Borrow(amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Borrow(&_Atlas.TransactOpts, amount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 amount) payable returns()
func (_Atlas *AtlasTransactorSession) Borrow(amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Borrow(&_Atlas.TransactOpts, amount)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_Atlas *AtlasTransactor) Contribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "contribute")
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_Atlas *AtlasSession) Contribute() (*types.Transaction, error) {
	return _Atlas.Contract.Contribute(&_Atlas.TransactOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_Atlas *AtlasTransactorSession) Contribute() (*types.Transaction, error) {
	return _Atlas.Contract.Contribute(&_Atlas.TransactOpts)
}

// CreateExecutionEnvironment is a paid mutator transaction binding the contract method 0x4626e4a2.
//
// Solidity: function createExecutionEnvironment(address control) returns(address executionEnvironment)
func (_Atlas *AtlasTransactor) CreateExecutionEnvironment(opts *bind.TransactOpts, control common.Address) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "createExecutionEnvironment", control)
}

// CreateExecutionEnvironment is a paid mutator transaction binding the contract method 0x4626e4a2.
//
// Solidity: function createExecutionEnvironment(address control) returns(address executionEnvironment)
func (_Atlas *AtlasSession) CreateExecutionEnvironment(control common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.CreateExecutionEnvironment(&_Atlas.TransactOpts, control)
}

// CreateExecutionEnvironment is a paid mutator transaction binding the contract method 0x4626e4a2.
//
// Solidity: function createExecutionEnvironment(address control) returns(address executionEnvironment)
func (_Atlas *AtlasTransactorSession) CreateExecutionEnvironment(control common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.CreateExecutionEnvironment(&_Atlas.TransactOpts, control)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Atlas *AtlasTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Atlas *AtlasSession) Deposit() (*types.Transaction, error) {
	return _Atlas.Contract.Deposit(&_Atlas.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Atlas *AtlasTransactorSession) Deposit() (*types.Transaction, error) {
	return _Atlas.Contract.Deposit(&_Atlas.TransactOpts)
}

// DepositAndBond is a paid mutator transaction binding the contract method 0xf05f88e0.
//
// Solidity: function depositAndBond(uint256 amountToBond) payable returns()
func (_Atlas *AtlasTransactor) DepositAndBond(opts *bind.TransactOpts, amountToBond *big.Int) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "depositAndBond", amountToBond)
}

// DepositAndBond is a paid mutator transaction binding the contract method 0xf05f88e0.
//
// Solidity: function depositAndBond(uint256 amountToBond) payable returns()
func (_Atlas *AtlasSession) DepositAndBond(amountToBond *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.DepositAndBond(&_Atlas.TransactOpts, amountToBond)
}

// DepositAndBond is a paid mutator transaction binding the contract method 0xf05f88e0.
//
// Solidity: function depositAndBond(uint256 amountToBond) payable returns()
func (_Atlas *AtlasTransactorSession) DepositAndBond(amountToBond *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.DepositAndBond(&_Atlas.TransactOpts, amountToBond)
}

// Execute is a paid mutator transaction binding the contract method 0x476c47f8.
//
// Solidity: function execute((address,uint32,address,uint32) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, address executionEnvironment, address bundler, bytes32 userOpHash) payable returns((bytes32,address,uint24,uint8,uint8,uint8,uint8,bool,bool,bool,bool,address) ctx)
func (_Atlas *AtlasTransactor) Execute(opts *bind.TransactOpts, dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, executionEnvironment common.Address, bundler common.Address, userOpHash [32]byte) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "execute", dConfig, userOp, solverOps, executionEnvironment, bundler, userOpHash)
}

// Execute is a paid mutator transaction binding the contract method 0x476c47f8.
//
// Solidity: function execute((address,uint32,address,uint32) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, address executionEnvironment, address bundler, bytes32 userOpHash) payable returns((bytes32,address,uint24,uint8,uint8,uint8,uint8,bool,bool,bool,bool,address) ctx)
func (_Atlas *AtlasSession) Execute(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, executionEnvironment common.Address, bundler common.Address, userOpHash [32]byte) (*types.Transaction, error) {
	return _Atlas.Contract.Execute(&_Atlas.TransactOpts, dConfig, userOp, solverOps, executionEnvironment, bundler, userOpHash)
}

// Execute is a paid mutator transaction binding the contract method 0x476c47f8.
//
// Solidity: function execute((address,uint32,address,uint32) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, address executionEnvironment, address bundler, bytes32 userOpHash) payable returns((bytes32,address,uint24,uint8,uint8,uint8,uint8,bool,bool,bool,bool,address) ctx)
func (_Atlas *AtlasTransactorSession) Execute(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, executionEnvironment common.Address, bundler common.Address, userOpHash [32]byte) (*types.Transaction, error) {
	return _Atlas.Contract.Execute(&_Atlas.TransactOpts, dConfig, userOp, solverOps, executionEnvironment, bundler, userOpHash)
}

// Metacall is a paid mutator transaction binding the contract method 0x4683d90f.
//
// Solidity: function metacall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool auctionWon)
func (_Atlas *AtlasTransactor) Metacall(opts *bind.TransactOpts, userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "metacall", userOp, solverOps, dAppOp)
}

// Metacall is a paid mutator transaction binding the contract method 0x4683d90f.
//
// Solidity: function metacall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool auctionWon)
func (_Atlas *AtlasSession) Metacall(userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Atlas.Contract.Metacall(&_Atlas.TransactOpts, userOp, solverOps, dAppOp)
}

// Metacall is a paid mutator transaction binding the contract method 0x4683d90f.
//
// Solidity: function metacall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool auctionWon)
func (_Atlas *AtlasTransactorSession) Metacall(userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Atlas.Contract.Metacall(&_Atlas.TransactOpts, userOp, solverOps, dAppOp)
}

// Reconcile is a paid mutator transaction binding the contract method 0xf68b84f7.
//
// Solidity: function reconcile(uint256 maxApprovedGasSpend) payable returns(uint256 owed)
func (_Atlas *AtlasTransactor) Reconcile(opts *bind.TransactOpts, maxApprovedGasSpend *big.Int) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "reconcile", maxApprovedGasSpend)
}

// Reconcile is a paid mutator transaction binding the contract method 0xf68b84f7.
//
// Solidity: function reconcile(uint256 maxApprovedGasSpend) payable returns(uint256 owed)
func (_Atlas *AtlasSession) Reconcile(maxApprovedGasSpend *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Reconcile(&_Atlas.TransactOpts, maxApprovedGasSpend)
}

// Reconcile is a paid mutator transaction binding the contract method 0xf68b84f7.
//
// Solidity: function reconcile(uint256 maxApprovedGasSpend) payable returns(uint256 owed)
func (_Atlas *AtlasTransactorSession) Reconcile(maxApprovedGasSpend *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Reconcile(&_Atlas.TransactOpts, maxApprovedGasSpend)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns()
func (_Atlas *AtlasTransactor) Redeem(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "redeem", amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns()
func (_Atlas *AtlasSession) Redeem(amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Redeem(&_Atlas.TransactOpts, amount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 amount) returns()
func (_Atlas *AtlasTransactorSession) Redeem(amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Redeem(&_Atlas.TransactOpts, amount)
}

// SolverCall is a paid mutator transaction binding the contract method 0x982d4f82.
//
// Solidity: function solverCall((bytes32,address,uint24,uint8,uint8,uint8,uint8,bool,bool,bool,bool,address) ctx, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, uint256 bidAmount, uint256 gasLimit, bytes returnData) payable returns((uint256,uint256,uint256,bool,bool) solverTracker)
func (_Atlas *AtlasTransactor) SolverCall(opts *bind.TransactOpts, ctx Context, solverOp SolverOperation, bidAmount *big.Int, gasLimit *big.Int, returnData []byte) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "solverCall", ctx, solverOp, bidAmount, gasLimit, returnData)
}

// SolverCall is a paid mutator transaction binding the contract method 0x982d4f82.
//
// Solidity: function solverCall((bytes32,address,uint24,uint8,uint8,uint8,uint8,bool,bool,bool,bool,address) ctx, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, uint256 bidAmount, uint256 gasLimit, bytes returnData) payable returns((uint256,uint256,uint256,bool,bool) solverTracker)
func (_Atlas *AtlasSession) SolverCall(ctx Context, solverOp SolverOperation, bidAmount *big.Int, gasLimit *big.Int, returnData []byte) (*types.Transaction, error) {
	return _Atlas.Contract.SolverCall(&_Atlas.TransactOpts, ctx, solverOp, bidAmount, gasLimit, returnData)
}

// SolverCall is a paid mutator transaction binding the contract method 0x982d4f82.
//
// Solidity: function solverCall((bytes32,address,uint24,uint8,uint8,uint8,uint8,bool,bool,bool,bool,address) ctx, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, uint256 bidAmount, uint256 gasLimit, bytes returnData) payable returns((uint256,uint256,uint256,bool,bool) solverTracker)
func (_Atlas *AtlasTransactorSession) SolverCall(ctx Context, solverOp SolverOperation, bidAmount *big.Int, gasLimit *big.Int, returnData []byte) (*types.Transaction, error) {
	return _Atlas.Contract.SolverCall(&_Atlas.TransactOpts, ctx, solverOp, bidAmount, gasLimit, returnData)
}

// TransferDAppERC20 is a paid mutator transaction binding the contract method 0xb2c5c510.
//
// Solidity: function transferDAppERC20(address token, address destination, uint256 amount, address user, address control) returns()
func (_Atlas *AtlasTransactor) TransferDAppERC20(opts *bind.TransactOpts, token common.Address, destination common.Address, amount *big.Int, user common.Address, control common.Address) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "transferDAppERC20", token, destination, amount, user, control)
}

// TransferDAppERC20 is a paid mutator transaction binding the contract method 0xb2c5c510.
//
// Solidity: function transferDAppERC20(address token, address destination, uint256 amount, address user, address control) returns()
func (_Atlas *AtlasSession) TransferDAppERC20(token common.Address, destination common.Address, amount *big.Int, user common.Address, control common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.TransferDAppERC20(&_Atlas.TransactOpts, token, destination, amount, user, control)
}

// TransferDAppERC20 is a paid mutator transaction binding the contract method 0xb2c5c510.
//
// Solidity: function transferDAppERC20(address token, address destination, uint256 amount, address user, address control) returns()
func (_Atlas *AtlasTransactorSession) TransferDAppERC20(token common.Address, destination common.Address, amount *big.Int, user common.Address, control common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.TransferDAppERC20(&_Atlas.TransactOpts, token, destination, amount, user, control)
}

// TransferSurchargeRecipient is a paid mutator transaction binding the contract method 0xa0531b02.
//
// Solidity: function transferSurchargeRecipient(address newRecipient) returns()
func (_Atlas *AtlasTransactor) TransferSurchargeRecipient(opts *bind.TransactOpts, newRecipient common.Address) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "transferSurchargeRecipient", newRecipient)
}

// TransferSurchargeRecipient is a paid mutator transaction binding the contract method 0xa0531b02.
//
// Solidity: function transferSurchargeRecipient(address newRecipient) returns()
func (_Atlas *AtlasSession) TransferSurchargeRecipient(newRecipient common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.TransferSurchargeRecipient(&_Atlas.TransactOpts, newRecipient)
}

// TransferSurchargeRecipient is a paid mutator transaction binding the contract method 0xa0531b02.
//
// Solidity: function transferSurchargeRecipient(address newRecipient) returns()
func (_Atlas *AtlasTransactorSession) TransferSurchargeRecipient(newRecipient common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.TransferSurchargeRecipient(&_Atlas.TransactOpts, newRecipient)
}

// TransferUserERC20 is a paid mutator transaction binding the contract method 0x234b7ede.
//
// Solidity: function transferUserERC20(address token, address destination, uint256 amount, address user, address control) returns()
func (_Atlas *AtlasTransactor) TransferUserERC20(opts *bind.TransactOpts, token common.Address, destination common.Address, amount *big.Int, user common.Address, control common.Address) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "transferUserERC20", token, destination, amount, user, control)
}

// TransferUserERC20 is a paid mutator transaction binding the contract method 0x234b7ede.
//
// Solidity: function transferUserERC20(address token, address destination, uint256 amount, address user, address control) returns()
func (_Atlas *AtlasSession) TransferUserERC20(token common.Address, destination common.Address, amount *big.Int, user common.Address, control common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.TransferUserERC20(&_Atlas.TransactOpts, token, destination, amount, user, control)
}

// TransferUserERC20 is a paid mutator transaction binding the contract method 0x234b7ede.
//
// Solidity: function transferUserERC20(address token, address destination, uint256 amount, address user, address control) returns()
func (_Atlas *AtlasTransactorSession) TransferUserERC20(token common.Address, destination common.Address, amount *big.Int, user common.Address, control common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.TransferUserERC20(&_Atlas.TransactOpts, token, destination, amount, user, control)
}

// Unbond is a paid mutator transaction binding the contract method 0x27de9e32.
//
// Solidity: function unbond(uint256 amount) returns()
func (_Atlas *AtlasTransactor) Unbond(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "unbond", amount)
}

// Unbond is a paid mutator transaction binding the contract method 0x27de9e32.
//
// Solidity: function unbond(uint256 amount) returns()
func (_Atlas *AtlasSession) Unbond(amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Unbond(&_Atlas.TransactOpts, amount)
}

// Unbond is a paid mutator transaction binding the contract method 0x27de9e32.
//
// Solidity: function unbond(uint256 amount) returns()
func (_Atlas *AtlasTransactorSession) Unbond(amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Unbond(&_Atlas.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Atlas *AtlasTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Atlas *AtlasSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Withdraw(&_Atlas.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Atlas *AtlasTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Withdraw(&_Atlas.TransactOpts, amount)
}

// WithdrawSurcharge is a paid mutator transaction binding the contract method 0xc41d54da.
//
// Solidity: function withdrawSurcharge() returns()
func (_Atlas *AtlasTransactor) WithdrawSurcharge(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "withdrawSurcharge")
}

// WithdrawSurcharge is a paid mutator transaction binding the contract method 0xc41d54da.
//
// Solidity: function withdrawSurcharge() returns()
func (_Atlas *AtlasSession) WithdrawSurcharge() (*types.Transaction, error) {
	return _Atlas.Contract.WithdrawSurcharge(&_Atlas.TransactOpts)
}

// WithdrawSurcharge is a paid mutator transaction binding the contract method 0xc41d54da.
//
// Solidity: function withdrawSurcharge() returns()
func (_Atlas *AtlasTransactorSession) WithdrawSurcharge() (*types.Transaction, error) {
	return _Atlas.Contract.WithdrawSurcharge(&_Atlas.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Atlas *AtlasTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Atlas.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Atlas *AtlasSession) Receive() (*types.Transaction, error) {
	return _Atlas.Contract.Receive(&_Atlas.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Atlas *AtlasTransactorSession) Receive() (*types.Transaction, error) {
	return _Atlas.Contract.Receive(&_Atlas.TransactOpts)
}

// AtlasApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Atlas contract.
type AtlasApprovalIterator struct {
	Event *AtlasApproval // Event containing the contract specifics and raw log

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
func (it *AtlasApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasApproval)
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
		it.Event = new(AtlasApproval)
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
func (it *AtlasApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasApproval represents a Approval event raised by the Atlas contract.
type AtlasApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Atlas *AtlasFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AtlasApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AtlasApprovalIterator{contract: _Atlas.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Atlas *AtlasFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AtlasApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasApproval)
				if err := _Atlas.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseApproval(log types.Log) (*AtlasApproval, error) {
	event := new(AtlasApproval)
	if err := _Atlas.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasBondIterator is returned from FilterBond and is used to iterate over the raw logs and unpacked data for Bond events raised by the Atlas contract.
type AtlasBondIterator struct {
	Event *AtlasBond // Event containing the contract specifics and raw log

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
func (it *AtlasBondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasBond)
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
		it.Event = new(AtlasBond)
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
func (it *AtlasBondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasBondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasBond represents a Bond event raised by the Atlas contract.
type AtlasBond struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBond is a free log retrieval operation binding the contract event 0x6b1d99469ed62a423d7e402bfa68a467261ca2229127c55045ee41e5d9a0f21d.
//
// Solidity: event Bond(address indexed owner, uint256 amount)
func (_Atlas *AtlasFilterer) FilterBond(opts *bind.FilterOpts, owner []common.Address) (*AtlasBondIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "Bond", ownerRule)
	if err != nil {
		return nil, err
	}
	return &AtlasBondIterator{contract: _Atlas.contract, event: "Bond", logs: logs, sub: sub}, nil
}

// WatchBond is a free log subscription operation binding the contract event 0x6b1d99469ed62a423d7e402bfa68a467261ca2229127c55045ee41e5d9a0f21d.
//
// Solidity: event Bond(address indexed owner, uint256 amount)
func (_Atlas *AtlasFilterer) WatchBond(opts *bind.WatchOpts, sink chan<- *AtlasBond, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "Bond", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasBond)
				if err := _Atlas.contract.UnpackLog(event, "Bond", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseBond(log types.Log) (*AtlasBond, error) {
	event := new(AtlasBond)
	if err := _Atlas.contract.UnpackLog(event, "Bond", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasDAppDisabledIterator is returned from FilterDAppDisabled and is used to iterate over the raw logs and unpacked data for DAppDisabled events raised by the Atlas contract.
type AtlasDAppDisabledIterator struct {
	Event *AtlasDAppDisabled // Event containing the contract specifics and raw log

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
func (it *AtlasDAppDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasDAppDisabled)
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
		it.Event = new(AtlasDAppDisabled)
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
func (it *AtlasDAppDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasDAppDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasDAppDisabled represents a DAppDisabled event raised by the Atlas contract.
type AtlasDAppDisabled struct {
	Control    common.Address
	Governance common.Address
	CallConfig uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDAppDisabled is a free log retrieval operation binding the contract event 0xfb402de0284ae6437f381fcd6b8da6639e155a56551b5a8b95d2ab6c4e007d0a.
//
// Solidity: event DAppDisabled(address indexed control, address indexed governance, uint32 callConfig)
func (_Atlas *AtlasFilterer) FilterDAppDisabled(opts *bind.FilterOpts, control []common.Address, governance []common.Address) (*AtlasDAppDisabledIterator, error) {

	var controlRule []interface{}
	for _, controlItem := range control {
		controlRule = append(controlRule, controlItem)
	}
	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "DAppDisabled", controlRule, governanceRule)
	if err != nil {
		return nil, err
	}
	return &AtlasDAppDisabledIterator{contract: _Atlas.contract, event: "DAppDisabled", logs: logs, sub: sub}, nil
}

// WatchDAppDisabled is a free log subscription operation binding the contract event 0xfb402de0284ae6437f381fcd6b8da6639e155a56551b5a8b95d2ab6c4e007d0a.
//
// Solidity: event DAppDisabled(address indexed control, address indexed governance, uint32 callConfig)
func (_Atlas *AtlasFilterer) WatchDAppDisabled(opts *bind.WatchOpts, sink chan<- *AtlasDAppDisabled, control []common.Address, governance []common.Address) (event.Subscription, error) {

	var controlRule []interface{}
	for _, controlItem := range control {
		controlRule = append(controlRule, controlItem)
	}
	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "DAppDisabled", controlRule, governanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasDAppDisabled)
				if err := _Atlas.contract.UnpackLog(event, "DAppDisabled", log); err != nil {
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
// Solidity: event DAppDisabled(address indexed control, address indexed governance, uint32 callConfig)
func (_Atlas *AtlasFilterer) ParseDAppDisabled(log types.Log) (*AtlasDAppDisabled, error) {
	event := new(AtlasDAppDisabled)
	if err := _Atlas.contract.UnpackLog(event, "DAppDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasDAppGovernanceChangedIterator is returned from FilterDAppGovernanceChanged and is used to iterate over the raw logs and unpacked data for DAppGovernanceChanged events raised by the Atlas contract.
type AtlasDAppGovernanceChangedIterator struct {
	Event *AtlasDAppGovernanceChanged // Event containing the contract specifics and raw log

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
func (it *AtlasDAppGovernanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasDAppGovernanceChanged)
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
		it.Event = new(AtlasDAppGovernanceChanged)
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
func (it *AtlasDAppGovernanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasDAppGovernanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasDAppGovernanceChanged represents a DAppGovernanceChanged event raised by the Atlas contract.
type AtlasDAppGovernanceChanged struct {
	Control       common.Address
	OldGovernance common.Address
	NewGovernance common.Address
	CallConfig    uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDAppGovernanceChanged is a free log retrieval operation binding the contract event 0xcc02b55b78457369d35dc189a074d3fc4f96a1897b405c0f1c57789f782d8bd2.
//
// Solidity: event DAppGovernanceChanged(address indexed control, address indexed oldGovernance, address indexed newGovernance, uint32 callConfig)
func (_Atlas *AtlasFilterer) FilterDAppGovernanceChanged(opts *bind.FilterOpts, control []common.Address, oldGovernance []common.Address, newGovernance []common.Address) (*AtlasDAppGovernanceChangedIterator, error) {

	var controlRule []interface{}
	for _, controlItem := range control {
		controlRule = append(controlRule, controlItem)
	}
	var oldGovernanceRule []interface{}
	for _, oldGovernanceItem := range oldGovernance {
		oldGovernanceRule = append(oldGovernanceRule, oldGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "DAppGovernanceChanged", controlRule, oldGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &AtlasDAppGovernanceChangedIterator{contract: _Atlas.contract, event: "DAppGovernanceChanged", logs: logs, sub: sub}, nil
}

// WatchDAppGovernanceChanged is a free log subscription operation binding the contract event 0xcc02b55b78457369d35dc189a074d3fc4f96a1897b405c0f1c57789f782d8bd2.
//
// Solidity: event DAppGovernanceChanged(address indexed control, address indexed oldGovernance, address indexed newGovernance, uint32 callConfig)
func (_Atlas *AtlasFilterer) WatchDAppGovernanceChanged(opts *bind.WatchOpts, sink chan<- *AtlasDAppGovernanceChanged, control []common.Address, oldGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var controlRule []interface{}
	for _, controlItem := range control {
		controlRule = append(controlRule, controlItem)
	}
	var oldGovernanceRule []interface{}
	for _, oldGovernanceItem := range oldGovernance {
		oldGovernanceRule = append(oldGovernanceRule, oldGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "DAppGovernanceChanged", controlRule, oldGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasDAppGovernanceChanged)
				if err := _Atlas.contract.UnpackLog(event, "DAppGovernanceChanged", log); err != nil {
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
// Solidity: event DAppGovernanceChanged(address indexed control, address indexed oldGovernance, address indexed newGovernance, uint32 callConfig)
func (_Atlas *AtlasFilterer) ParseDAppGovernanceChanged(log types.Log) (*AtlasDAppGovernanceChanged, error) {
	event := new(AtlasDAppGovernanceChanged)
	if err := _Atlas.contract.UnpackLog(event, "DAppGovernanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasExecutionEnvironmentCreatedIterator is returned from FilterExecutionEnvironmentCreated and is used to iterate over the raw logs and unpacked data for ExecutionEnvironmentCreated events raised by the Atlas contract.
type AtlasExecutionEnvironmentCreatedIterator struct {
	Event *AtlasExecutionEnvironmentCreated // Event containing the contract specifics and raw log

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
func (it *AtlasExecutionEnvironmentCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasExecutionEnvironmentCreated)
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
		it.Event = new(AtlasExecutionEnvironmentCreated)
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
func (it *AtlasExecutionEnvironmentCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasExecutionEnvironmentCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasExecutionEnvironmentCreated represents a ExecutionEnvironmentCreated event raised by the Atlas contract.
type AtlasExecutionEnvironmentCreated struct {
	User                 common.Address
	ExecutionEnvironment common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterExecutionEnvironmentCreated is a free log retrieval operation binding the contract event 0x6ed96358b086d2aca68c2e2e4dc23fb45421ac513a7fc3127e34569833b3c646.
//
// Solidity: event ExecutionEnvironmentCreated(address indexed user, address indexed executionEnvironment)
func (_Atlas *AtlasFilterer) FilterExecutionEnvironmentCreated(opts *bind.FilterOpts, user []common.Address, executionEnvironment []common.Address) (*AtlasExecutionEnvironmentCreatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var executionEnvironmentRule []interface{}
	for _, executionEnvironmentItem := range executionEnvironment {
		executionEnvironmentRule = append(executionEnvironmentRule, executionEnvironmentItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "ExecutionEnvironmentCreated", userRule, executionEnvironmentRule)
	if err != nil {
		return nil, err
	}
	return &AtlasExecutionEnvironmentCreatedIterator{contract: _Atlas.contract, event: "ExecutionEnvironmentCreated", logs: logs, sub: sub}, nil
}

// WatchExecutionEnvironmentCreated is a free log subscription operation binding the contract event 0x6ed96358b086d2aca68c2e2e4dc23fb45421ac513a7fc3127e34569833b3c646.
//
// Solidity: event ExecutionEnvironmentCreated(address indexed user, address indexed executionEnvironment)
func (_Atlas *AtlasFilterer) WatchExecutionEnvironmentCreated(opts *bind.WatchOpts, sink chan<- *AtlasExecutionEnvironmentCreated, user []common.Address, executionEnvironment []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var executionEnvironmentRule []interface{}
	for _, executionEnvironmentItem := range executionEnvironment {
		executionEnvironmentRule = append(executionEnvironmentRule, executionEnvironmentItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "ExecutionEnvironmentCreated", userRule, executionEnvironmentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasExecutionEnvironmentCreated)
				if err := _Atlas.contract.UnpackLog(event, "ExecutionEnvironmentCreated", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseExecutionEnvironmentCreated(log types.Log) (*AtlasExecutionEnvironmentCreated, error) {
	event := new(AtlasExecutionEnvironmentCreated)
	if err := _Atlas.contract.UnpackLog(event, "ExecutionEnvironmentCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasGovernanceTransferStartedIterator is returned from FilterGovernanceTransferStarted and is used to iterate over the raw logs and unpacked data for GovernanceTransferStarted events raised by the Atlas contract.
type AtlasGovernanceTransferStartedIterator struct {
	Event *AtlasGovernanceTransferStarted // Event containing the contract specifics and raw log

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
func (it *AtlasGovernanceTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasGovernanceTransferStarted)
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
		it.Event = new(AtlasGovernanceTransferStarted)
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
func (it *AtlasGovernanceTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasGovernanceTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasGovernanceTransferStarted represents a GovernanceTransferStarted event raised by the Atlas contract.
type AtlasGovernanceTransferStarted struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferStarted is a free log retrieval operation binding the contract event 0x1c4bb4e3cab7b72da7eb9f0ae62554dda85dc7fb907c946ad2776095b95ac1ad.
//
// Solidity: event GovernanceTransferStarted(address indexed previousGovernance, address indexed newGovernance)
func (_Atlas *AtlasFilterer) FilterGovernanceTransferStarted(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*AtlasGovernanceTransferStartedIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "GovernanceTransferStarted", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &AtlasGovernanceTransferStartedIterator{contract: _Atlas.contract, event: "GovernanceTransferStarted", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferStarted is a free log subscription operation binding the contract event 0x1c4bb4e3cab7b72da7eb9f0ae62554dda85dc7fb907c946ad2776095b95ac1ad.
//
// Solidity: event GovernanceTransferStarted(address indexed previousGovernance, address indexed newGovernance)
func (_Atlas *AtlasFilterer) WatchGovernanceTransferStarted(opts *bind.WatchOpts, sink chan<- *AtlasGovernanceTransferStarted, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "GovernanceTransferStarted", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasGovernanceTransferStarted)
				if err := _Atlas.contract.UnpackLog(event, "GovernanceTransferStarted", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseGovernanceTransferStarted(log types.Log) (*AtlasGovernanceTransferStarted, error) {
	event := new(AtlasGovernanceTransferStarted)
	if err := _Atlas.contract.UnpackLog(event, "GovernanceTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasGovernanceTransferredIterator is returned from FilterGovernanceTransferred and is used to iterate over the raw logs and unpacked data for GovernanceTransferred events raised by the Atlas contract.
type AtlasGovernanceTransferredIterator struct {
	Event *AtlasGovernanceTransferred // Event containing the contract specifics and raw log

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
func (it *AtlasGovernanceTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasGovernanceTransferred)
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
		it.Event = new(AtlasGovernanceTransferred)
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
func (it *AtlasGovernanceTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasGovernanceTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasGovernanceTransferred represents a GovernanceTransferred event raised by the Atlas contract.
type AtlasGovernanceTransferred struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferred is a free log retrieval operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_Atlas *AtlasFilterer) FilterGovernanceTransferred(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*AtlasGovernanceTransferredIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &AtlasGovernanceTransferredIterator{contract: _Atlas.contract, event: "GovernanceTransferred", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferred is a free log subscription operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_Atlas *AtlasFilterer) WatchGovernanceTransferred(opts *bind.WatchOpts, sink chan<- *AtlasGovernanceTransferred, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasGovernanceTransferred)
				if err := _Atlas.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseGovernanceTransferred(log types.Log) (*AtlasGovernanceTransferred, error) {
	event := new(AtlasGovernanceTransferred)
	if err := _Atlas.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasMetacallResultIterator is returned from FilterMetacallResult and is used to iterate over the raw logs and unpacked data for MetacallResult events raised by the Atlas contract.
type AtlasMetacallResultIterator struct {
	Event *AtlasMetacallResult // Event containing the contract specifics and raw log

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
func (it *AtlasMetacallResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasMetacallResult)
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
		it.Event = new(AtlasMetacallResult)
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
func (it *AtlasMetacallResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasMetacallResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasMetacallResult represents a MetacallResult event raised by the Atlas contract.
type AtlasMetacallResult struct {
	Bundler                common.Address
	User                   common.Address
	SolverSuccessful       bool
	DisbursementSuccessful bool
	EthPaidToBundler       *big.Int
	NetGasSurcharge        *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterMetacallResult is a free log retrieval operation binding the contract event 0xde0b67c553d23d5f8bb4c6305ceb7112ee27f974197bee8c02502a7e36cff063.
//
// Solidity: event MetacallResult(address indexed bundler, address indexed user, bool solverSuccessful, bool disbursementSuccessful, uint256 ethPaidToBundler, uint256 netGasSurcharge)
func (_Atlas *AtlasFilterer) FilterMetacallResult(opts *bind.FilterOpts, bundler []common.Address, user []common.Address) (*AtlasMetacallResultIterator, error) {

	var bundlerRule []interface{}
	for _, bundlerItem := range bundler {
		bundlerRule = append(bundlerRule, bundlerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "MetacallResult", bundlerRule, userRule)
	if err != nil {
		return nil, err
	}
	return &AtlasMetacallResultIterator{contract: _Atlas.contract, event: "MetacallResult", logs: logs, sub: sub}, nil
}

// WatchMetacallResult is a free log subscription operation binding the contract event 0xde0b67c553d23d5f8bb4c6305ceb7112ee27f974197bee8c02502a7e36cff063.
//
// Solidity: event MetacallResult(address indexed bundler, address indexed user, bool solverSuccessful, bool disbursementSuccessful, uint256 ethPaidToBundler, uint256 netGasSurcharge)
func (_Atlas *AtlasFilterer) WatchMetacallResult(opts *bind.WatchOpts, sink chan<- *AtlasMetacallResult, bundler []common.Address, user []common.Address) (event.Subscription, error) {

	var bundlerRule []interface{}
	for _, bundlerItem := range bundler {
		bundlerRule = append(bundlerRule, bundlerItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "MetacallResult", bundlerRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasMetacallResult)
				if err := _Atlas.contract.UnpackLog(event, "MetacallResult", log); err != nil {
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

// ParseMetacallResult is a log parse operation binding the contract event 0xde0b67c553d23d5f8bb4c6305ceb7112ee27f974197bee8c02502a7e36cff063.
//
// Solidity: event MetacallResult(address indexed bundler, address indexed user, bool solverSuccessful, bool disbursementSuccessful, uint256 ethPaidToBundler, uint256 netGasSurcharge)
func (_Atlas *AtlasFilterer) ParseMetacallResult(log types.Log) (*AtlasMetacallResult, error) {
	event := new(AtlasMetacallResult)
	if err := _Atlas.contract.UnpackLog(event, "MetacallResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasNewDAppSignatoryIterator is returned from FilterNewDAppSignatory and is used to iterate over the raw logs and unpacked data for NewDAppSignatory events raised by the Atlas contract.
type AtlasNewDAppSignatoryIterator struct {
	Event *AtlasNewDAppSignatory // Event containing the contract specifics and raw log

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
func (it *AtlasNewDAppSignatoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasNewDAppSignatory)
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
		it.Event = new(AtlasNewDAppSignatory)
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
func (it *AtlasNewDAppSignatoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasNewDAppSignatoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasNewDAppSignatory represents a NewDAppSignatory event raised by the Atlas contract.
type AtlasNewDAppSignatory struct {
	Control    common.Address
	Governance common.Address
	Signatory  common.Address
	CallConfig uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewDAppSignatory is a free log retrieval operation binding the contract event 0xcacb2beeeba676b69bddabab0d5f66e2733cabc804f82afd92ea5beae71934b4.
//
// Solidity: event NewDAppSignatory(address indexed control, address indexed governance, address indexed signatory, uint32 callConfig)
func (_Atlas *AtlasFilterer) FilterNewDAppSignatory(opts *bind.FilterOpts, control []common.Address, governance []common.Address, signatory []common.Address) (*AtlasNewDAppSignatoryIterator, error) {

	var controlRule []interface{}
	for _, controlItem := range control {
		controlRule = append(controlRule, controlItem)
	}
	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}
	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "NewDAppSignatory", controlRule, governanceRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return &AtlasNewDAppSignatoryIterator{contract: _Atlas.contract, event: "NewDAppSignatory", logs: logs, sub: sub}, nil
}

// WatchNewDAppSignatory is a free log subscription operation binding the contract event 0xcacb2beeeba676b69bddabab0d5f66e2733cabc804f82afd92ea5beae71934b4.
//
// Solidity: event NewDAppSignatory(address indexed control, address indexed governance, address indexed signatory, uint32 callConfig)
func (_Atlas *AtlasFilterer) WatchNewDAppSignatory(opts *bind.WatchOpts, sink chan<- *AtlasNewDAppSignatory, control []common.Address, governance []common.Address, signatory []common.Address) (event.Subscription, error) {

	var controlRule []interface{}
	for _, controlItem := range control {
		controlRule = append(controlRule, controlItem)
	}
	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}
	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "NewDAppSignatory", controlRule, governanceRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasNewDAppSignatory)
				if err := _Atlas.contract.UnpackLog(event, "NewDAppSignatory", log); err != nil {
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
// Solidity: event NewDAppSignatory(address indexed control, address indexed governance, address indexed signatory, uint32 callConfig)
func (_Atlas *AtlasFilterer) ParseNewDAppSignatory(log types.Log) (*AtlasNewDAppSignatory, error) {
	event := new(AtlasNewDAppSignatory)
	if err := _Atlas.contract.UnpackLog(event, "NewDAppSignatory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Atlas contract.
type AtlasRedeemIterator struct {
	Event *AtlasRedeem // Event containing the contract specifics and raw log

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
func (it *AtlasRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasRedeem)
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
		it.Event = new(AtlasRedeem)
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
func (it *AtlasRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasRedeem represents a Redeem event raised by the Atlas contract.
type AtlasRedeem struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0x222838db2794d11532d940e8dec38ae307ed0b63cd97c233322e221f998767a6.
//
// Solidity: event Redeem(address indexed owner, uint256 amount)
func (_Atlas *AtlasFilterer) FilterRedeem(opts *bind.FilterOpts, owner []common.Address) (*AtlasRedeemIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "Redeem", ownerRule)
	if err != nil {
		return nil, err
	}
	return &AtlasRedeemIterator{contract: _Atlas.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0x222838db2794d11532d940e8dec38ae307ed0b63cd97c233322e221f998767a6.
//
// Solidity: event Redeem(address indexed owner, uint256 amount)
func (_Atlas *AtlasFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *AtlasRedeem, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "Redeem", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasRedeem)
				if err := _Atlas.contract.UnpackLog(event, "Redeem", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseRedeem(log types.Log) (*AtlasRedeem, error) {
	event := new(AtlasRedeem)
	if err := _Atlas.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasRemovedDAppSignatoryIterator is returned from FilterRemovedDAppSignatory and is used to iterate over the raw logs and unpacked data for RemovedDAppSignatory events raised by the Atlas contract.
type AtlasRemovedDAppSignatoryIterator struct {
	Event *AtlasRemovedDAppSignatory // Event containing the contract specifics and raw log

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
func (it *AtlasRemovedDAppSignatoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasRemovedDAppSignatory)
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
		it.Event = new(AtlasRemovedDAppSignatory)
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
func (it *AtlasRemovedDAppSignatoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasRemovedDAppSignatoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasRemovedDAppSignatory represents a RemovedDAppSignatory event raised by the Atlas contract.
type AtlasRemovedDAppSignatory struct {
	Control    common.Address
	Governance common.Address
	Signatory  common.Address
	CallConfig uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemovedDAppSignatory is a free log retrieval operation binding the contract event 0x90a400fe7fe676a24908d5f730fe6d38072340bc8f74e18a543132d840bfe5ce.
//
// Solidity: event RemovedDAppSignatory(address indexed control, address indexed governance, address indexed signatory, uint32 callConfig)
func (_Atlas *AtlasFilterer) FilterRemovedDAppSignatory(opts *bind.FilterOpts, control []common.Address, governance []common.Address, signatory []common.Address) (*AtlasRemovedDAppSignatoryIterator, error) {

	var controlRule []interface{}
	for _, controlItem := range control {
		controlRule = append(controlRule, controlItem)
	}
	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}
	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "RemovedDAppSignatory", controlRule, governanceRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return &AtlasRemovedDAppSignatoryIterator{contract: _Atlas.contract, event: "RemovedDAppSignatory", logs: logs, sub: sub}, nil
}

// WatchRemovedDAppSignatory is a free log subscription operation binding the contract event 0x90a400fe7fe676a24908d5f730fe6d38072340bc8f74e18a543132d840bfe5ce.
//
// Solidity: event RemovedDAppSignatory(address indexed control, address indexed governance, address indexed signatory, uint32 callConfig)
func (_Atlas *AtlasFilterer) WatchRemovedDAppSignatory(opts *bind.WatchOpts, sink chan<- *AtlasRemovedDAppSignatory, control []common.Address, governance []common.Address, signatory []common.Address) (event.Subscription, error) {

	var controlRule []interface{}
	for _, controlItem := range control {
		controlRule = append(controlRule, controlItem)
	}
	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}
	var signatoryRule []interface{}
	for _, signatoryItem := range signatory {
		signatoryRule = append(signatoryRule, signatoryItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "RemovedDAppSignatory", controlRule, governanceRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasRemovedDAppSignatory)
				if err := _Atlas.contract.UnpackLog(event, "RemovedDAppSignatory", log); err != nil {
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
// Solidity: event RemovedDAppSignatory(address indexed control, address indexed governance, address indexed signatory, uint32 callConfig)
func (_Atlas *AtlasFilterer) ParseRemovedDAppSignatory(log types.Log) (*AtlasRemovedDAppSignatory, error) {
	event := new(AtlasRemovedDAppSignatory)
	if err := _Atlas.contract.UnpackLog(event, "RemovedDAppSignatory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasSolverTxResultIterator is returned from FilterSolverTxResult and is used to iterate over the raw logs and unpacked data for SolverTxResult events raised by the Atlas contract.
type AtlasSolverTxResultIterator struct {
	Event *AtlasSolverTxResult // Event containing the contract specifics and raw log

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
func (it *AtlasSolverTxResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasSolverTxResult)
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
		it.Event = new(AtlasSolverTxResult)
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
func (it *AtlasSolverTxResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasSolverTxResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasSolverTxResult represents a SolverTxResult event raised by the Atlas contract.
type AtlasSolverTxResult struct {
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
func (_Atlas *AtlasFilterer) FilterSolverTxResult(opts *bind.FilterOpts, solverTo []common.Address, solverFrom []common.Address) (*AtlasSolverTxResultIterator, error) {

	var solverToRule []interface{}
	for _, solverToItem := range solverTo {
		solverToRule = append(solverToRule, solverToItem)
	}
	var solverFromRule []interface{}
	for _, solverFromItem := range solverFrom {
		solverFromRule = append(solverFromRule, solverFromItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "SolverTxResult", solverToRule, solverFromRule)
	if err != nil {
		return nil, err
	}
	return &AtlasSolverTxResultIterator{contract: _Atlas.contract, event: "SolverTxResult", logs: logs, sub: sub}, nil
}

// WatchSolverTxResult is a free log subscription operation binding the contract event 0xcb73ceef18491eb9eafd886c9b1e14b4cce79f4e50eb1168717912ff6174ef8a.
//
// Solidity: event SolverTxResult(address indexed solverTo, address indexed solverFrom, bool executed, bool success, uint256 result)
func (_Atlas *AtlasFilterer) WatchSolverTxResult(opts *bind.WatchOpts, sink chan<- *AtlasSolverTxResult, solverTo []common.Address, solverFrom []common.Address) (event.Subscription, error) {

	var solverToRule []interface{}
	for _, solverToItem := range solverTo {
		solverToRule = append(solverToRule, solverToItem)
	}
	var solverFromRule []interface{}
	for _, solverFromItem := range solverFrom {
		solverFromRule = append(solverFromRule, solverFromItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "SolverTxResult", solverToRule, solverFromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasSolverTxResult)
				if err := _Atlas.contract.UnpackLog(event, "SolverTxResult", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseSolverTxResult(log types.Log) (*AtlasSolverTxResult, error) {
	event := new(AtlasSolverTxResult)
	if err := _Atlas.contract.UnpackLog(event, "SolverTxResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasSurchargeRecipientTransferStartedIterator is returned from FilterSurchargeRecipientTransferStarted and is used to iterate over the raw logs and unpacked data for SurchargeRecipientTransferStarted events raised by the Atlas contract.
type AtlasSurchargeRecipientTransferStartedIterator struct {
	Event *AtlasSurchargeRecipientTransferStarted // Event containing the contract specifics and raw log

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
func (it *AtlasSurchargeRecipientTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasSurchargeRecipientTransferStarted)
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
		it.Event = new(AtlasSurchargeRecipientTransferStarted)
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
func (it *AtlasSurchargeRecipientTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasSurchargeRecipientTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasSurchargeRecipientTransferStarted represents a SurchargeRecipientTransferStarted event raised by the Atlas contract.
type AtlasSurchargeRecipientTransferStarted struct {
	CurrentRecipient common.Address
	NewRecipient     common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSurchargeRecipientTransferStarted is a free log retrieval operation binding the contract event 0xfc722bcd56a71612cd14b287fbf50545429e6c9e8de86ea7c3febdecd34882fd.
//
// Solidity: event SurchargeRecipientTransferStarted(address currentRecipient, address newRecipient)
func (_Atlas *AtlasFilterer) FilterSurchargeRecipientTransferStarted(opts *bind.FilterOpts) (*AtlasSurchargeRecipientTransferStartedIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "SurchargeRecipientTransferStarted")
	if err != nil {
		return nil, err
	}
	return &AtlasSurchargeRecipientTransferStartedIterator{contract: _Atlas.contract, event: "SurchargeRecipientTransferStarted", logs: logs, sub: sub}, nil
}

// WatchSurchargeRecipientTransferStarted is a free log subscription operation binding the contract event 0xfc722bcd56a71612cd14b287fbf50545429e6c9e8de86ea7c3febdecd34882fd.
//
// Solidity: event SurchargeRecipientTransferStarted(address currentRecipient, address newRecipient)
func (_Atlas *AtlasFilterer) WatchSurchargeRecipientTransferStarted(opts *bind.WatchOpts, sink chan<- *AtlasSurchargeRecipientTransferStarted) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "SurchargeRecipientTransferStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasSurchargeRecipientTransferStarted)
				if err := _Atlas.contract.UnpackLog(event, "SurchargeRecipientTransferStarted", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseSurchargeRecipientTransferStarted(log types.Log) (*AtlasSurchargeRecipientTransferStarted, error) {
	event := new(AtlasSurchargeRecipientTransferStarted)
	if err := _Atlas.contract.UnpackLog(event, "SurchargeRecipientTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasSurchargeRecipientTransferredIterator is returned from FilterSurchargeRecipientTransferred and is used to iterate over the raw logs and unpacked data for SurchargeRecipientTransferred events raised by the Atlas contract.
type AtlasSurchargeRecipientTransferredIterator struct {
	Event *AtlasSurchargeRecipientTransferred // Event containing the contract specifics and raw log

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
func (it *AtlasSurchargeRecipientTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasSurchargeRecipientTransferred)
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
		it.Event = new(AtlasSurchargeRecipientTransferred)
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
func (it *AtlasSurchargeRecipientTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasSurchargeRecipientTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasSurchargeRecipientTransferred represents a SurchargeRecipientTransferred event raised by the Atlas contract.
type AtlasSurchargeRecipientTransferred struct {
	NewRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSurchargeRecipientTransferred is a free log retrieval operation binding the contract event 0x53960c2e64e72b2c1326635f0c002d5cf63e7844d12ed107404693fedde43985.
//
// Solidity: event SurchargeRecipientTransferred(address newRecipient)
func (_Atlas *AtlasFilterer) FilterSurchargeRecipientTransferred(opts *bind.FilterOpts) (*AtlasSurchargeRecipientTransferredIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "SurchargeRecipientTransferred")
	if err != nil {
		return nil, err
	}
	return &AtlasSurchargeRecipientTransferredIterator{contract: _Atlas.contract, event: "SurchargeRecipientTransferred", logs: logs, sub: sub}, nil
}

// WatchSurchargeRecipientTransferred is a free log subscription operation binding the contract event 0x53960c2e64e72b2c1326635f0c002d5cf63e7844d12ed107404693fedde43985.
//
// Solidity: event SurchargeRecipientTransferred(address newRecipient)
func (_Atlas *AtlasFilterer) WatchSurchargeRecipientTransferred(opts *bind.WatchOpts, sink chan<- *AtlasSurchargeRecipientTransferred) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "SurchargeRecipientTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasSurchargeRecipientTransferred)
				if err := _Atlas.contract.UnpackLog(event, "SurchargeRecipientTransferred", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseSurchargeRecipientTransferred(log types.Log) (*AtlasSurchargeRecipientTransferred, error) {
	event := new(AtlasSurchargeRecipientTransferred)
	if err := _Atlas.contract.UnpackLog(event, "SurchargeRecipientTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasSurchargeWithdrawnIterator is returned from FilterSurchargeWithdrawn and is used to iterate over the raw logs and unpacked data for SurchargeWithdrawn events raised by the Atlas contract.
type AtlasSurchargeWithdrawnIterator struct {
	Event *AtlasSurchargeWithdrawn // Event containing the contract specifics and raw log

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
func (it *AtlasSurchargeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasSurchargeWithdrawn)
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
		it.Event = new(AtlasSurchargeWithdrawn)
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
func (it *AtlasSurchargeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasSurchargeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasSurchargeWithdrawn represents a SurchargeWithdrawn event raised by the Atlas contract.
type AtlasSurchargeWithdrawn struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSurchargeWithdrawn is a free log retrieval operation binding the contract event 0x87fa2ce024d3bdae31517696c13905fc0882bc1b4f4508060eb29358056de22b.
//
// Solidity: event SurchargeWithdrawn(address to, uint256 amount)
func (_Atlas *AtlasFilterer) FilterSurchargeWithdrawn(opts *bind.FilterOpts) (*AtlasSurchargeWithdrawnIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "SurchargeWithdrawn")
	if err != nil {
		return nil, err
	}
	return &AtlasSurchargeWithdrawnIterator{contract: _Atlas.contract, event: "SurchargeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchSurchargeWithdrawn is a free log subscription operation binding the contract event 0x87fa2ce024d3bdae31517696c13905fc0882bc1b4f4508060eb29358056de22b.
//
// Solidity: event SurchargeWithdrawn(address to, uint256 amount)
func (_Atlas *AtlasFilterer) WatchSurchargeWithdrawn(opts *bind.WatchOpts, sink chan<- *AtlasSurchargeWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "SurchargeWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasSurchargeWithdrawn)
				if err := _Atlas.contract.UnpackLog(event, "SurchargeWithdrawn", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseSurchargeWithdrawn(log types.Log) (*AtlasSurchargeWithdrawn, error) {
	event := new(AtlasSurchargeWithdrawn)
	if err := _Atlas.contract.UnpackLog(event, "SurchargeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Atlas contract.
type AtlasTransferIterator struct {
	Event *AtlasTransfer // Event containing the contract specifics and raw log

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
func (it *AtlasTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasTransfer)
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
		it.Event = new(AtlasTransfer)
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
func (it *AtlasTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasTransfer represents a Transfer event raised by the Atlas contract.
type AtlasTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Atlas *AtlasFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AtlasTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AtlasTransferIterator{contract: _Atlas.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Atlas *AtlasFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AtlasTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasTransfer)
				if err := _Atlas.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseTransfer(log types.Log) (*AtlasTransfer, error) {
	event := new(AtlasTransfer)
	if err := _Atlas.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasUnbondIterator is returned from FilterUnbond and is used to iterate over the raw logs and unpacked data for Unbond events raised by the Atlas contract.
type AtlasUnbondIterator struct {
	Event *AtlasUnbond // Event containing the contract specifics and raw log

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
func (it *AtlasUnbondIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasUnbond)
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
		it.Event = new(AtlasUnbond)
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
func (it *AtlasUnbondIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasUnbondIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasUnbond represents a Unbond event raised by the Atlas contract.
type AtlasUnbond struct {
	Owner             common.Address
	Amount            *big.Int
	EarliestAvailable *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUnbond is a free log retrieval operation binding the contract event 0x7659747cd8571f1071eea946fdc648adcf181cad916f32a05f82c3a525976048.
//
// Solidity: event Unbond(address indexed owner, uint256 amount, uint256 earliestAvailable)
func (_Atlas *AtlasFilterer) FilterUnbond(opts *bind.FilterOpts, owner []common.Address) (*AtlasUnbondIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "Unbond", ownerRule)
	if err != nil {
		return nil, err
	}
	return &AtlasUnbondIterator{contract: _Atlas.contract, event: "Unbond", logs: logs, sub: sub}, nil
}

// WatchUnbond is a free log subscription operation binding the contract event 0x7659747cd8571f1071eea946fdc648adcf181cad916f32a05f82c3a525976048.
//
// Solidity: event Unbond(address indexed owner, uint256 amount, uint256 earliestAvailable)
func (_Atlas *AtlasFilterer) WatchUnbond(opts *bind.WatchOpts, sink chan<- *AtlasUnbond, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "Unbond", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasUnbond)
				if err := _Atlas.contract.UnpackLog(event, "Unbond", log); err != nil {
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
func (_Atlas *AtlasFilterer) ParseUnbond(log types.Log) (*AtlasUnbond, error) {
	event := new(AtlasUnbond)
	if err := _Atlas.contract.UnpackLog(event, "Unbond", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
