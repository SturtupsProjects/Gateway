FROM golang:1.23.3 AS builder

WORKDIR /app

# Устанавливаем зависимости
RUN apt-get update && apt-get install -y \
    librdkafka-dev \
    gcc \
    libc-dev

COPY go.mod go.sum ./
COPY .env /
COPY internal/casbin /app/internal/casbin
RUN go mod download

COPY . .

# ВАЖНО: Включаем CGO, иначе не слинкуется с Kafka
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./cmd/app/main.go

# Финальный контейнер
FROM alpine:latest

RUN apk --no-cache add ca-certificates librdkafka libc6-compat

WORKDIR /app

COPY --from=builder /app/main /app/main

RUN mkdir -p /app/pkg/logs

EXPOSE 1111

RUN chmod +x /app/main

CMD ["/app/main"]
