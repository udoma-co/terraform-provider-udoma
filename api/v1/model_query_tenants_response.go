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

// checks if the QueryTenantsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryTenantsResponse{}

// QueryTenantsResponse Result of a tenants query
type QueryTenantsResponse struct {
	// The maximum number of entities to return from the query
	Limit *int32 `json:"limit,omitempty"`
	// The number of entities to skip before returning the result
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
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTenantsResponse) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryTenantsResponse) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
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
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTenantsResponse) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *QueryTenantsResponse) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
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
	if o == nil || IsNil(o.Tenants) {
		var ret []Tenant
		return ret
	}
	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryTenantsResponse) GetTenantsOk() ([]Tenant, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *QueryTenantsResponse) HasTenants() bool {
	if o != nil && !IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given []Tenant and assigns it to the Tenants field.
func (o *QueryTenantsResponse) SetTenants(v []Tenant) {
	o.Tenants = v
}

func (o QueryTenantsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryTenantsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Tenants) {
		toSerialize["tenants"] = o.Tenants
	}
	return toSerialize, nil
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
