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

// checks if the TenantChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantChange{}

// TenantChange A change to the tenants of a property. This is used to indicate that a tenant has moved out or moved in. To be able to know whe stayed in the property, we also provide the list of tenants that remained.
type TenantChange struct {
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
	Tenant         Tenant  `json:"tenant"`
	// The ID of the tenancy that this update is for
	TenancyRef string `json:"tenancy_ref"`
	// The date when the update will enter into force
	EntryIntoForce int64                  `json:"entry_into_force"`
	Action         TenantChangeActionEnum `json:"action"`
}

type _TenantChange TenantChange

// NewTenantChange instantiates a new TenantChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantChange(id string, createdAt int64, updatedAt int64, tenant Tenant, tenancyRef string, entryIntoForce int64, action TenantChangeActionEnum) *TenantChange {
	this := TenantChange{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Tenant = tenant
	this.TenancyRef = tenancyRef
	this.EntryIntoForce = entryIntoForce
	this.Action = action
	return &this
}

// NewTenantChangeWithDefaults instantiates a new TenantChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantChangeWithDefaults() *TenantChange {
	this := TenantChange{}
	return &this
}

// GetId returns the Id field value
func (o *TenantChange) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TenantChange) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TenantChange) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TenantChange) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TenantChange) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TenantChange) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TenantChange) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TenantChange) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TenantChange) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *TenantChange) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantChange) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *TenantChange) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *TenantChange) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetExternalSource returns the ExternalSource field value if set, zero value otherwise.
func (o *TenantChange) GetExternalSource() string {
	if o == nil || IsNil(o.ExternalSource) {
		var ret string
		return ret
	}
	return *o.ExternalSource
}

// GetExternalSourceOk returns a tuple with the ExternalSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantChange) GetExternalSourceOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalSource) {
		return nil, false
	}
	return o.ExternalSource, true
}

// HasExternalSource returns a boolean if a field has been set.
func (o *TenantChange) HasExternalSource() bool {
	if o != nil && !IsNil(o.ExternalSource) {
		return true
	}

	return false
}

// SetExternalSource gets a reference to the given string and assigns it to the ExternalSource field.
func (o *TenantChange) SetExternalSource(v string) {
	o.ExternalSource = &v
}

// GetTenant returns the Tenant field value
func (o *TenantChange) GetTenant() Tenant {
	if o == nil {
		var ret Tenant
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *TenantChange) GetTenantOk() (*Tenant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *TenantChange) SetTenant(v Tenant) {
	o.Tenant = v
}

// GetTenancyRef returns the TenancyRef field value
func (o *TenantChange) GetTenancyRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenancyRef
}

// GetTenancyRefOk returns a tuple with the TenancyRef field value
// and a boolean to check if the value has been set.
func (o *TenantChange) GetTenancyRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenancyRef, true
}

// SetTenancyRef sets field value
func (o *TenantChange) SetTenancyRef(v string) {
	o.TenancyRef = v
}

// GetEntryIntoForce returns the EntryIntoForce field value
func (o *TenantChange) GetEntryIntoForce() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EntryIntoForce
}

// GetEntryIntoForceOk returns a tuple with the EntryIntoForce field value
// and a boolean to check if the value has been set.
func (o *TenantChange) GetEntryIntoForceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryIntoForce, true
}

// SetEntryIntoForce sets field value
func (o *TenantChange) SetEntryIntoForce(v int64) {
	o.EntryIntoForce = v
}

// GetAction returns the Action field value
func (o *TenantChange) GetAction() TenantChangeActionEnum {
	if o == nil {
		var ret TenantChangeActionEnum
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *TenantChange) GetActionOk() (*TenantChangeActionEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *TenantChange) SetAction(v TenantChangeActionEnum) {
	o.Action = v
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
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if !IsNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	if !IsNil(o.ExternalSource) {
		toSerialize["external_source"] = o.ExternalSource
	}
	toSerialize["tenant"] = o.Tenant
	toSerialize["tenancy_ref"] = o.TenancyRef
	toSerialize["entry_into_force"] = o.EntryIntoForce
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

func (o *TenantChange) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"tenant",
		"tenancy_ref",
		"entry_into_force",
		"action",
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

	varTenantChange := _TenantChange{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTenantChange)

	if err != nil {
		return err
	}

	*o = TenantChange(varTenantChange)

	return err
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
