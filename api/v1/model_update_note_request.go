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

// UpdateNoteRequest A request to update an existing note
type UpdateNoteRequest struct {
	// the contents of the note
	Text string `json:"text"`
}

// NewUpdateNoteRequest instantiates a new UpdateNoteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNoteRequest(text string) *UpdateNoteRequest {
	this := UpdateNoteRequest{}
	this.Text = text
	return &this
}

// NewUpdateNoteRequestWithDefaults instantiates a new UpdateNoteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNoteRequestWithDefaults() *UpdateNoteRequest {
	this := UpdateNoteRequest{}
	return &this
}

// GetText returns the Text field value
func (o *UpdateNoteRequest) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *UpdateNoteRequest) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *UpdateNoteRequest) SetText(v string) {
	o.Text = v
}

func (o UpdateNoteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["text"] = o.Text
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNoteRequest struct {
	value *UpdateNoteRequest
	isSet bool
}

func (v NullableUpdateNoteRequest) Get() *UpdateNoteRequest {
	return v.value
}

func (v *NullableUpdateNoteRequest) Set(val *UpdateNoteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNoteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNoteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNoteRequest(val *UpdateNoteRequest) *NullableUpdateNoteRequest {
	return &NullableUpdateNoteRequest{value: val, isSet: true}
}

func (v NullableUpdateNoteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNoteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
