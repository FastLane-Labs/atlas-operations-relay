package operation

import (
	"context"
	"math/big"

	relayCrypto "github.com/FastLane-Labs/atlas-operations-relay/crypto"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrUserOpInvalidSignature = relayerror.NewError(5001, "user operation has invalid signature")
	ErrUserOpInvalidToField   = relayerror.NewError(5002, "user operation's 'to' field must be atlas contract address")
	ErrUserOpDeadlineExceeded = relayerror.NewError(5003, "user operation's deadline exceeded")
	ErrUserOpComputeHash      = relayerror.NewError(5004, "failed to compute user operation hash")
	ErrUserOpComputeProofHash = relayerror.NewError(5005, "failed to compute user proof hash")
	ErrUserOpSignatureInvalid = relayerror.NewError(5006, "user operation has invalid signature")
)

var (
	USER_TYPE_HASH = crypto.Keccak256Hash([]byte("UserOperation(address from,address to,uint256 value,uint256 gas,uint256 maxFeePerGas,uint256 nonce,uint256 deadline,address dapp,address control,address sessionKey,bytes32 data)"))
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
		{Name: "sessionKey", Type: "address", InternalType: "address"},
		{Name: "data", Type: "bytes", InternalType: "bytes"},
		{Name: "signature", Type: "bytes", InternalType: "bytes"},
	})

	userProofHashSolType, _ = abi.NewType("tuple", "struct UserProofHash", []abi.ArgumentMarshaling{
		{Name: "userTypeHash", Type: "bytes32", InternalType: "bytes32"},
		{Name: "from", Type: "address", InternalType: "address"},
		{Name: "to", Type: "address", InternalType: "address"},
		{Name: "value", Type: "uint256", InternalType: "uint256"},
		{Name: "gas", Type: "uint256", InternalType: "uint256"},
		{Name: "maxFeePerGas", Type: "uint256", InternalType: "uint256"},
		{Name: "nonce", Type: "uint256", InternalType: "uint256"},
		{Name: "deadline", Type: "uint256", InternalType: "uint256"},
		{Name: "dapp", Type: "address", InternalType: "address"},
		{Name: "control", Type: "address", InternalType: "address"},
		{Name: "sessionKey", Type: "address", InternalType: "address"},
		{Name: "data", Type: "bytes32", InternalType: "bytes32"},
	})

	userOpArgs = abi.Arguments{
		{Type: userOpSolType, Name: "userOperation"},
	}

	userProofHashArgs = abi.Arguments{
		{Type: userProofHashSolType, Name: "proofHash"},
	}
)

type UserOperation struct {
	From         common.Address `json:"from"`
	To           common.Address `json:"to"`
	Value        *big.Int       `json:"value"`
	Gas          *big.Int       `json:"gas"`
	MaxFeePerGas *big.Int       `json:"maxFeePerGas"`
	Nonce        *big.Int       `json:"nonce"`
	Deadline     *big.Int       `json:"deadline"`
	Dapp         common.Address `json:"dapp"`
	Control      common.Address `json:"control"`
	SessionKey   common.Address `json:"sessionKey"`
	Data         []byte         `json:"data"`
	Signature    []byte         `json:"signature"`
}

func (u *UserOperation) Validate(ethClient *ethclient.Client, atlas common.Address, atlasDomainSeparator common.Hash) *relayerror.Error {
	if u.To != atlas {
		return ErrUserOpInvalidToField
	}

	// gas limit check?

	currentBlock, err := ethClient.BlockNumber(context.Background())
	if err != nil {
		log.Info("failed to get current block number", "err", err)
		return relayerror.ErrServerInternal
	}

	if u.Deadline.Uint64() <= currentBlock {
		return ErrUserOpDeadlineExceeded
	}

	relayErr := u.checkSignature(atlasDomainSeparator)
	if relayErr != nil {
		return relayErr
	}

	return nil
}

func (u *UserOperation) Hash() (common.Hash, *relayerror.Error) {
	packed, err := u.abiEncode()
	if err != nil {
		return common.Hash{}, ErrUserOpComputeHash.AddError(err)
	}
	return crypto.Keccak256Hash(packed), nil
}

func (u *UserOperation) abiEncode() ([]byte, error) {
	return userOpArgs.Pack(&u)
}

func (u *UserOperation) proofHash() (common.Hash, error) {
	proofHash := struct {
		UserTypeHash common.Hash
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
		Data         common.Hash
	}{
		USER_TYPE_HASH,
		u.From,
		u.To,
		u.Value,
		u.Gas,
		u.MaxFeePerGas,
		u.Nonce,
		u.Deadline,
		u.Dapp,
		u.Control,
		u.SessionKey,
		crypto.Keccak256Hash(u.Data),
	}

	packed, err := userProofHashArgs.Pack(&proofHash)
	if err != nil {
		return common.Hash{}, err
	}
	return crypto.Keccak256Hash(packed), nil
}

func (u *UserOperation) checkSignature(domainSeparator common.Hash) *relayerror.Error {
	proofHash, err := u.proofHash()
	if err != nil {
		log.Info("failed to compute user proof hash", "err", err)
		return ErrUserOpComputeProofHash.AddError(err)
	}

	signer, err := relayCrypto.GetSigner(domainSeparator, proofHash, u.Signature)
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
