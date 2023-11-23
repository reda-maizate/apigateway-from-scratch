package grpc

import "api-gateway/internal/core/services"

type GrpcPermissionHandler struct {
	svc services.PermissionService
}

func NewGrpcPermissionHandler(svc services.PermissionService) *GrpcPermissionHandler {
	return &GrpcPermissionHandler{
		svc: svc,
	}
}

func (gph *GrpcPermissionHandler) CheckPermission(UserUuid, Service, Resource, Action string) (bool, error) {
	return gph.svc.CheckPermission(UserUuid, Service, Resource, Action)
}
