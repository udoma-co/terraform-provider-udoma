package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCustomIDGeneratorResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: `
resource udoma_custom_id_generator "test" {
	name 							= "orders"
	generation_script = "Number(lastID) + 1"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify base addtributes
					resource.TestCheckResourceAttr("udoma_custom_id_generator.test", "name", "orders"),
					resource.TestCheckResourceAttr("udoma_custom_id_generator.test", "generation_script", "Number(lastID) + 1"),

					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("udoma_custom_id_generator.test", "id"),
					resource.TestCheckResourceAttrSet("udoma_custom_id_generator.test", "last_updated"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "udoma_custom_id_generator.test",
				ImportState:       true,
				ImportStateVerify: true,
				// The last_updated attribute does not exist in the Udoma
				// API, therefore there is no value for it during import.
				ImportStateVerifyIgnore: []string{"last_updated"},
			},
			// Update and Read testing
			{
				Config: `
resource "udoma_custom_id_generator" "test" {
	name 							= "order_ids"
	generation_script = "Number(lastID) + 5"
}
`,

				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_custom_id_generator.test", "name", "order_ids"),
					resource.TestCheckResourceAttr("udoma_custom_id_generator.test", "generation_script", "Number(lastID) + 5"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}
