FROM golang:1.23-alpine

# Set non-interactive mode for apt-get
ENV DEBIAN_FRONTEND=noninteractive

# Install additional required packages
RUN apk add --no-cache \
    python3 \
    gcc \
    g++ \
    openjdk11 \
    bash

WORKDIR /app

# Copy go.mod and main.go to download dependencies
COPY go.mod main.go ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o code-execution-service .

EXPOSE 8081 

CMD ["./code-execution-service"]
