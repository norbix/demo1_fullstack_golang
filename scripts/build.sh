#!/bin/bash

if [ "$1" == "frontend" ]; then
    echo "Building frontend..."
    GOOS=js GOARCH=wasm go build -o ./build/main.wasm ./cmd/ui/main.go
    cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./build/
    cp ./templates/index.html ./build/
    cp ./assets/* ./build/
    echo "Frontend build complete!"
else
    echo "Unknown build target: $1"
    exit 1
fi
