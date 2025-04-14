terraform {
  required_providers {
    algorithm = {
      source = "hashicorp.com/marcjanek/algorithm"
    }
  }
}

provider "algorithm" {}

output "timestamp" {
  value = provider: : algorithm : : bin_packing_string(
[
"dom", "ania", "kasia"
], 5)
}
