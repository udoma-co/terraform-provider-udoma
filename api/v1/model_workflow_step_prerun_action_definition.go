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

// checks if the WorkflowStepPrerunActionDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowStepPrerunActionDefinition{}

// WorkflowStepPrerunActionDefinition an action that can be performed before a workflow step is executed
type WorkflowStepPrerunActionDefinition struct {
	// the ID of the action
	Id string `json:"id"`
	// a parameter of a workflow step or step action. The value of the parameter is contextual and can vary in type and meaning depending on the step or action that uses it. If used in a step, the parameter will be available in the UI and will not be interpreted, i.e. JS expressions are not allowed. In actions however, the parameter might be interpreted as a JS expression, if the action type requires it.
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	// a map of values, where the key and values are strings
	DynamicParameters *map[string]string `json:"dynamic_parameters,omitempty"`
	// Optional JS expression that will be executed before any step- and action-specific logic is executed.
	ExecBefore *string `json:"exec_before,omitempty"`
	// Optional JS expression that will be executed after the step- and action-specific logic is executed.
	ExecAfter *string `json:"exec_after,omitempty"`
}

type _WorkflowStepPrerunActionDefinition WorkflowStepPrerunActionDefinition

// NewWorkflowStepPrerunActionDefinition instantiates a new WorkflowStepPrerunActionDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStepPrerunActionDefinition(id string) *WorkflowStepPrerunActionDefinition {
	this := WorkflowStepPrerunActionDefinition{}
	this.Id = id
	return &this
}

// NewWorkflowStepPrerunActionDefinitionWithDefaults instantiates a new WorkflowStepPrerunActionDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStepPrerunActionDefinitionWithDefaults() *WorkflowStepPrerunActionDefinition {
	this := WorkflowStepPrerunActionDefinition{}
	return &this
}

// GetId returns the Id field value
func (o *WorkflowStepPrerunActionDefinition) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkflowStepPrerunActionDefinition) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkflowStepPrerunActionDefinition) SetId(v string) {
	o.Id = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *WorkflowStepPrerunActionDefinition) GetParameters() map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepPrerunActionDefinition) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *WorkflowStepPrerunActionDefinition) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *WorkflowStepPrerunActionDefinition) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

// GetDynamicParameters returns the DynamicParameters field value if set, zero value otherwise.
func (o *WorkflowStepPrerunActionDefinition) GetDynamicParameters() map[string]string {
	if o == nil || IsNil(o.DynamicParameters) {
		var ret map[string]string
		return ret
	}
	return *o.DynamicParameters
}

// GetDynamicParametersOk returns a tuple with the DynamicParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepPrerunActionDefinition) GetDynamicParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.DynamicParameters) {
		return nil, false
	}
	return o.DynamicParameters, true
}

// HasDynamicParameters returns a boolean if a field has been set.
func (o *WorkflowStepPrerunActionDefinition) HasDynamicParameters() bool {
	if o != nil && !IsNil(o.DynamicParameters) {
		return true
	}

	return false
}

// SetDynamicParameters gets a reference to the given map[string]string and assigns it to the DynamicParameters field.
func (o *WorkflowStepPrerunActionDefinition) SetDynamicParameters(v map[string]string) {
	o.DynamicParameters = &v
}

// GetExecBefore returns the ExecBefore field value if set, zero value otherwise.
func (o *WorkflowStepPrerunActionDefinition) GetExecBefore() string {
	if o == nil || IsNil(o.ExecBefore) {
		var ret string
		return ret
	}
	return *o.ExecBefore
}

// GetExecBeforeOk returns a tuple with the ExecBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepPrerunActionDefinition) GetExecBeforeOk() (*string, bool) {
	if o == nil || IsNil(o.ExecBefore) {
		return nil, false
	}
	return o.ExecBefore, true
}

// HasExecBefore returns a boolean if a field has been set.
func (o *WorkflowStepPrerunActionDefinition) HasExecBefore() bool {
	if o != nil && !IsNil(o.ExecBefore) {
		return true
	}

	return false
}

// SetExecBefore gets a reference to the given string and assigns it to the ExecBefore field.
func (o *WorkflowStepPrerunActionDefinition) SetExecBefore(v string) {
	o.ExecBefore = &v
}

// GetExecAfter returns the ExecAfter field value if set, zero value otherwise.
func (o *WorkflowStepPrerunActionDefinition) GetExecAfter() string {
	if o == nil || IsNil(o.ExecAfter) {
		var ret string
		return ret
	}
	return *o.ExecAfter
}

// GetExecAfterOk returns a tuple with the ExecAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepPrerunActionDefinition) GetExecAfterOk() (*string, bool) {
	if o == nil || IsNil(o.ExecAfter) {
		return nil, false
	}
	return o.ExecAfter, true
}

// HasExecAfter returns a boolean if a field has been set.
func (o *WorkflowStepPrerunActionDefinition) HasExecAfter() bool {
	if o != nil && !IsNil(o.ExecAfter) {
		return true
	}

	return false
}

// SetExecAfter gets a reference to the given string and assigns it to the ExecAfter field.
func (o *WorkflowStepPrerunActionDefinition) SetExecAfter(v string) {
	o.ExecAfter = &v
}

func (o WorkflowStepPrerunActionDefinition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowStepPrerunActionDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.DynamicParameters) {
		toSerialize["dynamic_parameters"] = o.DynamicParameters
	}
	if !IsNil(o.ExecBefore) {
		toSerialize["exec_before"] = o.ExecBefore
	}
	if !IsNil(o.ExecAfter) {
		toSerialize["exec_after"] = o.ExecAfter
	}
	return toSerialize, nil
}

func (o *WorkflowStepPrerunActionDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varWorkflowStepPrerunActionDefinition := _WorkflowStepPrerunActionDefinition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkflowStepPrerunActionDefinition)

	if err != nil {
		return err
	}

	*o = WorkflowStepPrerunActionDefinition(varWorkflowStepPrerunActionDefinition)

	return err
}

type NullableWorkflowStepPrerunActionDefinition struct {
	value *WorkflowStepPrerunActionDefinition
	isSet bool
}

func (v NullableWorkflowStepPrerunActionDefinition) Get() *WorkflowStepPrerunActionDefinition {
	return v.value
}

func (v *NullableWorkflowStepPrerunActionDefinition) Set(val *WorkflowStepPrerunActionDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStepPrerunActionDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStepPrerunActionDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStepPrerunActionDefinition(val *WorkflowStepPrerunActionDefinition) *NullableWorkflowStepPrerunActionDefinition {
	return &NullableWorkflowStepPrerunActionDefinition{value: val, isSet: true}
}

func (v NullableWorkflowStepPrerunActionDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStepPrerunActionDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
