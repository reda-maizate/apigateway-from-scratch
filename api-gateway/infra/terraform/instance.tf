resource "scaleway_instance_ip" "public_ip" {}

resource "scaleway_instance_server" "my-instance" {
  depends_on = [
  	scaleway_instance_ip.public_ip,
   	null_resource.push_to_users_service_registry,
   	null_resource.push_to_users_service_registry,
   	null_resource.push_to_permissions_service_registry,
   	null_resource.push_to_gateway_service_registry,
   	null_resource.push_to_db_registry
  ]

  name              = "apigateway-from-scratch-instance"
  type              = "DEV1-L"
  image             = "ubuntu_focal"
  ip_id             = scaleway_instance_ip.public_ip.id
  security_group_id = scaleway_instance_security_group.my-security-group.id
}