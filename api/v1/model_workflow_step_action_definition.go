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

// checks if the WorkflowStepActionDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowStepActionDefinition{}

// WorkflowStepActionDefinition an action that can be performed on a workflow step
type WorkflowStepActionDefinition struct {
	// the ID of the action
	Id *string `json:"id,omitempty"`
	// the icon of the action
	Icon *string `json:"icon,omitempty"`
	// the button label for the action
	Label *string `json:"label,omitempty"`
	// optional button modifier of the action
	ButtonModifier *string `json:"button_modifier,omitempty"`
	// optional filter that can be used by the UI to show or hide the action
	UiFilter *string `json:"ui_filter,omitempty"`
	// a parameter of a workflow step or step action. The value of the parameter is contextual and can vary in type and meaning depending on the step or action that uses it. If used in a step, the parameter will be available in the UI and will not be interpreted, i.e. JS expressions are not allowed. In actions however, the parameter might be interpreted as a JS expression, if the action type requires it.
	Parameters map[string]interface{} `json:"parameters,omitempty"`
	// a map of values, where the key and values are strings
	DynamicParameters *map[string]string `json:"dynamic_parameters,omitempty"`
	// An optional JS expression that determines whether the action can be executed or  not. If not set, this will default to true. If the expression returns false, the action will not show up in the UI.
	CanBeExecutedExpression *string `json:"can_be_executed_expression,omitempty"`
	// the ID of the next step of the workflow
	NextStepId *string `json:"next_step_id,omitempty"`
	// Optional JS expression that will be executed before any step- and action-specific logic is executed.
	ExecBefore *string `json:"exec_before,omitempty"`
	// Optional JS expression that will be executed after the step- and action-specific logic is executed.
	ExecAfter *string `json:"exec_after,omitempty"`
	// optional ID of a workflow definition that should be executed
	StartChildWorkflow *string `json:"start_child_workflow,omitempty"`
	// An optional JS expression that returns an array of entities. For each entity, a new workflow execution will be started. The entity will be available as a parameter `iterator` in the child execution context expression.
	ChildExecutionIteratorExpression *string `json:"child_execution_iterator_expression,omitempty"`
	// An optional JS expression that determines the context of the child workflow execution. If not set, the context of the parent workflow execution will be used.
	ChildExecutionContextExpression *string `json:"child_execution_context_expression,omitempty"`
}

// NewWorkflowStepActionDefinition instantiates a new WorkflowStepActionDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStepActionDefinition() *WorkflowStepActionDefinition {
	this := WorkflowStepActionDefinition{}
	return &this
}

// NewWorkflowStepActionDefinitionWithDefaults instantiates a new WorkflowStepActionDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStepActionDefinitionWithDefaults() *WorkflowStepActionDefinition {
	this := WorkflowStepActionDefinition{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkflowStepActionDefinition) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepActionDefinition) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkflowStepActionDefinition) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkflowStepActionDefinition) SetId(v string) {
	o.Id = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *WorkflowStepActionDefinition) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepActionDefinition) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *WorkflowStepActionDefinition) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *WorkflowStepActionDefinition) SetIcon(v string) {
	o.Icon = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowStepActionDefinition) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepActionDefinition) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowStepActionDefinition) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowStepActionDefinition) SetLabel(v string) {
	o.Label = &v
}

// GetButtonModifier returns the ButtonModifier field value if set, zero value otherwise.
func (o *WorkflowStepActionDefinition) GetButtonModifier() string {
	if o == nil || IsNil(o.ButtonModifier) {
		var ret string
		return ret
	}
	return *o.ButtonModifier
}

// GetButtonModifierOk returns a tuple with the ButtonModifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepActionDefinition) GetButtonModifierOk() (*string, bool) {
	if o == nil || IsNil(o.ButtonModifier) {
		return nil, false
	}
	return o.ButtonModifier, true
}

// HasButtonModifier returns a boolean if a field has been set.
func (o *WorkflowStepActionDefinition) HasButtonModifier() bool {
	if o != nil && !IsNil(o.ButtonModifier) {
		return true
	}

	return false
}

