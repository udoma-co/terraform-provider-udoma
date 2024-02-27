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

// checks if the Case type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Case{}

// Case Case represents a single case, typically raised by a tenant
type Case struct {
	Id        *string `json:"id,omitempty"`
	CreatedAt *int64  `json:"created_at,omitempty"`
	UpdatedAt *int64  `json:"updated_at,omitempty"`
	Name      *string `json:"name,omitempty"`
	// Indicates if the case has been archived
	Archived *bool `json:"archived,omitempty"`
	// Reference to property
	PropertyRef     *string  `json:"property_ref,omitempty"`
	PropertyAddress *Address `json:"property_address,omitempty"`
	// The ID of the case template used to create this case
	Template *string         `json:"template,omitempty"`
	Status   *CaseStatusEnum `json:"status,omitempty"`
	// List of all accounts that have access to the case (e.g. manager, tenenat, and service provider)
	Parties []CaseParty `json:"parties,omitempty"`
	// List of all status changes. Latest one is the current one
	StatusHistory []CaseStatus  `json:"status_history,omitempty"`
	Comments      []CaseComment `json:"comments,omitempty"`
	// List of feedback that was provided by the parties when changing the status of the case
	Feedback []CaseFeedbackEntry `json:"feedback,omitempty"`
	Assignee *string             `json:"assignee,omitempty"`
	// Input provided by the user when raising the case as a key-value map
	Data map[string]interface{} `json:"data,omitempty"`
}

// NewCase instantiates a new Case object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCase() *Case {
	this := Case{}
	return &this
}

// NewCaseWithDefaults instantiates a new Case object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaseWithDefaults() *Case {
	this := Case{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Case) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Case) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Case) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Case) SetId(v string) {
	o.Id = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Case) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Case) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Case) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *Case) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Case) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Case) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Case) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *Case) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Case) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Case) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Case) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Case) SetName(v string) {
	o.Name = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *Case) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Case) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *Case) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *Case) SetArchived(v bool) {
	o.Archived = &v
}

// GetPropertyRef returns the PropertyRef field value if set, zero value otherwise.
func (o *Case) GetPropertyRef() string {
	if o == nil || IsNil(o.PropertyRef) {
		var ret string
		return ret
	}
	return *o.PropertyRef
}

// GetPropertyRefOk returns a tuple with the PropertyRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Case) GetPropertyRefOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyRef) {
		return nil, false
	}
	return o.PropertyRef, true
}

// HasPropertyRef returns a boolean if a field has been set.
func (o *Case) HasPropertyRef() bool {
	if o != nil && !IsNil(o.PropertyRef) {
		return true
	}

	return false
}

// SetPropertyRef gets a reference to the given string and assigns it to the PropertyRef field.
func (o *Case) SetPropertyRef(v string) {
	o.PropertyRef = &v
}

// GetPropertyAddress returns the PropertyAddress field value if set, zero value otherwise.
func (o *Case) GetPropertyAddress() Address {
	if o == nil || IsNil(o.PropertyAddress) {
		var ret Address
		return ret
	}
	return *o.PropertyAddress
}

// GetPropertyAddressOk returns a tuple with the PropertyAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Case) GetPropertyAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.PropertyAddress) {
		return nil, false
	}
	return o.PropertyAddress, true
}

// HasPropertyAddress returns a boolean if a field has been set.
func (o *Case) HasPropertyAddress() bool {
	if o != nil && !IsNil(o.PropertyAddress) {
		return true
	}

	return false
}

// SetPropertyAddress gets a reference to the given Address and assigns it to the PropertyAddress field.
func (o *Case) SetPropertyAddress(v Address) {
	o.PropertyAddress = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *Case) GetTemplate() string {
	if o == nil || IsNil(o.Template) {
		var ret string
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Case) GetTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *Case) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given string and assigns it to the Template field.
func (o *Case) SetTemplate(v string) {
	o.Template = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Case) GetStatus() CaseStatusEnum {
	if o == nil || IsNil(o.Status) {
		var ret CaseStatusEnum
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Case) GetStatusOk() (*CaseStatusEnum, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Case) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CaseStatusEnum and assigns it to the Status field.
func (o *Case) SetStatus(v CaseStatusEnum) {
	o.Status = &v
}

// GetParties returns the Parties field value if set, zero value otherwise.
func (o *Case) GetParties() []CaseParty {
	if o == nil || IsNil(o.Parties) {
		var ret []CaseParty
		return ret
	}
	return o.Parties
}

// GetPartiesOk returns a tuple with the Parties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Case) GetPartiesOk() ([]CaseParty, bool) {
	if o == nil || IsNil(o.Parties) {
		return nil, false
	}
	return o.Parties, true
}

