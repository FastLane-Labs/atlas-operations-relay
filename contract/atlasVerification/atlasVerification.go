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
	Value         *big.Int
	Gas           *big.Int
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

// AtlasVerificationMetaData contains all meta data concerning the AtlasVerification contract.
var AtlasVerificationMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_atlas\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ATLAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addSignatory\",\"inputs\":[{\"name\":\"controller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signatory\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeDAppGovernance\",\"inputs\":[{\"name\":\"oldGovernance\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newGovernance\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dAppSignatories\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disableDApp\",\"inputs\":[{\"name\":\"dAppControl\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDAppOperationPayload\",\"inputs\":[{\"name\":\"dAppOp\",\"type\":\"tuple\",\"internalType\":\"structDAppOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bundler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callChainHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"payload\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDAppSignatories\",\"inputs\":[{\"name\":\"dAppControl\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDomainSeparator\",\"inputs\":[],\"outputs\":[{\"name\":\"domainSeparator\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGovFromControl\",\"inputs\":[{\"name\":\"dAppControl\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNextNonce\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sequenced\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSolverPayload\",\"inputs\":[{\"name\":\"solverOp\",\"type\":\"tuple\",\"internalType\":\"structSolverOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"payload\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUserOperationPayload\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structUserOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"payload\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initializeGovernance\",\"inputs\":[{\"name\":\"controller\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDAppSignatory\",\"inputs\":[{\"name\":\"dAppControl\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signatory\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"manuallyUpdateNonceTracker\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nonceBitmaps\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"highestUsedNonce\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"bitmap\",\"type\":\"uint240\",\"internalType\":\"uint240\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonceTrackers\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"lastUsedSeqNonce\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"highestFullAsyncBitmap\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeSignatory\",\"inputs\":[{\"name\":\"controller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signatory\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"signatories\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateCalls\",\"inputs\":[{\"name\":\"dConfig\",\"type\":\"tuple\",\"internalType\":\"structDAppConfig\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solverGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structUserOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"solverOps\",\"type\":\"tuple[]\",\"internalType\":\"structSolverOperation[]\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"dAppOp\",\"type\":\"tuple\",\"internalType\":\"structDAppOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bundler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callChainHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"msgSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isSimulation\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumValidCallsResult\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifySolverOp\",\"inputs\":[{\"name\":\"solverOp\",\"type\":\"tuple\",\"internalType\":\"structSolverOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"userMaxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bundler\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DAppDisabled\",\"inputs\":[{\"name\":\"controller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"governance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DAppGovernanceChanged\",\"inputs\":[{\"name\":\"controller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"oldGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newGovernance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NewDAppSignatory\",\"inputs\":[{\"name\":\"controller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"governance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"signatory\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RemovedDAppSignatory\",\"inputs\":[{\"name\":\"controller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"governance\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"signatory\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AtlasLockActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DAppNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAccount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyGovernance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatoryActive\",\"inputs\":[]}]",
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

// DAppSignatories is a free data retrieval call binding the contract method 0x0d001b37.
//
// Solidity: function dAppSignatories(address , uint256 ) view returns(address)
func (_AtlasVerification *AtlasVerificationCaller) DAppSignatories(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "dAppSignatories", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DAppSignatories is a free data retrieval call binding the contract method 0x0d001b37.
//
// Solidity: function dAppSignatories(address , uint256 ) view returns(address)
func (_AtlasVerification *AtlasVerificationSession) DAppSignatories(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _AtlasVerification.Contract.DAppSignatories(&_AtlasVerification.CallOpts, arg0, arg1)
}

// DAppSignatories is a free data retrieval call binding the contract method 0x0d001b37.
//
// Solidity: function dAppSignatories(address , uint256 ) view returns(address)
func (_AtlasVerification *AtlasVerificationCallerSession) DAppSignatories(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _AtlasVerification.Contract.DAppSignatories(&_AtlasVerification.CallOpts, arg0, arg1)
}

// GetDAppOperationPayload is a free data retrieval call binding the contract method 0x6d8c7b04.
//
// Solidity: function getDAppOperationPayload((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationCaller) GetDAppOperationPayload(opts *bind.CallOpts, dAppOp DAppOperation) ([32]byte, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getDAppOperationPayload", dAppOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDAppOperationPayload is a free data retrieval call binding the contract method 0x6d8c7b04.
//
// Solidity: function getDAppOperationPayload((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationSession) GetDAppOperationPayload(dAppOp DAppOperation) ([32]byte, error) {
	return _AtlasVerification.Contract.GetDAppOperationPayload(&_AtlasVerification.CallOpts, dAppOp)
}

