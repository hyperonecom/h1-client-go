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

// StorageProjectImageUpdate struct for StorageProjectImageUpdate
type StorageProjectImageUpdate struct {
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewStorageProjectImageUpdate instantiates a new StorageProjectImageUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageProjectImageUpdate() *StorageProjectImageUpdate {
	this := StorageProjectImageUpdate{}
	return &this
}

// NewStorageProjectImageUpdateWithDefaults instantiates a new StorageProjectImageUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageProjectImageUpdateWithDefaults() *StorageProjectImageUpdate {
	this := StorageProjectImageUpdate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StorageProjectImageUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProjectImageUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StorageProjectImageUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StorageProjectImageUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageProjectImageUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProjectImageUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageProjectImageUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageProjectImageUpdate) SetName(v string) {
	o.Name = &v
}

func (o StorageProjectImageUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableStorageProjectImageUpdate struct {
	value *StorageProjectImageUpdate
	isSet bool
}

func (v NullableStorageProjectImageUpdate) Get() *StorageProjectImageUpdate {
	return v.value
}

func (v *NullableStorageProjectImageUpdate) Set(val *StorageProjectImageUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageProjectImageUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageProjectImageUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageProjectImageUpdate(val *StorageProjectImageUpdate) *NullableStorageProjectImageUpdate {
	return &NullableStorageProjectImageUpdate{value: val, isSet: true}
}

func (v NullableStorageProjectImageUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageProjectImageUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


