package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	api "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/tf"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &caseTemplate{}
	_ resource.ResourceWithConfigure   = &caseTemplate{}
	_ resource.ResourceWithImportState = &caseTemplate{}
)

func NewCaseTemplate() resource.Resource {
	return &caseTemplate{}
}

// caseTemplate defines the resource implementation.
type caseTemplate struct {
	client *client.UdomaClient
}

// caseTemplateModel describes the resource data model.
type caseTemplateModel struct {
	ID             types.String       `tfsdk:"id"`
	LastUpdated    types.String       `tfsdk:"last_updated"`
	CreatedAt      types.Int64        `tfsdk:"created_at"`
	UpdatedAt      types.Int64        `tfsdk:"updated_at"`
	Name           types.String       `tfsdk:"name"`
	NameExpression types.String       `tfsdk:"name_expression"`
	Label          types.Map          `tfsdk:"label"`
	Description    types.Map          `tfsdk:"description"`
	InfoText       types.Map          `tfsdk:"info_text"`
	Icon           types.String       `tfsdk:"icon"`
	CustomInputs   tf.JsonObjectValue `tfsdk:"custom_inputs"`
	Config         tf.JsonObjectValue `tfsdk:"config"`
}

func (r *caseTemplate) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_case_template"
}

func (r *caseTemplate) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "Resource represents a template for raising cases",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the case template",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the template was created",
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the template was last modified",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the case template, shown in the admin page",
			},
			"name_expression": schema.StringAttribute{
				Optional:    true,
				Description: "The name expression used to generate the name of cases based on this template",
			},
			"label": schema.MapAttribute{
				Optional:    true,
				Description: "The short label to be displayed on the reporting page",
				ElementType: types.StringType,
			},
			"description": schema.MapAttribute{
				Optional:    true,
				Description: "The description of the case template, shown in the reporting page",
				ElementType: types.StringType,
			},
			"info_text": schema.MapAttribute{
				Optional:    true,
				Description: "A longer introduction text, shown in the case specific reporting page",
				ElementType: types.StringType,
			},
			"icon": schema.StringAttribute{
				Optional:    true,
				Description: "The icon to be displayed on the reporting page",
			},
			"custom_inputs": schema.StringAttribute{
				Required:    true,
				CustomType:  tf.JsonObjectType{},
				Description: "The custom inputs to be displayed on the reporting page (as serialized JSON)",
			},
			"config": schema.StringAttribute{
				Optional:   true,
				CustomType: tf.JsonObjectType{},
				Description: "Defines custom behaviour of a case, based on the case template that was " +
					"used to create it (as serialized JSON)",
			},
		},
	}
}

func (r *caseTemplate) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *caseTemplate) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Retrieve values from plan
	var plan caseTemplateModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Case Template",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newEndpoint, _, err := r.client.GetApi().CreateCaseTemplate(ctx).CreateOrUpdateCaseTemplateRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Case Template",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newEndpoint); err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Case Template",
			"Could not process API response, unexpected error: "+err.Error(),
		)
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

func (r *caseTemplate) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	// Retrieve values from state
	var state caseTemplateModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	template, _, err := r.client.GetApi().GetCaseTemplate(ctx, state.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Case Template",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := state.fromAPI(template); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Case Template",
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

func (r *caseTemplate) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan caseTemplateModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Case Template",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	id := plan.ID.ValueString()

	newEndpoint, _, err := r.client.GetApi().UpdateCaseTemplate(ctx, id).CreateOrUpdateCaseTemplateRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Case Template",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newEndpoint); err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Case Template",
			"Could not process API response, unexpected error: "+err.Error(),
		)
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

func (r *caseTemplate) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	// Retrieve values from state
	var state caseTemplateModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.GetApi().DeleteCaseTemplate(ctx, state.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Case Template",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}
}

func (r *caseTemplate) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *caseTemplateModel) fromAPI(template *api.CaseTemplate) error {

	if template == nil {
		return fmt.Errorf("case template is nil")
	}

	model.ID = types.StringValue(sdp(template.Id))
	model.CreatedAt = types.Int64Value(idp(template.CreatedAt))
	model.UpdatedAt = types.Int64Value(idp(template.UpdatedAt))
	model.Name = types.StringValue(sdp(template.Name))
	model.NameExpression = types.StringValue(sdp(template.NameExpression))
	if template.Label != nil {
		in := stringMapToValueMap(*template.Label)
		modelValue, diags := types.MapValue(types.StringType, in)
		if diags.HasError() {
			return fmt.Errorf("error converting label to map: %v", diags)
		}
		model.Label = modelValue
	}
	if template.Description != nil {
		in := stringMapToValueMap(*template.Description)
		modelValue, diags := types.MapValue(types.StringType, in)
		if diags.HasError() {
			return fmt.Errorf("error converting description to map: %v", diags)
		}
		model.Description = modelValue
	}
	if template.InfoText != nil {
		in := stringMapToValueMap(*template.InfoText)
		modelValue, diags := types.MapValue(types.StringType, in)
		if diags.HasError() {
			return fmt.Errorf("error converting info text to map: %v", diags)
		}
		model.InfoText = modelValue
	}
	model.Icon = types.StringValue(sdp(template.Icon))

	if template.CustomInputs != nil {
		customInputs, err := json.Marshal(template.CustomInputs)
		if err != nil {
			return fmt.Errorf("failed to marshal custom inputs: %w", err)
		}
		model.CustomInputs = tf.NewJsonObjectValue(string(customInputs))
	}

	return nil
}

func (model *caseTemplateModel) toAPIRequest() (api.CreateOrUpdateCaseTemplateRequest, error) {

	template := api.CreateOrUpdateCaseTemplateRequest{
		Name:           model.Name.ValueStringPointer(),
		NameExpression: model.NameExpression.ValueStringPointer(),
	}

	template.Label = modelMapToStringMap(model.Label)
	template.Description = modelMapToStringMap(model.Description)
	template.InfoText = modelMapToStringMap(model.InfoText)

	template.Icon = model.Icon.ValueStringPointer()

	if !model.CustomInputs.IsNull() && !model.CustomInputs.IsUnknown() {
		if diag := model.CustomInputs.Unmarshal(&template.CustomInputs); diag.HasError() {
			return template, fmt.Errorf("failed to unmarshal custom inputs: %v", diag)
		}

		// update model
		customInputs, err := json.Marshal(template.CustomInputs)
		if err != nil {
			return template, fmt.Errorf("failed to marshal custom inputs: %w", err)
		}
		model.CustomInputs = tf.NewJsonObjectValue(string(customInputs))
	}

	if !model.Config.IsNull() && !model.Config.IsUnknown() {
		if diag := model.Config.Unmarshal(&template.Config); diag.HasError() {
			return template, fmt.Errorf("failed to unmarshal config: %v", diag)
		}

		// update model
		config, err := json.Marshal(template.Config)
		if err != nil {
			return template, fmt.Errorf("failed to marshal config: %w", err)
		}
		model.Config = tf.NewJsonObjectValue(string(config))
	}

	return template, nil
}
