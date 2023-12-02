buf:
	cd ./proto/ && buf generate && cd ./../;

sqlc:
	cd ./internal/db/ && sqlc generate && cd ./../../;

run:
	 docker-compose build --no-cache && docker-compose up