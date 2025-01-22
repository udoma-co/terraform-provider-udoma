package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCaseReportingEndpointResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: resourceDefinitionCaseTemplate("example template", "A case", 1) + resourceDefinitionFAQ() +
					`
resource udoma_case_reporting_endpoint "test" {
	name 							= "test endpoint"
	active 						= true
	case_categories   = [
		{
			priority = 1
			name = {
				"en" = "A case"
				"de" = "Ein Fall"
			}
			templates = [
				{
					priority = 1
					id = udoma_case_template.test.id
				}
			]
		}
	]
	faqs = [ udoma_faq.test.id ]
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify number of items
					resource.TestCheckResourceAttr("udoma_case_reporting_endpoint.test", "name", "test endpoint"),
					resource.TestCheckResourceAttr("udoma_case_reporting_endpoint.test", "active", "true"),

					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("udoma_case_reporting_endpoint.test", "id"),
					resource.TestCheckResourceAttrSet("udoma_case_reporting_endpoint.test", "last_updated"),
					resource.TestCheckResourceAttrSet("udoma_case_reporting_endpoint.test", "faqs.0"),
					resource.TestCheckResourceAttrSet("udoma_case_reporting_endpoint.test", "url"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "udoma_case_reporting_endpoint.test",
				ImportState:       true,
				ImportStateVerify: true,
				// The last_updated attribute does not exist in the Udoma
				// API, therefore there is no value for it during import.
				ImportStateVerifyIgnore: []string{"last_updated"},
			},
			// Update and Read testing
			{
				Config: resourceDefinitionCaseTemplate("example template", "A case", 1) + resourceDefinitionFAQ() + `
resource "udoma_case_reporting_endpoint" "test" {
	name 							= "updated endpoint"
	active 						= false
	case_categories   = [
		{
			priority = 1
			name = {
				en = "A case"
				de = "Ein Fall"
			}
			templates = [
				{
					priority = 1
					id = udoma_case_template.test.id
				}
			]
		}
	]
	faqs = [ udoma_faq.test.id ]
}
`,

				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_case_reporting_endpoint.test", "name", "updated endpoint"),
					resource.TestCheckResourceAttr("udoma_case_reporting_endpoint.test", "active", "false"),
					resource.TestCheckResourceAttrSet("udoma_case_reporting_endpoint.test", "faqs.0"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}
