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

// CaseReportingEndpointCategory struct for CaseReportingEndpointCategory
type CaseReportingEndpointCategory struct {
	// a map of values, where the key and values are strings
	Name *map[string]string `json:"name,omitempty"`
	// The priority of the Category, used to decide the order in which to display the categories in the initial Case Reporting Endpoint page.
	Priority *int32 `json:"priority,omitempty"`
	// A list of Template objects, that contain the unique IDs of the template and priorities.
	Templates []CaseReportingEndpointCategoryTemplatesInner `json:"templates,omitempty"`
}

// NewCaseReportingEndpointCategory instantiates a new CaseReportingEndpointCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaseReportingEndpointCategory() *CaseReportingEndpointCategory {
	this := CaseReportingEndpointCategory{}
	return &this
}

// NewCaseReportingEndpointCategoryWithDefaults instantiates a new CaseReportingEndpointCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaseReportingEndpointCategoryWithDefaults() *CaseReportingEndpointCategory {
	this := CaseReportingEndpointCategory{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CaseReportingEndpointCategory) GetName() map[string]string {
	if o == nil || o.Name == nil {
		var ret map[string]string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseReportingEndpointCategory) GetNameOk() (*map[string]string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CaseReportingEndpointCategory) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given map[string]string and assigns it to the Name field.
func (o *CaseReportingEndpointCategory) SetName(v map[string]string) {
	o.Name = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *CaseReportingEndpointCategory) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseReportingEndpointCategory) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *CaseReportingEndpointCategory) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *CaseReportingEndpointCategory) SetPriority(v int32) {
	o.Priority = &v
}

// GetTemplates returns the Templates field value if set, zero value otherwise.
func (o *CaseReportingEndpointCategory) GetTemplates() []CaseReportingEndpointCategoryTemplatesInner {
	if o == nil || o.Templates == nil {
		var ret []CaseReportingEndpointCategoryTemplatesInner
		return ret
	}
	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseReportingEndpointCategory) GetTemplatesOk() ([]CaseReportingEndpointCategoryTemplatesInner, bool) {
	if o == nil || o.Templates == nil {
		return nil, false
	}
	return o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *CaseReportingEndpointCategory) HasTemplates() bool {
	if o != nil && o.Templates != nil {
		return true
	}

	return false
}

// SetTemplates gets a reference to the given []CaseReportingEndpointCategoryTemplatesInner and assigns it to the Templates field.
func (o *CaseReportingEndpointCategory) SetTemplates(v []CaseReportingEndpointCategoryTemplatesInner) {
	o.Templates = v
}

func (o CaseReportingEndpointCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Templates != nil {
		toSerialize["templates"] = o.Templates
	}
	return json.Marshal(toSerialize)
}

type NullableCaseReportingEndpointCategory struct {
	value *CaseReportingEndpointCategory
	isSet bool
}

func (v NullableCaseReportingEndpointCategory) Get() *CaseReportingEndpointCategory {
	return v.value
}

func (v *NullableCaseReportingEndpointCategory) Set(val *CaseReportingEndpointCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableCaseReportingEndpointCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableCaseReportingEndpointCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaseReportingEndpointCategory(val *CaseReportingEndpointCategory) *NullableCaseReportingEndpointCategory {
	return &NullableCaseReportingEndpointCategory{value: val, isSet: true}
}

func (v NullableCaseReportingEndpointCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaseReportingEndpointCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}