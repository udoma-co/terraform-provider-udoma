package provider

import (
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestConnectorQueryResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: resourceDefinitionConnectorQuery(true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_connector_query.query", "name", "test-query"),
					resource.TestCheckResourceAttr("udoma_connector_query.query", "connector", "test-connector"),
					resource.TestCheckResourceAttr("udoma_connector_query.query", "enabled", "true"),
					resource.TestCheckResourceAttr("udoma_connector_query.query", "priority", "10"),
					resource.TestCheckResourceAttr("udoma_connector_query.query", "limit", "100"),
					resource.TestCheckResourceAttr("udoma_connector_query.query", "statement", "SELECT * FROM BANK_ACCOUNTS AS B;"),
					resource.TestCheckResourceAttr("udoma_connector_query.query", "script", "info(data);"),

					resource.TestCheckResourceAttrSet("udoma_connector_query.query", "id"),
				),
			},
			{
				ResourceName:      "udoma_connector_query.query",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: resourceDefinitionConnectorQuery(false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_connector_query.query", "name", "test-query"),
					resource.TestCheckResourceAttr("udoma_connector_query.query", "connector", "test-connector"),
					resource.TestCheckResourceAttr("udoma_connector_query.query", "enabled", "false"),
					resource.TestCheckResourceAttr("udoma_connector_query.query", "priority", "10"),
					resource.TestCheckResourceAttr("udoma_connector_query.query", "limit", "100"),
					resource.TestCheckResourceAttr("udoma_connector_query.query", "statement", "SELECT * FROM BANK_ACCOUNTS AS B;"),
					resource.TestCheckResourceAttr("udoma_connector_query.query", "script", "info(data);"),

					resource.TestCheckResourceAttrSet("udoma_connector_query.query", "id"),
				),
			},
		},
	})
}

func resourceDefinitionConnectorQuery(enabled bool) string {
	return resourceDefinitionConnectorConfig(true) + `


resource "udoma_connector_query" "query" {
	name = "test-query"
	connector = udoma_connector_config.connector.name
	enabled = "` + strconv.FormatBool(enabled) + `"

	priority = 10
	limit = 100

	statement = "SELECT * FROM BANK_ACCOUNTS AS B;"
	script = "info(data);"
}
`
}
