#!/bin/bash

set -e  # Exit immediately if a command exits with a non-zero status

# Helper function to print usage
usage() {
    echo "Usage: $0 {backend|frontend}"
    exit 1
}

# Ensure an argument is provided
if [ $# -ne 1 ]; then
    usage
fi

# Get the target component
TARGET=$1

# Function to build the backend
build_backend() {
    echo "Building backend binary..."
    go build -o build/backend/main .
    go build -o build/backend/main ./backend/cmd/api/main.go
    echo "Backend build complete!"
}

# Function to build the frontend
build_frontend() {
    echo "Building frontend WebAssembly..."
    mkdir -p build/frontend
    GOOS=js GOARCH=wasm go build -o build/frontend/main.wasm ./frontend/cmd/ui/main.go
    echo "Frontend build complete!"
}

# Main logic
case "$TARGET" in
    backend)
        build_backend
        ;;
    frontend)
        build_frontend
        ;;
    *)
        usage
        ;;
esac
