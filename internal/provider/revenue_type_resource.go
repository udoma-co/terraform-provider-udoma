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
	_ resource.ResourceWithConfigure   = &RevenueType{}
	_ resource.ResourceWithImportState = &RevenueType{}
)

func NewRevenueType() resource.Resource {
	return &RevenueType{}
}

type RevenueType struct {
	client *client.UdomaClient
}

type RevenueTypeModel struct {
	ID            types.String `tfsdk:"id"`
	CreatedAt     types.Int64  `tfsdk:"created_at"`
	UpdatedAt     types.Int64  `tfsdk:"updated_at"`
	Name          types.String `tfsdk:"name"`
	Description   types.String `tfsdk:"description"`
	Scope         types.String `tfsdk:"scope"`
	SystemDefault types.String `tfsdk:"system_default"`
	IsCatalogItem types.Bool   `tfsdk:"is_catalog_item"`
}

func (rt *RevenueType) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_revenue_type"
}

func (rt *RevenueType) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents a revenue type used for categorizing income on accounts.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The unique identifier for the revenue type.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The date and time the revenue type was created.",
			},
			"updated_at": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The date and time the revenue type was last modified.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the revenue type.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "An optional description of the revenue type.",
			},
			"scope": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The scope where this revenue type is used.",
				Validators: []validator.String{
					stringvalidator.OneOf(stringSlice(api.AllowedRevenueTypeScopeEnumEnumValues)...),
				},
			},
			"system_default": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The system default category of this revenue type.",
				Validators: []validator.String{
					stringvalidator.OneOf(stringSlice(api.AllowedRevenueTypeSystemDefaultEnumEnumValues)...),
				},
			},
			"is_catalog_item": schema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether this entity is represented as a catalog item.",
			},
		},
	}
}

func (rt *RevenueType) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	rt.client = cl
}

func (rt *RevenueType) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RevenueTypeModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Revenue Type Request",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newRevenueType, _, err := rt.client.GetApi().CreateRevenueType(ctx).CreateOrUpdateRevenueTypeRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Revenue Type",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	diags = plan.fromAPI(newRevenueType)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	resp.Diagnostics.Append(
		resp.State.Set(ctx, plan)...,
	)
}

func (rt *RevenueType) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RevenueTypeModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	revenueType, httpResp, err := rt.client.GetApi().GetRevenueType(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Revenue Type",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	diags = state.fromAPI(revenueType)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	resp.Diagnostics.Append(
		resp.State.Set(ctx, &state)...,
	)
}

func (rt *RevenueType) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan RevenueTypeModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Revenue Type",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	updatedRevenueType, _, err := rt.client.GetApi().UpdateRevenueType(ctx, plan.ID.ValueString()).CreateOrUpdateRevenueTypeRequest(updateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Revenue Type",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	diags = plan.fromAPI(updatedRevenueType)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	resp.Diagnostics.Append(
		resp.State.Set(ctx, plan)...,
	)
}

func (rt *RevenueType) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RevenueTypeModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := rt.client.GetApi().DeleteRevenueType(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Revenue Type",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (rt *RevenueType) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *RevenueTypeModel) fromAPI(revenueType *api.RevenueType) (diags diag.Diagnostics) {
	if revenueType == nil {
		diags.AddError("API response is nil", "The API response provided to RevenueTypeModel.fromAPI is nil")
		return
	}

	model.ID = types.StringValue(revenueType.Id)
	model.IsCatalogItem = types.BoolPointerValue(revenueType.IsCatalogItem)
	model.CreatedAt = types.Int64Value(revenueType.CreatedAt)
	model.UpdatedAt = types.Int64Value(revenueType.UpdatedAt)
	model.Name = types.StringValue(revenueType.Name)
	model.Description = omittableStringValue(revenueType.Description, model.Description)
	model.Scope = types.StringValue(string(revenueType.Scope))
	model.SystemDefault = types.StringValue(string(revenueType.SystemDefault))

	return
}

func (model *RevenueTypeModel) toAPIRequest() (api.CreateOrUpdateRevenueTypeRequest, error) {
	req := api.CreateOrUpdateRevenueTypeRequest{
		Name:          model.Name.ValueString(),
		Description:   model.Description.ValueStringPointer(),
		Scope:         api.RevenueTypeScopeEnum(model.Scope.ValueString()),
		SystemDefault: api.RevenueTypeSystemDefaultEnum(model.SystemDefault.ValueString()),
	}

	return req, nil
}
