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

// checks if the Property type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Property{}

// Property Property holds the information about a single accommodation unit (Wohneinheit).
type Property struct {
	// Unique and immutable ID attribute of the entity that is generated when the instance is created. The ID is unique within the system accross all accounts and it can be used to reference the entity in other entities or to retrieve it from the backend.
	Id string `json:"id"`
	// The date and time the entity was created
	CreatedAt int64 `json:"created_at"`
	// The date and time the entity was last updated
	UpdatedAt int64 `json:"updated_at"`
	// Optional external ID, in case entity was created via backend integration
	ExternalId *string `json:"external_id,omitempty"`
	// Optional external source, in case entity was created via backend integration
	ExternalSource *string `json:"external_source,omitempty"`
	// meaningful name of the property, e.g. 'Whg 12', or 'Etage 2/Links'
	Name    string       `json:"name"`
	Type    PropertyType `json:"type"`
	Address *Address     `json:"address,omitempty"`
	// Optional reference to the property owner
	OwnerRef *string `json:"owner_ref,omitempty"`
	// Optional reference to the parent property
	ParentRef *string `json:"parent_ref,omitempty"`
	// Address of the suite.
	Suite   *string          `json:"suite,omitempty"`
	Details *PropertyDetails `json:"details,omitempty"`
	// Optional name of the parent property
	ParentName    *string  `json:"parent_name,omitempty"`
	ParentAddress *Address `json:"parent_address,omitempty"`
}

type _Property Property

// NewProperty instantiates a new Property object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProperty(id string, createdAt int64, updatedAt int64, name string, type_ PropertyType) *Property {
	this := Property{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Name = name
	this.Type = type_
	return &this
}

// NewPropertyWithDefaults instantiates a new Property object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyWithDefaults() *Property {
	this := Property{}
	return &this
}

// GetId returns the Id field value
func (o *Property) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Property) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Property) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Property) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Property) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Property) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Property) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Property) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Property) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *Property) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *Property) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *Property) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetExternalSource returns the ExternalSource field value if set, zero value otherwise.
func (o *Property) GetExternalSource() string {
	if o == nil || IsNil(o.ExternalSource) {
		var ret string
		return ret
	}
	return *o.ExternalSource
}

// GetExternalSourceOk returns a tuple with the ExternalSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetExternalSourceOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalSource) {
		return nil, false
	}
	return o.ExternalSource, true
}

// HasExternalSource returns a boolean if a field has been set.
func (o *Property) HasExternalSource() bool {
	if o != nil && !IsNil(o.ExternalSource) {
		return true
	}

	return false
}

// SetExternalSource gets a reference to the given string and assigns it to the ExternalSource field.
func (o *Property) SetExternalSource(v string) {
	o.ExternalSource = &v
}

// GetName returns the Name field value
func (o *Property) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Property) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Property) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *Property) GetType() PropertyType {
	if o == nil {
		var ret PropertyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Property) GetTypeOk() (*PropertyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Property) SetType(v PropertyType) {
	o.Type = v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Property) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Property) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *Property) SetAddress(v Address) {
	o.Address = &v
}

// GetOwnerRef returns the OwnerRef field value if set, zero value otherwise.
func (o *Property) GetOwnerRef() string {
	if o == nil || IsNil(o.OwnerRef) {
		var ret string
		return ret
	}
	return *o.OwnerRef
}

// GetOwnerRefOk returns a tuple with the OwnerRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetOwnerRefOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerRef) {
		return nil, false
	}
	return o.OwnerRef, true
}

// HasOwnerRef returns a boolean if a field has been set.
func (o *Property) HasOwnerRef() bool {
	if o != nil && !IsNil(o.OwnerRef) {
		return true
	}

	return false
}

// SetOwnerRef gets a reference to the given string and assigns it to the OwnerRef field.
func (o *Property) SetOwnerRef(v string) {
	o.OwnerRef = &v
}

// GetParentRef returns the ParentRef field value if set, zero value otherwise.
func (o *Property) GetParentRef() string {
	if o == nil || IsNil(o.ParentRef) {
		var ret string
		return ret
	}
	return *o.ParentRef
}

// GetParentRefOk returns a tuple with the ParentRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetParentRefOk() (*string, bool) {
	if o == nil || IsNil(o.ParentRef) {
		return nil, false
	}
	return o.ParentRef, true
}

