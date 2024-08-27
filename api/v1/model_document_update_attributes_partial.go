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

// checks if the DocumentUpdateAttributesPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentUpdateAttributesPartial{}

// DocumentUpdateAttributesPartial Attributes of a document that can be updated
type DocumentUpdateAttributesPartial struct {
	// The name of the document entry
	Name string `json:"name"`
	// Optional path that can be used to nest documents in folder like structures
	Path *string `json:"path,omitempty"`
	// Indicator, controlling whether tenants can see the document
	Public *bool `json:"public,omitempty"`
}

type _DocumentUpdateAttributesPartial DocumentUpdateAttributesPartial

// NewDocumentUpdateAttributesPartial instantiates a new DocumentUpdateAttributesPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentUpdateAttributesPartial(name string) *DocumentUpdateAttributesPartial {
	this := DocumentUpdateAttributesPartial{}
	this.Name = name
	return &this
}

// NewDocumentUpdateAttributesPartialWithDefaults instantiates a new DocumentUpdateAttributesPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentUpdateAttributesPartialWithDefaults() *DocumentUpdateAttributesPartial {
	this := DocumentUpdateAttributesPartial{}
	return &this
}

// GetName returns the Name field value
func (o *DocumentUpdateAttributesPartial) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DocumentUpdateAttributesPartial) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DocumentUpdateAttributesPartial) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *DocumentUpdateAttributesPartial) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentUpdateAttributesPartial) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *DocumentUpdateAttributesPartial) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *DocumentUpdateAttributesPartial) SetPath(v string) {
	o.Path = &v
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *DocumentUpdateAttributesPartial) GetPublic() bool {
	if o == nil || IsNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentUpdateAttributesPartial) GetPublicOk() (*bool, bool) {
	if o == nil || IsNil(o.Public) {
		return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *DocumentUpdateAttributesPartial) HasPublic() bool {
	if o != nil && !IsNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *DocumentUpdateAttributesPartial) SetPublic(v bool) {
	o.Public = &v
}

func (o DocumentUpdateAttributesPartial) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentUpdateAttributesPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	return toSerialize, nil
}

func (o *DocumentUpdateAttributesPartial) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varDocumentUpdateAttributesPartial := _DocumentUpdateAttributesPartial{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDocumentUpdateAttributesPartial)

	if err != nil {
		return err
	}

	*o = DocumentUpdateAttributesPartial(varDocumentUpdateAttributesPartial)

	return err
}

type NullableDocumentUpdateAttributesPartial struct {
	value *DocumentUpdateAttributesPartial
	isSet bool
}

func (v NullableDocumentUpdateAttributesPartial) Get() *DocumentUpdateAttributesPartial {
	return v.value
}

func (v *NullableDocumentUpdateAttributesPartial) Set(val *DocumentUpdateAttributesPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentUpdateAttributesPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentUpdateAttributesPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentUpdateAttributesPartial(val *DocumentUpdateAttributesPartial) *NullableDocumentUpdateAttributesPartial {
	return &NullableDocumentUpdateAttributesPartial{value: val, isSet: true}
}

func (v NullableDocumentUpdateAttributesPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentUpdateAttributesPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}