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

// checks if the AccountBookingAllocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountBookingAllocation{}

// AccountBookingAllocation The allocation of (part of) the amount of a booking to a financial account
type AccountBookingAllocation struct {
	// The unique identifier of the financial account
	AccountRef string `json:"account_ref"`
	// The split amount that is assigned to the referenced account. The sum of all allocations must add up to the total amount of the booking.
	Amount float64 `json:"amount"`
	// A serialized version of dimensions containing account value for the first value and dimensions for the rest in downwards order. They should be separated with '.' Example - \"10000.20.001.02\", where 10000 is the account value, 20, 001, and 02 are dimension values for dimensions that belong to the account.
	FlatNumber *string `json:"flat_number,omitempty"`
	// The values of the dimensions that are assigned to the booking for the allocation
	Dimensions []AccountDimensionValueRef `json:"dimensions,omitempty"`
}

type _AccountBookingAllocation AccountBookingAllocation

// NewAccountBookingAllocation instantiates a new AccountBookingAllocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountBookingAllocation(accountRef string, amount float64) *AccountBookingAllocation {
	this := AccountBookingAllocation{}
	this.AccountRef = accountRef
	this.Amount = amount
	return &this
}

// NewAccountBookingAllocationWithDefaults instantiates a new AccountBookingAllocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountBookingAllocationWithDefaults() *AccountBookingAllocation {
	this := AccountBookingAllocation{}
	return &this
}

// GetAccountRef returns the AccountRef field value
func (o *AccountBookingAllocation) GetAccountRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountRef
}

// GetAccountRefOk returns a tuple with the AccountRef field value
// and a boolean to check if the value has been set.
func (o *AccountBookingAllocation) GetAccountRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountRef, true
}

// SetAccountRef sets field value
func (o *AccountBookingAllocation) SetAccountRef(v string) {
	o.AccountRef = v
}

// GetAmount returns the Amount field value
func (o *AccountBookingAllocation) GetAmount() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AccountBookingAllocation) GetAmountOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AccountBookingAllocation) SetAmount(v float64) {
	o.Amount = v
}

// GetFlatNumber returns the FlatNumber field value if set, zero value otherwise.
func (o *AccountBookingAllocation) GetFlatNumber() string {
	if o == nil || IsNil(o.FlatNumber) {
		var ret string
		return ret
	}
	return *o.FlatNumber
}

// GetFlatNumberOk returns a tuple with the FlatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBookingAllocation) GetFlatNumberOk() (*string, bool) {
	if o == nil || IsNil(o.FlatNumber) {
		return nil, false
	}
	return o.FlatNumber, true
}

// HasFlatNumber returns a boolean if a field has been set.
func (o *AccountBookingAllocation) HasFlatNumber() bool {
	if o != nil && !IsNil(o.FlatNumber) {
		return true
	}

	return false
}

// SetFlatNumber gets a reference to the given string and assigns it to the FlatNumber field.
func (o *AccountBookingAllocation) SetFlatNumber(v string) {
	o.FlatNumber = &v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *AccountBookingAllocation) GetDimensions() []AccountDimensionValueRef {
	if o == nil || IsNil(o.Dimensions) {
		var ret []AccountDimensionValueRef
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountBookingAllocation) GetDimensionsOk() ([]AccountDimensionValueRef, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *AccountBookingAllocation) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given []AccountDimensionValueRef and assigns it to the Dimensions field.
func (o *AccountBookingAllocation) SetDimensions(v []AccountDimensionValueRef) {
	o.Dimensions = v
}

func (o AccountBookingAllocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountBookingAllocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_ref"] = o.AccountRef
	toSerialize["amount"] = o.Amount
	if !IsNil(o.FlatNumber) {
		toSerialize["flat_number"] = o.FlatNumber
	}
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	return toSerialize, nil
}

func (o *AccountBookingAllocation) UnmarshalJSON(data []byte) (err error) {
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

	varAccountBookingAllocation := _AccountBookingAllocation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountBookingAllocation)

	if err != nil {
		return err
	}

	*o = AccountBookingAllocation(varAccountBookingAllocation)

	return err
}

type NullableAccountBookingAllocation struct {
	value *AccountBookingAllocation
	isSet bool
}

func (v NullableAccountBookingAllocation) Get() *AccountBookingAllocation {
	return v.value
}

func (v *NullableAccountBookingAllocation) Set(val *AccountBookingAllocation) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountBookingAllocation) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountBookingAllocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountBookingAllocation(val *AccountBookingAllocation) *NullableAccountBookingAllocation {
	return &NullableAccountBookingAllocation{value: val, isSet: true}
}

func (v NullableAccountBookingAllocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountBookingAllocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
