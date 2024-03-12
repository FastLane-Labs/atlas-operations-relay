package tests

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/core"
	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/FastLane-Labs/atlas-operations-relay/operation"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/websocket"
)

func runBundler(bundlerPk *ecdsa.PrivateKey, bundlerReceiveChan chan []byte, bundlerSendChan chan []byte) {
	bundlerAddr := crypto.PubkeyToAddress(bundlerPk.PublicKey)
	timestamp := time.Now().Unix()
	signatureContent := fmt.Sprintf("%s:%d", bundlerAddr, timestamp)
	signature, err := signMessage([]byte(signatureContent), bundlerPk)
	if err != nil {
		log.Error("failed to sign message:", err)
		return
	}

	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/ws/bundler",
		RawQuery: url.Values{
			"address":   []string{bundlerAddr.Hex()},
			"timestamp": []string{fmt.Sprintf("%d", timestamp)},
			"signature": []string{hexutil.Encode(signature)},
		}.Encode(),
	}

	conn, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Error("dial:", err)
		return
	}

	if resp.StatusCode != 101 {
		log.Error("Expected status code 101, got", resp.StatusCode)
		return
	}

	log.Info("bundler connected", "bundlerEoa", bundlerAddr.Hex())

	// start listening on connection
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Error("bundler ws error:", err)
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Error("error:", err)
				}
				break
			}

			bundlerReceiveChan <- message
		}
	}()

	// start listening on bundlerSendChan
	go func() {
		for {
			message := <-bundlerSendChan
			err := conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Error("bundler ws error:", err)
				break
			}
		}
	}()
}

func sendBundleOperation(t *testing.T, userOp *operation.UserOperation, solverOps []*operation.SolverOperation, dAppOp *operation.DAppOperation) error {
	bundleOps := &operation.BundleOperations{
		UserOperation:    userOp,
		SolverOperations: solverOps,
		DAppOperation:    dAppOp,
	}

	bundleOpsJSON, err := json.Marshal(bundleOps)
	if err != nil {
		t.Errorf("failed to marshal bundle operations: %v", err)
	}

	_, err = http.Post("http://localhost:8080/bundleOperations", "application/json", bytes.NewReader(bundleOpsJSON))
	if err != nil {
		return err
	}

	userOpHash, _ := userOp.Hash()
	log.Info("relay sent bundleOps", "userOpHash", userOpHash.Hex(), "nSolverOps", len(bundleOps.SolverOperations))

	return nil
}

func sendBundlerResposne(txHash common.Hash, id string, bundlerSendChan chan []byte) error {
	bundleResponse := &core.BundleResponse{
		Id:     id,
		Result: txHash,
	}

	bundleResponseJSON, err := json.Marshal(bundleResponse)
	if err != nil {
		return err
	}

	bundlerSendChan <- bundleResponseJSON

	log.Info("bundler sent bundleResponse", "txHash", txHash.Hex())
	return nil
}

func signMessage(data []byte, privKey *ecdsa.PrivateKey) ([]byte, error) {
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(data))
	prefixedData := []byte(prefix + string(data))

	hash := crypto.Keccak256Hash(prefixedData)

	signature, err := crypto.Sign(hash.Bytes(), privKey)
	if err != nil {
		return nil, err
	}

	// According to the yellow paper, the V value in signature (27 or 28) is expected
	// by most libraries, including web3. However, Go Ethereum `Sign` method produces
	// 0 or 1 for V. So we adjust V back to 27 or 28 to ensure compatibility.
	if signature[64] < 27 {
		signature[64] += 27
	}

	return signature, nil
}
