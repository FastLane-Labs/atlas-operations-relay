package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/bundle"
	relayCrypto "github.com/FastLane-Labs/atlas-operations-relay/crypto"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	ErrMalformedRequest      = relayerror.NewError(3000, "malformed request")
	ErrMalformedJson         = relayerror.NewError(3001, "malformed json")
	ErrInvalidParameter      = relayerror.NewError(3002, "invalid parameter")
	ErrServerCorruptedData   = relayerror.NewError(3003, "server corrupted data")
	ErrInvalidUserOpHash     = relayerror.NewError(3004, "invalid user operation hash")
	ErrInvalidBundlerAddress = relayerror.NewError(3005, "invalid bundler address")
	ErrInvalidTimestamp      = relayerror.NewError(3006, "invalid timestamp")
	ErrExpiredSignature      = relayerror.NewError(3007, "expired signature")
	ErrBadSignature          = relayerror.NewError(3008, "bad signature (decode/recover error)")
	ErrSignatureMismatch     = relayerror.NewError(3009, "signature mismatch")
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

func getRetrieveRequestData(r *http.Request) (*RetrieveRequest, *relayerror.Error) {
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

func getPostRequestData(r *http.Request, v interface{}) *relayerror.Error {
	body, err := io.ReadAll(r.Body)
	log.Info("post request body", r.Body)
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
	userOp := &UserOperationArgs{}
	if relayErr := getPostRequestData(r, userOp); relayErr != nil {
		log.Info("got to submituseroperation", "getpostdataerr", relayErr)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(relayErr.Marshal())
		return
	}
	userOpHash, relayErr := api.relay.submitUserOperation(userOp)
	if relayErr != nil {
		log.Info("got to submituseroperation", "submitoperr", relayErr)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(relayErr.Marshal())
		return
	}

	writeResponseData(w, userOpHash)
}

func (api *Api) GetSolverOperations(w http.ResponseWriter, r *http.Request) {
	retrieveReq, relayErr := getRetrieveRequestData(r)
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
	bundleOps := &operation.BundleOperations{}
	if relayErr := getPostRequestData(r, bundleOps); relayErr != nil {
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
	retrieveReq, relayErr := getRetrieveRequestData(r)
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
	solverOp := &operation.SolverOperation{}
	if relayErr := getPostRequestData(r, solverOp); relayErr != nil {
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
	q := r.URL.Query()

	address := q.Get("address")
	if !common.IsHexAddress(address) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrInvalidBundlerAddress.Marshal())
		return
	}
	bundler := common.HexToAddress(address)

	timestamp, err := strconv.ParseInt(q.Get("timestamp"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrInvalidTimestamp.Marshal())
		return
	}

	// 60 seconds window past and future for timestamp
	minTimestamp, maxTimestamp := time.Now().Unix()-60, time.Now().Unix()+60
	if timestamp < minTimestamp || timestamp > maxTimestamp {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrExpiredSignature.Marshal())
		return
	}

	signature, err := hexutil.Decode(q.Get("signature"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrBadSignature.Marshal())
		return
	}

	// Validate the signature
	signatureContent := fmt.Sprintf("%s:%d", bundler, timestamp)
	signer, err := relayCrypto.RecoverEthereumSigner(signatureContent, signature)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrBadSignature.Marshal())
		return
	}

	if bundler != signer {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(ErrSignatureMismatch.Marshal())
		return
	}

	api.relay.server.ServeWsBundler(w, r, bundler)
}
