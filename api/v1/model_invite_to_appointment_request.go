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

// checks if the InviteToAppointmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InviteToAppointmentRequest{}

// InviteToAppointmentRequest struct for InviteToAppointmentRequest
type InviteToAppointmentRequest struct {
	Recipients []ContactData `json:"recipients,omitempty"`
	// The text that is sent with the email to the recepients
	InvitationText *string `json:"invitation_text,omitempty"`
}

// NewInviteToAppointmentRequest instantiates a new InviteToAppointmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInviteToAppointmentRequest() *InviteToAppointmentRequest {
	this := InviteToAppointmentRequest{}
	return &this
}

// NewInviteToAppointmentRequestWithDefaults instantiates a new InviteToAppointmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteToAppointmentRequestWithDefaults() *InviteToAppointmentRequest {
	this := InviteToAppointmentRequest{}
	return &this
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *InviteToAppointmentRequest) GetRecipients() []ContactData {
	if o == nil || IsNil(o.Recipients) {
		var ret []ContactData
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteToAppointmentRequest) GetRecipientsOk() ([]ContactData, bool) {
	if o == nil || IsNil(o.Recipients) {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *InviteToAppointmentRequest) HasRecipients() bool {
	if o != nil && !IsNil(o.Recipients) {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given []ContactData and assigns it to the Recipients field.
func (o *InviteToAppointmentRequest) SetRecipients(v []ContactData) {
	o.Recipients = v
}

// GetInvitationText returns the InvitationText field value if set, zero value otherwise.
func (o *InviteToAppointmentRequest) GetInvitationText() string {
	if o == nil || IsNil(o.InvitationText) {
		var ret string
		return ret
	}
	return *o.InvitationText
}

// GetInvitationTextOk returns a tuple with the InvitationText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InviteToAppointmentRequest) GetInvitationTextOk() (*string, bool) {
	if o == nil || IsNil(o.InvitationText) {
		return nil, false
	}
	return o.InvitationText, true
}

// HasInvitationText returns a boolean if a field has been set.
func (o *InviteToAppointmentRequest) HasInvitationText() bool {
	if o != nil && !IsNil(o.InvitationText) {
		return true
	}

	return false
}

// SetInvitationText gets a reference to the given string and assigns it to the InvitationText field.
func (o *InviteToAppointmentRequest) SetInvitationText(v string) {
	o.InvitationText = &v
}

func (o InviteToAppointmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InviteToAppointmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Recipients) {
		toSerialize["recipients"] = o.Recipients
	}
	if !IsNil(o.InvitationText) {
		toSerialize["invitation_text"] = o.InvitationText
	}
	return toSerialize, nil
}

type NullableInviteToAppointmentRequest struct {
	value *InviteToAppointmentRequest
	isSet bool
}

func (v NullableInviteToAppointmentRequest) Get() *InviteToAppointmentRequest {
	return v.value
}

func (v *NullableInviteToAppointmentRequest) Set(val *InviteToAppointmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInviteToAppointmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInviteToAppointmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInviteToAppointmentRequest(val *InviteToAppointmentRequest) *NullableInviteToAppointmentRequest {
	return &NullableInviteToAppointmentRequest{value: val, isSet: true}
}

func (v NullableInviteToAppointmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInviteToAppointmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
