package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	datasourceschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"
)

var _ datasource.DataSourceWithConfigure = &catalogItemDataSource{}

func NewCatalogItemDataSource() datasource.DataSource {
	return &catalogItemDataSource{}
}

type catalogItemDataSource struct {
	client *client.UdomaClient
}

func (d *catalogItemDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_catalog_item"
}

func (d *catalogItemDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = datasourceschema.Schema{
		MarkdownDescription: "Data source represents a catalog item",
		Attributes: map[string]datasourceschema.Attribute{
			"id": datasourceschema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the catalog item.",
			},
			"code": datasourceschema.StringAttribute{
				Required:    true,
				Description: "User provided unique identifier for the catalog item.",
			},
			"groups": datasourceschema.ListAttribute{
				Computed:    true,
				Description: "Catalog group codes this item belongs to.",
				ElementType: types.StringType,
			},
			"ref_type": datasourceschema.StringAttribute{
				Computed:    true,
				Description: "Type of referenced entity that this catalog item points to.",
			},
			"ref_id": datasourceschema.StringAttribute{
				Computed:    true,
				Description: "ID of the referenced entity.",
			},
			"account_ref": datasourceschema.Int64Attribute{
				Computed:    true,
				Description: "Account number of the referenced financial account.",
			},
			"description": datasourceschema.MapAttribute{
				Computed:    true,
				Description: "Optional localized description of the catalog item.",
				ElementType: types.StringType,
			},
			"published": datasourceschema.BoolAttribute{
				Computed:    true,
				Description: "Whether the catalog item is published and available for use.",
			},
			"required_scopes": datasourceschema.ListAttribute{
				Computed:    true,
				Description: "Account scopes required for this item to be available for subscription.",
				ElementType: types.StringType,
			},
			"dependencies": datasourceschema.ListAttribute{
				Computed:    true,
				Description: "Codes of catalog items this item depends on.",
				ElementType: types.StringType,
			},
			"patches": datasourceschema.ListNestedAttribute{
				Computed:    true,
				Description: "Patch definitions that can override attributes of the referenced entity.",
				NestedObject: datasourceschema.NestedAttributeObject{
					Attributes: map[string]datasourceschema.Attribute{
						"code":   datasourceschema.StringAttribute{Computed: true, Description: "Unique code for the patch."},
						"format": datasourceschema.StringAttribute{Computed: true, Description: "Patch format (for example JSON patch or unified diff)."},
						"patch":  datasourceschema.StringAttribute{Computed: true, Description: "Patch payload."},
					},
				},
			},
		},
	}
}

func (d *catalogItemDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *catalogItemDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config catalogItemModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	item, httpResp, err := d.client.GetApi().GetCatalogItem(ctx, config.Code.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Catalog Item",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	if err := config.fromAPI(item); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Catalog Item",
			"Could not process API response, unexpected error: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &config)...)
}
