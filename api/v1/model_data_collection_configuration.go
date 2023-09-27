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

// DataCollectionConfiguration the configuration of the data collection setup
type DataCollectionConfiguration struct {
	// if true, the default endpoint will be enabled for this setup
	DefaultEndpointEnabled *bool `json:"default_endpoint_enabled,omitempty"`
	// the code that is used for the default endpoint
	DefaultEndpointCode *string              `json:"default_endpoint_code,omitempty"`
	Inputs              []DefaultInputConfig `json:"inputs,omitempty"`
	AdditionalQuestions *CustomForm          `json:"additional_questions,omitempty"`
}

// NewDataCollectionConfiguration instantiates a new DataCollectionConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataCollectionConfiguration() *DataCollectionConfiguration {
	this := DataCollectionConfiguration{}
	return &this
}

// NewDataCollectionConfigurationWithDefaults instantiates a new DataCollectionConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataCollectionConfigurationWithDefaults() *DataCollectionConfiguration {
	this := DataCollectionConfiguration{}
	return &this
}

// GetDefaultEndpointEnabled returns the DefaultEndpointEnabled field value if set, zero value otherwise.
func (o *DataCollectionConfiguration) GetDefaultEndpointEnabled() bool {
	if o == nil || o.DefaultEndpointEnabled == nil {
		var ret bool
		return ret
	}
	return *o.DefaultEndpointEnabled
}

// GetDefaultEndpointEnabledOk returns a tuple with the DefaultEndpointEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCollectionConfiguration) GetDefaultEndpointEnabledOk() (*bool, bool) {
	if o == nil || o.DefaultEndpointEnabled == nil {
		return nil, false
	}
	return o.DefaultEndpointEnabled, true
}

// HasDefaultEndpointEnabled returns a boolean if a field has been set.
func (o *DataCollectionConfiguration) HasDefaultEndpointEnabled() bool {
	if o != nil && o.DefaultEndpointEnabled != nil {
		return true
	}

	return false
}

// SetDefaultEndpointEnabled gets a reference to the given bool and assigns it to the DefaultEndpointEnabled field.
func (o *DataCollectionConfiguration) SetDefaultEndpointEnabled(v bool) {
	o.DefaultEndpointEnabled = &v
}

// GetDefaultEndpointCode returns the DefaultEndpointCode field value if set, zero value otherwise.
func (o *DataCollectionConfiguration) GetDefaultEndpointCode() string {
	if o == nil || o.DefaultEndpointCode == nil {
		var ret string
		return ret
	}
	return *o.DefaultEndpointCode
}

// GetDefaultEndpointCodeOk returns a tuple with the DefaultEndpointCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCollectionConfiguration) GetDefaultEndpointCodeOk() (*string, bool) {
	if o == nil || o.DefaultEndpointCode == nil {
		return nil, false
	}
	return o.DefaultEndpointCode, true
}

// HasDefaultEndpointCode returns a boolean if a field has been set.
func (o *DataCollectionConfiguration) HasDefaultEndpointCode() bool {
	if o != nil && o.DefaultEndpointCode != nil {
		return true
	}

	return false
}

// SetDefaultEndpointCode gets a reference to the given string and assigns it to the DefaultEndpointCode field.
func (o *DataCollectionConfiguration) SetDefaultEndpointCode(v string) {
	o.DefaultEndpointCode = &v
}

// GetInputs returns the Inputs field value if set, zero value otherwise.
func (o *DataCollectionConfiguration) GetInputs() []DefaultInputConfig {
	if o == nil || o.Inputs == nil {
		var ret []DefaultInputConfig
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCollectionConfiguration) GetInputsOk() ([]DefaultInputConfig, bool) {
	if o == nil || o.Inputs == nil {
		return nil, false
	}
	return o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *DataCollectionConfiguration) HasInputs() bool {
	if o != nil && o.Inputs != nil {
		return true
	}

	return false
}

// SetInputs gets a reference to the given []DefaultInputConfig and assigns it to the Inputs field.
func (o *DataCollectionConfiguration) SetInputs(v []DefaultInputConfig) {
	o.Inputs = v
}

// GetAdditionalQuestions returns the AdditionalQuestions field value if set, zero value otherwise.
func (o *DataCollectionConfiguration) GetAdditionalQuestions() CustomForm {
	if o == nil || o.AdditionalQuestions == nil {
		var ret CustomForm
		return ret
	}
	return *o.AdditionalQuestions
}

// GetAdditionalQuestionsOk returns a tuple with the AdditionalQuestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataCollectionConfiguration) GetAdditionalQuestionsOk() (*CustomForm, bool) {
	if o == nil || o.AdditionalQuestions == nil {
		return nil, false
	}
	return o.AdditionalQuestions, true
}

// HasAdditionalQuestions returns a boolean if a field has been set.
func (o *DataCollectionConfiguration) HasAdditionalQuestions() bool {
	if o != nil && o.AdditionalQuestions != nil {
		return true
	}

	return false
}

// SetAdditionalQuestions gets a reference to the given CustomForm and assigns it to the AdditionalQuestions field.
func (o *DataCollectionConfiguration) SetAdditionalQuestions(v CustomForm) {
	o.AdditionalQuestions = &v
}

func (o DataCollectionConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultEndpointEnabled != nil {
		toSerialize["default_endpoint_enabled"] = o.DefaultEndpointEnabled
	}
	if o.DefaultEndpointCode != nil {
		toSerialize["default_endpoint_code"] = o.DefaultEndpointCode
	}
	if o.Inputs != nil {
		toSerialize["inputs"] = o.Inputs
	}
	if o.AdditionalQuestions != nil {
		toSerialize["additional_questions"] = o.AdditionalQuestions
	}
	return json.Marshal(toSerialize)
}

type NullableDataCollectionConfiguration struct {
	value *DataCollectionConfiguration
	isSet bool
}

func (v NullableDataCollectionConfiguration) Get() *DataCollectionConfiguration {
	return v.value
}

func (v *NullableDataCollectionConfiguration) Set(val *DataCollectionConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableDataCollectionConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableDataCollectionConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataCollectionConfiguration(val *DataCollectionConfiguration) *NullableDataCollectionConfiguration {
	return &NullableDataCollectionConfiguration{value: val, isSet: true}
}

func (v NullableDataCollectionConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataCollectionConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
