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

// PropertyFloorTypeEnum the model 'PropertyFloorTypeEnum'
type PropertyFloorTypeEnum string

// List of PropertyFloorTypeEnum
const (
	PROPERTYFLOORTYPEENUM_LAMINATE PropertyFloorTypeEnum = "LAMINATE"
	PROPERTYFLOORTYPEENUM_TILES    PropertyFloorTypeEnum = "TILES"
	PROPERTYFLOORTYPEENUM_PARQUET  PropertyFloorTypeEnum = "PARQUET"
)

// All allowed values of PropertyFloorTypeEnum enum
var AllowedPropertyFloorTypeEnumEnumValues = []PropertyFloorTypeEnum{
	"LAMINATE",
	"TILES",
	"PARQUET",
}

func (v *PropertyFloorTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PropertyFloorTypeEnum(value)
	for _, existing := range AllowedPropertyFloorTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PropertyFloorTypeEnum", value)
}

// NewPropertyFloorTypeEnumFromValue returns a pointer to a valid PropertyFloorTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPropertyFloorTypeEnumFromValue(v string) (*PropertyFloorTypeEnum, error) {
	ev := PropertyFloorTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PropertyFloorTypeEnum: valid values are %v", v, AllowedPropertyFloorTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PropertyFloorTypeEnum) IsValid() bool {
	for _, existing := range AllowedPropertyFloorTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PropertyFloorTypeEnum value
func (v PropertyFloorTypeEnum) Ptr() *PropertyFloorTypeEnum {
	return &v
}

type NullablePropertyFloorTypeEnum struct {
	value *PropertyFloorTypeEnum
	isSet bool
}

func (v NullablePropertyFloorTypeEnum) Get() *PropertyFloorTypeEnum {
	return v.value
}

func (v *NullablePropertyFloorTypeEnum) Set(val *PropertyFloorTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyFloorTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyFloorTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyFloorTypeEnum(val *PropertyFloorTypeEnum) *NullablePropertyFloorTypeEnum {
	return &NullablePropertyFloorTypeEnum{value: val, isSet: true}
}

func (v NullablePropertyFloorTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyFloorTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
