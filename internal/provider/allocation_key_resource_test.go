package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAllocationKeyResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAllocationKeyConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_allocation_key.test", "name", "By Area"),
					resource.TestCheckResourceAttr("udoma_allocation_key.test", "description", "Distribute by area"),
					resource.TestCheckResourceAttr("udoma_allocation_key.test", "formula", "unit_area / total_area"),
				),
			},
			{
				ResourceName:      "udoma_allocation_key.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: `
			resource "udoma_allocation_key" "test" {
				name        = "By Occupants"
				description = "Distribute by occupants"
				formula     = "unit_occupants / total_occupants"
			}
			`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_allocation_key.test", "name", "By Occupants"),
					resource.TestCheckResourceAttr("udoma_allocation_key.test", "description", "Distribute by occupants"),
					resource.TestCheckResourceAttr("udoma_allocation_key.test", "formula", "unit_occupants / total_occupants"),
				),
			},
		},
	})
}

func testAccAllocationKeyConfig() string {
	return `
resource "udoma_allocation_key" "test" {
  name        = "By Area"
  description = "Distribute by area"
  formula     = "unit_area / total_area"
}
`
}
