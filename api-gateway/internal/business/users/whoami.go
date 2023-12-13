package users

import (
	database "api-gateway/internal/db"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserFromTokenParams struct {
	Token string
}

func (b *UsersBusiness) UserFromToken(ctx context.Context, params UserFromTokenParams) (database.User, error) {
	user, err := b.queries.GetUserFromAuthToken(ctx, params.Token)

	if err != nil {
		return database.User{}, status.Errorf(codes.NotFound, "User not found")
	}

	return user, nil
}
