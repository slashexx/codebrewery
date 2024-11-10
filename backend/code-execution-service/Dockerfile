FROM golang:1.23-alpine


ENV DEBIAN_FRONTEND=noninteractive

RUN apk add --no-cache \
    python3 \
    gcc \
    g++ \
    openjdk11 \
    bash \
    rust \
    cargo


RUN javac -version || echo "javac not found"

WORKDIR /app

COPY go.mod main.go ./
RUN go mod download

COPY . .

RUN go build -o code-execution-service .

EXPOSE 8081 

CMD ["./code-execution-service"]
