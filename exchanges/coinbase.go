// exchanges/coinbase.go
package exchanges

import (
	"bytes"
	"encoding/json"
	"log/slog"
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
	const wsURL = "wss://ws-feed.exchange.coinbase.com"

	// Coinbase requires a valid Origin header for the handshake
	header := http.Header{
		"Origin": []string{"https://exchange.coinbase.com"},
	}

	conn, _, err := websocket.DefaultDialer.Dial(wsURL, header)
	if err != nil {
		slog.Error("Coinbase dial failed", "err", err)
		return // Don’t proceed if connection failed
	}
	defer conn.Close()

	slog.Info("[Coinbase] Connected", "url", wsURL)

	// Subscribe to matches channel
	sub := map[string]interface{}{
		"type":        "subscribe",
		"product_ids": symbols,
		"channels": []interface{}{
			map[string]interface{}{"name": "matches", "product_ids": symbols},
		},
	}

	if err := conn.WriteJSON(sub); err != nil {
		slog.Error("Coinbase subscribe failed", "err", err)
		return
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			slog.Warn("Coinbase read error", "err", err)
			return
		}

		// If Coinbase returns a top-level error object, log and skip.
		if bytes.Contains(msg, []byte(`"error"`)) {
			slog.Warn("Received error from Coinbase", "payload", string(msg))
			continue
		}

		// Ignore non-match messages
		var meta struct {
			Type string `json:"type"`
		}
		if json.Unmarshal(msg, &meta) == nil && meta.Type != "match" {
			continue
		}

		// Parse the actual trade
		var raw struct {
			Type      string      `json:"type"`
			TradeID   json.Number `json:"trade_id"`
			ProductID string      `json:"product_id"`
			Price     float64     `json:"price,string"`
			Size      float64     `json:"size,string"`
			Time      string      `json:"time"`
		}
		if err := json.Unmarshal(msg, &raw); err != nil {
			slog.Warn("JSON unmarshal error", "err", err)
			continue
		}

		// Convert trade ID to string
		tradeIDStr := raw.TradeID.String()

		// Build MarketData proto
		ts, err := time.Parse(time.RFC3339Nano, raw.Time)
		if err != nil {
			slog.Warn("time parse error", "err", err, "value", raw.Time)
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
			slog.Warn("protojson marshal error", "err", err)
			continue
		}

		// Log the final proto JSON
		slog.Info("Proto JSON", "data", string(jsonOut))
	}
}
