/*
Udoma

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CaseConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaseConfig{}

// CaseConfig Defines custom behaviour of a case, based on the case template that was  used to create it.
type CaseConfig struct {
	// The configuration for the status transition of a case. This is used to  determine which status changes are allowed by which party at which time.
	StatusConfig []CaseStatusConfig `json:"status_config"`
	// The configuration for sending out reminders for a case.
	Reminders []CaseReminderConfig `json:"reminders,omitempty"`
	// The configuration for automatic actions to be performed for case. This is used to perform complex actions, like changing the status of a case or adding an automatic comment.
	AutomaticActions []CaseAutomaticActionConfig `json:"automatic_actions,omitempty"`
}

type _CaseConfig CaseConfig

// NewCaseConfig instantiates a new CaseConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaseConfig(statusConfig []CaseStatusConfig) *CaseConfig {
	this := CaseConfig{}
	this.StatusConfig = statusConfig
	return &this
}

// NewCaseConfigWithDefaults instantiates a new CaseConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaseConfigWithDefaults() *CaseConfig {
	this := CaseConfig{}
	return &this
}

// GetStatusConfig returns the StatusConfig field value
func (o *CaseConfig) GetStatusConfig() []CaseStatusConfig {
	if o == nil {
		var ret []CaseStatusConfig
		return ret
	}

	return o.StatusConfig
}

// GetStatusConfigOk returns a tuple with the StatusConfig field value
// and a boolean to check if the value has been set.
func (o *CaseConfig) GetStatusConfigOk() ([]CaseStatusConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.StatusConfig, true
}

// SetStatusConfig sets field value
func (o *CaseConfig) SetStatusConfig(v []CaseStatusConfig) {
	o.StatusConfig = v
}

// GetReminders returns the Reminders field value if set, zero value otherwise.
func (o *CaseConfig) GetReminders() []CaseReminderConfig {
	if o == nil || IsNil(o.Reminders) {
		var ret []CaseReminderConfig
		return ret
	}
	return o.Reminders
}

// GetRemindersOk returns a tuple with the Reminders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseConfig) GetRemindersOk() ([]CaseReminderConfig, bool) {
	if o == nil || IsNil(o.Reminders) {
		return nil, false
	}
	return o.Reminders, true
}

// HasReminders returns a boolean if a field has been set.
func (o *CaseConfig) HasReminders() bool {
	if o != nil && !IsNil(o.Reminders) {
		return true
	}

	return false
}

// SetReminders gets a reference to the given []CaseReminderConfig and assigns it to the Reminders field.
func (o *CaseConfig) SetReminders(v []CaseReminderConfig) {
	o.Reminders = v
}

// GetAutomaticActions returns the AutomaticActions field value if set, zero value otherwise.
func (o *CaseConfig) GetAutomaticActions() []CaseAutomaticActionConfig {
	if o == nil || IsNil(o.AutomaticActions) {
		var ret []CaseAutomaticActionConfig
		return ret
	}
	return o.AutomaticActions
}

// GetAutomaticActionsOk returns a tuple with the AutomaticActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseConfig) GetAutomaticActionsOk() ([]CaseAutomaticActionConfig, bool) {
	if o == nil || IsNil(o.AutomaticActions) {
		return nil, false
	}
	return o.AutomaticActions, true
}

// HasAutomaticActions returns a boolean if a field has been set.
func (o *CaseConfig) HasAutomaticActions() bool {
	if o != nil && !IsNil(o.AutomaticActions) {
		return true
	}

	return false
}

// SetAutomaticActions gets a reference to the given []CaseAutomaticActionConfig and assigns it to the AutomaticActions field.
func (o *CaseConfig) SetAutomaticActions(v []CaseAutomaticActionConfig) {
	o.AutomaticActions = v
}

func (o CaseConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaseConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status_config"] = o.StatusConfig
	if !IsNil(o.Reminders) {
		toSerialize["reminders"] = o.Reminders
	}
	if !IsNil(o.AutomaticActions) {
		toSerialize["automatic_actions"] = o.AutomaticActions
	}
	return toSerialize, nil
}

func (o *CaseConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status_config",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCaseConfig := _CaseConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCaseConfig)

	if err != nil {
		return err
	}

	*o = CaseConfig(varCaseConfig)

	return err
}

type NullableCaseConfig struct {
	value *CaseConfig
	isSet bool
}

func (v NullableCaseConfig) Get() *CaseConfig {
	return v.value
}

func (v *NullableCaseConfig) Set(val *CaseConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCaseConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCaseConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaseConfig(val *CaseConfig) *NullableCaseConfig {
	return &NullableCaseConfig{value: val, isSet: true}
}

func (v NullableCaseConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaseConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
