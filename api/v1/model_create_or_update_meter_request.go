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

// checks if the CreateOrUpdateMeterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateMeterRequest{}

// CreateOrUpdateMeterRequest All the information required for creating a new meter
type CreateOrUpdateMeterRequest struct {
	PropertyRef         string        `json:"property_ref"`
	MeterType           MeterTypeEnum `json:"meter_type"`
	Unit                string        `json:"unit"`
	Name                string        `json:"name"`
	Location            *string       `json:"location,omitempty"`
	SerialNumber        string        `json:"serial_number"`
	CalibrationDate     *int64        `json:"calibration_date,omitempty"`
	CalibrationValidity *int64        `json:"calibration_validity,omitempty"`
}

type _CreateOrUpdateMeterRequest CreateOrUpdateMeterRequest

// NewCreateOrUpdateMeterRequest instantiates a new CreateOrUpdateMeterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateMeterRequest(propertyRef string, meterType MeterTypeEnum, unit string, name string, serialNumber string) *CreateOrUpdateMeterRequest {
	this := CreateOrUpdateMeterRequest{}
	this.PropertyRef = propertyRef
	this.MeterType = meterType
	this.Unit = unit
	this.Name = name
	this.SerialNumber = serialNumber
	return &this
}

// NewCreateOrUpdateMeterRequestWithDefaults instantiates a new CreateOrUpdateMeterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateMeterRequestWithDefaults() *CreateOrUpdateMeterRequest {
	this := CreateOrUpdateMeterRequest{}
	return &this
}

// GetPropertyRef returns the PropertyRef field value
func (o *CreateOrUpdateMeterRequest) GetPropertyRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PropertyRef
}

// GetPropertyRefOk returns a tuple with the PropertyRef field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateMeterRequest) GetPropertyRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PropertyRef, true
}

// SetPropertyRef sets field value
func (o *CreateOrUpdateMeterRequest) SetPropertyRef(v string) {
	o.PropertyRef = v
}

// GetMeterType returns the MeterType field value
func (o *CreateOrUpdateMeterRequest) GetMeterType() MeterTypeEnum {
	if o == nil {
		var ret MeterTypeEnum
		return ret
	}

	return o.MeterType
}

// GetMeterTypeOk returns a tuple with the MeterType field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateMeterRequest) GetMeterTypeOk() (*MeterTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MeterType, true
}

// SetMeterType sets field value
func (o *CreateOrUpdateMeterRequest) SetMeterType(v MeterTypeEnum) {
	o.MeterType = v
}

// GetUnit returns the Unit field value
func (o *CreateOrUpdateMeterRequest) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateMeterRequest) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *CreateOrUpdateMeterRequest) SetUnit(v string) {
	o.Unit = v
}

// GetName returns the Name field value
func (o *CreateOrUpdateMeterRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateMeterRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrUpdateMeterRequest) SetName(v string) {
	o.Name = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *CreateOrUpdateMeterRequest) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateMeterRequest) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *CreateOrUpdateMeterRequest) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *CreateOrUpdateMeterRequest) SetLocation(v string) {
	o.Location = &v
}

// GetSerialNumber returns the SerialNumber field value
func (o *CreateOrUpdateMeterRequest) GetSerialNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateMeterRequest) GetSerialNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SerialNumber, true
}

// SetSerialNumber sets field value
func (o *CreateOrUpdateMeterRequest) SetSerialNumber(v string) {
	o.SerialNumber = v
}

// GetCalibrationDate returns the CalibrationDate field value if set, zero value otherwise.
func (o *CreateOrUpdateMeterRequest) GetCalibrationDate() int64 {
	if o == nil || IsNil(o.CalibrationDate) {
		var ret int64
		return ret
	}
	return *o.CalibrationDate
}

// GetCalibrationDateOk returns a tuple with the CalibrationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateMeterRequest) GetCalibrationDateOk() (*int64, bool) {
	if o == nil || IsNil(o.CalibrationDate) {
		return nil, false
	}
	return o.CalibrationDate, true
}

// HasCalibrationDate returns a boolean if a field has been set.
func (o *CreateOrUpdateMeterRequest) HasCalibrationDate() bool {
	if o != nil && !IsNil(o.CalibrationDate) {
		return true
	}

	return false
}

// SetCalibrationDate gets a reference to the given int64 and assigns it to the CalibrationDate field.
func (o *CreateOrUpdateMeterRequest) SetCalibrationDate(v int64) {
	o.CalibrationDate = &v
}

// GetCalibrationValidity returns the CalibrationValidity field value if set, zero value otherwise.
func (o *CreateOrUpdateMeterRequest) GetCalibrationValidity() int64 {
	if o == nil || IsNil(o.CalibrationValidity) {
		var ret int64
		return ret
	}
	return *o.CalibrationValidity
}

// GetCalibrationValidityOk returns a tuple with the CalibrationValidity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateMeterRequest) GetCalibrationValidityOk() (*int64, bool) {
	if o == nil || IsNil(o.CalibrationValidity) {
		return nil, false
	}
	return o.CalibrationValidity, true
}

// HasCalibrationValidity returns a boolean if a field has been set.
func (o *CreateOrUpdateMeterRequest) HasCalibrationValidity() bool {
	if o != nil && !IsNil(o.CalibrationValidity) {
		return true
	}

	return false
}

// SetCalibrationValidity gets a reference to the given int64 and assigns it to the CalibrationValidity field.
func (o *CreateOrUpdateMeterRequest) SetCalibrationValidity(v int64) {
	o.CalibrationValidity = &v
}

func (o CreateOrUpdateMeterRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateMeterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["property_ref"] = o.PropertyRef
	toSerialize["meter_type"] = o.MeterType
	toSerialize["unit"] = o.Unit
	toSerialize["name"] = o.Name
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	toSerialize["serial_number"] = o.SerialNumber
	if !IsNil(o.CalibrationDate) {
		toSerialize["calibration_date"] = o.CalibrationDate
	}
	if !IsNil(o.CalibrationValidity) {
		toSerialize["calibration_validity"] = o.CalibrationValidity
	}
	return toSerialize, nil
}

func (o *CreateOrUpdateMeterRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"property_ref",
		"meter_type",
		"unit",
		"name",
		"serial_number",
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

	varCreateOrUpdateMeterRequest := _CreateOrUpdateMeterRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrUpdateMeterRequest)

	if err != nil {
		return err
	}

	*o = CreateOrUpdateMeterRequest(varCreateOrUpdateMeterRequest)

	return err
}

type NullableCreateOrUpdateMeterRequest struct {
	value *CreateOrUpdateMeterRequest
	isSet bool
}

func (v NullableCreateOrUpdateMeterRequest) Get() *CreateOrUpdateMeterRequest {
	return v.value
}

func (v *NullableCreateOrUpdateMeterRequest) Set(val *CreateOrUpdateMeterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateMeterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateMeterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateMeterRequest(val *CreateOrUpdateMeterRequest) *NullableCreateOrUpdateMeterRequest {
	return &NullableCreateOrUpdateMeterRequest{value: val, isSet: true}
}

func (v NullableCreateOrUpdateMeterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateMeterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
