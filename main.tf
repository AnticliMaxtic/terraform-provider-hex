# Example usage of the hex provider
terraform {
  required_providers {
    hex = {
      source = "AnticliMaxtic/hex"
    }
  }
}

resource "hex_string" "example" {
  data = "hello world"
}

output "hex_result" {
  value = hex_string.example.result
}

output "original_data" {
  value = hex_string.example.data
}
