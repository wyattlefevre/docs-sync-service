# Dockerfile
# Start with the official Go image
FROM golang:1.21 as builder

# Set up working directory
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app
COPY . .

# Build the Go binary
RUN go build -o main .

# Use a minimal image to run the binary
FROM gcr.io/distroless/base-debian10

COPY --from=builder /app/main /main

# Run the binary
ENTRYPOINT ["/main"]

