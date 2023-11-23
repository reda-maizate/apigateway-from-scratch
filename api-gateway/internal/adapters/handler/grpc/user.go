package grpc

import (
	"api-gateway/internal/core/ports"
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

func (h *GrpcUserHandler) Login(loginParams ports.UserParams) (ports.UserResponse, error) {
	return h.userService.Login(loginParams)
}

func (h *GrpcUserHandler) SignUp(signupParams ports.UserParams) (ports.UserResponse, error) {
	return h.userService.SignUp(signupParams)
}

func (h *GrpcUserHandler) UserFromToken(userFromTokenParams ports.UserFromTokenParams) (ports.UserFromTokenResponse, error) {
	return h.userService.UserFromToken(userFromTokenParams)
}
