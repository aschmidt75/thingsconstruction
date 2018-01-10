FROM alpine:latest

RUN apk update \
    && apk add jq \
    && adduser -h /app -s /bin/ash -D -u 766 tc

VOLUME /in
VOLUME /out

USER tc
WORKDIR /app

ADD main.sh /app

ENTRYPOINT [ "/app/main.sh" ]