# Stage 1: Build the Go binary
FROM golang:1.22 AS builder

WORKDIR /app

# Copy the Go modules manifests
COPY go.mod ./
# Download the Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o echo-server .

# Stage 2: Create the final image
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/echo-server .

# Expose the port the server will run on
EXPOSE 8080

# Command to run the executable
CMD ["./echo-server"]
