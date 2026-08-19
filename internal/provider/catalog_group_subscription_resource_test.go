package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCatalogGroupSubscriptionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCatalogGroupSubscriptionResourceConfig(true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_catalog_group_subscription.test", "code", "catalog-group-test"),
					resource.TestCheckResourceAttr("udoma_catalog_group_subscription.test", "exclude_list.0", "item-code-1"),
					resource.TestCheckResourceAttrSet("udoma_catalog_group_subscription.test", "id"),
					resource.TestCheckResourceAttrSet("udoma_catalog_group_subscription.test", "created_at"),
					resource.TestCheckResourceAttrSet("udoma_catalog_group_subscription.test", "updated_at"),
				),
			},
			{
				ResourceName:      "udoma_catalog_group_subscription.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCatalogGroupSubscriptionResourceConfig(false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_catalog_group_subscription.test", "exclude_list.0", "item-code-2"),
				),
			},
		},
	})
}

func testAccCatalogGroupSubscriptionResourceConfig(initial bool) string {
	excludedItemCode := "item-code-2"
	if initial {
		excludedItemCode = "item-code-1"
	}

	return `
resource "udoma_catalog_group" "test" {
  code = "catalog-group-test"
  description = {
    en = "Catalog group test"
  }
  published = true
}

resource "udoma_catalog_group_subscription" "test" {
  code = udoma_catalog_group.test.code
	exclude_list = ["` + excludedItemCode + `"]
}
`
}
