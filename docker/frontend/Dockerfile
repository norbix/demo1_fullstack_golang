# Stage 1: Build
FROM golang:1.23.3 AS builder

# Set the working directory
WORKDIR /app

# Copy the build script and make it executable
COPY scripts/build.sh /app/scripts/build.sh
RUN chmod +x /app/scripts/build.sh

# Copy the Go module files to leverage Docker caching
COPY ./frontend/go.mod /app/frontend/
WORKDIR /app/frontend
RUN go mod download

# Copy the source code
COPY frontend/ /app/frontend

# Build the WebAssembly frontend using the build script
RUN /app/scripts/build.sh frontend

# Stage 2: Run
FROM nginx:alpine

# Set the working directory
WORKDIR /usr/share/nginx/html

# Copy static assets and the Wasm binary from the builder stage
COPY --from=builder /app/frontend/build/ /usr/share/nginx/html

# Copy the nginx configuration if needed (optional)
COPY docker/frontend/nginx.conf /etc/nginx/nginx.conf

# Expose the frontend port
EXPOSE 3000

# Start the nginx server
ENTRYPOINT ["nginx", "-g", "daemon off;"]
