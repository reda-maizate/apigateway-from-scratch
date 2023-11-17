package postgres

import (
	"api-gateway/internal/adapters/repository/postgres/gen"
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

type APIGatewayRepository struct {
	db  *pgx.Conn
	ctx context.Context
}

func (p *APIGatewayRepository) Login(email, password string) (string, error) {
	queries := gen.New(p.db)

	res, err := queries.GetUser(p.ctx, email)

	if err != nil {
		return "", err
	} else {
		if res.Password != password {
			return "", nil
		}
	}

	log.Println(res)
	return res.Email, nil
}

func (p *APIGatewayRepository) SignUp(email, password string) error {
	log.Println("SignUp", email, password)
	queries := gen.New(p.db)
	log.Println("SignUp 2", queries, p.db, p.ctx)

	_, err := queries.GetUser(p.ctx, email)

	if err != nil {
		log.Println("SignUp 2.1", err)
		return err
	}

	log.Println("SignUp 3", email, password)

	authToken := "1234567890"
	params := gen.CreateUserParams{
		Email:     email,
		Password:  password,
		AuthToken: authToken,
	}

	newUserCreated, err := queries.CreateUser(p.ctx, params)

	if err != nil {
		return err
	}

	log.Println("New user created :", newUserCreated)

	return nil
}
