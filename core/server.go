package core

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/auction"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"golang.org/x/time/rate"
)

const (
	BundlerTimeout      = 2 * time.Second
	SignatoryTimeout    = 2 * time.Second
	websocketRateLimit  = 2
	websocketBurstLimit = 32

	// Channels
	ChannelSolver    = "solver"
	ChannelBundler   = "bundler"
	ChannelSignatory = "signatory"

	// Methods
	MethodPing                  = "ping"
	MethodSubscribe             = "subscribe"
	MethodUnsubscribe           = "unsubscribe"
	MethodSubmitSolverOperation = "submitSolverOperation"
	MethodSolverOperationStatus = "solverOperationStatus"

	// Subscriptions topics
	TopicNewUserOperations = "newUserOperations"

	// Events
	EventUpdate              = "update"
	EventNewBundle           = "newBundle"
	EventNewSignatoryRequest = "newSignatoryRequest"

	// Messages
	ConnUpgradeFailed      = "failed upgrading connection"
	Pong                   = "pong"
	Subscribed             = "subscribed"
	Unsubscribed           = "unsubscribed"
	AlreadySubscribed      = "already subscribed"
	NotSubscribed          = "not subscribed"
	InvalidMessage         = "invalid message"
	InvalidChannel         = "invalid channel"
	InvalidMethod          = "invalid method"
	InvalidTopic           = "invalid topic"
	InvalidSolverOperation = "invalid solver operation"
	RateLimitExceeded      = "rate limit exceeded"

	// Websocket config
	pongWait        = 60 * time.Second
	pingPeriod      = (pongWait * 9) / 10
	writeWait       = 10 * time.Second
	maxMessageSize  = 2048
	readBufferSize  = 1024
	writeBufferSize = 1024
)

