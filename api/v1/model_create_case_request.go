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

// checks if the CreateCaseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateCaseRequest{}

// CreateCaseRequest All the information required for raising a new case
type CreateCaseRequest struct {
	// Reference to property
	PropertyRef     *string  `json:"property_ref,omitempty"`
	PropertyAddress *Address `json:"property_address,omitempty"`
	// Input provided by the user when updating the case as a key-value map
	Data         map[string]interface{} `json:"data"`
	ReporterInfo NullableContactData    `json:"reporter_info"`
	// The ID of the case template used to create this case
	TemplateRef string `json:"template_ref"`
}

type _CreateCaseRequest CreateCaseRequest

// NewCreateCaseRequest instantiates a new CreateCaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCaseRequest(data map[string]interface{}, reporterInfo NullableContactData, templateRef string) *CreateCaseRequest {
	this := CreateCaseRequest{}
	this.Data = data
	this.ReporterInfo = reporterInfo
	this.TemplateRef = templateRef
	return &this
}

// NewCreateCaseRequestWithDefaults instantiates a new CreateCaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCaseRequestWithDefaults() *CreateCaseRequest {
	this := CreateCaseRequest{}
	return &this
}

// GetPropertyRef returns the PropertyRef field value if set, zero value otherwise.
func (o *CreateCaseRequest) GetPropertyRef() string {
	if o == nil || IsNil(o.PropertyRef) {
		var ret string
		return ret
	}
	return *o.PropertyRef
}

// GetPropertyRefOk returns a tuple with the PropertyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCaseRequest) GetPropertyRefOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyRef) {
		return nil, false
	}
	return o.PropertyRef, true
}

// HasPropertyRef returns a boolean if a field has been set.
func (o *CreateCaseRequest) HasPropertyRef() bool {
	if o != nil && !IsNil(o.PropertyRef) {
		return true
	}

	return false
}

// SetPropertyRef gets a reference to the given string and assigns it to the PropertyRef field.
func (o *CreateCaseRequest) SetPropertyRef(v string) {
	o.PropertyRef = &v
}

// GetPropertyAddress returns the PropertyAddress field value if set, zero value otherwise.
func (o *CreateCaseRequest) GetPropertyAddress() Address {
	if o == nil || IsNil(o.PropertyAddress) {
		var ret Address
		return ret
	}
	return *o.PropertyAddress
}

// GetPropertyAddressOk returns a tuple with the PropertyAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCaseRequest) GetPropertyAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.PropertyAddress) {
		return nil, false
	}
	return o.PropertyAddress, true
}

// HasPropertyAddress returns a boolean if a field has been set.
func (o *CreateCaseRequest) HasPropertyAddress() bool {
	if o != nil && !IsNil(o.PropertyAddress) {
		return true
	}

	return false
}

// SetPropertyAddress gets a reference to the given Address and assigns it to the PropertyAddress field.
func (o *CreateCaseRequest) SetPropertyAddress(v Address) {
	o.PropertyAddress = &v
}

// GetData returns the Data field value
func (o *CreateCaseRequest) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateCaseRequest) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CreateCaseRequest) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetReporterInfo returns the ReporterInfo field value
// If the value is explicit nil, the zero value for ContactData will be returned
func (o *CreateCaseRequest) GetReporterInfo() ContactData {
	if o == nil || o.ReporterInfo.Get() == nil {
		var ret ContactData
		return ret
	}

	return *o.ReporterInfo.Get()
}

// GetReporterInfoOk returns a tuple with the ReporterInfo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateCaseRequest) GetReporterInfoOk() (*ContactData, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReporterInfo.Get(), o.ReporterInfo.IsSet()
}

// SetReporterInfo sets field value
func (o *CreateCaseRequest) SetReporterInfo(v ContactData) {
	o.ReporterInfo.Set(&v)
}

// GetTemplateRef returns the TemplateRef field value
func (o *CreateCaseRequest) GetTemplateRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateRef
}

// GetTemplateRefOk returns a tuple with the TemplateRef field value
// and a boolean to check if the value has been set.
func (o *CreateCaseRequest) GetTemplateRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateRef, true
}

// SetTemplateRef sets field value
func (o *CreateCaseRequest) SetTemplateRef(v string) {
	o.TemplateRef = v
}

func (o CreateCaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCaseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PropertyRef) {
		toSerialize["property_ref"] = o.PropertyRef
	}
	if !IsNil(o.PropertyAddress) {
		toSerialize["property_address"] = o.PropertyAddress
	}
	toSerialize["data"] = o.Data
	toSerialize["reporter_info"] = o.ReporterInfo.Get()
	toSerialize["template_ref"] = o.TemplateRef
	return toSerialize, nil
}

func (o *CreateCaseRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"reporter_info",
		"template_ref",
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

	varCreateCaseRequest := _CreateCaseRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateCaseRequest)

	if err != nil {
		return err
	}

	*o = CreateCaseRequest(varCreateCaseRequest)

	return err
}

type NullableCreateCaseRequest struct {
	value *CreateCaseRequest
	isSet bool
}

func (v NullableCreateCaseRequest) Get() *CreateCaseRequest {
	return v.value
}

func (v *NullableCreateCaseRequest) Set(val *CreateCaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCaseRequest(val *CreateCaseRequest) *NullableCreateCaseRequest {
	return &NullableCreateCaseRequest{value: val, isSet: true}
}

func (v NullableCreateCaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
