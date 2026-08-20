package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	api "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
)

func TestCatalogGroupResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCatalogGroupResourceConfig("catalog-group-test", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_catalog_group.test", "code", "catalog-group-test"),
					resource.TestCheckResourceAttr("udoma_catalog_group.test", "description.en", "Catalog group test"),
					resource.TestCheckResourceAttr("udoma_catalog_group.test", "required_scopes.0", "owner_management"),
					resource.TestCheckResourceAttr("udoma_catalog_group.test", "published", "true"),
					resource.TestCheckResourceAttrSet("udoma_catalog_group.test", "id"),
				),
			},
			{
				ResourceName:      "udoma_catalog_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCatalogGroupResourceConfig("catalog-group-test", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_catalog_group.test", "published", "false"),
				),
			},
		},
	})
}

func testAccCatalogGroupResourceConfig(code string, published bool) string {
	if published {
		return `
resource "udoma_catalog_group" "test" {
  code = "` + code + `"
  description = {
    en = "Catalog group test"
  }
  required_scopes = ["owner_management"]
  published = true
}
`
	}

	return `
resource "udoma_catalog_group" "test" {
  code = "` + code + `"
  description = {
    en = "Catalog group test"
  }
  required_scopes = ["owner_management"]
  published = false
}
`
}

func TestCatalogGroupFromAPIPreservesEmptyRequiredScopes(t *testing.T) {
	model := catalogGroupModel{
		RequiredScopes: types.ListValueMust(types.StringType, nil),
	}

	group := &api.CatalogGroup{Code: "catalog-group-test"}

	if err := model.fromAPI(group); err != nil {
		t.Fatalf("fromAPI returned error: %v", err)
	}

	if model.RequiredScopes.IsNull() {
		t.Fatal("expected required_scopes to remain an empty list, got null")
	}

	if len(model.RequiredScopes.Elements()) != 0 {
		t.Fatalf("expected empty required_scopes, got %d elements", len(model.RequiredScopes.Elements()))
	}
}
