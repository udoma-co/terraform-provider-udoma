package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAppointmentTemplateResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: resourceDefinitionAppointmentTemplate("basic template", "Basic template"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_appointment_template.test", "name", "basic template"),
					resource.TestCheckResourceAttr("udoma_appointment_template.test", "name_expression", "\"Basic template\""),
					resource.TestCheckResourceAttr("udoma_appointment_template.test", "description", "Basic template description"),

					resource.TestCheckResourceAttr("udoma_appointment_template.test", "inputs.layout.0.ref_id", "test"),
					resource.TestCheckResourceAttr("udoma_appointment_template.test", "inputs.layout.0.ref_type", "group"),
					resource.TestCheckResourceAttr("udoma_appointment_template.test", "inputs.inputs.0.id", "test"),
					resource.TestCheckResourceAttr("udoma_appointment_template.test", "inputs.inputs.0.type", "text"),
					resource.TestCheckResourceAttr("udoma_appointment_template.test", "inputs.inputs.0.label.de", "Test Eingabe"),
					resource.TestCheckResourceAttr("udoma_appointment_template.test", "inputs.inputs.0.label.en", "Test input"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "udoma_appointment_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: resourceDefinitionAppointmentTemplate("updated template", "Special appointment"),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_appointment_template.test", "name", "updated template"),
					resource.TestCheckResourceAttr("udoma_appointment_template.test", "name_expression", "\"Special appointment\""),
				),
			},
		},
	})
}

func resourceDefinitionAppointmentTemplate(name, nameExpression string) string {
	return `
	resource udoma_appointment_template "test" {
		name = "` + name + `"
		name_expression = "\"` + nameExpression + `\""
		description = "Basic template description"

		require_confirmation = true
		confirmation_reminders = [
			1, 2, 3, 1
		]

		inputs = {
			"layout" = [
				{
					"ref_id" = "test",
					"ref_type" = "group"
				}
			]
			"inputs" = [
				{
					"id" = "test",
					"type" = "text",
					"label" = {
						de = "Test Eingabe",
						en = "Test input"
					}
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
			"validation" = [
				{
					"id" = "test"
					"expression" = "1 == 1"
					"target" = "test"
					"message" = {
						"en" = "Helloooo"
					}
				}
			]
		}
	}
	`
}
