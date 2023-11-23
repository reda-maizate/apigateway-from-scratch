package postgres

import (
	gen "api-gateway/internal/adapters/repository/postgres/gen"
	"api-gateway/internal/core/ports"
	"slices"
)

func (r *APIGatewayRepository) CheckPermission(checkPermissionParams ports.CheckPermissionParams) (ports.CheckPermissionResponse, error) {
	queries := gen.New(r.db)

	params := gen.GetUserPermissionsParams{
		UserUuid: checkPermissionParams.UserUuid,
		Service:  checkPermissionParams.Service,
		Resource: checkPermissionParams.Resource,
	}

	authorizedActions, err := queries.GetUserPermissions(r.ctx, params)

	if err != nil {
		return ports.CheckPermissionResponse{Authorized: false}, err
	}

	if slices.Contains(authorizedActions, checkPermissionParams.Action) {
		return ports.CheckPermissionResponse{Authorized: true}, nil
	}

	return ports.CheckPermissionResponse{Authorized: false}, nil
}
