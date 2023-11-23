package services

import (
	"api-gateway/internal/core/domain"
	"api-gateway/internal/core/ports"
)

type UserService struct {
	userRepository ports.UserRepository
}

func NewUserService(repo ports.UserService) *UserService {
	return &UserService{
		userRepository: repo,
	}
}

func (s *UserService) Login(email string, password string) (string, error) {
	return s.userRepository.Login(email, password)
}

func (s *UserService) SignUp(email string, password string) (string, error) {
	return s.userRepository.SignUp(email, password)
}

func (s *UserService) UserFromToken(token string) (*domain.User, error) {
	return s.userRepository.UserFromToken(token)
}
