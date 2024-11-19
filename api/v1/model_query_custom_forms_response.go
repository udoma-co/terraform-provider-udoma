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

// checks if the QueryCustomFormsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryCustomFormsResponse{}

// QueryCustomFormsResponse Result of a custom form query
type QueryCustomFormsResponse struct {
	CustomForms []Form `json:"custom_forms,omitempty"`
}

// NewQueryCustomFormsResponse instantiates a new QueryCustomFormsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCustomFormsResponse() *QueryCustomFormsResponse {
	this := QueryCustomFormsResponse{}
	return &this
}

// NewQueryCustomFormsResponseWithDefaults instantiates a new QueryCustomFormsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCustomFormsResponseWithDefaults() *QueryCustomFormsResponse {
	this := QueryCustomFormsResponse{}
	return &this
}

// GetCustomForms returns the CustomForms field value if set, zero value otherwise.
func (o *QueryCustomFormsResponse) GetCustomForms() []Form {
	if o == nil || IsNil(o.CustomForms) {
		var ret []Form
		return ret
	}
	return o.CustomForms
}

// GetCustomFormsOk returns a tuple with the CustomForms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryCustomFormsResponse) GetCustomFormsOk() ([]Form, bool) {
	if o == nil || IsNil(o.CustomForms) {
		return nil, false
	}
	return o.CustomForms, true
}

// HasCustomForms returns a boolean if a field has been set.
func (o *QueryCustomFormsResponse) HasCustomForms() bool {
	if o != nil && !IsNil(o.CustomForms) {
		return true
	}

	return false
}

// SetCustomForms gets a reference to the given []Form and assigns it to the CustomForms field.
func (o *QueryCustomFormsResponse) SetCustomForms(v []Form) {
	o.CustomForms = v
}

func (o QueryCustomFormsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryCustomFormsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomForms) {
		toSerialize["custom_forms"] = o.CustomForms
	}
	return toSerialize, nil
}

type NullableQueryCustomFormsResponse struct {
	value *QueryCustomFormsResponse
	isSet bool
}

func (v NullableQueryCustomFormsResponse) Get() *QueryCustomFormsResponse {
	return v.value
}

func (v *NullableQueryCustomFormsResponse) Set(val *QueryCustomFormsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCustomFormsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCustomFormsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCustomFormsResponse(val *QueryCustomFormsResponse) *NullableQueryCustomFormsResponse {
	return &NullableQueryCustomFormsResponse{value: val, isSet: true}
}

func (v NullableQueryCustomFormsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCustomFormsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}