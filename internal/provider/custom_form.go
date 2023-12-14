package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	v1 "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
)

type CustomFormModel struct {
	Layout []CustomFormItemModel `tfsdk:"layout"`	
	Group []CustomFormGroupModel `tfsdk:"groups"`	
	Input []CustomFormInputModel `tfsdk:"inputs"`
	Validation []CustomFormValidationModel `tfsdk:"validation"`	
}

type CustomFormItemModel struct {
	RefID types.String `tfsdk:"ref_id"`
	RefType types.String `tfsdk:"ref_type"`
}

func NewCustomFormGroupNull() CustomFormGroupModel {
	return CustomFormGroupModel{
		Label: basetypes.NewMapNull(types.StringType),
		Info: basetypes.NewMapNull(types.StringType),
	}
}

type CustomFormGroupModel struct {
	ID types.String `tfsdk:"id"`
	Type types.String `tfsdk:"type"`
	Label types.Map `tfsdk:"label"`
	Info types.Map `tfsdk:"info"`
	Items []CustomFormItemModel `tfsdk:"items"`
	Target types.String `tfsdk:"target"`
	TopDivider types.Bool `tfsdk:"top_divider"`
	BottomDivider types.Bool `tfsdk:"bottom_divider"`
	UseItemGroup types.Bool `tfsdk:"use_item_group"`
	MinSize types.Int64 `tfsdk:"min_size"`
}

func NewCustomFormInputNull() CustomFormInputModel {
	return CustomFormInputModel{
		Label: basetypes.NewMapNull(types.StringType),
		ViewLabel: basetypes.NewMapNull(types.StringType),
		Placeholder: basetypes.NewMapNull(types.StringType),
		Attributes: basetypes.NewMapNull(types.StringType),
		Items: basetypes.NewListNull(types.StringType),
	}
}

type CustomFormInputModel struct {
	ID types.String `tfsdk:"id"`
	Label types.Map `tfsdk:"label"`
	ViewLabel types.Map `tfsdk:"view_label"`
	Placeholder types.Map `tfsdk:"placeholder"`
	Type types.String `tfsdk:"type"`
	DefaultValue types.String `tfsdk:"default_value"`
	Required types.Bool `tfsdk:"required"`
	Ephemeral types.Bool `tfsdk:"ephemeral"`
	Target types.String `tfsdk:"target"`
	Attributes types.Map `tfsdk:"attributes"`
	Items types.List `tfsdk:"items"`
}

func NewCustomFormValidationNull() CustomFormValidationModel {
	return CustomFormValidationModel{
		Message: basetypes.NewMapNull(types.StringType),
	}
}

type CustomFormValidationModel struct {
	ID types.String `tfsdk:"id"`
	Expression types.String `tfsdk:"expression"`
	Target types.String `tfsdk:"target"`
	Message types.Map `tfsdk:"message"`
}

func CustomFormNestedSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"layout": schema.ListNestedAttribute{
			Required: true,
			NestedObject: customFormItemNestedSchema(),
			Description: "The layout of the form, which groups and inputs will be displayed",
		},
		"groups": schema.ListNestedAttribute{
			Optional: true,
			NestedObject: customFormGroupNestedSchema(),
			Description: "The groups of inputs that will be displayed to the user",
		},
		"inputs": schema.ListNestedAttribute{
			Required: true,
			NestedObject: customFormInputNestedSchema(),
			Description: "The inputs that will be displayed to the user",
		},
		"validation": schema.ListNestedAttribute{
			Optional: true,
			NestedObject: customFormValidationNestedSchema(),
			Description: "The validations that will be performed on the data provided by the user",
		},
	}
}

func customFormItemNestedSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"ref_id": schema.StringAttribute{
				Required: true,
				Description: "The ID of the entity that will be referenced",
			},
			"ref_type": schema.StringAttribute{
				Required: true,
				Description: "The type of the entity that will be referenced",
			},
		},
	}
}

func customFormGroupNestedSchema() schema.NestedAttributeObject {
	// TODO: Figure out what's actaully required and what isn't
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
				Description: "The ID of the group",
			},
			"type": schema.StringAttribute{
				Required: true,
				Description: "The type of the group",
			},
			"label": schema.MapAttribute{
				Required: true,
				ElementType: types.StringType,
				Description: "The label of the group",
			},
			"info": schema.MapAttribute{
				Required: true,
				ElementType: types.StringType,
				Description: "The info of the group",
			},
			"items": schema.ListNestedAttribute{
				Required: true,
				NestedObject: customFormItemNestedSchema(),
				Description: "The items of the group",
			},
			"target": schema.StringAttribute{
				Required: true,
				Description: "The target of the group",
			},
			"top_divider": schema.BoolAttribute{
				Required: true,
				Description: "The top divider of the group",
			},
			"bottom_divider": schema.BoolAttribute{
				Required: true,
				Description: "The bottom divider of the group",
			},
			"use_item_group": schema.BoolAttribute{
				Required: true,
				Description: "The use item group of the group",
			},
			"min_size": schema.Int64Attribute{
				Required: true,
				Description: "The min size of the group",
			},
		},
	}
}

func customFormInputNestedSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
				Description: "The ID of the input",
			},
			"label": schema.MapAttribute{
				Required: true,
				ElementType: types.StringType,
				Description: "The label of the input",
			},
			"view_label": schema.MapAttribute{
				Optional: true,
				ElementType: types.StringType,
				Description: "The view label of the input",
			},
			"placeholder": schema.MapAttribute{
				Optional: true,
				ElementType: types.StringType,
				Description: "The placeholder of the input",
			},
			"type": schema.StringAttribute{
				Required: true,
				Description: "The type of the input",
			},
			"default_value": schema.StringAttribute{
				Optional: true,
				Description: "The default value of the input",
			},
			"required": schema.BoolAttribute{
				Optional: true,
				Description: "The required of the input",
			},
			"ephemeral": schema.BoolAttribute{
				Optional: true,
				Description: "The ephemeral of the input",
			},
			"target": schema.StringAttribute{
				Optional: true,
				Description: "The target of the input",
			},
			"attributes": schema.MapAttribute{
				Optional: true,
				ElementType: types.StringType,
				Description: "The attributes of the input",
			},
			"items": schema.ListAttribute{
				Optional: true,
				ElementType: types.StringType,
				Description: "The items of the input",
			},
		},
	}
}

func customFormValidationNestedSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
				Description: "The ID of the validation",
			},
			"expression": schema.StringAttribute{
				Required: true,
				Description: "The expression of the validation",
			},
			"target": schema.StringAttribute{
				Required: true,
				Description: "The target of the validation",
			},
			"message": schema.MapAttribute{
				Required: true,
				ElementType: types.StringType,
				Description: "The message of the validation",
			},
		},
	}
}

func (form *CustomFormModel) toApiRequest() (*v1.CustomForm) {

	layout := make([]v1.FormItem, len(form.Layout))
	for i := range form.Layout {
		layout[i] = *form.Layout[i].toApiRequest()
	}

	groups := make([]v1.FormGroup, len(form.Group))
	for i := range form.Group {
		groups[i] = *form.Group[i].toApiRequest()
	}

	inputs := make([]v1.FormInput, len(form.Input))
	for i := range form.Input {
		inputs[i] = *form.Input[i].toApiRequest()
	}

	validations := make([]v1.FormValidation, len(form.Validation))
	for i := range form.Validation {
		validations[i] = *form.Validation[i].toApiRequest()
	}

	return &v1.CustomForm{
		Layout: layout,
		Groups: groups,
		Inputs: inputs,
		Validations: validations,
	}
}

func (form *CustomFormModel) fromApiResponse(resp *v1.CustomForm) (diags diag.Diagnostics) {

	for i := range resp.Layout {
		if len(form.Layout) <= i {
			form.Layout = append(form.Layout, CustomFormItemModel{})
		}
		form.Layout[i].fromApiResponse(&resp.Layout[i])
	}

	for i := range resp.Groups {
		if len(form.Group) <= i {
			form.Group = append(form.Group, NewCustomFormGroupNull())
		}
		moreDiags := form.Group[i].fromApiResponse(&resp.Groups[i])
		diags.Append(moreDiags...)
		if diags.HasError() {
			return
		}
	}

	for i := range resp.Inputs {
		if len(form.Input) <= i {
			form.Input = append(form.Input, NewCustomFormInputNull())
		}
		moreDiags := form.Input[i].fromApiResponse(&resp.Inputs[i])
		diags.Append(moreDiags...)
		if diags.HasError() {
			return
		}
	}

	for i := range resp.Validations {
		if len(form.Validation) <= i {
			form.Validation = append(form.Validation, NewCustomFormValidationNull())
		}
		moreDiags := form.Validation[i].fromApiResponse(&resp.Validations[i])
		diags.Append(moreDiags...)
		if diags.HasError() {
			return
		}
	}

	return
}

func (item *CustomFormItemModel) toApiRequest() (*v1.FormItem) {

	refType := v1.FormItemType(item.RefType.ValueString())

	return &v1.FormItem{
		RefId:   item.RefID.ValueStringPointer(),
		RefType: &refType,
	}	
}

func (item *CustomFormItemModel) fromApiResponse(resp *v1.FormItem) {
	
	item.RefID = types.StringValue(sdp(resp.RefId))
	if resp.RefType != nil {
		item.RefType = types.StringValue(string(*resp.RefType))
	} else {
		item.RefType = types.StringValue("")
	}
}

func (group *CustomFormGroupModel) toApiRequest() (*v1.FormGroup) {

	groupType := v1.FormGroupType(group.Type.ValueString())

	items := make([]v1.FormItem, len(group.Items))
	for i := range group.Items {
		items[i] = *group.Items[i].toApiRequest()
	}

	minSize := int32(group.MinSize.ValueInt64())

	return &v1.FormGroup{
		Id: group.ID.ValueStringPointer(),
		Type: &groupType,
		Label: modelMapToStringMap(group.Label),
		Info: modelMapToStringMap(group.Info),
		Items: items,
		Target: group.Target.ValueStringPointer(),
		TopDivider: group.TopDivider.ValueBoolPointer(),
		BottomDivider: group.BottomDivider.ValueBoolPointer(),
		UseItemGroup: group.UseItemGroup.ValueBoolPointer(),
		MinSize: &minSize,
	}
}

