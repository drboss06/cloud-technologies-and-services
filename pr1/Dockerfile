FROM alpine:latest

RUN apk add --no-cache bash coreutils procps

COPY menu.sh /usr/local/bin/menu.sh

RUN chmod +x /usr/local/bin/menu.sh

CMD ["sh", "/usr/local/bin/menu.sh"]
