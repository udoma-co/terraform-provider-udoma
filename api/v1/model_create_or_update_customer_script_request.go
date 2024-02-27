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

// checks if the CreateOrUpdateCustomerScriptRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateCustomerScriptRequest{}

// CreateOrUpdateCustomerScriptRequest struct for CreateOrUpdateCustomerScriptRequest
type CreateOrUpdateCustomerScriptRequest struct {
	// The function name under which the script will be available in the scripting environment
	Name *string `json:"name,omitempty"`
	// A description of the script
	Description *string              `json:"description,omitempty"`
	Scope       *CustomerScriptScope `json:"scope,omitempty"`
	// The actual JS script code
	Script *string `json:"script,omitempty"`
}

// NewCreateOrUpdateCustomerScriptRequest instantiates a new CreateOrUpdateCustomerScriptRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateCustomerScriptRequest() *CreateOrUpdateCustomerScriptRequest {
	this := CreateOrUpdateCustomerScriptRequest{}
	return &this
}

// NewCreateOrUpdateCustomerScriptRequestWithDefaults instantiates a new CreateOrUpdateCustomerScriptRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateCustomerScriptRequestWithDefaults() *CreateOrUpdateCustomerScriptRequest {
	this := CreateOrUpdateCustomerScriptRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateOrUpdateCustomerScriptRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateCustomerScriptRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateOrUpdateCustomerScriptRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateOrUpdateCustomerScriptRequest) SetName(v string) {
	o.Name = &v
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

// GetScript returns the Script field value if set, zero value otherwise.
func (o *CreateOrUpdateCustomerScriptRequest) GetScript() string {
	if o == nil || IsNil(o.Script) {
		var ret string
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateCustomerScriptRequest) GetScriptOk() (*string, bool) {
	if o == nil || IsNil(o.Script) {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *CreateOrUpdateCustomerScriptRequest) HasScript() bool {
	if o != nil && !IsNil(o.Script) {
		return true
	}

	return false
}

// SetScript gets a reference to the given string and assigns it to the Script field.
func (o *CreateOrUpdateCustomerScriptRequest) SetScript(v string) {
	o.Script = &v
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
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.Script) {
		toSerialize["script"] = o.Script
	}
	return toSerialize, nil
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
