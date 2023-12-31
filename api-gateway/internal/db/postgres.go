package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
)

type DBConfig struct {
	Host     string
	Username string
	Password string
	Dbname   string
	Port     string
}

type DB struct {
	Db      *pgx.Conn
	Ctx     context.Context
	Queries *Queries
}

func NewDB(config *DBConfig) *DB {
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Host,
		config.Port,
		config.Username,
		config.Dbname,
		config.Password,
	)

	ctx := context.Background()
	db, err := pgx.Connect(ctx, conn)

	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}

	return &DB{
		Db:      db,
		Ctx:     ctx,
		Queries: New(db),
	}
}
