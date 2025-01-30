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

// checks if the CreateOrUpdateDocumentTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateDocumentTemplateRequest{}

// CreateOrUpdateDocumentTemplateRequest Create or update a document template
type CreateOrUpdateDocumentTemplateRequest struct {
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
	PlaceholdersScript *string                                        `json:"placeholders_script,omitempty"`
	Signatures         NullableDocumentTemplateSignatureConfiguration `json:"signatures,omitempty"`
	Version            *int32                                         `json:"version,omitempty"`
}

type _CreateOrUpdateDocumentTemplateRequest CreateOrUpdateDocumentTemplateRequest

// NewCreateOrUpdateDocumentTemplateRequest instantiates a new CreateOrUpdateDocumentTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateDocumentTemplateRequest(name string, content string, inputs NullableCustomForm) *CreateOrUpdateDocumentTemplateRequest {
	this := CreateOrUpdateDocumentTemplateRequest{}
	this.Name = name
	this.Content = content
	this.Inputs = inputs
	return &this
}

// NewCreateOrUpdateDocumentTemplateRequestWithDefaults instantiates a new CreateOrUpdateDocumentTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateDocumentTemplateRequestWithDefaults() *CreateOrUpdateDocumentTemplateRequest {
	this := CreateOrUpdateDocumentTemplateRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateOrUpdateDocumentTemplateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateDocumentTemplateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrUpdateDocumentTemplateRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOrUpdateDocumentTemplateRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateDocumentTemplateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrUpdateDocumentTemplateRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOrUpdateDocumentTemplateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetOptions returns the Options field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateDocumentTemplateRequest) GetOptions() DocumentTemplateOptions {
	if o == nil || IsNil(o.Options.Get()) {
		var ret DocumentTemplateOptions
		return ret
	}
	return *o.Options.Get()
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateDocumentTemplateRequest) GetOptionsOk() (*DocumentTemplateOptions, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options.Get(), o.Options.IsSet()
}

// HasOptions returns a boolean if a field has been set.
func (o *CreateOrUpdateDocumentTemplateRequest) HasOptions() bool {
	if o != nil && o.Options.IsSet() {
		return true
	}

	return false
}

// SetOptions gets a reference to the given NullableDocumentTemplateOptions and assigns it to the Options field.
func (o *CreateOrUpdateDocumentTemplateRequest) SetOptions(v DocumentTemplateOptions) {
	o.Options.Set(&v)
}

// SetOptionsNil sets the value for Options to be an explicit nil
func (o *CreateOrUpdateDocumentTemplateRequest) SetOptionsNil() {
	o.Options.Set(nil)
}

// UnsetOptions ensures that no value is present for Options, not even an explicit nil
func (o *CreateOrUpdateDocumentTemplateRequest) UnsetOptions() {
	o.Options.Unset()
}

// GetNameExpression returns the NameExpression field value if set, zero value otherwise.
func (o *CreateOrUpdateDocumentTemplateRequest) GetNameExpression() string {
	if o == nil || IsNil(o.NameExpression) {
		var ret string
		return ret
	}
	return *o.NameExpression
}

// GetNameExpressionOk returns a tuple with the NameExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateDocumentTemplateRequest) GetNameExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.NameExpression) {
		return nil, false
	}
	return o.NameExpression, true
}

// HasNameExpression returns a boolean if a field has been set.
func (o *CreateOrUpdateDocumentTemplateRequest) HasNameExpression() bool {
	if o != nil && !IsNil(o.NameExpression) {
		return true
	}

	return false
}

// SetNameExpression gets a reference to the given string and assigns it to the NameExpression field.
func (o *CreateOrUpdateDocumentTemplateRequest) SetNameExpression(v string) {
	o.NameExpression = &v
}

// GetContent returns the Content field value
func (o *CreateOrUpdateDocumentTemplateRequest) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateDocumentTemplateRequest) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *CreateOrUpdateDocumentTemplateRequest) SetContent(v string) {
	o.Content = v
}

// GetInputs returns the Inputs field value
// If the value is explicit nil, the zero value for CustomForm will be returned
func (o *CreateOrUpdateDocumentTemplateRequest) GetInputs() CustomForm {
	if o == nil || o.Inputs.Get() == nil {
		var ret CustomForm
		return ret
	}

	return *o.Inputs.Get()
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateDocumentTemplateRequest) GetInputsOk() (*CustomForm, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs.Get(), o.Inputs.IsSet()
}

