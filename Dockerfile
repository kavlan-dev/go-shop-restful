# Сборка
FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/app

# Запуск
FROM alpine:latest
WORKDIR /app

# Копирование бинарника
COPY --from=builder /app/app .

# Порты и запуск
EXPOSE 8080
CMD ["./app"]
