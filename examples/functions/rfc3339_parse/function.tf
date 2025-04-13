terraform {
  required_providers {
    algorithm = {
      source = "hashicorp.com/marcjnek/algorithm"
    }
  }
}

provider "algorithm" {}

output "timestamp" {
  value = provider::algorithm::rfc3339_parse("2023-07-25T23:43:16Z")
}
