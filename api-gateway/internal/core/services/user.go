package services

import (
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

func (s *UserService) Login(loginParams ports.UserParams) (ports.UserResponse, error) {
	return s.userRepository.Login(loginParams)
}

func (s *UserService) SignUp(signupParams ports.UserParams) (ports.UserResponse, error) {
	return s.userRepository.SignUp(signupParams)
}

func (s *UserService) UserFromToken(userFromTokenParams ports.UserFromTokenParams) (ports.UserFromTokenResponse, error) {
	return s.userRepository.UserFromToken(userFromTokenParams)
}
