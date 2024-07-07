// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swapIntentDappControl

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

// CallConfig is an auto generated low-level Go binding around an user-defined struct.
type CallConfig struct {
	UserNoncesSequential      bool
	DappNoncesSequential      bool
	RequirePreOps             bool
	TrackPreOpsReturnData     bool
	TrackUserReturnData       bool
	DelegateUser              bool
	PreSolver                 bool
	PostSolver                bool
	RequirePostOps            bool
	ZeroSolvers               bool
	ReuseUserOp               bool
	UserAuctioneer            bool
	SolverAuctioneer          bool
	UnknownAuctioneer         bool
	VerifyCallChainHash       bool
	ForwardReturnData         bool
	RequireFulfillment        bool
	TrustedOpHash             bool
	InvertBidValue            bool
	ExPostBids                bool
	AllowAllocateValueFailure bool
}

// Condition is an auto generated low-level Go binding around an user-defined struct.
type Condition struct {
	Antecedent common.Address
	Context    []byte
}

// DAppConfig is an auto generated low-level Go binding around an user-defined struct.
type DAppConfig struct {
	To             common.Address
	CallConfig     uint32
	BidToken       common.Address
	SolverGasLimit uint32
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

// SwapData is an auto generated low-level Go binding around an user-defined struct.
type SwapData struct {
	TokenUserBuys       common.Address
	AmountUserBuys      *big.Int
	TokenUserSells      common.Address
	AmountUserSells     *big.Int
	AuctionBaseCurrency common.Address
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

// SwapIntentDappControlMetaData contains all meta data concerning the SwapIntentDappControl contract.
var SwapIntentDappControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_atlas\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BothPreOpsAndUserReturnDataCannotBeTracked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BothUserAndDAppNoncesCannotBeSequential\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidControl\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvertBidValueCannotBeExPostBids\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustBeDelegatecalled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoDelegatecall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotImplemented\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAtlas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ATLAS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ATLAS_VERIFICATION\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CALL_CONFIG\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONTROL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_USER_CONDITIONS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SOURCE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"USER_CONDITION_GAS_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"allocateValueCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getBidFormat\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"}],\"name\":\"getBidValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCallConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"userNoncesSequential\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"dappNoncesSequential\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requirePreOps\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"trackPreOpsReturnData\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"trackUserReturnData\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"delegateUser\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"preSolver\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"postSolver\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requirePostOps\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"zeroSolvers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reuseUserOp\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"userAuctioneer\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"solverAuctioneer\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"unknownAuctioneer\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"verifyCallChainHash\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"forwardReturnData\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"requireFulfillment\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"trustedOpHash\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"invertBidValue\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exPostBids\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"allowAllocateValueFailure\",\"type\":\"bool\"}],\"internalType\":\"structCallConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getDAppConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"solverGasLimit\",\"type\":\"uint32\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDAppSignatory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSolverGasLimit\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingGovernance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"solved\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"postOpsCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"postSolverCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"preOpsCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"preSolverCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireSequentialDAppNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSequential\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireSequentialUserNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isSequential\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenUserBuys\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserBuys\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenUserSells\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserSells\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"auctionBaseCurrency\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"antecedent\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"context\",\"type\":\"bytes\"}],\"internalType\":\"structCondition[]\",\"name\":\"conditions\",\"type\":\"tuple[]\"}],\"internalType\":\"structSwapIntent\",\"name\":\"swapIntent\",\"type\":\"tuple\"}],\"name\":\"swap\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"tokenUserBuys\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserBuys\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenUserSells\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountUserSells\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"auctionBaseCurrency\",\"type\":\"address\"}],\"internalType\":\"structSwapData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"transferGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userDelegated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"delegated\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SwapIntentDappControlABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapIntentDappControlMetaData.ABI instead.
var SwapIntentDappControlABI = SwapIntentDappControlMetaData.ABI

// SwapIntentDappControl is an auto generated Go binding around an Ethereum contract.
type SwapIntentDappControl struct {
	SwapIntentDappControlCaller     // Read-only binding to the contract
	SwapIntentDappControlTransactor // Write-only binding to the contract
	SwapIntentDappControlFilterer   // Log filterer for contract events
}

// SwapIntentDappControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapIntentDappControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapIntentDappControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapIntentDappControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapIntentDappControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapIntentDappControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapIntentDappControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapIntentDappControlSession struct {
	Contract     *SwapIntentDappControl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SwapIntentDappControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapIntentDappControlCallerSession struct {
	Contract *SwapIntentDappControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// SwapIntentDappControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapIntentDappControlTransactorSession struct {
	Contract     *SwapIntentDappControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// SwapIntentDappControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapIntentDappControlRaw struct {
	Contract *SwapIntentDappControl // Generic contract binding to access the raw methods on
}

// SwapIntentDappControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapIntentDappControlCallerRaw struct {
	Contract *SwapIntentDappControlCaller // Generic read-only contract binding to access the raw methods on
}

// SwapIntentDappControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapIntentDappControlTransactorRaw struct {
	Contract *SwapIntentDappControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapIntentDappControl creates a new instance of SwapIntentDappControl, bound to a specific deployed contract.
func NewSwapIntentDappControl(address common.Address, backend bind.ContractBackend) (*SwapIntentDappControl, error) {
	contract, err := bindSwapIntentDappControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapIntentDappControl{SwapIntentDappControlCaller: SwapIntentDappControlCaller{contract: contract}, SwapIntentDappControlTransactor: SwapIntentDappControlTransactor{contract: contract}, SwapIntentDappControlFilterer: SwapIntentDappControlFilterer{contract: contract}}, nil
}

// NewSwapIntentDappControlCaller creates a new read-only instance of SwapIntentDappControl, bound to a specific deployed contract.
func NewSwapIntentDappControlCaller(address common.Address, caller bind.ContractCaller) (*SwapIntentDappControlCaller, error) {
	contract, err := bindSwapIntentDappControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapIntentDappControlCaller{contract: contract}, nil
}

// NewSwapIntentDappControlTransactor creates a new write-only instance of SwapIntentDappControl, bound to a specific deployed contract.
func NewSwapIntentDappControlTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapIntentDappControlTransactor, error) {
	contract, err := bindSwapIntentDappControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapIntentDappControlTransactor{contract: contract}, nil
}

