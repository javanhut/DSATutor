# Stage 1: Build Go binary
FROM golang:1.22-bookworm AS builder

WORKDIR /app

# Copy go.mod first for better layer caching
COPY go.mod ./

# Copy source code
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o dsatutor ./cmd/tutor

# Stage 2: Runtime with Python for sandbox execution
FROM python:3.12-slim-bookworm

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/dsatutor /app/dsatutor

# Expose default port
EXPOSE 8080

# Run the application
ENTRYPOINT ["/app/dsatutor"]
CMD ["-serve", "-addr=:8080"]
