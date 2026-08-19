package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestCatalogItemSubscriptionResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCatalogItemSubscriptionResourceConfig(true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_catalog_item_subscription.test", "code", "catalog-item-test"),
					resource.TestCheckResourceAttr("udoma_catalog_item_subscription.test", "patches.0.patch_ref", "patch-ref-1"),
					resource.TestCheckResourceAttrSet("udoma_catalog_item_subscription.test", "id"),
					resource.TestCheckResourceAttrSet("udoma_catalog_item_subscription.test", "created_at"),
					resource.TestCheckResourceAttrSet("udoma_catalog_item_subscription.test", "updated_at"),
				),
			},
			{
				ResourceName:      "udoma_catalog_item_subscription.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCatalogItemSubscriptionResourceConfig(false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_catalog_item_subscription.test", "patches.0.patch_ref", "patch-ref-2"),
				),
			},
		},
	})
}

func testAccCatalogItemSubscriptionResourceConfig(initial bool) string {
	patchRef := "patch-ref-2"
	if initial {
		patchRef = "patch-ref-1"
	}

	return `
resource "udoma_cost_type" "test" {
  name        = "Catalog Item Subscription Cost Type"
  category    = "maintenance"
  description = "Used by catalog item subscription test"
}

resource "udoma_catalog_item" "test" {
  code        = "catalog-item-test"
  ref_type    = "cost_type"
  ref_id      = udoma_cost_type.test.id
  account_ref = 1
  published   = true
}

resource "udoma_catalog_item_subscription" "test" {
  code = udoma_catalog_item.test.code
  patches = [
    {
      patch_ref = "` + patchRef + `"
    }
  ]
}
`
}
