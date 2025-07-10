#!/usr/bin/env bash
set -euo pipefail

# Determine script and project directories
script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROTO_DIR="$script_dir/../pb"
OUT_DIR="$script_dir/../model"

# Derive your Go module path from go.mod (unused here, but kept if needed)
MODULE_PATH="$(go list -m)"

# Create output directory if it doesn't exist
mkdir -p "$OUT_DIR"

echo "Generating protobufs from '$PROTO_DIR' into '$OUT_DIR' (module: $MODULE_PATH)"

# Build import-mapping flags for all your protos
# Each mapping: M<proto_filename>=<import_path>
MFLAGS="Mbinance.proto=${MODULE_PATH}/model;model","Mcoinbase.proto=${MODULE_PATH}/model;model","Mkraken.proto=${MODULE_PATH}/model;model","Mtrade.proto=${MODULE_PATH}/model;model"

# Compile each .proto into the flat model/ directory with override mappings .proto into the flat model/ directory with override mappings
for f in "$PROTO_DIR"/*.proto; do
  fname="$(basename "$f")"
  echo "🛠 protoc → $fname"
  protoc \
    -I "$PROTO_DIR" \
    --go_out="$MFLAGS,paths=source_relative:$OUT_DIR" \
    "$f"
done

echo "✅ Protobufs generated successfully into '$OUT_DIR'."
