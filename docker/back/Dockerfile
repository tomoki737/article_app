FROM golang:1.20.2-alpine

RUN apk update && apk add git

WORKDIR /src/back

COPY ../../go.mod ./
COPY ../../go.sum ./

RUN go mod download