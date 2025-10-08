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
	_ resource.ResourceWithConfigure   = &AccountDimension{}
	_ resource.ResourceWithImportState = &AccountDimension{}
)

func NewAccountDimension() resource.Resource {
	return &AccountDimension{}
}

type AccountDimension struct {
	client *client.UdomaClient
}

type AccountDimensionModel struct {
	ID              types.String `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	Description     types.String `tfsdk:"description"`
	ParentDimension types.String `tfsdk:"parent_dimension"`
	RefType         types.String `tfsdk:"ref_type"`
	Required        types.Bool   `tfsdk:"required"`
	PadToSize       types.Int32  `tfsdk:"pad_to_size"`
	ValueGenerator  types.String `tfsdk:"value_generator"`
}

func (faq *AccountDimension) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account_dimension"
}

func (faq *AccountDimension) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents a financial account",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The unique identifier for the account dimension",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the account",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The description of the account",
			},
			"parent_dimension": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The parent dimension of the account",
			},
			"ref_type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The reference type of the account",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(25),
					stringvalidator.OneOfCaseInsensitive(stringSlice(api.AllowedAccountDimensionReferenceTypeEnumEnumValues)...),
				},
			},
			"required": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "The required status of the account",
			},
			"pad_to_size": schema.Int32Attribute{
				Optional:            true,
				MarkdownDescription: "The pad to size of the account",
			},
			"value_generator": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The value generator of the account",
			},
		},
	}
}

func (dimension *AccountDimension) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	dimension.client = cl
}

func (dimension *AccountDimension) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Retrieve values from plan
	var plan AccountDimensionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Account Dimension",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newDimension, _, err := dimension.client.GetApi().CreateAccountDimension(ctx).CreateOrUpdateAccountDimensionRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Account Dimension",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newDimension); err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Account Dimension",
			"Could not process API response, unexpected error: "+err.Error(),
		)
		return
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (dimension *AccountDimension) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state AccountDimensionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	accountDimension, httpResp, err := dimension.client.GetApi().GetAccountDimension(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Account Dimension",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := state.fromAPI(accountDimension); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Account Dimension",
			"Could not process API response, unexpected error: "+err.Error(),
		)
		return
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (dimension *AccountDimension) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan AccountDimensionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Account Dimension",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	updatedDimension, _, err := dimension.client.GetApi().UpdateAccountDimension(ctx, plan.ID.ValueString()).CreateOrUpdateAccountDimensionRequest(updateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Account Dimension",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(updatedDimension); err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Account Dimension",
			"Could not process API response, unexpected error: "+err.Error(),
		)
		return
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (dimension *AccountDimension) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	// Retrieve values from state
	var state AccountDimensionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := dimension.client.GetApi().DeleteAccountDimension(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Account Dimension",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (value *AccountDimension) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *AccountDimensionModel) fromAPI(dimension *api.AccountDimension) error {

	if dimension == nil {
		return fmt.Errorf("account dimension is nil")
	}

	model.ID = types.StringValue(dimension.Id)
	model.Name = types.StringValue(dimension.Name)
	model.Description = types.StringPointerValue(dimension.Description)
	model.PadToSize = types.Int32PointerValue(dimension.PadToSize)
	model.RefType = types.StringValue(string(dimension.RefType))
	model.ParentDimension = types.StringPointerValue(dimension.ParentDimensionRef)
	model.Required = types.BoolPointerValue(dimension.Required)
	model.ValueGenerator = types.StringPointerValue(dimension.ValueGenerator)

	return nil
}

func (model *AccountDimensionModel) toAPIRequest() (api.CreateOrUpdateAccountDimensionRequest, error) {

	dimension := api.CreateOrUpdateAccountDimensionRequest{
		Name:               model.Name.ValueString(),
		Description:        model.Description.ValueStringPointer(),
		ParentDimensionRef: model.ParentDimension.ValueStringPointer(),
		PadToSize:          model.PadToSize.ValueInt32Pointer(),
		RefType:            api.AccountDimensionReferenceTypeEnum(model.RefType.ValueString()),
		Required:           model.Required.ValueBoolPointer(),
		ValueGenerator:     model.ValueGenerator.ValueStringPointer(),
	}

	return dimension, nil
}
