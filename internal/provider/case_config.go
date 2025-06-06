package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	v1 "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
)

type CaseConfigModel struct {
	StatusConfig     []CaseStatusConfigModel          `tfsdk:"status_config"`
	Reminders        []CaseReminderConfigModel        `tfsdk:"reminders"`
	AutomaticActions []CaseAutomaticActionConfigModel `tfsdk:"automatic_actions"`
}

func CaseConfigModelNestedSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"status_config": schema.ListNestedAttribute{
			Optional:     true,
			NestedObject: caseStatusConfigModelNestedSchema(),
			Description: `The configuration for the status transition of a case. This is 
			used to determine which status changes are allowed by which party at which time.`,
		},
		"reminders": schema.ListNestedAttribute{
			Optional:     true,
			NestedObject: caseReminderConfigModelNestedSchema(),
			Description:  `The configuration for sending out reminders for a case.`,
		},
		"automatic_actions": schema.ListNestedAttribute{
			Optional:     true,
			NestedObject: caseAutomaticStatusChangeConfigModelNestedSchema(),
			Description: `The configuration for automatic status changes of a case. This is 
			used to determine which status changes are allowed by which party at which time.`,
		},
	}
}

func (cfg *CaseConfigModel) toApiRequest() v1.CaseConfig {

	statusConfig := make([]v1.CaseStatusConfig, len(cfg.StatusConfig))
	for i := range cfg.StatusConfig {
		statusConfig[i] = *cfg.StatusConfig[i].toApiRequest()
	}

	reminders := make([]v1.CaseReminderConfig, len(cfg.Reminders))
	for i := range cfg.Reminders {
		reminders[i] = *cfg.Reminders[i].toApiRequest()
	}

	automaticActions := make([]v1.CaseAutomaticActionConfig, len(cfg.AutomaticActions))
	for i := range cfg.AutomaticActions {
		automaticActions[i] = *cfg.AutomaticActions[i].toApiRequest()
	}

	return v1.CaseConfig{
		StatusConfig:     statusConfig,
		Reminders:        reminders,
		AutomaticActions: automaticActions,
	}
}

func (cfg *CaseConfigModel) fromApiResponse(resp v1.CaseConfig) bool {

	isEmpty := true

	if len(resp.StatusConfig) != 0 {
		cfg.StatusConfig = make([]CaseStatusConfigModel, len(resp.StatusConfig))
		isEmpty = false
		for i := range resp.StatusConfig {
			cfg.StatusConfig[i].fromApiResponse(&resp.StatusConfig[i])
		}
	} else {
		cfg.StatusConfig = nil
	}

	if len(resp.Reminders) != 0 {
		cfg.Reminders = make([]CaseReminderConfigModel, len(resp.Reminders))
		isEmpty = false
		for i := range resp.Reminders {
			cfg.Reminders[i].fromApiResponse(&resp.Reminders[i])
		}
	} else {
		cfg.Reminders = nil
	}

	if len(resp.AutomaticActions) != 0 {
		cfg.AutomaticActions = make([]CaseAutomaticActionConfigModel, len(resp.AutomaticActions))
		isEmpty = false
		for i := range resp.AutomaticActions {
			cfg.AutomaticActions[i].fromApiResponse(&resp.AutomaticActions[i])
		}
	} else {
		cfg.AutomaticActions = nil
	}
	isEmpty = true

	return isEmpty
}

// ##################################     CaseStatusConfig 	 		##################################

type CaseStatusConfigModel struct {
	Action       types.String              `tfsdk:"action"`
	SourceStatus types.List                `tfsdk:"source_status"`
	Parties      types.List                `tfsdk:"parties"`
	Feedback     []CaseFeedbackConfigModel `tfsdk:"feedback"`
	Notify       types.List                `tfsdk:"notify"`
}

func caseStatusConfigModelNestedSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"action": schema.StringAttribute{
				Required: true,
				Description: `The action that is allowed to be executed for the below statuses 
				and parties.`,
			},
			"source_status": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
				Description: `A list of possible statuses in which the case must be, so that the 
				configuration applies.`,
			},
			"parties": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
				Description: "The parties that are allowed to execute the action.",
			},
			"feedback": schema.ListNestedAttribute{
				Optional:     true,
				NestedObject: caseFeedbackConfigModelNestedSchema(),
				Description: `Optional list of feedback that is requested from the party 
				executing the action.`,
			},
			"notify": schema.ListAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: `The list of parties that should be notified of the status change.`,
			},
		},
	}
}

