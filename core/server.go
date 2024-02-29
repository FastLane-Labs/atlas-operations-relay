package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/FastLane-Labs/atlas-operations-relay/relayerror"
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

const (
	BundlerTimeout = 300 * time.Millisecond

	// Channels
	ChannelSolver  = "solver"
	ChannelBundler = "bundler"

	// Methods
	MethodPing                  = "ping"
	MethodSubscribe             = "subscribe"
	MethodUnsubscribe           = "unsubscribe"
	MethodSubmitSolverOperation = "submitSolverOperation"

	// Subscriptions topics
	TopicNewUserOperations = "newUserOperations"

	// Events
	EventUpdate    = "update"
	EventNewBundle = "newBundle"

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

	// Websocket config
	pongWait        = 60 * time.Second
	pingPeriod      = (pongWait * 9) / 10
	writeWait       = 10 * time.Second
	maxMessageSize  = 512
	readBufferSize  = 1024
	writeBufferSize = 1024
)

var (
	Channels = map[string]struct{}{
		ChannelSolver:  {},
		ChannelBundler: {},
	}

	Methods = map[string]map[string]struct{}{
		ChannelSolver:  {MethodPing: {}, MethodSubscribe: {}, MethodUnsubscribe: {}, MethodSubmitSolverOperation: {}},
		ChannelBundler: {MethodPing: {}},
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
	ErrBundlerOffline    = relayerror.NewError(1300, "bundler is offline")
	ErrCloggedConnection = relayerror.NewError(1301, "clogged connection")
	ErrBundlingFailure   = relayerror.NewError(1302, "bundling failure")
)

type newSolverOperationFn func(*operation.SolverOperation) *relayerror.Error
type getDAppSignatoriesFn func(common.Address) ([]common.Address, *relayerror.Error)
type setAtlasTxHashFn func(common.Hash) *relayerror.Error
type setRelayErrorFn func(*relayerror.Error) *relayerror.Error

type RequestParams struct {
	Topic           string                      `json:"topic,omitempty"`
	SolverOperation *operation.SolverOperation  `json:"solverOperation,omitempty"`
	Bundle          *operation.BundleOperations `json:"bundle,omitempty"`
}

type Request struct {
	Id     string         `json:"id"`
	Method string         `json:"method"`
	Params *RequestParams `json:"params,omitempty"`
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
	UserOperation *operation.UserOperation `json:"userOperation,omitempty"`
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
	Id     string                      `json:"id"`
	Event  string                      `json:"event"`
	Bundle *operation.BundleOperations `json:"bundle"`
}

func (br *BundleRequest) Marshal() []byte {
	b, _ := json.Marshal(br)
	return b
}

type BundleResponse struct {
	Id     string      `json:"id"`
	Result common.Hash `json:"result,omitempty"`
	Error  string      `json:"error,omitempty"`
}

type Marshaler interface {
	Marshal() []byte
}

type Conn struct {
	*websocket.Conn

	uuid     string
	channel  string
	bundler  common.Address
	sendChan chan []byte
}

func NewConn(conn *websocket.Conn, channel string, bundler common.Address) *Conn {
	return &Conn{
		Conn:     conn,
		uuid:     uuid.New().String(),
		channel:  channel,
		bundler:  bundler,
		sendChan: make(chan []byte, 256),
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

type bundlingRequest struct {
	candidatesBundlers map[common.Address]*Conn
	offlineBundlers    map[common.Address]bool
	setAtlasTxHash     setAtlasTxHashFn
	setRelayError      setRelayErrorFn
	doneChan           chan struct{}
}

type Server struct {
	router *mux.Router

	// Indexed by [topic][uuid]
	subscriptions map[string]map[string]*Conn

	// Indexed by [address]
	bundlers map[common.Address]*Conn

	// Indexed by [id]
	bundlingRequests map[string]*bundlingRequest

	newSolverOperation newSolverOperationFn
	getDAppSignatories getDAppSignatoriesFn

	mu sync.RWMutex
}

func NewServer(router *mux.Router, newSolverOperation newSolverOperationFn, getDAppSignatories getDAppSignatoriesFn) *Server {
	return &Server{
		router:             router,
		subscriptions:      make(map[string]map[string]*Conn),
		bundlers:           make(map[common.Address]*Conn),
		bundlingRequests:   make(map[string]*bundlingRequest),
		newSolverOperation: newSolverOperation,
		getDAppSignatories: getDAppSignatories,
	}
}

func (s *Server) ListenAndServe() {
	log.Info("server started")
	err := http.ListenAndServe(":8080", s.router)
	log.Info("server stopped", "err", err)
}

func (s *Server) ServeWs(w http.ResponseWriter, r *http.Request, channel string, bundler common.Address) (*Conn, error) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Info("failed upgrading connection", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("%s: %s", ConnUpgradeFailed, err.Error())))
		return nil, err
	}

	conn := NewConn(c, channel, bundler)
	doneChan := make(chan struct{})

	go s.writePump(conn, doneChan)
	go s.readPump(conn, doneChan)

	return conn, nil
}

func (s *Server) ServeWsSolver(w http.ResponseWriter, r *http.Request) {
	s.ServeWs(w, r, ChannelSolver, common.Address{})
}

func (s *Server) ServeWsBundler(w http.ResponseWriter, r *http.Request, bundler common.Address) {
	conn, err := s.ServeWs(w, r, ChannelBundler, bundler)
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

func (s *Server) registerBundler(conn *Conn) {
	if conn.channel != ChannelBundler || conn.bundler == (common.Address{}) {
		log.Info("invalid bundler connection")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.bundlers[conn.bundler] = conn
}

func (s *Server) unregisterBundler(conn *Conn) {
	if conn.bundler == (common.Address{}) {
		log.Info("invalid bundler connection")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.bundlers, conn.bundler)
}

func (s *Server) BroadcastUserOperation(userOp *operation.UserOperation) {
	broadcast := &Broadcast{
		Event: EventUpdate,
		Topic: TopicNewUserOperations,
		Data: &BroadcastParams{
			UserOperation: userOp,
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
		Bundle: bundleOps,
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

func (s *Server) processSolverMessage(conn *Conn, msg []byte) {
	var req *Request
	if err := json.Unmarshal(msg, &req); err != nil {
		log.Info("failed to unmarshal solver message", "err", err)
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
		s.processNewSolverOperation(conn, req.Params.SolverOperation, resp)

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

func (s *Server) processNewSolverOperation(conn *Conn, solverOp *operation.SolverOperation, resp *Response) {
	if solverOp == nil {
		log.Info("invalid solver operation")
		resp.Error = InvalidSolverOperation
		return
	}

	relayErr := s.newSolverOperation(solverOp)
	if relayErr != nil {
		log.Info("failed to submit solver operation", "err", relayErr.Error())
		resp.Error = relayErr.Error()
		return
	}

	resp.Result = SolverOpSuccessfullySubmitted
}

func (s *Server) readPump(conn *Conn, doneChan chan<- struct{}) {
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

		if conn.isBundler() {
			s.processBundlerMessage(msg)
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
		if conn.isBundler() {
			s.unregisterBundler(conn)
		}
	}()

	for {
		select {
		case <-ticker.C:
			if err := conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(writeWait)); err != nil {
				return
			}

		case msg := <-conn.sendChan:
			conn.WriteMessage(websocket.TextMessage, msg)
			conn.SetWriteDeadline(time.Now().Add(writeWait))

		case <-doneChan:
			return
		}
	}
}
