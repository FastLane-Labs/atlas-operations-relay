package operation

import (
	"math/big"

	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/FastLane-Labs/atlas-operations-relay/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

var (
	ErrSolverOpInvalidToField      = relayerror.NewError(2100, "solver operation's 'to' field must be atlas contract address")
	ErrSolverOpGasLimitExceeded    = relayerror.NewError(2101, "solver operation's gas limit exceeded")
	ErrSolverOpMaxFeePerGasTooLow  = relayerror.NewError(2102, "solver operation's maxFeePerGas must be equal or higher the user operation")
	ErrSolverOpDeadlineTooLow      = relayerror.NewError(2103, "solver operation's deadline exceeded or lower than user operation's")
	ErrSolverOpDAppControlMismatch = relayerror.NewError(2104, "solver operation's dApp control does not match the user operation's")
	ErrSolverOpSignatureInvalid    = relayerror.NewError(2106, "solver operation has invalid signature")
	ErrSolverOpComputeHash         = relayerror.NewError(2107, "failed to compute solver operation hash")
)

var (
	solverOpStructArgs = []abi.ArgumentMarshaling{
		{Name: "from", Type: "address", InternalType: "address"},
		{Name: "to", Type: "address", InternalType: "address"},
		{Name: "value", Type: "uint256", InternalType: "uint256"},
		{Name: "gas", Type: "uint256", InternalType: "uint256"},
		{Name: "maxFeePerGas", Type: "uint256", InternalType: "uint256"},
		{Name: "deadline", Type: "uint256", InternalType: "uint256"},
		{Name: "solver", Type: "address", InternalType: "address"},
		{Name: "control", Type: "address", InternalType: "address"},
		{Name: "userOpHash", Type: "bytes32", InternalType: "bytes32"},
		{Name: "bidToken", Type: "address", InternalType: "address"},
		{Name: "bidAmount", Type: "uint256", InternalType: "uint256"},
		{Name: "data", Type: "bytes", InternalType: "bytes"},
		{Name: "signature", Type: "bytes", InternalType: "bytes"},
	}

	solverOpSolType, _      = abi.NewType("tuple", "struct SolverOperation", solverOpStructArgs)
	solverOpArraySolType, _ = abi.NewType("tuple[]", "struct SolverOperation[]", solverOpStructArgs)

	solverOpArgs = abi.Arguments{
		{Type: solverOpSolType, Name: "solverOperation"},
	}

	solverOpArrayArgs = abi.Arguments{
		{Type: solverOpArraySolType, Name: "solverOperations"},
	}
)

// External representation of a solver operation,
// the relay receives and broadcasts solver operations in this format
type SolverOperationRaw struct {
	From         common.Address `json:"from" validate:"required"`
	To           common.Address `json:"to" validate:"required"`
	Value        *hexutil.Big   `json:"value" validate:"required"`
	Gas          *hexutil.Big   `json:"gas" validate:"required"`
	MaxFeePerGas *hexutil.Big   `json:"maxFeePerGas" validate:"required"`
	Deadline     *hexutil.Big   `json:"deadline" validate:"required"`
	Solver       common.Address `json:"solver" validate:"required"`
	Control      common.Address `json:"control" validate:"required"`
	UserOpHash   common.Hash    `json:"userOpHash" validate:"required"`
	BidToken     common.Address `json:"bidToken"` // Optional (address(0) = ETH)
	BidAmount    *hexutil.Big   `json:"bidAmount" validate:"required"`
	Data         hexutil.Bytes  `json:"data" validate:"required"`
	Signature    hexutil.Bytes  `json:"signature" validate:"required"`
}

func (s *SolverOperationRaw) Decode() *SolverOperation {
	return &SolverOperation{
		From:         s.From,
		To:           s.To,
		Value:        s.Value.ToInt(),
		Gas:          s.Gas.ToInt(),
		MaxFeePerGas: s.MaxFeePerGas.ToInt(),
		Deadline:     s.Deadline.ToInt(),
		Solver:       s.Solver,
		Control:      s.Control,
		UserOpHash:   s.UserOpHash,
		BidToken:     s.BidToken,
		BidAmount:    s.BidAmount.ToInt(),
		Data:         s.Data,
		Signature:    s.Signature,
	}
}

type SolverOperationWithScoreRaw struct {
	SolverOperation *SolverOperationRaw `json:"solverOperation"`
	Score           int                 `json:"score"`
}

type SolverOperationWithScore struct {
	SolverOpHash    common.Hash
	SolverOperation *SolverOperation
	Score           int
}

func (sop *SolverOperationWithScore) EncodeToRaw() *SolverOperationWithScoreRaw {
	return &SolverOperationWithScoreRaw{
		SolverOperation: sop.SolverOperation.EncodeToRaw(),
		Score:           sop.Score,
	}
}

type SolverOperationsWithScore []*SolverOperationWithScore

func (sops SolverOperationsWithScore) EncodeToRaw() []*SolverOperationWithScoreRaw {
	sopsRaw := make([]*SolverOperationWithScoreRaw, len(sops))
	for i, sop := range sops {
		sopsRaw[i] = sop.EncodeToRaw()
	}
	return sopsRaw
}

// Internal representation of a solver operation
type SolverOperation struct {
	From         common.Address
	To           common.Address
	Value        *big.Int
	Gas          *big.Int
	MaxFeePerGas *big.Int
	Deadline     *big.Int
	Solver       common.Address
	Control      common.Address
	UserOpHash   common.Hash
	BidToken     common.Address
	BidAmount    *big.Int
	Data         []byte
	Signature    []byte
}

func (s *SolverOperation) EncodeToRaw() *SolverOperationRaw {
	return &SolverOperationRaw{
		From:         s.From,
		To:           s.To,
		Value:        (*hexutil.Big)(s.Value),
		Gas:          (*hexutil.Big)(s.Gas),
		MaxFeePerGas: (*hexutil.Big)(s.MaxFeePerGas),
		Deadline:     (*hexutil.Big)(s.Deadline),
		Solver:       s.Solver,
		Control:      s.Control,
		UserOpHash:   s.UserOpHash,
		BidToken:     s.BidToken,
		BidAmount:    (*hexutil.Big)(s.BidAmount),
		Data:         hexutil.Bytes(s.Data),
		Signature:    hexutil.Bytes(s.Signature),
	}
}

func (s *SolverOperation) Validate(userOperation *UserOperation, atlas common.Address, eip712Domain *apitypes.TypedDataDomain, gasLimit uint32) *relayerror.Error {
	if s.To != atlas {
		return ErrSolverOpInvalidToField
	}

	if s.Gas.Cmp(big.NewInt(int64(gasLimit))) > 0 {
		return ErrSolverOpGasLimitExceeded
	}

	if s.MaxFeePerGas.Cmp(userOperation.MaxFeePerGas) < 0 {
		return ErrSolverOpMaxFeePerGasTooLow
	}

	if s.Deadline.Cmp(userOperation.Deadline) < 0 {
		return ErrSolverOpDeadlineTooLow
	}

	if s.Control != userOperation.Control {
		return ErrSolverOpDAppControlMismatch
	}

	relayErr := s.checkSignature(eip712Domain)
	if relayErr != nil {
		return relayErr
	}

	return nil
}

func (s *SolverOperation) Hash(eip712Domain *apitypes.TypedDataDomain) (common.Hash, *relayerror.Error) {
	hash, _, err := apitypes.TypedDataAndHash(apitypes.TypedData{
		Types:       s.toTypedDataTypes(),
		PrimaryType: "SolverOperation",
		Domain:      *eip712Domain,
		Message:     s.toTypedDataMessage(),
	})

	if err != nil {
		return common.Hash{}, ErrSolverOpComputeHash.AddError(err)
	}

	return common.BytesToHash(hash), nil
}

func (s *SolverOperation) AbiEncode() ([]byte, error) {
	return solverOpArgs.Pack(&s)
}

func (s *SolverOperation) toTypedDataTypes() apitypes.Types {
	return apitypes.Types{
		"EIP712Domain": []apitypes.Type{
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
		"SolverOperation": []apitypes.Type{
			{Name: "from", Type: "address"},
			{Name: "to", Type: "address"},
			{Name: "value", Type: "uint256"},
			{Name: "gas", Type: "uint256"},
			{Name: "maxFeePerGas", Type: "uint256"},
			{Name: "deadline", Type: "uint256"},
			{Name: "solver", Type: "address"},
			{Name: "control", Type: "address"},
			{Name: "userOpHash", Type: "bytes32"},
			{Name: "bidToken", Type: "address"},
			{Name: "bidAmount", Type: "uint256"},
			{Name: "data", Type: "bytes"},
		},
	}
}

func (s *SolverOperation) toTypedDataMessage() apitypes.TypedDataMessage {
	return apitypes.TypedDataMessage{
		"from":         s.From.Hex(),
		"to":           s.To.Hex(),
		"value":        new(big.Int).Set(s.Value),
		"gas":          new(big.Int).Set(s.Gas),
		"maxFeePerGas": new(big.Int).Set(s.MaxFeePerGas),
		"deadline":     new(big.Int).Set(s.Deadline),
		"solver":       s.Solver.Hex(),
		"control":      s.Control.Hex(),
		"userOpHash":   s.UserOpHash.Hex(),
		"bidToken":     s.BidToken.Hex(),
		"bidAmount":    new(big.Int).Set(s.BidAmount),
		"data":         s.Data,
	}
}

func (s *SolverOperation) checkSignature(eip712Domain *apitypes.TypedDataDomain) *relayerror.Error {
	if len(s.Signature) != 65 {
		log.Info("invalid solver operation signature length", "length", len(s.Signature))
		return ErrSolverOpSignatureInvalid
	}

	solverOpHash, relayErr := s.Hash(eip712Domain)
	if relayErr != nil {
		log.Info("failed to compute solver operation hash", "err", relayErr.Message)
		return relayErr
	}

	signer, err := utils.RecoverSigner(solverOpHash.Bytes(), s.Signature)
	if err != nil {
		log.Info("failed to recover solver public key", "err", err)
		return ErrSolverOpSignatureInvalid.AddError(err)
	}

	if signer != s.From {
		log.Info("invalid solver operation signature", "signer", signer.Hex(), "from", s.From.Hex())
		return ErrUserOpInvalidSignature
	}

	return nil
}

// Array of solver operations
type SolverOperations []*SolverOperation

func (s SolverOperations) AbiEncode() ([]byte, error) {
	solverOps := make([]SolverOperation, len(s))
	for i, solverOp := range s {
		solverOps[i] = *solverOp
	}
	return solverOpArrayArgs.Pack(solverOps)
}
