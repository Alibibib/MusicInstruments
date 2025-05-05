FROM ubuntu:latest
LABEL authors="Алиби"

ENTRYPOINT ["top", "-b"]

# Используем официальный образ Golang
FROM golang:1.23 AS stage-2

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum отдельно, чтобы использовать кеш
COPY go.mod go.sum ./
RUN go mod download

# Копируем все остальные файлы
COPY . .

# Собираем бинарник
RUN go build -o musicinstruments .

# Открываем порт, например 8080
EXPOSE 8080

# Запуск приложения
CMD ["./musicinstruments"]
