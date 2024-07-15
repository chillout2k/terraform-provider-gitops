terraform {
  cloud {
    organization = "zwackl"

    workspaces {
      name = "ws-cli-01"
    }
  }
}
