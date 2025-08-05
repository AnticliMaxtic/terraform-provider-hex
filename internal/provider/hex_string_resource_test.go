// Copyright (c) 2025 AnticliMaxtic
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccHexStringResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccHexStringResourceConfig("hello"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"hex_string.test",
						tfjsonpath.New("id"),
						knownvalue.StringExact("68656c6c6f"),
					),
					statecheck.ExpectKnownValue(
						"hex_string.test",
						tfjsonpath.New("data"),
						knownvalue.StringExact("hello"),
					),
					statecheck.ExpectKnownValue(
						"hex_string.test",
						tfjsonpath.New("result"),
						knownvalue.StringExact("68656c6c6f"),
					),
				},
			},
			// ImportState testing
			{
				ResourceName:      "hex_string.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: testAccHexStringResourceConfig("world"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"hex_string.test",
						tfjsonpath.New("id"),
						knownvalue.StringExact("776f726c64"),
					),
					statecheck.ExpectKnownValue(
						"hex_string.test",
						tfjsonpath.New("data"),
						knownvalue.StringExact("world"),
					),
					statecheck.ExpectKnownValue(
						"hex_string.test",
						tfjsonpath.New("result"),
						knownvalue.StringExact("776f726c64"),
					),
				},
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccHexStringResourceConfig(data string) string {
	return fmt.Sprintf(`
resource "hex_string" "test" {
  data = %[1]q
}
`, data)
}
