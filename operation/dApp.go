package operation

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type DAppOperation struct {
	From          common.Address `json:"from"`
	To            common.Address `json:"to"`
	Value         *big.Int       `json:"value"`
	Gas           *big.Int       `json:"gas"`
	MaxFeePerGas  *big.Int       `json:"maxFeePerGas"`
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
		Control: userOp.Control,
	}
}
