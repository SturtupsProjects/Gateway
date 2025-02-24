FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
COPY .env ./
RUN go mod download

COPY . .

# Установка зависимостей для Kafka
RUN apk add --no-cache librdkafka-dev gcc libc-dev

RUN CGO_ENABLED=1 GOOS=linux go build -o main ./cmd/app/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates librdkafka

WORKDIR /app

COPY --from=builder /app .
COPY --from=builder /app/main .

RUN mkdir -p pkg/logs

EXPOSE 1111

RUN chmod +x ./main

CMD ["./main"]
