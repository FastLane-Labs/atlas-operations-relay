package operation

import (
	"context"
	"math/big"
	"math/rand"

	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/FastLane-Labs/atlas-operations-relay/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

var (
	ErrUserOpInvalidSignature = relayerror.NewError(2001, "user operation has invalid signature")
	ErrUserOpInvalidToField   = relayerror.NewError(2002, "user operation's 'to' field must be atlas contract address")
	ErrUserOpDeadlineExceeded = relayerror.NewError(2003, "user operation's deadline exceeded")
	ErrUserOpComputeHash      = relayerror.NewError(2004, "failed to compute user operation hash")
	ErrUserOpSignatureInvalid = relayerror.NewError(2006, "user operation has invalid signature")
	ErrUserOpGasLimitExceeded = relayerror.NewError(2007, "user operation's gas limit exceeded")
)

var (
	userOpSolType, _ = abi.NewType("tuple", "struct UserOperation", []abi.ArgumentMarshaling{
		{Name: "from", Type: "address", InternalType: "address"},
		{Name: "to", Type: "address", InternalType: "address"},
		{Name: "value", Type: "uint256", InternalType: "uint256"},
		{Name: "gas", Type: "uint256", InternalType: "uint256"},
		{Name: "maxFeePerGas", Type: "uint256", InternalType: "uint256"},
		{Name: "nonce", Type: "uint256", InternalType: "uint256"},
		{Name: "deadline", Type: "uint256", InternalType: "uint256"},
		{Name: "dapp", Type: "address", InternalType: "address"},
		{Name: "control", Type: "address", InternalType: "address"},
		{Name: "callConfig", Type: "uint32", InternalType: "uint32"},
		{Name: "sessionKey", Type: "address", InternalType: "address"},
		{Name: "data", Type: "bytes", InternalType: "bytes"},
		{Name: "signature", Type: "bytes", InternalType: "bytes"},
	})

	userOpArgs = abi.Arguments{
		{Type: userOpSolType, Name: "userOperation"},
	}
)

// External representation of a user operation,
// the relay receives and broadcasts user operations in this format
type UserOperationRaw struct {
	From         common.Address `json:"from" validate:"required"`
	To           common.Address `json:"to" validate:"required"`
	Value        *hexutil.Big   `json:"value" validate:"required"`
	Gas          *hexutil.Big   `json:"gas" validate:"required"`
	MaxFeePerGas *hexutil.Big   `json:"maxFeePerGas" validate:"required"`
	Nonce        *hexutil.Big   `json:"nonce" validate:"required"`
	Deadline     *hexutil.Big   `json:"deadline" validate:"required"`
	Dapp         common.Address `json:"dapp" validate:"required"`
	Control      common.Address `json:"control" validate:"required"`
	CallConfig   *hexutil.Big   `json:"callConfig" validate:"required"`
	SessionKey   common.Address `json:"sessionKey"` // Optional
	Data         hexutil.Bytes  `json:"data" validate:"required"`
	Signature    hexutil.Bytes  `json:"signature" validate:"required"`
}

func (u *UserOperationRaw) Decode() *UserOperation {
	return &UserOperation{
		From:         u.From,
		To:           u.To,
		Value:        u.Value.ToInt(),
		Gas:          u.Gas.ToInt(),
		MaxFeePerGas: u.MaxFeePerGas.ToInt(),
		Nonce:        u.Nonce.ToInt(),
		Deadline:     u.Deadline.ToInt(),
		Dapp:         u.Dapp,
		Control:      u.Control,
		CallConfig:   uint32(u.CallConfig.ToInt().Uint64()),
		SessionKey:   u.SessionKey,
		Data:         u.Data,
		Signature:    u.Signature,
	}
}

type UserOperationWithHintsRaw struct {
	UserOperation *UserOperationRaw `json:"userOperation" validate:"required"`
	Hints         []common.Address  `json:"hints"` // Optional
}

func (uop *UserOperationWithHintsRaw) Decode() (*UserOperation, []common.Address) {
	return uop.UserOperation.Decode(), uop.Hints
}

// Internal representation of a user operation
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

func (u *UserOperation) EncodeToRaw() *UserOperationRaw {
	return &UserOperationRaw{
		From:         u.From,
		To:           u.To,
		Value:        (*hexutil.Big)(u.Value),
		Gas:          (*hexutil.Big)(u.Gas),
		MaxFeePerGas: (*hexutil.Big)(u.MaxFeePerGas),
		Nonce:        (*hexutil.Big)(u.Nonce),
		Deadline:     (*hexutil.Big)(u.Deadline),
		Dapp:         u.Dapp,
		Control:      u.Control,
		CallConfig:   (*hexutil.Big)(big.NewInt(int64(u.CallConfig))),
		SessionKey:   u.SessionKey,
		Data:         hexutil.Bytes(u.Data),
		Signature:    hexutil.Bytes(u.Signature),
	}
}

func (u *UserOperation) Validate(ethClient *ethclient.Client, atlas common.Address, eip712Domain *apitypes.TypedDataDomain, gasLimit *big.Int) *relayerror.Error {
	if u.To != atlas {
		return ErrUserOpInvalidToField
	}

	if u.Gas.Cmp(gasLimit) > 0 {
		return ErrUserOpGasLimitExceeded
	}

	currentBlock, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		log.Info("failed to get current block number", "err", err)
		return relayerror.ErrServerInternal
	}

	if u.Deadline.Uint64() <= currentBlock {
		return ErrUserOpDeadlineExceeded
	}

	// Check the signature only if it's set
	if len(u.Signature) > 0 {
		if relayErr := u.checkSignature(eip712Domain); relayErr != nil {
			return relayErr
		}
	}

	return nil
}

