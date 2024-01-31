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
	_ resource.Resource                = &CaseReportingEndpoint{}
	_ resource.ResourceWithConfigure   = &CaseReportingEndpoint{}
	_ resource.ResourceWithImportState = &CaseReportingEndpoint{}
)

func NewCaseReportingEndpoint() resource.Resource {
	return &CaseReportingEndpoint{}
}

// CaseReportingEndpoint defines the resource implementation.
type CaseReportingEndpoint struct {
	client *client.UdomaClient
}

// CaseReportingEndpointModel describes the resource data model.
type CaseReportingEndpointModel struct {
	ID            types.String   `tfsdk:"id"`
	Code          types.String   `tfsdk:"code"`
	CreatedAt     types.Int64    `tfsdk:"created_at"`
	UpdatedAt     types.Int64    `tfsdk:"updated_at"`
	Name          types.String   `tfsdk:"name"`
	Active        types.Bool     `tfsdk:"active"`
	CaseTemplates []types.String `tfsdk:"case_templates"`
	// PropertyRefs []string `tfsdk:"property_refs"`
	LastUpdated types.String `tfsdk:"last_updated"`
}

func (r *CaseReportingEndpoint) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_case_reporting_endpoint"
}

func (r *CaseReportingEndpoint) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "Resource represents an endpoint where tenants can raise cases",

		Attributes: map[string]schema.Attribute{
			// required by terraform, set to code
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"code": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The unique code used for accessing the endpoint",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The date and time the endpoint was created",
			},
			"updated_at": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The date and time the endpoint was last modified",
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The name of the endpoint, shown in the admin page",
			},
			"active": schema.BoolAttribute{
				Optional:            true,
				MarkdownDescription: "Whether the endpoint is active or not",
			},
			"case_templates": schema.ListAttribute{
				Required:            true,
				MarkdownDescription: "Reference to the case templates that can be raised via this endpoint",
				ElementType:         types.StringType,
			},
		},
	}
}

func (r *CaseReportingEndpoint) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CaseReportingEndpoint) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Retrieve values from plan
	var plan CaseReportingEndpointModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Case Reporting Endpoint",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newEndpoint, _, err := r.client.GetApi().CreateCaseReportingEndpoint(ctx).CreateOrUpdateCaseReportingEndpointRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Case Reporting Endpoint",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newEndpoint); err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Case Reporting Endpoint",
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

func (r *CaseReportingEndpoint) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	// Retrieve values from state
	var state CaseReportingEndpointModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	endpoint, httpResp, err := r.client.GetApi().GetCaseReportingEndpoint(ctx, state.Code.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Case Reporting Endpoint",
			"Could not read entity from Udoma, unexpected error: "+err.Error(),
		)
		return
	}

	// update the tf struct with the new values
	if err := state.fromAPI(endpoint); err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Case Reporting Endpoint",
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

func (r *CaseReportingEndpoint) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan CaseReportingEndpointModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Case Reporting Endpoint",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newEndpoint, _, err := r.client.GetApi().UpdateCaseReportingEndpoint(ctx, plan.Code.ValueString()).CreateOrUpdateCaseReportingEndpointRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Case Reporting Endpoint",
			"Could not update entity in Udoma, unexpected error: "+err.Error(),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newEndpoint); err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Case Reporting Endpoint",
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

func (r *CaseReportingEndpoint) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	// Retrieve values from state
	var state CaseReportingEndpointModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.GetApi().DeleteCaseReportingEndpoint(ctx, state.Code.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Case Reporting Endpoint",
			"Could not delete entity in Udoma, unexpected error: "+err.Error(),
		)
		return
	}
}

func (r *CaseReportingEndpoint) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("code"), req, resp)
}

func (model *CaseReportingEndpointModel) fromAPI(endpoint *api.CaseReportingEndpoint) error {

	if endpoint == nil {
		return fmt.Errorf("endpoint is nil")
	}

	model.ID = types.StringValue(sdp(endpoint.Code))
	model.Code = types.StringValue(sdp(endpoint.Code))
	model.CreatedAt = types.Int64Value(idp(endpoint.CreatedAt))
	model.UpdatedAt = types.Int64Value(idp(endpoint.UpdatedAt))
	model.Name = types.StringValue(sdp(endpoint.Name))
	model.Active = types.BoolValue(bdp(endpoint.Active))
	model.CaseTemplates = []types.String{}
	for _, ct := range endpoint.CaseTemplates {
		model.CaseTemplates = append(model.CaseTemplates, types.StringValue(ct))
	}

	return nil
}

func (model *CaseReportingEndpointModel) toAPIRequest() (api.CreateOrUpdateCaseReportingEndpointRequest, error) {

	endpoint := api.CreateOrUpdateCaseReportingEndpointRequest{
		Name:          model.Name.ValueStringPointer(),
		Active:        model.Active.ValueBoolPointer(),
		CaseTemplates: []string{},
	}

	for _, ct := range model.CaseTemplates {
		endpoint.CaseTemplates = append(endpoint.CaseTemplates, ct.ValueString())
	}

	return endpoint, nil
}
