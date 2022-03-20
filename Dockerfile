FROM golang:1.18

WORKDIR /app

COPY ./ /app

RUN go mod download

ENTRYPOINT go run cmd/analyzer/main.go