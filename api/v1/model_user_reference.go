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

// checks if the UserReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserReference{}

// UserReference Contains information about a user referenced in other entities
type UserReference struct {
	// The ID of the user, can be undefined, if user is not registered, or was already deleted
	UserId      *string       `json:"user_id,omitempty"`
	UserRole    *UserTypeEnum `json:"user_role,omitempty"`
	ContactData *ContactData  `json:"contact_data,omitempty"`
}

// NewUserReference instantiates a new UserReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserReference() *UserReference {
	this := UserReference{}
	return &this
}

// NewUserReferenceWithDefaults instantiates a new UserReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserReferenceWithDefaults() *UserReference {
	this := UserReference{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UserReference) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserReference) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UserReference) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *UserReference) SetUserId(v string) {
	o.UserId = &v
}

// GetUserRole returns the UserRole field value if set, zero value otherwise.
func (o *UserReference) GetUserRole() UserTypeEnum {
	if o == nil || IsNil(o.UserRole) {
		var ret UserTypeEnum
		return ret
	}
	return *o.UserRole
}

// GetUserRoleOk returns a tuple with the UserRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserReference) GetUserRoleOk() (*UserTypeEnum, bool) {
	if o == nil || IsNil(o.UserRole) {
		return nil, false
	}
	return o.UserRole, true
}

// HasUserRole returns a boolean if a field has been set.
func (o *UserReference) HasUserRole() bool {
	if o != nil && !IsNil(o.UserRole) {
		return true
	}

	return false
}

// SetUserRole gets a reference to the given UserTypeEnum and assigns it to the UserRole field.
func (o *UserReference) SetUserRole(v UserTypeEnum) {
	o.UserRole = &v
}

// GetContactData returns the ContactData field value if set, zero value otherwise.
func (o *UserReference) GetContactData() ContactData {
	if o == nil || IsNil(o.ContactData) {
		var ret ContactData
		return ret
	}
	return *o.ContactData
}

// GetContactDataOk returns a tuple with the ContactData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserReference) GetContactDataOk() (*ContactData, bool) {
	if o == nil || IsNil(o.ContactData) {
		return nil, false
	}
	return o.ContactData, true
}

// HasContactData returns a boolean if a field has been set.
func (o *UserReference) HasContactData() bool {
	if o != nil && !IsNil(o.ContactData) {
		return true
	}

	return false
}

// SetContactData gets a reference to the given ContactData and assigns it to the ContactData field.
func (o *UserReference) SetContactData(v ContactData) {
	o.ContactData = &v
}

func (o UserReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.UserRole) {
		toSerialize["user_role"] = o.UserRole
	}
	if !IsNil(o.ContactData) {
		toSerialize["contact_data"] = o.ContactData
	}
	return toSerialize, nil
}

type NullableUserReference struct {
	value *UserReference
	isSet bool
}

func (v NullableUserReference) Get() *UserReference {
	return v.value
}

func (v *NullableUserReference) Set(val *UserReference) {
	v.value = val
	v.isSet = true
}

func (v NullableUserReference) IsSet() bool {
	return v.isSet
}

func (v *NullableUserReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserReference(val *UserReference) *NullableUserReference {
	return &NullableUserReference{value: val, isSet: true}
}

func (v NullableUserReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
