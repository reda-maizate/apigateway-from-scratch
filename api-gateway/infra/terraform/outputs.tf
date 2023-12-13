output "instance_ip" {
  value = scaleway_instance_server.my-instance.public_ip
}