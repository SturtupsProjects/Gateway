FROM golang:1.23.3 AS builder

WORKDIR /app

# Устанавливаем зависимости, включая библиотеку Kafka
RUN apt-get update && apt-get install -y librdkafka-dev

COPY go.mod go.sum ./
COPY .env ./
RUN go mod download

COPY . .

# Указываем, что нам нужна поддержка CGO для Kafka
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./cmd/app/main.go

FROM alpine:latest

# Устанавливаем необходимые пакеты для работы Kafka в контейнере
RUN apk --no-cache add ca-certificates librdkafka-dev

WORKDIR /app

COPY --from=builder /app .
COPY --from=builder /app/main .

RUN mkdir -p pkg/logs

EXPOSE 1111

RUN chmod +x ./main

CMD ["./main"]
