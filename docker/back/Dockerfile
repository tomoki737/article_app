FROM golang:1.20.2-alpine

RUN apk update && apk add git

WORKDIR /src/back

RUN go install github.com/cosmtrek/air@v1.29.0
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY ../../go.mod ./
COPY ../../go.sum ./

RUN go mod download

# CMD [ "air" ]