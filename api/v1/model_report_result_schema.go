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

// checks if the ReportResultSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportResultSchema{}

// ReportResultSchema struct for ReportResultSchema
type ReportResultSchema struct {
	ResultType ReportResultTypeEnum `json:"result_type"`
	// The attribute that will be used as the ID of each item in the result set.  When set, the table will be clickable and the value of the respective  attribute can be used to navigate to the detail view of the item.
	TableRowIdAttribute *string `json:"table_row_id_attribute,omitempty"`
	// The attributes within the result. Each attribute will be displayed  as a column in the UI, if the result is a list, or as a field if not.
	Attributes []ReportResultSchemaAttribute `json:"attributes"`
}

type _ReportResultSchema ReportResultSchema

// NewReportResultSchema instantiates a new ReportResultSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportResultSchema(resultType ReportResultTypeEnum, attributes []ReportResultSchemaAttribute) *ReportResultSchema {
	this := ReportResultSchema{}
	this.ResultType = resultType
	this.Attributes = attributes
	return &this
}

// NewReportResultSchemaWithDefaults instantiates a new ReportResultSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportResultSchemaWithDefaults() *ReportResultSchema {
	this := ReportResultSchema{}
	return &this
}

// GetResultType returns the ResultType field value
func (o *ReportResultSchema) GetResultType() ReportResultTypeEnum {
	if o == nil {
		var ret ReportResultTypeEnum
		return ret
	}

	return o.ResultType
}

// GetResultTypeOk returns a tuple with the ResultType field value
// and a boolean to check if the value has been set.
func (o *ReportResultSchema) GetResultTypeOk() (*ReportResultTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultType, true
}

// SetResultType sets field value
func (o *ReportResultSchema) SetResultType(v ReportResultTypeEnum) {
	o.ResultType = v
}

// GetTableRowIdAttribute returns the TableRowIdAttribute field value if set, zero value otherwise.
func (o *ReportResultSchema) GetTableRowIdAttribute() string {
	if o == nil || IsNil(o.TableRowIdAttribute) {
		var ret string
		return ret
	}
	return *o.TableRowIdAttribute
}

// GetTableRowIdAttributeOk returns a tuple with the TableRowIdAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportResultSchema) GetTableRowIdAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.TableRowIdAttribute) {
		return nil, false
	}
	return o.TableRowIdAttribute, true
}

// HasTableRowIdAttribute returns a boolean if a field has been set.
func (o *ReportResultSchema) HasTableRowIdAttribute() bool {
	if o != nil && !IsNil(o.TableRowIdAttribute) {
		return true
	}

	return false
}

// SetTableRowIdAttribute gets a reference to the given string and assigns it to the TableRowIdAttribute field.
func (o *ReportResultSchema) SetTableRowIdAttribute(v string) {
	o.TableRowIdAttribute = &v
}

// GetAttributes returns the Attributes field value
func (o *ReportResultSchema) GetAttributes() []ReportResultSchemaAttribute {
	if o == nil {
		var ret []ReportResultSchemaAttribute
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ReportResultSchema) GetAttributesOk() ([]ReportResultSchemaAttribute, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *ReportResultSchema) SetAttributes(v []ReportResultSchemaAttribute) {
	o.Attributes = v
}

func (o ReportResultSchema) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportResultSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result_type"] = o.ResultType
	if !IsNil(o.TableRowIdAttribute) {
		toSerialize["table_row_id_attribute"] = o.TableRowIdAttribute
	}
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *ReportResultSchema) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result_type",
		"attributes",
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

	varReportResultSchema := _ReportResultSchema{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReportResultSchema)

	if err != nil {
		return err
	}

	*o = ReportResultSchema(varReportResultSchema)

	return err
}

type NullableReportResultSchema struct {
	value *ReportResultSchema
	isSet bool
}

func (v NullableReportResultSchema) Get() *ReportResultSchema {
	return v.value
}

func (v *NullableReportResultSchema) Set(val *ReportResultSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableReportResultSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableReportResultSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportResultSchema(val *ReportResultSchema) *NullableReportResultSchema {
	return &NullableReportResultSchema{value: val, isSet: true}
}

func (v NullableReportResultSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportResultSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
