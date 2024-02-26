package operation

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
		{Name: "dApp", Type: "address", InternalType: "address"},
		{Name: "control", Type: "address", InternalType: "address"},
		{Name: "sessionKey", Type: "address", InternalType: "address"},
		{Name: "data", Type: "bytes", InternalType: "bytes"},
		{Name: "signature", Type: "bytes", InternalType: "bytes"},
	})

	proofHashSolType, _ = abi.NewType("tuple", "struct ProofHash", []abi.ArgumentMarshaling{
		{Name: "userTypeHash", Type: "bytes32", InternalType: "bytes32"},
		{Name: "from", Type: "address", InternalType: "address"},
		{Name: "to", Type: "address", InternalType: "address"},
		{Name: "value", Type: "uint256", InternalType: "uint256"},
		{Name: "gas", Type: "uint256", InternalType: "uint256"},
		{Name: "maxFeePerGas", Type: "uint256", InternalType: "uint256"},
		{Name: "nonce", Type: "uint256", InternalType: "uint256"},
		{Name: "deadline", Type: "uint256", InternalType: "uint256"},
		{Name: "dApp", Type: "address", InternalType: "address"},
		{Name: "control", Type: "address", InternalType: "address"},
		{Name: "sessionKey", Type: "address", InternalType: "address"},
		{Name: "data", Type: "bytes32", InternalType: "bytes32"},
	})

	userOpArgs = abi.Arguments{
		{Type: userOpSolType, Name: "userOperation"},
	}

	proofHashArgs = abi.Arguments{
		{Type: proofHashSolType, Name: "proofHash"},
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

func (u *UserOperation) Hash() (common.Hash, error) {
	packed, err := userOpArgs.Pack(&u)
	if err != nil {
		return common.Hash{}, err
	}
	return common.BytesToHash(packed), nil
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

	packed, err := proofHashArgs.Pack(&proofHash)
	if err != nil {
		return common.Hash{}, err
	}
	return crypto.Keccak256Hash(packed), nil
}
