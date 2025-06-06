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

// SignerNotificationTypeEnum The type of the notificaiton sent to the signer. If \"YOUSIGN\", then the notification is sent by yousign. If \"UDOMA\", then we send the notification.
type SignerNotificationTypeEnum string

// List of SignerNotificationTypeEnum
const (
	SIGNERNOTIFICATIONTYPEENUM_UDOMA   SignerNotificationTypeEnum = "udoma"
	SIGNERNOTIFICATIONTYPEENUM_YOUSIGN SignerNotificationTypeEnum = "yousign"
)

// All allowed values of SignerNotificationTypeEnum enum
var AllowedSignerNotificationTypeEnumEnumValues = []SignerNotificationTypeEnum{
	"udoma",
	"yousign",
}

func (v *SignerNotificationTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SignerNotificationTypeEnum(value)
	for _, existing := range AllowedSignerNotificationTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SignerNotificationTypeEnum", value)
}

// NewSignerNotificationTypeEnumFromValue returns a pointer to a valid SignerNotificationTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSignerNotificationTypeEnumFromValue(v string) (*SignerNotificationTypeEnum, error) {
	ev := SignerNotificationTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SignerNotificationTypeEnum: valid values are %v", v, AllowedSignerNotificationTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SignerNotificationTypeEnum) IsValid() bool {
	for _, existing := range AllowedSignerNotificationTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SignerNotificationTypeEnum value
func (v SignerNotificationTypeEnum) Ptr() *SignerNotificationTypeEnum {
	return &v
}

type NullableSignerNotificationTypeEnum struct {
	value *SignerNotificationTypeEnum
	isSet bool
}

func (v NullableSignerNotificationTypeEnum) Get() *SignerNotificationTypeEnum {
	return v.value
}

func (v *NullableSignerNotificationTypeEnum) Set(val *SignerNotificationTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSignerNotificationTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSignerNotificationTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignerNotificationTypeEnum(val *SignerNotificationTypeEnum) *NullableSignerNotificationTypeEnum {
	return &NullableSignerNotificationTypeEnum{value: val, isSet: true}
}

func (v NullableSignerNotificationTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignerNotificationTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
