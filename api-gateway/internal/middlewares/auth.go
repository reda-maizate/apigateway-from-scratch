package middlewares

import (
	userstubs "api-gateway/stubs/go/apigateway-from-scratch/users/v1"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "Missing context metadata")
	}

	token := md["Authorization"]
	if len(token) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "Missing Authorization token")
	}

	usersClient, err := grpc.Dial("users_service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial Users-service gRPC server: %v", err)
	}
	defer usersClient.Close()

	usersConn := userstubs.NewUserClient(usersClient)
	userUuid, err := usersConn.UserFromToken(ctx, &userstubs.MeUserRequest{Token: token[0]})

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid token")
	}

	ctx = context.WithValue(ctx, "userUuid", userUuid)
	//log.Println("Inserted userUuid in context: ", userUuid)

	return handler(ctx, req)
}
