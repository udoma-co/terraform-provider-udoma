package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataImportTemplateResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: resourceDefinitionDataImportTemplate("Customer Import", "CSV", true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_data_import_template.test", "name", "Customer Import"),
					resource.TestCheckResourceAttr("udoma_data_import_template.test", "file_type", "CSV"),
					resource.TestCheckResourceAttr("udoma_data_import_template.test", "has_header", "true"),
					resource.TestCheckResourceAttr("udoma_data_import_template.test", "data_mapper", "CUSTOMER_MAPPER"),
					resource.TestCheckResourceAttr("udoma_data_import_template.test", "description", "Template for importing customer data"),
					resource.TestCheckResourceAttr("udoma_data_import_template.test", "icon", "fa fa-upload"),

					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("udoma_data_import_template.test", "id"),
					resource.TestCheckResourceAttrSet("udoma_data_import_template.test", "created_at"),
					resource.TestCheckResourceAttrSet("udoma_data_import_template.test", "updated_at"),
				),
			},

			// ImportState testing
			{
				ResourceName:      "udoma_data_import_template.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"last_updated",
				},
			},

			{
				Config: resourceDefinitionDataImportTemplate("Customer Import Updated", "XLSX", false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_data_import_template.test", "name", "Customer Import Updated"),
					resource.TestCheckResourceAttr("udoma_data_import_template.test", "file_type", "XLSX"),
					resource.TestCheckResourceAttr("udoma_data_import_template.test", "has_header", "false"),
				),
			},
		},
	})
}

func resourceDefinitionDataImportTemplate(name, fileType string, hasHeader bool) string {
	return fmt.Sprintf(`
resource "udoma_data_import_template" "test" {
	name        = "%s"
	description = "Template for importing customer data"
	icon        = "fa fa-upload"
	file_type   = "%s"
	has_header  = %t
	data_mapper = "CUSTOMER_MAPPER"
}
`, name, fileType, hasHeader)
}
