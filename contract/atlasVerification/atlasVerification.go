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
	_ = abi.ConvertType
)

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

// AtlasVerificationMetaData contains all meta data concerning the AtlasVerification contract.
var AtlasVerificationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"atlas\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AtlasLockActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DAppNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignatory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SignatoryActive\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"name\":\"DAppDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldGovernance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"name\":\"DAppGovernanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"name\":\"NewDAppSignatory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"governance\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"}],\"name\":\"RemovedDAppSignatory\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ATLAS\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"S_dAppSequentialNonceTrackers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint248\",\"name\":\"\",\"type\":\"uint248\"}],\"name\":\"S_userNonSequentialNonceTrackers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"S_userSequentialNonceTrackers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"}],\"name\":\"addSignatory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newGovernance\",\"type\":\"address\"}],\"name\":\"changeDAppGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"dAppSequentialNonceTrackers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastUsedSeqNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"}],\"name\":\"dAppSignatories\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"}],\"name\":\"disableDApp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dApp\",\"type\":\"address\"}],\"name\":\"getDAppNextNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nextNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bundler\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"callChainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structDAppOperation\",\"name\":\"dAppOp\",\"type\":\"tuple\"}],\"name\":\"getDAppOperationPayload\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"domainSeparator\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"}],\"name\":\"getGovFromControl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"}],\"name\":\"getSolverPayload\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"refNonce\",\"type\":\"uint256\"}],\"name\":\"getUserNextNonSeqNonceAfter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"sequential\",\"type\":\"bool\"}],\"name\":\"getUserNextNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nextNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOperationHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"}],\"name\":\"getUserOperationPayload\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"payload\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"}],\"name\":\"initializeGovernance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"}],\"name\":\"isDAppSignatory\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signatory\",\"type\":\"address\"}],\"name\":\"removeSignatory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"signatories\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint248\",\"name\":\"wordIndex\",\"type\":\"uint248\"}],\"name\":\"userNonSequentialNonceTrackers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bitmap\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"userSequentialNonceTrackers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastUsedSeqNonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"solverGasLimit\",\"type\":\"uint32\"}],\"internalType\":\"structDAppConfig\",\"name\":\"dConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dapp\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"callConfig\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sessionKey\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"userOp\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation[]\",\"name\":\"solverOps\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bundler\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"callChainHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structDAppOperation\",\"name\":\"dAppOp\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"msgValue\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSimulation\",\"type\":\"bool\"}],\"name\":\"validateCalls\",\"outputs\":[{\"internalType\":\"enumValidCallsResult\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"solver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structSolverOperation\",\"name\":\"solverOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"userMaxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bundler\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowsTrustedOpHash\",\"type\":\"bool\"}],\"name\":\"verifySolverOp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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
	parsed, err := AtlasVerificationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_AtlasVerification *AtlasVerificationCaller) ATLAS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "ATLAS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_AtlasVerification *AtlasVerificationSession) ATLAS() (common.Address, error) {
	return _AtlasVerification.Contract.ATLAS(&_AtlasVerification.CallOpts)
}

// ATLAS is a free data retrieval call binding the contract method 0xe982ae92.
//
// Solidity: function ATLAS() view returns(address)
func (_AtlasVerification *AtlasVerificationCallerSession) ATLAS() (common.Address, error) {
	return _AtlasVerification.Contract.ATLAS(&_AtlasVerification.CallOpts)
}

