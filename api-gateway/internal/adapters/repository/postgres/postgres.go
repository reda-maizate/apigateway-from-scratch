package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

var (
	User     = os.Getenv("POSTGRES_USER")
	Password = os.Getenv("POSTGRES_PASSWORD")
	Dbname   = os.Getenv("POSTGRES_DB")
)

type APIGatewayRepository struct {
	db  *pgx.Conn
	ctx context.Context
}

func NewAPIGatewayRepository() *APIGatewayRepository {
	host := "host.docker.internal"
	port := "5432"
	user := User
	password := Password
	dbname := Dbname

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	log.Println("Trying to connect to database : ", conn, " ...")
	ctx := context.Background()
	db, err := pgx.Connect(ctx, conn)

	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}

	return &APIGatewayRepository{
		db:  db,
		ctx: ctx,
	}
}
