package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestPropertyHandoverTemplateResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: propertyHandoverTemplateResource("test-template", 1),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify base addtributes
					resource.TestCheckResourceAttr("udoma_property_handover_template.test_handover_template", "name", "test-template"),
					resource.TestCheckResourceAttr("udoma_property_handover_template.test_handover_template", "description", "Test Handover Template Description"),
					resource.TestCheckResourceAttr("udoma_property_handover_template.test_handover_template", "inputs.layout.0.ref_id", "test"),
					resource.TestCheckResourceAttr("udoma_property_handover_template.test_handover_template", "version", "1"),
				),
			},
			{
				ResourceName:      "udoma_property_handover_template.test_handover_template",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: propertyHandoverTemplateResource("updated-template", 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify base addtributes
					resource.TestCheckResourceAttr("udoma_property_handover_template.test_handover_template", "name", "updated-template"),
					resource.TestCheckResourceAttr("udoma_property_handover_template.test_handover_template", "description", "Test Handover Template Description"),
					resource.TestCheckResourceAttr("udoma_property_handover_template.test_handover_template", "inputs.layout.0.ref_id", "test"),
					resource.TestCheckResourceAttr("udoma_property_handover_template.test_handover_template", "version", "2"),
				),
			},
		},
	})
}

func propertyHandoverTemplateResource(name string, version int32) string {
	return `
resource "udoma_property_handover_template" "test_handover_template" {
	name = "` + name + `"
	description = "Test Handover Template Description"
	version = ` + fmt.Sprint(version) + `
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
