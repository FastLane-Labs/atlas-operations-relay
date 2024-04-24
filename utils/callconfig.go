package utils

const (
	indexUserNoncesSequenced = uint32(iota)
	indexDAppNoncesSequenced
	indexRequirePreOps
	indexTrackPreOpsReturnData
	indexTrackUserReturnData
	indexDelegateUser
	indexPreSolver
	indexPostSolver
	indexPostOpsCall
	indexZeroSolvers
	indexReuseUserOp
	indexUserAuctioneer
	indexSolverAuctioneer
	indexUnknownAuctioneer
	indexVerifyCallChainHash
	indexForwardReturnData
	indexNeedsFulfillment
	indexTrustedOpHash
	indexInvertBidValue
	indexExPostBids
)

func FlagUserNoncesSequenced(callConfig uint32) bool {
	return callConfig&(1<<indexUserNoncesSequenced) != 0
}

func FlagDAppNoncesSequenced(callConfig uint32) bool {
	return callConfig&(1<<indexDAppNoncesSequenced) != 0
}

func FlagRequirePreOps(callConfig uint32) bool {
	return callConfig&(1<<indexRequirePreOps) != 0
}

func FlagTrackPreOpsReturnData(callConfig uint32) bool {
	return callConfig&(1<<indexTrackPreOpsReturnData) != 0
}

func FlagTrackUserReturnData(callConfig uint32) bool {
	return callConfig&(1<<indexTrackUserReturnData) != 0
}

func FlagDelegateUser(callConfig uint32) bool {
	return callConfig&(1<<indexDelegateUser) != 0
}

func FlagPreSolver(callConfig uint32) bool {
	return callConfig&(1<<indexPreSolver) != 0
}

func FlagPostSolver(callConfig uint32) bool {
	return callConfig&(1<<indexPostSolver) != 0
}

func FlagPostOpsCall(callConfig uint32) bool {
	return callConfig&(1<<indexPostOpsCall) != 0
}

func FlagZeroSolvers(callConfig uint32) bool {
	return callConfig&(1<<indexZeroSolvers) != 0
}

func FlagReuseUserOp(callConfig uint32) bool {
	return callConfig&(1<<indexReuseUserOp) != 0
}

func FlagUserAuctioneer(callConfig uint32) bool {
	return callConfig&(1<<indexUserAuctioneer) != 0
}

func FlagSolverAuctioneer(callConfig uint32) bool {
	return callConfig&(1<<indexSolverAuctioneer) != 0
}

func FlagUnknownAuctioneer(callConfig uint32) bool {
	return callConfig&(1<<indexUnknownAuctioneer) != 0
}

func FlagVerifyCallChainHash(callConfig uint32) bool {
	return callConfig&(1<<indexVerifyCallChainHash) != 0
}

func FlagForwardReturnData(callConfig uint32) bool {
	return callConfig&(1<<indexForwardReturnData) != 0
}

func FlagNeedsFulfillment(callConfig uint32) bool {
	return callConfig&(1<<indexNeedsFulfillment) != 0
}

func FlagTrustedOpHash(callConfig uint32) bool {
	return callConfig&(1<<indexTrustedOpHash) != 0
}

func FlagInvertBidValue(callConfig uint32) bool {
	return callConfig&(1<<indexInvertBidValue) != 0
}

func FlagExPostBids(callConfig uint32) bool {
	return callConfig&(1<<indexExPostBids) != 0
}
