# Stage 1: Build the binary
FROM golang:1.23.4-alpine AS builder

# Set up a working directory
WORKDIR /app

# Cache dependencies by copying go.mod and go.sum first
# COPY go.mod go.sum ./
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the binary (static binary)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/main ./cmd/main.go

# Stage 2: Create a minimal image
FROM scratch

# Copy the compiled binary from the builder
COPY --from=builder /app/bin/main /main

# Expose port (optional, only needed if app listens on a port)
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/main"]
