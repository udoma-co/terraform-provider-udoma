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

// checks if the TenantChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantChange{}

// TenantChange A change to the tenants of a property. This is used to indicate that a tenant has moved out or moved in. To be able to know whe stayed in the property, we also provide the list of tenants that remained.
type TenantChange struct {
	Tenant *Tenant `json:"tenant,omitempty"`
	// Optional date when the change will enter into force
	EntryIntoForce *int64                  `json:"entry_into_force,omitempty"`
	Action         *TenantChangeActionEnum `json:"action,omitempty"`
}

// NewTenantChange instantiates a new TenantChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantChange() *TenantChange {
	this := TenantChange{}
	return &this
}

// NewTenantChangeWithDefaults instantiates a new TenantChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantChangeWithDefaults() *TenantChange {
	this := TenantChange{}
	return &this
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *TenantChange) GetTenant() Tenant {
	if o == nil || IsNil(o.Tenant) {
		var ret Tenant
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantChange) GetTenantOk() (*Tenant, bool) {
	if o == nil || IsNil(o.Tenant) {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *TenantChange) HasTenant() bool {
	if o != nil && !IsNil(o.Tenant) {
		return true
	}

	return false
}

// SetTenant gets a reference to the given Tenant and assigns it to the Tenant field.
func (o *TenantChange) SetTenant(v Tenant) {
	o.Tenant = &v
}

// GetEntryIntoForce returns the EntryIntoForce field value if set, zero value otherwise.
func (o *TenantChange) GetEntryIntoForce() int64 {
	if o == nil || IsNil(o.EntryIntoForce) {
		var ret int64
		return ret
	}
	return *o.EntryIntoForce
}

// GetEntryIntoForceOk returns a tuple with the EntryIntoForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantChange) GetEntryIntoForceOk() (*int64, bool) {
	if o == nil || IsNil(o.EntryIntoForce) {
		return nil, false
	}
	return o.EntryIntoForce, true
}

// HasEntryIntoForce returns a boolean if a field has been set.
func (o *TenantChange) HasEntryIntoForce() bool {
	if o != nil && !IsNil(o.EntryIntoForce) {
		return true
	}

	return false
}

// SetEntryIntoForce gets a reference to the given int64 and assigns it to the EntryIntoForce field.
func (o *TenantChange) SetEntryIntoForce(v int64) {
	o.EntryIntoForce = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *TenantChange) GetAction() TenantChangeActionEnum {
	if o == nil || IsNil(o.Action) {
		var ret TenantChangeActionEnum
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantChange) GetActionOk() (*TenantChangeActionEnum, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *TenantChange) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given TenantChangeActionEnum and assigns it to the Action field.
func (o *TenantChange) SetAction(v TenantChangeActionEnum) {
	o.Action = &v
}

func (o TenantChange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tenant) {
		toSerialize["tenant"] = o.Tenant
	}
	if !IsNil(o.EntryIntoForce) {
		toSerialize["entry_into_force"] = o.EntryIntoForce
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	return toSerialize, nil
}

type NullableTenantChange struct {
	value *TenantChange
	isSet bool
}

func (v NullableTenantChange) Get() *TenantChange {
	return v.value
}

func (v *NullableTenantChange) Set(val *TenantChange) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantChange) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantChange(val *TenantChange) *NullableTenantChange {
	return &NullableTenantChange{value: val, isSet: true}
}

func (v NullableTenantChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
