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

// checks if the PriceIndexPointQueryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PriceIndexPointQueryRequest{}

// PriceIndexPointQueryRequest The request used to query a price index point
type PriceIndexPointQueryRequest struct {
	CountryCode *string `json:"country_code,omitempty"`
	// The date to look for.
	Date *int64 `json:"date,omitempty"`
}

// NewPriceIndexPointQueryRequest instantiates a new PriceIndexPointQueryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPriceIndexPointQueryRequest() *PriceIndexPointQueryRequest {
	this := PriceIndexPointQueryRequest{}
	return &this
}

// NewPriceIndexPointQueryRequestWithDefaults instantiates a new PriceIndexPointQueryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPriceIndexPointQueryRequestWithDefaults() *PriceIndexPointQueryRequest {
	this := PriceIndexPointQueryRequest{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *PriceIndexPointQueryRequest) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceIndexPointQueryRequest) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *PriceIndexPointQueryRequest) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *PriceIndexPointQueryRequest) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *PriceIndexPointQueryRequest) GetDate() int64 {
	if o == nil || IsNil(o.Date) {
		var ret int64
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PriceIndexPointQueryRequest) GetDateOk() (*int64, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *PriceIndexPointQueryRequest) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given int64 and assigns it to the Date field.
func (o *PriceIndexPointQueryRequest) SetDate(v int64) {
	o.Date = &v
}

func (o PriceIndexPointQueryRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PriceIndexPointQueryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	return toSerialize, nil
}

type NullablePriceIndexPointQueryRequest struct {
	value *PriceIndexPointQueryRequest
	isSet bool
}

func (v NullablePriceIndexPointQueryRequest) Get() *PriceIndexPointQueryRequest {
	return v.value
}

func (v *NullablePriceIndexPointQueryRequest) Set(val *PriceIndexPointQueryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceIndexPointQueryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceIndexPointQueryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceIndexPointQueryRequest(val *PriceIndexPointQueryRequest) *NullablePriceIndexPointQueryRequest {
	return &NullablePriceIndexPointQueryRequest{value: val, isSet: true}
}

func (v NullablePriceIndexPointQueryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceIndexPointQueryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
