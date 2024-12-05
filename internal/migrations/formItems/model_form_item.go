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

// checks if the FormItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormItem{}

// FormItem an item in a form group, referencing another entity in the form
type FormItem struct {
	// the ID of the entity that will be referenced
	RefId   string       `json:"ref_id"`
	RefType FormItemType `json:"ref_type"`
}

type _FormItem FormItem

// NewFormItem instantiates a new FormItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormItem(refId string, refType FormItemType) *FormItem {
	this := FormItem{}
	this.RefId = refId
	this.RefType = refType
	return &this
}

// NewFormItemWithDefaults instantiates a new FormItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormItemWithDefaults() *FormItem {
	this := FormItem{}
	return &this
}

// GetRefId returns the RefId field value
func (o *FormItem) GetRefId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value
// and a boolean to check if the value has been set.
func (o *FormItem) GetRefIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefId, true
}

// SetRefId sets field value
func (o *FormItem) SetRefId(v string) {
	o.RefId = v
}

// GetRefType returns the RefType field value
func (o *FormItem) GetRefType() FormItemType {
	if o == nil {
		var ret FormItemType
		return ret
	}

	return o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value
// and a boolean to check if the value has been set.
func (o *FormItem) GetRefTypeOk() (*FormItemType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefType, true
}

// SetRefType sets field value
func (o *FormItem) SetRefType(v FormItemType) {
	o.RefType = v
}

func (o FormItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ref_id"] = o.RefId
	toSerialize["ref_type"] = o.RefType
	return toSerialize, nil
}

func (o *FormItem) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ref_id",
		"ref_type",
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

	varFormItem := _FormItem{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFormItem)

	if err != nil {
		return err
	}

	*o = FormItem(varFormItem)

	return err
}

type NullableFormItem struct {
	value *FormItem
	isSet bool
}

func (v NullableFormItem) Get() *FormItem {
	return v.value
}

func (v *NullableFormItem) Set(val *FormItem) {
	v.value = val
	v.isSet = true
}

func (v NullableFormItem) IsSet() bool {
	return v.isSet
}

func (v *NullableFormItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormItem(val *FormItem) *NullableFormItem {
	return &NullableFormItem{value: val, isSet: true}
}

func (v NullableFormItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}