// HasParentRef returns a boolean if a field has been set.
func (o *Property) HasParentRef() bool {
	if o != nil && !IsNil(o.ParentRef) {
		return true
	}

	return false
}

// SetParentRef gets a reference to the given string and assigns it to the ParentRef field.
func (o *Property) SetParentRef(v string) {
	o.ParentRef = &v
}

// GetSuite returns the Suite field value if set, zero value otherwise.
func (o *Property) GetSuite() string {
	if o == nil || IsNil(o.Suite) {
		var ret string
		return ret
	}
	return *o.Suite
}

// GetSuiteOk returns a tuple with the Suite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetSuiteOk() (*string, bool) {
	if o == nil || IsNil(o.Suite) {
		return nil, false
	}
	return o.Suite, true
}

// HasSuite returns a boolean if a field has been set.
func (o *Property) HasSuite() bool {
	if o != nil && !IsNil(o.Suite) {
		return true
	}

	return false
}

// SetSuite gets a reference to the given string and assigns it to the Suite field.
func (o *Property) SetSuite(v string) {
	o.Suite = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Property) GetDetails() PropertyDetails {
	if o == nil || IsNil(o.Details) {
		var ret PropertyDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetDetailsOk() (*PropertyDetails, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Property) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given PropertyDetails and assigns it to the Details field.
func (o *Property) SetDetails(v PropertyDetails) {
	o.Details = &v
}

// GetParentName returns the ParentName field value if set, zero value otherwise.
func (o *Property) GetParentName() string {
	if o == nil || IsNil(o.ParentName) {
		var ret string
		return ret
	}
	return *o.ParentName
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetParentNameOk() (*string, bool) {
	if o == nil || IsNil(o.ParentName) {
		return nil, false
	}
	return o.ParentName, true
}

// HasParentName returns a boolean if a field has been set.
func (o *Property) HasParentName() bool {
	if o != nil && !IsNil(o.ParentName) {
		return true
	}

	return false
}

// SetParentName gets a reference to the given string and assigns it to the ParentName field.
func (o *Property) SetParentName(v string) {
	o.ParentName = &v
}

// GetParentAddress returns the ParentAddress field value if set, zero value otherwise.
func (o *Property) GetParentAddress() Address {
	if o == nil || IsNil(o.ParentAddress) {
		var ret Address
		return ret
	}
	return *o.ParentAddress
}

// GetParentAddressOk returns a tuple with the ParentAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Property) GetParentAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.ParentAddress) {
		return nil, false
	}
	return o.ParentAddress, true
}

// HasParentAddress returns a boolean if a field has been set.
func (o *Property) HasParentAddress() bool {
	if o != nil && !IsNil(o.ParentAddress) {
		return true
	}

	return false
}

// SetParentAddress gets a reference to the given Address and assigns it to the ParentAddress field.
func (o *Property) SetParentAddress(v Address) {
	o.ParentAddress = &v
}

func (o Property) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Property) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if !IsNil(o.ExternalId) {
		toSerialize["external_id"] = o.ExternalId
	}
	if !IsNil(o.ExternalSource) {
		toSerialize["external_source"] = o.ExternalSource
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.OwnerRef) {
		toSerialize["owner_ref"] = o.OwnerRef
	}
	if !IsNil(o.ParentRef) {
		toSerialize["parent_ref"] = o.ParentRef
	}
	if !IsNil(o.Suite) {
		toSerialize["suite"] = o.Suite
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.ParentName) {
		toSerialize["parent_name"] = o.ParentName
	}
	if !IsNil(o.ParentAddress) {
		toSerialize["parent_address"] = o.ParentAddress
	}
	return toSerialize, nil
}

func (o *Property) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"name",
		"type",
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

	varProperty := _Property{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProperty)

	if err != nil {
		return err
	}

	*o = Property(varProperty)

	return err
}

type NullableProperty struct {
	value *Property
	isSet bool
}

func (v NullableProperty) Get() *Property {
	return v.value
}

func (v *NullableProperty) Set(val *Property) {
	v.value = val
	v.isSet = true
}

func (v NullableProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProperty(val *Property) *NullableProperty {
	return &NullableProperty{value: val, isSet: true}
}

func (v NullableProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
