package operation

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type SolverOperation struct {
	From         common.Address `json:"from"`
	To           common.Address `json:"to"`
	Value        *big.Int       `json:"value"`
	Gas          *big.Int       `json:"gas"`
	MaxFeePerGas *big.Int       `json:"maxFeePerGas"`
	Deadline     *big.Int       `json:"deadline"`
	Solver       common.Address `json:"solver"`
	Control      common.Address `json:"control"`
	UserOpHash   common.Hash    `json:"userOpHash"`
	BidToken     common.Address `json:"bidToken"`
	BidAmount    *big.Int       `json:"bidAmount"`
	Data         []byte         `json:"data"`
	Signature    []byte         `json:"signature"`
}
