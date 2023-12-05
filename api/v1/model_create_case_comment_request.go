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

// CreateCaseCommentRequest Request that will create a new comment for a case
type CreateCaseCommentRequest struct {
	Content string `json:"content"`
	// list of attachments that should be linked to the comment
	Attachments []Attachment `json:"attachments,omitempty"`
	// List of paries that should have access to the comment
	Visibility []CaseParty `json:"visibility,omitempty"`
}

// NewCreateCaseCommentRequest instantiates a new CreateCaseCommentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCaseCommentRequest(content string) *CreateCaseCommentRequest {
	this := CreateCaseCommentRequest{}
	this.Content = content
	return &this
}

// NewCreateCaseCommentRequestWithDefaults instantiates a new CreateCaseCommentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCaseCommentRequestWithDefaults() *CreateCaseCommentRequest {
	this := CreateCaseCommentRequest{}
	return &this
}

// GetContent returns the Content field value
func (o *CreateCaseCommentRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *CreateCaseCommentRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *CreateCaseCommentRequest) SetContent(v string) {
	o.Content = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *CreateCaseCommentRequest) GetAttachments() []Attachment {
	if o == nil || o.Attachments == nil {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCaseCommentRequest) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || o.Attachments == nil {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *CreateCaseCommentRequest) HasAttachments() bool {
	if o != nil && o.Attachments != nil {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *CreateCaseCommentRequest) SetAttachments(v []Attachment) {
	o.Attachments = v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *CreateCaseCommentRequest) GetVisibility() []CaseParty {
	if o == nil || o.Visibility == nil {
		var ret []CaseParty
		return ret
	}
	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCaseCommentRequest) GetVisibilityOk() ([]CaseParty, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *CreateCaseCommentRequest) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given []CaseParty and assigns it to the Visibility field.
func (o *CreateCaseCommentRequest) SetVisibility(v []CaseParty) {
	o.Visibility = v
}

func (o CreateCaseCommentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content"] = o.Content
	}
	if o.Attachments != nil {
		toSerialize["attachments"] = o.Attachments
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	return json.Marshal(toSerialize)
}

type NullableCreateCaseCommentRequest struct {
	value *CreateCaseCommentRequest
	isSet bool
}

func (v NullableCreateCaseCommentRequest) Get() *CreateCaseCommentRequest {
	return v.value
}

func (v *NullableCreateCaseCommentRequest) Set(val *CreateCaseCommentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCaseCommentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCaseCommentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCaseCommentRequest(val *CreateCaseCommentRequest) *NullableCreateCaseCommentRequest {
	return &NullableCreateCaseCommentRequest{value: val, isSet: true}
}

func (v NullableCreateCaseCommentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCaseCommentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
