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

// UnssignCaseRequest Request used to remove a service provider from a case
type UnssignCaseRequest struct {
	Party *CaseParty `json:"party,omitempty"`
}

// NewUnssignCaseRequest instantiates a new UnssignCaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnssignCaseRequest() *UnssignCaseRequest {
	this := UnssignCaseRequest{}
	return &this
}

// NewUnssignCaseRequestWithDefaults instantiates a new UnssignCaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnssignCaseRequestWithDefaults() *UnssignCaseRequest {
	this := UnssignCaseRequest{}
	return &this
}

// GetParty returns the Party field value if set, zero value otherwise.
func (o *UnssignCaseRequest) GetParty() CaseParty {
	if o == nil || o.Party == nil {
		var ret CaseParty
		return ret
	}
	return *o.Party
}

// GetPartyOk returns a tuple with the Party field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnssignCaseRequest) GetPartyOk() (*CaseParty, bool) {
	if o == nil || o.Party == nil {
		return nil, false
	}
	return o.Party, true
}

// HasParty returns a boolean if a field has been set.
func (o *UnssignCaseRequest) HasParty() bool {
	if o != nil && o.Party != nil {
		return true
	}

	return false
}

// SetParty gets a reference to the given CaseParty and assigns it to the Party field.
func (o *UnssignCaseRequest) SetParty(v CaseParty) {
	o.Party = &v
}

func (o UnssignCaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Party != nil {
		toSerialize["party"] = o.Party
	}
	return json.Marshal(toSerialize)
}

type NullableUnssignCaseRequest struct {
	value *UnssignCaseRequest
	isSet bool
}

func (v NullableUnssignCaseRequest) Get() *UnssignCaseRequest {
	return v.value
}

func (v *NullableUnssignCaseRequest) Set(val *UnssignCaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnssignCaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnssignCaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnssignCaseRequest(val *UnssignCaseRequest) *NullableUnssignCaseRequest {
	return &NullableUnssignCaseRequest{value: val, isSet: true}
}

func (v NullableUnssignCaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnssignCaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}