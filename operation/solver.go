package operation

import (
	"math/big"

	relayCrypto "github.com/FastLane-Labs/atlas-operations-relay/crypto"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ErrSolverOpInvalidToField      = relayerror.NewError(2201, "solver operation's 'to' field must be atlas contract address")
	ErrSolverOpGasLimitExceeded    = relayerror.NewError(2202, "solver operation's gas limit exceeded")
	ErrSolverOpMaxFeePerGasTooLow  = relayerror.NewError(2203, "solver operation's maxFeePerGas must be equal or higher the user operation")
	ErrSolverOpDeadlineTooLow      = relayerror.NewError(2204, "solver operation's deadline exceeded or lower than user operation's")
	ErrSolverOpDAppControlMismatch = relayerror.NewError(2205, "solver operation's dApp control does not match the user operation's")
	ErrSolverOpComputeProofHash    = relayerror.NewError(2206, "failed to compute solver proof hash")
	ErrSolverOpSignatureInvalid    = relayerror.NewError(2207, "solver operation has invalid signature")
)

var (
	solverGasLimit   = big.NewInt(1000000)
	SOLVER_TYPE_HASH = crypto.Keccak256Hash([]byte("UserOperation(address from,address to,uint256 value,uint256 gas,uint256 maxFeePerGas,uint256 nonce,uint256 deadline,address dapp,address control,address sessionKey,bytes32 data)"))
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

func (s *SolverOperation) Validate(userOp *UserOperation, atlas common.Address, atlasDomainSeparator common.Hash) *relayerror.Error {
	if s.To != atlas {
		return ErrSolverOpInvalidToField
	}

	if s.Gas.Cmp(solverGasLimit) > 0 {
		return ErrSolverOpGasLimitExceeded
	}

	if s.MaxFeePerGas.Cmp(userOp.MaxFeePerGas) < 0 {
		return ErrSolverOpMaxFeePerGasTooLow
	}

	if s.Deadline.Cmp(userOp.Deadline) < 0 {
		return ErrSolverOpDeadlineTooLow
	}

	if s.Control != userOp.Control {
		return ErrSolverOpDAppControlMismatch
	}

	relayErr := s.checkSignature(atlasDomainSeparator)
	if relayErr != nil {
		return relayErr
	}

	return nil
}

func (s *SolverOperation) abiEncode() ([]byte, error) {
	return solverOpArgs.Pack(&s)
}

func (s *SolverOperation) proofHash() (common.Hash, error) {
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
		USER_TYPE_HASH,
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
	proofHash, err := s.proofHash()
	if err != nil {
		log.Info("failed to compute solver proof hash", "err", err)
		return ErrSolverOpComputeProofHash.AddError(err)
	}

	signer, err := relayCrypto.GetSigner(domainSeparator, proofHash, s.Signature)
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
