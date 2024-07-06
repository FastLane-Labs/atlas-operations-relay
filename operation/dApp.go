package operation

import (
	"fmt"
	"math/big"

	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/FastLane-Labs/atlas-operations-relay/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

var (
	ErrDAppOpFromDoesNotMatchSessionKey = relayerror.NewError(2200, "dApp operation's 'from' field does not match user operation's session key")
	ErrDAppOpInvalidToField             = relayerror.NewError(2201, "dApp operation's 'to' field must be atlas contract address")
	ErrDAppOpDeadlineTooLow             = relayerror.NewError(2202, "dApp operation's deadline exceeded or lower than user operation's")
	ErrDAppOpDAppControlMismatch        = relayerror.NewError(2203, "dApp operation's dApp control does not match the user operation's")
	ErrDAppOpUserOpHashMismatch         = relayerror.NewError(2204, "dApp operation's user operation hash does not match the user operation's")
	ErrDAppOpInvalidCallChainHash       = relayerror.NewError(2205, "dApp operation's call chain hash is invalid")
	ErrDappOpSignatureInvalid           = relayerror.NewError(2207, "dApp operation has invalid signature")
	ErrDAppOpGasLimitExceeded           = relayerror.NewError(2208, "dApp operation's gas limit exceeded")
	ErrDAppOpComputeHash                = relayerror.NewError(2209, "failed to compute dApp operation hash")
)

// External representation of a dApp operation,
// the relay receives and broadcasts dApp operations in this format
type DAppOperationRaw struct {
	From          common.Address `json:"from" validate:"required"`
	To            common.Address `json:"to" validate:"required"`
	Nonce         *hexutil.Big   `json:"nonce" validate:"required"`
	Deadline      *hexutil.Big   `json:"deadline" validate:"required"`
	Control       common.Address `json:"control" validate:"required"`
	Bundler       common.Address `json:"bundler"` // Optional (address(0) = any bundler)
	UserOpHash    common.Hash    `json:"userOpHash" validate:"required"`
	CallChainHash common.Hash    `json:"callChainHash" validate:"required"`
	Signature     hexutil.Bytes  `json:"signature" validate:"required"`
}

func (d *DAppOperationRaw) Decode() *DAppOperation {
	return &DAppOperation{
		From:          d.From,
		To:            d.To,
		Nonce:         d.Nonce.ToInt(),
		Deadline:      d.Deadline.ToInt(),
		Control:       d.Control,
		Bundler:       d.Bundler,
		UserOpHash:    d.UserOpHash,
		CallChainHash: d.CallChainHash,
		Signature:     d.Signature,
	}
}

// Internal representation of a dApp operation
type DAppOperation struct {
	From          common.Address
	To            common.Address
	Nonce         *big.Int
	Deadline      *big.Int
	Control       common.Address
	Bundler       common.Address
	UserOpHash    common.Hash
	CallChainHash common.Hash
	Signature     []byte
}

func GenerateSimulationDAppOperation(userOpHash common.Hash, userOp *UserOperation, solverOps []*SolverOperation) (*DAppOperation, error) {
	dAppOp := &DAppOperation{
		From:          userOp.SessionKey,
		To:            userOp.To,
		Nonce:         big.NewInt(0),
		Deadline:      userOp.Deadline,
		Control:       userOp.Control,
		Bundler:       common.Address{},
		UserOpHash:    userOpHash,
		CallChainHash: common.Hash{},
		Signature:     []byte(""),
	}

	if utils.FlagVerifyCallChainHash(userOp.CallConfig) {
		callChainHash, err := (&BundleOperations{UserOperation: userOp, SolverOperations: solverOps}).CallChainHash(userOp.CallConfig, userOp.Control)
		if err != nil {
			return nil, fmt.Errorf("failed to compute call chain hash: %w", err)
		}
		dAppOp.CallChainHash = callChainHash
	}
	return dAppOp, nil
}

func (d *DAppOperation) EncodeToRaw() *DAppOperationRaw {
	return &DAppOperationRaw{
		From:          d.From,
		To:            d.To,
		Nonce:         (*hexutil.Big)(d.Nonce),
		Deadline:      (*hexutil.Big)(d.Deadline),
		Control:       d.Control,
		Bundler:       d.Bundler,
		UserOpHash:    d.UserOpHash,
		CallChainHash: d.CallChainHash,
		Signature:     d.Signature,
	}
}

func (d *DAppOperation) Validate(userOpHash common.Hash, userOp *UserOperation, callChainHash common.Hash, atlas common.Address, eip712Domain *apitypes.TypedDataDomain, gasLimit *big.Int) *relayerror.Error {
	if userOp.SessionKey != (common.Address{}) && d.From != userOp.SessionKey {
		return ErrDAppOpFromDoesNotMatchSessionKey
	}

	if d.To != atlas {
		return ErrDAppOpInvalidToField
	}

	if d.Deadline.Cmp(userOp.Deadline) < 0 {
		return ErrDAppOpDeadlineTooLow
	}

	if d.Control != userOp.Control {
		return ErrDAppOpDAppControlMismatch
	}

	if d.UserOpHash != userOpHash {
		return ErrDAppOpUserOpHashMismatch
	}

	// Validate callChainHash only if it's provided
	if callChainHash != (common.Hash{}) && d.CallChainHash != callChainHash {
		return ErrDAppOpInvalidCallChainHash
	}

	relayErr := d.checkSignature(eip712Domain)
	if relayErr != nil {
		return relayErr
	}

	return nil
}

func (d *DAppOperation) Hash(eip712Domain *apitypes.TypedDataDomain) (common.Hash, *relayerror.Error) {
	hash, _, err := apitypes.TypedDataAndHash(apitypes.TypedData{
		Types:       d.toTypedDataTypes(),
		PrimaryType: "DAppOperation",
		Domain:      *eip712Domain,
		Message:     d.toTypedDataMessage(),
	})

	if err != nil {
		return common.Hash{}, ErrDAppOpComputeHash.AddError(err)
	}

	return common.BytesToHash(hash), nil
}

func (d *DAppOperation) toTypedDataTypes() apitypes.Types {
	return apitypes.Types{
		"EIP712Domain": []apitypes.Type{
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
		"DAppOperation": []apitypes.Type{
			{Name: "from", Type: "address"},
			{Name: "to", Type: "address"},
			{Name: "nonce", Type: "uint256"},
			{Name: "deadline", Type: "uint256"},
			{Name: "control", Type: "address"},
			{Name: "bundler", Type: "address"},
			{Name: "userOpHash", Type: "bytes32"},
			{Name: "callChainHash", Type: "bytes32"},
		},
	}
}

func (d *DAppOperation) toTypedDataMessage() apitypes.TypedDataMessage {
	return apitypes.TypedDataMessage{
		"from":          d.From.Hex(),
		"to":            d.To.Hex(),
		"nonce":         new(big.Int).Set(d.Nonce),
		"deadline":      new(big.Int).Set(d.Deadline),
		"control":       d.Control.Hex(),
		"bundler":       d.Bundler.Hex(),
		"userOpHash":    d.UserOpHash.Hex(),
		"callChainHash": d.CallChainHash.Hex(),
	}
}

func (d *DAppOperation) checkSignature(eip712Domain *apitypes.TypedDataDomain) *relayerror.Error {
	if len(d.Signature) != 65 {
		log.Info("invalid dappOp signature length", "length", len(d.Signature))
		return ErrDappOpSignatureInvalid
	}

	dAppOpHash, relayErr := d.Hash(eip712Domain)
	if relayErr != nil {
		log.Info("failed to compute dApp operation hash", "err", relayErr.Message)
		return relayErr
	}

	signer, err := utils.RecoverSigner(dAppOpHash.Bytes(), d.Signature)
	if err != nil {
		log.Info("failed to recover dApp public key", "err", err)
		return ErrDappOpSignatureInvalid.AddError(err)
	}

	if signer != d.From {
		log.Info("invalid dApp operation signature", "signer", signer.Hex(), "from", d.From.Hex())
		return ErrDappOpSignatureInvalid
	}

	return nil
}
