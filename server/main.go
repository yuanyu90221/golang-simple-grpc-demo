package main

import (
	"context"
	"log"
	"net"

	"github.com/yuanyu90221/golang-simple-grpc-demo/server/pb/test"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type server struct {
	test.UnimplementedAddServiceServer
}

func main() {
	// create grpc server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("grpc server is running")
	test.RegisterAddServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
func (s *server) Add(ctx context.Context, in *test.TestRequest) (*test.TestResponse, error) {
	log.Printf("receive: first=%v, second=%v", in.First, in.Second)
	sum := in.First + in.Second
	return &test.TestResponse{Sum: sum}, nil
}
