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

// ComputeProjectVmUpdate struct for ComputeProjectVmUpdate
type ComputeProjectVmUpdate struct {
	UserMetadata *string `json:"userMetadata,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewComputeProjectVmUpdate instantiates a new ComputeProjectVmUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeProjectVmUpdate() *ComputeProjectVmUpdate {
	this := ComputeProjectVmUpdate{}
	return &this
}

// NewComputeProjectVmUpdateWithDefaults instantiates a new ComputeProjectVmUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeProjectVmUpdateWithDefaults() *ComputeProjectVmUpdate {
	this := ComputeProjectVmUpdate{}
	return &this
}

// GetUserMetadata returns the UserMetadata field value if set, zero value otherwise.
func (o *ComputeProjectVmUpdate) GetUserMetadata() string {
	if o == nil || o.UserMetadata == nil {
		var ret string
		return ret
	}
	return *o.UserMetadata
}

// GetUserMetadataOk returns a tuple with the UserMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeProjectVmUpdate) GetUserMetadataOk() (*string, bool) {
	if o == nil || o.UserMetadata == nil {
		return nil, false
	}
	return o.UserMetadata, true
}

// HasUserMetadata returns a boolean if a field has been set.
func (o *ComputeProjectVmUpdate) HasUserMetadata() bool {
	if o != nil && o.UserMetadata != nil {
		return true
	}

	return false
}

// SetUserMetadata gets a reference to the given string and assigns it to the UserMetadata field.
func (o *ComputeProjectVmUpdate) SetUserMetadata(v string) {
	o.UserMetadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ComputeProjectVmUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeProjectVmUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ComputeProjectVmUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ComputeProjectVmUpdate) SetName(v string) {
	o.Name = &v
}

func (o ComputeProjectVmUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserMetadata != nil {
		toSerialize["userMetadata"] = o.UserMetadata
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableComputeProjectVmUpdate struct {
	value *ComputeProjectVmUpdate
	isSet bool
}

func (v NullableComputeProjectVmUpdate) Get() *ComputeProjectVmUpdate {
	return v.value
}

func (v *NullableComputeProjectVmUpdate) Set(val *ComputeProjectVmUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableComputeProjectVmUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableComputeProjectVmUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputeProjectVmUpdate(val *ComputeProjectVmUpdate) *NullableComputeProjectVmUpdate {
	return &NullableComputeProjectVmUpdate{value: val, isSet: true}
}

func (v NullableComputeProjectVmUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputeProjectVmUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


