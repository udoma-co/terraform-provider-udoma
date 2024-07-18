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

// checks if the GetNumberOfInteractionsForCasesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNumberOfInteractionsForCasesResponse{}

// GetNumberOfInteractionsForCasesResponse struct for GetNumberOfInteractionsForCasesResponse
type GetNumberOfInteractionsForCasesResponse struct {
	NumberOfInteractions *int64 `json:"number_of_interactions,omitempty"`
}

// NewGetNumberOfInteractionsForCasesResponse instantiates a new GetNumberOfInteractionsForCasesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNumberOfInteractionsForCasesResponse() *GetNumberOfInteractionsForCasesResponse {
	this := GetNumberOfInteractionsForCasesResponse{}
	return &this
}

// NewGetNumberOfInteractionsForCasesResponseWithDefaults instantiates a new GetNumberOfInteractionsForCasesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNumberOfInteractionsForCasesResponseWithDefaults() *GetNumberOfInteractionsForCasesResponse {
	this := GetNumberOfInteractionsForCasesResponse{}
	return &this
}

// GetNumberOfInteractions returns the NumberOfInteractions field value if set, zero value otherwise.
func (o *GetNumberOfInteractionsForCasesResponse) GetNumberOfInteractions() int64 {
	if o == nil || IsNil(o.NumberOfInteractions) {
		var ret int64
		return ret
	}
	return *o.NumberOfInteractions
}

// GetNumberOfInteractionsOk returns a tuple with the NumberOfInteractions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNumberOfInteractionsForCasesResponse) GetNumberOfInteractionsOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfInteractions) {
		return nil, false
	}
	return o.NumberOfInteractions, true
}

// HasNumberOfInteractions returns a boolean if a field has been set.
func (o *GetNumberOfInteractionsForCasesResponse) HasNumberOfInteractions() bool {
	if o != nil && !IsNil(o.NumberOfInteractions) {
		return true
	}

	return false
}

// SetNumberOfInteractions gets a reference to the given int64 and assigns it to the NumberOfInteractions field.
func (o *GetNumberOfInteractionsForCasesResponse) SetNumberOfInteractions(v int64) {
	o.NumberOfInteractions = &v
}

func (o GetNumberOfInteractionsForCasesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNumberOfInteractionsForCasesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumberOfInteractions) {
		toSerialize["number_of_interactions"] = o.NumberOfInteractions
	}
	return toSerialize, nil
}

type NullableGetNumberOfInteractionsForCasesResponse struct {
	value *GetNumberOfInteractionsForCasesResponse
	isSet bool
}

func (v NullableGetNumberOfInteractionsForCasesResponse) Get() *GetNumberOfInteractionsForCasesResponse {
	return v.value
}

func (v *NullableGetNumberOfInteractionsForCasesResponse) Set(val *GetNumberOfInteractionsForCasesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNumberOfInteractionsForCasesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNumberOfInteractionsForCasesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNumberOfInteractionsForCasesResponse(val *GetNumberOfInteractionsForCasesResponse) *NullableGetNumberOfInteractionsForCasesResponse {
	return &NullableGetNumberOfInteractionsForCasesResponse{value: val, isSet: true}
}

func (v NullableGetNumberOfInteractionsForCasesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNumberOfInteractionsForCasesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
