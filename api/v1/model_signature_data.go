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

// checks if the SignatureData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignatureData{}

// SignatureData A signature of a user, consisting of a photo of the signature and the name of the user
type SignatureData struct {
	// the name of the user
	Name string     `json:"name"`
	Data Attachment `json:"data"`
}

type _SignatureData SignatureData

// NewSignatureData instantiates a new SignatureData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureData(name string, data Attachment) *SignatureData {
	this := SignatureData{}
	this.Name = name
	this.Data = data
	return &this
}

// NewSignatureDataWithDefaults instantiates a new SignatureData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureDataWithDefaults() *SignatureData {
	this := SignatureData{}
	return &this
}

// GetName returns the Name field value
func (o *SignatureData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SignatureData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SignatureData) SetName(v string) {
	o.Name = v
}

// GetData returns the Data field value
func (o *SignatureData) GetData() Attachment {
	if o == nil {
		var ret Attachment
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SignatureData) GetDataOk() (*Attachment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SignatureData) SetData(v Attachment) {
	o.Data = v
}

func (o SignatureData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignatureData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SignatureData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varSignatureData := _SignatureData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSignatureData)

	if err != nil {
		return err
	}

	*o = SignatureData(varSignatureData)

	return err
}

type NullableSignatureData struct {
	value *SignatureData
	isSet bool
}

func (v NullableSignatureData) Get() *SignatureData {
	return v.value
}

func (v *NullableSignatureData) Set(val *SignatureData) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureData) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureData(val *SignatureData) *NullableSignatureData {
	return &NullableSignatureData{value: val, isSet: true}
}

func (v NullableSignatureData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
