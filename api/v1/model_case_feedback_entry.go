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

// CaseFeedbackEntry Contains the feedback that was provided by a party when changing the status of a case.
type CaseFeedbackEntry struct {
	// The unique ID of the feedback entry
	Id           *string         `json:"id,omitempty"`
	CreatedAt    *int64          `json:"created_at,omitempty"`
	AuthorRef    *UserReference  `json:"author_ref,omitempty"`
	SourceStatus *CaseStatusEnum `json:"source_status,omitempty"`
	Action       *CaseActionEnum `json:"action,omitempty"`
	// The ID of the feedback entry, used to map responses against the config
	FeedbackId *string `json:"feedback_id,omitempty"`
	// List of paries that should have access to the comment
	Visibility []CaseParty `json:"visibility,omitempty"`
	// Form data that was provided by the party executing the action.
	FormData map[string]interface{} `json:"form_data,omitempty"`
}

// NewCaseFeedbackEntry instantiates a new CaseFeedbackEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaseFeedbackEntry() *CaseFeedbackEntry {
	this := CaseFeedbackEntry{}
	return &this
}

// NewCaseFeedbackEntryWithDefaults instantiates a new CaseFeedbackEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaseFeedbackEntryWithDefaults() *CaseFeedbackEntry {
	this := CaseFeedbackEntry{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CaseFeedbackEntry) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseFeedbackEntry) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CaseFeedbackEntry) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CaseFeedbackEntry) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CaseFeedbackEntry) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseFeedbackEntry) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CaseFeedbackEntry) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *CaseFeedbackEntry) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetAuthorRef returns the AuthorRef field value if set, zero value otherwise.
func (o *CaseFeedbackEntry) GetAuthorRef() UserReference {
	if o == nil || o.AuthorRef == nil {
		var ret UserReference
		return ret
	}
	return *o.AuthorRef
}

// GetAuthorRefOk returns a tuple with the AuthorRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseFeedbackEntry) GetAuthorRefOk() (*UserReference, bool) {
	if o == nil || o.AuthorRef == nil {
		return nil, false
	}
	return o.AuthorRef, true
}

// HasAuthorRef returns a boolean if a field has been set.
func (o *CaseFeedbackEntry) HasAuthorRef() bool {
	if o != nil && o.AuthorRef != nil {
		return true
	}

	return false
}

// SetAuthorRef gets a reference to the given UserReference and assigns it to the AuthorRef field.
func (o *CaseFeedbackEntry) SetAuthorRef(v UserReference) {
	o.AuthorRef = &v
}

// GetSourceStatus returns the SourceStatus field value if set, zero value otherwise.
func (o *CaseFeedbackEntry) GetSourceStatus() CaseStatusEnum {
	if o == nil || o.SourceStatus == nil {
		var ret CaseStatusEnum
		return ret
	}
	return *o.SourceStatus
}

// GetSourceStatusOk returns a tuple with the SourceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseFeedbackEntry) GetSourceStatusOk() (*CaseStatusEnum, bool) {
	if o == nil || o.SourceStatus == nil {
		return nil, false
	}
	return o.SourceStatus, true
}

// HasSourceStatus returns a boolean if a field has been set.
func (o *CaseFeedbackEntry) HasSourceStatus() bool {
	if o != nil && o.SourceStatus != nil {
		return true
	}

	return false
}

// SetSourceStatus gets a reference to the given CaseStatusEnum and assigns it to the SourceStatus field.
func (o *CaseFeedbackEntry) SetSourceStatus(v CaseStatusEnum) {
	o.SourceStatus = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *CaseFeedbackEntry) GetAction() CaseActionEnum {
	if o == nil || o.Action == nil {
		var ret CaseActionEnum
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseFeedbackEntry) GetActionOk() (*CaseActionEnum, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *CaseFeedbackEntry) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given CaseActionEnum and assigns it to the Action field.
func (o *CaseFeedbackEntry) SetAction(v CaseActionEnum) {
	o.Action = &v
}

// GetFeedbackId returns the FeedbackId field value if set, zero value otherwise.
func (o *CaseFeedbackEntry) GetFeedbackId() string {
	if o == nil || o.FeedbackId == nil {
		var ret string
		return ret
	}
	return *o.FeedbackId
}

// GetFeedbackIdOk returns a tuple with the FeedbackId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseFeedbackEntry) GetFeedbackIdOk() (*string, bool) {
	if o == nil || o.FeedbackId == nil {
		return nil, false
	}
	return o.FeedbackId, true
}

// HasFeedbackId returns a boolean if a field has been set.
func (o *CaseFeedbackEntry) HasFeedbackId() bool {
	if o != nil && o.FeedbackId != nil {
		return true
	}

	return false
}

// SetFeedbackId gets a reference to the given string and assigns it to the FeedbackId field.
func (o *CaseFeedbackEntry) SetFeedbackId(v string) {
	o.FeedbackId = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *CaseFeedbackEntry) GetVisibility() []CaseParty {
	if o == nil || o.Visibility == nil {
		var ret []CaseParty
		return ret
	}
	return o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseFeedbackEntry) GetVisibilityOk() ([]CaseParty, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *CaseFeedbackEntry) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given []CaseParty and assigns it to the Visibility field.
func (o *CaseFeedbackEntry) SetVisibility(v []CaseParty) {
	o.Visibility = v
}

// GetFormData returns the FormData field value if set, zero value otherwise.
func (o *CaseFeedbackEntry) GetFormData() map[string]interface{} {
	if o == nil || o.FormData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.FormData
}

// GetFormDataOk returns a tuple with the FormData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseFeedbackEntry) GetFormDataOk() (map[string]interface{}, bool) {
	if o == nil || o.FormData == nil {
		return nil, false
	}
	return o.FormData, true
}

// HasFormData returns a boolean if a field has been set.
func (o *CaseFeedbackEntry) HasFormData() bool {
	if o != nil && o.FormData != nil {
		return true
	}

	return false
}

// SetFormData gets a reference to the given map[string]interface{} and assigns it to the FormData field.
func (o *CaseFeedbackEntry) SetFormData(v map[string]interface{}) {
	o.FormData = v
}

func (o CaseFeedbackEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.AuthorRef != nil {
		toSerialize["author_ref"] = o.AuthorRef
	}
	if o.SourceStatus != nil {
		toSerialize["source_status"] = o.SourceStatus
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.FeedbackId != nil {
		toSerialize["feedback_id"] = o.FeedbackId
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	if o.FormData != nil {
		toSerialize["form_data"] = o.FormData
	}
	return json.Marshal(toSerialize)
}

type NullableCaseFeedbackEntry struct {
	value *CaseFeedbackEntry
	isSet bool
}

func (v NullableCaseFeedbackEntry) Get() *CaseFeedbackEntry {
	return v.value
}

func (v *NullableCaseFeedbackEntry) Set(val *CaseFeedbackEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableCaseFeedbackEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableCaseFeedbackEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaseFeedbackEntry(val *CaseFeedbackEntry) *NullableCaseFeedbackEntry {
	return &NullableCaseFeedbackEntry{value: val, isSet: true}
}

func (v NullableCaseFeedbackEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaseFeedbackEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
