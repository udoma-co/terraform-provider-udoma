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

// checks if the CreateOrUpdatePropertyHandoverTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdatePropertyHandoverTemplateRequest{}

// CreateOrUpdatePropertyHandoverTemplateRequest struct for CreateOrUpdatePropertyHandoverTemplateRequest
type CreateOrUpdatePropertyHandoverTemplateRequest struct {
	// The name of the property handover template.
	Name string `json:"name"`
	// The description of the property handover template.
	Description *string    `json:"description,omitempty"`
	CustomForm  CustomForm `json:"custom_form"`
}

type _CreateOrUpdatePropertyHandoverTemplateRequest CreateOrUpdatePropertyHandoverTemplateRequest

// NewCreateOrUpdatePropertyHandoverTemplateRequest instantiates a new CreateOrUpdatePropertyHandoverTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdatePropertyHandoverTemplateRequest(name string, customForm CustomForm) *CreateOrUpdatePropertyHandoverTemplateRequest {
	this := CreateOrUpdatePropertyHandoverTemplateRequest{}
	this.Name = name
	this.CustomForm = customForm
	return &this
}

// NewCreateOrUpdatePropertyHandoverTemplateRequestWithDefaults instantiates a new CreateOrUpdatePropertyHandoverTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdatePropertyHandoverTemplateRequestWithDefaults() *CreateOrUpdatePropertyHandoverTemplateRequest {
	this := CreateOrUpdatePropertyHandoverTemplateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateOrUpdatePropertyHandoverTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdatePropertyHandoverTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrUpdatePropertyHandoverTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOrUpdatePropertyHandoverTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdatePropertyHandoverTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrUpdatePropertyHandoverTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOrUpdatePropertyHandoverTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCustomForm returns the CustomForm field value
func (o *CreateOrUpdatePropertyHandoverTemplateRequest) GetCustomForm() CustomForm {
	if o == nil {
		var ret CustomForm
		return ret
	}

	return o.CustomForm
}

// GetCustomFormOk returns a tuple with the CustomForm field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdatePropertyHandoverTemplateRequest) GetCustomFormOk() (*CustomForm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomForm, true
}

// SetCustomForm sets field value
func (o *CreateOrUpdatePropertyHandoverTemplateRequest) SetCustomForm(v CustomForm) {
	o.CustomForm = v
}

func (o CreateOrUpdatePropertyHandoverTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdatePropertyHandoverTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["custom_form"] = o.CustomForm
	return toSerialize, nil
}

func (o *CreateOrUpdatePropertyHandoverTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"custom_form",
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

	varCreateOrUpdatePropertyHandoverTemplateRequest := _CreateOrUpdatePropertyHandoverTemplateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrUpdatePropertyHandoverTemplateRequest)

	if err != nil {
		return err
	}

	*o = CreateOrUpdatePropertyHandoverTemplateRequest(varCreateOrUpdatePropertyHandoverTemplateRequest)

	return err
}

type NullableCreateOrUpdatePropertyHandoverTemplateRequest struct {
	value *CreateOrUpdatePropertyHandoverTemplateRequest
	isSet bool
}

func (v NullableCreateOrUpdatePropertyHandoverTemplateRequest) Get() *CreateOrUpdatePropertyHandoverTemplateRequest {
	return v.value
}

func (v *NullableCreateOrUpdatePropertyHandoverTemplateRequest) Set(val *CreateOrUpdatePropertyHandoverTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdatePropertyHandoverTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdatePropertyHandoverTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdatePropertyHandoverTemplateRequest(val *CreateOrUpdatePropertyHandoverTemplateRequest) *NullableCreateOrUpdatePropertyHandoverTemplateRequest {
	return &NullableCreateOrUpdatePropertyHandoverTemplateRequest{value: val, isSet: true}
}

func (v NullableCreateOrUpdatePropertyHandoverTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdatePropertyHandoverTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
