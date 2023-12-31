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

// CreateTenantRequest Request issued by a property manager to create a new tenant, who is not yet  assigned to a property
type CreateTenantRequest struct {
	// The salutation to be used for the tenant (Miss, Mister, Family)
	Title *string `json:"title,omitempty"`
	// First name(s) of the tenant. In case of a company, this is the first name of the contact person within the company.
	FirstName *string `json:"first_name,omitempty"`
	// Last name of the tenant. In case of a company, this is the last name of the contact person within the company.
	LastName *string `json:"last_name,omitempty"`
	// Name of the company. Only used in case tenancy is between lessor and a company,  otherwise this field is empty.
	Company *string `json:"company,omitempty"`
	// The email address of the tenant. Used for sending communication to them, as well as used for invitations.
	Email *string `json:"email,omitempty"`
	// The phone number of the tenant. Not directly used by the system, but can be used by managers and service providers to contact the tenant.
	PhoneNumber *string `json:"phone_number,omitempty"`
}

// NewCreateTenantRequest instantiates a new CreateTenantRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTenantRequest() *CreateTenantRequest {
	this := CreateTenantRequest{}
	return &this
}

// NewCreateTenantRequestWithDefaults instantiates a new CreateTenantRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTenantRequestWithDefaults() *CreateTenantRequest {
	this := CreateTenantRequest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *CreateTenantRequest) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenantRequest) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *CreateTenantRequest) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *CreateTenantRequest) SetTitle(v string) {
	o.Title = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *CreateTenantRequest) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenantRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *CreateTenantRequest) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *CreateTenantRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *CreateTenantRequest) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenantRequest) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *CreateTenantRequest) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *CreateTenantRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *CreateTenantRequest) GetCompany() string {
	if o == nil || o.Company == nil {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenantRequest) GetCompanyOk() (*string, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *CreateTenantRequest) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *CreateTenantRequest) SetCompany(v string) {
	o.Company = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CreateTenantRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenantRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CreateTenantRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CreateTenantRequest) SetEmail(v string) {
	o.Email = &v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *CreateTenantRequest) GetPhoneNumber() string {
	if o == nil || o.PhoneNumber == nil {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTenantRequest) GetPhoneNumberOk() (*string, bool) {
	if o == nil || o.PhoneNumber == nil {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *CreateTenantRequest) HasPhoneNumber() bool {
	if o != nil && o.PhoneNumber != nil {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *CreateTenantRequest) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

func (o CreateTenantRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.PhoneNumber != nil {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTenantRequest struct {
	value *CreateTenantRequest
	isSet bool
}

func (v NullableCreateTenantRequest) Get() *CreateTenantRequest {
	return v.value
}

func (v *NullableCreateTenantRequest) Set(val *CreateTenantRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTenantRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTenantRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTenantRequest(val *CreateTenantRequest) *NullableCreateTenantRequest {
	return &NullableCreateTenantRequest{value: val, isSet: true}
}

func (v NullableCreateTenantRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTenantRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
