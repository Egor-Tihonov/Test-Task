FROM golang:alpine AS builder

WORKDIR /cmd

ADD go.mod .

COPY . .

RUN go build -o main main.go

FROM alpine

WORKDIR /cmd

COPY --from=builder /cmd/main /cmd/main

CMD [". /main"]

#FROM golang:latest

#RUN go version
#ENV GOPATH=/

#COPY ./ ./

#RUN go mod download
#RUN go build -o test-task ./cmd/main.go

#CMD ["./test-task"]