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

// RentUpdateMask the model 'RentUpdateMask'
type RentUpdateMask int64

// List of RentUpdateMask
const (
	RENTUPDATEMASK_ADDITIONAL_COSTS RentUpdateMask = 1
	RENTUPDATEMASK_RENT             RentUpdateMask = 2
	RENTUPDATEMASK_INCIDENTALS      RentUpdateMask = 4
	RENTUPDATEMASK_INDEX_MONTH      RentUpdateMask = 8
	RENTUPDATEMASK_INDEX_POINTS     RentUpdateMask = 16
)

// All allowed values of RentUpdateMask enum
var AllowedRentUpdateMaskEnumValues = []RentUpdateMask{
	1,
	2,
	4,
	8,
	16,
}

func (v *RentUpdateMask) UnmarshalJSON(src []byte) error {
	var value int64
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RentUpdateMask(value)
	for _, existing := range AllowedRentUpdateMaskEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RentUpdateMask", value)
}

// NewRentUpdateMaskFromValue returns a pointer to a valid RentUpdateMask
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRentUpdateMaskFromValue(v int64) (*RentUpdateMask, error) {
	ev := RentUpdateMask(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RentUpdateMask: valid values are %v", v, AllowedRentUpdateMaskEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RentUpdateMask) IsValid() bool {
	for _, existing := range AllowedRentUpdateMaskEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RentUpdateMask value
func (v RentUpdateMask) Ptr() *RentUpdateMask {
	return &v
}

type NullableRentUpdateMask struct {
	value *RentUpdateMask
	isSet bool
}

func (v NullableRentUpdateMask) Get() *RentUpdateMask {
	return v.value
}

func (v *NullableRentUpdateMask) Set(val *RentUpdateMask) {
	v.value = val
	v.isSet = true
}

func (v NullableRentUpdateMask) IsSet() bool {
	return v.isSet
}

func (v *NullableRentUpdateMask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRentUpdateMask(val *RentUpdateMask) *NullableRentUpdateMask {
	return &NullableRentUpdateMask{value: val, isSet: true}
}

func (v NullableRentUpdateMask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRentUpdateMask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
