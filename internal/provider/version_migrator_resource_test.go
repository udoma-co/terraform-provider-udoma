package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccVersionMigratorResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: resourceDefinitionVersionMigrato(),
				Check: resource.ComposeAggregateTestCheckFunc(

					resource.TestCheckResourceAttr("udoma_version_migrator.test", "target_version", "2"),
					resource.TestCheckResourceAttr("udoma_version_migrator.test", "source_version", "1"),
					resource.TestCheckResourceAttr("udoma_version_migrator.test", "script", "test script"),

					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("udoma_version_migrator.test", "id"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "udoma_version_migrator.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: resourceDefinitionCaseTemplate("example template", "A case", 1) + `
				resource udoma_version_migrator "test" {
					ref_id = udoma_case_template.test.id
					target_version = 2
					source_version = 1
					script = "updated script"
				}
				`,

				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_version_migrator.test", "script", "updated script"),
				),
			},

			// Delete testing automatically occurs in TestCase
		},
	})
}

func resourceDefinitionVersionMigrato() string {
	return resourceDefinitionCaseTemplate("example template", "A case", 1) + `
	resource udoma_version_migrator "test" {
		ref_id = udoma_case_template.test.id
		target_version = 2
		source_version = 1
		script = "test script"
	}
	`
}
