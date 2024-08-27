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

// checks if the PropertyOwnerAttributesPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyOwnerAttributesPartial{}

// PropertyOwnerAttributesPartial Request used to create a new property owner
type PropertyOwnerAttributesPartial struct {
	// the title of the owner
	Title *string `json:"title,omitempty"`
	// First name(s) of the property owners. In case of a company, this is the first name of the contact person within the company.
	FirstName *string `json:"first_name,omitempty"`
	// Last name of the property owners. In case of a company, this is the last name of the contact person within the company.
	LastName *string `json:"last_name,omitempty"`
	// Optional name of the company.
	Company *string `json:"company,omitempty"`
	// the email of the owner
	Email *string `json:"email,omitempty"`
	// the phone number of the owner
	PhoneNumber *string  `json:"phone_number,omitempty"`
	Address     *Address `json:"address,omitempty"`
	// the reference to the bank account of the owner
	BankAccountRef *string `json:"bank_account_ref,omitempty"`
}

// NewPropertyOwnerAttributesPartial instantiates a new PropertyOwnerAttributesPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyOwnerAttributesPartial() *PropertyOwnerAttributesPartial {
	this := PropertyOwnerAttributesPartial{}
	return &this
}

// NewPropertyOwnerAttributesPartialWithDefaults instantiates a new PropertyOwnerAttributesPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyOwnerAttributesPartialWithDefaults() *PropertyOwnerAttributesPartial {
	this := PropertyOwnerAttributesPartial{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PropertyOwnerAttributesPartial) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyOwnerAttributesPartial) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PropertyOwnerAttributesPartial) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PropertyOwnerAttributesPartial) SetTitle(v string) {
	o.Title = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *PropertyOwnerAttributesPartial) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyOwnerAttributesPartial) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *PropertyOwnerAttributesPartial) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *PropertyOwnerAttributesPartial) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *PropertyOwnerAttributesPartial) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyOwnerAttributesPartial) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *PropertyOwnerAttributesPartial) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *PropertyOwnerAttributesPartial) SetLastName(v string) {
	o.LastName = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *PropertyOwnerAttributesPartial) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyOwnerAttributesPartial) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *PropertyOwnerAttributesPartial) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *PropertyOwnerAttributesPartial) SetCompany(v string) {
	o.Company = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *PropertyOwnerAttributesPartial) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyOwnerAttributesPartial) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *PropertyOwnerAttributesPartial) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *PropertyOwnerAttributesPartial) SetEmail(v string) {
	o.Email = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *PropertyOwnerAttributesPartial) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyOwnerAttributesPartial) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *PropertyOwnerAttributesPartial) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *PropertyOwnerAttributesPartial) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *PropertyOwnerAttributesPartial) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyOwnerAttributesPartial) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *PropertyOwnerAttributesPartial) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *PropertyOwnerAttributesPartial) SetAddress(v Address) {
	o.Address = &v
}

// GetBankAccountRef returns the BankAccountRef field value if set, zero value otherwise.
func (o *PropertyOwnerAttributesPartial) GetBankAccountRef() string {
	if o == nil || IsNil(o.BankAccountRef) {
		var ret string
		return ret
	}
	return *o.BankAccountRef
}

// GetBankAccountRefOk returns a tuple with the BankAccountRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyOwnerAttributesPartial) GetBankAccountRefOk() (*string, bool) {
	if o == nil || IsNil(o.BankAccountRef) {
		return nil, false
	}
	return o.BankAccountRef, true
}

// HasBankAccountRef returns a boolean if a field has been set.
func (o *PropertyOwnerAttributesPartial) HasBankAccountRef() bool {
	if o != nil && !IsNil(o.BankAccountRef) {
		return true
	}

	return false
}

// SetBankAccountRef gets a reference to the given string and assigns it to the BankAccountRef field.
func (o *PropertyOwnerAttributesPartial) SetBankAccountRef(v string) {
	o.BankAccountRef = &v
}

func (o PropertyOwnerAttributesPartial) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyOwnerAttributesPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.BankAccountRef) {
		toSerialize["bank_account_ref"] = o.BankAccountRef
	}
	return toSerialize, nil
}

type NullablePropertyOwnerAttributesPartial struct {
	value *PropertyOwnerAttributesPartial
	isSet bool
}

func (v NullablePropertyOwnerAttributesPartial) Get() *PropertyOwnerAttributesPartial {
	return v.value
}

func (v *NullablePropertyOwnerAttributesPartial) Set(val *PropertyOwnerAttributesPartial) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyOwnerAttributesPartial) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyOwnerAttributesPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyOwnerAttributesPartial(val *PropertyOwnerAttributesPartial) *NullablePropertyOwnerAttributesPartial {
	return &NullablePropertyOwnerAttributesPartial{value: val, isSet: true}
}

func (v NullablePropertyOwnerAttributesPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyOwnerAttributesPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}