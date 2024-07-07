package utils

const (
	IndexUserNoncesSequenced = uint32(iota)
	IndexDAppNoncesSequenced
	IndexRequirePreOps
	IndexTrackPreOpsReturnData
	IndexTrackUserReturnData
	IndexDelegateUser
	IndexPreSolver
	IndexPostSolver
	IndexPostOpsCall
	IndexZeroSolvers
	IndexReuseUserOp
	IndexUserAuctioneer
	IndexSolverAuctioneer
	IndexUnknownAuctioneer
	IndexVerifyCallChainHash
	IndexForwardReturnData
	IndexNeedsFulfillment
	IndexTrustedOpHash
	IndexInvertBidValue
	IndexExPostBids
	IndexAllowAllocateValueFailure
)

func FlagUserNoncesSequenced(callConfig uint32) bool {
	return callConfig&(1<<IndexUserNoncesSequenced) != 0
}

func FlagDAppNoncesSequenced(callConfig uint32) bool {
	return callConfig&(1<<IndexDAppNoncesSequenced) != 0
}

func FlagRequirePreOps(callConfig uint32) bool {
	return callConfig&(1<<IndexRequirePreOps) != 0
}

func FlagTrackPreOpsReturnData(callConfig uint32) bool {
	return callConfig&(1<<IndexTrackPreOpsReturnData) != 0
}

func FlagTrackUserReturnData(callConfig uint32) bool {
	return callConfig&(1<<IndexTrackUserReturnData) != 0
}

func FlagDelegateUser(callConfig uint32) bool {
	return callConfig&(1<<IndexDelegateUser) != 0
}

func FlagPreSolver(callConfig uint32) bool {
	return callConfig&(1<<IndexPreSolver) != 0
}

func FlagPostSolver(callConfig uint32) bool {
	return callConfig&(1<<IndexPostSolver) != 0
}

func FlagPostOpsCall(callConfig uint32) bool {
	return callConfig&(1<<IndexPostOpsCall) != 0
}

func FlagZeroSolvers(callConfig uint32) bool {
	return callConfig&(1<<IndexZeroSolvers) != 0
}

func FlagReuseUserOp(callConfig uint32) bool {
	return callConfig&(1<<IndexReuseUserOp) != 0
}

func FlagUserAuctioneer(callConfig uint32) bool {
	return callConfig&(1<<IndexUserAuctioneer) != 0
}

func FlagSolverAuctioneer(callConfig uint32) bool {
	return callConfig&(1<<IndexSolverAuctioneer) != 0
}

func FlagUnknownAuctioneer(callConfig uint32) bool {
	return callConfig&(1<<IndexUnknownAuctioneer) != 0
}

func FlagVerifyCallChainHash(callConfig uint32) bool {
	return callConfig&(1<<IndexVerifyCallChainHash) != 0
}

func FlagForwardReturnData(callConfig uint32) bool {
	return callConfig&(1<<IndexForwardReturnData) != 0
}

func FlagNeedsFulfillment(callConfig uint32) bool {
	return callConfig&(1<<IndexNeedsFulfillment) != 0
}

func FlagTrustedOpHash(callConfig uint32) bool {
	return callConfig&(1<<IndexTrustedOpHash) != 0
}

func FlagInvertBidValue(callConfig uint32) bool {
	return callConfig&(1<<IndexInvertBidValue) != 0
}

func FlagExPostBids(callConfig uint32) bool {
	return callConfig&(1<<IndexExPostBids) != 0
}

func FlagAllowAllocateValueFailure(callConfig uint32) bool {
	return callConfig&(1<<IndexAllowAllocateValueFailure) != 0
}
