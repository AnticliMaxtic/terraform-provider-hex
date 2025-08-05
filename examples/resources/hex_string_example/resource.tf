locals {
  example_string = "this is an example string"
}
resource "hex_string" "example" {
  data = local.example_string
}