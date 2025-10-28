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
	CreatedAt   types.String `tfsdk:"created_at"`
	UpdatedAt   types.String `tfsdk:"updated_at"`
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
			"created_at": schema.StringAttribute{
				Computed:    true,
				Description: "The creation timestamp of the data import template.",
			},
			"updated_at": schema.StringAttribute{
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
				Computed:    true,
				Description: "The file type of the data import template (csv, txt, json, xml, xlsx, xls).",
				Validators: []validator.String{
					stringvalidator.OneOf("csv", "txt", "json", "xml", "xlsx", "xls"),
				},
			},
			"data_mapper": schema.StringAttribute{
				Computed:    true,
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

	// Retrieve values from plan
	var plan DataImportTemplateModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	templateReq := plan.toApiRequest()
	id := plan.ID.ValueString()

	newTemplate, _, err := r.client.GetApi().UpdateDataImportTemplate(ctx, id).CreateOrUpdateDataImportTemplateRequest(*templateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Data Import Template",
			fmt.Sprintf("Could not update entity in Udoma, unexpected error: %s", getApiErrorMessage(err)),
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
	ret := &v1.CreateOrUpdateDataImportTemplateRequest{
		Name:       model.Name.ValueString(),
		FileType:   v1.ImportDataTypeEnum(model.FileType.ValueString()),
		DataMapper: model.DataMapper.ValueString(),
	}

	if !model.Description.IsNull() && !model.Description.IsUnknown() {
		ret.Description = model.Description.ValueStringPointer()
	}

	if !model.Icon.IsNull() && !model.Icon.IsUnknown() {
		ret.Icon = model.Icon.ValueStringPointer()
	}

	return ret
}

func (template *DataImportTemplateModel) fromApiResponse(resp *v1.DataImportTemplate) (diags diag.Diagnostics) {
	if resp == nil {
		diags.AddError("Invalid API response", "Received nil DataImportTemplate response from API")
		return
	}

	template.ID = types.StringValue(resp.Id)
	template.Name = types.StringValue(resp.Name)
	template.FileType = types.StringValue(string(resp.FileType))
	template.DataMapper = types.StringValue(resp.DataMapper)

	if resp.Description != nil {
		template.Description = types.StringValue(*resp.Description)
	} else {
		template.Description = types.StringNull()
	}

	if resp.Icon != nil {
		template.Icon = types.StringValue(*resp.Icon)
	} else {
		template.Icon = types.StringNull()
	}

	return diags
}

func (r *DataImportTemplate) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {

	type DataImportTemplateModelV0 struct {
		ID          types.String `tfsdk:"id"`
		Name        types.String `tfsdk:"name"`
		Description types.String `tfsdk:"description"`
		Icon        types.String `tfsdk:"icon"`
		FileType    types.String `tfsdk:"file_type"`
		DataMapper  types.String `tfsdk:"data_mapper"`
		CreatedAt   types.String `tfsdk:"created_at"`
		UpdatedAt   types.String `tfsdk:"updated_at"`
	}

	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"created_at": schema.StringAttribute{
						Computed: true,
					},
					"updated_at": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Required: true,
					},
					"description": schema.StringAttribute{
						Optional: true,
					},
					"icon": schema.StringAttribute{
						Optional: true,
					},
					"file_type": schema.StringAttribute{
						Computed: true,
					},
					"data_mapper": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var priorState DataImportTemplateModelV0
				diags := req.State.Get(ctx, &priorState)
				resp.Diagnostics.Append(diags...)
				if resp.Diagnostics.HasError() {
					return
				}

				upgradedState := DataImportTemplateModel{
					ID:          priorState.ID,
					CreatedAt:   priorState.CreatedAt,
					UpdatedAt:   priorState.UpdatedAt,
					Name:        priorState.Name,
					Description: priorState.Description,
					Icon:        priorState.Icon,
					FileType:    priorState.FileType,
					DataMapper:  priorState.DataMapper,
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, upgradedState)...)
			},
		},
	}
}
