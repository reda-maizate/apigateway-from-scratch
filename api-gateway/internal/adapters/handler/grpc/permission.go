package grpc

import "api-gateway/internal/core/services"

type GrpcPermissionHandler struct {
	permissionService services.PermissionService
}

func NewGrpcPermissionHandler(permissionService services.PermissionService) *GrpcPermissionHandler {
	return &GrpcPermissionHandler{
		permissionService: permissionService,
	}
}

func (h *GrpcPermissionHandler) CheckPermission(UserUuid, Service, Resource, Action string) (bool, error) {
	return h.permissionService.CheckPermission(UserUuid, Service, Resource, Action)
}
