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

// checks if the QueryNotificationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryNotificationsResponse{}

// QueryNotificationsResponse A list of all notifications that are requested by the user
type QueryNotificationsResponse struct {
	// The maximum number of entities to return from the query
	Limit *int32 `json:"limit,omitempty"`
	// The number of entities to skip before returning the result
	Offset        *int32         `json:"offset,omitempty"`
	Total         *int32         `json:"total,omitempty"`
	Notifications []Notification `json:"notifications,omitempty"`
}

// NewQueryNotificationsResponse instantiates a new QueryNotificationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryNotificationsResponse() *QueryNotificationsResponse {
	this := QueryNotificationsResponse{}
	return &this
}

// NewQueryNotificationsResponseWithDefaults instantiates a new QueryNotificationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryNotificationsResponseWithDefaults() *QueryNotificationsResponse {
	this := QueryNotificationsResponse{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *QueryNotificationsResponse) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryNotificationsResponse) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *QueryNotificationsResponse) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *QueryNotificationsResponse) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *QueryNotificationsResponse) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryNotificationsResponse) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *QueryNotificationsResponse) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *QueryNotificationsResponse) SetOffset(v int32) {
	o.Offset = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *QueryNotificationsResponse) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryNotificationsResponse) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *QueryNotificationsResponse) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *QueryNotificationsResponse) SetTotal(v int32) {
	o.Total = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *QueryNotificationsResponse) GetNotifications() []Notification {
	if o == nil || IsNil(o.Notifications) {
		var ret []Notification
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryNotificationsResponse) GetNotificationsOk() ([]Notification, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *QueryNotificationsResponse) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []Notification and assigns it to the Notifications field.
func (o *QueryNotificationsResponse) SetNotifications(v []Notification) {
	o.Notifications = v
}

func (o QueryNotificationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryNotificationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	return toSerialize, nil
}

type NullableQueryNotificationsResponse struct {
	value *QueryNotificationsResponse
	isSet bool
}

func (v NullableQueryNotificationsResponse) Get() *QueryNotificationsResponse {
	return v.value
}

func (v *NullableQueryNotificationsResponse) Set(val *QueryNotificationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryNotificationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryNotificationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryNotificationsResponse(val *QueryNotificationsResponse) *NullableQueryNotificationsResponse {
	return &NullableQueryNotificationsResponse{value: val, isSet: true}
}

func (v NullableQueryNotificationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryNotificationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
