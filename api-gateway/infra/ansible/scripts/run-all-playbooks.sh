#!/bin/zsh

# Run wait until reachable playbook
ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -u root -i '${terraform output -json | jq},' --private-key ${var.scaleway_private_key_path} ../ansible/playbooks/wait-until-reachable.yml