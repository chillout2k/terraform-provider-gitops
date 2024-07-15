# terraform-provider-gitops
```
terraform {
  required_providers {
    gitops = {
      source = "chillout2k/gitops"
      version = "0.0.2"
    }
  }
}

/*variable "gitops_username" {
  description = "Gitops provider username"
  type        = string
  sensitive   = false
}

variable "gitops_password" {
  description = "Gitops provider password"
  type        = string
  sensitive   = true
}*/

provider "gitops" {
  gitops_api_uri = "http://localhost:8000"
  cache_path = "/tmp/.gitops-tf-provider"
  grant_type = "auth_code"
  #username = var.gitops_username
  #password = var.gitops_password
  client_id = "gitops-playground"
  client_secret = ""
  token_uri = "https://idp.example.com/protocol/openid-connect/token"
  auth_uri = "https://idp.example.com.de/protocol/openid-connect/auth"
  jwks_uri = "https://idp.example.com./protocol/openid-connect/certs"
  redirect_uri = "http://localhost:12345/authz"
  authz_listener_socket = "localhost:12345"
  scopes = "openid email"
  #debug = true
}

resource "gitops_instance" "test1" {
  instance_name = "terraform provisioned test1"
  orderer_id    = "deine.email@adresse"
  bits_account  = 12341
  service_id    = 4322
  replica_count = 3
  version       = "3.2.*"
  some_value    = "test instance 1"
}

output "instance1" {
  value = gitops_instance.test1
}
```