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

// checks if the DocumentTemplateAttributesPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentTemplateAttributesPartial{}

// DocumentTemplateAttributesPartial Create or update a document template
type DocumentTemplateAttributesPartial struct {
	// The name of the template
	Name string `json:"name"`
	// A description of the template
	Description *string                         `json:"description,omitempty"`
	Options     NullableDocumentTemplateOptions `json:"options,omitempty"`
	// An optional JS expression to be used to compute the name of the template. If not set, the name of the template will be used for new documents.
	NameExpression *string `json:"name_expression,omitempty"`
	// The source of the template, used to generate the document
	Content string             `json:"content"`
	Inputs  NullableCustomForm `json:"inputs"`
	// The script we run to generate the object used in the template
	PlaceholdersScript *string                                       `json:"placeholders_script,omitempty"`
	Signatures         NullableDocumentTemplateSignatureConfguration `json:"signatures,omitempty"`
}

type _DocumentTemplateAttributesPartial DocumentTemplateAttributesPartial

// NewDocumentTemplateAttributesPartial instantiates a new DocumentTemplateAttributesPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentTemplateAttributesPartial(name string, content string, inputs NullableCustomForm) *DocumentTemplateAttributesPartial {
	this := DocumentTemplateAttributesPartial{}
	this.Name = name
	this.Content = content
	this.Inputs = inputs
	return &this
}

// NewDocumentTemplateAttributesPartialWithDefaults instantiates a new DocumentTemplateAttributesPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentTemplateAttributesPartialWithDefaults() *DocumentTemplateAttributesPartial {
	this := DocumentTemplateAttributesPartial{}
	return &this
}

// GetName returns the Name field value
func (o *DocumentTemplateAttributesPartial) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DocumentTemplateAttributesPartial) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DocumentTemplateAttributesPartial) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DocumentTemplateAttributesPartial) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentTemplateAttributesPartial) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DocumentTemplateAttributesPartial) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DocumentTemplateAttributesPartial) SetDescription(v string) {
	o.Description = &v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentTemplateAttributesPartial) GetOptions() DocumentTemplateOptions {
	if o == nil || IsNil(o.Options.Get()) {
		var ret DocumentTemplateOptions
		return ret
	}
	return *o.Options.Get()
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentTemplateAttributesPartial) GetOptionsOk() (*DocumentTemplateOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options.Get(), o.Options.IsSet()
}

// HasOptions returns a boolean if a field has been set.
func (o *DocumentTemplateAttributesPartial) HasOptions() bool {
	if o != nil && o.Options.IsSet() {
		return true
	}

	return false
}

// SetOptions gets a reference to the given NullableDocumentTemplateOptions and assigns it to the Options field.
func (o *DocumentTemplateAttributesPartial) SetOptions(v DocumentTemplateOptions) {
	o.Options.Set(&v)
}

// SetOptionsNil sets the value for Options to be an explicit nil
func (o *DocumentTemplateAttributesPartial) SetOptionsNil() {
	o.Options.Set(nil)
}

// UnsetOptions ensures that no value is present for Options, not even an explicit nil
func (o *DocumentTemplateAttributesPartial) UnsetOptions() {
	o.Options.Unset()
}

// GetNameExpression returns the NameExpression field value if set, zero value otherwise.
func (o *DocumentTemplateAttributesPartial) GetNameExpression() string {
	if o == nil || IsNil(o.NameExpression) {
		var ret string
		return ret
	}
	return *o.NameExpression
}

// GetNameExpressionOk returns a tuple with the NameExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentTemplateAttributesPartial) GetNameExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.NameExpression) {
		return nil, false
	}
	return o.NameExpression, true
}

// HasNameExpression returns a boolean if a field has been set.
func (o *DocumentTemplateAttributesPartial) HasNameExpression() bool {
	if o != nil && !IsNil(o.NameExpression) {
		return true
	}

	return false
}

// SetNameExpression gets a reference to the given string and assigns it to the NameExpression field.
func (o *DocumentTemplateAttributesPartial) SetNameExpression(v string) {
	o.NameExpression = &v
}

// GetContent returns the Content field value
func (o *DocumentTemplateAttributesPartial) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *DocumentTemplateAttributesPartial) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *DocumentTemplateAttributesPartial) SetContent(v string) {
	o.Content = v
}

// GetInputs returns the Inputs field value
// If the value is explicit nil, the zero value for CustomForm will be returned
func (o *DocumentTemplateAttributesPartial) GetInputs() CustomForm {
	if o == nil || o.Inputs.Get() == nil {
		var ret CustomForm
		return ret
	}

	return *o.Inputs.Get()
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentTemplateAttributesPartial) GetInputsOk() (*CustomForm, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs.Get(), o.Inputs.IsSet()
}

