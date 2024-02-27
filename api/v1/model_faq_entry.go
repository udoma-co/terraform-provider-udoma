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

// checks if the FAQEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FAQEntry{}

// FAQEntry struct for FAQEntry
type FAQEntry struct {
	Id *string `json:"id,omitempty"`
	// a map of values, where the key and values are strings
	Question *map[string]string `json:"question,omitempty"`
	// a map of values, where the key and values are strings
	Answer   *map[string]string `json:"answer,omitempty"`
	Keywords []string           `json:"keywords,omitempty"`
}

// NewFAQEntry instantiates a new FAQEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFAQEntry() *FAQEntry {
	this := FAQEntry{}
	return &this
}

// NewFAQEntryWithDefaults instantiates a new FAQEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFAQEntryWithDefaults() *FAQEntry {
	this := FAQEntry{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FAQEntry) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FAQEntry) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FAQEntry) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *FAQEntry) SetId(v string) {
	o.Id = &v
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *FAQEntry) GetQuestion() map[string]string {
	if o == nil || IsNil(o.Question) {
		var ret map[string]string
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FAQEntry) GetQuestionOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Question) {
		return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *FAQEntry) HasQuestion() bool {
	if o != nil && !IsNil(o.Question) {
		return true
	}

	return false
}

// SetQuestion gets a reference to the given map[string]string and assigns it to the Question field.
func (o *FAQEntry) SetQuestion(v map[string]string) {
	o.Question = &v
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *FAQEntry) GetAnswer() map[string]string {
	if o == nil || IsNil(o.Answer) {
		var ret map[string]string
		return ret
	}
	return *o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FAQEntry) GetAnswerOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Answer) {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *FAQEntry) HasAnswer() bool {
	if o != nil && !IsNil(o.Answer) {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given map[string]string and assigns it to the Answer field.
func (o *FAQEntry) SetAnswer(v map[string]string) {
	o.Answer = &v
}

// GetKeywords returns the Keywords field value if set, zero value otherwise.
func (o *FAQEntry) GetKeywords() []string {
	if o == nil || IsNil(o.Keywords) {
		var ret []string
		return ret
	}
	return o.Keywords
}

// GetKeywordsOk returns a tuple with the Keywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FAQEntry) GetKeywordsOk() ([]string, bool) {
	if o == nil || IsNil(o.Keywords) {
		return nil, false
	}
	return o.Keywords, true
}

// HasKeywords returns a boolean if a field has been set.
func (o *FAQEntry) HasKeywords() bool {
	if o != nil && !IsNil(o.Keywords) {
		return true
	}

	return false
}

// SetKeywords gets a reference to the given []string and assigns it to the Keywords field.
func (o *FAQEntry) SetKeywords(v []string) {
	o.Keywords = v
}

func (o FAQEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FAQEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Question) {
		toSerialize["question"] = o.Question
	}
	if !IsNil(o.Answer) {
		toSerialize["answer"] = o.Answer
	}
	if !IsNil(o.Keywords) {
		toSerialize["keywords"] = o.Keywords
	}
	return toSerialize, nil
}

type NullableFAQEntry struct {
	value *FAQEntry
	isSet bool
}

func (v NullableFAQEntry) Get() *FAQEntry {
	return v.value
}

func (v *NullableFAQEntry) Set(val *FAQEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableFAQEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableFAQEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFAQEntry(val *FAQEntry) *NullableFAQEntry {
	return &NullableFAQEntry{value: val, isSet: true}
}

func (v NullableFAQEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFAQEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
