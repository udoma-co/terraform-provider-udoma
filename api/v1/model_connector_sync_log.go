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

// checks if the ConnectorSyncLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorSyncLog{}

// ConnectorSyncLog Used by connectors to push error logs to the backend.
type ConnectorSyncLog struct {
	// timestamp of the error
	Timestamp *int64                 `json:"timestamp,omitempty"`
	Severity  *ConnectorLogLevelEnum `json:"severity,omitempty"`
	// log message
	Message *string `json:"message,omitempty"`
}

// NewConnectorSyncLog instantiates a new ConnectorSyncLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorSyncLog() *ConnectorSyncLog {
	this := ConnectorSyncLog{}
	return &this
}

// NewConnectorSyncLogWithDefaults instantiates a new ConnectorSyncLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorSyncLogWithDefaults() *ConnectorSyncLog {
	this := ConnectorSyncLog{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ConnectorSyncLog) GetTimestamp() int64 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSyncLog) GetTimestampOk() (*int64, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ConnectorSyncLog) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ConnectorSyncLog) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ConnectorSyncLog) GetSeverity() ConnectorLogLevelEnum {
	if o == nil || IsNil(o.Severity) {
		var ret ConnectorLogLevelEnum
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSyncLog) GetSeverityOk() (*ConnectorLogLevelEnum, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ConnectorSyncLog) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given ConnectorLogLevelEnum and assigns it to the Severity field.
func (o *ConnectorSyncLog) SetSeverity(v ConnectorLogLevelEnum) {
	o.Severity = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ConnectorSyncLog) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorSyncLog) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ConnectorSyncLog) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ConnectorSyncLog) SetMessage(v string) {
	o.Message = &v
}

func (o ConnectorSyncLog) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorSyncLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableConnectorSyncLog struct {
	value *ConnectorSyncLog
	isSet bool
}

func (v NullableConnectorSyncLog) Get() *ConnectorSyncLog {
	return v.value
}

func (v *NullableConnectorSyncLog) Set(val *ConnectorSyncLog) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorSyncLog) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorSyncLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorSyncLog(val *ConnectorSyncLog) *NullableConnectorSyncLog {
	return &NullableConnectorSyncLog{value: val, isSet: true}
}

func (v NullableConnectorSyncLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorSyncLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
