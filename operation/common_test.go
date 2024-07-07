package operation

import (
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

var (
	Eip712DomainTest = apitypes.TypedDataDomain{
		Name:              "AtlasVerification",
		Version:           "1.0",
		ChainId:           math.NewHexOrDecimal256(1),
		VerifyingContract: "0x8Be503bcdEd90ED42Eff31f56199399B2b0154CA",
	}
)
