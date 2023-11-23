package main

import (
	pb "api-gateway/api/v1/gen/go"
	repository "api-gateway/internal/adapters/repository/postgres"
	"api-gateway/internal/core/services"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type UserServiceServer struct {
	userSvc services.UserService
	pb.UnimplementedUserServer
}

func NewUserServiceServer(userSvc *services.UserService) pb.UserServer {
	return &UserServiceServer{userSvc: *userSvc}
}

func (uss *UserServiceServer) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.UserResponse, error) {
	token, err := uss.userSvc.SignUp(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{Token: token}, nil
}

func (uss *UserServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserResponse, error) {
	token, err := uss.userSvc.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.UserResponse{Token: token}, nil
}

func (uss *UserServiceServer) UserFromToken(ctx context.Context, req *pb.MeUserRequest) (*pb.MeUserResponse, error) {
	//log.Println("UserFromToken", req.Token)
	user, err := uss.userSvc.UserFromToken(req.Token)
	if err != nil {
		return nil, err
	}

	return &pb.MeUserResponse{Id: user.Uuid}, nil
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

	pb.RegisterUserServer(grpcServer, server)

	log.Println("Serving User-service in gRPC Server on :50052")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
