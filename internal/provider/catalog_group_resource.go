package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	api "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"
)

var (
	_ resource.Resource                = &CatalogGroup{}
	_ resource.ResourceWithConfigure   = &CatalogGroup{}
	_ resource.ResourceWithImportState = &CatalogGroup{}
)

func NewCatalogGroup() resource.Resource {
	return &CatalogGroup{}
}

type CatalogGroup struct {
	client *client.UdomaClient
}

type catalogGroupModel struct {
	ID             types.String `tfsdk:"id"`
	Code           types.String `tfsdk:"code"`
	Description    types.Map    `tfsdk:"description"`
	RequiredScopes types.List   `tfsdk:"required_scopes"`
	Published      types.Bool   `tfsdk:"published"`
}

func (r *CatalogGroup) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_catalog_group"
}

func (r *CatalogGroup) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents a catalog group",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the catalog group.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"code": schema.StringAttribute{
				Required:    true,
				Description: "User provided unique identifier for the catalog group.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
			},
			"description": schema.MapAttribute{
				Optional:    true,
				Description: "Optional localized description of the catalog group.",
				ElementType: types.StringType,
			},
			"required_scopes": schema.ListAttribute{
				Optional:    true,
				Description: "Account scopes required for this group to be available for subscription.",
				ElementType: types.StringType,
				Validators: []validator.List{
					listvalidator.UniqueValues(),
					listvalidator.ValueStringsAre(stringvalidator.OneOf(stringSlice(api.AllowedAccountScopeCodeEnumValues)...)),
				},
			},
			"published": schema.BoolAttribute{Optional: true, Description: "Whether the group is published and available for use."},
		},
	}
}

func (r *CatalogGroup) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CatalogGroup) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan catalogGroupModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError("Error Creating Catalog Group", "Could not create API request, unexpected error: "+err.Error())
		return
	}

	group, _, err := r.client.GetApi().CreateCatalogGroup(ctx).CreateOrUpdateCatalogGroup(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Creating Catalog Group", "Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err))
		return
	}

	if err := plan.fromAPI(group); err != nil {
		resp.Diagnostics.AddError("Error Creating Catalog Group", "Could not process API response, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *CatalogGroup) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state catalogGroupModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	group, httpResp, err := r.client.GetApi().GetCatalogGroup(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error Reading Catalog Group", "Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err))
		return
	}

	if err := state.fromAPI(group); err != nil {
		resp.Diagnostics.AddError("Error Reading Catalog Group", "Could not process API response, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *CatalogGroup) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan catalogGroupModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError("Error Updating Catalog Group", "Could not create API request, unexpected error: "+err.Error())
		return
	}

	group, _, err := r.client.GetApi().UpdateCatalogGroup(ctx, plan.ID.ValueString()).CreateOrUpdateCatalogGroup(updateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Updating Catalog Group", "Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err))
		return
	}

	if err := plan.fromAPI(group); err != nil {
		resp.Diagnostics.AddError("Error Updating Catalog Group", "Could not process API response, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *CatalogGroup) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state catalogGroupModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.client.GetApi().DeleteCatalogGroup(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error Deleting Catalog Group", "Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err))
	}
}

func (r *CatalogGroup) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *catalogGroupModel) fromAPI(group *api.CatalogGroup) error {
	if group == nil {
		return fmt.Errorf("catalog group is nil")
	}

	model.ID = types.StringValue(group.Code)
	model.Code = types.StringValue(group.Code)

	if group.Description != nil {
		mapped, diags := types.MapValue(types.StringType, stringMapToValueMap(*group.Description))
		if diags.HasError() {
			return fmt.Errorf("could not map description")
		}
		model.Description = mapped
	} else {
		model.Description = types.MapNull(types.StringType)
	}

	model.RequiredScopes = omittableEnumListValue(group.RequiredScopes, model.RequiredScopes)

	model.Published = omittableBooleanValue(group.Published, model.Published)
	return nil
}

func (model *catalogGroupModel) toAPIRequest() (api.CreateOrUpdateCatalogGroup, error) {
	group := api.CreateOrUpdateCatalogGroup{Code: model.Code.ValueString()}

	if description := modelMapToStringMap(model.Description); len(description) > 0 {
		group.Description = &description
	}

	if requiredScopes := modelListToEnumSlice[api.AccountScopeCode](model.RequiredScopes); len(requiredScopes) > 0 {
		group.RequiredScopes = requiredScopes
	}

	group.Published = model.Published.ValueBoolPointer()
	return group, nil
}
