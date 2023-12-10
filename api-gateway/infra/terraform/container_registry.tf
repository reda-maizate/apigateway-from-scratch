## Users Service
resource "scaleway_registry_namespace" "users-service-registry-namespace" {
  name        = "users-service-registry-namespace"
  description = "Contains all images for the users service"
}

## Notes Service
resource "scaleway_registry_namespace" "notes-service-registry-namespace" {
  name        = "notes-service-registry-namespace"
  description = "Contains all images for the notes service"
}

## Permissions Service
resource "scaleway_registry_namespace" "permissions-service-registry-namespace" {
  name        = "permissions-service-registry-namespace"
  description = "Contains all images for the permissions service"
}

## Gateway Service
resource "scaleway_registry_namespace" "gateway-service-registry-namespace" {
  name        = "gateway-service-registry-namespace"
  description = "Contains all images for the gateway entrypoint"
}
