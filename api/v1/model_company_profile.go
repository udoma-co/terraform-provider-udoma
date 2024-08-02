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

// checks if the CompanyProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompanyProfile{}

// CompanyProfile Details about a company (property management, or service provider)
type CompanyProfile struct {
	Name    string   `json:"name"`
	Address *Address `json:"address,omitempty"`
	// The default email address at which the company should be contacted. This information can potentially be visible to other users of the platform.
	Email *string `json:"email,omitempty"`
	// The default phone number at which the company should be contacted. This information can potentially be visible to other users of the platform.
	PhoneNumber *string `json:"phone_number,omitempty"`
	// URL to the company website. This information can potentially be visible to other users of the platform.
	Website *string `json:"website,omitempty"`
}

type _CompanyProfile CompanyProfile

// NewCompanyProfile instantiates a new CompanyProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompanyProfile(name string) *CompanyProfile {
	this := CompanyProfile{}
	this.Name = name
	return &this
}

// NewCompanyProfileWithDefaults instantiates a new CompanyProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompanyProfileWithDefaults() *CompanyProfile {
	this := CompanyProfile{}
	return &this
}

// GetName returns the Name field value
func (o *CompanyProfile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CompanyProfile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CompanyProfile) SetName(v string) {
	o.Name = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *CompanyProfile) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfile) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *CompanyProfile) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *CompanyProfile) SetAddress(v Address) {
	o.Address = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CompanyProfile) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfile) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CompanyProfile) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CompanyProfile) SetEmail(v string) {
	o.Email = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *CompanyProfile) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfile) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *CompanyProfile) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *CompanyProfile) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *CompanyProfile) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompanyProfile) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *CompanyProfile) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *CompanyProfile) SetWebsite(v string) {
	o.Website = &v
}

func (o CompanyProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompanyProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	return toSerialize, nil
}

func (o *CompanyProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varCompanyProfile := _CompanyProfile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompanyProfile)

	if err != nil {
		return err
	}

	*o = CompanyProfile(varCompanyProfile)

	return err
}

type NullableCompanyProfile struct {
	value *CompanyProfile
	isSet bool
}

func (v NullableCompanyProfile) Get() *CompanyProfile {
	return v.value
}

func (v *NullableCompanyProfile) Set(val *CompanyProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableCompanyProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableCompanyProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompanyProfile(val *CompanyProfile) *NullableCompanyProfile {
	return &NullableCompanyProfile{value: val, isSet: true}
}

func (v NullableCompanyProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompanyProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
