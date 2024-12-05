FROM golang:1.23-alpine

ENV DEBIAN_FRONTEND=noninteractive

# Install dependencies
RUN apk add --no-cache \
    python3 \
    gcc \
    g++ \
    openjdk11 \
    bash \
    curl

RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y

ENV PATH="/root/.cargo/bin:${PATH}"

RUN rustc --version || echo "rustc not found"

WORKDIR /app

COPY go.mod main.go ./
RUN go mod download

COPY . .

RUN go build -o code-execution-service .

EXPOSE 8081

CMD ["./code-execution-service"]
