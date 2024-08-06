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
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"

	api "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
)

var (
	_ resource.ResourceWithConfigure   = &ConnectorQuery{}
	_ resource.ResourceWithImportState = &ConnectorQuery{}
)

func NewConnectorQuery() resource.Resource {
	return &ConnectorQuery{}
}

type ConnectorQuery struct {
	client *client.UdomaClient
}

type ConnectorQueryModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Connector   types.String `tfsdk:"connector"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	Priority    types.Int64  `tfsdk:"priority"`
	Limit       types.Int64  `tfsdk:"limit"`
	Statement   types.String `tfsdk:"statement"`
	Script      types.String `tfsdk:"script"`
}

func (c *ConnectorQuery) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_connector_query"
}

func (c *ConnectorQuery) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents a connector query",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the connector query",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the connector query",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "The display name of the connector query",
			},
			"enabled": schema.BoolAttribute{
				Required:    true,
				Description: "Whether the query is enabled or not",
			},
			"connector": schema.StringAttribute{
				Required:    true,
				Description: "The connector ID to link this query to",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Description: "The priority of the query, higher values are executed first",
			},
			"limit": schema.Int64Attribute{
				Required:    true,
				Description: "The maximum number of results to return",
			},
			"statement": schema.StringAttribute{
				Required:    true,
				Description: "The statement to query the data with",
			},
			"script": schema.StringAttribute{
				Required:    true,
				Description: "The script to process the data with",
			},
		},
	}
}

func (q *ConnectorQuery) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (c *ConnectorQuery) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ConnectorQueryModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiReq := plan.toAPIRequest()
	query, _, err := c.client.GetApi().CreateConnectorQuery(ctx).CreateOrUpdateConnectorQueryRequest(apiReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Connector Query",
			fmt.Sprintf("Could not create entity in Udoma, unexpected error: %s", getApiErrorMessage(err)),
		)
		return
	}

	plan.fromApiResponse(query)
	resp.Diagnostics.Append(
		resp.State.Set(ctx, plan)...,
	)
}

func (c *ConnectorQuery) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ConnectorQueryModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newQuery, httpResp, err := c.client.GetApi().GetConnectorQuery(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Connector Query",
			fmt.Sprintf("Could not read entity from Udoma, unexpected error: %s", getApiErrorMessage(err)),
		)
		return
	}

	state.fromApiResponse(newQuery)
	resp.Diagnostics.Append(
		resp.State.Set(ctx, state)...,
	)
}

func (c *ConnectorQuery) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ConnectorQueryModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiReq := plan.toAPIRequest()
	query, _, err := c.client.GetApi().UpdateConnectorQuery(ctx, plan.ID.ValueString()).CreateOrUpdateConnectorQueryRequest(apiReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Connector Query",
			fmt.Sprintf("Could not update entity in Udoma, unexpected error: %s", getApiErrorMessage(err)),
		)
		return
	}

	plan.fromApiResponse(query)
	resp.Diagnostics.Append(
		resp.State.Set(ctx, plan)...,
	)
}

func (c *ConnectorQuery) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ConnectorQueryModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := c.client.GetApi().DeleteConnectorQuery(ctx, state.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Connector Query",
			fmt.Sprintf("Could not delete entity from Udoma, unexpected error: %s", getApiErrorMessage(err)),
		)
		return
	}
}

func (r *ConnectorQuery) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (c *ConnectorQueryModel) fromApiResponse(query *api.ConnectorQuery) {
	c.ID = types.StringValue(query.Id)
	c.Name = types.StringValue(query.Name)
	c.Description = omittableStringValue(query.Description, c.Description)
	c.Connector = types.StringValue(query.ConnectorRef)
	c.Enabled = omittableBooleanValue(query.Enabled, c.Enabled)
	c.Priority = types.Int64Value(int64(query.Priority))
	c.Limit = types.Int64Value(int64(query.EntityLimit))
	c.Statement = types.StringValue(query.Statement)
	c.Script = types.StringValue(query.ProcessScript)
}

func (c *ConnectorQueryModel) toAPIRequest() api.CreateOrUpdateConnectorQueryRequest {
	return api.CreateOrUpdateConnectorQueryRequest{
		Name:          c.Name.ValueString(),
		Description:   c.Description.ValueStringPointer(),
		ConnectorRef:  c.Connector.ValueString(),
		Enabled:       c.Enabled.ValueBoolPointer(),
		Priority:      int32(c.Priority.ValueInt64()),
		EntityLimit:   int32(c.Limit.ValueInt64()),
		Statement:     c.Statement.ValueString(),
		ProcessScript: c.Script.ValueString(),
	}
}
