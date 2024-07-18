package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	api "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &CommentTemplate{}
	_ resource.ResourceWithConfigure   = &CommentTemplate{}
	_ resource.ResourceWithImportState = &CommentTemplate{}
)

func NewCommentTemplate() resource.Resource {
	return &CommentTemplate{}
}

// CommentTemplate defines the resource implementation.
type CommentTemplate struct {
	client *client.UdomaClient
}

// CommentTemplateModel defines the model for the resource.
type CommentTemplateModel struct {
	ID          types.String `tfsdk:"id"`
	CreatedAt   types.Int64  `tfsdk:"created_at"`
	UpdatedAt   types.Int64  `tfsdk:"updated_at"`
	DisplayName types.String `tfsdk:"display_name"`
	AccessList  types.List   `tfsdk:"access_list"`
	IsDenyList  types.Bool   `tfsdk:"is_deny_list"`
	Script      types.String `tfsdk:"script"`
	Template    types.String `tfsdk:"template"`
}

func (c *CommentTemplate) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_comment_template"
}

func (c *CommentTemplate) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Comment template resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier of the comment template",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the template was created",
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The date and time the template was last updated",
			},
			"display_name": schema.StringAttribute{
				Required:    true,
				Description: "The display name of the comment template",
			},
			"access_list": schema.ListAttribute{
				Optional:    true,
				Description: "The list of users that have access to the comment template",
				ElementType: types.StringType,
			},
			"is_deny_list": schema.BoolAttribute{
				Optional:    true,
				Description: "If set true the access list will act as a deny list instead of an allow list",
			},
			"script": schema.StringAttribute{
				Optional:    true,
				Description: "The script that will be executed to generate the comment",
			},
			"template": schema.StringAttribute{
				Required:    true,
				Description: "The template that will be used to generate the comment",
			},
		},
	}
}

func (c *CommentTemplate) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	c.client = cl
}

func (c *CommentTemplate) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan CommentTemplateModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq := plan.toAPIRequest()
	ct, _, err := c.client.GetApi().CreateCommentTemplate(ctx).CreateOrUpdateCommentTemplateRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to create comment template",
			fmt.Sprintf("Could not create comment template entity in Udoma, unexpected error: %s", getApiErrorMessage(err)),
		)
	}

	plan.fromAPIResponse(ct)
	resp.Diagnostics.Append(
		resp.State.Set(ctx, plan)...,
	)
}

func (c *CommentTemplate) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CommentTemplateModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	newCommentTemplate, httpResp, err := c.client.GetApi().GetCommentTemplate(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Comment Template",
			fmt.Sprintf("Could not read entity in Udoma, unexpected error: %s", getApiErrorMessage(err)),
		)
		return
	}

	state.fromAPIResponse(newCommentTemplate)
	resp.Diagnostics.Append(
		resp.State.Set(ctx, state)...,
	)
}

func (c *CommentTemplate) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan CommentTemplateModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiReq := plan.toAPIRequest()
	commentTemplate, _, err := c.client.GetApi().UpdateCommentTemplate(ctx, plan.ID.ValueString()).CreateOrUpdateCommentTemplateRequest(apiReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Comment Template",
			fmt.Sprintf("Could not update entity in Udoma, unexpected error: %s", getApiErrorMessage(err)),
		)
		return
	}

	plan.fromAPIResponse(commentTemplate)
	resp.Diagnostics.Append(
		resp.State.Set(ctx, plan)...,
	)
}

func (c *CommentTemplate) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CommentTemplateModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := c.client.GetApi().DeleteCommentTemplate(ctx, state.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Comment Template",
			fmt.Sprintf("Could not delete entity in Udoma, unexpected error: %s", getApiErrorMessage(err)),
		)
		return
	}
}

func (r *CommentTemplate) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Retrieve import ID and save to id attribute
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model CommentTemplateModel) toAPIRequest() api.CreateOrUpdateCommentTemplateRequest {
	return api.CreateOrUpdateCommentTemplateRequest{
		DisplayName: model.DisplayName.ValueString(),
		AccessList:  modelListToStringSlice(model.AccessList),
		IsDenyList:  model.IsDenyList.ValueBoolPointer(),
		Script:      model.Script.ValueStringPointer(),
		Template:    model.Template.ValueString(),
	}
}

func (model *CommentTemplateModel) fromAPIResponse(ct *api.CommentTemplate) (diags diag.Diagnostics) {
	if ct == nil {
		diags.AddError("API response is nil", "The API response provided to CommentTemplateModel.fromAPIResponse is nil")
		return
	}

	model.ID = types.StringValue(ct.Id)
	model.CreatedAt = types.Int64Value(ct.CreatedAt)
	model.UpdatedAt = types.Int64Value(ct.UpdatedAt)
	model.DisplayName = types.StringValue(ct.DisplayName)

	model.AccessList, diags = types.ListValue(types.StringType, stringSliceToValueList(ct.AccessList))
	if diags.HasError() {
		return
	}

	model.IsDenyList = omittableBooleanValue(ct.IsDenyList, model.IsDenyList)
	model.Script = omittableStringValue(ct.Script, model.Script)
	model.Template = types.StringValue(ct.Template)

	return nil
}
