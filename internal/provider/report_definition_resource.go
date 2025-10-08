package provider

import (
	"context"
	"fmt"
	"time"

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

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &reportDefinition{}
	_ resource.ResourceWithConfigure   = &reportDefinition{}
	_ resource.ResourceWithImportState = &reportDefinition{}
)

func NewReportDefinition() resource.Resource {
	return &reportDefinition{}
}

// reportDefinition defines the resource implementation.
type reportDefinition struct {
	client *client.UdomaClient
}

type resultSchemaAttributeModel struct {
	ID          types.String `tfsdk:"id"`
	Label       types.Map    `tfsdk:"label"`
	Sequence    types.Int32  `tfsdk:"sequence"`
	IsReference types.Bool   `tfsdk:"is_reference"`
}

type resultSchemaModel struct {
	ResultType          types.String                 `tfsdk:"result_type"`
	TableRowIdAttribute types.String                 `tfsdk:"table_row_id_attribute"`
	Attributes          []resultSchemaAttributeModel `tfsdk:"attributes"`
}

// reportDefinitionModel describes the resource data model.
type reportDefinitionModel struct {
	ID           types.String       `tfsdk:"id"`
	LastUpdated  types.String       `tfsdk:"last_updated"`
	CreatedAt    types.Int64        `tfsdk:"created_at"`
	UpdatedAt    types.Int64        `tfsdk:"updated_at"`
	Name         types.String       `tfsdk:"name"`
	Description  types.String       `tfsdk:"description"`
	EnvVars      types.Map          `tfsdk:"env_vars"`
	ResultSchema *resultSchemaModel `tfsdk:"result_schema"`
	Parameters   *CustomFormModel   `tfsdk:"parameters"`
	Script       types.String       `tfsdk:"script"`
	Version      types.Int32        `tfsdk:"version"`
}

func NewReportDefinitionModelNull() *reportDefinitionModel {
	return &reportDefinitionModel{
		ResultSchema: &resultSchemaModel{},
	}
}

func (r *reportDefinition) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_report_definition"
}

func (r *reportDefinition) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "Resource represents a defintion of a report",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the report definition",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the report definition was created",
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the report definition was last modified",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the report definition, shown in the admin page",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(250),
				},
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "The description of the report definition",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(250),
				},
			},
			"env_vars": schema.MapAttribute{
				Optional:    true,
				Description: "Environment variables that will be available to the report script",
				ElementType: types.StringType,
			},
			"result_schema": schema.SingleNestedAttribute{
				Required:    true,
				Description: "The schema of the result, which will be used to display it in the UI",
				Attributes: map[string]schema.Attribute{
					"result_type": schema.StringAttribute{
						Required:    true,
						Description: "The type of the result (object, table).",
					},
					"table_row_id_attribute": schema.StringAttribute{
						Optional:    true,
						Description: "The attribute that will be used as the row identifier in the table",
					},
					"attributes": schema.ListNestedAttribute{
						Required:    true,
						Description: "The attributes of the result",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Required:    true,
									Description: "The unique identifier for the attribute",
								},
								"label": schema.MapAttribute{
									Optional:            true,
									MarkdownDescription: "The label of the attribute, which will be displayed in the UI",
									ElementType:         types.StringType,
								},
								"sequence": schema.Int32Attribute{
									Optional:    true,
									Description: "The sequence of the attribute, which will be used to order the attributes in the UI",
								},
								"is_reference": schema.BoolAttribute{
									Optional:    true,
									Description: "Whether the attribute is a reference to another entity",
								},
							},
						},
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "Optional input definition",
				Attributes:  CustomFormNestedSchema(),
			},
			"script": schema.StringAttribute{
				Required:    true,
				Description: "JS script that is executed to generate the report",
			},
			"version": schema.Int32Attribute{
				Optional:    true,
				Description: "The version of the report definition",
			},
		},
	}
}

func (r *reportDefinition) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
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