func (u *UserOperation) Hash(trusted bool, eip712Domain *apitypes.TypedDataDomain) (common.Hash, *relayerror.Error) {
	hash, _, err := apitypes.TypedDataAndHash(apitypes.TypedData{
		Types:       u.toTypedDataTypes(trusted),
		PrimaryType: "UserOperation",
		Domain:      *eip712Domain,
		Message:     u.toTypedDataMessage(trusted),
	})

	if err != nil {
		return common.Hash{}, ErrUserOpComputeHash.AddError(err)
	}

	return common.BytesToHash(hash), nil
}

func (u *UserOperation) AbiEncode() ([]byte, error) {
	return userOpArgs.Pack(&u)
}

func (u *UserOperation) toTypedDataTypes(trusted bool) apitypes.Types {
	t := apitypes.Types{
		"EIP712Domain": []apitypes.Type{
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
	}

	if trusted {
		t["UserOperation"] = []apitypes.Type{
			{Name: "from", Type: "address"},
			{Name: "to", Type: "address"},
			{Name: "dapp", Type: "address"},
			{Name: "control", Type: "address"},
			{Name: "callConfig", Type: "uint32"},
			{Name: "sessionKey", Type: "address"},
		}
	} else {
		t["UserOperation"] = []apitypes.Type{
			{Name: "from", Type: "address"},
			{Name: "to", Type: "address"},
			{Name: "value", Type: "uint256"},
			{Name: "gas", Type: "uint256"},
			{Name: "maxFeePerGas", Type: "uint256"},
			{Name: "nonce", Type: "uint256"},
			{Name: "deadline", Type: "uint256"},
			{Name: "dapp", Type: "address"},
			{Name: "control", Type: "address"},
			{Name: "callConfig", Type: "uint32"},
			{Name: "sessionKey", Type: "address"},
			{Name: "data", Type: "bytes"},
		}
	}

	return t
}

func (u *UserOperation) toTypedDataMessage(trusted bool) apitypes.TypedDataMessage {
	if trusted {
		return apitypes.TypedDataMessage{
			"from":       u.From.Hex(),
			"to":         u.To.Hex(),
			"dapp":       u.Dapp.Hex(),
			"control":    u.Control.Hex(),
			"callConfig": big.NewInt(int64(u.CallConfig)),
			"sessionKey": u.SessionKey.Hex(),
		}
	}

	return apitypes.TypedDataMessage{
		"from":         u.From.Hex(),
		"to":           u.To.Hex(),
		"value":        new(big.Int).Set(u.Value),
		"gas":          new(big.Int).Set(u.Gas),
		"maxFeePerGas": new(big.Int).Set(u.MaxFeePerGas),
		"nonce":        new(big.Int).Set(u.Nonce),
		"deadline":     new(big.Int).Set(u.Deadline),
		"dapp":         u.Dapp.Hex(),
		"control":      u.Control.Hex(),
		"callConfig":   big.NewInt(int64(u.CallConfig)),
		"sessionKey":   u.SessionKey.Hex(),
		"data":         u.Data,
	}
}

func (u *UserOperation) checkSignature(eip712Domain *apitypes.TypedDataDomain) *relayerror.Error {
	if len(u.Signature) != 65 {
		log.Info("invalid user operation signature length", "length", len(u.Signature))
		return ErrUserOpInvalidSignature
	}

	trustedOpHash := utils.FlagTrustedOpHash(u.CallConfig)

	userOpHash, relayErr := u.Hash(trustedOpHash, eip712Domain)
	if relayErr != nil {
		log.Info("failed to compute user operation hash", "err", relayErr.Message)
		return relayErr
	}

	signer, err := utils.RecoverSigner(userOpHash.Bytes(), u.Signature)
	if err != nil {
		log.Info("failed to recover user public key", "err", err)
		return ErrUserOpSignatureInvalid.AddError(err)
	}

	if signer != u.From {
		log.Info("invalid user operation signature", "signer", signer.Hex(), "from", u.From.Hex())
		return ErrUserOpInvalidSignature
	}

	return nil
}

type UserOperationPartialRaw struct {
	UserOpHash   common.Hash    `json:"userOpHash"`
	To           common.Address `json:"to"`
	Gas          *hexutil.Big   `json:"gas"`
	MaxFeePerGas *hexutil.Big   `json:"maxFeePerGas"`
	Deadline     *hexutil.Big   `json:"deadline"`
	Dapp         common.Address `json:"dapp"`
	Control      common.Address `json:"control"`

	//Exactly one of 1. Hints  2. (Value, Data, From) must be set
	Hints []common.Address `json:"hints,omitempty"`

	Value *hexutil.Big   `json:"value,omitempty"`
	Data  hexutil.Bytes  `json:"data,omitempty"`
	From  common.Address `json:"from,omitempty"`
}

func NewUserOperationPartialRaw(userOpHash common.Hash, userOp *UserOperation, hints []common.Address) *UserOperationPartialRaw {
	userOpPartial := &UserOperationPartialRaw{
		UserOpHash:   userOpHash,
		To:           userOp.To,
		Gas:          (*hexutil.Big)(userOp.Gas),
		MaxFeePerGas: (*hexutil.Big)(userOp.MaxFeePerGas),
		Deadline:     (*hexutil.Big)(userOp.Deadline),
		Dapp:         userOp.Dapp,
		Control:      userOp.Control,
	}

	if len(hints) > 0 {
		//randomize hints
		rand.Shuffle(len(hints), func(i, j int) { hints[i], hints[j] = hints[j], hints[i] })
		userOpPartial.Hints = hints
	} else {
		userOpPartial.Data = hexutil.Bytes(userOp.Data)
		userOpPartial.From = userOp.From
		userOpPartial.Value = (*hexutil.Big)(userOp.Value)
	}

	return userOpPartial
}
