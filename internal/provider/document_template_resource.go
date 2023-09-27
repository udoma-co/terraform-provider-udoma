package provider

import (
	"context"
	"encoding/json"
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
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/tf"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &documentTemplate{}
	_ resource.ResourceWithConfigure   = &documentTemplate{}
	_ resource.ResourceWithImportState = &documentTemplate{}
)

func NewDocumentTemplate() resource.Resource {
	return &documentTemplate{}
}

// documentTemplate defines the resource implementation.
type documentTemplate struct {
	client *client.UdomaClient
}

type documentTemplateOptionsModel struct {
	AllowTextEdit         types.Bool `tfsdk:"allow_text_edit"`
	IncludeFooterBranding types.Bool `tfsdk:"include_footer_branding"`
	IncludePageNumbers    types.Bool `tfsdk:"include_page_numbers"`
}

// documentTemplateModel describes the resource data model.
type documentTemplateModel struct {
	ID             types.String                  `tfsdk:"id"`
	LastUpdated    types.String                  `tfsdk:"last_updated"`
	Name           types.String                  `tfsdk:"name"`
	Description    types.String                  `tfsdk:"description"`
	Options        *documentTemplateOptionsModel `tfsdk:"options"`
	NameExpression types.String                  `tfsdk:"name_expression"`
	Content        tf.JsonObjectValue            `tfsdk:"content"`
	Inputs         tf.JsonObjectValue            `tfsdk:"inputs"`
	Placeholders   tf.JsonObjectValue            `tfsdk:"placeholders"`
	Signatures     tf.JsonObjectValue            `tfsdk:"signatures"`
}

func (r *documentTemplate) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_document_template"
}

func (r *documentTemplate) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{

		MarkdownDescription: "Resource represents a template for generating documents",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the document template",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the document template, shown in the admin page",
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "The description of the document template",
			},
			"options": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "The options of the document template",
				Attributes: map[string]schema.Attribute{
					"allow_text_edit": schema.BoolAttribute{
						Optional:    true,
						Description: "Controls if the generated document may be edited by the user",
					},
					"include_footer_branding": schema.BoolAttribute{
						Optional:    true,
						Description: "Controls if the footer branding should be included in the document PDF",
					},
					"include_page_numbers": schema.BoolAttribute{
						Optional:    true,
						Description: "Controls if the page numbers should be included in the document PDF",
					},
				},
			},
			"name_expression": schema.StringAttribute{
				Required: true,
				Description: `An optional JS expression to be used to compute the name of 
				the template. If not set, the name of the template will be used for new documents`,
			},
			"content": schema.StringAttribute{
				CustomType:  tf.JsonObjectType{},
				Required:    true,
				Description: "The content of the template",
			},
			"inputs": schema.StringAttribute{
				CustomType:  tf.JsonObjectType{},
				Required:    true,
				Description: "The JSON serialised form definition of the template",
			},
			"placeholders": schema.StringAttribute{
				CustomType:  tf.JsonObjectType{},
				Required:    true,
				Description: "The JSON serialised placeholder definition of the template",
			},
			"signatures": schema.StringAttribute{
				CustomType:  tf.JsonObjectType{},
				Optional:    true,
				Description: "The JSON serialised signature configuration of the template",
			},
		},
	}
}

