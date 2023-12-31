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

// CaseReminderConfig Defines the configuration for sending out reminders for a case.
type CaseReminderConfig struct {
	Status *CaseStatusEnum `json:"status,omitempty"`
	// The number of days after which the reminder should be sent out. The updated_at attribute is considered for the calculation.
	Schedule []int32 `json:"schedule,omitempty"`
}

// NewCaseReminderConfig instantiates a new CaseReminderConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaseReminderConfig() *CaseReminderConfig {
	this := CaseReminderConfig{}
	return &this
}

// NewCaseReminderConfigWithDefaults instantiates a new CaseReminderConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaseReminderConfigWithDefaults() *CaseReminderConfig {
	this := CaseReminderConfig{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CaseReminderConfig) GetStatus() CaseStatusEnum {
	if o == nil || o.Status == nil {
		var ret CaseStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseReminderConfig) GetStatusOk() (*CaseStatusEnum, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CaseReminderConfig) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CaseStatusEnum and assigns it to the Status field.
func (o *CaseReminderConfig) SetStatus(v CaseStatusEnum) {
	o.Status = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *CaseReminderConfig) GetSchedule() []int32 {
	if o == nil || o.Schedule == nil {
		var ret []int32
		return ret
	}
	return o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseReminderConfig) GetScheduleOk() ([]int32, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *CaseReminderConfig) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given []int32 and assigns it to the Schedule field.
func (o *CaseReminderConfig) SetSchedule(v []int32) {
	o.Schedule = v
}

func (o CaseReminderConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	return json.Marshal(toSerialize)
}

type NullableCaseReminderConfig struct {
	value *CaseReminderConfig
	isSet bool
}

func (v NullableCaseReminderConfig) Get() *CaseReminderConfig {
	return v.value
}

func (v *NullableCaseReminderConfig) Set(val *CaseReminderConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCaseReminderConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCaseReminderConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaseReminderConfig(val *CaseReminderConfig) *NullableCaseReminderConfig {
	return &NullableCaseReminderConfig{value: val, isSet: true}
}

func (v NullableCaseReminderConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaseReminderConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
