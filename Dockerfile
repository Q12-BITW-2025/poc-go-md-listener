# syntax=docker/dockerfile:1

# === Builder Stage ===
FROM golang:1.23-bullseye AS builder

# Enable CGO for dynamic linking
ENV CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# Install system dependencies for protoc and building
RUN apt-get update && apt-get install -y --no-install-recommends \
    protobuf-compiler \
    libprotobuf-dev \
    bash \
    git \
    libc6-dev \
  && rm -rf /var/lib/apt/lists/*

# Install Go protobuf plugin
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1

WORKDIR /app/market-ws-listener

# Copy go.mod and go.sum, download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy proto definitions and generation script
COPY pb/ ./pb/
COPY scripts/gen-proto-docker.sh ./scripts/gen-proto-docker.sh
RUN chmod +x scripts/gen-proto-docker.sh \
    && ./scripts/gen-proto-docker.sh

# Copy application code
COPY main.go ./
COPY exchanges/ ./exchanges/

# Build the application (dynamic binary)
RUN go mod tidy
RUN go build -o market-ws-listener ./

# === Final Stage ===
FROM debian:bullseye-slim

# Install CA certificates for TLS
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
# Copy the dynamically-linked binary
COPY --from=builder /app/market-ws-listener ./

# Default environment variables
ENV EXCHANGE=BINANCE \
    SYMBOLS=BTCUSDT,BTCBRL \
    OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4317 \
    OTEL_SERVICE_NAME=market-ws-listener

ENTRYPOINT ["/app/market-ws-listener"]