// GetDAppOperationPayload is a free data retrieval call binding the contract method 0x6d8c7b04.
//
// Solidity: function getDAppOperationPayload((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationCallerSession) GetDAppOperationPayload(dAppOp DAppOperation) ([32]byte, error) {
	return _AtlasVerification.Contract.GetDAppOperationPayload(&_AtlasVerification.CallOpts, dAppOp)
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

// GetGovFromControl is a free data retrieval call binding the contract method 0xa55660da.
//
// Solidity: function getGovFromControl(address dAppControl) view returns(address)
func (_AtlasVerification *AtlasVerificationCaller) GetGovFromControl(opts *bind.CallOpts, dAppControl common.Address) (common.Address, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getGovFromControl", dAppControl)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGovFromControl is a free data retrieval call binding the contract method 0xa55660da.
//
// Solidity: function getGovFromControl(address dAppControl) view returns(address)
func (_AtlasVerification *AtlasVerificationSession) GetGovFromControl(dAppControl common.Address) (common.Address, error) {
	return _AtlasVerification.Contract.GetGovFromControl(&_AtlasVerification.CallOpts, dAppControl)
}

// GetGovFromControl is a free data retrieval call binding the contract method 0xa55660da.
//
// Solidity: function getGovFromControl(address dAppControl) view returns(address)
func (_AtlasVerification *AtlasVerificationCallerSession) GetGovFromControl(dAppControl common.Address) (common.Address, error) {
	return _AtlasVerification.Contract.GetGovFromControl(&_AtlasVerification.CallOpts, dAppControl)
}

// GetNextNonce is a free data retrieval call binding the contract method 0xf4e2a29d.
//
// Solidity: function getNextNonce(address account, bool sequenced) view returns(uint256)
func (_AtlasVerification *AtlasVerificationCaller) GetNextNonce(opts *bind.CallOpts, account common.Address, sequenced bool) (*big.Int, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getNextNonce", account, sequenced)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextNonce is a free data retrieval call binding the contract method 0xf4e2a29d.
//
// Solidity: function getNextNonce(address account, bool sequenced) view returns(uint256)
func (_AtlasVerification *AtlasVerificationSession) GetNextNonce(account common.Address, sequenced bool) (*big.Int, error) {
	return _AtlasVerification.Contract.GetNextNonce(&_AtlasVerification.CallOpts, account, sequenced)
}

// GetNextNonce is a free data retrieval call binding the contract method 0xf4e2a29d.
//
// Solidity: function getNextNonce(address account, bool sequenced) view returns(uint256)
func (_AtlasVerification *AtlasVerificationCallerSession) GetNextNonce(account common.Address, sequenced bool) (*big.Int, error) {
	return _AtlasVerification.Contract.GetNextNonce(&_AtlasVerification.CallOpts, account, sequenced)
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

// GetUserOperationPayload is a free data retrieval call binding the contract method 0xa99d599b.
//
// Solidity: function getUserOperationPayload((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationCaller) GetUserOperationPayload(opts *bind.CallOpts, userOp UserOperation) ([32]byte, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "getUserOperationPayload", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetUserOperationPayload is a free data retrieval call binding the contract method 0xa99d599b.
//
// Solidity: function getUserOperationPayload((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationSession) GetUserOperationPayload(userOp UserOperation) ([32]byte, error) {
	return _AtlasVerification.Contract.GetUserOperationPayload(&_AtlasVerification.CallOpts, userOp)
}

// GetUserOperationPayload is a free data retrieval call binding the contract method 0xa99d599b.
//
// Solidity: function getUserOperationPayload((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp) view returns(bytes32 payload)
func (_AtlasVerification *AtlasVerificationCallerSession) GetUserOperationPayload(userOp UserOperation) ([32]byte, error) {
	return _AtlasVerification.Contract.GetUserOperationPayload(&_AtlasVerification.CallOpts, userOp)
}

// IsDAppSignatory is a free data retrieval call binding the contract method 0x091d7b96.
//
// Solidity: function isDAppSignatory(address dAppControl, address signatory) view returns(bool)
func (_AtlasVerification *AtlasVerificationCaller) IsDAppSignatory(opts *bind.CallOpts, dAppControl common.Address, signatory common.Address) (bool, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "isDAppSignatory", dAppControl, signatory)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDAppSignatory is a free data retrieval call binding the contract method 0x091d7b96.
//
// Solidity: function isDAppSignatory(address dAppControl, address signatory) view returns(bool)
func (_AtlasVerification *AtlasVerificationSession) IsDAppSignatory(dAppControl common.Address, signatory common.Address) (bool, error) {
	return _AtlasVerification.Contract.IsDAppSignatory(&_AtlasVerification.CallOpts, dAppControl, signatory)
}

// IsDAppSignatory is a free data retrieval call binding the contract method 0x091d7b96.
//
// Solidity: function isDAppSignatory(address dAppControl, address signatory) view returns(bool)
func (_AtlasVerification *AtlasVerificationCallerSession) IsDAppSignatory(dAppControl common.Address, signatory common.Address) (bool, error) {
	return _AtlasVerification.Contract.IsDAppSignatory(&_AtlasVerification.CallOpts, dAppControl, signatory)
}

// NonceBitmaps is a free data retrieval call binding the contract method 0x2965a094.
//
// Solidity: function nonceBitmaps(bytes32 ) view returns(uint8 highestUsedNonce, uint240 bitmap)
func (_AtlasVerification *AtlasVerificationCaller) NonceBitmaps(opts *bind.CallOpts, arg0 [32]byte) (struct {
	HighestUsedNonce uint8
	Bitmap           *big.Int
}, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "nonceBitmaps", arg0)

	outstruct := new(struct {
		HighestUsedNonce uint8
		Bitmap           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.HighestUsedNonce = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Bitmap = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// NonceBitmaps is a free data retrieval call binding the contract method 0x2965a094.
//
// Solidity: function nonceBitmaps(bytes32 ) view returns(uint8 highestUsedNonce, uint240 bitmap)
func (_AtlasVerification *AtlasVerificationSession) NonceBitmaps(arg0 [32]byte) (struct {
	HighestUsedNonce uint8
	Bitmap           *big.Int
}, error) {
	return _AtlasVerification.Contract.NonceBitmaps(&_AtlasVerification.CallOpts, arg0)
}

// NonceBitmaps is a free data retrieval call binding the contract method 0x2965a094.
//
// Solidity: function nonceBitmaps(bytes32 ) view returns(uint8 highestUsedNonce, uint240 bitmap)
func (_AtlasVerification *AtlasVerificationCallerSession) NonceBitmaps(arg0 [32]byte) (struct {
	HighestUsedNonce uint8
	Bitmap           *big.Int
}, error) {
	return _AtlasVerification.Contract.NonceBitmaps(&_AtlasVerification.CallOpts, arg0)
}

// NonceTrackers is a free data retrieval call binding the contract method 0x4e65cbd2.
//
// Solidity: function nonceTrackers(address ) view returns(uint128 lastUsedSeqNonce, uint128 highestFullAsyncBitmap)
func (_AtlasVerification *AtlasVerificationCaller) NonceTrackers(opts *bind.CallOpts, arg0 common.Address) (struct {
	LastUsedSeqNonce       *big.Int
	HighestFullAsyncBitmap *big.Int
}, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "nonceTrackers", arg0)

	outstruct := new(struct {
		LastUsedSeqNonce       *big.Int
		HighestFullAsyncBitmap *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastUsedSeqNonce = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.HighestFullAsyncBitmap = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// NonceTrackers is a free data retrieval call binding the contract method 0x4e65cbd2.
//
// Solidity: function nonceTrackers(address ) view returns(uint128 lastUsedSeqNonce, uint128 highestFullAsyncBitmap)
func (_AtlasVerification *AtlasVerificationSession) NonceTrackers(arg0 common.Address) (struct {
	LastUsedSeqNonce       *big.Int
	HighestFullAsyncBitmap *big.Int
}, error) {
	return _AtlasVerification.Contract.NonceTrackers(&_AtlasVerification.CallOpts, arg0)
}

// NonceTrackers is a free data retrieval call binding the contract method 0x4e65cbd2.
//
// Solidity: function nonceTrackers(address ) view returns(uint128 lastUsedSeqNonce, uint128 highestFullAsyncBitmap)
func (_AtlasVerification *AtlasVerificationCallerSession) NonceTrackers(arg0 common.Address) (struct {
	LastUsedSeqNonce       *big.Int
	HighestFullAsyncBitmap *big.Int
}, error) {
	return _AtlasVerification.Contract.NonceTrackers(&_AtlasVerification.CallOpts, arg0)
}

// Signatories is a free data retrieval call binding the contract method 0xbfcf51ec.
//
// Solidity: function signatories(bytes32 ) view returns(bool)
func (_AtlasVerification *AtlasVerificationCaller) Signatories(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "signatories", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Signatories is a free data retrieval call binding the contract method 0xbfcf51ec.
//
// Solidity: function signatories(bytes32 ) view returns(bool)
func (_AtlasVerification *AtlasVerificationSession) Signatories(arg0 [32]byte) (bool, error) {
	return _AtlasVerification.Contract.Signatories(&_AtlasVerification.CallOpts, arg0)
}

// Signatories is a free data retrieval call binding the contract method 0xbfcf51ec.
//
// Solidity: function signatories(bytes32 ) view returns(bool)
func (_AtlasVerification *AtlasVerificationCallerSession) Signatories(arg0 [32]byte) (bool, error) {
	return _AtlasVerification.Contract.Signatories(&_AtlasVerification.CallOpts, arg0)
}

// VerifySolverOp is a free data retrieval call binding the contract method 0xb37b3887.
//
// Solidity: function verifySolverOp((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes32 userOpHash, uint256 userMaxFeePerGas, address bundler) view returns(uint256 result)
func (_AtlasVerification *AtlasVerificationCaller) VerifySolverOp(opts *bind.CallOpts, solverOp SolverOperation, userOpHash [32]byte, userMaxFeePerGas *big.Int, bundler common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AtlasVerification.contract.Call(opts, &out, "verifySolverOp", solverOp, userOpHash, userMaxFeePerGas, bundler)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VerifySolverOp is a free data retrieval call binding the contract method 0xb37b3887.
//
// Solidity: function verifySolverOp((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes32 userOpHash, uint256 userMaxFeePerGas, address bundler) view returns(uint256 result)
func (_AtlasVerification *AtlasVerificationSession) VerifySolverOp(solverOp SolverOperation, userOpHash [32]byte, userMaxFeePerGas *big.Int, bundler common.Address) (*big.Int, error) {
	return _AtlasVerification.Contract.VerifySolverOp(&_AtlasVerification.CallOpts, solverOp, userOpHash, userMaxFeePerGas, bundler)
}

// VerifySolverOp is a free data retrieval call binding the contract method 0xb37b3887.
//
// Solidity: function verifySolverOp((address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes) solverOp, bytes32 userOpHash, uint256 userMaxFeePerGas, address bundler) view returns(uint256 result)
func (_AtlasVerification *AtlasVerificationCallerSession) VerifySolverOp(solverOp SolverOperation, userOpHash [32]byte, userMaxFeePerGas *big.Int, bundler common.Address) (*big.Int, error) {
	return _AtlasVerification.Contract.VerifySolverOp(&_AtlasVerification.CallOpts, solverOp, userOpHash, userMaxFeePerGas, bundler)
}

// AddSignatory is a paid mutator transaction binding the contract method 0x1170a503.
//
// Solidity: function addSignatory(address controller, address signatory) returns()
func (_AtlasVerification *AtlasVerificationTransactor) AddSignatory(opts *bind.TransactOpts, controller common.Address, signatory common.Address) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "addSignatory", controller, signatory)
}

// AddSignatory is a paid mutator transaction binding the contract method 0x1170a503.
//
// Solidity: function addSignatory(address controller, address signatory) returns()
func (_AtlasVerification *AtlasVerificationSession) AddSignatory(controller common.Address, signatory common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.AddSignatory(&_AtlasVerification.TransactOpts, controller, signatory)
}

// AddSignatory is a paid mutator transaction binding the contract method 0x1170a503.
//
// Solidity: function addSignatory(address controller, address signatory) returns()
func (_AtlasVerification *AtlasVerificationTransactorSession) AddSignatory(controller common.Address, signatory common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.AddSignatory(&_AtlasVerification.TransactOpts, controller, signatory)
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
// Solidity: function disableDApp(address dAppControl) returns()
func (_AtlasVerification *AtlasVerificationTransactor) DisableDApp(opts *bind.TransactOpts, dAppControl common.Address) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "disableDApp", dAppControl)
}

// DisableDApp is a paid mutator transaction binding the contract method 0x9bcf79b7.
//
// Solidity: function disableDApp(address dAppControl) returns()
func (_AtlasVerification *AtlasVerificationSession) DisableDApp(dAppControl common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.DisableDApp(&_AtlasVerification.TransactOpts, dAppControl)
}

// DisableDApp is a paid mutator transaction binding the contract method 0x9bcf79b7.
//
// Solidity: function disableDApp(address dAppControl) returns()
func (_AtlasVerification *AtlasVerificationTransactorSession) DisableDApp(dAppControl common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.DisableDApp(&_AtlasVerification.TransactOpts, dAppControl)
}

// InitializeGovernance is a paid mutator transaction binding the contract method 0x55d202a6.
//
// Solidity: function initializeGovernance(address controller) returns()
func (_AtlasVerification *AtlasVerificationTransactor) InitializeGovernance(opts *bind.TransactOpts, controller common.Address) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "initializeGovernance", controller)
}

// InitializeGovernance is a paid mutator transaction binding the contract method 0x55d202a6.
//
// Solidity: function initializeGovernance(address controller) returns()
func (_AtlasVerification *AtlasVerificationSession) InitializeGovernance(controller common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.InitializeGovernance(&_AtlasVerification.TransactOpts, controller)
}

// InitializeGovernance is a paid mutator transaction binding the contract method 0x55d202a6.
//
// Solidity: function initializeGovernance(address controller) returns()
func (_AtlasVerification *AtlasVerificationTransactorSession) InitializeGovernance(controller common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.InitializeGovernance(&_AtlasVerification.TransactOpts, controller)
}

// ManuallyUpdateNonceTracker is a paid mutator transaction binding the contract method 0x8e9ec0dd.
//
// Solidity: function manuallyUpdateNonceTracker(address account) returns()
func (_AtlasVerification *AtlasVerificationTransactor) ManuallyUpdateNonceTracker(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "manuallyUpdateNonceTracker", account)
}

// ManuallyUpdateNonceTracker is a paid mutator transaction binding the contract method 0x8e9ec0dd.
//
// Solidity: function manuallyUpdateNonceTracker(address account) returns()
func (_AtlasVerification *AtlasVerificationSession) ManuallyUpdateNonceTracker(account common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.ManuallyUpdateNonceTracker(&_AtlasVerification.TransactOpts, account)
}

// ManuallyUpdateNonceTracker is a paid mutator transaction binding the contract method 0x8e9ec0dd.
//
// Solidity: function manuallyUpdateNonceTracker(address account) returns()
func (_AtlasVerification *AtlasVerificationTransactorSession) ManuallyUpdateNonceTracker(account common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.ManuallyUpdateNonceTracker(&_AtlasVerification.TransactOpts, account)
}

// RemoveSignatory is a paid mutator transaction binding the contract method 0xe2e439ea.
//
// Solidity: function removeSignatory(address controller, address signatory) returns()
func (_AtlasVerification *AtlasVerificationTransactor) RemoveSignatory(opts *bind.TransactOpts, controller common.Address, signatory common.Address) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "removeSignatory", controller, signatory)
}

