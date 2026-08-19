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
	api "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
	"gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/client"
)

var (
	_ resource.ResourceWithConfigure   = &PublicChatbotInstruction{}
	_ resource.ResourceWithImportState = &PublicChatbotInstruction{}
)

func NewPublicChatbotInstruction() resource.Resource {
	return &PublicChatbotInstruction{}
}

type PublicChatbotInstruction struct {
	client *client.UdomaClient
}

type publicChatbotInstructionModel struct {
	ID          types.String `tfsdk:"id"`
	CreatedAt   types.Int64  `tfsdk:"created_at"`
	UpdatedAt   types.Int64  `tfsdk:"updated_at"`
	Name        types.String `tfsdk:"name"`
	Instruction types.String `tfsdk:"instruction"`
	Priority    types.Int32  `tfsdk:"priority"`
}

func (r *PublicChatbotInstruction) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_public_chatbot_instruction"
}

func (r *PublicChatbotInstruction) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Resource represents a single instruction in the public chatbot knowledge base.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The unique identifier for the public chatbot instruction.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.Int64Attribute{Computed: true, Description: "The date and time the instruction was created."},
			"updated_at": schema.Int64Attribute{Computed: true, Description: "The date and time the instruction was last modified."},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Short, descriptive name of the instruction topic.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
			},
			"instruction": schema.StringAttribute{
				Required:    true,
				Description: "The instruction text used by the chatbot as its knowledge base.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(20_000),
				},
			},
			"priority": schema.Int32Attribute{
				Optional:    true,
				Description: "Order in which instructions are passed to the chatbot; lower values are processed first.",
			},
		},
	}
}

func (r *PublicChatbotInstruction) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *PublicChatbotInstruction) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan publicChatbotInstructionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	createReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError("Error Creating Public Chatbot Instruction", "Could not create API request, unexpected error: "+err.Error())
		return
	}

	instruction, _, err := r.client.GetApi().CreatePublicChatbotInstruction(ctx).CreateOrUpdatePublicChatbotInstructionRequest(createReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Creating Public Chatbot Instruction", "Could not create entity in Udoma, unexpected error: "+getApiErrorMessage(err))
		return
	}

	if err := plan.fromAPI(instruction); err != nil {
		resp.Diagnostics.AddError("Error Creating Public Chatbot Instruction", "Could not process API response, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *PublicChatbotInstruction) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state publicChatbotInstructionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	instruction, httpResp, err := r.client.GetApi().GetPublicChatbotInstruction(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error Reading Public Chatbot Instruction", "Could not read entity from Udoma, unexpected error: "+getApiErrorMessage(err))
		return
	}

	if err := state.fromAPI(instruction); err != nil {
		resp.Diagnostics.AddError("Error Reading Public Chatbot Instruction", "Could not process API response, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *PublicChatbotInstruction) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan publicChatbotInstructionModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateReq, err := plan.toAPIRequest()
	if err != nil {
		resp.Diagnostics.AddError("Error Updating Public Chatbot Instruction", "Could not create API request, unexpected error: "+err.Error())
		return
	}

	instruction, _, err := r.client.GetApi().UpdatePublicChatbotInstruction(ctx, plan.ID.ValueString()).CreateOrUpdatePublicChatbotInstructionRequest(updateReq).Execute()
	if err != nil {
		resp.Diagnostics.AddError("Error Updating Public Chatbot Instruction", "Could not update entity in Udoma, unexpected error: "+getApiErrorMessage(err))
		return
	}

	if err := plan.fromAPI(instruction); err != nil {
		resp.Diagnostics.AddError("Error Updating Public Chatbot Instruction", "Could not process API response, unexpected error: "+err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *PublicChatbotInstruction) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state publicChatbotInstructionModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	httpResp, err := r.client.GetApi().DeletePublicChatbotInstruction(ctx, state.ID.ValueString()).Execute()
	if httpResp != nil && httpResp.StatusCode == 404 {
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error Deleting Public Chatbot Instruction", "Could not delete entity in Udoma, unexpected error: "+getApiErrorMessage(err))
	}
}

func (r *PublicChatbotInstruction) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (model *publicChatbotInstructionModel) fromAPI(instruction *api.PublicChatbotInstruction) error {
	if instruction == nil {
		return fmt.Errorf("public chatbot instruction is nil")
	}

	model.ID = types.StringValue(instruction.Id)
	model.CreatedAt = types.Int64Value(instruction.CreatedAt)
	model.UpdatedAt = types.Int64Value(instruction.UpdatedAt)
	model.Name = types.StringValue(instruction.Name)
	model.Instruction = types.StringValue(instruction.Instruction)
	model.Priority = types.Int32PointerValue(instruction.Priority)

	return nil
}

func (model *publicChatbotInstructionModel) toAPIRequest() (api.CreateOrUpdatePublicChatbotInstructionRequest, error) {
	req := api.CreateOrUpdatePublicChatbotInstructionRequest{
		Name:        model.Name.ValueString(),
		Instruction: model.Instruction.ValueString(),
		Priority:    model.Priority.ValueInt32Pointer(),
	}

	return req, nil
}
