# Builder
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download


COPY . .

# Production
RUN go build -o main ./cmd/main

FROM debian:bookworm-slim

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 9010

CMD ["./main"]