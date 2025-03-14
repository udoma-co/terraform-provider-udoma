package provider

import (
	"context"
	"fmt"

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
)

var (
	_ resource.ResourceWithConfigure   = &Hook{}
	_ resource.ResourceWithImportState = &Hook{}
)

func NewHook() resource.Resource {
	return &FAQ{}
}

type Hook struct {
	client *client.UdomaClient
}

type HookModel struct {
	ID           types.String `tfsdk:"id"`
	Entity       types.String `tfsdk:"entity"`
	Action       types.String `tfsdk:"action"`
	Priority     types.Int32  `tfsdk:"priority"`
	Enabled      types.Bool   `tfsdk:"enabled"`
	BreakOnError types.Bool   `tfsdk:"break_on_error"`
	Script       types.String `tfsdk:"script"`
}

func (hook *Hook) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hook"
}

func (hook *Hook) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents a hook",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the faq",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"entity": schema.StringAttribute{
				Required:    true,
				Description: "The type of entity the hook is targeting",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"CASE",
						"BANK_ACCOUNT",
						"PROPERTY",
						"PROPERTY_OWNER",
						"TENANT",
						"TENANCY",
						"SERVICE_PROVIDER",
					),
				},
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "The action that should triger the hook",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"CREATE",
						"UPDATE",
						"DELETE",
					),
				},
			},
			"prioirty": schema.Int32Attribute{
				Required:    true,
				Description: "The priority of the hook",
			},
			"break_on_error": schema.BoolAttribute{
				Required:    true,
				Description: "Whether the hook should break on error or not",
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether the hook is enabled or not",
			},
		},
	}
}

func (hook *Hook) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	hook.client = cl
}

func (hook *Hook) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Retrieve values from plan
	var plan HookModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := plan.toAPIRequest()

	newHook, _, err := hook.client.GetApi().CreateHook(ctx).CreateOrUpdateHookRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating FAQ Entry",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	plan.fromAPI(newHook)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (hook *Hook) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state HookModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	gotHook, httpResp, err := hook.client.GetApi().GetHook(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading FAQ Entry",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	state.fromAPI(gotHook)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (hook *Hook) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan HookModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := plan.toAPIRequest()

	newHook, _, err := hook.client.GetApi().UpdateHook(ctx, plan.ID.ValueString()).CreateOrUpdateHookRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating FAQ Entry",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	plan.fromAPI(newHook)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (hook *Hook) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	// Retrieve values from state
	var state HookModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := hook.client.GetApi().DeleteHook(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting FAQ Entry",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (hook *Hook) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (hook *HookModel) toAPIRequest() api.CreateOrUpdateHookRequest {
	return api.CreateOrUpdateHookRequest{
		Entity:   api.HookEntity(hook.Entity.ValueString()),
		Action:   api.HookAction(hook.Action.ValueString()),
		Script:   hook.Script.ValueString(),
		Priority: hook.Priority.ValueInt32(),
		Enabled:  hook.Enabled.ValueBool(),
	}
}

func (hook *HookModel) fromAPI(resp *api.Hook) {
	hook.ID = types.StringValue(resp.Id)
	hook.Entity = types.StringValue(string(resp.Entity))
	hook.Action = types.StringValue(string(resp.Action))
	hook.Script = types.StringValue(resp.Script)
	hook.Priority = types.Int32Value(resp.Priority)
	hook.Enabled = types.BoolValue(resp.Enabled)
}
