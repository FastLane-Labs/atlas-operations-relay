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

// DAppConfig is an auto generated low-level Go binding around an user-defined struct.
type DAppConfig struct {
	To         common.Address
	CallConfig uint32
	BidToken   common.Address
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

// AtlasMetaData contains all meta data concerning the Atlas contract.
var AtlasMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_escrowDuration\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_verification\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_simulator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_surchargeRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_executionTemplate\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ESCROW_DURATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SIMULATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SURCHARGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SURCHARGE_BASE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERIFICATION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accessData\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"bonded\",\"type\":\"uint112\",\"internalType\":\"uint112\"},{\"name\":\"lastAccessedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"auctionWins\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"auctionFails\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"totalGasUsed\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accountLastActiveBlock\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeEnvironment\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOfBonded\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"balanceOfUnbonding\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"becomeSurchargeRecipient\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bond\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bondedTotalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"borrow\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claims\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"contribute\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"createExecutionEnvironment\",\"inputs\":[{\"name\":\"dAppControl\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"executionEnvironment\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"depositAndBond\",\"inputs\":[{\"name\":\"amountToBond\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"deposits\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"dConfig\",\"type\":\"tuple\",\"internalType\":\"structDAppConfig\",\"components\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structUserOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"solverOps\",\"type\":\"tuple[]\",\"internalType\":\"structSolverOperation[]\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"executionEnvironment\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bundler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"auctionWon\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"winningSearcherIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"executionTemplate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExecutionEnvironment\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"dAppControl\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"executionEnvironment\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isUnlocked\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"lock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"metacall\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structUserOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dapp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sessionKey\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"solverOps\",\"type\":\"tuple[]\",\"internalType\":\"structSolverOperation[]\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"maxFeePerGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"solver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"dAppOp\",\"type\":\"tuple\",\"internalType\":\"structDAppOperation\",\"components\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"control\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"bundler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userOpHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callChainHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"auctionWon\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingSurchargeRecipient\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"reconcile\",\"inputs\":[{\"name\":\"environment\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solverFrom\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxApprovedGasSpend\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"owed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"salt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"shortfall\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"solver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"solverLockData\",\"inputs\":[],\"outputs\":[{\"name\":\"currentSolver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"calledBack\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"surcharge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"surchargeRecipient\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferDAppERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"controller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"lockState\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferSurchargeRecipient\",\"inputs\":[{\"name\":\"newRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferUserERC20\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destination\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"controller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"lockState\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unbond\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unbondingCompleteBlock\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validateBalances\",\"inputs\":[],\"outputs\":[{\"name\":\"valid\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawSurcharge\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Bond\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutionEnvironmentCreated\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executionEnvironment\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GasRefundSettled\",\"inputs\":[{\"name\":\"bundler\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"refundedETH\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MEVPaymentFailure\",\"inputs\":[{\"name\":\"controller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"callConfig\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"bidToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetacallResult\",\"inputs\":[{\"name\":\"bundler\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"winningSolver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PostOpsCall\",\"inputs\":[{\"name\":\"environment\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PreOpsCall\",\"inputs\":[{\"name\":\"environment\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Redeem\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SolverExecution\",\"inputs\":[{\"name\":\"solver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"isWin\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SolverTxResult\",\"inputs\":[{\"name\":\"solverTo\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"solverFrom\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"executed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"result\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SurchargeRecipientTransferStarted\",\"inputs\":[{\"name\":\"currentRecipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newRecipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SurchargeRecipientTransferred\",\"inputs\":[{\"name\":\"newRecipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SurchargeWithdrawn\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unbond\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"earliestAvailable\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserCall\",\"inputs\":[{\"name\":\"environment\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UserTxResult\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"valueReturned\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"gasRefunded\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlteredControl\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BalanceNotReconciled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DAppNotEnabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnvironmentMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EscrowLockActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientAtlETHBalance\",\"inputs\":[{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientAvailableBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientBalanceForDeduction\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientBondedBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientFunds\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientSolverBalance\",\"inputs\":[{\"name\":\"actual\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"msgValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"holds\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientSurchargeBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientTotalBalance\",\"inputs\":[{\"name\":\"shortfall\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientUnbondedBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"InsufficientWithdrawableBalance\",\"inputs\":[{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requested\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"IntentUnfulfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAccess\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCaller\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDAppControl\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEnvironment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidExecutionEnvironment\",\"inputs\":[{\"name\":\"correctEnvironment\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InvalidLockState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSolverFrom\",\"inputs\":[{\"name\":\"solverFrom\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"LedgerBalancing\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"LedgerFinalized\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"MissingFunds\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"NoAuctionWinner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoUnfilledRequests\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoUnusedNonceInBitmap\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyAccount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyGovernance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnerActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PermitDeadlineExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PostOpsFail\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PostOpsSimFail\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PostSolverFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PreOpsFail\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PreOpsSimFail\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PreSolverFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatoryActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SimulationPassed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SolverBidUnpaid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SolverMustReconcile\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SolverOperationReverted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SolverSimFail\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnbalancedAccounting\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UncoveredResult\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserNotFulfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserOpFail\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserOpSimFail\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserSimulationFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserSimulationSucceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UserUnexpectedSuccess\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ValidCalls\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumValidCallsResult\"}]},{\"type\":\"error\",\"name\":\"ValueTooLarge\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"VerificationSimFail\",\"inputs\":[]}]",
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

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Atlas *AtlasCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Atlas *AtlasSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Atlas.Contract.DOMAINSEPARATOR(&_Atlas.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Atlas *AtlasCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Atlas.Contract.DOMAINSEPARATOR(&_Atlas.CallOpts)
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

// SURCHARGE is a free data retrieval call binding the contract method 0x37642394.
//
// Solidity: function SURCHARGE() view returns(uint256)
func (_Atlas *AtlasCaller) SURCHARGE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "SURCHARGE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SURCHARGE is a free data retrieval call binding the contract method 0x37642394.
//
// Solidity: function SURCHARGE() view returns(uint256)
func (_Atlas *AtlasSession) SURCHARGE() (*big.Int, error) {
	return _Atlas.Contract.SURCHARGE(&_Atlas.CallOpts)
}

// SURCHARGE is a free data retrieval call binding the contract method 0x37642394.
//
// Solidity: function SURCHARGE() view returns(uint256)
func (_Atlas *AtlasCallerSession) SURCHARGE() (*big.Int, error) {
	return _Atlas.Contract.SURCHARGE(&_Atlas.CallOpts)
}

// SURCHARGEBASE is a free data retrieval call binding the contract method 0x478d8e53.
//
// Solidity: function SURCHARGE_BASE() view returns(uint256)
func (_Atlas *AtlasCaller) SURCHARGEBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "SURCHARGE_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SURCHARGEBASE is a free data retrieval call binding the contract method 0x478d8e53.
//
// Solidity: function SURCHARGE_BASE() view returns(uint256)
func (_Atlas *AtlasSession) SURCHARGEBASE() (*big.Int, error) {
	return _Atlas.Contract.SURCHARGEBASE(&_Atlas.CallOpts)
}

// SURCHARGEBASE is a free data retrieval call binding the contract method 0x478d8e53.
//
// Solidity: function SURCHARGE_BASE() view returns(uint256)
func (_Atlas *AtlasCallerSession) SURCHARGEBASE() (*big.Int, error) {
	return _Atlas.Contract.SURCHARGEBASE(&_Atlas.CallOpts)
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
// Solidity: function accessData(address ) view returns(uint112 bonded, uint32 lastAccessedBlock, uint24 auctionWins, uint24 auctionFails, uint64 totalGasUsed)
func (_Atlas *AtlasCaller) AccessData(opts *bind.CallOpts, arg0 common.Address) (struct {
	Bonded            *big.Int
	LastAccessedBlock uint32
	AuctionWins       *big.Int
	AuctionFails      *big.Int
	TotalGasUsed      uint64
}, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "accessData", arg0)

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
func (_Atlas *AtlasSession) AccessData(arg0 common.Address) (struct {
	Bonded            *big.Int
	LastAccessedBlock uint32
	AuctionWins       *big.Int
	AuctionFails      *big.Int
	TotalGasUsed      uint64
}, error) {
	return _Atlas.Contract.AccessData(&_Atlas.CallOpts, arg0)
}

// AccessData is a free data retrieval call binding the contract method 0x5e8edccc.
//
// Solidity: function accessData(address ) view returns(uint112 bonded, uint32 lastAccessedBlock, uint24 auctionWins, uint24 auctionFails, uint64 totalGasUsed)
func (_Atlas *AtlasCallerSession) AccessData(arg0 common.Address) (struct {
	Bonded            *big.Int
	LastAccessedBlock uint32
	AuctionWins       *big.Int
	AuctionFails      *big.Int
	TotalGasUsed      uint64
}, error) {
	return _Atlas.Contract.AccessData(&_Atlas.CallOpts, arg0)
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

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Atlas *AtlasCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Atlas *AtlasSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Atlas.Contract.Allowance(&_Atlas.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Atlas *AtlasCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Atlas.Contract.Allowance(&_Atlas.CallOpts, arg0, arg1)
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

// ExecutionTemplate is a free data retrieval call binding the contract method 0xe412ac3c.
//
// Solidity: function executionTemplate() view returns(address)
func (_Atlas *AtlasCaller) ExecutionTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "executionTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutionTemplate is a free data retrieval call binding the contract method 0xe412ac3c.
//
// Solidity: function executionTemplate() view returns(address)
func (_Atlas *AtlasSession) ExecutionTemplate() (common.Address, error) {
	return _Atlas.Contract.ExecutionTemplate(&_Atlas.CallOpts)
}

// ExecutionTemplate is a free data retrieval call binding the contract method 0xe412ac3c.
//
// Solidity: function executionTemplate() view returns(address)
func (_Atlas *AtlasCallerSession) ExecutionTemplate() (common.Address, error) {
	return _Atlas.Contract.ExecutionTemplate(&_Atlas.CallOpts)
}

// GetExecutionEnvironment is a free data retrieval call binding the contract method 0x45112906.
//
// Solidity: function getExecutionEnvironment(address user, address dAppControl) view returns(address executionEnvironment, uint32 callConfig, bool exists)
func (_Atlas *AtlasCaller) GetExecutionEnvironment(opts *bind.CallOpts, user common.Address, dAppControl common.Address) (struct {
	ExecutionEnvironment common.Address
	CallConfig           uint32
	Exists               bool
}, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "getExecutionEnvironment", user, dAppControl)

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
// Solidity: function getExecutionEnvironment(address user, address dAppControl) view returns(address executionEnvironment, uint32 callConfig, bool exists)
func (_Atlas *AtlasSession) GetExecutionEnvironment(user common.Address, dAppControl common.Address) (struct {
	ExecutionEnvironment common.Address
	CallConfig           uint32
	Exists               bool
}, error) {
	return _Atlas.Contract.GetExecutionEnvironment(&_Atlas.CallOpts, user, dAppControl)
}

// GetExecutionEnvironment is a free data retrieval call binding the contract method 0x45112906.
//
// Solidity: function getExecutionEnvironment(address user, address dAppControl) view returns(address executionEnvironment, uint32 callConfig, bool exists)
func (_Atlas *AtlasCallerSession) GetExecutionEnvironment(user common.Address, dAppControl common.Address) (struct {
	ExecutionEnvironment common.Address
	CallConfig           uint32
	Exists               bool
}, error) {
	return _Atlas.Contract.GetExecutionEnvironment(&_Atlas.CallOpts, user, dAppControl)
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
// Solidity: function lock() view returns(address)
func (_Atlas *AtlasCaller) Lock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "lock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() view returns(address)
func (_Atlas *AtlasSession) Lock() (common.Address, error) {
	return _Atlas.Contract.Lock(&_Atlas.CallOpts)
}

// Lock is a free data retrieval call binding the contract method 0xf83d08ba.
//
// Solidity: function lock() view returns(address)
func (_Atlas *AtlasCallerSession) Lock() (common.Address, error) {
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

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Atlas *AtlasCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Atlas *AtlasSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Atlas.Contract.Nonces(&_Atlas.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Atlas *AtlasCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Atlas.Contract.Nonces(&_Atlas.CallOpts, arg0)
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

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_Atlas *AtlasCaller) Salt(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "salt")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_Atlas *AtlasSession) Salt() ([32]byte, error) {
	return _Atlas.Contract.Salt(&_Atlas.CallOpts)
}

// Salt is a free data retrieval call binding the contract method 0xbfa0b133.
//
// Solidity: function salt() view returns(bytes32)
func (_Atlas *AtlasCallerSession) Salt() ([32]byte, error) {
	return _Atlas.Contract.Salt(&_Atlas.CallOpts)
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

// Surcharge is a free data retrieval call binding the contract method 0x2ab26a56.
//
// Solidity: function surcharge() view returns(uint256)
func (_Atlas *AtlasCaller) Surcharge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "surcharge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Surcharge is a free data retrieval call binding the contract method 0x2ab26a56.
//
// Solidity: function surcharge() view returns(uint256)
func (_Atlas *AtlasSession) Surcharge() (*big.Int, error) {
	return _Atlas.Contract.Surcharge(&_Atlas.CallOpts)
}

// Surcharge is a free data retrieval call binding the contract method 0x2ab26a56.
//
// Solidity: function surcharge() view returns(uint256)
func (_Atlas *AtlasCallerSession) Surcharge() (*big.Int, error) {
	return _Atlas.Contract.Surcharge(&_Atlas.CallOpts)
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

// ValidateBalances is a free data retrieval call binding the contract method 0x15827e08.
//
// Solidity: function validateBalances() view returns(bool valid)
func (_Atlas *AtlasCaller) ValidateBalances(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Atlas.contract.Call(opts, &out, "validateBalances")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateBalances is a free data retrieval call binding the contract method 0x15827e08.
//
// Solidity: function validateBalances() view returns(bool valid)
func (_Atlas *AtlasSession) ValidateBalances() (bool, error) {
	return _Atlas.Contract.ValidateBalances(&_Atlas.CallOpts)
}

// ValidateBalances is a free data retrieval call binding the contract method 0x15827e08.
//
// Solidity: function validateBalances() view returns(bool valid)
func (_Atlas *AtlasCallerSession) ValidateBalances() (bool, error) {
	return _Atlas.Contract.ValidateBalances(&_Atlas.CallOpts)
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

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Atlas *AtlasTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Atlas *AtlasSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Approve(&_Atlas.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Atlas *AtlasTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Approve(&_Atlas.TransactOpts, spender, amount)
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
// Solidity: function createExecutionEnvironment(address dAppControl) returns(address executionEnvironment)
func (_Atlas *AtlasTransactor) CreateExecutionEnvironment(opts *bind.TransactOpts, dAppControl common.Address) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "createExecutionEnvironment", dAppControl)
}

// CreateExecutionEnvironment is a paid mutator transaction binding the contract method 0x4626e4a2.
//
// Solidity: function createExecutionEnvironment(address dAppControl) returns(address executionEnvironment)
func (_Atlas *AtlasSession) CreateExecutionEnvironment(dAppControl common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.CreateExecutionEnvironment(&_Atlas.TransactOpts, dAppControl)
}

// CreateExecutionEnvironment is a paid mutator transaction binding the contract method 0x4626e4a2.
//
// Solidity: function createExecutionEnvironment(address dAppControl) returns(address executionEnvironment)
func (_Atlas *AtlasTransactorSession) CreateExecutionEnvironment(dAppControl common.Address) (*types.Transaction, error) {
	return _Atlas.Contract.CreateExecutionEnvironment(&_Atlas.TransactOpts, dAppControl)
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

// Execute is a paid mutator transaction binding the contract method 0xd3707689.
//
// Solidity: function execute((address,uint32,address) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, address executionEnvironment, address bundler, bytes32 userOpHash) payable returns(bool auctionWon, uint256 winningSearcherIndex)
func (_Atlas *AtlasTransactor) Execute(opts *bind.TransactOpts, dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, executionEnvironment common.Address, bundler common.Address, userOpHash [32]byte) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "execute", dConfig, userOp, solverOps, executionEnvironment, bundler, userOpHash)
}

// Execute is a paid mutator transaction binding the contract method 0xd3707689.
//
// Solidity: function execute((address,uint32,address) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, address executionEnvironment, address bundler, bytes32 userOpHash) payable returns(bool auctionWon, uint256 winningSearcherIndex)
func (_Atlas *AtlasSession) Execute(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, executionEnvironment common.Address, bundler common.Address, userOpHash [32]byte) (*types.Transaction, error) {
	return _Atlas.Contract.Execute(&_Atlas.TransactOpts, dConfig, userOp, solverOps, executionEnvironment, bundler, userOpHash)
}

// Execute is a paid mutator transaction binding the contract method 0xd3707689.
//
// Solidity: function execute((address,uint32,address) dConfig, (address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, address executionEnvironment, address bundler, bytes32 userOpHash) payable returns(bool auctionWon, uint256 winningSearcherIndex)
func (_Atlas *AtlasTransactorSession) Execute(dConfig DAppConfig, userOp UserOperation, solverOps []SolverOperation, executionEnvironment common.Address, bundler common.Address, userOpHash [32]byte) (*types.Transaction, error) {
	return _Atlas.Contract.Execute(&_Atlas.TransactOpts, dConfig, userOp, solverOps, executionEnvironment, bundler, userOpHash)
}

// Metacall is a paid mutator transaction binding the contract method 0xda91b284.
//
// Solidity: function metacall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool auctionWon)
func (_Atlas *AtlasTransactor) Metacall(opts *bind.TransactOpts, userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "metacall", userOp, solverOps, dAppOp)
}

// Metacall is a paid mutator transaction binding the contract method 0xda91b284.
//
// Solidity: function metacall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool auctionWon)
func (_Atlas *AtlasSession) Metacall(userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Atlas.Contract.Metacall(&_Atlas.TransactOpts, userOp, solverOps, dAppOp)
}

// Metacall is a paid mutator transaction binding the contract method 0xda91b284.
//
// Solidity: function metacall((address,address,uint256,uint256,uint256,uint256,uint256,address,address,address,bytes,bytes) userOp, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,address,uint256,bytes,bytes)[] solverOps, (address,address,uint256,uint256,uint256,uint256,address,address,bytes32,bytes32,bytes) dAppOp) payable returns(bool auctionWon)
func (_Atlas *AtlasTransactorSession) Metacall(userOp UserOperation, solverOps []SolverOperation, dAppOp DAppOperation) (*types.Transaction, error) {
	return _Atlas.Contract.Metacall(&_Atlas.TransactOpts, userOp, solverOps, dAppOp)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Atlas *AtlasTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Atlas *AtlasSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Atlas.Contract.Permit(&_Atlas.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Atlas *AtlasTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Atlas.Contract.Permit(&_Atlas.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Reconcile is a paid mutator transaction binding the contract method 0xc6490cc2.
//
// Solidity: function reconcile(address environment, address solverFrom, uint256 maxApprovedGasSpend) payable returns(uint256 owed)
func (_Atlas *AtlasTransactor) Reconcile(opts *bind.TransactOpts, environment common.Address, solverFrom common.Address, maxApprovedGasSpend *big.Int) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "reconcile", environment, solverFrom, maxApprovedGasSpend)
}

// Reconcile is a paid mutator transaction binding the contract method 0xc6490cc2.
//
// Solidity: function reconcile(address environment, address solverFrom, uint256 maxApprovedGasSpend) payable returns(uint256 owed)
func (_Atlas *AtlasSession) Reconcile(environment common.Address, solverFrom common.Address, maxApprovedGasSpend *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Reconcile(&_Atlas.TransactOpts, environment, solverFrom, maxApprovedGasSpend)
}

// Reconcile is a paid mutator transaction binding the contract method 0xc6490cc2.
//
// Solidity: function reconcile(address environment, address solverFrom, uint256 maxApprovedGasSpend) payable returns(uint256 owed)
func (_Atlas *AtlasTransactorSession) Reconcile(environment common.Address, solverFrom common.Address, maxApprovedGasSpend *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Reconcile(&_Atlas.TransactOpts, environment, solverFrom, maxApprovedGasSpend)
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

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Atlas *AtlasTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Atlas *AtlasSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Transfer(&_Atlas.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Atlas *AtlasTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.Transfer(&_Atlas.TransactOpts, to, amount)
}

// TransferDAppERC20 is a paid mutator transaction binding the contract method 0x6625f68b.
//
// Solidity: function transferDAppERC20(address token, address destination, uint256 amount, address user, address controller, uint32 callConfig, uint16 lockState) returns()
func (_Atlas *AtlasTransactor) TransferDAppERC20(opts *bind.TransactOpts, token common.Address, destination common.Address, amount *big.Int, user common.Address, controller common.Address, callConfig uint32, lockState uint16) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "transferDAppERC20", token, destination, amount, user, controller, callConfig, lockState)
}

// TransferDAppERC20 is a paid mutator transaction binding the contract method 0x6625f68b.
//
// Solidity: function transferDAppERC20(address token, address destination, uint256 amount, address user, address controller, uint32 callConfig, uint16 lockState) returns()
func (_Atlas *AtlasSession) TransferDAppERC20(token common.Address, destination common.Address, amount *big.Int, user common.Address, controller common.Address, callConfig uint32, lockState uint16) (*types.Transaction, error) {
	return _Atlas.Contract.TransferDAppERC20(&_Atlas.TransactOpts, token, destination, amount, user, controller, callConfig, lockState)
}

// TransferDAppERC20 is a paid mutator transaction binding the contract method 0x6625f68b.
//
// Solidity: function transferDAppERC20(address token, address destination, uint256 amount, address user, address controller, uint32 callConfig, uint16 lockState) returns()
func (_Atlas *AtlasTransactorSession) TransferDAppERC20(token common.Address, destination common.Address, amount *big.Int, user common.Address, controller common.Address, callConfig uint32, lockState uint16) (*types.Transaction, error) {
	return _Atlas.Contract.TransferDAppERC20(&_Atlas.TransactOpts, token, destination, amount, user, controller, callConfig, lockState)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Atlas *AtlasTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Atlas *AtlasSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.TransferFrom(&_Atlas.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Atlas *AtlasTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Atlas.Contract.TransferFrom(&_Atlas.TransactOpts, from, to, amount)
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

// TransferUserERC20 is a paid mutator transaction binding the contract method 0x2c63112c.
//
// Solidity: function transferUserERC20(address token, address destination, uint256 amount, address user, address controller, uint32 callConfig, uint16 lockState) returns()
func (_Atlas *AtlasTransactor) TransferUserERC20(opts *bind.TransactOpts, token common.Address, destination common.Address, amount *big.Int, user common.Address, controller common.Address, callConfig uint32, lockState uint16) (*types.Transaction, error) {
	return _Atlas.contract.Transact(opts, "transferUserERC20", token, destination, amount, user, controller, callConfig, lockState)
}

// TransferUserERC20 is a paid mutator transaction binding the contract method 0x2c63112c.
//
// Solidity: function transferUserERC20(address token, address destination, uint256 amount, address user, address controller, uint32 callConfig, uint16 lockState) returns()
func (_Atlas *AtlasSession) TransferUserERC20(token common.Address, destination common.Address, amount *big.Int, user common.Address, controller common.Address, callConfig uint32, lockState uint16) (*types.Transaction, error) {
	return _Atlas.Contract.TransferUserERC20(&_Atlas.TransactOpts, token, destination, amount, user, controller, callConfig, lockState)
}

// TransferUserERC20 is a paid mutator transaction binding the contract method 0x2c63112c.
//
// Solidity: function transferUserERC20(address token, address destination, uint256 amount, address user, address controller, uint32 callConfig, uint16 lockState) returns()
func (_Atlas *AtlasTransactorSession) TransferUserERC20(token common.Address, destination common.Address, amount *big.Int, user common.Address, controller common.Address, callConfig uint32, lockState uint16) (*types.Transaction, error) {
	return _Atlas.Contract.TransferUserERC20(&_Atlas.TransactOpts, token, destination, amount, user, controller, callConfig, lockState)
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

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Atlas *AtlasTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Atlas.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Atlas *AtlasSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Atlas.Contract.Fallback(&_Atlas.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Atlas *AtlasTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Atlas.Contract.Fallback(&_Atlas.TransactOpts, calldata)
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

// AtlasGasRefundSettledIterator is returned from FilterGasRefundSettled and is used to iterate over the raw logs and unpacked data for GasRefundSettled events raised by the Atlas contract.
type AtlasGasRefundSettledIterator struct {
	Event *AtlasGasRefundSettled // Event containing the contract specifics and raw log

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
func (it *AtlasGasRefundSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasGasRefundSettled)
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
		it.Event = new(AtlasGasRefundSettled)
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
func (it *AtlasGasRefundSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasGasRefundSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasGasRefundSettled represents a GasRefundSettled event raised by the Atlas contract.
type AtlasGasRefundSettled struct {
	Bundler     common.Address
	RefundedETH *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGasRefundSettled is a free log retrieval operation binding the contract event 0x390a4349501fe8d6d7f56fb3ed1ff7d9d28f12ce40b594155df3b7d313c1d5af.
//
// Solidity: event GasRefundSettled(address indexed bundler, uint256 refundedETH)
func (_Atlas *AtlasFilterer) FilterGasRefundSettled(opts *bind.FilterOpts, bundler []common.Address) (*AtlasGasRefundSettledIterator, error) {

	var bundlerRule []interface{}
	for _, bundlerItem := range bundler {
		bundlerRule = append(bundlerRule, bundlerItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "GasRefundSettled", bundlerRule)
	if err != nil {
		return nil, err
	}
	return &AtlasGasRefundSettledIterator{contract: _Atlas.contract, event: "GasRefundSettled", logs: logs, sub: sub}, nil
}

// WatchGasRefundSettled is a free log subscription operation binding the contract event 0x390a4349501fe8d6d7f56fb3ed1ff7d9d28f12ce40b594155df3b7d313c1d5af.
//
// Solidity: event GasRefundSettled(address indexed bundler, uint256 refundedETH)
func (_Atlas *AtlasFilterer) WatchGasRefundSettled(opts *bind.WatchOpts, sink chan<- *AtlasGasRefundSettled, bundler []common.Address) (event.Subscription, error) {

	var bundlerRule []interface{}
	for _, bundlerItem := range bundler {
		bundlerRule = append(bundlerRule, bundlerItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "GasRefundSettled", bundlerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasGasRefundSettled)
				if err := _Atlas.contract.UnpackLog(event, "GasRefundSettled", log); err != nil {
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

// ParseGasRefundSettled is a log parse operation binding the contract event 0x390a4349501fe8d6d7f56fb3ed1ff7d9d28f12ce40b594155df3b7d313c1d5af.
//
// Solidity: event GasRefundSettled(address indexed bundler, uint256 refundedETH)
func (_Atlas *AtlasFilterer) ParseGasRefundSettled(log types.Log) (*AtlasGasRefundSettled, error) {
	event := new(AtlasGasRefundSettled)
	if err := _Atlas.contract.UnpackLog(event, "GasRefundSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasMEVPaymentFailureIterator is returned from FilterMEVPaymentFailure and is used to iterate over the raw logs and unpacked data for MEVPaymentFailure events raised by the Atlas contract.
type AtlasMEVPaymentFailureIterator struct {
	Event *AtlasMEVPaymentFailure // Event containing the contract specifics and raw log

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
func (it *AtlasMEVPaymentFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasMEVPaymentFailure)
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
		it.Event = new(AtlasMEVPaymentFailure)
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
func (it *AtlasMEVPaymentFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasMEVPaymentFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasMEVPaymentFailure represents a MEVPaymentFailure event raised by the Atlas contract.
type AtlasMEVPaymentFailure struct {
	Controller common.Address
	CallConfig uint32
	BidToken   common.Address
	BidAmount  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMEVPaymentFailure is a free log retrieval operation binding the contract event 0x47ac7e361a60b2835f2a22c12ba42fe0bdcfc613a15f45574bce05ff38c04c13.
//
// Solidity: event MEVPaymentFailure(address indexed controller, uint32 callConfig, address bidToken, uint256 bidAmount)
func (_Atlas *AtlasFilterer) FilterMEVPaymentFailure(opts *bind.FilterOpts, controller []common.Address) (*AtlasMEVPaymentFailureIterator, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "MEVPaymentFailure", controllerRule)
	if err != nil {
		return nil, err
	}
	return &AtlasMEVPaymentFailureIterator{contract: _Atlas.contract, event: "MEVPaymentFailure", logs: logs, sub: sub}, nil
}

// WatchMEVPaymentFailure is a free log subscription operation binding the contract event 0x47ac7e361a60b2835f2a22c12ba42fe0bdcfc613a15f45574bce05ff38c04c13.
//
// Solidity: event MEVPaymentFailure(address indexed controller, uint32 callConfig, address bidToken, uint256 bidAmount)
func (_Atlas *AtlasFilterer) WatchMEVPaymentFailure(opts *bind.WatchOpts, sink chan<- *AtlasMEVPaymentFailure, controller []common.Address) (event.Subscription, error) {

	var controllerRule []interface{}
	for _, controllerItem := range controller {
		controllerRule = append(controllerRule, controllerItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "MEVPaymentFailure", controllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasMEVPaymentFailure)
				if err := _Atlas.contract.UnpackLog(event, "MEVPaymentFailure", log); err != nil {
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

// ParseMEVPaymentFailure is a log parse operation binding the contract event 0x47ac7e361a60b2835f2a22c12ba42fe0bdcfc613a15f45574bce05ff38c04c13.
//
// Solidity: event MEVPaymentFailure(address indexed controller, uint32 callConfig, address bidToken, uint256 bidAmount)
func (_Atlas *AtlasFilterer) ParseMEVPaymentFailure(log types.Log) (*AtlasMEVPaymentFailure, error) {
	event := new(AtlasMEVPaymentFailure)
	if err := _Atlas.contract.UnpackLog(event, "MEVPaymentFailure", log); err != nil {
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
	Bundler       common.Address
	User          common.Address
	WinningSolver common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMetacallResult is a free log retrieval operation binding the contract event 0xccaf53f24f4a968247a7982fa8d532542e3c75f86438a1e6d75add121f8f49ff.
//
// Solidity: event MetacallResult(address indexed bundler, address indexed user, address indexed winningSolver)
func (_Atlas *AtlasFilterer) FilterMetacallResult(opts *bind.FilterOpts, bundler []common.Address, user []common.Address, winningSolver []common.Address) (*AtlasMetacallResultIterator, error) {

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

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "MetacallResult", bundlerRule, userRule, winningSolverRule)
	if err != nil {
		return nil, err
	}
	return &AtlasMetacallResultIterator{contract: _Atlas.contract, event: "MetacallResult", logs: logs, sub: sub}, nil
}

// WatchMetacallResult is a free log subscription operation binding the contract event 0xccaf53f24f4a968247a7982fa8d532542e3c75f86438a1e6d75add121f8f49ff.
//
// Solidity: event MetacallResult(address indexed bundler, address indexed user, address indexed winningSolver)
func (_Atlas *AtlasFilterer) WatchMetacallResult(opts *bind.WatchOpts, sink chan<- *AtlasMetacallResult, bundler []common.Address, user []common.Address, winningSolver []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "MetacallResult", bundlerRule, userRule, winningSolverRule)
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

// ParseMetacallResult is a log parse operation binding the contract event 0xccaf53f24f4a968247a7982fa8d532542e3c75f86438a1e6d75add121f8f49ff.
//
// Solidity: event MetacallResult(address indexed bundler, address indexed user, address indexed winningSolver)
func (_Atlas *AtlasFilterer) ParseMetacallResult(log types.Log) (*AtlasMetacallResult, error) {
	event := new(AtlasMetacallResult)
	if err := _Atlas.contract.UnpackLog(event, "MetacallResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasPostOpsCallIterator is returned from FilterPostOpsCall and is used to iterate over the raw logs and unpacked data for PostOpsCall events raised by the Atlas contract.
type AtlasPostOpsCallIterator struct {
	Event *AtlasPostOpsCall // Event containing the contract specifics and raw log

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
func (it *AtlasPostOpsCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasPostOpsCall)
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
		it.Event = new(AtlasPostOpsCall)
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
func (it *AtlasPostOpsCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasPostOpsCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasPostOpsCall represents a PostOpsCall event raised by the Atlas contract.
type AtlasPostOpsCall struct {
	Environment common.Address
	Success     bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPostOpsCall is a free log retrieval operation binding the contract event 0x7e8cfbbc29f299b9f787b9420dcc5601fd9cf36849a9e5fe6214191b14d81a80.
//
// Solidity: event PostOpsCall(address environment, bool success)
func (_Atlas *AtlasFilterer) FilterPostOpsCall(opts *bind.FilterOpts) (*AtlasPostOpsCallIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "PostOpsCall")
	if err != nil {
		return nil, err
	}
	return &AtlasPostOpsCallIterator{contract: _Atlas.contract, event: "PostOpsCall", logs: logs, sub: sub}, nil
}

// WatchPostOpsCall is a free log subscription operation binding the contract event 0x7e8cfbbc29f299b9f787b9420dcc5601fd9cf36849a9e5fe6214191b14d81a80.
//
// Solidity: event PostOpsCall(address environment, bool success)
func (_Atlas *AtlasFilterer) WatchPostOpsCall(opts *bind.WatchOpts, sink chan<- *AtlasPostOpsCall) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "PostOpsCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasPostOpsCall)
				if err := _Atlas.contract.UnpackLog(event, "PostOpsCall", log); err != nil {
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

// ParsePostOpsCall is a log parse operation binding the contract event 0x7e8cfbbc29f299b9f787b9420dcc5601fd9cf36849a9e5fe6214191b14d81a80.
//
// Solidity: event PostOpsCall(address environment, bool success)
func (_Atlas *AtlasFilterer) ParsePostOpsCall(log types.Log) (*AtlasPostOpsCall, error) {
	event := new(AtlasPostOpsCall)
	if err := _Atlas.contract.UnpackLog(event, "PostOpsCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasPreOpsCallIterator is returned from FilterPreOpsCall and is used to iterate over the raw logs and unpacked data for PreOpsCall events raised by the Atlas contract.
type AtlasPreOpsCallIterator struct {
	Event *AtlasPreOpsCall // Event containing the contract specifics and raw log

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
func (it *AtlasPreOpsCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasPreOpsCall)
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
		it.Event = new(AtlasPreOpsCall)
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
func (it *AtlasPreOpsCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasPreOpsCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasPreOpsCall represents a PreOpsCall event raised by the Atlas contract.
type AtlasPreOpsCall struct {
	Environment common.Address
	Success     bool
	ReturnData  []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPreOpsCall is a free log retrieval operation binding the contract event 0x32943f96b9a2a67c93967eafd91c42ca73ee7460e53f411fb1a18c31c7e42f7f.
//
// Solidity: event PreOpsCall(address environment, bool success, bytes returnData)
func (_Atlas *AtlasFilterer) FilterPreOpsCall(opts *bind.FilterOpts) (*AtlasPreOpsCallIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "PreOpsCall")
	if err != nil {
		return nil, err
	}
	return &AtlasPreOpsCallIterator{contract: _Atlas.contract, event: "PreOpsCall", logs: logs, sub: sub}, nil
}

// WatchPreOpsCall is a free log subscription operation binding the contract event 0x32943f96b9a2a67c93967eafd91c42ca73ee7460e53f411fb1a18c31c7e42f7f.
//
// Solidity: event PreOpsCall(address environment, bool success, bytes returnData)
func (_Atlas *AtlasFilterer) WatchPreOpsCall(opts *bind.WatchOpts, sink chan<- *AtlasPreOpsCall) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "PreOpsCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasPreOpsCall)
				if err := _Atlas.contract.UnpackLog(event, "PreOpsCall", log); err != nil {
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

// ParsePreOpsCall is a log parse operation binding the contract event 0x32943f96b9a2a67c93967eafd91c42ca73ee7460e53f411fb1a18c31c7e42f7f.
//
// Solidity: event PreOpsCall(address environment, bool success, bytes returnData)
func (_Atlas *AtlasFilterer) ParsePreOpsCall(log types.Log) (*AtlasPreOpsCall, error) {
	event := new(AtlasPreOpsCall)
	if err := _Atlas.contract.UnpackLog(event, "PreOpsCall", log); err != nil {
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

// AtlasSolverExecutionIterator is returned from FilterSolverExecution and is used to iterate over the raw logs and unpacked data for SolverExecution events raised by the Atlas contract.
type AtlasSolverExecutionIterator struct {
	Event *AtlasSolverExecution // Event containing the contract specifics and raw log

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
func (it *AtlasSolverExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasSolverExecution)
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
		it.Event = new(AtlasSolverExecution)
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
func (it *AtlasSolverExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasSolverExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasSolverExecution represents a SolverExecution event raised by the Atlas contract.
type AtlasSolverExecution struct {
	Solver common.Address
	Index  *big.Int
	IsWin  bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSolverExecution is a free log retrieval operation binding the contract event 0x20ecaac325cd57efaf6df656e93d1f89ecc1b0a8d3a053542fc9a58745697d6c.
//
// Solidity: event SolverExecution(address indexed solver, uint256 index, bool isWin)
func (_Atlas *AtlasFilterer) FilterSolverExecution(opts *bind.FilterOpts, solver []common.Address) (*AtlasSolverExecutionIterator, error) {

	var solverRule []interface{}
	for _, solverItem := range solver {
		solverRule = append(solverRule, solverItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "SolverExecution", solverRule)
	if err != nil {
		return nil, err
	}
	return &AtlasSolverExecutionIterator{contract: _Atlas.contract, event: "SolverExecution", logs: logs, sub: sub}, nil
}

// WatchSolverExecution is a free log subscription operation binding the contract event 0x20ecaac325cd57efaf6df656e93d1f89ecc1b0a8d3a053542fc9a58745697d6c.
//
// Solidity: event SolverExecution(address indexed solver, uint256 index, bool isWin)
func (_Atlas *AtlasFilterer) WatchSolverExecution(opts *bind.WatchOpts, sink chan<- *AtlasSolverExecution, solver []common.Address) (event.Subscription, error) {

	var solverRule []interface{}
	for _, solverItem := range solver {
		solverRule = append(solverRule, solverItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "SolverExecution", solverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasSolverExecution)
				if err := _Atlas.contract.UnpackLog(event, "SolverExecution", log); err != nil {
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

// ParseSolverExecution is a log parse operation binding the contract event 0x20ecaac325cd57efaf6df656e93d1f89ecc1b0a8d3a053542fc9a58745697d6c.
//
// Solidity: event SolverExecution(address indexed solver, uint256 index, bool isWin)
func (_Atlas *AtlasFilterer) ParseSolverExecution(log types.Log) (*AtlasSolverExecution, error) {
	event := new(AtlasSolverExecution)
	if err := _Atlas.contract.UnpackLog(event, "SolverExecution", log); err != nil {
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

// AtlasUserCallIterator is returned from FilterUserCall and is used to iterate over the raw logs and unpacked data for UserCall events raised by the Atlas contract.
type AtlasUserCallIterator struct {
	Event *AtlasUserCall // Event containing the contract specifics and raw log

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
func (it *AtlasUserCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasUserCall)
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
		it.Event = new(AtlasUserCall)
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
func (it *AtlasUserCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasUserCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasUserCall represents a UserCall event raised by the Atlas contract.
type AtlasUserCall struct {
	Environment common.Address
	Success     bool
	ReturnData  []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUserCall is a free log retrieval operation binding the contract event 0xc753e45077e0333d261b6ae90f2f18d7c9cecc7bad6f004b2a483a3cb00b2b85.
//
// Solidity: event UserCall(address environment, bool success, bytes returnData)
func (_Atlas *AtlasFilterer) FilterUserCall(opts *bind.FilterOpts) (*AtlasUserCallIterator, error) {

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "UserCall")
	if err != nil {
		return nil, err
	}
	return &AtlasUserCallIterator{contract: _Atlas.contract, event: "UserCall", logs: logs, sub: sub}, nil
}

// WatchUserCall is a free log subscription operation binding the contract event 0xc753e45077e0333d261b6ae90f2f18d7c9cecc7bad6f004b2a483a3cb00b2b85.
//
// Solidity: event UserCall(address environment, bool success, bytes returnData)
func (_Atlas *AtlasFilterer) WatchUserCall(opts *bind.WatchOpts, sink chan<- *AtlasUserCall) (event.Subscription, error) {

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "UserCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasUserCall)
				if err := _Atlas.contract.UnpackLog(event, "UserCall", log); err != nil {
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

// ParseUserCall is a log parse operation binding the contract event 0xc753e45077e0333d261b6ae90f2f18d7c9cecc7bad6f004b2a483a3cb00b2b85.
//
// Solidity: event UserCall(address environment, bool success, bytes returnData)
func (_Atlas *AtlasFilterer) ParseUserCall(log types.Log) (*AtlasUserCall, error) {
	event := new(AtlasUserCall)
	if err := _Atlas.contract.UnpackLog(event, "UserCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AtlasUserTxResultIterator is returned from FilterUserTxResult and is used to iterate over the raw logs and unpacked data for UserTxResult events raised by the Atlas contract.
type AtlasUserTxResultIterator struct {
	Event *AtlasUserTxResult // Event containing the contract specifics and raw log

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
func (it *AtlasUserTxResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AtlasUserTxResult)
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
		it.Event = new(AtlasUserTxResult)
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
func (it *AtlasUserTxResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AtlasUserTxResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AtlasUserTxResult represents a UserTxResult event raised by the Atlas contract.
type AtlasUserTxResult struct {
	User          common.Address
	ValueReturned *big.Int
	GasRefunded   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserTxResult is a free log retrieval operation binding the contract event 0xea1c7d02ed469d8ee042fe1710cdd7406bc6568f46607fdb0af29fb03e4a82b5.
//
// Solidity: event UserTxResult(address indexed user, uint256 valueReturned, uint256 gasRefunded)
func (_Atlas *AtlasFilterer) FilterUserTxResult(opts *bind.FilterOpts, user []common.Address) (*AtlasUserTxResultIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Atlas.contract.FilterLogs(opts, "UserTxResult", userRule)
	if err != nil {
		return nil, err
	}
	return &AtlasUserTxResultIterator{contract: _Atlas.contract, event: "UserTxResult", logs: logs, sub: sub}, nil
}

// WatchUserTxResult is a free log subscription operation binding the contract event 0xea1c7d02ed469d8ee042fe1710cdd7406bc6568f46607fdb0af29fb03e4a82b5.
//
// Solidity: event UserTxResult(address indexed user, uint256 valueReturned, uint256 gasRefunded)
func (_Atlas *AtlasFilterer) WatchUserTxResult(opts *bind.WatchOpts, sink chan<- *AtlasUserTxResult, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Atlas.contract.WatchLogs(opts, "UserTxResult", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AtlasUserTxResult)
				if err := _Atlas.contract.UnpackLog(event, "UserTxResult", log); err != nil {
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

// ParseUserTxResult is a log parse operation binding the contract event 0xea1c7d02ed469d8ee042fe1710cdd7406bc6568f46607fdb0af29fb03e4a82b5.
//
// Solidity: event UserTxResult(address indexed user, uint256 valueReturned, uint256 gasRefunded)
func (_Atlas *AtlasFilterer) ParseUserTxResult(log types.Log) (*AtlasUserTxResult, error) {
	event := new(AtlasUserTxResult)
	if err := _Atlas.contract.UnpackLog(event, "UserTxResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
