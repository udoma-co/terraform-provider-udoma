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

// QueryTenantsRequest Search criteria for querying tenants
type QueryTenantsRequest struct {
	Limit  *int32 `json:"limit,omitempty"`
	Offset *int32 `json:"offset,omitempty"`
}

// NewQueryTenantsRequest instantiates a new QueryTenantsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryTenantsRequest() *QueryTenantsRequest {
	this := QueryTenantsRequest{}
	return &this
}

// NewQueryTenantsRequestWithDefaults instantiates a new QueryTenantsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryTenantsRequestWithDefaults() *QueryTenantsRequest {
	this := QueryTenantsRequest{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *QueryTenantsRequest) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTenantsRequest) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryTenantsRequest) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *QueryTenantsRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *QueryTenantsRequest) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTenantsRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *QueryTenantsRequest) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *QueryTenantsRequest) SetOffset(v int32) {
	o.Offset = &v
}

func (o QueryTenantsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	return json.Marshal(toSerialize)
}

type NullableQueryTenantsRequest struct {
	value *QueryTenantsRequest
	isSet bool
}

func (v NullableQueryTenantsRequest) Get() *QueryTenantsRequest {
	return v.value
}

func (v *NullableQueryTenantsRequest) Set(val *QueryTenantsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryTenantsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryTenantsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryTenantsRequest(val *QueryTenantsRequest) *NullableQueryTenantsRequest {
	return &NullableQueryTenantsRequest{value: val, isSet: true}
}

func (v NullableQueryTenantsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryTenantsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
