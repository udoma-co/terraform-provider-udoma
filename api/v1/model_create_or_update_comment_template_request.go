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

// checks if the CreateOrUpdateCommentTemplateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateCommentTemplateRequest{}

// CreateOrUpdateCommentTemplateRequest struct for CreateOrUpdateCommentTemplateRequest
type CreateOrUpdateCommentTemplateRequest struct {
	DisplayName string `json:"display_name"`
	// A list of case template IDs that either gives or restricts access in cetain cases, based on is_deny_list property.
	AccessList []string `json:"access_list,omitempty"`
	// If true - access_list behaves as deny list, if false it behaves as allow list
	IsDenyList *bool `json:"is_deny_list,omitempty"`
	// An optional script to process the reporter if needed. The script will have reporter under `reporter` and should return an object to be used in the template.
	Script *string `json:"script,omitempty"`
	// A golang template, that'll be passed reporter as Data and should return a nicely formatted string to use as a comment.
	Template string `json:"template"`
}

type _CreateOrUpdateCommentTemplateRequest CreateOrUpdateCommentTemplateRequest

// NewCreateOrUpdateCommentTemplateRequest instantiates a new CreateOrUpdateCommentTemplateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateCommentTemplateRequest(displayName string, template string) *CreateOrUpdateCommentTemplateRequest {
	this := CreateOrUpdateCommentTemplateRequest{}
	this.DisplayName = displayName
	this.Template = template
	return &this
}

// NewCreateOrUpdateCommentTemplateRequestWithDefaults instantiates a new CreateOrUpdateCommentTemplateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateCommentTemplateRequestWithDefaults() *CreateOrUpdateCommentTemplateRequest {
	this := CreateOrUpdateCommentTemplateRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value
func (o *CreateOrUpdateCommentTemplateRequest) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateCommentTemplateRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CreateOrUpdateCommentTemplateRequest) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetAccessList returns the AccessList field value if set, zero value otherwise.
func (o *CreateOrUpdateCommentTemplateRequest) GetAccessList() []string {
	if o == nil || IsNil(o.AccessList) {
		var ret []string
		return ret
	}
	return o.AccessList
}

// GetAccessListOk returns a tuple with the AccessList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateCommentTemplateRequest) GetAccessListOk() ([]string, bool) {
	if o == nil || IsNil(o.AccessList) {
		return nil, false
	}
	return o.AccessList, true
}

// HasAccessList returns a boolean if a field has been set.
func (o *CreateOrUpdateCommentTemplateRequest) HasAccessList() bool {
	if o != nil && !IsNil(o.AccessList) {
		return true
	}

	return false
}

// SetAccessList gets a reference to the given []string and assigns it to the AccessList field.
func (o *CreateOrUpdateCommentTemplateRequest) SetAccessList(v []string) {
	o.AccessList = v
}

// GetIsDenyList returns the IsDenyList field value if set, zero value otherwise.
func (o *CreateOrUpdateCommentTemplateRequest) GetIsDenyList() bool {
	if o == nil || IsNil(o.IsDenyList) {
		var ret bool
		return ret
	}
	return *o.IsDenyList
}

// GetIsDenyListOk returns a tuple with the IsDenyList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateCommentTemplateRequest) GetIsDenyListOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDenyList) {
		return nil, false
	}
	return o.IsDenyList, true
}

// HasIsDenyList returns a boolean if a field has been set.
func (o *CreateOrUpdateCommentTemplateRequest) HasIsDenyList() bool {
	if o != nil && !IsNil(o.IsDenyList) {
		return true
	}

	return false
}

// SetIsDenyList gets a reference to the given bool and assigns it to the IsDenyList field.
func (o *CreateOrUpdateCommentTemplateRequest) SetIsDenyList(v bool) {
	o.IsDenyList = &v
}

// GetScript returns the Script field value if set, zero value otherwise.
func (o *CreateOrUpdateCommentTemplateRequest) GetScript() string {
	if o == nil || IsNil(o.Script) {
		var ret string
		return ret
	}
	return *o.Script
}

// GetScriptOk returns a tuple with the Script field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateCommentTemplateRequest) GetScriptOk() (*string, bool) {
	if o == nil || IsNil(o.Script) {
		return nil, false
	}
	return o.Script, true
}

// HasScript returns a boolean if a field has been set.
func (o *CreateOrUpdateCommentTemplateRequest) HasScript() bool {
	if o != nil && !IsNil(o.Script) {
		return true
	}

	return false
}

// SetScript gets a reference to the given string and assigns it to the Script field.
func (o *CreateOrUpdateCommentTemplateRequest) SetScript(v string) {
	o.Script = &v
}

// GetTemplate returns the Template field value
func (o *CreateOrUpdateCommentTemplateRequest) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateCommentTemplateRequest) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *CreateOrUpdateCommentTemplateRequest) SetTemplate(v string) {
	o.Template = v
}

func (o CreateOrUpdateCommentTemplateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateCommentTemplateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *CreateOrUpdateCommentTemplateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varCreateOrUpdateCommentTemplateRequest := _CreateOrUpdateCommentTemplateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrUpdateCommentTemplateRequest)

	if err != nil {
		return err
	}

	*o = CreateOrUpdateCommentTemplateRequest(varCreateOrUpdateCommentTemplateRequest)

	return err
}

type NullableCreateOrUpdateCommentTemplateRequest struct {
	value *CreateOrUpdateCommentTemplateRequest
	isSet bool
}

func (v NullableCreateOrUpdateCommentTemplateRequest) Get() *CreateOrUpdateCommentTemplateRequest {
	return v.value
}

func (v *NullableCreateOrUpdateCommentTemplateRequest) Set(val *CreateOrUpdateCommentTemplateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateCommentTemplateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateCommentTemplateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateCommentTemplateRequest(val *CreateOrUpdateCommentTemplateRequest) *NullableCreateOrUpdateCommentTemplateRequest {
	return &NullableCreateOrUpdateCommentTemplateRequest{value: val, isSet: true}
}

func (v NullableCreateOrUpdateCommentTemplateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateCommentTemplateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}