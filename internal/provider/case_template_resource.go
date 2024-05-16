package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
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
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/tf"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &CaseTemplate{}
	_ resource.ResourceWithConfigure   = &CaseTemplate{}
	_ resource.ResourceWithImportState = &CaseTemplate{}
)

func NewCaseTemplate() resource.Resource {
	return &CaseTemplate{}
}

// CaseTemplate defines the resource implementation.
type CaseTemplate struct {
	client *client.UdomaClient
}

// CaseTemplateModel describes the resource data model.
type CaseTemplateModel struct {
	ID               types.String       `tfsdk:"id"`
	LastUpdated      types.String       `tfsdk:"last_updated"`
	CreatedAt        types.Int64        `tfsdk:"created_at"`
	UpdatedAt        types.Int64        `tfsdk:"updated_at"`
	Name             types.String       `tfsdk:"name"`
	Access           []types.String     `tfsdk:"access"`
	NameExpression   types.String       `tfsdk:"name_expression"`
	Label            types.Map          `tfsdk:"label"`
	Description      types.Map          `tfsdk:"description"`
	InfoText         types.Map          `tfsdk:"info_text"`
	Icon             types.String       `tfsdk:"icon"`
	CustomInputs     tf.JsonObjectValue `tfsdk:"custom_inputs"`
	Config           *CaseConfigModel   `tfsdk:"config"`
	AdCategories     []types.String     `tfsdk:"ad_categories"`
	ConfirmationText types.Map          `tfsdk:"confirmation_text"`
}

func (r *CaseTemplate) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_case_template"
}

func (r *CaseTemplate) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"access": schema.ListAttribute{
				Required:    true,
				Description: "The list of accessibility options for the case template",
				ElementType: types.StringType,
				Validators: []validator.List{
					listvalidator.UniqueValues(),
					listvalidator.ValueStringsAre(stringvalidator.OneOf(
						string(api.CASETEMPLATEACCESSIBILITY_PROPERTY_MANAGER),
						string(api.CASETEMPLATEACCESSIBILITY_TENANT),
						string(api.CASETEMPLATEACCESSIBILITY_PUBLIC),
					)),
				},
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
			"config": schema.SingleNestedAttribute{
				Optional: true,
				Description: "Defines custom behaviour of a case, based on the case template that was " +
					"used to create it",
				Attributes: CaseConfigModelNestedSchema(),
			},
			"ad_categories": schema.ListAttribute{
				Optional:    true,
				Description: "The list of ad categories for the case template",
				ElementType: types.StringType,
				Validators: []validator.List{
					listvalidator.UniqueValues(),
					listvalidator.ValueStringsAre(stringvalidator.OneOf(
						string(api.CASETEMPLATEADCATEGORYENUM_ELECTRICITY_PROVIDERS),
						string(api.CASETEMPLATEADCATEGORYENUM_GAS_PROVIDERS),
						string(api.CASETEMPLATEADCATEGORYENUM_HOME_INSURANCE),
						string(api.CASETEMPLATEADCATEGORYENUM_INTERNET_PROVIDERS),
					)),
				},
			},
			"confirmation_text": schema.MapAttribute{
				Optional:    true,
				Description: "Text for the confirmation window before submiting case",
				ElementType: types.StringType,
			},
		},
	}
}

func (r *CaseTemplate) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CaseTemplate) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Retrieve values from plan
	var plan CaseTemplateModel
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

func (r *CaseTemplate) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	// Retrieve values from state
	var state CaseTemplateModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	template, httpResp, err := r.client.GetApi().GetCaseTemplate(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
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

func (r *CaseTemplate) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan CaseTemplateModel
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

func (r *CaseTemplate) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	// Retrieve values from state
	var state CaseTemplateModel
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

func (r *CaseTemplate) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *CaseTemplateModel) fromAPI(template *api.CaseTemplate) error {

	if template == nil {
		return fmt.Errorf("case template is nil")
	}

	model.ID = types.StringValue(template.Id)
	model.CreatedAt = types.Int64Value(template.CreatedAt)
	model.UpdatedAt = types.Int64Value(template.UpdatedAt)
	model.Name = types.StringValue(template.Name)
	model.NameExpression = omittableStringValue(template.NameExpression, model.NameExpression)
	model.Icon = omittableStringValue(template.Icon, model.Icon)

	model.Access = make([]types.String, len(template.Access))
	for i := range template.Access {
		model.Access[i] = types.StringValue(string(template.Access[i]))
	}

	if len(template.AdCategories) > 0 {
		model.AdCategories = make([]types.String, len(template.AdCategories))
		for i := range template.AdCategories {
			model.AdCategories[i] = types.StringValue(string(template.AdCategories[i]))
		}
	} else {
		if model.AdCategories != nil {
			model.AdCategories = make([]types.String, 0)
		}
	}

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

	if template.ConfirmationText != nil {
		in := stringMapToValueMap(*template.ConfirmationText)
		modelValue, diags := types.MapValue(types.StringType, in)
		if diags.HasError() {
			return fmt.Errorf("error converting confirmation text to map: %v", diags)
		}
		model.ConfirmationText = modelValue
	}

	customInputs, err := json.Marshal(template.CustomInputs)
	if err != nil {
		return fmt.Errorf("failed to marshal custom inputs: %w", err)
	}
	model.CustomInputs = tf.NewJsonObjectValue(string(customInputs))

	cfg := &CaseConfigModel{}
	isEmpty := cfg.fromApiResponse(template.Config)

	// handle default {} return from API
	if !isEmpty {
		model.Config = cfg
	}

	return nil
}

func (model *CaseTemplateModel) toAPIRequest() (api.CreateOrUpdateCaseTemplateRequest, error) {

	template := api.CreateOrUpdateCaseTemplateRequest{
		Name:             model.Name.ValueString(),
		NameExpression:   model.NameExpression.ValueStringPointer(),
		Label:            modelMapToStringMap(model.Label),
		Description:      modelMapToStringMap(model.Description),
		InfoText:         modelMapToStringMap(model.InfoText),
		Icon:             model.Icon.ValueStringPointer(),
		ConfirmationText: modelMapToStringMap(model.ConfirmationText),
	}

	template.Access = make([]api.CaseTemplateAccessibility, len(model.Access))
	for i := range model.Access {
		template.Access[i] = api.CaseTemplateAccessibility(model.Access[i].ValueString())
	}

	template.AdCategories = make([]api.CaseTemplateAdCategoryEnum, len(model.AdCategories))
	for i := range model.AdCategories {
		template.AdCategories[i] = api.CaseTemplateAdCategoryEnum(model.AdCategories[i].ValueString())
	}

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

	if model.Config != nil {
		template.Config = model.Config.toApiRequest()
	}

	return template, nil
}
