# Stage 1: Build
FROM golang:1.21 AS builder

WORKDIR /app

COPY scripts/build.sh /app/scripts/build.sh
RUN chmod +x /app/scripts/build.sh

COPY . /app
RUN ./scripts/build.sh frontend

# Stage 2: Run
FROM nginx:alpine

WORKDIR /usr/share/nginx/html
COPY --from=builder /app/build/frontend/ /usr/share/nginx/html
EXPOSE 3000
