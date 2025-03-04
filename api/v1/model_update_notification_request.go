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

// checks if the UpdateNotificationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateNotificationRequest{}

// UpdateNotificationRequest Update request for a notification.
type UpdateNotificationRequest struct {
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

type _UpdateNotificationRequest UpdateNotificationRequest

// NewUpdateNotificationRequest instantiates a new UpdateNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNotificationRequest(templateHtml string, templateText string) *UpdateNotificationRequest {
	this := UpdateNotificationRequest{}
	this.TemplateHtml = templateHtml
	this.TemplateText = templateText
	return &this
}

// NewUpdateNotificationRequestWithDefaults instantiates a new UpdateNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNotificationRequestWithDefaults() *UpdateNotificationRequest {
	this := UpdateNotificationRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateNotificationRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotificationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateNotificationRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateNotificationRequest) SetDescription(v string) {
	o.Description = &v
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *UpdateNotificationRequest) GetScript() string {
	if o == nil || IsNil(o.Script) {
		var ret string
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotificationRequest) GetScriptOk() (*string, bool) {
	if o == nil || IsNil(o.Script) {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *UpdateNotificationRequest) HasScript() bool {
	if o != nil && !IsNil(o.Script) {
		return true
	}

	return false
}

// SetScript gets a reference to the given string and assigns it to the Script field.
func (o *UpdateNotificationRequest) SetScript(v string) {
	o.Script = &v
}

// GetIntermediateRef returns the IntermediateRef field value if set, zero value otherwise.
func (o *UpdateNotificationRequest) GetIntermediateRef() string {
	if o == nil || IsNil(o.IntermediateRef) {
		var ret string
		return ret
	}
	return *o.IntermediateRef
}

// GetIntermediateRefOk returns a tuple with the IntermediateRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotificationRequest) GetIntermediateRefOk() (*string, bool) {
	if o == nil || IsNil(o.IntermediateRef) {
		return nil, false
	}
	return o.IntermediateRef, true
}

// HasIntermediateRef returns a boolean if a field has been set.
func (o *UpdateNotificationRequest) HasIntermediateRef() bool {
	if o != nil && !IsNil(o.IntermediateRef) {
		return true
	}

	return false
}

// SetIntermediateRef gets a reference to the given string and assigns it to the IntermediateRef field.
func (o *UpdateNotificationRequest) SetIntermediateRef(v string) {
	o.IntermediateRef = &v
}

// GetTemplateHtml returns the TemplateHtml field value
func (o *UpdateNotificationRequest) GetTemplateHtml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateHtml
}

// GetTemplateHtmlOk returns a tuple with the TemplateHtml field value
// and a boolean to check if the value has been set.
func (o *UpdateNotificationRequest) GetTemplateHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateHtml, true
}

// SetTemplateHtml sets field value
func (o *UpdateNotificationRequest) SetTemplateHtml(v string) {
	o.TemplateHtml = v
}

// GetTemplateText returns the TemplateText field value
func (o *UpdateNotificationRequest) GetTemplateText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateText
}

// GetTemplateTextOk returns a tuple with the TemplateText field value
// and a boolean to check if the value has been set.
func (o *UpdateNotificationRequest) GetTemplateTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateText, true
}

// SetTemplateText sets field value
func (o *UpdateNotificationRequest) SetTemplateText(v string) {
	o.TemplateText = v
}

func (o UpdateNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *UpdateNotificationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varUpdateNotificationRequest := _UpdateNotificationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateNotificationRequest)

	if err != nil {
		return err
	}

	*o = UpdateNotificationRequest(varUpdateNotificationRequest)

	return err
}

type NullableUpdateNotificationRequest struct {
	value *UpdateNotificationRequest
	isSet bool
}

func (v NullableUpdateNotificationRequest) Get() *UpdateNotificationRequest {
	return v.value
}

func (v *NullableUpdateNotificationRequest) Set(val *UpdateNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNotificationRequest(val *UpdateNotificationRequest) *NullableUpdateNotificationRequest {
	return &NullableUpdateNotificationRequest{value: val, isSet: true}
}

func (v NullableUpdateNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
