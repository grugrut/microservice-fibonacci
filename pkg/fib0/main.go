package main

import (
	"context"
	"log"
	"net"

	pb "github.com/grugrut/microservice-fibonacci/fib"
	"google.golang.org/grpc"
)

const (
	port = ":50050"
)

type server struct{}

func (s *server) Calc(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %v", req.In)
	return &pb.Response{Out: 0}, nil
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
