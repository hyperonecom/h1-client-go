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

// StorageProjectBucketUpload struct for StorageProjectBucketUpload
type StorageProjectBucketUpload struct {
	Name string `json:"name"`
}

// NewStorageProjectBucketUpload instantiates a new StorageProjectBucketUpload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageProjectBucketUpload(name string, ) *StorageProjectBucketUpload {
	this := StorageProjectBucketUpload{}
	this.Name = name
	return &this
}

// NewStorageProjectBucketUploadWithDefaults instantiates a new StorageProjectBucketUpload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageProjectBucketUploadWithDefaults() *StorageProjectBucketUpload {
	this := StorageProjectBucketUpload{}
	return &this
}

// GetName returns the Name field value
func (o *StorageProjectBucketUpload) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StorageProjectBucketUpload) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StorageProjectBucketUpload) SetName(v string) {
	o.Name = v
}

func (o StorageProjectBucketUpload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableStorageProjectBucketUpload struct {
	value *StorageProjectBucketUpload
	isSet bool
}

func (v NullableStorageProjectBucketUpload) Get() *StorageProjectBucketUpload {
	return v.value
}

func (v *NullableStorageProjectBucketUpload) Set(val *StorageProjectBucketUpload) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageProjectBucketUpload) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageProjectBucketUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageProjectBucketUpload(val *StorageProjectBucketUpload) *NullableStorageProjectBucketUpload {
	return &NullableStorageProjectBucketUpload{value: val, isSet: true}
}

func (v NullableStorageProjectBucketUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageProjectBucketUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


