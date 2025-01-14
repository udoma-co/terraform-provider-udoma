package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	api "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"
)

var (
	_ resource.ResourceWithConfigure   = &AccountDimensionValue{}
	_ resource.ResourceWithImportState = &AccountDimensionValue{}
)

func NewAccountDimensionValue() resource.Resource {
	return &AccountDimensionValue{}
}

type AccountDimensionValue struct {
	client *client.UdomaClient
}

type AccountDimensionValueModel struct {
	ID           types.String `tfsdk:"id"`
	ParentId     types.String `tfsdk:"parent_id"`
	Value        types.Int32  `tfsdk:"value"`
	DimensionRef types.String `tfsdk:"dimension_ref"`
	Alias        types.String `tfsdk:"alias"`
}

func (faq *AccountDimensionValue) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account_dimension_value"
}

func (faq *AccountDimensionValue) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents a financial account",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The unique identifier for the account dimension value",
			},
			"parent_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The parent account dimension value",
			},
			"value": schema.Int32Attribute{
				Required:            true,
				MarkdownDescription: "The value of the account dimension",
			},
			"dimension_ref": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The reference to the account dimension",
			},
			"alias": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The alias of the account dimension value",
			},
		},
	}
}

func (value *AccountDimensionValue) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	value.client = cl
}

func (value *AccountDimensionValue) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Retrieve values from plan
	var plan AccountDimensionValueModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Account Dimension Value Request",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newValue, _, err := value.client.GetApi().CreateAccountDimensionValue(ctx, plan.DimensionRef.ValueString()).CreateOrUpdateAccountDimensionValueRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Account Dimension Value",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newValue); err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Account Dimension Value",
			"Could not process API response, unexpected error: "+err.Error(),
		)
		return
	}

	fmt.Println("CREATE", plan)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (value *AccountDimensionValue) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state AccountDimensionValueModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	getValue, httpResp, err := value.client.GetApi().GetAccountDimensionValue(ctx, state.DimensionRef.ValueString(), state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Account Dimension Value",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := state.fromAPI(getValue); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Account Dimension Value",
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

func (value *AccountDimensionValue) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan AccountDimensionValueModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Account Dimension Value Request",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	updatedValue, _, err := value.client.GetApi().UpdateAccountDimensionValue(ctx, plan.DimensionRef.ValueString(), plan.ID.ValueString()).CreateOrUpdateAccountDimensionValueRequest(updateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Account Dimension Value",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(updatedValue); err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Account Dimension Value",
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

func (value *AccountDimensionValue) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	// Retrieve values from state
	var state AccountDimensionValueModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := value.client.GetApi().DeleteAccountDimensionValue(ctx, state.DimensionRef.ValueString(), state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Account Dimension Value Entry",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (value *AccountDimensionValue) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *AccountDimensionValueModel) fromAPI(value *api.AccountDimensionValue) error {

	if value == nil {
		return fmt.Errorf("account dimension value is nil")
	}

	fmt.Println("API response:", value)

	model.ID = types.StringValue(value.Id)
	model.ParentId = types.StringPointerValue(value.ParentId)
	model.Value = types.Int32PointerValue(value.Value)
	model.DimensionRef = types.StringValue(value.DimensionRef)
	model.Alias = types.StringPointerValue(value.Alias)

	return nil
}

func (model *AccountDimensionValueModel) toAPIRequest() (api.CreateOrUpdateAccountDimensionValueRequest, error) {

	value := api.CreateOrUpdateAccountDimensionValueRequest{
		Id:       model.ID.ValueString(),
		ParentId: model.ParentId.ValueStringPointer(),
		Value:    model.Value.ValueInt32Pointer(),
	}

	return value, nil
}
