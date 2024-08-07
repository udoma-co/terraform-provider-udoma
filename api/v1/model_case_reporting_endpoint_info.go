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

// checks if the CaseReportingEndpointInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaseReportingEndpointInfo{}

// CaseReportingEndpointInfo Contains necessary information for displaying a publicly available form,  where tenants who don't have a login raise cases.
type CaseReportingEndpointInfo struct {
	Endpoint   *CaseReportingEndpoint `json:"endpoint,omitempty"`
	Properties []Property             `json:"properties,omitempty"`
	// The case templates that can be used to raise a case from this endpoint
	Templates []CaseTemplate `json:"templates,omitempty"`
	// Optional list of FAQs that should be displayed on the endpoint
	Faqs    []FAQEntry      `json:"faqs,omitempty"`
	Company *CompanyProfile `json:"company,omitempty"`
}

// NewCaseReportingEndpointInfo instantiates a new CaseReportingEndpointInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaseReportingEndpointInfo() *CaseReportingEndpointInfo {
	this := CaseReportingEndpointInfo{}
	return &this
}

// NewCaseReportingEndpointInfoWithDefaults instantiates a new CaseReportingEndpointInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaseReportingEndpointInfoWithDefaults() *CaseReportingEndpointInfo {
	this := CaseReportingEndpointInfo{}
	return &this
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *CaseReportingEndpointInfo) GetEndpoint() CaseReportingEndpoint {
	if o == nil || IsNil(o.Endpoint) {
		var ret CaseReportingEndpoint
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseReportingEndpointInfo) GetEndpointOk() (*CaseReportingEndpoint, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *CaseReportingEndpointInfo) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given CaseReportingEndpoint and assigns it to the Endpoint field.
func (o *CaseReportingEndpointInfo) SetEndpoint(v CaseReportingEndpoint) {
	o.Endpoint = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *CaseReportingEndpointInfo) GetProperties() []Property {
	if o == nil || IsNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseReportingEndpointInfo) GetPropertiesOk() ([]Property, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CaseReportingEndpointInfo) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *CaseReportingEndpointInfo) SetProperties(v []Property) {
	o.Properties = v
}

// GetTemplates returns the Templates field value if set, zero value otherwise.
func (o *CaseReportingEndpointInfo) GetTemplates() []CaseTemplate {
	if o == nil || IsNil(o.Templates) {
		var ret []CaseTemplate
		return ret
	}
	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseReportingEndpointInfo) GetTemplatesOk() ([]CaseTemplate, bool) {
	if o == nil || IsNil(o.Templates) {
		return nil, false
	}
	return o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *CaseReportingEndpointInfo) HasTemplates() bool {
	if o != nil && !IsNil(o.Templates) {
		return true
	}

	return false
}

// SetTemplates gets a reference to the given []CaseTemplate and assigns it to the Templates field.
func (o *CaseReportingEndpointInfo) SetTemplates(v []CaseTemplate) {
	o.Templates = v
}

// GetFaqs returns the Faqs field value if set, zero value otherwise.
func (o *CaseReportingEndpointInfo) GetFaqs() []FAQEntry {
	if o == nil || IsNil(o.Faqs) {
		var ret []FAQEntry
		return ret
	}
	return o.Faqs
}

// GetFaqsOk returns a tuple with the Faqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseReportingEndpointInfo) GetFaqsOk() ([]FAQEntry, bool) {
	if o == nil || IsNil(o.Faqs) {
		return nil, false
	}
	return o.Faqs, true
}

// HasFaqs returns a boolean if a field has been set.
func (o *CaseReportingEndpointInfo) HasFaqs() bool {
	if o != nil && !IsNil(o.Faqs) {
		return true
	}

	return false
}

// SetFaqs gets a reference to the given []FAQEntry and assigns it to the Faqs field.
func (o *CaseReportingEndpointInfo) SetFaqs(v []FAQEntry) {
	o.Faqs = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *CaseReportingEndpointInfo) GetCompany() CompanyProfile {
	if o == nil || IsNil(o.Company) {
		var ret CompanyProfile
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseReportingEndpointInfo) GetCompanyOk() (*CompanyProfile, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *CaseReportingEndpointInfo) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given CompanyProfile and assigns it to the Company field.
func (o *CaseReportingEndpointInfo) SetCompany(v CompanyProfile) {
	o.Company = &v
}

func (o CaseReportingEndpointInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaseReportingEndpointInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Templates) {
		toSerialize["templates"] = o.Templates
	}
	if !IsNil(o.Faqs) {
		toSerialize["faqs"] = o.Faqs
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	return toSerialize, nil
}

type NullableCaseReportingEndpointInfo struct {
	value *CaseReportingEndpointInfo
	isSet bool
}

func (v NullableCaseReportingEndpointInfo) Get() *CaseReportingEndpointInfo {
	return v.value
}

func (v *NullableCaseReportingEndpointInfo) Set(val *CaseReportingEndpointInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCaseReportingEndpointInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCaseReportingEndpointInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaseReportingEndpointInfo(val *CaseReportingEndpointInfo) *NullableCaseReportingEndpointInfo {
	return &NullableCaseReportingEndpointInfo{value: val, isSet: true}
}

func (v NullableCaseReportingEndpointInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaseReportingEndpointInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
