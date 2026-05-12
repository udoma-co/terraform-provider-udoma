package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCostTypeResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCostTypeConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_cost_type.test", "name", "Test Cost Type"),
					resource.TestCheckResourceAttr("udoma_cost_type.test", "category", "maintenance"),
					resource.TestCheckResourceAttr("udoma_cost_type.test", "is_fixed", "true"),
					resource.TestCheckResourceAttr("udoma_cost_type.test", "description", "A test cost type"),
					resource.TestCheckResourceAttr("udoma_cost_type.test", "tenant_billing_rate", "0"),
				),
			},
			{
				ResourceName:            "udoma_cost_type.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"tenant_billing_rate"},
			},
			{
				Config: `
			resource "udoma_cost_type" "test" {
				name        = "Updated Cost Type"
				category    = "utilities"
				description = "An updated cost type"
			}
			`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_cost_type.test", "name", "Updated Cost Type"),
					resource.TestCheckResourceAttr("udoma_cost_type.test", "category", "utilities"),
					resource.TestCheckNoResourceAttr("udoma_cost_type.test", "is_fixed"),
					resource.TestCheckResourceAttr("udoma_cost_type.test", "description", "An updated cost type"),
					resource.TestCheckNoResourceAttr("udoma_cost_type.test", "tenant_billing_rate"),
				),
			},
		},
	})
}

func testAccCostTypeConfig() string {
	return `
resource "udoma_cost_type" "test" {
  name        = "Test Cost Type"
  category    = "maintenance"
  is_fixed    = true
  description = "A test cost type"
  tenant_billing_rate = 0
}
`
}
