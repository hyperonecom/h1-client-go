/*
 * HyperOne
 *
 * HyperOne API
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package h1

import (
	"encoding/json"
)

// DatabaseProjectInstanceCreate struct for DatabaseProjectInstanceCreate
type DatabaseProjectInstanceCreate struct {
	Name string `json:"name"`
	Service string `json:"service"`
	Source *string `json:"source,omitempty"`
	Tag *[]Tag `json:"tag,omitempty"`
}

// NewDatabaseProjectInstanceCreate instantiates a new DatabaseProjectInstanceCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseProjectInstanceCreate(name string, service string, ) *DatabaseProjectInstanceCreate {
	this := DatabaseProjectInstanceCreate{}
	this.Name = name
	this.Service = service
	return &this
}

// NewDatabaseProjectInstanceCreateWithDefaults instantiates a new DatabaseProjectInstanceCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseProjectInstanceCreateWithDefaults() *DatabaseProjectInstanceCreate {
	this := DatabaseProjectInstanceCreate{}
	return &this
}

// GetName returns the Name field value
func (o *DatabaseProjectInstanceCreate) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatabaseProjectInstanceCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatabaseProjectInstanceCreate) SetName(v string) {
	o.Name = v
}

// GetService returns the Service field value
func (o *DatabaseProjectInstanceCreate) GetService() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *DatabaseProjectInstanceCreate) GetServiceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *DatabaseProjectInstanceCreate) SetService(v string) {
	o.Service = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *DatabaseProjectInstanceCreate) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseProjectInstanceCreate) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *DatabaseProjectInstanceCreate) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *DatabaseProjectInstanceCreate) SetSource(v string) {
	o.Source = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *DatabaseProjectInstanceCreate) GetTag() []Tag {
	if o == nil || o.Tag == nil {
		var ret []Tag
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseProjectInstanceCreate) GetTagOk() (*[]Tag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *DatabaseProjectInstanceCreate) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Tag and assigns it to the Tag field.
func (o *DatabaseProjectInstanceCreate) SetTag(v []Tag) {
	o.Tag = &v
}

func (o DatabaseProjectInstanceCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["service"] = o.Service
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableDatabaseProjectInstanceCreate struct {
	value *DatabaseProjectInstanceCreate
	isSet bool
}

func (v NullableDatabaseProjectInstanceCreate) Get() *DatabaseProjectInstanceCreate {
	return v.value
}

func (v *NullableDatabaseProjectInstanceCreate) Set(val *DatabaseProjectInstanceCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseProjectInstanceCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseProjectInstanceCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseProjectInstanceCreate(val *DatabaseProjectInstanceCreate) *NullableDatabaseProjectInstanceCreate {
	return &NullableDatabaseProjectInstanceCreate{value: val, isSet: true}
}

func (v NullableDatabaseProjectInstanceCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseProjectInstanceCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


