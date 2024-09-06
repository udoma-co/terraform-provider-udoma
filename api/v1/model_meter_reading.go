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

// checks if the MeterReading type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeterReading{}

// MeterReading a single reading of a meter
type MeterReading struct {
	// Unique and immutable ID attribute of the entity that is generated when the instance is created. The ID is unique within the system accross all accounts and it can be used to reference the entity in other entities or to retrieve it from the backend.
	Id string `json:"id"`
	// The date and time the entity was created
	CreatedAt int64 `json:"created_at"`
	// The date and time the entity was last updated
	UpdatedAt     int64        `json:"updated_at"`
	Date          *int64       `json:"date,omitempty"`
	Value         *FloatNumber `json:"value,omitempty"`
	AttachmentRef *string      `json:"attachment_ref,omitempty"`
	MeterRef      string       `json:"meter_ref"`
	ReporterRef   string       `json:"reporter_ref"`
}

type _MeterReading MeterReading

// NewMeterReading instantiates a new MeterReading object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeterReading(id string, createdAt int64, updatedAt int64, meterRef string, reporterRef string) *MeterReading {
	this := MeterReading{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.MeterRef = meterRef
	this.ReporterRef = reporterRef
	return &this
}

// NewMeterReadingWithDefaults instantiates a new MeterReading object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeterReadingWithDefaults() *MeterReading {
	this := MeterReading{}
	return &this
}

// GetId returns the Id field value
func (o *MeterReading) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MeterReading) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MeterReading) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *MeterReading) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MeterReading) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *MeterReading) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *MeterReading) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *MeterReading) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *MeterReading) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *MeterReading) GetDate() int64 {
	if o == nil || IsNil(o.Date) {
		var ret int64
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeterReading) GetDateOk() (*int64, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *MeterReading) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given int64 and assigns it to the Date field.
func (o *MeterReading) SetDate(v int64) {
	o.Date = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MeterReading) GetValue() FloatNumber {
	if o == nil || IsNil(o.Value) {
		var ret FloatNumber
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeterReading) GetValueOk() (*FloatNumber, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MeterReading) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given FloatNumber and assigns it to the Value field.
func (o *MeterReading) SetValue(v FloatNumber) {
	o.Value = &v
}

// GetAttachmentRef returns the AttachmentRef field value if set, zero value otherwise.
func (o *MeterReading) GetAttachmentRef() string {
	if o == nil || IsNil(o.AttachmentRef) {
		var ret string
		return ret
	}
	return *o.AttachmentRef
}

// GetAttachmentRefOk returns a tuple with the AttachmentRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeterReading) GetAttachmentRefOk() (*string, bool) {
	if o == nil || IsNil(o.AttachmentRef) {
		return nil, false
	}
	return o.AttachmentRef, true
}

// HasAttachmentRef returns a boolean if a field has been set.
func (o *MeterReading) HasAttachmentRef() bool {
	if o != nil && !IsNil(o.AttachmentRef) {
		return true
	}

	return false
}

// SetAttachmentRef gets a reference to the given string and assigns it to the AttachmentRef field.
func (o *MeterReading) SetAttachmentRef(v string) {
	o.AttachmentRef = &v
}

// GetMeterRef returns the MeterRef field value
func (o *MeterReading) GetMeterRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MeterRef
}

// GetMeterRefOk returns a tuple with the MeterRef field value
// and a boolean to check if the value has been set.
func (o *MeterReading) GetMeterRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeterRef, true
}

// SetMeterRef sets field value
func (o *MeterReading) SetMeterRef(v string) {
	o.MeterRef = v
}

// GetReporterRef returns the ReporterRef field value
func (o *MeterReading) GetReporterRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReporterRef
}

// GetReporterRefOk returns a tuple with the ReporterRef field value
// and a boolean to check if the value has been set.
func (o *MeterReading) GetReporterRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReporterRef, true
}

// SetReporterRef sets field value
func (o *MeterReading) SetReporterRef(v string) {
	o.ReporterRef = v
}

func (o MeterReading) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeterReading) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.AttachmentRef) {
		toSerialize["attachment_ref"] = o.AttachmentRef
	}
	toSerialize["meter_ref"] = o.MeterRef
	toSerialize["reporter_ref"] = o.ReporterRef
	return toSerialize, nil
}

func (o *MeterReading) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"meter_ref",
		"reporter_ref",
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

	varMeterReading := _MeterReading{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMeterReading)

	if err != nil {
		return err
	}

	*o = MeterReading(varMeterReading)

	return err
}

type NullableMeterReading struct {
	value *MeterReading
	isSet bool
}

func (v NullableMeterReading) Get() *MeterReading {
	return v.value
}

func (v *NullableMeterReading) Set(val *MeterReading) {
	v.value = val
	v.isSet = true
}

func (v NullableMeterReading) IsSet() bool {
	return v.isSet
}

func (v *NullableMeterReading) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeterReading(val *MeterReading) *NullableMeterReading {
	return &NullableMeterReading{value: val, isSet: true}
}

func (v NullableMeterReading) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeterReading) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
