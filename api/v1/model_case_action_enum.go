/*
Udoma

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"fmt"
)

// CaseActionEnum The list of actions that can be performed on a case. Each action has a predetermined target status that is going to be set once the action is performed.
type CaseActionEnum string

// List of CaseActionEnum
const (
	CASEACTIONENUM_ACCEPT_WORK      CaseActionEnum = "ACCEPT_WORK"
	CASEACTIONENUM_ACKNOWLEDGE      CaseActionEnum = "ACKNOWLEDGE"
	CASEACTIONENUM_ARCHIVE          CaseActionEnum = "ARCHIVE"
	CASEACTIONENUM_ASSIGN           CaseActionEnum = "ASSIGN"
	CASEACTIONENUM_CANCEL           CaseActionEnum = "CANCEL"
	CASEACTIONENUM_CANCEL_WORK      CaseActionEnum = "CANCEL_WORK"
	CASEACTIONENUM_COMMENT          CaseActionEnum = "COMMENT"
	CASEACTIONENUM_FINISH_WORK      CaseActionEnum = "FINISH_WORK"
	CASEACTIONENUM_SEND_REMINDER    CaseActionEnum = "SEND_REMINDER"
	CASEACTIONENUM_OPEN             CaseActionEnum = "OPEN"
	CASEACTIONENUM_PAUSE_WORK       CaseActionEnum = "PAUSE_WORK"
	CASEACTIONENUM_REJECT           CaseActionEnum = "REJECT"
	CASEACTIONENUM_REJECT_WORK      CaseActionEnum = "REJECT_WORK"
	CASEACTIONENUM_REOPEN           CaseActionEnum = "REOPEN"
	CASEACTIONENUM_REQUEST_FEEDBACK CaseActionEnum = "REQUEST_FEEDBACK"
	CASEACTIONENUM_REQUEST_WORK     CaseActionEnum = "REQUEST_WORK"
	CASEACTIONENUM_RESCHEDULE_WORK  CaseActionEnum = "RESCHEDULE_WORK"
	CASEACTIONENUM_RESOLVE          CaseActionEnum = "RESOLVE"
	CASEACTIONENUM_SCHEDULE_WORK    CaseActionEnum = "SCHEDULE_WORK"
	CASEACTIONENUM_START_WORK       CaseActionEnum = "START_WORK"
	CASEACTIONENUM_UNASSIGN         CaseActionEnum = "UNASSIGN"
	CASEACTIONENUM_UPDATE           CaseActionEnum = "UPDATE"
)

// All allowed values of CaseActionEnum enum
var AllowedCaseActionEnumEnumValues = []CaseActionEnum{
	"ACCEPT_WORK",
	"ACKNOWLEDGE",
	"ARCHIVE",
	"ASSIGN",
	"CANCEL",
	"CANCEL_WORK",
	"COMMENT",
	"FINISH_WORK",
	"SEND_REMINDER",
	"OPEN",
	"PAUSE_WORK",
	"REJECT",
	"REJECT_WORK",
	"REOPEN",
	"REQUEST_FEEDBACK",
	"REQUEST_WORK",
	"RESCHEDULE_WORK",
	"RESOLVE",
	"SCHEDULE_WORK",
	"START_WORK",
	"UNASSIGN",
	"UPDATE",
}

func (v *CaseActionEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CaseActionEnum(value)
	for _, existing := range AllowedCaseActionEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CaseActionEnum", value)
}

// NewCaseActionEnumFromValue returns a pointer to a valid CaseActionEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCaseActionEnumFromValue(v string) (*CaseActionEnum, error) {
	ev := CaseActionEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CaseActionEnum: valid values are %v", v, AllowedCaseActionEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CaseActionEnum) IsValid() bool {
	for _, existing := range AllowedCaseActionEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseActionEnum value
func (v CaseActionEnum) Ptr() *CaseActionEnum {
	return &v
}

type NullableCaseActionEnum struct {
	value *CaseActionEnum
	isSet bool
}

func (v NullableCaseActionEnum) Get() *CaseActionEnum {
	return v.value
}

func (v *NullableCaseActionEnum) Set(val *CaseActionEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableCaseActionEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableCaseActionEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaseActionEnum(val *CaseActionEnum) *NullableCaseActionEnum {
	return &NullableCaseActionEnum{value: val, isSet: true}
}

func (v NullableCaseActionEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaseActionEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
