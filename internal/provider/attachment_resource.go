package provider

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	v1 "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"
)

var (
	_ resource.ResourceWithConfigure   = &Attachment{}
	_ resource.ResourceWithImportState = &Attachment{}
)

func NewAttachment() resource.Resource {
	return &Attachment{}
}

// AppointmentSchedule defines the resource implementation
type Attachment struct {
	client *client.UdomaClient
}

// AppointmentScheduleModel describes the resource data model
type AttachmentModel struct {
	ID         types.String `tfsdk:"id"`
	Source     types.String `tfsdk:"source"`
	Created    types.Int64  `tfsdk:"created"`
	Ref        types.String `tfsdk:"ref"`
	FileType   types.String `tfsdk:"file_type"`
	FileSize   types.Int64  `tfsdk:"file_size"`
	FileName   types.String `tfsdk:"file_name"`
	FileSHA256 types.String `tfsdk:"file_sha256"`
	URL        types.String `tfsdk:"url"`
}

func (r *Attachment) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_attachment"
}

func (r *Attachment) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents an attachment",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the attachment",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"source": schema.StringAttribute{
				Required:    true,
				Description: "The file location to read the JSON from",
			},
			"created": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the attachment was created",
			},
			"ref": schema.StringAttribute{
				Computed:    true,
				Description: "Some reference to anything to whatever the attachment is linked to",
			},
			"file_type": schema.StringAttribute{
				Computed:    true,
				Description: "Metadata about the attachment",
			},
			"file_size": schema.Int64Attribute{
				Computed:    true,
				Description: "Size of the attachment",
			},
			"file_name": schema.StringAttribute{
				Computed:    true,
				Description: "The name of the attachment file",
			},
			"file_sha256": schema.StringAttribute{
				Computed:    true,
				Description: "The SHA256 sum of the file",
				PlanModifiers: []planmodifier.String{
					SHA256Modifier{fileSource: "source"},
				},
			},
			"url": schema.StringAttribute{
				Computed:    true,
				Description: "The URL to download the attachment",
			},
		},
	}
}

func (r *Attachment) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = cl
}

func (r *Attachment) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan AttachmentModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	f, err := os.Open(plan.Source.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Failed to open file", err.Error())
		return
	}

	attachment, _, err := r.client.GetApi().UploadAttachment(ctx).Data(f).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Failed to upload attachment", getApiErrorMessage(err))
		return
	}

	diags = plan.fromApiResponse(attachment)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *Attachment) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state AttachmentModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	attachment, httpResp, err := r.client.GetApi().GetAttachment(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Attachment",
			getApiErrorMessage(err),
		)
		return
	}

	diags = state.fromApiResponse(attachment)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *Attachment) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *Attachment) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state AttachmentModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.client.GetApi().DeleteAttachment(ctx, state.ID.ValueString()).Execute()
	if err != nil && (httpResp == nil || httpResp.StatusCode != 404) {
		resp.Diagnostics.AddError(
			"Error Deleting Attachment",
			getApiErrorMessage(err),
		)
	}
}

func (r *Attachment) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (m *AttachmentModel) fromApiResponse(api *v1.Attachment) (diags diag.Diagnostics) {
	if api == nil {
		diags.AddError(
			"API response is nil",
			"The argument is used to map values to the Attachment state and cannot be nil",
		)
		return
	}

	m.ID = types.StringPointerValue(api.Id)
	m.Created = types.Int64PointerValue(api.Created)
	m.Ref = types.StringPointerValue(api.Ref)
	m.FileType = types.StringPointerValue(api.FileType)
	m.FileSize = types.Int64PointerValue(api.FileSize)
	m.FileName = types.StringPointerValue(api.FileName)
	m.FileSHA256 = types.StringPointerValue(api.FileSha256)
	m.URL = types.StringPointerValue(api.Url)

	return
}

func getFileSHA256(file *os.File) (string, error) {
	h := sha256.New()

	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(h.Sum(nil)), nil
}
