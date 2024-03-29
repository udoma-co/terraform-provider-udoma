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

// checks if the WorkflowTrigger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowTrigger{}

// WorkflowTrigger information for the trigger of a workflow execution
type WorkflowTrigger struct {
	// the ID of the workflow execution
	ExecutionRef   *string                     `json:"execution_ref,omitempty"`
	EntrypointType *WorkflowEntrypointLocation `json:"entrypoint_type,omitempty"`
	// the ID of the source object
	SourceRef *string `json:"source_ref,omitempty"`
	// The timestamp of when the workflow was created
	CreatedAt *int64 `json:"created_at,omitempty"`
	// the name of the related workflow execution
	Label *string `json:"label,omitempty"`
}

// NewWorkflowTrigger instantiates a new WorkflowTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowTrigger() *WorkflowTrigger {
	this := WorkflowTrigger{}
	return &this
}

// NewWorkflowTriggerWithDefaults instantiates a new WorkflowTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowTriggerWithDefaults() *WorkflowTrigger {
	this := WorkflowTrigger{}
	return &this
}

// GetExecutionRef returns the ExecutionRef field value if set, zero value otherwise.
func (o *WorkflowTrigger) GetExecutionRef() string {
	if o == nil || IsNil(o.ExecutionRef) {
		var ret string
		return ret
	}
	return *o.ExecutionRef
}

// GetExecutionRefOk returns a tuple with the ExecutionRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTrigger) GetExecutionRefOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionRef) {
		return nil, false
	}
	return o.ExecutionRef, true
}

// HasExecutionRef returns a boolean if a field has been set.
func (o *WorkflowTrigger) HasExecutionRef() bool {
	if o != nil && !IsNil(o.ExecutionRef) {
		return true
	}

	return false
}

// SetExecutionRef gets a reference to the given string and assigns it to the ExecutionRef field.
func (o *WorkflowTrigger) SetExecutionRef(v string) {
	o.ExecutionRef = &v
}

// GetEntrypointType returns the EntrypointType field value if set, zero value otherwise.
func (o *WorkflowTrigger) GetEntrypointType() WorkflowEntrypointLocation {
	if o == nil || IsNil(o.EntrypointType) {
		var ret WorkflowEntrypointLocation
		return ret
	}
	return *o.EntrypointType
}

// GetEntrypointTypeOk returns a tuple with the EntrypointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTrigger) GetEntrypointTypeOk() (*WorkflowEntrypointLocation, bool) {
	if o == nil || IsNil(o.EntrypointType) {
		return nil, false
	}
	return o.EntrypointType, true
}

// HasEntrypointType returns a boolean if a field has been set.
func (o *WorkflowTrigger) HasEntrypointType() bool {
	if o != nil && !IsNil(o.EntrypointType) {
		return true
	}

	return false
}

// SetEntrypointType gets a reference to the given WorkflowEntrypointLocation and assigns it to the EntrypointType field.
func (o *WorkflowTrigger) SetEntrypointType(v WorkflowEntrypointLocation) {
	o.EntrypointType = &v
}

// GetSourceRef returns the SourceRef field value if set, zero value otherwise.
func (o *WorkflowTrigger) GetSourceRef() string {
	if o == nil || IsNil(o.SourceRef) {
		var ret string
		return ret
	}
	return *o.SourceRef
}

// GetSourceRefOk returns a tuple with the SourceRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTrigger) GetSourceRefOk() (*string, bool) {
	if o == nil || IsNil(o.SourceRef) {
		return nil, false
	}
	return o.SourceRef, true
}

// HasSourceRef returns a boolean if a field has been set.
func (o *WorkflowTrigger) HasSourceRef() bool {
	if o != nil && !IsNil(o.SourceRef) {
		return true
	}

	return false
}

// SetSourceRef gets a reference to the given string and assigns it to the SourceRef field.
func (o *WorkflowTrigger) SetSourceRef(v string) {
	o.SourceRef = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WorkflowTrigger) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTrigger) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WorkflowTrigger) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *WorkflowTrigger) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowTrigger) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowTrigger) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowTrigger) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowTrigger) SetLabel(v string) {
	o.Label = &v
}

func (o WorkflowTrigger) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowTrigger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExecutionRef) {
		toSerialize["execution_ref"] = o.ExecutionRef
	}
	if !IsNil(o.EntrypointType) {
		toSerialize["entrypoint_type"] = o.EntrypointType
	}
	if !IsNil(o.SourceRef) {
		toSerialize["source_ref"] = o.SourceRef
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullableWorkflowTrigger struct {
	value *WorkflowTrigger
	isSet bool
}

func (v NullableWorkflowTrigger) Get() *WorkflowTrigger {
	return v.value
}

func (v *NullableWorkflowTrigger) Set(val *WorkflowTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowTrigger(val *WorkflowTrigger) *NullableWorkflowTrigger {
	return &NullableWorkflowTrigger{value: val, isSet: true}
}

func (v NullableWorkflowTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
