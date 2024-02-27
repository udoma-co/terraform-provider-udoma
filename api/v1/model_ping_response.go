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

// checks if the PingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PingResponse{}

// PingResponse Response used for connector pings
type PingResponse struct {
	Config   *ConnectorConfig          `json:"config,omitempty"`
	Entities []ConnectorEntitySyncInfo `json:"entities,omitempty"`
}

// NewPingResponse instantiates a new PingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPingResponse() *PingResponse {
	this := PingResponse{}
	return &this
}

// NewPingResponseWithDefaults instantiates a new PingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPingResponseWithDefaults() *PingResponse {
	this := PingResponse{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *PingResponse) GetConfig() ConnectorConfig {
	if o == nil || IsNil(o.Config) {
		var ret ConnectorConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingResponse) GetConfigOk() (*ConnectorConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *PingResponse) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given ConnectorConfig and assigns it to the Config field.
func (o *PingResponse) SetConfig(v ConnectorConfig) {
	o.Config = &v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *PingResponse) GetEntities() []ConnectorEntitySyncInfo {
	if o == nil || IsNil(o.Entities) {
		var ret []ConnectorEntitySyncInfo
		return ret
	}
	return o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PingResponse) GetEntitiesOk() ([]ConnectorEntitySyncInfo, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *PingResponse) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given []ConnectorEntitySyncInfo and assigns it to the Entities field.
func (o *PingResponse) SetEntities(v []ConnectorEntitySyncInfo) {
	o.Entities = v
}

func (o PingResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Entities) {
		toSerialize["entities"] = o.Entities
	}
	return toSerialize, nil
}

type NullablePingResponse struct {
	value *PingResponse
	isSet bool
}

func (v NullablePingResponse) Get() *PingResponse {
	return v.value
}

func (v *NullablePingResponse) Set(val *PingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePingResponse(val *PingResponse) *NullablePingResponse {
	return &NullablePingResponse{value: val, isSet: true}
}

func (v NullablePingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
