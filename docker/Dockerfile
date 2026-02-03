FROM golang:1.25.5

ARG GROUP=1000
ARG USER=1000

ENV TZ=Europe/Moscow

WORKDIR /app

COPY board/go.mod board/go.sum ./board/
COPY shared/go.mod shared/go.sum ./shared/

# Скачиваем внешние зависимости
WORKDIR /app/board
RUN go mod download

WORKDIR /app

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN groupadd -g $GROUP user
RUN useradd -u $USER -ms /bin/bash -g user user

USER user
