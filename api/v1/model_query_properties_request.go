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

// checks if the QueryPropertiesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryPropertiesRequest{}

// QueryPropertiesRequest Request used to get a list of properties for the curent user
type QueryPropertiesRequest struct {
	// The maximum number of entities to return from the query
	Limit *int32 `json:"limit,omitempty"`
	// The number of entities to skip before returning the result
	Offset *int32 `json:"offset,omitempty"`
	// like search for the name of the property
	Name *string `json:"name,omitempty"`
	// optional property ID, if set only properties that are children of the given property are returned
	ParentRef *string `json:"parent_ref,omitempty"`
	// optional property owner ID, if set only properties that are owned by the given owner are returned
	OwnerRef *string `json:"owner_ref,omitempty"`
}

// NewQueryPropertiesRequest instantiates a new QueryPropertiesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryPropertiesRequest() *QueryPropertiesRequest {
	this := QueryPropertiesRequest{}
	return &this
}

// NewQueryPropertiesRequestWithDefaults instantiates a new QueryPropertiesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryPropertiesRequestWithDefaults() *QueryPropertiesRequest {
	this := QueryPropertiesRequest{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *QueryPropertiesRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPropertiesRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryPropertiesRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *QueryPropertiesRequest) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *QueryPropertiesRequest) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPropertiesRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *QueryPropertiesRequest) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *QueryPropertiesRequest) SetOffset(v int32) {
	o.Offset = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *QueryPropertiesRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPropertiesRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *QueryPropertiesRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *QueryPropertiesRequest) SetName(v string) {
	o.Name = &v
}

// GetParentRef returns the ParentRef field value if set, zero value otherwise.
func (o *QueryPropertiesRequest) GetParentRef() string {
	if o == nil || IsNil(o.ParentRef) {
		var ret string
		return ret
	}
	return *o.ParentRef
}

// GetParentRefOk returns a tuple with the ParentRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPropertiesRequest) GetParentRefOk() (*string, bool) {
	if o == nil || IsNil(o.ParentRef) {
		return nil, false
	}
	return o.ParentRef, true
}

// HasParentRef returns a boolean if a field has been set.
func (o *QueryPropertiesRequest) HasParentRef() bool {
	if o != nil && !IsNil(o.ParentRef) {
		return true
	}

	return false
}

// SetParentRef gets a reference to the given string and assigns it to the ParentRef field.
func (o *QueryPropertiesRequest) SetParentRef(v string) {
	o.ParentRef = &v
}

// GetOwnerRef returns the OwnerRef field value if set, zero value otherwise.
func (o *QueryPropertiesRequest) GetOwnerRef() string {
	if o == nil || IsNil(o.OwnerRef) {
		var ret string
		return ret
	}
	return *o.OwnerRef
}

// GetOwnerRefOk returns a tuple with the OwnerRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryPropertiesRequest) GetOwnerRefOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerRef) {
		return nil, false
	}
	return o.OwnerRef, true
}

// HasOwnerRef returns a boolean if a field has been set.
func (o *QueryPropertiesRequest) HasOwnerRef() bool {
	if o != nil && !IsNil(o.OwnerRef) {
		return true
	}

	return false
}

// SetOwnerRef gets a reference to the given string and assigns it to the OwnerRef field.
func (o *QueryPropertiesRequest) SetOwnerRef(v string) {
	o.OwnerRef = &v
}

func (o QueryPropertiesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryPropertiesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ParentRef) {
		toSerialize["parent_ref"] = o.ParentRef
	}
	if !IsNil(o.OwnerRef) {
		toSerialize["owner_ref"] = o.OwnerRef
	}
	return toSerialize, nil
}

type NullableQueryPropertiesRequest struct {
	value *QueryPropertiesRequest
	isSet bool
}

func (v NullableQueryPropertiesRequest) Get() *QueryPropertiesRequest {
	return v.value
}

func (v *NullableQueryPropertiesRequest) Set(val *QueryPropertiesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryPropertiesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryPropertiesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryPropertiesRequest(val *QueryPropertiesRequest) *NullableQueryPropertiesRequest {
	return &NullableQueryPropertiesRequest{value: val, isSet: true}
}

func (v NullableQueryPropertiesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryPropertiesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
