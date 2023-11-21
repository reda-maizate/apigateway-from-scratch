package services

import "api-gateway/internal/core/ports"

type PermissionService struct {
	repo ports.PermissionRepository
}

func NewPermissionService(repo ports.PermissionRepository) *PermissionService {
	return &PermissionService{
		repo: repo,
	}
}

func (s *PermissionService) CheckPermission(UserUuid, Service, Resource string) (bool, error) {
	return s.repo.CheckPermission(UserUuid, Service, Resource)
}