var (
	Channels = map[string]struct{}{
		ChannelSolver:    {},
		ChannelBundler:   {},
		ChannelSignatory: {},
	}

	Methods = map[string]map[string]struct{}{
		ChannelSolver:    {MethodPing: {}, MethodSubscribe: {}, MethodUnsubscribe: {}, MethodSubmitSolverOperation: {}, MethodSolverOperationStatus: {}},
		ChannelBundler:   {MethodPing: {}},
		ChannelSignatory: {MethodPing: {}},
	}

	Topics = map[string]struct{}{
		TopicNewUserOperations: {},
	}

	upgrader = websocket.Upgrader{
		ReadBufferSize:  readBufferSize,
		WriteBufferSize: writeBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

var (
	ErrBundlerOffline    = relayerror.NewError(3200, "bundler is offline")
	ErrCloggedConnection = relayerror.NewError(3201, "clogged connection")
	ErrBundlingFailure   = relayerror.NewError(3202, "bundling failure")
)

type newSolverOperationFn func(*operation.SolverOperation) (common.Hash, *relayerror.Error)
type getSolverOperationStatusFn func(common.Hash, chan *auction.SolverStatus) (*auction.SolverStatus, *relayerror.Error)
type getDAppSignatoriesFn func(common.Address) ([]common.Address, *relayerror.Error)
type setAtlasTxHashFn func(common.Hash) *relayerror.Error
type setRelayErrorFn func(*relayerror.Error) *relayerror.Error
type submitBundleOperationsFn func(*operation.BundleOperations) (string, *relayerror.Error)

type RequestParams struct {
	Topic           string                         `json:"topic"`
	SolverOperation *operation.SolverOperationRaw  `json:"solverOperation"`
	OpHash          common.Hash                    `json:"operationHash"`
	Bundle          *operation.BundleOperationsRaw `json:"bundle"`
}

type BareRequest struct {
	Id string `json:"id" validate:"required"`
}

type Request struct {
	Id     string         `json:"id" validate:"required"`
	Method string         `json:"method" validate:"required"`
	Params *RequestParams `json:"params" validate:"required"`
}

type Response struct {
	Id     string      `json:"id"`
	Result interface{} `json:"result,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func (r *Response) Marshal() []byte {
	b, _ := json.Marshal(r)
	return b
}

type BroadcastParams struct {
	UserOperationPartial *operation.UserOperationPartialRaw `json:"userOperation,omitempty"`
}

type Broadcast struct {
	Event string           `json:"event"`
	Topic string           `json:"topic"`
	Data  *BroadcastParams `json:"data"`
}

func (bc *Broadcast) Marshal() []byte {
	b, _ := json.Marshal(bc)
	return b
}

type BundleRequest struct {
	Id     string                         `json:"id"`
	Event  string                         `json:"event"`
	Bundle *operation.BundleOperationsRaw `json:"bundle"`
}

func (br *BundleRequest) Marshal() []byte {
	b, _ := json.Marshal(br)
	return b
}

type BundleResponse struct {
	Id     string      `json:"id" validate:"required"`
	Result common.Hash `json:"result,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func (br *BundleResponse) Marshal() []byte {
	b, _ := json.Marshal(br)
	return b
}

type SignatoryRequest struct {
	Id               string                          `json:"id"`
	Event            string                          `json:"event"`
	UserOperation    *operation.UserOperationRaw     `json:"userOperation"`
	SolverOperations []*operation.SolverOperationRaw `json:"solverOperations"`
}

func (sr *SignatoryRequest) Marshal() []byte {
	b, _ := json.Marshal(sr)
	return b
}

type SignatoryResponse struct {
	Id     string                      `json:"id"`
	Result *operation.DAppOperationRaw `json:"result"`
	Error  string                      `json:"error,omitempty"`
}

func (sr *SignatoryResponse) Marshal() []byte {
	b, _ := json.Marshal(sr)
	return b
}

type Marshaler interface {
	Marshal() []byte
}

type Conn struct {
	*websocket.Conn

	uuid      string
	channel   string
	bundler   common.Address
	signatory common.Address
	sendChan  chan []byte
}

func NewConn(conn *websocket.Conn, channel string, bundler common.Address, signatory common.Address) *Conn {
	return &Conn{
		Conn:      conn,
		uuid:      uuid.New().String(),
		channel:   channel,
		bundler:   bundler,
		signatory: signatory,
		sendChan:  make(chan []byte, 256),
	}
}

func (c *Conn) send(message Marshaler) *relayerror.Error {
	return c.sendBytes(message.Marshal())
}

func (c *Conn) sendString(message string) *relayerror.Error {
	return c.sendBytes([]byte(message))
}

func (c *Conn) sendBytes(message []byte) *relayerror.Error {
	select {
	case c.sendChan <- message:
	default:
		// Channel is full, drop message
		return ErrCloggedConnection
	}

	return nil
}

func (c *Conn) isBundler() bool {
	return c.bundler != (common.Address{})
}

func (c *Conn) isSignatory() bool {
	return c.signatory != (common.Address{})
}

type bundlingRequest struct {
	candidatesBundlers map[common.Address]*Conn
	offlineBundlers    map[common.Address]bool
	setAtlasTxHash     setAtlasTxHashFn
	setRelayError      setRelayErrorFn
	doneChan           chan struct{}
}

type SigningRequest struct {
	candidatesSignatories  map[common.Address]*Conn
	offlineSignatories     map[common.Address]bool
	submitBundleOperations submitBundleOperationsFn
	bundle                 *operation.BundleOperations
	doneChan               chan struct{}
}

type Server struct {
	router *mux.Router

	// Indexed by [topic][uuid]
	subscriptions map[string]map[string]*Conn

	// Indexed by [address]
	bundlers    map[common.Address]*Conn
	signatories map[common.Address]*Conn

	// Indexed by [id]
	bundlingRequests map[string]*bundlingRequest
	signingRequests  map[string]*SigningRequest

	newSolverOperation       newSolverOperationFn
	getSolverOperationStatus getSolverOperationStatusFn
	getDAppSignatories       getDAppSignatoriesFn

	mu sync.RWMutex
}

func NewServer(router *mux.Router, newSolverOperation newSolverOperationFn, getSolverOperationStatus getSolverOperationStatusFn, getDAppSignatories getDAppSignatoriesFn) *Server {
	return &Server{
		router:                   router,
		subscriptions:            make(map[string]map[string]*Conn),
		bundlers:                 make(map[common.Address]*Conn),
		signatories:              make(map[common.Address]*Conn),
		bundlingRequests:         make(map[string]*bundlingRequest),
		signingRequests:          make(map[string]*SigningRequest),
		newSolverOperation:       newSolverOperation,
		getSolverOperationStatus: getSolverOperationStatus,
		getDAppSignatories:       getDAppSignatories,
	}
}

func (s *Server) ListenAndServe(serverReadyChan chan struct{}) {
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: handlers.CORS(handlers.AllowedHeaders([]string{"Content-Type"}))(s.router),
	}
	ln, err := net.Listen("tcp", httpServer.Addr)
	if err != nil {
		panic(err)
	}

	log.Info("server started")
	if serverReadyChan != nil {
		close(serverReadyChan)
	}

	err = httpServer.Serve(ln)
	log.Info("server stopped", "err", err)
}

func (s *Server) ServeWs(w http.ResponseWriter, r *http.Request, channel string, bundler common.Address, signatory common.Address) (*Conn, error) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Info("failed upgrading connection", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%s: %s", ConnUpgradeFailed, err.Error())))
		return nil, err
	}

	conn := NewConn(c, channel, bundler, signatory)
	doneChan := make(chan struct{})

	go s.writePump(conn, doneChan)
	go s.readPump(conn, doneChan)

	return conn, nil
}

