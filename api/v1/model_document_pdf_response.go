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

// DocumentPDFResponse link to generated document
type DocumentPDFResponse struct {
	// The URL of the generated document
	DocumentUrl *string `json:"document_url,omitempty"`
}

// NewDocumentPDFResponse instantiates a new DocumentPDFResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentPDFResponse() *DocumentPDFResponse {
	this := DocumentPDFResponse{}
	return &this
}

// NewDocumentPDFResponseWithDefaults instantiates a new DocumentPDFResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentPDFResponseWithDefaults() *DocumentPDFResponse {
	this := DocumentPDFResponse{}
	return &this
}

// GetDocumentUrl returns the DocumentUrl field value if set, zero value otherwise.
func (o *DocumentPDFResponse) GetDocumentUrl() string {
	if o == nil || o.DocumentUrl == nil {
		var ret string
		return ret
	}
	return *o.DocumentUrl
}

// GetDocumentUrlOk returns a tuple with the DocumentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentPDFResponse) GetDocumentUrlOk() (*string, bool) {
	if o == nil || o.DocumentUrl == nil {
		return nil, false
	}
	return o.DocumentUrl, true
}

// HasDocumentUrl returns a boolean if a field has been set.
func (o *DocumentPDFResponse) HasDocumentUrl() bool {
	if o != nil && o.DocumentUrl != nil {
		return true
	}

	return false
}

// SetDocumentUrl gets a reference to the given string and assigns it to the DocumentUrl field.
func (o *DocumentPDFResponse) SetDocumentUrl(v string) {
	o.DocumentUrl = &v
}

func (o DocumentPDFResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentUrl != nil {
		toSerialize["document_url"] = o.DocumentUrl
	}
	return json.Marshal(toSerialize)
}

type NullableDocumentPDFResponse struct {
	value *DocumentPDFResponse
	isSet bool
}

func (v NullableDocumentPDFResponse) Get() *DocumentPDFResponse {
	return v.value
}

func (v *NullableDocumentPDFResponse) Set(val *DocumentPDFResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentPDFResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentPDFResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentPDFResponse(val *DocumentPDFResponse) *NullableDocumentPDFResponse {
	return &NullableDocumentPDFResponse{value: val, isSet: true}
}

func (v NullableDocumentPDFResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentPDFResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
