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

// checks if the TenancyAttributesPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenancyAttributesPartial{}

// TenancyAttributesPartial struct for TenancyAttributesPartial
type TenancyAttributesPartial struct {
	// The timestamp of when the tenancy has started
	StartDate int64 `json:"start_date"`
	// The timestamp of when the tenancy has ended or is scheduled to end (required depending on the duration_type)
	EndDate     *int64      `json:"end_date,omitempty"`
	RentDetails RentDetails `json:"rent_details"`
	// Options to extend the contract if it's fixed term.
	ExtensionOptions []int64 `json:"extension_options,omitempty"`
}

type _TenancyAttributesPartial TenancyAttributesPartial

// NewTenancyAttributesPartial instantiates a new TenancyAttributesPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenancyAttributesPartial(startDate int64, rentDetails RentDetails) *TenancyAttributesPartial {
	this := TenancyAttributesPartial{}
	this.StartDate = startDate
	this.RentDetails = rentDetails
	return &this
}

// NewTenancyAttributesPartialWithDefaults instantiates a new TenancyAttributesPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenancyAttributesPartialWithDefaults() *TenancyAttributesPartial {
	this := TenancyAttributesPartial{}
	return &this
}

// GetStartDate returns the StartDate field value
func (o *TenancyAttributesPartial) GetStartDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *TenancyAttributesPartial) GetStartDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *TenancyAttributesPartial) SetStartDate(v int64) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *TenancyAttributesPartial) GetEndDate() int64 {
	if o == nil || IsNil(o.EndDate) {
		var ret int64
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyAttributesPartial) GetEndDateOk() (*int64, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *TenancyAttributesPartial) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int64 and assigns it to the EndDate field.
func (o *TenancyAttributesPartial) SetEndDate(v int64) {
	o.EndDate = &v
}

// GetRentDetails returns the RentDetails field value
func (o *TenancyAttributesPartial) GetRentDetails() RentDetails {
	if o == nil {
		var ret RentDetails
		return ret
	}

	return o.RentDetails
}

// GetRentDetailsOk returns a tuple with the RentDetails field value
// and a boolean to check if the value has been set.
func (o *TenancyAttributesPartial) GetRentDetailsOk() (*RentDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RentDetails, true
}

// SetRentDetails sets field value
func (o *TenancyAttributesPartial) SetRentDetails(v RentDetails) {
	o.RentDetails = v
}

// GetExtensionOptions returns the ExtensionOptions field value if set, zero value otherwise.
func (o *TenancyAttributesPartial) GetExtensionOptions() []int64 {
	if o == nil || IsNil(o.ExtensionOptions) {
		var ret []int64
		return ret
	}
	return o.ExtensionOptions
}

// GetExtensionOptionsOk returns a tuple with the ExtensionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenancyAttributesPartial) GetExtensionOptionsOk() ([]int64, bool) {
	if o == nil || IsNil(o.ExtensionOptions) {
		return nil, false
	}
	return o.ExtensionOptions, true
}

// HasExtensionOptions returns a boolean if a field has been set.
func (o *TenancyAttributesPartial) HasExtensionOptions() bool {
	if o != nil && !IsNil(o.ExtensionOptions) {
		return true
	}

	return false
}

// SetExtensionOptions gets a reference to the given []int64 and assigns it to the ExtensionOptions field.
func (o *TenancyAttributesPartial) SetExtensionOptions(v []int64) {
	o.ExtensionOptions = v
}

func (o TenancyAttributesPartial) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenancyAttributesPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start_date"] = o.StartDate
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	toSerialize["rent_details"] = o.RentDetails
	if !IsNil(o.ExtensionOptions) {
		toSerialize["extension_options"] = o.ExtensionOptions
	}
	return toSerialize, nil
}

func (o *TenancyAttributesPartial) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"start_date",
		"rent_details",
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

	varTenancyAttributesPartial := _TenancyAttributesPartial{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTenancyAttributesPartial)

	if err != nil {
		return err
	}

	*o = TenancyAttributesPartial(varTenancyAttributesPartial)

	return err
}

type NullableTenancyAttributesPartial struct {
	value *TenancyAttributesPartial
	isSet bool
}

func (v NullableTenancyAttributesPartial) Get() *TenancyAttributesPartial {
	return v.value
}

func (v *NullableTenancyAttributesPartial) Set(val *TenancyAttributesPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableTenancyAttributesPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableTenancyAttributesPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenancyAttributesPartial(val *TenancyAttributesPartial) *NullableTenancyAttributesPartial {
	return &NullableTenancyAttributesPartial{value: val, isSet: true}
}

func (v NullableTenancyAttributesPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenancyAttributesPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
