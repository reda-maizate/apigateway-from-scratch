package services

import "api-gateway/internal/core/ports"

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Login(email, password string) (string, error) {
	return s.repo.Login(email, password)
}

func (s *UserService) SignUp(email, password string) error {
	return s.repo.SignUp(email, password)
}
