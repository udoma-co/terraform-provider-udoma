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

// BankAccount Detailed information about a bank account
type BankAccount struct {
	// The ID of the bank account (auto-generated)
	Id *string `json:"id,omitempty"`
	// optional external ID, in case acount was created via backend integration
	ExternalId *string `json:"external_id,omitempty"`
	// optional external source, in case acount was created via backend integration
	ExternalSource *string `json:"external_source,omitempty"`
	// The timestamp of when the account was created
	CreatedAt *int64 `json:"created_at,omitempty"`
	// The timestamp of when the account was last updated
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// The name of the account holder (required)
	AccountHolder *string `json:"account_holder,omitempty"`
	// The IBAN of the bank account (required)
	Iban *string `json:"iban,omitempty"`
	// The BIC of the bank account (optional)
	Bic *string `json:"bic,omitempty"`
	// The name of the bank (optional)
	BankName *string `json:"bank_name,omitempty"`
	// An optional user friendly label, used to identify the account (optional)
	Description *string `json:"description,omitempty"`
}

// NewBankAccount instantiates a new BankAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankAccount() *BankAccount {
	this := BankAccount{}
	return &this
}

// NewBankAccountWithDefaults instantiates a new BankAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankAccountWithDefaults() *BankAccount {
	this := BankAccount{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BankAccount) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BankAccount) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BankAccount) SetId(v string) {
	o.Id = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *BankAccount) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *BankAccount) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *BankAccount) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetExternalSource returns the ExternalSource field value if set, zero value otherwise.
func (o *BankAccount) GetExternalSource() string {
	if o == nil || o.ExternalSource == nil {
		var ret string
		return ret
	}
	return *o.ExternalSource
}

// GetExternalSourceOk returns a tuple with the ExternalSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetExternalSourceOk() (*string, bool) {
	if o == nil || o.ExternalSource == nil {
		return nil, false
	}
	return o.ExternalSource, true
}

// HasExternalSource returns a boolean if a field has been set.
func (o *BankAccount) HasExternalSource() bool {
	if o != nil && o.ExternalSource != nil {
		return true
	}

	return false
}

// SetExternalSource gets a reference to the given string and assigns it to the ExternalSource field.
func (o *BankAccount) SetExternalSource(v string) {
	o.ExternalSource = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BankAccount) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BankAccount) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *BankAccount) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *BankAccount) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *BankAccount) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *BankAccount) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetAccountHolder returns the AccountHolder field value if set, zero value otherwise.
func (o *BankAccount) GetAccountHolder() string {
	if o == nil || o.AccountHolder == nil {
		var ret string
		return ret
	}
	return *o.AccountHolder
}

// GetAccountHolderOk returns a tuple with the AccountHolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetAccountHolderOk() (*string, bool) {
	if o == nil || o.AccountHolder == nil {
		return nil, false
	}
	return o.AccountHolder, true
}

// HasAccountHolder returns a boolean if a field has been set.
func (o *BankAccount) HasAccountHolder() bool {
	if o != nil && o.AccountHolder != nil {
		return true
	}

	return false
}

// SetAccountHolder gets a reference to the given string and assigns it to the AccountHolder field.
func (o *BankAccount) SetAccountHolder(v string) {
	o.AccountHolder = &v
}

// GetIban returns the Iban field value if set, zero value otherwise.
func (o *BankAccount) GetIban() string {
	if o == nil || o.Iban == nil {
		var ret string
		return ret
	}
	return *o.Iban
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetIbanOk() (*string, bool) {
	if o == nil || o.Iban == nil {
		return nil, false
	}
	return o.Iban, true
}

// HasIban returns a boolean if a field has been set.
func (o *BankAccount) HasIban() bool {
	if o != nil && o.Iban != nil {
		return true
	}

	return false
}

// SetIban gets a reference to the given string and assigns it to the Iban field.
func (o *BankAccount) SetIban(v string) {
	o.Iban = &v
}

// GetBic returns the Bic field value if set, zero value otherwise.
func (o *BankAccount) GetBic() string {
	if o == nil || o.Bic == nil {
		var ret string
		return ret
	}
	return *o.Bic
}

// GetBicOk returns a tuple with the Bic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetBicOk() (*string, bool) {
	if o == nil || o.Bic == nil {
		return nil, false
	}
	return o.Bic, true
}

// HasBic returns a boolean if a field has been set.
func (o *BankAccount) HasBic() bool {
	if o != nil && o.Bic != nil {
		return true
	}

	return false
}

// SetBic gets a reference to the given string and assigns it to the Bic field.
func (o *BankAccount) SetBic(v string) {
	o.Bic = &v
}

// GetBankName returns the BankName field value if set, zero value otherwise.
func (o *BankAccount) GetBankName() string {
	if o == nil || o.BankName == nil {
		var ret string
		return ret
	}
	return *o.BankName
}

// GetBankNameOk returns a tuple with the BankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetBankNameOk() (*string, bool) {
	if o == nil || o.BankName == nil {
		return nil, false
	}
	return o.BankName, true
}

// HasBankName returns a boolean if a field has been set.
func (o *BankAccount) HasBankName() bool {
	if o != nil && o.BankName != nil {
		return true
	}

	return false
}

// SetBankName gets a reference to the given string and assigns it to the BankName field.
func (o *BankAccount) SetBankName(v string) {
	o.BankName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BankAccount) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BankAccount) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BankAccount) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BankAccount) SetDescription(v string) {
	o.Description = &v
}

func (o BankAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if o.ExternalSource != nil {
		toSerialize["external_source"] = o.ExternalSource
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.AccountHolder != nil {
		toSerialize["account_holder"] = o.AccountHolder
	}
	if o.Iban != nil {
		toSerialize["iban"] = o.Iban
	}
	if o.Bic != nil {
		toSerialize["bic"] = o.Bic
	}
	if o.BankName != nil {
		toSerialize["bank_name"] = o.BankName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableBankAccount struct {
	value *BankAccount
	isSet bool
}

func (v NullableBankAccount) Get() *BankAccount {
	return v.value
}

func (v *NullableBankAccount) Set(val *BankAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableBankAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableBankAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankAccount(val *BankAccount) *NullableBankAccount {
	return &NullableBankAccount{value: val, isSet: true}
}

func (v NullableBankAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
