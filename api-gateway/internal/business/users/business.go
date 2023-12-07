package users

import (
	core_business "api-gateway/internal/business/core"
	database "api-gateway/internal/db"
	"github.com/jackc/pgx/v5"
)

type UsersBusiness struct {
	db      *pgx.Conn
	queries *database.Queries
}

type UsersBusinessConfig struct{}

func NewUsersBusiness(coreConfig *core_business.CoreBusinessConfig, config *UsersBusinessConfig) *UsersBusiness {
	return &UsersBusiness{
		db:      coreConfig.DB,
		queries: coreConfig.Queries,
	}
}