// RemoveSignatory is a paid mutator transaction binding the contract method 0xe2e439ea.
//
// Solidity: function removeSignatory(address controller, address signatory) returns()
func (_AtlasVerification *AtlasVerificationSession) RemoveSignatory(controller common.Address, signatory common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.RemoveSignatory(&_AtlasVerification.TransactOpts, controller, signatory)
}

// RemoveSignatory is a paid mutator transaction binding the contract method 0xe2e439ea.
//
// Solidity: function removeSignatory(address controller, address signatory) returns()
func (_AtlasVerification *AtlasVerificationTransactorSession) RemoveSignatory(controller common.Address, signatory common.Address) (*types.Transaction, error) {
	return _AtlasVerification.Contract.RemoveSignatory(&_AtlasVerification.TransactOpts, controller, signatory)
}

// ValidateCalls is a paid mutator transaction binding the contract method 0x19ae91dc.
//
// Solidity: function validateCalls((address,uint32,address,uint32) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp, uint256 msgValue, address msgSender, bool isSimulation) returns(bytes32 userOpHash, uint8)
func (_AtlasVerification *AtlasVerificationTransactor) ValidateCalls(opts *bind.TransactOpts, dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation, msgValue *big.Int, msgSender common.Address, isSimulation bool) (*types.Transaction, error) {
	return _AtlasVerification.contract.Transact(opts, "validateCalls", dConfig, userOp, solverOps, dAppOp, msgValue, msgSender, isSimulation)
}

