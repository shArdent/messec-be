# Build stage
FROM golang:1.23 AS builder
WORKDIR /app

# Copy go mod dan download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code
COPY . .

# Build binary dari folder cmd
RUN go build -o server ./cmd

# Run stage
FROM debian:bookworm-slim
WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]

