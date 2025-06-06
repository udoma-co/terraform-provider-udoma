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

// checks if the CreateOrUpdateAccountBookingAllocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateAccountBookingAllocation{}

// CreateOrUpdateAccountBookingAllocation The allocation of (part of) the amount of a booking to a financial account
type CreateOrUpdateAccountBookingAllocation struct {
	// The unique identifier of the financial account
	AccountRef string `json:"account_ref"`
	// The split amount that is assigned to the referenced account. The sum of all allocations must add up to the total amount of the booking.
	Amount float64 `json:"amount"`
	// A serialized version of dimensions containing account value for the first value and dimensions for the rest in downwards order. They should be separated with '.' Example - \"10000.20.001.02\", where 10000 is the account value, 20, 001, and 02 are dimension values for dimensions that belong to the account.
	FlatNumber *string `json:"flat_number,omitempty"`
}

type _CreateOrUpdateAccountBookingAllocation CreateOrUpdateAccountBookingAllocation

// NewCreateOrUpdateAccountBookingAllocation instantiates a new CreateOrUpdateAccountBookingAllocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateAccountBookingAllocation(accountRef string, amount float64) *CreateOrUpdateAccountBookingAllocation {
	this := CreateOrUpdateAccountBookingAllocation{}
	this.AccountRef = accountRef
	this.Amount = amount
	return &this
}

// NewCreateOrUpdateAccountBookingAllocationWithDefaults instantiates a new CreateOrUpdateAccountBookingAllocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateAccountBookingAllocationWithDefaults() *CreateOrUpdateAccountBookingAllocation {
	this := CreateOrUpdateAccountBookingAllocation{}
	return &this
}

// GetAccountRef returns the AccountRef field value
func (o *CreateOrUpdateAccountBookingAllocation) GetAccountRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountRef
}

// GetAccountRefOk returns a tuple with the AccountRef field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAccountBookingAllocation) GetAccountRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountRef, true
}

// SetAccountRef sets field value
func (o *CreateOrUpdateAccountBookingAllocation) SetAccountRef(v string) {
	o.AccountRef = v
}

// GetAmount returns the Amount field value
func (o *CreateOrUpdateAccountBookingAllocation) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAccountBookingAllocation) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreateOrUpdateAccountBookingAllocation) SetAmount(v float64) {
	o.Amount = v
}

// GetFlatNumber returns the FlatNumber field value if set, zero value otherwise.
func (o *CreateOrUpdateAccountBookingAllocation) GetFlatNumber() string {
	if o == nil || IsNil(o.FlatNumber) {
		var ret string
		return ret
	}
	return *o.FlatNumber
}

// GetFlatNumberOk returns a tuple with the FlatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAccountBookingAllocation) GetFlatNumberOk() (*string, bool) {
	if o == nil || IsNil(o.FlatNumber) {
		return nil, false
	}
	return o.FlatNumber, true
}

// HasFlatNumber returns a boolean if a field has been set.
func (o *CreateOrUpdateAccountBookingAllocation) HasFlatNumber() bool {
	if o != nil && !IsNil(o.FlatNumber) {
		return true
	}

	return false
}

// SetFlatNumber gets a reference to the given string and assigns it to the FlatNumber field.
func (o *CreateOrUpdateAccountBookingAllocation) SetFlatNumber(v string) {
	o.FlatNumber = &v
}

func (o CreateOrUpdateAccountBookingAllocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateAccountBookingAllocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_ref"] = o.AccountRef
	toSerialize["amount"] = o.Amount
	if !IsNil(o.FlatNumber) {
		toSerialize["flat_number"] = o.FlatNumber
	}
	return toSerialize, nil
}

func (o *CreateOrUpdateAccountBookingAllocation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_ref",
		"amount",
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

	varCreateOrUpdateAccountBookingAllocation := _CreateOrUpdateAccountBookingAllocation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrUpdateAccountBookingAllocation)

	if err != nil {
		return err
	}

	*o = CreateOrUpdateAccountBookingAllocation(varCreateOrUpdateAccountBookingAllocation)

	return err
}

type NullableCreateOrUpdateAccountBookingAllocation struct {
	value *CreateOrUpdateAccountBookingAllocation
	isSet bool
}

func (v NullableCreateOrUpdateAccountBookingAllocation) Get() *CreateOrUpdateAccountBookingAllocation {
	return v.value
}

func (v *NullableCreateOrUpdateAccountBookingAllocation) Set(val *CreateOrUpdateAccountBookingAllocation) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateAccountBookingAllocation) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateAccountBookingAllocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateAccountBookingAllocation(val *CreateOrUpdateAccountBookingAllocation) *NullableCreateOrUpdateAccountBookingAllocation {
	return &NullableCreateOrUpdateAccountBookingAllocation{value: val, isSet: true}
}

func (v NullableCreateOrUpdateAccountBookingAllocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateAccountBookingAllocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
