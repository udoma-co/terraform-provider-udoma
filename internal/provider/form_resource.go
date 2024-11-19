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
	_ resource.ResourceWithConfigure   = &Form{}
	_ resource.ResourceWithImportState = &Form{}
)

func NewForm() resource.Resource {
	return &Form{}
}

type Form struct {
	client *client.UdomaClient
}

// AppointmentTemplateModel describes the resource data model
type FormModel struct {
	ID             types.String     `tfsdk:"id"`
	Name           types.String     `tfsdk:"name"`
	Description    types.String     `tfsdk:"description"`
	FormDefinition *CustomFormModel `tfsdk:"form_definition"`
}

func (r *Form) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_form"
}

func (r *Form) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the custom form.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the custom form.",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "The description of the custom form.",
			},
			"form_definition": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "A custom form to collect data with",
				Attributes:  CustomFormNestedSchema(),
			},
		},
	}
}

func (r *Form) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

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

func (r *Form) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan FormModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	formReq := plan.toApiRequest()
	newForm, _, err := r.client.GetApi().CreateCustomForm(ctx).CreateOrUpdateCustomFormRequest(*formReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to create custom form entry",
			err.Error(),
		)
		return
	}

	diags = plan.fromApiResponse(newForm)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *Form) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state FormModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newForm, httpResp, err := r.client.GetApi().GetCustomForm(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading custom form entry",
			"Could not read entity from Udooma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	diags = state.fromApiResponse(newForm)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *Form) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan FormModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	formReq := plan.toApiRequest()
	id := plan.ID.ValueString()

	newForm, _, err := r.client.GetApi().UpdateCustomForm(ctx, id).CreateOrUpdateCustomFormRequest(*formReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating custom form entry",
			fmt.Sprintf("Could not update entity in Udoma, unexpected error: %s", getApiErrorMessage(err)),
		)
		return
	}

	diags = plan.fromApiResponse(newForm)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *Form) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state FormModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.client.GetApi().DeleteCustomForm(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting custom form entry",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (r *Form) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (template *FormModel) toApiRequest() *v1.CreateOrUpdateCustomFormRequest {
	form := template.FormDefinition.toApiRequest()
	ret := &v1.CreateOrUpdateCustomFormRequest{
		Name:           template.Name.ValueStringPointer(),
		Description:    template.Description.ValueStringPointer(),
		FormDefinition: *v1.NewNullableCustomForm(&form),
	}

	return ret
}

func (template *FormModel) fromApiResponse(resp *v1.Form) (diags diag.Diagnostics) {
	template.ID = types.StringValue(resp.Id)
	template.Name = types.StringPointerValue(resp.Name)
	template.Description = omittableStringValue(resp.Description, template.Description)

	if template.FormDefinition == nil {
		template.FormDefinition = &CustomFormModel{}
	}
	diags = template.FormDefinition.fromApiResponse(resp.FormDefinition.Get())

	return
}
