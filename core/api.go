package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/auction"
	"github.com/FastLane-Labs/atlas-operations-relay/bundle"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/FastLane-Labs/atlas-operations-relay/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-playground/validator/v10"
)

var (
	ErrMalformedRequest    = relayerror.NewError(3000, "malformed request")
	ErrMalformedJson       = relayerror.NewError(3001, "malformed json")
	ErrInvalidParameter    = relayerror.NewError(3002, "invalid parameter")
	ErrServerCorruptedData = relayerror.NewError(3003, "server corrupted data")
	ErrInvalidOpHash       = relayerror.NewError(3004, "invalid operation hash")
	ErrInvalidAddress      = relayerror.NewError(3005, "invalid address")
	ErrInvalidTimestamp    = relayerror.NewError(3006, "invalid timestamp")
	ErrExpiredSignature    = relayerror.NewError(3007, "expired signature")
	ErrBadSignature        = relayerror.NewError(3008, "bad signature (decode/recover error)")
	ErrSignatureMismatch   = relayerror.NewError(3009, "signature mismatch")
)

type RetrieveRequest struct {
	OpHash common.Hash
	Wait   bool
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

	hashStr := q.Get("operationHash")
	if len(hashStr) < 64 || len(hashStr) > 66 {
		return nil, ErrInvalidOpHash
	}

	opHash := common.HexToHash(hashStr)

	waitStr := q.Get("wait")
	if len(waitStr) == 0 {
		waitStr = "false"
	}
	wait, err := strconv.ParseBool(waitStr)
	if err != nil {
		return nil, ErrInvalidParameter.AddError(err)
	}

	return &RetrieveRequest{OpHash: opHash, Wait: wait}, nil
}

func getPostRequestData(r *http.Request, v interface{}) *relayerror.Error {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return ErrMalformedRequest.AddError(err)
	}

	if err := json.Unmarshal(body, v); err != nil {
		return ErrMalformedJson.AddError(err)
	}

	validate := validator.New()
	if err := validate.Struct(v); err != nil {
		return ErrInvalidParameter.AddError(err)
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

func (api *Api) Ping(w http.ResponseWriter, r *http.Request) {
	writeResponseData(w, "pong")
}

func (api *Api) SubmitUserOperation(w http.ResponseWriter, r *http.Request) {
	userOpWithHintsRaw := &operation.UserOperationWithHintsRaw{}
	if relayErr := getPostRequestData(r, userOpWithHintsRaw); relayErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(relayErr.Marshal())
		return
	}

	userOpHash, relayErr := api.relay.submitUserOperation(userOpWithHintsRaw.Decode())
	if relayErr != nil {
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

	var completionChan chan []*operation.SolverOperationWithScore
	if retrieveReq.Wait {
		completionChan = make(chan []*operation.SolverOperationWithScore)
	}

	solverOpsWithScore, relayErr := api.relay.getSolverOperations(retrieveReq.OpHash, completionChan)
	if relayErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(relayErr.Marshal())
		return
	}

	if retrieveReq.Wait && solverOpsWithScore == nil {
		solverOpsWithScore = <-completionChan
		close(completionChan)
	}

	writeResponseData(w, solverOpsWithScore.EncodeToRaw())
}

func (api *Api) SubmitBundleOperations(w http.ResponseWriter, r *http.Request) {
	bundleOpsRaw := &operation.BundleOperationsRaw{}
	if relayErr := getPostRequestData(r, bundleOpsRaw); relayErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(relayErr.Marshal())
		return
	}

	result, relayErr := api.relay.submitBundleOperations(bundleOpsRaw.Decode())
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

	atlasTxHash, relayErr := api.relay.getBundleHash(retrieveReq.OpHash, completionChan)
	if relayErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(relayErr.Marshal())
		return
	}

	if retrieveReq.Wait && atlasTxHash == nil {
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
	solverOpRaw := &operation.SolverOperationRaw{}
	if relayErr := getPostRequestData(r, solverOpRaw); relayErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(relayErr.Marshal())
		return
	}

	solverOpHash, relayErr := api.relay.submitSolverOperation(solverOpRaw.Decode())
	if relayErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(relayErr.Marshal())
		return
	}

	writeResponseData(w, solverOpHash)
}

func (api *Api) GetSolverOperationStatus(w http.ResponseWriter, r *http.Request) {
	retrieveReq, relayErr := getRetrieveRequestData(r)
	if relayErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(relayErr.Marshal())
		return
	}

	var completionChan chan *auction.SolverStatus
	if retrieveReq.Wait {
		completionChan = make(chan *auction.SolverStatus)
	}

	solverOpsStatus, relayErr := api.relay.getSolverOperationStatus(retrieveReq.OpHash, completionChan)
	if relayErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(relayErr.Marshal())
		return
	}

	if retrieveReq.Wait && solverOpsStatus == nil {
		solverOpsStatus = <-completionChan
		close(completionChan)
	}

	writeResponseData(w, solverOpsStatus)
}

func (api *Api) WebsocketSolver(w http.ResponseWriter, r *http.Request) {
	api.relay.server.ServeWsSolver(w, r)
}

func (api *Api) WebsocketBundler(w http.ResponseWriter, r *http.Request) {
	bundler, relayErr := authenticate(r)
	if relayErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(relayErr.Marshal())
		return
	}

	api.relay.server.ServeWsBundler(w, r, bundler)
}

func (api *Api) WebsocketSignatory(w http.ResponseWriter, r *http.Request) {
	signatory, relayErr := authenticate(r)
	if relayErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(relayErr.Marshal())
		return
	}

	api.relay.server.ServeWsSignatory(w, r, signatory)
}

func authenticate(r *http.Request) (common.Address, *relayerror.Error) {
	q := r.URL.Query()

	address := q.Get("address")
	if !common.IsHexAddress(address) {
		return common.Address{}, ErrInvalidAddress
	}
	account := common.HexToAddress(address)

	timestamp, err := strconv.ParseInt(q.Get("timestamp"), 10, 64)
	if err != nil {
		return common.Address{}, ErrInvalidTimestamp
	}

	// 60 seconds window past and future for timestamp
	minTimestamp, maxTimestamp := time.Now().Unix()-60, time.Now().Unix()+60
	if timestamp < minTimestamp || timestamp > maxTimestamp {
		return common.Address{}, ErrExpiredSignature
	}

	signature, err := hexutil.Decode(q.Get("signature"))
	if err != nil {
		return common.Address{}, ErrBadSignature
	}

	// Validate the signature
	signatureContent := fmt.Sprintf("%s:%d", account, timestamp)
	signer, err := utils.RecoverEthereumSigner(signatureContent, signature)
	if err != nil {
		return common.Address{}, ErrBadSignature
	}

	if account != signer {
		return common.Address{}, ErrSignatureMismatch
	}

	return account, nil
}
