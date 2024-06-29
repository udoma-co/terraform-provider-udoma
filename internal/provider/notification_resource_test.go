package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestNotificationResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: resourceDefinitionNotification("basic_notification", "Test description"),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify base addtributes
					resource.TestCheckResourceAttr("udoma_notification.test", "name", "basic_notification"),
					resource.TestCheckResourceAttr("udoma_notification.test", "description", "Test description"),
					resource.TestCheckResourceAttr("udoma_notification.test", "script", "Test script"),
					resource.TestCheckResourceAttr("udoma_notification.test", "template_html", "Test html template"),
					resource.TestCheckResourceAttr("udoma_notification.test", "template_text", "Test text template"),
				),
			},
			// ImportState testing
			{
				ResourceName:                         "udoma_notification.test",
				ImportState:                          true,
				ImportStateVerifyIdentifierAttribute: "name",
				ImportStateId:                        "basic_notification",
			},
			{
				Config: resourceDefinitionNotification("basic_notification", "updated description"),

				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_notification.test", "description", "updated description"),
				),
			},

			// Delete testing automatically occurs in TestNotification
		},
	})
}

func resourceDefinitionNotification(name, description string) string {
	return `
resource udoma_notification "test" {
	name          = "` + name + `"
	description   = "` + description + `"
	script        = "Test script"
	template_html = "Test html template"
	template_text = "Test text template"
}`
}
