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

// UpdateConnectorConfigRequest Request used to update a connector config
type UpdateConnectorConfigRequest struct {
	// short description of the connector
	Description *string `json:"description,omitempty"`
	// indicates whether the connector should push data or not. While disabled the connector will not attempt to push data to the backend, however, it  will continue to ping the backend to check for updates to the config.
	Enabled *bool `json:"enabled,omitempty"`
	// cron expression that tells the connector how often to sync data
	SyncTimes *string `json:"sync_times,omitempty"`
	// cron expression that tells the connector how often to ping the server and  retrieve the config
	PingTimes *string `json:"ping_times,omitempty"`
	// list of entities to be synchronised
	Entities []string               `json:"entities,omitempty"`
	LogLevel *ConnectorLogLevelEnum `json:"log_level,omitempty"`
}

// NewUpdateConnectorConfigRequest instantiates a new UpdateConnectorConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateConnectorConfigRequest() *UpdateConnectorConfigRequest {
	this := UpdateConnectorConfigRequest{}
	return &this
}

// NewUpdateConnectorConfigRequestWithDefaults instantiates a new UpdateConnectorConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateConnectorConfigRequestWithDefaults() *UpdateConnectorConfigRequest {
	this := UpdateConnectorConfigRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateConnectorConfigRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectorConfigRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateConnectorConfigRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateConnectorConfigRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateConnectorConfigRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectorConfigRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateConnectorConfigRequest) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateConnectorConfigRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetSyncTimes returns the SyncTimes field value if set, zero value otherwise.
func (o *UpdateConnectorConfigRequest) GetSyncTimes() string {
	if o == nil || o.SyncTimes == nil {
		var ret string
		return ret
	}
	return *o.SyncTimes
}

// GetSyncTimesOk returns a tuple with the SyncTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectorConfigRequest) GetSyncTimesOk() (*string, bool) {
	if o == nil || o.SyncTimes == nil {
		return nil, false
	}
	return o.SyncTimes, true
}

// HasSyncTimes returns a boolean if a field has been set.
func (o *UpdateConnectorConfigRequest) HasSyncTimes() bool {
	if o != nil && o.SyncTimes != nil {
		return true
	}

	return false
}

// SetSyncTimes gets a reference to the given string and assigns it to the SyncTimes field.
func (o *UpdateConnectorConfigRequest) SetSyncTimes(v string) {
	o.SyncTimes = &v
}

// GetPingTimes returns the PingTimes field value if set, zero value otherwise.
func (o *UpdateConnectorConfigRequest) GetPingTimes() string {
	if o == nil || o.PingTimes == nil {
		var ret string
		return ret
	}
	return *o.PingTimes
}

// GetPingTimesOk returns a tuple with the PingTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectorConfigRequest) GetPingTimesOk() (*string, bool) {
	if o == nil || o.PingTimes == nil {
		return nil, false
	}
	return o.PingTimes, true
}

// HasPingTimes returns a boolean if a field has been set.
func (o *UpdateConnectorConfigRequest) HasPingTimes() bool {
	if o != nil && o.PingTimes != nil {
		return true
	}

	return false
}

// SetPingTimes gets a reference to the given string and assigns it to the PingTimes field.
func (o *UpdateConnectorConfigRequest) SetPingTimes(v string) {
	o.PingTimes = &v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *UpdateConnectorConfigRequest) GetEntities() []string {
	if o == nil || o.Entities == nil {
		var ret []string
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectorConfigRequest) GetEntitiesOk() ([]string, bool) {
	if o == nil || o.Entities == nil {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *UpdateConnectorConfigRequest) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []string and assigns it to the Entities field.
func (o *UpdateConnectorConfigRequest) SetEntities(v []string) {
	o.Entities = v
}

// GetLogLevel returns the LogLevel field value if set, zero value otherwise.
func (o *UpdateConnectorConfigRequest) GetLogLevel() ConnectorLogLevelEnum {
	if o == nil || o.LogLevel == nil {
		var ret ConnectorLogLevelEnum
		return ret
	}
	return *o.LogLevel
}

// GetLogLevelOk returns a tuple with the LogLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateConnectorConfigRequest) GetLogLevelOk() (*ConnectorLogLevelEnum, bool) {
	if o == nil || o.LogLevel == nil {
		return nil, false
	}
	return o.LogLevel, true
}

// HasLogLevel returns a boolean if a field has been set.
func (o *UpdateConnectorConfigRequest) HasLogLevel() bool {
	if o != nil && o.LogLevel != nil {
		return true
	}

	return false
}

// SetLogLevel gets a reference to the given ConnectorLogLevelEnum and assigns it to the LogLevel field.
func (o *UpdateConnectorConfigRequest) SetLogLevel(v ConnectorLogLevelEnum) {
	o.LogLevel = &v
}

func (o UpdateConnectorConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.SyncTimes != nil {
		toSerialize["sync_times"] = o.SyncTimes
	}
	if o.PingTimes != nil {
		toSerialize["ping_times"] = o.PingTimes
	}
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}
	if o.LogLevel != nil {
		toSerialize["log_level"] = o.LogLevel
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateConnectorConfigRequest struct {
	value *UpdateConnectorConfigRequest
	isSet bool
}

func (v NullableUpdateConnectorConfigRequest) Get() *UpdateConnectorConfigRequest {
	return v.value
}

func (v *NullableUpdateConnectorConfigRequest) Set(val *UpdateConnectorConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateConnectorConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateConnectorConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateConnectorConfigRequest(val *UpdateConnectorConfigRequest) *NullableUpdateConnectorConfigRequest {
	return &NullableUpdateConnectorConfigRequest{value: val, isSet: true}
}

func (v NullableUpdateConnectorConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateConnectorConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
