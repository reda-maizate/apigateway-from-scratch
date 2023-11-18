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

	// Notes
	conn, err = grpc.Dial("host.docker.internal:50053", grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	cn := _go.NewNoteClient(conn)

	if err != nil {
		log.Fatalf("could not login: %v", err)
	}

	_, err = cn.CreateNote(ctx, &_go.CreateNoteRequest{
		Note: &_go.NoteMessage{Title: "Test", Content: "Test"}},
	)

	if err != nil {
		log.Fatalf("could not create note: %v", err)
	}

	log.Printf("New Note Created")

	r_notes, err := cn.GetAllNotes(ctx, &_go.GetAllNotesRequest{})
	if err != nil {
		log.Fatalf("could not get all notes: %v", err)
	}

	log.Printf("All Notes: %v", r_notes.GetNotes())
}
