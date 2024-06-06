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

// checks if the RentDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RentDetails{}

// RentDetails Contract terms about the rent
type RentDetails struct {
	// Whether the tenancy is taxed or not
	Taxable  *bool         `json:"taxable,omitempty"`
	RentType *RentTypeEnum `json:"rent_type,omitempty"`
	// the deposit amount
	Deposit *float64 `json:"deposit,omitempty"`
	// the date the graduated rent expires at
	GraduatedLength *int64 `json:"graduated_length,omitempty"`
	// A list of date-value pairs, representing a map, since the generator doesn't allow for maps that don't have strings as keys.
	GraduatedUpdates []GraduatedUpdate `json:"graduated_updates,omitempty"`
}

// NewRentDetails instantiates a new RentDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRentDetails() *RentDetails {
	this := RentDetails{}
	return &this
}

// NewRentDetailsWithDefaults instantiates a new RentDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRentDetailsWithDefaults() *RentDetails {
	this := RentDetails{}
	return &this
}

// GetTaxable returns the Taxable field value if set, zero value otherwise.
func (o *RentDetails) GetTaxable() bool {
	if o == nil || IsNil(o.Taxable) {
		var ret bool
		return ret
	}
	return *o.Taxable
}

// GetTaxableOk returns a tuple with the Taxable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RentDetails) GetTaxableOk() (*bool, bool) {
	if o == nil || IsNil(o.Taxable) {
		return nil, false
	}
	return o.Taxable, true
}

// HasTaxable returns a boolean if a field has been set.
func (o *RentDetails) HasTaxable() bool {
	if o != nil && !IsNil(o.Taxable) {
		return true
	}

	return false
}

// SetTaxable gets a reference to the given bool and assigns it to the Taxable field.
func (o *RentDetails) SetTaxable(v bool) {
	o.Taxable = &v
}

// GetRentType returns the RentType field value if set, zero value otherwise.
func (o *RentDetails) GetRentType() RentTypeEnum {
	if o == nil || IsNil(o.RentType) {
		var ret RentTypeEnum
		return ret
	}
	return *o.RentType
}

// GetRentTypeOk returns a tuple with the RentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RentDetails) GetRentTypeOk() (*RentTypeEnum, bool) {
	if o == nil || IsNil(o.RentType) {
		return nil, false
	}
	return o.RentType, true
}

// HasRentType returns a boolean if a field has been set.
func (o *RentDetails) HasRentType() bool {
	if o != nil && !IsNil(o.RentType) {
		return true
	}

	return false
}

// SetRentType gets a reference to the given RentTypeEnum and assigns it to the RentType field.
func (o *RentDetails) SetRentType(v RentTypeEnum) {
	o.RentType = &v
}

// GetDeposit returns the Deposit field value if set, zero value otherwise.
func (o *RentDetails) GetDeposit() float64 {
	if o == nil || IsNil(o.Deposit) {
		var ret float64
		return ret
	}
	return *o.Deposit
}

// GetDepositOk returns a tuple with the Deposit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RentDetails) GetDepositOk() (*float64, bool) {
	if o == nil || IsNil(o.Deposit) {
		return nil, false
	}
	return o.Deposit, true
}

// HasDeposit returns a boolean if a field has been set.
func (o *RentDetails) HasDeposit() bool {
	if o != nil && !IsNil(o.Deposit) {
		return true
	}

	return false
}

// SetDeposit gets a reference to the given float64 and assigns it to the Deposit field.
func (o *RentDetails) SetDeposit(v float64) {
	o.Deposit = &v
}

// GetGraduatedLength returns the GraduatedLength field value if set, zero value otherwise.
func (o *RentDetails) GetGraduatedLength() int64 {
	if o == nil || IsNil(o.GraduatedLength) {
		var ret int64
		return ret
	}
	return *o.GraduatedLength
}

// GetGraduatedLengthOk returns a tuple with the GraduatedLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RentDetails) GetGraduatedLengthOk() (*int64, bool) {
	if o == nil || IsNil(o.GraduatedLength) {
		return nil, false
	}
	return o.GraduatedLength, true
}

// HasGraduatedLength returns a boolean if a field has been set.
func (o *RentDetails) HasGraduatedLength() bool {
	if o != nil && !IsNil(o.GraduatedLength) {
		return true
	}

	return false
}

// SetGraduatedLength gets a reference to the given int64 and assigns it to the GraduatedLength field.
func (o *RentDetails) SetGraduatedLength(v int64) {
	o.GraduatedLength = &v
}

// GetGraduatedUpdates returns the GraduatedUpdates field value if set, zero value otherwise.
func (o *RentDetails) GetGraduatedUpdates() []GraduatedUpdate {
	if o == nil || IsNil(o.GraduatedUpdates) {
		var ret []GraduatedUpdate
		return ret
	}
	return o.GraduatedUpdates
}

// GetGraduatedUpdatesOk returns a tuple with the GraduatedUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RentDetails) GetGraduatedUpdatesOk() ([]GraduatedUpdate, bool) {
	if o == nil || IsNil(o.GraduatedUpdates) {
		return nil, false
	}
	return o.GraduatedUpdates, true
}

// HasGraduatedUpdates returns a boolean if a field has been set.
func (o *RentDetails) HasGraduatedUpdates() bool {
	if o != nil && !IsNil(o.GraduatedUpdates) {
		return true
	}

	return false
}

// SetGraduatedUpdates gets a reference to the given []GraduatedUpdate and assigns it to the GraduatedUpdates field.
func (o *RentDetails) SetGraduatedUpdates(v []GraduatedUpdate) {
	o.GraduatedUpdates = v
}

func (o RentDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RentDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Taxable) {
		toSerialize["taxable"] = o.Taxable
	}
	if !IsNil(o.RentType) {
		toSerialize["rent_type"] = o.RentType
	}
	if !IsNil(o.Deposit) {
		toSerialize["deposit"] = o.Deposit
	}
	if !IsNil(o.GraduatedLength) {
		toSerialize["graduated_length"] = o.GraduatedLength
	}
	if !IsNil(o.GraduatedUpdates) {
		toSerialize["graduated_updates"] = o.GraduatedUpdates
	}
	return toSerialize, nil
}

type NullableRentDetails struct {
	value *RentDetails
	isSet bool
}

func (v NullableRentDetails) Get() *RentDetails {
	return v.value
}

func (v *NullableRentDetails) Set(val *RentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRentDetails(val *RentDetails) *NullableRentDetails {
	return &NullableRentDetails{value: val, isSet: true}
}

func (v NullableRentDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
