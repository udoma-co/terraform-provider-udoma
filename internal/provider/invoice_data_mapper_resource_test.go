package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	api "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
)

func TestInvoiceDataMapperResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: resourceDefinitionInvoiceDataMapper("AFTER_EXTRACT_DOCUMENTS"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_invoice_data_mapper.test", "name", "hello"),
					resource.TestCheckResourceAttr("udoma_invoice_data_mapper.test", "description", "Hello description"),
					resource.TestCheckResourceAttr("udoma_invoice_data_mapper.test", "priority", "100"),
					resource.TestCheckResourceAttr("udoma_invoice_data_mapper.test", "script", "console.log(42)"),
					resource.TestCheckResourceAttr("udoma_invoice_data_mapper.test", "entrypoint", "AFTER_EXTRACT_DOCUMENTS"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "udoma_invoice_data_mapper.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: resourceDefinitionInvoiceDataMapper("AFTER_PROCESS_DOCUMENTS"),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify attributes were updated
					resource.TestCheckResourceAttr("udoma_invoice_data_mapper.test", "name", "hello"),
					resource.TestCheckResourceAttr("udoma_invoice_data_mapper.test", "description", "Hello description"),
					resource.TestCheckResourceAttr("udoma_invoice_data_mapper.test", "priority", "100"),
					resource.TestCheckResourceAttr("udoma_invoice_data_mapper.test", "script", "console.log(42)"),
					resource.TestCheckResourceAttr("udoma_invoice_data_mapper.test", "entrypoint", "AFTER_PROCESS_DOCUMENTS"),
				),
			},
		},
	})
}

func resourceDefinitionInvoiceDataMapper(entrypoint api.InvoiceDataMapperEntrypointEnum) string {
	return `
	resource udoma_invoice_data_mapper "test" {
		name = "hello"
		description = "Hello description"
		priority = 100
		entrypoint = "` + string(entrypoint) + `"
		script = "console.log(42)"
	}`
}
