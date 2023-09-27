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

// CreateOrUpdatePropertyContact Request used for creating a new contact for a property.
type CreateOrUpdatePropertyContact struct {
	// human-friendly label, describing the relationship and context of the contact
	Label *string `json:"label,omitempty"`
	// identifies if the contact is visible to the tenants or on a public form
	Public *bool `json:"public,omitempty"`
	// reference to the service provider that is linked as a contact
	Contact *string `json:"contact,omitempty"`
}

// NewCreateOrUpdatePropertyContact instantiates a new CreateOrUpdatePropertyContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdatePropertyContact() *CreateOrUpdatePropertyContact {
	this := CreateOrUpdatePropertyContact{}
	return &this
}

// NewCreateOrUpdatePropertyContactWithDefaults instantiates a new CreateOrUpdatePropertyContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdatePropertyContactWithDefaults() *CreateOrUpdatePropertyContact {
	this := CreateOrUpdatePropertyContact{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CreateOrUpdatePropertyContact) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdatePropertyContact) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CreateOrUpdatePropertyContact) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *CreateOrUpdatePropertyContact) SetLabel(v string) {
	o.Label = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *CreateOrUpdatePropertyContact) GetPublic() bool {
	if o == nil || o.Public == nil {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdatePropertyContact) GetPublicOk() (*bool, bool) {
	if o == nil || o.Public == nil {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *CreateOrUpdatePropertyContact) HasPublic() bool {
	if o != nil && o.Public != nil {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *CreateOrUpdatePropertyContact) SetPublic(v bool) {
	o.Public = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *CreateOrUpdatePropertyContact) GetContact() string {
	if o == nil || o.Contact == nil {
		var ret string
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdatePropertyContact) GetContactOk() (*string, bool) {
	if o == nil || o.Contact == nil {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *CreateOrUpdatePropertyContact) HasContact() bool {
	if o != nil && o.Contact != nil {
		return true
	}

	return false
}

// SetContact gets a reference to the given string and assigns it to the Contact field.
func (o *CreateOrUpdatePropertyContact) SetContact(v string) {
	o.Contact = &v
}

func (o CreateOrUpdatePropertyContact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Public != nil {
		toSerialize["public"] = o.Public
	}
	if o.Contact != nil {
		toSerialize["contact"] = o.Contact
	}
	return json.Marshal(toSerialize)
}

type NullableCreateOrUpdatePropertyContact struct {
	value *CreateOrUpdatePropertyContact
	isSet bool
}

func (v NullableCreateOrUpdatePropertyContact) Get() *CreateOrUpdatePropertyContact {
	return v.value
}

func (v *NullableCreateOrUpdatePropertyContact) Set(val *CreateOrUpdatePropertyContact) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdatePropertyContact) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdatePropertyContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdatePropertyContact(val *CreateOrUpdatePropertyContact) *NullableCreateOrUpdatePropertyContact {
	return &NullableCreateOrUpdatePropertyContact{value: val, isSet: true}
}

func (v NullableCreateOrUpdatePropertyContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdatePropertyContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
