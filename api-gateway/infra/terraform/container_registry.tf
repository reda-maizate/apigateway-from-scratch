## Users Service
resource "scaleway_registry_namespace" "users-service-registry-namespace" {
  name        = "users-service-registry-namespace"
  description = "Contains all images for the users service"
  is_public   = true
}

resource "null_resource" "push_to_users_service_registry" {
  depends_on = [scaleway_registry_namespace.users-service-registry-namespace]

  triggers = {
    registry_namespace_id = scaleway_registry_namespace.users-service-registry-namespace.id
    content_hash = md5(join(",", [
      join(",", fileset("${path.module}/../../internal", "*")),
      join(",", fileset("${path.module}/../../cmd", "*")),
      join(",", fileset("${path.module}/../../proto", "*")),
      join(",", fileset("${path.module}/../../stubs", "*"))
    ]))
  }

  provisioner "local-exec" {
    command = <<EOF
      docker buildx build --platform linux/amd64 --provenance=false -t rg.fr-par.scw.cloud/${scaleway_registry_namespace.users-service-registry-namespace.name}/users:latest -f ../../Dockerfile-services --target users ../../. --push
    EOF
  }
}


## Notes Service
resource "scaleway_registry_namespace" "notes-service-registry-namespace" {
  name        = "notes-service-registry-namespace"
  description = "Contains all images for the notes service"
  is_public   = true
}

resource "null_resource" "push_to_notes_service_registry" {
  depends_on = [scaleway_registry_namespace.notes-service-registry-namespace]

  triggers = {
    registry_namespace_id = scaleway_registry_namespace.notes-service-registry-namespace.id
    content_hash = md5(join(",", [
      join(",", fileset("${path.module}/../../internal", "*")),
      join(",", fileset("${path.module}/../../cmd", "*")),
      join(",", fileset("${path.module}/../../proto", "*")),
      join(",", fileset("${path.module}/../../stubs", "*"))
    ]))
  }

  provisioner "local-exec" {
    command = <<EOF
      docker buildx build --platform linux/amd64 --provenance=false -t rg.fr-par.scw.cloud/${scaleway_registry_namespace.notes-service-registry-namespace.name}/notes:latest -f ../../Dockerfile-services --target notes ../../. --push
    EOF
  }
}

## Permissions Service
resource "scaleway_registry_namespace" "permissions-service-registry-namespace" {
  name        = "permissions-service-registry-namespace"
  description = "Contains all images for the permissions service"
  is_public   = true
}

resource "null_resource" "push_to_permissions_service_registry" {
  depends_on = [scaleway_registry_namespace.permissions-service-registry-namespace]

  triggers = {
    registry_namespace_id = scaleway_registry_namespace.permissions-service-registry-namespace.id
    content_hash = md5(join(",", [
      join(",", fileset("${path.module}/../../internal", "*")),
      join(",", fileset("${path.module}/../../cmd", "*")),
      join(",", fileset("${path.module}/../../proto", "*")),
      join(",", fileset("${path.module}/../../stubs", "*"))
    ]))
  }

  provisioner "local-exec" {
    command = <<EOF
      docker buildx build --platform linux/amd64 --provenance=false -t rg.fr-par.scw.cloud/${scaleway_registry_namespace.permissions-service-registry-namespace.name}/permissions:latest -f ../../Dockerfile-services --target permissions ../../. --push
    EOF
  }
}

## Gateway Service
resource "scaleway_registry_namespace" "gateway-service-registry-namespace" {
  name        = "gateway-service-registry-namespace"
  description = "Contains all images for the gateway entrypoint"
  is_public   = true
}

resource "null_resource" "push_to_gateway_service_registry" {
  depends_on = [scaleway_registry_namespace.gateway-service-registry-namespace]

  triggers = {
    registry_namespace_id = scaleway_registry_namespace.gateway-service-registry-namespace.id
    content_hash = md5(join(",", [
      join(",", fileset("${path.module}/../../internal", "*")),
      join(",", fileset("${path.module}/../../cmd", "*")),
      join(",", fileset("${path.module}/../../proto", "*")),
      join(",", fileset("${path.module}/../../stubs", "*"))
    ]))
  }

  provisioner "local-exec" {
    command = <<EOF
      docker buildx build --platform linux/amd64 --provenance=false -t rg.fr-par.scw.cloud/${scaleway_registry_namespace.gateway-service-registry-namespace.name}/gateway:latest -f ../../Dockerfile-services --target gateway ../../. --push
    EOF
  }
}

## DB
resource "scaleway_registry_namespace" "db-registry-namespace" {
  name        = "db-registry-namespace"
  description = "Contains all images for the database"
  is_public   = true
}

resource "null_resource" "push_to_db_registry" {
  depends_on = [scaleway_registry_namespace.db-registry-namespace]

  triggers = {
    registry_namespace_id = scaleway_registry_namespace.db-registry-namespace.id
    content_hash = md5(join(",", [
      join(",", fileset("${path.module}/../../internal", "*")),
      join(",", fileset("${path.module}/../../cmd", "*")),
      join(",", fileset("${path.module}/../../proto", "*")),
      join(",", fileset("${path.module}/../../stubs", "*"))
    ]))
  }

  provisioner "local-exec" {
    command = <<EOF
      docker buildx build --platform linux/amd64 --provenance=false -t rg.fr-par.scw.cloud/${scaleway_registry_namespace.db-registry-namespace.name}/db:latest -f ../../Dockerfile-db ../../. --push
    EOF
  }
}