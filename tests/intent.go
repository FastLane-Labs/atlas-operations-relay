package tests

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	swapIntentFuncSelector = "83a6992a"

	swapIntentSolType, _ = abi.NewType("tuple", "struct SwapIntent", []abi.ArgumentMarshaling{
		{Name: "tokenUserBuys", Type: "address", InternalType: "address"},
		{Name: "amountUserBuys", Type: "uint256", InternalType: "uint256"},
		{Name: "tokenUserSells", Type: "address", InternalType: "address"},
		{Name: "amountUserSells", Type: "uint256", InternalType: "uint256"},
		{Name: "auctionBaseCurrency", Type: "address", InternalType: "address"},
		{Name: "solverMustReimburseGas", Type: "bool", InternalType: "bool"},
		{
			Name:         "conditions",
			Type:         "tuple[]",
			InternalType: "struct Condition[]",
			Components: []abi.ArgumentMarshaling{
				{Name: "antecedent", Type: "address", InternalType: "address"},
				{Name: "context", Type: "bytes", InternalType: "bytes"},
			},
		},
	})

	swapIntentArgs = abi.Arguments{
		{Type: swapIntentSolType, Name: "swapIntent"},
	}
)

type Condition struct {
	Antecedent common.Address
	Context    []byte
}

type SwapIntent struct {
	TokenUserBuys          common.Address
	AmountUserBuys         *big.Int
	TokenUserSells         common.Address
	AmountUserSells        *big.Int
	AuctionBaseCurrency    common.Address
	SolverMustReimburseGas bool
	Conditions             []Condition
}

func (i *SwapIntent) abiEncode() ([]byte, error) {
	return swapIntentArgs.Pack(&i)
}

func (i *SwapIntent) abiEncodeWithSelector() ([]byte, error) {
	encoded, err := i.abiEncode()
	if err != nil {
		return nil, err
	}

	selector := common.Hex2Bytes(swapIntentFuncSelector)
	return append(selector, encoded...), nil
}
