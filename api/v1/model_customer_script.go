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

// checks if the CustomerScript type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerScript{}

// CustomerScript struct for CustomerScript
type CustomerScript struct {
	// Unique and immutable ID attribute of the entity that is generated when the instance is created. The ID is unique within the system accross all accounts and it can be used to reference the entity in other entities or to retrieve it from the backend.
	Id string `json:"id"`
	// The date and time the entity was created
	CreatedAt int64 `json:"created_at"`
	// The date and time the entity was last updated
	UpdatedAt int64 `json:"updated_at"`
	// The function name under which the script will be available in the scripting environment
	Name string `json:"name"`
	// A description of the script
	Description *string              `json:"description,omitempty"`
	Scope       *CustomerScriptScope `json:"scope,omitempty"`
	// The actual JS script code
	Script string `json:"script"`
}

type _CustomerScript CustomerScript

// NewCustomerScript instantiates a new CustomerScript object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerScript(id string, createdAt int64, updatedAt int64, name string, script string) *CustomerScript {
	this := CustomerScript{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Name = name
	this.Script = script
	return &this
}

// NewCustomerScriptWithDefaults instantiates a new CustomerScript object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerScriptWithDefaults() *CustomerScript {
	this := CustomerScript{}
	return &this
}

// GetId returns the Id field value
func (o *CustomerScript) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomerScript) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CustomerScript) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CustomerScript) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomerScript) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CustomerScript) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CustomerScript) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomerScript) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CustomerScript) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetName returns the Name field value
func (o *CustomerScript) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomerScript) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CustomerScript) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomerScript) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerScript) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomerScript) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomerScript) SetDescription(v string) {
	o.Description = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *CustomerScript) GetScope() CustomerScriptScope {
	if o == nil || IsNil(o.Scope) {
		var ret CustomerScriptScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerScript) GetScopeOk() (*CustomerScriptScope, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *CustomerScript) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given CustomerScriptScope and assigns it to the Scope field.
func (o *CustomerScript) SetScope(v CustomerScriptScope) {
	o.Scope = &v
}

// GetScript returns the Script field value
func (o *CustomerScript) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *CustomerScript) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *CustomerScript) SetScript(v string) {
	o.Script = v
}

func (o CustomerScript) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerScript) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	toSerialize["script"] = o.Script
	return toSerialize, nil
}

func (o *CustomerScript) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"name",
		"script",
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

	varCustomerScript := _CustomerScript{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCustomerScript)

	if err != nil {
		return err
	}

	*o = CustomerScript(varCustomerScript)

	return err
}

type NullableCustomerScript struct {
	value *CustomerScript
	isSet bool
}

func (v NullableCustomerScript) Get() *CustomerScript {
	return v.value
}

func (v *NullableCustomerScript) Set(val *CustomerScript) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerScript) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerScript) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerScript(val *CustomerScript) *NullableCustomerScript {
	return &NullableCustomerScript{value: val, isSet: true}
}

func (v NullableCustomerScript) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerScript) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
