package postgres

import gen "api-gateway/internal/adapters/repository/postgres/gen"

func (p *APIGatewayRepository) CheckPermission(UserUuid string, Service string, Resource string) (bool, error) {
	queries := gen.New(p.db)

	params := gen.GetUserPermissionsParams{
		UserUuid: UserUuid,
		Service:  Service,
		Resource: Resource,
	}
	_, err := queries.GetUserPermissions(p.ctx, params)

	if err != nil {
		return false, err
	}

	return true, nil
}
