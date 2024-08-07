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

// checks if the FormValidationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormValidationError{}

// FormValidationError a custom validation that is used to validate data provided by the user
type FormValidationError struct {
	// a map of values, where the key and values are strings
	Message map[string]string `json:"message"`
	// the index of the input that failed validation (nesting is supported via dot notation)
	Target string `json:"target"`
}

type _FormValidationError FormValidationError

// NewFormValidationError instantiates a new FormValidationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormValidationError(message map[string]string, target string) *FormValidationError {
	this := FormValidationError{}
	this.Message = message
	this.Target = target
	return &this
}

// NewFormValidationErrorWithDefaults instantiates a new FormValidationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormValidationErrorWithDefaults() *FormValidationError {
	this := FormValidationError{}
	return &this
}

// GetMessage returns the Message field value
func (o *FormValidationError) GetMessage() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *FormValidationError) GetMessageOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *FormValidationError) SetMessage(v map[string]string) {
	o.Message = v
}

// GetTarget returns the Target field value
func (o *FormValidationError) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *FormValidationError) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *FormValidationError) SetTarget(v string) {
	o.Target = v
}

func (o FormValidationError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormValidationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	toSerialize["target"] = o.Target
	return toSerialize, nil
}

func (o *FormValidationError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"target",
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

	varFormValidationError := _FormValidationError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFormValidationError)

	if err != nil {
		return err
	}

	*o = FormValidationError(varFormValidationError)

	return err
}

type NullableFormValidationError struct {
	value *FormValidationError
	isSet bool
}

func (v NullableFormValidationError) Get() *FormValidationError {
	return v.value
}

func (v *NullableFormValidationError) Set(val *FormValidationError) {
	v.value = val
	v.isSet = true
}

func (v NullableFormValidationError) IsSet() bool {
	return v.isSet
}

func (v *NullableFormValidationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormValidationError(val *FormValidationError) *NullableFormValidationError {
	return &NullableFormValidationError{value: val, isSet: true}
}

func (v NullableFormValidationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormValidationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
