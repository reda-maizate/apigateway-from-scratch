package v1

import (
	business "api-gateway/internal/business/permissions"
	permissionstubs "api-gateway/stubs/go/apigateway-from-scratch/permissions/v1"
	"context"
)

type PermissionServiceServer struct {
	permissionBusiness business.PermissionsBusiness
	permissionstubs.UnimplementedPermissionServer
}

func NewPermissionServiceServer(permissionBusiness *business.PermissionsBusiness) permissionstubs.PermissionServer {
	return &PermissionServiceServer{permissionBusiness: *permissionBusiness}
}

func (s *PermissionServiceServer) CheckPermission(ctx context.Context, req *permissionstubs.CheckPermissionRequest) (*permissionstubs.CheckPermissionResponse, error) {
	checkPermissionParams := &business.CheckPermissionParams{
		UserUuid: req.UserUuid,
		Service:  req.Service,
		Resource: req.Resource,
		Action:   req.Action,
	}

	HasPermission, err := s.permissionBusiness.CheckPermission(ctx, checkPermissionParams)
	if err != nil {
		return nil, err
	}

	return &permissionstubs.CheckPermissionResponse{
		Authorized: HasPermission.Authorized,
	}, nil
}
