package operation

type BundleOperations struct {
	UserOperation    *UserOperation     `json:"userOperation"`
	SolverOperations []*SolverOperation `json:"solverOperations"`
	DAppOperation    *DAppOperation     `json:"dAppOperation"`
}
