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

// checks if the QueryApprovalsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryApprovalsResponse{}

// QueryApprovalsResponse Contains the details of a response to a query for approvals.
type QueryApprovalsResponse struct {
	Approvals []Approval `json:"approvals,omitempty"`
}

// NewQueryApprovalsResponse instantiates a new QueryApprovalsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryApprovalsResponse() *QueryApprovalsResponse {
	this := QueryApprovalsResponse{}
	return &this
}

// NewQueryApprovalsResponseWithDefaults instantiates a new QueryApprovalsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryApprovalsResponseWithDefaults() *QueryApprovalsResponse {
	this := QueryApprovalsResponse{}
	return &this
}

// GetApprovals returns the Approvals field value if set, zero value otherwise.
func (o *QueryApprovalsResponse) GetApprovals() []Approval {
	if o == nil || IsNil(o.Approvals) {
		var ret []Approval
		return ret
	}
	return o.Approvals
}

// GetApprovalsOk returns a tuple with the Approvals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryApprovalsResponse) GetApprovalsOk() ([]Approval, bool) {
	if o == nil || IsNil(o.Approvals) {
		return nil, false
	}
	return o.Approvals, true
}

// HasApprovals returns a boolean if a field has been set.
func (o *QueryApprovalsResponse) HasApprovals() bool {
	if o != nil && !IsNil(o.Approvals) {
		return true
	}

	return false
}

// SetApprovals gets a reference to the given []Approval and assigns it to the Approvals field.
func (o *QueryApprovalsResponse) SetApprovals(v []Approval) {
	o.Approvals = v
}

func (o QueryApprovalsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryApprovalsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Approvals) {
		toSerialize["approvals"] = o.Approvals
	}
	return toSerialize, nil
}

type NullableQueryApprovalsResponse struct {
	value *QueryApprovalsResponse
	isSet bool
}

func (v NullableQueryApprovalsResponse) Get() *QueryApprovalsResponse {
	return v.value
}

func (v *NullableQueryApprovalsResponse) Set(val *QueryApprovalsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryApprovalsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryApprovalsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryApprovalsResponse(val *QueryApprovalsResponse) *NullableQueryApprovalsResponse {
	return &NullableQueryApprovalsResponse{value: val, isSet: true}
}

func (v NullableQueryApprovalsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryApprovalsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
