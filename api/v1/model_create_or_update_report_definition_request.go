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

// checks if the CreateOrUpdateReportDefinitionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateReportDefinitionRequest{}

// CreateOrUpdateReportDefinitionRequest struct for CreateOrUpdateReportDefinitionRequest
type CreateOrUpdateReportDefinitionRequest struct {
	// The name of the report
	Name string `json:"name"`
	// The description of the report
	Description  *string            `json:"description,omitempty"`
	ResultSchema ReportResultSchema `json:"result_schema"`
	Parameters   NullableCustomForm `json:"parameters,omitempty"`
	// The JS script to execute to generate the report result
	Script  string `json:"script"`
	Version *int32 `json:"version,omitempty"`
}

type _CreateOrUpdateReportDefinitionRequest CreateOrUpdateReportDefinitionRequest

// NewCreateOrUpdateReportDefinitionRequest instantiates a new CreateOrUpdateReportDefinitionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateReportDefinitionRequest(name string, resultSchema ReportResultSchema, script string) *CreateOrUpdateReportDefinitionRequest {
	this := CreateOrUpdateReportDefinitionRequest{}
	this.Name = name
	this.ResultSchema = resultSchema
	this.Script = script
	return &this
}

// NewCreateOrUpdateReportDefinitionRequestWithDefaults instantiates a new CreateOrUpdateReportDefinitionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateReportDefinitionRequestWithDefaults() *CreateOrUpdateReportDefinitionRequest {
	this := CreateOrUpdateReportDefinitionRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateOrUpdateReportDefinitionRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateReportDefinitionRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrUpdateReportDefinitionRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOrUpdateReportDefinitionRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateReportDefinitionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrUpdateReportDefinitionRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOrUpdateReportDefinitionRequest) SetDescription(v string) {
	o.Description = &v
}

// GetResultSchema returns the ResultSchema field value
func (o *CreateOrUpdateReportDefinitionRequest) GetResultSchema() ReportResultSchema {
	if o == nil {
		var ret ReportResultSchema
		return ret
	}

	return o.ResultSchema
}

// GetResultSchemaOk returns a tuple with the ResultSchema field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateReportDefinitionRequest) GetResultSchemaOk() (*ReportResultSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultSchema, true
}

// SetResultSchema sets field value
func (o *CreateOrUpdateReportDefinitionRequest) SetResultSchema(v ReportResultSchema) {
	o.ResultSchema = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateOrUpdateReportDefinitionRequest) GetParameters() CustomForm {
	if o == nil || IsNil(o.Parameters.Get()) {
		var ret CustomForm
		return ret
	}
	return *o.Parameters.Get()
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateOrUpdateReportDefinitionRequest) GetParametersOk() (*CustomForm, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parameters.Get(), o.Parameters.IsSet()
}

// HasParameters returns a boolean if a field has been set.
func (o *CreateOrUpdateReportDefinitionRequest) HasParameters() bool {
	if o != nil && o.Parameters.IsSet() {
		return true
	}

	return false
}

// SetParameters gets a reference to the given NullableCustomForm and assigns it to the Parameters field.
func (o *CreateOrUpdateReportDefinitionRequest) SetParameters(v CustomForm) {
	o.Parameters.Set(&v)
}

// SetParametersNil sets the value for Parameters to be an explicit nil
func (o *CreateOrUpdateReportDefinitionRequest) SetParametersNil() {
	o.Parameters.Set(nil)
}

// UnsetParameters ensures that no value is present for Parameters, not even an explicit nil
func (o *CreateOrUpdateReportDefinitionRequest) UnsetParameters() {
	o.Parameters.Unset()
}

// GetScript returns the Script field value
func (o *CreateOrUpdateReportDefinitionRequest) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateReportDefinitionRequest) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *CreateOrUpdateReportDefinitionRequest) SetScript(v string) {
	o.Script = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CreateOrUpdateReportDefinitionRequest) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateReportDefinitionRequest) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CreateOrUpdateReportDefinitionRequest) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *CreateOrUpdateReportDefinitionRequest) SetVersion(v int32) {
	o.Version = &v
}

func (o CreateOrUpdateReportDefinitionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateReportDefinitionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["result_schema"] = o.ResultSchema
	if o.Parameters.IsSet() {
		toSerialize["parameters"] = o.Parameters.Get()
	}
	toSerialize["script"] = o.Script
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

func (o *CreateOrUpdateReportDefinitionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"result_schema",
		"script",
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

	varCreateOrUpdateReportDefinitionRequest := _CreateOrUpdateReportDefinitionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrUpdateReportDefinitionRequest)

	if err != nil {
		return err
	}

	*o = CreateOrUpdateReportDefinitionRequest(varCreateOrUpdateReportDefinitionRequest)

	return err
}

type NullableCreateOrUpdateReportDefinitionRequest struct {
	value *CreateOrUpdateReportDefinitionRequest
	isSet bool
}

func (v NullableCreateOrUpdateReportDefinitionRequest) Get() *CreateOrUpdateReportDefinitionRequest {
	return v.value
}

func (v *NullableCreateOrUpdateReportDefinitionRequest) Set(val *CreateOrUpdateReportDefinitionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateReportDefinitionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateReportDefinitionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateReportDefinitionRequest(val *CreateOrUpdateReportDefinitionRequest) *NullableCreateOrUpdateReportDefinitionRequest {
	return &NullableCreateOrUpdateReportDefinitionRequest{value: val, isSet: true}
}

func (v NullableCreateOrUpdateReportDefinitionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateReportDefinitionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
