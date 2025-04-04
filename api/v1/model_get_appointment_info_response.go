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

// checks if the GetAppointmentInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAppointmentInfoResponse{}

// GetAppointmentInfoResponse struct for GetAppointmentInfoResponse
type GetAppointmentInfoResponse struct {
	// A descriptive name for the appointment, displayed on the calendar overview
	Name *string `json:"name,omitempty"`
	// The timestamp of the beginning of the time slot
	StartTime *int64 `json:"start_time,omitempty"`
	// The timestamp of the end of the time slot
	EndTime         *int64   `json:"end_time,omitempty"`
	PropertyAddress *Address `json:"property_address,omitempty"`
}

// NewGetAppointmentInfoResponse instantiates a new GetAppointmentInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAppointmentInfoResponse() *GetAppointmentInfoResponse {
	this := GetAppointmentInfoResponse{}
	return &this
}

// NewGetAppointmentInfoResponseWithDefaults instantiates a new GetAppointmentInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAppointmentInfoResponseWithDefaults() *GetAppointmentInfoResponse {
	this := GetAppointmentInfoResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetAppointmentInfoResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppointmentInfoResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetAppointmentInfoResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetAppointmentInfoResponse) SetName(v string) {
	o.Name = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *GetAppointmentInfoResponse) GetStartTime() int64 {
	if o == nil || IsNil(o.StartTime) {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppointmentInfoResponse) GetStartTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *GetAppointmentInfoResponse) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *GetAppointmentInfoResponse) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *GetAppointmentInfoResponse) GetEndTime() int64 {
	if o == nil || IsNil(o.EndTime) {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppointmentInfoResponse) GetEndTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *GetAppointmentInfoResponse) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *GetAppointmentInfoResponse) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetPropertyAddress returns the PropertyAddress field value if set, zero value otherwise.
func (o *GetAppointmentInfoResponse) GetPropertyAddress() Address {
	if o == nil || IsNil(o.PropertyAddress) {
		var ret Address
		return ret
	}
	return *o.PropertyAddress
}

// GetPropertyAddressOk returns a tuple with the PropertyAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAppointmentInfoResponse) GetPropertyAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.PropertyAddress) {
		return nil, false
	}
	return o.PropertyAddress, true
}

// HasPropertyAddress returns a boolean if a field has been set.
func (o *GetAppointmentInfoResponse) HasPropertyAddress() bool {
	if o != nil && !IsNil(o.PropertyAddress) {
		return true
	}

	return false
}

// SetPropertyAddress gets a reference to the given Address and assigns it to the PropertyAddress field.
func (o *GetAppointmentInfoResponse) SetPropertyAddress(v Address) {
	o.PropertyAddress = &v
}

func (o GetAppointmentInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAppointmentInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.PropertyAddress) {
		toSerialize["property_address"] = o.PropertyAddress
	}
	return toSerialize, nil
}

type NullableGetAppointmentInfoResponse struct {
	value *GetAppointmentInfoResponse
	isSet bool
}

func (v NullableGetAppointmentInfoResponse) Get() *GetAppointmentInfoResponse {
	return v.value
}

func (v *NullableGetAppointmentInfoResponse) Set(val *GetAppointmentInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAppointmentInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAppointmentInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAppointmentInfoResponse(val *GetAppointmentInfoResponse) *NullableGetAppointmentInfoResponse {
	return &NullableGetAppointmentInfoResponse{value: val, isSet: true}
}

func (v NullableGetAppointmentInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAppointmentInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
