package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	v1 "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/internal/migrations/formItems"
)

type CustomFormModelV0 struct {
	Layout     []CustomFormItemModelV0       `tfsdk:"layout"`
	Group      []CustomFormGroupModelV0      `tfsdk:"groups"`
	Input      []CustomFormInputModelV0      `tfsdk:"inputs"`
	Validation []CustomFormValidationModelV0 `tfsdk:"validation"`
}

type CustomFormItemModelV0 struct {
	RefID   types.String `tfsdk:"ref_id"`
	RefType types.String `tfsdk:"ref_type"`
}

func NewCustomFormGroupNullV0() CustomFormGroupModelV0 {
	return CustomFormGroupModelV0{
		Label: basetypes.NewMapNull(types.StringType),
		Info:  basetypes.NewMapNull(types.StringType),
	}
}

type CustomFormGroupModelV0 struct {
	ID            types.String            `tfsdk:"id"`
	Type          types.String            `tfsdk:"type"`
	Label         types.Map               `tfsdk:"label"`
	Info          types.Map               `tfsdk:"info"`
	Items         []CustomFormItemModelV0 `tfsdk:"items"`
	Target        types.String            `tfsdk:"target"`
	TopDivider    types.Bool              `tfsdk:"top_divider"`
	BottomDivider types.Bool              `tfsdk:"bottom_divider"`
	UseItemGroup  types.Bool              `tfsdk:"use_item_group"`
	MinSize       types.Int64             `tfsdk:"min_size"`
	Icon          types.String            `tfsdk:"icon"`
}

func NewCustomFormInputNullV0() CustomFormInputModelV0 {
	return CustomFormInputModelV0{
		Label:       basetypes.NewMapNull(types.StringType),
		ViewLabel:   basetypes.NewMapNull(types.StringType),
		Placeholder: basetypes.NewMapNull(types.StringType),
		Attributes:  basetypes.NewMapNull(types.StringType),
		Items:       basetypes.NewListNull(types.StringType),
	}
}

type CustomFormInputModelV0 struct {
	ID               types.String `tfsdk:"id"`
	Label            types.Map    `tfsdk:"label"`
	ViewLabel        types.Map    `tfsdk:"view_label"`
	Placeholder      types.Map    `tfsdk:"placeholder"`
	Type             types.String `tfsdk:"type"`
	DefaultValue     types.String `tfsdk:"default_value"`
	Required         types.Bool   `tfsdk:"required"`
	Ephemeral        types.Bool   `tfsdk:"ephemeral"`
	PropagateChanges types.Bool   `tfsdk:"propagate_changes"`
	Target           types.String `tfsdk:"target"`
	Attributes       types.Map    `tfsdk:"attributes"`
	Items            types.List   `tfsdk:"items"`
}

func NewCustomFormValidationNullV0() CustomFormValidationModelV0 {
	return CustomFormValidationModelV0{
		Message: basetypes.NewMapNull(types.StringType),
	}
}

type CustomFormValidationModelV0 struct {
	ID         types.String `tfsdk:"id"`
	Expression types.String `tfsdk:"expression"`
	Target     types.String `tfsdk:"target"`
	Message    types.Map    `tfsdk:"message"`
}

func CustomFormNestedSchemaV0() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"layout": schema.ListNestedAttribute{
			Required:     true,
			NestedObject: customFormItemNestedSchemaV0(),
		},
		"groups": schema.ListNestedAttribute{
			Optional:     true,
			NestedObject: customFormGroupNestedSchemaV0(),
		},
		"inputs": schema.ListNestedAttribute{
			Required:     true,
			NestedObject: customFormInputNestedSchemaV0(),
		},
		"validation": schema.ListNestedAttribute{
			Optional:     true,
			NestedObject: customFormValidationNestedSchemaV0(),
		},
	}
}

