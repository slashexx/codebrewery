FROM golang:1.23-alpine

# Install bash (or any other tools you need)
RUN apk add --no-cache bash

WORKDIR /app

COPY go.mod main.go ./
RUN go mod download


COPY . .


RUN go build -o code-execution-service .


EXPOSE 8081 


CMD ["./code-execution-service"]
