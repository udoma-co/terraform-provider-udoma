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

// checks if the WorkflowEntrypoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowEntrypoint{}

// WorkflowEntrypoint an app entrypoint to a workflow execution
type WorkflowEntrypoint struct {
	// Unique and immutable ID attribute of the entity that is generated when the instance is created. The ID is unique within the system accross all accounts and it can be used to reference the entity in other entities or to retrieve it from the backend.
	Id string `json:"id"`
	// The date and time the entity was created
	CreatedAt int64 `json:"created_at"`
	// The date and time the entity was last updated
	UpdatedAt int64 `json:"updated_at"`
	// the ID of the workflow definition
	WorkflowDefinitionRef string                     `json:"workflow_definition_ref"`
	AppLocation           WorkflowEntrypointLocation `json:"app_location"`
	// optional filters that can be used to limit where the entrypoint is shown, e.g. for cases this can be the case template, for reports this can be the report  definition, etc.
	LocationFilters []WorkflowEntrypointFilter `json:"location_filters,omitempty"`
	// optional validations that can be used to determine if the workflow can  be started via this entrypoint. Depending on the result, the workflow can be started or not.
	Validations []WorkflowEntrypointValidation `json:"validations,omitempty"`
	// Optional icon of the entrypoint
	Icon *string `json:"icon,omitempty"`
	// a map of values, where the key and values are strings
	Label map[string]string `json:"label"`
	// optional JS script to be executed before the workflow is started
	InitScript *string `json:"init_script,omitempty"`
}

type _WorkflowEntrypoint WorkflowEntrypoint

// NewWorkflowEntrypoint instantiates a new WorkflowEntrypoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowEntrypoint(id string, createdAt int64, updatedAt int64, workflowDefinitionRef string, appLocation WorkflowEntrypointLocation, label map[string]string) *WorkflowEntrypoint {
	this := WorkflowEntrypoint{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.WorkflowDefinitionRef = workflowDefinitionRef
	this.AppLocation = appLocation
	this.Label = label
	return &this
}

// NewWorkflowEntrypointWithDefaults instantiates a new WorkflowEntrypoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowEntrypointWithDefaults() *WorkflowEntrypoint {
	this := WorkflowEntrypoint{}
	return &this
}

// GetId returns the Id field value
func (o *WorkflowEntrypoint) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkflowEntrypoint) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkflowEntrypoint) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WorkflowEntrypoint) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkflowEntrypoint) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WorkflowEntrypoint) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *WorkflowEntrypoint) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkflowEntrypoint) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *WorkflowEntrypoint) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetWorkflowDefinitionRef returns the WorkflowDefinitionRef field value
func (o *WorkflowEntrypoint) GetWorkflowDefinitionRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowDefinitionRef
}

// GetWorkflowDefinitionRefOk returns a tuple with the WorkflowDefinitionRef field value
// and a boolean to check if the value has been set.
func (o *WorkflowEntrypoint) GetWorkflowDefinitionRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowDefinitionRef, true
}

// SetWorkflowDefinitionRef sets field value
func (o *WorkflowEntrypoint) SetWorkflowDefinitionRef(v string) {
	o.WorkflowDefinitionRef = v
}

// GetAppLocation returns the AppLocation field value
func (o *WorkflowEntrypoint) GetAppLocation() WorkflowEntrypointLocation {
	if o == nil {
		var ret WorkflowEntrypointLocation
		return ret
	}

	return o.AppLocation
}

// GetAppLocationOk returns a tuple with the AppLocation field value
// and a boolean to check if the value has been set.
func (o *WorkflowEntrypoint) GetAppLocationOk() (*WorkflowEntrypointLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppLocation, true
}

// SetAppLocation sets field value
func (o *WorkflowEntrypoint) SetAppLocation(v WorkflowEntrypointLocation) {
	o.AppLocation = v
}

// GetLocationFilters returns the LocationFilters field value if set, zero value otherwise.
func (o *WorkflowEntrypoint) GetLocationFilters() []WorkflowEntrypointFilter {
	if o == nil || IsNil(o.LocationFilters) {
		var ret []WorkflowEntrypointFilter
		return ret
	}
	return o.LocationFilters
}

// GetLocationFiltersOk returns a tuple with the LocationFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowEntrypoint) GetLocationFiltersOk() ([]WorkflowEntrypointFilter, bool) {
	if o == nil || IsNil(o.LocationFilters) {
		return nil, false
	}
	return o.LocationFilters, true
}

