FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o sspanel-metron-go cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/sspanel-metron-go .

COPY config/config.yaml ./config.yaml

EXPOSE 8080

CMD ["./sspanel-metron-go", "-c", "config.yaml"]
