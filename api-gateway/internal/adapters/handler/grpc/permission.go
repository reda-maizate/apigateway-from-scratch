package grpc

import (
	"api-gateway/internal/core/ports"
	"api-gateway/internal/core/services"
)

type GrpcPermissionHandler struct {
	permissionService services.PermissionService
}

func NewGrpcPermissionHandler(permissionService services.PermissionService) *GrpcPermissionHandler {
	return &GrpcPermissionHandler{
		permissionService: permissionService,
	}
}

func (h *GrpcPermissionHandler) CheckPermission(checkPermissionParams ports.CheckPermissionParams) (ports.CheckPermissionResponse, error) {
	return h.permissionService.CheckPermission(checkPermissionParams)
}
