package main

import (
	_go "api-gateway/api/v1/gen/go"
	repository "api-gateway/internal/adapters/repository/postgres"
	"api-gateway/internal/core/services"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserServiceServer struct {
	svc services.UserService
	_go.UnimplementedUserServer
}

func (uss *UserServiceServer) SignUp(ctx context.Context, req *_go.SignUpRequest) (*_go.UserResponse, error) {
	token, err := uss.svc.SignUp(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &_go.UserResponse{Token: token}, nil
}

func (uss *UserServiceServer) Login(ctx context.Context, req *_go.LoginRequest) (*_go.UserResponse, error) {
	token, err := uss.svc.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &_go.UserResponse{Token: token}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50052")

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	store := repository.NewAPIGatewayRepository()
	svc := services.NewUserService(store)

	server := UserServiceServer{svc: *svc}

	grpcServer := grpc.NewServer()

	_go.RegisterUserServer(grpcServer, &server)

	log.Println("Serving User-service in gRPC Server on :50052")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
