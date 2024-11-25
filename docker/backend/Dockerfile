# Stage 1: Build
FROM golang:1.21 AS builder

WORKDIR /app

# Copy the build script and make it executable
COPY scripts/build.sh /app/scripts/build.sh
RUN chmod +x /app/scripts/build.sh

# Copy the source code
COPY . /app

# Execute the build script for the backend
RUN ./scripts/build.sh backend

# Stage 2: Run (distroless)
FROM gcr.io/distroless/base-debian11 AS runner

WORKDIR /app

# Copy the backend binary from the build stage
COPY --from=builder /app/build/backend/main /app/main

# Expose backend port
EXPOSE 8080

# Use the distroless base entrypoint (no shell available)
ENTRYPOINT ["/app/main"]
