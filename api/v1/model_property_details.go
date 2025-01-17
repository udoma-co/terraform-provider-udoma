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

// checks if the PropertyDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropertyDetails{}

// PropertyDetails Holds extra details about a property
type PropertyDetails struct {
	// Floor/storey of a building.
	Floor            *string                  `json:"floor,omitempty"`
	Area             *float32                 `json:"area,omitempty"`
	Rooms            *float32                 `json:"rooms,omitempty"`
	Bedrooms         *float32                 `json:"bedrooms,omitempty"`
	Bathrooms        *float32                 `json:"bathrooms,omitempty"`
	Balcony          *float32                 `json:"balcony,omitempty"`
	ConstructionYear *float32                 `json:"construction_year,omitempty"`
	Parking          *PropertyParkingTypeEnum `json:"parking,omitempty"`
	// List of Floor types present in a property
	FloorType []PropertyFloorTypeEnum `json:"floor_type,omitempty"`
	// List of Heating types present in a property
	HeatingType []PropertyHeatingTypeEnum `json:"heating_type,omitempty"`
	// List of window types present in a property
	WindowType []PropertyWindowTypeEnum `json:"window_type,omitempty"`
	Furnishing *PropertyFurnishingEnum  `json:"furnishing,omitempty"`
	// \"A building with special architectural or historic interest, considered to be of national importance and therefor worth being preserved/protected\"
	HistoricProperty *bool `json:"historic_property,omitempty"`
}

// NewPropertyDetails instantiates a new PropertyDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyDetails() *PropertyDetails {
	this := PropertyDetails{}
	return &this
}

// NewPropertyDetailsWithDefaults instantiates a new PropertyDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyDetailsWithDefaults() *PropertyDetails {
	this := PropertyDetails{}
	return &this
}

// GetFloor returns the Floor field value if set, zero value otherwise.
func (o *PropertyDetails) GetFloor() string {
	if o == nil || IsNil(o.Floor) {
		var ret string
		return ret
	}
	return *o.Floor
}

// GetFloorOk returns a tuple with the Floor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDetails) GetFloorOk() (*string, bool) {
	if o == nil || IsNil(o.Floor) {
		return nil, false
	}
	return o.Floor, true
}

// HasFloor returns a boolean if a field has been set.
func (o *PropertyDetails) HasFloor() bool {
	if o != nil && !IsNil(o.Floor) {
		return true
	}

	return false
}

// SetFloor gets a reference to the given string and assigns it to the Floor field.
func (o *PropertyDetails) SetFloor(v string) {
	o.Floor = &v
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *PropertyDetails) GetArea() float32 {
	if o == nil || IsNil(o.Area) {
		var ret float32
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDetails) GetAreaOk() (*float32, bool) {
	if o == nil || IsNil(o.Area) {
		return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *PropertyDetails) HasArea() bool {
	if o != nil && !IsNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given float32 and assigns it to the Area field.
func (o *PropertyDetails) SetArea(v float32) {
	o.Area = &v
}

// GetRooms returns the Rooms field value if set, zero value otherwise.
func (o *PropertyDetails) GetRooms() float32 {
	if o == nil || IsNil(o.Rooms) {
		var ret float32
		return ret
	}
	return *o.Rooms
}

// GetRoomsOk returns a tuple with the Rooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDetails) GetRoomsOk() (*float32, bool) {
	if o == nil || IsNil(o.Rooms) {
		return nil, false
	}
	return o.Rooms, true
}

// HasRooms returns a boolean if a field has been set.
func (o *PropertyDetails) HasRooms() bool {
	if o != nil && !IsNil(o.Rooms) {
		return true
	}

	return false
}

// SetRooms gets a reference to the given float32 and assigns it to the Rooms field.
func (o *PropertyDetails) SetRooms(v float32) {
	o.Rooms = &v
}

// GetBedrooms returns the Bedrooms field value if set, zero value otherwise.
func (o *PropertyDetails) GetBedrooms() float32 {
	if o == nil || IsNil(o.Bedrooms) {
		var ret float32
		return ret
	}
	return *o.Bedrooms
}

// GetBedroomsOk returns a tuple with the Bedrooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDetails) GetBedroomsOk() (*float32, bool) {
	if o == nil || IsNil(o.Bedrooms) {
		return nil, false
	}
	return o.Bedrooms, true
}

// HasBedrooms returns a boolean if a field has been set.
func (o *PropertyDetails) HasBedrooms() bool {
	if o != nil && !IsNil(o.Bedrooms) {
		return true
	}

	return false
}

// SetBedrooms gets a reference to the given float32 and assigns it to the Bedrooms field.
func (o *PropertyDetails) SetBedrooms(v float32) {
	o.Bedrooms = &v
}

// GetBathrooms returns the Bathrooms field value if set, zero value otherwise.
func (o *PropertyDetails) GetBathrooms() float32 {
	if o == nil || IsNil(o.Bathrooms) {
		var ret float32
		return ret
	}
	return *o.Bathrooms
}

// GetBathroomsOk returns a tuple with the Bathrooms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDetails) GetBathroomsOk() (*float32, bool) {
	if o == nil || IsNil(o.Bathrooms) {
		return nil, false
	}
	return o.Bathrooms, true
}

