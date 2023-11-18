package postgres

import (
	gen "api-gateway/internal/adapters/repository/postgres/gen"
	"github.com/google/uuid"
	"log"
)

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
	//log.Println("SignUp", email, password)
	queries := gen.New(p.db)
	//log.Println("SignUp 2", queries, p.db, p.ctx)

	_, err := queries.GetUser(p.ctx, email)

	if err != nil {
		//log.Println("SignUp 2.1", err)
		return err
	}

	//log.Println("SignUp 3", email, password)

	user_uuid := uuid.New().String()
	authToken := uuid.New().String()

	params := gen.CreateUserParams{
		Uuid:      user_uuid,
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
