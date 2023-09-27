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

// Meter a meter installed in a property to measure some kind of utility amount
type Meter struct {
	Id *string `json:"id,omitempty"`
	// The timestamp of when the meter was created
	CreatedAt *int64 `json:"created_at,omitempty"`
	// The timestamp of when the meter was last updated
	UpdatedAt           *int64         `json:"updated_at,omitempty"`
	PropertyRef         *string        `json:"property_ref,omitempty"`
	MeterType           *MeterTypeEnum `json:"meter_type,omitempty"`
	Unit                *string        `json:"unit,omitempty"`
	Name                *string        `json:"name,omitempty"`
	Location            *string        `json:"location,omitempty"`
	SerialNumber        *string        `json:"serial_number,omitempty"`
	CalibrationDate     *int64         `json:"calibration_date,omitempty"`
	CalibrationValidity *int64         `json:"calibration_validity,omitempty"`
}

// NewMeter instantiates a new Meter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeter() *Meter {
	this := Meter{}
	return &this
}

// NewMeterWithDefaults instantiates a new Meter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeterWithDefaults() *Meter {
	this := Meter{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Meter) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meter) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Meter) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Meter) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Meter) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meter) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Meter) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *Meter) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Meter) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meter) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Meter) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *Meter) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetPropertyRef returns the PropertyRef field value if set, zero value otherwise.
func (o *Meter) GetPropertyRef() string {
	if o == nil || o.PropertyRef == nil {
		var ret string
		return ret
	}
	return *o.PropertyRef
}

// GetPropertyRefOk returns a tuple with the PropertyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meter) GetPropertyRefOk() (*string, bool) {
	if o == nil || o.PropertyRef == nil {
		return nil, false
	}
	return o.PropertyRef, true
}

// HasPropertyRef returns a boolean if a field has been set.
func (o *Meter) HasPropertyRef() bool {
	if o != nil && o.PropertyRef != nil {
		return true
	}

	return false
}

// SetPropertyRef gets a reference to the given string and assigns it to the PropertyRef field.
func (o *Meter) SetPropertyRef(v string) {
	o.PropertyRef = &v
}

// GetMeterType returns the MeterType field value if set, zero value otherwise.
func (o *Meter) GetMeterType() MeterTypeEnum {
	if o == nil || o.MeterType == nil {
		var ret MeterTypeEnum
		return ret
	}
	return *o.MeterType
}

// GetMeterTypeOk returns a tuple with the MeterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meter) GetMeterTypeOk() (*MeterTypeEnum, bool) {
	if o == nil || o.MeterType == nil {
		return nil, false
	}
	return o.MeterType, true
}

// HasMeterType returns a boolean if a field has been set.
func (o *Meter) HasMeterType() bool {
	if o != nil && o.MeterType != nil {
		return true
	}

	return false
}

// SetMeterType gets a reference to the given MeterTypeEnum and assigns it to the MeterType field.
func (o *Meter) SetMeterType(v MeterTypeEnum) {
	o.MeterType = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *Meter) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meter) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *Meter) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *Meter) SetUnit(v string) {
	o.Unit = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Meter) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meter) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Meter) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Meter) SetName(v string) {
	o.Name = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Meter) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meter) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Meter) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *Meter) SetLocation(v string) {
	o.Location = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *Meter) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meter) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *Meter) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *Meter) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetCalibrationDate returns the CalibrationDate field value if set, zero value otherwise.
func (o *Meter) GetCalibrationDate() int64 {
	if o == nil || o.CalibrationDate == nil {
		var ret int64
		return ret
	}
	return *o.CalibrationDate
}

// GetCalibrationDateOk returns a tuple with the CalibrationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meter) GetCalibrationDateOk() (*int64, bool) {
	if o == nil || o.CalibrationDate == nil {
		return nil, false
	}
	return o.CalibrationDate, true
}

// HasCalibrationDate returns a boolean if a field has been set.
func (o *Meter) HasCalibrationDate() bool {
	if o != nil && o.CalibrationDate != nil {
		return true
	}

	return false
}

// SetCalibrationDate gets a reference to the given int64 and assigns it to the CalibrationDate field.
func (o *Meter) SetCalibrationDate(v int64) {
	o.CalibrationDate = &v
}

// GetCalibrationValidity returns the CalibrationValidity field value if set, zero value otherwise.
func (o *Meter) GetCalibrationValidity() int64 {
	if o == nil || o.CalibrationValidity == nil {
		var ret int64
		return ret
	}
	return *o.CalibrationValidity
}

// GetCalibrationValidityOk returns a tuple with the CalibrationValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Meter) GetCalibrationValidityOk() (*int64, bool) {
	if o == nil || o.CalibrationValidity == nil {
		return nil, false
	}
	return o.CalibrationValidity, true
}

// HasCalibrationValidity returns a boolean if a field has been set.
func (o *Meter) HasCalibrationValidity() bool {
	if o != nil && o.CalibrationValidity != nil {
		return true
	}

	return false
}

// SetCalibrationValidity gets a reference to the given int64 and assigns it to the CalibrationValidity field.
func (o *Meter) SetCalibrationValidity(v int64) {
	o.CalibrationValidity = &v
}

func (o Meter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.PropertyRef != nil {
		toSerialize["property_ref"] = o.PropertyRef
	}
	if o.MeterType != nil {
		toSerialize["meter_type"] = o.MeterType
	}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.SerialNumber != nil {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if o.CalibrationDate != nil {
		toSerialize["calibration_date"] = o.CalibrationDate
	}
	if o.CalibrationValidity != nil {
		toSerialize["calibration_validity"] = o.CalibrationValidity
	}
	return json.Marshal(toSerialize)
}

type NullableMeter struct {
	value *Meter
	isSet bool
}

func (v NullableMeter) Get() *Meter {
	return v.value
}

func (v *NullableMeter) Set(val *Meter) {
	v.value = val
	v.isSet = true
}

func (v NullableMeter) IsSet() bool {
	return v.isSet
}

func (v *NullableMeter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeter(val *Meter) *NullableMeter {
	return &NullableMeter{value: val, isSet: true}
}

func (v NullableMeter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
