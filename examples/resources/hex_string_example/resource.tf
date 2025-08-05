# Copyright (c) AnticliMaxtic
# SPDX-License-Identifier: MPL-2.0

locals {
  example_string = "this is an example string"
}
resource "hex_string" "example" {
  data = local.example_string
}