func (cfg *CaseStatusConfigModel) toApiRequest() *v1.CaseStatusConfig {

	feedback := make([]v1.CaseFeedbackConfig, len(cfg.Feedback))
	for i := range cfg.Feedback {
		pt := cfg.Feedback[i].toApiRequest()
		if pt == nil {
			continue
		}
		feedback[i] = *pt
	}

	return &v1.CaseStatusConfig{
		Action:       v1.CaseActionEnum(cfg.Action.ValueString()),
		SourceStatus: modelListToEnumSlice[v1.CaseStatusEnum](cfg.SourceStatus),
		Parties:      modelListToEnumSlice[v1.UserTypeEnum](cfg.Parties),
		Feedback:     feedback,
		Notify:       modelListToEnumSlice[v1.UserTypeEnum](cfg.Notify),
	}
}

func (cfg *CaseStatusConfigModel) fromApiResponse(resp *v1.CaseStatusConfig) (diags diag.Diagnostics) {

	cfg.Action = types.StringValue(string(resp.Action))

	if len(resp.SourceStatus) != 0 {
		cfg.SourceStatus, diags = types.ListValue(types.StringType, enumSliceToValueList(resp.SourceStatus))
		if diags.HasError() {
			return
		}
	} else {
		cfg.SourceStatus = basetypes.NewListNull(types.StringType)
	}

	if len(resp.Parties) != 0 {
		cfg.Parties, diags = types.ListValue(types.StringType, enumSliceToValueList(resp.Parties))
		if diags.HasError() {
			return
		}
	} else {
		cfg.Parties = basetypes.NewListNull(types.StringType)
	}

	if len(resp.Feedback) != 0 {
		cfg.Feedback = make([]CaseFeedbackConfigModel, len(resp.Feedback))
		for i := range resp.Feedback {
			cfg.Feedback[i].fromApiResponse(&resp.Feedback[i])
		}
	} else {
		cfg.Feedback = nil
	}

	if len(resp.Notify) != 0 {
		cfg.Notify, diags = types.ListValue(types.StringType, enumSliceToValueList(resp.Notify))
		if diags.HasError() {
			return
		}
	} else {
		cfg.Notify = basetypes.NewListNull(types.StringType)
	}

	return
}

// #################################     CaseFeedbackConfig 	 		#################################

type CaseFeedbackConfigModel struct {
	ID         types.String     `tfsdk:"id"`
	Visibility types.List       `tfsdk:"visibility"`
	Mode       types.String     `tfsdk:"mode"`
	Form       *CustomFormModel `tfsdk:"form"`
}

func caseFeedbackConfigModelNestedSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    true,
				Description: "The ID of the feedback that is requested.",
			},
			"visibility": schema.ListAttribute{
				Required:    true,
				ElementType: types.StringType,
				Description: `The list of parties that should be able to see the feedback.`,
			},
			"mode": schema.StringAttribute{
				Required:    true,
				Description: `The mode in which the feedback should be requested.`,
			},
			"form": schema.SingleNestedAttribute{
				Optional:    true,
				Description: "A custom form to collect data with",
				Attributes:  CustomFormNestedSchema(),
			},
		},
	}
}

func (cfg *CaseFeedbackConfigModel) toApiRequest() *v1.CaseFeedbackConfig {

	ret := &v1.CaseFeedbackConfig{
		Id:         cfg.ID.ValueString(),
		Visibility: modelListToEnumSlice[v1.UserTypeEnum](cfg.Visibility),
		Mode:       v1.CaseFeedbackModeEnum(cfg.Mode.ValueString()),
	}

	ret.Form = v1.NullableCustomForm{}
	if cfg.Form != nil {
		form := cfg.Form.toApiRequest()
		ret.Form.Set(&form)
	}

	return ret
}

func (cfg *CaseFeedbackConfigModel) fromApiResponse(resp *v1.CaseFeedbackConfig) (diags diag.Diagnostics) {

	cfg.ID = types.StringValue(resp.Id)
	cfg.Mode = types.StringValue(string(resp.Mode))

	if len(resp.Visibility) != 0 {
		cfg.Visibility, diags = types.ListValue(types.StringType, enumSliceToValueList[v1.UserTypeEnum](resp.Visibility))
		if diags.HasError() {
			return
		}
	} else {
		cfg.Visibility = basetypes.NewListNull(types.StringType)
	}

	if formDef := resp.Form.Get(); resp.Form.IsSet() && !isEmptyCustomForm(formDef) {
		cfg.Form = &CustomFormModel{}
		cfg.Form.fromApiResponse(formDef)
	} else {
		cfg.Form = nil
	}

	return

}

