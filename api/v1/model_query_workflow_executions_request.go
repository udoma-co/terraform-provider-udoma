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

// checks if the QueryWorkflowExecutionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryWorkflowExecutionsRequest{}

// QueryWorkflowExecutionsRequest a request for querying workflow executions
type QueryWorkflowExecutionsRequest struct {
	WorkflowRef                NullableString         `json:"workflow_ref,omitempty"`
	ParentWorkflowExecutionRef NullableString         `json:"parent_workflow_execution_ref,omitempty"`
	Active                     *QueryBooleanParameter `json:"active,omitempty"`
	Finished                   *QueryBooleanParameter `json:"finished,omitempty"`
	Archived                   *QueryBooleanParameter `json:"archived,omitempty"`
	Limit                      *int32                 `json:"limit,omitempty"`
}

// NewQueryWorkflowExecutionsRequest instantiates a new QueryWorkflowExecutionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryWorkflowExecutionsRequest() *QueryWorkflowExecutionsRequest {
	this := QueryWorkflowExecutionsRequest{}
	return &this
}

// NewQueryWorkflowExecutionsRequestWithDefaults instantiates a new QueryWorkflowExecutionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryWorkflowExecutionsRequestWithDefaults() *QueryWorkflowExecutionsRequest {
	this := QueryWorkflowExecutionsRequest{}
	return &this
}

// GetWorkflowRef returns the WorkflowRef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueryWorkflowExecutionsRequest) GetWorkflowRef() string {
	if o == nil || IsNil(o.WorkflowRef.Get()) {
		var ret string
		return ret
	}
	return *o.WorkflowRef.Get()
}

// GetWorkflowRefOk returns a tuple with the WorkflowRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueryWorkflowExecutionsRequest) GetWorkflowRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkflowRef.Get(), o.WorkflowRef.IsSet()
}

// HasWorkflowRef returns a boolean if a field has been set.
func (o *QueryWorkflowExecutionsRequest) HasWorkflowRef() bool {
	if o != nil && o.WorkflowRef.IsSet() {
		return true
	}

	return false
}

// SetWorkflowRef gets a reference to the given NullableString and assigns it to the WorkflowRef field.
func (o *QueryWorkflowExecutionsRequest) SetWorkflowRef(v string) {
	o.WorkflowRef.Set(&v)
}

// SetWorkflowRefNil sets the value for WorkflowRef to be an explicit nil
func (o *QueryWorkflowExecutionsRequest) SetWorkflowRefNil() {
	o.WorkflowRef.Set(nil)
}

// UnsetWorkflowRef ensures that no value is present for WorkflowRef, not even an explicit nil
func (o *QueryWorkflowExecutionsRequest) UnsetWorkflowRef() {
	o.WorkflowRef.Unset()
}

// GetParentWorkflowExecutionRef returns the ParentWorkflowExecutionRef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QueryWorkflowExecutionsRequest) GetParentWorkflowExecutionRef() string {
	if o == nil || IsNil(o.ParentWorkflowExecutionRef.Get()) {
		var ret string
		return ret
	}
	return *o.ParentWorkflowExecutionRef.Get()
}

// GetParentWorkflowExecutionRefOk returns a tuple with the ParentWorkflowExecutionRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QueryWorkflowExecutionsRequest) GetParentWorkflowExecutionRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentWorkflowExecutionRef.Get(), o.ParentWorkflowExecutionRef.IsSet()
}

// HasParentWorkflowExecutionRef returns a boolean if a field has been set.
func (o *QueryWorkflowExecutionsRequest) HasParentWorkflowExecutionRef() bool {
	if o != nil && o.ParentWorkflowExecutionRef.IsSet() {
		return true
	}

	return false
}

// SetParentWorkflowExecutionRef gets a reference to the given NullableString and assigns it to the ParentWorkflowExecutionRef field.
func (o *QueryWorkflowExecutionsRequest) SetParentWorkflowExecutionRef(v string) {
	o.ParentWorkflowExecutionRef.Set(&v)
}

