package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	pb "api-gateway/api/v1/gen/go"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err_user := pb.RegisterUserHandlerFromEndpoint(ctx, mux, "users_service:50052", opts)
	err_note := pb.RegisterNoteHandlerFromEndpoint(ctx, mux, "notes_service:50053", opts)

	if err_user != nil {
		log.Fatalf("failed to start HTTP user server: %v", err_user)
	} else if err_note != nil {
		log.Fatalf("failed to start HTTP note server: %v", err_note)
	}

	http.ListenAndServe(":8080", mux)
}
