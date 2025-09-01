# Step 1: Build the Go binary
FROM golang:1.25.0 AS builder

# Set the working directory inside the container
WORKDIR /app/api-gateway

# Copy go.mod and go.sum files first (for caching dependencies)
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary
# RUN GOOS=linux GOARCH=amd64 go build -o api-gateway ./cmd/app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o api-gateway-binary ./cmd/app

# Step 2: Create a lightweight image to run the binary
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app/api-gateway

# # Copy the binary from the builder stage
COPY --from=builder /app/api-gateway/api-gateway-binary /app/api-gateway
COPY --from=builder /app/api-gateway/.env /app/api-gateway/.env

# Expose the port your application uses (optional)
EXPOSE 8080

# Command to run the binary
CMD ["/app/api-gateway/api-gateway-binary"]