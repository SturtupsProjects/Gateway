FROM golang:1.23.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# ✅ Используем apk вместо apt-get
RUN apk add --no-cache librdkafka-dev gcc musl-dev

# ✅ Статическая сборка (убирает зависимости на libc)
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./cmd/app/main.go

# Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates librdkafka

WORKDIR /app

COPY --from=builder /app/main .

RUN chmod +x ./main

CMD ["./main"]
