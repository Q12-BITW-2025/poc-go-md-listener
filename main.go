// main.go
package main

import (
	"log/slog"
	"os"
	"strings"

	"market-ws-listener/exchanges"
)

func getSymbols() []string {
	s := os.Getenv("SYMBOLS")
	if s == "" {
		s = "BTCUSDC,XRPUSDT,ETHUSDT,USDTUSD,BTCUSDT,XLMUSDT,DOGEUSDT,SOLUSDT,ADAUSDT,HBARUSDT"
		slog.Info("SYMBOLS not set, defaulting", "value", s)
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
	// Read symbols (defaulting inside each exchange handler)
	symbols := getSymbols()

	// Determine exchange from ENV (default Binance)
	exch := os.Getenv("EXCHANGE")
	if exch == "" {
		exch = "BINANCE"
		slog.Info("EXCHANGE not set, defaulting", "value", exch)
	}

	switch strings.ToUpper(exch) {
	case "BINANCE":
		exchanges.ConnectBinance(symbols)
	case "COINBASE":
		exchanges.ConnectCoinbase(symbols)
	case "KRAKEN":
		exchanges.ConnectKraken(symbols)
	default:
		slog.Error("Unsupported exchange", "value", exch)
	}
}
