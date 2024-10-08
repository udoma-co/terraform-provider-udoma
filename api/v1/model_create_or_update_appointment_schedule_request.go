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

// checks if the CreateOrUpdateAppointmentScheduleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateAppointmentScheduleRequest{}

// CreateOrUpdateAppointmentScheduleRequest struct for CreateOrUpdateAppointmentScheduleRequest
type CreateOrUpdateAppointmentScheduleRequest struct {
	// The name of the appointment schedule
	Name string `json:"name"`
	// a map of values, where the key and values are strings
	Description *map[string]string `json:"description,omitempty"`
	// The duration of the appointment slots in minutes
	SlotDuration int32 `json:"slot_duration"`
	// The duration of the gap between appointments in minutes
	GapDuration *int32 `json:"gap_duration,omitempty"`
	// The color to use when displaying the appointment window in a calendar view
	Color *string `json:"color,omitempty"`
	// The appointment windows
	Windows []AppointmentWindow `json:"windows,omitempty"`
	// The reference to the appointment template that is used to determine what data to collect when booking a slot.
	TemplateRef string `json:"template_ref"`
}

type _CreateOrUpdateAppointmentScheduleRequest CreateOrUpdateAppointmentScheduleRequest

// NewCreateOrUpdateAppointmentScheduleRequest instantiates a new CreateOrUpdateAppointmentScheduleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateAppointmentScheduleRequest(name string, slotDuration int32, templateRef string) *CreateOrUpdateAppointmentScheduleRequest {
	this := CreateOrUpdateAppointmentScheduleRequest{}
	this.Name = name
	this.SlotDuration = slotDuration
	this.TemplateRef = templateRef
	return &this
}

// NewCreateOrUpdateAppointmentScheduleRequestWithDefaults instantiates a new CreateOrUpdateAppointmentScheduleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateAppointmentScheduleRequestWithDefaults() *CreateOrUpdateAppointmentScheduleRequest {
	this := CreateOrUpdateAppointmentScheduleRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateOrUpdateAppointmentScheduleRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppointmentScheduleRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrUpdateAppointmentScheduleRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOrUpdateAppointmentScheduleRequest) GetDescription() map[string]string {
	if o == nil || IsNil(o.Description) {
		var ret map[string]string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppointmentScheduleRequest) GetDescriptionOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrUpdateAppointmentScheduleRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given map[string]string and assigns it to the Description field.
func (o *CreateOrUpdateAppointmentScheduleRequest) SetDescription(v map[string]string) {
	o.Description = &v
}

// GetSlotDuration returns the SlotDuration field value
func (o *CreateOrUpdateAppointmentScheduleRequest) GetSlotDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SlotDuration
}

// GetSlotDurationOk returns a tuple with the SlotDuration field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppointmentScheduleRequest) GetSlotDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlotDuration, true
}

// SetSlotDuration sets field value
func (o *CreateOrUpdateAppointmentScheduleRequest) SetSlotDuration(v int32) {
	o.SlotDuration = v
}

// GetGapDuration returns the GapDuration field value if set, zero value otherwise.
func (o *CreateOrUpdateAppointmentScheduleRequest) GetGapDuration() int32 {
	if o == nil || IsNil(o.GapDuration) {
		var ret int32
		return ret
	}
	return *o.GapDuration
}

// GetGapDurationOk returns a tuple with the GapDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppointmentScheduleRequest) GetGapDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.GapDuration) {
		return nil, false
	}
	return o.GapDuration, true
}

// HasGapDuration returns a boolean if a field has been set.
func (o *CreateOrUpdateAppointmentScheduleRequest) HasGapDuration() bool {
	if o != nil && !IsNil(o.GapDuration) {
		return true
	}

	return false
}

// SetGapDuration gets a reference to the given int32 and assigns it to the GapDuration field.
func (o *CreateOrUpdateAppointmentScheduleRequest) SetGapDuration(v int32) {
	o.GapDuration = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *CreateOrUpdateAppointmentScheduleRequest) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppointmentScheduleRequest) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *CreateOrUpdateAppointmentScheduleRequest) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *CreateOrUpdateAppointmentScheduleRequest) SetColor(v string) {
	o.Color = &v
}

// GetWindows returns the Windows field value if set, zero value otherwise.
func (o *CreateOrUpdateAppointmentScheduleRequest) GetWindows() []AppointmentWindow {
	if o == nil || IsNil(o.Windows) {
		var ret []AppointmentWindow
		return ret
	}
	return o.Windows
}

// GetWindowsOk returns a tuple with the Windows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppointmentScheduleRequest) GetWindowsOk() ([]AppointmentWindow, bool) {
	if o == nil || IsNil(o.Windows) {
		return nil, false
	}
	return o.Windows, true
}

// HasWindows returns a boolean if a field has been set.
func (o *CreateOrUpdateAppointmentScheduleRequest) HasWindows() bool {
	if o != nil && !IsNil(o.Windows) {
		return true
	}

	return false
}

// SetWindows gets a reference to the given []AppointmentWindow and assigns it to the Windows field.
func (o *CreateOrUpdateAppointmentScheduleRequest) SetWindows(v []AppointmentWindow) {
	o.Windows = v
}

// GetTemplateRef returns the TemplateRef field value
func (o *CreateOrUpdateAppointmentScheduleRequest) GetTemplateRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateRef
}

// GetTemplateRefOk returns a tuple with the TemplateRef field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateAppointmentScheduleRequest) GetTemplateRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateRef, true
}

// SetTemplateRef sets field value
func (o *CreateOrUpdateAppointmentScheduleRequest) SetTemplateRef(v string) {
	o.TemplateRef = v
}

func (o CreateOrUpdateAppointmentScheduleRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateAppointmentScheduleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["slot_duration"] = o.SlotDuration
	if !IsNil(o.GapDuration) {
		toSerialize["gap_duration"] = o.GapDuration
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Windows) {
		toSerialize["windows"] = o.Windows
	}
	toSerialize["template_ref"] = o.TemplateRef
	return toSerialize, nil
}

func (o *CreateOrUpdateAppointmentScheduleRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"slot_duration",
		"template_ref",
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

	varCreateOrUpdateAppointmentScheduleRequest := _CreateOrUpdateAppointmentScheduleRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrUpdateAppointmentScheduleRequest)

	if err != nil {
		return err
	}

	*o = CreateOrUpdateAppointmentScheduleRequest(varCreateOrUpdateAppointmentScheduleRequest)

	return err
}

type NullableCreateOrUpdateAppointmentScheduleRequest struct {
	value *CreateOrUpdateAppointmentScheduleRequest
	isSet bool
}

func (v NullableCreateOrUpdateAppointmentScheduleRequest) Get() *CreateOrUpdateAppointmentScheduleRequest {
	return v.value
}

func (v *NullableCreateOrUpdateAppointmentScheduleRequest) Set(val *CreateOrUpdateAppointmentScheduleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateAppointmentScheduleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateAppointmentScheduleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateAppointmentScheduleRequest(val *CreateOrUpdateAppointmentScheduleRequest) *NullableCreateOrUpdateAppointmentScheduleRequest {
	return &NullableCreateOrUpdateAppointmentScheduleRequest{value: val, isSet: true}
}

func (v NullableCreateOrUpdateAppointmentScheduleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateAppointmentScheduleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
