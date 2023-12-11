resource "null_resource" "ansible" {
  depends_on = [scaleway_instance_server.my-instance]

  triggers = {
    build_number = "${timestamp()}"
  }

  provisioner "local-exec" {
    command = "ANSIBLE_HOST_KEY_CHECKING=False ansible-playbook -u root -i '${scaleway_instance_server.my-instance.public_ip},' --private-key ${var.scaleway_private_key_path} ../ansible/playbook.yml"
  }
}