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
	_ resource.Resource                = &workflowEntrypoint{}
	_ resource.ResourceWithConfigure   = &workflowEntrypoint{}
	_ resource.ResourceWithImportState = &workflowEntrypoint{}
)

func NewWorkflowEntrypoint() resource.Resource {
	return &workflowEntrypoint{}
}

// workflowEntrypoint defines the resource implementation.
type workflowEntrypoint struct {
	client *client.UdomaClient
}

type workflowEntrypointFilter struct {
	Attribute types.String `tfsdk:"attribute"`
	Value     types.String `tfsdk:"value"`
}

type workflowEntrypointValidation struct {
	Name   types.String `tfsdk:"name"`
	Script types.String `tfsdk:"script"`
}

// workflowEntrypointModel describes the resource data model.
type workflowEntrypointModel struct {
	ID                    types.String                   `tfsdk:"id"`
	LastUpdated           types.String                   `tfsdk:"last_updated"`
	CreatedAt             types.Int64                    `tfsdk:"created_at"`
	UpdatedAt             types.Int64                    `tfsdk:"updated_at"`
	WorkflowDefinitionRef types.String                   `tfsdk:"workflow_definition_ref"`
	AppLocation           types.String                   `tfsdk:"app_location"`
	LocationFilters       []workflowEntrypointFilter     `tfsdk:"location_filters"`
	Validations           []workflowEntrypointValidation `tfsdk:"validations"`
	Icon                  types.String                   `tfsdk:"icon"`
	Label                 types.Map                      `tfsdk:"label"`
	InitScript            types.String                   `tfsdk:"init_script"`
}

func (r *workflowEntrypoint) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workflow_entrypoint"
}

func (r *workflowEntrypoint) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "Resource represents an entrypoint to a workflow",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the workflow entrypoint.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the workflow entrypoint was created",
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the workflow entrypoint was last modified",
			},
			"workflow_definition_ref": schema.StringAttribute{
				Required:    true,
				Description: "The ID of the workflow definition",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"app_location": schema.StringAttribute{
				Optional:    true,
				Description: "The location in the webapp where the workflow can be strated from",
			},
			"location_filters": schema.ListNestedAttribute{
				Optional:     true,
				NestedObject: workflowEntrypointFilterNestedSchema(),
				Description:  "Filters that can be used to limit where the entrypoint is shown",
			},
			"validations": schema.ListNestedAttribute{
				Optional:     true,
				NestedObject: workflowEntrypointValidationNestedSchema(),
				Description: `Optional validations that can be used to determine if the workflow can 
					be started via this entrypoint. Depending on the result, the workflow
					can be started or not.`,
			},
			"icon": schema.StringAttribute{
				Optional:    true,
				Description: "Optional icon to be displayed on the button that will start the workflow execution",
			},
			"label": schema.MapAttribute{
				Optional:    true,
				Description: "The label to be displayed on the button starting the workflow execution",
				ElementType: types.StringType,
			},
			"init_script": schema.StringAttribute{
				Optional:    true,
				Description: "Optional JS script to be executed before the workflow is started",
			},
		},
	}
}

func workflowEntrypointFilterNestedSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"attribute": schema.StringAttribute{
				Required:    true,
				Description: "The ID of the entity that will be referenced",
			},
			"value": schema.StringAttribute{
				Required:    true,
				Description: "The type of the entity that will be referenced",
			},
		},
	}
}

func workflowEntrypointValidationNestedSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the validation",
			},
			"script": schema.StringAttribute{
				Required:    true,
				Description: "JS script that will be executed to validate if the workflow can be started via this entrypoint",
			},
		},
	}
}

func (r *workflowEntrypoint) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *workflowEntrypoint) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Retrieve values from plan
	var plan workflowEntrypointModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, _ := plan.toAPIRequest()

	newlyCreatedWorkflowEntrypoint, _, err := r.client.GetApi().CreateWorkflowEntrypoint(ctx, createReq.WorkflowDefinitionRef).CreateOrUpdateWorkflowEntrypointRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Workflow Entrypoint",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newlyCreatedWorkflowEntrypoint); err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Workflow Entrypoint",
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

