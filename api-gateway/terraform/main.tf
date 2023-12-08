resource "scaleway_instance_ip" "public_ip" {}

resource "scaleway_instance_security_group" "my-security-group" {
  inbound_default_policy  = "drop"
  outbound_default_policy = "accept"
  inbound_rule {
    action = "accept"
    port   = "22" # SSH
  }
  inbound_rule {
    action = "accept"
    port   = "8080" # API Gateway entrypoint
  }
}

resource "scaleway_instance_server" "my-instance" {
  type              = "DEV1-L"
  image             = "ubuntu_focal"
  ip_id             = scaleway_instance_ip.public_ip.id
  security_group_id = scaleway_instance_security_group.my-security-group.id
}