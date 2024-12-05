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

// checks if the FormInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormInput{}

// FormInput a custom input that is used in dynamic forms to collect data from the user
type FormInput struct {
	// The ID of the input field, used to identify it within the form. If no  target value is set, this will also be how the value is stored and later accessed in the collected data.
	Id string `json:"id"`
	// Optional attribute name to use when exporting the value of this input. If not set, the ID will be used.
	Target *string       `json:"target,omitempty"`
	Type   FormInputType `json:"type"`
	// a map of values, where the key and values are strings
	Attributes *map[string]string `json:"attributes,omitempty"`
	// a map of values, where the key and values are strings
	Label *map[string]string `json:"label,omitempty"`
	// a map of values, where the key and values are strings
	ViewLabel *map[string]string `json:"view_label,omitempty"`
	// a map of values, where the key and values are strings
	Placeholder *map[string]string `json:"placeholder,omitempty"`
	// Optional default value for the input field (as a JSON string)
	DefaultValue *string `json:"default_value,omitempty"`
	// If true, the user will be required to provide a valid value for this input
	Required *bool `json:"required,omitempty"`
	// If true, the value of the input will not be persisted. This is useful for checkboxes for accepting terms and conditions, etc.
	Ephemeral *bool `json:"ephemeral,omitempty"`
	// If true, changes to the input will be propagated to event listeners for  the custom form.
	PropagateChanges *bool `json:"propagate_changes,omitempty"`
	// Only used when the type is select or multi select. This is a list of  values that the user can choose from.
	Items []InputItem `json:"items,omitempty"`
}

type _FormInput FormInput

// NewFormInput instantiates a new FormInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormInput(id string, type_ FormInputType) *FormInput {
	this := FormInput{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewFormInputWithDefaults instantiates a new FormInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormInputWithDefaults() *FormInput {
	this := FormInput{}
	return &this
}

// GetId returns the Id field value
func (o *FormInput) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FormInput) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FormInput) SetId(v string) {
	o.Id = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *FormInput) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInput) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *FormInput) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *FormInput) SetTarget(v string) {
	o.Target = &v
}

// GetType returns the Type field value
func (o *FormInput) GetType() FormInputType {
	if o == nil {
		var ret FormInputType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FormInput) GetTypeOk() (*FormInputType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FormInput) SetType(v FormInputType) {
	o.Type = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *FormInput) GetAttributes() map[string]string {
	if o == nil || IsNil(o.Attributes) {
		var ret map[string]string
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInput) GetAttributesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *FormInput) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]string and assigns it to the Attributes field.
func (o *FormInput) SetAttributes(v map[string]string) {
	o.Attributes = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FormInput) GetLabel() map[string]string {
	if o == nil || IsNil(o.Label) {
		var ret map[string]string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInput) GetLabelOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FormInput) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given map[string]string and assigns it to the Label field.
func (o *FormInput) SetLabel(v map[string]string) {
	o.Label = &v
}

// GetViewLabel returns the ViewLabel field value if set, zero value otherwise.
func (o *FormInput) GetViewLabel() map[string]string {
	if o == nil || IsNil(o.ViewLabel) {
		var ret map[string]string
		return ret
	}
	return *o.ViewLabel
}

// GetViewLabelOk returns a tuple with the ViewLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInput) GetViewLabelOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ViewLabel) {
		return nil, false
	}
	return o.ViewLabel, true
}

// HasViewLabel returns a boolean if a field has been set.
func (o *FormInput) HasViewLabel() bool {
	if o != nil && !IsNil(o.ViewLabel) {
		return true
	}

	return false
}

// SetViewLabel gets a reference to the given map[string]string and assigns it to the ViewLabel field.
func (o *FormInput) SetViewLabel(v map[string]string) {
	o.ViewLabel = &v
}

// GetPlaceholder returns the Placeholder field value if set, zero value otherwise.
func (o *FormInput) GetPlaceholder() map[string]string {
	if o == nil || IsNil(o.Placeholder) {
		var ret map[string]string
		return ret
	}
	return *o.Placeholder
}

