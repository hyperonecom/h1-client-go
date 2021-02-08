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

// StorageProjectImageTransfer struct for StorageProjectImageTransfer
type StorageProjectImageTransfer struct {
	Project string `json:"project"`
}

// NewStorageProjectImageTransfer instantiates a new StorageProjectImageTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageProjectImageTransfer(project string, ) *StorageProjectImageTransfer {
	this := StorageProjectImageTransfer{}
	this.Project = project
	return &this
}

// NewStorageProjectImageTransferWithDefaults instantiates a new StorageProjectImageTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageProjectImageTransferWithDefaults() *StorageProjectImageTransfer {
	this := StorageProjectImageTransfer{}
	return &this
}

// GetProject returns the Project field value
func (o *StorageProjectImageTransfer) GetProject() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *StorageProjectImageTransfer) GetProjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *StorageProjectImageTransfer) SetProject(v string) {
	o.Project = v
}

func (o StorageProjectImageTransfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["project"] = o.Project
	}
	return json.Marshal(toSerialize)
}

type NullableStorageProjectImageTransfer struct {
	value *StorageProjectImageTransfer
	isSet bool
}

func (v NullableStorageProjectImageTransfer) Get() *StorageProjectImageTransfer {
	return v.value
}

func (v *NullableStorageProjectImageTransfer) Set(val *StorageProjectImageTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageProjectImageTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageProjectImageTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageProjectImageTransfer(val *StorageProjectImageTransfer) *NullableStorageProjectImageTransfer {
	return &NullableStorageProjectImageTransfer{value: val, isSet: true}
}

func (v NullableStorageProjectImageTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageProjectImageTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


