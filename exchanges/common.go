package exchanges

import (
	"context"
	"github.com/gorilla/websocket"
	"log/slog"
	"os"
	"os/signal"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"

	"google.golang.org/protobuf/encoding/protojson"
	marketpb "market-ws-listener/model"
)

// toProto constructs a canonical MarketData pb
func toProto(exchange, symbol string, timestamp int64, price, size float64, tradeID string) *marketpb.MarketData {
	return &marketpb.MarketData{
		Exchange:  exchange,
		Symbol:    symbol,
		Timestamp: timestamp,
		Price:     price,
		Size:      size,
		TradeId:   tradeID,
	}
}

// listenLoop reads messages, converts to canonical pb, and logs
func listenLoop(conn *websocket.Conn, exchange string, rawHandler func([]byte) (*marketpb.MarketData, error)) {
	tracer := otel.Tracer("market-data-ws-listener")
	done := make(chan struct{})

	// Use background context here; no need to pass it in
	ctx := context.Background()

	go func() {
		defer close(done)
		for {
			// Start a new span per message

			_, span := tracer.Start(ctx, "ws.MessageReceived")
			defer span.End()

			span.SetAttributes(attribute.String("event", "messageReceived"))
			span.SetAttributes(attribute.String("exchange", exchange))

			_, msg, err := conn.ReadMessage()
			if err != nil {
				span.RecordError(err)
				span.SetAttributes(attribute.Bool("success", false))
				span.End()
				slog.Warn("read error:", err)
				return
			}
			span.SetAttributes(attribute.Int("bytes", len(msg)))

			md, err := rawHandler(msg)
			if err != nil {
				span.RecordError(err)
				span.SetAttributes(attribute.Bool("conversion.success", false))
				span.End()
				slog.Error("conversion error:", err)
				continue
			}
			span.SetAttributes(
				attribute.String("symbol", md.Symbol),
				attribute.Float64("price", md.Price),
				attribute.Float64("size", md.Size),
				attribute.Bool("conversion.success", true),
			)

			if jsonData, err := protojson.Marshal(md); err != nil {
				span.RecordError(err)
				span.SetAttributes(attribute.Bool("json.success", false))
			} else {
				span.SetAttributes(attribute.Bool("json.success", true))
				slog.Info("MarketData proto JSON", "exchange", exchange, "data", string(jsonData))
			}

			span.End()
		}
	}()

	// graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	select {
	case <-done:
	case <-interrupt:
		slog.Info("interrupt received, closing connection")
		conn.WriteMessage(
			websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""),
		)
		select {
		case <-done:
		case <-time.After(time.Second):
		}
	}
}
