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

type PermissionServiceServer struct {
	svc services.PermissionService
	_go.UnimplementedPermissionServer
}

func (pss *PermissionServiceServer) CheckPermission(ctx context.Context, req *_go.CheckPermissionRequest) (*_go.CheckPermissionResponse, error) {
	log.Println("CheckPermission 1:", req)
	log.Println("CheckPermission 2 ctx:", ctx)
	HasPermission, err := pss.svc.CheckPermission(req.UserUuid, req.Service, req.Resource, req.Action)
	if err != nil {
		return nil, err
	}

	return &_go.CheckPermissionResponse{HasPermission: HasPermission}, nil
}

//func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler, gatewayRepository *repository.APIGatewayRepository) (interface{}, error) {
//	md, ok := metadata.FromIncomingContext(ctx)
//	if !ok {
//		return nil, status.Errorf(codes.Unauthenticated, "Missing context metadata")
//	}
//
//	log.Println("ctx 2", ctx)
//	token := md["authorization"]
//	if len(token) == 0 {
//		return nil, status.Errorf(codes.Unauthenticated, "Missing authorization token")
//	}
//
//	usersClient, err := grpc.Dial("users_service:50052", grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("Failed to dial Users-service gRPC server: %v", err)
//	}
//	defer usersClient.Close()
//
//	usersService := _go.NewUserClient(usersClient)
//	userUuid, err := usersService.UserFromToken(ctx, &_go.MeUserRequest{Token: token[0]})
//	if err != nil {
//		return nil, status.Errorf(codes.Unauthenticated, "Invalid token")
//	}
//
//	ctx = context.WithValue(ctx, "userUuid", userUuid)
//
//	return handler(ctx, req)
//}

func main() {
	listener, err := net.Listen("tcp", ":50054")

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	store := repository.NewAPIGatewayRepository()
	svc := services.NewPermissionService(store)

	server := PermissionServiceServer{svc: *svc}

	//var opts []grpc.ServerOption
	//opts = append(opts, grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	//	return AuthInterceptor(ctx, req, info, handler, store)
	//}))
	grpcServer := grpc.NewServer()

	_go.RegisterPermissionServer(grpcServer, &server)

	log.Println("Serving Permissions-service in gRPC Server on :50054")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
