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

// StorageProjectVaultCreate struct for StorageProjectVaultCreate
type StorageProjectVaultCreate struct {
	Name string `json:"name"`
	Service *string `json:"service,omitempty"`
	Size float32 `json:"size"`
	Source *string `json:"source,omitempty"`
	Tag *[]Tag `json:"tag,omitempty"`
}

// NewStorageProjectVaultCreate instantiates a new StorageProjectVaultCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageProjectVaultCreate(name string, size float32) *StorageProjectVaultCreate {
	this := StorageProjectVaultCreate{}
	this.Name = name
	var service string = "5a0332c4eb8f4ed95c206a12"
	this.Service = &service
	this.Size = size
	return &this
}

// NewStorageProjectVaultCreateWithDefaults instantiates a new StorageProjectVaultCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageProjectVaultCreateWithDefaults() *StorageProjectVaultCreate {
	this := StorageProjectVaultCreate{}
	var service string = "5a0332c4eb8f4ed95c206a12"
	this.Service = &service
	return &this
}

// GetName returns the Name field value
func (o *StorageProjectVaultCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StorageProjectVaultCreate) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StorageProjectVaultCreate) SetName(v string) {
	o.Name = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *StorageProjectVaultCreate) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProjectVaultCreate) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *StorageProjectVaultCreate) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *StorageProjectVaultCreate) SetService(v string) {
	o.Service = &v
}

// GetSize returns the Size field value
func (o *StorageProjectVaultCreate) GetSize() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *StorageProjectVaultCreate) GetSizeOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *StorageProjectVaultCreate) SetSize(v float32) {
	o.Size = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *StorageProjectVaultCreate) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProjectVaultCreate) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *StorageProjectVaultCreate) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *StorageProjectVaultCreate) SetSource(v string) {
	o.Source = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *StorageProjectVaultCreate) GetTag() []Tag {
	if o == nil || o.Tag == nil {
		var ret []Tag
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageProjectVaultCreate) GetTagOk() (*[]Tag, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *StorageProjectVaultCreate) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given []Tag and assigns it to the Tag field.
func (o *StorageProjectVaultCreate) SetTag(v []Tag) {
	o.Tag = &v
}

func (o StorageProjectVaultCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	return json.Marshal(toSerialize)
}

type NullableStorageProjectVaultCreate struct {
	value *StorageProjectVaultCreate
	isSet bool
}

func (v NullableStorageProjectVaultCreate) Get() *StorageProjectVaultCreate {
	return v.value
}

func (v *NullableStorageProjectVaultCreate) Set(val *StorageProjectVaultCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageProjectVaultCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageProjectVaultCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageProjectVaultCreate(val *StorageProjectVaultCreate) *NullableStorageProjectVaultCreate {
	return &NullableStorageProjectVaultCreate{value: val, isSet: true}
}

func (v NullableStorageProjectVaultCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageProjectVaultCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


