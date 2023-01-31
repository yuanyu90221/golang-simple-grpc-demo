# golang with grpc sample

## first create proto file for grpc service

```proto
syntax = "proto3";

option go_package = "pb/test";
// define service name
service AddService {
  rpc Add (testRequest) returns (testResponse) {}
}

message testRequest {
  int64 first = 1;
  int64 second = 2;
}
message testResponse {
  int64 sum = 1;
}
```

## use protobuf cli to generate client/server callback

need to install protobuf compiler and golang-gen-grpc plugin

https://grpc.io/docs/languages/go/quickstart/

```makefile
gen-server:
	protoc --proto_path=proto \
		proto/test.proto \
		--go_out=server \
		--go-grpc_out=server
gen-client:
	protoc --proto_path=proto \
		proto/test.proto \
		--go_out=client \
		--go-grpc_out=client
```

## after gen relative code

for server side 

create service for to service

```go
func (s *server) Add(ctx context.Context, in *test.TestRequest) (*test.TestResponse, error) {
	log.Printf("receive: first=%v, second=%v", in.First, in.Second)
	sum := in.First + in.Second
	return &test.TestResponse{Sum: sum}, nil
}
```

for client create connect function to call server

```go
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
```