package main

import (
	"context"
	"log"
	"time"

	"github.com/yuanyu90221/golang-simple-grpc-demo/client/pb/test"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:8080"
)

func main() {
	// conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := test.NewAddServiceClient(conn)
	TestRPC(c, 5, 4)
}

func TestRPC(c test.AddServiceClient, first, second int64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.Add(ctx, &test.TestRequest{First: first, Second: second})
	if err != nil {
		log.Fatalf("failed to executed Add: %v", err)
	}
	log.Printf("grpc response: %v", res)
}
