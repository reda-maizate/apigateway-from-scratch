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

type PermissionServiceServer struct {
	svc services.PermissionService
	pb.UnimplementedPermissionServer
}

func (pss *PermissionServiceServer) CheckPermission(ctx context.Context, req *pb.CheckPermissionRequest) (*pb.CheckPermissionResponse, error) {
	HasPermission, err := pss.svc.CheckPermission(req.UserUuid, req.Service, req.Resource, req.Action)
	if err != nil {
		return nil, err
	}

	return &pb.CheckPermissionResponse{HasPermission: HasPermission}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50054")

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	store := repository.NewAPIGatewayRepository()
	svc := services.NewPermissionService(store)

	server := PermissionServiceServer{svc: *svc}

	grpcServer := grpc.NewServer()

	pb.RegisterPermissionServer(grpcServer, &server)

	log.Println("Serving Permissions-service in gRPC Server on :50054")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
