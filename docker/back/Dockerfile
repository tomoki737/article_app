FROM golang:1.20.2-alpine

RUN apk update && apk add git

WORKDIR /src/back

RUN go install github.com/cosmtrek/air@v1.29.0

COPY ../../go.mod ./
COPY ../../go.sum ./

RUN go mod download

CMD [ "air" ]