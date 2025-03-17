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

// checks if the CreateOrUpdateConnectorQueryRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateOrUpdateConnectorQueryRequest{}

// CreateOrUpdateConnectorQueryRequest Request used to create or update a connector query. If the query already exists it will be updated.
type CreateOrUpdateConnectorQueryRequest struct {
	// A unique and descriptive name given to the query.
	Name string `json:"name"`
	// The name to display in the webapp.
	Description *string `json:"description,omitempty"`
	// The Name of the connector this query belongs to.
	ConnectorRef string `json:"connector_ref"`
	// Whether the Query is enabled and should be used to sync data
	Enabled *bool `json:"enabled,omitempty"`
	// The priority used to tell the order in which to execute the queries
	Priority int32 `json:"priority"`
	// The amount of entities to query for at once.
	EntityLimit int32 `json:"entity_limit"`
	// The statement to be executed by the connector client against their DB.
	Statement string `json:"statement"`
	// JavaScript code that processes the response into an object that can be understood by the backend code and upserted into the database.
	ProcessScript string `json:"process_script"`
}

type _CreateOrUpdateConnectorQueryRequest CreateOrUpdateConnectorQueryRequest

// NewCreateOrUpdateConnectorQueryRequest instantiates a new CreateOrUpdateConnectorQueryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateOrUpdateConnectorQueryRequest(name string, connectorRef string, priority int32, entityLimit int32, statement string, processScript string) *CreateOrUpdateConnectorQueryRequest {
	this := CreateOrUpdateConnectorQueryRequest{}
	this.Name = name
	this.ConnectorRef = connectorRef
	this.Priority = priority
	this.EntityLimit = entityLimit
	this.Statement = statement
	this.ProcessScript = processScript
	return &this
}

// NewCreateOrUpdateConnectorQueryRequestWithDefaults instantiates a new CreateOrUpdateConnectorQueryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateOrUpdateConnectorQueryRequestWithDefaults() *CreateOrUpdateConnectorQueryRequest {
	this := CreateOrUpdateConnectorQueryRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateOrUpdateConnectorQueryRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateConnectorQueryRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateOrUpdateConnectorQueryRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateOrUpdateConnectorQueryRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateConnectorQueryRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateOrUpdateConnectorQueryRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateOrUpdateConnectorQueryRequest) SetDescription(v string) {
	o.Description = &v
}

// GetConnectorRef returns the ConnectorRef field value
func (o *CreateOrUpdateConnectorQueryRequest) GetConnectorRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorRef
}

// GetConnectorRefOk returns a tuple with the ConnectorRef field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateConnectorQueryRequest) GetConnectorRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorRef, true
}

// SetConnectorRef sets field value
func (o *CreateOrUpdateConnectorQueryRequest) SetConnectorRef(v string) {
	o.ConnectorRef = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateOrUpdateConnectorQueryRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateConnectorQueryRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateOrUpdateConnectorQueryRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateOrUpdateConnectorQueryRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPriority returns the Priority field value
func (o *CreateOrUpdateConnectorQueryRequest) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateConnectorQueryRequest) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *CreateOrUpdateConnectorQueryRequest) SetPriority(v int32) {
	o.Priority = v
}

// GetEntityLimit returns the EntityLimit field value
func (o *CreateOrUpdateConnectorQueryRequest) GetEntityLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EntityLimit
}

// GetEntityLimitOk returns a tuple with the EntityLimit field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateConnectorQueryRequest) GetEntityLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityLimit, true
}

// SetEntityLimit sets field value
func (o *CreateOrUpdateConnectorQueryRequest) SetEntityLimit(v int32) {
	o.EntityLimit = v
}

// GetStatement returns the Statement field value
func (o *CreateOrUpdateConnectorQueryRequest) GetStatement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Statement
}

// GetStatementOk returns a tuple with the Statement field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateConnectorQueryRequest) GetStatementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Statement, true
}

// SetStatement sets field value
func (o *CreateOrUpdateConnectorQueryRequest) SetStatement(v string) {
	o.Statement = v
}

// GetProcessScript returns the ProcessScript field value
func (o *CreateOrUpdateConnectorQueryRequest) GetProcessScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessScript
}

// GetProcessScriptOk returns a tuple with the ProcessScript field value
// and a boolean to check if the value has been set.
func (o *CreateOrUpdateConnectorQueryRequest) GetProcessScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessScript, true
}

// SetProcessScript sets field value
func (o *CreateOrUpdateConnectorQueryRequest) SetProcessScript(v string) {
	o.ProcessScript = v
}

func (o CreateOrUpdateConnectorQueryRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateOrUpdateConnectorQueryRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["connector_ref"] = o.ConnectorRef
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["priority"] = o.Priority
	toSerialize["entity_limit"] = o.EntityLimit
	toSerialize["statement"] = o.Statement
	toSerialize["process_script"] = o.ProcessScript
	return toSerialize, nil
}

func (o *CreateOrUpdateConnectorQueryRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"connector_ref",
		"priority",
		"entity_limit",
		"statement",
		"process_script",
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

	varCreateOrUpdateConnectorQueryRequest := _CreateOrUpdateConnectorQueryRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateOrUpdateConnectorQueryRequest)

	if err != nil {
		return err
	}

	*o = CreateOrUpdateConnectorQueryRequest(varCreateOrUpdateConnectorQueryRequest)

	return err
}

type NullableCreateOrUpdateConnectorQueryRequest struct {
	value *CreateOrUpdateConnectorQueryRequest
	isSet bool
}

func (v NullableCreateOrUpdateConnectorQueryRequest) Get() *CreateOrUpdateConnectorQueryRequest {
	return v.value
}

func (v *NullableCreateOrUpdateConnectorQueryRequest) Set(val *CreateOrUpdateConnectorQueryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateOrUpdateConnectorQueryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateOrUpdateConnectorQueryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateOrUpdateConnectorQueryRequest(val *CreateOrUpdateConnectorQueryRequest) *NullableCreateOrUpdateConnectorQueryRequest {
	return &NullableCreateOrUpdateConnectorQueryRequest{value: val, isSet: true}
}

func (v NullableCreateOrUpdateConnectorQueryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateOrUpdateConnectorQueryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