func customFormItemNestedSchemaV0() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"ref_id": schema.StringAttribute{
				Required: true,
			},
			"ref_type": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func customFormGroupNestedSchemaV0() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"type": schema.StringAttribute{
				Required: true,
			},
			"label": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"info": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"items": schema.ListNestedAttribute{
				Required:     true,
				NestedObject: customFormItemNestedSchemaV0(),
			},
			"target": schema.StringAttribute{
				Optional: true,
			},
			"top_divider": schema.BoolAttribute{
				Optional: true,
			},
			"bottom_divider": schema.BoolAttribute{
				Optional: true,
			},
			"use_item_group": schema.BoolAttribute{
				Optional: true,
			},
			"min_size": schema.Int64Attribute{
				Optional: true,
			},
			"icon": schema.StringAttribute{
				Optional: true,
			},
		},
	}
}

func customFormInputNestedSchemaV0() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"label": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"view_label": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"placeholder": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"type": schema.StringAttribute{
				Required: true,
			},
			"default_value": schema.StringAttribute{
				Optional: true,
			},
			"required": schema.BoolAttribute{
				Optional: true,
			},
			"ephemeral": schema.BoolAttribute{
				Optional: true,
			},
			"propagate_changes": schema.BoolAttribute{
				Optional: true,
			},
			"target": schema.StringAttribute{
				Optional: true,
			},
			"attributes": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
			"items": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func customFormValidationNestedSchemaV0() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
			},
			"expression": schema.StringAttribute{
				Required: true,
			},
			"target": schema.StringAttribute{
				Required: true,
			},
			"message": schema.MapAttribute{
				Required:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func (form *CustomFormModelV0) toApiRequest() v1.CustomForm {

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

	return v1.CustomForm{
		Layout:      layout,
		Groups:      groups,
		Inputs:      inputs,
		Validations: validations,
	}
}

func (item *CustomFormItemModelV0) toApiRequest() *v1.FormItem {
	return &v1.FormItem{
		RefId:   item.RefID.ValueString(),
		RefType: v1.FormItemType(item.RefType.ValueString()),
	}
}

func (group *CustomFormGroupModelV0) toApiRequest() *v1.FormGroup {

	items := make([]v1.FormItem, len(group.Items))
	for i := range group.Items {
		items[i] = *group.Items[i].toApiRequest()
	}

	minSize := int32(group.MinSize.ValueInt64())

	label := modelMapToStringMap(group.Label)
	info := modelMapToStringMap(group.Info)

	return &v1.FormGroup{
		Id:            group.ID.ValueString(),
		Type:          v1.FormGroupType(group.Type.ValueString()),
		Label:         &label,
		Info:          &info,
		Items:         items,
		Target:        group.Target.ValueStringPointer(),
		TopDivider:    group.TopDivider.ValueBoolPointer(),
		BottomDivider: group.BottomDivider.ValueBoolPointer(),
		UseItemGroup:  group.UseItemGroup.ValueBoolPointer(),
		MinSize:       &minSize,
	}
}

func (input *CustomFormInputModelV0) toApiRequest() *v1.FormInput {

	ret := &v1.FormInput{
		Id:               input.ID.ValueString(),
		Type:             v1.FormInputType(input.Type.ValueString()),
		DefaultValue:     input.DefaultValue.ValueStringPointer(),
		Required:         input.Required.ValueBoolPointer(),
		Ephemeral:        input.Ephemeral.ValueBoolPointer(),
		PropagateChanges: input.PropagateChanges.ValueBoolPointer(),
		Target:           input.Target.ValueStringPointer(),
		Items:            modelListToStringSlice(input.Items),
	}

	if label := modelMapToStringMap(input.Label); len(label) > 0 {
		ret.Label = &label
	}

	if viewLabel := modelMapToStringMap(input.ViewLabel); len(viewLabel) > 0 {
		ret.ViewLabel = &viewLabel
	}

	if placeholder := modelMapToStringMap(input.Placeholder); len(placeholder) > 0 {
		ret.Placeholder = &placeholder
	}

	if attributes := modelMapToStringMap(input.Attributes); len(attributes) > 0 {
		ret.Attributes = &attributes
	}

	return ret
}

func (validation *CustomFormValidationModelV0) toApiRequest() *v1.FormValidation {
	return &v1.FormValidation{
		Id:         validation.ID.ValueString(),
		Expression: validation.Expression.ValueString(),
		Target:     validation.Target.ValueString(),
		Message:    modelMapToStringMap(validation.Message),
	}
}
