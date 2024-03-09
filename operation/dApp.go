package operation

import (
	"math/big"

	relayCrypto "github.com/FastLane-Labs/atlas-operations-relay/crypto"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ErrDAppOpFromDoesNotMatchSessionKey = relayerror.NewError(3203, "dApp operation's 'from' field does not match user operation's session key")
	ErrDAppOpInvalidToField             = relayerror.NewError(3204, "dApp operation's 'to' field must be atlas contract address")
	ErrDAppOpDeadlineTooLow             = relayerror.NewError(3205, "dApp operation's deadline exceeded or lower than user operation's")
	ErrDAppOpDAppControlMismatch        = relayerror.NewError(3206, "dApp operation's dApp control does not match the user operation's")
	ErrDAppOpUserOpHashMismatch         = relayerror.NewError(3207, "dApp operation's user operation hash does not match the user operation's")
	ErrDAppOpInvalidCallChainHash       = relayerror.NewError(3208, "dApp operation's call chain hash is invalid")
	ErrDAppOpComputeProofHash           = relayerror.NewError(3209, "failed to compute dApp proof hash")
	ErrDappOpSignatureInvalid           = relayerror.NewError(3210, "dApp operation has invalid signature")
	ErrDAppOpGasLimitExceeded           = relayerror.NewError(3211, "dApp operation's gas limit exceeded")
)

var (
	dAppGasLimit   = big.NewInt(1000000)
	DAPP_TYPE_HASH = crypto.Keccak256Hash([]byte("DAppApproval(address from,address to,uint256 value,uint256 gas,uint256 nonce,uint256 deadline,address control,address bundler,bytes32 userOpHash,bytes32 callChainHash)"))
)

var (
	dAppProofHashSolType, _ = abi.NewType("tuple", "struct DappProofHash", []abi.ArgumentMarshaling{
		{Name: "dAppTypeHash", Type: "bytes32", InternalType: "bytes32"},
		{Name: "from", Type: "address", InternalType: "address"},
		{Name: "to", Type: "address", InternalType: "address"},
		{Name: "value", Type: "uint256", InternalType: "uint256"},
		{Name: "gas", Type: "uint256", InternalType: "uint256"},
		{Name: "nonce", Type: "uint256", InternalType: "uint256"},
		{Name: "deadline", Type: "uint256", InternalType: "uint256"},
		{Name: "control", Type: "address", InternalType: "address"},
		{Name: "userOpHash", Type: "bytes32", InternalType: "bytes32"},
		{Name: "callChainHash", Type: "bytes32", InternalType: "bytes32"},
	})

	dAppProofHashArgs = abi.Arguments{
		{Type: dAppProofHashSolType, Name: "proofHash"},
	}
)

type DAppOperation struct {
	From          common.Address `json:"from"`
	To            common.Address `json:"to"`
	Value         *big.Int       `json:"value"`
	Gas           *big.Int       `json:"gas"`
	Nonce         *big.Int       `json:"nonce"`
	Deadline      *big.Int       `json:"deadline"`
	Control       common.Address `json:"control"`
	Bundler       common.Address `json:"bundler"`
	UserOpHash    common.Hash    `json:"userOpHash"`
	CallChainHash common.Hash    `json:"callChainHash"`
	Signature     []byte         `json:"signature"`
}

func GenerateSimulationDAppOperation(userOp *UserOperation) *DAppOperation {
	return &DAppOperation{
		From:          common.HexToAddress("0x0"),
		To:            common.HexToAddress("0x0"),
		Value:         big.NewInt(0),
		Gas:           big.NewInt(100000),
		Nonce:         big.NewInt(0),
		Deadline:      userOp.Deadline,
		Control:       userOp.Control,
		Bundler:       common.HexToAddress("0x0"),
		UserOpHash:    common.HexToHash("0x0"),
		CallChainHash: common.HexToHash("0x0"),
		Signature:     []byte(""),
	}
}

func (d *DAppOperation) Validate(userOpHash common.Hash, userOp *UserOperation, callChainHash common.Hash, atlas common.Address, atlasDomainSeparator common.Hash, gasLimit *big.Int) *relayerror.Error {
	if userOp.SessionKey != (common.Address{}) && d.From != userOp.SessionKey {
		return ErrDAppOpFromDoesNotMatchSessionKey
	}

	if d.To != atlas {
		return ErrDAppOpInvalidToField
	}

	enforcedGasLimit := new(big.Int).Set(dAppGasLimit)
	if gasLimit != nil && gasLimit.Cmp(common.Big0) > 0 {
		enforcedGasLimit = gasLimit
	}

	if d.Gas.Cmp(enforcedGasLimit) > 0 {
		return ErrDAppOpGasLimitExceeded
	}

	if d.Deadline.Cmp(userOp.Deadline) < 0 {
		return ErrDAppOpDeadlineTooLow
	}

	if d.Control != userOp.To {
		return ErrDAppOpDAppControlMismatch
	}

	if d.UserOpHash != userOpHash {
		return ErrDAppOpUserOpHashMismatch
	}

	if d.CallChainHash != callChainHash {
		return ErrDAppOpInvalidCallChainHash
	}

	relayErr := d.checkSignature(atlasDomainSeparator)
	if relayErr != nil {
		return relayErr
	}

	return nil
}

func (d *DAppOperation) proofHash() (common.Hash, error) {
	proofHash := struct {
		DAppTypeHash  common.Hash
		From          common.Address
		To            common.Address
		Value         *big.Int
		Gas           *big.Int
		Nonce         *big.Int
		Deadline      *big.Int
		Control       common.Address
		UserOpHash    common.Hash
		CallChainHash common.Hash
	}{
		DAPP_TYPE_HASH,
		d.From,
		d.To,
		d.Value,
		d.Gas,
		d.Nonce,
		d.Deadline,
		d.Control,
		d.UserOpHash,
		d.CallChainHash,
	}

	packed, err := dAppProofHashArgs.Pack(&proofHash)
	if err != nil {
		return common.Hash{}, err
	}
	return crypto.Keccak256Hash(packed), nil
}

func (d *DAppOperation) checkSignature(domainSeparator common.Hash) *relayerror.Error {
	proofHash, err := d.proofHash()
	if err != nil {
		log.Info("failed to compute dApp proof hash", "err", err)
		return ErrDAppOpComputeProofHash.AddError(err)
	}

	signer, err := relayCrypto.RecoverEip712Signer(domainSeparator, proofHash, d.Signature)
	if err != nil {
		log.Info("failed to recover dApp public key", "err", err)
		return ErrDappOpSignatureInvalid.AddError(err)
	}

	if signer != d.From {
		log.Info("invalid dApp operation signature", "signer", signer.Hex(), "from", d.From.Hex())
		return ErrUserOpInvalidSignature
	}

	return nil
}
