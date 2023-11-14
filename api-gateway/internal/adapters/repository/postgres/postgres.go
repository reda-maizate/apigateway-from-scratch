package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type APIGatewayRepository struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewAPIGatewayRepository(db *pgx.Conn) *APIGatewayRepository {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "postgres"
	dbname := "postgres"

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	ctx := context.Background()
	db, err := pgx.Connect(ctx, conn)

	if err != nil {
		panic(err)
	}

	return &APIGatewayRepository{
		db:  db,
		ctx: ctx,
	}
}
