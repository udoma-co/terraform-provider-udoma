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

// checks if the WorkflowEntrypointAttributesPartial type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowEntrypointAttributesPartial{}

// WorkflowEntrypointAttributesPartial struct for WorkflowEntrypointAttributesPartial
type WorkflowEntrypointAttributesPartial struct {
	// the ID of the workflow definition
	WorkflowDefinitionRef string                     `json:"workflow_definition_ref"`
	AppLocation           WorkflowEntrypointLocation `json:"app_location"`
	// optional filter that can be used to limit where the entrypoint is shown, e.g. for cases this can be the case template, for reports this can be the report  definition, etc.
	LocationFilter *string `json:"location_filter,omitempty"`
	// Optional icon of the entrypoint
	Icon *string `json:"icon,omitempty"`
	// a map of values, where the key and values are strings
	Label map[string]string `json:"label"`
	// optional JS script to be executed before the workflow is started
	InitScript *string `json:"init_script,omitempty"`
	// whether the init step should be skipped or not
	SkipInitStep *bool `json:"skip_init_step,omitempty"`
}

type _WorkflowEntrypointAttributesPartial WorkflowEntrypointAttributesPartial

// NewWorkflowEntrypointAttributesPartial instantiates a new WorkflowEntrypointAttributesPartial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowEntrypointAttributesPartial(workflowDefinitionRef string, appLocation WorkflowEntrypointLocation, label map[string]string) *WorkflowEntrypointAttributesPartial {
	this := WorkflowEntrypointAttributesPartial{}
	this.WorkflowDefinitionRef = workflowDefinitionRef
	this.AppLocation = appLocation
	this.Label = label
	return &this
}

// NewWorkflowEntrypointAttributesPartialWithDefaults instantiates a new WorkflowEntrypointAttributesPartial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowEntrypointAttributesPartialWithDefaults() *WorkflowEntrypointAttributesPartial {
	this := WorkflowEntrypointAttributesPartial{}
	return &this
}

// GetWorkflowDefinitionRef returns the WorkflowDefinitionRef field value
func (o *WorkflowEntrypointAttributesPartial) GetWorkflowDefinitionRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkflowDefinitionRef
}

// GetWorkflowDefinitionRefOk returns a tuple with the WorkflowDefinitionRef field value
// and a boolean to check if the value has been set.
func (o *WorkflowEntrypointAttributesPartial) GetWorkflowDefinitionRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkflowDefinitionRef, true
}

// SetWorkflowDefinitionRef sets field value
func (o *WorkflowEntrypointAttributesPartial) SetWorkflowDefinitionRef(v string) {
	o.WorkflowDefinitionRef = v
}

// GetAppLocation returns the AppLocation field value
func (o *WorkflowEntrypointAttributesPartial) GetAppLocation() WorkflowEntrypointLocation {
	if o == nil {
		var ret WorkflowEntrypointLocation
		return ret
	}

	return o.AppLocation
}

// GetAppLocationOk returns a tuple with the AppLocation field value
// and a boolean to check if the value has been set.
func (o *WorkflowEntrypointAttributesPartial) GetAppLocationOk() (*WorkflowEntrypointLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppLocation, true
}

// SetAppLocation sets field value
func (o *WorkflowEntrypointAttributesPartial) SetAppLocation(v WorkflowEntrypointLocation) {
	o.AppLocation = v
}

// GetLocationFilter returns the LocationFilter field value if set, zero value otherwise.
func (o *WorkflowEntrypointAttributesPartial) GetLocationFilter() string {
	if o == nil || IsNil(o.LocationFilter) {
		var ret string
		return ret
	}
	return *o.LocationFilter
}

// GetLocationFilterOk returns a tuple with the LocationFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowEntrypointAttributesPartial) GetLocationFilterOk() (*string, bool) {
	if o == nil || IsNil(o.LocationFilter) {
		return nil, false
	}
	return o.LocationFilter, true
}

// HasLocationFilter returns a boolean if a field has been set.
func (o *WorkflowEntrypointAttributesPartial) HasLocationFilter() bool {
	if o != nil && !IsNil(o.LocationFilter) {
		return true
	}

	return false
}

