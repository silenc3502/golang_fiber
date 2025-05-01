# syntax=docker/dockerfile:1
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o fiber-app

# runtime image
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/fiber-app .

EXPOSE 3773

CMD ["./fiber-app"]