// SetInputs sets field value
func (o *CreateOrUpdateDocumentTemplateRequest) SetInputs(v CustomForm) {
	o.Inputs.Set(&v)
}

// GetPlaceholdersScript returns the PlaceholdersScript field value if set, zero value otherwise.
func (o *CreateOrUpdateDocumentTemplateRequest) GetPlaceholdersScript() string {
	if o == nil || IsNil(o.PlaceholdersScript) {
		var ret string
		return ret
	}
	return *o.PlaceholdersScript
}

// GetPlaceholdersScriptOk returns a tuple with the PlaceholdersScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateDocumentTemplateRequest) GetPlaceholdersScriptOk() (*string, bool) {
	if o == nil || IsNil(o.PlaceholdersScript) {
		return nil, false
	}
	return o.PlaceholdersScript, true
}

// HasPlaceholdersScript returns a boolean if a field has been set.
func (o *CreateOrUpdateDocumentTemplateRequest) HasPlaceholdersScript() bool {
	if o != nil && !IsNil(o.PlaceholdersScript) {
		return true
	}

	return false
}

// SetPlaceholdersScript gets a reference to the given string and assigns it to the PlaceholdersScript field.
func (o *CreateOrUpdateDocumentTemplateRequest) SetPlaceholdersScript(v string) {
	o.PlaceholdersScript = &v
}

// GetSignatures returns the Signatures field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateDocumentTemplateRequest) GetSignatures() DocumentTemplateSignatureConfiguration {
	if o == nil || IsNil(o.Signatures.Get()) {
		var ret DocumentTemplateSignatureConfiguration
		return ret
	}
	return *o.Signatures.Get()
}

// GetSignaturesOk returns a tuple with the Signatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateDocumentTemplateRequest) GetSignaturesOk() (*DocumentTemplateSignatureConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signatures.Get(), o.Signatures.IsSet()
}

// HasSignatures returns a boolean if a field has been set.
func (o *CreateOrUpdateDocumentTemplateRequest) HasSignatures() bool {
	if o != nil && o.Signatures.IsSet() {
		return true
	}

	return false
}

// SetSignatures gets a reference to the given NullableDocumentTemplateSignatureConfiguration and assigns it to the Signatures field.
func (o *CreateOrUpdateDocumentTemplateRequest) SetSignatures(v DocumentTemplateSignatureConfiguration) {
	o.Signatures.Set(&v)
}

// SetSignaturesNil sets the value for Signatures to be an explicit nil
func (o *CreateOrUpdateDocumentTemplateRequest) SetSignaturesNil() {
	o.Signatures.Set(nil)
}

// UnsetSignatures ensures that no value is present for Signatures, not even an explicit nil
func (o *CreateOrUpdateDocumentTemplateRequest) UnsetSignatures() {
	o.Signatures.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CreateOrUpdateDocumentTemplateRequest) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateDocumentTemplateRequest) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CreateOrUpdateDocumentTemplateRequest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *CreateOrUpdateDocumentTemplateRequest) SetVersion(v int32) {
	o.Version = &v
}

func (o CreateOrUpdateDocumentTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateDocumentTemplateRequest) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

func (o *CreateOrUpdateDocumentTemplateRequest) UnmarshalJSON(data []byte) (err error) {
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

	varCreateOrUpdateDocumentTemplateRequest := _CreateOrUpdateDocumentTemplateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrUpdateDocumentTemplateRequest)

	if err != nil {
		return err
	}

	*o = CreateOrUpdateDocumentTemplateRequest(varCreateOrUpdateDocumentTemplateRequest)

	return err
}

type NullableCreateOrUpdateDocumentTemplateRequest struct {
	value *CreateOrUpdateDocumentTemplateRequest
	isSet bool
}

func (v NullableCreateOrUpdateDocumentTemplateRequest) Get() *CreateOrUpdateDocumentTemplateRequest {
	return v.value
}

func (v *NullableCreateOrUpdateDocumentTemplateRequest) Set(val *CreateOrUpdateDocumentTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateDocumentTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateDocumentTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateDocumentTemplateRequest(val *CreateOrUpdateDocumentTemplateRequest) *NullableCreateOrUpdateDocumentTemplateRequest {
	return &NullableCreateOrUpdateDocumentTemplateRequest{value: val, isSet: true}
}

func (v NullableCreateOrUpdateDocumentTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateDocumentTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