func isEmptyCustomForm(form *v1.CustomForm) bool {
	return form == nil || (len(form.Layout) == 0 && len(form.Groups) == 0 && len(form.Inputs) == 0 && len(form.Validations) == 0)
}

// #################################     CaseReminderConfig 	 		#################################

type CaseReminderConfigModel struct {
	Status   types.String `tfsdk:"status"`
	Schedule types.List   `tfsdk:"schedule"`
}

func caseReminderConfigModelNestedSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"status": schema.StringAttribute{
				Required:    true,
				Description: "The status in which the case has to be, for the reminder to be sent out.",
			},
			"schedule": schema.ListAttribute{
				Required:    true,
				ElementType: types.Int64Type,
				Description: `The number of days after which a reminder should be sent out.`,
				MarkdownDescription: `The number of days after which a reminder should be sent out. The list is 
        taken as provided and a reminder will be set for the amount of days as set
        in the first element of the list. Once that time has passed, the reminder
        will be rescheduled for the next element in the list. This is repeated until
        the list is empty. The values in the list are considered relative to the
        previous reminder. So [2, 2] will send out a reminder after 2 and 4 days.
        The reminders will, however, not be sent out on weekends. So if the first
        reminder is sent out on a Friday, the second reminder will be sent out on
        Tuesday. All reminders are reset, whenever the case is updated.`,
			},
		},
	}
}

func (cfg *CaseReminderConfigModel) toApiRequest() *v1.CaseReminderConfig {
	return &v1.CaseReminderConfig{
		Status:   v1.CaseStatusEnum(cfg.Status.ValueString()),
		Schedule: modelListToInt32Slice(cfg.Schedule),
	}
}

func (cfg *CaseReminderConfigModel) fromApiResponse(resp *v1.CaseReminderConfig) (diags diag.Diagnostics) {

	cfg.Status = types.StringValue(string(resp.Status))

	if len(resp.Schedule) != 0 {
		cfg.Schedule, diags = types.ListValue(types.Int64Type, int32SliceToValueList(resp.Schedule))
		if diags.HasError() {
			return
		}
	} else {
		cfg.Schedule = basetypes.NewListNull(types.Int64Type)
	}

	return
}

// ###########################     CaseAutomaticStatusChangeConfig 	 		###########################

type CaseAutomaticActionConfigModel struct {
	Status   types.String `tfsdk:"status"`
	Schedule types.Int64  `tfsdk:"schedule"`
	Action   types.String `tfsdk:"action"`
}

func caseAutomaticStatusChangeConfigModelNestedSchema() schema.NestedAttributeObject {
	return schema.NestedAttributeObject{
		Attributes: map[string]schema.Attribute{
			"status": schema.StringAttribute{
				Required:    true,
				Description: "The status in which the case has to be, for the automatic action to be triggered.",
			},
			"schedule": schema.Int64Attribute{
				Required: true,
				Description: `The number of days after which the status change should be triggered. The 
        schedule is reset, whenever the case is updated. 0 means that the action
        should be triggered immediately.`,
			},
			"action": schema.StringAttribute{
				Required: true,
				Description: `JS script of the action that should be executed, if the schedule and execution
        check are successful.`,
			},
		},
	}
}

func (cfg *CaseAutomaticActionConfigModel) toApiRequest() *v1.CaseAutomaticActionConfig {

	return &v1.CaseAutomaticActionConfig{
		Status:   v1.CaseStatusEnum(cfg.Status.ValueString()),
		Schedule: int32(cfg.Schedule.ValueInt64()),
		Action:   cfg.Action.ValueString(),
	}
}

func (cfg *CaseAutomaticActionConfigModel) fromApiResponse(resp *v1.CaseAutomaticActionConfig) (diags diag.Diagnostics) {

	cfg.Status = types.StringValue(string(resp.Status))
	cfg.Schedule = types.Int64Value(int64(resp.Schedule))
	cfg.Action = types.StringValue(string(resp.Action))

	return
}

// ################################################################################################
