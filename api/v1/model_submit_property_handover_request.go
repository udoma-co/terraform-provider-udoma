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

// checks if the SubmitPropertyHandoverRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitPropertyHandoverRequest{}

// SubmitPropertyHandoverRequest struct for SubmitPropertyHandoverRequest
type SubmitPropertyHandoverRequest struct {
	// The data captured by of the property handover.
	Data map[string]interface{} `json:"data"`
}

type _SubmitPropertyHandoverRequest SubmitPropertyHandoverRequest

// NewSubmitPropertyHandoverRequest instantiates a new SubmitPropertyHandoverRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitPropertyHandoverRequest(data map[string]interface{}) *SubmitPropertyHandoverRequest {
	this := SubmitPropertyHandoverRequest{}
	this.Data = data
	return &this
}

// NewSubmitPropertyHandoverRequestWithDefaults instantiates a new SubmitPropertyHandoverRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitPropertyHandoverRequestWithDefaults() *SubmitPropertyHandoverRequest {
	this := SubmitPropertyHandoverRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SubmitPropertyHandoverRequest) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubmitPropertyHandoverRequest) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *SubmitPropertyHandoverRequest) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o SubmitPropertyHandoverRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubmitPropertyHandoverRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SubmitPropertyHandoverRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varSubmitPropertyHandoverRequest := _SubmitPropertyHandoverRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubmitPropertyHandoverRequest)

	if err != nil {
		return err
	}

	*o = SubmitPropertyHandoverRequest(varSubmitPropertyHandoverRequest)

	return err
}

type NullableSubmitPropertyHandoverRequest struct {
	value *SubmitPropertyHandoverRequest
	isSet bool
}

func (v NullableSubmitPropertyHandoverRequest) Get() *SubmitPropertyHandoverRequest {
	return v.value
}

func (v *NullableSubmitPropertyHandoverRequest) Set(val *SubmitPropertyHandoverRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitPropertyHandoverRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitPropertyHandoverRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitPropertyHandoverRequest(val *SubmitPropertyHandoverRequest) *NullableSubmitPropertyHandoverRequest {
	return &NullableSubmitPropertyHandoverRequest{value: val, isSet: true}
}

func (v NullableSubmitPropertyHandoverRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitPropertyHandoverRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
