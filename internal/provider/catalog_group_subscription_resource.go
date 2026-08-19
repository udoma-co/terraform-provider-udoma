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
	_ resource.Resource                = &CatalogGroupSubscription{}
	_ resource.ResourceWithConfigure   = &CatalogGroupSubscription{}
	_ resource.ResourceWithImportState = &CatalogGroupSubscription{}
)

func NewCatalogGroupSubscription() resource.Resource {
	return &CatalogGroupSubscription{}
}

type CatalogGroupSubscription struct {
	client *client.UdomaClient
}

type catalogGroupSubscriptionModel struct {
	ID          types.String   `tfsdk:"id"`
	CreatedAt   types.Int64    `tfsdk:"created_at"`
	UpdatedAt   types.Int64    `tfsdk:"updated_at"`
	Code        types.String   `tfsdk:"code"`
	ExcludeList []types.String `tfsdk:"exclude_list"`
}

func (r *CatalogGroupSubscription) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_catalog_group_subscription"
}

func (r *CatalogGroupSubscription) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents a catalog group subscription",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the catalog group subscription.",
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
				Description: "Code of the catalog group being subscribed to.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
			},
			"exclude_list": schema.ListAttribute{
				Optional:    true,
				Description: "Catalog item codes to exclude from the group subscription.",
				ElementType: types.StringType,
			},
		},
	}
}

func (r *CatalogGroupSubscription) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CatalogGroupSubscription) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan catalogGroupSubscriptionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError("Error Creating Catalog Group Subscription", "Could not create API request, unexpected error: "+err.Error())
		return
	}

	groupSubscription, _, err := r.client.GetApi().CreateGroupSubscription(ctx).CreateOrUpdateCatalogGroupSubscription(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Creating Catalog Group Subscription", "Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err))
		return
	}

	if err := plan.fromAPI(groupSubscription); err != nil {
		resp.Diagnostics.AddError("Error Creating Catalog Group Subscription", "Could not process API response, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *CatalogGroupSubscription) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state catalogGroupSubscriptionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	groupSubscription, httpResp, err := r.client.GetApi().GetGroupSubscription(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error Reading Catalog Group Subscription", "Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err))
		return
	}

	if err := state.fromAPI(groupSubscription); err != nil {
		resp.Diagnostics.AddError("Error Reading Catalog Group Subscription", "Could not process API response, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *CatalogGroupSubscription) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan catalogGroupSubscriptionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError("Error Updating Catalog Group Subscription", "Could not create API request, unexpected error: "+err.Error())
		return
	}

	groupSubscription, _, err := r.client.GetApi().UpdateGroupSubscription(ctx, plan.ID.ValueString()).CreateOrUpdateCatalogGroupSubscription(updateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Updating Catalog Group Subscription", "Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err))
		return
	}

	if err := plan.fromAPI(groupSubscription); err != nil {
		resp.Diagnostics.AddError("Error Updating Catalog Group Subscription", "Could not process API response, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *CatalogGroupSubscription) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state catalogGroupSubscriptionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.client.GetApi().DeleteGroupSubscription(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error Deleting Catalog Group Subscription", "Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err))
	}
}

func (r *CatalogGroupSubscription) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *catalogGroupSubscriptionModel) fromAPI(groupSubscription *api.CatalogGroupSubscription) error {
	if groupSubscription == nil {
		return fmt.Errorf("catalog group subscription is nil")
	}

	model.ID = types.StringValue(groupSubscription.Id)
	model.CreatedAt = types.Int64Value(groupSubscription.CreatedAt)
	model.UpdatedAt = types.Int64Value(groupSubscription.UpdatedAt)
	model.Code = types.StringValue(groupSubscription.Code)

	if len(groupSubscription.ExcludeList) > 0 {
		model.ExcludeList = make([]types.String, len(groupSubscription.ExcludeList))
		for i := range groupSubscription.ExcludeList {
			model.ExcludeList[i] = types.StringValue(groupSubscription.ExcludeList[i])
		}
	} else {
		if model.ExcludeList != nil {
			model.ExcludeList = make([]types.String, 0)
		}
	}

	return nil
}

func (model *catalogGroupSubscriptionModel) toAPIRequest() (api.CreateOrUpdateCatalogGroupSubscription, error) {
	subscription := api.CreateOrUpdateCatalogGroupSubscription{Code: model.Code.ValueString()}

	if len(model.ExcludeList) > 0 {
		subscription.ExcludeList = make([]string, len(model.ExcludeList))
		for i := range model.ExcludeList {
			subscription.ExcludeList[i] = model.ExcludeList[i].ValueString()
		}
	}

	return subscription, nil
}
