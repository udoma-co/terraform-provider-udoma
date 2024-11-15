package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
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

type CaseReportingEndpointTemplateModel struct {
	Priority types.Int64  `tfsdk:"priority"`
	ID       types.String `tfsdk:"id"`
}

type CaseReportingEndpointCategoryModel struct {
	Name      types.Map                            `tfsdk:"name"`
	Priority  types.Int64                          `tfsdk:"priority"`
	Templates []CaseReportingEndpointTemplateModel `tfsdk:"templates"`
}

// CaseReportingEndpointModel describes the resource data model.
type CaseReportingEndpointModel struct {
	ID             types.String                         `tfsdk:"id"`
	Code           types.String                         `tfsdk:"code"`
	CreatedAt      types.Int64                          `tfsdk:"created_at"`
	UpdatedAt      types.Int64                          `tfsdk:"updated_at"`
	Name           types.String                         `tfsdk:"name"`
	Active         types.Bool                           `tfsdk:"active"`
	CaseCategories []CaseReportingEndpointCategoryModel `tfsdk:"case_categories"`
	LastUpdated    types.String                         `tfsdk:"last_updated"`
	FAQs           []types.String                       `tfsdk:"faqs"`
	Url            types.String                         `tfsdk:"url"`
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
			"case_categories": schema.ListNestedAttribute{
				Required:            true,
				MarkdownDescription: "Reference to the case templates that can be raised via this endpoint",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.MapAttribute{
							Required:            true,
							MarkdownDescription: "The name of the category in different languages",
							ElementType:         types.StringType,
						},
						"priority": schema.Int64Attribute{
							Required:            true,
							MarkdownDescription: "The priority of the category - used to order the categories",
						},
						"templates": schema.ListNestedAttribute{
							Required:            true,
							MarkdownDescription: "Reference to the templates that would show up in this category",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"priority": schema.Int64Attribute{
										Required:            true,
										MarkdownDescription: "The priority of the template - used to order the templates in the category",
									},
									"id": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "A reference to the appointment template",
									},
								},
							},
						},
					},
				},
			},
			"faqs": schema.ListAttribute{
				Optional:            true,
				MarkdownDescription: "The faq by id that are linked to this endpoint",
				ElementType:         types.StringType,
			},
			"url": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The URL to access the endpoint",
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
	diags = plan.fromAPI(newEndpoint)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
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
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Case Reporting Endpoint",
			"Could not read entity from Udooma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	diags = state.fromAPI(endpoint)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
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
	diags = plan.fromAPI(newEndpoint)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
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

	httpResp, err := r.client.GetApi().DeleteCaseReportingEndpoint(ctx, state.Code.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Case Reporting Endpoint",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (r *CaseReportingEndpoint) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("code"), req, resp)
}

func (model *CaseReportingEndpointModel) fromAPI(endpoint *api.CaseReportingEndpoint) (diags diag.Diagnostics) {

	if endpoint == nil {
		diags.AddError(
			"endpoint is nil",
			"the endpoint is used for mapping values to the CaseReportingEndpoint state and cannot be nil",
		)
		return
	}

	model.ID = types.StringValue(endpoint.Code)
	model.Code = types.StringValue(endpoint.Code)
	model.CreatedAt = types.Int64Value(endpoint.CreatedAt)
	model.UpdatedAt = types.Int64Value(endpoint.UpdatedAt)
	model.Name = types.StringValue(endpoint.Name)
	model.Active = omittableBooleanValue(endpoint.Active, model.Active)
	model.Url = types.StringValue(endpoint.Url)
	model.CaseCategories = make([]CaseReportingEndpointCategoryModel, len(endpoint.CaseCategories))

	for i := range endpoint.CaseCategories {
		model.CaseCategories[i].Name = basetypes.NewMapNull(types.StringType)

		diags = model.CaseCategories[i].fromAPIResponse(&endpoint.CaseCategories[i])
		if diags.HasError() {
			return
		}
	}

	if len(endpoint.Faqs) > 0 {
		model.FAQs = make([]types.String, len(endpoint.Faqs))
		for i := range endpoint.Faqs {
			model.FAQs[i] = types.StringValue(endpoint.Faqs[i])
		}
	} else {
		if model.FAQs != nil {
			model.FAQs = make([]types.String, 0)
		}
	}

	return
}

func (model *CaseReportingEndpointCategoryModel) fromAPIResponse(category *api.CaseReportingEndpointCategory) (diags diag.Diagnostics) {
	if category == nil {
		diags.AddError(
			"category is nil",
			"the category is used for mapping values to the CaseReportingEndpoint state and cannot be nil",
		)
		return
	}

	model.Name, diags = types.MapValue(types.StringType, stringMapToValueMap(category.Name))
	if diags.HasError() {
		return
	}

	model.Priority = types.Int64Value(int64(idp32(category.Priority)))

	model.Templates = make([]CaseReportingEndpointTemplateModel, len(category.Templates))
	for i := range category.Templates {
		diags = model.Templates[i].fromAPIResponse(&category.Templates[i])
		if diags.HasError() {
			return
		}
	}

	return
}

func (model *CaseReportingEndpointTemplateModel) fromAPIResponse(template *api.CaseReportingEndpointCategoryTemplatesInner) (diags diag.Diagnostics) {
	if template == nil {
		diags.AddError(
			"template is nil",
			"the template is used for mapping values to the CaseReportingEndpoint state and cannot be nil",
		)
		return
	}

	model.ID = types.StringValue(template.Id)
	model.Priority = types.Int64Value(int64(idp32(template.Priority)))

	return
}

func (model *CaseReportingEndpointCategoryModel) toAPIRequest() api.CaseReportingEndpointCategory {

	category := api.CaseReportingEndpointCategory{
		Name:      modelMapToStringMap(model.Name),
		Priority:  i64ToI32Ptr(model.Priority.ValueInt64()),
		Templates: make([]api.CaseReportingEndpointCategoryTemplatesInner, len(model.Templates)),
	}

	for i := range model.Templates {
		category.Templates[i] = model.Templates[i].toAPIRequest()
	}

	return category
}

func (model *CaseReportingEndpointTemplateModel) toAPIRequest() api.CaseReportingEndpointCategoryTemplatesInner {
	return api.CaseReportingEndpointCategoryTemplatesInner{
		Id:       model.ID.ValueString(),
		Priority: i64ToI32Ptr(model.Priority.ValueInt64()),
	}
}

func (model *CaseReportingEndpointModel) toAPIRequest() (api.CreateOrUpdateCaseReportingEndpointRequest, error) {

	endpoint := api.CreateOrUpdateCaseReportingEndpointRequest{
		Name:           model.Name.ValueString(),
		Active:         model.Active.ValueBoolPointer(),
		CaseCategories: make([]api.CaseReportingEndpointCategory, len(model.CaseCategories)),
		Faqs:           make([]string, len(model.FAQs)),
	}

	for i := range model.CaseCategories {
		endpoint.CaseCategories[i] = model.CaseCategories[i].toAPIRequest()
	}

	for i := range model.FAQs {
		endpoint.Faqs[i] = model.FAQs[i].ValueString()
	}

	return endpoint, nil
}