// SetButtonModifier gets a reference to the given string and assigns it to the ButtonModifier field.
func (o *WorkflowStepActionDefinition) SetButtonModifier(v string) {
	o.ButtonModifier = &v
}

// GetUiFilter returns the UiFilter field value if set, zero value otherwise.
func (o *WorkflowStepActionDefinition) GetUiFilter() string {
	if o == nil || IsNil(o.UiFilter) {
		var ret string
		return ret
	}
	return *o.UiFilter
}

// GetUiFilterOk returns a tuple with the UiFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepActionDefinition) GetUiFilterOk() (*string, bool) {
	if o == nil || IsNil(o.UiFilter) {
		return nil, false
	}
	return o.UiFilter, true
}

// HasUiFilter returns a boolean if a field has been set.
func (o *WorkflowStepActionDefinition) HasUiFilter() bool {
	if o != nil && !IsNil(o.UiFilter) {
		return true
	}

	return false
}

// SetUiFilter gets a reference to the given string and assigns it to the UiFilter field.
func (o *WorkflowStepActionDefinition) SetUiFilter(v string) {
	o.UiFilter = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *WorkflowStepActionDefinition) GetParameters() map[string]interface{} {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]interface{}
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepActionDefinition) GetParametersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Parameters) {
		return map[string]interface{}{}, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *WorkflowStepActionDefinition) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *WorkflowStepActionDefinition) SetParameters(v map[string]interface{}) {
	o.Parameters = v
}

// GetDynamicParameters returns the DynamicParameters field value if set, zero value otherwise.
func (o *WorkflowStepActionDefinition) GetDynamicParameters() map[string]string {
	if o == nil || IsNil(o.DynamicParameters) {
		var ret map[string]string
		return ret
	}
	return *o.DynamicParameters
}

// GetDynamicParametersOk returns a tuple with the DynamicParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepActionDefinition) GetDynamicParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.DynamicParameters) {
		return nil, false
	}
	return o.DynamicParameters, true
}

// HasDynamicParameters returns a boolean if a field has been set.
func (o *WorkflowStepActionDefinition) HasDynamicParameters() bool {
	if o != nil && !IsNil(o.DynamicParameters) {
		return true
	}

	return false
}

// SetDynamicParameters gets a reference to the given map[string]string and assigns it to the DynamicParameters field.
func (o *WorkflowStepActionDefinition) SetDynamicParameters(v map[string]string) {
	o.DynamicParameters = &v
}

// GetCanBeExecutedExpression returns the CanBeExecutedExpression field value if set, zero value otherwise.
func (o *WorkflowStepActionDefinition) GetCanBeExecutedExpression() string {
	if o == nil || IsNil(o.CanBeExecutedExpression) {
		var ret string
		return ret
	}
	return *o.CanBeExecutedExpression
}

// GetCanBeExecutedExpressionOk returns a tuple with the CanBeExecutedExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepActionDefinition) GetCanBeExecutedExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.CanBeExecutedExpression) {
		return nil, false
	}
	return o.CanBeExecutedExpression, true
}

// HasCanBeExecutedExpression returns a boolean if a field has been set.
func (o *WorkflowStepActionDefinition) HasCanBeExecutedExpression() bool {
	if o != nil && !IsNil(o.CanBeExecutedExpression) {
		return true
	}

	return false
}

// SetCanBeExecutedExpression gets a reference to the given string and assigns it to the CanBeExecutedExpression field.
func (o *WorkflowStepActionDefinition) SetCanBeExecutedExpression(v string) {
	o.CanBeExecutedExpression = &v
}

// GetNextStepId returns the NextStepId field value if set, zero value otherwise.
func (o *WorkflowStepActionDefinition) GetNextStepId() string {
	if o == nil || IsNil(o.NextStepId) {
		var ret string
		return ret
	}
	return *o.NextStepId
}