func (group *CustomFormGroupModel) fromApiResponse(resp *v1.FormGroup) (diags diag.Diagnostics) {

	groupType := ""
	if resp.Type != nil {
		groupType = string(*resp.Type)
	}

	minSize := int64(idp32(resp.MinSize))

	if resp.Label != nil {
		group.Label, diags = types.MapValue(types.StringType, stringMapToValueMap(*resp.Label))
		if diags.HasError() {
			return
		}
	}

	if resp.Info != nil {
		group.Info, diags = types.MapValue(types.StringType, stringMapToValueMap(*resp.Info))
		if diags.HasError() {
			return
		}
	}

	for i := range resp.Items {
		if len(group.Items) <= i {
			group.Items = append(group.Items, CustomFormItemModel{})
		}
		group.Items[i].fromApiResponse(&resp.Items[i])
	}

	group.ID = types.StringValue(sdp(resp.Id))
	group.Type = types.StringValue(groupType)
	group.Target = types.StringValue(sdp(resp.Target))
	group.TopDivider = types.BoolValue(bdp(resp.TopDivider))
	group.BottomDivider = types.BoolValue(bdp(resp.BottomDivider))
	group.UseItemGroup = types.BoolValue(bdp(resp.UseItemGroup))
	group.MinSize = types.Int64Value(minSize)

	return
}

func (input *CustomFormInputModel) toApiRequest() *v1.FormInput {

	inputType := v1.FormInputType(input.Type.ValueString())

	return &v1.FormInput{
		Id: input.ID.ValueStringPointer(),
		Label: modelMapToStringMap(input.Label),
		ViewLabel: modelMapToStringMap(input.ViewLabel),
		Placeholder: modelMapToStringMap(input.Placeholder),
		Type: &inputType,
		DefaultValue: input.DefaultValue.ValueStringPointer(),
		Required: input.Required.ValueBoolPointer(),
		Ephemeral: input.Ephemeral.ValueBoolPointer(),
		Target: input.Target.ValueStringPointer(),
		Attributes: modelMapToStringMap(input.Attributes),
		Items: modelListToStringSlice(input.Items),
	}
}

func (input *CustomFormInputModel) fromApiResponse(resp *v1.FormInput) (diags diag.Diagnostics) {

	inputType := ""
	if resp.Type != nil {
		inputType = string(*resp.Type)
	}

	if resp.HasId() {
		input.ID = types.StringValue(sdp(resp.Id))
	}
	if resp.HasType() {
		input.Type = types.StringValue(inputType)
	}

	if resp != nil {
		input.Label, diags = types.MapValue(types.StringType, stringMapToValueMap(*resp.Label))
		if diags.HasError() {
			return
		}
	}

	if resp.ViewLabel != nil {
		input.ViewLabel, diags = types.MapValue(types.StringType, stringMapToValueMap(*resp.ViewLabel))
		if diags.HasError() {
			return
		}
	}

	if resp.Placeholder != nil {
		input.Placeholder, diags = types.MapValue(types.StringType, stringMapToValueMap(*resp.Placeholder))
		if diags.HasError() {
			return
		}
	}

	if resp.HasDefaultValue() {
		input.DefaultValue = types.StringValue(sdp(resp.DefaultValue))
	}
	if resp.HasRequired() {
		input.Required = types.BoolValue(bdp(resp.Required))
	}
	if resp.HasEphemeral() {
		input.Ephemeral = types.BoolValue(bdp(resp.Ephemeral))
	}
	if resp.HasTarget() {
		input.Target = types.StringValue(sdp(resp.Target))
	}

	if resp.Attributes != nil {
		input.Attributes, diags = types.MapValue(types.StringType, stringMapToValueMap(*resp.Attributes))
		if diags.HasError() {
			return
		}
	}

	if resp.HasItems() {
		input.Items, diags = types.ListValue(types.StringType, stringSliceToValueList(resp.Items))
		if diags.HasError() {
			return
		}
	}

	return
}

func (validation *CustomFormValidationModel) toApiRequest() *v1.FormValidation {

	message := make(map[string]string)
	for key, value := range validation.Message.Elements() {
		message[key] = value.String()
	}

	return &v1.FormValidation{
		Id: validation.ID.ValueStringPointer(),
		Expression: validation.Expression.ValueStringPointer(),
		Target: validation.Target.ValueStringPointer(),
		Message: &message,
	}
}

func (validation *CustomFormValidationModel) fromApiResponse(resp *v1.FormValidation) (diags diag.Diagnostics) {

	message := make(map[string]attr.Value)
	if resp.Message != nil {
		for key, value := range *resp.Message {
			message[key] = types.StringValue(value)
		}
	}

	validation.ID = types.StringValue(sdp(resp.Id))
	validation.Expression = types.StringValue(sdp(resp.Expression))
	validation.Target = types.StringValue(sdp(resp.Target))
	validation.Message, diags = types.MapValue(types.StringType, message)

	return
}
