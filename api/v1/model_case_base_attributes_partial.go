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

// checks if the CaseBaseAttributesPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaseBaseAttributesPartial{}

// CaseBaseAttributesPartial struct for CaseBaseAttributesPartial
type CaseBaseAttributesPartial struct {
	// Reference to property
	PropertyRef     *string  `json:"property_ref,omitempty"`
	PropertyAddress *Address `json:"property_address,omitempty"`
	// Input provided by the user when updating the case as a key-value map
	Data map[string]interface{} `json:"data"`
}

type _CaseBaseAttributesPartial CaseBaseAttributesPartial

// NewCaseBaseAttributesPartial instantiates a new CaseBaseAttributesPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaseBaseAttributesPartial(data map[string]interface{}) *CaseBaseAttributesPartial {
	this := CaseBaseAttributesPartial{}
	this.Data = data
	return &this
}

// NewCaseBaseAttributesPartialWithDefaults instantiates a new CaseBaseAttributesPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaseBaseAttributesPartialWithDefaults() *CaseBaseAttributesPartial {
	this := CaseBaseAttributesPartial{}
	return &this
}

// GetPropertyRef returns the PropertyRef field value if set, zero value otherwise.
func (o *CaseBaseAttributesPartial) GetPropertyRef() string {
	if o == nil || IsNil(o.PropertyRef) {
		var ret string
		return ret
	}
	return *o.PropertyRef
}

// GetPropertyRefOk returns a tuple with the PropertyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseBaseAttributesPartial) GetPropertyRefOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyRef) {
		return nil, false
	}
	return o.PropertyRef, true
}

// HasPropertyRef returns a boolean if a field has been set.
func (o *CaseBaseAttributesPartial) HasPropertyRef() bool {
	if o != nil && !IsNil(o.PropertyRef) {
		return true
	}

	return false
}

// SetPropertyRef gets a reference to the given string and assigns it to the PropertyRef field.
func (o *CaseBaseAttributesPartial) SetPropertyRef(v string) {
	o.PropertyRef = &v
}

// GetPropertyAddress returns the PropertyAddress field value if set, zero value otherwise.
func (o *CaseBaseAttributesPartial) GetPropertyAddress() Address {
	if o == nil || IsNil(o.PropertyAddress) {
		var ret Address
		return ret
	}
	return *o.PropertyAddress
}

// GetPropertyAddressOk returns a tuple with the PropertyAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseBaseAttributesPartial) GetPropertyAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.PropertyAddress) {
		return nil, false
	}
	return o.PropertyAddress, true
}

// HasPropertyAddress returns a boolean if a field has been set.
func (o *CaseBaseAttributesPartial) HasPropertyAddress() bool {
	if o != nil && !IsNil(o.PropertyAddress) {
		return true
	}

	return false
}

// SetPropertyAddress gets a reference to the given Address and assigns it to the PropertyAddress field.
func (o *CaseBaseAttributesPartial) SetPropertyAddress(v Address) {
	o.PropertyAddress = &v
}

// GetData returns the Data field value
func (o *CaseBaseAttributesPartial) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CaseBaseAttributesPartial) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CaseBaseAttributesPartial) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o CaseBaseAttributesPartial) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaseBaseAttributesPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyRef) {
		toSerialize["property_ref"] = o.PropertyRef
	}
	if !IsNil(o.PropertyAddress) {
		toSerialize["property_address"] = o.PropertyAddress
	}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CaseBaseAttributesPartial) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varCaseBaseAttributesPartial := _CaseBaseAttributesPartial{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCaseBaseAttributesPartial)

	if err != nil {
		return err
	}

	*o = CaseBaseAttributesPartial(varCaseBaseAttributesPartial)

	return err
}

type NullableCaseBaseAttributesPartial struct {
	value *CaseBaseAttributesPartial
	isSet bool
}

func (v NullableCaseBaseAttributesPartial) Get() *CaseBaseAttributesPartial {
	return v.value
}

func (v *NullableCaseBaseAttributesPartial) Set(val *CaseBaseAttributesPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableCaseBaseAttributesPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableCaseBaseAttributesPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaseBaseAttributesPartial(val *CaseBaseAttributesPartial) *NullableCaseBaseAttributesPartial {
	return &NullableCaseBaseAttributesPartial{value: val, isSet: true}
}

func (v NullableCaseBaseAttributesPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaseBaseAttributesPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
