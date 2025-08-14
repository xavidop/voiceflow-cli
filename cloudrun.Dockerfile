# Build stage
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o voiceflow-cli .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates tzdata

# Create a non-root user
RUN adduser -D -s /bin/sh voiceflow

WORKDIR /app

# Copy the binary from builder stage and set proper ownership
COPY --from=builder --chown=voiceflow:voiceflow /app/voiceflow-cli .

# Make sure the binary is executable
RUN chmod +x voiceflow-cli

USER voiceflow

# Expose port for the server
EXPOSE 8080

# Set environment variables
ENV PORT=8080
ENV GIN_MODE=release

# Run the binary
CMD ["/app/voiceflow-cli", "server", "--port", "8080"]