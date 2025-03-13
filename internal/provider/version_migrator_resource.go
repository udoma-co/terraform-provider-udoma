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
	api "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"
)

var (
	_ resource.ResourceWithConfigure   = &VersionMigrator{}
	_ resource.ResourceWithImportState = &VersionMigrator{}
)

func NewVersionMigrator() resource.Resource {
	return &VersionMigrator{}
}

type VersionMigrator struct {
	client *client.UdomaClient
}

type VersionMigratorModel struct {
	ID            types.String `tfsdk:"id"`
	CreatedAt     types.Int64  `tfsdk:"created_at"`
	UpdatedAt     types.Int64  `tfsdk:"updated_at"`
	RefId         types.String `tfsdk:"ref_id"`
	SourceVersion types.Int32  `tfsdk:"source_version"`
	TargetVersion types.Int32  `tfsdk:"target_version"`
	Script        types.String `tfsdk:"script"`
}

func (migrator *VersionMigrator) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_version_migrator"
}

func (migrator *VersionMigrator) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents a version migrator",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the faq",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the version migrators was created",
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the version migrator was last modified",
			},
			"ref_id": schema.StringAttribute{
				Required:    true,
				Description: "The id of the template for the version migrator",
			},
			"source_version": schema.Int32Attribute{
				Required:    true,
				Description: "The version of the entity to be updated",
			},
			"target_version": schema.Int32Attribute{
				Required:    true,
				Description: "The version of the template that the entity has to be migrated to",
			},
			"script": schema.StringAttribute{
				Required:    true,
				Description: "The script that is going to hanlde the migration",
			},
		},
	}
}

func (migrator *VersionMigrator) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	migrator.client = cl
}

func (migrator *VersionMigrator) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Retrieve values from plan
	var plan VersionMigratorModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Version Migrator Request",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newMigrator, _, err := migrator.client.GetApi().CreateVersionMigrator(ctx).CreateOrUpdateVersionMigratorRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Version Migrator",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newMigrator); err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Version Migrator",
			"Could not process API response, unexpected error: "+err.Error(),
		)
		return
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (migrator *VersionMigrator) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state VersionMigratorModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	versionMigrator, httpResp, err := migrator.client.GetApi().GetVersionMigrator(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Version Migrator",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := state.fromAPI(versionMigrator); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Version Migrator",
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

func (migrator *VersionMigrator) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan VersionMigratorModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Version Migrator",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newMigrator, _, err := migrator.client.GetApi().UpdateVersionMigrator(ctx, plan.ID.ValueString()).CreateOrUpdateVersionMigratorRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Version Migrator",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newMigrator); err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Version Migrator",
			"Could not process API response, unexpected error: "+err.Error(),
		)
		return
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (migrator *VersionMigrator) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	// Retrieve values from state
	var state VersionMigratorModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := migrator.client.GetApi().DeleteVersionMigrator(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Version Migrator",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (migrator *VersionMigrator) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *VersionMigratorModel) fromAPI(migrator *api.VersionMigrator) error {

	if migrator == nil {
		return fmt.Errorf("version migrator is nil")
	}

	model.ID = types.StringValue(migrator.Id)
	model.RefId = types.StringValue(migrator.RefId)
	model.SourceVersion = types.Int32Value(migrator.SourceVersion)
	model.TargetVersion = types.Int32Value(migrator.TargetVersion)
	model.Script = types.StringValue(migrator.Script)
	model.CreatedAt = types.Int64Value(migrator.CreatedAt)
	model.UpdatedAt = types.Int64Value(migrator.UpdatedAt)

	return nil
}

func (model *VersionMigratorModel) toAPIRequest() (api.CreateOrUpdateVersionMigratorRequest, error) {

	migrator := api.CreateOrUpdateVersionMigratorRequest{
		RefId:         model.RefId.ValueString(),
		TargetVersion: model.TargetVersion.ValueInt32(),
		SourceVersion: model.SourceVersion.ValueInt32(),
		Script:        model.Script.ValueString(),
	}

	return migrator, nil
}
