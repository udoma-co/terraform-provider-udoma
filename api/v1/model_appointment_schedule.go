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

// checks if the AppointmentSchedule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppointmentSchedule{}

// AppointmentSchedule struct for AppointmentSchedule
type AppointmentSchedule struct {
	// Unique and immutable ID attribute of the entity that is generated when  the instance is created. The ID is unique within the system accross all accounts and it can be used to reference the entity in other entities  or to retrieve it from the backend.
	Id string `json:"id"`
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
	Windows  []AppointmentWindow `json:"windows,omitempty"`
	Template AppointmentTemplate `json:"template"`
}

type _AppointmentSchedule AppointmentSchedule

// NewAppointmentSchedule instantiates a new AppointmentSchedule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppointmentSchedule(id string, name string, slotDuration int32, template AppointmentTemplate) *AppointmentSchedule {
	this := AppointmentSchedule{}
	this.Id = id
	this.Name = name
	this.SlotDuration = slotDuration
	this.Template = template
	return &this
}

// NewAppointmentScheduleWithDefaults instantiates a new AppointmentSchedule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppointmentScheduleWithDefaults() *AppointmentSchedule {
	this := AppointmentSchedule{}
	return &this
}

// GetId returns the Id field value
func (o *AppointmentSchedule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppointmentSchedule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppointmentSchedule) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AppointmentSchedule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppointmentSchedule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppointmentSchedule) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AppointmentSchedule) GetDescription() map[string]string {
	if o == nil || IsNil(o.Description) {
		var ret map[string]string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentSchedule) GetDescriptionOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AppointmentSchedule) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given map[string]string and assigns it to the Description field.
func (o *AppointmentSchedule) SetDescription(v map[string]string) {
	o.Description = &v
}

// GetSlotDuration returns the SlotDuration field value
func (o *AppointmentSchedule) GetSlotDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SlotDuration
}

// GetSlotDurationOk returns a tuple with the SlotDuration field value
// and a boolean to check if the value has been set.
func (o *AppointmentSchedule) GetSlotDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlotDuration, true
}

// SetSlotDuration sets field value
func (o *AppointmentSchedule) SetSlotDuration(v int32) {
	o.SlotDuration = v
}

// GetGapDuration returns the GapDuration field value if set, zero value otherwise.
func (o *AppointmentSchedule) GetGapDuration() int32 {
	if o == nil || IsNil(o.GapDuration) {
		var ret int32
		return ret
	}
	return *o.GapDuration
}

// GetGapDurationOk returns a tuple with the GapDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentSchedule) GetGapDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.GapDuration) {
		return nil, false
	}
	return o.GapDuration, true
}

// HasGapDuration returns a boolean if a field has been set.
func (o *AppointmentSchedule) HasGapDuration() bool {
	if o != nil && !IsNil(o.GapDuration) {
		return true
	}

	return false
}

// SetGapDuration gets a reference to the given int32 and assigns it to the GapDuration field.
func (o *AppointmentSchedule) SetGapDuration(v int32) {
	o.GapDuration = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *AppointmentSchedule) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentSchedule) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *AppointmentSchedule) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *AppointmentSchedule) SetColor(v string) {
	o.Color = &v
}

// GetWindows returns the Windows field value if set, zero value otherwise.
func (o *AppointmentSchedule) GetWindows() []AppointmentWindow {
	if o == nil || IsNil(o.Windows) {
		var ret []AppointmentWindow
		return ret
	}
	return o.Windows
}

// GetWindowsOk returns a tuple with the Windows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentSchedule) GetWindowsOk() ([]AppointmentWindow, bool) {
	if o == nil || IsNil(o.Windows) {
		return nil, false
	}
	return o.Windows, true
}

// HasWindows returns a boolean if a field has been set.
func (o *AppointmentSchedule) HasWindows() bool {
	if o != nil && !IsNil(o.Windows) {
		return true
	}

	return false
}

// SetWindows gets a reference to the given []AppointmentWindow and assigns it to the Windows field.
func (o *AppointmentSchedule) SetWindows(v []AppointmentWindow) {
	o.Windows = v
}

// GetTemplate returns the Template field value
func (o *AppointmentSchedule) GetTemplate() AppointmentTemplate {
	if o == nil {
		var ret AppointmentTemplate
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *AppointmentSchedule) GetTemplateOk() (*AppointmentTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *AppointmentSchedule) SetTemplate(v AppointmentTemplate) {
	o.Template = v
}

func (o AppointmentSchedule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppointmentSchedule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
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
	toSerialize["template"] = o.Template
	return toSerialize, nil
}

func (o *AppointmentSchedule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"slot_duration",
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

	varAppointmentSchedule := _AppointmentSchedule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppointmentSchedule)

	if err != nil {
		return err
	}

	*o = AppointmentSchedule(varAppointmentSchedule)

	return err
}

type NullableAppointmentSchedule struct {
	value *AppointmentSchedule
	isSet bool
}

func (v NullableAppointmentSchedule) Get() *AppointmentSchedule {
	return v.value
}

func (v *NullableAppointmentSchedule) Set(val *AppointmentSchedule) {
	v.value = val
	v.isSet = true
}

func (v NullableAppointmentSchedule) IsSet() bool {
	return v.isSet
}

func (v *NullableAppointmentSchedule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppointmentSchedule(val *AppointmentSchedule) *NullableAppointmentSchedule {
	return &NullableAppointmentSchedule{value: val, isSet: true}
}

func (v NullableAppointmentSchedule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppointmentSchedule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
