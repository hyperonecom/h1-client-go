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

// StorageProjectIsoCreate struct for StorageProjectIsoCreate
type StorageProjectIsoCreate struct {
	Name string `json:"name"`
	Source string `json:"source"`
	Tag *[]Tag `json:"tag,omitempty"`
}

// NewStorageProjectIsoCreate instantiates a new StorageProjectIsoCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageProjectIsoCreate(name string, source string) *StorageProjectIsoCreate {
	this := StorageProjectIsoCreate{}
	this.Name = name
	this.Source = source
	return &this
}

// NewStorageProjectIsoCreateWithDefaults instantiates a new StorageProjectIsoCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageProjectIsoCreateWithDefaults() *StorageProjectIsoCreate {
	this := StorageProjectIsoCreate{}
	return &this
}

// GetName returns the Name field value
func (o *StorageProjectIsoCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StorageProjectIsoCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StorageProjectIsoCreate) SetName(v string) {
	o.Name = v
}

// GetSource returns the Source field value
func (o *StorageProjectIsoCreate) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *StorageProjectIsoCreate) GetSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *StorageProjectIsoCreate) SetSource(v string) {
	o.Source = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *StorageProjectIsoCreate) GetTag() []Tag {
	if o == nil || o.Tag == nil {
		var ret []Tag
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProjectIsoCreate) GetTagOk() (*[]Tag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *StorageProjectIsoCreate) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Tag and assigns it to the Tag field.
func (o *StorageProjectIsoCreate) SetTag(v []Tag) {
	o.Tag = &v
}

func (o StorageProjectIsoCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableStorageProjectIsoCreate struct {
	value *StorageProjectIsoCreate
	isSet bool
}

func (v NullableStorageProjectIsoCreate) Get() *StorageProjectIsoCreate {
	return v.value
}

func (v *NullableStorageProjectIsoCreate) Set(val *StorageProjectIsoCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageProjectIsoCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageProjectIsoCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageProjectIsoCreate(val *StorageProjectIsoCreate) *NullableStorageProjectIsoCreate {
	return &NullableStorageProjectIsoCreate{value: val, isSet: true}
}

func (v NullableStorageProjectIsoCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageProjectIsoCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


