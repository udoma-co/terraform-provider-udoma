package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccWorkflowDefinitionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: resourceDefinitionWorkflowDefinition(),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify base addtributes
					resource.TestCheckResourceAttr("udoma_workflow_definition.test", "name", "basic workflow"),
					resource.TestCheckResourceAttr("udoma_workflow_definition.test", "name_expression", "\"Workflow\""),
					resource.TestCheckResourceAttr("udoma_workflow_definition.test", "description", "Test description"),
					resource.TestCheckResourceAttr("udoma_workflow_definition.test", "icon", "fa-solid fa-file-alt"),
					resource.TestCheckResourceAttr("udoma_workflow_definition.test", "first_step_id", "generate_document"),

					// Verify env_vars
					resource.TestCheckResourceAttr("udoma_workflow_definition.test", "env_vars.var1", "val1"),

					// Verify init_step, steps
					// TODO

					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("udoma_workflow_definition.test", "id"),
					resource.TestCheckResourceAttrSet("udoma_workflow_definition.test", "last_updated"),
					resource.TestCheckResourceAttrSet("udoma_workflow_definition.test", "created_at"),
					resource.TestCheckResourceAttrSet("udoma_workflow_definition.test", "updated_at"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "udoma_workflow_definition.test",
				ImportState:       true,
				ImportStateVerify: true,
				// The last_updated attribute does not exist in the HashiCups
				// API, therefore there is no value for it during import.
				ImportStateVerifyIgnore: []string{"last_updated"},
			},
			// Update and Read testing
			{
				Config: `
resource "udoma_workflow_definition" "test" {
	name 			= "basic updated workflow"
	description 	= "Test updated description"
	name_expression = "\"New workflow\""
	icon 			= "fa-solid fa-file-alt"

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
			type: "generate_document",
			name: "Generate Document",
			actions: [
				{
					id: "save",
					label: "Save",
					next_step_id: "finish"
				}
			]
		}
	])

}
`,

				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_workflow_definition.test", "name", "basic updated workflow"),
					resource.TestCheckResourceAttr("udoma_workflow_definition.test", "name_expression", "\"New workflow\""),
					resource.TestCheckResourceAttr("udoma_workflow_definition.test", "description", "Test updated description"),
					resource.TestCheckResourceAttr("udoma_workflow_definition.test", "first_step_id", "generate_document"),

					// Verify env_vars
					resource.TestCheckResourceAttr("udoma_workflow_definition.test", "env_vars.var1", "val1"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}

func resourceDefinitionWorkflowDefinition() string {
	return `
	resource udoma_workflow_definition "test" {
		name 			= "basic workflow"
		description 	= "Test description"
		name_expression = "\"Workflow\""
		icon 			= "fa-solid fa-file-alt"
	
		env_vars = {
			var1 = "val1"
		}
	
		first_step_id = "generate_document"
	
		steps = jsonencode([
			{
				id: "generate_document",
				type: "generate_document",
				name: "Generate Document",
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
				type: "finish_execution",
				name: "Finish",
				actions: [
					{
						id: "finish",
						label: "Finish"
					}
				]
			}
		])
	
	}
	`
}
