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

// checks if the GetNumberOfCasesForActionAndTimePeriodRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNumberOfCasesForActionAndTimePeriodRequest{}

// GetNumberOfCasesForActionAndTimePeriodRequest struct for GetNumberOfCasesForActionAndTimePeriodRequest
type GetNumberOfCasesForActionAndTimePeriodRequest struct {
	Action   *CaseActionEnum `json:"action,omitempty"`
	FromTime *int64          `json:"from_time,omitempty"`
	TillTime *int64          `json:"till_time,omitempty"`
}

// NewGetNumberOfCasesForActionAndTimePeriodRequest instantiates a new GetNumberOfCasesForActionAndTimePeriodRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNumberOfCasesForActionAndTimePeriodRequest() *GetNumberOfCasesForActionAndTimePeriodRequest {
	this := GetNumberOfCasesForActionAndTimePeriodRequest{}
	return &this
}

// NewGetNumberOfCasesForActionAndTimePeriodRequestWithDefaults instantiates a new GetNumberOfCasesForActionAndTimePeriodRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNumberOfCasesForActionAndTimePeriodRequestWithDefaults() *GetNumberOfCasesForActionAndTimePeriodRequest {
	this := GetNumberOfCasesForActionAndTimePeriodRequest{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *GetNumberOfCasesForActionAndTimePeriodRequest) GetAction() CaseActionEnum {
	if o == nil || IsNil(o.Action) {
		var ret CaseActionEnum
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNumberOfCasesForActionAndTimePeriodRequest) GetActionOk() (*CaseActionEnum, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *GetNumberOfCasesForActionAndTimePeriodRequest) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given CaseActionEnum and assigns it to the Action field.
func (o *GetNumberOfCasesForActionAndTimePeriodRequest) SetAction(v CaseActionEnum) {
	o.Action = &v
}

// GetFromTime returns the FromTime field value if set, zero value otherwise.
func (o *GetNumberOfCasesForActionAndTimePeriodRequest) GetFromTime() int64 {
	if o == nil || IsNil(o.FromTime) {
		var ret int64
		return ret
	}
	return *o.FromTime
}

// GetFromTimeOk returns a tuple with the FromTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNumberOfCasesForActionAndTimePeriodRequest) GetFromTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.FromTime) {
		return nil, false
	}
	return o.FromTime, true
}

// HasFromTime returns a boolean if a field has been set.
func (o *GetNumberOfCasesForActionAndTimePeriodRequest) HasFromTime() bool {
	if o != nil && !IsNil(o.FromTime) {
		return true
	}

	return false
}

// SetFromTime gets a reference to the given int64 and assigns it to the FromTime field.
func (o *GetNumberOfCasesForActionAndTimePeriodRequest) SetFromTime(v int64) {
	o.FromTime = &v
}

// GetTillTime returns the TillTime field value if set, zero value otherwise.
func (o *GetNumberOfCasesForActionAndTimePeriodRequest) GetTillTime() int64 {
	if o == nil || IsNil(o.TillTime) {
		var ret int64
		return ret
	}
	return *o.TillTime
}

// GetTillTimeOk returns a tuple with the TillTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNumberOfCasesForActionAndTimePeriodRequest) GetTillTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.TillTime) {
		return nil, false
	}
	return o.TillTime, true
}

// HasTillTime returns a boolean if a field has been set.
func (o *GetNumberOfCasesForActionAndTimePeriodRequest) HasTillTime() bool {
	if o != nil && !IsNil(o.TillTime) {
		return true
	}

	return false
}

// SetTillTime gets a reference to the given int64 and assigns it to the TillTime field.
func (o *GetNumberOfCasesForActionAndTimePeriodRequest) SetTillTime(v int64) {
	o.TillTime = &v
}

func (o GetNumberOfCasesForActionAndTimePeriodRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNumberOfCasesForActionAndTimePeriodRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.FromTime) {
		toSerialize["from_time"] = o.FromTime
	}
	if !IsNil(o.TillTime) {
		toSerialize["till_time"] = o.TillTime
	}
	return toSerialize, nil
}

type NullableGetNumberOfCasesForActionAndTimePeriodRequest struct {
	value *GetNumberOfCasesForActionAndTimePeriodRequest
	isSet bool
}

func (v NullableGetNumberOfCasesForActionAndTimePeriodRequest) Get() *GetNumberOfCasesForActionAndTimePeriodRequest {
	return v.value
}

func (v *NullableGetNumberOfCasesForActionAndTimePeriodRequest) Set(val *GetNumberOfCasesForActionAndTimePeriodRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNumberOfCasesForActionAndTimePeriodRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNumberOfCasesForActionAndTimePeriodRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNumberOfCasesForActionAndTimePeriodRequest(val *GetNumberOfCasesForActionAndTimePeriodRequest) *NullableGetNumberOfCasesForActionAndTimePeriodRequest {
	return &NullableGetNumberOfCasesForActionAndTimePeriodRequest{value: val, isSet: true}
}

func (v NullableGetNumberOfCasesForActionAndTimePeriodRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNumberOfCasesForActionAndTimePeriodRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