// HasLocationFilters returns a boolean if a field has been set.
func (o *WorkflowEntrypoint) HasLocationFilters() bool {
	if o != nil && !IsNil(o.LocationFilters) {
		return true
	}

	return false
}

// SetLocationFilters gets a reference to the given []WorkflowEntrypointFilter and assigns it to the LocationFilters field.
func (o *WorkflowEntrypoint) SetLocationFilters(v []WorkflowEntrypointFilter) {
	o.LocationFilters = v
}

// GetValidations returns the Validations field value if set, zero value otherwise.
func (o *WorkflowEntrypoint) GetValidations() []WorkflowEntrypointValidation {
	if o == nil || IsNil(o.Validations) {
		var ret []WorkflowEntrypointValidation
		return ret
	}
	return o.Validations
}

// GetValidationsOk returns a tuple with the Validations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowEntrypoint) GetValidationsOk() ([]WorkflowEntrypointValidation, bool) {
	if o == nil || IsNil(o.Validations) {
		return nil, false
	}
	return o.Validations, true
}

// HasValidations returns a boolean if a field has been set.
func (o *WorkflowEntrypoint) HasValidations() bool {
	if o != nil && !IsNil(o.Validations) {
		return true
	}

	return false
}

// SetValidations gets a reference to the given []WorkflowEntrypointValidation and assigns it to the Validations field.
func (o *WorkflowEntrypoint) SetValidations(v []WorkflowEntrypointValidation) {
	o.Validations = v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *WorkflowEntrypoint) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowEntrypoint) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *WorkflowEntrypoint) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *WorkflowEntrypoint) SetIcon(v string) {
	o.Icon = &v
}

// GetLabel returns the Label field value
func (o *WorkflowEntrypoint) GetLabel() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *WorkflowEntrypoint) GetLabelOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *WorkflowEntrypoint) SetLabel(v map[string]string) {
	o.Label = v
}

// GetInitScript returns the InitScript field value if set, zero value otherwise.
func (o *WorkflowEntrypoint) GetInitScript() string {
	if o == nil || IsNil(o.InitScript) {
		var ret string
		return ret
	}
	return *o.InitScript
}

// GetInitScriptOk returns a tuple with the InitScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowEntrypoint) GetInitScriptOk() (*string, bool) {
	if o == nil || IsNil(o.InitScript) {
		return nil, false
	}
	return o.InitScript, true
}

// HasInitScript returns a boolean if a field has been set.
func (o *WorkflowEntrypoint) HasInitScript() bool {
	if o != nil && !IsNil(o.InitScript) {
		return true
	}

	return false
}

// SetInitScript gets a reference to the given string and assigns it to the InitScript field.
func (o *WorkflowEntrypoint) SetInitScript(v string) {
	o.InitScript = &v
}

func (o WorkflowEntrypoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowEntrypoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["workflow_definition_ref"] = o.WorkflowDefinitionRef
	toSerialize["app_location"] = o.AppLocation
	if !IsNil(o.LocationFilters) {
		toSerialize["location_filters"] = o.LocationFilters
	}
	if !IsNil(o.Validations) {
		toSerialize["validations"] = o.Validations
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	toSerialize["label"] = o.Label
	if !IsNil(o.InitScript) {
		toSerialize["init_script"] = o.InitScript
	}
	return toSerialize, nil
}

func (o *WorkflowEntrypoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"workflow_definition_ref",
		"app_location",
		"label",
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

	varWorkflowEntrypoint := _WorkflowEntrypoint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkflowEntrypoint)

	if err != nil {
		return err
	}

	*o = WorkflowEntrypoint(varWorkflowEntrypoint)

	return err
}

type NullableWorkflowEntrypoint struct {
	value *WorkflowEntrypoint
	isSet bool
}

func (v NullableWorkflowEntrypoint) Get() *WorkflowEntrypoint {
	return v.value
}

func (v *NullableWorkflowEntrypoint) Set(val *WorkflowEntrypoint) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowEntrypoint) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowEntrypoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowEntrypoint(val *WorkflowEntrypoint) *NullableWorkflowEntrypoint {
	return &NullableWorkflowEntrypoint{value: val, isSet: true}
}

func (v NullableWorkflowEntrypoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowEntrypoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
