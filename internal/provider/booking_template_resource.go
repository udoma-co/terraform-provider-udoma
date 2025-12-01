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
	_ resource.ResourceWithConfigure   = &BookingTemplate{}
	_ resource.ResourceWithImportState = &BookingTemplate{}
)

func NewBookingTemplate() resource.Resource {
	return &BookingTemplate{}
}

type BookingTemplate struct {
	client *client.UdomaClient
}

type BookingTemplateModel struct {
	ID          types.String     `tfsdk:"id"`
	CreatedAt   types.Int64      `tfsdk:"created_at"`
	UpdatedAt   types.Int64      `tfsdk:"updated_at"`
	Name        types.String     `tfsdk:"name"`
	Description types.String     `tfsdk:"description"`
	Icon        types.String     `tfsdk:"icon"`
	TriggerSource types.String   `tfsdk:"trigger_source"`
	Inputs      *CustomFormModel `tfsdk:"inputs"`
	InitScript  types.String     `tfsdk:"init_script"`
	Script      types.String     `tfsdk:"script"`
}

func (faq *BookingTemplate) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_booking_template"
}

func (faq *BookingTemplate) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents a booking template",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the booking template",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the booking template was created",
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the booking template was last modified",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the booking template",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "An optional user friendly label, used to identify the booking template",
			},
			"icon": schema.StringAttribute{
				Optional:    true,
				Description: "An optional icon for the booking template",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
			},
			"trigger_source": schema.StringAttribute{
				Optional:    true,
				Description: "The source from which this template can be triggered. Defaults to 'manual' if not specified",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"manual",
						"invoice",
						"bank_transaction",
					),
				},
			},
			"inputs": schema.SingleNestedAttribute{
				Required:    true,
				Description: "A custom form to collect data with",
				Attributes:  CustomFormNestedSchema(),
			},
			"init_script": schema.StringAttribute{
				Optional:    true,
				Description: "An optional script that can be used to prepare data and populate the custom form when the template is triggered",
			},
			"script": schema.StringAttribute{
				Required:    true,
				Description: "The JS script that generates the bookings",
			},
		},
	}
}

func (res *BookingTemplate) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	res.client = cl
}

func (res *BookingTemplate) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Retrieve values from plan
	var plan BookingTemplateModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toApiRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Booking Template Request",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newBookingTemplate, _, err := res.client.GetApi().CreateBookingTemplate(ctx).CreateOrUpdateBookingTemplateRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Booking Template Entry",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	diags = plan.fromApiResponse(newBookingTemplate)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (res *BookingTemplate) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state BookingTemplateModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	bookingTemplate, httpResp, err := res.client.GetApi().GetBookingTemplate(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Booking Template Entry",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	diags = state.fromApiResponse(bookingTemplate)
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

func (res *BookingTemplate) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan BookingTemplateModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq, err := plan.toApiRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Booking Template",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	updatedBookingTemplate, _, err := res.client.GetApi().UpdateBookingTemplate(ctx, plan.ID.ValueString()).CreateOrUpdateBookingTemplateRequest(updateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Booking Template",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	diags = plan.fromApiResponse(updatedBookingTemplate)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (res *BookingTemplate) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	// Retrieve values from state
	var state BookingTemplateModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := res.client.GetApi().DeleteBookingTemplate(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Booking Template Entry",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (res *BookingTemplate) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *BookingTemplateModel) fromApiResponse(bookingTemplate *v1.BookingTemplate) (diags diag.Diagnostics) {

	if bookingTemplate == nil {
		diags.AddError("Error Reading Booking Template", "Booking template is nil")
		return
	}

	model.ID = types.StringValue(bookingTemplate.Id)
	model.CreatedAt = types.Int64Value(bookingTemplate.CreatedAt)
	model.UpdatedAt = types.Int64Value(bookingTemplate.UpdatedAt)
	model.Name = types.StringValue(bookingTemplate.Name)
	model.Description = omittableStringValue(bookingTemplate.Description, model.Description)
	model.Icon = omittableStringValue(bookingTemplate.Icon, model.Icon)
	model.Script = types.StringValue(bookingTemplate.Script)
	model.InitScript = omittableStringValue(bookingTemplate.InitScript, model.InitScript)

	var triggerSourcePtr *string
	if bookingTemplate.TriggerSource != nil {
		enumValue := string(*bookingTemplate.TriggerSource)
		triggerSourcePtr = &enumValue
	}

	model.TriggerSource = omittableStringValue(triggerSourcePtr, model.TriggerSource)
	
	if model.Inputs == nil {
		model.Inputs = &CustomFormModel{}
	}
	diags = model.Inputs.fromApiResponse(bookingTemplate.Inputs.Get())

	return nil
}

func (model *BookingTemplateModel) toApiRequest() (v1.CreateOrUpdateBookingTemplateRequest, error) {
	form := model.Inputs.toApiRequest()

	bookingTemplate := v1.CreateOrUpdateBookingTemplateRequest{
		Name:        model.Name.ValueString(),
		Description: model.Description.ValueStringPointer(),
		Icon:        model.Icon.ValueStringPointer(),
		Inputs:      *v1.NewNullableCustomForm(&form),
		InitScript:  model.InitScript.ValueStringPointer(),
		Script:      model.Script.ValueString(),
		TriggerSource: func() *v1.BookingTemplateTriggerSourceEnum {
			if model.TriggerSource.ValueString() == "" {
				return nil
			}
			enumValue := v1.BookingTemplateTriggerSourceEnum(model.TriggerSource.ValueString())
			return &enumValue
		}(),
	}

	return bookingTemplate, nil
}

