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

// checks if the WorkflowExecution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowExecution{}

// WorkflowExecution contains all information of the execution of a single workflow
type WorkflowExecution struct {
	// Unique and immutable ID attribute of the entity that is generated when the instance is created. The ID is unique within the system accross all accounts and it can be used to reference the entity in other entities or to retrieve it from the backend.
	Id string `json:"id"`
	// The date and time the entity was created
	CreatedAt int64 `json:"created_at"`
	// The date and time the entity was last updated
	UpdatedAt int64 `json:"updated_at"`
	// the ID of the workflow definition on which the workflow execution is based
	DefinitionRef string `json:"definition_ref"`
	// the name of the workflow execution generated by the name expression of the workflow definition
	Name   string                  `json:"name"`
	Status WorkflowExecutionStatus `json:"status"`
	// whether the workflow execution is archived or not
	Archived *bool `json:"archived,omitempty"`
	// the ID of the step that is the current step of the workflow execution, i.e. it's  the next step to be executed
	CurrentStep string `json:"current_step"`
	// the name of the step that is the current step of the workflow execution. This is used in the UI to display the current step, without loading the step definition.
	CurrentStepName *string `json:"current_step_name,omitempty"`
	// the data collected and generated during the workflow execution, serialiased as JSON
	Context *string `json:"context,omitempty"`
	// the steps of the workflow execution, including the result of the execution of each step
	Steps       []WorkflowExecutionStep       `json:"steps"`
	StepHistory []WorkflowExecutionStepResult `json:"step_history,omitempty"`
}

type _WorkflowExecution WorkflowExecution

// NewWorkflowExecution instantiates a new WorkflowExecution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowExecution(id string, createdAt int64, updatedAt int64, definitionRef string, name string, status WorkflowExecutionStatus, currentStep string, steps []WorkflowExecutionStep) *WorkflowExecution {
	this := WorkflowExecution{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.DefinitionRef = definitionRef
	this.Name = name
	this.Status = status
	this.CurrentStep = currentStep
	this.Steps = steps
	return &this
}

// NewWorkflowExecutionWithDefaults instantiates a new WorkflowExecution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowExecutionWithDefaults() *WorkflowExecution {
	this := WorkflowExecution{}
	return &this
}

// GetId returns the Id field value
func (o *WorkflowExecution) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecution) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkflowExecution) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WorkflowExecution) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecution) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WorkflowExecution) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *WorkflowExecution) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecution) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *WorkflowExecution) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetDefinitionRef returns the DefinitionRef field value
func (o *WorkflowExecution) GetDefinitionRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefinitionRef
}

// GetDefinitionRefOk returns a tuple with the DefinitionRef field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecution) GetDefinitionRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefinitionRef, true
}

// SetDefinitionRef sets field value
func (o *WorkflowExecution) SetDefinitionRef(v string) {
	o.DefinitionRef = v
}

// GetName returns the Name field value
func (o *WorkflowExecution) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecution) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkflowExecution) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *WorkflowExecution) GetStatus() WorkflowExecutionStatus {
	if o == nil {
		var ret WorkflowExecutionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecution) GetStatusOk() (*WorkflowExecutionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *WorkflowExecution) SetStatus(v WorkflowExecutionStatus) {
	o.Status = v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *WorkflowExecution) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecution) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *WorkflowExecution) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *WorkflowExecution) SetArchived(v bool) {
	o.Archived = &v
}

// GetCurrentStep returns the CurrentStep field value
func (o *WorkflowExecution) GetCurrentStep() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentStep
}

// GetCurrentStepOk returns a tuple with the CurrentStep field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecution) GetCurrentStepOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentStep, true
}

// SetCurrentStep sets field value
func (o *WorkflowExecution) SetCurrentStep(v string) {
	o.CurrentStep = v
}

// GetCurrentStepName returns the CurrentStepName field value if set, zero value otherwise.
func (o *WorkflowExecution) GetCurrentStepName() string {
	if o == nil || IsNil(o.CurrentStepName) {
		var ret string
		return ret
	}
	return *o.CurrentStepName
}

// GetCurrentStepNameOk returns a tuple with the CurrentStepName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecution) GetCurrentStepNameOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentStepName) {
		return nil, false
	}
	return o.CurrentStepName, true
}

// HasCurrentStepName returns a boolean if a field has been set.
func (o *WorkflowExecution) HasCurrentStepName() bool {
	if o != nil && !IsNil(o.CurrentStepName) {
		return true
	}

	return false
}

// SetCurrentStepName gets a reference to the given string and assigns it to the CurrentStepName field.
func (o *WorkflowExecution) SetCurrentStepName(v string) {
	o.CurrentStepName = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *WorkflowExecution) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecution) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *WorkflowExecution) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *WorkflowExecution) SetContext(v string) {
	o.Context = &v
}

// GetSteps returns the Steps field value
func (o *WorkflowExecution) GetSteps() []WorkflowExecutionStep {
	if o == nil {
		var ret []WorkflowExecutionStep
		return ret
	}

	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *WorkflowExecution) GetStepsOk() ([]WorkflowExecutionStep, bool) {
	if o == nil {
		return nil, false
	}
	return o.Steps, true
}

// SetSteps sets field value
func (o *WorkflowExecution) SetSteps(v []WorkflowExecutionStep) {
	o.Steps = v
}

// GetStepHistory returns the StepHistory field value if set, zero value otherwise.
func (o *WorkflowExecution) GetStepHistory() []WorkflowExecutionStepResult {
	if o == nil || IsNil(o.StepHistory) {
		var ret []WorkflowExecutionStepResult
		return ret
	}
	return o.StepHistory
}

// GetStepHistoryOk returns a tuple with the StepHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowExecution) GetStepHistoryOk() ([]WorkflowExecutionStepResult, bool) {
	if o == nil || IsNil(o.StepHistory) {
		return nil, false
	}
	return o.StepHistory, true
}

// HasStepHistory returns a boolean if a field has been set.
func (o *WorkflowExecution) HasStepHistory() bool {
	if o != nil && !IsNil(o.StepHistory) {
		return true
	}

	return false
}

// SetStepHistory gets a reference to the given []WorkflowExecutionStepResult and assigns it to the StepHistory field.
func (o *WorkflowExecution) SetStepHistory(v []WorkflowExecutionStepResult) {
	o.StepHistory = v
}

func (o WorkflowExecution) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowExecution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["definition_ref"] = o.DefinitionRef
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	toSerialize["current_step"] = o.CurrentStep
	if !IsNil(o.CurrentStepName) {
		toSerialize["current_step_name"] = o.CurrentStepName
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	toSerialize["steps"] = o.Steps
	if !IsNil(o.StepHistory) {
		toSerialize["step_history"] = o.StepHistory
	}
	return toSerialize, nil
}

func (o *WorkflowExecution) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"definition_ref",
		"name",
		"status",
		"current_step",
		"steps",
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

	varWorkflowExecution := _WorkflowExecution{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkflowExecution)

	if err != nil {
		return err
	}

	*o = WorkflowExecution(varWorkflowExecution)

	return err
}

type NullableWorkflowExecution struct {
	value *WorkflowExecution
	isSet bool
}

func (v NullableWorkflowExecution) Get() *WorkflowExecution {
	return v.value
}

func (v *NullableWorkflowExecution) Set(val *WorkflowExecution) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowExecution) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowExecution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowExecution(val *WorkflowExecution) *NullableWorkflowExecution {
	return &NullableWorkflowExecution{value: val, isSet: true}
}

func (v NullableWorkflowExecution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowExecution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
