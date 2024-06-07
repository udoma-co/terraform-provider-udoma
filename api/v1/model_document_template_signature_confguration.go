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

// checks if the DocumentTemplateSignatureConfguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentTemplateSignatureConfguration{}

// DocumentTemplateSignatureConfguration The configuration for electronic signatures
type DocumentTemplateSignatureConfguration struct {
	// The list of signers of the document
	Signers []DocumentTemplateSignerDefinition `json:"signers,omitempty"`
	// True if electronic signatures are enabled for the template
	EsignaturesEnabled *bool `json:"esignatures_enabled,omitempty"`
	// The number of days after which the signature request expires. If not specified, the default value is 7 days.
	ExpirationDays *int32 `json:"expiration_days,omitempty"`
	// A custom message to be included in the signature request. If not specified, the default message is used.
	CustomMessage *string `json:"custom_message,omitempty"`
}

// NewDocumentTemplateSignatureConfguration instantiates a new DocumentTemplateSignatureConfguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentTemplateSignatureConfguration() *DocumentTemplateSignatureConfguration {
	this := DocumentTemplateSignatureConfguration{}
	return &this
}

// NewDocumentTemplateSignatureConfgurationWithDefaults instantiates a new DocumentTemplateSignatureConfguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentTemplateSignatureConfgurationWithDefaults() *DocumentTemplateSignatureConfguration {
	this := DocumentTemplateSignatureConfguration{}
	return &this
}

// GetSigners returns the Signers field value if set, zero value otherwise.
func (o *DocumentTemplateSignatureConfguration) GetSigners() []DocumentTemplateSignerDefinition {
	if o == nil || IsNil(o.Signers) {
		var ret []DocumentTemplateSignerDefinition
		return ret
	}
	return o.Signers
}

// GetSignersOk returns a tuple with the Signers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentTemplateSignatureConfguration) GetSignersOk() ([]DocumentTemplateSignerDefinition, bool) {
	if o == nil || IsNil(o.Signers) {
		return nil, false
	}
	return o.Signers, true
}

// HasSigners returns a boolean if a field has been set.
func (o *DocumentTemplateSignatureConfguration) HasSigners() bool {
	if o != nil && !IsNil(o.Signers) {
		return true
	}

	return false
}

// SetSigners gets a reference to the given []DocumentTemplateSignerDefinition and assigns it to the Signers field.
func (o *DocumentTemplateSignatureConfguration) SetSigners(v []DocumentTemplateSignerDefinition) {
	o.Signers = v
}

// GetEsignaturesEnabled returns the EsignaturesEnabled field value if set, zero value otherwise.
func (o *DocumentTemplateSignatureConfguration) GetEsignaturesEnabled() bool {
	if o == nil || IsNil(o.EsignaturesEnabled) {
		var ret bool
		return ret
	}
	return *o.EsignaturesEnabled
}

// GetEsignaturesEnabledOk returns a tuple with the EsignaturesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentTemplateSignatureConfguration) GetEsignaturesEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EsignaturesEnabled) {
		return nil, false
	}
	return o.EsignaturesEnabled, true
}

// HasEsignaturesEnabled returns a boolean if a field has been set.
func (o *DocumentTemplateSignatureConfguration) HasEsignaturesEnabled() bool {
	if o != nil && !IsNil(o.EsignaturesEnabled) {
		return true
	}

	return false
}

// SetEsignaturesEnabled gets a reference to the given bool and assigns it to the EsignaturesEnabled field.
func (o *DocumentTemplateSignatureConfguration) SetEsignaturesEnabled(v bool) {
	o.EsignaturesEnabled = &v
}

// GetExpirationDays returns the ExpirationDays field value if set, zero value otherwise.
func (o *DocumentTemplateSignatureConfguration) GetExpirationDays() int32 {
	if o == nil || IsNil(o.ExpirationDays) {
		var ret int32
		return ret
	}
	return *o.ExpirationDays
}

// GetExpirationDaysOk returns a tuple with the ExpirationDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentTemplateSignatureConfguration) GetExpirationDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpirationDays) {
		return nil, false
	}
	return o.ExpirationDays, true
}

// HasExpirationDays returns a boolean if a field has been set.
func (o *DocumentTemplateSignatureConfguration) HasExpirationDays() bool {
	if o != nil && !IsNil(o.ExpirationDays) {
		return true
	}

	return false
}

// SetExpirationDays gets a reference to the given int32 and assigns it to the ExpirationDays field.
func (o *DocumentTemplateSignatureConfguration) SetExpirationDays(v int32) {
	o.ExpirationDays = &v
}

// GetCustomMessage returns the CustomMessage field value if set, zero value otherwise.
func (o *DocumentTemplateSignatureConfguration) GetCustomMessage() string {
	if o == nil || IsNil(o.CustomMessage) {
		var ret string
		return ret
	}
	return *o.CustomMessage
}

// GetCustomMessageOk returns a tuple with the CustomMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentTemplateSignatureConfguration) GetCustomMessageOk() (*string, bool) {
	if o == nil || IsNil(o.CustomMessage) {
		return nil, false
	}
	return o.CustomMessage, true
}

// HasCustomMessage returns a boolean if a field has been set.
func (o *DocumentTemplateSignatureConfguration) HasCustomMessage() bool {
	if o != nil && !IsNil(o.CustomMessage) {
		return true
	}

	return false
}

// SetCustomMessage gets a reference to the given string and assigns it to the CustomMessage field.
func (o *DocumentTemplateSignatureConfguration) SetCustomMessage(v string) {
	o.CustomMessage = &v
}

func (o DocumentTemplateSignatureConfguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentTemplateSignatureConfguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Signers) {
		toSerialize["signers"] = o.Signers
	}
	if !IsNil(o.EsignaturesEnabled) {
		toSerialize["esignatures_enabled"] = o.EsignaturesEnabled
	}
	if !IsNil(o.ExpirationDays) {
		toSerialize["expiration_days"] = o.ExpirationDays
	}
	if !IsNil(o.CustomMessage) {
		toSerialize["custom_message"] = o.CustomMessage
	}
	return toSerialize, nil
}

type NullableDocumentTemplateSignatureConfguration struct {
	value *DocumentTemplateSignatureConfguration
	isSet bool
}

func (v NullableDocumentTemplateSignatureConfguration) Get() *DocumentTemplateSignatureConfguration {
	return v.value
}

func (v *NullableDocumentTemplateSignatureConfguration) Set(val *DocumentTemplateSignatureConfguration) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentTemplateSignatureConfguration) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentTemplateSignatureConfguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentTemplateSignatureConfguration(val *DocumentTemplateSignatureConfguration) *NullableDocumentTemplateSignatureConfguration {
	return &NullableDocumentTemplateSignatureConfguration{value: val, isSet: true}
}

func (v NullableDocumentTemplateSignatureConfguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentTemplateSignatureConfguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
