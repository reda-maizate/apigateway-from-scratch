buf:
	protoc \
 	--proto_path=. \
 	--go_out=. \
 	--go-grpc_out=. \
 	-I /api/proto/compiled /api/proto/helloworld.proto

run-server:
	go run api/helloworld/server.go