// SetLocationFilter gets a reference to the given string and assigns it to the LocationFilter field.
func (o *WorkflowEntrypointAttributesPartial) SetLocationFilter(v string) {
	o.LocationFilter = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *WorkflowEntrypointAttributesPartial) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowEntrypointAttributesPartial) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *WorkflowEntrypointAttributesPartial) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *WorkflowEntrypointAttributesPartial) SetIcon(v string) {
	o.Icon = &v
}

// GetLabel returns the Label field value
func (o *WorkflowEntrypointAttributesPartial) GetLabel() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *WorkflowEntrypointAttributesPartial) GetLabelOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *WorkflowEntrypointAttributesPartial) SetLabel(v map[string]string) {
	o.Label = v
}

// GetInitScript returns the InitScript field value if set, zero value otherwise.
func (o *WorkflowEntrypointAttributesPartial) GetInitScript() string {
	if o == nil || IsNil(o.InitScript) {
		var ret string
		return ret
	}
	return *o.InitScript
}

// GetInitScriptOk returns a tuple with the InitScript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowEntrypointAttributesPartial) GetInitScriptOk() (*string, bool) {
	if o == nil || IsNil(o.InitScript) {
		return nil, false
	}
	return o.InitScript, true
}

// HasInitScript returns a boolean if a field has been set.
func (o *WorkflowEntrypointAttributesPartial) HasInitScript() bool {
	if o != nil && !IsNil(o.InitScript) {
		return true
	}

	return false
}

// SetInitScript gets a reference to the given string and assigns it to the InitScript field.
func (o *WorkflowEntrypointAttributesPartial) SetInitScript(v string) {
	o.InitScript = &v
}

// GetSkipInitStep returns the SkipInitStep field value if set, zero value otherwise.
func (o *WorkflowEntrypointAttributesPartial) GetSkipInitStep() bool {
	if o == nil || IsNil(o.SkipInitStep) {
		var ret bool
		return ret
	}
	return *o.SkipInitStep
}

// GetSkipInitStepOk returns a tuple with the SkipInitStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowEntrypointAttributesPartial) GetSkipInitStepOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipInitStep) {
		return nil, false
	}
	return o.SkipInitStep, true
}

// HasSkipInitStep returns a boolean if a field has been set.
func (o *WorkflowEntrypointAttributesPartial) HasSkipInitStep() bool {
	if o != nil && !IsNil(o.SkipInitStep) {
		return true
	}

	return false
}

// SetSkipInitStep gets a reference to the given bool and assigns it to the SkipInitStep field.
func (o *WorkflowEntrypointAttributesPartial) SetSkipInitStep(v bool) {
	o.SkipInitStep = &v
}

func (o WorkflowEntrypointAttributesPartial) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowEntrypointAttributesPartial) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["workflow_definition_ref"] = o.WorkflowDefinitionRef
	toSerialize["app_location"] = o.AppLocation
	if !IsNil(o.LocationFilter) {
		toSerialize["location_filter"] = o.LocationFilter
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	toSerialize["label"] = o.Label
	if !IsNil(o.InitScript) {
		toSerialize["init_script"] = o.InitScript
	}
	if !IsNil(o.SkipInitStep) {
		toSerialize["skip_init_step"] = o.SkipInitStep
	}
	return toSerialize, nil
}

func (o *WorkflowEntrypointAttributesPartial) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varWorkflowEntrypointAttributesPartial := _WorkflowEntrypointAttributesPartial{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkflowEntrypointAttributesPartial)

	if err != nil {
		return err
	}

	*o = WorkflowEntrypointAttributesPartial(varWorkflowEntrypointAttributesPartial)

	return err
}

type NullableWorkflowEntrypointAttributesPartial struct {
	value *WorkflowEntrypointAttributesPartial
	isSet bool
}

func (v NullableWorkflowEntrypointAttributesPartial) Get() *WorkflowEntrypointAttributesPartial {
	return v.value
}

func (v *NullableWorkflowEntrypointAttributesPartial) Set(val *WorkflowEntrypointAttributesPartial) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowEntrypointAttributesPartial) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowEntrypointAttributesPartial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowEntrypointAttributesPartial(val *WorkflowEntrypointAttributesPartial) *NullableWorkflowEntrypointAttributesPartial {
	return &NullableWorkflowEntrypointAttributesPartial{value: val, isSet: true}
}

func (v NullableWorkflowEntrypointAttributesPartial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowEntrypointAttributesPartial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}