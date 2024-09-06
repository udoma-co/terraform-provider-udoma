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

// checks if the Tenancy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tenancy{}

// Tenancy A tenancy is a contract between a tenant and a property manager.
type Tenancy struct {
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
	// The timestamp of when the tenancy has started
	StartDate int64 `json:"start_date"`
	// The timestamp of when the tenancy has ended or is scheduled to end (required depending on the duration_type)
	EndDate     *int64      `json:"end_date,omitempty"`
	RentDetails RentDetails `json:"rent_details"`
	// Options to extend the contract if it's fixed term.
	ExtensionOptions []int64 `json:"extension_options,omitempty"`
	// The ID of the property that is being rented out
	PropertyRef  string              `json:"property_ref"`
	DurationType TenancyDurationEnum `json:"duration_type"`
	CurrentRent  RentData            `json:"current_rent"`
	// The IDs of the tenants
	Tenants []Tenant `json:"tenants"`
}

type _Tenancy Tenancy

// NewTenancy instantiates a new Tenancy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenancy(id string, createdAt int64, updatedAt int64, startDate int64, rentDetails RentDetails, propertyRef string, durationType TenancyDurationEnum, currentRent RentData, tenants []Tenant) *Tenancy {
	this := Tenancy{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.StartDate = startDate
	this.RentDetails = rentDetails
	this.PropertyRef = propertyRef
	this.DurationType = durationType
	this.CurrentRent = currentRent
	this.Tenants = tenants
	return &this
}

// NewTenancyWithDefaults instantiates a new Tenancy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenancyWithDefaults() *Tenancy {
	this := Tenancy{}
	return &this
}

// GetId returns the Id field value
func (o *Tenancy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Tenancy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Tenancy) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Tenancy) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Tenancy) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Tenancy) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Tenancy) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Tenancy) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Tenancy) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *Tenancy) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenancy) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *Tenancy) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *Tenancy) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetExternalSource returns the ExternalSource field value if set, zero value otherwise.
func (o *Tenancy) GetExternalSource() string {
	if o == nil || IsNil(o.ExternalSource) {
		var ret string
		return ret
	}
	return *o.ExternalSource
}

// GetExternalSourceOk returns a tuple with the ExternalSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenancy) GetExternalSourceOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalSource) {
		return nil, false
	}
	return o.ExternalSource, true
}

// HasExternalSource returns a boolean if a field has been set.
func (o *Tenancy) HasExternalSource() bool {
	if o != nil && !IsNil(o.ExternalSource) {
		return true
	}

	return false
}

// SetExternalSource gets a reference to the given string and assigns it to the ExternalSource field.
func (o *Tenancy) SetExternalSource(v string) {
	o.ExternalSource = &v
}

// GetStartDate returns the StartDate field value
func (o *Tenancy) GetStartDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *Tenancy) GetStartDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *Tenancy) SetStartDate(v int64) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *Tenancy) GetEndDate() int64 {
	if o == nil || IsNil(o.EndDate) {
		var ret int64
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenancy) GetEndDateOk() (*int64, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *Tenancy) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given int64 and assigns it to the EndDate field.
func (o *Tenancy) SetEndDate(v int64) {
	o.EndDate = &v
}

// GetRentDetails returns the RentDetails field value
func (o *Tenancy) GetRentDetails() RentDetails {
	if o == nil {
		var ret RentDetails
		return ret
	}

	return o.RentDetails
}

// GetRentDetailsOk returns a tuple with the RentDetails field value
// and a boolean to check if the value has been set.
func (o *Tenancy) GetRentDetailsOk() (*RentDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RentDetails, true
}

// SetRentDetails sets field value
func (o *Tenancy) SetRentDetails(v RentDetails) {
	o.RentDetails = v
}

// GetExtensionOptions returns the ExtensionOptions field value if set, zero value otherwise.
func (o *Tenancy) GetExtensionOptions() []int64 {
	if o == nil || IsNil(o.ExtensionOptions) {
		var ret []int64
		return ret
	}
	return o.ExtensionOptions
}

// GetExtensionOptionsOk returns a tuple with the ExtensionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenancy) GetExtensionOptionsOk() ([]int64, bool) {
	if o == nil || IsNil(o.ExtensionOptions) {
		return nil, false
	}
	return o.ExtensionOptions, true
}

