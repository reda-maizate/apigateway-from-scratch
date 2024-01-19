resource "scaleway_instance_security_group" "my-security-group" {
  inbound_default_policy  = "drop"
  outbound_default_policy = "accept"
  inbound_rule {
    action = "accept"
    port   = "22" # SSH
  }
  inbound_rule {
    action = "accept"
    port   = "80" # HTTP API Gateway entrypoint
  }
  inbound_rule {
    action = "accept"
	port   = "8080" # API Gateway entrypoint
  }
  inbound_rule {
    action = "accept"
    port   = "443" # HTTPS API Gateway entrypoint
  }
}