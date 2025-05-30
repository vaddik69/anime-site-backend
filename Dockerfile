FROM golang:1.23-bullseye AS builder

WORKDIR /app

# Копируем go.mod и go.sum отдельно, чтобы кэш работал
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Копируем всё остальное
COPY . .

# Собираем бинарник с учетом правильного пути
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/server ./cmd/server

# Финальный образ
FROM alpine:latest

WORKDIR /app

# Копируем только бинарник и необходимые файлы
COPY --from=builder /app/server /app/server

# Для PostgreSQL может потребоваться дополнительные библиотеки
RUN apk add --no-cache libc6-compat

EXPOSE 8080

# Правильный путь к бинарнику
CMD ["/app/server"]