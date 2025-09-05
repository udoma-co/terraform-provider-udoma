package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAccountDimensionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccountDimensionConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_account_dimension.test", "name", "Test Account Dimension"),
					resource.TestCheckResourceAttr("udoma_account_dimension.test", "description", "A test account dimension"),
					resource.TestCheckResourceAttr("udoma_account_dimension.test", "ref_type", "none"),
					resource.TestCheckResourceAttr("udoma_account_dimension.test", "required", "true"),
					resource.TestCheckResourceAttr("udoma_account_dimension.test", "pad_to_size", "10"),
				),
			},
			{
				ResourceName:      "udoma_account_dimension.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: `
				resource "udoma_account_dimension" "test" {
					name            = "Updated Account Dimension"
					description     = "An updated account dimension"
					ref_type        = "none"
					required        = true
					pad_to_size     = 8
				}
				`,

				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_account_dimension.test", "name", "Updated Account Dimension"),
					resource.TestCheckResourceAttr("udoma_account_dimension.test", "description", "An updated account dimension"),
					resource.TestCheckResourceAttr("udoma_account_dimension.test", "pad_to_size", "8"),
				),
			},
		},
	})
}

func testAccAccountDimensionConfig() string {
	return `
resource "udoma_account_dimension" "test" {
  name            = "Test Account Dimension"
  description     = "A test account dimension"
  ref_type        = "none"
  required        = true
  pad_to_size     = 10
}
`
}
