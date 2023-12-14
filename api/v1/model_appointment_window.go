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

// AppointmentWindow struct for AppointmentWindow
type AppointmentWindow struct {
	// The id of the window - only provided as a response or in update requests
	Id *string `json:"id,omitempty"`
	// The timestamp of the beginning of the appointment window
	StartTime int64 `json:"start_time"`
	// The timestamp of the end of the appointment window
	EndTime int64 `json:"end_time"`
	// List of slots that are available or booked
	Slots []AppointmentSlot `json:"slots,omitempty"`
}

// NewAppointmentWindow instantiates a new AppointmentWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppointmentWindow(startTime int64, endTime int64) *AppointmentWindow {
	this := AppointmentWindow{}
	this.StartTime = startTime
	this.EndTime = endTime
	return &this
}

// NewAppointmentWindowWithDefaults instantiates a new AppointmentWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppointmentWindowWithDefaults() *AppointmentWindow {
	this := AppointmentWindow{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppointmentWindow) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentWindow) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppointmentWindow) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppointmentWindow) SetId(v string) {
	o.Id = &v
}

// GetStartTime returns the StartTime field value
func (o *AppointmentWindow) GetStartTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *AppointmentWindow) GetStartTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *AppointmentWindow) SetStartTime(v int64) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *AppointmentWindow) GetEndTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *AppointmentWindow) GetEndTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *AppointmentWindow) SetEndTime(v int64) {
	o.EndTime = v
}

// GetSlots returns the Slots field value if set, zero value otherwise.
func (o *AppointmentWindow) GetSlots() []AppointmentSlot {
	if o == nil || o.Slots == nil {
		var ret []AppointmentSlot
		return ret
	}
	return o.Slots
}

// GetSlotsOk returns a tuple with the Slots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentWindow) GetSlotsOk() ([]AppointmentSlot, bool) {
	if o == nil || o.Slots == nil {
		return nil, false
	}
	return o.Slots, true
}

// HasSlots returns a boolean if a field has been set.
func (o *AppointmentWindow) HasSlots() bool {
	if o != nil && o.Slots != nil {
		return true
	}

	return false
}

// SetSlots gets a reference to the given []AppointmentSlot and assigns it to the Slots field.
func (o *AppointmentWindow) SetSlots(v []AppointmentSlot) {
	o.Slots = v
}

func (o AppointmentWindow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["start_time"] = o.StartTime
	}
	if true {
		toSerialize["end_time"] = o.EndTime
	}
	if o.Slots != nil {
		toSerialize["slots"] = o.Slots
	}
	return json.Marshal(toSerialize)
}

type NullableAppointmentWindow struct {
	value *AppointmentWindow
	isSet bool
}

func (v NullableAppointmentWindow) Get() *AppointmentWindow {
	return v.value
}

func (v *NullableAppointmentWindow) Set(val *AppointmentWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableAppointmentWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableAppointmentWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppointmentWindow(val *AppointmentWindow) *NullableAppointmentWindow {
	return &NullableAppointmentWindow{value: val, isSet: true}
}

func (v NullableAppointmentWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppointmentWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
