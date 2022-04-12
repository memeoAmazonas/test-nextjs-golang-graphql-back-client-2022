# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o bin/client cmd/main.go

CMD [ "/app/bin/client" ]
