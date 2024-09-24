FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/main

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main /app/main

COPY .env /app/.env

WORKDIR /app

EXPOSE 8080

CMD ["/app/main"]
