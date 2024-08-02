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

// checks if the WorkflowStepDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowStepDefinition{}

// WorkflowStepDefinition the definition of a single step within a workflow
type WorkflowStepDefinition struct {
	// the ID of the step, unique within the workflow. This is used to handle  the flow of the workflow execution.
	Id string `json:"id"`
	// the type of the step, which defines how it will be rendered in the UI and how it will be executed in the backend.
	Type string `json:"type"`
	// a parameter of a workflow step or step action. The value of the parameter is contextual and can vary in type and meaning depending on the step or action that uses it. If used in a step, the parameter will be available in the UI and will not be interpreted, i.e. JS expressions are not allowed. In actions however, the parameter might be interpreted as a JS expression, if the action type requires it.
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	// a map of values, where the key and values are strings
	DynamicParameters *map[string]string `json:"dynamic_parameters,omitempty"`
	// The icon of the step (shown in the menu). If empty, the default icon  of the step type will be used.
	Icon *string `json:"icon,omitempty"`
	// The name of the step (shown as title and in the menu). If empty, the  default name of the step type will be used.
	Name string `json:"name"`
	// Optional name of the group of the step. If a group is provided, steps within the same group will be grouped together in the UI as a drawer.
	GroupName    *string                                    `json:"group_name,omitempty"`
	PrerunAction NullableWorkflowStepPrerunActionDefinition `json:"prerun_action,omitempty"`
	// An optional JS expression that determines whether the step can be executed or  not. If not set, this will default to true, once the previous step has been  executed.
	CanBeExecutedExpression *string                        `json:"can_be_executed_expression,omitempty"`
	Actions                 []WorkflowStepActionDefinition `json:"actions"`
}

type _WorkflowStepDefinition WorkflowStepDefinition

// NewWorkflowStepDefinition instantiates a new WorkflowStepDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStepDefinition(id string, type_ string, name string, actions []WorkflowStepActionDefinition) *WorkflowStepDefinition {
	this := WorkflowStepDefinition{}
	this.Id = id
	this.Type = type_
	this.Name = name
	this.Actions = actions
	return &this
}

// NewWorkflowStepDefinitionWithDefaults instantiates a new WorkflowStepDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStepDefinitionWithDefaults() *WorkflowStepDefinition {
	this := WorkflowStepDefinition{}
	return &this
}

// GetId returns the Id field value
func (o *WorkflowStepDefinition) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkflowStepDefinition) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkflowStepDefinition) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *WorkflowStepDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WorkflowStepDefinition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WorkflowStepDefinition) SetType(v string) {
	o.Type = v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *WorkflowStepDefinition) GetParameters() map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepDefinition) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *WorkflowStepDefinition) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *WorkflowStepDefinition) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

// GetDynamicParameters returns the DynamicParameters field value if set, zero value otherwise.
func (o *WorkflowStepDefinition) GetDynamicParameters() map[string]string {
	if o == nil || IsNil(o.DynamicParameters) {
		var ret map[string]string
		return ret
	}
	return *o.DynamicParameters
}

// GetDynamicParametersOk returns a tuple with the DynamicParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepDefinition) GetDynamicParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.DynamicParameters) {
		return nil, false
	}
	return o.DynamicParameters, true
}

// HasDynamicParameters returns a boolean if a field has been set.
func (o *WorkflowStepDefinition) HasDynamicParameters() bool {
	if o != nil && !IsNil(o.DynamicParameters) {
		return true
	}

	return false
}

// SetDynamicParameters gets a reference to the given map[string]string and assigns it to the DynamicParameters field.
func (o *WorkflowStepDefinition) SetDynamicParameters(v map[string]string) {
	o.DynamicParameters = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *WorkflowStepDefinition) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepDefinition) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *WorkflowStepDefinition) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *WorkflowStepDefinition) SetIcon(v string) {
	o.Icon = &v
}

// GetName returns the Name field value
func (o *WorkflowStepDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkflowStepDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkflowStepDefinition) SetName(v string) {
	o.Name = v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *WorkflowStepDefinition) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepDefinition) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *WorkflowStepDefinition) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *WorkflowStepDefinition) SetGroupName(v string) {
	o.GroupName = &v
}

