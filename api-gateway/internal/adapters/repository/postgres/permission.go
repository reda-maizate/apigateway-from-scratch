package postgres

import (
	gen "api-gateway/internal/adapters/repository/postgres/gen"
	"log"
	"slices"
)

func (p *APIGatewayRepository) CheckPermission(UserUuid, Service, Resource, Action string) (bool, error) {
	queries := gen.New(p.db)

	params := gen.GetUserPermissionsParams{
		UserUuid: UserUuid,
		Service:  Service,
		Resource: Resource,
	}

	log.Println("CheckPermission 3", UserUuid, Service, Resource)
	authorizedActions, err := queries.GetUserPermissions(p.ctx, params)

	log.Println("CheckPermission 4", authorizedActions)
	if err != nil {
		return false, err
	}

	if slices.Contains(authorizedActions, Action) {
		return true, nil
	}

	return true, nil
}
