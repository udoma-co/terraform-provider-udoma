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

// checks if the CreateCaseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCaseResponse{}

// CreateCaseResponse The data generated based on a new case requests
type CreateCaseResponse struct {
	Case       *Case   `json:"case,omitempty"`
	AccessCode *string `json:"access_code,omitempty"`
}

// NewCreateCaseResponse instantiates a new CreateCaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCaseResponse() *CreateCaseResponse {
	this := CreateCaseResponse{}
	return &this
}

// NewCreateCaseResponseWithDefaults instantiates a new CreateCaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCaseResponseWithDefaults() *CreateCaseResponse {
	this := CreateCaseResponse{}
	return &this
}

// GetCase returns the Case field value if set, zero value otherwise.
func (o *CreateCaseResponse) GetCase() Case {
	if o == nil || IsNil(o.Case) {
		var ret Case
		return ret
	}
	return *o.Case
}

// GetCaseOk returns a tuple with the Case field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCaseResponse) GetCaseOk() (*Case, bool) {
	if o == nil || IsNil(o.Case) {
		return nil, false
	}
	return o.Case, true
}

// HasCase returns a boolean if a field has been set.
func (o *CreateCaseResponse) HasCase() bool {
	if o != nil && !IsNil(o.Case) {
		return true
	}

	return false
}

// SetCase gets a reference to the given Case and assigns it to the Case field.
func (o *CreateCaseResponse) SetCase(v Case) {
	o.Case = &v
}

// GetAccessCode returns the AccessCode field value if set, zero value otherwise.
func (o *CreateCaseResponse) GetAccessCode() string {
	if o == nil || IsNil(o.AccessCode) {
		var ret string
		return ret
	}
	return *o.AccessCode
}

// GetAccessCodeOk returns a tuple with the AccessCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCaseResponse) GetAccessCodeOk() (*string, bool) {
	if o == nil || IsNil(o.AccessCode) {
		return nil, false
	}
	return o.AccessCode, true
}

// HasAccessCode returns a boolean if a field has been set.
func (o *CreateCaseResponse) HasAccessCode() bool {
	if o != nil && !IsNil(o.AccessCode) {
		return true
	}

	return false
}

// SetAccessCode gets a reference to the given string and assigns it to the AccessCode field.
func (o *CreateCaseResponse) SetAccessCode(v string) {
	o.AccessCode = &v
}

func (o CreateCaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCaseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Case) {
		toSerialize["case"] = o.Case
	}
	if !IsNil(o.AccessCode) {
		toSerialize["access_code"] = o.AccessCode
	}
	return toSerialize, nil
}

type NullableCreateCaseResponse struct {
	value *CreateCaseResponse
	isSet bool
}

func (v NullableCreateCaseResponse) Get() *CreateCaseResponse {
	return v.value
}

func (v *NullableCreateCaseResponse) Set(val *CreateCaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCaseResponse(val *CreateCaseResponse) *NullableCreateCaseResponse {
	return &NullableCreateCaseResponse{value: val, isSet: true}
}

func (v NullableCreateCaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
