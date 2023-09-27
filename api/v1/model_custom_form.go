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

// CustomForm a custom form that is used to collect data from the user
type CustomForm struct {
	// the layout of the form, which groups and inputs will be displayed
	Layout []FormItem `json:"layout,omitempty"`
	// the groups of inputs that will be displayed to the user
	Groups []FormGroup `json:"groups,omitempty"`
	// the inputs that will be displayed to the user
	Inputs []FormInput `json:"inputs,omitempty"`
	// the validations that will be performed on the data provided by the user
	Validations []FormValidation `json:"validations,omitempty"`
}

// NewCustomForm instantiates a new CustomForm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomForm() *CustomForm {
	this := CustomForm{}
	return &this
}

// NewCustomFormWithDefaults instantiates a new CustomForm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFormWithDefaults() *CustomForm {
	this := CustomForm{}
	return &this
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *CustomForm) GetLayout() []FormItem {
	if o == nil || o.Layout == nil {
		var ret []FormItem
		return ret
	}
	return o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomForm) GetLayoutOk() ([]FormItem, bool) {
	if o == nil || o.Layout == nil {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *CustomForm) HasLayout() bool {
	if o != nil && o.Layout != nil {
		return true
	}

	return false
}

// SetLayout gets a reference to the given []FormItem and assigns it to the Layout field.
func (o *CustomForm) SetLayout(v []FormItem) {
	o.Layout = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *CustomForm) GetGroups() []FormGroup {
	if o == nil || o.Groups == nil {
		var ret []FormGroup
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomForm) GetGroupsOk() ([]FormGroup, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *CustomForm) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []FormGroup and assigns it to the Groups field.
func (o *CustomForm) SetGroups(v []FormGroup) {
	o.Groups = v
}

// GetInputs returns the Inputs field value if set, zero value otherwise.
func (o *CustomForm) GetInputs() []FormInput {
	if o == nil || o.Inputs == nil {
		var ret []FormInput
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomForm) GetInputsOk() ([]FormInput, bool) {
	if o == nil || o.Inputs == nil {
		return nil, false
	}
	return o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *CustomForm) HasInputs() bool {
	if o != nil && o.Inputs != nil {
		return true
	}

	return false
}

// SetInputs gets a reference to the given []FormInput and assigns it to the Inputs field.
func (o *CustomForm) SetInputs(v []FormInput) {
	o.Inputs = v
}

// GetValidations returns the Validations field value if set, zero value otherwise.
func (o *CustomForm) GetValidations() []FormValidation {
	if o == nil || o.Validations == nil {
		var ret []FormValidation
		return ret
	}
	return o.Validations
}

// GetValidationsOk returns a tuple with the Validations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomForm) GetValidationsOk() ([]FormValidation, bool) {
	if o == nil || o.Validations == nil {
		return nil, false
	}
	return o.Validations, true
}

// HasValidations returns a boolean if a field has been set.
func (o *CustomForm) HasValidations() bool {
	if o != nil && o.Validations != nil {
		return true
	}

	return false
}

// SetValidations gets a reference to the given []FormValidation and assigns it to the Validations field.
func (o *CustomForm) SetValidations(v []FormValidation) {
	o.Validations = v
}

func (o CustomForm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Layout != nil {
		toSerialize["layout"] = o.Layout
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.Inputs != nil {
		toSerialize["inputs"] = o.Inputs
	}
	if o.Validations != nil {
		toSerialize["validations"] = o.Validations
	}
	return json.Marshal(toSerialize)
}

type NullableCustomForm struct {
	value *CustomForm
	isSet bool
}

func (v NullableCustomForm) Get() *CustomForm {
	return v.value
}

func (v *NullableCustomForm) Set(val *CustomForm) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomForm) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomForm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomForm(val *CustomForm) *NullableCustomForm {
	return &NullableCustomForm{value: val, isSet: true}
}

func (v NullableCustomForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomForm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
