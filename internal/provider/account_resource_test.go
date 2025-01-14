package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccountResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccountConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_account.test", "name", "Test Account"),
					resource.TestCheckResourceAttr("udoma_account.test", "currency", "USD"),
					resource.TestCheckResourceAttr("udoma_account.test", "number", "128"),
					resource.TestCheckResourceAttr("udoma_account.test", "is_balance", "true"),
					resource.TestCheckResourceAttr("udoma_account.test", "type", "asset"),
				),
			},
			{
				ResourceName:      "udoma_account.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: `
				resource udoma_account "test" {
					number   = 501
					name     = "Updated Account"
					currency = "EUR"
					is_balance = true
  					type = "asset"
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_account.test", "name", "Updated Account"),
					resource.TestCheckResourceAttr("udoma_account.test", "currency", "EUR"),
					resource.TestCheckResourceAttr("udoma_account.test", "number", "501"),
				),
			},
		},
	})
}

func testAccAccountConfig() string {
	return testAccAccountDimensionConfig() + `
resource "udoma_account" "test" {
  number   = 128
  name     = "Test Account"
  is_balance = true
  type = "asset"
  currency = "USD"
}
`
}
