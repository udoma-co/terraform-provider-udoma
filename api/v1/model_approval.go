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

// checks if the Approval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Approval{}

// Approval Contains the details of an approval requested from a user.
type Approval struct {
	// Unique and immutable ID attribute of the entity that is generated when the instance is created. The ID is unique within the system accross all accounts and it can be used to reference the entity in other entities or to retrieve it from the backend.
	Id string `json:"id"`
	// The date and time the entity was created
	CreatedAt int64 `json:"created_at"`
	// The date and time the entity was last updated
	UpdatedAt int64 `json:"updated_at"`
	// A user-friendly name of the approval.
	Name string `json:"name"`
	// The reference to the entity for which this approval is requested.
	SourceRef string `json:"source_ref"`
	// Optional URL that can be set that will link to the entity for  which this approval is requested. If not set, the URL will be generated based on the source_ref.
	Url *string `json:"url,omitempty"`
	// The message that is displayed to the user when the approval is requested.
	Message *string `json:"message,omitempty"`
	// The list of users that are requested to approve the entity.
	Approvers []Approver         `json:"approvers"`
	Status    ApprovalStatusEnum `json:"status"`
	// Whether the user who makes the call may review the approval or not.
	UserMayReview *bool `json:"user_may_review,omitempty"`
	// Whether the user who makes the call is the requester of the approval or not.
	UserIsRequester *bool               `json:"user_is_requester,omitempty"`
	Requestor       NullableContactData `json:"requestor"`
	Reviewer        NullableContactData `json:"reviewer,omitempty"`
	// The date and time when the approval was approved or rejected.
	ReviewedAt *int64 `json:"reviewed_at,omitempty"`
	// The feedback provided by the reviewer when approving or rejecting the approval.
	ReviewerFeedback *string `json:"reviewer_feedback,omitempty"`
	// Whether the approval is archived or not. If archived, it will not be displayed in the list of approvals and it will not be possible to approve or reject it.
	Archived *bool `json:"archived,omitempty"`
}

type _Approval Approval

// NewApproval instantiates a new Approval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApproval(id string, createdAt int64, updatedAt int64, name string, sourceRef string, approvers []Approver, status ApprovalStatusEnum, requestor NullableContactData) *Approval {
	this := Approval{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Name = name
	this.SourceRef = sourceRef
	this.Approvers = approvers
	this.Status = status
	this.Requestor = requestor
	return &this
}

// NewApprovalWithDefaults instantiates a new Approval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApprovalWithDefaults() *Approval {
	this := Approval{}
	return &this
}

// GetId returns the Id field value
func (o *Approval) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Approval) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Approval) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Approval) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Approval) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Approval) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Approval) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Approval) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Approval) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetName returns the Name field value
func (o *Approval) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Approval) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Approval) SetName(v string) {
	o.Name = v
}

// GetSourceRef returns the SourceRef field value
func (o *Approval) GetSourceRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceRef
}

// GetSourceRefOk returns a tuple with the SourceRef field value
// and a boolean to check if the value has been set.
func (o *Approval) GetSourceRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceRef, true
}

// SetSourceRef sets field value
func (o *Approval) SetSourceRef(v string) {
	o.SourceRef = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Approval) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Approval) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Approval) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Approval) SetUrl(v string) {
	o.Url = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Approval) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Approval) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Approval) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Approval) SetMessage(v string) {
	o.Message = &v
}

// GetApprovers returns the Approvers field value
func (o *Approval) GetApprovers() []Approver {
	if o == nil {
		var ret []Approver
		return ret
	}

	return o.Approvers
}

// GetApproversOk returns a tuple with the Approvers field value
// and a boolean to check if the value has been set.
func (o *Approval) GetApproversOk() ([]Approver, bool) {
	if o == nil {
		return nil, false
	}
	return o.Approvers, true
}

// SetApprovers sets field value
func (o *Approval) SetApprovers(v []Approver) {
	o.Approvers = v
}

// GetStatus returns the Status field value
func (o *Approval) GetStatus() ApprovalStatusEnum {
	if o == nil {
		var ret ApprovalStatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Approval) GetStatusOk() (*ApprovalStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Approval) SetStatus(v ApprovalStatusEnum) {
	o.Status = v
}

// GetUserMayReview returns the UserMayReview field value if set, zero value otherwise.
func (o *Approval) GetUserMayReview() bool {
	if o == nil || IsNil(o.UserMayReview) {
		var ret bool
		return ret
	}
	return *o.UserMayReview
}

// GetUserMayReviewOk returns a tuple with the UserMayReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Approval) GetUserMayReviewOk() (*bool, bool) {
	if o == nil || IsNil(o.UserMayReview) {
		return nil, false
	}
	return o.UserMayReview, true
}

// HasUserMayReview returns a boolean if a field has been set.
func (o *Approval) HasUserMayReview() bool {
	if o != nil && !IsNil(o.UserMayReview) {
		return true
	}

	return false
}

// SetUserMayReview gets a reference to the given bool and assigns it to the UserMayReview field.
func (o *Approval) SetUserMayReview(v bool) {
	o.UserMayReview = &v
}

// GetUserIsRequester returns the UserIsRequester field value if set, zero value otherwise.
func (o *Approval) GetUserIsRequester() bool {
	if o == nil || IsNil(o.UserIsRequester) {
		var ret bool
		return ret
	}
	return *o.UserIsRequester
}

// GetUserIsRequesterOk returns a tuple with the UserIsRequester field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Approval) GetUserIsRequesterOk() (*bool, bool) {
	if o == nil || IsNil(o.UserIsRequester) {
		return nil, false
	}
	return o.UserIsRequester, true
}