// SetParentWorkflowExecutionRefNil sets the value for ParentWorkflowExecutionRef to be an explicit nil
func (o *QueryWorkflowExecutionsRequest) SetParentWorkflowExecutionRefNil() {
	o.ParentWorkflowExecutionRef.Set(nil)
}

// UnsetParentWorkflowExecutionRef ensures that no value is present for ParentWorkflowExecutionRef, not even an explicit nil
func (o *QueryWorkflowExecutionsRequest) UnsetParentWorkflowExecutionRef() {
	o.ParentWorkflowExecutionRef.Unset()
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *QueryWorkflowExecutionsRequest) GetActive() QueryBooleanParameter {
	if o == nil || IsNil(o.Active) {
		var ret QueryBooleanParameter
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryWorkflowExecutionsRequest) GetActiveOk() (*QueryBooleanParameter, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *QueryWorkflowExecutionsRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given QueryBooleanParameter and assigns it to the Active field.
func (o *QueryWorkflowExecutionsRequest) SetActive(v QueryBooleanParameter) {
	o.Active = &v
}

// GetFinished returns the Finished field value if set, zero value otherwise.
func (o *QueryWorkflowExecutionsRequest) GetFinished() QueryBooleanParameter {
	if o == nil || IsNil(o.Finished) {
		var ret QueryBooleanParameter
		return ret
	}
	return *o.Finished
}

// GetFinishedOk returns a tuple with the Finished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryWorkflowExecutionsRequest) GetFinishedOk() (*QueryBooleanParameter, bool) {
	if o == nil || IsNil(o.Finished) {
		return nil, false
	}
	return o.Finished, true
}

// HasFinished returns a boolean if a field has been set.
func (o *QueryWorkflowExecutionsRequest) HasFinished() bool {
	if o != nil && !IsNil(o.Finished) {
		return true
	}

	return false
}

// SetFinished gets a reference to the given QueryBooleanParameter and assigns it to the Finished field.
func (o *QueryWorkflowExecutionsRequest) SetFinished(v QueryBooleanParameter) {
	o.Finished = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *QueryWorkflowExecutionsRequest) GetArchived() QueryBooleanParameter {
	if o == nil || IsNil(o.Archived) {
		var ret QueryBooleanParameter
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryWorkflowExecutionsRequest) GetArchivedOk() (*QueryBooleanParameter, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *QueryWorkflowExecutionsRequest) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given QueryBooleanParameter and assigns it to the Archived field.
func (o *QueryWorkflowExecutionsRequest) SetArchived(v QueryBooleanParameter) {
	o.Archived = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *QueryWorkflowExecutionsRequest) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryWorkflowExecutionsRequest) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryWorkflowExecutionsRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *QueryWorkflowExecutionsRequest) SetLimit(v int32) {
	o.Limit = &v
}

func (o QueryWorkflowExecutionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryWorkflowExecutionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.WorkflowRef.IsSet() {
		toSerialize["workflow_ref"] = o.WorkflowRef.Get()
	}
	if o.ParentWorkflowExecutionRef.IsSet() {
		toSerialize["parent_workflow_execution_ref"] = o.ParentWorkflowExecutionRef.Get()
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Finished) {
		toSerialize["finished"] = o.Finished
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	return toSerialize, nil
}

type NullableQueryWorkflowExecutionsRequest struct {
	value *QueryWorkflowExecutionsRequest
	isSet bool
}

func (v NullableQueryWorkflowExecutionsRequest) Get() *QueryWorkflowExecutionsRequest {
	return v.value
}

func (v *NullableQueryWorkflowExecutionsRequest) Set(val *QueryWorkflowExecutionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryWorkflowExecutionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryWorkflowExecutionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryWorkflowExecutionsRequest(val *QueryWorkflowExecutionsRequest) *NullableQueryWorkflowExecutionsRequest {
	return &NullableQueryWorkflowExecutionsRequest{value: val, isSet: true}
}

func (v NullableQueryWorkflowExecutionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryWorkflowExecutionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
