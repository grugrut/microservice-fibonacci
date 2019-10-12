FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GO111MODULE=on
WORKDIR /go/src/github.com/grugrut/microservice-fibonacci
COPY . .
RUN go build -o server pkg/server/main.go

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/grugrut/microservice-fibonacci/server /server
EXPOSE 8080
CMD ["/server"]
