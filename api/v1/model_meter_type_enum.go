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

// MeterTypeEnum the model 'MeterTypeEnum'
type MeterTypeEnum string

// List of MeterTypeEnum
const (
	METERTYPEENUM_OTHER       MeterTypeEnum = "OTHER"
	METERTYPEENUM_WARM_WATER  MeterTypeEnum = "WARM_WATER"
	METERTYPEENUM_COLD_WATER  MeterTypeEnum = "COLD_WATER"
	METERTYPEENUM_ELECTRICITY MeterTypeEnum = "ELECTRICITY"
	METERTYPEENUM_GAS         MeterTypeEnum = "GAS"
	METERTYPEENUM_HEATER      MeterTypeEnum = "HEATER"
)

// All allowed values of MeterTypeEnum enum
var AllowedMeterTypeEnumEnumValues = []MeterTypeEnum{
	"OTHER",
	"WARM_WATER",
	"COLD_WATER",
	"ELECTRICITY",
	"GAS",
	"HEATER",
}

func (v *MeterTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MeterTypeEnum(value)
	for _, existing := range AllowedMeterTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MeterTypeEnum", value)
}

// NewMeterTypeEnumFromValue returns a pointer to a valid MeterTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMeterTypeEnumFromValue(v string) (*MeterTypeEnum, error) {
	ev := MeterTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MeterTypeEnum: valid values are %v", v, AllowedMeterTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MeterTypeEnum) IsValid() bool {
	for _, existing := range AllowedMeterTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MeterTypeEnum value
func (v MeterTypeEnum) Ptr() *MeterTypeEnum {
	return &v
}

type NullableMeterTypeEnum struct {
	value *MeterTypeEnum
	isSet bool
}

func (v NullableMeterTypeEnum) Get() *MeterTypeEnum {
	return v.value
}

func (v *NullableMeterTypeEnum) Set(val *MeterTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableMeterTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableMeterTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeterTypeEnum(val *MeterTypeEnum) *NullableMeterTypeEnum {
	return &NullableMeterTypeEnum{value: val, isSet: true}
}

func (v NullableMeterTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeterTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
