package main

import (
	"api-gateway/internal/business/core"
	"api-gateway/internal/business/permissions"
	repository "api-gateway/internal/db"
	"api-gateway/internal/env"
	rpcPermissions "api-gateway/internal/rpc/permissions/v1"
	permissionstubs "api-gateway/stubs/go/apigateway-from-scratch/permissions/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	ENV_DB_USERNAME = env.Get("DB_USERNAME", "")
	ENV_DB_PASSWORD = env.Get("DB_PASSWORD", "")
	ENV_DB_NAME     = env.Get("DB_NAME", "")
	ENV_DB_PORT     = env.Get("DB_PORT", "")
)

func main() {
	listener, err := net.Listen("tcp", ":50054")

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	dbConfig := &repository.DBConfig{
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

	permissionsBusinessConfig := &permissions.PermissionsBusinessConfig{}
	permissionsBusiness := permissions.NewPermissionsBusiness(coreBusinessConfig, permissionsBusinessConfig)

	permissionService := rpcPermissions.NewPermissionServiceServer(permissionsBusiness)

	grpcServer := grpc.NewServer()

	permissionstubs.RegisterPermissionServer(grpcServer, permissionService)

	log.Println("Serving Permissions-service in gRPC Server on :50054")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
