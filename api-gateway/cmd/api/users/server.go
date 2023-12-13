package main

import (
	"api-gateway/internal/business/core"
	"api-gateway/internal/business/users"
	repository "api-gateway/internal/db"
	"api-gateway/internal/env"
	rpcUsers "api-gateway/internal/rpc/users/v1"
	userstubs "api-gateway/stubs/go/apigateway-from-scratch/users/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	ENV_DB_USERNAME = env.Get("DB_USERNAME", "")
	ENV_DB_PASSWORD = env.Get("DB_PASSWORD", "")
	ENV_DB_NAME     = env.Get("DB_NAME", "")
	ENV_DB_PORT     = env.Get("DB_PORT", "")
	EVN_DB_HOST     = env.Get("DB_HOST", "")
)

func main() {
	listener, err := net.Listen("tcp", ":50052")

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

	userBusinessConfig := &users.UsersBusinessConfig{}
	userBusiness := users.NewUsersBusiness(coreBusinessConfig, userBusinessConfig)

	userService := rpcUsers.NewUserServiceServer(userBusiness)

	grpcServer := grpc.NewServer()

	userstubs.RegisterUserServer(grpcServer, userService)

	log.Println("Serving User-service in gRPC Server on :50052")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
