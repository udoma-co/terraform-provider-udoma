package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	api "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &customerScript{}
	_ resource.ResourceWithConfigure   = &customerScript{}
	_ resource.ResourceWithImportState = &customerScript{}
)

func NewCustomerScript() resource.Resource {
	return &customerScript{}
}

// customerScript defines the resource implementation.
type customerScript struct {
	client *client.UdomaClient
}

// customerScriptModel describes the resource data model.
type customerScriptModel struct {
	ID          types.String `tfsdk:"id"`
	LastUpdated types.String `tfsdk:"last_updated"`
	CreatedAt   types.Int64  `tfsdk:"created_at"`
	UpdatedAt   types.Int64  `tfsdk:"updated_at"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Scope       types.String `tfsdk:"scope"`
	Script      types.String `tfsdk:"script"`
}

func (r *customerScript) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_customer_script"
}

func (r *customerScript) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "Resource represents a customer specific JS script",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the customer script",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the customer script was created",
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the customer script was last modified",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the customer script, shown in the admin page",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "The description of the customer script",
			},
			"scope": schema.StringAttribute{
				Optional:    true,
				Description: "The scope where the customer script will be available",
			},
			"script": schema.StringAttribute{
				Optional:    true,
				Description: "The actual JS script of the customer script",
			},
		},
	}
}

func (r *customerScript) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *customerScript) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Retrieve values from plan
	var plan customerScriptModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, _ := plan.toAPIRequest()

	newEndpoint, _, err := r.client.GetApi().CreateCustomerScript(ctx).CreateOrUpdateCustomerScriptRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Customer Script",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newEndpoint); err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Customer Script",
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

func (r *customerScript) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	// Retrieve values from state
	var state customerScriptModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	script, httpResp, err := r.client.GetApi().GetCustomerScript(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Customer Script",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := state.fromAPI(script); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Customer Script",
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

func (r *customerScript) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan customerScriptModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, _ := plan.toAPIRequest()

	id := plan.ID.ValueString()

	newEndpoint, _, err := r.client.GetApi().UpdateCustomerScript(ctx, id).CreateOrUpdateCustomerScriptRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Customer Script",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newEndpoint); err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Customer Script",
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

func (r *customerScript) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	// Retrieve values from state
	var state customerScriptModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.GetApi().DeleteCustomerScript(ctx, state.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Customer Script",
			"Could not delete entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}
}

func (r *customerScript) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *customerScriptModel) fromAPI(script *api.CustomerScript) error {

	if script == nil {
		return fmt.Errorf("customer script is nil")
	}

	model.ID = types.StringValue(sdp(script.Id))
	model.CreatedAt = types.Int64Value(idp(script.CreatedAt))
	model.UpdatedAt = types.Int64Value(idp(script.UpdatedAt))
	model.Name = types.StringValue(sdp(script.Name))
	model.Description = types.StringValue(sdp(script.Description))
	if script.Scope != nil {
		scope := *script.Scope
		model.Scope = types.StringValue(string(scope))
	}
	model.Script = types.StringValue(sdp(script.Script))

	return nil
}

func (model *customerScriptModel) toAPIRequest() (api.CreateOrUpdateCustomerScriptRequest, error) {

	script := api.CreateOrUpdateCustomerScriptRequest{
		Name:        model.Name.ValueStringPointer(),
		Description: model.Description.ValueStringPointer(),
		Script:      model.Script.ValueStringPointer(),
	}

	if !model.Scope.IsNull() && !model.Scope.IsUnknown() {
		scope := api.CustomerScriptScope(model.Scope.ValueString())
		script.Scope = &scope
	}

	return script, nil
}
