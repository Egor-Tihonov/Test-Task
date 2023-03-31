FROM golang:alpine AS builder

WORKDIR /cmd

ADD go.mod .

COPY . .

RUN go build -o main cmd/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /cmd/main /cmd/main

CMD [". /cmd"]
