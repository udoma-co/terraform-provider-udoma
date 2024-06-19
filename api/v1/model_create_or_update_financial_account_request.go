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

// checks if the CreateOrUpdateFinancialAccountRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateFinancialAccountRequest{}

// CreateOrUpdateFinancialAccountRequest The data required to create a new financial account
type CreateOrUpdateFinancialAccountRequest struct {
	// The unique account number, manually set
	Number int32 `json:"number"`
	// The name of the account
	Name string `json:"name"`
	// The currency of the booking
	Currency string `json:"currency"`
	// The IDs of the dimensions that are assigned to the account
	Dimensions []string `json:"dimensions,omitempty"`
}

type _CreateOrUpdateFinancialAccountRequest CreateOrUpdateFinancialAccountRequest

// NewCreateOrUpdateFinancialAccountRequest instantiates a new CreateOrUpdateFinancialAccountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateFinancialAccountRequest(number int32, name string, currency string) *CreateOrUpdateFinancialAccountRequest {
	this := CreateOrUpdateFinancialAccountRequest{}
	this.Number = number
	this.Name = name
	this.Currency = currency
	return &this
}

// NewCreateOrUpdateFinancialAccountRequestWithDefaults instantiates a new CreateOrUpdateFinancialAccountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateFinancialAccountRequestWithDefaults() *CreateOrUpdateFinancialAccountRequest {
	this := CreateOrUpdateFinancialAccountRequest{}
	return &this
}

// GetNumber returns the Number field value
func (o *CreateOrUpdateFinancialAccountRequest) GetNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateFinancialAccountRequest) GetNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *CreateOrUpdateFinancialAccountRequest) SetNumber(v int32) {
	o.Number = v
}

// GetName returns the Name field value
func (o *CreateOrUpdateFinancialAccountRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateFinancialAccountRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrUpdateFinancialAccountRequest) SetName(v string) {
	o.Name = v
}

// GetCurrency returns the Currency field value
func (o *CreateOrUpdateFinancialAccountRequest) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateFinancialAccountRequest) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *CreateOrUpdateFinancialAccountRequest) SetCurrency(v string) {
	o.Currency = v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *CreateOrUpdateFinancialAccountRequest) GetDimensions() []string {
	if o == nil || IsNil(o.Dimensions) {
		var ret []string
		return ret
	}
	return o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateFinancialAccountRequest) GetDimensionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *CreateOrUpdateFinancialAccountRequest) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given []string and assigns it to the Dimensions field.
func (o *CreateOrUpdateFinancialAccountRequest) SetDimensions(v []string) {
	o.Dimensions = v
}

func (o CreateOrUpdateFinancialAccountRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateFinancialAccountRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["number"] = o.Number
	toSerialize["name"] = o.Name
	toSerialize["currency"] = o.Currency
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	return toSerialize, nil
}

func (o *CreateOrUpdateFinancialAccountRequest) UnmarshalJSON(data []byte) (err error) {
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

	varCreateOrUpdateFinancialAccountRequest := _CreateOrUpdateFinancialAccountRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrUpdateFinancialAccountRequest)

	if err != nil {
		return err
	}

	*o = CreateOrUpdateFinancialAccountRequest(varCreateOrUpdateFinancialAccountRequest)

	return err
}

type NullableCreateOrUpdateFinancialAccountRequest struct {
	value *CreateOrUpdateFinancialAccountRequest
	isSet bool
}

func (v NullableCreateOrUpdateFinancialAccountRequest) Get() *CreateOrUpdateFinancialAccountRequest {
	return v.value
}

func (v *NullableCreateOrUpdateFinancialAccountRequest) Set(val *CreateOrUpdateFinancialAccountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateFinancialAccountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateFinancialAccountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateFinancialAccountRequest(val *CreateOrUpdateFinancialAccountRequest) *NullableCreateOrUpdateFinancialAccountRequest {
	return &NullableCreateOrUpdateFinancialAccountRequest{value: val, isSet: true}
}

func (v NullableCreateOrUpdateFinancialAccountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateFinancialAccountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
