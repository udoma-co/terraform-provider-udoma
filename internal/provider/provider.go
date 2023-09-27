package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"
)

// Ensure UdomaProvider satisfies various provider interfaces.
var _ provider.Provider = &UdomaProvider{}

// UdomaProvider defines the provider implementation.
type UdomaProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// udomaProviderModel describes the provider data model.
type udomaProviderModel struct {
	ApiKey    types.String `tfsdk:"api_key"`
	KeySecret types.String `tfsdk:"key_secret"`
	Endpoint  types.String `tfsdk:"endpoint"`
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &UdomaProvider{
			version: version,
		}
	}
}

func (p *UdomaProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "udoma"
	resp.Version = p.version
}

func (p *UdomaProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "The URL for the Udoma instance",
				Optional:            true,
			},
			"api_key": schema.StringAttribute{
				MarkdownDescription: "The API key part of authentication for accessing the Udoma APIs",
				Optional:            true,
			},
			"key_secret": schema.StringAttribute{
				MarkdownDescription: "The key secret part of the authentication for accessing the Udoma APIs",
				Optional:            true,
			},
		},
	}
}

func (p *UdomaProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

	tflog.Info(ctx, "Configuring the Udoma provider")

	// retrieve config from the configuration
	var config udomaProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// validate the configuration

	if config.Endpoint.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Unknown Udoma API Endpoint",
			"The provider cannot create the Udoma API client as there is an unknown configuration value for the Udoma API endpoint. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the UDOMA_ENDPOINT environment variable.",
		)
	}

	if config.ApiKey.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("api_key"),
			"Unknown Udoma API API Key",
			"The provider cannot create the Udoma API client as there is an unknown configuration value for the Udoma API API Key. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the UDOMA_API_KEY environment variable.",
		)
	}

	if config.KeySecret.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("key_secret"),
			"Unknown Udoma API Key Secret",
			"The provider cannot create the Udoma API client as there is an unknown configuration value for the Udoma API Key Secret. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the UDOMA_KEY_SECRET environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	endpoint := os.Getenv("UDOMA_ENDPOINT")
	apiKey := os.Getenv("UDOMA_API_KEY")
	keySecret := os.Getenv("UDOMA_KEY_SECRET")

	if !config.Endpoint.IsNull() {
		endpoint = config.Endpoint.ValueString()
	}

	if !config.ApiKey.IsNull() {
		apiKey = config.ApiKey.ValueString()
	}

	if !config.KeySecret.IsNull() {
		keySecret = config.KeySecret.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if endpoint == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Udoma API Endpoint",
			"The provider cannot create the Udoma API client as there is a missing or empty value for the Udoma API endpoint. "+
				"Set the endpoint value in the configuration or use the UDOMA_ENDPOINT environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if apiKey == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing Udoma API Key",
			"The provider cannot create the Udoma API client as there is a missing or empty value for the Udoma API key. "+
				"Set the api key value in the configuration or use the UDOMA_API_KEY environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if keySecret == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Udoma API Key Secret",
			"The provider cannot create the Udoma API client as there is a missing or empty value for the Udoma API key secret. "+
				"Set the key secret value in the configuration or use the UDOMA_KEY_SECRET environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	conf := client.Config{
		Endpoint:     endpoint,
		ApiKey:       apiKey,
		ApiSecretKey: keySecret,
	}

	// Create a new Udoma client using the configuration values
	client, err := client.NewClient(conf)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Udoma API Client",
			"An unexpected error occurred when creating the Udoma API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Udoma Client Error: "+err.Error(),
		)
		return
	}

	ctx = tflog.SetField(ctx, "udoma_endpoint", endpoint)

	tflog.Info(ctx, "Created Udoma API client")

	// Make the HashiCups client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *UdomaProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewCaseReportingEndpoint,
		NewCaseTemplate,
		NewCustomIDGenerator,
		NewCustomerScript,
		NewDocumentTemplate,
		NewWorkflowDefinition,
		NewWorkflowEntrypoint,
	}
}

func (p *UdomaProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewJsonSource,
	}
}