// SDAppSequentialNonceTrackers is a free data retrieval call binding the contract method 0x2fd4a9e5.
//
// Solidity: function S_dAppSequentialNonceTrackers(address ) view returns(uint256)
func (_AtlasVerification *AtlasVerificationCaller) SDAppSequentialNonceTrackers(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "S_dAppSequentialNonceTrackers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SDAppSequentialNonceTrackers is a free data retrieval call binding the contract method 0x2fd4a9e5.
//
// Solidity: function S_dAppSequentialNonceTrackers(address ) view returns(uint256)
func (_AtlasVerification *AtlasVerificationSession) SDAppSequentialNonceTrackers(arg0 common.Address) (*big.Int, error) {
	return _AtlasVerification.Contract.SDAppSequentialNonceTrackers(&_AtlasVerification.CallOpts, arg0)
}

// SDAppSequentialNonceTrackers is a free data retrieval call binding the contract method 0x2fd4a9e5.
//
// Solidity: function S_dAppSequentialNonceTrackers(address ) view returns(uint256)
func (_AtlasVerification *AtlasVerificationCallerSession) SDAppSequentialNonceTrackers(arg0 common.Address) (*big.Int, error) {
	return _AtlasVerification.Contract.SDAppSequentialNonceTrackers(&_AtlasVerification.CallOpts, arg0)
}

// SUserNonSequentialNonceTrackers is a free data retrieval call binding the contract method 0x643f637a.
//
// Solidity: function S_userNonSequentialNonceTrackers(address , uint248 ) view returns(uint256)
func (_AtlasVerification *AtlasVerificationCaller) SUserNonSequentialNonceTrackers(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "S_userNonSequentialNonceTrackers", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SUserNonSequentialNonceTrackers is a free data retrieval call binding the contract method 0x643f637a.
//
// Solidity: function S_userNonSequentialNonceTrackers(address , uint248 ) view returns(uint256)
func (_AtlasVerification *AtlasVerificationSession) SUserNonSequentialNonceTrackers(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _AtlasVerification.Contract.SUserNonSequentialNonceTrackers(&_AtlasVerification.CallOpts, arg0, arg1)
}

// SUserNonSequentialNonceTrackers is a free data retrieval call binding the contract method 0x643f637a.
//
// Solidity: function S_userNonSequentialNonceTrackers(address , uint248 ) view returns(uint256)
func (_AtlasVerification *AtlasVerificationCallerSession) SUserNonSequentialNonceTrackers(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _AtlasVerification.Contract.SUserNonSequentialNonceTrackers(&_AtlasVerification.CallOpts, arg0, arg1)
}

// SUserSequentialNonceTrackers is a free data retrieval call binding the contract method 0x091c4021.
//
// Solidity: function S_userSequentialNonceTrackers(address ) view returns(uint256)
func (_AtlasVerification *AtlasVerificationCaller) SUserSequentialNonceTrackers(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "S_userSequentialNonceTrackers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SUserSequentialNonceTrackers is a free data retrieval call binding the contract method 0x091c4021.
//
// Solidity: function S_userSequentialNonceTrackers(address ) view returns(uint256)
func (_AtlasVerification *AtlasVerificationSession) SUserSequentialNonceTrackers(arg0 common.Address) (*big.Int, error) {
	return _AtlasVerification.Contract.SUserSequentialNonceTrackers(&_AtlasVerification.CallOpts, arg0)
}

// SUserSequentialNonceTrackers is a free data retrieval call binding the contract method 0x091c4021.
//
// Solidity: function S_userSequentialNonceTrackers(address ) view returns(uint256)
func (_AtlasVerification *AtlasVerificationCallerSession) SUserSequentialNonceTrackers(arg0 common.Address) (*big.Int, error) {
	return _AtlasVerification.Contract.SUserSequentialNonceTrackers(&_AtlasVerification.CallOpts, arg0)
}

// DAppSequentialNonceTrackers is a free data retrieval call binding the contract method 0xea380b97.
//
// Solidity: function dAppSequentialNonceTrackers(address account) view returns(uint256 lastUsedSeqNonce)
func (_AtlasVerification *AtlasVerificationCaller) DAppSequentialNonceTrackers(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "dAppSequentialNonceTrackers", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DAppSequentialNonceTrackers is a free data retrieval call binding the contract method 0xea380b97.
//
// Solidity: function dAppSequentialNonceTrackers(address account) view returns(uint256 lastUsedSeqNonce)
func (_AtlasVerification *AtlasVerificationSession) DAppSequentialNonceTrackers(account common.Address) (*big.Int, error) {
	return _AtlasVerification.Contract.DAppSequentialNonceTrackers(&_AtlasVerification.CallOpts, account)
}

// DAppSequentialNonceTrackers is a free data retrieval call binding the contract method 0xea380b97.
//
// Solidity: function dAppSequentialNonceTrackers(address account) view returns(uint256 lastUsedSeqNonce)
func (_AtlasVerification *AtlasVerificationCallerSession) DAppSequentialNonceTrackers(account common.Address) (*big.Int, error) {
	return _AtlasVerification.Contract.DAppSequentialNonceTrackers(&_AtlasVerification.CallOpts, account)
}

// DAppSignatories is a free data retrieval call binding the contract method 0x6564524a.
//
// Solidity: function dAppSignatories(address control) view returns(address[])
func (_AtlasVerification *AtlasVerificationCaller) DAppSignatories(opts *bind.CallOpts, control common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "dAppSignatories", control)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// DAppSignatories is a free data retrieval call binding the contract method 0x6564524a.
//
// Solidity: function dAppSignatories(address control) view returns(address[])
func (_AtlasVerification *AtlasVerificationSession) DAppSignatories(control common.Address) ([]common.Address, error) {
	return _AtlasVerification.Contract.DAppSignatories(&_AtlasVerification.CallOpts, control)
}

// DAppSignatories is a free data retrieval call binding the contract method 0x6564524a.
//
// Solidity: function dAppSignatories(address control) view returns(address[])
func (_AtlasVerification *AtlasVerificationCallerSession) DAppSignatories(control common.Address) ([]common.Address, error) {
	return _AtlasVerification.Contract.DAppSignatories(&_AtlasVerification.CallOpts, control)
}

// GetDAppNextNonce is a free data retrieval call binding the contract method 0x9776f11e.
//
// Solidity: function getDAppNextNonce(address dApp) view returns(uint256 nextNonce)
func (_AtlasVerification *AtlasVerificationCaller) GetDAppNextNonce(opts *bind.CallOpts, dApp common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getDAppNextNonce", dApp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDAppNextNonce is a free data retrieval call binding the contract method 0x9776f11e.
//
// Solidity: function getDAppNextNonce(address dApp) view returns(uint256 nextNonce)
func (_AtlasVerification *AtlasVerificationSession) GetDAppNextNonce(dApp common.Address) (*big.Int, error) {
	return _AtlasVerification.Contract.GetDAppNextNonce(&_AtlasVerification.CallOpts, dApp)
}

// GetDAppNextNonce is a free data retrieval call binding the contract method 0x9776f11e.
//
// Solidity: function getDAppNextNonce(address dApp) view returns(uint256 nextNonce)
func (_AtlasVerification *AtlasVerificationCallerSession) GetDAppNextNonce(dApp common.Address) (*big.Int, error) {
	return _AtlasVerification.Contract.GetDAppNextNonce(&_AtlasVerification.CallOpts, dApp)
}

// GetDAppOperationPayload is a free data retrieval call binding the contract method 0x246c8c00.
//
// Solidity: function getDAppOperationPayload((address,address,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationCaller) GetDAppOperationPayload(opts *bind.CallOpts, dAppOp DAppOperation) ([32]byte, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getDAppOperationPayload", dAppOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDAppOperationPayload is a free data retrieval call binding the contract method 0x246c8c00.
//
// Solidity: function getDAppOperationPayload((address,address,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationSession) GetDAppOperationPayload(dAppOp DAppOperation) ([32]byte, error) {
	return _AtlasVerification.Contract.GetDAppOperationPayload(&_AtlasVerification.CallOpts, dAppOp)
}

// GetDAppOperationPayload is a free data retrieval call binding the contract method 0x246c8c00.
//
// Solidity: function getDAppOperationPayload((address,address,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationCallerSession) GetDAppOperationPayload(dAppOp DAppOperation) ([32]byte, error) {
	return _AtlasVerification.Contract.GetDAppOperationPayload(&_AtlasVerification.CallOpts, dAppOp)
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

// GetGovFromControl is a free data retrieval call binding the contract method 0xa55660da.
//
// Solidity: function getGovFromControl(address control) view returns(address)
func (_AtlasVerification *AtlasVerificationCaller) GetGovFromControl(opts *bind.CallOpts, control common.Address) (common.Address, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getGovFromControl", control)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGovFromControl is a free data retrieval call binding the contract method 0xa55660da.
//
// Solidity: function getGovFromControl(address control) view returns(address)
func (_AtlasVerification *AtlasVerificationSession) GetGovFromControl(control common.Address) (common.Address, error) {
	return _AtlasVerification.Contract.GetGovFromControl(&_AtlasVerification.CallOpts, control)
}

// GetGovFromControl is a free data retrieval call binding the contract method 0xa55660da.
//
// Solidity: function getGovFromControl(address control) view returns(address)
func (_AtlasVerification *AtlasVerificationCallerSession) GetGovFromControl(control common.Address) (common.Address, error) {
	return _AtlasVerification.Contract.GetGovFromControl(&_AtlasVerification.CallOpts, control)
}

// GetSolverPayload is a free data retrieval call binding the contract method 0xc5883fc1.
//
// Solidity: function getSolverPayload((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationCaller) GetSolverPayload(opts *bind.CallOpts, solverOp SolverOperation) ([32]byte, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getSolverPayload", solverOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSolverPayload is a free data retrieval call binding the contract method 0xc5883fc1.
//
// Solidity: function getSolverPayload((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationSession) GetSolverPayload(solverOp SolverOperation) ([32]byte, error) {
	return _AtlasVerification.Contract.GetSolverPayload(&_AtlasVerification.CallOpts, solverOp)
}

// GetSolverPayload is a free data retrieval call binding the contract method 0xc5883fc1.
//
// Solidity: function getSolverPayload((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationCallerSession) GetSolverPayload(solverOp SolverOperation) ([32]byte, error) {
	return _AtlasVerification.Contract.GetSolverPayload(&_AtlasVerification.CallOpts, solverOp)
}

// GetUserNextNonSeqNonceAfter is a free data retrieval call binding the contract method 0x34a99a03.
//
// Solidity: function getUserNextNonSeqNonceAfter(address user, uint256 refNonce) view returns(uint256)
func (_AtlasVerification *AtlasVerificationCaller) GetUserNextNonSeqNonceAfter(opts *bind.CallOpts, user common.Address, refNonce *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getUserNextNonSeqNonceAfter", user, refNonce)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserNextNonSeqNonceAfter is a free data retrieval call binding the contract method 0x34a99a03.
//
// Solidity: function getUserNextNonSeqNonceAfter(address user, uint256 refNonce) view returns(uint256)
func (_AtlasVerification *AtlasVerificationSession) GetUserNextNonSeqNonceAfter(user common.Address, refNonce *big.Int) (*big.Int, error) {
	return _AtlasVerification.Contract.GetUserNextNonSeqNonceAfter(&_AtlasVerification.CallOpts, user, refNonce)
}

// GetUserNextNonSeqNonceAfter is a free data retrieval call binding the contract method 0x34a99a03.
//
// Solidity: function getUserNextNonSeqNonceAfter(address user, uint256 refNonce) view returns(uint256)
func (_AtlasVerification *AtlasVerificationCallerSession) GetUserNextNonSeqNonceAfter(user common.Address, refNonce *big.Int) (*big.Int, error) {
	return _AtlasVerification.Contract.GetUserNextNonSeqNonceAfter(&_AtlasVerification.CallOpts, user, refNonce)
}

// GetUserNextNonce is a free data retrieval call binding the contract method 0x0c388878.
//
// Solidity: function getUserNextNonce(address user, bool sequential) view returns(uint256 nextNonce)
func (_AtlasVerification *AtlasVerificationCaller) GetUserNextNonce(opts *bind.CallOpts, user common.Address, sequential bool) (*big.Int, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getUserNextNonce", user, sequential)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserNextNonce is a free data retrieval call binding the contract method 0x0c388878.
//
// Solidity: function getUserNextNonce(address user, bool sequential) view returns(uint256 nextNonce)
func (_AtlasVerification *AtlasVerificationSession) GetUserNextNonce(user common.Address, sequential bool) (*big.Int, error) {
	return _AtlasVerification.Contract.GetUserNextNonce(&_AtlasVerification.CallOpts, user, sequential)
}

// GetUserNextNonce is a free data retrieval call binding the contract method 0x0c388878.
//
// Solidity: function getUserNextNonce(address user, bool sequential) view returns(uint256 nextNonce)
func (_AtlasVerification *AtlasVerificationCallerSession) GetUserNextNonce(user common.Address, sequential bool) (*big.Int, error) {
	return _AtlasVerification.Contract.GetUserNextNonce(&_AtlasVerification.CallOpts, user, sequential)
}

// GetUserOperationHash is a free data retrieval call binding the contract method 0x488c4b19.
//
// Solidity: function getUserOperationHash((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) view returns(bytes32 userOpHash)
func (_AtlasVerification *AtlasVerificationCaller) GetUserOperationHash(opts *bind.CallOpts, userOp UserOperation) ([32]byte, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getUserOperationHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOperationHash is a free data retrieval call binding the contract method 0x488c4b19.
//
// Solidity: function getUserOperationHash((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) view returns(bytes32 userOpHash)
func (_AtlasVerification *AtlasVerificationSession) GetUserOperationHash(userOp UserOperation) ([32]byte, error) {
	return _AtlasVerification.Contract.GetUserOperationHash(&_AtlasVerification.CallOpts, userOp)
}

// GetUserOperationHash is a free data retrieval call binding the contract method 0x488c4b19.
//
// Solidity: function getUserOperationHash((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) view returns(bytes32 userOpHash)
func (_AtlasVerification *AtlasVerificationCallerSession) GetUserOperationHash(userOp UserOperation) ([32]byte, error) {
	return _AtlasVerification.Contract.GetUserOperationHash(&_AtlasVerification.CallOpts, userOp)
}

// GetUserOperationPayload is a free data retrieval call binding the contract method 0x8b28829e.
//
// Solidity: function getUserOperationPayload((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationCaller) GetUserOperationPayload(opts *bind.CallOpts, userOp UserOperation) ([32]byte, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getUserOperationPayload", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOperationPayload is a free data retrieval call binding the contract method 0x8b28829e.
//
// Solidity: function getUserOperationPayload((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationSession) GetUserOperationPayload(userOp UserOperation) ([32]byte, error) {
	return _AtlasVerification.Contract.GetUserOperationPayload(&_AtlasVerification.CallOpts, userOp)
}

// GetUserOperationPayload is a free data retrieval call binding the contract method 0x8b28829e.
//
// Solidity: function getUserOperationPayload((address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationCallerSession) GetUserOperationPayload(userOp UserOperation) ([32]byte, error) {
	return _AtlasVerification.Contract.GetUserOperationPayload(&_AtlasVerification.CallOpts, userOp)
}

// IsDAppSignatory is a free data retrieval call binding the contract method 0x091d7b96.
//
// Solidity: function isDAppSignatory(address control, address signatory) view returns(bool)
func (_AtlasVerification *AtlasVerificationCaller) IsDAppSignatory(opts *bind.CallOpts, control common.Address, signatory common.Address) (bool, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "isDAppSignatory", control, signatory)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDAppSignatory is a free data retrieval call binding the contract method 0x091d7b96.
//
// Solidity: function isDAppSignatory(address control, address signatory) view returns(bool)
func (_AtlasVerification *AtlasVerificationSession) IsDAppSignatory(control common.Address, signatory common.Address) (bool, error) {
	return _AtlasVerification.Contract.IsDAppSignatory(&_AtlasVerification.CallOpts, control, signatory)
}

// IsDAppSignatory is a free data retrieval call binding the contract method 0x091d7b96.
//
// Solidity: function isDAppSignatory(address control, address signatory) view returns(bool)
func (_AtlasVerification *AtlasVerificationCallerSession) IsDAppSignatory(control common.Address, signatory common.Address) (bool, error) {
	return _AtlasVerification.Contract.IsDAppSignatory(&_AtlasVerification.CallOpts, control, signatory)
}

// Signatories is a free data retrieval call binding the contract method 0xbfcf51ec.
//
// Solidity: function signatories(bytes32 key) view returns(bool)
func (_AtlasVerification *AtlasVerificationCaller) Signatories(opts *bind.CallOpts, key [32]byte) (bool, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "signatories", key)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Signatories is a free data retrieval call binding the contract method 0xbfcf51ec.
//
// Solidity: function signatories(bytes32 key) view returns(bool)
func (_AtlasVerification *AtlasVerificationSession) Signatories(key [32]byte) (bool, error) {
	return _AtlasVerification.Contract.Signatories(&_AtlasVerification.CallOpts, key)
}

// Signatories is a free data retrieval call binding the contract method 0xbfcf51ec.
//
// Solidity: function signatories(bytes32 key) view returns(bool)
func (_AtlasVerification *AtlasVerificationCallerSession) Signatories(key [32]byte) (bool, error) {
	return _AtlasVerification.Contract.Signatories(&_AtlasVerification.CallOpts, key)
}

// UserNonSequentialNonceTrackers is a free data retrieval call binding the contract method 0x8caf11aa.
//
// Solidity: function userNonSequentialNonceTrackers(address account, uint248 wordIndex) view returns(uint256 bitmap)
func (_AtlasVerification *AtlasVerificationCaller) UserNonSequentialNonceTrackers(opts *bind.CallOpts, account common.Address, wordIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "userNonSequentialNonceTrackers", account, wordIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserNonSequentialNonceTrackers is a free data retrieval call binding the contract method 0x8caf11aa.
//
// Solidity: function userNonSequentialNonceTrackers(address account, uint248 wordIndex) view returns(uint256 bitmap)
func (_AtlasVerification *AtlasVerificationSession) UserNonSequentialNonceTrackers(account common.Address, wordIndex *big.Int) (*big.Int, error) {
	return _AtlasVerification.Contract.UserNonSequentialNonceTrackers(&_AtlasVerification.CallOpts, account, wordIndex)
}

// UserNonSequentialNonceTrackers is a free data retrieval call binding the contract method 0x8caf11aa.
//
// Solidity: function userNonSequentialNonceTrackers(address account, uint248 wordIndex) view returns(uint256 bitmap)
func (_AtlasVerification *AtlasVerificationCallerSession) UserNonSequentialNonceTrackers(account common.Address, wordIndex *big.Int) (*big.Int, error) {
	return _AtlasVerification.Contract.UserNonSequentialNonceTrackers(&_AtlasVerification.CallOpts, account, wordIndex)
}

// UserSequentialNonceTrackers is a free data retrieval call binding the contract method 0x7d7303d0.
//
// Solidity: function userSequentialNonceTrackers(address account) view returns(uint256 lastUsedSeqNonce)
func (_AtlasVerification *AtlasVerificationCaller) UserSequentialNonceTrackers(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "userSequentialNonceTrackers", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserSequentialNonceTrackers is a free data retrieval call binding the contract method 0x7d7303d0.
//
// Solidity: function userSequentialNonceTrackers(address account) view returns(uint256 lastUsedSeqNonce)
func (_AtlasVerification *AtlasVerificationSession) UserSequentialNonceTrackers(account common.Address) (*big.Int, error) {
	return _AtlasVerification.Contract.UserSequentialNonceTrackers(&_AtlasVerification.CallOpts, account)
}

// UserSequentialNonceTrackers is a free data retrieval call binding the contract method 0x7d7303d0.
//
// Solidity: function userSequentialNonceTrackers(address account) view returns(uint256 lastUsedSeqNonce)
func (_AtlasVerification *AtlasVerificationCallerSession) UserSequentialNonceTrackers(account common.Address) (*big.Int, error) {
	return _AtlasVerification.Contract.UserSequentialNonceTrackers(&_AtlasVerification.CallOpts, account)
}

// VerifySolverOp is a free data retrieval call binding the contract method 0x9f7e72b6.
//
// Solidity: function verifySolverOp((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes32 userOpHash, uint256 userMaxFeePerGas, address bundler, bool allowsTrustedOpHash) view returns(uint256 result)
func (_AtlasVerification *AtlasVerificationCaller) VerifySolverOp(opts *bind.CallOpts, solverOp SolverOperation, userOpHash [32]byte, userMaxFeePerGas *big.Int, bundler common.Address, allowsTrustedOpHash bool) (*big.Int, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "verifySolverOp", solverOp, userOpHash, userMaxFeePerGas, bundler, allowsTrustedOpHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VerifySolverOp is a free data retrieval call binding the contract method 0x9f7e72b6.
//
// Solidity: function verifySolverOp((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes32 userOpHash, uint256 userMaxFeePerGas, address bundler, bool allowsTrustedOpHash) view returns(uint256 result)
func (_AtlasVerification *AtlasVerificationSession) VerifySolverOp(solverOp SolverOperation, userOpHash [32]byte, userMaxFeePerGas *big.Int, bundler common.Address, allowsTrustedOpHash bool) (*big.Int, error) {
	return _AtlasVerification.Contract.VerifySolverOp(&_AtlasVerification.CallOpts, solverOp, userOpHash, userMaxFeePerGas, bundler, allowsTrustedOpHash)
}

// VerifySolverOp is a free data retrieval call binding the contract method 0x9f7e72b6.
//
// Solidity: function verifySolverOp((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes32 userOpHash, uint256 userMaxFeePerGas, address bundler, bool allowsTrustedOpHash) view returns(uint256 result)
func (_AtlasVerification *AtlasVerificationCallerSession) VerifySolverOp(solverOp SolverOperation, userOpHash [32]byte, userMaxFeePerGas *big.Int, bundler common.Address, allowsTrustedOpHash bool) (*big.Int, error) {
	return _AtlasVerification.Contract.VerifySolverOp(&_AtlasVerification.CallOpts, solverOp, userOpHash, userMaxFeePerGas, bundler, allowsTrustedOpHash)
}

// AddSignatory is a paid mutator transaction binding the contract method 0x1170a503.
//
// Solidity: function addSignatory(address control, address signatory) returns()
func (_AtlasVerification *AtlasVerificationTransactor) AddSignatory(opts *bind.TransactOpts, control common.Address, signatory common.Address) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "addSignatory", control, signatory)
}

// AddSignatory is a paid mutator transaction binding the contract method 0x1170a503.
//
// Solidity: function addSignatory(address control, address signatory) returns()
func (_AtlasVerification *AtlasVerificationSession) AddSignatory(control common.Address, signatory common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.AddSignatory(&_AtlasVerification.TransactOpts, control, signatory)
}

// AddSignatory is a paid mutator transaction binding the contract method 0x1170a503.
//
// Solidity: function addSignatory(address control, address signatory) returns()
func (_AtlasVerification *AtlasVerificationTransactorSession) AddSignatory(control common.Address, signatory common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.AddSignatory(&_AtlasVerification.TransactOpts, control, signatory)
}

// ChangeDAppGovernance is a paid mutator transaction binding the contract method 0x1364147e.
//
// Solidity: function changeDAppGovernance(address oldGovernance, address newGovernance) returns()
func (_AtlasVerification *AtlasVerificationTransactor) ChangeDAppGovernance(opts *bind.TransactOpts, oldGovernance common.Address, newGovernance common.Address) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "changeDAppGovernance", oldGovernance, newGovernance)
}

// ChangeDAppGovernance is a paid mutator transaction binding the contract method 0x1364147e.
//
// Solidity: function changeDAppGovernance(address oldGovernance, address newGovernance) returns()
func (_AtlasVerification *AtlasVerificationSession) ChangeDAppGovernance(oldGovernance common.Address, newGovernance common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.ChangeDAppGovernance(&_AtlasVerification.TransactOpts, oldGovernance, newGovernance)
}

// ChangeDAppGovernance is a paid mutator transaction binding the contract method 0x1364147e.
//
// Solidity: function changeDAppGovernance(address oldGovernance, address newGovernance) returns()
func (_AtlasVerification *AtlasVerificationTransactorSession) ChangeDAppGovernance(oldGovernance common.Address, newGovernance common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.ChangeDAppGovernance(&_AtlasVerification.TransactOpts, oldGovernance, newGovernance)
}

// DisableDApp is a paid mutator transaction binding the contract method 0x9bcf79b7.
//
// Solidity: function disableDApp(address control) returns()
func (_AtlasVerification *AtlasVerificationTransactor) DisableDApp(opts *bind.TransactOpts, control common.Address) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "disableDApp", control)
}

// DisableDApp is a paid mutator transaction binding the contract method 0x9bcf79b7.
//
// Solidity: function disableDApp(address control) returns()
func (_AtlasVerification *AtlasVerificationSession) DisableDApp(control common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.DisableDApp(&_AtlasVerification.TransactOpts, control)
}

// DisableDApp is a paid mutator transaction binding the contract method 0x9bcf79b7.
//
// Solidity: function disableDApp(address control) returns()
func (_AtlasVerification *AtlasVerificationTransactorSession) DisableDApp(control common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.DisableDApp(&_AtlasVerification.TransactOpts, control)
}

// InitializeGovernance is a paid mutator transaction binding the contract method 0x55d202a6.
//
// Solidity: function initializeGovernance(address control) returns()
func (_AtlasVerification *AtlasVerificationTransactor) InitializeGovernance(opts *bind.TransactOpts, control common.Address) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "initializeGovernance", control)
}

