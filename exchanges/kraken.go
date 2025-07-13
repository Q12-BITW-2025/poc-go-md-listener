package exchanges

import (
	"bytes"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"log/slog"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
	marketpb "market-ws-listener/model"
)

// ConnectKraken subscribes and converts trades
func ConnectKraken(symbols []string) {
	conn, _, err := websocket.DefaultDialer.Dial("wss://ws.kraken.com", nil)
	if err != nil {
		slog.Error("Kraken dial failed", "err", err)
		return // Don’t continue on failed dial
	}
	defer conn.Close()

	slog.Info("[Kraken] Connected", "symbols", symbols)

	// Send subscription request
	sub := map[string]interface{}{
		"event":        "subscribe",
		"pair":         symbols,
		"subscription": map[string]string{"name": "trade"},
	}
	if err := conn.WriteJSON(sub); err != nil {
		slog.Error("Kraken subscribe failed", "err", err)
		return
	}

	// Handler to convert raw Kraken messages into MarketData
	rawHandler := func(msg []byte) (*marketpb.MarketData, error) {
		slog.Debug("[Kraken] Received raw message", "msg", string(msg))

		if len(msg) == 0 {
			return nil, nil
		}

		switch msg[0] {
		case '{':
			// JSON object → status/heartbeat/subscription msg; skip
			return nil, nil
		case '[':
			// JSON array → actual trade payload
		default:
			return nil, nil
		}

		// Unmarshal into generic slice
		var packet []interface{}
		if err := json.Unmarshal(msg, &packet); err != nil {
			return nil, fmt.Errorf("json unmarshal error: %w", err)
		}

		// Expect packet[1] to be a slice of trades
		trades, ok := packet[1].([]interface{})
		if !ok || len(trades) == 0 {
			return nil, nil
		}

		// First trade entry
		entry, ok := trades[0].([]interface{})
		if !ok || len(entry) < 3 {
			return nil, nil
		}

		// Parse fields: [ price, volume, time, ... ]
		priceStr, _ := entry[0].(string)
		volStr, _ := entry[1].(string)
		timeF, _ := entry[2].(float64)

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			return nil, fmt.Errorf("parse price error: %w", err)
		}

		volume, err := strconv.ParseFloat(volStr, 64)
		if err != nil {
			return nil, fmt.Errorf("parse volume error: %w", err)
		}

		// Convert seconds to milliseconds
		ts := int64(timeF * 1000)

		// Normalize symbol: e.g. "BTC/USD" → "BTCUSD"
		sym := strings.ToUpper(strings.ReplaceAll(symbols[0], "/", ""))

		return &marketpb.MarketData{
			Exchange:  "KRAKEN",
			Symbol:    sym,
			Timestamp: ts,
			Price:     price,
			Size:      volume,
			TradeId:   "", // Kraken doesn't provide trade ID here
		}, nil
	}

	// Start the listen loop
	for {
		_, msg, err := conn.ReadMessage()

		if err != nil {
			slog.Warn("Kraken read error", "err", err)
			return
		}

		if bytes.Contains(msg, []byte("error")) {
			slog.Warn("Received error message from Kraken", "payload", string(msg))
			continue
		}

		md, err := rawHandler(msg)
		if err != nil {
			slog.Warn("Conversion error", "err", err)
			continue
		}

		if md == nil {
			continue // Not a trade message
		}

		jsonData, err := protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(md)
		if err != nil {
			slog.Warn("protojson marshal error", "err", err)
			continue
		}

		slog.Info("Proto JSON", "data", string(jsonData))
	}
}
