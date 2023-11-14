package main

import (
	_go "api-gateway/api/v1/gen/go"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := _go.NewUserClient(conn)

	mail := "test@test.com"
	password := "testset"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r_sign_up, err := c.SignUp(ctx, &_go.UserRequest{
		Mail:     mail,
		Password: password,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("New User Signed up: %s", r_sign_up.GetMessage())

	r_login, err := c.Login(ctx, &_go.UserRequest{
		Mail:     mail,
		Password: password,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("New User Logged in: %s / auth token: %s", r_login.GetMessage(), r_login.GetAuthToken())
}
