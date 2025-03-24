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

// checks if the UpdateUserPreferencesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserPreferencesRequest{}

// UpdateUserPreferencesRequest Change the preferences of the current user
type UpdateUserPreferencesRequest struct {
	// Notify when a new case has been raised for the account. For property  managers or service providers, this will be triggered when a new case is raised for their company/any of the properties the company manages. For tenants it is triggered, when a case is raised for the property  they live in.
	NotifyOnNewCase []NotificationType `json:"notify_on_new_case,omitempty"`
	// Notify when a new case has been updated, which the user or their  company has been active on.
	NotifyOnCaseUpdate []NotificationType `json:"notify_on_case_update,omitempty"`
	// Notify when an appointment has been created, confirmed, cancelled,  cancelled by participant or joined.
	NotifyOnAppointment []NotificationType `json:"notify_on_appointment,omitempty"`
	// Notify when a user has accessed a sent correspondence.
	NotifyOnCorrespondenceAccessed []NotificationType `json:"notify_on_correspondence_accessed,omitempty"`
}

// NewUpdateUserPreferencesRequest instantiates a new UpdateUserPreferencesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserPreferencesRequest() *UpdateUserPreferencesRequest {
	this := UpdateUserPreferencesRequest{}
	return &this
}

// NewUpdateUserPreferencesRequestWithDefaults instantiates a new UpdateUserPreferencesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserPreferencesRequestWithDefaults() *UpdateUserPreferencesRequest {
	this := UpdateUserPreferencesRequest{}
	return &this
}

// GetNotifyOnNewCase returns the NotifyOnNewCase field value if set, zero value otherwise.
func (o *UpdateUserPreferencesRequest) GetNotifyOnNewCase() []NotificationType {
	if o == nil || IsNil(o.NotifyOnNewCase) {
		var ret []NotificationType
		return ret
	}
	return o.NotifyOnNewCase
}

// GetNotifyOnNewCaseOk returns a tuple with the NotifyOnNewCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserPreferencesRequest) GetNotifyOnNewCaseOk() ([]NotificationType, bool) {
	if o == nil || IsNil(o.NotifyOnNewCase) {
		return nil, false
	}
	return o.NotifyOnNewCase, true
}

// HasNotifyOnNewCase returns a boolean if a field has been set.
func (o *UpdateUserPreferencesRequest) HasNotifyOnNewCase() bool {
	if o != nil && !IsNil(o.NotifyOnNewCase) {
		return true
	}

	return false
}

// SetNotifyOnNewCase gets a reference to the given []NotificationType and assigns it to the NotifyOnNewCase field.
func (o *UpdateUserPreferencesRequest) SetNotifyOnNewCase(v []NotificationType) {
	o.NotifyOnNewCase = v
}

// GetNotifyOnCaseUpdate returns the NotifyOnCaseUpdate field value if set, zero value otherwise.
func (o *UpdateUserPreferencesRequest) GetNotifyOnCaseUpdate() []NotificationType {
	if o == nil || IsNil(o.NotifyOnCaseUpdate) {
		var ret []NotificationType
		return ret
	}
	return o.NotifyOnCaseUpdate
}

// GetNotifyOnCaseUpdateOk returns a tuple with the NotifyOnCaseUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserPreferencesRequest) GetNotifyOnCaseUpdateOk() ([]NotificationType, bool) {
	if o == nil || IsNil(o.NotifyOnCaseUpdate) {
		return nil, false
	}
	return o.NotifyOnCaseUpdate, true
}

// HasNotifyOnCaseUpdate returns a boolean if a field has been set.
func (o *UpdateUserPreferencesRequest) HasNotifyOnCaseUpdate() bool {
	if o != nil && !IsNil(o.NotifyOnCaseUpdate) {
		return true
	}

	return false
}

// SetNotifyOnCaseUpdate gets a reference to the given []NotificationType and assigns it to the NotifyOnCaseUpdate field.
func (o *UpdateUserPreferencesRequest) SetNotifyOnCaseUpdate(v []NotificationType) {
	o.NotifyOnCaseUpdate = v
}

