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

// checks if the BankAccountAttributesPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BankAccountAttributesPartial{}

// BankAccountAttributesPartial Attributes of a bank account that must be provided upon creation
type BankAccountAttributesPartial struct {
	// The name of the account holder (required)
	AccountHolder string `json:"account_holder"`
	// The IBAN of the bank account (required)
	Iban string `json:"iban"`
	// The BIC of the bank account (optional)
	Bic *string `json:"bic,omitempty"`
	// The name of the bank (optional)
	BankName *string `json:"bank_name,omitempty"`
	// A user friendly label, used to identify the account (optional)
	Description *string `json:"description,omitempty"`
}

type _BankAccountAttributesPartial BankAccountAttributesPartial

// NewBankAccountAttributesPartial instantiates a new BankAccountAttributesPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankAccountAttributesPartial(accountHolder string, iban string) *BankAccountAttributesPartial {
	this := BankAccountAttributesPartial{}
	this.AccountHolder = accountHolder
	this.Iban = iban
	return &this
}

// NewBankAccountAttributesPartialWithDefaults instantiates a new BankAccountAttributesPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankAccountAttributesPartialWithDefaults() *BankAccountAttributesPartial {
	this := BankAccountAttributesPartial{}
	return &this
}

// GetAccountHolder returns the AccountHolder field value
func (o *BankAccountAttributesPartial) GetAccountHolder() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountHolder
}

// GetAccountHolderOk returns a tuple with the AccountHolder field value
// and a boolean to check if the value has been set.
func (o *BankAccountAttributesPartial) GetAccountHolderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountHolder, true
}

// SetAccountHolder sets field value
func (o *BankAccountAttributesPartial) SetAccountHolder(v string) {
	o.AccountHolder = v
}

// GetIban returns the Iban field value
func (o *BankAccountAttributesPartial) GetIban() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iban
}

// GetIbanOk returns a tuple with the Iban field value
// and a boolean to check if the value has been set.
func (o *BankAccountAttributesPartial) GetIbanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iban, true
}

// SetIban sets field value
func (o *BankAccountAttributesPartial) SetIban(v string) {
	o.Iban = v
}

// GetBic returns the Bic field value if set, zero value otherwise.
func (o *BankAccountAttributesPartial) GetBic() string {
	if o == nil || IsNil(o.Bic) {
		var ret string
		return ret
	}
	return *o.Bic
}

// GetBicOk returns a tuple with the Bic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountAttributesPartial) GetBicOk() (*string, bool) {
	if o == nil || IsNil(o.Bic) {
		return nil, false
	}
	return o.Bic, true
}

// HasBic returns a boolean if a field has been set.
func (o *BankAccountAttributesPartial) HasBic() bool {
	if o != nil && !IsNil(o.Bic) {
		return true
	}

	return false
}

// SetBic gets a reference to the given string and assigns it to the Bic field.
func (o *BankAccountAttributesPartial) SetBic(v string) {
	o.Bic = &v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *BankAccountAttributesPartial) GetBankName() string {
	if o == nil || IsNil(o.BankName) {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountAttributesPartial) GetBankNameOk() (*string, bool) {
	if o == nil || IsNil(o.BankName) {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *BankAccountAttributesPartial) HasBankName() bool {
	if o != nil && !IsNil(o.BankName) {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *BankAccountAttributesPartial) SetBankName(v string) {
	o.BankName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BankAccountAttributesPartial) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccountAttributesPartial) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BankAccountAttributesPartial) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BankAccountAttributesPartial) SetDescription(v string) {
	o.Description = &v
}

func (o BankAccountAttributesPartial) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BankAccountAttributesPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_holder"] = o.AccountHolder
	toSerialize["iban"] = o.Iban
	if !IsNil(o.Bic) {
		toSerialize["bic"] = o.Bic
	}
	if !IsNil(o.BankName) {
		toSerialize["bank_name"] = o.BankName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

func (o *BankAccountAttributesPartial) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_holder",
		"iban",
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

	varBankAccountAttributesPartial := _BankAccountAttributesPartial{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBankAccountAttributesPartial)

	if err != nil {
		return err
	}

	*o = BankAccountAttributesPartial(varBankAccountAttributesPartial)

	return err
}

type NullableBankAccountAttributesPartial struct {
	value *BankAccountAttributesPartial
	isSet bool
}

func (v NullableBankAccountAttributesPartial) Get() *BankAccountAttributesPartial {
	return v.value
}

func (v *NullableBankAccountAttributesPartial) Set(val *BankAccountAttributesPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccountAttributesPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccountAttributesPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccountAttributesPartial(val *BankAccountAttributesPartial) *NullableBankAccountAttributesPartial {
	return &NullableBankAccountAttributesPartial{value: val, isSet: true}
}

func (v NullableBankAccountAttributesPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankAccountAttributesPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
