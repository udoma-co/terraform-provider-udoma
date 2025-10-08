package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
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
	_ resource.ResourceWithConfigure   = &AppointmentSchedule{}
	_ resource.ResourceWithImportState = &AppointmentSchedule{}
)

func NewAppointmentSchedule() resource.Resource {
	return &AppointmentSchedule{}
}

// AppointmentSchedule defines the resource implementation
type AppointmentSchedule struct {
	client *client.UdomaClient
}

// AppointmentScheduleModel describes the resource data model
type AppointmentScheduleModel struct {
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Description  types.Map    `tfsdk:"description"`
	TemplateRef  types.String `tfsdk:"template_ref"`
	SlotDuration types.Int64  `tfsdk:"slot_duration"`
	GapDuration  types.Int64  `tfsdk:"gap_duration"`
	Color        types.String `tfsdk:"color"`
}

func (r *AppointmentSchedule) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appointment_schedule"
}

func (r *AppointmentSchedule) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents an appointment schedule",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the appointment schedule",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the appointment schedule",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
			},
			"description": schema.MapAttribute{
				Required:    true,
				Description: "Description of the appointment schedule in all different languages",
				ElementType: types.StringType,
			},
			"template_ref": schema.StringAttribute{
				Required:    true,
				Description: "ID of the template to be used in the schedule",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(25),
				},
			},
			"slot_duration": schema.Int64Attribute{
				Required:    true,
				Description: "The length of a single appointment",
			},
			"gap_duration": schema.Int64Attribute{
				Required:    true,
				Description: "The length of the gap in-between slots",
			},
			"color": schema.StringAttribute{
				Required:    true,
				Description: "The color to use when displaying the schedule",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(31),
				},
			},
		},
	}
}

func (r *AppointmentSchedule) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = cl
}

func (r *AppointmentSchedule) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan AppointmentScheduleModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiReq := plan.toAPIRequest()
	newApp, _, err := r.client.GetApi().CreateAppointmentSchedule(ctx).CreateOrUpdateAppointmentScheduleRequest(apiReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Appointment Schedule",
			fmt.Sprintf("Could not create entity in Udoma, unexpected error: %s", getApiErrorMessage(err)),
		)
	}

	diags = plan.fromApiResponse(newApp)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *AppointmentSchedule) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state AppointmentScheduleModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newApp, httpResp, err := r.client.GetApi().GetAppointmentSchedule(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Appointment Schedule",
			"Could not read entity from Udooma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	diags = state.fromApiResponse(newApp)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *AppointmentSchedule) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan AppointmentScheduleModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	appReq := plan.toAPIRequest()
	id := plan.ID.ValueString()

	newApp, _, err := r.client.GetApi().UpdateAppointmentSchedule(ctx, id).CreateOrUpdateAppointmentScheduleRequest(appReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Appointment Schedule",
			fmt.Sprintf("Could not update entity in Udoma, unexpected error: %s", getApiErrorMessage(err)),
		)
		return
	}

	diags = plan.fromApiResponse(newApp)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *AppointmentSchedule) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state AppointmentScheduleModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.client.GetApi().DeleteAppointmentSchedule(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Appointment Schedule",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (r *AppointmentSchedule) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (app *AppointmentScheduleModel) toAPIRequest() v1.CreateOrUpdateAppointmentScheduleRequest {
	schedule := v1.CreateOrUpdateAppointmentScheduleRequest{
		Name:         app.Name.ValueString(),
		TemplateRef:  app.TemplateRef.ValueString(),
		SlotDuration: int32(app.SlotDuration.ValueInt64()),
		GapDuration:  i64ToI32Ptr(app.GapDuration.ValueInt64()),
		Color:        app.Color.ValueStringPointer(),
	}

	if description := modelMapToStringMap(app.Description); len(description) > 0 {
		schedule.Description = &description
	}

	return schedule
}

func (app *AppointmentScheduleModel) fromApiResponse(resp *v1.AppointmentSchedule) (diags diag.Diagnostics) {
	if resp == nil {
		diags.AddError(
			"Appointment Schedule is nil",
			"The argument is used to map values to the Appointment Schedule and cannot be nil",
		)
		return
	}

	app.ID = types.StringValue(resp.Id)
	app.Name = types.StringValue(resp.Name)

	if resp.Description != nil {
		mappedDescription := make(map[string]attr.Value)
		for key, value := range *resp.Description {
			mappedDescription[key] = types.StringValue(value)
		}

		app.Description, diags = types.MapValue(types.StringType, mappedDescription)
		if diags.HasError() {
			return
		}
	}

	app.TemplateRef = types.StringValue(resp.Template.Id)
	app.SlotDuration = types.Int64Value(int64(resp.SlotDuration))
	app.GapDuration = types.Int64Value(int64(idp32(resp.GapDuration)))
	app.Color = types.StringPointerValue(resp.Color)

	return
}
