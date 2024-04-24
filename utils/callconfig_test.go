package utils

import (
	"testing"
)

func TestFlagUserNoncesSequenced(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "UserNoncesSequenced",
			callConfig: 1 << indexUserNoncesSequenced,
			want:       true,
		},
		{
			name:       "UserNoncesSequencedNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "UserNoncesSequencedAndDAppNoncesSequenced",
			callConfig: 1<<indexUserNoncesSequenced | 1<<indexDAppNoncesSequenced,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagUserNoncesSequenced(tt.callConfig); got != tt.want {
				t.Errorf("FlagUserNoncesSequenced() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagDAppNoncesSequenced(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "DAppNoncesSequenced",
			callConfig: 1 << indexDAppNoncesSequenced,
			want:       true,
		},
		{
			name:       "DAppNoncesSequencedNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "DAppNoncesSequencedAndUserNoncesSequenced",
			callConfig: 1<<indexDAppNoncesSequenced | 1<<indexUserNoncesSequenced,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagDAppNoncesSequenced(tt.callConfig); got != tt.want {
				t.Errorf("FlagDAppNoncesSequenced() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagRequirePreOps(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "RequirePreOps",
			callConfig: 1 << indexRequirePreOps,
			want:       true,
		},
		{
			name:       "RequirePreOpsNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "RequirePreOpsAndTrackPreOpsReturnData",
			callConfig: 1<<indexRequirePreOps | 1<<indexTrackPreOpsReturnData,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagRequirePreOps(tt.callConfig); got != tt.want {
				t.Errorf("FlagRequirePreOps() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagTrackPreOpsReturnData(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "TrackPreOpsReturnData",
			callConfig: 1 << indexTrackPreOpsReturnData,
			want:       true,
		},
		{
			name:       "TrackPreOpsReturnDataNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "TrackPreOpsReturnDataAndTrackUserReturnData",
			callConfig: 1<<indexTrackPreOpsReturnData | 1<<indexTrackUserReturnData,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagTrackPreOpsReturnData(tt.callConfig); got != tt.want {
				t.Errorf("FlagTrackPreOpsReturnData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagTrackUserReturnData(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "TrackUserReturnData",
			callConfig: 1 << indexTrackUserReturnData,
			want:       true,
		},
		{
			name:       "TrackUserReturnDataNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "TrackUserReturnDataAndDelegateUser",
			callConfig: 1<<indexTrackUserReturnData | 1<<indexDelegateUser,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagTrackUserReturnData(tt.callConfig); got != tt.want {
				t.Errorf("FlagTrackUserReturnData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagDelegateUser(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "DelegateUser",
			callConfig: 1 << indexDelegateUser,
			want:       true,
		},
		{
			name:       "DelegateUserNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "DelegateUserAndPreSolver",
			callConfig: 1<<indexDelegateUser | 1<<indexPreSolver,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagDelegateUser(tt.callConfig); got != tt.want {
				t.Errorf("FlagDelegateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagPreSolver(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "PreSolver",
			callConfig: 1 << indexPreSolver,
			want:       true,
		},
		{
			name:       "PreSolverNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "PreSolverAndPostSolver",
			callConfig: 1<<indexPreSolver | 1<<indexPostSolver,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagPreSolver(tt.callConfig); got != tt.want {
				t.Errorf("FlagPreSolver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagPostSolver(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "PostSolver",
			callConfig: 1 << indexPostSolver,
			want:       true,
		},
		{
			name:       "PostSolverNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "PostSolverAndPostOpsCall",
			callConfig: 1<<indexPostSolver | 1<<indexPostOpsCall,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagPostSolver(tt.callConfig); got != tt.want {
				t.Errorf("FlagPostSolver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagPostOpsCall(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "PostOpsCall",
			callConfig: 1 << indexPostOpsCall,
			want:       true,
		},
		{
			name:       "PostOpsCallNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "PostOpsCallAndZeroSolvers",
			callConfig: 1<<indexPostOpsCall | 1<<indexZeroSolvers,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagPostOpsCall(tt.callConfig); got != tt.want {
				t.Errorf("FlagPostOpsCall() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagZeroSolvers(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "ZeroSolvers",
			callConfig: 1 << indexZeroSolvers,
			want:       true,
		},
		{
			name:       "ZeroSolversNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "ZeroSolversAndReuseUserOp",
			callConfig: 1<<indexZeroSolvers | 1<<indexReuseUserOp,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagZeroSolvers(tt.callConfig); got != tt.want {
				t.Errorf("FlagZeroSolvers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagReuseUserOp(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "ReuseUserOp",
			callConfig: 1 << indexReuseUserOp,
			want:       true,
		},
		{
			name:       "ReuseUserOpNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "ReuseUserOpAndUserAuctioneer",
			callConfig: 1<<indexReuseUserOp | 1<<indexUserAuctioneer,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagReuseUserOp(tt.callConfig); got != tt.want {
				t.Errorf("FlagReuseUserOp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagUserAuctioneer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "UserAuctioneer",
			callConfig: 1 << indexUserAuctioneer,
			want:       true,
		},
		{
			name:       "UserAuctioneerNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "UserAuctioneerAndSolverAuctioneer",
			callConfig: 1<<indexUserAuctioneer | 1<<indexSolverAuctioneer,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagUserAuctioneer(tt.callConfig); got != tt.want {
				t.Errorf("FlagUserAuctioneer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagSolverAuctioneer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "SolverAuctioneer",
			callConfig: 1 << indexSolverAuctioneer,
			want:       true,
		},
		{
			name:       "SolverAuctioneerNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "SolverAuctioneerAndUnknownAuctioneer",
			callConfig: 1<<indexSolverAuctioneer | 1<<indexUnknownAuctioneer,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagSolverAuctioneer(tt.callConfig); got != tt.want {
				t.Errorf("FlagSolverAuctioneer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagUnknownAuctioneer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "UnknownAuctioneer",
			callConfig: 1 << indexUnknownAuctioneer,
			want:       true,
		},
		{
			name:       "UnknownAuctioneerNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "UnknownAuctioneerAndUserNoncesSequenced",
			callConfig: 1<<indexUnknownAuctioneer | 1<<indexUserNoncesSequenced,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagUnknownAuctioneer(tt.callConfig); got != tt.want {
				t.Errorf("FlagUnknownAuctioneer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagVerifyCallChainHash(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "VerifyCallChainHash",
			callConfig: 1 << indexVerifyCallChainHash,
			want:       true,
		},
		{
			name:       "VerifyCallChainHashNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "VerifyCallChainHashAndIndexInvertBidValue",
			callConfig: 1<<indexVerifyCallChainHash | 1<<indexInvertBidValue,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagVerifyCallChainHash(tt.callConfig); got != tt.want {
				t.Errorf("FlagVerifyCallChainHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagInvertBidValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "InvertBidValue",
			callConfig: 1 << indexInvertBidValue,
			want:       true,
		},
		{
			name:       "InvertBidValueNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "InvertBidValueAndUserNoncesSequenced",
			callConfig: 1<<indexInvertBidValue | 1<<indexUserNoncesSequenced,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagInvertBidValue(tt.callConfig); got != tt.want {
				t.Errorf("FlagInvertBidValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagNeedsFulfillment(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "NeedsFulfillment",
			callConfig: 1 << indexNeedsFulfillment,
			want:       true,
		},
		{
			name:       "NeedsFulfillmentNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "NeedsFulfillmentAndTrustedOpHash",
			callConfig: 1<<indexNeedsFulfillment | 1<<indexTrustedOpHash,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagNeedsFulfillment(tt.callConfig); got != tt.want {
				t.Errorf("FlagNeedsFulfillment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagTrustedOpHash(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "TrustedOpHash",
			callConfig: 1 << indexTrustedOpHash,
			want:       true,
		},
		{
			name:       "TrustedOpHashNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "TrustedOpHashAndFlagForwardReturnData",
			callConfig: 1<<indexTrustedOpHash | 1<<indexForwardReturnData,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagTrustedOpHash(tt.callConfig); got != tt.want {
				t.Errorf("FlagTrustedOpHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagForwardReturnData(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "ForwardReturnData",
			callConfig: 1 << indexForwardReturnData,
			want:       true,
		},
		{
			name:       "ForwardReturnDataNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "ForwardReturnDataAndFlagNeedsFulfillment",
			callConfig: 1<<indexForwardReturnData | 1<<indexNeedsFulfillment,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagForwardReturnData(tt.callConfig); got != tt.want {
				t.Errorf("FlagForwardReturnData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagExPostBids(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		callConfig uint32
		want       bool
	}{
		{
			name:       "ExPostBids",
			callConfig: 1 << indexExPostBids,
			want:       true,
		},
		{
			name:       "ExPostBidsNotSet",
			callConfig: 0,
			want:       false,
		},
		{
			name:       "ExPostBidsAndFlagTrustedOpHash",
			callConfig: 1<<indexExPostBids | 1<<indexTrustedOpHash,
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlagExPostBids(tt.callConfig); got != tt.want {
				t.Errorf("FlagExPostBids() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlagVerifyCallChainHashAndInvertBidValue(t *testing.T) {
	t.Parallel()

	callConfig := uint32(1<<indexVerifyCallChainHash | 1<<indexInvertBidValue)
	want := true

	if got := FlagVerifyCallChainHash(callConfig); got != want {
		t.Errorf("FlagVerifyCallChainHash() = %v, want %v", got, want)
	}

	want = true

	if got := FlagInvertBidValue(callConfig); got != want {
		t.Errorf("FlagInvertBidValue() = %v, want %v", got, want)
	}
}

func TestFlagTrackUserReturnDataAndDelegateUser(t *testing.T) {
	t.Parallel()

	callConfig := uint32(1<<indexTrackUserReturnData | 1<<indexDelegateUser)
	want := true

	if got := FlagTrackUserReturnData(callConfig); got != want {
		t.Errorf("FlagTrackUserReturnData() = %v, want %v", got, want)
	}

	want = true

	if got := FlagDelegateUser(callConfig); got != want {
		t.Errorf("FlagDelegateUser() = %v, want %v", got, want)
	}
}

func TestFlagPreSolverAndPostSolver(t *testing.T) {
	t.Parallel()

	callConfig := uint32(1<<indexPreSolver | 1<<indexPostSolver)
	want := true

	if got := FlagPreSolver(callConfig); got != want {
		t.Errorf("FlagPreSolver() = %v, want %v", got, want)
	}

	want = true

	if got := FlagPostSolver(callConfig); got != want {
		t.Errorf("FlagPostSolver() = %v, want %v", got, want)
	}
}
