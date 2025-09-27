FROM golang:latest AS builder

WORKDIR /dictionary-api

COPY . .

ENV GOOS=linux

RUN go build -o backend.dictionary-api . && \
    chmod +x backend.dictionary-api

FROM ubuntu:latest

WORKDIR /dictionary-api

COPY --from=builder /dictionary-api/backend.dictionary-api .

ARG APP_USER=dictionary-api
ARG UID=10001
ARG GID=10001

RUN apt-get update && \
    rm -rf /var/lib/apt/lists/* && \
    groupadd -g ${GID} ${APP_USER} && \
    useradd -u ${UID} -g ${GID} -m -s /usr/sbin/nologin ${APP_USER}

EXPOSE 8080

USER ${UID}:${GID}

ENV USERNAME=root \
    PASSWORD=mysql \
    HOSTNAME=mysql-0 \
    PORT=3306 \
    DATABASE=dictionary-api

ENTRYPOINT [ "./backend.dictionary-api" ]