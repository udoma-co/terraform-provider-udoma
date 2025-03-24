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

// checks if the CaseDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaseDetails{}

// CaseDetails All details of a case, raised via a case report endpoint
type CaseDetails struct {
	Case     Case           `json:"case"`
	Reporter *UserReference `json:"reporter,omitempty"`
	// The status of the email that was used to report the case.
	ReporterEmailStatus *string      `json:"reporter_email_status,omitempty"`
	CaseParty           *CaseParty   `json:"case_party,omitempty"`
	Template            CaseTemplate `json:"template"`
	Property            *Property    `json:"property,omitempty"`
	// Indicator of if the user still has access to the case
	AccessRevoked *bool `json:"access_revoked,omitempty"`
}

type _CaseDetails CaseDetails

// NewCaseDetails instantiates a new CaseDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaseDetails(case_ Case, template CaseTemplate) *CaseDetails {
	this := CaseDetails{}
	this.Case = case_
	this.Template = template
	return &this
}

// NewCaseDetailsWithDefaults instantiates a new CaseDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaseDetailsWithDefaults() *CaseDetails {
	this := CaseDetails{}
	return &this
}

// GetCase returns the Case field value
func (o *CaseDetails) GetCase() Case {
	if o == nil {
		var ret Case
		return ret
	}

	return o.Case
}

// GetCaseOk returns a tuple with the Case field value
// and a boolean to check if the value has been set.
func (o *CaseDetails) GetCaseOk() (*Case, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Case, true
}

// SetCase sets field value
func (o *CaseDetails) SetCase(v Case) {
	o.Case = v
}

// GetReporter returns the Reporter field value if set, zero value otherwise.
func (o *CaseDetails) GetReporter() UserReference {
	if o == nil || IsNil(o.Reporter) {
		var ret UserReference
		return ret
	}
	return *o.Reporter
}

// GetReporterOk returns a tuple with the Reporter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseDetails) GetReporterOk() (*UserReference, bool) {
	if o == nil || IsNil(o.Reporter) {
		return nil, false
	}
	return o.Reporter, true
}

// HasReporter returns a boolean if a field has been set.
func (o *CaseDetails) HasReporter() bool {
	if o != nil && !IsNil(o.Reporter) {
		return true
	}

	return false
}

// SetReporter gets a reference to the given UserReference and assigns it to the Reporter field.
func (o *CaseDetails) SetReporter(v UserReference) {
	o.Reporter = &v
}

// GetReporterEmailStatus returns the ReporterEmailStatus field value if set, zero value otherwise.
func (o *CaseDetails) GetReporterEmailStatus() string {
	if o == nil || IsNil(o.ReporterEmailStatus) {
		var ret string
		return ret
	}
	return *o.ReporterEmailStatus
}

// GetReporterEmailStatusOk returns a tuple with the ReporterEmailStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseDetails) GetReporterEmailStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ReporterEmailStatus) {
		return nil, false
	}
	return o.ReporterEmailStatus, true
}

// HasReporterEmailStatus returns a boolean if a field has been set.
func (o *CaseDetails) HasReporterEmailStatus() bool {
	if o != nil && !IsNil(o.ReporterEmailStatus) {
		return true
	}

	return false
}

// SetReporterEmailStatus gets a reference to the given string and assigns it to the ReporterEmailStatus field.
func (o *CaseDetails) SetReporterEmailStatus(v string) {
	o.ReporterEmailStatus = &v
}

// GetCaseParty returns the CaseParty field value if set, zero value otherwise.
func (o *CaseDetails) GetCaseParty() CaseParty {
	if o == nil || IsNil(o.CaseParty) {
		var ret CaseParty
		return ret
	}
	return *o.CaseParty
}

// GetCasePartyOk returns a tuple with the CaseParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseDetails) GetCasePartyOk() (*CaseParty, bool) {
	if o == nil || IsNil(o.CaseParty) {
		return nil, false
	}
	return o.CaseParty, true
}

// HasCaseParty returns a boolean if a field has been set.
func (o *CaseDetails) HasCaseParty() bool {
	if o != nil && !IsNil(o.CaseParty) {
		return true
	}

	return false
}

// SetCaseParty gets a reference to the given CaseParty and assigns it to the CaseParty field.
func (o *CaseDetails) SetCaseParty(v CaseParty) {
	o.CaseParty = &v
}

// GetTemplate returns the Template field value
func (o *CaseDetails) GetTemplate() CaseTemplate {
	if o == nil {
		var ret CaseTemplate
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *CaseDetails) GetTemplateOk() (*CaseTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *CaseDetails) SetTemplate(v CaseTemplate) {
	o.Template = v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *CaseDetails) GetProperty() Property {
	if o == nil || IsNil(o.Property) {
		var ret Property
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseDetails) GetPropertyOk() (*Property, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *CaseDetails) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given Property and assigns it to the Property field.
func (o *CaseDetails) SetProperty(v Property) {
	o.Property = &v
}

// GetAccessRevoked returns the AccessRevoked field value if set, zero value otherwise.
func (o *CaseDetails) GetAccessRevoked() bool {
	if o == nil || IsNil(o.AccessRevoked) {
		var ret bool
		return ret
	}
	return *o.AccessRevoked
}

// GetAccessRevokedOk returns a tuple with the AccessRevoked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseDetails) GetAccessRevokedOk() (*bool, bool) {
	if o == nil || IsNil(o.AccessRevoked) {
		return nil, false
	}
	return o.AccessRevoked, true
}

// HasAccessRevoked returns a boolean if a field has been set.
func (o *CaseDetails) HasAccessRevoked() bool {
	if o != nil && !IsNil(o.AccessRevoked) {
		return true
	}

	return false
}

// SetAccessRevoked gets a reference to the given bool and assigns it to the AccessRevoked field.
func (o *CaseDetails) SetAccessRevoked(v bool) {
	o.AccessRevoked = &v
}

func (o CaseDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaseDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["case"] = o.Case
	if !IsNil(o.Reporter) {
		toSerialize["reporter"] = o.Reporter
	}
	if !IsNil(o.ReporterEmailStatus) {
		toSerialize["reporter_email_status"] = o.ReporterEmailStatus
	}
	if !IsNil(o.CaseParty) {
		toSerialize["case_party"] = o.CaseParty
	}
	toSerialize["template"] = o.Template
	if !IsNil(o.Property) {
		toSerialize["property"] = o.Property
	}
	if !IsNil(o.AccessRevoked) {
		toSerialize["access_revoked"] = o.AccessRevoked
	}
	return toSerialize, nil
}

func (o *CaseDetails) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"case",
		"template",
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

	varCaseDetails := _CaseDetails{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCaseDetails)

	if err != nil {
		return err
	}

	*o = CaseDetails(varCaseDetails)

	return err
}

type NullableCaseDetails struct {
	value *CaseDetails
	isSet bool
}

func (v NullableCaseDetails) Get() *CaseDetails {
	return v.value
}

func (v *NullableCaseDetails) Set(val *CaseDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCaseDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCaseDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaseDetails(val *CaseDetails) *NullableCaseDetails {
	return &NullableCaseDetails{value: val, isSet: true}
}

func (v NullableCaseDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaseDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
