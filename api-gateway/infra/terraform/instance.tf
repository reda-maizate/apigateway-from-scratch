resource "scaleway_instance_ip" "public_ip" {}

resource "scaleway_instance_server" "my-instance" {
  depends_on = [scaleway_instance_ip.public_ip]

  name              = "apigateway-from-scratch-instance"
  type              = "DEV1-L"
  image             = "ubuntu_focal"
  ip_id             = scaleway_instance_ip.public_ip.id
  security_group_id = scaleway_instance_security_group.my-security-group.id
}