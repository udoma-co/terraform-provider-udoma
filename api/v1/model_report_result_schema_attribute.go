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

// checks if the ReportResultSchemaAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportResultSchemaAttribute{}

// ReportResultSchemaAttribute struct for ReportResultSchemaAttribute
type ReportResultSchemaAttribute struct {
	// The name of the attribute, via which it can be accessed from the result  object.
	Id string `json:"id"`
	// a map of values, where the key and values are strings
	Label map[string]string `json:"label"`
	// The sequence of the attribute, which will be used to order the attributes  in the UI.
	Sequence int32 `json:"sequence"`
	// Whether the attribute is used to link to another entity. If true, the  UI will parse the attribute and render a link to the referenced entity. If the result is a list, this attribute will be used as the target of each row, hence, it should be avoided to have more than one reference attribute in a table schema.
	IsReference *bool `json:"is_reference,omitempty"`
}

type _ReportResultSchemaAttribute ReportResultSchemaAttribute

// NewReportResultSchemaAttribute instantiates a new ReportResultSchemaAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportResultSchemaAttribute(id string, label map[string]string, sequence int32) *ReportResultSchemaAttribute {
	this := ReportResultSchemaAttribute{}
	this.Id = id
	this.Label = label
	this.Sequence = sequence
	return &this
}

// NewReportResultSchemaAttributeWithDefaults instantiates a new ReportResultSchemaAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportResultSchemaAttributeWithDefaults() *ReportResultSchemaAttribute {
	this := ReportResultSchemaAttribute{}
	return &this
}

// GetId returns the Id field value
func (o *ReportResultSchemaAttribute) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReportResultSchemaAttribute) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReportResultSchemaAttribute) SetId(v string) {
	o.Id = v
}

// GetLabel returns the Label field value
func (o *ReportResultSchemaAttribute) GetLabel() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ReportResultSchemaAttribute) GetLabelOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ReportResultSchemaAttribute) SetLabel(v map[string]string) {
	o.Label = v
}

// GetSequence returns the Sequence field value
func (o *ReportResultSchemaAttribute) GetSequence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Sequence
}

// GetSequenceOk returns a tuple with the Sequence field value
// and a boolean to check if the value has been set.
func (o *ReportResultSchemaAttribute) GetSequenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sequence, true
}

// SetSequence sets field value
func (o *ReportResultSchemaAttribute) SetSequence(v int32) {
	o.Sequence = v
}

// GetIsReference returns the IsReference field value if set, zero value otherwise.
func (o *ReportResultSchemaAttribute) GetIsReference() bool {
	if o == nil || IsNil(o.IsReference) {
		var ret bool
		return ret
	}
	return *o.IsReference
}

// GetIsReferenceOk returns a tuple with the IsReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResultSchemaAttribute) GetIsReferenceOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReference) {
		return nil, false
	}
	return o.IsReference, true
}

// HasIsReference returns a boolean if a field has been set.
func (o *ReportResultSchemaAttribute) HasIsReference() bool {
	if o != nil && !IsNil(o.IsReference) {
		return true
	}

	return false
}

// SetIsReference gets a reference to the given bool and assigns it to the IsReference field.
func (o *ReportResultSchemaAttribute) SetIsReference(v bool) {
	o.IsReference = &v
}

func (o ReportResultSchemaAttribute) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportResultSchemaAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["label"] = o.Label
	toSerialize["sequence"] = o.Sequence
	if !IsNil(o.IsReference) {
		toSerialize["is_reference"] = o.IsReference
	}
	return toSerialize, nil
}

func (o *ReportResultSchemaAttribute) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"label",
		"sequence",
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

	varReportResultSchemaAttribute := _ReportResultSchemaAttribute{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReportResultSchemaAttribute)

	if err != nil {
		return err
	}

	*o = ReportResultSchemaAttribute(varReportResultSchemaAttribute)

	return err
}

type NullableReportResultSchemaAttribute struct {
	value *ReportResultSchemaAttribute
	isSet bool
}

func (v NullableReportResultSchemaAttribute) Get() *ReportResultSchemaAttribute {
	return v.value
}

func (v *NullableReportResultSchemaAttribute) Set(val *ReportResultSchemaAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableReportResultSchemaAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableReportResultSchemaAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportResultSchemaAttribute(val *ReportResultSchemaAttribute) *NullableReportResultSchemaAttribute {
	return &NullableReportResultSchemaAttribute{value: val, isSet: true}
}

func (v NullableReportResultSchemaAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportResultSchemaAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
