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
	Layout     []CustomFormItemModel       `tfsdk:"layout"`
	Group      []CustomFormGroupModel      `tfsdk:"groups"`
	Input      []CustomFormInputModel      `tfsdk:"inputs"`
	Validation []CustomFormValidationModel `tfsdk:"validation"`
}

type CustomFormItemModel struct {
	RefID   types.String `tfsdk:"ref_id"`
	RefType types.String `tfsdk:"ref_type"`
}

func NewCustomFormGroupNull() CustomFormGroupModel {
	return CustomFormGroupModel{
		Label: basetypes.NewMapNull(types.StringType),
		Info:  basetypes.NewMapNull(types.StringType),
	}
}

type CustomFormGroupModel struct {
	ID            types.String          `tfsdk:"id"`
	Type          types.String          `tfsdk:"type"`
	Label         types.Map             `tfsdk:"label"`
	Info          types.Map             `tfsdk:"info"`
	Subtitle      types.List            `tfsdk:"subtitle"`
	NestedDisplay types.Bool            `tfsdk:"nested_display"`
	Items         []CustomFormItemModel `tfsdk:"items"`
	Target        types.String          `tfsdk:"target"`
	TopDivider    types.Bool            `tfsdk:"top_divider"`
	BottomDivider types.Bool            `tfsdk:"bottom_divider"`
	MinSize       types.Int64           `tfsdk:"min_size"`
	Icon          types.String          `tfsdk:"icon"`
}

func NewCustomFormInputItemNull() CustomFormInputItemModel {
	return CustomFormInputItemModel{
		Label: basetypes.NewMapNull(types.StringType),
	}
}

type CustomFormInputItemModel struct {
	ID    types.String `tfsdk:"id"`
	Label types.Map    `tfsdk:"label"`
}

func NewConditionalInputNull() DisplayConditionModel {
	return DisplayConditionModel{}
}

type DisplayConditionModel struct {
	Source  types.String `tfsdk:"source"`
	Operand types.String `tfsdk:"operand"`
	Value   types.String `tfsdk:"value"`
}

func NewCustomFormInputNull() CustomFormInputModel {
	return CustomFormInputModel{
		Label:       basetypes.NewMapNull(types.StringType),
		ViewLabel:   basetypes.NewMapNull(types.StringType),
		Placeholder: basetypes.NewMapNull(types.StringType),
		Tooltip:     basetypes.NewMapNull(types.StringType),
		Attributes:  basetypes.NewMapNull(types.StringType),
	}
}

type CustomFormInputModel struct {
	ID               types.String               `tfsdk:"id"`
	Label            types.Map                  `tfsdk:"label"`
	ViewLabel        types.Map                  `tfsdk:"view_label"`
	Placeholder      types.Map                  `tfsdk:"placeholder"`
	Tooltip          types.Map                  `tfsdk:"tooltip"`
	Type             types.String               `tfsdk:"type"`
	DefaultValue     types.String               `tfsdk:"default_value"`
	Required         types.Bool                 `tfsdk:"required"`
	Readonly         types.Bool                 `tfsdk:"readonly"`
	Ephemeral        types.Bool                 `tfsdk:"ephemeral"`
	PropagateChanges types.Bool                 `tfsdk:"propagate_changes"`
	Target           types.String               `tfsdk:"target"`
	Attributes       types.Map                  `tfsdk:"attributes"`
	Items            []CustomFormInputItemModel `tfsdk:"items"`
	DisplayCondition *DisplayConditionModel     `tfsdk:"display_condition"`
}

func NewCustomFormValidationNull() CustomFormValidationModel {
	return CustomFormValidationModel{
		Message: basetypes.NewMapNull(types.StringType),
	}
}

type CustomFormValidationModel struct {
	ID         types.String `tfsdk:"id"`
	Expression types.String `tfsdk:"expression"`
	Target     types.String `tfsdk:"target"`
	Message    types.Map    `tfsdk:"message"`
}

func CustomFormNestedSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"layout": schema.ListNestedAttribute{
			Required:     true,
			NestedObject: customFormItemNestedSchema(),
			Description:  "The layout of the form, which groups and inputs will be displayed",
		},
		"groups": schema.ListNestedAttribute{
			Optional:     true,
			NestedObject: customFormGroupNestedSchema(),
			Description:  "The groups of inputs that will be displayed to the user",
		},
		"inputs": schema.ListNestedAttribute{
			Required:     true,
			NestedObject: customFormInputNestedSchema(),
			Description:  "The inputs that will be displayed to the user",
		},
		"validation": schema.ListNestedAttribute{
			Optional:     true,
			NestedObject: customFormValidationNestedSchema(),
			Description:  "The validations that will be performed on the data provided by the user",
		},
	}
}

func customFormItemNestedSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"ref_id": schema.StringAttribute{
				Required:    true,
				Description: "The ID of the entity that will be referenced",
			},
			"ref_type": schema.StringAttribute{
				Required:    true,
				Description: "The type of the entity that will be referenced",
			},
		},
	}
}

func customFormGroupNestedSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    true,
				Description: "The ID of the group",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "The type of the group",
			},
			"label": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "The label of the group",
			},
			"subtitle": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Optional subtitle for repeat groups. This will be computed dynamically based on the " +
					"data provided by the user and will be displayed as a subtitle in the group accordion. The " +
					"values are interpreted as JSON paths based on the data for the group.",
			},
			"info": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "The info of the group",
			},
			"nested_display": schema.BoolAttribute{
				Optional: true,
				Description: "If true, the nested group will be displayed in a nested UI (only works on mobile " +
					"and tablet). This is useful for more complex group that require more space to be displayed. " +
					"The group will be displayed in a separate screen, and the user will be able to navigate back " +
					"and forth between the group and the main form.",
			},
			"items": schema.ListNestedAttribute{
				Required:     true,
				NestedObject: customFormItemNestedSchema(),
				Description:  "the IDs of the inputs that will be displayed in the group",
			},
			"target": schema.StringAttribute{
				Optional:    true,
				Description: "the attribute name to use when exporting the result of this group (only used for repeat groups)",
			},
			"top_divider": schema.BoolAttribute{
				Optional:    true,
				Description: "if true, a divider will be displayed above the group",
			},
			"bottom_divider": schema.BoolAttribute{
				Optional:    true,
				Description: "if true, a divider will be displayed below the group",
			},
			"min_size": schema.Int64Attribute{
				Optional:    true,
				Description: "the minimum number of items that must be submitted in the group (only used for repeat groups)",
			},
			"icon": schema.StringAttribute{
				Optional:    true,
				Description: "the icon to display (only used for pages)",
			},
		},
	}
}

func customFormInputItemNestedSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    true,
				Description: "The technical value that will be stored in the collected data",
			},
			"label": schema.MapAttribute{
				Required:    true,
				ElementType: types.StringType,
				Description: "The label to be displayed to the user as a language dictionary",
			},
		},
	}
}

func customFormInputNestedSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
				Description: "The ID of the input field, used to identify it within the form. If no " +
					"target value is set, this will also be how the value is stored and later " +
					"accessed in the collected data.",
			},
			"target": schema.StringAttribute{
				Optional: true,
				Description: "Optional attribute name to use when exporting the value of this input. If " +
					"not set, the ID will be used.",
			},
			"type": schema.StringAttribute{
				Required: true,
				Description: "The type of data that will be retrieved by that input. This also controls " +
					"how the input will be displayed to the user.",
			},
			"label": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "The label for the input field as a language dictionary (e.g. 'What is your name?')",
			},
			"view_label": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Optional label to be displayed when viewing the input field (e.g. 'Name'). " +
					"This is useful when the label is not suitable to be displayed in a view.",
			},
			"placeholder": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "An optional placeholder to use for the input field. If not set, a default " +
					"one will be used based on the label of the input.",
			},
			"tooltip": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Optional tooltip to be displayed next to the label (hover or click on icon). " +
					"This is useful for providing additional information to the user.",
			},
			"default_value": schema.StringAttribute{
				Optional:    true,
				Description: "Optional default value for the input field (as a JSON string)",
			},
			"required": schema.BoolAttribute{
				Optional:    true,
				Description: "If true, the user will be required to provide a valid value for this input",
			},
			"readonly": schema.BoolAttribute{
				Optional:    true,
				Description: "If true, the user will not be able to change the value of this input",
			},
			"ephemeral": schema.BoolAttribute{
				Optional: true,
				Description: "If true, the value of the input will not be persisted. This is useful for " +
					"checkboxes for accepting terms and conditions, etc.",
			},
			"propagate_changes": schema.BoolAttribute{
				Optional: true,
				Description: "If true, changes to the input will be propagated to event listeners for " +
					"the custom form.",
			},
			"attributes": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Optional set of settings for the input field, such as min/max values. " +
					"Those will vary depending on the type.",
			},
			"items": schema.ListNestedAttribute{
				Optional:     true,
				NestedObject: customFormInputItemNestedSchema(),
				Description: "Only used when the type is select or multi select. This is a list of " +
					"values that the user can choose from.",
			},
			"display_condition": schema.SingleNestedAttribute{
				Optional:   true,
				Attributes: conditionalInputNestedSchema(),
				Description: "Optional condition that must be met for the input to be displayed. If the " +
					"condition is not met, the input will be hidden. This overrides the required " +
					"attribute.",
			},
		},
	}
}

func conditionalInputNestedSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"source": schema.StringAttribute{
			Required:    true,
			Description: "the source field, which will be checked",
		},
		"operand": schema.StringAttribute{
			Required:    true,
			Description: "the operand of the condition",
		},
		"value": schema.StringAttribute{
			Required:    true,
			Description: "the value of the condition, as serialized JSON",
		},
	}
}

func customFormValidationNestedSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    true,
				Description: "the ID of the validation, used to identify it",
			},
			"expression": schema.StringAttribute{
				Required:    true,
				Description: "a JS expression that will be evaluated to determine if the validation passes or fails",
			},
			"target": schema.StringAttribute{
				Required:    true,
				Description: "the index of the input should be highlighted if the validation fails (nesting is supported via dot notation)",
			},
			"message": schema.MapAttribute{
				Required:    true,
				ElementType: types.StringType,
				Description: "a map of values, where the key and values are strings",
			},
		},
	}
}

func (form *CustomFormModel) toApiRequest() v1.CustomForm {

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

func (item *CustomFormItemModel) toApiRequest() *v1.FormItem {
	return &v1.FormItem{
		RefId:   item.RefID.ValueString(),
		RefType: v1.FormItemType(item.RefType.ValueString()),
	}
}

func (item *CustomFormItemModel) fromApiResponse(resp *v1.FormItem) {
	item.RefID = types.StringValue(resp.RefId)
	item.RefType = types.StringValue(string(resp.RefType))
}

func (group *CustomFormGroupModel) toApiRequest() *v1.FormGroup {

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
		Subtitle:      modelListToStringSlice(group.Subtitle),
		NestedDisplay: group.NestedDisplay.ValueBoolPointer(),
		Items:         items,
		Target:        group.Target.ValueStringPointer(),
		TopDivider:    group.TopDivider.ValueBoolPointer(),
		BottomDivider: group.BottomDivider.ValueBoolPointer(),
		MinSize:       &minSize,
		Icon:          group.Icon.ValueStringPointer(),
	}
}

func (group *CustomFormGroupModel) fromApiResponse(resp *v1.FormGroup) (diags diag.Diagnostics) {

	group.ID = types.StringValue(resp.Id)
	group.Type = types.StringValue(string(resp.Type))

	if resp.Subtitle != nil {
		group.Subtitle, diags = types.ListValue(types.StringType, stringSliceToValueList(resp.Subtitle))
		if diags.HasError() {
			return
		}
	} else {
		group.Subtitle = types.ListNull(types.StringType)
	}

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

	if resp.NestedDisplay != nil {
		group.NestedDisplay = types.BoolValue(*resp.NestedDisplay)
	}

	for i := range resp.Items {
		if len(group.Items) <= i {
			group.Items = append(group.Items, CustomFormItemModel{})
		}
		group.Items[i].fromApiResponse(&resp.Items[i])
	}

	if resp.Target != nil {
		group.Target = types.StringValue(*resp.Target)
	}
	if resp.TopDivider != nil {
		group.TopDivider = types.BoolValue(*resp.TopDivider)
	}
	if resp.BottomDivider != nil {
		group.BottomDivider = types.BoolValue(*resp.BottomDivider)
	}
	if resp.MinSize != nil {
		group.MinSize = types.Int64Value(int64(idp32(resp.MinSize)))
	}

	if resp.Icon != nil {
		group.Icon = types.StringValue(*resp.Icon)
	}

	return
}

func (input *CustomFormInputItemModel) toApiRequest() *v1.InputItem {
	return &v1.InputItem{
		Id:    input.ID.ValueString(),
		Label: modelMapToStringMap(input.Label),
	}
}

func (input *CustomFormInputItemModel) fromApiResponse(resp *v1.InputItem) (diags diag.Diagnostics) {

	input.ID = types.StringValue(resp.Id)
	input.Label, diags = types.MapValue(types.StringType, stringMapToValueMap(resp.Label))
	if diags.HasError() {
		return
	}

	return
}

