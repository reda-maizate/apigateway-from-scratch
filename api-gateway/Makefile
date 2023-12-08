buf:
	cd ./proto/ && buf generate && cd ./../;

sqlc:
	cd ./internal/db/ && sqlc generate && cd ./../../;

run:
	docker-compose build --no-cache && docker-compose up;

tf-init:
	cd ./terraform/ && terraform init && cd ./../;

tf-up:
	cd ./terraform/ && terraform apply -var-file=terraform.tfvars -auto-approve && cd ./../;

tf-down:
	cd ./terraform/ && terraform destroy -var-file=terraform.tfvars -auto-approve && cd ./../;