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

// CaseAutomaticActionConfig Defines the configuration for an automatic action a case.
type CaseAutomaticActionConfig struct {
	Status *CaseStatusEnum `json:"status,omitempty"`
	// The number of days after which the status change should be triggered. The  schedule is reset, whenever the case is updated. 0 means that the action should be triggered immediately.
	Schedule *int32 `json:"schedule,omitempty"`
	// JSON script of the action that should be executed, if the schedule and execution check are successful.
	Action *string `json:"action,omitempty"`
}

// NewCaseAutomaticActionConfig instantiates a new CaseAutomaticActionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaseAutomaticActionConfig() *CaseAutomaticActionConfig {
	this := CaseAutomaticActionConfig{}
	return &this
}

// NewCaseAutomaticActionConfigWithDefaults instantiates a new CaseAutomaticActionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaseAutomaticActionConfigWithDefaults() *CaseAutomaticActionConfig {
	this := CaseAutomaticActionConfig{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CaseAutomaticActionConfig) GetStatus() CaseStatusEnum {
	if o == nil || o.Status == nil {
		var ret CaseStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseAutomaticActionConfig) GetStatusOk() (*CaseStatusEnum, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CaseAutomaticActionConfig) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CaseStatusEnum and assigns it to the Status field.
func (o *CaseAutomaticActionConfig) SetStatus(v CaseStatusEnum) {
	o.Status = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *CaseAutomaticActionConfig) GetSchedule() int32 {
	if o == nil || o.Schedule == nil {
		var ret int32
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseAutomaticActionConfig) GetScheduleOk() (*int32, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *CaseAutomaticActionConfig) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given int32 and assigns it to the Schedule field.
func (o *CaseAutomaticActionConfig) SetSchedule(v int32) {
	o.Schedule = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *CaseAutomaticActionConfig) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseAutomaticActionConfig) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *CaseAutomaticActionConfig) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *CaseAutomaticActionConfig) SetAction(v string) {
	o.Action = &v
}

func (o CaseAutomaticActionConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	return json.Marshal(toSerialize)
}

type NullableCaseAutomaticActionConfig struct {
	value *CaseAutomaticActionConfig
	isSet bool
}

func (v NullableCaseAutomaticActionConfig) Get() *CaseAutomaticActionConfig {
	return v.value
}

func (v *NullableCaseAutomaticActionConfig) Set(val *CaseAutomaticActionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCaseAutomaticActionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCaseAutomaticActionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaseAutomaticActionConfig(val *CaseAutomaticActionConfig) *NullableCaseAutomaticActionConfig {
	return &NullableCaseAutomaticActionConfig{value: val, isSet: true}
}

func (v NullableCaseAutomaticActionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaseAutomaticActionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
