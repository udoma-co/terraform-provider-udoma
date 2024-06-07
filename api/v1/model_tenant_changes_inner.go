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

// checks if the TenantChangesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantChangesInner{}

// TenantChangesInner struct for TenantChangesInner
type TenantChangesInner struct {
	Tenant *ContactData            `json:"tenant,omitempty"`
	Action *TenantChangeActionEnum `json:"action,omitempty"`
}

// NewTenantChangesInner instantiates a new TenantChangesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantChangesInner() *TenantChangesInner {
	this := TenantChangesInner{}
	return &this
}

// NewTenantChangesInnerWithDefaults instantiates a new TenantChangesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantChangesInnerWithDefaults() *TenantChangesInner {
	this := TenantChangesInner{}
	return &this
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *TenantChangesInner) GetTenant() ContactData {
	if o == nil || IsNil(o.Tenant) {
		var ret ContactData
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantChangesInner) GetTenantOk() (*ContactData, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *TenantChangesInner) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given ContactData and assigns it to the Tenant field.
func (o *TenantChangesInner) SetTenant(v ContactData) {
	o.Tenant = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *TenantChangesInner) GetAction() TenantChangeActionEnum {
	if o == nil || IsNil(o.Action) {
		var ret TenantChangeActionEnum
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantChangesInner) GetActionOk() (*TenantChangeActionEnum, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *TenantChangesInner) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given TenantChangeActionEnum and assigns it to the Action field.
func (o *TenantChangesInner) SetAction(v TenantChangeActionEnum) {
	o.Action = &v
}

func (o TenantChangesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantChangesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	return toSerialize, nil
}

type NullableTenantChangesInner struct {
	value *TenantChangesInner
	isSet bool
}

func (v NullableTenantChangesInner) Get() *TenantChangesInner {
	return v.value
}

func (v *NullableTenantChangesInner) Set(val *TenantChangesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantChangesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantChangesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantChangesInner(val *TenantChangesInner) *NullableTenantChangesInner {
	return &NullableTenantChangesInner{value: val, isSet: true}
}

func (v NullableTenantChangesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantChangesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}