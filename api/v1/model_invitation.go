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

// Invitation struct for Invitation
type Invitation struct {
	// The unique ID of the invitation
	Id *string `json:"id,omitempty"`
	// The timestamp of when the invitation was created
	Created *int64 `json:"created,omitempty"`
	// The timestamp of when the invitation was accepted or rejected
	Processed *int64 `json:"processed,omitempty"`
	// The email of the tenant who was invited
	UserEmail *string           `json:"user_email,omitempty"`
	Status    *InvitationStatus `json:"status,omitempty"`
	Property  *Property         `json:"property,omitempty"`
}

// NewInvitation instantiates a new Invitation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitation() *Invitation {
	this := Invitation{}
	return &this
}

// NewInvitationWithDefaults instantiates a new Invitation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationWithDefaults() *Invitation {
	this := Invitation{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Invitation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Invitation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Invitation) SetId(v string) {
	o.Id = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Invitation) GetCreated() int64 {
	if o == nil || o.Created == nil {
		var ret int64
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetCreatedOk() (*int64, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Invitation) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int64 and assigns it to the Created field.
func (o *Invitation) SetCreated(v int64) {
	o.Created = &v
}

// GetProcessed returns the Processed field value if set, zero value otherwise.
func (o *Invitation) GetProcessed() int64 {
	if o == nil || o.Processed == nil {
		var ret int64
		return ret
	}
	return *o.Processed
}

// GetProcessedOk returns a tuple with the Processed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetProcessedOk() (*int64, bool) {
	if o == nil || o.Processed == nil {
		return nil, false
	}
	return o.Processed, true
}

// HasProcessed returns a boolean if a field has been set.
func (o *Invitation) HasProcessed() bool {
	if o != nil && o.Processed != nil {
		return true
	}

	return false
}

// SetProcessed gets a reference to the given int64 and assigns it to the Processed field.
func (o *Invitation) SetProcessed(v int64) {
	o.Processed = &v
}

// GetUserEmail returns the UserEmail field value if set, zero value otherwise.
func (o *Invitation) GetUserEmail() string {
	if o == nil || o.UserEmail == nil {
		var ret string
		return ret
	}
	return *o.UserEmail
}

// GetUserEmailOk returns a tuple with the UserEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetUserEmailOk() (*string, bool) {
	if o == nil || o.UserEmail == nil {
		return nil, false
	}
	return o.UserEmail, true
}

// HasUserEmail returns a boolean if a field has been set.
func (o *Invitation) HasUserEmail() bool {
	if o != nil && o.UserEmail != nil {
		return true
	}

	return false
}

// SetUserEmail gets a reference to the given string and assigns it to the UserEmail field.
func (o *Invitation) SetUserEmail(v string) {
	o.UserEmail = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Invitation) GetStatus() InvitationStatus {
	if o == nil || o.Status == nil {
		var ret InvitationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetStatusOk() (*InvitationStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Invitation) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given InvitationStatus and assigns it to the Status field.
func (o *Invitation) SetStatus(v InvitationStatus) {
	o.Status = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *Invitation) GetProperty() Property {
	if o == nil || o.Property == nil {
		var ret Property
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invitation) GetPropertyOk() (*Property, bool) {
	if o == nil || o.Property == nil {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *Invitation) HasProperty() bool {
	if o != nil && o.Property != nil {
		return true
	}

	return false
}

// SetProperty gets a reference to the given Property and assigns it to the Property field.
func (o *Invitation) SetProperty(v Property) {
	o.Property = &v
}

func (o Invitation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.Processed != nil {
		toSerialize["processed"] = o.Processed
	}
	if o.UserEmail != nil {
		toSerialize["user_email"] = o.UserEmail
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Property != nil {
		toSerialize["property"] = o.Property
	}
	return json.Marshal(toSerialize)
}

type NullableInvitation struct {
	value *Invitation
	isSet bool
}

func (v NullableInvitation) Get() *Invitation {
	return v.value
}

func (v *NullableInvitation) Set(val *Invitation) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitation) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitation(val *Invitation) *NullableInvitation {
	return &NullableInvitation{value: val, isSet: true}
}

func (v NullableInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
