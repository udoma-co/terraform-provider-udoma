/*
Udoma

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CaseComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaseComment{}

// CaseComment Comment represents a single user comment on a case
type CaseComment struct {
	// Unique and immutable ID attribute of the entity that is generated when  the instance is created. The ID is unique within the system accross all accounts and it can be used to reference the entity in other entities  or to retrieve it from the backend.
	Id string `json:"id"`
	// The date and time the entity was created
	CreatedAt int64 `json:"created_at"`
	// The date and time the entity was last updated
	UpdatedAt int64         `json:"updated_at"`
	AuthorRef UserReference `json:"author_ref"`
	// Indicates if the comment has been deleted
	Deleted *bool `json:"deleted,omitempty"`
	// List of paries that should have access to the comment
	Visibility []CaseParty `json:"visibility,omitempty"`
	Content    string      `json:"content"`
	// list of attachments that should be linked to the comment
	Attachments []Attachment `json:"attachments,omitempty"`
}

type _CaseComment CaseComment

// NewCaseComment instantiates a new CaseComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaseComment(id string, createdAt int64, updatedAt int64, authorRef UserReference, content string) *CaseComment {
	this := CaseComment{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.AuthorRef = authorRef
	this.Content = content
	return &this
}

// NewCaseCommentWithDefaults instantiates a new CaseComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaseCommentWithDefaults() *CaseComment {
	this := CaseComment{}
	return &this
}

// GetId returns the Id field value
func (o *CaseComment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CaseComment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CaseComment) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CaseComment) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CaseComment) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CaseComment) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CaseComment) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CaseComment) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CaseComment) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetAuthorRef returns the AuthorRef field value
func (o *CaseComment) GetAuthorRef() UserReference {
	if o == nil {
		var ret UserReference
		return ret
	}

	return o.AuthorRef
}

// GetAuthorRefOk returns a tuple with the AuthorRef field value
// and a boolean to check if the value has been set.
func (o *CaseComment) GetAuthorRefOk() (*UserReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorRef, true
}

// SetAuthorRef sets field value
func (o *CaseComment) SetAuthorRef(v UserReference) {
	o.AuthorRef = v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *CaseComment) GetDeleted() bool {
	if o == nil || IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseComment) GetDeletedOk() (*bool, bool) {
	if o == nil || IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *CaseComment) HasDeleted() bool {
	if o != nil && !IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *CaseComment) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *CaseComment) GetVisibility() []CaseParty {
	if o == nil || IsNil(o.Visibility) {
		var ret []CaseParty
		return ret
	}
	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseComment) GetVisibilityOk() ([]CaseParty, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *CaseComment) HasVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given []CaseParty and assigns it to the Visibility field.
func (o *CaseComment) SetVisibility(v []CaseParty) {
	o.Visibility = v
}

// GetContent returns the Content field value
func (o *CaseComment) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *CaseComment) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *CaseComment) SetContent(v string) {
	o.Content = v
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *CaseComment) GetAttachments() []Attachment {
	if o == nil || IsNil(o.Attachments) {
		var ret []Attachment
		return ret
	}
	return o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseComment) GetAttachmentsOk() ([]Attachment, bool) {
	if o == nil || IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *CaseComment) HasAttachments() bool {
	if o != nil && !IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given []Attachment and assigns it to the Attachments field.
func (o *CaseComment) SetAttachments(v []Attachment) {
	o.Attachments = v
}

func (o CaseComment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaseComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["author_ref"] = o.AuthorRef
	if !IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	toSerialize["content"] = o.Content
	if !IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	return toSerialize, nil
}

func (o *CaseComment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"author_ref",
		"content",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCaseComment := _CaseComment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCaseComment)

	if err != nil {
		return err
	}

	*o = CaseComment(varCaseComment)

	return err
}

type NullableCaseComment struct {
	value *CaseComment
	isSet bool
}

func (v NullableCaseComment) Get() *CaseComment {
	return v.value
}

func (v *NullableCaseComment) Set(val *CaseComment) {
	v.value = val
	v.isSet = true
}

func (v NullableCaseComment) IsSet() bool {
	return v.isSet
}

func (v *NullableCaseComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaseComment(val *CaseComment) *NullableCaseComment {
	return &NullableCaseComment{value: val, isSet: true}
}

func (v NullableCaseComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaseComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
