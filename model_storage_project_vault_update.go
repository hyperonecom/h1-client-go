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

// StorageProjectVaultUpdate struct for StorageProjectVaultUpdate
type StorageProjectVaultUpdate struct {
	Name *string `json:"name,omitempty"`
}

// NewStorageProjectVaultUpdate instantiates a new StorageProjectVaultUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageProjectVaultUpdate() *StorageProjectVaultUpdate {
	this := StorageProjectVaultUpdate{}
	return &this
}

// NewStorageProjectVaultUpdateWithDefaults instantiates a new StorageProjectVaultUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageProjectVaultUpdateWithDefaults() *StorageProjectVaultUpdate {
	this := StorageProjectVaultUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StorageProjectVaultUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProjectVaultUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StorageProjectVaultUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StorageProjectVaultUpdate) SetName(v string) {
	o.Name = &v
}

func (o StorageProjectVaultUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableStorageProjectVaultUpdate struct {
	value *StorageProjectVaultUpdate
	isSet bool
}

func (v NullableStorageProjectVaultUpdate) Get() *StorageProjectVaultUpdate {
	return v.value
}

func (v *NullableStorageProjectVaultUpdate) Set(val *StorageProjectVaultUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageProjectVaultUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageProjectVaultUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageProjectVaultUpdate(val *StorageProjectVaultUpdate) *NullableStorageProjectVaultUpdate {
	return &NullableStorageProjectVaultUpdate{value: val, isSet: true}
}

func (v NullableStorageProjectVaultUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageProjectVaultUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


