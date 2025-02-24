FROM golang:1.23.3 AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y \
    librdkafka-dev \
    gcc \
    libc-dev

COPY go.mod go.sum ./
COPY .env ./
RUN go mod download

COPY . .

# Статическая линковка, исправляет проблему "no such file or directory"
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/app/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates librdkafka libc6-compat

WORKDIR /app

COPY --from=builder /app/main /app/main

RUN mkdir -p /app/pkg/logs

EXPOSE 1111

RUN chmod +x /app/main

CMD ["/app/main"]