// GetPrerunAction returns the PrerunAction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowStepDefinition) GetPrerunAction() WorkflowStepPrerunActionDefinition {
	if o == nil || IsNil(o.PrerunAction.Get()) {
		var ret WorkflowStepPrerunActionDefinition
		return ret
	}
	return *o.PrerunAction.Get()
}

// GetPrerunActionOk returns a tuple with the PrerunAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowStepDefinition) GetPrerunActionOk() (*WorkflowStepPrerunActionDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrerunAction.Get(), o.PrerunAction.IsSet()
}

// HasPrerunAction returns a boolean if a field has been set.
func (o *WorkflowStepDefinition) HasPrerunAction() bool {
	if o != nil && o.PrerunAction.IsSet() {
		return true
	}

	return false
}

// SetPrerunAction gets a reference to the given NullableWorkflowStepPrerunActionDefinition and assigns it to the PrerunAction field.
func (o *WorkflowStepDefinition) SetPrerunAction(v WorkflowStepPrerunActionDefinition) {
	o.PrerunAction.Set(&v)
}

// SetPrerunActionNil sets the value for PrerunAction to be an explicit nil
func (o *WorkflowStepDefinition) SetPrerunActionNil() {
	o.PrerunAction.Set(nil)
}

// UnsetPrerunAction ensures that no value is present for PrerunAction, not even an explicit nil
func (o *WorkflowStepDefinition) UnsetPrerunAction() {
	o.PrerunAction.Unset()
}

// GetCanBeExecutedExpression returns the CanBeExecutedExpression field value if set, zero value otherwise.
func (o *WorkflowStepDefinition) GetCanBeExecutedExpression() string {
	if o == nil || IsNil(o.CanBeExecutedExpression) {
		var ret string
		return ret
	}
	return *o.CanBeExecutedExpression
}

// GetCanBeExecutedExpressionOk returns a tuple with the CanBeExecutedExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepDefinition) GetCanBeExecutedExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.CanBeExecutedExpression) {
		return nil, false
	}
	return o.CanBeExecutedExpression, true
}

// HasCanBeExecutedExpression returns a boolean if a field has been set.
func (o *WorkflowStepDefinition) HasCanBeExecutedExpression() bool {
	if o != nil && !IsNil(o.CanBeExecutedExpression) {
		return true
	}

	return false
}

// SetCanBeExecutedExpression gets a reference to the given string and assigns it to the CanBeExecutedExpression field.
func (o *WorkflowStepDefinition) SetCanBeExecutedExpression(v string) {
	o.CanBeExecutedExpression = &v
}

// GetActions returns the Actions field value
func (o *WorkflowStepDefinition) GetActions() []WorkflowStepActionDefinition {
	if o == nil {
		var ret []WorkflowStepActionDefinition
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *WorkflowStepDefinition) GetActionsOk() ([]WorkflowStepActionDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *WorkflowStepDefinition) SetActions(v []WorkflowStepActionDefinition) {
	o.Actions = v
}

func (o WorkflowStepDefinition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowStepDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.DynamicParameters) {
		toSerialize["dynamic_parameters"] = o.DynamicParameters
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.GroupName) {
		toSerialize["group_name"] = o.GroupName
	}
	if o.PrerunAction.IsSet() {
		toSerialize["prerun_action"] = o.PrerunAction.Get()
	}
	if !IsNil(o.CanBeExecutedExpression) {
		toSerialize["can_be_executed_expression"] = o.CanBeExecutedExpression
	}
	toSerialize["actions"] = o.Actions
	return toSerialize, nil
}

func (o *WorkflowStepDefinition) UnmarshalJSON(data []byte) (err error) {
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

	varWorkflowStepDefinition := _WorkflowStepDefinition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkflowStepDefinition)

	if err != nil {
		return err
	}

	*o = WorkflowStepDefinition(varWorkflowStepDefinition)

	return err
}

type NullableWorkflowStepDefinition struct {
	value *WorkflowStepDefinition
	isSet bool
}

func (v NullableWorkflowStepDefinition) Get() *WorkflowStepDefinition {
	return v.value
}

func (v *NullableWorkflowStepDefinition) Set(val *WorkflowStepDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStepDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStepDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStepDefinition(val *WorkflowStepDefinition) *NullableWorkflowStepDefinition {
	return &NullableWorkflowStepDefinition{value: val, isSet: true}
}

func (v NullableWorkflowStepDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStepDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
