package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/grugrut/microservice-fibonacci/fib"
	"google.golang.org/grpc"
)

const (
	fib0 = "fib0:50050"
	fib1 = "fib1:50050"
	fibn = "fibn:50050"
	port = ":50050"
)

type server struct{}

func (s *server) Calc(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %v", req.In)
	var address string

	if req.In == 1 || req.In == 0 {
		if req.In == 1 {
			address = fib1
		} else {
			address = fib0
		}
		conn, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewFibClient(conn)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.Calc(ctx, &pb.Request{In: 1})
		if err != nil {
			log.Fatalf("could not connect %v", err)
		}
		return &pb.Response{Out: r.Out}, nil
	}

	n1 := req.In - 1
	n2 := req.In - 2

	switch n1 {
	case 0:
		address = fib0
	case 1:
		address = fib1
	default:
		address = fibn
	}

	conn1, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn1.Close()
	c1 := pb.NewFibClient(conn1)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r1, err := c1.Calc(ctx, &pb.Request{In: n1})
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	cal1 := r1.Out

	switch n2 {
	case 0:
		address = fib0
	case 1:
		address = fib1
	default:
		address = fibn
	}

	conn2, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn2.Close()
	c2 := pb.NewFibClient(conn1)

	ctx, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()
	r2, err := c2.Calc(ctx, &pb.Request{In: n1})
	if err != nil {
		log.Fatalf("could not connect %v", err)
	}
	cal2 := r2.Out

	return &pb.Response{Out: (cal1 + cal2)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFibServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
