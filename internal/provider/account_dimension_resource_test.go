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
					resource.TestCheckResourceAttr("udoma_account_dimension.test", "ref_type", "owner"),
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
				resource udoma_bank_account "test" {
					account_holder = "updated_test_account"
					iban 					 = "DE29100100100987654321"
					bic 					 = "TESTDETTXXX"
					bank_name 		 = "updated_test_bank"
					description 	 = "updated_test_description"
				}
				`,

				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_bank_account.test", "account_holder", "updated_test_account"),
					resource.TestCheckResourceAttr("udoma_bank_account.test", "bank_name", "updated_test_bank"),
					resource.TestCheckResourceAttr("udoma_bank_account.test", "description", "updated_test_description"),
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
  ref_type        = "owner"
  required        = true
  pad_to_size     = 10
}
`
}
