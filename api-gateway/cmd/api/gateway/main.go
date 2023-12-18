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

	errUser := userstubs.RegisterUserHandlerFromEndpoint(ctx, mux, "users_service:50052", opts)
	errNote := notestubs.RegisterNoteHandlerFromEndpoint(ctx, mux, "notes_service:50053", opts)

	if errUser != nil {
		log.Fatalf("failed to start HTTP user server: %v", errUser)
	} else if errNote != nil {
		log.Fatalf("failed to start HTTP note server: %v", errNote)
	}

	http.ListenAndServe(":80", mux)
}
