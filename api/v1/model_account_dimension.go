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

// checks if the AccountDimension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountDimension{}

// AccountDimension An single dimension of a financial account
type AccountDimension struct {
	// Unique and immutable ID attribute of the entity that is generated when the instance is created. The ID is unique within the system accross all accounts and it can be used to reference the entity in other entities or to retrieve it from the backend.
	Id string `json:"id"`
	// The date and time the entity was created
	CreatedAt int64 `json:"created_at"`
	// The date and time the entity was last updated
	UpdatedAt int64 `json:"updated_at"`
	// The name of the dimension
	Name string `json:"name"`
	// A short description of the dimension
	Description *string `json:"description,omitempty"`
	// The ID of a parent dimension, if any
	ParentDimensionRef *string `json:"parent_dimension_ref,omitempty"`
	// If set to a value greater than 0, the values of the dimension will be padded with 0 to match the size. E.g. if this is set to 3, a dimension value '42' would be padded to '042'. If not set, values will not be padded.
	PadToSize *int32                            `json:"pad_to_size,omitempty"`
	RefType   AccountDimensionReferenceTypeEnum `json:"ref_type"`
	// Indicates whether providing a value for the dimension is required when adding an entry to an account
	Required *bool `json:"required,omitempty"`
	// A JS script that generates a value for the dimension.
	ValueGenerator *string `json:"value_generator,omitempty"`
	// Optional nested dimensions for this dimension. This is lazy loaded, that is, it is returned only if a single dimension is loaded.
	SubDimensions []AccountDimension `json:"sub_dimensions,omitempty"`
	// All available values for this dimension. This is lazy loaded, that is, only if a single dimension is loaded.
	Values []AccountDimensionValue `json:"values,omitempty"`
}

type _AccountDimension AccountDimension

// NewAccountDimension instantiates a new AccountDimension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountDimension(id string, createdAt int64, updatedAt int64, name string, refType AccountDimensionReferenceTypeEnum) *AccountDimension {
	this := AccountDimension{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Name = name
	this.RefType = refType
	return &this
}

// NewAccountDimensionWithDefaults instantiates a new AccountDimension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountDimensionWithDefaults() *AccountDimension {
	this := AccountDimension{}
	return &this
}

// GetId returns the Id field value
func (o *AccountDimension) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AccountDimension) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AccountDimension) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AccountDimension) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AccountDimension) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AccountDimension) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AccountDimension) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AccountDimension) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AccountDimension) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetName returns the Name field value
func (o *AccountDimension) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AccountDimension) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AccountDimension) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccountDimension) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDimension) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccountDimension) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccountDimension) SetDescription(v string) {
	o.Description = &v
}

// GetParentDimensionRef returns the ParentDimensionRef field value if set, zero value otherwise.
func (o *AccountDimension) GetParentDimensionRef() string {
	if o == nil || IsNil(o.ParentDimensionRef) {
		var ret string
		return ret
	}
	return *o.ParentDimensionRef
}

// GetParentDimensionRefOk returns a tuple with the ParentDimensionRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDimension) GetParentDimensionRefOk() (*string, bool) {
	if o == nil || IsNil(o.ParentDimensionRef) {
		return nil, false
	}
	return o.ParentDimensionRef, true
}

// HasParentDimensionRef returns a boolean if a field has been set.
func (o *AccountDimension) HasParentDimensionRef() bool {
	if o != nil && !IsNil(o.ParentDimensionRef) {
		return true
	}

	return false
}

// SetParentDimensionRef gets a reference to the given string and assigns it to the ParentDimensionRef field.
func (o *AccountDimension) SetParentDimensionRef(v string) {
	o.ParentDimensionRef = &v
}

// GetPadToSize returns the PadToSize field value if set, zero value otherwise.
func (o *AccountDimension) GetPadToSize() int32 {
	if o == nil || IsNil(o.PadToSize) {
		var ret int32
		return ret
	}
	return *o.PadToSize
}

// GetPadToSizeOk returns a tuple with the PadToSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDimension) GetPadToSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PadToSize) {
		return nil, false
	}
	return o.PadToSize, true
}

// HasPadToSize returns a boolean if a field has been set.
func (o *AccountDimension) HasPadToSize() bool {
	if o != nil && !IsNil(o.PadToSize) {
		return true
	}

	return false
}

