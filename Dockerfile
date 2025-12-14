# Build stage
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o /app/bin/cadastro-pessoas ./cmd/api

# Final stage
FROM scratch

# Copy SSL certificates for HTTPS
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy binary from builder
COPY --from=builder /app/bin/cadastro-pessoas /cadastro-pessoas

# Copy configuration files
COPY --from=builder /app/configs /configs

# Expose port
EXPOSE 8080

# Set entrypoint
ENTRYPOINT ["/cadastro-pessoas"]
