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

// checks if the WorkflowExecutionStep type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowExecutionStep{}

// WorkflowExecutionStep a single step within a workflow execution. This is derived from the step definition, however, dynamic data is computed and populated by the backend, before the result is  being sent to the client.
type WorkflowExecutionStep struct {
	// the ID of the step, unique within the workflow
	Id string `json:"id"`
	// the type of the step
	Type string `json:"type"`
	// The icon of the step (shown in the menu). If empty, the default icon  of the step type will be used.
	Icon *string `json:"icon,omitempty"`
	// The name of the step (shown as title and in the menu). If empty, the  default name of the step type will be used.
	Name string `json:"name"`
	// a parameter of a workflow step or step action. The value of the parameter is contextual and can vary in type and meaning depending on the step or action that uses it. If used in a step, the parameter will be available in the UI and will not be interpreted, i.e. JS expressions are not allowed. In actions however, the parameter might be interpreted as a JS expression, if the action type requires it.
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	// Optional name of the group of the step. If a group is provided, steps within the same group will be grouped together in the UI as a drawer.
	GroupName *string              `json:"group_name,omitempty"`
	Actions   []WorkflowStepAction `json:"actions"`
	// Indicates whether the step has been completed or not.
	IsCompleted *bool `json:"is_completed,omitempty"`
	// Indicates whether the step can be executed or not.
	CanBeExecuted *bool `json:"can_be_executed,omitempty"`
}

type _WorkflowExecutionStep WorkflowExecutionStep

// NewWorkflowExecutionStep instantiates a new WorkflowExecutionStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowExecutionStep(id string, type_ string, name string, actions []WorkflowStepAction) *WorkflowExecutionStep {
	this := WorkflowExecutionStep{}
	this.Id = id
	this.Type = type_
	this.Name = name
	this.Actions = actions
	return &this
}

// NewWorkflowExecutionStepWithDefaults instantiates a new WorkflowExecutionStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowExecutionStepWithDefaults() *WorkflowExecutionStep {
	this := WorkflowExecutionStep{}
	return &this
}

// GetId returns the Id field value
func (o *WorkflowExecutionStep) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionStep) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkflowExecutionStep) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *WorkflowExecutionStep) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionStep) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WorkflowExecutionStep) SetType(v string) {
	o.Type = v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *WorkflowExecutionStep) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionStep) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *WorkflowExecutionStep) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *WorkflowExecutionStep) SetIcon(v string) {
	o.Icon = &v
}

// GetName returns the Name field value
func (o *WorkflowExecutionStep) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionStep) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkflowExecutionStep) SetName(v string) {
	o.Name = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *WorkflowExecutionStep) GetParameters() map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionStep) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *WorkflowExecutionStep) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *WorkflowExecutionStep) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *WorkflowExecutionStep) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionStep) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *WorkflowExecutionStep) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *WorkflowExecutionStep) SetGroupName(v string) {
	o.GroupName = &v
}

// GetActions returns the Actions field value
func (o *WorkflowExecutionStep) GetActions() []WorkflowStepAction {
	if o == nil {
		var ret []WorkflowStepAction
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionStep) GetActionsOk() ([]WorkflowStepAction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *WorkflowExecutionStep) SetActions(v []WorkflowStepAction) {
	o.Actions = v
}

// GetIsCompleted returns the IsCompleted field value if set, zero value otherwise.
func (o *WorkflowExecutionStep) GetIsCompleted() bool {
	if o == nil || IsNil(o.IsCompleted) {
		var ret bool
		return ret
	}
	return *o.IsCompleted
}

// GetIsCompletedOk returns a tuple with the IsCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionStep) GetIsCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCompleted) {
		return nil, false
	}
	return o.IsCompleted, true
}

// HasIsCompleted returns a boolean if a field has been set.
func (o *WorkflowExecutionStep) HasIsCompleted() bool {
	if o != nil && !IsNil(o.IsCompleted) {
		return true
	}

	return false
}

// SetIsCompleted gets a reference to the given bool and assigns it to the IsCompleted field.
func (o *WorkflowExecutionStep) SetIsCompleted(v bool) {
	o.IsCompleted = &v
}

// GetCanBeExecuted returns the CanBeExecuted field value if set, zero value otherwise.
func (o *WorkflowExecutionStep) GetCanBeExecuted() bool {
	if o == nil || IsNil(o.CanBeExecuted) {
		var ret bool
		return ret
	}
	return *o.CanBeExecuted
}

// GetCanBeExecutedOk returns a tuple with the CanBeExecuted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecutionStep) GetCanBeExecutedOk() (*bool, bool) {
	if o == nil || IsNil(o.CanBeExecuted) {
		return nil, false
	}
	return o.CanBeExecuted, true
}

// HasCanBeExecuted returns a boolean if a field has been set.
func (o *WorkflowExecutionStep) HasCanBeExecuted() bool {
	if o != nil && !IsNil(o.CanBeExecuted) {
		return true
	}

	return false
}

// SetCanBeExecuted gets a reference to the given bool and assigns it to the CanBeExecuted field.
func (o *WorkflowExecutionStep) SetCanBeExecuted(v bool) {
	o.CanBeExecuted = &v
}

func (o WorkflowExecutionStep) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowExecutionStep) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.GroupName) {
		toSerialize["group_name"] = o.GroupName
	}
	toSerialize["actions"] = o.Actions
	if !IsNil(o.IsCompleted) {
		toSerialize["is_completed"] = o.IsCompleted
	}
	if !IsNil(o.CanBeExecuted) {
		toSerialize["can_be_executed"] = o.CanBeExecuted
	}
	return toSerialize, nil
}

func (o *WorkflowExecutionStep) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"name",
		"actions",
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

	varWorkflowExecutionStep := _WorkflowExecutionStep{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkflowExecutionStep)

	if err != nil {
		return err
	}

	*o = WorkflowExecutionStep(varWorkflowExecutionStep)

	return err
}

type NullableWorkflowExecutionStep struct {
	value *WorkflowExecutionStep
	isSet bool
}

func (v NullableWorkflowExecutionStep) Get() *WorkflowExecutionStep {
	return v.value
}

func (v *NullableWorkflowExecutionStep) Set(val *WorkflowExecutionStep) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowExecutionStep) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowExecutionStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowExecutionStep(val *WorkflowExecutionStep) *NullableWorkflowExecutionStep {
	return &NullableWorkflowExecutionStep{value: val, isSet: true}
}

func (v NullableWorkflowExecutionStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowExecutionStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
