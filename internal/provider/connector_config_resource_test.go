package provider

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestConnectorConfigResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: resourceDefinitionConnectorConfig(false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_connector_config.connector", "name", "test-connector"),
					resource.TestCheckResourceAttr("udoma_connector_config.connector", "enabled", "false"),
					resource.TestCheckResourceAttr("udoma_connector_config.connector", "sync_time", "*/5 * * * *"),
					resource.TestCheckResourceAttr("udoma_connector_config.connector", "ping_time", "*/5 * * * *"),
				),
			},
			{
				ResourceName:                         "udoma_connector_config.connector",
				ImportState:                          true,
				ImportStateVerifyIdentifierAttribute: "name",
				ImportStateId:                        "test-connector",
			},
			{
				Config: resourceDefinitionConnectorConfig(true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_connector_config.connector", "name", "test-connector"),
					resource.TestCheckResourceAttr("udoma_connector_config.connector", "enabled", "true"),
					resource.TestCheckResourceAttr("udoma_connector_config.connector", "sync_time", "*/5 * * * *"),
					resource.TestCheckResourceAttr("udoma_connector_config.connector", "ping_time", "*/5 * * * *"),
				),
			},
		},
	})
}

func resourceDefinitionConnectorConfig(enabled bool) string {
	return `
resource "udoma_connector_config" "connector" {
	name = "test-connector"
	enabled = "` + strconv.FormatBool(enabled) + `"
	sync_time = "*/5 * * * *"
	ping_time = "*/5 * * * *"
}
`
}