// NewSwapIntentDappControlFilterer creates a new log filterer instance of SwapIntentDappControl, bound to a specific deployed contract.
func NewSwapIntentDappControlFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapIntentDappControlFilterer, error) {
	contract, err := bindSwapIntentDappControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapIntentDappControlFilterer{contract: contract}, nil
}

// bindSwapIntentDappControl binds a generic wrapper to an already deployed contract.
func bindSwapIntentDappControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapIntentDappControlMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapIntentDappControl *SwapIntentDappControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapIntentDappControl.Contract.SwapIntentDappControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapIntentDappControl *SwapIntentDappControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.SwapIntentDappControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapIntentDappControl *SwapIntentDappControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.SwapIntentDappControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapIntentDappControl *SwapIntentDappControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapIntentDappControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapIntentDappControl *SwapIntentDappControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapIntentDappControl *SwapIntentDappControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.contract.Transact(opts, method, params...)
}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlCaller) ATLAS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapIntentDappControl.contract.Call(opts, &out, "ATLAS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlSession) ATLAS() (common.Address, error) {
	return _SwapIntentDappControl.Contract.ATLAS(&_SwapIntentDappControl.CallOpts)
}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlCallerSession) ATLAS() (common.Address, error) {
	return _SwapIntentDappControl.Contract.ATLAS(&_SwapIntentDappControl.CallOpts)
}

// ATLASVERIFICATION is a free data retrieval call binding the contract method 0xbf230cfb.
//
// Solidity: function ATLAS_VERIFICATION() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlCaller) ATLASVERIFICATION(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapIntentDappControl.contract.Call(opts, &out, "ATLAS_VERIFICATION")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ATLASVERIFICATION is a free data retrieval call binding the contract method 0xbf230cfb.
//
// Solidity: function ATLAS_VERIFICATION() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlSession) ATLASVERIFICATION() (common.Address, error) {
	return _SwapIntentDappControl.Contract.ATLASVERIFICATION(&_SwapIntentDappControl.CallOpts)
}

// ATLASVERIFICATION is a free data retrieval call binding the contract method 0xbf230cfb.
//
// Solidity: function ATLAS_VERIFICATION() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlCallerSession) ATLASVERIFICATION() (common.Address, error) {
	return _SwapIntentDappControl.Contract.ATLASVERIFICATION(&_SwapIntentDappControl.CallOpts)
}

// CALLCONFIG is a free data retrieval call binding the contract method 0x8d212978.
//
// Solidity: function CALL_CONFIG() view returns(uint32)
func (_SwapIntentDappControl *SwapIntentDappControlCaller) CALLCONFIG(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SwapIntentDappControl.contract.Call(opts, &out, "CALL_CONFIG")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CALLCONFIG is a free data retrieval call binding the contract method 0x8d212978.
//
// Solidity: function CALL_CONFIG() view returns(uint32)
func (_SwapIntentDappControl *SwapIntentDappControlSession) CALLCONFIG() (uint32, error) {
	return _SwapIntentDappControl.Contract.CALLCONFIG(&_SwapIntentDappControl.CallOpts)
}

// CALLCONFIG is a free data retrieval call binding the contract method 0x8d212978.
//
// Solidity: function CALL_CONFIG() view returns(uint32)
func (_SwapIntentDappControl *SwapIntentDappControlCallerSession) CALLCONFIG() (uint32, error) {
	return _SwapIntentDappControl.Contract.CALLCONFIG(&_SwapIntentDappControl.CallOpts)
}

// CONTROL is a free data retrieval call binding the contract method 0x4953ecc7.
//
// Solidity: function CONTROL() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlCaller) CONTROL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapIntentDappControl.contract.Call(opts, &out, "CONTROL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CONTROL is a free data retrieval call binding the contract method 0x4953ecc7.
//
// Solidity: function CONTROL() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlSession) CONTROL() (common.Address, error) {
	return _SwapIntentDappControl.Contract.CONTROL(&_SwapIntentDappControl.CallOpts)
}

