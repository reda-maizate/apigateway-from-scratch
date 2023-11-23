package grpc

import (
	"api-gateway/internal/core/domain"
	"api-gateway/internal/core/services"
)

type GrpcUserHandler struct {
	svc services.UserService
}

func NewGrpcUserHandler(svc services.UserService) *GrpcUserHandler {
	return &GrpcUserHandler{
		svc: svc,
	}
}

func (guh *GrpcUserHandler) Login(email, password string) (string, error) {
	return guh.svc.Login(email, password)
}

func (guh *GrpcUserHandler) SignUp(email, password string) (string, error) {
	return guh.svc.SignUp(email, password)
}

func (guh *GrpcUserHandler) UserFromToken(token string) (*domain.User, error) {
	return guh.svc.UserFromToken(token)
}
