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

// checks if the CreateCorrespondenceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCorrespondenceRequest{}

// CreateCorrespondenceRequest The request body object, used to create a correspondence
type CreateCorrespondenceRequest struct {
	// A descriptive display name
	DisplayName string `json:"display_name"`
	// A reference to the document generation that is being shared with the tenant
	DocumentRef *string `json:"document_ref,omitempty"`
	// A reference to the pdf attachment that'll be used for the correspondence
	AttachmentRef *string              `json:"attachment_ref,omitempty"`
	Recipient     *ContactData         `json:"recipient,omitempty"`
	Email         *CorrespondenceEmail `json:"email,omitempty"`
}

type _CreateCorrespondenceRequest CreateCorrespondenceRequest

// NewCreateCorrespondenceRequest instantiates a new CreateCorrespondenceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCorrespondenceRequest(displayName string) *CreateCorrespondenceRequest {
	this := CreateCorrespondenceRequest{}
	this.DisplayName = displayName
	return &this
}

// NewCreateCorrespondenceRequestWithDefaults instantiates a new CreateCorrespondenceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCorrespondenceRequestWithDefaults() *CreateCorrespondenceRequest {
	this := CreateCorrespondenceRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *CreateCorrespondenceRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CreateCorrespondenceRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CreateCorrespondenceRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetDocumentRef returns the DocumentRef field value if set, zero value otherwise.
func (o *CreateCorrespondenceRequest) GetDocumentRef() string {
	if o == nil || IsNil(o.DocumentRef) {
		var ret string
		return ret
	}
	return *o.DocumentRef
}

// GetDocumentRefOk returns a tuple with the DocumentRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCorrespondenceRequest) GetDocumentRefOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentRef) {
		return nil, false
	}
	return o.DocumentRef, true
}

// HasDocumentRef returns a boolean if a field has been set.
func (o *CreateCorrespondenceRequest) HasDocumentRef() bool {
	if o != nil && !IsNil(o.DocumentRef) {
		return true
	}

	return false
}

// SetDocumentRef gets a reference to the given string and assigns it to the DocumentRef field.
func (o *CreateCorrespondenceRequest) SetDocumentRef(v string) {
	o.DocumentRef = &v
}

// GetAttachmentRef returns the AttachmentRef field value if set, zero value otherwise.
func (o *CreateCorrespondenceRequest) GetAttachmentRef() string {
	if o == nil || IsNil(o.AttachmentRef) {
		var ret string
		return ret
	}
	return *o.AttachmentRef
}

// GetAttachmentRefOk returns a tuple with the AttachmentRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCorrespondenceRequest) GetAttachmentRefOk() (*string, bool) {
	if o == nil || IsNil(o.AttachmentRef) {
		return nil, false
	}
	return o.AttachmentRef, true
}

// HasAttachmentRef returns a boolean if a field has been set.
func (o *CreateCorrespondenceRequest) HasAttachmentRef() bool {
	if o != nil && !IsNil(o.AttachmentRef) {
		return true
	}

	return false
}

// SetAttachmentRef gets a reference to the given string and assigns it to the AttachmentRef field.
func (o *CreateCorrespondenceRequest) SetAttachmentRef(v string) {
	o.AttachmentRef = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *CreateCorrespondenceRequest) GetRecipient() ContactData {
	if o == nil || IsNil(o.Recipient) {
		var ret ContactData
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCorrespondenceRequest) GetRecipientOk() (*ContactData, bool) {
	if o == nil || IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *CreateCorrespondenceRequest) HasRecipient() bool {
	if o != nil && !IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given ContactData and assigns it to the Recipient field.
func (o *CreateCorrespondenceRequest) SetRecipient(v ContactData) {
	o.Recipient = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CreateCorrespondenceRequest) GetEmail() CorrespondenceEmail {
	if o == nil || IsNil(o.Email) {
		var ret CorrespondenceEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCorrespondenceRequest) GetEmailOk() (*CorrespondenceEmail, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CreateCorrespondenceRequest) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given CorrespondenceEmail and assigns it to the Email field.
func (o *CreateCorrespondenceRequest) SetEmail(v CorrespondenceEmail) {
	o.Email = &v
}

func (o CreateCorrespondenceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCorrespondenceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["display_name"] = o.DisplayName
	if !IsNil(o.DocumentRef) {
		toSerialize["document_ref"] = o.DocumentRef
	}
	if !IsNil(o.AttachmentRef) {
		toSerialize["attachment_ref"] = o.AttachmentRef
	}
	if !IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return toSerialize, nil
}

func (o *CreateCorrespondenceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"display_name",
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

	varCreateCorrespondenceRequest := _CreateCorrespondenceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateCorrespondenceRequest)

	if err != nil {
		return err
	}

	*o = CreateCorrespondenceRequest(varCreateCorrespondenceRequest)

	return err
}

type NullableCreateCorrespondenceRequest struct {
	value *CreateCorrespondenceRequest
	isSet bool
}

func (v NullableCreateCorrespondenceRequest) Get() *CreateCorrespondenceRequest {
	return v.value
}

func (v *NullableCreateCorrespondenceRequest) Set(val *CreateCorrespondenceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCorrespondenceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCorrespondenceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCorrespondenceRequest(val *CreateCorrespondenceRequest) *NullableCreateCorrespondenceRequest {
	return &NullableCreateCorrespondenceRequest{value: val, isSet: true}
}

func (v NullableCreateCorrespondenceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCorrespondenceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
