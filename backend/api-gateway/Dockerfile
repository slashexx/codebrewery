FROM golang:1.23-alpine

WORKDIR /app
COPY . .

RUN go build -o api-gateway .

CMD ["./api-gateway"]
