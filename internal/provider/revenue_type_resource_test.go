package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestRevenueTypeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRevenueTypeConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_revenue_type.test", "name", "Test Revenue Type"),
					resource.TestCheckResourceAttr("udoma_revenue_type.test", "scope", "property"),
					resource.TestCheckResourceAttr("udoma_revenue_type.test", "system_default", "NONE"),
					resource.TestCheckResourceAttr("udoma_revenue_type.test", "description", "A test revenue type"),
				),
			},
			{
				ResourceName:      "udoma_revenue_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: `
			resource "udoma_revenue_type" "test" {
				name           = "Updated Revenue Type"
				scope          = "ownership"
				system_default = "MAINTENANCE_FEE"
				description    = "An updated revenue type"
			}
			`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_revenue_type.test", "name", "Updated Revenue Type"),
					resource.TestCheckResourceAttr("udoma_revenue_type.test", "scope", "ownership"),
					resource.TestCheckResourceAttr("udoma_revenue_type.test", "system_default", "MAINTENANCE_FEE"),
					resource.TestCheckResourceAttr("udoma_revenue_type.test", "description", "An updated revenue type"),
				),
			},
		},
	})
}

func testAccRevenueTypeConfig() string {
	return `
resource "udoma_revenue_type" "test" {
  name           = "Test Revenue Type"
  scope          = "property"
  system_default = "NONE"
  description    = "A test revenue type"
}
`
}