// HasExtensionOptions returns a boolean if a field has been set.
func (o *Tenancy) HasExtensionOptions() bool {
	if o != nil && !IsNil(o.ExtensionOptions) {
		return true
	}

	return false
}

// SetExtensionOptions gets a reference to the given []int64 and assigns it to the ExtensionOptions field.
func (o *Tenancy) SetExtensionOptions(v []int64) {
	o.ExtensionOptions = v
}

// GetPropertyRef returns the PropertyRef field value
func (o *Tenancy) GetPropertyRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PropertyRef
}

// GetPropertyRefOk returns a tuple with the PropertyRef field value
// and a boolean to check if the value has been set.
func (o *Tenancy) GetPropertyRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertyRef, true
}

// SetPropertyRef sets field value
func (o *Tenancy) SetPropertyRef(v string) {
	o.PropertyRef = v
}

// GetDurationType returns the DurationType field value
func (o *Tenancy) GetDurationType() TenancyDurationEnum {
	if o == nil {
		var ret TenancyDurationEnum
		return ret
	}

	return o.DurationType
}

// GetDurationTypeOk returns a tuple with the DurationType field value
// and a boolean to check if the value has been set.
func (o *Tenancy) GetDurationTypeOk() (*TenancyDurationEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DurationType, true
}

// SetDurationType sets field value
func (o *Tenancy) SetDurationType(v TenancyDurationEnum) {
	o.DurationType = v
}

// GetCurrentRent returns the CurrentRent field value
func (o *Tenancy) GetCurrentRent() RentData {
	if o == nil {
		var ret RentData
		return ret
	}

	return o.CurrentRent
}

// GetCurrentRentOk returns a tuple with the CurrentRent field value
// and a boolean to check if the value has been set.
func (o *Tenancy) GetCurrentRentOk() (*RentData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentRent, true
}

// SetCurrentRent sets field value
func (o *Tenancy) SetCurrentRent(v RentData) {
	o.CurrentRent = v
}

// GetTenants returns the Tenants field value
func (o *Tenancy) GetTenants() []Tenant {
	if o == nil {
		var ret []Tenant
		return ret
	}

	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value
// and a boolean to check if the value has been set.
func (o *Tenancy) GetTenantsOk() ([]Tenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenants, true
}

// SetTenants sets field value
func (o *Tenancy) SetTenants(v []Tenant) {
	o.Tenants = v
}

func (o Tenancy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tenancy) ToMap() (map[string]interface{}, error) {
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
	toSerialize["start_date"] = o.StartDate
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	toSerialize["rent_details"] = o.RentDetails
	if !IsNil(o.ExtensionOptions) {
		toSerialize["extension_options"] = o.ExtensionOptions
	}
	toSerialize["property_ref"] = o.PropertyRef
	toSerialize["duration_type"] = o.DurationType
	toSerialize["current_rent"] = o.CurrentRent
	toSerialize["tenants"] = o.Tenants
	return toSerialize, nil
}

func (o *Tenancy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"start_date",
		"rent_details",
		"property_ref",
		"duration_type",
		"current_rent",
		"tenants",
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

	varTenancy := _Tenancy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTenancy)

	if err != nil {
		return err
	}

	*o = Tenancy(varTenancy)

	return err
}

type NullableTenancy struct {
	value *Tenancy
	isSet bool
}

func (v NullableTenancy) Get() *Tenancy {
	return v.value
}

func (v *NullableTenancy) Set(val *Tenancy) {
	v.value = val
	v.isSet = true
}

func (v NullableTenancy) IsSet() bool {
	return v.isSet
}

func (v *NullableTenancy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenancy(val *Tenancy) *NullableTenancy {
	return &NullableTenancy{value: val, isSet: true}
}

func (v NullableTenancy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenancy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
