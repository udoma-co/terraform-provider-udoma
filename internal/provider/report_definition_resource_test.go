package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccReportDefinitionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: resourceDefinitionReportDefinition(),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify base addtributes
					resource.TestCheckResourceAttr("udoma_report_definition.test", "name", "basic report"),
					resource.TestCheckResourceAttr("udoma_report_definition.test", "description", "Test description"),
					resource.TestCheckResourceAttr("udoma_report_definition.test", "script", `const ret = { "count": 1 }; return ret;`),

					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("udoma_report_definition.test", "id"),
					resource.TestCheckResourceAttrSet("udoma_report_definition.test", "last_updated"),
					resource.TestCheckResourceAttrSet("udoma_report_definition.test", "created_at"),
					resource.TestCheckResourceAttrSet("udoma_report_definition.test", "updated_at"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "udoma_report_definition.test",
				ImportState:       true,
				ImportStateVerify: true,
				// The last_updated attribute does not exist in the Udoma
				// API, therefore there is no value for it during import.
				ImportStateVerifyIgnore: []string{"last_updated"},
			},
			// Update and Read testing
			{
				Config: `
resource "udoma_report_definition" "test" {
	name 			= "basic updated report"
	description 	= "Test updated description"

	result_schema = {
		result_type 		   = "TABLE"
		table_row_id_attribute = "id"

		attributes = [
			{
				id = "count"
				label = {
					en = "Count"
					de = "Anzahl"
				}
				sequence = 1
			}
		]
	}

	script = "const ret = { \"count\": 2 }; return ret;"
}
`,

				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_report_definition.test", "name", "basic updated report"),
					resource.TestCheckResourceAttr("udoma_report_definition.test", "description", "Test updated description"),
					resource.TestCheckResourceAttr("udoma_report_definition.test", "script", `const ret = { "count": 2 }; return ret;`),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}

func resourceDefinitionReportDefinition() string {
	return `
	resource udoma_report_definition "test" {
		name 			= "basic report"
		description 	= "Test description"
		
		result_schema = {
			result_type 		   = "TABLE"
			table_row_id_attribute = "id"
			
			attributes = [
				{
					id = "count"
					label = {
						en = "Count"
						de = "Anzahl"
					}
					sequence = 1
				}
			]
		}

		script = "const ret = { \"count\": 1 }; return ret;"
	}
	`
}
