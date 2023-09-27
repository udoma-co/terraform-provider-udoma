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

// UpdateTenancyUpdateRequest Request issued by a property manager to update an existing update for a tenancy
type UpdateTenancyUpdateRequest struct {
	// The date when the update will enter into force
	EntryIntoForce *int64    `json:"entry_into_force,omitempty"`
	NewRent        *RentData `json:"new_rent,omitempty"`
	// Optional change to the tenants.
	TenantChanges []TenantChangeRequest `json:"tenant_changes,omitempty"`
}

// NewUpdateTenancyUpdateRequest instantiates a new UpdateTenancyUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTenancyUpdateRequest() *UpdateTenancyUpdateRequest {
	this := UpdateTenancyUpdateRequest{}
	return &this
}

// NewUpdateTenancyUpdateRequestWithDefaults instantiates a new UpdateTenancyUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTenancyUpdateRequestWithDefaults() *UpdateTenancyUpdateRequest {
	this := UpdateTenancyUpdateRequest{}
	return &this
}

// GetEntryIntoForce returns the EntryIntoForce field value if set, zero value otherwise.
func (o *UpdateTenancyUpdateRequest) GetEntryIntoForce() int64 {
	if o == nil || o.EntryIntoForce == nil {
		var ret int64
		return ret
	}
	return *o.EntryIntoForce
}

// GetEntryIntoForceOk returns a tuple with the EntryIntoForce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenancyUpdateRequest) GetEntryIntoForceOk() (*int64, bool) {
	if o == nil || o.EntryIntoForce == nil {
		return nil, false
	}
	return o.EntryIntoForce, true
}

// HasEntryIntoForce returns a boolean if a field has been set.
func (o *UpdateTenancyUpdateRequest) HasEntryIntoForce() bool {
	if o != nil && o.EntryIntoForce != nil {
		return true
	}

	return false
}

// SetEntryIntoForce gets a reference to the given int64 and assigns it to the EntryIntoForce field.
func (o *UpdateTenancyUpdateRequest) SetEntryIntoForce(v int64) {
	o.EntryIntoForce = &v
}

// GetNewRent returns the NewRent field value if set, zero value otherwise.
func (o *UpdateTenancyUpdateRequest) GetNewRent() RentData {
	if o == nil || o.NewRent == nil {
		var ret RentData
		return ret
	}
	return *o.NewRent
}

// GetNewRentOk returns a tuple with the NewRent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenancyUpdateRequest) GetNewRentOk() (*RentData, bool) {
	if o == nil || o.NewRent == nil {
		return nil, false
	}
	return o.NewRent, true
}

// HasNewRent returns a boolean if a field has been set.
func (o *UpdateTenancyUpdateRequest) HasNewRent() bool {
	if o != nil && o.NewRent != nil {
		return true
	}

	return false
}

// SetNewRent gets a reference to the given RentData and assigns it to the NewRent field.
func (o *UpdateTenancyUpdateRequest) SetNewRent(v RentData) {
	o.NewRent = &v
}

// GetTenantChanges returns the TenantChanges field value if set, zero value otherwise.
func (o *UpdateTenancyUpdateRequest) GetTenantChanges() []TenantChangeRequest {
	if o == nil || o.TenantChanges == nil {
		var ret []TenantChangeRequest
		return ret
	}
	return o.TenantChanges
}

// GetTenantChangesOk returns a tuple with the TenantChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenancyUpdateRequest) GetTenantChangesOk() ([]TenantChangeRequest, bool) {
	if o == nil || o.TenantChanges == nil {
		return nil, false
	}
	return o.TenantChanges, true
}

// HasTenantChanges returns a boolean if a field has been set.
func (o *UpdateTenancyUpdateRequest) HasTenantChanges() bool {
	if o != nil && o.TenantChanges != nil {
		return true
	}

	return false
}

// SetTenantChanges gets a reference to the given []TenantChangeRequest and assigns it to the TenantChanges field.
func (o *UpdateTenancyUpdateRequest) SetTenantChanges(v []TenantChangeRequest) {
	o.TenantChanges = v
}

func (o UpdateTenancyUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EntryIntoForce != nil {
		toSerialize["entry_into_force"] = o.EntryIntoForce
	}
	if o.NewRent != nil {
		toSerialize["new_rent"] = o.NewRent
	}
	if o.TenantChanges != nil {
		toSerialize["tenant_changes"] = o.TenantChanges
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateTenancyUpdateRequest struct {
	value *UpdateTenancyUpdateRequest
	isSet bool
}

func (v NullableUpdateTenancyUpdateRequest) Get() *UpdateTenancyUpdateRequest {
	return v.value
}

func (v *NullableUpdateTenancyUpdateRequest) Set(val *UpdateTenancyUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTenancyUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTenancyUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTenancyUpdateRequest(val *UpdateTenancyUpdateRequest) *NullableUpdateTenancyUpdateRequest {
	return &NullableUpdateTenancyUpdateRequest{value: val, isSet: true}
}

func (v NullableUpdateTenancyUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTenancyUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
