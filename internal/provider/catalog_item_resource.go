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
	_ resource.Resource                = &CatalogItem{}
	_ resource.ResourceWithConfigure   = &CatalogItem{}
	_ resource.ResourceWithImportState = &CatalogItem{}
)

func NewCatalogItem() resource.Resource {
	return &CatalogItem{}
}

type CatalogItem struct {
	client *client.UdomaClient
}

type catalogItemPatchDefinitionModel struct {
	Code   types.String `tfsdk:"code"`
	Format types.String `tfsdk:"format"`
	Patch  types.String `tfsdk:"patch"`
}

type catalogItemModel struct {
	ID             types.String                      `tfsdk:"id"`
	Code           types.String                      `tfsdk:"code"`
	Groups         types.List                        `tfsdk:"groups"`
	RefType        types.String                      `tfsdk:"ref_type"`
	RefId          types.String                      `tfsdk:"ref_id"`
	AccountRef     types.Int64                       `tfsdk:"account_ref"`
	Description    types.Map                         `tfsdk:"description"`
	Published      types.Bool                        `tfsdk:"published"`
	RequiredScopes types.List                        `tfsdk:"required_scopes"`
	Dependencies   types.List                        `tfsdk:"dependencies"`
	Patches        []catalogItemPatchDefinitionModel `tfsdk:"patches"`
}

func (r *CatalogItem) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_catalog_item"
}

func (r *CatalogItem) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents a catalog item",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the catalog item.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"code": schema.StringAttribute{
				Required:    true,
				Description: "User provided unique identifier for the catalog item.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
			},
			"groups": schema.ListAttribute{
				Optional:    true,
				Description: "Catalog group codes this item belongs to.",
				ElementType: types.StringType,
			},
			"ref_type": schema.StringAttribute{
				Required:    true,
				Description: "Type of referenced entity that this catalog item points to.",
				Validators: []validator.String{
					stringvalidator.OneOf(stringSlice(api.AllowedCatalogItemTypeEnumValues)...),
				},
			},
			"ref_id": schema.StringAttribute{
				Required:    true,
				Description: "ID of the referenced entity.",
			},
			"account_ref": schema.Int64Attribute{
				Required:    true,
				Description: "Account number of the referenced financial account.",
			},
			"description": schema.MapAttribute{
				Optional:    true,
				Description: "Optional localized description of the catalog item.",
				ElementType: types.StringType,
			},
			"published": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether the catalog item is published and available for use.",
			},
			"required_scopes": schema.ListAttribute{
				Optional:    true,
				Description: "Account scopes required for this item to be available for subscription.",
				ElementType: types.StringType,
				Validators: []validator.List{
					listvalidator.UniqueValues(),
					listvalidator.ValueStringsAre(stringvalidator.OneOf(stringSlice(api.AllowedAccountScopeCodeEnumValues)...)),
				},
			},
			"dependencies": schema.ListAttribute{
				Optional:    true,
				Description: "Codes of catalog items this item depends on.",
				ElementType: types.StringType,
			},
			"patches": schema.ListNestedAttribute{
				Optional:    true,
				Description: "Patch definitions that can override attributes of the referenced entity.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"code": schema.StringAttribute{Required: true, Description: "Unique code for the patch."},
						"format": schema.StringAttribute{
							Required:    true,
							Description: "Patch format (for example JSON patch or unified diff).",
							Validators: []validator.String{
								stringvalidator.OneOf(stringSlice(api.AllowedPatchFormatEnumValues)...),
							},
						},
						"patch": schema.StringAttribute{Required: true, Description: "Patch payload."},
					},
				},
			},
		},
	}
}

