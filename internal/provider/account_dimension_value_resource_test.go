package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccAccountDimensionValueResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: resourceDefinitionAccountDimensionValue(),
				Check: resource.ComposeAggregateTestCheckFunc(

					// Verify base addtributes
					resource.TestCheckResourceAttr("udoma_account_dimension_value.test", "value", "32"),
					resource.TestCheckResourceAttr("udoma_account_dimension_value.test", "id", "test"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "udoma_account_dimension_value.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: `
				resource udoma_account_dimension_value "test" {
					value = 33
					
				}
				`,

				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_account_dimension_value.test", "value", "33"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}

func resourceDefinitionAccountDimensionValue() string {
	return testAccAccountDimensionConfig() + `
	resource udoma_account_dimension_value "test" {	
		dimension_ref = udoma_account_dimension.test.id
		id = "test"
		value = 32
	}
	`
}
