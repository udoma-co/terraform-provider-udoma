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

// checks if the CreateOrUpdatePropertyHandoverRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdatePropertyHandoverRequest{}

// CreateOrUpdatePropertyHandoverRequest struct for CreateOrUpdatePropertyHandoverRequest
type CreateOrUpdatePropertyHandoverRequest struct {
	// The ID of the property handover template.
	TemplateRef     string   `json:"template_ref"`
	PropertyRef     *string  `json:"property_ref,omitempty"`
	PropertyAddress *Address `json:"property_address,omitempty"`
	// The time of the property handover.
	HandoverTime int64         `json:"handover_time"`
	Tenants      []ContactData `json:"tenants"`
}

type _CreateOrUpdatePropertyHandoverRequest CreateOrUpdatePropertyHandoverRequest

// NewCreateOrUpdatePropertyHandoverRequest instantiates a new CreateOrUpdatePropertyHandoverRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdatePropertyHandoverRequest(templateRef string, handoverTime int64, tenants []ContactData) *CreateOrUpdatePropertyHandoverRequest {
	this := CreateOrUpdatePropertyHandoverRequest{}
	this.TemplateRef = templateRef
	this.HandoverTime = handoverTime
	this.Tenants = tenants
	return &this
}

// NewCreateOrUpdatePropertyHandoverRequestWithDefaults instantiates a new CreateOrUpdatePropertyHandoverRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdatePropertyHandoverRequestWithDefaults() *CreateOrUpdatePropertyHandoverRequest {
	this := CreateOrUpdatePropertyHandoverRequest{}
	return &this
}

// GetTemplateRef returns the TemplateRef field value
func (o *CreateOrUpdatePropertyHandoverRequest) GetTemplateRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateRef
}

// GetTemplateRefOk returns a tuple with the TemplateRef field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdatePropertyHandoverRequest) GetTemplateRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateRef, true
}

// SetTemplateRef sets field value
func (o *CreateOrUpdatePropertyHandoverRequest) SetTemplateRef(v string) {
	o.TemplateRef = v
}

// GetPropertyRef returns the PropertyRef field value if set, zero value otherwise.
func (o *CreateOrUpdatePropertyHandoverRequest) GetPropertyRef() string {
	if o == nil || IsNil(o.PropertyRef) {
		var ret string
		return ret
	}
	return *o.PropertyRef
}

// GetPropertyRefOk returns a tuple with the PropertyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdatePropertyHandoverRequest) GetPropertyRefOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyRef) {
		return nil, false
	}
	return o.PropertyRef, true
}

// HasPropertyRef returns a boolean if a field has been set.
func (o *CreateOrUpdatePropertyHandoverRequest) HasPropertyRef() bool {
	if o != nil && !IsNil(o.PropertyRef) {
		return true
	}

	return false
}

// SetPropertyRef gets a reference to the given string and assigns it to the PropertyRef field.
func (o *CreateOrUpdatePropertyHandoverRequest) SetPropertyRef(v string) {
	o.PropertyRef = &v
}

// GetPropertyAddress returns the PropertyAddress field value if set, zero value otherwise.
func (o *CreateOrUpdatePropertyHandoverRequest) GetPropertyAddress() Address {
	if o == nil || IsNil(o.PropertyAddress) {
		var ret Address
		return ret
	}
	return *o.PropertyAddress
}

// GetPropertyAddressOk returns a tuple with the PropertyAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdatePropertyHandoverRequest) GetPropertyAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.PropertyAddress) {
		return nil, false
	}
	return o.PropertyAddress, true
}

// HasPropertyAddress returns a boolean if a field has been set.
func (o *CreateOrUpdatePropertyHandoverRequest) HasPropertyAddress() bool {
	if o != nil && !IsNil(o.PropertyAddress) {
		return true
	}

	return false
}

// SetPropertyAddress gets a reference to the given Address and assigns it to the PropertyAddress field.
func (o *CreateOrUpdatePropertyHandoverRequest) SetPropertyAddress(v Address) {
	o.PropertyAddress = &v
}

// GetHandoverTime returns the HandoverTime field value
func (o *CreateOrUpdatePropertyHandoverRequest) GetHandoverTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.HandoverTime
}

// GetHandoverTimeOk returns a tuple with the HandoverTime field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdatePropertyHandoverRequest) GetHandoverTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HandoverTime, true
}

// SetHandoverTime sets field value
func (o *CreateOrUpdatePropertyHandoverRequest) SetHandoverTime(v int64) {
	o.HandoverTime = v
}

// GetTenants returns the Tenants field value
func (o *CreateOrUpdatePropertyHandoverRequest) GetTenants() []ContactData {
	if o == nil {
		var ret []ContactData
		return ret
	}

	return o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdatePropertyHandoverRequest) GetTenantsOk() ([]ContactData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenants, true
}

// SetTenants sets field value
func (o *CreateOrUpdatePropertyHandoverRequest) SetTenants(v []ContactData) {
	o.Tenants = v
}

func (o CreateOrUpdatePropertyHandoverRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdatePropertyHandoverRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["template_ref"] = o.TemplateRef
	if !IsNil(o.PropertyRef) {
		toSerialize["property_ref"] = o.PropertyRef
	}
	if !IsNil(o.PropertyAddress) {
		toSerialize["property_address"] = o.PropertyAddress
	}
	toSerialize["handover_time"] = o.HandoverTime
	toSerialize["tenants"] = o.Tenants
	return toSerialize, nil
}

func (o *CreateOrUpdatePropertyHandoverRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"template_ref",
		"handover_time",
		"tenants",
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

	varCreateOrUpdatePropertyHandoverRequest := _CreateOrUpdatePropertyHandoverRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrUpdatePropertyHandoverRequest)

	if err != nil {
		return err
	}

	*o = CreateOrUpdatePropertyHandoverRequest(varCreateOrUpdatePropertyHandoverRequest)

	return err
}

type NullableCreateOrUpdatePropertyHandoverRequest struct {
	value *CreateOrUpdatePropertyHandoverRequest
	isSet bool
}

func (v NullableCreateOrUpdatePropertyHandoverRequest) Get() *CreateOrUpdatePropertyHandoverRequest {
	return v.value
}

func (v *NullableCreateOrUpdatePropertyHandoverRequest) Set(val *CreateOrUpdatePropertyHandoverRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdatePropertyHandoverRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdatePropertyHandoverRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdatePropertyHandoverRequest(val *CreateOrUpdatePropertyHandoverRequest) *NullableCreateOrUpdatePropertyHandoverRequest {
	return &NullableCreateOrUpdatePropertyHandoverRequest{value: val, isSet: true}
}

func (v NullableCreateOrUpdatePropertyHandoverRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdatePropertyHandoverRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
