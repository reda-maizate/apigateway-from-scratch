package core

import (
	database "api-gateway/internal/db"
	"context"
	pgx "github.com/jackc/pgx/v5"
)

type CoreBusinessConfig struct {
	DB      *pgx.Conn
	Ctx     context.Context
	Queries *database.Queries
}
