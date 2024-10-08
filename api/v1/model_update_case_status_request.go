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

// checks if the UpdateCaseStatusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCaseStatusRequest{}

// UpdateCaseStatusRequest All the information required for updating the status of a case
type UpdateCaseStatusRequest struct {
	Action   *CaseActionEnum `json:"action,omitempty"`
	Assignee *CaseAssignee   `json:"assignee,omitempty"`
	// Feedback that is provided from the party executing the action. This information will be visible only to the property manager
	Feedback []CaseFeedbackResponse `json:"feedback,omitempty"`
}

// NewUpdateCaseStatusRequest instantiates a new UpdateCaseStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCaseStatusRequest() *UpdateCaseStatusRequest {
	this := UpdateCaseStatusRequest{}
	return &this
}

// NewUpdateCaseStatusRequestWithDefaults instantiates a new UpdateCaseStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCaseStatusRequestWithDefaults() *UpdateCaseStatusRequest {
	this := UpdateCaseStatusRequest{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *UpdateCaseStatusRequest) GetAction() CaseActionEnum {
	if o == nil || IsNil(o.Action) {
		var ret CaseActionEnum
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCaseStatusRequest) GetActionOk() (*CaseActionEnum, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *UpdateCaseStatusRequest) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given CaseActionEnum and assigns it to the Action field.
func (o *UpdateCaseStatusRequest) SetAction(v CaseActionEnum) {
	o.Action = &v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *UpdateCaseStatusRequest) GetAssignee() CaseAssignee {
	if o == nil || IsNil(o.Assignee) {
		var ret CaseAssignee
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCaseStatusRequest) GetAssigneeOk() (*CaseAssignee, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *UpdateCaseStatusRequest) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given CaseAssignee and assigns it to the Assignee field.
func (o *UpdateCaseStatusRequest) SetAssignee(v CaseAssignee) {
	o.Assignee = &v
}

// GetFeedback returns the Feedback field value if set, zero value otherwise.
func (o *UpdateCaseStatusRequest) GetFeedback() []CaseFeedbackResponse {
	if o == nil || IsNil(o.Feedback) {
		var ret []CaseFeedbackResponse
		return ret
	}
	return o.Feedback
}

// GetFeedbackOk returns a tuple with the Feedback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCaseStatusRequest) GetFeedbackOk() ([]CaseFeedbackResponse, bool) {
	if o == nil || IsNil(o.Feedback) {
		return nil, false
	}
	return o.Feedback, true
}

// HasFeedback returns a boolean if a field has been set.
func (o *UpdateCaseStatusRequest) HasFeedback() bool {
	if o != nil && !IsNil(o.Feedback) {
		return true
	}

	return false
}

// SetFeedback gets a reference to the given []CaseFeedbackResponse and assigns it to the Feedback field.
func (o *UpdateCaseStatusRequest) SetFeedback(v []CaseFeedbackResponse) {
	o.Feedback = v
}

func (o UpdateCaseStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCaseStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	if !IsNil(o.Feedback) {
		toSerialize["feedback"] = o.Feedback
	}
	return toSerialize, nil
}

type NullableUpdateCaseStatusRequest struct {
	value *UpdateCaseStatusRequest
	isSet bool
}

func (v NullableUpdateCaseStatusRequest) Get() *UpdateCaseStatusRequest {
	return v.value
}

func (v *NullableUpdateCaseStatusRequest) Set(val *UpdateCaseStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCaseStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCaseStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCaseStatusRequest(val *UpdateCaseStatusRequest) *NullableUpdateCaseStatusRequest {
	return &NullableUpdateCaseStatusRequest{value: val, isSet: true}
}

func (v NullableUpdateCaseStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCaseStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
