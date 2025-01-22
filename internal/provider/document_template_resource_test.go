package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDocumentTemplateResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: `
resource udoma_document_template "test" {
	name 			= "doc-template"
	name_expression = "\"Document\""
	description 	= "Test description"
	version 		= 1

	options = {
		allow_text_edit			= true
		include_footer_branding = true
		include_page_numbers	= true
	}

	content = <<EOF
{
	"entries": [
		{
			"type": "paragraph",
			"value": {
				"id": "C1A98FAA",
				"text": [
					{
						"value": "Test"
					}
				]
			}
		}
	]
}
EOF

	placeholders_script = <<EOF
const something = {
	Key: "Value"
};

something
EOF

	signatures = <<EOF
	{
		"esignatures_enabled": true
	}
EOF

	inputs = <<EOF
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
}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify base addtributes
					resource.TestCheckResourceAttr("udoma_document_template.test", "name", "doc-template"),
					resource.TestCheckResourceAttr("udoma_document_template.test", "name_expression", "\"Document\""),
					resource.TestCheckResourceAttr("udoma_document_template.test", "description", "Test description"),

					// Verify options attributes
					resource.TestCheckResourceAttr("udoma_document_template.test", "options.allow_text_edit", "true"),
					resource.TestCheckResourceAttr("udoma_document_template.test", "options.include_footer_branding", "true"),
					resource.TestCheckResourceAttr("udoma_document_template.test", "options.include_page_numbers", "true"),
					resource.TestCheckResourceAttr("udoma_document_template.test", "version", "1"),

					// Verify custom_inputs
					// TODO

					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("udoma_document_template.test", "id"),
					resource.TestCheckResourceAttrSet("udoma_document_template.test", "last_updated"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "udoma_document_template.test",
				ImportState:       true,
				ImportStateVerify: true,
				// The last_updated attribute does not exist in the Udoma API,
				// therefore there is no value for it during import. JSON types
				// are matched literally, so we need to ignore them.
				ImportStateVerifyIgnore: []string{"last_updated", "content", "placeholders_script", "signatures", "inputs"},
			},
			// Update and Read testing
			{
				Config: `
resource "udoma_document_template" "test" {
	name 			= "doc-template-updated"
	name_expression = "\"Document name\""
	description 	= "Test description"
	version 		= 2

	options = {
		allow_text_edit 		= true
		include_footer_branding = true
		include_page_numbers	= false
	}

	content = <<EOF
{
	"entries": [
		{
			"type": "paragraph",
			"value": {
				"id": "C1A98FAA",
				"text": [
					{
						"value": "Test"
					}
				]
			}
		}
	]
}
EOF

	placeholders_script = <<EOF
const something = {
	Key: "Value"
};

something
EOF

	inputs = <<EOF
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
}
`,

				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_document_template.test", "name", "doc-template-updated"),
					resource.TestCheckResourceAttr("udoma_document_template.test", "name_expression", "\"Document name\""),
					resource.TestCheckResourceAttr("udoma_document_template.test", "description", "Test description"),
					resource.TestCheckResourceAttr("udoma_document_template.test", "options.include_page_numbers", "false"),
					resource.TestCheckResourceAttr("udoma_document_template.test", "version", "2"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}
