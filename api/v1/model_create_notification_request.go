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

// checks if the CreateNotificationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNotificationRequest{}

// CreateNotificationRequest Create request for a new notification.
type CreateNotificationRequest struct {
	// A unique, descriptive, and short identifier name.
	Name string `json:"name"`
	// Whether this template is used an intermediate template or not. If true it cannot be used as a regular template and must only be referenced by other templates.
	IsIntermediate *bool `json:"is_intermediate,omitempty"`
	// A short description of the notification.
	Description *string `json:"description,omitempty"`
	// A script that can run the notification given some initial data.
	Script *string `json:"script,omitempty"`
	// A potential reference to an intermediate template. Empty if no intermediate used.
	IntermediateRef *string `json:"intermediate_ref,omitempty"`
	// The golang html template to use for the notification.
	TemplateHtml string `json:"template_html"`
	// The golang text template to use for the notification.
	TemplateText string `json:"template_text"`
}

type _CreateNotificationRequest CreateNotificationRequest

// NewCreateNotificationRequest instantiates a new CreateNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNotificationRequest(name string, templateHtml string, templateText string) *CreateNotificationRequest {
	this := CreateNotificationRequest{}
	this.Name = name
	this.TemplateHtml = templateHtml
	this.TemplateText = templateText
	return &this
}

// NewCreateNotificationRequestWithDefaults instantiates a new CreateNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNotificationRequestWithDefaults() *CreateNotificationRequest {
	this := CreateNotificationRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNotificationRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNotificationRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNotificationRequest) SetName(v string) {
	o.Name = v
}

// GetIsIntermediate returns the IsIntermediate field value if set, zero value otherwise.
func (o *CreateNotificationRequest) GetIsIntermediate() bool {
	if o == nil || IsNil(o.IsIntermediate) {
		var ret bool
		return ret
	}
	return *o.IsIntermediate
}

// GetIsIntermediateOk returns a tuple with the IsIntermediate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationRequest) GetIsIntermediateOk() (*bool, bool) {
	if o == nil || IsNil(o.IsIntermediate) {
		return nil, false
	}
	return o.IsIntermediate, true
}

// HasIsIntermediate returns a boolean if a field has been set.
func (o *CreateNotificationRequest) HasIsIntermediate() bool {
	if o != nil && !IsNil(o.IsIntermediate) {
		return true
	}

	return false
}

// SetIsIntermediate gets a reference to the given bool and assigns it to the IsIntermediate field.
func (o *CreateNotificationRequest) SetIsIntermediate(v bool) {
	o.IsIntermediate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateNotificationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateNotificationRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateNotificationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *CreateNotificationRequest) GetScript() string {
	if o == nil || IsNil(o.Script) {
		var ret string
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationRequest) GetScriptOk() (*string, bool) {
	if o == nil || IsNil(o.Script) {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *CreateNotificationRequest) HasScript() bool {
	if o != nil && !IsNil(o.Script) {
		return true
	}

	return false
}

// SetScript gets a reference to the given string and assigns it to the Script field.
func (o *CreateNotificationRequest) SetScript(v string) {
	o.Script = &v
}

// GetIntermediateRef returns the IntermediateRef field value if set, zero value otherwise.
func (o *CreateNotificationRequest) GetIntermediateRef() string {
	if o == nil || IsNil(o.IntermediateRef) {
		var ret string
		return ret
	}
	return *o.IntermediateRef
}

// GetIntermediateRefOk returns a tuple with the IntermediateRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationRequest) GetIntermediateRefOk() (*string, bool) {
	if o == nil || IsNil(o.IntermediateRef) {
		return nil, false
	}
	return o.IntermediateRef, true
}

// HasIntermediateRef returns a boolean if a field has been set.
func (o *CreateNotificationRequest) HasIntermediateRef() bool {
	if o != nil && !IsNil(o.IntermediateRef) {
		return true
	}

	return false
}

// SetIntermediateRef gets a reference to the given string and assigns it to the IntermediateRef field.
func (o *CreateNotificationRequest) SetIntermediateRef(v string) {
	o.IntermediateRef = &v
}

// GetTemplateHtml returns the TemplateHtml field value
func (o *CreateNotificationRequest) GetTemplateHtml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateHtml
}

// GetTemplateHtmlOk returns a tuple with the TemplateHtml field value
// and a boolean to check if the value has been set.
func (o *CreateNotificationRequest) GetTemplateHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateHtml, true
}

// SetTemplateHtml sets field value
func (o *CreateNotificationRequest) SetTemplateHtml(v string) {
	o.TemplateHtml = v
}

// GetTemplateText returns the TemplateText field value
func (o *CreateNotificationRequest) GetTemplateText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateText
}

// GetTemplateTextOk returns a tuple with the TemplateText field value
// and a boolean to check if the value has been set.
func (o *CreateNotificationRequest) GetTemplateTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateText, true
}

// SetTemplateText sets field value
func (o *CreateNotificationRequest) SetTemplateText(v string) {
	o.TemplateText = v
}

func (o CreateNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.IsIntermediate) {
		toSerialize["is_intermediate"] = o.IsIntermediate
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Script) {
		toSerialize["script"] = o.Script
	}
	if !IsNil(o.IntermediateRef) {
		toSerialize["intermediate_ref"] = o.IntermediateRef
	}
	toSerialize["template_html"] = o.TemplateHtml
	toSerialize["template_text"] = o.TemplateText
	return toSerialize, nil
}

func (o *CreateNotificationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"template_html",
		"template_text",
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

	varCreateNotificationRequest := _CreateNotificationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNotificationRequest)

	if err != nil {
		return err
	}

	*o = CreateNotificationRequest(varCreateNotificationRequest)

	return err
}

type NullableCreateNotificationRequest struct {
	value *CreateNotificationRequest
	isSet bool
}

func (v NullableCreateNotificationRequest) Get() *CreateNotificationRequest {
	return v.value
}

func (v *NullableCreateNotificationRequest) Set(val *CreateNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNotificationRequest(val *CreateNotificationRequest) *NullableCreateNotificationRequest {
	return &NullableCreateNotificationRequest{value: val, isSet: true}
}

func (v NullableCreateNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
