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
	_ resource.ResourceWithConfigure   = &InvoiceDataMapper{}
	_ resource.ResourceWithImportState = &InvoiceDataMapper{}
)

func NewInvoiceDataMapper() resource.Resource {
	return &InvoiceDataMapper{}
}

type InvoiceDataMapper struct {
	client *client.UdomaClient
}

type InvoiceDataMapperModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Script      types.String `tfsdk:"script"`
	Entrypoint  types.String `tfsdk:"entrypoint"`
}

func (idm *InvoiceDataMapper) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_invoice_data_mapper"
}

func (idm *InvoiceDataMapper) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents an invoice data mapper",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the invoice data mapper",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Some name used to identify the invoice mapper. No unique requirement.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(64),
				},
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "A text description of what the invoice data mapper does.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(512),
				},
			},
			"script": schema.StringAttribute{
				Required:    true,
				Description: "The script that executes to remap invoice analyses",
			},
			"entrypoint": schema.StringAttribute{
				Required:    true,
				Description: "When to run the script, either before or after data processing",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(32),
					stringvalidator.OneOfCaseInsensitive(stringSlice(api.AllowedInvoiceDataMapperEntrypointEnumEnumValues)...),
				},
			},
		},
	}
}

func (idm *InvoiceDataMapper) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	idm.client = cl
}

func (idm *InvoiceDataMapper) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Retrieve values from plan
	var plan InvoiceDataMapperModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := plan.toAPIRequest()

	newIDM, _, err := idm.client.GetApi().CreateInvoiceDataMapper(ctx).CreateOrUpdateInvoiceDataMapperRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Invoice Data Mapper Entry",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newIDM); err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Invoice Data Mapper Entry",
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

func (idm *InvoiceDataMapper) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state InvoiceDataMapperModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	idmEntry, httpResp, err := idm.client.GetApi().GetInvoiceDataMapper(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Invoice Data Mapper Entry",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := state.fromAPI(idmEntry); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Invoice Data Mapper Entry",
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

func (idm *InvoiceDataMapper) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan InvoiceDataMapperModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := plan.toAPIRequest()

	newIDM, _, err := idm.client.GetApi().UpdateInvoiceDataMapper(ctx, plan.ID.ValueString()).CreateOrUpdateInvoiceDataMapperRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Invoice Data Mapper Entry",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newIDM); err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Invoice Data Mapper Entry",
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

func (idm *InvoiceDataMapper) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	// Retrieve values from state
	var state InvoiceDataMapperModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := idm.client.GetApi().DeleteInvoiceDataMapper(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Invoice Data Mapper Entry",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (idm *InvoiceDataMapper) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *InvoiceDataMapperModel) fromAPI(invoiceDataMapper *api.InvoiceDataMapper) error {

	if invoiceDataMapper == nil {
		return fmt.Errorf("invoice data mapper entry is nil")
	}

	model.ID = types.StringValue(invoiceDataMapper.Id)
	model.Name = types.StringValue(invoiceDataMapper.Name)
	model.Description = omittableStringValue(invoiceDataMapper.Description, model.Description)
	model.Script = types.StringValue(invoiceDataMapper.Script)
	model.Entrypoint = types.StringValue(string(invoiceDataMapper.Entrypoint))

	return nil
}

func (model *InvoiceDataMapperModel) toAPIRequest() api.CreateOrUpdateInvoiceDataMapperRequest {
	return api.CreateOrUpdateInvoiceDataMapperRequest{
		Script:      model.Script.ValueString(),
		Entrypoint:  api.InvoiceDataMapperEntrypointEnum(model.Entrypoint.ValueString()),
		Name:        model.Name.ValueString(),
		Description: model.Description.ValueStringPointer(),
	}
}
