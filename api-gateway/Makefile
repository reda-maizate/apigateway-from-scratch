buf:
	cd ./api/v1/ && buf generate && cd ./../../;

sqlc:
	cd ./internal/adapters/repository/postgres/ && sqlc generate && cd ./../../../../;

run-server:
	go run ./cmd/api/userServer.go

run-client:
	go run ./cmd/api/client/client.go