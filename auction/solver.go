package auction

const (
	SolverStatusCodeFailure = iota
	SolverStatusCodeSuccess
	SolverStatusCodePending
)

var (
	SolverStatusIncluded         = NewSolverStatus(SolverStatusCodeSuccess, "included in auction")
	SolverStatusAuctionPending   = NewSolverStatus(SolverStatusCodePending, "auction is pending")
	SolverStatusFailedSimulation = NewSolverStatus(SolverStatusCodeFailure, "failed simulation")
	SolverStatusNotIncluded      = NewSolverStatus(SolverStatusCodeFailure, "not included in auction")
)

type SolverStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewSolverStatus(code int, message string) *SolverStatus {
	return &SolverStatus{
		Code:    code,
		Message: message,
	}
}