// SetInputs sets field value
func (o *DocumentTemplateAttributesPartial) SetInputs(v CustomForm) {
	o.Inputs.Set(&v)
}

// GetPlaceholdersScript returns the PlaceholdersScript field value if set, zero value otherwise.
func (o *DocumentTemplateAttributesPartial) GetPlaceholdersScript() string {
	if o == nil || IsNil(o.PlaceholdersScript) {
		var ret string
		return ret
	}
	return *o.PlaceholdersScript
}

// GetPlaceholdersScriptOk returns a tuple with the PlaceholdersScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentTemplateAttributesPartial) GetPlaceholdersScriptOk() (*string, bool) {
	if o == nil || IsNil(o.PlaceholdersScript) {
		return nil, false
	}
	return o.PlaceholdersScript, true
}

// HasPlaceholdersScript returns a boolean if a field has been set.
func (o *DocumentTemplateAttributesPartial) HasPlaceholdersScript() bool {
	if o != nil && !IsNil(o.PlaceholdersScript) {
		return true
	}

	return false
}

// SetPlaceholdersScript gets a reference to the given string and assigns it to the PlaceholdersScript field.
func (o *DocumentTemplateAttributesPartial) SetPlaceholdersScript(v string) {
	o.PlaceholdersScript = &v
}

// GetSignatures returns the Signatures field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DocumentTemplateAttributesPartial) GetSignatures() DocumentTemplateSignatureConfguration {
	if o == nil || IsNil(o.Signatures.Get()) {
		var ret DocumentTemplateSignatureConfguration
		return ret
	}
	return *o.Signatures.Get()
}

// GetSignaturesOk returns a tuple with the Signatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DocumentTemplateAttributesPartial) GetSignaturesOk() (*DocumentTemplateSignatureConfguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signatures.Get(), o.Signatures.IsSet()
}

// HasSignatures returns a boolean if a field has been set.
func (o *DocumentTemplateAttributesPartial) HasSignatures() bool {
	if o != nil && o.Signatures.IsSet() {
		return true
	}

	return false
}

// SetSignatures gets a reference to the given NullableDocumentTemplateSignatureConfguration and assigns it to the Signatures field.
func (o *DocumentTemplateAttributesPartial) SetSignatures(v DocumentTemplateSignatureConfguration) {
	o.Signatures.Set(&v)
}

// SetSignaturesNil sets the value for Signatures to be an explicit nil
func (o *DocumentTemplateAttributesPartial) SetSignaturesNil() {
	o.Signatures.Set(nil)
}

// UnsetSignatures ensures that no value is present for Signatures, not even an explicit nil
func (o *DocumentTemplateAttributesPartial) UnsetSignatures() {
	o.Signatures.Unset()
}

func (o DocumentTemplateAttributesPartial) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentTemplateAttributesPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if o.Options.IsSet() {
		toSerialize["options"] = o.Options.Get()
	}
	if !IsNil(o.NameExpression) {
		toSerialize["name_expression"] = o.NameExpression
	}
	toSerialize["content"] = o.Content
	toSerialize["inputs"] = o.Inputs.Get()
	if !IsNil(o.PlaceholdersScript) {
		toSerialize["placeholders_script"] = o.PlaceholdersScript
	}
	if o.Signatures.IsSet() {
		toSerialize["signatures"] = o.Signatures.Get()
	}
	return toSerialize, nil
}

func (o *DocumentTemplateAttributesPartial) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"content",
		"inputs",
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

	varDocumentTemplateAttributesPartial := _DocumentTemplateAttributesPartial{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDocumentTemplateAttributesPartial)

	if err != nil {
		return err
	}

	*o = DocumentTemplateAttributesPartial(varDocumentTemplateAttributesPartial)

	return err
}

type NullableDocumentTemplateAttributesPartial struct {
	value *DocumentTemplateAttributesPartial
	isSet bool
}

func (v NullableDocumentTemplateAttributesPartial) Get() *DocumentTemplateAttributesPartial {
	return v.value
}

func (v *NullableDocumentTemplateAttributesPartial) Set(val *DocumentTemplateAttributesPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentTemplateAttributesPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentTemplateAttributesPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentTemplateAttributesPartial(val *DocumentTemplateAttributesPartial) *NullableDocumentTemplateAttributesPartial {
	return &NullableDocumentTemplateAttributesPartial{value: val, isSet: true}
}

func (v NullableDocumentTemplateAttributesPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentTemplateAttributesPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}