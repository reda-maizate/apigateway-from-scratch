package main

import (
	userstubs "api-gateway/stubs/go/apigateway-from-scratch/users/v1"
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	notestubs "api-gateway/stubs/go/apigateway-from-scratch/notes/v1"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err_user := userstubs.RegisterUserHandlerFromEndpoint(ctx, mux, "users_service:50052", opts)
	err_note := notestubs.RegisterNoteHandlerFromEndpoint(ctx, mux, "notes_service:50053", opts)

	if err_user != nil {
		log.Fatalf("failed to start HTTP user server: %v", err_user)
	} else if err_note != nil {
		log.Fatalf("failed to start HTTP note server: %v", err_note)
	}

	http.ListenAndServe(":8080", mux)
}
