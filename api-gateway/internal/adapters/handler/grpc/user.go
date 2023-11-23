package grpc

import (
	"api-gateway/internal/core/domain"
	"api-gateway/internal/core/services"
)

type GrpcUserHandler struct {
	userService services.UserService
}

func NewGrpcUserHandler(userService services.UserService) *GrpcUserHandler {
	return &GrpcUserHandler{
		userService: userService,
	}
}

func (h *GrpcUserHandler) Login(email, password string) (string, error) {
	return h.userService.Login(email, password)
}

func (h *GrpcUserHandler) SignUp(email, password string) (string, error) {
	return h.userService.SignUp(email, password)
}

func (h *GrpcUserHandler) UserFromToken(token string) (*domain.User, error) {
	return h.userService.UserFromToken(token)
}
