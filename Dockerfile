FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Устанавливаем зависимости
RUN apt-get update && apt-get install -y librdkafka-dev gcc libc-dev

# Включаем CGO для поддержки C-библиотек
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./cmd/app/main.go

# Финальный образ (Alpine)
FROM alpine:latest

# Устанавливаем зависимости для работы бинарника
RUN apk --no-cache add ca-certificates librdkafka

WORKDIR /app

# Копируем собранный бинарник
COPY --from=builder /app/main .

RUN chmod +x ./main

CMD ["./main"]
