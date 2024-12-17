package provider

import (
	"context"
	"encoding/json"
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
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/tf"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &workflowDefinition{}
	_ resource.ResourceWithConfigure   = &workflowDefinition{}
	_ resource.ResourceWithImportState = &workflowDefinition{}
)

func NewWorkflowDefinition() resource.Resource {
	return &workflowDefinition{}
}

// workflowDefinition defines the resource implementation.
type workflowDefinition struct {
	client *client.UdomaClient
}

// workflowDefinitionModel describes the resource data model.
type workflowDefinitionModel struct {
	ID             types.String       `tfsdk:"id"`
	LastUpdated    types.String       `tfsdk:"last_updated"`
	CreatedAt      types.Int64        `tfsdk:"created_at"`
	UpdatedAt      types.Int64        `tfsdk:"updated_at"`
	Name           types.String       `tfsdk:"name"`
	Description    types.String       `tfsdk:"description"`
	Icon           types.String       `tfsdk:"icon"`
	NameExpression types.String       `tfsdk:"name_expression"`
	EnvVars        types.Map          `tfsdk:"env_vars"`
	FirstStepID    types.String       `tfsdk:"first_step_id"`
	InitStep       tf.JsonObjectValue `tfsdk:"init_step"`
	Steps          tf.JsonObjectValue `tfsdk:"steps"`
}

func NewWorkflowDefinitionModelNull() *workflowDefinitionModel {
	return &workflowDefinitionModel{
		InitStep: tf.NewJsonObjectNull(),
		Steps:    tf.NewJsonObjectNull(),
	}
}

func (r *workflowDefinition) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workflow_definition"
}

func (r *workflowDefinition) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "Resource represents a defintion of a workflow",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the workflow definition",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the workflow definition was created",
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the workflow definition was last modified",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the workflow definition, shown in the admin page",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "The description of the workflow definition",
			},
			"icon": schema.StringAttribute{
				Optional:    true,
				Description: "The icon to be displayed on the manual workflow execution page",
			},
			"name_expression": schema.StringAttribute{
				Optional: true,
				Description: `An optional JS expression to be used to compute the name of 
				the workflow execution. If not set, the name of the definition will be used
				for new executions`,
			},
			"env_vars": schema.MapAttribute{
				Optional:    true,
				Description: "A map of environment variables to be set for the workflow execution",
				ElementType: types.StringType,
			},
			"first_step_id": schema.StringAttribute{
				Required:    true,
				Description: "The ID of the first step of the workflow",
			},
			"init_step": schema.StringAttribute{
				CustomType:  tf.JsonObjectType{},
				Optional:    true,
				Description: "Optional JSON serialised initial step definition",
			},
			"steps": schema.StringAttribute{
				CustomType:  tf.JsonObjectType{},
				Optional:    true,
				Description: "The JSON serialised step definitions",
			},
		},
	}
}

func (r *workflowDefinition) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *workflowDefinition) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Retrieve values from plan
	var plan workflowDefinitionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Workflow Definition",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newlyCreatedWorkflowDefinition, _, err := r.client.GetApi().CreateWorkflowDefinition(ctx).CreateOrUpdateWorkflowDefinitionRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Workflow Definition",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newlyCreatedWorkflowDefinition); err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Workflow Definition",
			"Could not process API response, unexpected error: "+err.Error(),
		)
		return
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *workflowDefinition) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	// Retrieve values from state
	var state workflowDefinitionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	workflow, httpResp, err := r.client.GetApi().GetWorkflowDefinition(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Workflow Definition",
			"Could not read entity from Udooma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := state.fromAPI(workflow); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Workflow Definition",
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

func (r *workflowDefinition) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan workflowDefinitionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Workflow Definition",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	id := plan.ID.ValueString()

	updatedWorkflowDefinition, _, err := r.client.GetApi().UpdateWorkflowDefinition(ctx, id).CreateOrUpdateWorkflowDefinitionRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Workflow Definition",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(updatedWorkflowDefinition); err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Workflow Definition",
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

func (r *workflowDefinition) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	// Retrieve values from state
	var state workflowDefinitionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.client.GetApi().DeleteWorkflowDefinition(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Workflow Definition",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (r *workflowDefinition) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *workflowDefinitionModel) fromAPI(workflowDefinition *api.WorkflowDefinition) error {

	if workflowDefinition == nil {
		return fmt.Errorf("workflow definition is nil")
	}

	model.ID = types.StringValue(workflowDefinition.Id)
	model.CreatedAt = types.Int64Value(workflowDefinition.CreatedAt)
	model.UpdatedAt = types.Int64Value(workflowDefinition.UpdatedAt)
	model.Name = types.StringValue(workflowDefinition.Name)
	model.Description = omittableStringValue(workflowDefinition.Description, model.Description)
	model.Icon = omittableStringValue(workflowDefinition.Icon, model.Icon)
	model.NameExpression = omittableStringValue(workflowDefinition.NameExpression, model.NameExpression)
	model.FirstStepID = types.StringValue(workflowDefinition.FirstStepId)

	if workflowDefinition.EnvVars != nil {
		in := stringMapToValueMap(*workflowDefinition.EnvVars)
		modelValue, diags := types.MapValue(types.StringType, in)
		if diags.HasError() {
			return fmt.Errorf("error converting env vars to map: %v", diags)
		}
		model.EnvVars = modelValue
	}

	if workflowDefinition.InitStep.IsSet() {
		step := workflowDefinition.InitStep.Get()
		jsonData, err := json.Marshal(step)
		if err != nil {
			return fmt.Errorf("failed to marshal init step: %w", err)
		}
		model.InitStep = tf.NewJsonObjectValue(string(jsonData))
	}

	steps, err := json.Marshal(workflowDefinition.Steps)
	if err != nil {
		return fmt.Errorf("failed to marshal steps: %w", err)
	}
	model.Steps = tf.NewJsonObjectValue(string(steps))

	return nil
}

func (model *workflowDefinitionModel) toAPIRequest() (api.CreateOrUpdateWorkflowDefinitionRequest, error) {

	workflowDefinition := api.CreateOrUpdateWorkflowDefinitionRequest{
		Name:           model.Name.ValueString(),
		Description:    model.Description.ValueStringPointer(),
		Icon:           model.Icon.ValueStringPointer(),
		NameExpression: model.NameExpression.ValueStringPointer(),
		FirstStepId:    model.FirstStepID.ValueString(),
	}

	if envVars := modelMapToStringMap(model.EnvVars); len(envVars) > 0 {
		workflowDefinition.EnvVars = &envVars
	}

	if !model.InitStep.IsNull() && !model.InitStep.IsUnknown() {
		var jsonVal *api.WorkflowInitStepDefinition
		if diag := model.InitStep.Unmarshal(&jsonVal); diag.HasError() {
			return workflowDefinition, fmt.Errorf("failed to unmarshal initial step: %v", diag)
		}
		if jsonVal != nil {
			workflowDefinition.InitStep = *api.NewNullableWorkflowInitStepDefinition(jsonVal)
		}
	}

	if !model.Steps.IsNull() && !model.Steps.IsUnknown() {
		if diag := model.Steps.Unmarshal(&workflowDefinition.Steps); diag.HasError() {
			return workflowDefinition, fmt.Errorf("failed to unmarshal steps: %v", diag)
		}
	}

	return workflowDefinition, nil
}