func (s *Server) ServeWsSolver(w http.ResponseWriter, r *http.Request) {
	s.ServeWs(w, r, ChannelSolver, common.Address{}, common.Address{})
}

func (s *Server) ServeWsBundler(w http.ResponseWriter, r *http.Request, bundler common.Address) {
	conn, err := s.ServeWs(w, r, ChannelBundler, bundler, common.Address{})
	if err != nil {
		log.Info("failed to serve ws bundler", "err", err)
		return
	}

	s.registerBundler(conn)
	conn.SetCloseHandler(func(code int, text string) error {
		s.unregisterBundler(conn)
		return nil
	})
}

func (s *Server) ServeWsSignatory(w http.ResponseWriter, r *http.Request, signatory common.Address) {
	conn, err := s.ServeWs(w, r, ChannelSignatory, common.Address{}, signatory)
	if err != nil {
		log.Info("failed to serve ws signatory", "err", err)
		return
	}

	s.registerSignatory(conn)
	conn.SetCloseHandler(func(code int, text string) error {
		s.unregisterSignatory(conn)
		return nil
	})
}

func (s *Server) registerBundler(conn *Conn) {
	if conn.channel != ChannelBundler || conn.bundler == (common.Address{}) {
		log.Info("invalid bundler connection")
		return
	}

	log.Info("bundler connected", "bundler", conn.bundler.Hex())

	s.mu.Lock()
	defer s.mu.Unlock()

	s.bundlers[conn.bundler] = conn
}

func (s *Server) unregisterBundler(conn *Conn) {
	if conn.bundler == (common.Address{}) {
		log.Info("invalid bundler connection")
		return
	}

	log.Info("bundler disconnected", "bundler", conn.bundler.Hex())

	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.bundlers, conn.bundler)
}

func (s *Server) registerSignatory(conn *Conn) {
	if conn.channel != ChannelSignatory || conn.signatory == (common.Address{}) {
		log.Info("invalid signatory connection")
		return
	}

	log.Info("signatory connected", "signatory", conn.signatory.Hex())

	s.mu.Lock()
	defer s.mu.Unlock()

	s.signatories[conn.signatory] = conn
}

func (s *Server) unregisterSignatory(conn *Conn) {
	if conn.signatory == (common.Address{}) {
		log.Info("invalid signatory connection")
		return
	}

	log.Info("signatory disconnected", "signatory", conn.signatory.Hex())

	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.signatories, conn.signatory)
}

func (s *Server) BroadcastUserOperationPartial(userOperationPartialRaw *operation.UserOperationPartialRaw) {
	broadcast := &Broadcast{
		Event: EventUpdate,
		Topic: TopicNewUserOperations,
		Data: &BroadcastParams{
			UserOperationPartial: userOperationPartialRaw,
		},
	}
	s.publish(broadcast)
}

