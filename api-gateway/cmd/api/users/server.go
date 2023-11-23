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
	userSvc services.UserService
	_go.UnimplementedUserServer
}

func NewUserServiceServer(userSvc *services.UserService) _go.UserServer {
	return &UserServiceServer{userSvc: *userSvc}
}

func (uss *UserServiceServer) SignUp(ctx context.Context, req *_go.SignUpRequest) (*_go.UserResponse, error) {
	token, err := uss.userSvc.SignUp(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &_go.UserResponse{Token: token}, nil
}

func (uss *UserServiceServer) Login(ctx context.Context, req *_go.LoginRequest) (*_go.UserResponse, error) {
	token, err := uss.userSvc.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &_go.UserResponse{Token: token}, nil
}

func (uss *UserServiceServer) UserFromToken(ctx context.Context, req *_go.MeUserRequest) (*_go.MeUserResponse, error) {
	//log.Println("UserFromToken", req.Token)
	user, err := uss.userSvc.UserFromToken(req.Token)
	if err != nil {
		return nil, err
	}

	return &_go.MeUserResponse{Id: user.Uuid}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50052")

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	store := repository.NewAPIGatewayRepository()
	userSvc := services.NewUserService(store)

	server := NewUserServiceServer(userSvc)

	grpcServer := grpc.NewServer()

	_go.RegisterUserServer(grpcServer, server)

	log.Println("Serving User-service in gRPC Server on :50052")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
