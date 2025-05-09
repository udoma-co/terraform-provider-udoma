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
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
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
	ID                         types.String     `tfsdk:"id"`
	Name                       types.String     `tfsdk:"name"`
	NameExpression             types.String     `tfsdk:"name_expression"`
	Description                types.String     `tfsdk:"description"`
	Inputs                     *CustomFormModel `tfsdk:"inputs"`
	RequireConfirmation        types.Bool       `tfsdk:"require_confirmation"`
	ConfirmationReminders      types.List       `tfsdk:"confirmation_reminders"`
	DefaultScheduleDescription types.Map        `tfsdk:"default_schedule_description"`
	InvitationText             types.String     `tfsdk:"invitation_text"`
	Icon                       types.String     `tfsdk:"icon"`
	Version                    types.Int32      `tfsdk:"version"`
}

func (r *AppointmentTemplate) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appointment_template"
}

func (r *AppointmentTemplate) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1, // update custom form input items
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appointment template.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the appointment template.",
			},
			"name_expression": schema.StringAttribute{
				Optional:    true,
				Description: "Optional JS expression used to generate the name of appointments.",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "The description of the appointment template.",
			},
			"inputs": schema.SingleNestedAttribute{
				Required:    true,
				Description: "A custom form to collect data with",
				Attributes:  CustomFormNestedSchema(),
			},
			"require_confirmation": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether the appointment requires confirmation.",
				MarkdownDescription: `Whether the appointment requires confirmation.
				If set to true, the Appointment will not be scheduled until the
				manager has confirmed it.`,
			},
			"confirmation_reminders": schema.ListAttribute{
				Optional:    true,
				ElementType: types.Int64Type,
				Description: `The number of days after which a reminder should be sent out.`,
				MarkdownDescription: `The number of days after which a reminder should
				be sent out. The list is taken as provided and a reminder will be set
				for the amount of days as set in the first element of the list. Once
				that time has passed, the reminder will be rescheduled for the next
				element in the list. This is repeated until the list is empty. The
				values in the list are considered relative to the previous reminder.
				So [2, 2] will send out a reminder after 2 and 4 days. The reminders
				will, however, not be sent out on weekends. So if the first reminder is
				sent out on a Friday, the second reminder will be sent out on Tuesday.`,
			},
			"default_schedule_description": schema.MapAttribute{
				Optional:    true,
				Description: "The default schedule description in all different languages.",
				ElementType: types.StringType,
			},
			"invitation_text": schema.StringAttribute{
				Optional:    true,
				Description: "The invitation text for inviting users to book appointment.",
			},
			"icon": schema.StringAttribute{
				Optional:    true,
				Description: "The icon of the appointment template.",
			},
			"version": schema.Int32Attribute{
				Optional:    true,
				Description: "The version of the appointment template.",
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

	newTemplate, httpResp, err := r.client.GetApi().GetAppointmentTemplate(ctx, state.ID.ValueString()).Execute()
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

	httpResp, err := r.client.GetApi().DeleteAppointmentTemplate(ctx, state.ID.ValueString()).Execute()
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

func (r *AppointmentTemplate) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (template *AppointmentTemplateModel) toApiRequest() *v1.CreateOrUpdateAppointmentTemplateRequest {
	form := template.Inputs.toApiRequest()
	ret := &v1.CreateOrUpdateAppointmentTemplateRequest{
		Name:                  template.Name.ValueString(),
		NameExpression:        template.NameExpression.ValueStringPointer(),
		Description:           template.Description.ValueStringPointer(),
		Form:                  *v1.NewNullableCustomForm(&form),
		RequireConfirmation:   template.RequireConfirmation.ValueBoolPointer(),
		ConfirmationReminders: modelListToInt32Slice(template.ConfirmationReminders),
		InvitationText:        template.InvitationText.ValueStringPointer(),
		Icon:                  template.Icon.ValueStringPointer(),
		Version:               template.Version.ValueInt32Pointer(),
	}

	if description := modelMapToStringMap(template.DefaultScheduleDescription); len(description) > 0 {
		ret.DefaultScheduleDescription = &description
	}

	return ret
}

func (template *AppointmentTemplateModel) fromApiResponse(resp *v1.AppointmentTemplate) (diags diag.Diagnostics) {
	template.ID = types.StringValue(resp.Id)
	template.Name = types.StringValue(resp.Name)
	template.NameExpression = omittableStringValue(resp.NameExpression, template.NameExpression)
	template.Description = omittableStringValue(resp.Description, template.Description)
	template.InvitationText = omittableStringValue(resp.InvitationText, template.InvitationText)
	template.Icon = omittableStringValue(resp.Icon, template.Icon)
	template.RequireConfirmation = omittableBooleanValue(resp.RequireConfirmation, template.RequireConfirmation)
	template.Version = types.Int32PointerValue(resp.Version)

	if len(resp.ConfirmationReminders) != 0 {
		template.ConfirmationReminders, diags = types.ListValue(types.Int64Type, int32SliceToValueList(resp.ConfirmationReminders))
		if diags.HasError() {
			return
		}
	} else {
		template.ConfirmationReminders = basetypes.NewListNull(types.Int64Type)
	}

	if resp.DefaultScheduleDescription != nil {
		in := stringMapToValueMap(*resp.DefaultScheduleDescription)
		modelValue, diags := types.MapValue(types.StringType, in)
		if diags.HasError() {
			return diags
		}
		template.DefaultScheduleDescription = modelValue
	}

	if template.Inputs == nil {
		template.Inputs = &CustomFormModel{}
	}
	diags = template.Inputs.fromApiResponse(resp.Form.Get())

	return
}

func (r *AppointmentTemplate) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {

	// AppointmentTemplateModel describes the resource data model with the old custom form model
	type AppointmentTemplateModelV0 struct {
		ID                         types.String       `tfsdk:"id"`
		Name                       types.String       `tfsdk:"name"`
		NameExpression             types.String       `tfsdk:"name_expression"`
		Description                types.String       `tfsdk:"description"`
		Inputs                     *CustomFormModelV0 `tfsdk:"inputs"`
		RequireConfirmation        types.Bool         `tfsdk:"require_confirmation"`
		ConfirmationReminders      types.List         `tfsdk:"confirmation_reminders"`
		DefaultScheduleDescription types.Map          `tfsdk:"default_schedule_description"`
		InvitationText             types.String       `tfsdk:"invitation_text"`
		Icon                       types.String       `tfsdk:"icon"`
		Version                    types.Int32        `tfsdk:"version"`
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
					"name_expression": schema.StringAttribute{
						Optional: true,
					},
					"description": schema.StringAttribute{
						Optional: true,
					},
					"inputs": schema.SingleNestedAttribute{
						Required:   true,
						Attributes: CustomFormNestedSchemaV0(), // old custom form schema
					},
					"require_confirmation": schema.BoolAttribute{
						Optional: true,
					},
					"confirmation_reminders": schema.ListAttribute{
						Optional:    true,
						ElementType: types.Int64Type,
					},
					"default_schedule_description": schema.MapAttribute{
						Optional:    true,
						ElementType: types.StringType,
					},
					"invitation_text": schema.StringAttribute{
						Optional: true,
					},
					"icon": schema.StringAttribute{
						Optional: true,
					},
					"version": schema.Int32Attribute{
						Optional: true,
					},
				},
			},
			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var priorStateData AppointmentTemplateModelV0

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

				upgradedStateData := AppointmentTemplateModel{
					ID:                         priorStateData.ID,
					Name:                       priorStateData.Name,
					NameExpression:             priorStateData.NameExpression,
					Description:                priorStateData.Description,
					Inputs:                     upgradedInputs,
					RequireConfirmation:        priorStateData.RequireConfirmation,
					ConfirmationReminders:      priorStateData.ConfirmationReminders,
					DefaultScheduleDescription: priorStateData.DefaultScheduleDescription,
					InvitationText:             priorStateData.InvitationText,
					Icon:                       priorStateData.Icon,
					Version:                    priorStateData.Version,
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, upgradedStateData)...)
			},
		},
	}
}
