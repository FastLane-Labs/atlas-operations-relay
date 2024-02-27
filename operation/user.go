package operation

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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

	userOpArgs = abi.Arguments{
		{Type: userOpSolType, Name: "userOperation"},
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
