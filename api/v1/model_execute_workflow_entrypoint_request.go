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

// ExecuteWorkflowEntrypointRequest a request for starting a workflow execution via an entry point
type ExecuteWorkflowEntrypointRequest struct {
	// optional ID of the source object, for which this entrypoint was triggered
	SourceRef *string `json:"source_ref,omitempty"`
	// the data to be passed to the workflow execution
	Data map[string]interface{} `json:"data,omitempty"`
}

// NewExecuteWorkflowEntrypointRequest instantiates a new ExecuteWorkflowEntrypointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecuteWorkflowEntrypointRequest() *ExecuteWorkflowEntrypointRequest {
	this := ExecuteWorkflowEntrypointRequest{}
	return &this
}

// NewExecuteWorkflowEntrypointRequestWithDefaults instantiates a new ExecuteWorkflowEntrypointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecuteWorkflowEntrypointRequestWithDefaults() *ExecuteWorkflowEntrypointRequest {
	this := ExecuteWorkflowEntrypointRequest{}
	return &this
}

// GetSourceRef returns the SourceRef field value if set, zero value otherwise.
func (o *ExecuteWorkflowEntrypointRequest) GetSourceRef() string {
	if o == nil || o.SourceRef == nil {
		var ret string
		return ret
	}
	return *o.SourceRef
}

// GetSourceRefOk returns a tuple with the SourceRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteWorkflowEntrypointRequest) GetSourceRefOk() (*string, bool) {
	if o == nil || o.SourceRef == nil {
		return nil, false
	}
	return o.SourceRef, true
}

// HasSourceRef returns a boolean if a field has been set.
func (o *ExecuteWorkflowEntrypointRequest) HasSourceRef() bool {
	if o != nil && o.SourceRef != nil {
		return true
	}

	return false
}

// SetSourceRef gets a reference to the given string and assigns it to the SourceRef field.
func (o *ExecuteWorkflowEntrypointRequest) SetSourceRef(v string) {
	o.SourceRef = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ExecuteWorkflowEntrypointRequest) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteWorkflowEntrypointRequest) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ExecuteWorkflowEntrypointRequest) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *ExecuteWorkflowEntrypointRequest) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o ExecuteWorkflowEntrypointRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceRef != nil {
		toSerialize["source_ref"] = o.SourceRef
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableExecuteWorkflowEntrypointRequest struct {
	value *ExecuteWorkflowEntrypointRequest
	isSet bool
}

func (v NullableExecuteWorkflowEntrypointRequest) Get() *ExecuteWorkflowEntrypointRequest {
	return v.value
}

func (v *NullableExecuteWorkflowEntrypointRequest) Set(val *ExecuteWorkflowEntrypointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteWorkflowEntrypointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteWorkflowEntrypointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteWorkflowEntrypointRequest(val *ExecuteWorkflowEntrypointRequest) *NullableExecuteWorkflowEntrypointRequest {
	return &NullableExecuteWorkflowEntrypointRequest{value: val, isSet: true}
}

func (v NullableExecuteWorkflowEntrypointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteWorkflowEntrypointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
