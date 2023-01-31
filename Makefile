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
		