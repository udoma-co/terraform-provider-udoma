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

// checks if the AccessRecordData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRecordData{}

// AccessRecordData An object containing all the data we can collect from the request
type AccessRecordData struct {
	// The IP Address used to retrieve the document.
	Ip string `json:"ip"`
	// The agent used
	UserAgent string `json:"user_agent"`
	// The referer header in the request made
	Referer string `json:"referer"`
	// The requested URI
	Requested string `json:"requested"`
}

type _AccessRecordData AccessRecordData

// NewAccessRecordData instantiates a new AccessRecordData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRecordData(ip string, userAgent string, referer string, requested string) *AccessRecordData {
	this := AccessRecordData{}
	this.Ip = ip
	this.UserAgent = userAgent
	this.Referer = referer
	this.Requested = requested
	return &this
}

// NewAccessRecordDataWithDefaults instantiates a new AccessRecordData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRecordDataWithDefaults() *AccessRecordData {
	this := AccessRecordData{}
	return &this
}

// GetIp returns the Ip field value
func (o *AccessRecordData) GetIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *AccessRecordData) GetIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *AccessRecordData) SetIp(v string) {
	o.Ip = v
}

// GetUserAgent returns the UserAgent field value
func (o *AccessRecordData) GetUserAgent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value
// and a boolean to check if the value has been set.
func (o *AccessRecordData) GetUserAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserAgent, true
}

// SetUserAgent sets field value
func (o *AccessRecordData) SetUserAgent(v string) {
	o.UserAgent = v
}

// GetReferer returns the Referer field value
func (o *AccessRecordData) GetReferer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Referer
}

// GetRefererOk returns a tuple with the Referer field value
// and a boolean to check if the value has been set.
func (o *AccessRecordData) GetRefererOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Referer, true
}

// SetReferer sets field value
func (o *AccessRecordData) SetReferer(v string) {
	o.Referer = v
}

// GetRequested returns the Requested field value
func (o *AccessRecordData) GetRequested() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value
// and a boolean to check if the value has been set.
func (o *AccessRecordData) GetRequestedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requested, true
}

// SetRequested sets field value
func (o *AccessRecordData) SetRequested(v string) {
	o.Requested = v
}

func (o AccessRecordData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRecordData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ip"] = o.Ip
	toSerialize["user_agent"] = o.UserAgent
	toSerialize["referer"] = o.Referer
	toSerialize["requested"] = o.Requested
	return toSerialize, nil
}

func (o *AccessRecordData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ip",
		"user_agent",
		"referer",
		"requested",
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

	varAccessRecordData := _AccessRecordData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessRecordData)

	if err != nil {
		return err
	}

	*o = AccessRecordData(varAccessRecordData)

	return err
}

type NullableAccessRecordData struct {
	value *AccessRecordData
	isSet bool
}

func (v NullableAccessRecordData) Get() *AccessRecordData {
	return v.value
}

func (v *NullableAccessRecordData) Set(val *AccessRecordData) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRecordData) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRecordData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRecordData(val *AccessRecordData) *NullableAccessRecordData {
	return &NullableAccessRecordData{value: val, isSet: true}
}

func (v NullableAccessRecordData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRecordData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}