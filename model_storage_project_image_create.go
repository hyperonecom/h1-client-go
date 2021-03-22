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

// StorageProjectImageCreate struct for StorageProjectImageCreate
type StorageProjectImageCreate struct {
	Name string `json:"name"`
	Service *string `json:"service,omitempty"`
	Vm *string `json:"vm,omitempty"`
	Replica *string `json:"replica,omitempty"`
	Description *string `json:"description,omitempty"`
	Tag *[]Tag `json:"tag,omitempty"`
}

// NewStorageProjectImageCreate instantiates a new StorageProjectImageCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageProjectImageCreate(name string) *StorageProjectImageCreate {
	this := StorageProjectImageCreate{}
	this.Name = name
	var service string = "564639bc052c084e2f2e3266"
	this.Service = &service
	return &this
}

// NewStorageProjectImageCreateWithDefaults instantiates a new StorageProjectImageCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageProjectImageCreateWithDefaults() *StorageProjectImageCreate {
	this := StorageProjectImageCreate{}
	var service string = "564639bc052c084e2f2e3266"
	this.Service = &service
	return &this
}

// GetName returns the Name field value
func (o *StorageProjectImageCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StorageProjectImageCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StorageProjectImageCreate) SetName(v string) {
	o.Name = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *StorageProjectImageCreate) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProjectImageCreate) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *StorageProjectImageCreate) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *StorageProjectImageCreate) SetService(v string) {
	o.Service = &v
}

// GetVm returns the Vm field value if set, zero value otherwise.
func (o *StorageProjectImageCreate) GetVm() string {
	if o == nil || o.Vm == nil {
		var ret string
		return ret
	}
	return *o.Vm
}

// GetVmOk returns a tuple with the Vm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProjectImageCreate) GetVmOk() (*string, bool) {
	if o == nil || o.Vm == nil {
		return nil, false
	}
	return o.Vm, true
}

// HasVm returns a boolean if a field has been set.
func (o *StorageProjectImageCreate) HasVm() bool {
	if o != nil && o.Vm != nil {
		return true
	}

	return false
}

// SetVm gets a reference to the given string and assigns it to the Vm field.
func (o *StorageProjectImageCreate) SetVm(v string) {
	o.Vm = &v
}

// GetReplica returns the Replica field value if set, zero value otherwise.
func (o *StorageProjectImageCreate) GetReplica() string {
	if o == nil || o.Replica == nil {
		var ret string
		return ret
	}
	return *o.Replica
}

// GetReplicaOk returns a tuple with the Replica field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProjectImageCreate) GetReplicaOk() (*string, bool) {
	if o == nil || o.Replica == nil {
		return nil, false
	}
	return o.Replica, true
}

// HasReplica returns a boolean if a field has been set.
func (o *StorageProjectImageCreate) HasReplica() bool {
	if o != nil && o.Replica != nil {
		return true
	}

	return false
}

// SetReplica gets a reference to the given string and assigns it to the Replica field.
func (o *StorageProjectImageCreate) SetReplica(v string) {
	o.Replica = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StorageProjectImageCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProjectImageCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StorageProjectImageCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StorageProjectImageCreate) SetDescription(v string) {
	o.Description = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *StorageProjectImageCreate) GetTag() []Tag {
	if o == nil || o.Tag == nil {
		var ret []Tag
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProjectImageCreate) GetTagOk() (*[]Tag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *StorageProjectImageCreate) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Tag and assigns it to the Tag field.
func (o *StorageProjectImageCreate) SetTag(v []Tag) {
	o.Tag = &v
}

func (o StorageProjectImageCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Vm != nil {
		toSerialize["vm"] = o.Vm
	}
	if o.Replica != nil {
		toSerialize["replica"] = o.Replica
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableStorageProjectImageCreate struct {
	value *StorageProjectImageCreate
	isSet bool
}

func (v NullableStorageProjectImageCreate) Get() *StorageProjectImageCreate {
	return v.value
}

func (v *NullableStorageProjectImageCreate) Set(val *StorageProjectImageCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageProjectImageCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageProjectImageCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageProjectImageCreate(val *StorageProjectImageCreate) *NullableStorageProjectImageCreate {
	return &NullableStorageProjectImageCreate{value: val, isSet: true}
}

func (v NullableStorageProjectImageCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageProjectImageCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