// GetPlaceholderOk returns a tuple with the Placeholder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInput) GetPlaceholderOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Placeholder) {
		return nil, false
	}
	return o.Placeholder, true
}

// HasPlaceholder returns a boolean if a field has been set.
func (o *FormInput) HasPlaceholder() bool {
	if o != nil && !IsNil(o.Placeholder) {
		return true
	}

	return false
}

// SetPlaceholder gets a reference to the given map[string]string and assigns it to the Placeholder field.
func (o *FormInput) SetPlaceholder(v map[string]string) {
	o.Placeholder = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *FormInput) GetDefaultValue() string {
	if o == nil || IsNil(o.DefaultValue) {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInput) GetDefaultValueOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *FormInput) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *FormInput) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *FormInput) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInput) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *FormInput) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *FormInput) SetRequired(v bool) {
	o.Required = &v
}

// GetEphemeral returns the Ephemeral field value if set, zero value otherwise.
func (o *FormInput) GetEphemeral() bool {
	if o == nil || IsNil(o.Ephemeral) {
		var ret bool
		return ret
	}
	return *o.Ephemeral
}

// GetEphemeralOk returns a tuple with the Ephemeral field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInput) GetEphemeralOk() (*bool, bool) {
	if o == nil || IsNil(o.Ephemeral) {
		return nil, false
	}
	return o.Ephemeral, true
}

// HasEphemeral returns a boolean if a field has been set.
func (o *FormInput) HasEphemeral() bool {
	if o != nil && !IsNil(o.Ephemeral) {
		return true
	}

	return false
}

// SetEphemeral gets a reference to the given bool and assigns it to the Ephemeral field.
func (o *FormInput) SetEphemeral(v bool) {
	o.Ephemeral = &v
}

// GetPropagateChanges returns the PropagateChanges field value if set, zero value otherwise.
func (o *FormInput) GetPropagateChanges() bool {
	if o == nil || IsNil(o.PropagateChanges) {
		var ret bool
		return ret
	}
	return *o.PropagateChanges
}

// GetPropagateChangesOk returns a tuple with the PropagateChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInput) GetPropagateChangesOk() (*bool, bool) {
	if o == nil || IsNil(o.PropagateChanges) {
		return nil, false
	}
	return o.PropagateChanges, true
}

// HasPropagateChanges returns a boolean if a field has been set.
func (o *FormInput) HasPropagateChanges() bool {
	if o != nil && !IsNil(o.PropagateChanges) {
		return true
	}

	return false
}

// SetPropagateChanges gets a reference to the given bool and assigns it to the PropagateChanges field.
func (o *FormInput) SetPropagateChanges(v bool) {
	o.PropagateChanges = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *FormInput) GetItems() []InputItem {
	if o == nil || IsNil(o.Items) {
		var ret []InputItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormInput) GetItemsOk() ([]InputItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *FormInput) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []InputItem and assigns it to the Items field.
func (o *FormInput) SetItems(v []InputItem) {
	o.Items = v
}

func (o FormInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.ViewLabel) {
		toSerialize["view_label"] = o.ViewLabel
	}
	if !IsNil(o.Placeholder) {
		toSerialize["placeholder"] = o.Placeholder
	}
	if !IsNil(o.DefaultValue) {
		toSerialize["default_value"] = o.DefaultValue
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.Ephemeral) {
		toSerialize["ephemeral"] = o.Ephemeral
	}
	if !IsNil(o.PropagateChanges) {
		toSerialize["propagate_changes"] = o.PropagateChanges
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

func (o *FormInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
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

	varFormInput := _FormInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFormInput)

	if err != nil {
		return err
	}

	*o = FormInput(varFormInput)

	return err
}

type NullableFormInput struct {
	value *FormInput
	isSet bool
}

func (v NullableFormInput) Get() *FormInput {
	return v.value
}

func (v *NullableFormInput) Set(val *FormInput) {
	v.value = val
	v.isSet = true
}

func (v NullableFormInput) IsSet() bool {
	return v.isSet
}

func (v *NullableFormInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormInput(val *FormInput) *NullableFormInput {
	return &NullableFormInput{value: val, isSet: true}
}

func (v NullableFormInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
