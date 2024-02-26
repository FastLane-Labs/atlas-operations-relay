package core

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/FastLane-Labs/atlas-operations-relay/bundle"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrMalformedRequest    = relayerror.NewError(1100, "malformed request")
	ErrMalformedJson       = relayerror.NewError(1101, "malformed json")
	ErrInvalidParameter    = relayerror.NewError(1102, "invalid parameter")
	ErrServerCorruptedData = relayerror.NewError(1103, "server corrupted data")
	ErrInvalidUserOpHash   = relayerror.NewError(1104, "invalid user operation hash")
)

type RetrieveRequest struct {
	UserOpHash common.Hash `json:"userOpHash"`
	Wait       bool        `json:"wait"`
}

func NewRetrieveRequest(userOpHash common.Hash, wait bool) *RetrieveRequest {
	return &RetrieveRequest{
		UserOpHash: userOpHash,
		Wait:       wait,
	}
}

type Api struct {
	relay *Relay
}

func NewApi(relay *Relay) *Api {
	return &Api{
		relay: relay,
	}
}

func getRetrieveRequestData(w http.ResponseWriter, r *http.Request) (*RetrieveRequest, *relayerror.Error) {
	q := r.URL.Query()

	userOpHashStr := q.Get("userOpHash")
	if len(userOpHashStr) < 64 || len(userOpHashStr) > 66 {
		return nil, ErrInvalidUserOpHash
	}

	userOpHash := common.HexToHash(q.Get("userOpHash"))

	waitStr := q.Get("wait")
	if len(waitStr) == 0 {
		waitStr = "false"
	}
	wait, err := strconv.ParseBool(waitStr)
	if err != nil {
		return nil, ErrInvalidParameter.AddError(err)
	}

	return NewRetrieveRequest(userOpHash, wait), nil
}

func getPostRequestData(w http.ResponseWriter, r *http.Request, v interface{}) *relayerror.Error {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return ErrMalformedRequest.AddError(err)
	}

	if err := json.Unmarshal(body, v); err != nil {
		return ErrMalformedJson.AddError(err)
	}

	return nil
}

func writeResponseData(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(ErrServerCorruptedData.AddError(err).Marshal())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(b)
}

func (api *Api) SubmitUserOperation(w http.ResponseWriter, r *http.Request) {
	var userpOp *operation.UserOperation
	if relayErr := getPostRequestData(w, r, userpOp); relayErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(relayErr.Marshal())
		return
	}

	userOpHash, relayErr := api.relay.submitUserOperation(userpOp)
	if relayErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(relayErr.Marshal())
		return
	}

	writeResponseData(w, userOpHash)
}

func (api *Api) GetSolverOperations(w http.ResponseWriter, r *http.Request) {
	retrieveReq, relayErr := getRetrieveRequestData(w, r)
	if relayErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(relayErr.Marshal())
		return
	}

	var completionChan chan []*operation.SolverOperation
	if retrieveReq.Wait {
		completionChan = make(chan []*operation.SolverOperation)
	}

	solverOps, relayErr := api.relay.getSolverOperations(retrieveReq.UserOpHash, completionChan)
	if relayErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(relayErr.Marshal())
		return
	}

	if retrieveReq.Wait && solverOps == nil {
		solverOps = <-completionChan
		close(completionChan)
	}

	writeResponseData(w, solverOps)
}

func (api *Api) SubmitBundleOperations(w http.ResponseWriter, r *http.Request) {
	var bundleOps *operation.BundleOperations
	if relayErr := getPostRequestData(w, r, bundleOps); relayErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(relayErr.Marshal())
		return
	}

	result, relayErr := api.relay.submitBundleOperations(bundleOps)
	if relayErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(relayErr.Marshal())
		return
	}

	writeResponseData(w, result)
}

func (api *Api) GetBundleHash(w http.ResponseWriter, r *http.Request) {
	retrieveReq, relayErr := getRetrieveRequestData(w, r)
	if relayErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(relayErr.Marshal())
		return
	}

	var completionChan chan *bundle.Bundle
	if retrieveReq.Wait {
		completionChan = make(chan *bundle.Bundle)
	}

	atlasTxHash, relayErr := api.relay.getBundleHash(retrieveReq.UserOpHash, completionChan)
	if relayErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(relayErr.Marshal())
		return
	}

	if retrieveReq.Wait && atlasTxHash == (common.Hash{}) {
		bundle := <-completionChan
		atlasTxHash, relayErr = bundle.GetResult()
		if relayErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(relayErr.Marshal())
			return
		}
		close(completionChan)
	}

	writeResponseData(w, atlasTxHash)
}

func (api *Api) SubmitSolverOperation(w http.ResponseWriter, r *http.Request) {
	var solverOp *operation.SolverOperation
	if relayErr := getPostRequestData(w, r, solverOp); relayErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(relayErr.Marshal())
		return
	}

	result, relayErr := api.relay.submitSolverOperation(solverOp)
	if relayErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(relayErr.Marshal())
		return
	}

	writeResponseData(w, result)
}

func (api *Api) WebsocketSolver(w http.ResponseWriter, r *http.Request) {
	api.relay.server.ServeWsSolver(w, r)
}

func (api *Api) WebsocketBundler(w http.ResponseWriter, r *http.Request) {
	// TODO: bundler authentication
	var bundler common.Address

	api.relay.server.ServeWsBundler(w, r, bundler)
}
