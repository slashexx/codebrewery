FROM golang:1.20-alpine

WORKDIR /app
COPY . .

RUN go build -o code-execution-service .

CMD ["./code-execution-service"]
