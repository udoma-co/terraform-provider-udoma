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

// checks if the CaseReportingEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaseReportingEndpoint{}

// CaseReportingEndpoint Contains necessary information for displaying a publicly available form,  where tenants who don't have a login can raise cases.
type CaseReportingEndpoint struct {
	// Unique, immutable, alpha-numeric random code that is generated by the backend and can be used to access the entity via a publicly available URL. The code is unique within the system accross all accounts and it can be also be used to reference the entity in other entities or to retrieve it from the backend.
	Code string `json:"code"`
	// The date and time the entity was created
	CreatedAt int64 `json:"created_at"`
	// The date and time the entity was last updated
	UpdatedAt int64 `json:"updated_at"`
	// The name of the endpoint (only used in admin pages)
	Name string `json:"name"`
	// Whether the endpoint is enabled or not
	Active *bool `json:"active,omitempty"`
	// The IDs of the properties for which the endpoint can be used to raise cases
	PropertyRefs []string `json:"property_refs,omitempty"`
	// Categories, used to group the templates, in order not to clutter the Case Reporting Endpoint page.
	CaseCategories []CaseReportingEndpointCategory `json:"case_categories,omitempty"`
	// Optional list of FAQs that should be displayed on the endpoint
	Faqs []string `json:"faqs,omitempty"`
	// The full URL that can be used to access the endpoint
	Url string `json:"url"`
}

type _CaseReportingEndpoint CaseReportingEndpoint

// NewCaseReportingEndpoint instantiates a new CaseReportingEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaseReportingEndpoint(code string, createdAt int64, updatedAt int64, name string, url string) *CaseReportingEndpoint {
	this := CaseReportingEndpoint{}
	this.Code = code
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Name = name
	this.Url = url
	return &this
}

// NewCaseReportingEndpointWithDefaults instantiates a new CaseReportingEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaseReportingEndpointWithDefaults() *CaseReportingEndpoint {
	this := CaseReportingEndpoint{}
	return &this
}

// GetCode returns the Code field value
func (o *CaseReportingEndpoint) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *CaseReportingEndpoint) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *CaseReportingEndpoint) SetCode(v string) {
	o.Code = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CaseReportingEndpoint) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CaseReportingEndpoint) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CaseReportingEndpoint) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CaseReportingEndpoint) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CaseReportingEndpoint) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CaseReportingEndpoint) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetName returns the Name field value
func (o *CaseReportingEndpoint) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CaseReportingEndpoint) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CaseReportingEndpoint) SetName(v string) {
	o.Name = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CaseReportingEndpoint) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseReportingEndpoint) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CaseReportingEndpoint) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *CaseReportingEndpoint) SetActive(v bool) {
	o.Active = &v
}

// GetPropertyRefs returns the PropertyRefs field value if set, zero value otherwise.
func (o *CaseReportingEndpoint) GetPropertyRefs() []string {
	if o == nil || IsNil(o.PropertyRefs) {
		var ret []string
		return ret
	}
	return o.PropertyRefs
}

// GetPropertyRefsOk returns a tuple with the PropertyRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseReportingEndpoint) GetPropertyRefsOk() ([]string, bool) {
	if o == nil || IsNil(o.PropertyRefs) {
		return nil, false
	}
	return o.PropertyRefs, true
}

// HasPropertyRefs returns a boolean if a field has been set.
func (o *CaseReportingEndpoint) HasPropertyRefs() bool {
	if o != nil && !IsNil(o.PropertyRefs) {
		return true
	}

	return false
}

// SetPropertyRefs gets a reference to the given []string and assigns it to the PropertyRefs field.
func (o *CaseReportingEndpoint) SetPropertyRefs(v []string) {
	o.PropertyRefs = v
}

// GetCaseCategories returns the CaseCategories field value if set, zero value otherwise.
func (o *CaseReportingEndpoint) GetCaseCategories() []CaseReportingEndpointCategory {
	if o == nil || IsNil(o.CaseCategories) {
		var ret []CaseReportingEndpointCategory
		return ret
	}
	return o.CaseCategories
}

// GetCaseCategoriesOk returns a tuple with the CaseCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseReportingEndpoint) GetCaseCategoriesOk() ([]CaseReportingEndpointCategory, bool) {
	if o == nil || IsNil(o.CaseCategories) {
		return nil, false
	}
	return o.CaseCategories, true
}

// HasCaseCategories returns a boolean if a field has been set.
func (o *CaseReportingEndpoint) HasCaseCategories() bool {
	if o != nil && !IsNil(o.CaseCategories) {
		return true
	}

	return false
}

// SetCaseCategories gets a reference to the given []CaseReportingEndpointCategory and assigns it to the CaseCategories field.
func (o *CaseReportingEndpoint) SetCaseCategories(v []CaseReportingEndpointCategory) {
	o.CaseCategories = v
}

// GetFaqs returns the Faqs field value if set, zero value otherwise.
func (o *CaseReportingEndpoint) GetFaqs() []string {
	if o == nil || IsNil(o.Faqs) {
		var ret []string
		return ret
	}
	return o.Faqs
}

// GetFaqsOk returns a tuple with the Faqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseReportingEndpoint) GetFaqsOk() ([]string, bool) {
	if o == nil || IsNil(o.Faqs) {
		return nil, false
	}
	return o.Faqs, true
}

// HasFaqs returns a boolean if a field has been set.
func (o *CaseReportingEndpoint) HasFaqs() bool {
	if o != nil && !IsNil(o.Faqs) {
		return true
	}

	return false
}

// SetFaqs gets a reference to the given []string and assigns it to the Faqs field.
func (o *CaseReportingEndpoint) SetFaqs(v []string) {
	o.Faqs = v
}

// GetUrl returns the Url field value
func (o *CaseReportingEndpoint) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *CaseReportingEndpoint) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *CaseReportingEndpoint) SetUrl(v string) {
	o.Url = v
}

func (o CaseReportingEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaseReportingEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["name"] = o.Name
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.PropertyRefs) {
		toSerialize["property_refs"] = o.PropertyRefs
	}
	if !IsNil(o.CaseCategories) {
		toSerialize["case_categories"] = o.CaseCategories
	}
	if !IsNil(o.Faqs) {
		toSerialize["faqs"] = o.Faqs
	}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

func (o *CaseReportingEndpoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"created_at",
		"updated_at",
		"name",
		"url",
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

	varCaseReportingEndpoint := _CaseReportingEndpoint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCaseReportingEndpoint)

	if err != nil {
		return err
	}

	*o = CaseReportingEndpoint(varCaseReportingEndpoint)

	return err
}

type NullableCaseReportingEndpoint struct {
	value *CaseReportingEndpoint
	isSet bool
}

func (v NullableCaseReportingEndpoint) Get() *CaseReportingEndpoint {
	return v.value
}

func (v *NullableCaseReportingEndpoint) Set(val *CaseReportingEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableCaseReportingEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableCaseReportingEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaseReportingEndpoint(val *CaseReportingEndpoint) *NullableCaseReportingEndpoint {
	return &NullableCaseReportingEndpoint{value: val, isSet: true}
}

func (v NullableCaseReportingEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaseReportingEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
