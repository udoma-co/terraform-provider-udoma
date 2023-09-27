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

// QueryTenantsResponse Result of a tenants query
type QueryTenantsResponse struct {
	Limit   *int32   `json:"limit,omitempty"`
	Offset  *int32   `json:"offset,omitempty"`
	Tenants []Tenant `json:"tenants,omitempty"`
}

// NewQueryTenantsResponse instantiates a new QueryTenantsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryTenantsResponse() *QueryTenantsResponse {
	this := QueryTenantsResponse{}
	return &this
}

// NewQueryTenantsResponseWithDefaults instantiates a new QueryTenantsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryTenantsResponseWithDefaults() *QueryTenantsResponse {
	this := QueryTenantsResponse{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *QueryTenantsResponse) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTenantsResponse) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryTenantsResponse) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *QueryTenantsResponse) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *QueryTenantsResponse) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTenantsResponse) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *QueryTenantsResponse) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *QueryTenantsResponse) SetOffset(v int32) {
	o.Offset = &v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *QueryTenantsResponse) GetTenants() []Tenant {
	if o == nil || o.Tenants == nil {
		var ret []Tenant
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTenantsResponse) GetTenantsOk() ([]Tenant, bool) {
	if o == nil || o.Tenants == nil {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *QueryTenantsResponse) HasTenants() bool {
	if o != nil && o.Tenants != nil {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []Tenant and assigns it to the Tenants field.
func (o *QueryTenantsResponse) SetTenants(v []Tenant) {
	o.Tenants = v
}

func (o QueryTenantsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Tenants != nil {
		toSerialize["tenants"] = o.Tenants
	}
	return json.Marshal(toSerialize)
}

type NullableQueryTenantsResponse struct {
	value *QueryTenantsResponse
	isSet bool
}

func (v NullableQueryTenantsResponse) Get() *QueryTenantsResponse {
	return v.value
}

func (v *NullableQueryTenantsResponse) Set(val *QueryTenantsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryTenantsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryTenantsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryTenantsResponse(val *QueryTenantsResponse) *NullableQueryTenantsResponse {
	return &NullableQueryTenantsResponse{value: val, isSet: true}
}

func (v NullableQueryTenantsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryTenantsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