func (r *CatalogItem) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *CatalogItem) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan catalogItemModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError("Error Creating Catalog Item", "Could not create API request, unexpected error: "+err.Error())
		return
	}

	item, _, err := r.client.GetApi().CreateCatalogItem(ctx).CreateOrUpdateCatalogItem(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Creating Catalog Item", "Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err))
		return
	}

	if err := plan.fromAPI(item); err != nil {
		resp.Diagnostics.AddError("Error Creating Catalog Item", "Could not process API response, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *CatalogItem) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state catalogItemModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	item, httpResp, err := r.client.GetApi().GetCatalogItem(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error Reading Catalog Item", "Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err))
		return
	}

	if err := state.fromAPI(item); err != nil {
		resp.Diagnostics.AddError("Error Reading Catalog Item", "Could not process API response, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *CatalogItem) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan catalogItemModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError("Error Updating Catalog Item", "Could not create API request, unexpected error: "+err.Error())
		return
	}

	item, _, err := r.client.GetApi().UpdateCatalogItem(ctx, plan.ID.ValueString()).CreateOrUpdateCatalogItem(updateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Updating Catalog Item", "Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err))
		return
	}

	if err := plan.fromAPI(item); err != nil {
		resp.Diagnostics.AddError("Error Updating Catalog Item", "Could not process API response, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *CatalogItem) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state catalogItemModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.client.GetApi().DeleteCatalogItem(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error Deleting Catalog Item", "Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err))
	}
}

func (r *CatalogItem) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("code"), req, resp)
}

func (model *catalogItemModel) fromAPI(item *api.CatalogItem) error {
	if item == nil {
		return fmt.Errorf("catalog item is nil")
	}

	model.ID = types.StringValue(item.Code)
	model.Code = types.StringValue(item.Code)
	model.RefType = types.StringValue(string(item.RefType))
	model.RefId = types.StringValue(item.RefId)
	model.AccountRef = types.Int64Value(item.AccountRef)

	model.Groups = omittableStringListValue(item.Groups, model.Groups)

	if item.Description != nil {
		mapped, diags := types.MapValue(types.StringType, stringMapToValueMap(*item.Description))
		if diags.HasError() {
			return fmt.Errorf("could not map description")
		}
		model.Description = mapped
	} else {
		model.Description = types.MapNull(types.StringType)
	}

	model.Published = omittableBooleanValue(item.Published, model.Published)

	model.RequiredScopes = omittableEnumListValue(item.RequiredScopes, model.RequiredScopes)

	model.Dependencies = omittableStringListValue(item.Dependencies, model.Dependencies)

	if item.Patches != nil {
		model.Patches = make([]catalogItemPatchDefinitionModel, len(item.Patches))
		for i := range item.Patches {
			model.Patches[i] = catalogItemPatchDefinitionModel{
				Code:   types.StringValue(item.Patches[i].Code),
				Format: types.StringValue(string(item.Patches[i].Format)),
				Patch:  types.StringValue(item.Patches[i].Patch),
			}
		}
	} else {
		if model.Patches != nil {
			model.Patches = make([]catalogItemPatchDefinitionModel, 0)
		} else {
			model.Patches = nil
		}
	}

	return nil
}

func (model *catalogItemModel) toAPIRequest() (api.CreateOrUpdateCatalogItem, error) {
	item := api.CreateOrUpdateCatalogItem{
		Code:       model.Code.ValueString(),
		RefType:    api.CatalogItemType(model.RefType.ValueString()),
		RefId:      model.RefId.ValueString(),
		AccountRef: model.AccountRef.ValueInt64(),
	}

	if groups := modelListToStringSlice(model.Groups); len(groups) > 0 {
		item.Groups = groups
	}

	if description := modelMapToStringMap(model.Description); len(description) > 0 {
		item.Description = &description
	}

	item.Published = model.Published.ValueBoolPointer()

	if requiredScopes := modelListToEnumSlice[api.AccountScopeCode](model.RequiredScopes); len(requiredScopes) > 0 {
		item.RequiredScopes = requiredScopes
	}

	if dependencies := modelListToStringSlice(model.Dependencies); len(dependencies) > 0 {
		item.Dependencies = dependencies
	}

	if len(model.Patches) > 0 {
		item.Patches = make([]api.CatalogItemPatchDefinition, len(model.Patches))
		for i, patch := range model.Patches {
			item.Patches[i] = api.CatalogItemPatchDefinition{
				Code:   patch.Code.ValueString(),
				Format: api.PatchFormat(patch.Format.ValueString()),
				Patch:  patch.Patch.ValueString(),
			}
		}
	}

	return item, nil
}
