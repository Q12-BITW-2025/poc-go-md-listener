// main.go
package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"market-ws-listener/exchanges"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func initTracer(ctx context.Context) func(context.Context) error {
	// Create OTLP gRPC exporter pointing at localhost:4317 (Odigos sidecar)
	exporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint("localhost:4317"),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("failed to create OTLP exporter: %v", err)
	}

	// Build a TracerProvider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(os.Getenv("OTEL_SERVICE_NAME")), // e.g. "market-ws-listener"
		)),
	)
	otel.SetTracerProvider(tp)

	// Use W3C Trace Context propagation
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Return shutdown func
	return tp.Shutdown
}

func getSymbols() []string {
	s := os.Getenv("SYMBOLS")
	if s == "" {
		s = "BTCUSDT, BTCBRL"
		log.Printf("SYMBOLS not set, defaulting to %s", s)
	}
	raw := strings.Split(s, ",")
	out := make([]string, 0, len(raw))
	for _, sym := range raw {
		o := strings.ToUpper(strings.TrimSpace(sym))
		out = append(out, o)
	}
	return out
}

func main() {
	
	ctx := context.Background()
	shutdown := initTracer(ctx)
	defer func() {
		// give exporter up to 5s to flush
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	// Read symbols (defaulting inside each exchange handler)
	symbols := getSymbols() // helpers in exchanges

	// Determine exchange from ENV (default Binance)
	exch := os.Getenv("EXCHANGE")
	if exch == "" {
		exch = "BINANCE"
		log.Printf("EXCHANGE not set, defaulting to %s", exch)
	}

	switch exch {
	case "BINANCE":
		exchanges.ConnectBinance(symbols)
	case "COINBASE":
		exchanges.ConnectCoinbase(symbols)
	case "KRAKEN":
		exchanges.ConnectKraken(symbols)
	default:
		log.Fatalf("unsupported exchange: %s", exch)
	}
}
