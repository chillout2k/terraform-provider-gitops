provider "gitops" {
  gitops_api_uri        = "http://localhost:8000"
  cache_path            = "/tmp/.gitops-tf-provider"
  grant_type            = "device_code"
  username              = "blah"
  password              = "blubb"
  client_id             = "gitops-playground"
  client_secret         = ""
  token_uri             = "https://idp.example.com/protocol/openid-connect/token"
  auth_uri              = "https://idp.example.com.de/protocol/openid-connect/auth"
  jwks_uri              = "https://idp.example.com./protocol/openid-connect/certs"
  redirect_uri          = "http://localhost:12345/authz"
  authz_listener_socket = "localhost:12345"
  scopes                = "openid email"
  #debug = true
}
