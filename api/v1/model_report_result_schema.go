/*
Udoma

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// ReportResultSchema struct for ReportResultSchema
type ReportResultSchema struct {
	// Whether the result is a list of items
	IsList     *bool                        `json:"is_list,omitempty"`
	Properties []ReportResultSchemaProperty `json:"properties,omitempty"`
}

// NewReportResultSchema instantiates a new ReportResultSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportResultSchema() *ReportResultSchema {
	this := ReportResultSchema{}
	return &this
}

// NewReportResultSchemaWithDefaults instantiates a new ReportResultSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportResultSchemaWithDefaults() *ReportResultSchema {
	this := ReportResultSchema{}
	return &this
}

// GetIsList returns the IsList field value if set, zero value otherwise.
func (o *ReportResultSchema) GetIsList() bool {
	if o == nil || o.IsList == nil {
		var ret bool
		return ret
	}
	return *o.IsList
}

// GetIsListOk returns a tuple with the IsList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResultSchema) GetIsListOk() (*bool, bool) {
	if o == nil || o.IsList == nil {
		return nil, false
	}
	return o.IsList, true
}

// HasIsList returns a boolean if a field has been set.
func (o *ReportResultSchema) HasIsList() bool {
	if o != nil && o.IsList != nil {
		return true
	}

	return false
}

// SetIsList gets a reference to the given bool and assigns it to the IsList field.
func (o *ReportResultSchema) SetIsList(v bool) {
	o.IsList = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *ReportResultSchema) GetProperties() []ReportResultSchemaProperty {
	if o == nil || o.Properties == nil {
		var ret []ReportResultSchemaProperty
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResultSchema) GetPropertiesOk() ([]ReportResultSchemaProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *ReportResultSchema) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []ReportResultSchemaProperty and assigns it to the Properties field.
func (o *ReportResultSchema) SetProperties(v []ReportResultSchemaProperty) {
	o.Properties = v
}

func (o ReportResultSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsList != nil {
		toSerialize["is_list"] = o.IsList
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableReportResultSchema struct {
	value *ReportResultSchema
	isSet bool
}

func (v NullableReportResultSchema) Get() *ReportResultSchema {
	return v.value
}

func (v *NullableReportResultSchema) Set(val *ReportResultSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableReportResultSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableReportResultSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportResultSchema(val *ReportResultSchema) *NullableReportResultSchema {
	return &NullableReportResultSchema{value: val, isSet: true}
}

func (v NullableReportResultSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportResultSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