// ValidateCalls is a paid mutator transaction binding the contract method 0x19ae91dc.
//
// Solidity: function validateCalls((address,uint32,address,uint32) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp, uint256 msgValue, address msgSender, bool isSimulation) returns(bytes32 userOpHash, uint8)
func (_AtlasVerification *AtlasVerificationSession) ValidateCalls(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation, msgValue *big.Int, msgSender common.Address, isSimulation bool) (*types.Transaction, error) {
	return _AtlasVerification.Contract.ValidateCalls(&_AtlasVerification.TransactOpts, dConfig, userOp, solverOps, dAppOp, msgValue, msgSender, isSimulation)
}

// ValidateCalls is a paid mutator transaction binding the contract method 0x19ae91dc.
//
// Solidity: function validateCalls((address,uint32,address,uint32) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp, uint256 msgValue, address msgSender, bool isSimulation) returns(bytes32 userOpHash, uint8)
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
	Controller common.Address
	Governance common.Address
	CallConfig uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDAppDisabled is a free log retrieval operation binding the contract event 0xfb402de0284ae6437f381fcd6b8da6639e155a56551b5a8b95d2ab6c4e007d0a.
//
// Solidity: event DAppDisabled(address indexed controller, address indexed governance, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) FilterDAppDisabled(opts *bind.FilterOpts, controller []common.Address, governance []common.Address) (*AtlasVerificationDAppDisabledIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}
	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}

	logs, sub, err := _AtlasVerification.contract.FilterLogs(opts, "DAppDisabled", controllerRule, governanceRule)
	if err != nil {
		return nil, err
	}
	return &AtlasVerificationDAppDisabledIterator{contract: _AtlasVerification.contract, event: "DAppDisabled", logs: logs, sub: sub}, nil
}

