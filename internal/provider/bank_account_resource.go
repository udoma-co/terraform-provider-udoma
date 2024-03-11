package provider

import (
	"context"
	"fmt"

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
	_ resource.ResourceWithConfigure   = &FAQ{}
	_ resource.ResourceWithImportState = &FAQ{}
)

func NewBankAccount() resource.Resource {
	return &BankAccount{}
}

type BankAccount struct {
	client *client.UdomaClient
}

type BankAccountModel struct {
	ID             types.String `tfsdk:"id"`
	ExternalID     types.String `tfsdk:"external_id"`
	ExternalSource types.String `tfsdk:"external_source"`
	CreatedAt      types.Int64  `tfsdk:"created_at"`
	UpdatedAt      types.Int64  `tfsdk:"updated_at"`
	AccountHolder  types.String `tfsdk:"account_holder"`
	Iban           types.String `tfsdk:"iban"`
	Bic            types.String `tfsdk:"bic"`
	BankName       types.String `tfsdk:"bank_name"`
	Description    types.String `tfsdk:"description"`
}

func (faq *BankAccount) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bank_account"
}

func (faq *BankAccount) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents a bank account",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the bank account",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"external_id": schema.StringAttribute{
				Optional:    true,
				Description: "Optional external ID, in case acount was created via backend integration",
			},
			"external_source": schema.StringAttribute{
				Optional:    true,
				Description: "Optional external source, in case acount was created via backend integration",
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the bank account was created",
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the bank account was last modified",
			},
			"account_holder": schema.StringAttribute{
				Required:    true,
				Description: "The name of the account holder",
			},
			"iban": schema.StringAttribute{
				Required:    true,
				Description: "The IBAN of the bank account",
			},
			"bic": schema.StringAttribute{
				Optional:    true,
				Description: "The BIC of the bank account",
			},
			"bank_name": schema.StringAttribute{
				Optional:    true,
				Description: "The name of the bank",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "An optional user friendly label, used to identify the account",
			},
		},
	}
}

func (bank_account *BankAccount) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	bank_account.client = cl
}

func (bank_account *BankAccount) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Retrieve values from plan
	var plan BankAccountModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Bank Account Request",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newBankAccount, _, err := bank_account.client.GetApi().CreateBankAccount(ctx).CreateOrUpdateBankAccountRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Bank Account Entry",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newBankAccount); err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Bank Account Entry",
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

func (bank_account *BankAccount) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state BankAccountModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	bankAccount, httpResp, err := bank_account.client.GetApi().GetBankAccount(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Bank Account Entry",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := state.fromAPI(bankAccount); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Bank Account Entry",
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

func (bank_account *BankAccount) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan BankAccountModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Bank Account Entry",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newBankAccount, _, err := bank_account.client.GetApi().UpdateBankAccount(ctx, plan.ID.ValueString()).CreateOrUpdateBankAccountRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Bank Account Entry",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newBankAccount); err != nil {
		resp.Diagnostics.AddError(
			"Error Updating FAQ Entry",
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

func (bank_account *BankAccount) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	// Retrieve values from state
	var state BankAccountModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := bank_account.client.GetApi().DeleteBankAccount(ctx, state.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Bank Account Entry",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}
}

func (bank_account *BankAccount) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *BankAccountModel) fromAPI(bankAccount *api.BankAccount) error {

	if bankAccount == nil {
		return fmt.Errorf("bank account is nil")
	}

	model.ID = types.StringValue(sdp(bankAccount.Id))
	model.ExternalID = types.StringValue(sdp(bankAccount.ExternalId))
	model.ExternalSource = types.StringValue(sdp(bankAccount.ExternalSource))
	model.AccountHolder = types.StringValue(sdp(bankAccount.AccountHolder))
	model.Iban = types.StringValue(sdp(bankAccount.Iban))
	model.Bic = types.StringValue(sdp(bankAccount.Bic))
	model.BankName = types.StringValue(sdp(bankAccount.BankName))
	model.Description = types.StringValue(sdp(bankAccount.Description))
	model.CreatedAt = types.Int64Value(idp(bankAccount.CreatedAt))
	model.UpdatedAt = types.Int64Value(idp(bankAccount.UpdatedAt))

	return nil
}

func (model *BankAccountModel) toAPIRequest() (api.CreateOrUpdateBankAccountRequest, error) {

	bankAccount := api.CreateOrUpdateBankAccountRequest{
		AccountHolder: model.AccountHolder.ValueString(),
		Iban:          model.Iban.ValueString(),
		Bic:           model.Bic.ValueStringPointer(),
		BankName:      model.BankName.ValueStringPointer(),
		Description:   model.Description.ValueStringPointer(),
	}

	return bankAccount, nil
}
