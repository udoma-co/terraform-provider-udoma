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

// AdditionalTenant additional tenants who will not be included in the contract
type AdditionalTenant struct {
	Name         *string `json:"name,omitempty"`
	DayOfBirth   *string `json:"day_of_birth,omitempty"`
	Relationship *string `json:"relationship,omitempty"`
}

// NewAdditionalTenant instantiates a new AdditionalTenant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalTenant() *AdditionalTenant {
	this := AdditionalTenant{}
	return &this
}

// NewAdditionalTenantWithDefaults instantiates a new AdditionalTenant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalTenantWithDefaults() *AdditionalTenant {
	this := AdditionalTenant{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdditionalTenant) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalTenant) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdditionalTenant) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdditionalTenant) SetName(v string) {
	o.Name = &v
}

// GetDayOfBirth returns the DayOfBirth field value if set, zero value otherwise.
func (o *AdditionalTenant) GetDayOfBirth() string {
	if o == nil || o.DayOfBirth == nil {
		var ret string
		return ret
	}
	return *o.DayOfBirth
}

// GetDayOfBirthOk returns a tuple with the DayOfBirth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalTenant) GetDayOfBirthOk() (*string, bool) {
	if o == nil || o.DayOfBirth == nil {
		return nil, false
	}
	return o.DayOfBirth, true
}

// HasDayOfBirth returns a boolean if a field has been set.
func (o *AdditionalTenant) HasDayOfBirth() bool {
	if o != nil && o.DayOfBirth != nil {
		return true
	}

	return false
}

// SetDayOfBirth gets a reference to the given string and assigns it to the DayOfBirth field.
func (o *AdditionalTenant) SetDayOfBirth(v string) {
	o.DayOfBirth = &v
}

// GetRelationship returns the Relationship field value if set, zero value otherwise.
func (o *AdditionalTenant) GetRelationship() string {
	if o == nil || o.Relationship == nil {
		var ret string
		return ret
	}
	return *o.Relationship
}

// GetRelationshipOk returns a tuple with the Relationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdditionalTenant) GetRelationshipOk() (*string, bool) {
	if o == nil || o.Relationship == nil {
		return nil, false
	}
	return o.Relationship, true
}

// HasRelationship returns a boolean if a field has been set.
func (o *AdditionalTenant) HasRelationship() bool {
	if o != nil && o.Relationship != nil {
		return true
	}

	return false
}

// SetRelationship gets a reference to the given string and assigns it to the Relationship field.
func (o *AdditionalTenant) SetRelationship(v string) {
	o.Relationship = &v
}

func (o AdditionalTenant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DayOfBirth != nil {
		toSerialize["day_of_birth"] = o.DayOfBirth
	}
	if o.Relationship != nil {
		toSerialize["relationship"] = o.Relationship
	}
	return json.Marshal(toSerialize)
}

type NullableAdditionalTenant struct {
	value *AdditionalTenant
	isSet bool
}

func (v NullableAdditionalTenant) Get() *AdditionalTenant {
	return v.value
}

func (v *NullableAdditionalTenant) Set(val *AdditionalTenant) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalTenant) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalTenant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalTenant(val *AdditionalTenant) *NullableAdditionalTenant {
	return &NullableAdditionalTenant{value: val, isSet: true}
}

func (v NullableAdditionalTenant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalTenant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
