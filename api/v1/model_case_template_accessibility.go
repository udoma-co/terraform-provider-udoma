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

// CaseTemplateAccessibility The list of options that can be used to define where a case template can be used in the webapp.
type CaseTemplateAccessibility string

// List of CaseTemplateAccessibility
const (
	CASETEMPLATEACCESSIBILITY_PROPERTY_MANAGER CaseTemplateAccessibility = "PROPERTY_MANAGER"
	CASETEMPLATEACCESSIBILITY_TENANT           CaseTemplateAccessibility = "TENANT"
	CASETEMPLATEACCESSIBILITY_PUBLIC           CaseTemplateAccessibility = "PUBLIC"
)

// All allowed values of CaseTemplateAccessibility enum
var AllowedCaseTemplateAccessibilityEnumValues = []CaseTemplateAccessibility{
	"PROPERTY_MANAGER",
	"TENANT",
	"PUBLIC",
}

func (v *CaseTemplateAccessibility) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CaseTemplateAccessibility(value)
	for _, existing := range AllowedCaseTemplateAccessibilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CaseTemplateAccessibility", value)
}

// NewCaseTemplateAccessibilityFromValue returns a pointer to a valid CaseTemplateAccessibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCaseTemplateAccessibilityFromValue(v string) (*CaseTemplateAccessibility, error) {
	ev := CaseTemplateAccessibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CaseTemplateAccessibility: valid values are %v", v, AllowedCaseTemplateAccessibilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CaseTemplateAccessibility) IsValid() bool {
	for _, existing := range AllowedCaseTemplateAccessibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CaseTemplateAccessibility value
func (v CaseTemplateAccessibility) Ptr() *CaseTemplateAccessibility {
	return &v
}

type NullableCaseTemplateAccessibility struct {
	value *CaseTemplateAccessibility
	isSet bool
}

func (v NullableCaseTemplateAccessibility) Get() *CaseTemplateAccessibility {
	return v.value
}

func (v *NullableCaseTemplateAccessibility) Set(val *CaseTemplateAccessibility) {
	v.value = val
	v.isSet = true
}

func (v NullableCaseTemplateAccessibility) IsSet() bool {
	return v.isSet
}

func (v *NullableCaseTemplateAccessibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaseTemplateAccessibility(val *CaseTemplateAccessibility) *NullableCaseTemplateAccessibility {
	return &NullableCaseTemplateAccessibility{value: val, isSet: true}
}

func (v NullableCaseTemplateAccessibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaseTemplateAccessibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
