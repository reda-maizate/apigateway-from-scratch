resource "cloudflare_record" "terraform" {
  zone_id = var.cloudflare_zone_id
  name    = "api"
  value   = scaleway_instance_server.my-instance.public_ip
  type    = "A"
}