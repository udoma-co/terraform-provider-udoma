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

// checks if the CreateOrUpdateRentUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateRentUpdateRequest{}

// CreateOrUpdateRentUpdateRequest Request issued by a property manager to create a new update for a tenancy
type CreateOrUpdateRentUpdateRequest struct {
	// The date when the update will enter into force
	EntryIntoForce int64    `json:"entry_into_force"`
	Rent           RentData `json:"rent"`
	// A mask on which fields of rent to update. We cannot leave them empty because the generator doesn't differentiate between 0 and empty.
	Mask *int32 `json:"mask,omitempty"`
}

type _CreateOrUpdateRentUpdateRequest CreateOrUpdateRentUpdateRequest

// NewCreateOrUpdateRentUpdateRequest instantiates a new CreateOrUpdateRentUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateRentUpdateRequest(entryIntoForce int64, rent RentData) *CreateOrUpdateRentUpdateRequest {
	this := CreateOrUpdateRentUpdateRequest{}
	this.EntryIntoForce = entryIntoForce
	this.Rent = rent
	return &this
}

// NewCreateOrUpdateRentUpdateRequestWithDefaults instantiates a new CreateOrUpdateRentUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateRentUpdateRequestWithDefaults() *CreateOrUpdateRentUpdateRequest {
	this := CreateOrUpdateRentUpdateRequest{}
	return &this
}

// GetEntryIntoForce returns the EntryIntoForce field value
func (o *CreateOrUpdateRentUpdateRequest) GetEntryIntoForce() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EntryIntoForce
}

// GetEntryIntoForceOk returns a tuple with the EntryIntoForce field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateRentUpdateRequest) GetEntryIntoForceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryIntoForce, true
}

// SetEntryIntoForce sets field value
func (o *CreateOrUpdateRentUpdateRequest) SetEntryIntoForce(v int64) {
	o.EntryIntoForce = v
}

// GetRent returns the Rent field value
func (o *CreateOrUpdateRentUpdateRequest) GetRent() RentData {
	if o == nil {
		var ret RentData
		return ret
	}

	return o.Rent
}

// GetRentOk returns a tuple with the Rent field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateRentUpdateRequest) GetRentOk() (*RentData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rent, true
}

// SetRent sets field value
func (o *CreateOrUpdateRentUpdateRequest) SetRent(v RentData) {
	o.Rent = v
}

// GetMask returns the Mask field value if set, zero value otherwise.
func (o *CreateOrUpdateRentUpdateRequest) GetMask() int32 {
	if o == nil || IsNil(o.Mask) {
		var ret int32
		return ret
	}
	return *o.Mask
}

// GetMaskOk returns a tuple with the Mask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateRentUpdateRequest) GetMaskOk() (*int32, bool) {
	if o == nil || IsNil(o.Mask) {
		return nil, false
	}
	return o.Mask, true
}

// HasMask returns a boolean if a field has been set.
func (o *CreateOrUpdateRentUpdateRequest) HasMask() bool {
	if o != nil && !IsNil(o.Mask) {
		return true
	}

	return false
}

// SetMask gets a reference to the given int32 and assigns it to the Mask field.
func (o *CreateOrUpdateRentUpdateRequest) SetMask(v int32) {
	o.Mask = &v
}

func (o CreateOrUpdateRentUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateRentUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entry_into_force"] = o.EntryIntoForce
	toSerialize["rent"] = o.Rent
	if !IsNil(o.Mask) {
		toSerialize["mask"] = o.Mask
	}
	return toSerialize, nil
}

func (o *CreateOrUpdateRentUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varCreateOrUpdateRentUpdateRequest := _CreateOrUpdateRentUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrUpdateRentUpdateRequest)

	if err != nil {
		return err
	}

	*o = CreateOrUpdateRentUpdateRequest(varCreateOrUpdateRentUpdateRequest)

	return err
}

type NullableCreateOrUpdateRentUpdateRequest struct {
	value *CreateOrUpdateRentUpdateRequest
	isSet bool
}

func (v NullableCreateOrUpdateRentUpdateRequest) Get() *CreateOrUpdateRentUpdateRequest {
	return v.value
}

func (v *NullableCreateOrUpdateRentUpdateRequest) Set(val *CreateOrUpdateRentUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateRentUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateRentUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateRentUpdateRequest(val *CreateOrUpdateRentUpdateRequest) *NullableCreateOrUpdateRentUpdateRequest {
	return &NullableCreateOrUpdateRentUpdateRequest{value: val, isSet: true}
}

func (v NullableCreateOrUpdateRentUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateRentUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