func (s *Server) ForwardBundle(bundleOps *operation.BundleOperations, setAtlasTxHash setAtlasTxHashFn, setRelayError setRelayErrorFn) *relayerror.Error {
	bundlingRequest := &bundlingRequest{
		candidatesBundlers: make(map[common.Address]*Conn),
		offlineBundlers:    make(map[common.Address]bool),
		setAtlasTxHash:     setAtlasTxHash,
		setRelayError:      setRelayError,
		doneChan:           make(chan struct{}),
	}

	var firstCandidate *Conn

	s.mu.Lock()
	defer s.mu.Unlock()

	primaryBundler, primaryOnline := s.bundlers[bundleOps.DAppOperation.Bundler]
	if primaryOnline {
		firstCandidate = primaryBundler
		bundlingRequest.candidatesBundlers[bundleOps.DAppOperation.Bundler] = primaryBundler
	}

	// Retrieve dApp signatories, which are allowed bundlers
	signatories, relayErr := s.getDAppSignatories(bundleOps.DAppOperation.Control)
	if relayErr != nil {
		log.Info("failed to get dApp signatories", "control", bundleOps.DAppOperation.Control.Hex(), "err", relayErr.Message)
		if !primaryOnline {
			return ErrBundlerOffline
		}
	}

	for _, signatory := range signatories {
		conn, online := s.bundlers[signatory]
		if online {
			bundlingRequest.candidatesBundlers[signatory] = conn
		}
		if firstCandidate == nil {
			firstCandidate = conn
		}
	}

	if len(bundlingRequest.candidatesBundlers) == 0 {
		log.Info("no online bundler", "control", bundleOps.DAppOperation.Control.Hex())
		return ErrBundlerOffline
	}

	id := uuid.New().String()
	s.bundlingRequests[id] = bundlingRequest

	bundleReq := &BundleRequest{
		Id:     id,
		Event:  EventNewBundle,
		Bundle: bundleOps.EncodeToRaw(),
	}

	go s.runBundlingRequest(id, bundleReq, firstCandidate, bundlingRequest)
	return nil
}

func (s *Server) runBundlingRequest(id string, bundleReq *BundleRequest, firstCandidate *Conn, bundlingRequest *bundlingRequest) {
	nextCandidate := firstCandidate

	getNextCandidate := func() *Conn {
		for candidate, conn := range bundlingRequest.candidatesBundlers {
			if bundlingRequest.offlineBundlers[candidate] {
				continue
			}
			return conn
		}
		return nil
	}

	for {
		if nextCandidate == nil {
			// No more candidates
			s.mu.Lock()
			delete(s.bundlingRequests, id)
			s.mu.Unlock()
			bundlingRequest.setRelayError(ErrBundlerOffline)
			return
		}

		if relayErr := nextCandidate.send(bundleReq); relayErr != nil {
			log.Info("failed to send bundle request", "bundler", nextCandidate.bundler.Hex(), "err", relayErr.Message)
			bundlingRequest.offlineBundlers[nextCandidate.bundler] = true
			nextCandidate = getNextCandidate()
			continue
		}

		select {
		case <-time.After(BundlerTimeout):
			log.Info("bundler timed out", "bundler", nextCandidate.bundler.Hex())
			bundlingRequest.offlineBundlers[nextCandidate.bundler] = true
			nextCandidate = getNextCandidate()

		case <-bundlingRequest.doneChan:
			// Got a response, exit
			return
		}
	}
}

func (s *Server) NewSignatoryRequest(userOp *operation.UserOperation, solverOps []*operation.SolverOperation, submitBundleOperations submitBundleOperationsFn) *relayerror.Error {
	signingRequest := &SigningRequest{
		candidatesSignatories:  make(map[common.Address]*Conn),
		offlineSignatories:     make(map[common.Address]bool),
		submitBundleOperations: submitBundleOperations,
		bundle:                 &operation.BundleOperations{UserOperation: userOp, SolverOperations: solverOps},
		doneChan:               make(chan struct{}),
	}

	var firstCandidate *Conn

	s.mu.Lock()
	defer s.mu.Unlock()

	// Retrieve dApp signatories
	signatories, relayErr := s.getDAppSignatories(userOp.Control)
	if relayErr != nil {
		log.Info("failed to get dApp signatories", "control", userOp.Control.Hex(), "err", relayErr.Message)
	}

	for _, signatory := range signatories {
		conn, online := s.signatories[signatory]
		if online {
			signingRequest.candidatesSignatories[signatory] = conn
		}
		if firstCandidate == nil {
			firstCandidate = conn
		}
	}

	if len(signingRequest.candidatesSignatories) == 0 {
		log.Info("no online signatory", "control", userOp.Control.Hex())
		return relayErr
	}

	id := uuid.New().String()
	s.signingRequests[id] = signingRequest

	solverOpsRaw := make([]*operation.SolverOperationRaw, 0, len(solverOps))
	for _, solverOp := range solverOps {
		solverOpsRaw = append(solverOpsRaw, solverOp.EncodeToRaw())
	}

	signatoryReq := &SignatoryRequest{
		Id:               id,
		Event:            EventNewSignatoryRequest,
		UserOperation:    userOp.EncodeToRaw(),
		SolverOperations: solverOpsRaw,
	}

	go s.runSignatoryRequest(id, signatoryReq, firstCandidate, signingRequest)
	return nil
}

