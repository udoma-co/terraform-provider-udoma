package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccBookingTemplateResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: resourceDefinitionBookingTemplate("test_template", "test_description"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("udoma_booking_template.test", "name", "test_template"),
					resource.TestCheckResourceAttr("udoma_booking_template.test", "name_expression", "test_expression"),
					resource.TestCheckResourceAttr("udoma_booking_template.test", "description", "test_description"),
					resource.TestCheckResourceAttr("udoma_booking_template.test", "icon", "fa-solid fa-plus"),

					resource.TestCheckResourceAttr("udoma_booking_template.test", "inputs.layout.0.ref_id", "test"),
					resource.TestCheckResourceAttr("udoma_booking_template.test", "inputs.layout.0.ref_type", "group"),
					resource.TestCheckResourceAttr("udoma_booking_template.test", "inputs.inputs.0.id", "test"),
					resource.TestCheckResourceAttr("udoma_booking_template.test", "inputs.inputs.0.type", "text"),
					resource.TestCheckResourceAttr("udoma_booking_template.test", "inputs.inputs.0.label.en", "Test input"),

					resource.TestCheckResourceAttr("udoma_booking_template.test", "env_vars.FOO", "bar"),
					resource.TestCheckResourceAttr("udoma_booking_template.test", "env_vars.TEST2", "value"),

					resource.TestCheckResourceAttr("udoma_booking_template.test", "script", "// This is a test generation script"),

					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("udoma_booking_template.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "udoma_booking_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: resourceDefinitionBookingTemplate("updated_name", "updated_description"),

				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_booking_template.test", "name", "updated_name"),
					resource.TestCheckResourceAttr("udoma_booking_template.test", "name_expression", "test_expression"),
					resource.TestCheckResourceAttr("udoma_booking_template.test", "description", "updated_description"),
					resource.TestCheckResourceAttr("udoma_booking_template.test", "icon", "fa-solid fa-plus"),

					resource.TestCheckResourceAttr("udoma_booking_template.test", "inputs.layout.0.ref_id", "test"),
					resource.TestCheckResourceAttr("udoma_booking_template.test", "inputs.layout.0.ref_type", "group"),
					resource.TestCheckResourceAttr("udoma_booking_template.test", "inputs.inputs.0.id", "test"),
					resource.TestCheckResourceAttr("udoma_booking_template.test", "inputs.inputs.0.type", "text"),
					resource.TestCheckResourceAttr("udoma_booking_template.test", "inputs.inputs.0.label.en", "Test input"),

					resource.TestCheckResourceAttr("udoma_booking_template.test", "env_vars.FOO", "bar"),
					resource.TestCheckResourceAttr("udoma_booking_template.test", "env_vars.TEST2", "value"),

					resource.TestCheckResourceAttr("udoma_booking_template.test", "trigger_source", "invoice"),

					resource.TestCheckResourceAttr("udoma_booking_template.test", "init_script", "// This is a test init script"),

					resource.TestCheckResourceAttr("udoma_booking_template.test", "script", "// This is a test generation script"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}

func resourceDefinitionBookingTemplate(name, description string) string {
	return `
	resource udoma_booking_template "test" {
		name        = "` + name + `"
		name_expression = "test_expression"
		description = "` + description + `"
		icon		= "fa-solid fa-plus"
		trigger_source = "invoice"
		init_script = "// This is a test init script"
		
		inputs = {
			"layout" = [
				{
					"ref_id" = "test",
					"ref_type" = "group"
				}
			]
			"groups" = [
				{
					"id" = "test"
					"type" = "row"
					"items" = [
						{
							"ref_id" = "test"
							"ref_type" = "input"
						}
					]
				}
			]
			"inputs" = [
				{
					"id" = "test",
					"type" = "text",
					"label" = {
						en = "Test input"
					}
				}
			]
		}

		env_vars = {
			FOO   = "bar"
			TEST2 = "value"
		}

		script = "// This is a test generation script"
	}
	`
}
