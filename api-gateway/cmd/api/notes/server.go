package main

import (
	"api-gateway/internal/business/core"
	"api-gateway/internal/business/notes"
	repository "api-gateway/internal/db"
	"api-gateway/internal/env"
	"api-gateway/internal/middlewares"
	rpcNotes "api-gateway/internal/rpc/notes/v1"
	notestubs "api-gateway/stubs/go/apigateway-from-scratch/notes/v1"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"log"
	"net"
)

//const (
//	Service  = "notes"
//	Resource = "note"
//)

var (
	ENV_DB_USERNAME = env.Get("DB_USERNAME", "")
	ENV_DB_PASSWORD = env.Get("DB_PASSWORD", "")
	ENV_DB_NAME     = env.Get("DB_NAME", "")
	ENV_DB_PORT     = env.Get("DB_PORT", "")
	EVN_DB_HOST     = env.Get("DB_HOST", "")
)

func main() {
	listener, err := net.Listen("tcp", ":50053")

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	dbConfig := &repository.DBConfig{
		Host:     EVN_DB_HOST,
		Username: ENV_DB_USERNAME,
		Password: ENV_DB_PASSWORD,
		Dbname:   ENV_DB_NAME,
		Port:     ENV_DB_PORT,
	}

	store := repository.NewDB(dbConfig)

	coreBusinessConfig := &core.CoreBusinessConfig{
		DB:      store.Db,
		Queries: store.Queries,
		Ctx:     store.Ctx,
	}

	notesBusinessConfig := &notes.NotesBusinessConfig{}
	notesBusiness := notes.NewNotesBusiness(coreBusinessConfig, notesBusinessConfig)

	notesService := rpcNotes.NewNoteServiceServer(notesBusiness)

	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc.UnaryServerInterceptor(
			middlewares.AuthInterceptor,
		),
	)))
	grpcServer := grpc.NewServer(opts...)

	notestubs.RegisterNoteServer(grpcServer, &notesService)

	log.Println("Serving Notes-service in gRPC Server on :50053")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
