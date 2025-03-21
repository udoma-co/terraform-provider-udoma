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

// checks if the DocumentTemplateSignerDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentTemplateSignerDefinition{}

// DocumentTemplateSignerDefinition Indicates where to get data to generate info of a signer
type DocumentTemplateSignerDefinition struct {
	ContactData NullableContactData `json:"contact_data,omitempty"`
	// JS expression to compute the contact data of the signer. Must be set if repeated is true, or if contact_data is not set.
	ContactDataExpression *string `json:"contact_data_expression,omitempty"`
	// If set to true, the signer definition will be repeated for each value of the name and contract data expressions.
	Repeated *bool `json:"repeated,omitempty"`
}

// NewDocumentTemplateSignerDefinition instantiates a new DocumentTemplateSignerDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentTemplateSignerDefinition() *DocumentTemplateSignerDefinition {
	this := DocumentTemplateSignerDefinition{}
	return &this
}

// NewDocumentTemplateSignerDefinitionWithDefaults instantiates a new DocumentTemplateSignerDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentTemplateSignerDefinitionWithDefaults() *DocumentTemplateSignerDefinition {
	this := DocumentTemplateSignerDefinition{}
	return &this
}

// GetContactData returns the ContactData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentTemplateSignerDefinition) GetContactData() ContactData {
	if o == nil || IsNil(o.ContactData.Get()) {
		var ret ContactData
		return ret
	}
	return *o.ContactData.Get()
}

// GetContactDataOk returns a tuple with the ContactData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentTemplateSignerDefinition) GetContactDataOk() (*ContactData, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContactData.Get(), o.ContactData.IsSet()
}

// HasContactData returns a boolean if a field has been set.
func (o *DocumentTemplateSignerDefinition) HasContactData() bool {
	if o != nil && o.ContactData.IsSet() {
		return true
	}

	return false
}

// SetContactData gets a reference to the given NullableContactData and assigns it to the ContactData field.
func (o *DocumentTemplateSignerDefinition) SetContactData(v ContactData) {
	o.ContactData.Set(&v)
}

// SetContactDataNil sets the value for ContactData to be an explicit nil
func (o *DocumentTemplateSignerDefinition) SetContactDataNil() {
	o.ContactData.Set(nil)
}

// UnsetContactData ensures that no value is present for ContactData, not even an explicit nil
func (o *DocumentTemplateSignerDefinition) UnsetContactData() {
	o.ContactData.Unset()
}

// GetContactDataExpression returns the ContactDataExpression field value if set, zero value otherwise.
func (o *DocumentTemplateSignerDefinition) GetContactDataExpression() string {
	if o == nil || IsNil(o.ContactDataExpression) {
		var ret string
		return ret
	}
	return *o.ContactDataExpression
}

// GetContactDataExpressionOk returns a tuple with the ContactDataExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentTemplateSignerDefinition) GetContactDataExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.ContactDataExpression) {
		return nil, false
	}
	return o.ContactDataExpression, true
}

// HasContactDataExpression returns a boolean if a field has been set.
func (o *DocumentTemplateSignerDefinition) HasContactDataExpression() bool {
	if o != nil && !IsNil(o.ContactDataExpression) {
		return true
	}

	return false
}

// SetContactDataExpression gets a reference to the given string and assigns it to the ContactDataExpression field.
func (o *DocumentTemplateSignerDefinition) SetContactDataExpression(v string) {
	o.ContactDataExpression = &v
}

// GetRepeated returns the Repeated field value if set, zero value otherwise.
func (o *DocumentTemplateSignerDefinition) GetRepeated() bool {
	if o == nil || IsNil(o.Repeated) {
		var ret bool
		return ret
	}
	return *o.Repeated
}

// GetRepeatedOk returns a tuple with the Repeated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentTemplateSignerDefinition) GetRepeatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Repeated) {
		return nil, false
	}
	return o.Repeated, true
}

// HasRepeated returns a boolean if a field has been set.
func (o *DocumentTemplateSignerDefinition) HasRepeated() bool {
	if o != nil && !IsNil(o.Repeated) {
		return true
	}

	return false
}

// SetRepeated gets a reference to the given bool and assigns it to the Repeated field.
func (o *DocumentTemplateSignerDefinition) SetRepeated(v bool) {
	o.Repeated = &v
}

func (o DocumentTemplateSignerDefinition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentTemplateSignerDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContactData.IsSet() {
		toSerialize["contact_data"] = o.ContactData.Get()
	}
	if !IsNil(o.ContactDataExpression) {
		toSerialize["contact_data_expression"] = o.ContactDataExpression
	}
	if !IsNil(o.Repeated) {
		toSerialize["repeated"] = o.Repeated
	}
	return toSerialize, nil
}

type NullableDocumentTemplateSignerDefinition struct {
	value *DocumentTemplateSignerDefinition
	isSet bool
}

func (v NullableDocumentTemplateSignerDefinition) Get() *DocumentTemplateSignerDefinition {
	return v.value
}

func (v *NullableDocumentTemplateSignerDefinition) Set(val *DocumentTemplateSignerDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentTemplateSignerDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentTemplateSignerDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentTemplateSignerDefinition(val *DocumentTemplateSignerDefinition) *NullableDocumentTemplateSignerDefinition {
	return &NullableDocumentTemplateSignerDefinition{value: val, isSet: true}
}

func (v NullableDocumentTemplateSignerDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentTemplateSignerDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
