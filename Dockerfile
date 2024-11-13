FROM golang:1.23.2-alpine

RUN go build -o pony ./
