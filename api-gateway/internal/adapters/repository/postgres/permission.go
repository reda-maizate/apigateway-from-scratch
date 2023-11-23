package postgres

import (
	gen "api-gateway/internal/adapters/repository/postgres/gen"
	"slices"
)

func (r *APIGatewayRepository) CheckPermission(UserUuid, Service, Resource, Action string) (bool, error) {
	queries := gen.New(r.db)

	params := gen.GetUserPermissionsParams{
		UserUuid: UserUuid,
		Service:  Service,
		Resource: Resource,
	}

	authorizedActions, err := queries.GetUserPermissions(r.ctx, params)

	if err != nil {
		return false, err
	}

	if slices.Contains(authorizedActions, Action) {
		return true, nil
	}

	return true, nil
}
