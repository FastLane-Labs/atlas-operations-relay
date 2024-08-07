package tests

import (
	"fmt"
	"math/big"

	"github.com/FastLane-Labs/atlas-operations-relay/contract/swapIntentDappControl"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	swapIntentSolType, _ = abi.NewType("tuple", "struct SwapIntent", []abi.ArgumentMarshaling{
		{Name: "tokenUserBuys", Type: "address", InternalType: "address"},
		{Name: "amountUserBuys", Type: "uint256", InternalType: "uint256"},
		{Name: "tokenUserSells", Type: "address", InternalType: "address"},
		{Name: "amountUserSells", Type: "uint256", InternalType: "uint256"},
		{Name: "auctionBaseCurrency", Type: "address", InternalType: "address"},
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
	TokenUserBuys       common.Address
	AmountUserBuys      *big.Int
	TokenUserSells      common.Address
	AmountUserSells     *big.Int
	AuctionBaseCurrency common.Address
	Conditions          []Condition
}

func (i *SwapIntent) abiEncode() ([]byte, error) {
	return swapIntentArgs.Pack(&i)
}

func (i *SwapIntent) abiEncodeWithSelector() ([]byte, error) {
	encoded, err := i.abiEncode()
	if err != nil {
		return nil, err
	}

	swapIntentDappControlAbi, err := swapIntentDappControl.SwapIntentDappControlMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get swapIntentDappControl ABI: %w", err)
	}
	methodName := "swap"
	method, exists := swapIntentDappControlAbi.Methods[methodName]
	if !exists {
		return nil, fmt.Errorf("method %s not found in swapIntentDappControl ABI", methodName)
	}
	selector := method.ID

	return append(selector, encoded...), nil
}

func swapIntentAbiDecode(input []byte) (*SwapIntent, error) {
	if len(input) < 4 {
		return nil, fmt.Errorf("input too short to contain selector")
	}

	data := input[4:]

	// Unpack data into swapIntent using the swapIntentArgs ABI definition.
	unpacked, err := swapIntentArgs.Unpack(data)
	if err != nil {
		return nil, fmt.Errorf("failed to unpack data: %w", err)
	}

	if len(unpacked) != 1 {
		return nil, fmt.Errorf("expected 1 unpacked value, got %d", len(unpacked))
	}

	decoded_struct := unpacked[0].(struct {
		TokenUserBuys       common.Address `json:"tokenUserBuys"`
		AmountUserBuys      *big.Int       `json:"amountUserBuys"`
		TokenUserSells      common.Address `json:"tokenUserSells"`
		AmountUserSells     *big.Int       `json:"amountUserSells"`
		AuctionBaseCurrency common.Address `json:"auctionBaseCurrency"`
		Conditions          []struct {
			Antecedent common.Address `json:"antecedent"`
			Context    []uint8        `json:"context"`
		} `json:"conditions"`
	})

	return &SwapIntent{
		TokenUserBuys:       decoded_struct.TokenUserBuys,
		AmountUserBuys:      decoded_struct.AmountUserBuys,
		TokenUserSells:      decoded_struct.TokenUserSells,
		AmountUserSells:     decoded_struct.AmountUserSells,
		AuctionBaseCurrency: decoded_struct.AuctionBaseCurrency,
		Conditions:          make([]Condition, len(decoded_struct.Conditions)),
	}, nil
}
