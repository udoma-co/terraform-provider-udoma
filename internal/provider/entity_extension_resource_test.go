package provider

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccEntityExtensionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: resourceDefinitionEntityExtension("Parent Tenancy", fmt.Sprintf("parent_tenancy_%d", time.Now().UnixNano())),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_entity_extension.test", "name", "Parent Tenancy"),
					resource.TestCheckResourceAttr("udoma_entity_extension.test", "description", "Manage the tenancy for the property with the owner"),
					resource.TestCheckResourceAttr("udoma_entity_extension.test", "entity", "property"),
					resource.TestCheckResourceAttr("udoma_entity_extension.test", "sequence", "1"),
					resource.TestCheckResourceAttrSet("udoma_entity_extension.test", "id"),
				),
			},

			// ImportState testing
			{
				ResourceName:      "udoma_entity_extension.test",
				ImportState:       true,
				ImportStateVerify: true,
			},

			// Update testing
			{
				Config: resourceDefinitionEntityExtension("Parent Tenancy Updated", fmt.Sprintf("parent_tenancy_%d", time.Now().UnixNano())),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_entity_extension.test", "name", "Parent Tenancy Updated"),
				),
			},
		},
	})
}


func resourceDefinitionEntityExtension(name string, key string) string {
    return fmt.Sprintf(`
resource "udoma_entity_extension" "test" {
  name        = "%s"
  description = "Manage the tenancy for the property with the owner"
  entity      = "property"
  sequence    = 1
  key         = "%s"
  version     = 2

  form_definition = {
    layout = [
      { ref_id = "property_123", ref_type = "slot" },
      { ref_id = "tenant_456",   ref_type = "slot" }
    ]
    inputs      = []
    groups      = []
    validation  = []
  }
}
`, name, key)
}

