# Stage 1: Build
FROM golang:1.23.3 AS builder

# Set the working directory
WORKDIR /app

# Copy the build script and make it executable
COPY scripts/build.sh /app/scripts/build.sh
RUN chmod +x /app/scripts/build.sh

# Copy the Go module files first to leverage Docker caching
COPY ./frontend/go.mod /app/frontend/
WORKDIR /app/frontend
RUN go mod download

# Copy the source code
COPY frontend/ /app/frontend

# Execute the build script for the frontend
RUN /app/scripts/build.sh frontend

# Stage 2: Run
FROM nginx:alpine

# Set the working directory
WORKDIR /usr/share/nginx/html

# Copy the built frontend files from the builder stage
COPY --from=builder /app/frontend/build/ /usr/share/nginx/html

# Expose the frontend port
EXPOSE 3000

# Start the nginx server
ENTRYPOINT ["nginx", "-g", "daemon off;"]
