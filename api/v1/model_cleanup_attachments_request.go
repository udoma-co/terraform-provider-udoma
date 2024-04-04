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

// checks if the CleanupAttachmentsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CleanupAttachmentsRequest{}

// CleanupAttachmentsRequest Request used to cleanup attachments
type CleanupAttachmentsRequest struct {
	// list of IDs of attachments to be deleted
	AttachmentRefs []string `json:"attachment_refs,omitempty"`
}

// NewCleanupAttachmentsRequest instantiates a new CleanupAttachmentsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCleanupAttachmentsRequest() *CleanupAttachmentsRequest {
	this := CleanupAttachmentsRequest{}
	return &this
}

// NewCleanupAttachmentsRequestWithDefaults instantiates a new CleanupAttachmentsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCleanupAttachmentsRequestWithDefaults() *CleanupAttachmentsRequest {
	this := CleanupAttachmentsRequest{}
	return &this
}

// GetAttachmentRefs returns the AttachmentRefs field value if set, zero value otherwise.
func (o *CleanupAttachmentsRequest) GetAttachmentRefs() []string {
	if o == nil || IsNil(o.AttachmentRefs) {
		var ret []string
		return ret
	}
	return o.AttachmentRefs
}

// GetAttachmentRefsOk returns a tuple with the AttachmentRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CleanupAttachmentsRequest) GetAttachmentRefsOk() ([]string, bool) {
	if o == nil || IsNil(o.AttachmentRefs) {
		return nil, false
	}
	return o.AttachmentRefs, true
}

// HasAttachmentRefs returns a boolean if a field has been set.
func (o *CleanupAttachmentsRequest) HasAttachmentRefs() bool {
	if o != nil && !IsNil(o.AttachmentRefs) {
		return true
	}

	return false
}

// SetAttachmentRefs gets a reference to the given []string and assigns it to the AttachmentRefs field.
func (o *CleanupAttachmentsRequest) SetAttachmentRefs(v []string) {
	o.AttachmentRefs = v
}

func (o CleanupAttachmentsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CleanupAttachmentsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AttachmentRefs) {
		toSerialize["attachment_refs"] = o.AttachmentRefs
	}
	return toSerialize, nil
}

type NullableCleanupAttachmentsRequest struct {
	value *CleanupAttachmentsRequest
	isSet bool
}

func (v NullableCleanupAttachmentsRequest) Get() *CleanupAttachmentsRequest {
	return v.value
}

func (v *NullableCleanupAttachmentsRequest) Set(val *CleanupAttachmentsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCleanupAttachmentsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCleanupAttachmentsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCleanupAttachmentsRequest(val *CleanupAttachmentsRequest) *NullableCleanupAttachmentsRequest {
	return &NullableCleanupAttachmentsRequest{value: val, isSet: true}
}

func (v NullableCleanupAttachmentsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCleanupAttachmentsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
