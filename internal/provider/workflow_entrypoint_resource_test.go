package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccWorkflowEntrypointResource(t *testing.T) {

	workflowDefinition := `resource udoma_workflow_definition "test" {
		name 							= "basic workflow"
		description 			= "Test description"
		name_expression 	= "\"Workflow\""
		icon 							= "fa-solid fa-file-alt"
	
		env_vars = {
			var1 = "val1"
		}
	
		first_step_id = "generate_document"
	
		init_step = jsonencode({
			id: "init",
			type: "select_property"
		})
	
		steps = jsonencode([
			{
				id: "generate_document",
				type: "generate_document"
				actions: [
					{
						id: "save",
						label: "Save",
						next_step_id: "finish"
					}
				]
			},
	
			{
				id: "finish",
				type: "finish_execution"
				actions: [
					{
						id: "finish"
					}
				]
			}
		])
	
	}
	`

	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: workflowDefinition + `
resource udoma_workflow_entrypoint "test" {
	workflow_definition_ref = udoma_workflow_definition.test.id
	app_location 						= "case"
	location_filter 				= "ct-JsDnoGWuXrFXqJWNFGBXUN"
	icon 										= "fa-solid fa-file-alt"
	skip_init_step 					= true

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
					resource.TestCheckResourceAttr("udoma_workflow_entrypoint.test", "skip_init_step", "true"),

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
				Config: workflowDefinition + `
resource "udoma_workflow_entrypoint" "test" {
	workflow_definition_ref = udoma_workflow_definition.test.id
	app_location 						= "manual"
	icon 										= "fa-solid fa-file-alt"
	skip_init_step 					= true

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
					resource.TestCheckResourceAttr("udoma_workflow_entrypoint.test", "skip_init_step", "true"),

					// Verify label
					resource.TestCheckResourceAttr("udoma_workflow_entrypoint.test", "label.de", "Starte Workflow"),
					resource.TestCheckResourceAttr("udoma_workflow_entrypoint.test", "label.en", "Start workflow"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}
