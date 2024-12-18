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

// FormSourceType the model 'FormSourceType'
type FormSourceType string

// List of FormSourceType
const (
	FORMSOURCETYPE_APPOINTMENT         FormSourceType = "appointment"
	FORMSOURCETYPE_CASE                FormSourceType = "case"
	FORMSOURCETYPE_CASE_FEEDBACK       FormSourceType = "case_feedback"
	FORMSOURCETYPE_DOCUMENT_GENERATION FormSourceType = "document_generation"
	FORMSOURCETYPE_HANDOVER            FormSourceType = "handover"
	FORMSOURCETYPE_REPORT              FormSourceType = "report"
	FORMSOURCETYPE_WORKFLOW            FormSourceType = "workflow"
)

// All allowed values of FormSourceType enum
var AllowedFormSourceTypeEnumValues = []FormSourceType{
	"appointment",
	"case",
	"case_feedback",
	"document_generation",
	"handover",
	"report",
	"workflow",
}

func (v *FormSourceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FormSourceType(value)
	for _, existing := range AllowedFormSourceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FormSourceType", value)
}

// NewFormSourceTypeFromValue returns a pointer to a valid FormSourceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFormSourceTypeFromValue(v string) (*FormSourceType, error) {
	ev := FormSourceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FormSourceType: valid values are %v", v, AllowedFormSourceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FormSourceType) IsValid() bool {
	for _, existing := range AllowedFormSourceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FormSourceType value
func (v FormSourceType) Ptr() *FormSourceType {
	return &v
}

type NullableFormSourceType struct {
	value *FormSourceType
	isSet bool
}

func (v NullableFormSourceType) Get() *FormSourceType {
	return v.value
}

func (v *NullableFormSourceType) Set(val *FormSourceType) {
	v.value = val
	v.isSet = true
}

func (v NullableFormSourceType) IsSet() bool {
	return v.isSet
}

func (v *NullableFormSourceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormSourceType(val *FormSourceType) *NullableFormSourceType {
	return &NullableFormSourceType{value: val, isSet: true}
}

func (v NullableFormSourceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormSourceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
