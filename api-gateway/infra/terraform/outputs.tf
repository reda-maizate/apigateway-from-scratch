output "instance_ip" {
  depends_on = [scaleway_instance_server.my-instance]
  value = scaleway_instance_server.my-instance.public_ip
}