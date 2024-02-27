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

// checks if the AppointmentTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppointmentTemplate{}

// AppointmentTemplate struct for AppointmentTemplate
type AppointmentTemplate struct {
	// The unique identifier of the appointment template
	Id string `json:"id"`
	// The name of the appointment template
	Name string `json:"name"`
	// The expression that is used to generate the name of the appointment from the data that the user has entered when booking the appointment
	NameExpression string `json:"name_expression"`
	// The description of the appointment template
	Description *string `json:"description,omitempty"`
	// Constrain the minimum amount of time between the scheduling and the beginning of the appointment in minutes. Set this to 10 and people won't be able to schedule an appointment that's in less than 10 minutes.
	ScheduleBefore *int32     `json:"schedule_before,omitempty"`
	Form           CustomForm `json:"form"`
}

type _AppointmentTemplate AppointmentTemplate

// NewAppointmentTemplate instantiates a new AppointmentTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppointmentTemplate(id string, name string, nameExpression string, form CustomForm) *AppointmentTemplate {
	this := AppointmentTemplate{}
	this.Id = id
	this.Name = name
	this.NameExpression = nameExpression
	this.Form = form
	return &this
}

// NewAppointmentTemplateWithDefaults instantiates a new AppointmentTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppointmentTemplateWithDefaults() *AppointmentTemplate {
	this := AppointmentTemplate{}
	return &this
}

// GetId returns the Id field value
func (o *AppointmentTemplate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppointmentTemplate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppointmentTemplate) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AppointmentTemplate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppointmentTemplate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppointmentTemplate) SetName(v string) {
	o.Name = v
}

// GetNameExpression returns the NameExpression field value
func (o *AppointmentTemplate) GetNameExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameExpression
}

// GetNameExpressionOk returns a tuple with the NameExpression field value
// and a boolean to check if the value has been set.
func (o *AppointmentTemplate) GetNameExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameExpression, true
}

// SetNameExpression sets field value
func (o *AppointmentTemplate) SetNameExpression(v string) {
	o.NameExpression = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AppointmentTemplate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentTemplate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AppointmentTemplate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AppointmentTemplate) SetDescription(v string) {
	o.Description = &v
}

// GetScheduleBefore returns the ScheduleBefore field value if set, zero value otherwise.
func (o *AppointmentTemplate) GetScheduleBefore() int32 {
	if o == nil || IsNil(o.ScheduleBefore) {
		var ret int32
		return ret
	}
	return *o.ScheduleBefore
}

// GetScheduleBeforeOk returns a tuple with the ScheduleBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppointmentTemplate) GetScheduleBeforeOk() (*int32, bool) {
	if o == nil || IsNil(o.ScheduleBefore) {
		return nil, false
	}
	return o.ScheduleBefore, true
}

// HasScheduleBefore returns a boolean if a field has been set.
func (o *AppointmentTemplate) HasScheduleBefore() bool {
	if o != nil && !IsNil(o.ScheduleBefore) {
		return true
	}

	return false
}

// SetScheduleBefore gets a reference to the given int32 and assigns it to the ScheduleBefore field.
func (o *AppointmentTemplate) SetScheduleBefore(v int32) {
	o.ScheduleBefore = &v
}

// GetForm returns the Form field value
func (o *AppointmentTemplate) GetForm() CustomForm {
	if o == nil {
		var ret CustomForm
		return ret
	}

	return o.Form
}

// GetFormOk returns a tuple with the Form field value
// and a boolean to check if the value has been set.
func (o *AppointmentTemplate) GetFormOk() (*CustomForm, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Form, true
}

// SetForm sets field value
func (o *AppointmentTemplate) SetForm(v CustomForm) {
	o.Form = v
}

func (o AppointmentTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppointmentTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["name_expression"] = o.NameExpression
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ScheduleBefore) {
		toSerialize["schedule_before"] = o.ScheduleBefore
	}
	toSerialize["form"] = o.Form
	return toSerialize, nil
}

func (o *AppointmentTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"name_expression",
		"form",
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

	varAppointmentTemplate := _AppointmentTemplate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAppointmentTemplate)

	if err != nil {
		return err
	}

	*o = AppointmentTemplate(varAppointmentTemplate)

	return err
}

type NullableAppointmentTemplate struct {
	value *AppointmentTemplate
	isSet bool
}

func (v NullableAppointmentTemplate) Get() *AppointmentTemplate {
	return v.value
}

func (v *NullableAppointmentTemplate) Set(val *AppointmentTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableAppointmentTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableAppointmentTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppointmentTemplate(val *AppointmentTemplate) *NullableAppointmentTemplate {
	return &NullableAppointmentTemplate{value: val, isSet: true}
}

func (v NullableAppointmentTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppointmentTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
