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

// QueryApprovalType the model 'QueryApprovalType'
type QueryApprovalType string

// List of QueryApprovalType
const (
	QUERYAPPROVALTYPE_ALL      QueryApprovalType = "ALL"
	QUERYAPPROVALTYPE_REJECTED QueryApprovalType = "REJECTED"
	QUERYAPPROVALTYPE_PENDING  QueryApprovalType = "PENDING"
	QUERYAPPROVALTYPE_ARCHIVED QueryApprovalType = "ARCHIVED"
)

// All allowed values of QueryApprovalType enum
var AllowedQueryApprovalTypeEnumValues = []QueryApprovalType{
	"ALL",
	"REJECTED",
	"PENDING",
	"ARCHIVED",
}

func (v *QueryApprovalType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QueryApprovalType(value)
	for _, existing := range AllowedQueryApprovalTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QueryApprovalType", value)
}

// NewQueryApprovalTypeFromValue returns a pointer to a valid QueryApprovalType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQueryApprovalTypeFromValue(v string) (*QueryApprovalType, error) {
	ev := QueryApprovalType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QueryApprovalType: valid values are %v", v, AllowedQueryApprovalTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QueryApprovalType) IsValid() bool {
	for _, existing := range AllowedQueryApprovalTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QueryApprovalType value
func (v QueryApprovalType) Ptr() *QueryApprovalType {
	return &v
}

type NullableQueryApprovalType struct {
	value *QueryApprovalType
	isSet bool
}

func (v NullableQueryApprovalType) Get() *QueryApprovalType {
	return v.value
}

func (v *NullableQueryApprovalType) Set(val *QueryApprovalType) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryApprovalType) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryApprovalType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryApprovalType(val *QueryApprovalType) *NullableQueryApprovalType {
	return &NullableQueryApprovalType{value: val, isSet: true}
}

func (v NullableQueryApprovalType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryApprovalType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
