FROM golang:1.23.1 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/tmp/main ./cmd/main.go

FROM debian:bookworm

WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /app/tmp/main .

CMD ["./main"]