func (r *reportDefinition) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Retrieve values from plan
	var plan reportDefinitionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Report Definition",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newDefinition, _, err := r.client.GetApi().CreateReportDefinition(ctx).CreateOrUpdateReportDefinitionRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Report Definition",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	diags = plan.fromAPI(newDefinition)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *reportDefinition) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	// Retrieve values from state
	var state reportDefinitionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	report, httpResp, err := r.client.GetApi().GetReportDefinition(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Report Definition",
			"Could not read entity from Udooma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	diags = state.fromAPI(report)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *reportDefinition) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan reportDefinitionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Report Definition",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	id := plan.ID.ValueString()

	updatedReportDefinition, _, err := r.client.GetApi().UpdateReportDefinition(ctx, id).CreateOrUpdateReportDefinitionRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Report Definition",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	diags = plan.fromAPI(updatedReportDefinition)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *reportDefinition) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	// Retrieve values from state
	var state reportDefinitionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.client.GetApi().DeleteReportDefinition(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Report Definition",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (r *reportDefinition) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *resultSchemaModel) fromAPI(resp *api.ReportResultSchema) (diags diag.Diagnostics) {

	if model == nil {
		return
	}

	model.ResultType = types.StringValue(string(resp.ResultType))
	model.TableRowIdAttribute = types.StringPointerValue(resp.TableRowIdAttribute)
	oldAttributes := model.Attributes
	model.Attributes = make([]resultSchemaAttributeModel, len(resp.Attributes))

	for i, attribute := range resp.Attributes {

		var oldAttribute *resultSchemaAttributeModel
		if i < len(oldAttributes) {
			oldAttribute = &oldAttributes[i]
		} else {
			oldAttribute = &resultSchemaAttributeModel{}
		}

		newAttr := resultSchemaAttributeModel{
			ID:          types.StringValue(attribute.Id),
			Sequence:    types.Int32Value(attribute.Sequence),
			IsReference: omittableBooleanValue(attribute.IsReference, oldAttribute.IsReference),
		}

		if attribute.Label != nil {
			newAttr.Label, diags = types.MapValue(types.StringType, stringMapToValueMap(attribute.Label))
			if diags.HasError() {
				return
			}
		} else {
			newAttr.Label = types.MapNull(types.StringType)
		}

		model.Attributes[i] = newAttr
	}

	return
}

func (model *reportDefinitionModel) fromAPI(reportDefinition *api.ReportDefinition) (diags diag.Diagnostics) {

	if reportDefinition == nil {
		diags.AddError("Could not read report definition from API", "report definition is nil")
		return
	}

	model.ID = types.StringValue(reportDefinition.Id)
	model.CreatedAt = types.Int64Value(reportDefinition.CreatedAt)
	model.UpdatedAt = types.Int64Value(reportDefinition.UpdatedAt)
	model.Name = types.StringValue(reportDefinition.Name)
	model.Description = omittableStringValue(reportDefinition.Description, model.Description)
	model.Script = types.StringValue(reportDefinition.Script)
	model.Version = types.Int32PointerValue(reportDefinition.Version)

	if reportDefinition.EnvVars != nil {
		in := stringMapToValueMap(*reportDefinition.EnvVars)
		modelValue, diags := types.MapValue(types.StringType, in)
		if diags.HasError() {
			return diags
		}
		model.EnvVars = modelValue
	} else {
		model.EnvVars = types.MapNull(types.StringType)
	}

	if params := reportDefinition.Parameters.Get(); params != nil {
		if model.Parameters == nil {
			model.Parameters = &CustomFormModel{}
		}
		diags = model.Parameters.fromApiResponse(params)
		if diags.HasError() {
			return
		}
	}

	if model.ResultSchema == nil {
		model.ResultSchema = &resultSchemaModel{}
	}

	diags = model.ResultSchema.fromAPI(&reportDefinition.ResultSchema)

	return diags
}

func (model resultSchemaModel) toAPIRequest() api.ReportResultSchema {

	resultSchema := api.ReportResultSchema{
		ResultType:          api.ReportResultTypeEnum(model.ResultType.ValueString()),
		TableRowIdAttribute: model.TableRowIdAttribute.ValueStringPointer(),
		Attributes:          make([]api.ReportResultSchemaAttribute, len(model.Attributes)),
	}

	for i, attribute := range model.Attributes {

		resultSchema.Attributes[i] = api.ReportResultSchemaAttribute{
			Id:          attribute.ID.ValueString(),
			Sequence:    attribute.Sequence.ValueInt32(),
			Label:       modelMapToStringMap(attribute.Label),
			IsReference: attribute.IsReference.ValueBoolPointer(),
		}

	}

	return resultSchema
}

func (model *reportDefinitionModel) toAPIRequest() (api.CreateOrUpdateReportDefinitionRequest, error) {

	reportDefinition := api.CreateOrUpdateReportDefinitionRequest{
		Name:         model.Name.ValueString(),
		Description:  model.Description.ValueStringPointer(),
		ResultSchema: model.ResultSchema.toAPIRequest(),
		Script:       model.Script.ValueString(),
		Version:      model.Version.ValueInt32Pointer(),
	}

	if envVars := modelMapToStringMap(model.EnvVars); len(envVars) > 0 {
		reportDefinition.EnvVars = &envVars
	}

	if model.Parameters != nil {
		parameters := model.Parameters.toApiRequest()
		reportDefinition.Parameters = *api.NewNullableCustomForm(&parameters)
	}

	return reportDefinition, nil
}
