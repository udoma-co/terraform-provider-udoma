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

// checks if the AccountDimensionValueAttributesPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountDimensionValueAttributesPartial{}

// AccountDimensionValueAttributesPartial The attributes of a value of a dimension
type AccountDimensionValueAttributesPartial struct {
	// The unique ID this this value (typically the ID of the referenced  entity that is mapped, e.g. property ID). This value must be unique for the given dimension, regardless of the parent dimension value.
	Id string `json:"id"`
	// The numeric value under the given dimension. This value is unique for the given dimension. If a parent dimension reference is set,  then the value is only unique for the parent dimension value and  may be repeated for other parent dimension values.
	Value *int32 `json:"value,omitempty"`
	// Optional ID for the parent dimension value. If set, the value can only be set, when the parent dimension was set to the given ID.
	ParentDimensionId *string `json:"parent_dimension_id,omitempty"`
}

type _AccountDimensionValueAttributesPartial AccountDimensionValueAttributesPartial

// NewAccountDimensionValueAttributesPartial instantiates a new AccountDimensionValueAttributesPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountDimensionValueAttributesPartial(id string) *AccountDimensionValueAttributesPartial {
	this := AccountDimensionValueAttributesPartial{}
	this.Id = id
	return &this
}

// NewAccountDimensionValueAttributesPartialWithDefaults instantiates a new AccountDimensionValueAttributesPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountDimensionValueAttributesPartialWithDefaults() *AccountDimensionValueAttributesPartial {
	this := AccountDimensionValueAttributesPartial{}
	return &this
}

// GetId returns the Id field value
func (o *AccountDimensionValueAttributesPartial) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountDimensionValueAttributesPartial) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountDimensionValueAttributesPartial) SetId(v string) {
	o.Id = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AccountDimensionValueAttributesPartial) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDimensionValueAttributesPartial) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AccountDimensionValueAttributesPartial) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *AccountDimensionValueAttributesPartial) SetValue(v int32) {
	o.Value = &v
}

// GetParentDimensionId returns the ParentDimensionId field value if set, zero value otherwise.
func (o *AccountDimensionValueAttributesPartial) GetParentDimensionId() string {
	if o == nil || IsNil(o.ParentDimensionId) {
		var ret string
		return ret
	}
	return *o.ParentDimensionId
}

// GetParentDimensionIdOk returns a tuple with the ParentDimensionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDimensionValueAttributesPartial) GetParentDimensionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentDimensionId) {
		return nil, false
	}
	return o.ParentDimensionId, true
}

// HasParentDimensionId returns a boolean if a field has been set.
func (o *AccountDimensionValueAttributesPartial) HasParentDimensionId() bool {
	if o != nil && !IsNil(o.ParentDimensionId) {
		return true
	}

	return false
}

// SetParentDimensionId gets a reference to the given string and assigns it to the ParentDimensionId field.
func (o *AccountDimensionValueAttributesPartial) SetParentDimensionId(v string) {
	o.ParentDimensionId = &v
}

func (o AccountDimensionValueAttributesPartial) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountDimensionValueAttributesPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.ParentDimensionId) {
		toSerialize["parent_dimension_id"] = o.ParentDimensionId
	}
	return toSerialize, nil
}

func (o *AccountDimensionValueAttributesPartial) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varAccountDimensionValueAttributesPartial := _AccountDimensionValueAttributesPartial{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountDimensionValueAttributesPartial)

	if err != nil {
		return err
	}

	*o = AccountDimensionValueAttributesPartial(varAccountDimensionValueAttributesPartial)

	return err
}

type NullableAccountDimensionValueAttributesPartial struct {
	value *AccountDimensionValueAttributesPartial
	isSet bool
}

func (v NullableAccountDimensionValueAttributesPartial) Get() *AccountDimensionValueAttributesPartial {
	return v.value
}

func (v *NullableAccountDimensionValueAttributesPartial) Set(val *AccountDimensionValueAttributesPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountDimensionValueAttributesPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountDimensionValueAttributesPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountDimensionValueAttributesPartial(val *AccountDimensionValueAttributesPartial) *NullableAccountDimensionValueAttributesPartial {
	return &NullableAccountDimensionValueAttributesPartial{value: val, isSet: true}
}

func (v NullableAccountDimensionValueAttributesPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountDimensionValueAttributesPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
