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

// checks if the QueryWorkflowDefinitionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryWorkflowDefinitionsRequest{}

// QueryWorkflowDefinitionsRequest a request for querying workflow definitions
type QueryWorkflowDefinitionsRequest struct {
	Limit *int32 `json:"limit,omitempty"`
}

// NewQueryWorkflowDefinitionsRequest instantiates a new QueryWorkflowDefinitionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryWorkflowDefinitionsRequest() *QueryWorkflowDefinitionsRequest {
	this := QueryWorkflowDefinitionsRequest{}
	return &this
}

// NewQueryWorkflowDefinitionsRequestWithDefaults instantiates a new QueryWorkflowDefinitionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryWorkflowDefinitionsRequestWithDefaults() *QueryWorkflowDefinitionsRequest {
	this := QueryWorkflowDefinitionsRequest{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *QueryWorkflowDefinitionsRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryWorkflowDefinitionsRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryWorkflowDefinitionsRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *QueryWorkflowDefinitionsRequest) SetLimit(v int32) {
	o.Limit = &v
}

func (o QueryWorkflowDefinitionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryWorkflowDefinitionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	return toSerialize, nil
}

type NullableQueryWorkflowDefinitionsRequest struct {
	value *QueryWorkflowDefinitionsRequest
	isSet bool
}

func (v NullableQueryWorkflowDefinitionsRequest) Get() *QueryWorkflowDefinitionsRequest {
	return v.value
}

func (v *NullableQueryWorkflowDefinitionsRequest) Set(val *QueryWorkflowDefinitionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryWorkflowDefinitionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryWorkflowDefinitionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryWorkflowDefinitionsRequest(val *QueryWorkflowDefinitionsRequest) *NullableQueryWorkflowDefinitionsRequest {
	return &NullableQueryWorkflowDefinitionsRequest{value: val, isSet: true}
}

func (v NullableQueryWorkflowDefinitionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryWorkflowDefinitionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
