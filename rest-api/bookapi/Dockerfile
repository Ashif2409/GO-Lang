# Use the official Golang image to build the binary
FROM golang:1.23.4 AS builder 

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency resolution
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application files
COPY . .

# Build the Go binary for Linux (ensures compatibility)
RUN GOOS=linux GOARCH=amd64 go build -o bookapi .

# Use Ubuntu as the base image for the final container
FROM ubuntu:latest

# Install required dependencies
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Set the working directory in the final container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/bookapi ./bookapi

# Expose the application port
EXPOSE 8000

# Command to run the application
CMD ["./bookapi"]