// GetNextStepIdOk returns a tuple with the NextStepId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepActionDefinition) GetNextStepIdOk() (*string, bool) {
	if o == nil || IsNil(o.NextStepId) {
		return nil, false
	}
	return o.NextStepId, true
}

// HasNextStepId returns a boolean if a field has been set.
func (o *WorkflowStepActionDefinition) HasNextStepId() bool {
	if o != nil && !IsNil(o.NextStepId) {
		return true
	}

	return false
}

// SetNextStepId gets a reference to the given string and assigns it to the NextStepId field.
func (o *WorkflowStepActionDefinition) SetNextStepId(v string) {
	o.NextStepId = &v
}

// GetExecBefore returns the ExecBefore field value if set, zero value otherwise.
func (o *WorkflowStepActionDefinition) GetExecBefore() string {
	if o == nil || IsNil(o.ExecBefore) {
		var ret string
		return ret
	}
	return *o.ExecBefore
}

// GetExecBeforeOk returns a tuple with the ExecBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepActionDefinition) GetExecBeforeOk() (*string, bool) {
	if o == nil || IsNil(o.ExecBefore) {
		return nil, false
	}
	return o.ExecBefore, true
}

// HasExecBefore returns a boolean if a field has been set.
func (o *WorkflowStepActionDefinition) HasExecBefore() bool {
	if o != nil && !IsNil(o.ExecBefore) {
		return true
	}

	return false
}

// SetExecBefore gets a reference to the given string and assigns it to the ExecBefore field.
func (o *WorkflowStepActionDefinition) SetExecBefore(v string) {
	o.ExecBefore = &v
}

// GetExecAfter returns the ExecAfter field value if set, zero value otherwise.
func (o *WorkflowStepActionDefinition) GetExecAfter() string {
	if o == nil || IsNil(o.ExecAfter) {
		var ret string
		return ret
	}
	return *o.ExecAfter
}

// GetExecAfterOk returns a tuple with the ExecAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepActionDefinition) GetExecAfterOk() (*string, bool) {
	if o == nil || IsNil(o.ExecAfter) {
		return nil, false
	}
	return o.ExecAfter, true
}

// HasExecAfter returns a boolean if a field has been set.
func (o *WorkflowStepActionDefinition) HasExecAfter() bool {
	if o != nil && !IsNil(o.ExecAfter) {
		return true
	}

	return false
}

// SetExecAfter gets a reference to the given string and assigns it to the ExecAfter field.
func (o *WorkflowStepActionDefinition) SetExecAfter(v string) {
	o.ExecAfter = &v
}

// GetStartChildWorkflow returns the StartChildWorkflow field value if set, zero value otherwise.
func (o *WorkflowStepActionDefinition) GetStartChildWorkflow() string {
	if o == nil || IsNil(o.StartChildWorkflow) {
		var ret string
		return ret
	}
	return *o.StartChildWorkflow
}

// GetStartChildWorkflowOk returns a tuple with the StartChildWorkflow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepActionDefinition) GetStartChildWorkflowOk() (*string, bool) {
	if o == nil || IsNil(o.StartChildWorkflow) {
		return nil, false
	}
	return o.StartChildWorkflow, true
}

// HasStartChildWorkflow returns a boolean if a field has been set.
func (o *WorkflowStepActionDefinition) HasStartChildWorkflow() bool {
	if o != nil && !IsNil(o.StartChildWorkflow) {
		return true
	}

	return false
}

// SetStartChildWorkflow gets a reference to the given string and assigns it to the StartChildWorkflow field.
func (o *WorkflowStepActionDefinition) SetStartChildWorkflow(v string) {
	o.StartChildWorkflow = &v
}

// GetChildExecutionIteratorExpression returns the ChildExecutionIteratorExpression field value if set, zero value otherwise.
func (o *WorkflowStepActionDefinition) GetChildExecutionIteratorExpression() string {
	if o == nil || IsNil(o.ChildExecutionIteratorExpression) {
		var ret string
		return ret
	}
	return *o.ChildExecutionIteratorExpression
}