// HasBathrooms returns a boolean if a field has been set.
func (o *PropertyDetails) HasBathrooms() bool {
	if o != nil && !IsNil(o.Bathrooms) {
		return true
	}

	return false
}

// SetBathrooms gets a reference to the given float32 and assigns it to the Bathrooms field.
func (o *PropertyDetails) SetBathrooms(v float32) {
	o.Bathrooms = &v
}

// GetBalcony returns the Balcony field value if set, zero value otherwise.
func (o *PropertyDetails) GetBalcony() float32 {
	if o == nil || IsNil(o.Balcony) {
		var ret float32
		return ret
	}
	return *o.Balcony
}

// GetBalconyOk returns a tuple with the Balcony field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDetails) GetBalconyOk() (*float32, bool) {
	if o == nil || IsNil(o.Balcony) {
		return nil, false
	}
	return o.Balcony, true
}

// HasBalcony returns a boolean if a field has been set.
func (o *PropertyDetails) HasBalcony() bool {
	if o != nil && !IsNil(o.Balcony) {
		return true
	}

	return false
}

// SetBalcony gets a reference to the given float32 and assigns it to the Balcony field.
func (o *PropertyDetails) SetBalcony(v float32) {
	o.Balcony = &v
}

// GetConstructionYear returns the ConstructionYear field value if set, zero value otherwise.
func (o *PropertyDetails) GetConstructionYear() float32 {
	if o == nil || IsNil(o.ConstructionYear) {
		var ret float32
		return ret
	}
	return *o.ConstructionYear
}

// GetConstructionYearOk returns a tuple with the ConstructionYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDetails) GetConstructionYearOk() (*float32, bool) {
	if o == nil || IsNil(o.ConstructionYear) {
		return nil, false
	}
	return o.ConstructionYear, true
}

// HasConstructionYear returns a boolean if a field has been set.
func (o *PropertyDetails) HasConstructionYear() bool {
	if o != nil && !IsNil(o.ConstructionYear) {
		return true
	}

	return false
}

// SetConstructionYear gets a reference to the given float32 and assigns it to the ConstructionYear field.
func (o *PropertyDetails) SetConstructionYear(v float32) {
	o.ConstructionYear = &v
}

// GetParking returns the Parking field value if set, zero value otherwise.
func (o *PropertyDetails) GetParking() PropertyParkingTypeEnum {
	if o == nil || IsNil(o.Parking) {
		var ret PropertyParkingTypeEnum
		return ret
	}
	return *o.Parking
}

// GetParkingOk returns a tuple with the Parking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDetails) GetParkingOk() (*PropertyParkingTypeEnum, bool) {
	if o == nil || IsNil(o.Parking) {
		return nil, false
	}
	return o.Parking, true
}

// HasParking returns a boolean if a field has been set.
func (o *PropertyDetails) HasParking() bool {
	if o != nil && !IsNil(o.Parking) {
		return true
	}

	return false
}

// SetParking gets a reference to the given PropertyParkingTypeEnum and assigns it to the Parking field.
func (o *PropertyDetails) SetParking(v PropertyParkingTypeEnum) {
	o.Parking = &v
}

// GetFloorType returns the FloorType field value if set, zero value otherwise.
func (o *PropertyDetails) GetFloorType() []PropertyFloorTypeEnum {
	if o == nil || IsNil(o.FloorType) {
		var ret []PropertyFloorTypeEnum
		return ret
	}
	return o.FloorType
}

// GetFloorTypeOk returns a tuple with the FloorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDetails) GetFloorTypeOk() ([]PropertyFloorTypeEnum, bool) {
	if o == nil || IsNil(o.FloorType) {
		return nil, false
	}
	return o.FloorType, true
}

// HasFloorType returns a boolean if a field has been set.
func (o *PropertyDetails) HasFloorType() bool {
	if o != nil && !IsNil(o.FloorType) {
		return true
	}

	return false
}

// SetFloorType gets a reference to the given []PropertyFloorTypeEnum and assigns it to the FloorType field.
func (o *PropertyDetails) SetFloorType(v []PropertyFloorTypeEnum) {
	o.FloorType = v
}

// GetHeatingType returns the HeatingType field value if set, zero value otherwise.
func (o *PropertyDetails) GetHeatingType() []PropertyHeatingTypeEnum {
	if o == nil || IsNil(o.HeatingType) {
		var ret []PropertyHeatingTypeEnum
		return ret
	}
	return o.HeatingType
}

// GetHeatingTypeOk returns a tuple with the HeatingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDetails) GetHeatingTypeOk() ([]PropertyHeatingTypeEnum, bool) {
	if o == nil || IsNil(o.HeatingType) {
		return nil, false
	}
	return o.HeatingType, true
}

// HasHeatingType returns a boolean if a field has been set.
func (o *PropertyDetails) HasHeatingType() bool {
	if o != nil && !IsNil(o.HeatingType) {
		return true
	}

	return false
}

