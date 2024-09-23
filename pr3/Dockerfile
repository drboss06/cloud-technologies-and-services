FROM alpine:latest

# Обновление индекса пакетов и установка netcat
RUN apk update && apk add --no-cache netcat-openbsd

# Копируем серверный скрипт
COPY server.sh /usr/local/bin/server.sh

# Делаем скрипт исполняемым
RUN chmod +x /usr/local/bin/server.sh

# Запускаем сервер
CMD ["/usr/local/bin/server.sh"]
