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

// checks if the AdditionalRentCost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalRentCost{}

// AdditionalRentCost Additional costs that are not part of the rent, but are paid by the tenant, e.g. garage rent, parking space rent, etc.
type AdditionalRentCost struct {
	// The name of the cost, e.g. \"garage\"
	Name *string `json:"name,omitempty"`
	// The monthly amount due
	Amount *float64 `json:"amount,omitempty"`
}

// NewAdditionalRentCost instantiates a new AdditionalRentCost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalRentCost() *AdditionalRentCost {
	this := AdditionalRentCost{}
	return &this
}

// NewAdditionalRentCostWithDefaults instantiates a new AdditionalRentCost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalRentCostWithDefaults() *AdditionalRentCost {
	this := AdditionalRentCost{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdditionalRentCost) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalRentCost) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdditionalRentCost) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdditionalRentCost) SetName(v string) {
	o.Name = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *AdditionalRentCost) GetAmount() float64 {
	if o == nil || IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalRentCost) GetAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *AdditionalRentCost) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *AdditionalRentCost) SetAmount(v float64) {
	o.Amount = &v
}

func (o AdditionalRentCost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalRentCost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableAdditionalRentCost struct {
	value *AdditionalRentCost
	isSet bool
}

func (v NullableAdditionalRentCost) Get() *AdditionalRentCost {
	return v.value
}

func (v *NullableAdditionalRentCost) Set(val *AdditionalRentCost) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalRentCost) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalRentCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalRentCost(val *AdditionalRentCost) *NullableAdditionalRentCost {
	return &NullableAdditionalRentCost{value: val, isSet: true}
}

func (v NullableAdditionalRentCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalRentCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
