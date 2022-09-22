# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /proxy

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /dockerproxy
EXPOSE 8081
CMD ["/dockerproxy"]