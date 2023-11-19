package services

import (
	"api-gateway/internal/core/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserService) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Login(email string, password string) (string, error) {
	return s.repo.Login(email, password)
}

func (s *UserService) SignUp(email string, password string) (string, error) {
	return s.repo.SignUp(email, password)
}
