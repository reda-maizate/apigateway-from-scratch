package main

import (
	_go "api-gateway/api/v1/gen/go"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("host.docker.internal:50052", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := _go.NewUserClient(conn)

	email := "test@test.com"
	password := "testset"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.SignUp(ctx, &_go.SignUpRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.Fatalf("could not signup: %v", err)
	}
	log.Printf("New User Signed up with email: %s", email)

	r_login, err := c.Login(ctx, &_go.LoginRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.Fatalf("could not login: %v", err)
	}
	log.Printf("New User Logged in: %s / auth token: %s", email, r_login.GetToken())
}
