FROM golang:1.23.2-alpine as builder
WORKDIR /code
COPY . .
RUN go build -o pony .

FROM alpine:3.20.3
WORKDIR /home
COPY --from=builder /code/pony /home/pony
CMD [./pony]