// HasParties returns a boolean if a field has been set.
func (o *Case) HasParties() bool {
	if o != nil && !IsNil(o.Parties) {
		return true
	}

	return false
}

// SetParties gets a reference to the given []CaseParty and assigns it to the Parties field.
func (o *Case) SetParties(v []CaseParty) {
	o.Parties = v
}

// GetStatusHistory returns the StatusHistory field value if set, zero value otherwise.
func (o *Case) GetStatusHistory() []CaseStatus {
	if o == nil || IsNil(o.StatusHistory) {
		var ret []CaseStatus
		return ret
	}
	return o.StatusHistory
}

// GetStatusHistoryOk returns a tuple with the StatusHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Case) GetStatusHistoryOk() ([]CaseStatus, bool) {
	if o == nil || IsNil(o.StatusHistory) {
		return nil, false
	}
	return o.StatusHistory, true
}

// HasStatusHistory returns a boolean if a field has been set.
func (o *Case) HasStatusHistory() bool {
	if o != nil && !IsNil(o.StatusHistory) {
		return true
	}

	return false
}

// SetStatusHistory gets a reference to the given []CaseStatus and assigns it to the StatusHistory field.
func (o *Case) SetStatusHistory(v []CaseStatus) {
	o.StatusHistory = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *Case) GetComments() []CaseComment {
	if o == nil || IsNil(o.Comments) {
		var ret []CaseComment
		return ret
	}
	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Case) GetCommentsOk() ([]CaseComment, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *Case) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given []CaseComment and assigns it to the Comments field.
func (o *Case) SetComments(v []CaseComment) {
	o.Comments = v
}

// GetFeedback returns the Feedback field value if set, zero value otherwise.
func (o *Case) GetFeedback() []CaseFeedbackEntry {
	if o == nil || IsNil(o.Feedback) {
		var ret []CaseFeedbackEntry
		return ret
	}
	return o.Feedback
}

// GetFeedbackOk returns a tuple with the Feedback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Case) GetFeedbackOk() ([]CaseFeedbackEntry, bool) {
	if o == nil || IsNil(o.Feedback) {
		return nil, false
	}
	return o.Feedback, true
}

// HasFeedback returns a boolean if a field has been set.
func (o *Case) HasFeedback() bool {
	if o != nil && !IsNil(o.Feedback) {
		return true
	}

	return false
}

// SetFeedback gets a reference to the given []CaseFeedbackEntry and assigns it to the Feedback field.
func (o *Case) SetFeedback(v []CaseFeedbackEntry) {
	o.Feedback = v
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *Case) GetAssignee() string {
	if o == nil || IsNil(o.Assignee) {
		var ret string
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Case) GetAssigneeOk() (*string, bool) {
	if o == nil || IsNil(o.Assignee) {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *Case) HasAssignee() bool {
	if o != nil && !IsNil(o.Assignee) {
		return true
	}

	return false
}

// SetAssignee gets a reference to the given string and assigns it to the Assignee field.
func (o *Case) SetAssignee(v string) {
	o.Assignee = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Case) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Case) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Case) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *Case) SetData(v map[string]interface{}) {
	o.Data = v
}

func (o Case) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Case) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !IsNil(o.PropertyRef) {
		toSerialize["property_ref"] = o.PropertyRef
	}
	if !IsNil(o.PropertyAddress) {
		toSerialize["property_address"] = o.PropertyAddress
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Parties) {
		toSerialize["parties"] = o.Parties
	}
	if !IsNil(o.StatusHistory) {
		toSerialize["status_history"] = o.StatusHistory
	}
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Feedback) {
		toSerialize["feedback"] = o.Feedback
	}
	if !IsNil(o.Assignee) {
		toSerialize["assignee"] = o.Assignee
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCase struct {
	value *Case
	isSet bool
}

func (v NullableCase) Get() *Case {
	return v.value
}

func (v *NullableCase) Set(val *Case) {
	v.value = val
	v.isSet = true
}

func (v NullableCase) IsSet() bool {
	return v.isSet
}

func (v *NullableCase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCase(val *Case) *NullableCase {
	return &NullableCase{value: val, isSet: true}
}

func (v NullableCase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
