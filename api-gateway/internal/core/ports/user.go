package ports

import "api-gateway/internal/core/domain"

type UserService interface {
	Login(email, password string) (string, error)
	SignUp(email, password string) (string, error)
	UserFromToken(token string) (*domain.User, error)
}

type UserRepository interface {
	Login(email, password string) (string, error)
	SignUp(email, password string) (string, error)
	UserFromToken(token string) (*domain.User, error)
}
