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

// checks if the ConnectorQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectorQuery{}

// ConnectorQuery This is the query that is sent to be ran by the client and return data from their database.
type ConnectorQuery struct {
	// Unique and immutable ID attribute of the entity that is generated when  the instance is created. The ID is unique within the system accross all accounts and it can be used to reference the entity in other entities  or to retrieve it from the backend.
	Id string `json:"id"`
	// The date and time the entity was created
	CreatedAt int64 `json:"created_at"`
	// The date and time the entity was last updated
	UpdatedAt int64 `json:"updated_at"`
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
	// A timestamp of the last updated object. This should be updated by the process script, to track the last update to an object we performed.
	UpdatedLast *int64 `json:"updated_last,omitempty"`
}

type _ConnectorQuery ConnectorQuery

// NewConnectorQuery instantiates a new ConnectorQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorQuery(id string, createdAt int64, updatedAt int64, name string, connectorRef string, priority int32, entityLimit int32, statement string, processScript string) *ConnectorQuery {
	this := ConnectorQuery{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Name = name
	this.ConnectorRef = connectorRef
	this.Priority = priority
	this.EntityLimit = entityLimit
	this.Statement = statement
	this.ProcessScript = processScript
	return &this
}

// NewConnectorQueryWithDefaults instantiates a new ConnectorQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorQueryWithDefaults() *ConnectorQuery {
	this := ConnectorQuery{}
	return &this
}

// GetId returns the Id field value
func (o *ConnectorQuery) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ConnectorQuery) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ConnectorQuery) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ConnectorQuery) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ConnectorQuery) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ConnectorQuery) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ConnectorQuery) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ConnectorQuery) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ConnectorQuery) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

// GetName returns the Name field value
func (o *ConnectorQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConnectorQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConnectorQuery) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ConnectorQuery) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorQuery) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ConnectorQuery) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ConnectorQuery) SetDescription(v string) {
	o.Description = &v
}

// GetConnectorRef returns the ConnectorRef field value
func (o *ConnectorQuery) GetConnectorRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorRef
}

// GetConnectorRefOk returns a tuple with the ConnectorRef field value
// and a boolean to check if the value has been set.
func (o *ConnectorQuery) GetConnectorRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorRef, true
}

// SetConnectorRef sets field value
func (o *ConnectorQuery) SetConnectorRef(v string) {
	o.ConnectorRef = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ConnectorQuery) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorQuery) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ConnectorQuery) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ConnectorQuery) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPriority returns the Priority field value
func (o *ConnectorQuery) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *ConnectorQuery) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *ConnectorQuery) SetPriority(v int32) {
	o.Priority = v
}

// GetEntityLimit returns the EntityLimit field value
func (o *ConnectorQuery) GetEntityLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EntityLimit
}

// GetEntityLimitOk returns a tuple with the EntityLimit field value
// and a boolean to check if the value has been set.
func (o *ConnectorQuery) GetEntityLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityLimit, true
}

// SetEntityLimit sets field value
func (o *ConnectorQuery) SetEntityLimit(v int32) {
	o.EntityLimit = v
}

// GetStatement returns the Statement field value
func (o *ConnectorQuery) GetStatement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Statement
}

// GetStatementOk returns a tuple with the Statement field value
// and a boolean to check if the value has been set.
func (o *ConnectorQuery) GetStatementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Statement, true
}

// SetStatement sets field value
func (o *ConnectorQuery) SetStatement(v string) {
	o.Statement = v
}

// GetProcessScript returns the ProcessScript field value
func (o *ConnectorQuery) GetProcessScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessScript
}

// GetProcessScriptOk returns a tuple with the ProcessScript field value
// and a boolean to check if the value has been set.
func (o *ConnectorQuery) GetProcessScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessScript, true
}

// SetProcessScript sets field value
func (o *ConnectorQuery) SetProcessScript(v string) {
	o.ProcessScript = v
}

// GetUpdatedLast returns the UpdatedLast field value if set, zero value otherwise.
func (o *ConnectorQuery) GetUpdatedLast() int64 {
	if o == nil || IsNil(o.UpdatedLast) {
		var ret int64
		return ret
	}
	return *o.UpdatedLast
}

// GetUpdatedLastOk returns a tuple with the UpdatedLast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorQuery) GetUpdatedLastOk() (*int64, bool) {
	if o == nil || IsNil(o.UpdatedLast) {
		return nil, false
	}
	return o.UpdatedLast, true
}

// HasUpdatedLast returns a boolean if a field has been set.
func (o *ConnectorQuery) HasUpdatedLast() bool {
	if o != nil && !IsNil(o.UpdatedLast) {
		return true
	}

	return false
}

// SetUpdatedLast gets a reference to the given int64 and assigns it to the UpdatedLast field.
func (o *ConnectorQuery) SetUpdatedLast(v int64) {
	o.UpdatedLast = &v
}

func (o ConnectorQuery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectorQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
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
	if !IsNil(o.UpdatedLast) {
		toSerialize["updated_last"] = o.UpdatedLast
	}
	return toSerialize, nil
}

func (o *ConnectorQuery) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
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

	varConnectorQuery := _ConnectorQuery{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConnectorQuery)

	if err != nil {
		return err
	}

	*o = ConnectorQuery(varConnectorQuery)

	return err
}

type NullableConnectorQuery struct {
	value *ConnectorQuery
	isSet bool
}

func (v NullableConnectorQuery) Get() *ConnectorQuery {
	return v.value
}

func (v *NullableConnectorQuery) Set(val *ConnectorQuery) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorQuery) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorQuery(val *ConnectorQuery) *NullableConnectorQuery {
	return &NullableConnectorQuery{value: val, isSet: true}
}

func (v NullableConnectorQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
