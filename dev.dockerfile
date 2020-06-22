FROM golang:1.14.4-buster
RUN apt-get install git

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go get github.com/cespare/reflex
RUN go get -v github.com/rubenv/sql-migrate/...