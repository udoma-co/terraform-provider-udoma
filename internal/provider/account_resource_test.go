package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	v1 "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
)

func TestAccountResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAccountConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_account.test", "name", "Test Account"),
					resource.TestCheckResourceAttr("udoma_account.test", "currency", "USD"),
					resource.TestCheckResourceAttr("udoma_account.test", "number", "128"),
					resource.TestCheckResourceAttr("udoma_account.test", "type", "asset"),
				),
			},
			{
				ResourceName:      "udoma_account.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: `
				resource udoma_account "test" {
					number   = 501
					name     = "Updated Account"
					currency = "EUR"
					type 	   = "asset"
					cadence  = "monthly_cadence"
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("udoma_account.test", "name", "Updated Account"),
					resource.TestCheckResourceAttr("udoma_account.test", "currency", "EUR"),
					resource.TestCheckResourceAttr("udoma_account.test", "number", "501"),
				),
			},
		},
	})
}

func TestAccountModelRevenueTypeRefRoundTrip(t *testing.T) {
	model := AccountModel{
		Number:         types.Int32Value(128),
		Name:           types.StringValue("Revenue account"),
		Type:           types.StringValue("asset"),
		Currency:       types.StringValue("USD"),
		Cadence:        types.StringValue("monthly_cadence"),
		RevenueTypeRef: types.StringValue("revenue-type-ref"),
	}

	request, err := model.toAPIRequest()
	if err != nil {
		t.Fatalf("toAPIRequest returned error: %v", err)
	}
	if request.RevenueTypeRef == nil || *request.RevenueTypeRef != "revenue-type-ref" {
		t.Fatalf("expected revenue_type_ref to be set to revenue-type-ref, got %#v", request.RevenueTypeRef)
	}
	if request.CostTypeRef != nil {
		t.Fatalf("expected cost_type_ref to stay unset when revenue_type_ref is set, got %#v", request.CostTypeRef)
	}

	cadence := v1.BALANCECADENCEENUM_MONTHLY_CADENCE
	payload := &v1.FinancialAccount{
		Id:             "account-id",
		Number:         128,
		Name:           "Revenue account",
		Currency:       "USD",
		CreatedAt:      1,
		UpdatedAt:      1,
		Type:           v1.ACCOUNTTYPESENUM_ASSET,
		Cadence:        &cadence,
		RevenueTypeRef: stringPtr("revenue-type-ref"),
	}

	if diags := model.fromAPI(payload); diags.HasError() {
		t.Fatalf("fromAPI returned diagnostics: %v", diags.Errors())
	}
	if model.RevenueTypeRef.IsNull() || model.RevenueTypeRef.ValueString() != "revenue-type-ref" {
		t.Fatalf("expected fromAPI to hydrate revenue_type_ref, got %#v", model.RevenueTypeRef)
	}
}

func stringPtr(s string) *string { return &s }

func testAccAccountConfig() string {
	return testAccAccountDimensionConfig() + `
resource "udoma_account" "test" {
  number   = 128
  name     = "Test Account"
  type     = "asset"
  currency = "USD"
	cadence  = "monthly_cadence"
}
`
}
