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

// checks if the PropertyContact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyContact{}

// PropertyContact A contact that is linked to a property. Can either be made  available to tenants, or just used as internal reference.
type PropertyContact struct {
	// unique ID of this contact
	Id *string `json:"id,omitempty"`
	// human-friendly label, describing the relationship and context of the contact
	Label *string `json:"label,omitempty"`
	// identifies if the contact is visible to the tenants or on a public form
	Public  *bool            `json:"public,omitempty"`
	Contact *ServiceProvider `json:"contact,omitempty"`
	// the ID of the property for which the contact is maintained
	PropertyRef *string `json:"property_ref,omitempty"`
}

// NewPropertyContact instantiates a new PropertyContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyContact() *PropertyContact {
	this := PropertyContact{}
	return &this
}

// NewPropertyContactWithDefaults instantiates a new PropertyContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyContactWithDefaults() *PropertyContact {
	this := PropertyContact{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PropertyContact) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyContact) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PropertyContact) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PropertyContact) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PropertyContact) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyContact) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PropertyContact) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PropertyContact) SetLabel(v string) {
	o.Label = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *PropertyContact) GetPublic() bool {
	if o == nil || IsNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyContact) GetPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *PropertyContact) HasPublic() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *PropertyContact) SetPublic(v bool) {
	o.Public = &v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *PropertyContact) GetContact() ServiceProvider {
	if o == nil || IsNil(o.Contact) {
		var ret ServiceProvider
		return ret
	}
	return *o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyContact) GetContactOk() (*ServiceProvider, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *PropertyContact) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given ServiceProvider and assigns it to the Contact field.
func (o *PropertyContact) SetContact(v ServiceProvider) {
	o.Contact = &v
}

// GetPropertyRef returns the PropertyRef field value if set, zero value otherwise.
func (o *PropertyContact) GetPropertyRef() string {
	if o == nil || IsNil(o.PropertyRef) {
		var ret string
		return ret
	}
	return *o.PropertyRef
}

// GetPropertyRefOk returns a tuple with the PropertyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyContact) GetPropertyRefOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyRef) {
		return nil, false
	}
	return o.PropertyRef, true
}

// HasPropertyRef returns a boolean if a field has been set.
func (o *PropertyContact) HasPropertyRef() bool {
	if o != nil && !IsNil(o.PropertyRef) {
		return true
	}

	return false
}

// SetPropertyRef gets a reference to the given string and assigns it to the PropertyRef field.
func (o *PropertyContact) SetPropertyRef(v string) {
	o.PropertyRef = &v
}

func (o PropertyContact) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyContact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.PropertyRef) {
		toSerialize["property_ref"] = o.PropertyRef
	}
	return toSerialize, nil
}

type NullablePropertyContact struct {
	value *PropertyContact
	isSet bool
}

func (v NullablePropertyContact) Get() *PropertyContact {
	return v.value
}

func (v *NullablePropertyContact) Set(val *PropertyContact) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyContact) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyContact(val *PropertyContact) *NullablePropertyContact {
	return &NullablePropertyContact{value: val, isSet: true}
}

func (v NullablePropertyContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
