package provider

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
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
	_ resource.ResourceWithConfigure   = &EntityExtension{}
	_ resource.ResourceWithImportState = &EntityExtension{}
)

func NewEntityExtension() resource.Resource {
	return &EntityExtension{}
}

type EntityExtension struct {
	client *client.UdomaClient
}

type EntityExtensionModel struct {
	ID             types.String `tfsdk:"id"`
	CreatedAt      types.Int64  `tfsdk:"created_at"`
	UpdatedAt      types.Int64  `tfsdk:"updated_at"`
	Name           types.String `tfsdk:"name"`
	Description    types.String `tfsdk:"description"`
	Entity         types.String `tfsdk:"entity"`
	Sequence       types.Int64  `tfsdk:"sequence"`
	Key            types.String `tfsdk:"key"`
	Version        types.Int64  `tfsdk:"version"`
	FormDefinition *CustomFormModel `tfsdk:"form_definition"`
}

func (r *EntityExtension) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_entity_extension"
}

func (r *EntityExtension) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the entity extension.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The creation timestamp of the entity extension.",
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: "The last update timestamp of the entity extension.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "A name for the entity extension.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "A description of what this extension is used for.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(1024),
				},
			},
			"entity": schema.StringAttribute{
				Required:    true,
				Description: "The type of entity this extension applies to (property, tenant, etc.).",
				Validators: []validator.String{
					stringvalidator.OneOf("property", "tenant", "service_provider", "tenancy", "owner"),
				},
			},
			"sequence": schema.Int64Attribute{
				Optional:    true,
				Description: "The order in which extensions should be displayed in the UI. Lower numbers appear first.",
			},
			"key": schema.StringAttribute{
				Required:    true,
				Description: "A unique identifier for this extension within the account and entity type.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
			},
			"version": schema.Int64Attribute{
				Optional:    true,
				Description: "The version number of this extension definition.",
			},
			"form_definition": schema.SingleNestedAttribute{
				Required: true,
				Description: "The custom form definition.",
				Attributes: map[string]schema.Attribute{
					"layout": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"ref_id": schema.StringAttribute{
									Required: true,
								},
								"ref_type": schema.StringAttribute{
									Required: true,
								},
							},
						},
					},
					"inputs": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
							},
						},
					},

					"groups": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
							},
						},
					},

					"validation": schema.ListNestedAttribute{
						Optional: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
							},
						},
					},
				},
			},
		},
	}
}

func (r *EntityExtension) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

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

func (r *EntityExtension) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	var plan EntityExtensionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiReq, toApiDiags := plan.toApiRequest()
	resp.Diagnostics.Append(toApiDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	log.Printf("API Request: %+v", apiReq)

	created, _, err := r.client.GetApi().
		CreateEntityExtension(ctx).
		CreateOrUpdateEntityExtensionRequest(*apiReq).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to create entity extension",
			err.Error(),
		)
		return
	}

	fromApiDiags := plan.fromApiResponse(created)
	resp.Diagnostics.Append(fromApiDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	setDiags := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(setDiags...)
}

func (r *EntityExtension) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state EntityExtensionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	found, httpResp, err := r.client.GetApi().GetEntityExtension(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Entity Extension",
			"Could not read entity extension from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	diags = state.fromApiResponse(found)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func (r *EntityExtension) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan EntityExtensionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	apiReq, toApiDiags := plan.toApiRequest()
	resp.Diagnostics.Append(toApiDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	id := plan.ID.ValueString()

	updated, _, err := r.client.GetApi().
		UpdateEntityExtension(ctx, id).
		CreateOrUpdateEntityExtensionRequest(*apiReq).
		Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Entity Extension",
			fmt.Sprintf("Could not update entity extension in Udoma, unexpected error: %s", getApiErrorMessage(err)),
		)
		return
	}

	fromApiDiags := plan.fromApiResponse(updated)
	resp.Diagnostics.Append(fromApiDiags...)
	if resp.Diagnostics.HasError() {
		return
	}

	setDiags := resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(setDiags...)
}

func (r *EntityExtension) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	var state EntityExtensionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.client.GetApi().DeleteEntityExtension(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Entity Extension",
			"Could not delete entity extension in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
	}
}

func (r *EntityExtension) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *EntityExtensionModel) toApiRequest() (*v1.CreateOrUpdateEntityExtensionRequest, diag.Diagnostics) {
	var diags diag.Diagnostics

	// Convert form_definition to API type
	cf := model.FormDefinition.toApiRequest()

	// Ensure slices are not nil
	if cf.Layout == nil {
		cf.Layout = []v1.FormItem{}
	}
	if cf.Inputs == nil {
		cf.Inputs = []v1.FormInput{}
	}
	if cf.Groups == nil {
		cf.Groups = []v1.FormGroup{}
	}
	if cf.Validations == nil {
		cf.Validations = []v1.FormValidation{}
	}

	req := &v1.CreateOrUpdateEntityExtensionRequest{
		Name:           model.Name.ValueString(),
		Key:            model.Key.ValueString(),
		Entity:         v1.EntityExtensionTypeEnum(model.Entity.ValueString()),
		FormDefinition: *v1.NewNullableCustomForm(&cf),
	}

	if !model.Description.IsNull() {
		desc := model.Description.ValueString()
		req.Description = &desc
	}
	if !model.Sequence.IsNull() {
		seq := int32(model.Sequence.ValueInt64())
		req.Sequence = &seq
	}
	if !model.Version.IsNull() {
		ver := int32(model.Version.ValueInt64())
		req.Version = &ver
	}

	return req, diags
}

func (model *EntityExtensionModel) fromApiResponse(resp *v1.EntityExtension) diag.Diagnostics {
	var diags diag.Diagnostics

	model.ID = types.StringValue(resp.Id)
	model.Name = types.StringValue(resp.Name)
	model.Entity = types.StringValue(string(resp.Entity))
	model.Key = types.StringValue(resp.Key)
	model.CreatedAt = types.Int64Value(int64(resp.CreatedAt))
	model.UpdatedAt = types.Int64Value(int64(resp.UpdatedAt))

	if resp.Description != nil {
		model.Description = types.StringValue(*resp.Description)
	} else {
		model.Description = types.StringNull()
	}

	if resp.Sequence != nil {
		model.Sequence = types.Int64Value(int64(*resp.Sequence))
	} else {
		model.Sequence = types.Int64Null()
	}

	if resp.Version != nil {
		model.Version = types.Int64Value(int64(*resp.Version))
	} else {
		model.Version = types.Int64Null()
	}

	form := &CustomFormModel{
		Layout:     []CustomFormItemModel{},
		Input:      []CustomFormInputModel{},
		Group:      []CustomFormGroupModel{},
		Validation: []CustomFormValidationModel{},
	}

	if resp.FormDefinition.IsSet() && resp.FormDefinition.Get() != nil {
		cf := resp.FormDefinition.Get()
		form.fromApiResponse(cf)

		// Ensure slices are initialized to empty
		if form.Layout == nil {
			form.Layout = []CustomFormItemModel{}
		}
		if form.Input == nil {
			form.Input = []CustomFormInputModel{}
		}
		if form.Group == nil {
			form.Group = []CustomFormGroupModel{}
		}
		if form.Validation == nil {
			form.Validation = []CustomFormValidationModel{}
		}
	}

	model.FormDefinition = form

	return diags
}
