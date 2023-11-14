package main

import (
	"api-gateway/api/v1/gen/go"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	_go.UnimplementedUserServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) SignUp(ctx context.Context, in *_go.UserRequest) (*_go.SignUpReply, error) {
	return &_go.SignUpReply{Message: in.GetMail()}, nil
}

func (s *server) Login(ctx context.Context, in *_go.UserRequest) (*_go.LoginUserReply, error) {
	return &_go.LoginUserReply{Message: "User logged in", AuthToken: "dsfjfkshsf7898987djoo"}, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	//_go.RegisterGreeterServer(s, &server{})
	_go.RegisterUserServer(s, &server{})

	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(lis))
}
