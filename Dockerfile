FROM golang:1.23.3 AS builder

WORKDIR /app

# Устанавливаем необходимые зависимости
RUN apt-get update && apt-get install -y \
    librdkafka-dev \
    gcc \
    libc-dev

COPY go.mod go.sum ./
COPY .env ./
RUN go mod download

COPY . .

# Включаем CGO, теперь сборка будет работать
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./cmd/app/main.go

FROM alpine:latest

# Добавляем runtime-зависимости для Kafka
RUN apk --no-cache add ca-certificates librdkafka

WORKDIR /app

COPY --from=builder /app .
COPY --from=builder /app/main .

RUN mkdir -p pkg/logs

EXPOSE 1111

RUN chmod +x ./main

CMD ["./main"]
