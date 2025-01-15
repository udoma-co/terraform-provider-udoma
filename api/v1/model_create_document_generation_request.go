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

// checks if the CreateDocumentGenerationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDocumentGenerationRequest{}

// CreateDocumentGenerationRequest Start a new document generation based on the given template. This will also trigger the initialisation of the inputs, defined in the template and store the result in the new generation object.
type CreateDocumentGenerationRequest struct {
	// The template based on which a document shall be generated
	TemplateRef string `json:"template_ref"`
	// Optional initialisation data for the inputs of the template. If not set, the default values of the inputs will be used. To be provided as JSON serialised object, where the values are NOT JSON serialised themselves.
	InitData *string `json:"init_data,omitempty"`
	// Optional static data to be used in the generation process.
	StaticData *string `json:"static_data,omitempty"`
}

type _CreateDocumentGenerationRequest CreateDocumentGenerationRequest

// NewCreateDocumentGenerationRequest instantiates a new CreateDocumentGenerationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDocumentGenerationRequest(templateRef string) *CreateDocumentGenerationRequest {
	this := CreateDocumentGenerationRequest{}
	this.TemplateRef = templateRef
	return &this
}

// NewCreateDocumentGenerationRequestWithDefaults instantiates a new CreateDocumentGenerationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDocumentGenerationRequestWithDefaults() *CreateDocumentGenerationRequest {
	this := CreateDocumentGenerationRequest{}
	return &this
}

// GetTemplateRef returns the TemplateRef field value
func (o *CreateDocumentGenerationRequest) GetTemplateRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateRef
}

// GetTemplateRefOk returns a tuple with the TemplateRef field value
// and a boolean to check if the value has been set.
func (o *CreateDocumentGenerationRequest) GetTemplateRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateRef, true
}

// SetTemplateRef sets field value
func (o *CreateDocumentGenerationRequest) SetTemplateRef(v string) {
	o.TemplateRef = v
}

// GetInitData returns the InitData field value if set, zero value otherwise.
func (o *CreateDocumentGenerationRequest) GetInitData() string {
	if o == nil || IsNil(o.InitData) {
		var ret string
		return ret
	}
	return *o.InitData
}

// GetInitDataOk returns a tuple with the InitData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDocumentGenerationRequest) GetInitDataOk() (*string, bool) {
	if o == nil || IsNil(o.InitData) {
		return nil, false
	}
	return o.InitData, true
}

// HasInitData returns a boolean if a field has been set.
func (o *CreateDocumentGenerationRequest) HasInitData() bool {
	if o != nil && !IsNil(o.InitData) {
		return true
	}

	return false
}

// SetInitData gets a reference to the given string and assigns it to the InitData field.
func (o *CreateDocumentGenerationRequest) SetInitData(v string) {
	o.InitData = &v
}

// GetStaticData returns the StaticData field value if set, zero value otherwise.
func (o *CreateDocumentGenerationRequest) GetStaticData() string {
	if o == nil || IsNil(o.StaticData) {
		var ret string
		return ret
	}
	return *o.StaticData
}

// GetStaticDataOk returns a tuple with the StaticData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDocumentGenerationRequest) GetStaticDataOk() (*string, bool) {
	if o == nil || IsNil(o.StaticData) {
		return nil, false
	}
	return o.StaticData, true
}

// HasStaticData returns a boolean if a field has been set.
func (o *CreateDocumentGenerationRequest) HasStaticData() bool {
	if o != nil && !IsNil(o.StaticData) {
		return true
	}

	return false
}

// SetStaticData gets a reference to the given string and assigns it to the StaticData field.
func (o *CreateDocumentGenerationRequest) SetStaticData(v string) {
	o.StaticData = &v
}

func (o CreateDocumentGenerationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDocumentGenerationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["template_ref"] = o.TemplateRef
	if !IsNil(o.InitData) {
		toSerialize["init_data"] = o.InitData
	}
	if !IsNil(o.StaticData) {
		toSerialize["static_data"] = o.StaticData
	}
	return toSerialize, nil
}

func (o *CreateDocumentGenerationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"template_ref",
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

	varCreateDocumentGenerationRequest := _CreateDocumentGenerationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateDocumentGenerationRequest)

	if err != nil {
		return err
	}

	*o = CreateDocumentGenerationRequest(varCreateDocumentGenerationRequest)

	return err
}

type NullableCreateDocumentGenerationRequest struct {
	value *CreateDocumentGenerationRequest
	isSet bool
}

func (v NullableCreateDocumentGenerationRequest) Get() *CreateDocumentGenerationRequest {
	return v.value
}

func (v *NullableCreateDocumentGenerationRequest) Set(val *CreateDocumentGenerationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDocumentGenerationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDocumentGenerationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDocumentGenerationRequest(val *CreateDocumentGenerationRequest) *NullableCreateDocumentGenerationRequest {
	return &NullableCreateDocumentGenerationRequest{value: val, isSet: true}
}

func (v NullableCreateDocumentGenerationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDocumentGenerationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