// SetHeatingType gets a reference to the given []PropertyHeatingTypeEnum and assigns it to the HeatingType field.
func (o *PropertyDetails) SetHeatingType(v []PropertyHeatingTypeEnum) {
	o.HeatingType = v
}

// GetWindowType returns the WindowType field value if set, zero value otherwise.
func (o *PropertyDetails) GetWindowType() []PropertyWindowTypeEnum {
	if o == nil || IsNil(o.WindowType) {
		var ret []PropertyWindowTypeEnum
		return ret
	}
	return o.WindowType
}

// GetWindowTypeOk returns a tuple with the WindowType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDetails) GetWindowTypeOk() ([]PropertyWindowTypeEnum, bool) {
	if o == nil || IsNil(o.WindowType) {
		return nil, false
	}
	return o.WindowType, true
}

// HasWindowType returns a boolean if a field has been set.
func (o *PropertyDetails) HasWindowType() bool {
	if o != nil && !IsNil(o.WindowType) {
		return true
	}

	return false
}

// SetWindowType gets a reference to the given []PropertyWindowTypeEnum and assigns it to the WindowType field.
func (o *PropertyDetails) SetWindowType(v []PropertyWindowTypeEnum) {
	o.WindowType = v
}

// GetFurnishing returns the Furnishing field value if set, zero value otherwise.
func (o *PropertyDetails) GetFurnishing() PropertyFurnishingEnum {
	if o == nil || IsNil(o.Furnishing) {
		var ret PropertyFurnishingEnum
		return ret
	}
	return *o.Furnishing
}

// GetFurnishingOk returns a tuple with the Furnishing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDetails) GetFurnishingOk() (*PropertyFurnishingEnum, bool) {
	if o == nil || IsNil(o.Furnishing) {
		return nil, false
	}
	return o.Furnishing, true
}

// HasFurnishing returns a boolean if a field has been set.
func (o *PropertyDetails) HasFurnishing() bool {
	if o != nil && !IsNil(o.Furnishing) {
		return true
	}

	return false
}

// SetFurnishing gets a reference to the given PropertyFurnishingEnum and assigns it to the Furnishing field.
func (o *PropertyDetails) SetFurnishing(v PropertyFurnishingEnum) {
	o.Furnishing = &v
}

// GetHistoricProperty returns the HistoricProperty field value if set, zero value otherwise.
func (o *PropertyDetails) GetHistoricProperty() bool {
	if o == nil || IsNil(o.HistoricProperty) {
		var ret bool
		return ret
	}
	return *o.HistoricProperty
}

// GetHistoricPropertyOk returns a tuple with the HistoricProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyDetails) GetHistoricPropertyOk() (*bool, bool) {
	if o == nil || IsNil(o.HistoricProperty) {
		return nil, false
	}
	return o.HistoricProperty, true
}

// HasHistoricProperty returns a boolean if a field has been set.
func (o *PropertyDetails) HasHistoricProperty() bool {
	if o != nil && !IsNil(o.HistoricProperty) {
		return true
	}

	return false
}

// SetHistoricProperty gets a reference to the given bool and assigns it to the HistoricProperty field.
func (o *PropertyDetails) SetHistoricProperty(v bool) {
	o.HistoricProperty = &v
}

func (o PropertyDetails) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropertyDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Floor) {
		toSerialize["floor"] = o.Floor
	}
	if !IsNil(o.Area) {
		toSerialize["area"] = o.Area
	}
	if !IsNil(o.Rooms) {
		toSerialize["rooms"] = o.Rooms
	}
	if !IsNil(o.Bedrooms) {
		toSerialize["bedrooms"] = o.Bedrooms
	}
	if !IsNil(o.Bathrooms) {
		toSerialize["bathrooms"] = o.Bathrooms
	}
	if !IsNil(o.Balcony) {
		toSerialize["balcony"] = o.Balcony
	}
	if !IsNil(o.ConstructionYear) {
		toSerialize["construction_year"] = o.ConstructionYear
	}
	if !IsNil(o.Parking) {
		toSerialize["parking"] = o.Parking
	}
	if !IsNil(o.FloorType) {
		toSerialize["floor_type"] = o.FloorType
	}
	if !IsNil(o.HeatingType) {
		toSerialize["heating_type"] = o.HeatingType
	}
	if !IsNil(o.WindowType) {
		toSerialize["window_type"] = o.WindowType
	}
	if !IsNil(o.Furnishing) {
		toSerialize["furnishing"] = o.Furnishing
	}
	if !IsNil(o.HistoricProperty) {
		toSerialize["historic_property"] = o.HistoricProperty
	}
	return toSerialize, nil
}

type NullablePropertyDetails struct {
	value *PropertyDetails
	isSet bool
}

func (v NullablePropertyDetails) Get() *PropertyDetails {
	return v.value
}

func (v *NullablePropertyDetails) Set(val *PropertyDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyDetails(val *PropertyDetails) *NullablePropertyDetails {
	return &NullablePropertyDetails{value: val, isSet: true}
}

func (v NullablePropertyDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