// InitializeGovernance is a paid mutator transaction binding the contract method 0x55d202a6.
//
// Solidity: function initializeGovernance(address control) returns()
func (_AtlasVerification *AtlasVerificationSession) InitializeGovernance(control common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.InitializeGovernance(&_AtlasVerification.TransactOpts, control)
}

// InitializeGovernance is a paid mutator transaction binding the contract method 0x55d202a6.
//
// Solidity: function initializeGovernance(address control) returns()
func (_AtlasVerification *AtlasVerificationTransactorSession) InitializeGovernance(control common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.InitializeGovernance(&_AtlasVerification.TransactOpts, control)
}

// RemoveSignatory is a paid mutator transaction binding the contract method 0xe2e439ea.
//
// Solidity: function removeSignatory(address control, address signatory) returns()
func (_AtlasVerification *AtlasVerificationTransactor) RemoveSignatory(opts *bind.TransactOpts, control common.Address, signatory common.Address) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "removeSignatory", control, signatory)
}

// RemoveSignatory is a paid mutator transaction binding the contract method 0xe2e439ea.
//
// Solidity: function removeSignatory(address control, address signatory) returns()
func (_AtlasVerification *AtlasVerificationSession) RemoveSignatory(control common.Address, signatory common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.RemoveSignatory(&_AtlasVerification.TransactOpts, control, signatory)
}

