package postgres

import (
	gen "api-gateway/internal/adapters/repository/postgres/gen"
	"api-gateway/internal/core/domain"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *APIGatewayRepository) Login(email, password string) (string, error) {
	queries := gen.New(p.db)

	res, err := queries.GetUser(p.ctx, email)

	if err != nil {
		return "", status.Errorf(codes.InvalidArgument, "Invalid email or Password")
	} else {
		if res.Password != password {
			return "", status.Errorf(codes.InvalidArgument, "Invalid email or Password")
		}
	}

	return res.AuthToken, nil
}

func (p *APIGatewayRepository) SignUp(email, password string) (string, error) {
	queries := gen.New(p.db)

	_, err := queries.GetUser(p.ctx, email)

	if err == nil {
		return "", status.Errorf(codes.AlreadyExists, "Email already exists")
	}

	user_uuid := uuid.New().String()
	authToken := uuid.New().String()

	params := gen.CreateUserParams{
		Uuid:      user_uuid,
		Email:     email,
		Password:  password,
		AuthToken: authToken,
	}

	_, err = queries.CreateUser(p.ctx, params)

	if err != nil {
		return "", status.Errorf(codes.Internal, "Internal error while creating user")
	}

	return authToken, nil
}

func (p *APIGatewayRepository) UserFromToken(token string) (*domain.User, error) {
	queries := gen.New(p.db)

	user, err := queries.GetUserFromAuthToken(p.ctx, token)

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	return &domain.User{
		Uuid:      user.Uuid,
		Email:     user.Email,
		Password:  user.Password,
		AuthToken: user.AuthToken,
	}, nil
}
