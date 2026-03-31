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
	_ resource.ResourceWithConfigure   = &CostType{}
	_ resource.ResourceWithImportState = &CostType{}
)

func NewCostType() resource.Resource {
	return &CostType{}
}

type CostType struct {
	client *client.UdomaClient
}

type CostTypeModel struct {
	ID          types.String `tfsdk:"id"`
	CreatedAt   types.Int64  `tfsdk:"created_at"`
	UpdatedAt   types.Int64  `tfsdk:"updated_at"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	IsFixed     types.Bool   `tfsdk:"is_fixed"`
	Category    types.String `tfsdk:"category"`
}

func (ct *CostType) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cost_type"
}

func (ct *CostType) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents a cost type used for categorizing costs in operating cost statements",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The unique identifier for the cost type",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The date and time the cost type was created",
			},
			"updated_at": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The date and time the cost type was last modified",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the cost type",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "An optional description of the cost type",
			},
			"is_fixed": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "If true, this cost is fixed (not dependent on actual usage). If false, the cost is variable.",
			},
			"category": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The category of the cost type. Must be one of: operating_costs, maintenance, administrative_costs, utilities, insurance, taxes_fees, cap_ex, others",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"operating_costs",
						"maintenance",
						"administrative_costs",
						"utilities",
						"insurance",
						"taxes_fees",
						"cap_ex",
						"others",
					),
				},
			},
		},
	}
}

func (ct *CostType) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	ct.client = cl
}

func (ct *CostType) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CostTypeModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Cost Type Request",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newCostType, _, err := ct.client.GetApi().CreateCostType(ctx).CreateOrUpdateCostTypeRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Cost Type",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	diags = plan.fromAPI(newCostType)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	resp.Diagnostics.Append(
		resp.State.Set(ctx, plan)...,
	)
}

func (ct *CostType) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CostTypeModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	costType, httpResp, err := ct.client.GetApi().GetCostType(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Cost Type",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	diags = state.fromAPI(costType)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	resp.Diagnostics.Append(
		resp.State.Set(ctx, &state)...,
	)
}

func (ct *CostType) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan CostTypeModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Cost Type",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	updatedCostType, _, err := ct.client.GetApi().UpdateCostType(ctx, plan.ID.ValueString()).CreateOrUpdateCostTypeRequest(updateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Cost Type",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	diags = plan.fromAPI(updatedCostType)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	resp.Diagnostics.Append(
		resp.State.Set(ctx, plan)...,
	)
}

func (ct *CostType) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CostTypeModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := ct.client.GetApi().DeleteCostType(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Cost Type",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (ct *CostType) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *CostTypeModel) fromAPI(costType *api.CostType) (diags diag.Diagnostics) {
	if costType == nil {
		diags.AddError("API response is nil", "The API response provided to CostTypeModel.fromAPI is nil")
		return
	}

	model.ID = types.StringValue(costType.Id)
	model.CreatedAt = types.Int64Value(costType.CreatedAt)
	model.UpdatedAt = types.Int64Value(costType.UpdatedAt)
	model.Name = types.StringValue(costType.Name)
	model.Category = types.StringValue(string(costType.Category))

	if costType.Description != nil {
		model.Description = types.StringValue(*costType.Description)
	} else {
		model.Description = types.StringNull()
	}

	if costType.IsFixed != nil {
		model.IsFixed = types.BoolValue(*costType.IsFixed)
	} else {
		model.IsFixed = types.BoolNull()
	}

	return
}

func (model *CostTypeModel) toAPIRequest() (api.CreateOrUpdateCostTypeRequest, error) {
	req := api.CreateOrUpdateCostTypeRequest{
		Name:     model.Name.ValueString(),
		Category: api.CostTypeCategoryEnum(model.Category.ValueString()),
	}

	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		desc := model.Description.ValueString()
		req.Description = &desc
	}

	if !model.IsFixed.IsNull() && !model.IsFixed.IsUnknown() {
		isFixed := model.IsFixed.ValueBool()
		req.IsFixed = &isFixed
	}

	return req, nil
}
