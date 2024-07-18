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

// checks if the Notification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Notification{}

// Notification A notification template that can be sent to a user.
type Notification struct {
	// A unique, descriptive, and short identifier name.
	Name string `json:"name"`
	// A short description of the notification.
	Description *string `json:"description,omitempty"`
	// The timestamp of when the notification was created.
	CreatedAt int64 `json:"created_at"`
	// The timestamp of when the notification was last updated.
	UpdatedAt int64 `json:"updated_at"`
	// A script that can run the notification given some initial data.
	Script string `json:"script"`
	// Whether this template is used an intermediate template or not. If true it cannot be used as a regular template and must only be referenced by others.
	IsIntermediate *bool `json:"is_intermediate,omitempty"`
	// A potential reference to an intermediate template. Empty if no intermediate used.
	IntermediateRef *string `json:"intermediate_ref,omitempty"`
	// The golang html template to use for the notification.
	TemplateHtml string `json:"template_html"`
	// The golang text template to use for the notification.
	TemplateText string `json:"template_text"`
}

type _Notification Notification

// NewNotification instantiates a new Notification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotification(name string, createdAt int64, updatedAt int64, script string, templateHtml string, templateText string) *Notification {
	this := Notification{}
	this.Name = name
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Script = script
	this.TemplateHtml = templateHtml
	this.TemplateText = templateText
	return &this
}

// NewNotificationWithDefaults instantiates a new Notification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationWithDefaults() *Notification {
	this := Notification{}
	return &this
}

// GetName returns the Name field value
func (o *Notification) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Notification) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Notification) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Notification) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Notification) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Notification) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Notification) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Notification) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Notification) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Notification) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Notification) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Notification) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetScript returns the Script field value
func (o *Notification) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *Notification) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *Notification) SetScript(v string) {
	o.Script = v
}

// GetIsIntermediate returns the IsIntermediate field value if set, zero value otherwise.
func (o *Notification) GetIsIntermediate() bool {
	if o == nil || IsNil(o.IsIntermediate) {
		var ret bool
		return ret
	}
	return *o.IsIntermediate
}

// GetIsIntermediateOk returns a tuple with the IsIntermediate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetIsIntermediateOk() (*bool, bool) {
	if o == nil || IsNil(o.IsIntermediate) {
		return nil, false
	}
	return o.IsIntermediate, true
}

// HasIsIntermediate returns a boolean if a field has been set.
func (o *Notification) HasIsIntermediate() bool {
	if o != nil && !IsNil(o.IsIntermediate) {
		return true
	}

	return false
}

// SetIsIntermediate gets a reference to the given bool and assigns it to the IsIntermediate field.
func (o *Notification) SetIsIntermediate(v bool) {
	o.IsIntermediate = &v
}

// GetIntermediateRef returns the IntermediateRef field value if set, zero value otherwise.
func (o *Notification) GetIntermediateRef() string {
	if o == nil || IsNil(o.IntermediateRef) {
		var ret string
		return ret
	}
	return *o.IntermediateRef
}

// GetIntermediateRefOk returns a tuple with the IntermediateRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Notification) GetIntermediateRefOk() (*string, bool) {
	if o == nil || IsNil(o.IntermediateRef) {
		return nil, false
	}
	return o.IntermediateRef, true
}

// HasIntermediateRef returns a boolean if a field has been set.
func (o *Notification) HasIntermediateRef() bool {
	if o != nil && !IsNil(o.IntermediateRef) {
		return true
	}

	return false
}

// SetIntermediateRef gets a reference to the given string and assigns it to the IntermediateRef field.
func (o *Notification) SetIntermediateRef(v string) {
	o.IntermediateRef = &v
}

// GetTemplateHtml returns the TemplateHtml field value
func (o *Notification) GetTemplateHtml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateHtml
}

// GetTemplateHtmlOk returns a tuple with the TemplateHtml field value
// and a boolean to check if the value has been set.
func (o *Notification) GetTemplateHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateHtml, true
}

// SetTemplateHtml sets field value
func (o *Notification) SetTemplateHtml(v string) {
	o.TemplateHtml = v
}

// GetTemplateText returns the TemplateText field value
func (o *Notification) GetTemplateText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateText
}

// GetTemplateTextOk returns a tuple with the TemplateText field value
// and a boolean to check if the value has been set.
func (o *Notification) GetTemplateTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateText, true
}

// SetTemplateText sets field value
func (o *Notification) SetTemplateText(v string) {
	o.TemplateText = v
}

func (o Notification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Notification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["script"] = o.Script
	if !IsNil(o.IsIntermediate) {
		toSerialize["is_intermediate"] = o.IsIntermediate
	}
	if !IsNil(o.IntermediateRef) {
		toSerialize["intermediate_ref"] = o.IntermediateRef
	}
	toSerialize["template_html"] = o.TemplateHtml
	toSerialize["template_text"] = o.TemplateText
	return toSerialize, nil
}

func (o *Notification) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"created_at",
		"updated_at",
		"script",
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

	varNotification := _Notification{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNotification)

	if err != nil {
		return err
	}

	*o = Notification(varNotification)

	return err
}

type NullableNotification struct {
	value *Notification
	isSet bool
}

func (v NullableNotification) Get() *Notification {
	return v.value
}

func (v *NullableNotification) Set(val *Notification) {
	v.value = val
	v.isSet = true
}

func (v NullableNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotification(val *Notification) *NullableNotification {
	return &NullableNotification{value: val, isSet: true}
}

func (v NullableNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
