package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCustomerScriptResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: `
resource udoma_customer_script "test" {
	name 				= "test-script"
	description = "test-script"
	scope       = "GLOBAL"
	script      = "function test() { return \"test\"; }"
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify base addtributes
					resource.TestCheckResourceAttr("udoma_customer_script.test", "name", "test-script"),
					resource.TestCheckResourceAttr("udoma_customer_script.test", "description", "test-script"),
					resource.TestCheckResourceAttr("udoma_customer_script.test", "scope", "GLOBAL"),
					resource.TestCheckResourceAttr("udoma_customer_script.test", "script", "function test() { return \"test\"; }"),

					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("udoma_customer_script.test", "id"),
					resource.TestCheckResourceAttrSet("udoma_customer_script.test", "last_updated"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "udoma_customer_script.test",
				ImportState:       true,
				ImportStateVerify: true,
				// The last_updated attribute does not exist in the Udoma
				// API, therefore there is no value for it during import.
				ImportStateVerifyIgnore: []string{"last_updated"},
			},
			// Update and Read testing
			{
				Config: `
resource "udoma_customer_script" "test" {
	name 				= "updated-script"
	description = "updated-script"
	scope       = "WORKFLOW"
	script      = "function test() { return \"test-workflow\"; }"
}
`,

				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_customer_script.test", "name", "updated-script"),
					resource.TestCheckResourceAttr("udoma_customer_script.test", "description", "updated-script"),
					resource.TestCheckResourceAttr("udoma_customer_script.test", "scope", "WORKFLOW"),
					resource.TestCheckResourceAttr("udoma_customer_script.test", "script", "function test() { return \"test-workflow\"; }"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}
