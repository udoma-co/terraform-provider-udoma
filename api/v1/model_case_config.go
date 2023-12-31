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

// CaseConfig Defines custom behaviour of a case, based on the case template that was  used to create it.
type CaseConfig struct {
	BaseConfig *BaseCaseConfig `json:"base_config,omitempty"`
	// Indicates if the default status configuration should be extended with  the custom configuration. If false, the custom configuration will  replace the default configuration. If true, only the actions that are  defined in the custom configuration will override the default ones.
	ExtendDefaultStatusConfig *bool `json:"extend_default_status_config,omitempty"`
	// The configuration for the status transition of a case. This is used to  determine which status changes are allowed by which party at which time.
	StatusConfig []CaseStatusConfig `json:"status_config,omitempty"`
	// The configuration for sending out reminders for a case.
	Reminders []CaseReminderConfig `json:"reminders,omitempty"`
	// The configuration for automatic status changes of a case. This is used to  determine which status changes are allowed by which party at which time.
	AutomaticStatusChanges []CaseAutomaticStatusChangeConfig `json:"automatic_status_changes,omitempty"`
}

// NewCaseConfig instantiates a new CaseConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaseConfig() *CaseConfig {
	this := CaseConfig{}
	return &this
}

// NewCaseConfigWithDefaults instantiates a new CaseConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaseConfigWithDefaults() *CaseConfig {
	this := CaseConfig{}
	return &this
}

// GetBaseConfig returns the BaseConfig field value if set, zero value otherwise.
func (o *CaseConfig) GetBaseConfig() BaseCaseConfig {
	if o == nil || o.BaseConfig == nil {
		var ret BaseCaseConfig
		return ret
	}
	return *o.BaseConfig
}

// GetBaseConfigOk returns a tuple with the BaseConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseConfig) GetBaseConfigOk() (*BaseCaseConfig, bool) {
	if o == nil || o.BaseConfig == nil {
		return nil, false
	}
	return o.BaseConfig, true
}

// HasBaseConfig returns a boolean if a field has been set.
func (o *CaseConfig) HasBaseConfig() bool {
	if o != nil && o.BaseConfig != nil {
		return true
	}

	return false
}

// SetBaseConfig gets a reference to the given BaseCaseConfig and assigns it to the BaseConfig field.
func (o *CaseConfig) SetBaseConfig(v BaseCaseConfig) {
	o.BaseConfig = &v
}

// GetExtendDefaultStatusConfig returns the ExtendDefaultStatusConfig field value if set, zero value otherwise.
func (o *CaseConfig) GetExtendDefaultStatusConfig() bool {
	if o == nil || o.ExtendDefaultStatusConfig == nil {
		var ret bool
		return ret
	}
	return *o.ExtendDefaultStatusConfig
}

// GetExtendDefaultStatusConfigOk returns a tuple with the ExtendDefaultStatusConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseConfig) GetExtendDefaultStatusConfigOk() (*bool, bool) {
	if o == nil || o.ExtendDefaultStatusConfig == nil {
		return nil, false
	}
	return o.ExtendDefaultStatusConfig, true
}

// HasExtendDefaultStatusConfig returns a boolean if a field has been set.
func (o *CaseConfig) HasExtendDefaultStatusConfig() bool {
	if o != nil && o.ExtendDefaultStatusConfig != nil {
		return true
	}

	return false
}

// SetExtendDefaultStatusConfig gets a reference to the given bool and assigns it to the ExtendDefaultStatusConfig field.
func (o *CaseConfig) SetExtendDefaultStatusConfig(v bool) {
	o.ExtendDefaultStatusConfig = &v
}

// GetStatusConfig returns the StatusConfig field value if set, zero value otherwise.
func (o *CaseConfig) GetStatusConfig() []CaseStatusConfig {
	if o == nil || o.StatusConfig == nil {
		var ret []CaseStatusConfig
		return ret
	}
	return o.StatusConfig
}

// GetStatusConfigOk returns a tuple with the StatusConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseConfig) GetStatusConfigOk() ([]CaseStatusConfig, bool) {
	if o == nil || o.StatusConfig == nil {
		return nil, false
	}
	return o.StatusConfig, true
}

// HasStatusConfig returns a boolean if a field has been set.
func (o *CaseConfig) HasStatusConfig() bool {
	if o != nil && o.StatusConfig != nil {
		return true
	}

	return false
}

// SetStatusConfig gets a reference to the given []CaseStatusConfig and assigns it to the StatusConfig field.
func (o *CaseConfig) SetStatusConfig(v []CaseStatusConfig) {
	o.StatusConfig = v
}

// GetReminders returns the Reminders field value if set, zero value otherwise.
func (o *CaseConfig) GetReminders() []CaseReminderConfig {
	if o == nil || o.Reminders == nil {
		var ret []CaseReminderConfig
		return ret
	}
	return o.Reminders
}

// GetRemindersOk returns a tuple with the Reminders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseConfig) GetRemindersOk() ([]CaseReminderConfig, bool) {
	if o == nil || o.Reminders == nil {
		return nil, false
	}
	return o.Reminders, true
}

// HasReminders returns a boolean if a field has been set.
func (o *CaseConfig) HasReminders() bool {
	if o != nil && o.Reminders != nil {
		return true
	}

	return false
}

// SetReminders gets a reference to the given []CaseReminderConfig and assigns it to the Reminders field.
func (o *CaseConfig) SetReminders(v []CaseReminderConfig) {
	o.Reminders = v
}

// GetAutomaticStatusChanges returns the AutomaticStatusChanges field value if set, zero value otherwise.
func (o *CaseConfig) GetAutomaticStatusChanges() []CaseAutomaticStatusChangeConfig {
	if o == nil || o.AutomaticStatusChanges == nil {
		var ret []CaseAutomaticStatusChangeConfig
		return ret
	}
	return o.AutomaticStatusChanges
}

// GetAutomaticStatusChangesOk returns a tuple with the AutomaticStatusChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseConfig) GetAutomaticStatusChangesOk() ([]CaseAutomaticStatusChangeConfig, bool) {
	if o == nil || o.AutomaticStatusChanges == nil {
		return nil, false
	}
	return o.AutomaticStatusChanges, true
}

// HasAutomaticStatusChanges returns a boolean if a field has been set.
func (o *CaseConfig) HasAutomaticStatusChanges() bool {
	if o != nil && o.AutomaticStatusChanges != nil {
		return true
	}

	return false
}

// SetAutomaticStatusChanges gets a reference to the given []CaseAutomaticStatusChangeConfig and assigns it to the AutomaticStatusChanges field.
func (o *CaseConfig) SetAutomaticStatusChanges(v []CaseAutomaticStatusChangeConfig) {
	o.AutomaticStatusChanges = v
}

func (o CaseConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BaseConfig != nil {
		toSerialize["base_config"] = o.BaseConfig
	}
	if o.ExtendDefaultStatusConfig != nil {
		toSerialize["extend_default_status_config"] = o.ExtendDefaultStatusConfig
	}
	if o.StatusConfig != nil {
		toSerialize["status_config"] = o.StatusConfig
	}
	if o.Reminders != nil {
		toSerialize["reminders"] = o.Reminders
	}
	if o.AutomaticStatusChanges != nil {
		toSerialize["automatic_status_changes"] = o.AutomaticStatusChanges
	}
	return json.Marshal(toSerialize)
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
