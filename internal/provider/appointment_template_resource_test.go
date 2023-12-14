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

					// Verify custom_inputs
					// TODO
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

		inputs = {
			"layout" = [
				{
					"ref_id" = "test",
					"ref_type" = "input"
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
		}
	}
	`
}