// HasUserIsRequester returns a boolean if a field has been set.
func (o *Approval) HasUserIsRequester() bool {
	if o != nil && !IsNil(o.UserIsRequester) {
		return true
	}

	return false
}

// SetUserIsRequester gets a reference to the given bool and assigns it to the UserIsRequester field.
func (o *Approval) SetUserIsRequester(v bool) {
	o.UserIsRequester = &v
}

// GetRequestor returns the Requestor field value
// If the value is explicit nil, the zero value for ContactData will be returned
func (o *Approval) GetRequestor() ContactData {
	if o == nil || o.Requestor.Get() == nil {
		var ret ContactData
		return ret
	}

	return *o.Requestor.Get()
}

// GetRequestorOk returns a tuple with the Requestor field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Approval) GetRequestorOk() (*ContactData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Requestor.Get(), o.Requestor.IsSet()
}

// SetRequestor sets field value
func (o *Approval) SetRequestor(v ContactData) {
	o.Requestor.Set(&v)
}

// GetReviewer returns the Reviewer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Approval) GetReviewer() ContactData {
	if o == nil || IsNil(o.Reviewer.Get()) {
		var ret ContactData
		return ret
	}
	return *o.Reviewer.Get()
}

// GetReviewerOk returns a tuple with the Reviewer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Approval) GetReviewerOk() (*ContactData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reviewer.Get(), o.Reviewer.IsSet()
}

// HasReviewer returns a boolean if a field has been set.
func (o *Approval) HasReviewer() bool {
	if o != nil && o.Reviewer.IsSet() {
		return true
	}

	return false
}

// SetReviewer gets a reference to the given NullableContactData and assigns it to the Reviewer field.
func (o *Approval) SetReviewer(v ContactData) {
	o.Reviewer.Set(&v)
}

// SetReviewerNil sets the value for Reviewer to be an explicit nil
func (o *Approval) SetReviewerNil() {
	o.Reviewer.Set(nil)
}

// UnsetReviewer ensures that no value is present for Reviewer, not even an explicit nil
func (o *Approval) UnsetReviewer() {
	o.Reviewer.Unset()
}

// GetReviewedAt returns the ReviewedAt field value if set, zero value otherwise.
func (o *Approval) GetReviewedAt() int64 {
	if o == nil || IsNil(o.ReviewedAt) {
		var ret int64
		return ret
	}
	return *o.ReviewedAt
}

// GetReviewedAtOk returns a tuple with the ReviewedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Approval) GetReviewedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.ReviewedAt) {
		return nil, false
	}
	return o.ReviewedAt, true
}

// HasReviewedAt returns a boolean if a field has been set.
func (o *Approval) HasReviewedAt() bool {
	if o != nil && !IsNil(o.ReviewedAt) {
		return true
	}

	return false
}

// SetReviewedAt gets a reference to the given int64 and assigns it to the ReviewedAt field.
func (o *Approval) SetReviewedAt(v int64) {
	o.ReviewedAt = &v
}

// GetReviewerFeedback returns the ReviewerFeedback field value if set, zero value otherwise.
func (o *Approval) GetReviewerFeedback() string {
	if o == nil || IsNil(o.ReviewerFeedback) {
		var ret string
		return ret
	}
	return *o.ReviewerFeedback
}

// GetReviewerFeedbackOk returns a tuple with the ReviewerFeedback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Approval) GetReviewerFeedbackOk() (*string, bool) {
	if o == nil || IsNil(o.ReviewerFeedback) {
		return nil, false
	}
	return o.ReviewerFeedback, true
}

// HasReviewerFeedback returns a boolean if a field has been set.
func (o *Approval) HasReviewerFeedback() bool {
	if o != nil && !IsNil(o.ReviewerFeedback) {
		return true
	}

	return false
}

// SetReviewerFeedback gets a reference to the given string and assigns it to the ReviewerFeedback field.
func (o *Approval) SetReviewerFeedback(v string) {
	o.ReviewerFeedback = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *Approval) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Approval) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *Approval) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *Approval) SetArchived(v bool) {
	o.Archived = &v
}

func (o Approval) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Approval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["name"] = o.Name
	toSerialize["source_ref"] = o.SourceRef
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	toSerialize["approvers"] = o.Approvers
	toSerialize["status"] = o.Status
	if !IsNil(o.UserMayReview) {
		toSerialize["user_may_review"] = o.UserMayReview
	}
	if !IsNil(o.UserIsRequester) {
		toSerialize["user_is_requester"] = o.UserIsRequester
	}
	toSerialize["requestor"] = o.Requestor.Get()
	if o.Reviewer.IsSet() {
		toSerialize["reviewer"] = o.Reviewer.Get()
	}
	if !IsNil(o.ReviewedAt) {
		toSerialize["reviewed_at"] = o.ReviewedAt
	}
	if !IsNil(o.ReviewerFeedback) {
		toSerialize["reviewer_feedback"] = o.ReviewerFeedback
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	return toSerialize, nil
}

func (o *Approval) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"name",
		"source_ref",
		"approvers",
		"status",
		"requestor",
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

	varApproval := _Approval{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApproval)

	if err != nil {
		return err
	}

	*o = Approval(varApproval)

	return err
}

type NullableApproval struct {
	value *Approval
	isSet bool
}

func (v NullableApproval) Get() *Approval {
	return v.value
}

func (v *NullableApproval) Set(val *Approval) {
	v.value = val
	v.isSet = true
}

func (v NullableApproval) IsSet() bool {
	return v.isSet
}

func (v *NullableApproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApproval(val *Approval) *NullableApproval {
	return &NullableApproval{value: val, isSet: true}
}

func (v NullableApproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
