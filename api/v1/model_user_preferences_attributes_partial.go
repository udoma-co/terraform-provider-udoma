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

// checks if the UserPreferencesAttributesPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPreferencesAttributesPartial{}

// UserPreferencesAttributesPartial Notification preferences for individual users
type UserPreferencesAttributesPartial struct {
	// Notify when a new case has been raised for the account. For property  managers or service providers, this will be triggered when a new case is raised for their company/any of the properties the company manages. For tenants it is triggered, when a case is raised for the property  they live in.
	NotifyOnNewCase []NotificationType `json:"notify_on_new_case,omitempty"`
	// Notify when a new case has been updated, which the user or their  company has been active on.
	NotifyOnCaseUpdate []NotificationType `json:"notify_on_case_update,omitempty"`
}

// NewUserPreferencesAttributesPartial instantiates a new UserPreferencesAttributesPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPreferencesAttributesPartial() *UserPreferencesAttributesPartial {
	this := UserPreferencesAttributesPartial{}
	return &this
}

// NewUserPreferencesAttributesPartialWithDefaults instantiates a new UserPreferencesAttributesPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPreferencesAttributesPartialWithDefaults() *UserPreferencesAttributesPartial {
	this := UserPreferencesAttributesPartial{}
	return &this
}

// GetNotifyOnNewCase returns the NotifyOnNewCase field value if set, zero value otherwise.
func (o *UserPreferencesAttributesPartial) GetNotifyOnNewCase() []NotificationType {
	if o == nil || IsNil(o.NotifyOnNewCase) {
		var ret []NotificationType
		return ret
	}
	return o.NotifyOnNewCase
}

// GetNotifyOnNewCaseOk returns a tuple with the NotifyOnNewCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferencesAttributesPartial) GetNotifyOnNewCaseOk() ([]NotificationType, bool) {
	if o == nil || IsNil(o.NotifyOnNewCase) {
		return nil, false
	}
	return o.NotifyOnNewCase, true
}

// HasNotifyOnNewCase returns a boolean if a field has been set.
func (o *UserPreferencesAttributesPartial) HasNotifyOnNewCase() bool {
	if o != nil && !IsNil(o.NotifyOnNewCase) {
		return true
	}

	return false
}

// SetNotifyOnNewCase gets a reference to the given []NotificationType and assigns it to the NotifyOnNewCase field.
func (o *UserPreferencesAttributesPartial) SetNotifyOnNewCase(v []NotificationType) {
	o.NotifyOnNewCase = v
}

// GetNotifyOnCaseUpdate returns the NotifyOnCaseUpdate field value if set, zero value otherwise.
func (o *UserPreferencesAttributesPartial) GetNotifyOnCaseUpdate() []NotificationType {
	if o == nil || IsNil(o.NotifyOnCaseUpdate) {
		var ret []NotificationType
		return ret
	}
	return o.NotifyOnCaseUpdate
}

// GetNotifyOnCaseUpdateOk returns a tuple with the NotifyOnCaseUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferencesAttributesPartial) GetNotifyOnCaseUpdateOk() ([]NotificationType, bool) {
	if o == nil || IsNil(o.NotifyOnCaseUpdate) {
		return nil, false
	}
	return o.NotifyOnCaseUpdate, true
}

// HasNotifyOnCaseUpdate returns a boolean if a field has been set.
func (o *UserPreferencesAttributesPartial) HasNotifyOnCaseUpdate() bool {
	if o != nil && !IsNil(o.NotifyOnCaseUpdate) {
		return true
	}

	return false
}

// SetNotifyOnCaseUpdate gets a reference to the given []NotificationType and assigns it to the NotifyOnCaseUpdate field.
func (o *UserPreferencesAttributesPartial) SetNotifyOnCaseUpdate(v []NotificationType) {
	o.NotifyOnCaseUpdate = v
}

func (o UserPreferencesAttributesPartial) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPreferencesAttributesPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotifyOnNewCase) {
		toSerialize["notify_on_new_case"] = o.NotifyOnNewCase
	}
	if !IsNil(o.NotifyOnCaseUpdate) {
		toSerialize["notify_on_case_update"] = o.NotifyOnCaseUpdate
	}
	return toSerialize, nil
}

type NullableUserPreferencesAttributesPartial struct {
	value *UserPreferencesAttributesPartial
	isSet bool
}

func (v NullableUserPreferencesAttributesPartial) Get() *UserPreferencesAttributesPartial {
	return v.value
}

func (v *NullableUserPreferencesAttributesPartial) Set(val *UserPreferencesAttributesPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPreferencesAttributesPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPreferencesAttributesPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPreferencesAttributesPartial(val *UserPreferencesAttributesPartial) *NullableUserPreferencesAttributesPartial {
	return &NullableUserPreferencesAttributesPartial{value: val, isSet: true}
}

func (v NullableUserPreferencesAttributesPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPreferencesAttributesPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}