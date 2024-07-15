resource "gitops_instance" "test1" {
  instance_name = "terraform provisioned test1"
  orderer_id    = "your.email@address"
  bits_account  = 12341
  service_id    = 4322
  replica_count = 3
  version       = "3.2.*"
  some_value    = "test instance 1"
}