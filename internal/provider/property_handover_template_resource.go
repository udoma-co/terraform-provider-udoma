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
	v1 "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"
)

var (
	_ resource.ResourceWithConfigure   = &PropertyHandoverTemplate{}
	_ resource.ResourceWithImportState = &PropertyHandoverTemplate{}
)

func NewPropertyHandoverTemplate() resource.Resource {
	return &PropertyHandoverTemplate{}
}

// PropertyHandoverTemplate defines the resource implementation
type PropertyHandoverTemplate struct {
	client *client.UdomaClient
}

// PropertyTemplateTemplateModel describes the resource data model
type PropertyHandoverTemplateModel struct {
	ID          types.String     `tfsdk:"id"`
	Name        types.String     `tfsdk:"name"`
	Description types.String     `tfsdk:"description"`
	Inputs      *CustomFormModel `tfsdk:"inputs"`
	Version     types.Int32      `tfsdk:"version"`
}

func (r *PropertyHandoverTemplate) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_property_handover_template"
}

func (r *PropertyHandoverTemplate) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1, // update custom form input items
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the property handover template.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the property handover template.",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "The description of the property handover template.",
			},
			"inputs": schema.SingleNestedAttribute{
				Required:    true,
				Description: "A custom form to collect data with",
				Attributes:  CustomFormNestedSchema(),
			},
			"version": schema.Int32Attribute{
				Computed:    true,
				Description: "The version of the property handover template.",
			},
		},
	}
}

func (r *PropertyHandoverTemplate) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

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

func (r *PropertyHandoverTemplate) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan PropertyHandoverTemplateModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	templateReq := plan.toApiRequest()
	newTemplate, _, err := r.client.GetApi().CreatePropertyHandoverTemplate(ctx).CreateOrUpdatePropertyHandoverTemplateRequest(*templateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to create property handover template",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
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

func (r *PropertyHandoverTemplate) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state PropertyHandoverTemplateModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newTemplate, httpResp, err := r.client.GetApi().GetPropertyHandoverTemplate(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Property Handover Template",
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

func (r *PropertyHandoverTemplate) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan PropertyHandoverTemplateModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	templateReq := plan.toApiRequest()
	id := plan.ID.ValueString()

	newTemplate, _, err := r.client.GetApi().UpdatePropertyHandoverTemplate(ctx, id).CreateOrUpdatePropertyHandoverTemplateRequest(*templateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating PropertyHandover Template",
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

func (r *PropertyHandoverTemplate) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state PropertyHandoverTemplateModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.client.GetApi().DeletePropertyHandoverTemplate(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Property Handover Template",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (r *PropertyHandoverTemplate) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (template *PropertyHandoverTemplateModel) toApiRequest() *v1.CreateOrUpdatePropertyHandoverTemplateRequest {
	form := template.Inputs.toApiRequest()
	return &v1.CreateOrUpdatePropertyHandoverTemplateRequest{
		Name:        template.Name.ValueString(),
		Description: template.Description.ValueStringPointer(),
		CustomForm:  *v1.NewNullableCustomForm(&form),
		Version:     template.Version.ValueInt32Pointer(),
	}
}

func (template *PropertyHandoverTemplateModel) fromApiResponse(resp *v1.PropertyHandoverTemplate) (diags diag.Diagnostics) {
	template.ID = types.StringValue(resp.Id)
	template.Name = types.StringValue(resp.Name)
	template.Description = omittableStringValue(resp.Description, template.Description)
	template.Version = types.Int32PointerValue(resp.Version)

	if template.Inputs == nil {
		template.Inputs = &CustomFormModel{}
	}
	diags = template.Inputs.fromApiResponse(resp.CustomForm.Get())

	return
}

func (r *PropertyHandoverTemplate) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {

	// PropertyTemplateTemplateModel describes the resource data model with the prior version of the custom form
	type PropertyHandoverTemplateModelV0 struct {
		ID          types.String       `tfsdk:"id"`
		Name        types.String       `tfsdk:"name"`
		Description types.String       `tfsdk:"description"`
		Inputs      *CustomFormModelV0 `tfsdk:"inputs"`
		Version     types.Int32        `tfsdk:"version"`
	}

	return map[int64]resource.StateUpgrader{
		// State upgrade implementation from 0 (prior state version) to 1 (Schema.Version)
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						Required: true,
					},
					"description": schema.StringAttribute{
						Optional: true,
					},
					"inputs": schema.SingleNestedAttribute{
						Required:   true,
						Attributes: CustomFormNestedSchemaV0(), // use the prior version of the custom form schema
					},
					"version": schema.Int32Attribute{
						Computed: true,
					},
				},
			},
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var priorStateData PropertyHandoverTemplateModelV0

				resp.Diagnostics.Append(req.State.Get(ctx, &priorStateData)...)

				if resp.Diagnostics.HasError() {
					return
				}

				priorCustomForm := priorStateData.Inputs.toApiRequest()
				newCustomForm, err := priorCustomForm.Migrate()
				if err != nil {
					resp.Diagnostics.AddError(
						"Error Migrating Custom Form",
						"Could not migrate custom form: "+err.Error(),
					)
					return
				}

				upgradedInputs := &CustomFormModel{}
				diags := upgradedInputs.fromApiResponse(newCustomForm)
				if diags.HasError() {
					resp.Diagnostics.Append(diags...)
					return
				}

				upgradedStateData := PropertyHandoverTemplateModel{
					ID:          priorStateData.ID,
					Name:        priorStateData.Name,
					Description: priorStateData.Description,
					Inputs:      upgradedInputs,
					Version:     priorStateData.Version,
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, upgradedStateData)...)
			},
		},
	}
}
