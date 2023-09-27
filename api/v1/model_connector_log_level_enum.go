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

// ConnectorLogLevelEnum Log level to be used by the connector. If not provided, the connector will use the default log level
type ConnectorLogLevelEnum string

// List of ConnectorLogLevelEnum
const (
	CONNECTORLOGLEVELENUM_DEBUG   ConnectorLogLevelEnum = "DEBUG"
	CONNECTORLOGLEVELENUM_INFO    ConnectorLogLevelEnum = "INFO"
	CONNECTORLOGLEVELENUM_WARNING ConnectorLogLevelEnum = "WARNING"
	CONNECTORLOGLEVELENUM_ERROR   ConnectorLogLevelEnum = "ERROR"
)

// All allowed values of ConnectorLogLevelEnum enum
var AllowedConnectorLogLevelEnumEnumValues = []ConnectorLogLevelEnum{
	"DEBUG",
	"INFO",
	"WARNING",
	"ERROR",
}

func (v *ConnectorLogLevelEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConnectorLogLevelEnum(value)
	for _, existing := range AllowedConnectorLogLevelEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConnectorLogLevelEnum", value)
}

// NewConnectorLogLevelEnumFromValue returns a pointer to a valid ConnectorLogLevelEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConnectorLogLevelEnumFromValue(v string) (*ConnectorLogLevelEnum, error) {
	ev := ConnectorLogLevelEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConnectorLogLevelEnum: valid values are %v", v, AllowedConnectorLogLevelEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConnectorLogLevelEnum) IsValid() bool {
	for _, existing := range AllowedConnectorLogLevelEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConnectorLogLevelEnum value
func (v ConnectorLogLevelEnum) Ptr() *ConnectorLogLevelEnum {
	return &v
}

type NullableConnectorLogLevelEnum struct {
	value *ConnectorLogLevelEnum
	isSet bool
}

func (v NullableConnectorLogLevelEnum) Get() *ConnectorLogLevelEnum {
	return v.value
}

func (v *NullableConnectorLogLevelEnum) Set(val *ConnectorLogLevelEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorLogLevelEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorLogLevelEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorLogLevelEnum(val *ConnectorLogLevelEnum) *NullableConnectorLogLevelEnum {
	return &NullableConnectorLogLevelEnum{value: val, isSet: true}
}

func (v NullableConnectorLogLevelEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorLogLevelEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
