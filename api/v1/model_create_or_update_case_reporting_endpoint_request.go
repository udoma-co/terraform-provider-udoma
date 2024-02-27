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

// checks if the CreateOrUpdateCaseReportingEndpointRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateCaseReportingEndpointRequest{}

// CreateOrUpdateCaseReportingEndpointRequest struct for CreateOrUpdateCaseReportingEndpointRequest
type CreateOrUpdateCaseReportingEndpointRequest struct {
	// The name of the endpoint (only used in admin pages)
	Name *string `json:"name,omitempty"`
	// Whether the endpoint is enabled or not
	Active *bool `json:"active,omitempty"`
	// The IDs of the properties for which the endpoint can be used to raise cases
	PropertyRefs []string `json:"property_refs,omitempty"`
	// Categories, used to group the templates, in order not to clutter the Case Reporting Endpoint page.
	CaseCategories []CaseReportingEndpointCategory `json:"case_categories,omitempty"`
}

// NewCreateOrUpdateCaseReportingEndpointRequest instantiates a new CreateOrUpdateCaseReportingEndpointRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateCaseReportingEndpointRequest() *CreateOrUpdateCaseReportingEndpointRequest {
	this := CreateOrUpdateCaseReportingEndpointRequest{}
	return &this
}

// NewCreateOrUpdateCaseReportingEndpointRequestWithDefaults instantiates a new CreateOrUpdateCaseReportingEndpointRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateCaseReportingEndpointRequestWithDefaults() *CreateOrUpdateCaseReportingEndpointRequest {
	this := CreateOrUpdateCaseReportingEndpointRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateOrUpdateCaseReportingEndpointRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateCaseReportingEndpointRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateOrUpdateCaseReportingEndpointRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateOrUpdateCaseReportingEndpointRequest) SetName(v string) {
	o.Name = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CreateOrUpdateCaseReportingEndpointRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateCaseReportingEndpointRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CreateOrUpdateCaseReportingEndpointRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *CreateOrUpdateCaseReportingEndpointRequest) SetActive(v bool) {
	o.Active = &v
}

// GetPropertyRefs returns the PropertyRefs field value if set, zero value otherwise.
func (o *CreateOrUpdateCaseReportingEndpointRequest) GetPropertyRefs() []string {
	if o == nil || IsNil(o.PropertyRefs) {
		var ret []string
		return ret
	}
	return o.PropertyRefs
}

// GetPropertyRefsOk returns a tuple with the PropertyRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateCaseReportingEndpointRequest) GetPropertyRefsOk() ([]string, bool) {
	if o == nil || IsNil(o.PropertyRefs) {
		return nil, false
	}
	return o.PropertyRefs, true
}

// HasPropertyRefs returns a boolean if a field has been set.
func (o *CreateOrUpdateCaseReportingEndpointRequest) HasPropertyRefs() bool {
	if o != nil && !IsNil(o.PropertyRefs) {
		return true
	}

	return false
}

// SetPropertyRefs gets a reference to the given []string and assigns it to the PropertyRefs field.
func (o *CreateOrUpdateCaseReportingEndpointRequest) SetPropertyRefs(v []string) {
	o.PropertyRefs = v
}

// GetCaseCategories returns the CaseCategories field value if set, zero value otherwise.
func (o *CreateOrUpdateCaseReportingEndpointRequest) GetCaseCategories() []CaseReportingEndpointCategory {
	if o == nil || IsNil(o.CaseCategories) {
		var ret []CaseReportingEndpointCategory
		return ret
	}
	return o.CaseCategories
}

// GetCaseCategoriesOk returns a tuple with the CaseCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateCaseReportingEndpointRequest) GetCaseCategoriesOk() ([]CaseReportingEndpointCategory, bool) {
	if o == nil || IsNil(o.CaseCategories) {
		return nil, false
	}
	return o.CaseCategories, true
}

// HasCaseCategories returns a boolean if a field has been set.
func (o *CreateOrUpdateCaseReportingEndpointRequest) HasCaseCategories() bool {
	if o != nil && !IsNil(o.CaseCategories) {
		return true
	}

	return false
}

// SetCaseCategories gets a reference to the given []CaseReportingEndpointCategory and assigns it to the CaseCategories field.
func (o *CreateOrUpdateCaseReportingEndpointRequest) SetCaseCategories(v []CaseReportingEndpointCategory) {
	o.CaseCategories = v
}

func (o CreateOrUpdateCaseReportingEndpointRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateCaseReportingEndpointRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.PropertyRefs) {
		toSerialize["property_refs"] = o.PropertyRefs
	}
	if !IsNil(o.CaseCategories) {
		toSerialize["case_categories"] = o.CaseCategories
	}
	return toSerialize, nil
}

type NullableCreateOrUpdateCaseReportingEndpointRequest struct {
	value *CreateOrUpdateCaseReportingEndpointRequest
	isSet bool
}

func (v NullableCreateOrUpdateCaseReportingEndpointRequest) Get() *CreateOrUpdateCaseReportingEndpointRequest {
	return v.value
}

func (v *NullableCreateOrUpdateCaseReportingEndpointRequest) Set(val *CreateOrUpdateCaseReportingEndpointRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateCaseReportingEndpointRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateCaseReportingEndpointRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateCaseReportingEndpointRequest(val *CreateOrUpdateCaseReportingEndpointRequest) *NullableCreateOrUpdateCaseReportingEndpointRequest {
	return &NullableCreateOrUpdateCaseReportingEndpointRequest{value: val, isSet: true}
}

func (v NullableCreateOrUpdateCaseReportingEndpointRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateCaseReportingEndpointRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
