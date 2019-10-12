package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	pb "github.com/grugrut/microservice-fibonacci/fib"
	"google.golang.org/grpc"
)

const (
	address = "fibn:50050"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		num, err := strconv.Atoi(r.URL.Query().Get("num"))
		if err != nil {
			num = 0
		}

		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewFibClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		res, err := c.Calc(ctx, &pb.Request{In: int32(num)})
		if err != nil {
			log.Fatalf("could not connect %v", err)
		}
		fmt.Fprintf(w, "%v", res.Out)
	})

	http.ListenAndServe(":8080", nil)
}
