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

// PropertyHistory struct for PropertyHistory
type PropertyHistory struct {
	// Reference to the property for which this history object applies
	PropertyRef *string                `json:"property_ref,omitempty"`
	Entries     []PropertyHistoryEntry `json:"entries,omitempty"`
}

// NewPropertyHistory instantiates a new PropertyHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyHistory() *PropertyHistory {
	this := PropertyHistory{}
	return &this
}

// NewPropertyHistoryWithDefaults instantiates a new PropertyHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyHistoryWithDefaults() *PropertyHistory {
	this := PropertyHistory{}
	return &this
}

// GetPropertyRef returns the PropertyRef field value if set, zero value otherwise.
func (o *PropertyHistory) GetPropertyRef() string {
	if o == nil || o.PropertyRef == nil {
		var ret string
		return ret
	}
	return *o.PropertyRef
}

// GetPropertyRefOk returns a tuple with the PropertyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyHistory) GetPropertyRefOk() (*string, bool) {
	if o == nil || o.PropertyRef == nil {
		return nil, false
	}
	return o.PropertyRef, true
}

// HasPropertyRef returns a boolean if a field has been set.
func (o *PropertyHistory) HasPropertyRef() bool {
	if o != nil && o.PropertyRef != nil {
		return true
	}

	return false
}

// SetPropertyRef gets a reference to the given string and assigns it to the PropertyRef field.
func (o *PropertyHistory) SetPropertyRef(v string) {
	o.PropertyRef = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *PropertyHistory) GetEntries() []PropertyHistoryEntry {
	if o == nil || o.Entries == nil {
		var ret []PropertyHistoryEntry
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyHistory) GetEntriesOk() ([]PropertyHistoryEntry, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *PropertyHistory) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []PropertyHistoryEntry and assigns it to the Entries field.
func (o *PropertyHistory) SetEntries(v []PropertyHistoryEntry) {
	o.Entries = v
}

func (o PropertyHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PropertyRef != nil {
		toSerialize["property_ref"] = o.PropertyRef
	}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyHistory struct {
	value *PropertyHistory
	isSet bool
}

func (v NullablePropertyHistory) Get() *PropertyHistory {
	return v.value
}

func (v *NullablePropertyHistory) Set(val *PropertyHistory) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyHistory) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyHistory(val *PropertyHistory) *NullablePropertyHistory {
	return &NullablePropertyHistory{value: val, isSet: true}
}

func (v NullablePropertyHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