// CONTROL is a free data retrieval call binding the contract method 0x4953ecc7.
//
// Solidity: function CONTROL() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlCallerSession) CONTROL() (common.Address, error) {
	return _SwapIntentDappControl.Contract.CONTROL(&_SwapIntentDappControl.CallOpts)
}

// MAXUSERCONDITIONS is a free data retrieval call binding the contract method 0x1bcda8c5.
//
// Solidity: function MAX_USER_CONDITIONS() view returns(uint256)
func (_SwapIntentDappControl *SwapIntentDappControlCaller) MAXUSERCONDITIONS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapIntentDappControl.contract.Call(opts, &out, "MAX_USER_CONDITIONS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXUSERCONDITIONS is a free data retrieval call binding the contract method 0x1bcda8c5.
//
// Solidity: function MAX_USER_CONDITIONS() view returns(uint256)
func (_SwapIntentDappControl *SwapIntentDappControlSession) MAXUSERCONDITIONS() (*big.Int, error) {
	return _SwapIntentDappControl.Contract.MAXUSERCONDITIONS(&_SwapIntentDappControl.CallOpts)
}

// MAXUSERCONDITIONS is a free data retrieval call binding the contract method 0x1bcda8c5.
//
// Solidity: function MAX_USER_CONDITIONS() view returns(uint256)
func (_SwapIntentDappControl *SwapIntentDappControlCallerSession) MAXUSERCONDITIONS() (*big.Int, error) {
	return _SwapIntentDappControl.Contract.MAXUSERCONDITIONS(&_SwapIntentDappControl.CallOpts)
}

// SOURCE is a free data retrieval call binding the contract method 0xf230b4c2.
//
// Solidity: function SOURCE() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlCaller) SOURCE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapIntentDappControl.contract.Call(opts, &out, "SOURCE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SOURCE is a free data retrieval call binding the contract method 0xf230b4c2.
//
// Solidity: function SOURCE() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlSession) SOURCE() (common.Address, error) {
	return _SwapIntentDappControl.Contract.SOURCE(&_SwapIntentDappControl.CallOpts)
}

// SOURCE is a free data retrieval call binding the contract method 0xf230b4c2.
//
// Solidity: function SOURCE() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlCallerSession) SOURCE() (common.Address, error) {
	return _SwapIntentDappControl.Contract.SOURCE(&_SwapIntentDappControl.CallOpts)
}

