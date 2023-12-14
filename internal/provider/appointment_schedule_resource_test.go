package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAppointmentScheduleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: resourceDefinitionAppointmentSchedule("basic schedule", "Basic schedule"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_appointment_schedule.test_schedule", "name", "basic schedule"),
					resource.TestCheckResourceAttr("udoma_appointment_schedule.test_schedule", "description.en", "Basic schedule"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "udoma_appointment_schedule.test_schedule",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: resourceDefinitionAppointmentSchedule("updated schedule", "Special appointment schedule"),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_appointment_schedule.test_schedule", "name", "updated schedule"),
					resource.TestCheckResourceAttr("udoma_appointment_schedule.test_schedule", "description.en", "Special appointment schedule"),
				),
			},
		},
	})
}

func resourceDefinitionAppointmentSchedule(name, description string) string {
	return `
	locals {
		inputs = <<-EOT
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
						EOT
	}

	resource udoma_appointment_template "test_template" {
		name = "Test template"
		name_expression = "\"Hello appointment\""
		description = "Basic template description"

		inputs = jsondecode(local.inputs)
	}

	resource udoma_appointment_schedule "test_schedule" {
		name = "` + name + `"
		description = {
			en = "` + description + `"
		}
		template_ref = udoma_appointment_template.test_template.id
		slot_duration = 45
		gap_duration = 15
		color = "green"
	}
	`
}
