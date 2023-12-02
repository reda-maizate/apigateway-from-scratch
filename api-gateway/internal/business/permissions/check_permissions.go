package permissions

import (
	database "api-gateway/internal/db"
	"context"
	"slices"
)

type CheckPermissionParams struct {
	UserUuid string
	Service  string
	Resource string
	Action   string
}

type CheckPermissionResponse struct {
	Authorized bool
}

func (b *PermissionsBusiness) CheckPermission(ctx context.Context, params *CheckPermissionParams) (CheckPermissionResponse, error) {

	authorizedActions, err := b.queries.GetUserPermissions(ctx, database.GetUserPermissionsParams{
		UserUuid: params.UserUuid,
		Service:  params.Service,
		Resource: params.Resource,
	})

	if err != nil {
		return CheckPermissionResponse{Authorized: false}, err
	}

	if slices.Contains(authorizedActions, params.Action) {
		return CheckPermissionResponse{Authorized: true}, nil
	}

	return CheckPermissionResponse{Authorized: false}, nil
}
