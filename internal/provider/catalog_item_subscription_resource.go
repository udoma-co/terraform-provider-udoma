package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	api "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"
)

var (
	_ resource.Resource                = &CatalogItemSubscription{}
	_ resource.ResourceWithConfigure   = &CatalogItemSubscription{}
	_ resource.ResourceWithImportState = &CatalogItemSubscription{}
)

func NewCatalogItemSubscription() resource.Resource {
	return &CatalogItemSubscription{}
}

type CatalogItemSubscription struct {
	client *client.UdomaClient
}

type catalogItemSubscriptionPatchModel struct {
	PatchRef types.String `tfsdk:"patch_ref"`
	Format   types.String `tfsdk:"format"`
	Patch    types.String `tfsdk:"patch"`
}

type catalogItemSubscriptionModel struct {
	ID        types.String                        `tfsdk:"id"`
	CreatedAt types.Int64                         `tfsdk:"created_at"`
	UpdatedAt types.Int64                         `tfsdk:"updated_at"`
	Code      types.String                        `tfsdk:"code"`
	Patches   []catalogItemSubscriptionPatchModel `tfsdk:"patches"`
}

func (r *CatalogItemSubscription) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_catalog_item_subscription"
}

func (r *CatalogItemSubscription) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents a catalog item subscription",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the catalog item subscription.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the subscription was created.",
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the subscription was last updated.",
			},
			"code": schema.StringAttribute{
				Required:    true,
				Description: "Code of the catalog item being subscribed to.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
			},
			"patches": schema.ListNestedAttribute{
				Optional:    true,
				Description: "Optional patches that override attributes of the subscribed catalog item.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"patch_ref": schema.StringAttribute{Optional: true, Description: "Reference key identifying a patch target."},
						"format": schema.StringAttribute{
							Optional:    true,
							Description: "Patch format (for example JSON patch or unified diff).",
							Validators: []validator.String{
								stringvalidator.OneOf(stringSlice(api.AllowedPatchFormatEnumValues)...),
							},
						},
						"patch": schema.StringAttribute{Optional: true, Description: "Patch payload."},
					},
				},
			},
		},
	}
}

func (r *CatalogItemSubscription) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	cl, ok := req.ProviderData.(*client.UdomaClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Data Type",
			fmt.Sprintf("Expected *client.UdomaClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = cl
}

func (r *CatalogItemSubscription) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan catalogItemSubscriptionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError("Error Creating Catalog Item Subscription", "Could not create API request, unexpected error: "+err.Error())
		return
	}

	itemSubscription, _, err := r.client.GetApi().CreateCatalogItemSubscription(ctx).CreateOrUpdateCatalogItemSubscription(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Creating Catalog Item Subscription", "Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err))
		return
	}

	if err := plan.fromAPI(itemSubscription); err != nil {
		resp.Diagnostics.AddError("Error Creating Catalog Item Subscription", "Could not process API response, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *CatalogItemSubscription) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state catalogItemSubscriptionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	itemSubscription, httpResp, err := r.client.GetApi().GetCatalogItemSubscription(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error Reading Catalog Item Subscription", "Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err))
		return
	}

	if err := state.fromAPI(itemSubscription); err != nil {
		resp.Diagnostics.AddError("Error Reading Catalog Item Subscription", "Could not process API response, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *CatalogItemSubscription) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan catalogItemSubscriptionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError("Error Updating Catalog Item Subscription", "Could not create API request, unexpected error: "+err.Error())
		return
	}

	itemSubscription, _, err := r.client.GetApi().UpdateCatalogItemSubscription(ctx, plan.ID.ValueString()).CreateOrUpdateCatalogItemSubscription(updateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Updating Catalog Item Subscription", "Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err))
		return
	}

	if err := plan.fromAPI(itemSubscription); err != nil {
		resp.Diagnostics.AddError("Error Updating Catalog Item Subscription", "Could not process API response, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *CatalogItemSubscription) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state catalogItemSubscriptionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.client.GetApi().DeleteCatalogItemSubscription(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error Deleting Catalog Item Subscription", "Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err))
	}
}

func (r *CatalogItemSubscription) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *catalogItemSubscriptionModel) fromAPI(itemSubscription *api.CatalogItemSubscription) error {
	if itemSubscription == nil {
		return fmt.Errorf("catalog item subscription is nil")
	}

	model.ID = types.StringValue(itemSubscription.Id)
	model.CreatedAt = types.Int64Value(itemSubscription.CreatedAt)
	model.UpdatedAt = types.Int64Value(itemSubscription.UpdatedAt)
	model.Code = types.StringValue(itemSubscription.Code)

	if itemSubscription.Patches != nil {
		model.Patches = make([]catalogItemSubscriptionPatchModel, len(itemSubscription.Patches))
		for i := range itemSubscription.Patches {
			patch := itemSubscription.Patches[i]
			model.Patches[i] = catalogItemSubscriptionPatchModel{
				PatchRef: types.StringPointerValue(patch.PatchRef),
				Patch:    types.StringPointerValue(patch.Patch),
			}
			if patch.Format != nil {
				model.Patches[i].Format = types.StringValue(string(*patch.Format))
			} else {
				model.Patches[i].Format = types.StringNull()
			}
		}
	} else {
		if model.Patches != nil {
			model.Patches = make([]catalogItemSubscriptionPatchModel, 0)
		} else {
			model.Patches = nil
		}
	}

	return nil
}

func (model *catalogItemSubscriptionModel) toAPIRequest() (api.CreateOrUpdateCatalogItemSubscription, error) {
	subscription := api.CreateOrUpdateCatalogItemSubscription{Code: model.Code.ValueString()}

	if len(model.Patches) > 0 {
		subscription.Patches = make([]api.CatalogItemPatch, len(model.Patches))
		for i, patch := range model.Patches {
			subscription.Patches[i] = api.CatalogItemPatch{
				PatchRef: patch.PatchRef.ValueStringPointer(),
				Patch:    patch.Patch.ValueStringPointer(),
			}
			if !patch.Format.IsNull() && !patch.Format.IsUnknown() {
				format := api.PatchFormat(patch.Format.ValueString())
				subscription.Patches[i].Format = &format
			}
		}
	}

	return subscription, nil
}