func (r *documentTemplate) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *documentTemplate) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Retrieve values from plan
	var plan documentTemplateModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := documentTemplateFromModel(&plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Document Template",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newEndpoint, _, err := r.client.GetApi().CreateDocumentTemplate(ctx).CreateOrUpdateDocumentTemplateRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Document Template",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := documentTemplateToModel(&plan, newEndpoint); err != nil {
		resp.Diagnostics.AddError(
			"Error Creating Document Template",
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

func (r *documentTemplate) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	// Retrieve values from state
	var state documentTemplateModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	template, _, err := r.client.GetApi().GetDocumentTemplate(ctx, state.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Document Template",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := documentTemplateToModel(&state, template); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Document Template",
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

func (r *documentTemplate) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan documentTemplateModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := documentTemplateFromModel(&plan)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Document Template",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	id := plan.ID.ValueString()

	newEndpoint, _, err := r.client.GetApi().UpdateDocumentTemplate(ctx, id).CreateOrUpdateDocumentTemplateRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Document Template",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := documentTemplateToModel(&plan, newEndpoint); err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Document Template",
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

func (r *documentTemplate) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	// Retrieve values from state
	var state documentTemplateModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := r.client.GetApi().DeleteDocumentTemplate(ctx, state.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Document Template",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}
}

func (r *documentTemplate) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func documentTemplateToModel(model *documentTemplateModel, template *api.DocumentTemplate) error {

	if template == nil {
		return fmt.Errorf("document template is nil")
	}

	model.ID = types.StringValue(sdp(template.Id))
	model.Name = types.StringValue(sdp(template.Name))
	model.Description = types.StringValue(sdp(template.Description))
	if template.Options != nil {
		model.Options = &documentTemplateOptionsModel{
			AllowTextEdit:         types.BoolValue(bdp(template.Options.AllowTextEdit)),
			IncludeFooterBranding: types.BoolValue(bdp(template.Options.IncludeFooterBranding)),
			IncludePageNumbers:    types.BoolValue(bdp(template.Options.IncludePageNumbers)),
		}
	}
	model.NameExpression = types.StringValue(sdp(template.NameExpression))
	model.Content = tf.NewJsonObjectValue(sdp(template.Content))

	if template.Inputs != nil {
		jsonData, err := json.Marshal(template.Inputs)
		if err != nil {
			return fmt.Errorf("failed to marshal inputs: %w", err)
		}
		model.Inputs = tf.NewJsonObjectValue(string(jsonData))
	}

	if template.Placeholders != nil {
		jsonData, err := json.Marshal(template.Placeholders)
		if err != nil {
			return fmt.Errorf("failed to marshal placeholders: %w", err)
		}
		model.Placeholders = tf.NewJsonObjectValue(string(jsonData))
	}

	if template.Signatures != nil {
		jsonData, err := json.Marshal(template.Signatures)
		if err != nil {
			return fmt.Errorf("failed to marshal signatures: %w", err)
		}
		if string(jsonData) != "{}" {
			// Signatures are optional, but we need to set an empty object to avoid
			// inconsistencies in the state because of server side serialization.
			model.Signatures = tf.NewJsonObjectValue(string(jsonData))
		}
	}

	return nil
}

func documentTemplateFromModel(model *documentTemplateModel) (api.CreateOrUpdateDocumentTemplateRequest, error) {

	template := api.CreateOrUpdateDocumentTemplateRequest{
		Name:           model.Name.ValueStringPointer(),
		NameExpression: model.NameExpression.ValueStringPointer(),
		Description:    model.Description.ValueStringPointer(),
		Content:        model.Content.ValueStringPointer(),
	}
	if model.Options != nil {
		template.Options = &api.DocumentTemplateOptions{
			AllowTextEdit:         model.Options.AllowTextEdit.ValueBoolPointer(),
			IncludeFooterBranding: model.Options.IncludeFooterBranding.ValueBoolPointer(),
			IncludePageNumbers:    model.Options.IncludePageNumbers.ValueBoolPointer(),
		}
	}

	if !model.Inputs.IsNull() && !model.Inputs.IsUnknown() {
		if diag := model.Inputs.Unmarshal(&template.Inputs); diag.HasError() {
			return template, fmt.Errorf("failed to unmarshal inputs: %v", diag)
		}
	}

	if !model.Placeholders.IsNull() && !model.Placeholders.IsUnknown() {
		if diag := model.Placeholders.Unmarshal(&template.Placeholders); diag.HasError() {
			return template, fmt.Errorf("failed to unmarshal placeholders: %v", diag)
		}
	}

	if !model.Signatures.IsNull() && !model.Signatures.IsUnknown() {
		if diag := model.Signatures.Unmarshal(&template.Signatures); diag.HasError() {
			return template, fmt.Errorf("failed to unmarshal signatures: %v", diag)
		}
	}

	return template, nil
}