// WatchDAppDisabled is a free log subscription operation binding the contract event 0xfb402de0284ae6437f381fcd6b8da6639e155a56551b5a8b95d2ab6c4e007d0a.
//
// Solidity: event DAppDisabled(address indexed controller, address indexed governance, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) WatchDAppDisabled(opts *bind.WatchOpts, sink chan<- *AtlasVerificationDAppDisabled, controller []common.Address, governance []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}
	var governanceRule []interface{}
	for _, governanceItem := range governance {
		governanceRule = append(governanceRule, governanceItem)
	}

	logs, sub, err := _AtlasVerification.contract.WatchLogs(opts, "DAppDisabled", controllerRule, governanceRule)
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
// Solidity: event DAppDisabled(address indexed controller, address indexed governance, uint32 callConfig)
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
	Controller    common.Address
	OldGovernance common.Address
	NewGovernance common.Address
	CallConfig    uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDAppGovernanceChanged is a free log retrieval operation binding the contract event 0xcc02b55b78457369d35dc189a074d3fc4f96a1897b405c0f1c57789f782d8bd2.
//
// Solidity: event DAppGovernanceChanged(address indexed controller, address indexed oldGovernance, address indexed newGovernance, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) FilterDAppGovernanceChanged(opts *bind.FilterOpts, controller []common.Address, oldGovernance []common.Address, newGovernance []common.Address) (*AtlasVerificationDAppGovernanceChangedIterator, error) {

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

	logs, sub, err := _AtlasVerification.contract.FilterLogs(opts, "DAppGovernanceChanged", controllerRule, oldGovernanceRule, newGovernanceRule)
	if err != nil {
		return nil, err
	}
	return &AtlasVerificationDAppGovernanceChangedIterator{contract: _AtlasVerification.contract, event: "DAppGovernanceChanged", logs: logs, sub: sub}, nil
}

