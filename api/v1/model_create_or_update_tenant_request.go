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

// checks if the CreateOrUpdateTenantRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateTenantRequest{}

// CreateOrUpdateTenantRequest Request issued by a property manager to create a new tenant, who is not yet assigned to a property
type CreateOrUpdateTenantRequest struct {
	// The salutation to be used for the tenant (Miss, Mister, Family)
	Title *string `json:"title,omitempty"`
	// First name(s) of the tenant. In case of a company, this is the first name of the contact person within the company.
	FirstName string `json:"first_name"`
	// Last name of the tenant. In case of a company, this is the last name of the contact person within the company.
	LastName string `json:"last_name"`
	// Name of the company. Only used in case tenancy is between lessor and a company, otherwise this field is empty.
	Company *string `json:"company,omitempty"`
	// The email address of the tenant. Used for sending communication to them, as well as used for invitations.
	Email string `json:"email"`
	// The phone number of the tenant. Not directly used by the system, but can be used by managers and service providers to contact the tenant.
	PhoneNumber string `json:"phone_number"`
}

type _CreateOrUpdateTenantRequest CreateOrUpdateTenantRequest

// NewCreateOrUpdateTenantRequest instantiates a new CreateOrUpdateTenantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateTenantRequest(firstName string, lastName string, email string, phoneNumber string) *CreateOrUpdateTenantRequest {
	this := CreateOrUpdateTenantRequest{}
	this.FirstName = firstName
	this.LastName = lastName
	this.Email = email
	this.PhoneNumber = phoneNumber
	return &this
}

// NewCreateOrUpdateTenantRequestWithDefaults instantiates a new CreateOrUpdateTenantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateTenantRequestWithDefaults() *CreateOrUpdateTenantRequest {
	this := CreateOrUpdateTenantRequest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CreateOrUpdateTenantRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateTenantRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CreateOrUpdateTenantRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CreateOrUpdateTenantRequest) SetTitle(v string) {
	o.Title = &v
}

// GetFirstName returns the FirstName field value
func (o *CreateOrUpdateTenantRequest) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateTenantRequest) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *CreateOrUpdateTenantRequest) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *CreateOrUpdateTenantRequest) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateTenantRequest) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *CreateOrUpdateTenantRequest) SetLastName(v string) {
	o.LastName = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *CreateOrUpdateTenantRequest) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateTenantRequest) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *CreateOrUpdateTenantRequest) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *CreateOrUpdateTenantRequest) SetCompany(v string) {
	o.Company = &v
}

// GetEmail returns the Email field value
func (o *CreateOrUpdateTenantRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateTenantRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *CreateOrUpdateTenantRequest) SetEmail(v string) {
	o.Email = v
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *CreateOrUpdateTenantRequest) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateTenantRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *CreateOrUpdateTenantRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

func (o CreateOrUpdateTenantRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateTenantRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	toSerialize["first_name"] = o.FirstName
	toSerialize["last_name"] = o.LastName
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	toSerialize["email"] = o.Email
	toSerialize["phone_number"] = o.PhoneNumber
	return toSerialize, nil
}

func (o *CreateOrUpdateTenantRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"first_name",
		"last_name",
		"email",
		"phone_number",
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

	varCreateOrUpdateTenantRequest := _CreateOrUpdateTenantRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrUpdateTenantRequest)

	if err != nil {
		return err
	}

	*o = CreateOrUpdateTenantRequest(varCreateOrUpdateTenantRequest)

	return err
}

type NullableCreateOrUpdateTenantRequest struct {
	value *CreateOrUpdateTenantRequest
	isSet bool
}

func (v NullableCreateOrUpdateTenantRequest) Get() *CreateOrUpdateTenantRequest {
	return v.value
}

func (v *NullableCreateOrUpdateTenantRequest) Set(val *CreateOrUpdateTenantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateTenantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateTenantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateTenantRequest(val *CreateOrUpdateTenantRequest) *NullableCreateOrUpdateTenantRequest {
	return &NullableCreateOrUpdateTenantRequest{value: val, isSet: true}
}

func (v NullableCreateOrUpdateTenantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateTenantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
