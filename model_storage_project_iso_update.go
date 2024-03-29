/*
HyperOne

HyperOne API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// StorageProjectIsoUpdate struct for StorageProjectIsoUpdate
type StorageProjectIsoUpdate struct {
	Name *string `json:"name,omitempty"`
}

// NewStorageProjectIsoUpdate instantiates a new StorageProjectIsoUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageProjectIsoUpdate() *StorageProjectIsoUpdate {
	this := StorageProjectIsoUpdate{}
	return &this
}

// NewStorageProjectIsoUpdateWithDefaults instantiates a new StorageProjectIsoUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageProjectIsoUpdateWithDefaults() *StorageProjectIsoUpdate {
	this := StorageProjectIsoUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageProjectIsoUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProjectIsoUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageProjectIsoUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageProjectIsoUpdate) SetName(v string) {
	o.Name = &v
}

func (o StorageProjectIsoUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableStorageProjectIsoUpdate struct {
	value *StorageProjectIsoUpdate
	isSet bool
}

func (v NullableStorageProjectIsoUpdate) Get() *StorageProjectIsoUpdate {
	return v.value
}

func (v *NullableStorageProjectIsoUpdate) Set(val *StorageProjectIsoUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageProjectIsoUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageProjectIsoUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageProjectIsoUpdate(val *StorageProjectIsoUpdate) *NullableStorageProjectIsoUpdate {
	return &NullableStorageProjectIsoUpdate{value: val, isSet: true}
}

func (v NullableStorageProjectIsoUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageProjectIsoUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


