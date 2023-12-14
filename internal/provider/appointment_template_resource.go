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
	_ resource.ResourceWithConfigure   = &AppointmentTemplate{}
	_ resource.ResourceWithImportState = &AppointmentTemplate{}
)

func NewAppointmentTemplate() resource.Resource {
	return &AppointmentTemplate{}
}

// AppointmentTemplate defines the resource implementation
type AppointmentTemplate struct {
	client *client.UdomaClient
}

// AppointmentTemplateModel describes the resource data model
type AppointmentTemplateModel struct {
	ID types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	NameExpression types.String `tfsdk:"name_expression"`
	Description types.String `tfsdk:"description"`
	Inputs *CustomFormModel `tfsdk:"inputs"`
}

func (r *AppointmentTemplate) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appointment_template"
}

func (r *AppointmentTemplate) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				Description: "The ID of the appointment template.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				Description: "The name of the appointment template.",
			},
			"name_expression": schema.StringAttribute{
				Required: true,
				Description: "The name expression of the appointment template.",
			},
			"description": schema.StringAttribute{
				Required: true,
				Description: "The description of the appointment template.",
			},
			"inputs": schema.SingleNestedAttribute{
				Required: true,
				Description: "A custom form to collect data with",
				Attributes: CustomFormNestedSchema(),
			},
		},
	}
}

func (r *AppointmentTemplate) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

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

func (r *AppointmentTemplate) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan AppointmentTemplateModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	
	templateReq := plan.toApiRequest()
	newTemplate, _, err := r.client.GetApi().CreateAppointmentTemplate(ctx).CreateOrUpdateAppointmentTemplateRequest(*templateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to create appointment template",
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

func (r *AppointmentTemplate) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state AppointmentTemplateModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newTemplate, _, err := r.client.GetApi().GetAppointmentTemplate(ctx, state.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Appointment Schedule",
			fmt.Sprintf("Could not read entity in Udoma, unexpected error: %s", getApiErrorMessage(err)),
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

func (r *AppointmentTemplate) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan AppointmentTemplateModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	templateReq := plan.toApiRequest()
	id := plan.ID.ValueString()

	newTemplate, _, err := r.client.GetApi().UpdateAppointmentTemplate(ctx, id).CreateOrUpdateAppointmentTemplateRequest(*templateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Appointment Schedule",
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

func (r *AppointmentTemplate) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state AppointmentTemplateModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.GetApi().DeleteAppointmentTemplate(ctx, state.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Appointment Schedule",
			fmt.Sprintf("Could not delete entity in Udoma, unexpected error: %s", getApiErrorMessage(err)),
		)
		return
	}
}

func (r *AppointmentTemplate) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (template *AppointmentTemplateModel) toApiRequest() *v1.CreateOrUpdateAppointmentTemplateRequest {
	return &v1.CreateOrUpdateAppointmentTemplateRequest{
		Name:           template.Name.ValueStringPointer(),
		NameExpression: template.NameExpression.ValueStringPointer(),
		Description:    template.Description.ValueStringPointer(),
		Form:						template.Inputs.toApiRequest(),
	}
}

func (template *AppointmentTemplateModel) fromApiResponse(resp *v1.AppointmentTemplate) (diags diag.Diagnostics) {
	template.ID = types.StringValue(resp.Id)
	template.Name = types.StringValue(resp.Name)
	template.NameExpression = types.StringValue(resp.NameExpression)
	template.Description = types.StringValue(sdp(resp.Description))

	if template.Inputs == nil {
		template.Inputs = &CustomFormModel{}
	}
	diags = template.Inputs.fromApiResponse(&resp.Form)

	return
}
