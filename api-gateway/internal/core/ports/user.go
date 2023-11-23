package ports

import "api-gateway/internal/core/domain"

type UserParams struct {
	Email    string
	Password string
}

type UserResponse struct {
	Token string
}

type UserFromTokenParams struct {
	Token string
}

type UserFromTokenResponse struct {
	User *domain.User
}

type UserService interface {
	Login(UserParams) (UserResponse, error)
	SignUp(UserParams) (UserResponse, error)
	UserFromToken(UserFromTokenParams) (UserFromTokenResponse, error)
}

type UserRepository interface {
	Login(UserParams) (UserResponse, error)
	SignUp(UserParams) (UserResponse, error)
	UserFromToken(UserFromTokenParams) (UserFromTokenResponse, error)
}
