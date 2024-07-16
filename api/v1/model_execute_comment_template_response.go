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

// checks if the ExecuteCommentTemplateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExecuteCommentTemplateResponse{}

// ExecuteCommentTemplateResponse struct for ExecuteCommentTemplateResponse
type ExecuteCommentTemplateResponse struct {
	// The processed comment.
	Comment *string `json:"comment,omitempty"`
}

// NewExecuteCommentTemplateResponse instantiates a new ExecuteCommentTemplateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExecuteCommentTemplateResponse() *ExecuteCommentTemplateResponse {
	this := ExecuteCommentTemplateResponse{}
	return &this
}

// NewExecuteCommentTemplateResponseWithDefaults instantiates a new ExecuteCommentTemplateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExecuteCommentTemplateResponseWithDefaults() *ExecuteCommentTemplateResponse {
	this := ExecuteCommentTemplateResponse{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ExecuteCommentTemplateResponse) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExecuteCommentTemplateResponse) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ExecuteCommentTemplateResponse) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ExecuteCommentTemplateResponse) SetComment(v string) {
	o.Comment = &v
}

func (o ExecuteCommentTemplateResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExecuteCommentTemplateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableExecuteCommentTemplateResponse struct {
	value *ExecuteCommentTemplateResponse
	isSet bool
}

func (v NullableExecuteCommentTemplateResponse) Get() *ExecuteCommentTemplateResponse {
	return v.value
}

func (v *NullableExecuteCommentTemplateResponse) Set(val *ExecuteCommentTemplateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExecuteCommentTemplateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExecuteCommentTemplateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExecuteCommentTemplateResponse(val *ExecuteCommentTemplateResponse) *NullableExecuteCommentTemplateResponse {
	return &NullableExecuteCommentTemplateResponse{value: val, isSet: true}
}

func (v NullableExecuteCommentTemplateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExecuteCommentTemplateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}