package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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
	_ resource.ResourceWithConfigure   = &AllocationKey{}
	_ resource.ResourceWithImportState = &AllocationKey{}
)

func NewAllocationKey() resource.Resource {
	return &AllocationKey{}
}

type AllocationKey struct {
	client *client.UdomaClient
}

type AllocationKeyModel struct {
	ID            types.String `tfsdk:"id"`
	CreatedAt     types.Int64  `tfsdk:"created_at"`
	UpdatedAt     types.Int64  `tfsdk:"updated_at"`
	Name          types.String `tfsdk:"name"`
	Description   types.String `tfsdk:"description"`
	Formula       types.String `tfsdk:"formula"`
	IsCatalogItem types.Bool   `tfsdk:"is_catalog_item"`
}

func (ak *AllocationKey) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_allocation_key"
}

func (ak *AllocationKey) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents how a shared cost is distributed across properties or units.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The unique identifier for the allocation key.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The date and time the allocation key was created.",
			},
			"updated_at": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The date and time the allocation key was last modified.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Human-readable name of the allocation key, for example `by area`.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Optional explanation of the allocation key.",
			},
			"formula": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Pseudocode or mathematical expression defining how costs are distributed.",
			},
			"is_catalog_item": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether this entity is represented as a catalog item.",
			},
		},
	}
}

func (ak *AllocationKey) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	ak.client = cl
}

func (ak *AllocationKey) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan AllocationKeyModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Allocation Key Request",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newAllocationKey, _, err := ak.client.GetApi().CreateAllocationKey(ctx).CreateOrUpdateAllocationKeyRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Allocation Key",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	diags = plan.fromAPI(newAllocationKey)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (ak *AllocationKey) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state AllocationKeyModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	allocationKey, httpResp, err := ak.client.GetApi().GetAllocationKey(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Allocation Key",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	diags = state.fromAPI(allocationKey)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (ak *AllocationKey) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan AllocationKeyModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Allocation Key",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	updatedAllocationKey, _, err := ak.client.GetApi().UpdateAllocationKey(ctx, plan.ID.ValueString()).CreateOrUpdateAllocationKeyRequest(updateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Allocation Key",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	diags = plan.fromAPI(updatedAllocationKey)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (ak *AllocationKey) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state AllocationKeyModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := ak.client.GetApi().DeleteAllocationKey(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Allocation Key",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (ak *AllocationKey) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *AllocationKeyModel) fromAPI(allocationKey *api.AllocationKey) (diags diag.Diagnostics) {
	if allocationKey == nil {
		diags.AddError("API response is nil", "The API response provided to AllocationKeyModel.fromAPI is nil")
		return
	}

	model.ID = types.StringValue(allocationKey.Id)
	model.CreatedAt = types.Int64Value(allocationKey.CreatedAt)
	model.UpdatedAt = types.Int64Value(allocationKey.UpdatedAt)
	model.Name = types.StringValue(allocationKey.Name)
	model.Description = omittableStringValue(allocationKey.Description, model.Description)
	model.Formula = omittableStringValue(allocationKey.Formula, model.Formula)
	model.IsCatalogItem = types.BoolPointerValue(allocationKey.IsCatalogItem)

	return
}

func (model *AllocationKeyModel) toAPIRequest() (api.CreateOrUpdateAllocationKeyRequest, error) {
	req := api.CreateOrUpdateAllocationKeyRequest{
		Name:        model.Name.ValueString(),
		Description: model.Description.ValueStringPointer(),
		Formula:     model.Formula.ValueStringPointer(),
	}

	return req, nil
}
