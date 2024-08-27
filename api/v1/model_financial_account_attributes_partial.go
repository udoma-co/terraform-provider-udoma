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

// checks if the FinancialAccountAttributesPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialAccountAttributesPartial{}

// FinancialAccountAttributesPartial Attributes of a financial account that can be updated
type FinancialAccountAttributesPartial struct {
	// The unique account number, manually set
	Number int32 `json:"number"`
	// The name of the account
	Name string `json:"name"`
	// The currency of the booking
	Currency string `json:"currency"`
}

type _FinancialAccountAttributesPartial FinancialAccountAttributesPartial

// NewFinancialAccountAttributesPartial instantiates a new FinancialAccountAttributesPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialAccountAttributesPartial(number int32, name string, currency string) *FinancialAccountAttributesPartial {
	this := FinancialAccountAttributesPartial{}
	this.Number = number
	this.Name = name
	this.Currency = currency
	return &this
}

// NewFinancialAccountAttributesPartialWithDefaults instantiates a new FinancialAccountAttributesPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialAccountAttributesPartialWithDefaults() *FinancialAccountAttributesPartial {
	this := FinancialAccountAttributesPartial{}
	return &this
}

// GetNumber returns the Number field value
func (o *FinancialAccountAttributesPartial) GetNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountAttributesPartial) GetNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *FinancialAccountAttributesPartial) SetNumber(v int32) {
	o.Number = v
}

// GetName returns the Name field value
func (o *FinancialAccountAttributesPartial) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountAttributesPartial) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FinancialAccountAttributesPartial) SetName(v string) {
	o.Name = v
}

// GetCurrency returns the Currency field value
func (o *FinancialAccountAttributesPartial) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *FinancialAccountAttributesPartial) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *FinancialAccountAttributesPartial) SetCurrency(v string) {
	o.Currency = v
}

func (o FinancialAccountAttributesPartial) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialAccountAttributesPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["number"] = o.Number
	toSerialize["name"] = o.Name
	toSerialize["currency"] = o.Currency
	return toSerialize, nil
}

func (o *FinancialAccountAttributesPartial) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"number",
		"name",
		"currency",
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

	varFinancialAccountAttributesPartial := _FinancialAccountAttributesPartial{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFinancialAccountAttributesPartial)

	if err != nil {
		return err
	}

	*o = FinancialAccountAttributesPartial(varFinancialAccountAttributesPartial)

	return err
}

type NullableFinancialAccountAttributesPartial struct {
	value *FinancialAccountAttributesPartial
	isSet bool
}

func (v NullableFinancialAccountAttributesPartial) Get() *FinancialAccountAttributesPartial {
	return v.value
}

func (v *NullableFinancialAccountAttributesPartial) Set(val *FinancialAccountAttributesPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialAccountAttributesPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialAccountAttributesPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialAccountAttributesPartial(val *FinancialAccountAttributesPartial) *NullableFinancialAccountAttributesPartial {
	return &NullableFinancialAccountAttributesPartial{value: val, isSet: true}
}

func (v NullableFinancialAccountAttributesPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialAccountAttributesPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}