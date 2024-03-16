package operation

import (
	"context"
	"math/big"

	relayCrypto "github.com/FastLane-Labs/atlas-operations-relay/crypto"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrUserOpInvalidSignature = relayerror.NewError(2001, "user operation has invalid signature")
	ErrUserOpInvalidToField   = relayerror.NewError(2002, "user operation's 'to' field must be atlas contract address")
	ErrUserOpDeadlineExceeded = relayerror.NewError(2003, "user operation's deadline exceeded")
	ErrUserOpComputeHash      = relayerror.NewError(2004, "failed to compute user operation hash")
	ErrUserOpComputeProofHash = relayerror.NewError(2005, "failed to compute user proof hash")
	ErrUserOpSignatureInvalid = relayerror.NewError(2006, "user operation has invalid signature")
	ErrUserOpGasLimitExceeded = relayerror.NewError(2007, "user operation's gas limit exceeded")
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

// External representation of a user operation,
// the relay receives and broadcasts user operations in this format
type UserOperationRaw struct {
	From         common.Address `json:"from"`
	To           common.Address `json:"to"`
	Value        *hexutil.Big   `json:"value"`
	Gas          hexutil.Uint64 `json:"gas"`
	MaxFeePerGas *hexutil.Big   `json:"maxFeePerGas"`
	Nonce        hexutil.Uint64 `json:"nonce"`
	Deadline     hexutil.Uint64 `json:"deadline"`
	Dapp         common.Address `json:"dapp"`
	Control      common.Address `json:"control"`
	SessionKey   common.Address `json:"sessionKey"`
	Data         hexutil.Bytes  `json:"data"`
	Signature    hexutil.Bytes  `json:"signature"`
}

func (u *UserOperationRaw) Decode() *UserOperation {
	return &UserOperation{
		From:         u.From,
		To:           u.To,
		Value:        u.Value.ToInt(),
		Gas:          new(big.Int).SetUint64(uint64(u.Gas)),
		MaxFeePerGas: u.MaxFeePerGas.ToInt(),
		Nonce:        new(big.Int).SetUint64(uint64(u.Nonce)),
		Deadline:     new(big.Int).SetUint64(uint64(u.Deadline)),
		Dapp:         u.Dapp,
		Control:      u.Control,
		SessionKey:   u.SessionKey,
		Data:         u.Data,
		Signature:    u.Signature,
	}
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
	SessionKey   common.Address
	Data         []byte
	Signature    []byte
}

func (u *UserOperation) EncodeToRaw() *UserOperationRaw {
	return &UserOperationRaw{
		From:         u.From,
		To:           u.To,
		Value:        (*hexutil.Big)(u.Value),
		Gas:          hexutil.Uint64(u.Gas.Uint64()),
		MaxFeePerGas: (*hexutil.Big)(u.MaxFeePerGas),
		Nonce:        hexutil.Uint64(u.Nonce.Uint64()),
		Deadline:     hexutil.Uint64(u.Deadline.Uint64()),
		Dapp:         u.Dapp,
		Control:      u.Control,
		SessionKey:   u.SessionKey,
		Data:         hexutil.Bytes(u.Data),
		Signature:    hexutil.Bytes(u.Signature),
	}
}

func (u *UserOperation) Validate(ethClient *ethclient.Client, atlas common.Address, atlasDomainSeparator common.Hash, gasLimit *big.Int) *relayerror.Error {
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

	relayErr := u.checkSignature(atlasDomainSeparator)
	if relayErr != nil {
		return relayErr
	}

	return nil
}

func (u *UserOperation) Hash() (common.Hash, *relayerror.Error) {
	packed, err := u.AbiEncode()
	if err != nil {
		return common.Hash{}, ErrUserOpComputeHash.AddError(err)
	}
	return crypto.Keccak256Hash(packed), nil
}

func (u *UserOperation) AbiEncode() ([]byte, error) {
	return userOpArgs.Pack(&u)
}

func (u *UserOperation) ProofHash() (common.Hash, error) {
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
	proofHash, err := u.ProofHash()
	if err != nil {
		log.Info("failed to compute user proof hash", "err", err)
		return ErrUserOpComputeProofHash.AddError(err)
	}

	signer, err := relayCrypto.RecoverEip712Signer(domainSeparator, proofHash, u.Signature)
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
