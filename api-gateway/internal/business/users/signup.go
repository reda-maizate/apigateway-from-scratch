package users

import (
	database "api-gateway/internal/db"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SignupParams struct {
	Email    string
	Password string
}

type SignupResponse struct {
	Token string
}

func (b *UsersBusiness) SignUp(ctx context.Context, params SignupParams) (SignupResponse, error) {
	_, err := b.queries.GetUser(ctx, params.Email)

	if err == nil {
		return SignupResponse{}, status.Errorf(codes.AlreadyExists, "User already exists")
	}

	userUuid := uuid.New().String()
	authToken := uuid.New().String() // Uuid used as auth token here for simplicity

	_, err = b.queries.CreateUser(ctx, database.CreateUserParams{
		Uuid:      userUuid,
		Email:     params.Email,
		Password:  params.Password,
		AuthToken: authToken,
	})

	if err != nil {
		return SignupResponse{}, status.Errorf(codes.Internal, "Internal error")
	}

	return SignupResponse{
		Token: authToken,
	}, nil
}
