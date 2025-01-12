# Builder
FROM golang:1.23.4-bullseye AS builder

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