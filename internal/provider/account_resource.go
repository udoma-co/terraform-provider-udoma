package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	api "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"
)

var (
	_ resource.ResourceWithConfigure   = &Account{}
	_ resource.ResourceWithImportState = &Account{}
)

func NewAccount() resource.Resource {
	return &Account{}
}

type Account struct {
	client *client.UdomaClient
}

type AccountModel struct {
	ID         types.String `tfsdk:"id"`
	CreatedAt  types.Int64  `tfsdk:"created_at"`
	UpdatedAt  types.Int64  `tfsdk:"updated_at"`
	Number     types.Int32  `tfsdk:"number"`
	Name       types.String `tfsdk:"name"`
	Type       types.String `tfsdk:"type"`
	IsBalance  types.Bool   `tfsdk:"is_balance"`
	Currency   types.String `tfsdk:"currency"`
	Dimensions types.List   `tfsdk:"dimensions"`
}

func (faq *Account) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account"
}

func (faq *Account) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents a financial account",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The unique identifier for the financial account",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The date and time the account was created",
			},
			"updated_at": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The date and time the account was last modified",
			},
			"number": schema.Int32Attribute{
				Required:            true,
				MarkdownDescription: "The unique account number, manually set",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the account",
			},
			"type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The type of the account",
			},
			"is_balance": schema.BoolAttribute{
				Required:            true,
				MarkdownDescription: "Whether the account is a balance account",
			},
			"currency": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The currency of the account",
			},
			"dimensions": schema.ListAttribute{
				Optional:    true,
				Description: "The sub dimensions of the account.",
				ElementType: types.StringType,
			},
		},
	}
}

func (account *Account) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	account.client = cl
}

func (account *Account) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Retrieve values from plan
	var plan AccountModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Financial Account Request",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newAccount, _, err := account.client.GetApi().CreateFinancialAccount(ctx).CreateOrUpdateFinancialAccountRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Financial Account Entry",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	diags = plan.fromAPI(newAccount)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	// Set state to fully populated data
	resp.Diagnostics.Append(
		resp.State.Set(ctx, plan)...,
	)
}

func (account *Account) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state AccountModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	financialAccount, httpResp, err := account.client.GetApi().GetFinancialAccount(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Financial Account Entry",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	diags = state.fromAPI(financialAccount)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	// Set refreshed state
	resp.Diagnostics.Append(
		resp.State.Set(ctx, &state)...,
	)

}

func (account *Account) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan AccountModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Financial Account",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	updatedAccount, _, err := account.client.GetApi().UpdateFinancialAccount(ctx, plan.ID.ValueString()).CreateOrUpdateFinancialAccountRequest(updateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Financial Account",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	diags = plan.fromAPI(updatedAccount)
	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}

	// Set state to fully populated data
	resp.Diagnostics.Append(
		resp.State.Set(ctx, plan)...,
	)
}

func (account *Account) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	// Retrieve values from state
	var state AccountModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := account.client.GetApi().DeleteFinancialAccount(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Financial Account Entry",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (account *Account) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *AccountModel) fromAPI(account *api.FinancialAccount) (diags diag.Diagnostics) {

	if account == nil {
		diags.AddError("API response is nil", "The API response provided to AccountModel.fromAPIResponse is nil")
		return
	}

	model.ID = types.StringValue(account.Id)
	model.Number = types.Int32Value(account.Number)
	model.Name = types.StringValue(account.Name)
	model.Currency = types.StringValue(account.Currency)
	model.CreatedAt = types.Int64Value(account.CreatedAt)
	model.UpdatedAt = types.Int64Value(account.UpdatedAt)
	model.Type = types.StringValue(string(account.Type))
	model.IsBalance = omittableBooleanValue(account.IsBalance, model.IsBalance)

	dimensionIDs := make([]string, len(account.Dimensions))
	for i := range account.Dimensions {
		dimensionIDs[i] = account.Dimensions[i].Id
	}

	if model.Dimensions.IsNull() && len(dimensionIDs) > 0 {
		model.Dimensions, diags = types.ListValue(types.StringType, stringSliceToValueList(dimensionIDs))
		if diags.HasError() {
			return
		}
	}

	return nil
}

func (model *AccountModel) toAPIRequest() (api.CreateOrUpdateFinancialAccountRequest, error) {

	account := api.CreateOrUpdateFinancialAccountRequest{
		Number:     model.Number.ValueInt32(),
		Name:       model.Name.ValueString(),
		Type:       api.AccountTypesEnum(model.Type.ValueString()),
		IsBalance:  model.IsBalance.ValueBoolPointer(),
		Currency:   model.Currency.ValueString(),
		Dimensions: modelListToStringSlice(model.Dimensions),
	}

	return account, nil
}
