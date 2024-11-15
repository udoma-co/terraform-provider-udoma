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
	v1 "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"
)

var (
	_ resource.ResourceWithConfigure   = &Document{}
	_ resource.ResourceWithImportState = &Document{}
)

func NewDocument() resource.Resource {
	return &Document{}
}

type Document struct {
	client *client.UdomaClient
}

type DocumentModel struct {
	ID         types.String `tfsdk:"id"`
	CreatedAt  types.Int64  `tfsdk:"created_at"`
	UpdatedAt  types.Int64  `tfsdk:"updated_at"`
	Name       types.String `tfsdk:"name"`
	Path       types.String `tfsdk:"path"`
	Public     types.Bool   `tfsdk:"public"`
	Attachment types.String `tfsdk:"attachment"`
	RefType    types.String `tfsdk:"ref_type"`
	RefID      types.String `tfsdk:"ref_id"`
}

func (d *Document) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_document"
}

func (d *Document) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents a document",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the document",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The time the document was created",
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The time the document was last updated",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the document",
			},
			"attachment": schema.StringAttribute{
				Required:    true,
				Description: "The attachment to be used in the document",
			},
			"path": schema.StringAttribute{
				Optional:    true,
				Description: "The path of the document in the storage",
			},
			"public": schema.BoolAttribute{
				Optional:    true,
				Description: "The public status of the document",
			},
			"ref_type": schema.StringAttribute{
				Required:    true,
				Description: "The type of the reference",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("static", "proprty"),
				},
			},
			"ref_id": schema.StringAttribute{
				Optional:    true,
				Description: "The identifier of the reference",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (d *Document) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	d.client = cl
}

func (d *Document) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var model DocumentModel
	diags := req.Plan.Get(ctx, &model)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	document, _, err := d.client.GetApi().CreateDocument(ctx).CreateDocumentRequest(*model.toCreateAPIRequest()).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Failed to create document", getApiErrorMessage(err))
		return
	}

	model.fromAPIResponse(document)
	diags = resp.State.Set(ctx, &model)
	resp.Diagnostics.Append(diags...)
}

func (d *Document) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var model DocumentModel
	diags := req.State.Get(ctx, &model)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	document, httpResp, err := d.client.GetApi().GetDocument(ctx, model.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Document",
			"Could not read entity from Udooma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	model.fromAPIResponse(document)
	diags = resp.State.Set(ctx, &model)
	resp.Diagnostics.Append(diags...)
}

func (d *Document) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	var model DocumentModel
	diags := req.Plan.Get(ctx, &model)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	document, httpResp, err := d.client.GetApi().UpdateDocument(ctx, model.ID.ValueString()).UpdateDocumentRequest(*model.toUpdateAPIRequest()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Failed to update document", getApiErrorMessage(err))
		return
	}

	model.fromAPIResponse(document)
	resp.State.Set(ctx, &model)
	resp.Diagnostics.Append(diags...)
}

func (d *Document) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var model DocumentModel
	diags := req.State.Get(ctx, &model)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := d.client.GetApi().DeleteDocument(ctx, model.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		// if resource is not found, we consider it already deleted
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting document",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (r *Document) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *DocumentModel) toCreateAPIRequest() *v1.CreateDocumentRequest {

	req := &v1.CreateDocumentRequest{
		Name:          m.Name.ValueString(),
		Path:          m.Path.ValueStringPointer(),
		Public:        m.Public.ValueBoolPointer(),
		RefType:       v1.DocumentRefTypeEnum(m.RefType.ValueString()),
		RefId:         m.RefID.ValueStringPointer(),
		AttachmentRef: m.Attachment.ValueString(),
	}

	return req
}

func (m *DocumentModel) toUpdateAPIRequest() *v1.UpdateDocumentRequest {

	req := &v1.UpdateDocumentRequest{
		Name:   m.Name.ValueString(),
		Path:   m.Path.ValueStringPointer(),
		Public: m.Public.ValueBoolPointer(),
	}

	return req
}

func (m *DocumentModel) fromAPIResponse(document *v1.Document) {
	m.ID = types.StringValue(document.Id)
	m.CreatedAt = types.Int64Value(document.CreatedAt)
	m.UpdatedAt = types.Int64Value(document.UpdatedAt)
	m.Name = types.StringValue(document.Name)
	m.Path = omittableStringValue(document.Path, m.Path)
	m.Public = omittableBooleanValue(document.Public, m.Public)
	m.Attachment = types.StringValue(document.Attachment.Id)
}
