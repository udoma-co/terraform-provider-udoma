/*
Udoma

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// checks if the DocumentTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentTemplate{}

// DocumentTemplate A template for generating documents
type DocumentTemplate struct {
	// The unique ID of the template
	Id *string `json:"id,omitempty"`
	// The name of the template
	Name *string `json:"name,omitempty"`
	// A description of the template
	Description *string                  `json:"description,omitempty"`
	Options     *DocumentTemplateOptions `json:"options,omitempty"`
	// An optional JS expression to be used to compute the name of the template. If not set, the name of the template will be used for new documents.
	NameExpression *string `json:"name_expression,omitempty"`
	// The source of the template, used to generate the document
	Content *string     `json:"content,omitempty"`
	Inputs  *CustomForm `json:"inputs,omitempty"`
	// The script we run to generate the object used in the template
	PlaceholdersScript *string                                `json:"placeholders_script,omitempty"`
	Signatures         *DocumentTemplateSignatureConfguration `json:"signatures,omitempty"`
}

// NewDocumentTemplate instantiates a new DocumentTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentTemplate() *DocumentTemplate {
	this := DocumentTemplate{}
	return &this
}

// NewDocumentTemplateWithDefaults instantiates a new DocumentTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentTemplateWithDefaults() *DocumentTemplate {
	this := DocumentTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DocumentTemplate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentTemplate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DocumentTemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DocumentTemplate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DocumentTemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DocumentTemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DocumentTemplate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DocumentTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DocumentTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DocumentTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *DocumentTemplate) GetOptions() DocumentTemplateOptions {
	if o == nil || IsNil(o.Options) {
		var ret DocumentTemplateOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentTemplate) GetOptionsOk() (*DocumentTemplateOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *DocumentTemplate) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given DocumentTemplateOptions and assigns it to the Options field.
func (o *DocumentTemplate) SetOptions(v DocumentTemplateOptions) {
	o.Options = &v
}

// GetNameExpression returns the NameExpression field value if set, zero value otherwise.
func (o *DocumentTemplate) GetNameExpression() string {
	if o == nil || IsNil(o.NameExpression) {
		var ret string
		return ret
	}
	return *o.NameExpression
}

// GetNameExpressionOk returns a tuple with the NameExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentTemplate) GetNameExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.NameExpression) {
		return nil, false
	}
	return o.NameExpression, true
}

// HasNameExpression returns a boolean if a field has been set.
func (o *DocumentTemplate) HasNameExpression() bool {
	if o != nil && !IsNil(o.NameExpression) {
		return true
	}

	return false
}

// SetNameExpression gets a reference to the given string and assigns it to the NameExpression field.
func (o *DocumentTemplate) SetNameExpression(v string) {
	o.NameExpression = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *DocumentTemplate) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentTemplate) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *DocumentTemplate) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *DocumentTemplate) SetContent(v string) {
	o.Content = &v
}

// GetInputs returns the Inputs field value if set, zero value otherwise.
func (o *DocumentTemplate) GetInputs() CustomForm {
	if o == nil || IsNil(o.Inputs) {
		var ret CustomForm
		return ret
	}
	return *o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentTemplate) GetInputsOk() (*CustomForm, bool) {
	if o == nil || IsNil(o.Inputs) {
		return nil, false
	}
	return o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *DocumentTemplate) HasInputs() bool {
	if o != nil && !IsNil(o.Inputs) {
		return true
	}

	return false
}

// SetInputs gets a reference to the given CustomForm and assigns it to the Inputs field.
func (o *DocumentTemplate) SetInputs(v CustomForm) {
	o.Inputs = &v
}

// GetPlaceholdersScript returns the PlaceholdersScript field value if set, zero value otherwise.
func (o *DocumentTemplate) GetPlaceholdersScript() string {
	if o == nil || IsNil(o.PlaceholdersScript) {
		var ret string
		return ret
	}
	return *o.PlaceholdersScript
}

// GetPlaceholdersScriptOk returns a tuple with the PlaceholdersScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentTemplate) GetPlaceholdersScriptOk() (*string, bool) {
	if o == nil || IsNil(o.PlaceholdersScript) {
		return nil, false
	}
	return o.PlaceholdersScript, true
}

// HasPlaceholdersScript returns a boolean if a field has been set.
func (o *DocumentTemplate) HasPlaceholdersScript() bool {
	if o != nil && !IsNil(o.PlaceholdersScript) {
		return true
	}

	return false
}

// SetPlaceholdersScript gets a reference to the given string and assigns it to the PlaceholdersScript field.
func (o *DocumentTemplate) SetPlaceholdersScript(v string) {
	o.PlaceholdersScript = &v
}

// GetSignatures returns the Signatures field value if set, zero value otherwise.
func (o *DocumentTemplate) GetSignatures() DocumentTemplateSignatureConfguration {
	if o == nil || IsNil(o.Signatures) {
		var ret DocumentTemplateSignatureConfguration
		return ret
	}
	return *o.Signatures
}

// GetSignaturesOk returns a tuple with the Signatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentTemplate) GetSignaturesOk() (*DocumentTemplateSignatureConfguration, bool) {
	if o == nil || IsNil(o.Signatures) {
		return nil, false
	}
	return o.Signatures, true
}

// HasSignatures returns a boolean if a field has been set.
func (o *DocumentTemplate) HasSignatures() bool {
	if o != nil && !IsNil(o.Signatures) {
		return true
	}

	return false
}

// SetSignatures gets a reference to the given DocumentTemplateSignatureConfguration and assigns it to the Signatures field.
func (o *DocumentTemplate) SetSignatures(v DocumentTemplateSignatureConfguration) {
	o.Signatures = &v
}

func (o DocumentTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.NameExpression) {
		toSerialize["name_expression"] = o.NameExpression
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Inputs) {
		toSerialize["inputs"] = o.Inputs
	}
	if !IsNil(o.PlaceholdersScript) {
		toSerialize["placeholders_script"] = o.PlaceholdersScript
	}
	if !IsNil(o.Signatures) {
		toSerialize["signatures"] = o.Signatures
	}
	return toSerialize, nil
}

type NullableDocumentTemplate struct {
	value *DocumentTemplate
	isSet bool
}

func (v NullableDocumentTemplate) Get() *DocumentTemplate {
	return v.value
}

func (v *NullableDocumentTemplate) Set(val *DocumentTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentTemplate(val *DocumentTemplate) *NullableDocumentTemplate {
	return &NullableDocumentTemplate{value: val, isSet: true}
}

func (v NullableDocumentTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
