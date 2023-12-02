package users

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LoginParams struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Token string
}

func (b *UsersBusiness) Login(ctx context.Context, params LoginParams) (LoginResponse, error) {
	res, err := b.queries.GetUser(ctx, params.Email)

	if err != nil {
		return LoginResponse{}, status.Errorf(codes.NotFound, "User not found")
	} else {
		if res.Password != params.Password {
			return LoginResponse{}, status.Errorf(codes.Unauthenticated, "Wrong password")
		}
	}

	return LoginResponse{
		Token: res.AuthToken,
	}, nil
}
