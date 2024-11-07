package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccWorkflowEntrypointResource(t *testing.T) {

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: resourceDefinitionWorkflowDefinition() + `
resource udoma_workflow_entrypoint "test" {
	workflow_definition_ref = udoma_workflow_definition.test.id
	app_location 			= "case"
	location_filter 		= "ct-JsDnoGWuXrFXqJWNFGBXUN"
	icon 					= "fa-solid fa-file-alt"

	label = {
		en = "Start workflow"
		de = "Starte Workflow"
	}

	init_script = <<EOF
	// This is a comment
	// This is another comment
EOF

}
`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify base addtributes
					resource.TestCheckResourceAttr("udoma_workflow_entrypoint.test", "app_location", "case"),
					resource.TestCheckResourceAttr("udoma_workflow_entrypoint.test", "location_filter", "ct-JsDnoGWuXrFXqJWNFGBXUN"),
					resource.TestCheckResourceAttr("udoma_workflow_entrypoint.test", "icon", "fa-solid fa-file-alt"),

					// Verify label
					resource.TestCheckResourceAttr("udoma_workflow_entrypoint.test", "label.de", "Starte Workflow"),
					resource.TestCheckResourceAttr("udoma_workflow_entrypoint.test", "label.en", "Start workflow"),

					// Verify init_script
					// TODO

					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("udoma_workflow_entrypoint.test", "id"),
					resource.TestCheckResourceAttrSet("udoma_workflow_entrypoint.test", "last_updated"),
					resource.TestCheckResourceAttrSet("udoma_workflow_entrypoint.test", "created_at"),
					resource.TestCheckResourceAttrSet("udoma_workflow_entrypoint.test", "updated_at"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "udoma_workflow_entrypoint.test",
				ImportState:       true,
				ImportStateVerify: true,
				// The last_updated attribute does not exist in the HashiCups
				// API, therefore there is no value for it during import.
				ImportStateVerifyIgnore: []string{"last_updated"},
			},
			// Update and Read testing
			{
				Config: resourceDefinitionWorkflowDefinition() + `
resource "udoma_workflow_entrypoint" "test" {
	workflow_definition_ref = udoma_workflow_definition.test.id
	app_location 			= "manual"
	icon 					= "fa-solid fa-file-alt"

	label = {
		en = "Start workflow"
		de = "Starte Workflow"
	}

	init_script = <<EOF
	// This is a comment
	// This is another comment
EOF

}
`,

				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_workflow_entrypoint.test", "app_location", "manual"),
					resource.TestCheckResourceAttr("udoma_workflow_entrypoint.test", "icon", "fa-solid fa-file-alt"),

					// Verify label
					resource.TestCheckResourceAttr("udoma_workflow_entrypoint.test", "label.de", "Starte Workflow"),
					resource.TestCheckResourceAttr("udoma_workflow_entrypoint.test", "label.en", "Start workflow"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}
