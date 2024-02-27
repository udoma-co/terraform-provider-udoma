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

// checks if the CaseTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaseTemplate{}

// CaseTemplate Contains the information necessary to create a report of a certain  type from a report endpoint.
type CaseTemplate struct {
	// The ID of the cases template
	Id *string `json:"id,omitempty"`
	// The timestamp of when the template was created
	CreatedAt *int64 `json:"created_at,omitempty"`
	// The timestamp of when the template was last updated
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// The name of the cases template (only used in admin pages). If not set, the name of the template will be used instead.
	Name *string `json:"name,omitempty"`
	// Optional JS expression used to generate the name of the actual case
	NameExpression *string `json:"name_expression,omitempty"`
	// a map of values, where the key and values are strings
	Label *map[string]string `json:"label,omitempty"`
	// a map of values, where the key and values are strings
	Description *map[string]string `json:"description,omitempty"`
	// a map of values, where the key and values are strings
	InfoText *map[string]string `json:"info_text,omitempty"`
	// The font-awesome icon to use for this template
	Icon         *string     `json:"icon,omitempty"`
	CustomInputs *CustomForm `json:"custom_inputs,omitempty"`
	Config       *CaseConfig `json:"config,omitempty"`
}

// NewCaseTemplate instantiates a new CaseTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaseTemplate() *CaseTemplate {
	this := CaseTemplate{}
	return &this
}

// NewCaseTemplateWithDefaults instantiates a new CaseTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaseTemplateWithDefaults() *CaseTemplate {
	this := CaseTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CaseTemplate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseTemplate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CaseTemplate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CaseTemplate) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CaseTemplate) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseTemplate) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CaseTemplate) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *CaseTemplate) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *CaseTemplate) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseTemplate) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CaseTemplate) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *CaseTemplate) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CaseTemplate) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseTemplate) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CaseTemplate) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CaseTemplate) SetName(v string) {
	o.Name = &v
}

// GetNameExpression returns the NameExpression field value if set, zero value otherwise.
func (o *CaseTemplate) GetNameExpression() string {
	if o == nil || IsNil(o.NameExpression) {
		var ret string
		return ret
	}
	return *o.NameExpression
}

// GetNameExpressionOk returns a tuple with the NameExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseTemplate) GetNameExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.NameExpression) {
		return nil, false
	}
	return o.NameExpression, true
}

// HasNameExpression returns a boolean if a field has been set.
func (o *CaseTemplate) HasNameExpression() bool {
	if o != nil && !IsNil(o.NameExpression) {
		return true
	}

	return false
}

// SetNameExpression gets a reference to the given string and assigns it to the NameExpression field.
func (o *CaseTemplate) SetNameExpression(v string) {
	o.NameExpression = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *CaseTemplate) GetLabel() map[string]string {
	if o == nil || IsNil(o.Label) {
		var ret map[string]string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseTemplate) GetLabelOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *CaseTemplate) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given map[string]string and assigns it to the Label field.
func (o *CaseTemplate) SetLabel(v map[string]string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CaseTemplate) GetDescription() map[string]string {
	if o == nil || IsNil(o.Description) {
		var ret map[string]string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseTemplate) GetDescriptionOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CaseTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given map[string]string and assigns it to the Description field.
func (o *CaseTemplate) SetDescription(v map[string]string) {
	o.Description = &v
}

// GetInfoText returns the InfoText field value if set, zero value otherwise.
func (o *CaseTemplate) GetInfoText() map[string]string {
	if o == nil || IsNil(o.InfoText) {
		var ret map[string]string
		return ret
	}
	return *o.InfoText
}

// GetInfoTextOk returns a tuple with the InfoText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseTemplate) GetInfoTextOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.InfoText) {
		return nil, false
	}
	return o.InfoText, true
}

// HasInfoText returns a boolean if a field has been set.
func (o *CaseTemplate) HasInfoText() bool {
	if o != nil && !IsNil(o.InfoText) {
		return true
	}

	return false
}

// SetInfoText gets a reference to the given map[string]string and assigns it to the InfoText field.
func (o *CaseTemplate) SetInfoText(v map[string]string) {
	o.InfoText = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *CaseTemplate) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseTemplate) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *CaseTemplate) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *CaseTemplate) SetIcon(v string) {
	o.Icon = &v
}

// GetCustomInputs returns the CustomInputs field value if set, zero value otherwise.
func (o *CaseTemplate) GetCustomInputs() CustomForm {
	if o == nil || IsNil(o.CustomInputs) {
		var ret CustomForm
		return ret
	}
	return *o.CustomInputs
}

// GetCustomInputsOk returns a tuple with the CustomInputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseTemplate) GetCustomInputsOk() (*CustomForm, bool) {
	if o == nil || IsNil(o.CustomInputs) {
		return nil, false
	}
	return o.CustomInputs, true
}

// HasCustomInputs returns a boolean if a field has been set.
func (o *CaseTemplate) HasCustomInputs() bool {
	if o != nil && !IsNil(o.CustomInputs) {
		return true
	}

	return false
}

// SetCustomInputs gets a reference to the given CustomForm and assigns it to the CustomInputs field.
func (o *CaseTemplate) SetCustomInputs(v CustomForm) {
	o.CustomInputs = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *CaseTemplate) GetConfig() CaseConfig {
	if o == nil || IsNil(o.Config) {
		var ret CaseConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseTemplate) GetConfigOk() (*CaseConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *CaseTemplate) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given CaseConfig and assigns it to the Config field.
func (o *CaseTemplate) SetConfig(v CaseConfig) {
	o.Config = &v
}

func (o CaseTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaseTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NameExpression) {
		toSerialize["name_expression"] = o.NameExpression
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.InfoText) {
		toSerialize["info_text"] = o.InfoText
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.CustomInputs) {
		toSerialize["custom_inputs"] = o.CustomInputs
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return toSerialize, nil
}

type NullableCaseTemplate struct {
	value *CaseTemplate
	isSet bool
}

func (v NullableCaseTemplate) Get() *CaseTemplate {
	return v.value
}

func (v *NullableCaseTemplate) Set(val *CaseTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableCaseTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableCaseTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaseTemplate(val *CaseTemplate) *NullableCaseTemplate {
	return &NullableCaseTemplate{value: val, isSet: true}
}

func (v NullableCaseTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaseTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
