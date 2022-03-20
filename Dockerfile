FROM golang:1.18 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build analyzer app
RUN cd /app/cmd/analyzer && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./analyzer

# Build from scratch to minimize image size
FROM scratch
# Copy only pre-build binary
COPY --from=builder /app/cmd/analyzer/analyzer /app/

ENTRYPOINT ["/app/analyzer"]