// RemoveSignatory is a paid mutator transaction binding the contract method 0xe2e439ea.
//
// Solidity: function removeSignatory(address control, address signatory) returns()
func (_AtlasVerification *AtlasVerificationTransactorSession) RemoveSignatory(control common.Address, signatory common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.RemoveSignatory(&_AtlasVerification.TransactOpts, control, signatory)
}

// ValidateCalls is a paid mutator transaction binding the contract method 0xab0a4b96.
//
// Solidity: function validateCalls((address,uint32,address,uint32) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp, uint256 msgValue, address msgSender, bool isSimulation) returns(uint8)
func (_AtlasVerification *AtlasVerificationTransactor) ValidateCalls(opts *bind.TransactOpts, dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation, msgValue *big.Int, msgSender common.Address, isSimulation bool) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "validateCalls", dConfig, userOp, solverOps, dAppOp, msgValue, msgSender, isSimulation)
}

// ValidateCalls is a paid mutator transaction binding the contract method 0xab0a4b96.
//
// Solidity: function validateCalls((address,uint32,address,uint32) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp, uint256 msgValue, address msgSender, bool isSimulation) returns(uint8)
func (_AtlasVerification *AtlasVerificationSession) ValidateCalls(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation, msgValue *big.Int, msgSender common.Address, isSimulation bool) (*types.Transaction, error) {
	return _AtlasVerification.Contract.ValidateCalls(&_AtlasVerification.TransactOpts, dConfig, userOp, solverOps, dAppOp, msgValue, msgSender, isSimulation)
}

