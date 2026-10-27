# Oxigen release to ship for Oxigen solutions (see docs/oxigen.md). Override
# with --build-arg OXIGEN_VERSION=<tag> from https://hub.docker.com/r/javanhut/oxigen.
ARG OXIGEN_VERSION=0.2.1

# Oxigen's release image: a static binary plus its standard library
FROM docker.io/javanhut/oxigen:${OXIGEN_VERSION} AS oxigen

# Stage 1: Build Go binary
FROM golang:1.22-bookworm AS builder

WORKDIR /app

# Copy go.mod first for better layer caching
COPY go.mod ./

# Copy source code
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o dsatutor ./cmd/tutor

# Stage 2: Runtime with Python and Oxigen for sandbox execution
FROM python:3.12-slim-bookworm

WORKDIR /app

# oxigen finds its stdlib at <prefix>/lib/oxigen/stdlib next to <prefix>/bin
COPY --from=oxigen /usr/local/bin/oxigen /usr/local/bin/oxigen
COPY --from=oxigen /usr/local/lib/oxigen/stdlib /usr/local/lib/oxigen/stdlib

# Copy the binary from builder stage
COPY --from=builder /app/dsatutor /app/dsatutor

# Expose default port
EXPOSE 8080

# Run the application
ENTRYPOINT ["/app/dsatutor"]
CMD ["-serve", "-addr=:8080"]
