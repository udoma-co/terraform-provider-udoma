package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccHookResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: resourceDefinitionHook("BANK_ACCOUNT"),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("udoma_hook.test", "entity", "BANK_ACCOUNT"),
					resource.TestCheckResourceAttr("udoma_hook.test", "action", "CREATE"),
					resource.TestCheckResourceAttr("udoma_hook.test", "priority", "1"),
					resource.TestCheckResourceAttr("udoma_hook.test", "enabled", "true"),
					resource.TestCheckResourceAttr("udoma_hook.test", "break_on_error", "true"),
					resource.TestCheckResourceAttr("udoma_hook.test", "script", "data"),

					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("udoma_hook.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "udoma_hook.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: resourceDefinitionHook("SERVICE_PROVIDER"),

				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_hook.test", "entity", "SERVICE_PROVIDER"),
				),
			},
		},
	})
}

func resourceDefinitionHook(entity string) string {
	return `
	resource udoma_hook "test" {
		entity = "` + entity + `"
		action = "CREATE"
		priority = 1
		break_on_error = true
		enabled = true
		script = "data"
	}
	`
}
