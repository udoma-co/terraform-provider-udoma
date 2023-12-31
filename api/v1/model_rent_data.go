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

// RentData Information about the rent of a property
type RentData struct {
	// the monthly rent amount
	Rent *float64 `json:"rent,omitempty"`
	// the amount of incidentals paid monthly
	Incidentals *float64 `json:"incidentals,omitempty"`
}

// NewRentData instantiates a new RentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRentData() *RentData {
	this := RentData{}
	return &this
}

// NewRentDataWithDefaults instantiates a new RentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRentDataWithDefaults() *RentData {
	this := RentData{}
	return &this
}

// GetRent returns the Rent field value if set, zero value otherwise.
func (o *RentData) GetRent() float64 {
	if o == nil || o.Rent == nil {
		var ret float64
		return ret
	}
	return *o.Rent
}

// GetRentOk returns a tuple with the Rent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RentData) GetRentOk() (*float64, bool) {
	if o == nil || o.Rent == nil {
		return nil, false
	}
	return o.Rent, true
}

// HasRent returns a boolean if a field has been set.
func (o *RentData) HasRent() bool {
	if o != nil && o.Rent != nil {
		return true
	}

	return false
}

// SetRent gets a reference to the given float64 and assigns it to the Rent field.
func (o *RentData) SetRent(v float64) {
	o.Rent = &v
}

// GetIncidentals returns the Incidentals field value if set, zero value otherwise.
func (o *RentData) GetIncidentals() float64 {
	if o == nil || o.Incidentals == nil {
		var ret float64
		return ret
	}
	return *o.Incidentals
}

// GetIncidentalsOk returns a tuple with the Incidentals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RentData) GetIncidentalsOk() (*float64, bool) {
	if o == nil || o.Incidentals == nil {
		return nil, false
	}
	return o.Incidentals, true
}

// HasIncidentals returns a boolean if a field has been set.
func (o *RentData) HasIncidentals() bool {
	if o != nil && o.Incidentals != nil {
		return true
	}

	return false
}

// SetIncidentals gets a reference to the given float64 and assigns it to the Incidentals field.
func (o *RentData) SetIncidentals(v float64) {
	o.Incidentals = &v
}

func (o RentData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rent != nil {
		toSerialize["rent"] = o.Rent
	}
	if o.Incidentals != nil {
		toSerialize["incidentals"] = o.Incidentals
	}
	return json.Marshal(toSerialize)
}

type NullableRentData struct {
	value *RentData
	isSet bool
}

func (v NullableRentData) Get() *RentData {
	return v.value
}

func (v *NullableRentData) Set(val *RentData) {
	v.value = val
	v.isSet = true
}

func (v NullableRentData) IsSet() bool {
	return v.isSet
}

func (v *NullableRentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRentData(val *RentData) *NullableRentData {
	return &NullableRentData{value: val, isSet: true}
}

func (v NullableRentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
