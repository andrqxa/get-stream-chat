# ======================
# Build stage
# ======================
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install ca-certificates for HTTPS (Go + Stream)
RUN apk add --no-cache ca-certificates

# Go deps
COPY go.mod go.sum ./
RUN go mod download

# Source
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o server ./cmd/server

# ======================
# Runtime stage
# ======================
FROM alpine:3.19

# Certificates for HTTPS calls (GetStream)
RUN apk add --no-cache ca-certificates

# Non-root user
RUN adduser -D appuser
USER appuser

WORKDIR /app

# Binary
COPY --from=builder /app/server /app/server

# PocketBase schema (NOT runtime data)
COPY pb/schema /app/pb/schema

# Runtime port (overridden by env / compose)
EXPOSE 3000

CMD ["./server"]
