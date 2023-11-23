package main

import (
	pb "api-gateway/api/v1/gen/go"
	repository "api-gateway/internal/adapters/repository/postgres"
	"api-gateway/internal/core/ports"
	"api-gateway/internal/core/services"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserServiceServer struct {
	userService services.UserService
	pb.UnimplementedUserServer
}

func NewUserServiceServer(userService *services.UserService) pb.UserServer {
	return &UserServiceServer{userService: *userService}
}

func (s *UserServiceServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.UserResponse, error) {
	signUpParams := ports.UserParams{
		Email:    req.Email,
		Password: req.Password,
	}

	userResponse, err := s.userService.SignUp(signUpParams)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{Token: userResponse.Token}, nil
}

func (s *UserServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserResponse, error) {
	logInParams := ports.UserParams{
		Email:    req.Email,
		Password: req.Password,
	}

	userResponse, err := s.userService.Login(logInParams)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{Token: userResponse.Token}, nil
}

func (s *UserServiceServer) UserFromToken(ctx context.Context, req *pb.MeUserRequest) (*pb.MeUserResponse, error) {
	userFromTokenParams := ports.UserFromTokenParams{
		Token: req.Token,
	}
	userFromTokenResponse, err := s.userService.UserFromToken(userFromTokenParams)
	if err != nil {
		return nil, err
	}

	return &pb.MeUserResponse{Id: userFromTokenResponse.User.Uuid}, nil
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

	pb.RegisterUserServer(grpcServer, server)

	log.Println("Serving User-service in gRPC Server on :50052")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
