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

// StorageProjectDiskCreate struct for StorageProjectDiskCreate
type StorageProjectDiskCreate struct {
	Name string `json:"name"`
	Service string `json:"service"`
	Size float32 `json:"size"`
	Source *OneOfAnyTypeAnyType `json:"source,omitempty"`
	Vm *string `json:"vm,omitempty"`
	Tag []Tag `json:"tag,omitempty"`
}

// NewStorageProjectDiskCreate instantiates a new StorageProjectDiskCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageProjectDiskCreate(name string, service string, size float32) *StorageProjectDiskCreate {
	this := StorageProjectDiskCreate{}
	this.Name = name
	this.Service = service
	this.Size = size
	return &this
}

// NewStorageProjectDiskCreateWithDefaults instantiates a new StorageProjectDiskCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageProjectDiskCreateWithDefaults() *StorageProjectDiskCreate {
	this := StorageProjectDiskCreate{}
	return &this
}

// GetName returns the Name field value
func (o *StorageProjectDiskCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StorageProjectDiskCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StorageProjectDiskCreate) SetName(v string) {
	o.Name = v
}

// GetService returns the Service field value
func (o *StorageProjectDiskCreate) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *StorageProjectDiskCreate) GetServiceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *StorageProjectDiskCreate) SetService(v string) {
	o.Service = v
}

// GetSize returns the Size field value
func (o *StorageProjectDiskCreate) GetSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *StorageProjectDiskCreate) GetSizeOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *StorageProjectDiskCreate) SetSize(v float32) {
	o.Size = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *StorageProjectDiskCreate) GetSource() OneOfAnyTypeAnyType {
	if o == nil || o.Source == nil {
		var ret OneOfAnyTypeAnyType
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProjectDiskCreate) GetSourceOk() (*OneOfAnyTypeAnyType, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *StorageProjectDiskCreate) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given OneOfAnyTypeAnyType and assigns it to the Source field.
func (o *StorageProjectDiskCreate) SetSource(v OneOfAnyTypeAnyType) {
	o.Source = &v
}

// GetVm returns the Vm field value if set, zero value otherwise.
func (o *StorageProjectDiskCreate) GetVm() string {
	if o == nil || o.Vm == nil {
		var ret string
		return ret
	}
	return *o.Vm
}

// GetVmOk returns a tuple with the Vm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProjectDiskCreate) GetVmOk() (*string, bool) {
	if o == nil || o.Vm == nil {
		return nil, false
	}
	return o.Vm, true
}

// HasVm returns a boolean if a field has been set.
func (o *StorageProjectDiskCreate) HasVm() bool {
	if o != nil && o.Vm != nil {
		return true
	}

	return false
}

// SetVm gets a reference to the given string and assigns it to the Vm field.
func (o *StorageProjectDiskCreate) SetVm(v string) {
	o.Vm = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *StorageProjectDiskCreate) GetTag() []Tag {
	if o == nil || o.Tag == nil {
		var ret []Tag
		return ret
	}
	return o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProjectDiskCreate) GetTagOk() ([]Tag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *StorageProjectDiskCreate) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Tag and assigns it to the Tag field.
func (o *StorageProjectDiskCreate) SetTag(v []Tag) {
	o.Tag = v
}

func (o StorageProjectDiskCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["service"] = o.Service
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Vm != nil {
		toSerialize["vm"] = o.Vm
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableStorageProjectDiskCreate struct {
	value *StorageProjectDiskCreate
	isSet bool
}

func (v NullableStorageProjectDiskCreate) Get() *StorageProjectDiskCreate {
	return v.value
}

func (v *NullableStorageProjectDiskCreate) Set(val *StorageProjectDiskCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageProjectDiskCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageProjectDiskCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageProjectDiskCreate(val *StorageProjectDiskCreate) *NullableStorageProjectDiskCreate {
	return &NullableStorageProjectDiskCreate{value: val, isSet: true}
}

func (v NullableStorageProjectDiskCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageProjectDiskCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


