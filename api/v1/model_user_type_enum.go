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

// UserTypeEnum the model 'UserTypeEnum'
type UserTypeEnum string

// List of UserTypeEnum
const (
	USERTYPEENUM_ANY              UserTypeEnum = "ANY"
	USERTYPEENUM_TENANT           UserTypeEnum = "TENANT"
	USERTYPEENUM_MANAGER          UserTypeEnum = "MANAGER"
	USERTYPEENUM_SERVICE_PROVIDER UserTypeEnum = "SERVICE_PROVIDER"
	USERTYPEENUM_AUTOMATION       UserTypeEnum = "AUTOMATION"
)

// All allowed values of UserTypeEnum enum
var AllowedUserTypeEnumEnumValues = []UserTypeEnum{
	"ANY",
	"TENANT",
	"MANAGER",
	"SERVICE_PROVIDER",
	"AUTOMATION",
}

func (v *UserTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UserTypeEnum(value)
	for _, existing := range AllowedUserTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UserTypeEnum", value)
}

// NewUserTypeEnumFromValue returns a pointer to a valid UserTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUserTypeEnumFromValue(v string) (*UserTypeEnum, error) {
	ev := UserTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UserTypeEnum: valid values are %v", v, AllowedUserTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UserTypeEnum) IsValid() bool {
	for _, existing := range AllowedUserTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UserTypeEnum value
func (v UserTypeEnum) Ptr() *UserTypeEnum {
	return &v
}

type NullableUserTypeEnum struct {
	value *UserTypeEnum
	isSet bool
}

func (v NullableUserTypeEnum) Get() *UserTypeEnum {
	return v.value
}

func (v *NullableUserTypeEnum) Set(val *UserTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableUserTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableUserTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserTypeEnum(val *UserTypeEnum) *NullableUserTypeEnum {
	return &NullableUserTypeEnum{value: val, isSet: true}
}

func (v NullableUserTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
