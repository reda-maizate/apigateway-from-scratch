buf:
	cd ./api/v1/ && buf generate && cd ./../../;

sqlc:
	cd ./internal/adapters/repository/postgres/ && sqlc generate && cd ./../../../../;

run:
	 docker-compose build --no-cache && docker-compose up