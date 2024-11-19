package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCustomFormResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: resourceDefinitionCustomForm("basic custom form"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_form.test", "name", "basic custom form"),
					resource.TestCheckResourceAttr("udoma_form.test", "description", "Basic custom form description"),

					resource.TestCheckResourceAttr("udoma_form.test", "form_definition.layout.0.ref_id", "test"),
					resource.TestCheckResourceAttr("udoma_form.test", "form_definition.layout.0.ref_type", "group"),
					resource.TestCheckResourceAttr("udoma_form.test", "form_definition.inputs.0.id", "test"),
					resource.TestCheckResourceAttr("udoma_form.test", "form_definition.inputs.0.type", "text"),
					resource.TestCheckResourceAttr("udoma_form.test", "form_definition.inputs.0.label.de", "Test Eingabe"),
					resource.TestCheckResourceAttr("udoma_form.test", "form_definition.inputs.0.label.en", "Test input"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "udoma_form.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: resourceDefinitionCustomForm("updated custom form"),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_form.test", "name", "updated custom form"),
				),
			},
		},
	})
}

func resourceDefinitionCustomForm(name string) string {
	return `
	resource udoma_form "test" {
		name = "` + name + `"
		description = "Basic custom form description"
		form_definition = {
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