// USERCONDITIONGASLIMIT is a free data retrieval call binding the contract method 0x13f32cf5.
//
// Solidity: function USER_CONDITION_GAS_LIMIT() view returns(uint256)
func (_SwapIntentDappControl *SwapIntentDappControlCaller) USERCONDITIONGASLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapIntentDappControl.contract.Call(opts, &out, "USER_CONDITION_GAS_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// USERCONDITIONGASLIMIT is a free data retrieval call binding the contract method 0x13f32cf5.
//
// Solidity: function USER_CONDITION_GAS_LIMIT() view returns(uint256)
func (_SwapIntentDappControl *SwapIntentDappControlSession) USERCONDITIONGASLIMIT() (*big.Int, error) {
	return _SwapIntentDappControl.Contract.USERCONDITIONGASLIMIT(&_SwapIntentDappControl.CallOpts)
}

// USERCONDITIONGASLIMIT is a free data retrieval call binding the contract method 0x13f32cf5.
//
// Solidity: function USER_CONDITION_GAS_LIMIT() view returns(uint256)
func (_SwapIntentDappControl *SwapIntentDappControlCallerSession) USERCONDITIONGASLIMIT() (*big.Int, error) {
	return _SwapIntentDappControl.Contract.USERCONDITIONGASLIMIT(&_SwapIntentDappControl.CallOpts)
}

// GetBidFormat is a free data retrieval call binding the contract method 0x8831b924.
//
// Solidity: function getBidFormat((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) pure returns(address bidToken)
func (_SwapIntentDappControl *SwapIntentDappControlCaller) GetBidFormat(opts *bind.CallOpts, userOp UserOperation) (common.Address, error) {
	var out []interface{}
	err := _SwapIntentDappControl.contract.Call(opts, &out, "getBidFormat", userOp)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBidFormat is a free data retrieval call binding the contract method 0x8831b924.
//
// Solidity: function getBidFormat((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) pure returns(address bidToken)
func (_SwapIntentDappControl *SwapIntentDappControlSession) GetBidFormat(userOp UserOperation) (common.Address, error) {
	return _SwapIntentDappControl.Contract.GetBidFormat(&_SwapIntentDappControl.CallOpts, userOp)
}

// GetBidFormat is a free data retrieval call binding the contract method 0x8831b924.
//
// Solidity: function getBidFormat((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) pure returns(address bidToken)
func (_SwapIntentDappControl *SwapIntentDappControlCallerSession) GetBidFormat(userOp UserOperation) (common.Address, error) {
	return _SwapIntentDappControl.Contract.GetBidFormat(&_SwapIntentDappControl.CallOpts, userOp)
}

// GetBidValue is a free data retrieval call binding the contract method 0x6d25fc9a.
//
// Solidity: function getBidValue((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) pure returns(uint256)
func (_SwapIntentDappControl *SwapIntentDappControlCaller) GetBidValue(opts *bind.CallOpts, solverOp SolverOperation) (*big.Int, error) {
	var out []interface{}
	err := _SwapIntentDappControl.contract.Call(opts, &out, "getBidValue", solverOp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidValue is a free data retrieval call binding the contract method 0x6d25fc9a.
//
// Solidity: function getBidValue((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) pure returns(uint256)
func (_SwapIntentDappControl *SwapIntentDappControlSession) GetBidValue(solverOp SolverOperation) (*big.Int, error) {
	return _SwapIntentDappControl.Contract.GetBidValue(&_SwapIntentDappControl.CallOpts, solverOp)
}

// GetBidValue is a free data retrieval call binding the contract method 0x6d25fc9a.
//
// Solidity: function getBidValue((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) pure returns(uint256)
func (_SwapIntentDappControl *SwapIntentDappControlCallerSession) GetBidValue(solverOp SolverOperation) (*big.Int, error) {
	return _SwapIntentDappControl.Contract.GetBidValue(&_SwapIntentDappControl.CallOpts, solverOp)
}

// GetCallConfig is a free data retrieval call binding the contract method 0x6e1ccfc4.
//
// Solidity: function getCallConfig() view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_SwapIntentDappControl *SwapIntentDappControlCaller) GetCallConfig(opts *bind.CallOpts) (CallConfig, error) {
	var out []interface{}
	err := _SwapIntentDappControl.contract.Call(opts, &out, "getCallConfig")

	if err != nil {
		return *new(CallConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(CallConfig)).(*CallConfig)

	return out0, err

}

// GetCallConfig is a free data retrieval call binding the contract method 0x6e1ccfc4.
//
// Solidity: function getCallConfig() view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_SwapIntentDappControl *SwapIntentDappControlSession) GetCallConfig() (CallConfig, error) {
	return _SwapIntentDappControl.Contract.GetCallConfig(&_SwapIntentDappControl.CallOpts)
}

// GetCallConfig is a free data retrieval call binding the contract method 0x6e1ccfc4.
//
// Solidity: function getCallConfig() view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,bool))
func (_SwapIntentDappControl *SwapIntentDappControlCallerSession) GetCallConfig() (CallConfig, error) {
	return _SwapIntentDappControl.Contract.GetCallConfig(&_SwapIntentDappControl.CallOpts)
}

// GetDAppConfig is a free data retrieval call binding the contract method 0x44912b6e.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) view returns((address,uint32,address,uint32) dConfig)
func (_SwapIntentDappControl *SwapIntentDappControlCaller) GetDAppConfig(opts *bind.CallOpts, userOp UserOperation) (DAppConfig, error) {
	var out []interface{}
	err := _SwapIntentDappControl.contract.Call(opts, &out, "getDAppConfig", userOp)

	if err != nil {
		return *new(DAppConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(DAppConfig)).(*DAppConfig)

	return out0, err

}

// GetDAppConfig is a free data retrieval call binding the contract method 0x44912b6e.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) view returns((address,uint32,address,uint32) dConfig)
func (_SwapIntentDappControl *SwapIntentDappControlSession) GetDAppConfig(userOp UserOperation) (DAppConfig, error) {
	return _SwapIntentDappControl.Contract.GetDAppConfig(&_SwapIntentDappControl.CallOpts, userOp)
}

// GetDAppConfig is a free data retrieval call binding the contract method 0x44912b6e.
//
// Solidity: function getDAppConfig((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) view returns((address,uint32,address,uint32) dConfig)
func (_SwapIntentDappControl *SwapIntentDappControlCallerSession) GetDAppConfig(userOp UserOperation) (DAppConfig, error) {
	return _SwapIntentDappControl.Contract.GetDAppConfig(&_SwapIntentDappControl.CallOpts, userOp)
}

// GetDAppSignatory is a free data retrieval call binding the contract method 0xee418488.
//
// Solidity: function getDAppSignatory() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlCaller) GetDAppSignatory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapIntentDappControl.contract.Call(opts, &out, "getDAppSignatory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDAppSignatory is a free data retrieval call binding the contract method 0xee418488.
//
// Solidity: function getDAppSignatory() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlSession) GetDAppSignatory() (common.Address, error) {
	return _SwapIntentDappControl.Contract.GetDAppSignatory(&_SwapIntentDappControl.CallOpts)
}

// GetDAppSignatory is a free data retrieval call binding the contract method 0xee418488.
//
// Solidity: function getDAppSignatory() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlCallerSession) GetDAppSignatory() (common.Address, error) {
	return _SwapIntentDappControl.Contract.GetDAppSignatory(&_SwapIntentDappControl.CallOpts)
}

// GetSolverGasLimit is a free data retrieval call binding the contract method 0x99218be5.
//
// Solidity: function getSolverGasLimit() view returns(uint32)
func (_SwapIntentDappControl *SwapIntentDappControlCaller) GetSolverGasLimit(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SwapIntentDappControl.contract.Call(opts, &out, "getSolverGasLimit")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetSolverGasLimit is a free data retrieval call binding the contract method 0x99218be5.
//
// Solidity: function getSolverGasLimit() view returns(uint32)
func (_SwapIntentDappControl *SwapIntentDappControlSession) GetSolverGasLimit() (uint32, error) {
	return _SwapIntentDappControl.Contract.GetSolverGasLimit(&_SwapIntentDappControl.CallOpts)
}

// GetSolverGasLimit is a free data retrieval call binding the contract method 0x99218be5.
//
// Solidity: function getSolverGasLimit() view returns(uint32)
func (_SwapIntentDappControl *SwapIntentDappControlCallerSession) GetSolverGasLimit() (uint32, error) {
	return _SwapIntentDappControl.Contract.GetSolverGasLimit(&_SwapIntentDappControl.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapIntentDappControl.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlSession) Governance() (common.Address, error) {
	return _SwapIntentDappControl.Contract.Governance(&_SwapIntentDappControl.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlCallerSession) Governance() (common.Address, error) {
	return _SwapIntentDappControl.Contract.Governance(&_SwapIntentDappControl.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlCaller) PendingGovernance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapIntentDappControl.contract.Call(opts, &out, "pendingGovernance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlSession) PendingGovernance() (common.Address, error) {
	return _SwapIntentDappControl.Contract.PendingGovernance(&_SwapIntentDappControl.CallOpts)
}

// PendingGovernance is a free data retrieval call binding the contract method 0xf39c38a0.
//
// Solidity: function pendingGovernance() view returns(address)
func (_SwapIntentDappControl *SwapIntentDappControlCallerSession) PendingGovernance() (common.Address, error) {
	return _SwapIntentDappControl.Contract.PendingGovernance(&_SwapIntentDappControl.CallOpts)
}

// RequireSequentialDAppNonces is a free data retrieval call binding the contract method 0x72d91684.
//
// Solidity: function requireSequentialDAppNonces() view returns(bool isSequential)
func (_SwapIntentDappControl *SwapIntentDappControlCaller) RequireSequentialDAppNonces(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SwapIntentDappControl.contract.Call(opts, &out, "requireSequentialDAppNonces")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequireSequentialDAppNonces is a free data retrieval call binding the contract method 0x72d91684.
//
// Solidity: function requireSequentialDAppNonces() view returns(bool isSequential)
func (_SwapIntentDappControl *SwapIntentDappControlSession) RequireSequentialDAppNonces() (bool, error) {
	return _SwapIntentDappControl.Contract.RequireSequentialDAppNonces(&_SwapIntentDappControl.CallOpts)
}

// RequireSequentialDAppNonces is a free data retrieval call binding the contract method 0x72d91684.
//
// Solidity: function requireSequentialDAppNonces() view returns(bool isSequential)
func (_SwapIntentDappControl *SwapIntentDappControlCallerSession) RequireSequentialDAppNonces() (bool, error) {
	return _SwapIntentDappControl.Contract.RequireSequentialDAppNonces(&_SwapIntentDappControl.CallOpts)
}

// RequireSequentialUserNonces is a free data retrieval call binding the contract method 0xe2c0c30f.
//
// Solidity: function requireSequentialUserNonces() view returns(bool isSequential)
func (_SwapIntentDappControl *SwapIntentDappControlCaller) RequireSequentialUserNonces(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SwapIntentDappControl.contract.Call(opts, &out, "requireSequentialUserNonces")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequireSequentialUserNonces is a free data retrieval call binding the contract method 0xe2c0c30f.
//
// Solidity: function requireSequentialUserNonces() view returns(bool isSequential)
func (_SwapIntentDappControl *SwapIntentDappControlSession) RequireSequentialUserNonces() (bool, error) {
	return _SwapIntentDappControl.Contract.RequireSequentialUserNonces(&_SwapIntentDappControl.CallOpts)
}

// RequireSequentialUserNonces is a free data retrieval call binding the contract method 0xe2c0c30f.
//
// Solidity: function requireSequentialUserNonces() view returns(bool isSequential)
func (_SwapIntentDappControl *SwapIntentDappControlCallerSession) RequireSequentialUserNonces() (bool, error) {
	return _SwapIntentDappControl.Contract.RequireSequentialUserNonces(&_SwapIntentDappControl.CallOpts)
}

// UserDelegated is a free data retrieval call binding the contract method 0x1e151167.
//
// Solidity: function userDelegated() view returns(bool delegated)
func (_SwapIntentDappControl *SwapIntentDappControlCaller) UserDelegated(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SwapIntentDappControl.contract.Call(opts, &out, "userDelegated")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserDelegated is a free data retrieval call binding the contract method 0x1e151167.
//
// Solidity: function userDelegated() view returns(bool delegated)
func (_SwapIntentDappControl *SwapIntentDappControlSession) UserDelegated() (bool, error) {
	return _SwapIntentDappControl.Contract.UserDelegated(&_SwapIntentDappControl.CallOpts)
}

// UserDelegated is a free data retrieval call binding the contract method 0x1e151167.
//
// Solidity: function userDelegated() view returns(bool delegated)
func (_SwapIntentDappControl *SwapIntentDappControlCallerSession) UserDelegated() (bool, error) {
	return _SwapIntentDappControl.Contract.UserDelegated(&_SwapIntentDappControl.CallOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_SwapIntentDappControl *SwapIntentDappControlTransactor) AcceptGovernance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapIntentDappControl.contract.Transact(opts, "acceptGovernance")
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_SwapIntentDappControl *SwapIntentDappControlSession) AcceptGovernance() (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.AcceptGovernance(&_SwapIntentDappControl.TransactOpts)
}

// AcceptGovernance is a paid mutator transaction binding the contract method 0x238efcbc.
//
// Solidity: function acceptGovernance() returns()
func (_SwapIntentDappControl *SwapIntentDappControlTransactorSession) AcceptGovernance() (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.AcceptGovernance(&_SwapIntentDappControl.TransactOpts)
}

// AllocateValueCall is a paid mutator transaction binding the contract method 0x2f5e0d16.
//
// Solidity: function allocateValueCall(address bidToken, uint256 bidAmount, bytes data) returns()
func (_SwapIntentDappControl *SwapIntentDappControlTransactor) AllocateValueCall(opts *bind.TransactOpts, bidToken common.Address, bidAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _SwapIntentDappControl.contract.Transact(opts, "allocateValueCall", bidToken, bidAmount, data)
}

// AllocateValueCall is a paid mutator transaction binding the contract method 0x2f5e0d16.
//
// Solidity: function allocateValueCall(address bidToken, uint256 bidAmount, bytes data) returns()
func (_SwapIntentDappControl *SwapIntentDappControlSession) AllocateValueCall(bidToken common.Address, bidAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.AllocateValueCall(&_SwapIntentDappControl.TransactOpts, bidToken, bidAmount, data)
}

// AllocateValueCall is a paid mutator transaction binding the contract method 0x2f5e0d16.
//
// Solidity: function allocateValueCall(address bidToken, uint256 bidAmount, bytes data) returns()
func (_SwapIntentDappControl *SwapIntentDappControlTransactorSession) AllocateValueCall(bidToken common.Address, bidAmount *big.Int, data []byte) (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.AllocateValueCall(&_SwapIntentDappControl.TransactOpts, bidToken, bidAmount, data)
}

// PostOpsCall is a paid mutator transaction binding the contract method 0x836a611b.
//
// Solidity: function postOpsCall(bool solved, bytes data) payable returns(bool)
func (_SwapIntentDappControl *SwapIntentDappControlTransactor) PostOpsCall(opts *bind.TransactOpts, solved bool, data []byte) (*types.Transaction, error) {
	return _SwapIntentDappControl.contract.Transact(opts, "postOpsCall", solved, data)
}

// PostOpsCall is a paid mutator transaction binding the contract method 0x836a611b.
//
// Solidity: function postOpsCall(bool solved, bytes data) payable returns(bool)
func (_SwapIntentDappControl *SwapIntentDappControlSession) PostOpsCall(solved bool, data []byte) (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.PostOpsCall(&_SwapIntentDappControl.TransactOpts, solved, data)
}

// PostOpsCall is a paid mutator transaction binding the contract method 0x836a611b.
//
// Solidity: function postOpsCall(bool solved, bytes data) payable returns(bool)
func (_SwapIntentDappControl *SwapIntentDappControlTransactorSession) PostOpsCall(solved bool, data []byte) (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.PostOpsCall(&_SwapIntentDappControl.TransactOpts, solved, data)
}

// PostSolverCall is a paid mutator transaction binding the contract method 0x6d4d6b2e.
//
// Solidity: function postSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_SwapIntentDappControl *SwapIntentDappControlTransactor) PostSolverCall(opts *bind.TransactOpts, solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _SwapIntentDappControl.contract.Transact(opts, "postSolverCall", solverOp, returnData)
}

// PostSolverCall is a paid mutator transaction binding the contract method 0x6d4d6b2e.
//
// Solidity: function postSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_SwapIntentDappControl *SwapIntentDappControlSession) PostSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.PostSolverCall(&_SwapIntentDappControl.TransactOpts, solverOp, returnData)
}

// PostSolverCall is a paid mutator transaction binding the contract method 0x6d4d6b2e.
//
// Solidity: function postSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_SwapIntentDappControl *SwapIntentDappControlTransactorSession) PostSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.PostSolverCall(&_SwapIntentDappControl.TransactOpts, solverOp, returnData)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0x77bceb1b.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) payable returns(bytes)
func (_SwapIntentDappControl *SwapIntentDappControlTransactor) PreOpsCall(opts *bind.TransactOpts, userOp UserOperation) (*types.Transaction, error) {
	return _SwapIntentDappControl.contract.Transact(opts, "preOpsCall", userOp)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0x77bceb1b.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) payable returns(bytes)
func (_SwapIntentDappControl *SwapIntentDappControlSession) PreOpsCall(userOp UserOperation) (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.PreOpsCall(&_SwapIntentDappControl.TransactOpts, userOp)
}

// PreOpsCall is a paid mutator transaction binding the contract method 0x77bceb1b.
//
// Solidity: function preOpsCall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) payable returns(bytes)
func (_SwapIntentDappControl *SwapIntentDappControlTransactorSession) PreOpsCall(userOp UserOperation) (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.PreOpsCall(&_SwapIntentDappControl.TransactOpts, userOp)
}

// PreSolverCall is a paid mutator transaction binding the contract method 0x283ee1cf.
//
// Solidity: function preSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_SwapIntentDappControl *SwapIntentDappControlTransactor) PreSolverCall(opts *bind.TransactOpts, solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _SwapIntentDappControl.contract.Transact(opts, "preSolverCall", solverOp, returnData)
}

// PreSolverCall is a paid mutator transaction binding the contract method 0x283ee1cf.
//
// Solidity: function preSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_SwapIntentDappControl *SwapIntentDappControlSession) PreSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.PreSolverCall(&_SwapIntentDappControl.TransactOpts, solverOp, returnData)
}

// PreSolverCall is a paid mutator transaction binding the contract method 0x283ee1cf.
//
// Solidity: function preSolverCall((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes returnData) payable returns()
func (_SwapIntentDappControl *SwapIntentDappControlTransactorSession) PreSolverCall(solverOp SolverOperation, returnData []byte) (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.PreSolverCall(&_SwapIntentDappControl.TransactOpts, solverOp, returnData)
}

// Swap is a paid mutator transaction binding the contract method 0x4a9de849.
//
// Solidity: function swap((address,uint256,address,uint256,address,(address,bytes)[]) swapIntent) payable returns((address,uint256,address,uint256,address))
func (_SwapIntentDappControl *SwapIntentDappControlTransactor) Swap(opts *bind.TransactOpts, swapIntent SwapIntent) (*types.Transaction, error) {
	return _SwapIntentDappControl.contract.Transact(opts, "swap", swapIntent)
}

// Swap is a paid mutator transaction binding the contract method 0x4a9de849.
//
// Solidity: function swap((address,uint256,address,uint256,address,(address,bytes)[]) swapIntent) payable returns((address,uint256,address,uint256,address))
func (_SwapIntentDappControl *SwapIntentDappControlSession) Swap(swapIntent SwapIntent) (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.Swap(&_SwapIntentDappControl.TransactOpts, swapIntent)
}

// Swap is a paid mutator transaction binding the contract method 0x4a9de849.
//
// Solidity: function swap((address,uint256,address,uint256,address,(address,bytes)[]) swapIntent) payable returns((address,uint256,address,uint256,address))
func (_SwapIntentDappControl *SwapIntentDappControlTransactorSession) Swap(swapIntent SwapIntent) (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.Swap(&_SwapIntentDappControl.TransactOpts, swapIntent)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address newGovernance) returns()
func (_SwapIntentDappControl *SwapIntentDappControlTransactor) TransferGovernance(opts *bind.TransactOpts, newGovernance common.Address) (*types.Transaction, error) {
	return _SwapIntentDappControl.contract.Transact(opts, "transferGovernance", newGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address newGovernance) returns()
func (_SwapIntentDappControl *SwapIntentDappControlSession) TransferGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.TransferGovernance(&_SwapIntentDappControl.TransactOpts, newGovernance)
}

// TransferGovernance is a paid mutator transaction binding the contract method 0xd38bfff4.
//
// Solidity: function transferGovernance(address newGovernance) returns()
func (_SwapIntentDappControl *SwapIntentDappControlTransactorSession) TransferGovernance(newGovernance common.Address) (*types.Transaction, error) {
	return _SwapIntentDappControl.Contract.TransferGovernance(&_SwapIntentDappControl.TransactOpts, newGovernance)
}

// SwapIntentDappControlGovernanceTransferStartedIterator is returned from FilterGovernanceTransferStarted and is used to iterate over the raw logs and unpacked data for GovernanceTransferStarted events raised by the SwapIntentDappControl contract.
type SwapIntentDappControlGovernanceTransferStartedIterator struct {
	Event *SwapIntentDappControlGovernanceTransferStarted // Event containing the contract specifics and raw log

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
func (it *SwapIntentDappControlGovernanceTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentDappControlGovernanceTransferStarted)
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
		it.Event = new(SwapIntentDappControlGovernanceTransferStarted)
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
func (it *SwapIntentDappControlGovernanceTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentDappControlGovernanceTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentDappControlGovernanceTransferStarted represents a GovernanceTransferStarted event raised by the SwapIntentDappControl contract.
type SwapIntentDappControlGovernanceTransferStarted struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferStarted is a free log retrieval operation binding the contract event 0x1c4bb4e3cab7b72da7eb9f0ae62554dda85dc7fb907c946ad2776095b95ac1ad.
//
// Solidity: event GovernanceTransferStarted(address indexed previousGovernance, address indexed newGovernance)
func (_SwapIntentDappControl *SwapIntentDappControlFilterer) FilterGovernanceTransferStarted(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*SwapIntentDappControlGovernanceTransferStartedIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _SwapIntentDappControl.contract.FilterLogs(opts, "GovernanceTransferStarted", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &SwapIntentDappControlGovernanceTransferStartedIterator{contract: _SwapIntentDappControl.contract, event: "GovernanceTransferStarted", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferStarted is a free log subscription operation binding the contract event 0x1c4bb4e3cab7b72da7eb9f0ae62554dda85dc7fb907c946ad2776095b95ac1ad.
//
// Solidity: event GovernanceTransferStarted(address indexed previousGovernance, address indexed newGovernance)
func (_SwapIntentDappControl *SwapIntentDappControlFilterer) WatchGovernanceTransferStarted(opts *bind.WatchOpts, sink chan<- *SwapIntentDappControlGovernanceTransferStarted, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _SwapIntentDappControl.contract.WatchLogs(opts, "GovernanceTransferStarted", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentDappControlGovernanceTransferStarted)
				if err := _SwapIntentDappControl.contract.UnpackLog(event, "GovernanceTransferStarted", log); err != nil {
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
func (_SwapIntentDappControl *SwapIntentDappControlFilterer) ParseGovernanceTransferStarted(log types.Log) (*SwapIntentDappControlGovernanceTransferStarted, error) {
	event := new(SwapIntentDappControlGovernanceTransferStarted)
	if err := _SwapIntentDappControl.contract.UnpackLog(event, "GovernanceTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapIntentDappControlGovernanceTransferredIterator is returned from FilterGovernanceTransferred and is used to iterate over the raw logs and unpacked data for GovernanceTransferred events raised by the SwapIntentDappControl contract.
type SwapIntentDappControlGovernanceTransferredIterator struct {
	Event *SwapIntentDappControlGovernanceTransferred // Event containing the contract specifics and raw log

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
func (it *SwapIntentDappControlGovernanceTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapIntentDappControlGovernanceTransferred)
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
		it.Event = new(SwapIntentDappControlGovernanceTransferred)
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
func (it *SwapIntentDappControlGovernanceTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapIntentDappControlGovernanceTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapIntentDappControlGovernanceTransferred represents a GovernanceTransferred event raised by the SwapIntentDappControl contract.
type SwapIntentDappControlGovernanceTransferred struct {
	PreviousGovernance common.Address
	NewGovernance      common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernanceTransferred is a free log retrieval operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_SwapIntentDappControl *SwapIntentDappControlFilterer) FilterGovernanceTransferred(opts *bind.FilterOpts, previousGovernance []common.Address, newGovernance []common.Address) (*SwapIntentDappControlGovernanceTransferredIterator, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _SwapIntentDappControl.contract.FilterLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &SwapIntentDappControlGovernanceTransferredIterator{contract: _SwapIntentDappControl.contract, event: "GovernanceTransferred", logs: logs, sub: sub}, nil
}

// WatchGovernanceTransferred is a free log subscription operation binding the contract event 0x5f56bee8cffbe9a78652a74a60705edede02af10b0bbb888ca44b79a0d42ce80.
//
// Solidity: event GovernanceTransferred(address indexed previousGovernance, address indexed newGovernance)
func (_SwapIntentDappControl *SwapIntentDappControlFilterer) WatchGovernanceTransferred(opts *bind.WatchOpts, sink chan<- *SwapIntentDappControlGovernanceTransferred, previousGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

	var previousGovernanceRule []interface{}
	for _, previousGovernanceItem := range previousGovernance {
		previousGovernanceRule = append(previousGovernanceRule, previousGovernanceItem)
	}
	var newGovernanceRule []interface{}
	for _, newGovernanceItem := range newGovernance {
		newGovernanceRule = append(newGovernanceRule, newGovernanceItem)
	}

	logs, sub, err := _SwapIntentDappControl.contract.WatchLogs(opts, "GovernanceTransferred", previousGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapIntentDappControlGovernanceTransferred)
				if err := _SwapIntentDappControl.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
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
func (_SwapIntentDappControl *SwapIntentDappControlFilterer) ParseGovernanceTransferred(log types.Log) (*SwapIntentDappControlGovernanceTransferred, error) {
	event := new(SwapIntentDappControlGovernanceTransferred)
	if err := _SwapIntentDappControl.contract.UnpackLog(event, "GovernanceTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
