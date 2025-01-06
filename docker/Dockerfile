# Используем официальный образ Golang 1.22 для сборки
FROM golang:1.22 AS build

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum в контейнер
COPY ../go.mod ../go.sum ./

# Загружаем зависимости
RUN go mod tidy

# Копируем все исходники проекта
COPY ../ ./

## Копируем файл конфигурации
#COPY ../config/config.json /config/config.json

# Переходим в папку с исходниками main.go (например, cmd)
WORKDIR /app/cmd

# Собираем приложение
RUN go build -o yourapp .

# Используем более новый образ с обновленной версией glibc (Debian Bookworm)
FROM debian:bookworm-slim

# Устанавливаем необходимые зависимости (например, libc)
RUN apt-get update && apt-get install -y libc6

# Копируем скомпилированное приложение из предыдущего шага
COPY --from=build /app/cmd/yourapp /usr/local/bin/yourapp

# Копируем файл конфигурации из сборки (если нужно)
COPY --from=build /app/config/config.json /config/config.json

# Порт, на котором будет слушать ваше приложение
EXPOSE 8080

# Команда для запуска приложения
CMD ["/usr/local/bin/yourapp"]
