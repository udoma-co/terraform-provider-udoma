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

// checks if the EntityNameAttributePartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityNameAttributePartial{}

// EntityNameAttributePartial struct for EntityNameAttributePartial
type EntityNameAttributePartial struct {
	// A unique, descriptive name of the entity. This attribute is used to  reference the entity in other entities or to retrieve it from the  backend. The name is unique within the account, but not the entire  system, and it's not generated by the backend, but provided by the  user when creating the entity.
	Name string `json:"name"`
}

type _EntityNameAttributePartial EntityNameAttributePartial

// NewEntityNameAttributePartial instantiates a new EntityNameAttributePartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityNameAttributePartial(name string) *EntityNameAttributePartial {
	this := EntityNameAttributePartial{}
	this.Name = name
	return &this
}

// NewEntityNameAttributePartialWithDefaults instantiates a new EntityNameAttributePartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityNameAttributePartialWithDefaults() *EntityNameAttributePartial {
	this := EntityNameAttributePartial{}
	return &this
}

// GetName returns the Name field value
func (o *EntityNameAttributePartial) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntityNameAttributePartial) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntityNameAttributePartial) SetName(v string) {
	o.Name = v
}

func (o EntityNameAttributePartial) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityNameAttributePartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *EntityNameAttributePartial) UnmarshalJSON(data []byte) (err error) {
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

	varEntityNameAttributePartial := _EntityNameAttributePartial{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEntityNameAttributePartial)

	if err != nil {
		return err
	}

	*o = EntityNameAttributePartial(varEntityNameAttributePartial)

	return err
}

type NullableEntityNameAttributePartial struct {
	value *EntityNameAttributePartial
	isSet bool
}

func (v NullableEntityNameAttributePartial) Get() *EntityNameAttributePartial {
	return v.value
}

func (v *NullableEntityNameAttributePartial) Set(val *EntityNameAttributePartial) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityNameAttributePartial) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityNameAttributePartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityNameAttributePartial(val *EntityNameAttributePartial) *NullableEntityNameAttributePartial {
	return &NullableEntityNameAttributePartial{value: val, isSet: true}
}

func (v NullableEntityNameAttributePartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityNameAttributePartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}