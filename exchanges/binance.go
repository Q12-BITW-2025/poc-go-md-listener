package exchanges

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	marketpb "market-ws-listener/model"
	"net/url"
	"strconv"
	"strings"
)

// ConnectBinance subscribes to trade streams and converts to MarketData
func ConnectBinance(symbols []string) {
	// build param
	streams := []string{}
	for _, sym := range symbols {
		streams = append(streams, strings.ToLower(sym)+"@trade")
	}
	param := strings.Join(streams, "/")
	u := url.URL{Scheme: "wss", Host: "stream.binance.us:9443", Path: "/stream", RawQuery: "streams=" + param}
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Binance dial error:", err)
	}
	defer conn.Close()
	log.Printf("[Binance] Connected to %s", u.String())

	// Handler to convert raw Binance JSON into MarketData proto
	rawHandler := func(msg []byte) (*marketpb.MarketData, error) {
		// Binance combined stream payload
		var wrapper struct {
			Stream string `json:"stream"`
			Data   struct {
				EventType    string `json:"e"`
				EventTime    int64  `json:"E"`
				Symbol       string `json:"s"`
				TradeID      int64  `json:"t"`
				PriceStr     string `json:"p"`
				QuantityStr  string `json:"q"`
				TradeTime    int64  `json:"T"`
				IsBuyerMaker bool   `json:"m"`
			} `json:"data"`
		}
		if err := json.Unmarshal(msg, &wrapper); err != nil {
			return nil, fmt.Errorf("json unmarshal error: %w", err)
		}

		// Parse numeric fields
		price, err := strconv.ParseFloat(wrapper.Data.PriceStr, 64)
		if err != nil {
			return nil, fmt.Errorf("parse price error: %w", err)
		}
		qty, err := strconv.ParseFloat(wrapper.Data.QuantityStr, 64)
		if err != nil {
			return nil, fmt.Errorf("parse quantity error: %w", err)
		}

		// Build MarketData proto
		md := &marketpb.MarketData{
			Exchange:  "BINANCE",
			Symbol:    wrapper.Data.Symbol,
			Timestamp: wrapper.Data.TradeTime,
			Price:     price,
			Size:      qty,
			TradeId:   fmt.Sprint(wrapper.Data.TradeID),
		}
		return md, nil
	}
	// Start listening loop with JSON-to-proto conversion
	listenLoop(conn, "BINANCE", rawHandler)
}
