FROM golang:1.19-alpine as builder

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /go/src/app
COPY . .

RUN apk update && apk add git

RUN go get
RUN go build -o /app

FROM golang:1.19-alpine as dev
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/app
COPY . .

RUN apk update && apk add git bash

RUN go get

EXPOSE 8080

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml", "server"]