# Этап сборки
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o musicinstruments .

# Финальный образ
FROM ubuntu:latest

LABEL authors="Алиби"
WORKDIR /app

COPY templates/ ./templates/

# Копируем бинарник из builder
COPY --from=builder /app/musicinstruments .

# Копируем wait-for-it.sh и даём права на выполнение
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh

# Ждём PostgreSQL перед запуском приложения
#CMD ["/wait-for-it.sh", "postgres:5432", "--", "./musicinstruments"]
