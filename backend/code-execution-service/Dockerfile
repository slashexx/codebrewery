FROM ubuntu:20.04

# Update package lists and install necessary tools
RUN apt-get update && apt-get install -y \
    golang-go \
    python3 \
    python3-pip \
    gcc \
    g++ \
    default-jdk \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copy Go mod and main.go to download dependencies
COPY go.mod main.go ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o code-execution-service .

EXPOSE 8081 

CMD ["./code-execution-service"]
