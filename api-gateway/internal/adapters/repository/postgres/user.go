package postgres

import (
	"api-gateway/internal/adapters/repository/postgres/gen"
	"log"
)

func (p *APIGatewayRepository) Login(email, password string) (string, error) {
	queries := gen.New(p.db)

	res, err := queries.GetUser(p.ctx, email)

	if err != nil {
		return "", err
	}

	log.Println(res)
	return res.Email, nil
}

func (p *APIGatewayRepository) SignUp(email, password string) error {
	queries := gen.New(p.db)

	if res, err := queries.GetUser(p.ctx, email); err == nil {
		if res.Email == email {
			log.Println("User", email, "already exists")
			return nil
		}
	}

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
