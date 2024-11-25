# Stage 1: Build
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy only the go.mod and go.sum files first to leverage Docker's layer caching
COPY ./backend/go.mod ./
RUN go mod download

# Copy the entire backend source code
COPY ./backend/ ./

# Enable CGO and build the Go application
ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64
# Build the Go application
RUN go build -o backend ./cmd/api

# Stage 2: Run
# Non-distroless image with the required GLIBC
FROM debian:bookworm-slim

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/backend .

# Expose the application port
EXPOSE 8080

# Use the distroless base entrypoint
ENTRYPOINT ["/app/backend"]
