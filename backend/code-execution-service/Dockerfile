# Stage 1: Build
FROM golang:1.18-alpine AS build

WORKDIR /app


COPY go.mod main.go ./
RUN go mod download


RUN go build -o codeexec ./main.go  


FROM alpine:latest

WORKDIR /app


COPY --from=build /app/codeexec .

CMD ["./codeexec"]
