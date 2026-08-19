package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCatalogGroupDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "udoma_catalog_group" "test" {
  code = "catalog-group-data-test"
  description = {
    en = "Catalog group data source"
  }
  required_scopes = ["owner_management"]
  published = true
}

data "udoma_catalog_group" "test" {
  code = udoma_catalog_group.test.code
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.udoma_catalog_group.test", "code", "catalog-group-data-test"),
					resource.TestCheckResourceAttr("data.udoma_catalog_group.test", "description.en", "Catalog group data source"),
					resource.TestCheckResourceAttr("data.udoma_catalog_group.test", "required_scopes.0", "owner_management"),
					resource.TestCheckResourceAttr("data.udoma_catalog_group.test", "published", "true"),
					resource.TestCheckResourceAttrSet("data.udoma_catalog_group.test", "id"),
				),
			},
		},
	})
}
