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

// ContainerRepository struct for ContainerRepository
type ContainerRepository struct {
	Id *string `json:"id,omitempty"`
	Name string `json:"name"`
}

// NewContainerRepository instantiates a new ContainerRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerRepository(name string) *ContainerRepository {
	this := ContainerRepository{}
	this.Name = name
	return &this
}

// NewContainerRepositoryWithDefaults instantiates a new ContainerRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerRepositoryWithDefaults() *ContainerRepository {
	this := ContainerRepository{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContainerRepository) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerRepository) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContainerRepository) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ContainerRepository) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *ContainerRepository) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ContainerRepository) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ContainerRepository) SetName(v string) {
	o.Name = v
}

func (o ContainerRepository) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableContainerRepository struct {
	value *ContainerRepository
	isSet bool
}

func (v NullableContainerRepository) Get() *ContainerRepository {
	return v.value
}

func (v *NullableContainerRepository) Set(val *ContainerRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableContainerRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableContainerRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainerRepository(val *ContainerRepository) *NullableContainerRepository {
	return &NullableContainerRepository{value: val, isSet: true}
}

func (v NullableContainerRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainerRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


