package main

import (
	_go "api-gateway/api/v1/gen/go"
	repository "api-gateway/internal/adapters/repository/postgres"
	"api-gateway/internal/core/services"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

/*
type UserServer struct {
	_go.UnimplementedUserServer
}
*/

type UserServiceServer struct {
	svc services.UserService
	_go.UnimplementedUserServer
}

func (uss *UserServiceServer) SignUp(ctx context.Context, req *_go.SignUpRequest) (*emptypb.Empty, error) {
	err := uss.svc.SignUp(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (uss *UserServiceServer) Login(ctx context.Context, req *_go.LoginRequest) (*_go.LoginResponse, error) {
	token, err := uss.svc.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &_go.LoginResponse{Token: token}, nil
}

/*
func NewUserServer() *UserServer {
	return &UserServer{}
}
*/

func main() {
	// Add a log to indicate that the server is starting
	listener, err := net.Listen("tcp", ":50052")

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	/*
		s := grpc.NewServer()

		_go.RegisterUserServer(s, &UserServer{})

		log.Println("Serving gRPC Server on 0.0.0.0:8080")
		log.Fatal(s.Serve(lis))
	*/

	store := repository.NewAPIGatewayRepository()
	svc := services.NewUserService(store)

	server := UserServiceServer{svc: *svc}

	grpcServer := grpc.NewServer()

	_go.RegisterUserServer(grpcServer, &server)

	log.Println("Serving gRPC Server on :50052")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
