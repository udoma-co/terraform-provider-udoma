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

// checks if the GetExternalUserCasesCountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetExternalUserCasesCountResponse{}

// GetExternalUserCasesCountResponse struct for GetExternalUserCasesCountResponse
type GetExternalUserCasesCountResponse struct {
	NumberOfCases *int64 `json:"numberOfCases,omitempty"`
}

// NewGetExternalUserCasesCountResponse instantiates a new GetExternalUserCasesCountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetExternalUserCasesCountResponse() *GetExternalUserCasesCountResponse {
	this := GetExternalUserCasesCountResponse{}
	return &this
}

// NewGetExternalUserCasesCountResponseWithDefaults instantiates a new GetExternalUserCasesCountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetExternalUserCasesCountResponseWithDefaults() *GetExternalUserCasesCountResponse {
	this := GetExternalUserCasesCountResponse{}
	return &this
}

// GetNumberOfCases returns the NumberOfCases field value if set, zero value otherwise.
func (o *GetExternalUserCasesCountResponse) GetNumberOfCases() int64 {
	if o == nil || IsNil(o.NumberOfCases) {
		var ret int64
		return ret
	}
	return *o.NumberOfCases
}

// GetNumberOfCasesOk returns a tuple with the NumberOfCases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetExternalUserCasesCountResponse) GetNumberOfCasesOk() (*int64, bool) {
	if o == nil || IsNil(o.NumberOfCases) {
		return nil, false
	}
	return o.NumberOfCases, true
}

// HasNumberOfCases returns a boolean if a field has been set.
func (o *GetExternalUserCasesCountResponse) HasNumberOfCases() bool {
	if o != nil && !IsNil(o.NumberOfCases) {
		return true
	}

	return false
}

// SetNumberOfCases gets a reference to the given int64 and assigns it to the NumberOfCases field.
func (o *GetExternalUserCasesCountResponse) SetNumberOfCases(v int64) {
	o.NumberOfCases = &v
}

func (o GetExternalUserCasesCountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetExternalUserCasesCountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumberOfCases) {
		toSerialize["numberOfCases"] = o.NumberOfCases
	}
	return toSerialize, nil
}

type NullableGetExternalUserCasesCountResponse struct {
	value *GetExternalUserCasesCountResponse
	isSet bool
}

func (v NullableGetExternalUserCasesCountResponse) Get() *GetExternalUserCasesCountResponse {
	return v.value
}

func (v *NullableGetExternalUserCasesCountResponse) Set(val *GetExternalUserCasesCountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetExternalUserCasesCountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetExternalUserCasesCountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetExternalUserCasesCountResponse(val *GetExternalUserCasesCountResponse) *NullableGetExternalUserCasesCountResponse {
	return &NullableGetExternalUserCasesCountResponse{value: val, isSet: true}
}

func (v NullableGetExternalUserCasesCountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetExternalUserCasesCountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
