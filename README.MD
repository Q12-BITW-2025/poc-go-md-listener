# POC-GO-MD-LISTENER

A Go application that connects to multiple cryptocurrency exchanges (Binance, Coinbase, Kraken) via WebSocket to listen for live trade data, converts it into a canonical Protobuf `MarketData` model, and logs the JSON output.

---

## Features

* **Exchange support**: Binance, Coinbase, Kraken
* **Dynamic symbols**: Pass a comma-separated list of trading pairs via `SYMBOLS` environment variable
* **Exchange selection**: Choose the exchange via `EXCHANGE` environment variable (`BINANCE` by default)
* **Canonical Protobuf model**: All trades are normalized into a single `MarketData` schema
* **Docker-ready**: Includes Dockerfile for containerized builds with on-the-fly Protobuf generation

---

## Prerequisites

* Go 1.23+
* `protoc` compiler (v3.0+)
* `protoc-gen-go` plugin (v1.28+)
* Bash (for generation script)

On Alpine-based Linux you can install:

```bash
aapk add protobuf protobuf-dev bash git
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
```

---

## Project Layout

```plaintext
├── go.mod
├── go.sum
├── main.go              # entrypoint
├── helpers.go           # SYMBOLS env parsing
├── exchanges/           # exchange-specific connectors
│   ├── binance.go
│   ├── coinbase.go
│   ├── kraken.go
│   └── common.go
├── pb/                  # Protobuf definitions
│   ├── marketdata.proto
│   ├── binance.proto
│   ├── coinbase.proto
│   └── kraken.proto
├── model/               # generated Go bindings (via script)
│   ├── market/
│   ├── binance/
│   ├── coinbase/
│   └── kraken/
├── scripts/
│   └── gen-proto-docker.sh  # Protobuf generation helper
├── Dockerfile           # multi-stage container build
└── README.md
```

---

## Building & Protobuf Generation

1. **Generate Protobuf bindings**:

   ```bash
   scripts/gen-proto-docker.sh
   ```

   This compiles `pb/*.proto` into Go packages under `model/`.

2. **Compile the binary**:

   ```bash
   go build -o market-ws-listener ./
   ```

---

## Running Locally

```bash
# Default: Binance, BTCUSDT & BTCBRL
EXCHANGE=BINANCE SYMBOLS=BTCUSDT,BTCBRL ./market-ws-listener
```

Examples:

```bash
# Listen to ETH-USDT on Binance
EXCHANGE=BINANCE SYMBOLS=ETHUSDT ./market-ws-listener

# Listen to BTC-USD on Coinbase Exchange
EXCHANGE=COINBASE SYMBOLS=BTC-USD ./market-ws-listener

# Listen to XBT/USD on Kraken
EXCHANGE=KRAKEN SYMBOLS=XBT/USD ./market-ws-listener
```

---

## Docker

1. **Build the Docker image**:

   ```bash
   docker build -t poc-go-md-listener:latest .
   ```

2. **Run the container**:

   ```bash
   docker run -e EXCHANGE=KRAKEN -e SYMBOLS=XBT/USD poc-go-md-listener:latest
   ```

---

## Customization

* **Add new exchanges**: Implement a new file in `exchanges/` following existing patterns and register in `main.go`.
* **Extend Protobuf**: Modify `pb/*.proto` and rerun `scripts/gen-proto-docker.sh`.

---
