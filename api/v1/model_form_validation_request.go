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

// checks if the FormValidationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormValidationRequest{}

// FormValidationRequest struct for FormValidationRequest
type FormValidationRequest struct {
	SourceType FormSourceType `json:"source_type"`
	// the ID of the source of the form that will be validated
	SourceId string `json:"source_id"`
	// the data that will be validated
	Data map[string]interface{} `json:"data"`
}

type _FormValidationRequest FormValidationRequest

// NewFormValidationRequest instantiates a new FormValidationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormValidationRequest(sourceType FormSourceType, sourceId string, data map[string]interface{}) *FormValidationRequest {
	this := FormValidationRequest{}
	this.SourceType = sourceType
	this.SourceId = sourceId
	this.Data = data
	return &this
}

// NewFormValidationRequestWithDefaults instantiates a new FormValidationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormValidationRequestWithDefaults() *FormValidationRequest {
	this := FormValidationRequest{}
	return &this
}

// GetSourceType returns the SourceType field value
func (o *FormValidationRequest) GetSourceType() FormSourceType {
	if o == nil {
		var ret FormSourceType
		return ret
	}

	return o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value
// and a boolean to check if the value has been set.
func (o *FormValidationRequest) GetSourceTypeOk() (*FormSourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceType, true
}

// SetSourceType sets field value
func (o *FormValidationRequest) SetSourceType(v FormSourceType) {
	o.SourceType = v
}

// GetSourceId returns the SourceId field value
func (o *FormValidationRequest) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *FormValidationRequest) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *FormValidationRequest) SetSourceId(v string) {
	o.SourceId = v
}

// GetData returns the Data field value
func (o *FormValidationRequest) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *FormValidationRequest) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *FormValidationRequest) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o FormValidationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormValidationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["source_type"] = o.SourceType
	toSerialize["source_id"] = o.SourceId
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *FormValidationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"source_type",
		"source_id",
		"data",
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

	varFormValidationRequest := _FormValidationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFormValidationRequest)

	if err != nil {
		return err
	}

	*o = FormValidationRequest(varFormValidationRequest)

	return err
}

type NullableFormValidationRequest struct {
	value *FormValidationRequest
	isSet bool
}

func (v NullableFormValidationRequest) Get() *FormValidationRequest {
	return v.value
}

func (v *NullableFormValidationRequest) Set(val *FormValidationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFormValidationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFormValidationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormValidationRequest(val *FormValidationRequest) *NullableFormValidationRequest {
	return &NullableFormValidationRequest{value: val, isSet: true}
}

func (v NullableFormValidationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormValidationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
