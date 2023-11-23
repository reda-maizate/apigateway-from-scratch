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

type PermissionServiceServer struct {
	permissionService services.PermissionService
	pb.UnimplementedPermissionServer
}

func NewPermissionServiceServer(permissionService *services.PermissionService) pb.PermissionServer {
	return &PermissionServiceServer{permissionService: *permissionService}
}

func (s *PermissionServiceServer) CheckPermission(ctx context.Context, req *pb.CheckPermissionRequest) (*pb.CheckPermissionResponse, error) {
	checkPermissionParams := ports.CheckPermissionParams{
		UserUuid: req.UserUuid,
		Service:  req.Service,
		Resource: req.Resource,
		Action:   req.Action,
	}

	HasPermission, err := s.permissionService.CheckPermission(checkPermissionParams)
	if err != nil {
		return nil, err
	}

	return &pb.CheckPermissionResponse{
		HasPermission: HasPermission.Authorized,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50054")

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	store := repository.NewAPIGatewayRepository()
	permissionService := services.NewPermissionService(store)

	server := NewPermissionServiceServer(permissionService)

	grpcServer := grpc.NewServer()

	pb.RegisterPermissionServer(grpcServer, server)

	log.Println("Serving Permissions-service in gRPC Server on :50054")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
