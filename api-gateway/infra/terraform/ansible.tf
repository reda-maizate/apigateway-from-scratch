resource "null_resource" "ansible" {
  depends_on = [scaleway_instance_server.my-instance]

  provisioner "local-exec" {
    command = "ansible-playbook -u root -i '${scaleway_instance_server.my-instance.public_ip},' --private-key ${var.scaleway_private_key_path} ../ansible/playbook.yml"
  }
}