func (s *Server) runSignatoryRequest(id string, signatoryReq *SignatoryRequest, firstCandidate *Conn, signingRequest *SigningRequest) {
	nextCandidate := firstCandidate

	getNextCandidate := func() *Conn {
		for candidate, conn := range signingRequest.candidatesSignatories {
			if signingRequest.offlineSignatories[candidate] {
				continue
			}
			return conn
		}
		return nil
	}

	for {
		if nextCandidate == nil {
			// No more candidates
			s.mu.Lock()
			delete(s.signingRequests, id)
			s.mu.Unlock()
			return
		}

		if relayErr := nextCandidate.send(signatoryReq); relayErr != nil {
			log.Info("failed to send signatory request", "signatory", nextCandidate.signatory.Hex(), "err", relayErr.Message)
			signingRequest.offlineSignatories[nextCandidate.signatory] = true
			nextCandidate = getNextCandidate()
			continue
		}

		select {
		case <-time.After(SignatoryTimeout):
			log.Info("signatory timed out", "signatory", nextCandidate.signatory.Hex())
			signingRequest.offlineSignatories[nextCandidate.signatory] = true
			nextCandidate = getNextCandidate()

		case <-signingRequest.doneChan:
			// Got a response, exit
			return
		}
	}
}

func (s *Server) processSolverMessage(conn *Conn, msg []byte) {
	var req *Request
	if err := json.Unmarshal(msg, &req); err != nil {
		log.Info("failed to unmarshal solver message", "err", err)
		conn.sendString(fmt.Sprintf("%s: %s", InvalidMessage, err.Error()))
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Info("invalid solver message", "err", err)
		conn.sendString(fmt.Sprintf("%s: %s", InvalidMessage, err.Error()))
		return
	}

	resp := &Response{
		Id: req.Id,
	}

	if _, valid := Methods[conn.channel][req.Method]; !valid {
		resp.Error = InvalidMethod
		conn.send(resp)
		return
	}

	switch req.Method {
	case MethodPing:
		resp.Result = Pong

	case MethodSubscribe:
		s.subscribe(conn, req.Params.Topic, resp)

	case MethodUnsubscribe:
		s.unsubscribe(conn, req.Params.Topic, resp)

	case MethodSubmitSolverOperation:
		s.processNewSolverOperation(req.Params.SolverOperation.Decode(), resp)

	case MethodSolverOperationStatus:
		s.processGetSolverOperationStatus(req.Params.OpHash, resp)

	default:
		resp.Error = InvalidMethod
	}

	conn.send(resp)
}