func (r *workflowEntrypoint) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	// Retrieve values from state
	var state workflowEntrypointModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	workflow, httpResp, err := r.client.GetApi().GetWorkflowEntrypoint(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Workflow Entrypoint",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := state.fromAPI(workflow); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Workflow Entrypoint",
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

func (r *workflowEntrypoint) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan workflowEntrypointModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, _ := plan.toAPIRequest()

	id := plan.ID.ValueString()

	updatedWorkflowEntrypoint, _, err := r.client.GetApi().UpdateWorkflowEntrypoint(ctx, id).CreateOrUpdateWorkflowEntrypointRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Workflow Entrypoint",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(updatedWorkflowEntrypoint); err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Workflow Entrypoint",
			"Could process API response, unexpected error: "+err.Error(),
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

func (r *workflowEntrypoint) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	// Retrieve values from state
	var state workflowEntrypointModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.client.GetApi().DeleteWorkflowEntrypoint(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Workflow Entrypoint",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (r *workflowEntrypoint) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (item *workflowEntrypointFilter) fromApi(resp *api.WorkflowEntrypointFilter) {
	item.Attribute = types.StringValue(resp.Attribute)
	item.Value = types.StringValue(resp.Value)
}

func (item *workflowEntrypointValidation) fromApi(resp *api.WorkflowEntrypointValidation) {
	item.Name = types.StringValue(resp.Name)
	item.Script = types.StringValue(resp.Script)
}

func (model *workflowEntrypointModel) fromAPI(workflowEntrypoint *api.WorkflowEntrypoint) error {

	if workflowEntrypoint == nil {
		return fmt.Errorf("workflow entrypoint is nil")
	}

	model.ID = types.StringValue(workflowEntrypoint.Id)
	model.CreatedAt = types.Int64Value(workflowEntrypoint.CreatedAt)
	model.UpdatedAt = types.Int64Value(workflowEntrypoint.UpdatedAt)
	model.WorkflowDefinitionRef = types.StringValue(workflowEntrypoint.WorkflowDefinitionRef)
	model.AppLocation = types.StringValue(string(workflowEntrypoint.AppLocation))
	model.Icon = omittableStringValue(workflowEntrypoint.Icon, model.Icon)

	in := stringMapToValueMap(workflowEntrypoint.Label)
	modelValue, diags := types.MapValue(types.StringType, in)
	if diags.HasError() {
		return fmt.Errorf("error converting label to map: %v", diags)
	}
	model.Label = modelValue

	model.InitScript = omittableStringValue(workflowEntrypoint.InitScript, model.InitScript)

	for i := range workflowEntrypoint.LocationFilters {
		if len(model.LocationFilters) <= i {
			model.LocationFilters = append(model.LocationFilters, workflowEntrypointFilter{})
		}
		model.LocationFilters[i].fromApi(&workflowEntrypoint.LocationFilters[i])
	}

	for i := range workflowEntrypoint.Validations {
		if len(model.Validations) <= i {
			model.Validations = append(model.Validations, workflowEntrypointValidation{})
		}
		model.Validations[i].fromApi(&workflowEntrypoint.Validations[i])
	}

	return nil
}

func (filter *workflowEntrypointFilter) toApiRequest() *api.WorkflowEntrypointFilter {
	return &api.WorkflowEntrypointFilter{
		Attribute: filter.Attribute.ValueString(),
		Value:     filter.Value.ValueString(),
	}
}

func (filter *workflowEntrypointValidation) toApiRequest() *api.WorkflowEntrypointValidation {
	return &api.WorkflowEntrypointValidation{
		Name:   filter.Name.ValueString(),
		Script: filter.Script.ValueString(),
	}
}

func (model *workflowEntrypointModel) toAPIRequest() (api.CreateOrUpdateWorkflowEntrypointRequest, error) {

	workflowEntrypoint := api.CreateOrUpdateWorkflowEntrypointRequest{
		WorkflowDefinitionRef: model.WorkflowDefinitionRef.ValueString(),
		Icon:                  model.Icon.ValueStringPointer(),
		Label:                 modelMapToStringMap(model.Label),
		InitScript:            model.InitScript.ValueStringPointer(),
	}

	filters := make([]api.WorkflowEntrypointFilter, len(model.LocationFilters))
	for i := range model.LocationFilters {
		filters[i] = *model.LocationFilters[i].toApiRequest()
	}
	workflowEntrypoint.LocationFilters = filters

	validations := make([]api.WorkflowEntrypointValidation, len(model.Validations))
	for i := range model.Validations {
		validations[i] = *model.Validations[i].toApiRequest()
	}
	workflowEntrypoint.Validations = validations

	if !model.AppLocation.IsNull() && !model.AppLocation.IsUnknown() {
		location := api.WorkflowEntrypointLocation(model.AppLocation.ValueString())
		workflowEntrypoint.AppLocation = location
	}

	return workflowEntrypoint, nil
}
