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

// checks if the QueryVersionMigratorsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryVersionMigratorsRequest{}

// QueryVersionMigratorsRequest a request for querying version migrators
type QueryVersionMigratorsRequest struct {
	// The maximum number of entities to return from the query
	Limit *int32 `json:"limit,omitempty"`
	// The number of entities to skip before returning the result
	Offset *int32 `json:"offset,omitempty"`
}

// NewQueryVersionMigratorsRequest instantiates a new QueryVersionMigratorsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryVersionMigratorsRequest() *QueryVersionMigratorsRequest {
	this := QueryVersionMigratorsRequest{}
	return &this
}

// NewQueryVersionMigratorsRequestWithDefaults instantiates a new QueryVersionMigratorsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryVersionMigratorsRequestWithDefaults() *QueryVersionMigratorsRequest {
	this := QueryVersionMigratorsRequest{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *QueryVersionMigratorsRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryVersionMigratorsRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryVersionMigratorsRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *QueryVersionMigratorsRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *QueryVersionMigratorsRequest) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryVersionMigratorsRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *QueryVersionMigratorsRequest) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *QueryVersionMigratorsRequest) SetOffset(v int32) {
	o.Offset = &v
}

func (o QueryVersionMigratorsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryVersionMigratorsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	return toSerialize, nil
}

type NullableQueryVersionMigratorsRequest struct {
	value *QueryVersionMigratorsRequest
	isSet bool
}

func (v NullableQueryVersionMigratorsRequest) Get() *QueryVersionMigratorsRequest {
	return v.value
}

func (v *NullableQueryVersionMigratorsRequest) Set(val *QueryVersionMigratorsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryVersionMigratorsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryVersionMigratorsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryVersionMigratorsRequest(val *QueryVersionMigratorsRequest) *NullableQueryVersionMigratorsRequest {
	return &NullableQueryVersionMigratorsRequest{value: val, isSet: true}
}

func (v NullableQueryVersionMigratorsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryVersionMigratorsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