// WatchDAppGovernanceChanged is a free log subscription operation binding the contract event 0xcc02b55b78457369d35dc189a074d3fc4f96a1897b405c0f1c57789f782d8bd2.
//
// Solidity: event DAppGovernanceChanged(address indexed controller, address indexed oldGovernance, address indexed newGovernance, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) WatchDAppGovernanceChanged(opts *bind.WatchOpts, sink chan<- *AtlasVerificationDAppGovernanceChanged, controller []common.Address, oldGovernance []common.Address, newGovernance []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AtlasVerification.contract.WatchLogs(opts, "DAppGovernanceChanged", controllerRule, oldGovernanceRule, newGovernanceRule)
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
// Solidity: event DAppGovernanceChanged(address indexed controller, address indexed oldGovernance, address indexed newGovernance, uint32 callConfig)
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
	Controller common.Address
	Governance common.Address
	Signatory  common.Address
	CallConfig uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewDAppSignatory is a free log retrieval operation binding the contract event 0xcacb2beeeba676b69bddabab0d5f66e2733cabc804f82afd92ea5beae71934b4.
//
// Solidity: event NewDAppSignatory(address indexed controller, address indexed governance, address indexed signatory, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) FilterNewDAppSignatory(opts *bind.FilterOpts, controller []common.Address, governance []common.Address, signatory []common.Address) (*AtlasVerificationNewDAppSignatoryIterator, error) {

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

	logs, sub, err := _AtlasVerification.contract.FilterLogs(opts, "NewDAppSignatory", controllerRule, governanceRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return &AtlasVerificationNewDAppSignatoryIterator{contract: _AtlasVerification.contract, event: "NewDAppSignatory", logs: logs, sub: sub}, nil
}

