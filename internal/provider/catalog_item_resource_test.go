package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	api "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
)

func TestCatalogItemResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccCatalogItemResourceConfig(true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_catalog_item.test", "code", "catalog-item-test"),
					resource.TestCheckResourceAttr("udoma_catalog_item.test", "ref_type", "cost_type"),
					resource.TestCheckResourceAttr("udoma_catalog_item.test", "description.en", "Catalog item test"),
					resource.TestCheckResourceAttr("udoma_catalog_item.test", "groups.0", "catalog-group-test"),
					resource.TestCheckResourceAttr("udoma_catalog_item.test", "published", "true"),
					resource.TestCheckResourceAttrSet("udoma_catalog_item.test", "id"),
				),
			},
			{
				ResourceName:      "udoma_catalog_item.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCatalogItemResourceConfig(false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_catalog_item.test", "published", "false"),
				),
			},
		},
	})
}

func testAccCatalogItemResourceConfig(published bool) string {
	publishedLiteral := "false"
	if published {
		publishedLiteral = "true"
	}

	return `
resource "udoma_cost_type" "test" {
  name        = "Catalog Item Cost Type"
  category    = "maintenance"
  description = "Used by catalog item test"
}

resource "udoma_catalog_group" "test" {
  code = "catalog-group-test"
  description = {
    en = "Catalog group test"
  }
  required_scopes = ["owner_management"]
  published = true
}

resource "udoma_catalog_item" "test" {
  code        = "catalog-item-test"
  ref_type    = "cost_type"
  ref_id      = udoma_cost_type.test.id
  account_ref = ` + testAccAccountRef() + `
  groups      = [udoma_catalog_group.test.code]
  description = {
    en = "Catalog item test"
  }
  published = ` + publishedLiteral + `
  required_scopes = ["owner_management"]
  dependencies = ["dep-a"]
  patches = [
    {
      code   = "patch-1"
      format = "json_patch"
      patch  = "[]"
    }
  ]
}
`
}

func TestTestAccAccountRefUsesEnvironment(t *testing.T) {
	t.Setenv("UDOMA_ACCOUNT_REF", "42")
	if got := testAccAccountRef(); got != "42" {
		t.Fatalf("expected UDOMA_ACCOUNT_REF to be used, got %q", got)
	}
}

func TestCatalogItemFromAPIPreservesEmptyConfiguredLists(t *testing.T) {
	model := catalogItemModel{
		Groups:         types.ListValueMust(types.StringType, nil),
		RequiredScopes: types.ListValueMust(types.StringType, nil),
		Dependencies:   types.ListValueMust(types.StringType, nil),
		Patches:        make([]catalogItemPatchDefinitionModel, 0),
	}

	item := &api.CatalogItem{
		Code:       "catalog-item-test",
		RefType:    api.CATALOGITEMTYPE_COST_TYPE,
		RefId:      "ref-1",
		AccountRef: 1,
	}

	if err := model.fromAPI(item); err != nil {
		t.Fatalf("fromAPI returned error: %v", err)
	}

	if model.Groups.IsNull() {
		t.Fatal("expected groups to remain an empty list, got null")
	}
	if len(model.Groups.Elements()) != 0 {
		t.Fatalf("expected empty groups, got %d elements", len(model.Groups.Elements()))
	}

	if model.RequiredScopes.IsNull() {
		t.Fatal("expected required_scopes to remain an empty list, got null")
	}
	if len(model.RequiredScopes.Elements()) != 0 {
		t.Fatalf("expected empty required_scopes, got %d elements", len(model.RequiredScopes.Elements()))
	}

	if model.Dependencies.IsNull() {
		t.Fatal("expected dependencies to remain an empty list, got null")
	}
	if len(model.Dependencies.Elements()) != 0 {
		t.Fatalf("expected empty dependencies, got %d elements", len(model.Dependencies.Elements()))
	}

	if model.Patches == nil {
		t.Fatal("expected patches to remain an empty list, got nil")
	}
	if len(model.Patches) != 0 {
		t.Fatalf("expected empty patches, got %d elements", len(model.Patches))
	}
}
