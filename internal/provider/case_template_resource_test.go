package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCaseTemplateResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: resourceDefinitionCaseTemplate("basic template", "Basic case"),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify base addtributes
					resource.TestCheckResourceAttr("udoma_case_template.test", "name", "basic template"),
					resource.TestCheckResourceAttr("udoma_case_template.test", "name_expression", "\"Basic case\""),
					resource.TestCheckResourceAttr("udoma_case_template.test", "icon", "fa fa-bug"),
					resource.TestCheckResourceAttr("udoma_case_template.test", "ad_categories.0", "ELECTRICITY_PROVIDERS"),

					// Verify map attributes
					resource.TestCheckResourceAttr("udoma_case_template.test", "label.de", "Einfache Meldung"),
					resource.TestCheckResourceAttr("udoma_case_template.test", "label.en", "Basic case"),

					resource.TestCheckResourceAttr("udoma_case_template.test", "description.de", "Test Beschreibung"),
					resource.TestCheckResourceAttr("udoma_case_template.test", "description.en", "Test description"),

					resource.TestCheckResourceAttr("udoma_case_template.test", "info_text.de", "Test Info Text"),
					resource.TestCheckResourceAttr("udoma_case_template.test", "info_text.en", "Test info text"),

					// Verify custom_inputs
					// TODO

					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("udoma_case_template.test", "id"),
					resource.TestCheckResourceAttrSet("udoma_case_template.test", "last_updated"),
					resource.TestCheckResourceAttrSet("udoma_case_template.test", "created_at"),
					resource.TestCheckResourceAttrSet("udoma_case_template.test", "updated_at"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "udoma_case_template.test",
				ImportState:       true,
				ImportStateVerify: true,
				// The last_updated attribute does not exist in the Udoma API,
				// therefore there is no value for it during import. JSON types
				// are matched literally, so we need to ignore them.
				ImportStateVerifyIgnore: []string{"last_updated", "custom_inputs", "config"},
			},
			// Update and Read testing
			{
				Config: resourceDefinitionCaseTemplate("updated template", "Special case"),

				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_case_template.test", "name", "updated template"),
					resource.TestCheckResourceAttr("udoma_case_template.test", "name_expression", "\"Special case\""),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}

func resourceDefinitionCaseTemplate(name, nameExpression string) string {
	return `
	resource udoma_case_template "test" {
		name 							= "` + name + `"
		access = ["PROPERTY_MANAGER"]
		name_expression 	= "\"` + nameExpression + `\""
		icon 							= "fa fa-bug"
	
		label = {
			de = "Einfache Meldung"
			en = "Basic case"
		}
	
		description = {
			de = "Test Beschreibung"
			en = "Test description"
		}
	
		info_text = {
			de = "Test Info Text"
			en = "Test info text"
		}

	
		custom_inputs = <<EOF
		{
			"layout": [
				{
					"ref_id": "test",
					"ref_type": "input"
				}
			],
			"inputs": [
				{
					"id": "test",
					"type": "text",
					"label": {
						"de": "Test Eingabe",
						"en": "Test input"
					}
				}
			]
		}
		EOF

		config = {
			base_config = "SIMPLE",
			extend_default_status_config = true
		}

		ad_categories = ["ELECTRICITY_PROVIDERS"]	
	}
	`
}
