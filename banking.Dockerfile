# syntax=docker/dockerfile:1

FROM golang
ENV GO111MODULE=on
WORKDIR /banking/build

COPY . .

RUN go mod download
RUN go build -o banking-app

