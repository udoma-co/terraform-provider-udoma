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

// ConnectorEntitySyncInfo Used in connector configurations to indicate what data to be synchronised
type ConnectorEntitySyncInfo struct {
	Type *ConnectorEntityType `json:"type,omitempty"`
	// total count of the entities that are currently available in the Udoma database,  that were synced by the connector.
	TotalCount *int64 `json:"total_count,omitempty"`
	// optional timestamp of the last update. If the timestamp is not provided, the connector will sync all data for the entity
	LastUpdate *int64 `json:"last_update,omitempty"`
	// a progress indicator that can be used by the connector to know how far it is in the sync process. The backend will not use this value, but it will be returned in the response to the connector.
	Progress *string `json:"progress,omitempty"`
}

// NewConnectorEntitySyncInfo instantiates a new ConnectorEntitySyncInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorEntitySyncInfo() *ConnectorEntitySyncInfo {
	this := ConnectorEntitySyncInfo{}
	return &this
}

// NewConnectorEntitySyncInfoWithDefaults instantiates a new ConnectorEntitySyncInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorEntitySyncInfoWithDefaults() *ConnectorEntitySyncInfo {
	this := ConnectorEntitySyncInfo{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConnectorEntitySyncInfo) GetType() ConnectorEntityType {
	if o == nil || o.Type == nil {
		var ret ConnectorEntityType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorEntitySyncInfo) GetTypeOk() (*ConnectorEntityType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConnectorEntitySyncInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given ConnectorEntityType and assigns it to the Type field.
func (o *ConnectorEntitySyncInfo) SetType(v ConnectorEntityType) {
	o.Type = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ConnectorEntitySyncInfo) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorEntitySyncInfo) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ConnectorEntitySyncInfo) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *ConnectorEntitySyncInfo) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise.
func (o *ConnectorEntitySyncInfo) GetLastUpdate() int64 {
	if o == nil || o.LastUpdate == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdate
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorEntitySyncInfo) GetLastUpdateOk() (*int64, bool) {
	if o == nil || o.LastUpdate == nil {
		return nil, false
	}
	return o.LastUpdate, true
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *ConnectorEntitySyncInfo) HasLastUpdate() bool {
	if o != nil && o.LastUpdate != nil {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given int64 and assigns it to the LastUpdate field.
func (o *ConnectorEntitySyncInfo) SetLastUpdate(v int64) {
	o.LastUpdate = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *ConnectorEntitySyncInfo) GetProgress() string {
	if o == nil || o.Progress == nil {
		var ret string
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorEntitySyncInfo) GetProgressOk() (*string, bool) {
	if o == nil || o.Progress == nil {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *ConnectorEntitySyncInfo) HasProgress() bool {
	if o != nil && o.Progress != nil {
		return true
	}

	return false
}

// SetProgress gets a reference to the given string and assigns it to the Progress field.
func (o *ConnectorEntitySyncInfo) SetProgress(v string) {
	o.Progress = &v
}

func (o ConnectorEntitySyncInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}
	if o.LastUpdate != nil {
		toSerialize["last_update"] = o.LastUpdate
	}
	if o.Progress != nil {
		toSerialize["progress"] = o.Progress
	}
	return json.Marshal(toSerialize)
}

type NullableConnectorEntitySyncInfo struct {
	value *ConnectorEntitySyncInfo
	isSet bool
}

func (v NullableConnectorEntitySyncInfo) Get() *ConnectorEntitySyncInfo {
	return v.value
}

func (v *NullableConnectorEntitySyncInfo) Set(val *ConnectorEntitySyncInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorEntitySyncInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorEntitySyncInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorEntitySyncInfo(val *ConnectorEntitySyncInfo) *NullableConnectorEntitySyncInfo {
	return &NullableConnectorEntitySyncInfo{value: val, isSet: true}
}

func (v NullableConnectorEntitySyncInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorEntitySyncInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
