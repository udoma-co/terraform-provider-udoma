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

// TenancyDurationEnum The duration of the tenancy. This is used to determine which fields are required for the tenancy contract.
type TenancyDurationEnum string

// List of TenancyDurationEnum
const (
	TENANCYDURATIONENUM_OPEN_ENDED TenancyDurationEnum = "OPEN_ENDED"
	TENANCYDURATIONENUM_FIXED_TERM TenancyDurationEnum = "FIXED_TERM"
)

// All allowed values of TenancyDurationEnum enum
var AllowedTenancyDurationEnumEnumValues = []TenancyDurationEnum{
	"OPEN_ENDED",
	"FIXED_TERM",
}

func (v *TenancyDurationEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TenancyDurationEnum(value)
	for _, existing := range AllowedTenancyDurationEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TenancyDurationEnum", value)
}

// NewTenancyDurationEnumFromValue returns a pointer to a valid TenancyDurationEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTenancyDurationEnumFromValue(v string) (*TenancyDurationEnum, error) {
	ev := TenancyDurationEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TenancyDurationEnum: valid values are %v", v, AllowedTenancyDurationEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TenancyDurationEnum) IsValid() bool {
	for _, existing := range AllowedTenancyDurationEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TenancyDurationEnum value
func (v TenancyDurationEnum) Ptr() *TenancyDurationEnum {
	return &v
}

type NullableTenancyDurationEnum struct {
	value *TenancyDurationEnum
	isSet bool
}

func (v NullableTenancyDurationEnum) Get() *TenancyDurationEnum {
	return v.value
}

func (v *NullableTenancyDurationEnum) Set(val *TenancyDurationEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableTenancyDurationEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableTenancyDurationEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenancyDurationEnum(val *TenancyDurationEnum) *NullableTenancyDurationEnum {
	return &NullableTenancyDurationEnum{value: val, isSet: true}
}

func (v NullableTenancyDurationEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenancyDurationEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
