FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
WORKDIR /go/src/github.com/grugrut/microservice-fibonacci
COPY . .
RUN go build -o fib0 pkg/fib0/main.go

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/grugrut/microservice-fibonacci/fib0 /fib0
EXPOSE 50050
CMD ["/fib0"]
