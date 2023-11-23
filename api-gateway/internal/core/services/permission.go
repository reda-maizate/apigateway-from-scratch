package services

import "api-gateway/internal/core/ports"

type PermissionService struct {
	permissionRepository ports.PermissionRepository
}

func NewPermissionService(repo ports.PermissionRepository) *PermissionService {
	return &PermissionService{
		permissionRepository: repo,
	}
}

func (s *PermissionService) CheckPermission(UserUuid, Service, Resource, Action string) (bool, error) {
	return s.permissionRepository.CheckPermission(UserUuid, Service, Resource, Action)
}
