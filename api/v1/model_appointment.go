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

// Appointment An appointment is a time slot that has already been booked by a user.
type Appointment struct {
	// The unique identifier of the appointment
	Id *string `json:"id,omitempty"`
	// The code of the appointment if the appointment is public
	Code *string `json:"code,omitempty"`
	// The timestamp of when the appointment was created
	CreatedAt *int64 `json:"created_at,omitempty"`
	// The timestamp of when the appointment was last updated
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// The timestamp of the appointment, i.e. when it starts
	StartTime *int64 `json:"start_time,omitempty"`
	// The timestamp of the appointment, i.e. when it ends
	EndTime *int64 `json:"end_time,omitempty"`
	// A descriptive name for the appointment, displayed on the calendar overview
	Name    *string      `json:"name,omitempty"`
	Contact *ContactData `json:"contact,omitempty"`
	// Input provided by the user when booking the appointment
	Data map[string]interface{} `json:"data,omitempty"`
}

// NewAppointment instantiates a new Appointment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppointment() *Appointment {
	this := Appointment{}
	return &this
}

// NewAppointmentWithDefaults instantiates a new Appointment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppointmentWithDefaults() *Appointment {
	this := Appointment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Appointment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appointment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Appointment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Appointment) SetId(v string) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *Appointment) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appointment) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *Appointment) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *Appointment) SetCode(v string) {
	o.Code = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Appointment) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appointment) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Appointment) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *Appointment) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Appointment) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appointment) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Appointment) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *Appointment) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *Appointment) GetStartTime() int64 {
	if o == nil || o.StartTime == nil {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appointment) GetStartTimeOk() (*int64, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *Appointment) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *Appointment) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *Appointment) GetEndTime() int64 {
	if o == nil || o.EndTime == nil {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appointment) GetEndTimeOk() (*int64, bool) {
	if o == nil || o.EndTime == nil {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *Appointment) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *Appointment) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Appointment) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appointment) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Appointment) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Appointment) SetName(v string) {
	o.Name = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *Appointment) GetContact() ContactData {
	if o == nil || o.Contact == nil {
		var ret ContactData
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appointment) GetContactOk() (*ContactData, bool) {
	if o == nil || o.Contact == nil {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *Appointment) HasContact() bool {
	if o != nil && o.Contact != nil {
		return true
	}

	return false
}

// SetContact gets a reference to the given ContactData and assigns it to the Contact field.
func (o *Appointment) SetContact(v ContactData) {
	o.Contact = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Appointment) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Appointment) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Appointment) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *Appointment) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o Appointment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	if o.EndTime != nil {
		toSerialize["end_time"] = o.EndTime
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Contact != nil {
		toSerialize["contact"] = o.Contact
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableAppointment struct {
	value *Appointment
	isSet bool
}

func (v NullableAppointment) Get() *Appointment {
	return v.value
}

func (v *NullableAppointment) Set(val *Appointment) {
	v.value = val
	v.isSet = true
}

func (v NullableAppointment) IsSet() bool {
	return v.isSet
}

func (v *NullableAppointment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppointment(val *Appointment) *NullableAppointment {
	return &NullableAppointment{value: val, isSet: true}
}

func (v NullableAppointment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppointment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
