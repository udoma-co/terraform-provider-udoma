package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	v1 "gitlab.com/zestlabs-io/udoma/terraform-provider-udoma/api/v1"
)

func TestBankAccountFromAPIPreservesEmptyCategories(t *testing.T) {
	model := BankAccountModel{
		Categories: types.ListValueMust(types.StringType, nil),
	}

	err := model.fromAPI(&v1.BankAccount{
		Id:            "bank-account-id",
		AccountHolder: "holder",
		Iban:          "DE29100100100987654321",
		CreatedAt:     1,
		UpdatedAt:     1,
		Cadence:       v1.BALANCECADENCEENUM_MONTHLY_CADENCE,
	})
	if err != nil {
		t.Fatalf("fromAPI returned error: %v", err)
	}

	if model.Categories.IsNull() {
		t.Fatal("expected categories to remain an empty list, got null")
	}
	if len(model.Categories.Elements()) != 0 {
		t.Fatalf("expected empty categories, got %d elements", len(model.Categories.Elements()))
	}
}

func TestAppointmentTemplateFromAPIPreservesEmptyConfirmationReminders(t *testing.T) {
	template := AppointmentTemplateModel{
		ConfirmationReminders: types.ListValueMust(types.Int64Type, nil),
		Inputs:                &CustomFormModel{},
	}

	resp := &v1.AppointmentTemplate{
		Id:   "appointment-template-id",
		Name: "template",
		Form: *v1.NewNullableCustomForm(&v1.CustomForm{}),
	}

	if diags := template.fromApiResponse(resp); diags.HasError() {
		t.Fatalf("fromApiResponse returned diagnostics: %v", diags.Errors())
	}

	if template.ConfirmationReminders.IsNull() {
		t.Fatal("expected confirmation_reminders to remain an empty list, got null")
	}
	if len(template.ConfirmationReminders.Elements()) != 0 {
		t.Fatalf("expected empty confirmation_reminders, got %d elements", len(template.ConfirmationReminders.Elements()))
	}
}

func TestCaseConfigFromAPIPreservesEmptyConfiguredLists(t *testing.T) {
	cfg := CaseConfigModel{
		StatusConfig:     make([]CaseStatusConfigModel, 0),
		Reminders:        make([]CaseReminderConfigModel, 0),
		AutomaticActions: make([]CaseAutomaticActionConfigModel, 0),
	}

	cfg.fromApiResponse(v1.CaseConfig{})

	if cfg.StatusConfig == nil {
		t.Fatal("expected status_config to remain an empty list, got nil")
	}
	if cfg.Reminders == nil {
		t.Fatal("expected reminders to remain an empty list, got nil")
	}
	if cfg.AutomaticActions == nil {
		t.Fatal("expected automatic_actions to remain an empty list, got nil")
	}
}

func TestCaseStatusConfigFromAPIPreservesEmptyConfiguredLists(t *testing.T) {
	cfg := CaseStatusConfigModel{
		SourceStatus: types.ListValueMust(types.StringType, nil),
		Parties:      types.ListValueMust(types.StringType, nil),
		Notify:       types.ListValueMust(types.StringType, nil),
		Feedback:     make([]CaseFeedbackConfigModel, 0),
	}

	resp := &v1.CaseStatusConfig{Action: v1.CASEACTIONENUM_OPEN}
	if diags := cfg.fromApiResponse(resp); diags.HasError() {
		t.Fatalf("fromApiResponse returned diagnostics: %v", diags.Errors())
	}

	if cfg.SourceStatus.IsNull() {
		t.Fatal("expected source_status to remain an empty list, got null")
	}
	if cfg.Parties.IsNull() {
		t.Fatal("expected parties to remain an empty list, got null")
	}
	if cfg.Notify.IsNull() {
		t.Fatal("expected notify to remain an empty list, got null")
	}
	if cfg.Feedback == nil {
		t.Fatal("expected feedback to remain an empty list, got nil")
	}
}

func TestCaseFeedbackConfigFromAPIPreservesEmptyVisibility(t *testing.T) {
	cfg := CaseFeedbackConfigModel{
		Visibility: types.ListValueMust(types.StringType, nil),
	}

	resp := &v1.CaseFeedbackConfig{
		Id:   "feedback-id",
		Mode: v1.CASEFEEDBACKMODEENUM_OPTIONAL_COMMENT,
	}
	if diags := cfg.fromApiResponse(resp); diags.HasError() {
		t.Fatalf("fromApiResponse returned diagnostics: %v", diags.Errors())
	}

	if cfg.Visibility.IsNull() {
		t.Fatal("expected visibility to remain an empty list, got null")
	}
	if len(cfg.Visibility.Elements()) != 0 {
		t.Fatalf("expected empty visibility, got %d elements", len(cfg.Visibility.Elements()))
	}
}

func TestCaseReminderConfigFromAPIPreservesEmptySchedule(t *testing.T) {
	cfg := CaseReminderConfigModel{
		Schedule: types.ListValueMust(types.Int64Type, nil),
	}

	resp := &v1.CaseReminderConfig{Status: v1.CASESTATUSENUM_OPEN}
	if diags := cfg.fromApiResponse(resp); diags.HasError() {
		t.Fatalf("fromApiResponse returned diagnostics: %v", diags.Errors())
	}

	if cfg.Schedule.IsNull() {
		t.Fatal("expected schedule to remain an empty list, got null")
	}
	if len(cfg.Schedule.Elements()) != 0 {
		t.Fatalf("expected empty schedule, got %d elements", len(cfg.Schedule.Elements()))
	}
}

func TestCustomFormGroupFromAPIPreservesEmptySubtitle(t *testing.T) {
	group := CustomFormGroupModel{
		Subtitle: types.ListValueMust(types.StringType, nil),
	}

	resp := &v1.FormGroup{
		Id:    "group-id",
		Type:  v1.FORMGROUPTYPE_ROW,
		Items: []v1.FormItem{},
	}
	if diags := group.fromApiResponse(resp); diags.HasError() {
		t.Fatalf("fromApiResponse returned diagnostics: %v", diags.Errors())
	}

	if group.Subtitle.IsNull() {
		t.Fatal("expected subtitle to remain an empty list, got null")
	}
	if len(group.Subtitle.Elements()) != 0 {
		t.Fatalf("expected empty subtitle, got %d elements", len(group.Subtitle.Elements()))
	}
}
