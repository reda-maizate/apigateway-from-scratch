package main

import (
	"api-gateway/internal/core/ports"
	"api-gateway/internal/core/services"
	repository "api-gateway/internal/db"
	userstubs "api-gateway/stubs/go/apigateway-from-scratch/users/v1"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserServiceServer struct {
	userService services.UserService
	userstubs.UnimplementedUserServer
}

func NewUserServiceServer(userService *services.UserService) userstubs.UserServer {
	return &UserServiceServer{userService: *userService}
}

func (s *UserServiceServer) SignUp(ctx context.Context, req *userstubs.SignUpRequest) (*userstubs.UserResponse, error) {
	signUpParams := ports.UserParams{
		Email:    req.Email,
		Password: req.Password,
	}

	userResponse, err := s.userService.SignUp(signUpParams)
	if err != nil {
		return nil, err
	}

	return &userstubs.UserResponse{Token: userResponse.Token}, nil
}

func (s *UserServiceServer) Login(ctx context.Context, req *userstubs.LoginRequest) (*userstubs.UserResponse, error) {
	logInParams := ports.UserParams{
		Email:    req.Email,
		Password: req.Password,
	}

	userResponse, err := s.userService.Login(logInParams)
	if err != nil {
		return nil, err
	}

	return &userstubs.UserResponse{Token: userResponse.Token}, nil
}

func (s *UserServiceServer) UserFromToken(ctx context.Context, req *userstubs.MeUserRequest) (*userstubs.MeUserResponse, error) {
	userFromTokenParams := ports.UserFromTokenParams{
		Token: req.Token,
	}
	userFromTokenResponse, err := s.userService.UserFromToken(userFromTokenParams)
	if err != nil {
		return nil, err
	}

	return &userstubs.MeUserResponse{Id: userFromTokenResponse.User.Uuid}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50052")

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	store := repository.NewAPIGatewayRepository()
	userService := services.NewUserService(store)

	server := NewUserServiceServer(userService)

	grpcServer := grpc.NewServer()

	userstubs.RegisterUserServer(grpcServer, server)

	log.Println("Serving User-service in gRPC Server on :50052")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
