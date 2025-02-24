# Используем Go 1.23.3 как билдовый образ
FROM golang:1.23.3 AS builder

# Устанавливаем нужные зависимости для Kafka
RUN apt-get update && apt-get install -y git gcc libc-dev librdkafka-dev

WORKDIR /app

# Копируем файлы зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY . .

# Включаем CGO для работы с Kafka
ENV CGO_ENABLED=1

# Сборка бинарника
RUN go build -o main ./cmd/app/main.go

# Финальный образ с Alpine
FROM alpine:latest

RUN apk --no-cache add ca-certificates librdkafka

WORKDIR /app

# Копируем скомпилированный бинарник
COPY --from=builder /app/main .

# Даем права на запуск
RUN chmod +x ./main

# Открываем порт
EXPOSE 1111

# Запускаем приложение
CMD ["./main"]