func (s *Server) processBundlerMessage(msg []byte) {
	var resp *BundleResponse
	if err := json.Unmarshal(msg, &resp); err != nil {
		log.Info("failed to unmarshal bundler message", "err", err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(resp); err != nil {
		log.Info("invalid bundler message", "err", err)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	bundlingRequest, exist := s.bundlingRequests[resp.Id]
	if !exist {
		log.Info("bundler response not found", "id", resp.Id)
		return
	}

	close(bundlingRequest.doneChan)
	delete(s.bundlingRequests, resp.Id)

	if len(resp.Error) > 0 {
		log.Info("bundler response error", "err", resp.Error)
		bundlingRequest.setRelayError(ErrBundlingFailure.AddMessage(resp.Error))
		return
	}

	if resp.Result == (common.Hash{}) {
		log.Info("bundler response invalid tx hash", "txHash", resp.Result.Hex())
		bundlingRequest.setRelayError(ErrBundlingFailure)
		return
	}

	bundlingRequest.setAtlasTxHash(resp.Result)
}

func (s *Server) processSignatoryMessage(msg []byte) {
	var resp *SignatoryResponse
	if err := json.Unmarshal(msg, &resp); err != nil {
		log.Info("failed to unmarshal signatory message", "err", err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(resp); err != nil {
		log.Info("invalid signatory message", "err", err)
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	signingRequest, exist := s.signingRequests[resp.Id]
	if !exist {
		log.Info("signatory response not found", "id", resp.Id)
		return
	}

	close(signingRequest.doneChan)
	delete(s.signingRequests, resp.Id)

	if len(resp.Error) > 0 {
		log.Info("signatory response error", "err", resp.Error)
		return
	}

	if resp.Result == nil {
		log.Info("signatory response invalid dApp operation")
		return
	}

	signingRequest.bundle.DAppOperation = resp.Result.Decode()
	signingRequest.submitBundleOperations(signingRequest.bundle)
}

func (s *Server) publish(broadcast *Broadcast) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, exist := s.subscriptions[broadcast.Topic]; !exist {
		log.Info("no subscribers for this topic", "topic", broadcast.Topic)
		return
	}

	for _, conn := range s.subscriptions[broadcast.Topic] {
		conn.send(broadcast)
	}
}

func (s *Server) subscribe(conn *Conn, topic string, resp *Response) {
	if _, valid := Topics[topic]; !valid {
		log.Info("invalid topic", "topic", topic)
		resp.Error = InvalidTopic
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exist := s.subscriptions[topic]; !exist {
		s.subscriptions[topic] = make(map[string]*Conn)
	}

	if _, subbed := s.subscriptions[topic][conn.uuid]; subbed {
		log.Info("already subscribed", "topic", topic, "uuid", conn.uuid)
		resp.Error = AlreadySubscribed
		return
	}

	s.subscriptions[topic][conn.uuid] = conn
	resp.Result = Subscribed
}

func (s *Server) unsubscribe(conn *Conn, topic string, resp *Response) {
	if _, valid := Topics[topic]; !valid {
		log.Info("invalid topic", "topic", topic)
		resp.Error = InvalidTopic
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exist := s.subscriptions[topic][conn.uuid]; !exist {
		log.Info("not subscribed", "topic", topic, "uuid", conn.uuid)
		resp.Error = NotSubscribed
		return
	}

	delete(s.subscriptions[topic], conn.uuid)
	resp.Result = Unsubscribed
}

func (s *Server) removeSubscriber(uuid string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, client := range s.subscriptions {
		delete(client, uuid)
	}
}

func (s *Server) processNewSolverOperation(solverOp *operation.SolverOperation, resp *Response) {
	if solverOp == nil {
		log.Info("invalid solver operation")
		resp.Error = InvalidSolverOperation
		return
	}

	solverOpHash, relayErr := s.newSolverOperation(solverOp)
	if relayErr != nil {
		log.Info("failed to submit solver operation", "err", relayErr.Error())
		resp.Error = relayErr.Error()
		return
	}

	resp.Result = solverOpHash
}

func (s *Server) processGetSolverOperationStatus(solverOpHash common.Hash, resp *Response) {
	solverStatus, relayErr := s.getSolverOperationStatus(solverOpHash, nil)
	if relayErr != nil {
		log.Info("failed to get solver operation status", "err", relayErr.Error())
		resp.Error = relayErr.Error()
		return
	}

	resp.Result = solverStatus
}

func (s *Server) rateLimit(conn *Conn, msg []byte) {
	resp := &Response{
		Error: RateLimitExceeded,
	}

	defer conn.send(resp)

	var req *BareRequest
	if err := json.Unmarshal(msg, &req); err != nil {
		log.Info("failed to unmarshal rate limited message", "err", err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Info("invalid rate limited message", "err", err)
		return
	}

	resp.Id = req.Id
}

func (s *Server) readPump(conn *Conn, doneChan chan<- struct{}) {
	limiter := rate.NewLimiter(rate.Limit(websocketRateLimit), websocketBurstLimit)
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			close(doneChan)
			break
		}

		if !limiter.Allow() {
			s.rateLimit(conn, msg)
			continue
		}

		if conn.isBundler() {
			s.processBundlerMessage(msg)
		} else if conn.isSignatory() {
			s.processSignatoryMessage(msg)
		} else {
			s.processSolverMessage(conn, msg)
		}
	}
}

func (s *Server) writePump(conn *Conn, doneChan <-chan struct{}) {
	ticker := time.NewTicker(pingPeriod)
	defer ticker.Stop()

	defer func() {
		conn.Close()
		s.removeSubscriber(conn.uuid)
	}()

	for {
		select {
		case <-ticker.C:
			if err := conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait)); err != nil {
				return
			}

		case msg := <-conn.sendChan:
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				return
			}

		case <-doneChan:
			return
		}
	}
}
