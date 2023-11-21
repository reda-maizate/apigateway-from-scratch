package main

import (
	_go "api-gateway/api/v1/gen/go"
	repository "api-gateway/internal/adapters/repository/postgres"
	"api-gateway/internal/core/services"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type PermissionServiceServer struct {
	svc services.PermissionService
	_go.UnimplementedPermissionServer
}

func (pss *PermissionServiceServer) CheckPermission(ctx context.Context, req *_go.CheckPermissionRequest) (*_go.CheckPermissionResponse, error) {
	HasPermission, err := pss.svc.CheckPermission(req.UserUuid, req.Service, req.Resource)
	if err != nil {
		return nil, err
	}

	return &_go.CheckPermissionResponse{HasPermission: HasPermission}, nil
}

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler, gatewayRepository *repository.APIGatewayRepository) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Missing context metadata")
	}

	token := md["authorization"]
	if len(token) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "Missing authorization token")
	}

	if !gatewayRepository.AuthTokenExists(token[0]) {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid token")
	}

	return handler(ctx, req)
}

func main() {
	listener, err := net.Listen("tcp", ":50054")

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	store := repository.NewAPIGatewayRepository()
	svc := services.NewPermissionService(store)

	server := PermissionServiceServer{svc: *svc}

	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return AuthInterceptor(ctx, req, info, handler, store)
	}))
	grpcServer := grpc.NewServer(opts...)

	_go.RegisterPermissionServer(grpcServer, &server)

	log.Println("Serving Permissions-service in gRPC Server on :50054")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
