package tests

import (
	"crypto/ecdsa"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/FastLane-Labs/atlas-operations-relay/log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/websocket"
)

func TestBundlerListen(t *testing.T) {
	sampleBundlerPk, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	runBundler(sampleBundlerPk, make(chan []byte), make(chan []byte))
}

func runBundler(bundlerPk *ecdsa.PrivateKey, bundlerReceiveChan chan []byte, bundlerSendChan chan []byte) {
	bundlerAddr := crypto.PubkeyToAddress(bundlerPk.PublicKey)
	timestamp := time.Now().Unix()
	signatureContent := fmt.Sprintf("%s:%d", bundlerAddr, timestamp)
	signature, err := signMessage([]byte(signatureContent), bundlerPk)
	if err != nil {
		panic(err)
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

	fmt.Printf("Connecting to %s\n", u.String())

	conn, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Error("dial:", err)
		return
	}

	if resp.StatusCode != 101 {
		log.Error("Expected status code 101, got", resp.StatusCode)
		return
	}
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
