package provider

import (
	"context"
	"fmt"
	"time"

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

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &reportEntrypoint{}
	_ resource.ResourceWithConfigure   = &reportEntrypoint{}
	_ resource.ResourceWithImportState = &reportEntrypoint{}
)

func NewReportEntrypoint() resource.Resource {
	return &reportEntrypoint{}
}

// reportEntrypoint defines the resource implementation.
type reportEntrypoint struct {
	client *client.UdomaClient
}

// reportEntrypointModel describes the resource data model.
type reportEntrypointModel struct {
	ID                  types.String `tfsdk:"id"`
	LastUpdated         types.String `tfsdk:"last_updated"`
	CreatedAt           types.Int64  `tfsdk:"created_at"`
	UpdatedAt           types.Int64  `tfsdk:"updated_at"`
	AppLocation         types.String `tfsdk:"app_location"`
	Icon                types.String `tfsdk:"icon"`
	Label               types.Map    `tfsdk:"label"`
	QuickExecute        types.Bool   `tfsdk:"quick_execute"`
	ReportDefinitionRef types.String `tfsdk:"report_definition_ref"`
}

func NewReportEntrypointModelNull() *reportEntrypointModel {
	return &reportEntrypointModel{}
}

func (r *reportEntrypoint) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_report_entrypoint"
}

func (r *reportEntrypoint) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents an entrypoint for a report execution",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the report entrypoint",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the report entrypoint was created",
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the report entrypoint was last modified",
			},
			"app_location": schema.StringAttribute{
				Required:    true,
				Description: "The location of the entrypoint in the app",
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(stringSlice(api.AllowedReportEntrypointLocationEnumValues)...),
				},
			},
			"icon": schema.StringAttribute{
				Optional:    true,
				Description: "The icon for the entrypoint",
			},
			"label": schema.MapAttribute{
				Required:    true,
				Description: "The label for the entrypoint, localized",
				ElementType: types.StringType,
			},
			"quick_execute": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether the report should be executed immediately",
			},
			"report_definition_ref": schema.StringAttribute{
				Required:    true,
				Description: "The reference to the report definition",
			},
		},
	}
}

func (r *reportEntrypoint) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *reportEntrypoint) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan reportEntrypointModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Report Entrypoint",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	// Use the report_definition_ref to construct the correct API endpoint
	definitionID := plan.ReportDefinitionRef.ValueString()
	newEntrypoint, _, err := r.client.GetApi().CreateReportEntrypoint(ctx, definitionID).CreateOrUpdateReportEntrypointRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Report Entrypoint",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// Update the Terraform state with the new values
	diags = plan.fromAPI(newEntrypoint)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *reportEntrypoint) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state reportEntrypointModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	entrypoint, httpResp, err := r.client.GetApi().GetReportEntrypoint(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Report Entrypoint",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	diags = state.fromAPI(entrypoint)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *reportEntrypoint) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan reportEntrypointModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Report Entrypoint",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	updatedEntrypoint, _, err := r.client.GetApi().UpdateReportEntrypoint(ctx, plan.ID.ValueString()).CreateOrUpdateReportEntrypointRequest(updateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Report Entrypoint",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	diags = plan.fromAPI(updatedEntrypoint)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *reportEntrypoint) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state reportEntrypointModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.client.GetApi().DeleteReportEntrypoint(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Report Entrypoint",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (r *reportEntrypoint) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *reportEntrypointModel) fromAPI(entrypoint *api.ReportEntrypoint) (diags diag.Diagnostics) {
	if entrypoint == nil {
		diags.AddError("Could not read report entrypoint from API", "report entrypoint is nil")
		return
	}

	model.ID = types.StringValue(entrypoint.Id)
	model.CreatedAt = types.Int64Value(entrypoint.CreatedAt)
	model.UpdatedAt = types.Int64Value(entrypoint.UpdatedAt)
	model.AppLocation = types.StringValue(string(entrypoint.AppLocation))
	model.Icon = omittableStringValue(entrypoint.Icon, model.Icon)
	model.QuickExecute = omittableBooleanValue(entrypoint.QuickExecute, model.QuickExecute)
	model.ReportDefinitionRef = types.StringValue(entrypoint.ReportDefinitionRef)

	if entrypoint.Label != nil {
		model.Label, diags = types.MapValue(types.StringType, stringMapToValueMap(entrypoint.Label))
		if diags.HasError() {
			return
		}
	} else {
		model.Label = types.MapNull(types.StringType)
	}

	return
}

func (model *reportEntrypointModel) toAPIRequest() (api.CreateOrUpdateReportEntrypointRequest, error) {
	entrypoint := api.CreateOrUpdateReportEntrypointRequest{
		AppLocation: api.ReportEntrypointLocation(model.AppLocation.ValueString()),
		Label:       modelMapToStringMap(model.Label),
	}

	if !model.Icon.IsNull() {
		icon := model.Icon.ValueString()
		entrypoint.Icon = &icon
	}

	if !model.QuickExecute.IsNull() {
		quickExecute := model.QuickExecute.ValueBool()
		entrypoint.QuickExecute = &quickExecute
	}

	return entrypoint, nil
}
