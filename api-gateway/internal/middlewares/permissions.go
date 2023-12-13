package middlewares

import (
	permissionstubs "api-gateway/stubs/go/apigateway-from-scratch/permissions/v1"
	userstubs "api-gateway/stubs/go/apigateway-from-scratch/users/v1"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

const (
	Service  = "notes"
	Resource = "note"
)

func CheckPermission(ctx context.Context, Action string) (bool, error) {
	userUuid, ok := ctx.Value("userUuid").(*userstubs.MeUserResponse)

	if !ok {
		return false, status.Errorf(codes.Unauthenticated, "Missing userUuid")
	}

	permissionsClient, err := grpc.Dial("permissions_service:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial Permissions-service gRPC server: %v", err)
	}
	defer permissionsClient.Close()

	permissionsService := permissionstubs.NewPermissionClient(permissionsClient)
	authorized, err := permissionsService.CheckPermission(ctx, &permissionstubs.CheckPermissionRequest{
		UserUuid: userUuid.GetId(),
		Service:  Service,
		Resource: Resource,
		Action:   Action,
	})

	//log.Println("PermissionsInterceptor hasPermission:", hasPermission)
	if err != nil || !authorized.GetAuthorized() {
		log.Println("PermissionsInterceptor hasPermission:", authorized, "/ err:", err)
		return false, status.Errorf(codes.PermissionDenied, "You don't have permission to create note")
	}

	return true, nil
}
