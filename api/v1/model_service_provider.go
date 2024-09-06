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

// checks if the ServiceProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceProvider{}

// ServiceProvider Service providers are used by property managers to keep a reference of companies they work with in the respective category and subcategory
type ServiceProvider struct {
	// Unique and immutable ID attribute of the entity that is generated when the instance is created. The ID is unique within the system accross all accounts and it can be used to reference the entity in other entities or to retrieve it from the backend.
	Id string `json:"id"`
	// The date and time the entity was created
	CreatedAt int64 `json:"created_at"`
	// The date and time the entity was last updated
	UpdatedAt int64 `json:"updated_at"`
	// Optional external ID, in case entity was created via backend integration
	ExternalId *string `json:"external_id,omitempty"`
	// Optional external source, in case entity was created via backend integration
	ExternalSource *string `json:"external_source,omitempty"`
	// The salutation to be used for the contant person (Miss, Mister)
	Title *string `json:"title,omitempty"`
	// optional first name of contact person
	FirstName *string `json:"first_name,omitempty"`
	// optional last name of contact person
	LastName *string `json:"last_name,omitempty"`
	// the name of the company
	Company *string `json:"company,omitempty"`
	// the email address at which the company can be contacted by the platform
	Email string `json:"email"`
	// a phone number at which the company can be reached
	PhoneNumber *string              `json:"phone_number,omitempty"`
	Category    *ServiceCategoryEnum `json:"category,omitempty"`
	Address     *Address             `json:"address,omitempty"`
	// Indicates whether the service provider has an active account with Udoma. If so, it will be possible to assign cases directly to them.
	Connected *bool `json:"connected,omitempty"`
}

type _ServiceProvider ServiceProvider

// NewServiceProvider instantiates a new ServiceProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceProvider(id string, createdAt int64, updatedAt int64, email string) *ServiceProvider {
	this := ServiceProvider{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Email = email
	return &this
}

// NewServiceProviderWithDefaults instantiates a new ServiceProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceProviderWithDefaults() *ServiceProvider {
	this := ServiceProvider{}
	return &this
}

// GetId returns the Id field value
func (o *ServiceProvider) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ServiceProvider) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ServiceProvider) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ServiceProvider) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ServiceProvider) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ServiceProvider) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ServiceProvider) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ServiceProvider) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ServiceProvider) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetExternalSource returns the ExternalSource field value if set, zero value otherwise.
func (o *ServiceProvider) GetExternalSource() string {
	if o == nil || IsNil(o.ExternalSource) {
		var ret string
		return ret
	}
	return *o.ExternalSource
}

// GetExternalSourceOk returns a tuple with the ExternalSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetExternalSourceOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalSource) {
		return nil, false
	}
	return o.ExternalSource, true
}

// HasExternalSource returns a boolean if a field has been set.
func (o *ServiceProvider) HasExternalSource() bool {
	if o != nil && !IsNil(o.ExternalSource) {
		return true
	}

	return false
}

// SetExternalSource gets a reference to the given string and assigns it to the ExternalSource field.
func (o *ServiceProvider) SetExternalSource(v string) {
	o.ExternalSource = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ServiceProvider) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ServiceProvider) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ServiceProvider) SetTitle(v string) {
	o.Title = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ServiceProvider) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ServiceProvider) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ServiceProvider) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ServiceProvider) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ServiceProvider) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ServiceProvider) SetLastName(v string) {
	o.LastName = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *ServiceProvider) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *ServiceProvider) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *ServiceProvider) SetCompany(v string) {
	o.Company = &v
}

// GetEmail returns the Email field value
func (o *ServiceProvider) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ServiceProvider) SetEmail(v string) {
	o.Email = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *ServiceProvider) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *ServiceProvider) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *ServiceProvider) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *ServiceProvider) GetCategory() ServiceCategoryEnum {
	if o == nil || IsNil(o.Category) {
		var ret ServiceCategoryEnum
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetCategoryOk() (*ServiceCategoryEnum, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *ServiceProvider) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given ServiceCategoryEnum and assigns it to the Category field.
func (o *ServiceProvider) SetCategory(v ServiceCategoryEnum) {
	o.Category = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ServiceProvider) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ServiceProvider) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *ServiceProvider) SetAddress(v Address) {
	o.Address = &v
}

// GetConnected returns the Connected field value if set, zero value otherwise.
func (o *ServiceProvider) GetConnected() bool {
	if o == nil || IsNil(o.Connected) {
		var ret bool
		return ret
	}
	return *o.Connected
}

// GetConnectedOk returns a tuple with the Connected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetConnectedOk() (*bool, bool) {
	if o == nil || IsNil(o.Connected) {
		return nil, false
	}
	return o.Connected, true
}

// HasConnected returns a boolean if a field has been set.
func (o *ServiceProvider) HasConnected() bool {
	if o != nil && !IsNil(o.Connected) {
		return true
	}

	return false
}

// SetConnected gets a reference to the given bool and assigns it to the Connected field.
func (o *ServiceProvider) SetConnected(v bool) {
	o.Connected = &v
}

func (o ServiceProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if !IsNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	if !IsNil(o.ExternalSource) {
		toSerialize["external_source"] = o.ExternalSource
	}
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
	toSerialize["email"] = o.Email
	if !IsNil(o.PhoneNumber) {
		toSerialize["phone_number"] = o.PhoneNumber
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Connected) {
		toSerialize["connected"] = o.Connected
	}
	return toSerialize, nil
}

func (o *ServiceProvider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"email",
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

	varServiceProvider := _ServiceProvider{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceProvider)

	if err != nil {
		return err
	}

	*o = ServiceProvider(varServiceProvider)

	return err
}

type NullableServiceProvider struct {
	value *ServiceProvider
	isSet bool
}

func (v NullableServiceProvider) Get() *ServiceProvider {
	return v.value
}

func (v *NullableServiceProvider) Set(val *ServiceProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProvider(val *ServiceProvider) *NullableServiceProvider {
	return &NullableServiceProvider{value: val, isSet: true}
}

func (v NullableServiceProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
