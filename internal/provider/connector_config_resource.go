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
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"

	api "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
)

var (
	_ resource.ResourceWithConfigure   = &ConnectorConfig{}
	_ resource.ResourceWithImportState = &ConnectorConfig{}
)

func NewConnectorConfig() resource.Resource {
	return &ConnectorConfig{}
}

type ConnectorConfig struct {
	client *client.UdomaClient
}

type ConnectorConfigModel struct {
	Name        types.String `tfsdk:"name"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	Description types.String `tfsdk:"description"`
	EnableLogs  types.Bool   `tfsdk:"enable_logs"`
	PingTime    types.String `tfsdk:"ping_time"`
	SyncTime    types.String `tfsdk:"sync_time"`
}

func (c *ConnectorConfig) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connector_config"
}

func (c *ConnectorConfig) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents a connector config",

		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the connector config",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.LengthAtMost(127),
				},
			},
			"enabled": schema.BoolAttribute{
				Required:    true,
				Description: "Whether the connector config is enabled",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "The description of the connector config",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(250),
				},
			},
			"enable_logs": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether to enable logs for the connector config",
			},
			"ping_time": schema.StringAttribute{
				Required:    true,
				Description: "The ping time of the connector config",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(25),
				},
			},
			"sync_time": schema.StringAttribute{
				Required:    true,
				Description: "The sync time of the connector config",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(25),
				},
			},
		},
	}
}

func (q *ConnectorConfig) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	q.client = cl
}

func (c *ConnectorConfig) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ConnectorConfigModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiReq := plan.toAPICreateRequest()
	config, _, err := c.client.GetApi().CreateConnectorConfig(ctx).CreateConnectorConfigRequest(apiReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Connector Config",
			fmt.Sprintf("Could not create entity in Udoma, unexpected error: %s", getApiErrorMessage(err)),
		)
		return
	}

	plan.fromApiResponse(config)
	resp.Diagnostics.Append(
		resp.State.Set(ctx, plan)...,
	)
}

func (c *ConnectorConfig) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ConnectorConfigModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newConfig, httpResp, err := c.client.GetApi().GetConnectorConfig(ctx, state.Name.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Connector Config",
			"Could not read entity from Udooma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	state.fromApiResponse(newConfig)
	resp.Diagnostics.Append(
		resp.State.Set(ctx, state)...,
	)
}

func (c *ConnectorConfig) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ConnectorConfigModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiReq := plan.toAPIUpdateRequest()
	config, _, err := c.client.GetApi().UpdateConnectorConfig(ctx, plan.Name.ValueString()).UpdateConnectorConfigRequest(apiReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Connector Config",
			fmt.Sprintf("Could not update entity in Udoma, unexpected error: %s", getApiErrorMessage(err)),
		)
		return
	}

	plan.fromApiResponse(config)
	resp.Diagnostics.Append(
		resp.State.Set(ctx, plan)...,
	)
}

func (c *ConnectorConfig) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ConnectorConfigModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := c.client.GetApi().DeleteConnectorConfig(ctx, state.Name.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Connector Config",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (r *ConnectorConfig) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("name"), req, resp)
}

func (c *ConnectorConfigModel) fromApiResponse(config *api.ConnectorConfig) {
	c.Name = types.StringValue(config.Name)
	c.Enabled = omittableBooleanValue(config.Enabled, c.Enabled)
	c.Description = omittableStringValue(config.Description, c.Description)
	c.EnableLogs = omittableBooleanValue(config.EnableLogs, c.EnableLogs)
	c.PingTime = types.StringValue(config.PingTime)
	c.SyncTime = types.StringValue(config.SyncTime)
}

func (c *ConnectorConfigModel) toAPICreateRequest() api.CreateConnectorConfigRequest {
	return api.CreateConnectorConfigRequest{
		Name:        c.Name.ValueString(),
		Enabled:     c.Enabled.ValueBoolPointer(),
		Description: c.Description.ValueStringPointer(),
		EnableLogs:  c.EnableLogs.ValueBoolPointer(),
		SyncTime:    c.SyncTime.ValueString(),
		PingTime:    c.PingTime.ValueString(),
	}
}

func (c *ConnectorConfigModel) toAPIUpdateRequest() api.UpdateConnectorConfigRequest {
	return api.UpdateConnectorConfigRequest{
		Enabled:     c.Enabled.ValueBoolPointer(),
		Description: c.Description.ValueStringPointer(),
		EnableLogs:  c.EnableLogs.ValueBoolPointer(),
		SyncTime:    c.SyncTime.ValueString(),
		PingTime:    c.PingTime.ValueString(),
	}
}
