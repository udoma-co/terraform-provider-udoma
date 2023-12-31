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

// DocumentRepositoryEntry A single document uploaded in a document repository
type DocumentRepositoryEntry struct {
	// the unique ID of the document entry
	Id   *string `json:"id,omitempty"`
	Path *string `json:"path,omitempty"`
	// The name of the document entry
	Name       *string     `json:"name,omitempty"`
	Public     *bool       `json:"public,omitempty"`
	Attachment *Attachment `json:"attachment,omitempty"`
}

// NewDocumentRepositoryEntry instantiates a new DocumentRepositoryEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentRepositoryEntry() *DocumentRepositoryEntry {
	this := DocumentRepositoryEntry{}
	return &this
}

// NewDocumentRepositoryEntryWithDefaults instantiates a new DocumentRepositoryEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentRepositoryEntryWithDefaults() *DocumentRepositoryEntry {
	this := DocumentRepositoryEntry{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DocumentRepositoryEntry) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentRepositoryEntry) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DocumentRepositoryEntry) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DocumentRepositoryEntry) SetId(v string) {
	o.Id = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *DocumentRepositoryEntry) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentRepositoryEntry) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *DocumentRepositoryEntry) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *DocumentRepositoryEntry) SetPath(v string) {
	o.Path = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DocumentRepositoryEntry) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentRepositoryEntry) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DocumentRepositoryEntry) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DocumentRepositoryEntry) SetName(v string) {
	o.Name = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *DocumentRepositoryEntry) GetPublic() bool {
	if o == nil || o.Public == nil {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentRepositoryEntry) GetPublicOk() (*bool, bool) {
	if o == nil || o.Public == nil {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *DocumentRepositoryEntry) HasPublic() bool {
	if o != nil && o.Public != nil {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *DocumentRepositoryEntry) SetPublic(v bool) {
	o.Public = &v
}

// GetAttachment returns the Attachment field value if set, zero value otherwise.
func (o *DocumentRepositoryEntry) GetAttachment() Attachment {
	if o == nil || o.Attachment == nil {
		var ret Attachment
		return ret
	}
	return *o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentRepositoryEntry) GetAttachmentOk() (*Attachment, bool) {
	if o == nil || o.Attachment == nil {
		return nil, false
	}
	return o.Attachment, true
}

// HasAttachment returns a boolean if a field has been set.
func (o *DocumentRepositoryEntry) HasAttachment() bool {
	if o != nil && o.Attachment != nil {
		return true
	}

	return false
}

// SetAttachment gets a reference to the given Attachment and assigns it to the Attachment field.
func (o *DocumentRepositoryEntry) SetAttachment(v Attachment) {
	o.Attachment = &v
}

func (o DocumentRepositoryEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Public != nil {
		toSerialize["public"] = o.Public
	}
	if o.Attachment != nil {
		toSerialize["attachment"] = o.Attachment
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentRepositoryEntry struct {
	value *DocumentRepositoryEntry
	isSet bool
}

func (v NullableDocumentRepositoryEntry) Get() *DocumentRepositoryEntry {
	return v.value
}

func (v *NullableDocumentRepositoryEntry) Set(val *DocumentRepositoryEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentRepositoryEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentRepositoryEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentRepositoryEntry(val *DocumentRepositoryEntry) *NullableDocumentRepositoryEntry {
	return &NullableDocumentRepositoryEntry{value: val, isSet: true}
}

func (v NullableDocumentRepositoryEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentRepositoryEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
