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

// CustomIDGenerator struct for CustomIDGenerator
type CustomIDGenerator struct {
	Id *string `json:"id,omitempty"`
	// The timestamp of when the script was created
	CreatedAt *int64 `json:"created_at,omitempty"`
	// The timestamp of when the script was last updated
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// The user friendly name of the ID generator.
	Name *string `json:"name,omitempty"`
	// A JS script code that will be used to generate the ID.
	GenerationScript *string `json:"generation_script,omitempty"`
	// The last generated ID. This is used to keep track of the last generated ID.
	LastGeneratedId *string `json:"last_generated_id,omitempty"`
}

// NewCustomIDGenerator instantiates a new CustomIDGenerator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomIDGenerator() *CustomIDGenerator {
	this := CustomIDGenerator{}
	return &this
}

// NewCustomIDGeneratorWithDefaults instantiates a new CustomIDGenerator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomIDGeneratorWithDefaults() *CustomIDGenerator {
	this := CustomIDGenerator{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomIDGenerator) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomIDGenerator) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomIDGenerator) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomIDGenerator) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CustomIDGenerator) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomIDGenerator) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CustomIDGenerator) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *CustomIDGenerator) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CustomIDGenerator) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomIDGenerator) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CustomIDGenerator) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *CustomIDGenerator) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomIDGenerator) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomIDGenerator) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomIDGenerator) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomIDGenerator) SetName(v string) {
	o.Name = &v
}

// GetGenerationScript returns the GenerationScript field value if set, zero value otherwise.
func (o *CustomIDGenerator) GetGenerationScript() string {
	if o == nil || o.GenerationScript == nil {
		var ret string
		return ret
	}
	return *o.GenerationScript
}

// GetGenerationScriptOk returns a tuple with the GenerationScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomIDGenerator) GetGenerationScriptOk() (*string, bool) {
	if o == nil || o.GenerationScript == nil {
		return nil, false
	}
	return o.GenerationScript, true
}

// HasGenerationScript returns a boolean if a field has been set.
func (o *CustomIDGenerator) HasGenerationScript() bool {
	if o != nil && o.GenerationScript != nil {
		return true
	}

	return false
}

// SetGenerationScript gets a reference to the given string and assigns it to the GenerationScript field.
func (o *CustomIDGenerator) SetGenerationScript(v string) {
	o.GenerationScript = &v
}

// GetLastGeneratedId returns the LastGeneratedId field value if set, zero value otherwise.
func (o *CustomIDGenerator) GetLastGeneratedId() string {
	if o == nil || o.LastGeneratedId == nil {
		var ret string
		return ret
	}
	return *o.LastGeneratedId
}

// GetLastGeneratedIdOk returns a tuple with the LastGeneratedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomIDGenerator) GetLastGeneratedIdOk() (*string, bool) {
	if o == nil || o.LastGeneratedId == nil {
		return nil, false
	}
	return o.LastGeneratedId, true
}

// HasLastGeneratedId returns a boolean if a field has been set.
func (o *CustomIDGenerator) HasLastGeneratedId() bool {
	if o != nil && o.LastGeneratedId != nil {
		return true
	}

	return false
}

// SetLastGeneratedId gets a reference to the given string and assigns it to the LastGeneratedId field.
func (o *CustomIDGenerator) SetLastGeneratedId(v string) {
	o.LastGeneratedId = &v
}

func (o CustomIDGenerator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.GenerationScript != nil {
		toSerialize["generation_script"] = o.GenerationScript
	}
	if o.LastGeneratedId != nil {
		toSerialize["last_generated_id"] = o.LastGeneratedId
	}
	return json.Marshal(toSerialize)
}

type NullableCustomIDGenerator struct {
	value *CustomIDGenerator
	isSet bool
}

func (v NullableCustomIDGenerator) Get() *CustomIDGenerator {
	return v.value
}

func (v *NullableCustomIDGenerator) Set(val *CustomIDGenerator) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomIDGenerator) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomIDGenerator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomIDGenerator(val *CustomIDGenerator) *NullableCustomIDGenerator {
	return &NullableCustomIDGenerator{value: val, isSet: true}
}

func (v NullableCustomIDGenerator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomIDGenerator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
