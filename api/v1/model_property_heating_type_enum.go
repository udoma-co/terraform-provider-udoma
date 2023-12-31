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

// PropertyHeatingTypeEnum the model 'PropertyHeatingTypeEnum'
type PropertyHeatingTypeEnum string

// List of PropertyHeatingTypeEnum
const (
	PROPERTYHEATINGTYPEENUM_CENTRAL_HEATING    PropertyHeatingTypeEnum = "CENTRAL_HEATING"
	PROPERTYHEATINGTYPEENUM_FLOOR_HEATING      PropertyHeatingTypeEnum = "FLOOR_HEATING"
	PROPERTYHEATINGTYPEENUM_GAS_HEATING        PropertyHeatingTypeEnum = "GAS_HEATING"
	PROPERTYHEATINGTYPEENUM_ELECTRICAL_HEATING PropertyHeatingTypeEnum = "ELECTRICAL_HEATING"
	PROPERTYHEATINGTYPEENUM_HEAT_PUMP          PropertyHeatingTypeEnum = "HEAT_PUMP"
	PROPERTYHEATINGTYPEENUM_AC                 PropertyHeatingTypeEnum = "AC"
)

// All allowed values of PropertyHeatingTypeEnum enum
var AllowedPropertyHeatingTypeEnumEnumValues = []PropertyHeatingTypeEnum{
	"CENTRAL_HEATING",
	"FLOOR_HEATING",
	"GAS_HEATING",
	"ELECTRICAL_HEATING",
	"HEAT_PUMP",
	"AC",
}

func (v *PropertyHeatingTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PropertyHeatingTypeEnum(value)
	for _, existing := range AllowedPropertyHeatingTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PropertyHeatingTypeEnum", value)
}

// NewPropertyHeatingTypeEnumFromValue returns a pointer to a valid PropertyHeatingTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPropertyHeatingTypeEnumFromValue(v string) (*PropertyHeatingTypeEnum, error) {
	ev := PropertyHeatingTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PropertyHeatingTypeEnum: valid values are %v", v, AllowedPropertyHeatingTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PropertyHeatingTypeEnum) IsValid() bool {
	for _, existing := range AllowedPropertyHeatingTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PropertyHeatingTypeEnum value
func (v PropertyHeatingTypeEnum) Ptr() *PropertyHeatingTypeEnum {
	return &v
}

type NullablePropertyHeatingTypeEnum struct {
	value *PropertyHeatingTypeEnum
	isSet bool
}

func (v NullablePropertyHeatingTypeEnum) Get() *PropertyHeatingTypeEnum {
	return v.value
}

func (v *NullablePropertyHeatingTypeEnum) Set(val *PropertyHeatingTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyHeatingTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyHeatingTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyHeatingTypeEnum(val *PropertyHeatingTypeEnum) *NullablePropertyHeatingTypeEnum {
	return &NullablePropertyHeatingTypeEnum{value: val, isSet: true}
}

func (v NullablePropertyHeatingTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyHeatingTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