func (input *CustomFormInputModel) toApiRequest() *v1.FormInput {

	items := make([]v1.InputItem, len(input.Items))
	for i := range input.Items {
		items[i] = *input.Items[i].toApiRequest()
	}

	ret := &v1.FormInput{
		Id:               input.ID.ValueString(),
		Type:             v1.FormInputType(input.Type.ValueString()),
		DefaultValue:     input.DefaultValue.ValueStringPointer(),
		Required:         input.Required.ValueBoolPointer(),
		Readonly:         input.Readonly.ValueBoolPointer(),
		Ephemeral:        input.Ephemeral.ValueBoolPointer(),
		PropagateChanges: input.PropagateChanges.ValueBoolPointer(),
		Target:           input.Target.ValueStringPointer(),
		Items:            items,
	}

	if input.DisplayCondition != nil {
		ret.DisplayCondition = input.DisplayCondition.toApiRequest()
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

	if tooltip := modelMapToStringMap(input.Tooltip); len(tooltip) > 0 {
		ret.Tooltip = &tooltip
	}

	if attributes := modelMapToStringMap(input.Attributes); len(attributes) > 0 {
		ret.Attributes = &attributes
	}

	return ret
}

func (input *CustomFormInputModel) fromApiResponse(resp *v1.FormInput) (diags diag.Diagnostics) {

	input.ID = types.StringValue(resp.Id)
	input.Type = types.StringValue(string(resp.Type))

	if resp.Label != nil {
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

	if resp.Tooltip != nil {
		input.Tooltip, diags = types.MapValue(types.StringType, stringMapToValueMap(*resp.Tooltip))
		if diags.HasError() {
			return
		}
	}

	if resp.DefaultValue != nil {
		input.DefaultValue = types.StringValue(*resp.DefaultValue)
	}
	if resp.Required != nil {
		input.Required = types.BoolValue(*resp.Required)
	}
	if resp.Readonly != nil {
		input.Readonly = types.BoolValue(*resp.Readonly)
	}
	if resp.PropagateChanges != nil {
		input.PropagateChanges = types.BoolValue(*resp.PropagateChanges)
	}
	if resp.Ephemeral != nil {
		input.Ephemeral = types.BoolValue(*resp.Ephemeral)
	}
	if resp.PropagateChanges != nil {
		input.PropagateChanges = types.BoolValue(*resp.PropagateChanges)
	}
	if resp.Target != nil {
		input.Target = types.StringValue(*resp.Target)
	}

	if resp.Attributes != nil {
		input.Attributes, diags = types.MapValue(types.StringType, stringMapToValueMap(*resp.Attributes))
		if diags.HasError() {
			return
		}
	}

	for i := range resp.Items {
		if len(input.Items) <= i {
			input.Items = append(input.Items, CustomFormInputItemModel{})
		}
		input.Items[i].fromApiResponse(&resp.Items[i])
	}

	if resp.DisplayCondition.IsSet() {
		if input.DisplayCondition == nil {
			input.DisplayCondition = &DisplayConditionModel{}
		}
		diags = input.DisplayCondition.fromApiResponse(resp.DisplayCondition.Get())
	}

	return
}

func (validation *CustomFormValidationModel) toApiRequest() *v1.FormValidation {
	return &v1.FormValidation{
		Id:         validation.ID.ValueString(),
		Expression: validation.Expression.ValueString(),
		Target:     validation.Target.ValueString(),
		Message:    modelMapToStringMap(validation.Message),
	}
}

func (validation *CustomFormValidationModel) fromApiResponse(resp *v1.FormValidation) (diags diag.Diagnostics) {

	message := make(map[string]attr.Value)
	for key, value := range resp.Message {
		message[key] = types.StringValue(value)
	}

	validation.ID = types.StringValue(resp.Id)
	validation.Expression = types.StringValue(resp.Expression)
	validation.Target = types.StringValue(resp.Target)
	validation.Message, diags = types.MapValue(types.StringType, message)

	return
}

func (conditional *DisplayConditionModel) toApiRequest() v1.NullableDisplayCondition {
	return *v1.NewNullableDisplayCondition(&v1.DisplayCondition{
		Source:  conditional.Source.ValueString(),
		Operand: v1.ConditionOperandEnum(conditional.Operand.ValueString()),
		Value:   conditional.Value.ValueString(),
	})
}

func (conditional *DisplayConditionModel) fromApiResponse(resp *v1.DisplayCondition) (diags diag.Diagnostics) {

	if resp == nil {
		return
	}

	conditional.Source = types.StringValue(resp.Source)
	conditional.Operand = types.StringValue(string(resp.Operand))
	conditional.Value = types.StringValue(resp.Value)

	return
}
