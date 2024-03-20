package operation

import (
	"math/big"

	relayCrypto "github.com/FastLane-Labs/atlas-operations-relay/crypto"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ErrSolverOpInvalidToField      = relayerror.NewError(2100, "solver operation's 'to' field must be atlas contract address")
	ErrSolverOpGasLimitExceeded    = relayerror.NewError(2101, "solver operation's gas limit exceeded")
	ErrSolverOpMaxFeePerGasTooLow  = relayerror.NewError(2102, "solver operation's maxFeePerGas must be equal or higher the user operation")
	ErrSolverOpDeadlineTooLow      = relayerror.NewError(2103, "solver operation's deadline exceeded or lower than user operation's")
	ErrSolverOpDAppControlMismatch = relayerror.NewError(2104, "solver operation's dApp control does not match the user operation's")
	ErrSolverOpComputeProofHash    = relayerror.NewError(2105, "failed to compute solver proof hash")
	ErrSolverOpSignatureInvalid    = relayerror.NewError(2106, "solver operation has invalid signature")
)

var (
	SOLVER_TYPE_HASH = crypto.Keccak256Hash([]byte("SolverOperation(address from,address to,uint256 value,uint256 gas,uint256 maxFeePerGas,uint256 deadline,address dapp,address control,bytes32 userOpHash,address bidToken,uint256 bidAmount,bytes32 data)"))
)

var (
	solverOpSolType, _ = abi.NewType("tuple", "struct SolverOperation", []abi.ArgumentMarshaling{
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
	})

	solverProofHashSolType, _ = abi.NewType("tuple", "struct SolverProofHash", []abi.ArgumentMarshaling{
		{Name: "solverTypeHash", Type: "bytes32", InternalType: "bytes32"},
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
		{Name: "data", Type: "bytes32", InternalType: "bytes32"},
	})

	solverOpArgs = abi.Arguments{
		{Type: solverOpSolType, Name: "solverOperation"},
	}

	solverProofHashArgs = abi.Arguments{
		{Type: solverProofHashSolType, Name: "proofHash"},
	}
)

type SolverInput struct {
	UserOpHash   common.Hash      `json:"userOpHash"`
	From         common.Address   `json:"from"`
	To           common.Address   `json:"to"`
	Value        hexutil.Big      `json:"value"`
	Gas          hexutil.Big      `json:"gas"`
	MaxFeePerGas hexutil.Big      `json:"maxFeePerGas"`
	Deadline     hexutil.Big      `json:"deadline"`
	Dapp         common.Address   `json:"dapp"`
	Control      common.Address   `json:"control"`
	Hints        []common.Address `json:"hints,omitempty"`
	Data         hexutil.Bytes    `json:"data,omitempty"`
}

func NewSolverInput(userOp *UserOperation, hints []common.Address) *SolverInput {
	userOpHash, _ := userOp.Hash()
	solverInput := &SolverInput{
		UserOpHash:   userOpHash,
		To:           userOp.To,
		Value:        hexutil.Big(*userOp.Value),
		Gas:          hexutil.Big(*userOp.Gas),
		MaxFeePerGas: hexutil.Big(*userOp.MaxFeePerGas),
		Deadline:     hexutil.Big(*userOp.Deadline),
		Dapp:         userOp.Dapp,
		Control:      userOp.Control,
	}

	if len(hints) > 0 {
		solverInput.Hints = hints
	} else {
		solverInput.Data = hexutil.Bytes(userOp.Data)
	}

	return solverInput
}

// External representation of a solver operation,
// the relay receives and broadcasts solver operations in this format
type SolverOperationRaw struct {
	From         common.Address `json:"from"`
	To           common.Address `json:"to"`
	Value        *hexutil.Big   `json:"value"`
	Gas          *hexutil.Big   `json:"gas"`
	MaxFeePerGas *hexutil.Big   `json:"maxFeePerGas"`
	Deadline     *hexutil.Big   `json:"deadline"`
	Solver       common.Address `json:"solver"`
	Control      common.Address `json:"control"`
	UserOpHash   common.Hash    `json:"userOpHash"`
	BidToken     common.Address `json:"bidToken"`
	BidAmount    *hexutil.Big   `json:"bidAmount"`
	Data         hexutil.Bytes  `json:"data"`
	Signature    hexutil.Bytes  `json:"signature"`
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

func (s *SolverOperation) Validate(solverInput *SolverInput, atlas common.Address, atlasDomainSeparator common.Hash, gasLimit *big.Int) *relayerror.Error {
	if s.To != atlas {
		return ErrSolverOpInvalidToField
	}

	if s.Gas.Cmp(gasLimit) > 0 {
		return ErrSolverOpGasLimitExceeded
	}

	if s.MaxFeePerGas.Cmp(solverInput.MaxFeePerGas.ToInt()) < 0 {
		return ErrSolverOpMaxFeePerGasTooLow
	}

	if s.Deadline.Cmp(solverInput.Deadline.ToInt()) < 0 {
		return ErrSolverOpDeadlineTooLow
	}

	if s.Control != solverInput.Control {
		return ErrSolverOpDAppControlMismatch
	}

	relayErr := s.checkSignature(atlasDomainSeparator)
	if relayErr != nil {
		return relayErr
	}

	return nil
}

func (s *SolverOperation) AbiEncode() ([]byte, error) {
	return solverOpArgs.Pack(&s)
}

func (s *SolverOperation) ProofHash() (common.Hash, error) {
	proofHash := struct {
		SolverTypeHash common.Hash
		From           common.Address
		To             common.Address
		Value          *big.Int
		Gas            *big.Int
		MaxFeePerGas   *big.Int
		Deadline       *big.Int
		Solver         common.Address
		Control        common.Address
		UserOpHash     common.Hash
		BidToken       common.Address
		BidAmount      *big.Int
		Data           common.Hash
	}{
		SOLVER_TYPE_HASH,
		s.From,
		s.To,
		s.Value,
		s.Gas,
		s.MaxFeePerGas,
		s.Deadline,
		s.Solver,
		s.Control,
		s.UserOpHash,
		s.BidToken,
		s.BidAmount,
		crypto.Keccak256Hash(s.Data),
	}

	packed, err := solverProofHashArgs.Pack(&proofHash)
	if err != nil {
		return common.Hash{}, err
	}
	return crypto.Keccak256Hash(packed), nil
}

func (s *SolverOperation) checkSignature(domainSeparator common.Hash) *relayerror.Error {
	if len(s.Signature) != 65 {
		log.Info("invalid user operation signature length", "length", len(s.Signature))
		return ErrUserOpInvalidSignature
	}

	proofHash, err := s.ProofHash()
	if err != nil {
		log.Info("failed to compute solver proof hash", "err", err)
		return ErrSolverOpComputeProofHash.AddError(err)
	}

	signer, err := relayCrypto.RecoverEip712Signer(domainSeparator, proofHash, s.Signature)
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
