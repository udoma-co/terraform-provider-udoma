package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCatalogItemDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
resource "udoma_cost_type" "test" {
  name        = "Catalog Item Data Source Cost Type"
  category    = "maintenance"
  description = "Used by catalog item data source test"
}

resource "udoma_catalog_group" "test" {
  code = "catalog-group-data-test"
  description = {
    en = "Catalog group data source"
  }
  published = true
}

resource "udoma_catalog_item" "test" {
  code        = "catalog-item-data-test"
  ref_type    = "cost_type"
  ref_id      = udoma_cost_type.test.id
  account_ref = ` + testAccAccountRef() + `
  groups      = [udoma_catalog_group.test.code]
  description = {
    en = "Catalog item data source"
  }
  published = true
}

data "udoma_catalog_item" "test" {
  code = udoma_catalog_item.test.code
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.udoma_catalog_item.test", "code", "catalog-item-data-test"),
					resource.TestCheckResourceAttr("data.udoma_catalog_item.test", "ref_type", "cost_type"),
					resource.TestCheckResourceAttr("data.udoma_catalog_item.test", "groups.0", "catalog-group-data-test"),
					resource.TestCheckResourceAttr("data.udoma_catalog_item.test", "description.en", "Catalog item data source"),
					resource.TestCheckResourceAttr("data.udoma_catalog_item.test", "published", "true"),
					resource.TestCheckResourceAttrSet("data.udoma_catalog_item.test", "id"),
				),
			},
		},
	})
}
