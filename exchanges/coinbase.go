// exchanges/coinbase.go
package exchanges

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/encoding/protojson"

	marketpb "market-ws-listener/model"
)

// ConnectCoinbase subscribes to Coinbase Exchange market trades
// and logs them as JSON.
func ConnectCoinbase(symbols []string) {
	const url = "wss://ws-feed.exchange.coinbase.com"

	// Coinbase will reject handshakes without a matching Origin:
	header := http.Header{
		"Origin": []string{"https://exchange.coinbase.com"},
	}

	conn, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		log.Fatal("Coinbase dial error:", err)
	}
	defer conn.Close()
	log.Printf("[Coinbase] Connected to %s", url)

	// Subscribe to matches channel
	sub := map[string]interface{}{
		"type":        "subscribe",
		"product_ids": symbols,
		"channels":    []interface{}{map[string]interface{}{"name": "matches", "product_ids": symbols}},
	}
	if err := conn.WriteJSON(sub); err != nil {
		log.Fatal("Coinbase subscribe error:", err)
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			return
		}

		// If Coinbase returns a top-level error object, log and skip.
		if bytes.Contains(msg, []byte(`"error"`)) {
			log.Printf("Received error from Coinbase: %s", msg)
			continue
		}

		// Ignore non-match messages
		var meta struct {
			Type string `json:"type"`
		}
		if json.Unmarshal(msg, &meta) == nil && meta.Type != "match" {
			continue
		}

		// Now unmarshal the actual trade
		var raw struct {
			Type      string      `json:"type"`
			TradeID   json.Number `json:"trade_id"`
			ProductID string      `json:"product_id"`
			Price     float64     `json:"price,string"`
			Size      float64     `json:"size,string"`
			Time      string      `json:"time"`
		}
		if err := json.Unmarshal(msg, &raw); err != nil {
			log.Println("json unmarshal error:", err)
			continue
		}

		// convert trade_id to string
		tradeIDStr := raw.TradeID.String()

		// Build and log canonical MarketData
		ts, err := time.Parse(time.RFC3339Nano, raw.Time)
		if err != nil {
			log.Println("time parse error:", err)
			continue
		}
		md := &marketpb.MarketData{
			Exchange:  "COINBASE",
			Symbol:    strings.ReplaceAll(raw.ProductID, "-", ""),
			Timestamp: ts.UnixNano() / int64(time.Millisecond),
			Price:     raw.Price,
			Size:      raw.Size,
			TradeId:   tradeIDStr,
		}
		jsonOut, err := protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(md)
		if err != nil {
			log.Println("json marshal error:", err)
			continue
		}
		log.Printf("Proto JSON: %s", jsonOut)
	}
}
