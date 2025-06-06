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

// checks if the Correspondence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Correspondence{}

// Correspondence struct for Correspondence
type Correspondence struct {
	// Unique and immutable ID attribute of the entity that is generated when the instance is created. The ID is unique within the system accross all accounts and it can be used to reference the entity in other entities or to retrieve it from the backend.
	Id string `json:"id"`
	// The date and time the entity was created
	CreatedAt int64 `json:"created_at"`
	// The date and time the entity was last updated
	UpdatedAt int64 `json:"updated_at"`
	// Unique, immutable, alpha-numeric random code that is generated by the backend and can be used to access the entity via a publicly available URL. The code is unique within the system accross all accounts and it can be also be used to reference the entity in other entities or to retrieve it from the backend.
	Code string `json:"code"`
	// A descriptive display name
	DisplayName string `json:"display_name"`
	// A reference to the document generation that is being shared with the tenant
	DocumentRef string `json:"document_ref"`
	// A reference to the pdf attachment that'll be used for the correspondence
	AttachmentRef *string             `json:"attachment_ref,omitempty"`
	Recipient     NullableContactData `json:"recipient"`
	// Whether the document has been seen by the tenant or not(true if AccessRecord exists for the correspondence)
	Seen bool `json:"seen"`
	// Whether the Correspondence has been archived and shouldn't appear in query requests
	Archived bool `json:"archived"`
	// Whether the Correspondence has been cancelled
	Cancelled *bool `json:"cancelled,omitempty"`
}

type _Correspondence Correspondence

// NewCorrespondence instantiates a new Correspondence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorrespondence(id string, createdAt int64, updatedAt int64, code string, displayName string, documentRef string, recipient NullableContactData, seen bool, archived bool) *Correspondence {
	this := Correspondence{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Code = code
	this.DisplayName = displayName
	this.DocumentRef = documentRef
	this.Recipient = recipient
	this.Seen = seen
	this.Archived = archived
	return &this
}

// NewCorrespondenceWithDefaults instantiates a new Correspondence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorrespondenceWithDefaults() *Correspondence {
	this := Correspondence{}
	return &this
}

// GetId returns the Id field value
func (o *Correspondence) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Correspondence) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Correspondence) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Correspondence) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Correspondence) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Correspondence) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Correspondence) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Correspondence) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Correspondence) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetCode returns the Code field value
func (o *Correspondence) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *Correspondence) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *Correspondence) SetCode(v string) {
	o.Code = v
}

// GetDisplayName returns the DisplayName field value
func (o *Correspondence) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *Correspondence) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *Correspondence) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDocumentRef returns the DocumentRef field value
func (o *Correspondence) GetDocumentRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentRef
}

// GetDocumentRefOk returns a tuple with the DocumentRef field value
// and a boolean to check if the value has been set.
func (o *Correspondence) GetDocumentRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentRef, true
}

// SetDocumentRef sets field value
func (o *Correspondence) SetDocumentRef(v string) {
	o.DocumentRef = v
}

// GetAttachmentRef returns the AttachmentRef field value if set, zero value otherwise.
func (o *Correspondence) GetAttachmentRef() string {
	if o == nil || IsNil(o.AttachmentRef) {
		var ret string
		return ret
	}
	return *o.AttachmentRef
}

// GetAttachmentRefOk returns a tuple with the AttachmentRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Correspondence) GetAttachmentRefOk() (*string, bool) {
	if o == nil || IsNil(o.AttachmentRef) {
		return nil, false
	}
	return o.AttachmentRef, true
}

// HasAttachmentRef returns a boolean if a field has been set.
func (o *Correspondence) HasAttachmentRef() bool {
	if o != nil && !IsNil(o.AttachmentRef) {
		return true
	}

	return false
}

// SetAttachmentRef gets a reference to the given string and assigns it to the AttachmentRef field.
func (o *Correspondence) SetAttachmentRef(v string) {
	o.AttachmentRef = &v
}

// GetRecipient returns the Recipient field value
// If the value is explicit nil, the zero value for ContactData will be returned
func (o *Correspondence) GetRecipient() ContactData {
	if o == nil || o.Recipient.Get() == nil {
		var ret ContactData
		return ret
	}

	return *o.Recipient.Get()
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Correspondence) GetRecipientOk() (*ContactData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Recipient.Get(), o.Recipient.IsSet()
}

// SetRecipient sets field value
func (o *Correspondence) SetRecipient(v ContactData) {
	o.Recipient.Set(&v)
}

// GetSeen returns the Seen field value
func (o *Correspondence) GetSeen() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Seen
}

// GetSeenOk returns a tuple with the Seen field value
// and a boolean to check if the value has been set.
func (o *Correspondence) GetSeenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seen, true
}

// SetSeen sets field value
func (o *Correspondence) SetSeen(v bool) {
	o.Seen = v
}

// GetArchived returns the Archived field value
func (o *Correspondence) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *Correspondence) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *Correspondence) SetArchived(v bool) {
	o.Archived = v
}

// GetCancelled returns the Cancelled field value if set, zero value otherwise.
func (o *Correspondence) GetCancelled() bool {
	if o == nil || IsNil(o.Cancelled) {
		var ret bool
		return ret
	}
	return *o.Cancelled
}

// GetCancelledOk returns a tuple with the Cancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Correspondence) GetCancelledOk() (*bool, bool) {
	if o == nil || IsNil(o.Cancelled) {
		return nil, false
	}
	return o.Cancelled, true
}

// HasCancelled returns a boolean if a field has been set.
func (o *Correspondence) HasCancelled() bool {
	if o != nil && !IsNil(o.Cancelled) {
		return true
	}

	return false
}

// SetCancelled gets a reference to the given bool and assigns it to the Cancelled field.
func (o *Correspondence) SetCancelled(v bool) {
	o.Cancelled = &v
}

func (o Correspondence) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Correspondence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["code"] = o.Code
	toSerialize["display_name"] = o.DisplayName
	toSerialize["document_ref"] = o.DocumentRef
	if !IsNil(o.AttachmentRef) {
		toSerialize["attachment_ref"] = o.AttachmentRef
	}
	toSerialize["recipient"] = o.Recipient.Get()
	toSerialize["seen"] = o.Seen
	toSerialize["archived"] = o.Archived
	if !IsNil(o.Cancelled) {
		toSerialize["cancelled"] = o.Cancelled
	}
	return toSerialize, nil
}

func (o *Correspondence) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"code",
		"display_name",
		"document_ref",
		"recipient",
		"seen",
		"archived",
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

	varCorrespondence := _Correspondence{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCorrespondence)

	if err != nil {
		return err
	}

	*o = Correspondence(varCorrespondence)

	return err
}

type NullableCorrespondence struct {
	value *Correspondence
	isSet bool
}

func (v NullableCorrespondence) Get() *Correspondence {
	return v.value
}

func (v *NullableCorrespondence) Set(val *Correspondence) {
	v.value = val
	v.isSet = true
}

func (v NullableCorrespondence) IsSet() bool {
	return v.isSet
}

func (v *NullableCorrespondence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorrespondence(val *Correspondence) *NullableCorrespondence {
	return &NullableCorrespondence{value: val, isSet: true}
}

func (v NullableCorrespondence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorrespondence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
