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

// checks if the QueryPropertiesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryPropertiesResponse{}

// QueryPropertiesResponse A list of all properties that are requested by the user
type QueryPropertiesResponse struct {
	// list of properties that match the query request
	Properties []Property `json:"properties,omitempty"`
	// total number of properties available
	TotalCount *int64 `json:"total_count,omitempty"`
	// number of items that were skipped in the result
	Offset *int32 `json:"offset,omitempty"`
}

// NewQueryPropertiesResponse instantiates a new QueryPropertiesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPropertiesResponse() *QueryPropertiesResponse {
	this := QueryPropertiesResponse{}
	return &this
}

// NewQueryPropertiesResponseWithDefaults instantiates a new QueryPropertiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPropertiesResponseWithDefaults() *QueryPropertiesResponse {
	this := QueryPropertiesResponse{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *QueryPropertiesResponse) GetProperties() []Property {
	if o == nil || IsNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPropertiesResponse) GetPropertiesOk() ([]Property, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *QueryPropertiesResponse) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *QueryPropertiesResponse) SetProperties(v []Property) {
	o.Properties = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *QueryPropertiesResponse) GetTotalCount() int64 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPropertiesResponse) GetTotalCountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *QueryPropertiesResponse) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *QueryPropertiesResponse) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *QueryPropertiesResponse) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPropertiesResponse) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *QueryPropertiesResponse) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *QueryPropertiesResponse) SetOffset(v int32) {
	o.Offset = &v
}

func (o QueryPropertiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPropertiesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.TotalCount) {
		toSerialize["total_count"] = o.TotalCount
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	return toSerialize, nil
}

type NullableQueryPropertiesResponse struct {
	value *QueryPropertiesResponse
	isSet bool
}

func (v NullableQueryPropertiesResponse) Get() *QueryPropertiesResponse {
	return v.value
}

func (v *NullableQueryPropertiesResponse) Set(val *QueryPropertiesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPropertiesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPropertiesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPropertiesResponse(val *QueryPropertiesResponse) *NullableQueryPropertiesResponse {
	return &NullableQueryPropertiesResponse{value: val, isSet: true}
}

func (v NullableQueryPropertiesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPropertiesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
