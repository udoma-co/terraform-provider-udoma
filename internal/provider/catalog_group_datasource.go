package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"
)

var _ datasource.DataSourceWithConfigure = &catalogGroupDataSource{}

func NewCatalogGroupDataSource() datasource.DataSource {
	return &catalogGroupDataSource{}
}

type catalogGroupDataSource struct {
	client *client.UdomaClient
}

func (d *catalogGroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_catalog_group"
}

func (d *catalogGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasourceschema.Schema{
		MarkdownDescription: "Data source represents a catalog group",
		Attributes: map[string]datasourceschema.Attribute{
			"id": datasourceschema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the catalog group.",
			},
			"code": datasourceschema.StringAttribute{
				Required:    true,
				Description: "User provided unique identifier for the catalog group.",
			},
			"description": datasourceschema.MapAttribute{
				Computed:    true,
				Description: "Optional localized description of the catalog group.",
				ElementType: types.StringType,
			},
			"required_scopes": datasourceschema.ListAttribute{
				Computed:    true,
				Description: "Account scopes required for this group to be available for subscription.",
				ElementType: types.StringType,
			},
			"published": datasourceschema.BoolAttribute{
				Computed:    true,
				Description: "Whether the group is published and available for use.",
			},
		},
	}
}

func (d *catalogGroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	cl, ok := req.ProviderData.(*client.UdomaClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Data Type",
			fmt.Sprintf("Expected *client.UdomaClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.client = cl
}

func (d *catalogGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config catalogGroupModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	group, httpResp, err := d.client.GetApi().GetCatalogGroup(ctx, config.Code.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Catalog Group",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	if err := config.fromAPI(group); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Catalog Group",
			"Could not process API response, unexpected error: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
