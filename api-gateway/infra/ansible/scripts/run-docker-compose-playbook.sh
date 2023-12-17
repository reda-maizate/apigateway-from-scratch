#!/bin/bash

# Check if the script has been called with an argument
if [ $# -eq 0 ]; then
  echo "No arguments supplied, please use -pkp followed by the private key path."
  exit 1
fi

# Check if the first argument is -pkp
if [ "$1" != "-pkp" ]; then
  echo "Invalid argument. Please use -pkp followed by the private key path."
  exit 1
fi

# Check if a second argument is provided
if [ $# -eq 1 ]; then
  echo "No private key path provided. Please provide the path after -pkp."
  exit 1
fi

# Assign the second argument to a variable
pkp=$2
instance_ip=$(cd ./infra/terraform/ && terraform output -json instance_ip | jq -r . && cd /../../)

#echo "Instance IP: $instance_ip"
#echo "Private key path: $pkp"

echo "Running all playbooks..."

# Run wait until reachable playbook
ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -u root -i $instance_ip, --private-key $pkp infra/ansible/playbooks/wait-until-reachable.yml && \
ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -u root -i $instance_ip, --private-key $pkp infra/ansible/playbooks/run-docker-compose.yml