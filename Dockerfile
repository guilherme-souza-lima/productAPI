FROM golang:1.18-alpine3.14 AS build
WORKDIR /app

COPY cmd /app/cmd
COPY entitie /app/entitie
COPY infra /app/infra
COPY migration /app/migration
COPY user_case /app/user_case
COPY utils /app/utils
COPY .env /app/.env
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
COPY main.go /app/main.go

RUN go mod download
ENTRYPOINT go run main.go