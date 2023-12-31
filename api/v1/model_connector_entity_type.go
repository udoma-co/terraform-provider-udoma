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

// ConnectorEntityType The type of entity that can be synchronised. This is used to indicate what data to be synchronised by the connector
type ConnectorEntityType string

// List of ConnectorEntityType
const (
	CONNECTORENTITYTYPE_BANK_ACCOUNT     ConnectorEntityType = "BANK_ACCOUNT"
	CONNECTORENTITYTYPE_TENANT           ConnectorEntityType = "TENANT"
	CONNECTORENTITYTYPE_TENANCY          ConnectorEntityType = "TENANCY"
	CONNECTORENTITYTYPE_OWNER            ConnectorEntityType = "OWNER"
	CONNECTORENTITYTYPE_PROPERTY         ConnectorEntityType = "PROPERTY"
	CONNECTORENTITYTYPE_SERVICE_PROVIDER ConnectorEntityType = "SERVICE_PROVIDER"
)

// All allowed values of ConnectorEntityType enum
var AllowedConnectorEntityTypeEnumValues = []ConnectorEntityType{
	"BANK_ACCOUNT",
	"TENANT",
	"TENANCY",
	"OWNER",
	"PROPERTY",
	"SERVICE_PROVIDER",
}

func (v *ConnectorEntityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConnectorEntityType(value)
	for _, existing := range AllowedConnectorEntityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConnectorEntityType", value)
}

// NewConnectorEntityTypeFromValue returns a pointer to a valid ConnectorEntityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConnectorEntityTypeFromValue(v string) (*ConnectorEntityType, error) {
	ev := ConnectorEntityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConnectorEntityType: valid values are %v", v, AllowedConnectorEntityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConnectorEntityType) IsValid() bool {
	for _, existing := range AllowedConnectorEntityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConnectorEntityType value
func (v ConnectorEntityType) Ptr() *ConnectorEntityType {
	return &v
}

type NullableConnectorEntityType struct {
	value *ConnectorEntityType
	isSet bool
}

func (v NullableConnectorEntityType) Get() *ConnectorEntityType {
	return v.value
}

func (v *NullableConnectorEntityType) Set(val *ConnectorEntityType) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorEntityType) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorEntityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorEntityType(val *ConnectorEntityType) *NullableConnectorEntityType {
	return &NullableConnectorEntityType{value: val, isSet: true}
}

func (v NullableConnectorEntityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorEntityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
