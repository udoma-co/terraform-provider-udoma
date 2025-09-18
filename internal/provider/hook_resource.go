package provider

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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
	return &Hook{}
}

type Hook struct {
	client *client.UdomaClient
}

type HookModel struct {
	ID             types.String         `tfsdk:"id"`
	Name           types.String         `tfsdk:"name"`
	Entity         types.String         `tfsdk:"entity"`
	RunOnCreate    types.Bool           `tfsdk:"run_on_create"`
	RunOnUpdate    types.Bool           `tfsdk:"run_on_update"`
	RunOnDelete    types.Bool           `tfsdk:"run_on_delete"`
	Pre            types.Bool           `tfsdk:"pre"`
	Post           types.Bool           `tfsdk:"post"`
	Priority       types.Int32          `tfsdk:"priority"`
	Enabled        types.Bool           `tfsdk:"enabled"`
	BreakOnError   types.Bool           `tfsdk:"break_on_error"`
	Script         types.String         `tfsdk:"script"`
	AdditionalData jsontypes.Normalized `tfsdk:"additional_data"`
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
				Description: "The unique identifier for the hook",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "A name used to identify the hook, mainly for debugging.",
			},
			"entity": schema.StringAttribute{
				Required:    true,
				Description: "The type of entity the hook is targeting",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"BANK_ACCOUNT",
						"PROPERTY",
						"PROPERTY_OWNER",
						"TENANT",
						"TENANCY",
						"SERVICE_PROVIDER",
					),
				},
			},
			"run_on_create": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether the hook should run on entity create",
			},
			"run_on_update": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether the hook should run on entity update",
			},
			"run_on_delete": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether the hook should run on entity delete",
			},
			"pre": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether to run the script before the action, both pre and post and can true then the script runs twice. At least one must be true.",
			},
			"post": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether to run the script after the action, both pre and post and can true then the script runs twice. At least one must be true.",
			},
			"priority": schema.Int32Attribute{
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
			"script": schema.StringAttribute{
				Required:    true,
				Description: "The script that should be executed when the hook is triggered",
			},
			"additional_data": schema.StringAttribute{
				Optional:    true,
				CustomType:  jsontypes.NormalizedType{},
				Description: "Additional data that can be used by the hook",
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

	createReq, moreDiags := plan.toAPIRequest()
	diags.Append(moreDiags...)
	if diags.HasError() {
		return
	}

	if (createReq.Pre == nil || !*createReq.Pre) && (createReq.Post == nil || !*createReq.Post) {
		diags.AddError(
			"Invalid Hook Configuration",
			"At least one of 'pre' or 'post' must be true.",
		)
	}

	newHook, _, err := hook.client.GetApi().CreateHook(ctx).CreateOrUpdateHookRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Hook Entry",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err = plan.fromAPI(newHook); err != nil {
		diags.AddError(
			"could not read API response",
			err.Error(),
		)
	}

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
			"Error Reading Hook Entry",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err = state.fromAPI(gotHook); err != nil {
		diags.AddError(
			"could not read API response",
			err.Error(),
		)
	}

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

	updateReq, moreDiags := plan.toAPIRequest()
	diags.Append(moreDiags...)
	if diags.HasError() {
		return
	}

	if (updateReq.Pre == nil || !*updateReq.Pre) && (updateReq.Post == nil || !*updateReq.Post) {
		diags.AddError(
			"Invalid Hook Configuration",
			"At least one of 'pre' or 'post' must be true.",
		)
	}

	newHook, _, err := hook.client.GetApi().UpdateHook(ctx, plan.ID.ValueString()).CreateOrUpdateHookRequest(updateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Hook Entry",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err = plan.fromAPI(newHook); err != nil {
		diags.AddError(
			"could not read API response",
			err.Error(),
		)
	}

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
			"Error Deleting Hook Entry",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (hook *Hook) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (hook *HookModel) toAPIRequest() (api.CreateOrUpdateHookRequest, diag.Diagnostics) {
	result := api.CreateOrUpdateHookRequest{
		Name:         hook.Name.ValueString(),
		Entity:       api.HookEntity(hook.Entity.ValueString()),
		RunOnCreate:  hook.RunOnCreate.ValueBoolPointer(),
		RunOnUpdate:  hook.RunOnUpdate.ValueBoolPointer(),
		RunOnDelete:  hook.RunOnDelete.ValueBoolPointer(),
		Pre:          hook.Pre.ValueBoolPointer(),
		Post:         hook.Post.ValueBoolPointer(),
		Script:       hook.Script.ValueString(),
		Priority:     hook.Priority.ValueInt32Pointer(),
		Enabled:      hook.Enabled.ValueBoolPointer(),
		BreakOnError: hook.BreakOnError.ValueBoolPointer(),
	}

	diags := hook.AdditionalData.Unmarshal(&result.AdditionalData)

	return result, diags
}

func (hook *HookModel) fromAPI(resp *api.Hook) (err error) {
	hook.ID = types.StringValue(resp.Id)
	hook.Name = types.StringValue(resp.Name)
	hook.Entity = types.StringValue(string(resp.Entity))
	hook.RunOnCreate = omittableBooleanValue(resp.RunOnCreate, hook.RunOnCreate)
	hook.RunOnUpdate = omittableBooleanValue(resp.RunOnUpdate, hook.RunOnUpdate)
	hook.RunOnDelete = omittableBooleanValue(resp.RunOnDelete, hook.RunOnDelete)
	hook.Pre = omittableBooleanValue(resp.Pre, hook.Pre)
	hook.Post = omittableBooleanValue(resp.Post, hook.Post)
	hook.Script = types.StringValue(resp.Script)
	hook.Priority = omittableInt32Value(resp.Priority, hook.Priority)
	hook.Enabled = omittableBooleanValue(resp.Enabled, hook.Enabled)
	hook.BreakOnError = omittableBooleanValue(resp.BreakOnError, hook.BreakOnError)

	bAdditionalData, err := json.Marshal(resp.AdditionalData)
	if err != nil {
		return fmt.Errorf("could not marshal additional_data: %w", err)
	}

	hook.AdditionalData = jsontypes.NewNormalizedValue(string(bAdditionalData))

	return
}
