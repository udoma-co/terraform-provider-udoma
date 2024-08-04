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

// checks if the CommentTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommentTemplate{}

// CommentTemplate struct for CommentTemplate
type CommentTemplate struct {
	// Unique and immutable ID attribute of the entity that is generated when  the instance is created. The ID is unique within the system accross all accounts and it can be used to reference the entity in other entities  or to retrieve it from the backend.
	Id string `json:"id"`
	// The date and time the entity was created
	CreatedAt int64 `json:"created_at"`
	// The date and time the entity was last updated
	UpdatedAt   int64  `json:"updated_at"`
	DisplayName string `json:"display_name"`
	// A list of case template IDs that either gives or restricts access in cetain cases, based on is_deny_list property.
	AccessList []string `json:"access_list,omitempty"`
	// If true - access list behaves as deny list, if false it behaves as allow list
	IsDenyList *bool `json:"is_deny_list,omitempty"`
	// An optional script to process the reporter if needed. The script will have reporter under `reporter` and should return an object to be used in the template.
	Script *string `json:"script,omitempty"`
	// A golang template, that'll be passed reporter as Data and should return a nicely formatted string to use as a comment.
	Template string `json:"template"`
}

type _CommentTemplate CommentTemplate

// NewCommentTemplate instantiates a new CommentTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommentTemplate(id string, createdAt int64, updatedAt int64, displayName string, template string) *CommentTemplate {
	this := CommentTemplate{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.DisplayName = displayName
	this.Template = template
	return &this
}

// NewCommentTemplateWithDefaults instantiates a new CommentTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentTemplateWithDefaults() *CommentTemplate {
	this := CommentTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *CommentTemplate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CommentTemplate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CommentTemplate) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CommentTemplate) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CommentTemplate) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CommentTemplate) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *CommentTemplate) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *CommentTemplate) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *CommentTemplate) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetDisplayName returns the DisplayName field value
func (o *CommentTemplate) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CommentTemplate) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CommentTemplate) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetAccessList returns the AccessList field value if set, zero value otherwise.
func (o *CommentTemplate) GetAccessList() []string {
	if o == nil || IsNil(o.AccessList) {
		var ret []string
		return ret
	}
	return o.AccessList
}

// GetAccessListOk returns a tuple with the AccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentTemplate) GetAccessListOk() ([]string, bool) {
	if o == nil || IsNil(o.AccessList) {
		return nil, false
	}
	return o.AccessList, true
}

// HasAccessList returns a boolean if a field has been set.
func (o *CommentTemplate) HasAccessList() bool {
	if o != nil && !IsNil(o.AccessList) {
		return true
	}

	return false
}

// SetAccessList gets a reference to the given []string and assigns it to the AccessList field.
func (o *CommentTemplate) SetAccessList(v []string) {
	o.AccessList = v
}

// GetIsDenyList returns the IsDenyList field value if set, zero value otherwise.
func (o *CommentTemplate) GetIsDenyList() bool {
	if o == nil || IsNil(o.IsDenyList) {
		var ret bool
		return ret
	}
	return *o.IsDenyList
}

// GetIsDenyListOk returns a tuple with the IsDenyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentTemplate) GetIsDenyListOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDenyList) {
		return nil, false
	}
	return o.IsDenyList, true
}

// HasIsDenyList returns a boolean if a field has been set.
func (o *CommentTemplate) HasIsDenyList() bool {
	if o != nil && !IsNil(o.IsDenyList) {
		return true
	}

	return false
}

// SetIsDenyList gets a reference to the given bool and assigns it to the IsDenyList field.
func (o *CommentTemplate) SetIsDenyList(v bool) {
	o.IsDenyList = &v
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *CommentTemplate) GetScript() string {
	if o == nil || IsNil(o.Script) {
		var ret string
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommentTemplate) GetScriptOk() (*string, bool) {
	if o == nil || IsNil(o.Script) {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *CommentTemplate) HasScript() bool {
	if o != nil && !IsNil(o.Script) {
		return true
	}

	return false
}

// SetScript gets a reference to the given string and assigns it to the Script field.
func (o *CommentTemplate) SetScript(v string) {
	o.Script = &v
}

// GetTemplate returns the Template field value
func (o *CommentTemplate) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *CommentTemplate) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *CommentTemplate) SetTemplate(v string) {
	o.Template = v
}

func (o CommentTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommentTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["display_name"] = o.DisplayName
	if !IsNil(o.AccessList) {
		toSerialize["access_list"] = o.AccessList
	}
	if !IsNil(o.IsDenyList) {
		toSerialize["is_deny_list"] = o.IsDenyList
	}
	if !IsNil(o.Script) {
		toSerialize["script"] = o.Script
	}
	toSerialize["template"] = o.Template
	return toSerialize, nil
}

func (o *CommentTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"display_name",
		"template",
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

	varCommentTemplate := _CommentTemplate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommentTemplate)

	if err != nil {
		return err
	}

	*o = CommentTemplate(varCommentTemplate)

	return err
}

type NullableCommentTemplate struct {
	value *CommentTemplate
	isSet bool
}

func (v NullableCommentTemplate) Get() *CommentTemplate {
	return v.value
}

func (v *NullableCommentTemplate) Set(val *CommentTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableCommentTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableCommentTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommentTemplate(val *CommentTemplate) *NullableCommentTemplate {
	return &NullableCommentTemplate{value: val, isSet: true}
}

func (v NullableCommentTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommentTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
