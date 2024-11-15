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
	_ resource.Resource                = &customIDGenerator{}
	_ resource.ResourceWithConfigure   = &customIDGenerator{}
	_ resource.ResourceWithImportState = &customIDGenerator{}
)

func NewCustomIDGenerator() resource.Resource {
	return &customIDGenerator{}
}

// customIDGenerator defines the resource implementation.
type customIDGenerator struct {
	client *client.UdomaClient
}

// customIDGeneratorModel describes the resource data model.
type customIDGeneratorModel struct {
	ID               types.String `tfsdk:"id"`
	LastUpdated      types.String `tfsdk:"last_updated"`
	CreatedAt        types.Int64  `tfsdk:"created_at"`
	UpdatedAt        types.Int64  `tfsdk:"updated_at"`
	Name             types.String `tfsdk:"name"`
	GenerationScript types.String `tfsdk:"generation_script"`
	LastGeneratedID  types.String `tfsdk:"last_generated_id"`
}

func (r *customIDGenerator) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_id_generator"
}

func (r *customIDGenerator) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "Resource represents a custom ID generator",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the custom ID generator",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the custom ID generatorwas created",
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the custom ID generatorwas last modified",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the custom ID generator",
			},
			"generation_script": schema.StringAttribute{
				Required:    true,
				Description: "The JS script used to generate a new unique ID",
			},
			"last_generated_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The last generated ID",
			},
		},
	}
}

func (r *customIDGenerator) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *customIDGenerator) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Retrieve values from plan
	var plan customIDGeneratorModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Custom ID Generator",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	generator, _, err := r.client.GetApi().CreateCustomIDGenerator(ctx).CreateOrUpdateCustomIDGeneratorRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Custom ID Generator",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(generator); err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Custom ID Generator",
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

func (r *customIDGenerator) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	// Retrieve values from state
	var state customIDGeneratorModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	generator, httpResp, err := r.client.GetApi().GetCustomIDGenerator(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Custom ID Generator",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := state.fromAPI(generator); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Custom ID Generator",
			"Could not create entity in Udoma, unexpected error: "+err.Error(),
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

func (r *customIDGenerator) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan customIDGeneratorModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, _ := plan.toAPIRequest()

	id := plan.ID.ValueString()

	newEndpoint, _, err := r.client.GetApi().UpdateCustomIDGenerator(ctx, id).CreateOrUpdateCustomIDGeneratorRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Custom ID Generator",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newEndpoint); err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Custom ID Generator",
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

func (r *customIDGenerator) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	// Retrieve values from state
	var state customIDGeneratorModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.client.GetApi().DeleteCustomIDGenerator(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Custom ID Generator",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (r *customIDGenerator) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *customIDGeneratorModel) fromAPI(generator *api.CustomIDGenerator) error {

	if generator == nil {
		return fmt.Errorf("custom ID generator is nil")
	}

	model.ID = types.StringValue(generator.Id)
	model.CreatedAt = types.Int64Value(generator.CreatedAt)
	model.UpdatedAt = types.Int64Value(generator.UpdatedAt)
	model.Name = types.StringValue(generator.Name)
	model.GenerationScript = types.StringValue(generator.GenerationScript)
	model.LastGeneratedID = omittableStringValue(generator.LastGeneratedId, model.LastGeneratedID)

	return nil
}

func (model *customIDGeneratorModel) toAPIRequest() (api.CreateOrUpdateCustomIDGeneratorRequest, error) {

	script := api.CreateOrUpdateCustomIDGeneratorRequest{
		Name:             model.Name.ValueString(),
		GenerationScript: model.GenerationScript.ValueString(),
		LastGeneratedId:  model.LastGeneratedID.ValueStringPointer(),
	}

	return script, nil
}
