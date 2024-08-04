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

// checks if the RentUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RentUpdate{}

// RentUpdate A rent update is a change to the monthly amount due for a rent of a property.
type RentUpdate struct {
	// Unique and immutable ID attribute of the entity that is generated when  the instance is created. The ID is unique within the system accross all accounts and it can be used to reference the entity in other entities  or to retrieve it from the backend.
	Id string `json:"id"`
	// The date and time the entity was created
	CreatedAt int64 `json:"created_at"`
	// The date and time the entity was last updated
	UpdatedAt int64 `json:"updated_at"`
	// Optional external ID, in case entity was created via backend integration
	ExternalId *string `json:"external_id,omitempty"`
	// Optional external source, in case entity was created via backend integration
	ExternalSource *string `json:"external_source,omitempty"`
	// The ID of the tenancy that this update is for
	TenancyRef string `json:"tenancy_ref"`
	// The date when the update will enter into force
	EntryIntoForce int64    `json:"entry_into_force"`
	Rent           RentData `json:"rent"`
	// A mask on which fields of rent to update. We cannot leave them empty because the generator doesn't differentiate between 0 and empty.
	Mask *int32 `json:"mask,omitempty"`
}

type _RentUpdate RentUpdate

// NewRentUpdate instantiates a new RentUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRentUpdate(id string, createdAt int64, updatedAt int64, tenancyRef string, entryIntoForce int64, rent RentData) *RentUpdate {
	this := RentUpdate{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.TenancyRef = tenancyRef
	this.EntryIntoForce = entryIntoForce
	this.Rent = rent
	return &this
}

// NewRentUpdateWithDefaults instantiates a new RentUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRentUpdateWithDefaults() *RentUpdate {
	this := RentUpdate{}
	return &this
}

// GetId returns the Id field value
func (o *RentUpdate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RentUpdate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RentUpdate) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *RentUpdate) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *RentUpdate) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *RentUpdate) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *RentUpdate) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *RentUpdate) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *RentUpdate) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *RentUpdate) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RentUpdate) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *RentUpdate) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *RentUpdate) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetExternalSource returns the ExternalSource field value if set, zero value otherwise.
func (o *RentUpdate) GetExternalSource() string {
	if o == nil || IsNil(o.ExternalSource) {
		var ret string
		return ret
	}
	return *o.ExternalSource
}

// GetExternalSourceOk returns a tuple with the ExternalSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RentUpdate) GetExternalSourceOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalSource) {
		return nil, false
	}
	return o.ExternalSource, true
}

// HasExternalSource returns a boolean if a field has been set.
func (o *RentUpdate) HasExternalSource() bool {
	if o != nil && !IsNil(o.ExternalSource) {
		return true
	}

	return false
}

// SetExternalSource gets a reference to the given string and assigns it to the ExternalSource field.
func (o *RentUpdate) SetExternalSource(v string) {
	o.ExternalSource = &v
}

// GetTenancyRef returns the TenancyRef field value
func (o *RentUpdate) GetTenancyRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenancyRef
}

// GetTenancyRefOk returns a tuple with the TenancyRef field value
// and a boolean to check if the value has been set.
func (o *RentUpdate) GetTenancyRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenancyRef, true
}

// SetTenancyRef sets field value
func (o *RentUpdate) SetTenancyRef(v string) {
	o.TenancyRef = v
}

// GetEntryIntoForce returns the EntryIntoForce field value
func (o *RentUpdate) GetEntryIntoForce() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EntryIntoForce
}

// GetEntryIntoForceOk returns a tuple with the EntryIntoForce field value
// and a boolean to check if the value has been set.
func (o *RentUpdate) GetEntryIntoForceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryIntoForce, true
}

// SetEntryIntoForce sets field value
func (o *RentUpdate) SetEntryIntoForce(v int64) {
	o.EntryIntoForce = v
}

// GetRent returns the Rent field value
func (o *RentUpdate) GetRent() RentData {
	if o == nil {
		var ret RentData
		return ret
	}

	return o.Rent
}

// GetRentOk returns a tuple with the Rent field value
// and a boolean to check if the value has been set.
func (o *RentUpdate) GetRentOk() (*RentData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rent, true
}

// SetRent sets field value
func (o *RentUpdate) SetRent(v RentData) {
	o.Rent = v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *RentUpdate) GetMask() int32 {
	if o == nil || IsNil(o.Mask) {
		var ret int32
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RentUpdate) GetMaskOk() (*int32, bool) {
	if o == nil || IsNil(o.Mask) {
		return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *RentUpdate) HasMask() bool {
	if o != nil && !IsNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given int32 and assigns it to the Mask field.
func (o *RentUpdate) SetMask(v int32) {
	o.Mask = &v
}

func (o RentUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RentUpdate) ToMap() (map[string]interface{}, error) {
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
	toSerialize["tenancy_ref"] = o.TenancyRef
	toSerialize["entry_into_force"] = o.EntryIntoForce
	toSerialize["rent"] = o.Rent
	if !IsNil(o.Mask) {
		toSerialize["mask"] = o.Mask
	}
	return toSerialize, nil
}

func (o *RentUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"tenancy_ref",
		"entry_into_force",
		"rent",
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

	varRentUpdate := _RentUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRentUpdate)

	if err != nil {
		return err
	}

	*o = RentUpdate(varRentUpdate)

	return err
}

type NullableRentUpdate struct {
	value *RentUpdate
	isSet bool
}

func (v NullableRentUpdate) Get() *RentUpdate {
	return v.value
}

func (v *NullableRentUpdate) Set(val *RentUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableRentUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableRentUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRentUpdate(val *RentUpdate) *NullableRentUpdate {
	return &NullableRentUpdate{value: val, isSet: true}
}

func (v NullableRentUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRentUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
