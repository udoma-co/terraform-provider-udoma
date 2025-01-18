package provider

import (
	"context"
	"fmt"

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
	_ resource.ResourceWithConfigure   = &Notification{}
	_ resource.ResourceWithImportState = &Notification{}
)

func NewNotification() resource.Resource {
	return &Notification{}
}

type Notification struct {
	client *client.UdomaClient
}

type NotificationModel struct {
	Name            types.String `tfsdk:"name"`
	Description     types.String `tfsdk:"description"`
	Script          types.String `tfsdk:"script"`
	IsIntermediate  types.Bool   `tfsdk:"is_intermediate"`
	IntermediateRef types.String `tfsdk:"intermediate_ref"`
	TemplateHtml    types.String `tfsdk:"template_html"`
	TemplateText    types.String `tfsdk:"template_text"`
}

func (n *Notification) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_notification"
}

func (n *Notification) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the notification.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "The description of the notification.",
			},
			"script": schema.StringAttribute{
				Optional:    true,
				Description: "The script of the notification.",
			},
			"is_intermediate": schema.BoolAttribute{
				Optional: true,
				Description: `Whether this template is used an intermediate template or not. 
					If true it cannot be used as a regular template and must only be referenced 
					by other templates.`,
			},
			"intermediate_ref": schema.StringAttribute{
				Optional:    true,
				Description: "Optional reference to a intermediate template.",
			},
			"template_html": schema.StringAttribute{
				Required:    true,
				Description: "The html template of the notification.",
			},
			"template_text": schema.StringAttribute{
				Required:    true,
				Description: "The text template of the notification.",
			},
		},
	}
}

func (n *Notification) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

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

	n.client = cl
}

func (r *Notification) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan NotificationModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	notificationReq := plan.toCreateApiRequest()
	newNotification, _, err := r.client.GetApi().CreateNotification(ctx).CreateNotificationRequest(notificationReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Notification",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	plan.fromApiResponse(*newNotification)
	resp.Diagnostics.Append(
		resp.State.Set(ctx, &plan)...,
	)
}

func (r *Notification) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var plan NotificationModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	notificationReq := plan.toUpdateApiRequest()
	updatedNotification, _, err := r.client.GetApi().UpdateNotification(ctx, plan.Name.ValueString()).UpdateNotificationRequest(notificationReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Notification",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	plan.fromApiResponse(*updatedNotification)

	resp.Diagnostics.Append(
		resp.State.Set(ctx, &plan)...,
	)
}

func (r *Notification) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state NotificationModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newNotification, httpResp, err := r.client.GetApi().GetNotification(ctx, state.Name.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Notification",
			fmt.Sprintf("Could not read entity in Udoma, unexpected error: %s", getApiErrorMessage(err)),
		)
		return
	}

	state.fromApiResponse(*newNotification)

	resp.Diagnostics.Append(
		resp.State.Set(ctx, state)...,
	)
}

func (r *Notification) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state NotificationModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.client.GetApi().DeleteNotification(ctx, state.Name.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Notification",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err))
	}
}

func (n *NotificationModel) toCreateApiRequest() v1.CreateNotificationRequest {
	return v1.CreateNotificationRequest{
		Name:            n.Name.ValueString(),
		Description:     n.Description.ValueStringPointer(),
		IsIntermediate:  n.IsIntermediate.ValueBoolPointer(),
		IntermediateRef: n.IntermediateRef.ValueStringPointer(),
		Script:          n.Script.ValueStringPointer(),
		TemplateHtml:    n.TemplateHtml.ValueString(),
		TemplateText:    n.TemplateText.ValueString(),
	}
}

func (n *NotificationModel) toUpdateApiRequest() v1.UpdateNotificationRequest {
	return v1.UpdateNotificationRequest{
		IntermediateRef: n.IntermediateRef.ValueStringPointer(),
		Description:     n.Description.ValueStringPointer(),
		Script:          n.Script.ValueStringPointer(),
		TemplateHtml:    n.TemplateHtml.ValueString(),
		TemplateText:    n.TemplateText.ValueString(),
	}
}

func (n *NotificationModel) fromApiResponse(notification v1.Notification) {
	n.Name = types.StringValue(notification.Name)
	n.Description = omittableStringValue(notification.Description, n.Description)
	n.IsIntermediate = omittableBooleanValue(notification.IsIntermediate, n.IsIntermediate)
	n.IntermediateRef = omittableStringValue(notification.IntermediateRef, n.IntermediateRef)
	n.Script = omittableStringValue(notification.Script, n.Script)
	n.TemplateHtml = types.StringValue(notification.TemplateHtml)
	n.TemplateText = types.StringValue(notification.TemplateText)
}

func (n *Notification) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("name"), req, resp)
}
