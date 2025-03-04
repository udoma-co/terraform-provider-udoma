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

// checks if the WorkflowDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowDefinition{}

// WorkflowDefinition struct for WorkflowDefinition
type WorkflowDefinition struct {
	// Unique and immutable ID attribute of the entity that is generated when the instance is created. The ID is unique within the system accross all accounts and it can be used to reference the entity in other entities or to retrieve it from the backend.
	Id string `json:"id"`
	// The date and time the entity was created
	CreatedAt int64 `json:"created_at"`
	// The date and time the entity was last updated
	UpdatedAt   int64   `json:"updated_at"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// a JS expression that determines the name of the workflow execution
	NameExpression *string `json:"name_expression,omitempty"`
	// The icon of the workflow (shown in the menu). If empty, the default icon  of the workflow type will be used.
	Icon *string `json:"icon,omitempty"`
	// a map of values, where the key and values are strings
	EnvVars *map[string]string `json:"env_vars,omitempty"`
	// ID of the first step of the workflow, which will be executed when the  workflow is started. If the workflow is started via a manual trigger, the init_step will be executed ahead of that.
	FirstStepId string                             `json:"first_step_id"`
	InitStep    NullableWorkflowInitStepDefinition `json:"init_step,omitempty"`
	Steps       []WorkflowStepDefinition           `json:"steps"`
	Version     *int32                             `json:"version,omitempty"`
}

type _WorkflowDefinition WorkflowDefinition

// NewWorkflowDefinition instantiates a new WorkflowDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowDefinition(id string, createdAt int64, updatedAt int64, name string, firstStepId string, steps []WorkflowStepDefinition) *WorkflowDefinition {
	this := WorkflowDefinition{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Name = name
	this.FirstStepId = firstStepId
	this.Steps = steps
	return &this
}

// NewWorkflowDefinitionWithDefaults instantiates a new WorkflowDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowDefinitionWithDefaults() *WorkflowDefinition {
	this := WorkflowDefinition{}
	return &this
}

// GetId returns the Id field value
func (o *WorkflowDefinition) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkflowDefinition) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkflowDefinition) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WorkflowDefinition) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkflowDefinition) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WorkflowDefinition) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *WorkflowDefinition) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkflowDefinition) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *WorkflowDefinition) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetName returns the Name field value
func (o *WorkflowDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkflowDefinition) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkflowDefinition) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowDefinition) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowDefinition) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetNameExpression returns the NameExpression field value if set, zero value otherwise.
func (o *WorkflowDefinition) GetNameExpression() string {
	if o == nil || IsNil(o.NameExpression) {
		var ret string
		return ret
	}
	return *o.NameExpression
}

// GetNameExpressionOk returns a tuple with the NameExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDefinition) GetNameExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.NameExpression) {
		return nil, false
	}
	return o.NameExpression, true
}

// HasNameExpression returns a boolean if a field has been set.
func (o *WorkflowDefinition) HasNameExpression() bool {
	if o != nil && !IsNil(o.NameExpression) {
		return true
	}

	return false
}

// SetNameExpression gets a reference to the given string and assigns it to the NameExpression field.
func (o *WorkflowDefinition) SetNameExpression(v string) {
	o.NameExpression = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *WorkflowDefinition) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDefinition) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *WorkflowDefinition) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *WorkflowDefinition) SetIcon(v string) {
	o.Icon = &v
}

// GetEnvVars returns the EnvVars field value if set, zero value otherwise.
func (o *WorkflowDefinition) GetEnvVars() map[string]string {
	if o == nil || IsNil(o.EnvVars) {
		var ret map[string]string
		return ret
	}
	return *o.EnvVars
}

// GetEnvVarsOk returns a tuple with the EnvVars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDefinition) GetEnvVarsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.EnvVars) {
		return nil, false
	}
	return o.EnvVars, true
}

// HasEnvVars returns a boolean if a field has been set.
func (o *WorkflowDefinition) HasEnvVars() bool {
	if o != nil && !IsNil(o.EnvVars) {
		return true
	}

	return false
}

// SetEnvVars gets a reference to the given map[string]string and assigns it to the EnvVars field.
func (o *WorkflowDefinition) SetEnvVars(v map[string]string) {
	o.EnvVars = &v
}

// GetFirstStepId returns the FirstStepId field value
func (o *WorkflowDefinition) GetFirstStepId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstStepId
}

// GetFirstStepIdOk returns a tuple with the FirstStepId field value
// and a boolean to check if the value has been set.
func (o *WorkflowDefinition) GetFirstStepIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstStepId, true
}

// SetFirstStepId sets field value
func (o *WorkflowDefinition) SetFirstStepId(v string) {
	o.FirstStepId = v
}

// GetInitStep returns the InitStep field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowDefinition) GetInitStep() WorkflowInitStepDefinition {
	if o == nil || IsNil(o.InitStep.Get()) {
		var ret WorkflowInitStepDefinition
		return ret
	}
	return *o.InitStep.Get()
}

// GetInitStepOk returns a tuple with the InitStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowDefinition) GetInitStepOk() (*WorkflowInitStepDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.InitStep.Get(), o.InitStep.IsSet()
}

// HasInitStep returns a boolean if a field has been set.
func (o *WorkflowDefinition) HasInitStep() bool {
	if o != nil && o.InitStep.IsSet() {
		return true
	}

	return false
}

// SetInitStep gets a reference to the given NullableWorkflowInitStepDefinition and assigns it to the InitStep field.
func (o *WorkflowDefinition) SetInitStep(v WorkflowInitStepDefinition) {
	o.InitStep.Set(&v)
}

// SetInitStepNil sets the value for InitStep to be an explicit nil
func (o *WorkflowDefinition) SetInitStepNil() {
	o.InitStep.Set(nil)
}

// UnsetInitStep ensures that no value is present for InitStep, not even an explicit nil
func (o *WorkflowDefinition) UnsetInitStep() {
	o.InitStep.Unset()
}

// GetSteps returns the Steps field value
func (o *WorkflowDefinition) GetSteps() []WorkflowStepDefinition {
	if o == nil {
		var ret []WorkflowStepDefinition
		return ret
	}

	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *WorkflowDefinition) GetStepsOk() ([]WorkflowStepDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Steps, true
}

// SetSteps sets field value
func (o *WorkflowDefinition) SetSteps(v []WorkflowStepDefinition) {
	o.Steps = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowDefinition) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDefinition) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowDefinition) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *WorkflowDefinition) SetVersion(v int32) {
	o.Version = &v
}

func (o WorkflowDefinition) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.NameExpression) {
		toSerialize["name_expression"] = o.NameExpression
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !IsNil(o.EnvVars) {
		toSerialize["env_vars"] = o.EnvVars
	}
	toSerialize["first_step_id"] = o.FirstStepId
	if o.InitStep.IsSet() {
		toSerialize["init_step"] = o.InitStep.Get()
	}
	toSerialize["steps"] = o.Steps
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

func (o *WorkflowDefinition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"name",
		"first_step_id",
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

	varWorkflowDefinition := _WorkflowDefinition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkflowDefinition)

	if err != nil {
		return err
	}

	*o = WorkflowDefinition(varWorkflowDefinition)

	return err
}

type NullableWorkflowDefinition struct {
	value *WorkflowDefinition
	isSet bool
}

func (v NullableWorkflowDefinition) Get() *WorkflowDefinition {
	return v.value
}

func (v *NullableWorkflowDefinition) Set(val *WorkflowDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowDefinition(val *WorkflowDefinition) *NullableWorkflowDefinition {
	return &NullableWorkflowDefinition{value: val, isSet: true}
}

func (v NullableWorkflowDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
