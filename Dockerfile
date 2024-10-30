# syntax=docker/dockerfile:1

FROM golang:1.23

# Set destination for COPY
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping ./cmd/main.go # Adjust this path to your main Go file

# Expose the port your application runs on
EXPOSE 8080

# Run the compiled binary
CMD ["/docker-gs-ping"]

