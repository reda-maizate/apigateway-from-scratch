package postgres

import (
	gen "api-gateway/internal/adapters/repository/postgres/gen"
	"log"
)

func (p *APIGatewayRepository) AuthTokenExists(token string) bool {
	queries := gen.New(p.db)

	_, err := queries.GetUserByAuthToken(p.ctx, token)
	log.Println("token", token, err)
	if err != nil {
		return false
	}

	return true
}
