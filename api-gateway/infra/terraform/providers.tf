terraform {
  required_providers {
    scaleway = {
      source = "scaleway/scaleway"
    }
    null = {
      source = "hashicorp/null"
    }
  }
}

provider "scaleway" {
  access_key = var.scaleway_access_key
  secret_key = var.scaleway_secret_key
  project_id = var.scaleway_project_id
  region     = var.scaleway_region
}