// WatchNewDAppSignatory is a free log subscription operation binding the contract event 0xcacb2beeeba676b69bddabab0d5f66e2733cabc804f82afd92ea5beae71934b4.
//
// Solidity: event NewDAppSignatory(address indexed controller, address indexed governance, address indexed signatory, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) WatchNewDAppSignatory(opts *bind.WatchOpts, sink chan<- *AtlasVerificationNewDAppSignatory, controller []common.Address, governance []common.Address, signatory []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AtlasVerification.contract.WatchLogs(opts, "NewDAppSignatory", controllerRule, governanceRule, signatoryRule)
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
// Solidity: event NewDAppSignatory(address indexed controller, address indexed governance, address indexed signatory, uint32 callConfig)
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
	Controller common.Address
	Governance common.Address
	Signatory  common.Address
	CallConfig uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemovedDAppSignatory is a free log retrieval operation binding the contract event 0x90a400fe7fe676a24908d5f730fe6d38072340bc8f74e18a543132d840bfe5ce.
//
// Solidity: event RemovedDAppSignatory(address indexed controller, address indexed governance, address indexed signatory, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) FilterRemovedDAppSignatory(opts *bind.FilterOpts, controller []common.Address, governance []common.Address, signatory []common.Address) (*AtlasVerificationRemovedDAppSignatoryIterator, error) {

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

	logs, sub, err := _AtlasVerification.contract.FilterLogs(opts, "RemovedDAppSignatory", controllerRule, governanceRule, signatoryRule)
	if err != nil {
		return nil, err
	}
	return &AtlasVerificationRemovedDAppSignatoryIterator{contract: _AtlasVerification.contract, event: "RemovedDAppSignatory", logs: logs, sub: sub}, nil
}

// WatchRemovedDAppSignatory is a free log subscription operation binding the contract event 0x90a400fe7fe676a24908d5f730fe6d38072340bc8f74e18a543132d840bfe5ce.
//
// Solidity: event RemovedDAppSignatory(address indexed controller, address indexed governance, address indexed signatory, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) WatchRemovedDAppSignatory(opts *bind.WatchOpts, sink chan<- *AtlasVerificationRemovedDAppSignatory, controller []common.Address, governance []common.Address, signatory []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AtlasVerification.contract.WatchLogs(opts, "RemovedDAppSignatory", controllerRule, governanceRule, signatoryRule)
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
// Solidity: event RemovedDAppSignatory(address indexed controller, address indexed governance, address indexed signatory, uint32 callConfig)
func (_AtlasVerification *AtlasVerificationFilterer) ParseRemovedDAppSignatory(log types.Log) (*AtlasVerificationRemovedDAppSignatory, error) {
	event := new(AtlasVerificationRemovedDAppSignatory)
	if err := _AtlasVerification.contract.UnpackLog(event, "RemovedDAppSignatory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
