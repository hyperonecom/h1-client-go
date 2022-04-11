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

// StorageProjectIsoTransfer struct for StorageProjectIsoTransfer
type StorageProjectIsoTransfer struct {
	Project string `json:"project"`
}

// NewStorageProjectIsoTransfer instantiates a new StorageProjectIsoTransfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageProjectIsoTransfer(project string) *StorageProjectIsoTransfer {
	this := StorageProjectIsoTransfer{}
	this.Project = project
	return &this
}

// NewStorageProjectIsoTransferWithDefaults instantiates a new StorageProjectIsoTransfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageProjectIsoTransferWithDefaults() *StorageProjectIsoTransfer {
	this := StorageProjectIsoTransfer{}
	return &this
}

// GetProject returns the Project field value
func (o *StorageProjectIsoTransfer) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *StorageProjectIsoTransfer) GetProjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *StorageProjectIsoTransfer) SetProject(v string) {
	o.Project = v
}

func (o StorageProjectIsoTransfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["project"] = o.Project
	}
	return json.Marshal(toSerialize)
}

type NullableStorageProjectIsoTransfer struct {
	value *StorageProjectIsoTransfer
	isSet bool
}

func (v NullableStorageProjectIsoTransfer) Get() *StorageProjectIsoTransfer {
	return v.value
}

func (v *NullableStorageProjectIsoTransfer) Set(val *StorageProjectIsoTransfer) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageProjectIsoTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageProjectIsoTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageProjectIsoTransfer(val *StorageProjectIsoTransfer) *NullableStorageProjectIsoTransfer {
	return &NullableStorageProjectIsoTransfer{value: val, isSet: true}
}

func (v NullableStorageProjectIsoTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageProjectIsoTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


