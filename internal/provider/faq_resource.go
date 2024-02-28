package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	api "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"
)

var (
	_ resource.ResourceWithConfigure   = &FAQ{}
	_ resource.ResourceWithImportState = &FAQ{}
)

func NewFAQ() resource.Resource {
	return &FAQ{}
}

type FAQ struct {
	client *client.UdomaClient
}

type FAQModel struct {
	ID       types.String   `tfsdk:"id"`
	Question types.Map      `tfsdk:"question"`
	Answer   types.Map      `tfsdk:"answer"`
	Keywords []types.String `tfsdk:"keywords"`
}

func (faq *FAQ) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_faq"
}

func (faq *FAQ) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents an faq",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the faq",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"question": schema.MapAttribute{
				Required:    true,
				Description: "The question to be answered in the faq in different languages",
				ElementType: types.StringType,
			},
			"answer": schema.MapAttribute{
				Required:    true,
				Description: "The answer for the faq question in different languages",
				ElementType: types.StringType,
			},
			"keywords": schema.ListAttribute{
				Optional:    true,
				Description: "The keywords for the faq for easier searching",
				ElementType: types.StringType,
			},
		},
	}
}

func (faq *FAQ) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	faq.client = cl
}

func (faq *FAQ) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {

	// Retrieve values from plan
	var plan FAQModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating FAQ Request",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newFAQ, _, err := faq.client.GetApi().CreateFAQEntry(ctx).CreateOrUpdateFAQEntryRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Creating FAQ Entry",
			"Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newFAQ); err != nil {
		resp.Diagnostics.AddError(
			"Error Creating FAQ Entry",
			"Could not process API response, unexpected error: "+err.Error(),
		)
		return
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (faq *FAQ) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {

	var state FAQModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	faqEntry, httpResp, err := faq.client.GetApi().GetFAQEntry(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading FAQ Entry",
			"Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := state.fromAPI(faqEntry); err != nil {
		resp.Diagnostics.AddError(
			"Error Reading FAQ Entry",
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

func (faq *FAQ) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {

	// Retrieve values from plan
	var plan FAQModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating FAQ Entry",
			"Could not create API request, unexpected error: "+err.Error(),
		)
		return
	}

	newFAQ, _, err := faq.client.GetApi().UpdateFAQEntry(ctx, plan.ID.ValueString()).CreateOrUpdateFAQEntryRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating FAQ Entry",
			"Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}

	// update the tf struct with the new values
	if err := plan.fromAPI(newFAQ); err != nil {
		resp.Diagnostics.AddError(
			"Error Updating FAQ Entry",
			"Could not process API response, unexpected error: "+err.Error(),
		)
		return
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (faq *FAQ) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

	// Retrieve values from state
	var state FAQModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := faq.client.GetApi().DeleteFAQEntry(ctx, state.ID.ValueString()).Execute()
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting FAQ Entry",
			"Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err),
		)
		return
	}
}

func (faq *FAQ) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *FAQModel) fromAPI(faq *api.FAQEntry) error {

	if faq == nil {
		return fmt.Errorf("faq entry is nil")
	}

	model.ID = types.StringValue(sdp(faq.Id))

	if faq.Question != nil {
		in := stringMapToValueMap(*faq.Question)
		modelValue, diags := types.MapValue(types.StringType, in)
		if diags.HasError() {
			return fmt.Errorf("error converting question to map: %v", diags)
		}
		model.Question = modelValue
	}
	if faq.Answer != nil {
		in := stringMapToValueMap(*faq.Answer)
		modelValue, diags := types.MapValue(types.StringType, in)
		if diags.HasError() {
			return fmt.Errorf("error converting answer to map: %v", diags)
		}
		model.Answer = modelValue
	}

	model.Keywords = make([]types.String, len(faq.Keywords))
	for i := range faq.Keywords {
		model.Keywords[i] = types.StringValue(faq.Keywords[i])
	}

	return nil
}

func (model *FAQModel) toAPIRequest() (api.CreateOrUpdateFAQEntryRequest, error) {

	faq := api.CreateOrUpdateFAQEntryRequest{
		Keywords: make([]string, len(model.Keywords)),
	}

	faq.Question = modelMapToStringMap(model.Question)
	faq.Answer = modelMapToStringMap(model.Answer)

	for i := range model.Keywords {
		faq.Keywords[i] = model.Keywords[i].ValueString()
	}

	return faq, nil
}
