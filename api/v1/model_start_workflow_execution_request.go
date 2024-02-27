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

// checks if the StartWorkflowExecutionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartWorkflowExecutionRequest{}

// StartWorkflowExecutionRequest a request for starting a workflow execution
type StartWorkflowExecutionRequest struct {
	WorkflowRef                *string `json:"workflow_ref,omitempty"`
	ParentWorkflowExecutionRef *string `json:"parent_workflow_execution_ref,omitempty"`
	// the initial context of the workflow execution as JSON
	Context *string `json:"context,omitempty"`
}

// NewStartWorkflowExecutionRequest instantiates a new StartWorkflowExecutionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartWorkflowExecutionRequest() *StartWorkflowExecutionRequest {
	this := StartWorkflowExecutionRequest{}
	return &this
}

// NewStartWorkflowExecutionRequestWithDefaults instantiates a new StartWorkflowExecutionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartWorkflowExecutionRequestWithDefaults() *StartWorkflowExecutionRequest {
	this := StartWorkflowExecutionRequest{}
	return &this
}

// GetWorkflowRef returns the WorkflowRef field value if set, zero value otherwise.
func (o *StartWorkflowExecutionRequest) GetWorkflowRef() string {
	if o == nil || IsNil(o.WorkflowRef) {
		var ret string
		return ret
	}
	return *o.WorkflowRef
}

// GetWorkflowRefOk returns a tuple with the WorkflowRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartWorkflowExecutionRequest) GetWorkflowRefOk() (*string, bool) {
	if o == nil || IsNil(o.WorkflowRef) {
		return nil, false
	}
	return o.WorkflowRef, true
}

// HasWorkflowRef returns a boolean if a field has been set.
func (o *StartWorkflowExecutionRequest) HasWorkflowRef() bool {
	if o != nil && !IsNil(o.WorkflowRef) {
		return true
	}

	return false
}

// SetWorkflowRef gets a reference to the given string and assigns it to the WorkflowRef field.
func (o *StartWorkflowExecutionRequest) SetWorkflowRef(v string) {
	o.WorkflowRef = &v
}

// GetParentWorkflowExecutionRef returns the ParentWorkflowExecutionRef field value if set, zero value otherwise.
func (o *StartWorkflowExecutionRequest) GetParentWorkflowExecutionRef() string {
	if o == nil || IsNil(o.ParentWorkflowExecutionRef) {
		var ret string
		return ret
	}
	return *o.ParentWorkflowExecutionRef
}

// GetParentWorkflowExecutionRefOk returns a tuple with the ParentWorkflowExecutionRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartWorkflowExecutionRequest) GetParentWorkflowExecutionRefOk() (*string, bool) {
	if o == nil || IsNil(o.ParentWorkflowExecutionRef) {
		return nil, false
	}
	return o.ParentWorkflowExecutionRef, true
}

// HasParentWorkflowExecutionRef returns a boolean if a field has been set.
func (o *StartWorkflowExecutionRequest) HasParentWorkflowExecutionRef() bool {
	if o != nil && !IsNil(o.ParentWorkflowExecutionRef) {
		return true
	}

	return false
}

// SetParentWorkflowExecutionRef gets a reference to the given string and assigns it to the ParentWorkflowExecutionRef field.
func (o *StartWorkflowExecutionRequest) SetParentWorkflowExecutionRef(v string) {
	o.ParentWorkflowExecutionRef = &v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *StartWorkflowExecutionRequest) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartWorkflowExecutionRequest) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *StartWorkflowExecutionRequest) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *StartWorkflowExecutionRequest) SetContext(v string) {
	o.Context = &v
}

func (o StartWorkflowExecutionRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartWorkflowExecutionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WorkflowRef) {
		toSerialize["workflow_ref"] = o.WorkflowRef
	}
	if !IsNil(o.ParentWorkflowExecutionRef) {
		toSerialize["parent_workflow_execution_ref"] = o.ParentWorkflowExecutionRef
	}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	return toSerialize, nil
}

type NullableStartWorkflowExecutionRequest struct {
	value *StartWorkflowExecutionRequest
	isSet bool
}

func (v NullableStartWorkflowExecutionRequest) Get() *StartWorkflowExecutionRequest {
	return v.value
}

func (v *NullableStartWorkflowExecutionRequest) Set(val *StartWorkflowExecutionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStartWorkflowExecutionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStartWorkflowExecutionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartWorkflowExecutionRequest(val *StartWorkflowExecutionRequest) *NullableStartWorkflowExecutionRequest {
	return &NullableStartWorkflowExecutionRequest{value: val, isSet: true}
}

func (v NullableStartWorkflowExecutionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartWorkflowExecutionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