// GetChildExecutionIteratorExpressionOk returns a tuple with the ChildExecutionIteratorExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepActionDefinition) GetChildExecutionIteratorExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.ChildExecutionIteratorExpression) {
		return nil, false
	}
	return o.ChildExecutionIteratorExpression, true
}

// HasChildExecutionIteratorExpression returns a boolean if a field has been set.
func (o *WorkflowStepActionDefinition) HasChildExecutionIteratorExpression() bool {
	if o != nil && !IsNil(o.ChildExecutionIteratorExpression) {
		return true
	}

	return false
}

// SetChildExecutionIteratorExpression gets a reference to the given string and assigns it to the ChildExecutionIteratorExpression field.
func (o *WorkflowStepActionDefinition) SetChildExecutionIteratorExpression(v string) {
	o.ChildExecutionIteratorExpression = &v
}

// GetChildExecutionContextExpression returns the ChildExecutionContextExpression field value if set, zero value otherwise.
func (o *WorkflowStepActionDefinition) GetChildExecutionContextExpression() string {
	if o == nil || IsNil(o.ChildExecutionContextExpression) {
		var ret string
		return ret
	}
	return *o.ChildExecutionContextExpression
}

// GetChildExecutionContextExpressionOk returns a tuple with the ChildExecutionContextExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepActionDefinition) GetChildExecutionContextExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.ChildExecutionContextExpression) {
		return nil, false
	}
	return o.ChildExecutionContextExpression, true
}

// HasChildExecutionContextExpression returns a boolean if a field has been set.
func (o *WorkflowStepActionDefinition) HasChildExecutionContextExpression() bool {
	if o != nil && !IsNil(o.ChildExecutionContextExpression) {
		return true
	}

	return false
}

// SetChildExecutionContextExpression gets a reference to the given string and assigns it to the ChildExecutionContextExpression field.
func (o *WorkflowStepActionDefinition) SetChildExecutionContextExpression(v string) {
	o.ChildExecutionContextExpression = &v
}

func (o WorkflowStepActionDefinition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowStepActionDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.ButtonModifier) {
		toSerialize["button_modifier"] = o.ButtonModifier
	}
	if !IsNil(o.UiFilter) {
		toSerialize["ui_filter"] = o.UiFilter
	}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.DynamicParameters) {
		toSerialize["dynamic_parameters"] = o.DynamicParameters
	}
	if !IsNil(o.CanBeExecutedExpression) {
		toSerialize["can_be_executed_expression"] = o.CanBeExecutedExpression
	}
	if !IsNil(o.NextStepId) {
		toSerialize["next_step_id"] = o.NextStepId
	}
	if !IsNil(o.ExecBefore) {
		toSerialize["exec_before"] = o.ExecBefore
	}
	if !IsNil(o.ExecAfter) {
		toSerialize["exec_after"] = o.ExecAfter
	}
	if !IsNil(o.StartChildWorkflow) {
		toSerialize["start_child_workflow"] = o.StartChildWorkflow
	}
	if !IsNil(o.ChildExecutionIteratorExpression) {
		toSerialize["child_execution_iterator_expression"] = o.ChildExecutionIteratorExpression
	}
	if !IsNil(o.ChildExecutionContextExpression) {
		toSerialize["child_execution_context_expression"] = o.ChildExecutionContextExpression
	}
	return toSerialize, nil
}

type NullableWorkflowStepActionDefinition struct {
	value *WorkflowStepActionDefinition
	isSet bool
}

func (v NullableWorkflowStepActionDefinition) Get() *WorkflowStepActionDefinition {
	return v.value
}

func (v *NullableWorkflowStepActionDefinition) Set(val *WorkflowStepActionDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStepActionDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStepActionDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStepActionDefinition(val *WorkflowStepActionDefinition) *NullableWorkflowStepActionDefinition {
	return &NullableWorkflowStepActionDefinition{value: val, isSet: true}
}

func (v NullableWorkflowStepActionDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStepActionDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
