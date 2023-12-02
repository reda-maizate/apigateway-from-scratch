package permissions

import (
	core_business "api-gateway/internal/business/core"
	database "api-gateway/internal/db"
	"github.com/jackc/pgx/v5"
)

type PermissionsBusiness struct {
	db      *pgx.Conn
	queries *database.Queries
}

type PermissionsBusinessConfig struct{}

func NewPermissionsBusiness(coreConfig *core_business.CoreBusinessConfig, config *PermissionsBusinessConfig) *PermissionsBusiness {
	return &PermissionsBusiness{
		db:      coreConfig.DB,
		queries: coreConfig.Queries,
	}
}
