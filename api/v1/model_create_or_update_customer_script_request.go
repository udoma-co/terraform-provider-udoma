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

// checks if the CreateOrUpdateCustomerScriptRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateCustomerScriptRequest{}

// CreateOrUpdateCustomerScriptRequest struct for CreateOrUpdateCustomerScriptRequest
type CreateOrUpdateCustomerScriptRequest struct {
	// The function name under which the script will be available in the scripting environment
	Name string `json:"name"`
	// A description of the script
	Description *string              `json:"description,omitempty"`
	Scope       *CustomerScriptScope `json:"scope,omitempty"`
	// The actual JS script code
	Script string `json:"script"`
}

type _CreateOrUpdateCustomerScriptRequest CreateOrUpdateCustomerScriptRequest

// NewCreateOrUpdateCustomerScriptRequest instantiates a new CreateOrUpdateCustomerScriptRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateCustomerScriptRequest(name string, script string) *CreateOrUpdateCustomerScriptRequest {
	this := CreateOrUpdateCustomerScriptRequest{}
	this.Name = name
	this.Script = script
	return &this
}

// NewCreateOrUpdateCustomerScriptRequestWithDefaults instantiates a new CreateOrUpdateCustomerScriptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateCustomerScriptRequestWithDefaults() *CreateOrUpdateCustomerScriptRequest {
	this := CreateOrUpdateCustomerScriptRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateOrUpdateCustomerScriptRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateCustomerScriptRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrUpdateCustomerScriptRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOrUpdateCustomerScriptRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateCustomerScriptRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrUpdateCustomerScriptRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOrUpdateCustomerScriptRequest) SetDescription(v string) {
	o.Description = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *CreateOrUpdateCustomerScriptRequest) GetScope() CustomerScriptScope {
	if o == nil || IsNil(o.Scope) {
		var ret CustomerScriptScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateCustomerScriptRequest) GetScopeOk() (*CustomerScriptScope, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *CreateOrUpdateCustomerScriptRequest) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given CustomerScriptScope and assigns it to the Scope field.
func (o *CreateOrUpdateCustomerScriptRequest) SetScope(v CustomerScriptScope) {
	o.Scope = &v
}

// GetScript returns the Script field value
func (o *CreateOrUpdateCustomerScriptRequest) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateCustomerScriptRequest) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *CreateOrUpdateCustomerScriptRequest) SetScript(v string) {
	o.Script = v
}

func (o CreateOrUpdateCustomerScriptRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateCustomerScriptRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *CreateOrUpdateCustomerScriptRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varCreateOrUpdateCustomerScriptRequest := _CreateOrUpdateCustomerScriptRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrUpdateCustomerScriptRequest)

	if err != nil {
		return err
	}

	*o = CreateOrUpdateCustomerScriptRequest(varCreateOrUpdateCustomerScriptRequest)

	return err
}

type NullableCreateOrUpdateCustomerScriptRequest struct {
	value *CreateOrUpdateCustomerScriptRequest
	isSet bool
}

func (v NullableCreateOrUpdateCustomerScriptRequest) Get() *CreateOrUpdateCustomerScriptRequest {
	return v.value
}

func (v *NullableCreateOrUpdateCustomerScriptRequest) Set(val *CreateOrUpdateCustomerScriptRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateCustomerScriptRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateCustomerScriptRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateCustomerScriptRequest(val *CreateOrUpdateCustomerScriptRequest) *NullableCreateOrUpdateCustomerScriptRequest {
	return &NullableCreateOrUpdateCustomerScriptRequest{value: val, isSet: true}
}

func (v NullableCreateOrUpdateCustomerScriptRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateCustomerScriptRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