// GetNotifyOnAppointment returns the NotifyOnAppointment field value if set, zero value otherwise.
func (o *UpdateUserPreferencesRequest) GetNotifyOnAppointment() []NotificationType {
	if o == nil || IsNil(o.NotifyOnAppointment) {
		var ret []NotificationType
		return ret
	}
	return o.NotifyOnAppointment
}

// GetNotifyOnAppointmentOk returns a tuple with the NotifyOnAppointment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserPreferencesRequest) GetNotifyOnAppointmentOk() ([]NotificationType, bool) {
	if o == nil || IsNil(o.NotifyOnAppointment) {
		return nil, false
	}
	return o.NotifyOnAppointment, true
}

// HasNotifyOnAppointment returns a boolean if a field has been set.
func (o *UpdateUserPreferencesRequest) HasNotifyOnAppointment() bool {
	if o != nil && !IsNil(o.NotifyOnAppointment) {
		return true
	}

	return false
}

// SetNotifyOnAppointment gets a reference to the given []NotificationType and assigns it to the NotifyOnAppointment field.
func (o *UpdateUserPreferencesRequest) SetNotifyOnAppointment(v []NotificationType) {
	o.NotifyOnAppointment = v
}

// GetNotifyOnCorrespondenceAccessed returns the NotifyOnCorrespondenceAccessed field value if set, zero value otherwise.
func (o *UpdateUserPreferencesRequest) GetNotifyOnCorrespondenceAccessed() []NotificationType {
	if o == nil || IsNil(o.NotifyOnCorrespondenceAccessed) {
		var ret []NotificationType
		return ret
	}
	return o.NotifyOnCorrespondenceAccessed
}

// GetNotifyOnCorrespondenceAccessedOk returns a tuple with the NotifyOnCorrespondenceAccessed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserPreferencesRequest) GetNotifyOnCorrespondenceAccessedOk() ([]NotificationType, bool) {
	if o == nil || IsNil(o.NotifyOnCorrespondenceAccessed) {
		return nil, false
	}
	return o.NotifyOnCorrespondenceAccessed, true
}

// HasNotifyOnCorrespondenceAccessed returns a boolean if a field has been set.
func (o *UpdateUserPreferencesRequest) HasNotifyOnCorrespondenceAccessed() bool {
	if o != nil && !IsNil(o.NotifyOnCorrespondenceAccessed) {
		return true
	}

	return false
}

// SetNotifyOnCorrespondenceAccessed gets a reference to the given []NotificationType and assigns it to the NotifyOnCorrespondenceAccessed field.
func (o *UpdateUserPreferencesRequest) SetNotifyOnCorrespondenceAccessed(v []NotificationType) {
	o.NotifyOnCorrespondenceAccessed = v
}

func (o UpdateUserPreferencesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserPreferencesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotifyOnNewCase) {
		toSerialize["notify_on_new_case"] = o.NotifyOnNewCase
	}
	if !IsNil(o.NotifyOnCaseUpdate) {
		toSerialize["notify_on_case_update"] = o.NotifyOnCaseUpdate
	}
	if !IsNil(o.NotifyOnAppointment) {
		toSerialize["notify_on_appointment"] = o.NotifyOnAppointment
	}
	if !IsNil(o.NotifyOnCorrespondenceAccessed) {
		toSerialize["notify_on_correspondence_accessed"] = o.NotifyOnCorrespondenceAccessed
	}
	return toSerialize, nil
}

type NullableUpdateUserPreferencesRequest struct {
	value *UpdateUserPreferencesRequest
	isSet bool
}

func (v NullableUpdateUserPreferencesRequest) Get() *UpdateUserPreferencesRequest {
	return v.value
}

func (v *NullableUpdateUserPreferencesRequest) Set(val *UpdateUserPreferencesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserPreferencesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserPreferencesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserPreferencesRequest(val *UpdateUserPreferencesRequest) *NullableUpdateUserPreferencesRequest {
	return &NullableUpdateUserPreferencesRequest{value: val, isSet: true}
}

func (v NullableUpdateUserPreferencesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserPreferencesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
