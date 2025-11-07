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
	v1 "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"
)

var (
	_ resource.ResourceWithConfigure   = &DataImportTemplate{}
	_ resource.ResourceWithImportState = &DataImportTemplate{}
)

func NewDataImportTemplate() resource.Resource {
	return &DataImportTemplate{}
}

// DataImportTemplate defines the resource implementation
type DataImportTemplate struct {
	client *client.UdomaClient
}

// DataImportTemplateModel describes the resource data model
type DataImportTemplateModel struct {
	ID          types.String `tfsdk:"id"`
	CreatedAt   types.Int64  `tfsdk:"created_at"`
	UpdatedAt   types.Int64  `tfsdk:"updated_at"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Icon        types.String `tfsdk:"icon"`
	FileType    types.String `tfsdk:"file_type"`
	DataMapper  types.String `tfsdk:"data_mapper"`
}

func (r *DataImportTemplate) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_data_import_template"
}

func (r *DataImportTemplate) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the data import template.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The creation timestamp of the data import template.",
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The last update timestamp of the data import template.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the data import template.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "The description of the data import template.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
			},
			"icon": schema.StringAttribute{
				Optional:    true,
				Description: "The icon of the data import template.",
			},
			"file_type": schema.StringAttribute{
				Required:    true,
				Description: "The file type of the data import template (csv, txt, json, xml, xlsx, xls).",
				Validators: []validator.String{
					stringvalidator.OneOf("csv", "txt", "json", "xml", "xlsx", "xls"),
				},
			},
			"data_mapper": schema.StringAttribute{
				Required:    true,
				Description: "A JS expression that maps imported data to the system’s data model.",
			},
		},
	}
}

func (r *DataImportTemplate) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	if req.ProviderData == nil {
		return
	}

	cl, ok := req.ProviderData.(*client.UdomaClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.UdomaClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = cl
}

func (r *DataImportTemplate) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan DataImportTemplateModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	templateReq := plan.toApiRequest()
	newTemplate, _, err := r.client.GetApi().CreateDataImportTemplate(ctx).CreateOrUpdateDataImportTemplateRequest(*templateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to create data import template",
			err.Error(),
		)
		return
	}

	diags = plan.fromApiResponse(newTemplate)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *DataImportTemplate) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state DataImportTemplateModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newTemplate, httpResp, err := r.client.GetApi().GetDataImportTemplate(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Data Import Template",
			"Could not read entity from Udooma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	diags = state.fromApiResponse(newTemplate)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *DataImportTemplate) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DataImportTemplateModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	templateReq := plan.toApiRequest()
	id := plan.ID.ValueString()

	updatedTemplate, _, err := r.client.GetApi().
		UpdateDataImportTemplate(ctx, id).
		CreateOrUpdateDataImportTemplateRequest(*templateReq).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Data Import Template",
			fmt.Sprintf("Could not update entity in Udoma, unexpected error: %s", getApiErrorMessage(err)),
		)
		return
	}

	diags = plan.fromApiResponse(updatedTemplate)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *DataImportTemplate) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state DataImportTemplateModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.client.GetApi().DeleteDataImportTemplate(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Data Import Template",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (r *DataImportTemplate) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *DataImportTemplateModel) toApiRequest() *v1.CreateOrUpdateDataImportTemplateRequest {

	req := &v1.CreateOrUpdateDataImportTemplateRequest{
		Name:        model.Name.ValueString(),
		FileType:    v1.ImportDataTypeEnum(model.FileType.ValueString()),
		Icon:        model.Icon.ValueStringPointer(),
		Description: model.Description.ValueStringPointer(),
		DataMapper:  model.DataMapper.ValueString(),
	}

	return req
}

func (template *DataImportTemplateModel) fromApiResponse(resp *v1.DataImportTemplate) (diags diag.Diagnostics) {

	if resp == nil {
		diags.AddError("Invalid API response", "Received nil DataImportTemplate response from API")
		return
	}

	template.ID = types.StringValue(resp.Id)
	template.CreatedAt = types.Int64Value(resp.CreatedAt)
	template.UpdatedAt = types.Int64Value(resp.UpdatedAt)
	template.Name = types.StringValue(resp.Name)
	template.FileType = types.StringValue(string(resp.FileType))
	template.DataMapper = types.StringValue(resp.DataMapper)
	template.Icon = omittableStringValue(resp.Icon, template.Icon)
	template.Description = omittableStringValue(resp.Description, template.Description)

	return diags
}