// SetPadToSize gets a reference to the given int32 and assigns it to the PadToSize field.
func (o *AccountDimension) SetPadToSize(v int32) {
	o.PadToSize = &v
}

// GetRefType returns the RefType field value
func (o *AccountDimension) GetRefType() AccountDimensionReferenceTypeEnum {
	if o == nil {
		var ret AccountDimensionReferenceTypeEnum
		return ret
	}

	return o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value
// and a boolean to check if the value has been set.
func (o *AccountDimension) GetRefTypeOk() (*AccountDimensionReferenceTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefType, true
}

// SetRefType sets field value
func (o *AccountDimension) SetRefType(v AccountDimensionReferenceTypeEnum) {
	o.RefType = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *AccountDimension) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDimension) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *AccountDimension) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *AccountDimension) SetRequired(v bool) {
	o.Required = &v
}

// GetValueGenerator returns the ValueGenerator field value if set, zero value otherwise.
func (o *AccountDimension) GetValueGenerator() string {
	if o == nil || IsNil(o.ValueGenerator) {
		var ret string
		return ret
	}
	return *o.ValueGenerator
}

// GetValueGeneratorOk returns a tuple with the ValueGenerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDimension) GetValueGeneratorOk() (*string, bool) {
	if o == nil || IsNil(o.ValueGenerator) {
		return nil, false
	}
	return o.ValueGenerator, true
}

// HasValueGenerator returns a boolean if a field has been set.
func (o *AccountDimension) HasValueGenerator() bool {
	if o != nil && !IsNil(o.ValueGenerator) {
		return true
	}

	return false
}

// SetValueGenerator gets a reference to the given string and assigns it to the ValueGenerator field.
func (o *AccountDimension) SetValueGenerator(v string) {
	o.ValueGenerator = &v
}

// GetSubDimensions returns the SubDimensions field value if set, zero value otherwise.
func (o *AccountDimension) GetSubDimensions() []AccountDimension {
	if o == nil || IsNil(o.SubDimensions) {
		var ret []AccountDimension
		return ret
	}
	return o.SubDimensions
}

// GetSubDimensionsOk returns a tuple with the SubDimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDimension) GetSubDimensionsOk() ([]AccountDimension, bool) {
	if o == nil || IsNil(o.SubDimensions) {
		return nil, false
	}
	return o.SubDimensions, true
}

// HasSubDimensions returns a boolean if a field has been set.
func (o *AccountDimension) HasSubDimensions() bool {
	if o != nil && !IsNil(o.SubDimensions) {
		return true
	}

	return false
}

// SetSubDimensions gets a reference to the given []AccountDimension and assigns it to the SubDimensions field.
func (o *AccountDimension) SetSubDimensions(v []AccountDimension) {
	o.SubDimensions = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *AccountDimension) GetValues() []AccountDimensionValue {
	if o == nil || IsNil(o.Values) {
		var ret []AccountDimensionValue
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDimension) GetValuesOk() ([]AccountDimensionValue, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *AccountDimension) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []AccountDimensionValue and assigns it to the Values field.
func (o *AccountDimension) SetValues(v []AccountDimensionValue) {
	o.Values = v
}

func (o AccountDimension) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountDimension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ParentDimensionRef) {
		toSerialize["parent_dimension_ref"] = o.ParentDimensionRef
	}
	if !IsNil(o.PadToSize) {
		toSerialize["pad_to_size"] = o.PadToSize
	}
	toSerialize["ref_type"] = o.RefType
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.ValueGenerator) {
		toSerialize["value_generator"] = o.ValueGenerator
	}
	if !IsNil(o.SubDimensions) {
		toSerialize["sub_dimensions"] = o.SubDimensions
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

func (o *AccountDimension) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"name",
		"ref_type",
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

	varAccountDimension := _AccountDimension{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountDimension)

	if err != nil {
		return err
	}

	*o = AccountDimension(varAccountDimension)

	return err
}

type NullableAccountDimension struct {
	value *AccountDimension
	isSet bool
}

func (v NullableAccountDimension) Get() *AccountDimension {
	return v.value
}

func (v *NullableAccountDimension) Set(val *AccountDimension) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountDimension) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountDimension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountDimension(val *AccountDimension) *NullableAccountDimension {
	return &NullableAccountDimension{value: val, isSet: true}
}

func (v NullableAccountDimension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountDimension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
