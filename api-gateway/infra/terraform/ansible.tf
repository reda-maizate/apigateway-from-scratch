resource "null_resource" "ansible_wait_until_reachable" {
  depends_on = [scaleway_instance_server.my-instance]

  triggers = {
    build_number = timestamp()
  }

  provisioner "local-exec" {
    command = "ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -u root -i '${scaleway_instance_server.my-instance.public_ip},' --private-key ${var.scaleway_private_key_path} ../ansible/playbooks/wait-until-reachable.yml"
  }
}

resource "null_resource" "ansible_setup_docker" {
  depends_on = [null_resource.ansible_wait_until_reachable]

  triggers = {
    build_number = timestamp()
  }

  provisioner "local-exec" {
    command = "ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -u root -i '${scaleway_instance_server.my-instance.public_ip},' --private-key ${var.scaleway_private_key_path} ../ansible/playbooks/setup-docker.yml"
  }
}

resource "null_resource" "ansible_setup_docker_compose" {
  depends_on = [
    null_resource.ansible_setup_docker,
    null_resource.push_to_users_service_registry,
    null_resource.push_to_db_registry,
    null_resource.push_to_gateway_service_registry,
    null_resource.push_to_notes_service_registry,
    null_resource.push_to_permissions_service_registry,
  ]

  triggers = {
    build_number = timestamp()
  }

  provisioner "local-exec" {
    command = "ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -u root -i '${scaleway_instance_server.my-instance.public_ip},' --private-key ${var.scaleway_private_key_path} ../ansible/playbooks/run-docker-compose.yml"
  }
}