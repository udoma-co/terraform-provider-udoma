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

// checks if the EntityExternalReferenceAttributesPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityExternalReferenceAttributesPartial{}

// EntityExternalReferenceAttributesPartial Attributes of an entity that can be synchronized with an external system.
type EntityExternalReferenceAttributesPartial struct {
	// Optional external ID, in case entity was created via backend integration
	ExternalId *string `json:"external_id,omitempty"`
	// Optional external source, in case entity was created via backend integration
	ExternalSource *string `json:"external_source,omitempty"`
}

// NewEntityExternalReferenceAttributesPartial instantiates a new EntityExternalReferenceAttributesPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityExternalReferenceAttributesPartial() *EntityExternalReferenceAttributesPartial {
	this := EntityExternalReferenceAttributesPartial{}
	return &this
}

// NewEntityExternalReferenceAttributesPartialWithDefaults instantiates a new EntityExternalReferenceAttributesPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityExternalReferenceAttributesPartialWithDefaults() *EntityExternalReferenceAttributesPartial {
	this := EntityExternalReferenceAttributesPartial{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *EntityExternalReferenceAttributesPartial) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityExternalReferenceAttributesPartial) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *EntityExternalReferenceAttributesPartial) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *EntityExternalReferenceAttributesPartial) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetExternalSource returns the ExternalSource field value if set, zero value otherwise.
func (o *EntityExternalReferenceAttributesPartial) GetExternalSource() string {
	if o == nil || IsNil(o.ExternalSource) {
		var ret string
		return ret
	}
	return *o.ExternalSource
}

// GetExternalSourceOk returns a tuple with the ExternalSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityExternalReferenceAttributesPartial) GetExternalSourceOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalSource) {
		return nil, false
	}
	return o.ExternalSource, true
}

// HasExternalSource returns a boolean if a field has been set.
func (o *EntityExternalReferenceAttributesPartial) HasExternalSource() bool {
	if o != nil && !IsNil(o.ExternalSource) {
		return true
	}

	return false
}

// SetExternalSource gets a reference to the given string and assigns it to the ExternalSource field.
func (o *EntityExternalReferenceAttributesPartial) SetExternalSource(v string) {
	o.ExternalSource = &v
}

func (o EntityExternalReferenceAttributesPartial) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityExternalReferenceAttributesPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	if !IsNil(o.ExternalSource) {
		toSerialize["external_source"] = o.ExternalSource
	}
	return toSerialize, nil
}

type NullableEntityExternalReferenceAttributesPartial struct {
	value *EntityExternalReferenceAttributesPartial
	isSet bool
}

func (v NullableEntityExternalReferenceAttributesPartial) Get() *EntityExternalReferenceAttributesPartial {
	return v.value
}

func (v *NullableEntityExternalReferenceAttributesPartial) Set(val *EntityExternalReferenceAttributesPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityExternalReferenceAttributesPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityExternalReferenceAttributesPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityExternalReferenceAttributesPartial(val *EntityExternalReferenceAttributesPartial) *NullableEntityExternalReferenceAttributesPartial {
	return &NullableEntityExternalReferenceAttributesPartial{value: val, isSet: true}
}

func (v NullableEntityExternalReferenceAttributesPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityExternalReferenceAttributesPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