// ValidateCalls is a paid mutator transaction binding the contract method 0xab0a4b96.
//
// Solidity: function validateCalls((address,uint32,address,uint32) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,uint32,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp, uint256 msgValue, address msgSender, bool isSimulation) returns(uint8)
func (_AtlasVerification *AtlasVerificationTransactorSession) ValidateCalls(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation, msgValue *big.Int, msgSender common.Address, isSimulation bool) (*types.Transaction, error) {
	return _AtlasVerification.Contract.ValidateCalls(&_AtlasVerification.TransactOpts, dConfig, userOp, solverOps, dAppOp, msgValue, msgSender, isSimulation)
}

// AtlasVerificationDAppDisabledIterator is returned from FilterDAppDisabled and is used to iterate over the raw logs and unpacked data for DAppDisabled events raised by the AtlasVerification contract.
type AtlasVerificationDAppDisabledIterator struct {
	Event *AtlasVerificationDAppDisabled // Event containing the contract specifics and raw log

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
func (it *AtlasVerificationDAppDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasVerificationDAppDisabled)
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
		it.Event = new(AtlasVerificationDAppDisabled)
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
func (it *AtlasVerificationDAppDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasVerificationDAppDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasVerificationDAppDisabled represents a DAppDisabled event raised by the AtlasVerification contract.
type AtlasVerificationDAppDisabled struct {
	Control    common.Address
	Governance common.Address
	CallConfig uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDAppDisabled is a free log retrieval operation binding the contract event 0xfb402de0284ae6437f381fcd6b8da6639e155a56551b5a8b95d2ab6c4e007d0a.
//
// Solidity: event DAppDisabled(address indexed control, address indexed governance, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) FilterDAppDisabled(opts *bind.FilterOpts, control []common.Address, governance []common.Address) (*AtlasVerificationDAppDisabledIterator, error) {

	var controlRule []interface{}
	for _, controlItem := range control {
		controlRule = append(controlRule, controlItem)
	}
	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}

	logs, sub, err := _AtlasVerification.contract.FilterLogs(opts, "DAppDisabled", controlRule, governanceRule)
	if err != nil {
		return nil, err
	}
	return &AtlasVerificationDAppDisabledIterator{contract: _AtlasVerification.contract, event: "DAppDisabled", logs: logs, sub: sub}, nil
}

// WatchDAppDisabled is a free log subscription operation binding the contract event 0xfb402de0284ae6437f381fcd6b8da6639e155a56551b5a8b95d2ab6c4e007d0a.
//
// Solidity: event DAppDisabled(address indexed control, address indexed governance, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) WatchDAppDisabled(opts *bind.WatchOpts, sink chan<- *AtlasVerificationDAppDisabled, control []common.Address, governance []common.Address) (event.Subscription, error) {

	var controlRule []interface{}
	for _, controlItem := range control {
		controlRule = append(controlRule, controlItem)
	}
	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}

	logs, sub, err := _AtlasVerification.contract.WatchLogs(opts, "DAppDisabled", controlRule, governanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasVerificationDAppDisabled)
				if err := _AtlasVerification.contract.UnpackLog(event, "DAppDisabled", log); err != nil {
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
func (_AtlasVerification *AtlasVerificationFilterer) ParseDAppDisabled(log types.Log) (*AtlasVerificationDAppDisabled, error) {
	event := new(AtlasVerificationDAppDisabled)
	if err := _AtlasVerification.contract.UnpackLog(event, "DAppDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasVerificationDAppGovernanceChangedIterator is returned from FilterDAppGovernanceChanged and is used to iterate over the raw logs and unpacked data for DAppGovernanceChanged events raised by the AtlasVerification contract.
type AtlasVerificationDAppGovernanceChangedIterator struct {
	Event *AtlasVerificationDAppGovernanceChanged // Event containing the contract specifics and raw log

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
func (it *AtlasVerificationDAppGovernanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasVerificationDAppGovernanceChanged)
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
		it.Event = new(AtlasVerificationDAppGovernanceChanged)
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
func (it *AtlasVerificationDAppGovernanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasVerificationDAppGovernanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasVerificationDAppGovernanceChanged represents a DAppGovernanceChanged event raised by the AtlasVerification contract.
type AtlasVerificationDAppGovernanceChanged struct {
	Control       common.Address
	OldGovernance common.Address
	NewGovernance common.Address
	CallConfig    uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDAppGovernanceChanged is a free log retrieval operation binding the contract event 0xcc02b55b78457369d35dc189a074d3fc4f96a1897b405c0f1c57789f782d8bd2.
//
// Solidity: event DAppGovernanceChanged(address indexed control, address indexed oldGovernance, address indexed newGovernance, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) FilterDAppGovernanceChanged(opts *bind.FilterOpts, control []common.Address, oldGovernance []common.Address, newGovernance []common.Address) (*AtlasVerificationDAppGovernanceChangedIterator, error) {

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

	logs, sub, err := _AtlasVerification.contract.FilterLogs(opts, "DAppGovernanceChanged", controlRule, oldGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &AtlasVerificationDAppGovernanceChangedIterator{contract: _AtlasVerification.contract, event: "DAppGovernanceChanged", logs: logs, sub: sub}, nil
}

// WatchDAppGovernanceChanged is a free log subscription operation binding the contract event 0xcc02b55b78457369d35dc189a074d3fc4f96a1897b405c0f1c57789f782d8bd2.
//
// Solidity: event DAppGovernanceChanged(address indexed control, address indexed oldGovernance, address indexed newGovernance, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) WatchDAppGovernanceChanged(opts *bind.WatchOpts, sink chan<- *AtlasVerificationDAppGovernanceChanged, control []common.Address, oldGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AtlasVerification.contract.WatchLogs(opts, "DAppGovernanceChanged", controlRule, oldGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasVerificationDAppGovernanceChanged)
				if err := _AtlasVerification.contract.UnpackLog(event, "DAppGovernanceChanged", log); err != nil {
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
func (_AtlasVerification *AtlasVerificationFilterer) ParseDAppGovernanceChanged(log types.Log) (*AtlasVerificationDAppGovernanceChanged, error) {
	event := new(AtlasVerificationDAppGovernanceChanged)
	if err := _AtlasVerification.contract.UnpackLog(event, "DAppGovernanceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasVerificationNewDAppSignatoryIterator is returned from FilterNewDAppSignatory and is used to iterate over the raw logs and unpacked data for NewDAppSignatory events raised by the AtlasVerification contract.
type AtlasVerificationNewDAppSignatoryIterator struct {
	Event *AtlasVerificationNewDAppSignatory // Event containing the contract specifics and raw log

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
func (it *AtlasVerificationNewDAppSignatoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasVerificationNewDAppSignatory)
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
		it.Event = new(AtlasVerificationNewDAppSignatory)
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
func (it *AtlasVerificationNewDAppSignatoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasVerificationNewDAppSignatoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasVerificationNewDAppSignatory represents a NewDAppSignatory event raised by the AtlasVerification contract.
type AtlasVerificationNewDAppSignatory struct {
	Control    common.Address
	Governance common.Address
	Signatory  common.Address
	CallConfig uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewDAppSignatory is a free log retrieval operation binding the contract event 0xcacb2beeeba676b69bddabab0d5f66e2733cabc804f82afd92ea5beae71934b4.
//
// Solidity: event NewDAppSignatory(address indexed control, address indexed governance, address indexed signatory, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) FilterNewDAppSignatory(opts *bind.FilterOpts, control []common.Address, governance []common.Address, signatory []common.Address) (*AtlasVerificationNewDAppSignatoryIterator, error) {

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

	logs, sub, err := _AtlasVerification.contract.FilterLogs(opts, "NewDAppSignatory", controlRule, governanceRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return &AtlasVerificationNewDAppSignatoryIterator{contract: _AtlasVerification.contract, event: "NewDAppSignatory", logs: logs, sub: sub}, nil
}

// WatchNewDAppSignatory is a free log subscription operation binding the contract event 0xcacb2beeeba676b69bddabab0d5f66e2733cabc804f82afd92ea5beae71934b4.
//
// Solidity: event NewDAppSignatory(address indexed control, address indexed governance, address indexed signatory, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) WatchNewDAppSignatory(opts *bind.WatchOpts, sink chan<- *AtlasVerificationNewDAppSignatory, control []common.Address, governance []common.Address, signatory []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AtlasVerification.contract.WatchLogs(opts, "NewDAppSignatory", controlRule, governanceRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasVerificationNewDAppSignatory)
				if err := _AtlasVerification.contract.UnpackLog(event, "NewDAppSignatory", log); err != nil {
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
func (_AtlasVerification *AtlasVerificationFilterer) ParseNewDAppSignatory(log types.Log) (*AtlasVerificationNewDAppSignatory, error) {
	event := new(AtlasVerificationNewDAppSignatory)
	if err := _AtlasVerification.contract.UnpackLog(event, "NewDAppSignatory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasVerificationRemovedDAppSignatoryIterator is returned from FilterRemovedDAppSignatory and is used to iterate over the raw logs and unpacked data for RemovedDAppSignatory events raised by the AtlasVerification contract.
type AtlasVerificationRemovedDAppSignatoryIterator struct {
	Event *AtlasVerificationRemovedDAppSignatory // Event containing the contract specifics and raw log

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
func (it *AtlasVerificationRemovedDAppSignatoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasVerificationRemovedDAppSignatory)
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
		it.Event = new(AtlasVerificationRemovedDAppSignatory)
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
func (it *AtlasVerificationRemovedDAppSignatoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasVerificationRemovedDAppSignatoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasVerificationRemovedDAppSignatory represents a RemovedDAppSignatory event raised by the AtlasVerification contract.
type AtlasVerificationRemovedDAppSignatory struct {
	Control    common.Address
	Governance common.Address
	Signatory  common.Address
	CallConfig uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemovedDAppSignatory is a free log retrieval operation binding the contract event 0x90a400fe7fe676a24908d5f730fe6d38072340bc8f74e18a543132d840bfe5ce.
//
// Solidity: event RemovedDAppSignatory(address indexed control, address indexed governance, address indexed signatory, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) FilterRemovedDAppSignatory(opts *bind.FilterOpts, control []common.Address, governance []common.Address, signatory []common.Address) (*AtlasVerificationRemovedDAppSignatoryIterator, error) {

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

	logs, sub, err := _AtlasVerification.contract.FilterLogs(opts, "RemovedDAppSignatory", controlRule, governanceRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return &AtlasVerificationRemovedDAppSignatoryIterator{contract: _AtlasVerification.contract, event: "RemovedDAppSignatory", logs: logs, sub: sub}, nil
}

// WatchRemovedDAppSignatory is a free log subscription operation binding the contract event 0x90a400fe7fe676a24908d5f730fe6d38072340bc8f74e18a543132d840bfe5ce.
//
// Solidity: event RemovedDAppSignatory(address indexed control, address indexed governance, address indexed signatory, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) WatchRemovedDAppSignatory(opts *bind.WatchOpts, sink chan<- *AtlasVerificationRemovedDAppSignatory, control []common.Address, governance []common.Address, signatory []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AtlasVerification.contract.WatchLogs(opts, "RemovedDAppSignatory", controlRule, governanceRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasVerificationRemovedDAppSignatory)
				if err := _AtlasVerification.contract.UnpackLog(event, "RemovedDAppSignatory", log); err != nil {
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
func (_AtlasVerification *AtlasVerificationFilterer) ParseRemovedDAppSignatory(log types.Log) (*AtlasVerificationRemovedDAppSignatory, error) {
	event := new(AtlasVerificationRemovedDAppSignatory)
	if err := _AtlasVerification.contract.UnpackLog(event, "RemovedDAppSignatory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
