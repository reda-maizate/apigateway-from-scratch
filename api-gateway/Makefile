buf:
	cd ./proto/ && buf generate && cd ./../;

sqlc:
	cd ./internal/db/ && sqlc generate && cd ./../../;

run:
	docker-compose build --no-cache && docker-compose up;

tf-init:
	cd ./infra/terraform/ && terraform init && cd ./../../;

tf-up:
	cd ./infra/terraform/ && terraform apply -var-file=terraform.tfvars -auto-approve && cd ./../../;

tf-down:
	cd ./infra/terraform/ && terraform destroy -var-file=terraform.tfvars -auto-approve && cd ./../../;

ansible-run-all-playbooks:
	cd ./infra/ansible/scripts && ./run-all-playbooks.sh && cd ./../../../;

ansible-run-docker-compose:
	cd ./infra/ansible/ && ansible-playbook -i inventory.ini playbook.yml --tags "docker-compose" && cd ./../../;