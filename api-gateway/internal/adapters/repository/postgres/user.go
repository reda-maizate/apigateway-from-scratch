package postgres

import (
	gen "api-gateway/internal/adapters/repository/postgres/gen"
	"api-gateway/internal/core/domain"
	"api-gateway/internal/core/ports"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *APIGatewayRepository) Login(loginParams ports.UserParams) (ports.UserResponse, error) {
	queries := gen.New(r.db)

	res, err := queries.GetUser(r.ctx, loginParams.Email)

	if err != nil {
		return ports.UserResponse{}, status.Errorf(codes.NotFound, "User not found")
	} else {
		if res.Password != loginParams.Password {
			return ports.UserResponse{}, status.Errorf(codes.Unauthenticated, "Wrong password")
		}
	}

	return ports.UserResponse{
		Token: res.AuthToken,
	}, nil
}

func (r *APIGatewayRepository) SignUp(signupParams ports.UserParams) (ports.UserResponse, error) {
	queries := gen.New(r.db)

	_, err := queries.GetUser(r.ctx, signupParams.Email)

	if err == nil {
		return ports.UserResponse{}, status.Errorf(codes.AlreadyExists, "User already exists")
	}

	user_uuid := uuid.New().String()
	authToken := uuid.New().String()

	params := gen.CreateUserParams{
		Uuid:      user_uuid,
		Email:     signupParams.Email,
		Password:  signupParams.Password,
		AuthToken: authToken,
	}

	_, err = queries.CreateUser(r.ctx, params)

	if err != nil {
		return ports.UserResponse{}, status.Errorf(codes.Internal, "Error creating user")
	}

	return ports.UserResponse{
		Token: authToken,
	}, nil
}

func (r *APIGatewayRepository) UserFromToken(userFromTokenParams ports.UserFromTokenParams) (ports.UserFromTokenResponse, error) {
	queries := gen.New(r.db)

	user, err := queries.GetUserFromAuthToken(r.ctx, userFromTokenParams.Token)

	if err != nil {
		return ports.UserFromTokenResponse{}, status.Errorf(codes.NotFound, "User not found")
	}

	return ports.UserFromTokenResponse{
		User: &domain.User{
			Uuid:  user.Uuid,
			Email: user.Email,
		},
	}, nil
}
