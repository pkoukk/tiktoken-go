#!/usr/bin/env bash

set -euo pipefail

# Copied from https://github.com/zurawiki/tiktoken-rs
# but vocab.bpe and encoder.json are not needed
#https://openaipublic.blob.core.windows.net/gpt-2/encodings/main/vocab.bpe
#https://openaipublic.blob.core.windows.net/gpt-2/encodings/main/encoder.json

ASSETS=$(cat <<EOF
https://openaipublic.blob.core.windows.net/encodings/r50k_base.tiktoken
https://openaipublic.blob.core.windows.net/encodings/p50k_base.tiktoken
https://openaipublic.blob.core.windows.net/encodings/cl100k_base.tiktoken
EOF
)

# Download assets
echo "Downloading assets..."
for asset in $ASSETS; do
    echo "Downloading $asset"
    curl -sSL -f -o "./assets/$(basename $asset)" $asset
done
