package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccReportEntrypointResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: resourceDefinitionReportEntrypoint("properties", "Generate Report", "Bericht generieren", true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_report_entrypoint.test", "app_location", "properties"),
					resource.TestCheckResourceAttr("udoma_report_entrypoint.test", "label.en", "Generate Report"),
					resource.TestCheckResourceAttr("udoma_report_entrypoint.test", "label.de", "Bericht generieren"),
					resource.TestCheckResourceAttr("udoma_report_entrypoint.test", "quick_execute", "true"),
					resource.TestCheckResourceAttrSet("udoma_report_entrypoint.test", "id"),
					resource.TestCheckResourceAttrSet("udoma_report_entrypoint.test", "last_updated"),
					resource.TestCheckResourceAttrSet("udoma_report_entrypoint.test", "created_at"),
					resource.TestCheckResourceAttrSet("udoma_report_entrypoint.test", "updated_at"),
				),
			},

			// ImportState testing
			{
				ResourceName:         "udoma_report_entrypoint.test",
				ImportState:          true,
				ImportStateVerify:    true,
				ImportStateVerifyIgnore: []string{"last_updated"},
			},

			// Update and Read testing
			{
				Config: resourceDefinitionReportEntrypoint("tenancies", "Generate Updated Report", "Aktualisierten Bericht generieren", false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_report_entrypoint.test", "app_location", "tenancies"),
					resource.TestCheckResourceAttr("udoma_report_entrypoint.test", "label.en", "Generate Updated Report"),
					resource.TestCheckResourceAttr("udoma_report_entrypoint.test", "label.de", "Aktualisierten Bericht generieren"),
					resource.TestCheckResourceAttr("udoma_report_entrypoint.test", "quick_execute", "false"),
				),
			},
		},
	})
}

func resourceDefinitionReportEntrypoint(appLocation string, labelEn string, labelDe string, quickExecute bool) string {
	return fmt.Sprintf(`
resource "udoma_report_definition" "test" {
    name        = "test report definition"
    description = "Test description"
    script      = "const ret = { count: 1 }; return ret;"
    result_schema = {
        result_type = "TABLE"
        attributes = [
            {
                id       = "count"
                label    = { en = "Count", de = "Anzahl" }
                sequence = 1
            }
        ]
    }
}

resource "udoma_report_entrypoint" "test" {
    app_location         = "%s"
    label = {
        en = "%s"
        de = "%s"
    }
    quick_execute = %t
    report_definition_ref = udoma_report_definition.test.id
}
`, appLocation, labelEn, labelDe, quickExecute)
}
